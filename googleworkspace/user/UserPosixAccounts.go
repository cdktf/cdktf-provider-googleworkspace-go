package user


type UserPosixAccounts struct {
	// A POSIX account field identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#account_id User#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The GECOS (user information) for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#gecos User#gecos}
	Gecos *string `field:"optional" json:"gecos" yaml:"gecos"`
	// The default group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#gid User#gid}
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// The path to the home directory for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#home_directory User#home_directory}
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// The operating system type for this account. Acceptable values: `linux`, `unspecified`, `windows`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#operating_system_type User#operating_system_type}
	OperatingSystemType *string `field:"optional" json:"operatingSystemType" yaml:"operatingSystemType"`
	// If this is user's primary account within the SystemId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#primary User#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The path to the login shell for this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#shell User#shell}
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// System identifier for which account Username or Uid apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#system_id User#system_id}
	SystemId *string `field:"optional" json:"systemId" yaml:"systemId"`
	// The POSIX compliant user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#uid User#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// The username of the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user#username User#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

