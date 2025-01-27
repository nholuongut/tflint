package formatter

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/fatih/color"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/nholuongut/tflint/tflint"
)

func Test_prettyPrint(t *testing.T) {
	// Disable color
	color.NoColor = true

	warningColor := "\x1b[33m"
	highlightColor := "\x1b[1;4m"
	resetColor := "\x1b[0m"

	cases := []struct {
		Name    string
		Issues  tflint.Issues
		Fix     bool
		Error   error
		Sources map[string][]byte
		Stdout  string
		Stderr  string
	}{
		{
			Name:   "no issues",
			Issues: tflint.Issues{},
			Stdout: "",
		},
		{
			Name: "issues",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
					Callers: []hcl.Range{
						{
							Filename: "test.tf",
							Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
							End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
						},
						{
							Filename: "module.tf",
							Start:    hcl.Pos{Line: 2, Column: 3, Byte: 0},
							End:      hcl.Pos{Line: 2, Column: 6, Byte: 3},
						},
					},
				},
			},
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stdout: `1 issue(s) found:

Error: test (test_rule)

  on test.tf line 1:
   1: foo = 1

Callers:
   test.tf:1,1-4
   module.tf:2,3-6

Reference: https://github.com

`,
		},
		{
			Name: "no sources",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
				},
			},
			Stdout: `1 issue(s) found:

Error: test (test_rule)

  on test.tf line 1:
   (source code not available)

Reference: https://github.com

`,
		},
		{
			Name: "fixable",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Fixable: true,
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
				},
			},
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stdout: `1 issue(s) found:

Error: [Fixable] test (test_rule)

  on test.tf line 1:
   1: foo = 1

Reference: https://github.com

`,
		},
		{
			Name: "fixed",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Fixable: true,
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
				},
			},
			Fix: true,
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stdout: `1 issue(s) found:

Error: [Fixed] test (test_rule)

  on test.tf line 1:
   1: foo = 1

Reference: https://github.com

`,
		},
		{
			Name: "issue with source",
			Issues: tflint.Issues{
				{
					Rule:    &testRule{},
					Message: "test",
					Range: hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
					Source: []byte("bar = 1"),
				},
			},
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stdout: `1 issue(s) found:

Error: test (test_rule)

  on test.tf line 1:
   1: bar = 1

Reference: https://github.com

`,
		},
		{
			Name:   "error",
			Issues: tflint.Issues{},
			Error:  fmt.Errorf("Failed to work; %w", errors.New("I don't feel like working")),
			Stderr: "Failed to work; I don't feel like working\n",
		},
		{
			Name:   "diagnostics",
			Issues: tflint.Issues{},
			Error: hcl.Diagnostics{
				&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "summary",
					Detail:   "detail",
					Subject: &hcl.Range{
						Filename: "test.tf",
						Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
						End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
					},
				},
			},
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stderr: fmt.Sprintf(`test.tf:1,1-4: summary; detail:

%sWarning%s: summary

  on test.tf line 1:
   1: %sfoo%s = 1

detail

`, warningColor, resetColor, highlightColor, resetColor),
		},
		{
			Name:   "joined errors",
			Issues: tflint.Issues{},
			Error: errors.Join(
				errors.New("an error occurred"),
				errors.New("failed"),
				hcl.Diagnostics{
					&hcl.Diagnostic{
						Severity: hcl.DiagWarning,
						Summary:  "summary",
						Detail:   "detail",
						Subject: &hcl.Range{
							Filename: "test.tf",
							Start:    hcl.Pos{Line: 1, Column: 1, Byte: 0},
							End:      hcl.Pos{Line: 1, Column: 4, Byte: 3},
						},
					},
				},
			),
			Sources: map[string][]byte{
				"test.tf": []byte("foo = 1"),
			},
			Stderr: fmt.Sprintf(`an error occurred
failed
test.tf:1,1-4: summary; detail:

%sWarning%s: summary

  on test.tf line 1:
   1: %sfoo%s = 1

detail

`, warningColor, resetColor, highlightColor, resetColor),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			formatter := &Formatter{Stdout: stdout, Stderr: stderr, Fix: tc.Fix}

			formatter.prettyPrint(tc.Issues, tc.Error, tc.Sources)

			if stdout.String() != tc.Stdout {
				t.Fatalf("expected=%s, stdout=%s", tc.Stdout, stdout.String())
			}
			if stderr.String() != tc.Stderr {
				t.Fatalf("expected=%s, stderr=%s", tc.Stderr, stderr.String())
			}
		})
	}
}
