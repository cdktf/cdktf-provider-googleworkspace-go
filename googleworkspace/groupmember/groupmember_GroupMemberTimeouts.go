package groupmember


type GroupMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_member#create GroupMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_member#update GroupMember#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

