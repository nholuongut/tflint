package main

import (
	"github.com/nholuongut/tflint-plugin-sdk/plugin"
	"github.com/nholuongut/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:       "incompatiblehost",
			Version:    "0.1.0",
			Constraint: ">= 1.0",
		},
	})
}
