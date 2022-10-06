package user


type UserAddresses struct {
	// The address type. Acceptable values: `custom`, `home`, `other`, `work`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#type User#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Country.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#country User#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// The country code. Uses the ISO 3166-1 standard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#country_code User#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// If the address type is custom, this property contains the custom value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#custom_type User#custom_type}
	CustomType *string `field:"optional" json:"customType" yaml:"customType"`
	// For extended addresses, such as an address that includes a sub-region.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#extended_address User#extended_address}
	ExtendedAddress *string `field:"optional" json:"extendedAddress" yaml:"extendedAddress"`
	// A full and unstructured postal address. This is not synced with the structured address fields.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#formatted User#formatted}
	Formatted *string `field:"optional" json:"formatted" yaml:"formatted"`
	// The town or city of the address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#locality User#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The post office box, if present.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#po_box User#po_box}
	PoBox *string `field:"optional" json:"poBox" yaml:"poBox"`
	// The ZIP or postal code, if applicable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#postal_code User#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// If this is the user's primary address. The addresses list may contain only one primary address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The abbreviated province or state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#region User#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Indicates if the user-supplied address was formatted. Formatted addresses are not currently supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#source_is_structured User#source_is_structured}
	SourceIsStructured interface{} `field:"optional" json:"sourceIsStructured" yaml:"sourceIsStructured"`
	// The street address, such as 1600 Amphitheatre Parkway. Whitespace within the string is ignored; however, newlines are significant.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#street_address User#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
}

