package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v4/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v4/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserPosixAccountsOutputReference interface {
	cdktf.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
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
	// Experimental.
	Fqn() *string
	Gecos() *string
	SetGecos(val *string)
	GecosInput() *string
	Gid() *string
	SetGid(val *string)
	GidInput() *string
	HomeDirectory() *string
	SetHomeDirectory(val *string)
	HomeDirectoryInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperatingSystemType() *string
	SetOperatingSystemType(val *string)
	OperatingSystemTypeInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Shell() *string
	SetShell(val *string)
	ShellInput() *string
	SystemId() *string
	SetSystemId(val *string)
	SystemIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uid() *string
	SetUid(val *string)
	UidInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetAccountId()
	ResetGecos()
	ResetGid()
	ResetHomeDirectory()
	ResetOperatingSystemType()
	ResetPrimary()
	ResetShell()
	ResetSystemId()
	ResetUid()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UserPosixAccountsOutputReference
type jsiiProxy_UserPosixAccountsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Gecos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gecos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) GecosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gecosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) HomeDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) HomeDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) OperatingSystemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) OperatingSystemTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Shell() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shell",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) ShellInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shellInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) SystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) SystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPosixAccountsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewUserPosixAccountsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) UserPosixAccountsOutputReference {
	_init_.Initialize()

	if err := validateNewUserPosixAccountsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPosixAccountsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserPosixAccountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewUserPosixAccountsOutputReference_Override(u UserPosixAccountsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserPosixAccountsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		u,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetGecos(val *string) {
	if err := j.validateSetGecosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gecos",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetHomeDirectory(val *string) {
	if err := j.validateSetHomeDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homeDirectory",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetOperatingSystemType(val *string) {
	if err := j.validateSetOperatingSystemTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystemType",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetShell(val *string) {
	if err := j.validateSetShellParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shell",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetSystemId(val *string) {
	if err := j.validateSetSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemId",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (j *jsiiProxy_UserPosixAccountsOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UserPosixAccountsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := u.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetAccountId() {
	_jsii_.InvokeVoid(
		u,
		"resetAccountId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetGecos() {
	_jsii_.InvokeVoid(
		u,
		"resetGecos",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		u,
		"resetGid",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetHomeDirectory() {
	_jsii_.InvokeVoid(
		u,
		"resetHomeDirectory",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetOperatingSystemType() {
	_jsii_.InvokeVoid(
		u,
		"resetOperatingSystemType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		u,
		"resetPrimary",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetShell() {
	_jsii_.InvokeVoid(
		u,
		"resetShell",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetSystemId() {
	_jsii_.InvokeVoid(
		u,
		"resetSystemId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		u,
		"resetUid",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		u,
		"resetUsername",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := u.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		u,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPosixAccountsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

