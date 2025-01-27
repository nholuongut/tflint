package custom

import (
	"github.com/nholuongut/tflint-plugin-sdk/hclext"
	"github.com/nholuongut/tflint-plugin-sdk/tflint"
)

type Runner struct {
	tflint.Runner
	CustomConfig *Config
}

func NewRunner(runner tflint.Runner, config *Config) (*Runner, error) {
	providers, err := runner.GetModuleContent(
		&hclext.BodySchema{
			Blocks: []hclext.BlockSchema{
				{
					Type:       "provider",
					LabelNames: []string{"name"},
					Body: &hclext.BodySchema{
						Attributes: []hclext.AttributeSchema{
							{Name: "zone"},
						},
						Blocks: []hclext.BlockSchema{
							{
								Type: "annotation",
								Body: &hclext.BodySchema{
									Attributes: []hclext.AttributeSchema{
										{Name: "value"},
									},
								},
							},
						},
					},
				},
			},
		},
		&tflint.GetModuleContentOption{ModuleCtx: tflint.RootModuleCtxType},
	)
	if err != nil {
		return nil, err
	}

	for _, provider := range providers.Blocks {
		if provider.Labels[0] != "custom" {
			continue
		}

		opts := &tflint.EvaluateExprOption{ModuleCtx: tflint.RootModuleCtxType}

		if attr, exists := provider.Body.Attributes["zone"]; exists {
			err := runner.EvaluateExpr(attr.Expr, func(zone string) error {
				config.Zone = zone
				return nil
			}, opts)
			if err != nil {
				return nil, err
			}
		}

		for _, annotation := range provider.Body.Blocks {
			if attr, exists := annotation.Body.Attributes["value"]; exists {
				err := runner.EvaluateExpr(attr.Expr, func(val string) error {
					config.Annotation = val
					return nil
				}, opts)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return &Runner{
		Runner:       runner,
		CustomConfig: config,
	}, nil
}
