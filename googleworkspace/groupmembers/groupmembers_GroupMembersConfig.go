package groupmembers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupMembersConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifies the group in the API request.
	//
	// The value can be the group's email address, group alias, or the unique group ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_members#group_id GroupMembers#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// members block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_members#members GroupMembers#members}
	Members interface{} `field:"optional" json:"members" yaml:"members"`
}

