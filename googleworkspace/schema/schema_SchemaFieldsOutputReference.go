package schema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v2/jsii"

	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v2/schema/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchemaFieldsOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Etag() *string
	FieldId() *string
	FieldName() *string
	SetFieldName(val *string)
	FieldNameInput() *string
	FieldType() *string
	SetFieldType(val *string)
	FieldTypeInput() *string
	// Experimental.
	Fqn() *string
	Indexed() interface{}
	SetIndexed(val interface{})
	IndexedInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiValued() interface{}
	SetMultiValued(val interface{})
	MultiValuedInput() interface{}
	NumericIndexingSpec() SchemaFieldsNumericIndexingSpecOutputReference
	NumericIndexingSpecInput() *SchemaFieldsNumericIndexingSpec
	ReadAccessType() *string
	SetReadAccessType(val *string)
	ReadAccessTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutNumericIndexingSpec(value *SchemaFieldsNumericIndexingSpec)
	ResetDisplayName()
	ResetIndexed()
	ResetMultiValued()
	ResetNumericIndexingSpec()
	ResetReadAccessType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SchemaFieldsOutputReference
type jsiiProxy_SchemaFieldsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SchemaFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) FieldId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) FieldType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) FieldTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) Indexed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"indexed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) IndexedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"indexedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) MultiValued() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValued",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) MultiValuedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiValuedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) NumericIndexingSpec() SchemaFieldsNumericIndexingSpecOutputReference {
	var returns SchemaFieldsNumericIndexingSpecOutputReference
	_jsii_.Get(
		j,
		"numericIndexingSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) NumericIndexingSpecInput() *SchemaFieldsNumericIndexingSpec {
	var returns *SchemaFieldsNumericIndexingSpec
	_jsii_.Get(
		j,
		"numericIndexingSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) ReadAccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) ReadAccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readAccessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SchemaFieldsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSchemaFieldsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SchemaFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewSchemaFieldsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchemaFieldsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.schema.SchemaFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSchemaFieldsOutputReference_Override(s SchemaFieldsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.schema.SchemaFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetFieldType(val *string) {
	if err := j.validateSetFieldTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldType",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetIndexed(val interface{}) {
	if err := j.validateSetIndexedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexed",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetMultiValued(val interface{}) {
	if err := j.validateSetMultiValuedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiValued",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetReadAccessType(val *string) {
	if err := j.validateSetReadAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readAccessType",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SchemaFieldsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) PutNumericIndexingSpec(value *SchemaFieldsNumericIndexingSpec) {
	if err := s.validatePutNumericIndexingSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNumericIndexingSpec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		s,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ResetIndexed() {
	_jsii_.InvokeVoid(
		s,
		"resetIndexed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ResetMultiValued() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiValued",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ResetNumericIndexingSpec() {
	_jsii_.InvokeVoid(
		s,
		"resetNumericIndexingSpec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ResetReadAccessType() {
	_jsii_.InvokeVoid(
		s,
		"resetReadAccessType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SchemaFieldsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

