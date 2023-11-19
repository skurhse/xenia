package res

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
)

func NewVirtualNetwork(stk cdktf.TerraformStack, naming naming.Naming, rg resourcegroup.ResourceGroup, addrSpace []*string, token string) vnet.VirtualNetwork {

	id := fmt.Sprintf("%s_%s", Ids.VirtualNetwork, token)

	input := vnet.VirtualNetworkConfig{
		Name:              naming.VirtualNetworkOutput(),
		AddressSpace:      &addrSpace,
		Location:          rg.Location(),
		ResourceGroupName: rg.Name(),
	}

	return vnet.NewVirtualNetwork(stk, &id, &input)
}
