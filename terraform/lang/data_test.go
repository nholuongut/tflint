package lang

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/nholuongut/tflint/terraform/addrs"
	"github.com/zclconf/go-cty/cty"
)

type dataForTests struct {
	CountAttrs     map[string]cty.Value
	ForEachAttrs   map[string]cty.Value
	LocalValues    map[string]cty.Value
	PathAttrs      map[string]cty.Value
	TerraformAttrs map[string]cty.Value
	InputVariables map[string]cty.Value
}

var _ Data = &dataForTests{}

func (d *dataForTests) GetCountAttr(addr addrs.CountAttr, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.CountAttrs[addr.Name], nil
}

func (d *dataForTests) GetForEachAttr(addr addrs.ForEachAttr, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.ForEachAttrs[addr.Name], nil
}

func (d *dataForTests) GetInputVariable(addr addrs.InputVariable, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.InputVariables[addr.Name], nil
}

func (d *dataForTests) GetLocalValue(addr addrs.LocalValue, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.LocalValues[addr.Name], nil
}

func (d *dataForTests) GetPathAttr(addr addrs.PathAttr, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.PathAttrs[addr.Name], nil
}

func (d *dataForTests) GetTerraformAttr(addr addrs.TerraformAttr, rng hcl.Range) (cty.Value, hcl.Diagnostics) {
	return d.TerraformAttrs[addr.Name], nil
}
