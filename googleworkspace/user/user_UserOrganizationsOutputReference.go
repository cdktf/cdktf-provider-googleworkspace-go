package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserOrganizationsOutputReference interface {
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
	CostCenter() *string
	SetCostCenter(val *string)
	CostCenterInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomType() *string
	SetCustomType(val *string)
	CustomTypeInput() *string
	Department() *string
	SetDepartment(val *string)
	DepartmentInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	FullTimeEquivalent() *float64
	SetFullTimeEquivalent(val *float64)
	FullTimeEquivalentInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Symbol() *string
	SetSymbol(val *string)
	SymbolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	ResetCostCenter()
	ResetCustomType()
	ResetDepartment()
	ResetDescription()
	ResetDomain()
	ResetFullTimeEquivalent()
	ResetLocation()
	ResetName()
	ResetPrimary()
	ResetSymbol()
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UserOrganizationsOutputReference
type jsiiProxy_UserOrganizationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_UserOrganizationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) CostCenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) CustomType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) CustomTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) DepartmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"departmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) FullTimeEquivalent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullTimeEquivalent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) FullTimeEquivalentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullTimeEquivalentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Symbol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symbol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) SymbolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symbolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserOrganizationsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewUserOrganizationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) UserOrganizationsOutputReference {
	_init_.Initialize()

	if err := validateNewUserOrganizationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserOrganizationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserOrganizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewUserOrganizationsOutputReference_Override(u UserOrganizationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserOrganizationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		u,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetCostCenter(val *string) {
	if err := j.validateSetCostCenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"costCenter",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetCustomType(val *string) {
	if err := j.validateSetCustomTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customType",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetDepartment(val *string) {
	if err := j.validateSetDepartmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"department",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetFullTimeEquivalent(val *float64) {
	if err := j.validateSetFullTimeEquivalentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullTimeEquivalent",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetSymbol(val *string) {
	if err := j.validateSetSymbolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symbol",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_UserOrganizationsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserOrganizationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserOrganizationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetCostCenter() {
	_jsii_.InvokeVoid(
		u,
		"resetCostCenter",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetCustomType() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetDepartment() {
	_jsii_.InvokeVoid(
		u,
		"resetDepartment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		u,
		"resetDescription",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetDomain() {
	_jsii_.InvokeVoid(
		u,
		"resetDomain",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetFullTimeEquivalent() {
	_jsii_.InvokeVoid(
		u,
		"resetFullTimeEquivalent",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		u,
		"resetLocation",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		u,
		"resetName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		u,
		"resetPrimary",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetSymbol() {
	_jsii_.InvokeVoid(
		u,
		"resetSymbol",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		u,
		"resetTitle",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserOrganizationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (u *jsiiProxy_UserOrganizationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

