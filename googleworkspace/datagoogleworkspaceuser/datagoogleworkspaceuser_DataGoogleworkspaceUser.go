package datagoogleworkspaceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v2/datagoogleworkspaceuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/googleworkspace/d/user googleworkspace_user}.
type DataGoogleworkspaceUser interface {
	cdktf.TerraformDataSource
	Addresses() DataGoogleworkspaceUserAddressesList
	AgreedToTerms() cdktf.IResolvable
	Aliases() *[]*string
	Archived() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChangePasswordAtNextLogin() cdktf.IResolvable
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreationTime() *string
	CustomerId() *string
	CustomSchemas() DataGoogleworkspaceUserCustomSchemasList
	DeletionTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Emails() DataGoogleworkspaceUserEmailsList
	Etag() *string
	ExternalIds() DataGoogleworkspaceUserExternalIdsList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HashFunction() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Ims() DataGoogleworkspaceUserImsList
	IncludeInGlobalAddressList() cdktf.IResolvable
	IpAllowlist() cdktf.IResolvable
	IsAdmin() cdktf.IResolvable
	IsDelegatedAdmin() cdktf.IResolvable
	IsEnforcedIn2StepVerification() cdktf.IResolvable
	IsEnrolledIn2StepVerification() cdktf.IResolvable
	IsMailboxSetup() cdktf.IResolvable
	Keywords() DataGoogleworkspaceUserKeywordsList
	Languages() DataGoogleworkspaceUserLanguagesList
	LastLoginTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locations() DataGoogleworkspaceUserLocationsList
	Name() DataGoogleworkspaceUserNameList
	// The tree node.
	Node() constructs.Node
	NonEditableAliases() *[]*string
	Organizations() DataGoogleworkspaceUserOrganizationsList
	OrgUnitPath() *string
	Password() *string
	Phones() DataGoogleworkspaceUserPhonesList
	PosixAccounts() DataGoogleworkspaceUserPosixAccountsList
	PrimaryEmail() *string
	SetPrimaryEmail(val *string)
	PrimaryEmailInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RecoveryEmail() *string
	RecoveryPhone() *string
	Relations() DataGoogleworkspaceUserRelationsList
	SshPublicKeys() DataGoogleworkspaceUserSshPublicKeysList
	Suspended() cdktf.IResolvable
	SuspensionReason() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThumbnailPhotoEtag() *string
	ThumbnailPhotoUrl() *string
	Websites() DataGoogleworkspaceUserWebsitesList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryEmail()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleworkspaceUser
type jsiiProxy_DataGoogleworkspaceUser struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Addresses() DataGoogleworkspaceUserAddressesList {
	var returns DataGoogleworkspaceUserAddressesList
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) AgreedToTerms() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"agreedToTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Archived() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ChangePasswordAtNextLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"changePasswordAtNextLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) CustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) CustomSchemas() DataGoogleworkspaceUserCustomSchemasList {
	var returns DataGoogleworkspaceUserCustomSchemasList
	_jsii_.Get(
		j,
		"customSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) DeletionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Emails() DataGoogleworkspaceUserEmailsList {
	var returns DataGoogleworkspaceUserEmailsList
	_jsii_.Get(
		j,
		"emails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ExternalIds() DataGoogleworkspaceUserExternalIdsList {
	var returns DataGoogleworkspaceUserExternalIdsList
	_jsii_.Get(
		j,
		"externalIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) HashFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Ims() DataGoogleworkspaceUserImsList {
	var returns DataGoogleworkspaceUserImsList
	_jsii_.Get(
		j,
		"ims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IncludeInGlobalAddressList() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeInGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IpAllowlist() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IsAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IsDelegatedAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isDelegatedAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IsEnforcedIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnforcedIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IsEnrolledIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnrolledIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) IsMailboxSetup() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isMailboxSetup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Keywords() DataGoogleworkspaceUserKeywordsList {
	var returns DataGoogleworkspaceUserKeywordsList
	_jsii_.Get(
		j,
		"keywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Languages() DataGoogleworkspaceUserLanguagesList {
	var returns DataGoogleworkspaceUserLanguagesList
	_jsii_.Get(
		j,
		"languages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) LastLoginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastLoginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Locations() DataGoogleworkspaceUserLocationsList {
	var returns DataGoogleworkspaceUserLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Name() DataGoogleworkspaceUserNameList {
	var returns DataGoogleworkspaceUserNameList
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) NonEditableAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonEditableAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Organizations() DataGoogleworkspaceUserOrganizationsList {
	var returns DataGoogleworkspaceUserOrganizationsList
	_jsii_.Get(
		j,
		"organizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) OrgUnitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Phones() DataGoogleworkspaceUserPhonesList {
	var returns DataGoogleworkspaceUserPhonesList
	_jsii_.Get(
		j,
		"phones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) PosixAccounts() DataGoogleworkspaceUserPosixAccountsList {
	var returns DataGoogleworkspaceUserPosixAccountsList
	_jsii_.Get(
		j,
		"posixAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) PrimaryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) PrimaryEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) RecoveryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) RecoveryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Relations() DataGoogleworkspaceUserRelationsList {
	var returns DataGoogleworkspaceUserRelationsList
	_jsii_.Get(
		j,
		"relations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) SshPublicKeys() DataGoogleworkspaceUserSshPublicKeysList {
	var returns DataGoogleworkspaceUserSshPublicKeysList
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Suspended() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) SuspensionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspensionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ThumbnailPhotoEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) ThumbnailPhotoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUser) Websites() DataGoogleworkspaceUserWebsitesList {
	var returns DataGoogleworkspaceUserWebsitesList
	_jsii_.Get(
		j,
		"websites",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/googleworkspace/d/user googleworkspace_user} Data Source.
func NewDataGoogleworkspaceUser(scope constructs.Construct, id *string, config *DataGoogleworkspaceUserConfig) DataGoogleworkspaceUser {
	_init_.Initialize()

	if err := validateNewDataGoogleworkspaceUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleworkspaceUser{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/googleworkspace/d/user googleworkspace_user} Data Source.
func NewDataGoogleworkspaceUser_Override(d DataGoogleworkspaceUser, scope constructs.Construct, id *string, config *DataGoogleworkspaceUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetPrimaryEmail(val *string) {
	if err := j.validateSetPrimaryEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryEmail",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
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
func DataGoogleworkspaceUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleworkspaceUser_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceUser_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleworkspaceUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleworkspaceUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUser.DataGoogleworkspaceUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ResetPrimaryEmail() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryEmail",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

