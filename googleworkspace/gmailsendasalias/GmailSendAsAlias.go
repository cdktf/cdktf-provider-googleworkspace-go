// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gmailsendasalias

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/gmailsendasalias/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias googleworkspace_gmail_send_as_alias}.
type GmailSendAsAlias interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	IsPrimary() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	ReplyToAddress() *string
	SetReplyToAddress(val *string)
	ReplyToAddressInput() *string
	SendAsEmail() *string
	SetSendAsEmail(val *string)
	SendAsEmailInput() *string
	Signature() *string
	SetSignature(val *string)
	SignatureInput() *string
	SmtpMsa() GmailSendAsAliasSmtpMsaOutputReference
	SmtpMsaInput() *GmailSendAsAliasSmtpMsa
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TreatAsAlias() interface{}
	SetTreatAsAlias(val interface{})
	TreatAsAliasInput() interface{}
	VerificationStatus() *string
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
	PutSmtpMsa(value *GmailSendAsAliasSmtpMsa)
	ResetDisplayName()
	ResetIsDefault()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplyToAddress()
	ResetSignature()
	ResetSmtpMsa()
	ResetTreatAsAlias()
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

// The jsii proxy struct for GmailSendAsAlias
type jsiiProxy_GmailSendAsAlias struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GmailSendAsAlias) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) IsPrimary() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) PrimaryEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) PrimaryEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) ReplyToAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) ReplyToAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) SendAsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sendAsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) SendAsEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sendAsEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) Signature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) SignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) SmtpMsa() GmailSendAsAliasSmtpMsaOutputReference {
	var returns GmailSendAsAliasSmtpMsaOutputReference
	_jsii_.Get(
		j,
		"smtpMsa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) SmtpMsaInput() *GmailSendAsAliasSmtpMsa {
	var returns *GmailSendAsAliasSmtpMsa
	_jsii_.Get(
		j,
		"smtpMsaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) TreatAsAlias() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatAsAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) TreatAsAliasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatAsAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GmailSendAsAlias) VerificationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationStatus",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias googleworkspace_gmail_send_as_alias} Resource.
func NewGmailSendAsAlias(scope constructs.Construct, id *string, config *GmailSendAsAliasConfig) GmailSendAsAlias {
	_init_.Initialize()

	if err := validateNewGmailSendAsAliasParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GmailSendAsAlias{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/gmail_send_as_alias googleworkspace_gmail_send_as_alias} Resource.
func NewGmailSendAsAlias_Override(g GmailSendAsAlias, scope constructs.Construct, id *string, config *GmailSendAsAliasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetIsDefault(val interface{}) {
	if err := j.validateSetIsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetPrimaryEmail(val *string) {
	if err := j.validateSetPrimaryEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryEmail",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetReplyToAddress(val *string) {
	if err := j.validateSetReplyToAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyToAddress",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetSendAsEmail(val *string) {
	if err := j.validateSetSendAsEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendAsEmail",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetSignature(val *string) {
	if err := j.validateSetSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signature",
		val,
	)
}

func (j *jsiiProxy_GmailSendAsAlias)SetTreatAsAlias(val interface{}) {
	if err := j.validateSetTreatAsAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatAsAlias",
		val,
	)
}

// Generates CDKTF code for importing a GmailSendAsAlias resource upon running "cdktf plan <stack-name>".
func GmailSendAsAlias_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGmailSendAsAlias_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
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
func GmailSendAsAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGmailSendAsAlias_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GmailSendAsAlias_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGmailSendAsAlias_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GmailSendAsAlias_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGmailSendAsAlias_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GmailSendAsAlias_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googleworkspace.gmailSendAsAlias.GmailSendAsAlias",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) PutSmtpMsa(value *GmailSendAsAliasSmtpMsa) {
	if err := g.validatePutSmtpMsaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSmtpMsa",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetIsDefault() {
	_jsii_.InvokeVoid(
		g,
		"resetIsDefault",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetReplyToAddress() {
	_jsii_.InvokeVoid(
		g,
		"resetReplyToAddress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetSignature() {
	_jsii_.InvokeVoid(
		g,
		"resetSignature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetSmtpMsa() {
	_jsii_.InvokeVoid(
		g,
		"resetSmtpMsa",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) ResetTreatAsAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetTreatAsAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GmailSendAsAlias) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GmailSendAsAlias) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

