package datagoogleworkspaceusers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/datagoogleworkspaceusers/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleworkspaceUsersUsersOutputReference interface {
	cdktf.ComplexObject
	Addresses() DataGoogleworkspaceUsersUsersAddressesList
	AgreedToTerms() cdktf.IResolvable
	Aliases() *[]*string
	Archived() cdktf.IResolvable
	ChangePasswordAtNextLogin() cdktf.IResolvable
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CreationTime() *string
	CustomerId() *string
	CustomSchemas() DataGoogleworkspaceUsersUsersCustomSchemasList
	DeletionTime() *string
	Emails() DataGoogleworkspaceUsersUsersEmailsList
	Etag() *string
	ExternalIds() DataGoogleworkspaceUsersUsersExternalIdsList
	// Experimental.
	Fqn() *string
	HashFunction() *string
	Id() *string
	Ims() DataGoogleworkspaceUsersUsersImsList
	IncludeInGlobalAddressList() cdktf.IResolvable
	InternalValue() *DataGoogleworkspaceUsersUsers
	SetInternalValue(val *DataGoogleworkspaceUsersUsers)
	IpAllowlist() cdktf.IResolvable
	IsAdmin() cdktf.IResolvable
	IsDelegatedAdmin() cdktf.IResolvable
	IsEnforcedIn2StepVerification() cdktf.IResolvable
	IsEnrolledIn2StepVerification() cdktf.IResolvable
	IsMailboxSetup() cdktf.IResolvable
	Keywords() DataGoogleworkspaceUsersUsersKeywordsList
	Languages() DataGoogleworkspaceUsersUsersLanguagesList
	LastLoginTime() *string
	Locations() DataGoogleworkspaceUsersUsersLocationsList
	Name() DataGoogleworkspaceUsersUsersNameList
	NonEditableAliases() *[]*string
	Organizations() DataGoogleworkspaceUsersUsersOrganizationsList
	OrgUnitPath() *string
	Password() *string
	Phones() DataGoogleworkspaceUsersUsersPhonesList
	PosixAccounts() DataGoogleworkspaceUsersUsersPosixAccountsList
	PrimaryEmail() *string
	RecoveryEmail() *string
	RecoveryPhone() *string
	Relations() DataGoogleworkspaceUsersUsersRelationsList
	SshPublicKeys() DataGoogleworkspaceUsersUsersSshPublicKeysList
	Suspended() cdktf.IResolvable
	SuspensionReason() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThumbnailPhotoEtag() *string
	ThumbnailPhotoUrl() *string
	Websites() DataGoogleworkspaceUsersUsersWebsitesList
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleworkspaceUsersUsersOutputReference
type jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Addresses() DataGoogleworkspaceUsersUsersAddressesList {
	var returns DataGoogleworkspaceUsersUsersAddressesList
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) AgreedToTerms() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"agreedToTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Archived() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ChangePasswordAtNextLogin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"changePasswordAtNextLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) CustomerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) CustomSchemas() DataGoogleworkspaceUsersUsersCustomSchemasList {
	var returns DataGoogleworkspaceUsersUsersCustomSchemasList
	_jsii_.Get(
		j,
		"customSchemas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) DeletionTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Emails() DataGoogleworkspaceUsersUsersEmailsList {
	var returns DataGoogleworkspaceUsersUsersEmailsList
	_jsii_.Get(
		j,
		"emails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ExternalIds() DataGoogleworkspaceUsersUsersExternalIdsList {
	var returns DataGoogleworkspaceUsersUsersExternalIdsList
	_jsii_.Get(
		j,
		"externalIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) HashFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Ims() DataGoogleworkspaceUsersUsersImsList {
	var returns DataGoogleworkspaceUsersUsersImsList
	_jsii_.Get(
		j,
		"ims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IncludeInGlobalAddressList() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeInGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) InternalValue() *DataGoogleworkspaceUsersUsers {
	var returns *DataGoogleworkspaceUsersUsers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IpAllowlist() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ipAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IsAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IsDelegatedAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isDelegatedAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IsEnforcedIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnforcedIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IsEnrolledIn2StepVerification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isEnrolledIn2StepVerification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) IsMailboxSetup() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isMailboxSetup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Keywords() DataGoogleworkspaceUsersUsersKeywordsList {
	var returns DataGoogleworkspaceUsersUsersKeywordsList
	_jsii_.Get(
		j,
		"keywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Languages() DataGoogleworkspaceUsersUsersLanguagesList {
	var returns DataGoogleworkspaceUsersUsersLanguagesList
	_jsii_.Get(
		j,
		"languages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) LastLoginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastLoginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Locations() DataGoogleworkspaceUsersUsersLocationsList {
	var returns DataGoogleworkspaceUsersUsersLocationsList
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Name() DataGoogleworkspaceUsersUsersNameList {
	var returns DataGoogleworkspaceUsersUsersNameList
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) NonEditableAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonEditableAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Organizations() DataGoogleworkspaceUsersUsersOrganizationsList {
	var returns DataGoogleworkspaceUsersUsersOrganizationsList
	_jsii_.Get(
		j,
		"organizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) OrgUnitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Phones() DataGoogleworkspaceUsersUsersPhonesList {
	var returns DataGoogleworkspaceUsersUsersPhonesList
	_jsii_.Get(
		j,
		"phones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) PosixAccounts() DataGoogleworkspaceUsersUsersPosixAccountsList {
	var returns DataGoogleworkspaceUsersUsersPosixAccountsList
	_jsii_.Get(
		j,
		"posixAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) PrimaryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) RecoveryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) RecoveryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Relations() DataGoogleworkspaceUsersUsersRelationsList {
	var returns DataGoogleworkspaceUsersUsersRelationsList
	_jsii_.Get(
		j,
		"relations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) SshPublicKeys() DataGoogleworkspaceUsersUsersSshPublicKeysList {
	var returns DataGoogleworkspaceUsersUsersSshPublicKeysList
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Suspended() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) SuspensionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspensionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ThumbnailPhotoEtag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoEtag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ThumbnailPhotoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbnailPhotoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Websites() DataGoogleworkspaceUsersUsersWebsitesList {
	var returns DataGoogleworkspaceUsersUsersWebsitesList
	_jsii_.Get(
		j,
		"websites",
		&returns,
	)
	return returns
}


func NewDataGoogleworkspaceUsersUsersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleworkspaceUsersUsersOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleworkspaceUsersUsersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUsers.DataGoogleworkspaceUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleworkspaceUsersUsersOutputReference_Override(d DataGoogleworkspaceUsersUsersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceUsers.DataGoogleworkspaceUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference)SetInternalValue(val *DataGoogleworkspaceUsersUsers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceUsersUsersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

