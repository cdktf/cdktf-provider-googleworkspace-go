package group


type GroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group#create Group#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group#update Group#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

