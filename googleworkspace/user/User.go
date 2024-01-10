// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user googleworkspace_user}.
type User interface {
	cdktf.TerraformResource
	Addresses() UserAddressesList
	AddressesInput() interface{}
	AgreedToTerms() cdktf.IResolvable
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	Archived() interface{}
	SetArchived(val interface{})
	ArchivedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChangePasswordAtNextLogin() interface{}
	SetChangePasswordAtNextLogin(val interface{})
	ChangePasswordAtNextLoginInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	CustomerId() *string
	CustomSchemas() UserCustomSchemasList
	CustomSchemasInput() interface{}
	DeletionTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Emails() UserEmailsList
	EmailsInput() interface{}
	Etag() *string
	ExternalIds() UserExternalIdsList
	ExternalIdsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HashFunction() *string
	SetHashFunction(val *string)
	HashFunctionInput() *string
	Id() *string
	Ims() UserImsList
	ImsInput() interface{}
	IncludeInGlobalAddressList() interface{}
	SetIncludeInGlobalAddressList(val interface{})
	IncludeInGlobalAddressListInput() interface{}
	IpAllowlist() interface{}
	SetIpAllowlist(val interface{})
	IpAllowlistInput() interface{}
	IsAdmin() interface{}
	SetIsAdmin(val interface{})
	IsAdminInput() interface{}
	IsDelegatedAdmin() cdktf.IResolvable
	IsEnforcedIn2StepVerification() cdktf.IResolvable
	IsEnrolledIn2StepVerification() cdktf.IResolvable
	IsMailboxSetup() cdktf.IResolvable
	Keywords() UserKeywordsList
	KeywordsInput() interface{}
	Languages() UserLanguagesList
	LanguagesInput() interface{}
	LastLoginTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locations() UserLocationsList
	LocationsInput() interface{}
	Name() UserNameOutputReference
	NameInput() *UserName
	// The tree node.
	Node() constructs.Node
	NonEditableAliases() *[]*string
	Organizations() UserOrganizationsList
	OrganizationsInput() interface{}
	OrgUnitPath() *string
	SetOrgUnitPath(val *string)
	OrgUnitPathInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Phones() UserPhonesList
	PhonesInput() interface{}
	PosixAccounts() UserPosixAccountsList
	PosixAccountsInput() interface{}
	PrimaryEmail() *string
	SetPrimaryEmail(val *string)
	PrimaryEmailInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecoveryEmail() *string
	SetRecoveryEmail(val *string)
	RecoveryEmailInput() *string
	RecoveryPhone() *string
	SetRecoveryPhone(val *string)
	RecoveryPhoneInput() *string
	Relations() UserRelationsList
	RelationsInput() interface{}
	SshPublicKeys() UserSshPublicKeysList
	SshPublicKeysInput() interface{}
	Suspended() interface{}
	SetSuspended(val interface{})
	SuspendedInput() interface{}
	SuspensionReason() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThumbnailPhotoEtag() *string
	ThumbnailPhotoUrl() *string
	Timeouts() UserTimeoutsOutputReference
	TimeoutsInput() interface{}
	Websites() UserWebsitesList
	WebsitesInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAddresses(value interface{})
	PutCustomSchemas(value interface{})
	PutEmails(value interface{})
	PutExternalIds(value interface{})
	PutIms(value interface{})
	PutKeywords(value interface{})
	PutLanguages(value interface{})
	PutLocations(value interface{})
	PutName(value *UserName)
	PutOrganizations(value interface{})
	PutPhones(value interface{})
	PutPosixAccounts(value interface{})
	PutRelations(value interface{})
	PutSshPublicKeys(value interface{})
	PutTimeouts(value *UserTimeouts)
	PutWebsites(value interface{})
	ResetAddresses()
	ResetAliases()
	ResetArchived()
	ResetChangePasswordAtNextLogin()
	ResetCustomSchemas()
	ResetEmails()
	ResetExternalIds()
	ResetHashFunction()
	ResetIms()
	ResetIncludeInGlobalAddressList()
	ResetIpAllowlist()
	ResetIsAdmin()
	ResetKeywords()
	ResetLanguages()
	ResetLocations()
	ResetOrganizations()
	ResetOrgUnitPath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPhones()
	ResetPosixAccounts()
	ResetRecoveryEmail()
	ResetRecoveryPhone()
	ResetRelations()
	ResetSshPublicKeys()
	ResetSuspended()
	ResetTimeouts()
	ResetWebsites()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_User) Addresses() UserAddressesList {
	var returns UserAddressesList
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AgreedToTerms() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"agreedToTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Archived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ChangePasswordAtNextLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changePasswordAtNextLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ChangePasswordAtNextLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changePasswordAtNextLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomSchemas() UserCustomSchemasList {
	var returns UserCustomSchemasList
	_jsii_.Get(
		j,
		"customSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) CustomSchemasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customSchemasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DeletionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Emails() UserEmailsList {
	var returns UserEmailsList
	_jsii_.Get(
		j,
		"emails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) EmailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ExternalIds() UserExternalIdsList {
	var returns UserExternalIdsList
	_jsii_.Get(
		j,
		"externalIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ExternalIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HashFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) HashFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Ims() UserImsList {
	var returns UserImsList
	_jsii_.Get(
		j,
		"ims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ImsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IncludeInGlobalAddressList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IncludeInGlobalAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInGlobalAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IpAllowlist() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IpAllowlistInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsDelegatedAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isDelegatedAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsEnforcedIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnforcedIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsEnrolledIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnrolledIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) IsMailboxSetup() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isMailboxSetup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Keywords() UserKeywordsList {
	var returns UserKeywordsList
	_jsii_.Get(
		j,
		"keywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) KeywordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Languages() UserLanguagesList {
	var returns UserLanguagesList
	_jsii_.Get(
		j,
		"languages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LanguagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"languagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LastLoginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastLoginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Locations() UserLocationsList {
	var returns UserLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) LocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Name() UserNameOutputReference {
	var returns UserNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NameInput() *UserName {
	var returns *UserName
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) NonEditableAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonEditableAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Organizations() UserOrganizationsList {
	var returns UserOrganizationsList
	_jsii_.Get(
		j,
		"organizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OrganizationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OrgUnitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) OrgUnitPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnitPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Phones() UserPhonesList {
	var returns UserPhonesList
	_jsii_.Get(
		j,
		"phones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PhonesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"phonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PosixAccounts() UserPosixAccountsList {
	var returns UserPosixAccountsList
	_jsii_.Get(
		j,
		"posixAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PosixAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"posixAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrimaryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) PrimaryEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RecoveryPhoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPhoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Relations() UserRelationsList {
	var returns UserRelationsList
	_jsii_.Get(
		j,
		"relations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) RelationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SshPublicKeys() UserSshPublicKeysList {
	var returns UserSshPublicKeysList
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SshPublicKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Suspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) SuspensionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspensionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ThumbnailPhotoEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) ThumbnailPhotoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Timeouts() UserTimeoutsOutputReference {
	var returns UserTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Websites() UserWebsitesList {
	var returns UserWebsitesList
	_jsii_.Get(
		j,
		"websites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) WebsitesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websitesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user googleworkspace_user} Resource.
func NewUser(scope constructs.Construct, id *string, config *UserConfig) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.User",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/user googleworkspace_user} Resource.
func NewUser_Override(u User, scope constructs.Construct, id *string, config *UserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.User",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_User)SetAliases(val *[]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_User)SetArchived(val interface{}) {
	if err := j.validateSetArchivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archived",
		val,
	)
}

func (j *jsiiProxy_User)SetChangePasswordAtNextLogin(val interface{}) {
	if err := j.validateSetChangePasswordAtNextLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changePasswordAtNextLogin",
		val,
	)
}

func (j *jsiiProxy_User)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_User)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_User)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_User)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_User)SetHashFunction(val *string) {
	if err := j.validateSetHashFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashFunction",
		val,
	)
}

func (j *jsiiProxy_User)SetIncludeInGlobalAddressList(val interface{}) {
	if err := j.validateSetIncludeInGlobalAddressListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInGlobalAddressList",
		val,
	)
}

func (j *jsiiProxy_User)SetIpAllowlist(val interface{}) {
	if err := j.validateSetIpAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAllowlist",
		val,
	)
}

func (j *jsiiProxy_User)SetIsAdmin(val interface{}) {
	if err := j.validateSetIsAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAdmin",
		val,
	)
}

func (j *jsiiProxy_User)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_User)SetOrgUnitPath(val *string) {
	if err := j.validateSetOrgUnitPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgUnitPath",
		val,
	)
}

func (j *jsiiProxy_User)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_User)SetPrimaryEmail(val *string) {
	if err := j.validateSetPrimaryEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryEmail",
		val,
	)
}

func (j *jsiiProxy_User)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_User)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_User)SetRecoveryEmail(val *string) {
	if err := j.validateSetRecoveryEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryEmail",
		val,
	)
}

