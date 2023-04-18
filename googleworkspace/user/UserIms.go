package user


type UserIms struct {
	// An IM protocol identifies the IM network.
	//
	// The value can be a custom network or the standard network. Acceptable values: `aim`, `custom_protocol`, `gtalk`, `icq`, `jabber`, `msn`, `net_meeting`, `qq`, `skype`, `yahoo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#protocol User#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Acceptable values: `custom`, `home`, `other`, `work`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If the protocol value is custom_protocol, this property holds the custom protocol's string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_protocol User#custom_protocol}
	CustomProtocol *string `field:"optional" json:"customProtocol" yaml:"customProtocol"`
	// If the IM type is custom, this property holds the custom type string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// The user's IM network ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#im User#im}
	Im *string `field:"optional" json:"im" yaml:"im"`
	// If this is the user's primary IM.
	//
	// Only one entry in the IM list can have a value of true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

