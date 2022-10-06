package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#name User#name}
	Name *UserName `field:"required" json:"name" yaml:"name"`
	// The user's primary email address. The primaryEmail must be unique and cannot be an alias of another user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#primary_email User#primary_email}
	PrimaryEmail *string `field:"required" json:"primaryEmail" yaml:"primaryEmail"`
	// addresses block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#addresses User#addresses}
	Addresses interface{} `field:"optional" json:"addresses" yaml:"addresses"`
	// asps.list of the user's alias email addresses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#aliases User#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Indicates if user is archived.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#archived User#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// Indicates if the user is forced to change their password at next login.
	//
	// This setting doesn't apply when the user signs in via a third-party identity provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#change_password_at_next_login User#change_password_at_next_login}
	ChangePasswordAtNextLogin interface{} `field:"optional" json:"changePasswordAtNextLogin" yaml:"changePasswordAtNextLogin"`
	// custom_schemas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#custom_schemas User#custom_schemas}
	CustomSchemas interface{} `field:"optional" json:"customSchemas" yaml:"customSchemas"`
	// emails block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#emails User#emails}
	Emails interface{} `field:"optional" json:"emails" yaml:"emails"`
	// external_ids block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#external_ids User#external_ids}
	ExternalIds interface{} `field:"optional" json:"externalIds" yaml:"externalIds"`
	// Stores the hash format of the password property.
	//
	// We recommend sending the password property value as a base 16 bit hexadecimal-encoded hash value. Set the hashFunction values as either the SHA-1, MD5, or crypt hash format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#hash_function User#hash_function}
	HashFunction *string `field:"optional" json:"hashFunction" yaml:"hashFunction"`
	// ims block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#ims User#ims}
	Ims interface{} `field:"optional" json:"ims" yaml:"ims"`
	// Defaults to `true`.
	//
	// Indicates if the user's profile is visible in the Google Workspace global address list when the contact sharing feature is enabled for the domain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#include_in_global_address_list User#include_in_global_address_list}
	IncludeInGlobalAddressList interface{} `field:"optional" json:"includeInGlobalAddressList" yaml:"includeInGlobalAddressList"`
	// If true, the user's IP address is added to the allow list.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#ip_allowlist User#ip_allowlist}
	IpAllowlist interface{} `field:"optional" json:"ipAllowlist" yaml:"ipAllowlist"`
	// Indicates a user with super admininistrator privileges.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#is_admin User#is_admin}
	IsAdmin interface{} `field:"optional" json:"isAdmin" yaml:"isAdmin"`
	// keywords block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#keywords User#keywords}
	Keywords interface{} `field:"optional" json:"keywords" yaml:"keywords"`
	// languages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#languages User#languages}
	Languages interface{} `field:"optional" json:"languages" yaml:"languages"`
	// locations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#locations User#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// organizations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#organizations User#organizations}
	Organizations interface{} `field:"optional" json:"organizations" yaml:"organizations"`
	// The full path of the parent organization associated with the user.
	//
	// If the parent organization is the top-level, it is represented as a forward slash (/).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#org_unit_path User#org_unit_path}
	OrgUnitPath *string `field:"optional" json:"orgUnitPath" yaml:"orgUnitPath"`
	// Stores the password for the user account.
	//
	// A password can contain any combination of ASCII characters. A minimum of 8 characters is required. The maximum length is 100 characters. As the API does not return the value of password, this field is write-only, and the value stored in the state will be what is provided in the configuration. The field is required on create and will be empty on import.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#password User#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// phones block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#phones User#phones}
	Phones interface{} `field:"optional" json:"phones" yaml:"phones"`
	// posix_accounts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#posix_accounts User#posix_accounts}
	PosixAccounts interface{} `field:"optional" json:"posixAccounts" yaml:"posixAccounts"`
	// Recovery email of the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#recovery_email User#recovery_email}
	RecoveryEmail *string `field:"optional" json:"recoveryEmail" yaml:"recoveryEmail"`
	// Recovery phone of the user.
	//
	// The phone number must be in the E.164 format, starting with the plus sign (+). Example: +16506661212.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#recovery_phone User#recovery_phone}
	RecoveryPhone *string `field:"optional" json:"recoveryPhone" yaml:"recoveryPhone"`
	// relations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#relations User#relations}
	Relations interface{} `field:"optional" json:"relations" yaml:"relations"`
	// ssh_public_keys block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#ssh_public_keys User#ssh_public_keys}
	SshPublicKeys interface{} `field:"optional" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// Indicates if user is suspended.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#suspended User#suspended}
	Suspended interface{} `field:"optional" json:"suspended" yaml:"suspended"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#timeouts User#timeouts}
	Timeouts *UserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// websites block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/googleworkspace/r/user#websites User#websites}
	Websites interface{} `field:"optional" json:"websites" yaml:"websites"`
}

