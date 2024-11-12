package cmd

import (
	"fmt"

	"github.com/nholuongut/tflint-plugin-sdk/plugin"
	"github.com/nholuongut/tflint-plugin-sdk/tflint"
	"github.com/nholuongut/tflint-ruleset-terraform/project"
	"github.com/nholuongut/tflint-ruleset-terraform/rules"
	"github.com/nholuongut/tflint-ruleset-terraform/terraform"
)

func (cli *CLI) actAsBundledPlugin() int {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &terraform.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "terraform",
				Version: fmt.Sprintf("%s-bundled", project.Version),
			},
			PresetRules: rules.PresetRules,
		},
	})
	return ExitCodeOK
}
