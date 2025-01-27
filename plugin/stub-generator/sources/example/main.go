package main

import (
	"fmt"

	"github.com/nholuongut/tflint-plugin-sdk/hclext"
	"github.com/nholuongut/tflint-plugin-sdk/plugin"
	"github.com/nholuongut/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "example",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				NewAwsInstanceExampleTypeRule(),
			},
		},
	})
}

// AwsInstanceExampleTypeRule checks whether ...
type AwsInstanceExampleTypeRule struct {
	tflint.DefaultRule
}

// NewAwsInstanceExampleTypeRule returns a new rule
func NewAwsInstanceExampleTypeRule() *AwsInstanceExampleTypeRule {
	return &AwsInstanceExampleTypeRule{}
}

// Name returns the rule name
func (r *AwsInstanceExampleTypeRule) Name() string {
	return "aws_instance_example_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceExampleTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceExampleTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceExampleTypeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *AwsInstanceExampleTypeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("aws_instance", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: "instance_type"}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes["instance_type"]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(instanceType string) error {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("instance type is %s", instanceType),
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
