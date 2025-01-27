package tflint

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

// Version is application version
var Version *version.Version = version.Must(version.NewVersion("0.54.0"))

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/nholuongut/tflint/blob/v%s/docs/rules/%s.md", Version, name)
}
