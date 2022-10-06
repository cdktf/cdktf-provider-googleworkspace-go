package user


type UserPhones struct {
	// The type of phone number.
	//
	// Acceptable values: `assistant`, `callback`, `car`, `company_main` , `custom`, `grand_central`, `home`, `home_fax`, `isdn`, `main`, `mobile`, `other`, `other_fax`, `pager`, `radio`, `telex`, `tty_tdd`, `work`, `work_fax`, `work_mobile`, `work_pager`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A human-readable phone number. It may be in any telephone number format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#value User#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// If the phone number type is custom, this property contains the custom value and must be set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// Indicates if this is the user's primary phone number. A user may only have one primary phone number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

