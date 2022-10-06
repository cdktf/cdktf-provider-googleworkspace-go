package user

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/user/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserLocationsOutputReference interface {
	cdktf.ComplexObject
	Area() *string
	SetArea(val *string)
	AreaInput() *string
	BuildingId() *string
	SetBuildingId(val *string)
	BuildingIdInput() *string
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
	CustomType() *string
	SetCustomType(val *string)
	CustomTypeInput() *string
	DeskCode() *string
	SetDeskCode(val *string)
	DeskCodeInput() *string
	FloorName() *string
	SetFloorName(val *string)
	FloorNameInput() *string
	FloorSection() *string
	SetFloorSection(val *string)
	FloorSectionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetArea()
	ResetBuildingId()
	ResetCustomType()
	ResetDeskCode()
	ResetFloorName()
	ResetFloorSection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UserLocationsOutputReference
type jsiiProxy_UserLocationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_UserLocationsOutputReference) Area() *string {
	var returns *string
	_jsii_.Get(
		j,
		"area",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) AreaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"areaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) BuildingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) BuildingIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildingIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) CustomType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) CustomTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) DeskCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deskCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) DeskCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deskCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) FloorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) FloorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) FloorSection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floorSection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) FloorSectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"floorSectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserLocationsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewUserLocationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) UserLocationsOutputReference {
	_init_.Initialize()

	if err := validateNewUserLocationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserLocationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewUserLocationsOutputReference_Override(u UserLocationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.user.UserLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		u,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetArea(val *string) {
	if err := j.validateSetAreaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"area",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetBuildingId(val *string) {
	if err := j.validateSetBuildingIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildingId",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetCustomType(val *string) {
	if err := j.validateSetCustomTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customType",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetDeskCode(val *string) {
	if err := j.validateSetDeskCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deskCode",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetFloorName(val *string) {
	if err := j.validateSetFloorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"floorName",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetFloorSection(val *string) {
	if err := j.validateSetFloorSectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"floorSection",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_UserLocationsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserLocationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (u *jsiiProxy_UserLocationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (u *jsiiProxy_UserLocationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		u,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserLocationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (u *jsiiProxy_UserLocationsOutputReference) ResetArea() {
	_jsii_.InvokeVoid(
		u,
		"resetArea",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ResetBuildingId() {
	_jsii_.InvokeVoid(
		u,
		"resetBuildingId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ResetCustomType() {
	_jsii_.InvokeVoid(
		u,
		"resetCustomType",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ResetDeskCode() {
	_jsii_.InvokeVoid(
		u,
		"resetDeskCode",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ResetFloorName() {
	_jsii_.InvokeVoid(
		u,
		"resetFloorName",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) ResetFloorSection() {
	_jsii_.InvokeVoid(
		u,
		"resetFloorSection",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UserLocationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (u *jsiiProxy_UserLocationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

