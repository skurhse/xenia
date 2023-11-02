package res

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	aks "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/kubernetescluster"
	rg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/resourcegroup"
	vnet "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/virtualnetwork"
	"github.com/transprogrammer/xenia/generated/naming"
	"github.com/transprogrammer/xenia/pkg/cfg"
)

// TODO <rbt 2023-10-28>
// - Pipe jumpbox ip
// - Pipe subnet id

func NewCluster(scope constructs.Construct, cfg cfg.Config, naming naming.Naming, rg rg.ResourceGroup, subnet vnet.VirtualNetworkSubnetOutputReference) aks.KubernetesCluster {

	accessProfile := KubernetesClusterApiServerAccessProfile{
		AuthorizedIpRange:      jump.NIC().IP(),
		SubnetId:               jump.Subnet().Id(),
		VnetIntegrationEnabled: jsii.Bool(true),
	}

	roleBasedAccessControl := AzureActiveDirectoryRoleBasedAccessControl{
		Managed:             jsii.Bool(true),
		AzureRbacEnabled:    jsii.Bool(true),
		AdminGroupObjectIds: *[]*string{group.ObjectId()},
	}

	input := KubernetesClusterConfig{
		Name:                    naming.KubernetesClusterOutput(),
		Location:                cfg.Regions().Primary(),
		ResourceGroupName:       rg.Name(),
		DnsPrefixPrivateCluster: cfg.Name(),
		AutomaticChannelUpgrade: jsii.String("node-image"),
		APIServerAccessProfile:  accessProfile,
		AzureActiveDirectoryRoleBasedAccessControl: roleBasedAccessControl,
	}

	id := Ids.KubernetesCluster
}