func (j *jsiiProxy_User)SetRecoveryPhone(val *string) {
	if err := j.validateSetRecoveryPhoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPhone",
		val,
	)
}

func (j *jsiiProxy_User)SetSuspended(val interface{}) {
	if err := j.validateSetSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspended",
		val,
	)
}

// Generates CDKTF code for importing a User resource upon running "cdktf plan <stack-name>".
func User_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.user.User",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.user.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.user.User",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func User_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.user.User",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func User_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googleworkspace.user.User",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_User) AddMoveTarget(moveTarget *string) {
	if err := u.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (u *jsiiProxy_User) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_User) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := u.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := u.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		u,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := u.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := u.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		u,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := u.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		u,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := u.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		u,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringAttribute(terraformAttribute *string) *string {
	if err := u.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := u.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		u,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := u.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (u *jsiiProxy_User) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) MoveFromId(id *string) {
	if err := u.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveFromId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_User) MoveTo(moveTarget *string, index interface{}) {
	if err := u.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (u *jsiiProxy_User) MoveToId(id *string) {
	if err := u.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"moveToId",
		[]interface{}{id},
	)
}

func (u *jsiiProxy_User) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_User) PutAddresses(value interface{}) {
	if err := u.validatePutAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putAddresses",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutCustomSchemas(value interface{}) {
	if err := u.validatePutCustomSchemasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putCustomSchemas",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutEmails(value interface{}) {
	if err := u.validatePutEmailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putEmails",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutExternalIds(value interface{}) {
	if err := u.validatePutExternalIdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putExternalIds",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutIms(value interface{}) {
	if err := u.validatePutImsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putIms",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutKeywords(value interface{}) {
	if err := u.validatePutKeywordsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putKeywords",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutLanguages(value interface{}) {
	if err := u.validatePutLanguagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putLanguages",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutLocations(value interface{}) {
	if err := u.validatePutLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putLocations",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutName(value *UserName) {
	if err := u.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putName",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutOrganizations(value interface{}) {
	if err := u.validatePutOrganizationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putOrganizations",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutPhones(value interface{}) {
	if err := u.validatePutPhonesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putPhones",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutPosixAccounts(value interface{}) {
	if err := u.validatePutPosixAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putPosixAccounts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutRelations(value interface{}) {
	if err := u.validatePutRelationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putRelations",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutSshPublicKeys(value interface{}) {
	if err := u.validatePutSshPublicKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putSshPublicKeys",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutTimeouts(value *UserTimeouts) {
	if err := u.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) PutWebsites(value interface{}) {
	if err := u.validatePutWebsitesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"putWebsites",
		[]interface{}{value},
	)
}

