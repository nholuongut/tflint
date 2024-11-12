package main

import (
	"github.com/nholuongut/tflint-plugin-sdk/plugin"
	"github.com/nholuongut/tflint-plugin-sdk/tflint"
	"github.com/nholuongut/tflint/plugin/stub-generator/sources/customrulesettesting/custom"
	"github.com/nholuongut/tflint/plugin/stub-generator/sources/customrulesettesting/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &custom.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "customrulesettesting",
				Version: "0.1.0",
				Rules: []tflint.Rule{
					rules.NewAwsInstanceExampleTypeRule(),
				},
			},
		},
	})
}
