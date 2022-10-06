package groupsettings


type GroupSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_settings#create GroupSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/group_settings#update GroupSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