func (u *jsiiProxy_User) ResetAddresses() {
	_jsii_.InvokeVoid(
		u,
		"resetAddresses",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetAliases() {
	_jsii_.InvokeVoid(
		u,
		"resetAliases",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetArchived() {
	_jsii_.InvokeVoid(
		u,
		"resetArchived",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetChangePasswordAtNextLogin() {
	_jsii_.InvokeVoid(
		u,
		"resetChangePasswordAtNextLogin",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetCustomSchemas() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomSchemas",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetEmails() {
	_jsii_.InvokeVoid(
		u,
		"resetEmails",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetExternalIds() {
	_jsii_.InvokeVoid(
		u,
		"resetExternalIds",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetHashFunction() {
	_jsii_.InvokeVoid(
		u,
		"resetHashFunction",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetIms() {
	_jsii_.InvokeVoid(
		u,
		"resetIms",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetIncludeInGlobalAddressList() {
	_jsii_.InvokeVoid(
		u,
		"resetIncludeInGlobalAddressList",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetIpAllowlist() {
	_jsii_.InvokeVoid(
		u,
		"resetIpAllowlist",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetIsAdmin() {
	_jsii_.InvokeVoid(
		u,
		"resetIsAdmin",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetKeywords() {
	_jsii_.InvokeVoid(
		u,
		"resetKeywords",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLanguages() {
	_jsii_.InvokeVoid(
		u,
		"resetLanguages",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetLocations() {
	_jsii_.InvokeVoid(
		u,
		"resetLocations",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOrganizations() {
	_jsii_.InvokeVoid(
		u,
		"resetOrganizations",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOrgUnitPath() {
	_jsii_.InvokeVoid(
		u,
		"resetOrgUnitPath",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPhones() {
	_jsii_.InvokeVoid(
		u,
		"resetPhones",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetPosixAccounts() {
	_jsii_.InvokeVoid(
		u,
		"resetPosixAccounts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRecoveryEmail() {
	_jsii_.InvokeVoid(
		u,
		"resetRecoveryEmail",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRecoveryPhone() {
	_jsii_.InvokeVoid(
		u,
		"resetRecoveryPhone",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetRelations() {
	_jsii_.InvokeVoid(
		u,
		"resetRelations",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		u,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetSuspended() {
	_jsii_.InvokeVoid(
		u,
		"resetSuspended",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetTimeouts() {
	_jsii_.InvokeVoid(
		u,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) ResetWebsites() {
	_jsii_.InvokeVoid(
		u,
		"resetWebsites",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

