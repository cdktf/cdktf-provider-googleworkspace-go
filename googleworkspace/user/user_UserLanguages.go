package user


type UserLanguages struct {
	// Other language.
	//
	// A user can provide their own language name if there is no corresponding Google III language code. If this is set, LanguageCode can't be set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#custom_language User#custom_language}
	CustomLanguage *string `field:"optional" json:"customLanguage" yaml:"customLanguage"`
	// Defaults to `en`.
	//
	// Language Code. Should be used for storing Google III LanguageCode string representation for language. Illegal values cause SchemaException.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#language_code User#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Defaults to `preferred`.
	//
	// If present, controls whether the specified languageCode is the user's preferred language. Allowed values are `preferred` and `not_preferred`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#preference User#preference}
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
}

