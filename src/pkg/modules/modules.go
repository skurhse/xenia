package modules

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/apps"
)

func NewNaming(stack cdktf.TerraformStack, config apps.Config, suffix *[]*string) *naming.Naming {

	id := ModuleIds.Naming

	prefix := &[]*string{apps.AppName}

	input := naming.NamingConfig{
		Prefix:               prefix,
		UniqueIncludeNumbers: jsii.Bool(false),
		Suffix:               suffix,
	}

	naming := naming.NewNaming(stack, id, &input)

	return &naming
}
