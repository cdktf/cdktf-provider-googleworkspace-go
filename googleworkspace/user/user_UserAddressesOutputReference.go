package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v3/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v3/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserAddressesOutputReference interface {
	cdktf.ComplexObject
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
	Country() *string
	SetCountry(val *string)
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	CountryInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomType() *string
	SetCustomType(val *string)
	CustomTypeInput() *string
	ExtendedAddress() *string
	SetExtendedAddress(val *string)
	ExtendedAddressInput() *string
	Formatted() *string
	SetFormatted(val *string)
	FormattedInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	PoBox() *string
	SetPoBox(val *string)
	PoBoxInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SourceIsStructured() interface{}
	SetSourceIsStructured(val interface{})
	SourceIsStructuredInput() interface{}
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetCountry()
	ResetCountryCode()
	ResetCustomType()
	ResetExtendedAddress()
	ResetFormatted()
	ResetLocality()
	ResetPoBox()
	ResetPostalCode()
	ResetPrimary()
	ResetRegion()
	ResetSourceIsStructured()
	ResetStreetAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UserAddressesOutputReference
type jsiiProxy_UserAddressesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_UserAddressesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CustomType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) CustomTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) ExtendedAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) ExtendedAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Formatted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) FormattedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formattedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) PoBox() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poBox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) PoBoxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poBoxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) SourceIsStructured() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIsStructured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) SourceIsStructuredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceIsStructuredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserAddressesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewUserAddressesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) UserAddressesOutputReference {
	_init_.Initialize()

	if err := validateNewUserAddressesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserAddressesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewUserAddressesOutputReference_Override(u UserAddressesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserAddressesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		u,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetCustomType(val *string) {
	if err := j.validateSetCustomTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customType",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetExtendedAddress(val *string) {
	if err := j.validateSetExtendedAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedAddress",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetFormatted(val *string) {
	if err := j.validateSetFormattedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatted",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetPoBox(val *string) {
	if err := j.validateSetPoBoxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poBox",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetSourceIsStructured(val interface{}) {
	if err := j.validateSetSourceIsStructuredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIsStructured",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetStreetAddress(val *string) {
	if err := j.validateSetStreetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_UserAddressesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserAddressesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UserAddressesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UserAddressesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserAddressesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserAddressesOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		u,
		"resetCountry",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		u,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetCustomType() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetExtendedAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetExtendedAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetFormatted() {
	_jsii_.InvokeVoid(
		u,
		"resetFormatted",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		u,
		"resetLocality",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetPoBox() {
	_jsii_.InvokeVoid(
		u,
		"resetPoBox",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		u,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		u,
		"resetPrimary",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		u,
		"resetRegion",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetSourceIsStructured() {
	_jsii_.InvokeVoid(
		u,
		"resetSourceIsStructured",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		u,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserAddressesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (u *jsiiProxy_UserAddressesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

