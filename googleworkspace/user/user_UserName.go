package user


type UserName struct {
	// The user's last name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#family_name User#family_name}
	FamilyName *string `field:"required" json:"familyName" yaml:"familyName"`
	// The user's first name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#given_name User#given_name}
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
}

