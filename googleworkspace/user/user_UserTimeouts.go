package user


type UserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#create User#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#update User#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

