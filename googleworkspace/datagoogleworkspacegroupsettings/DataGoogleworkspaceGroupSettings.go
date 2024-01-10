// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleworkspacegroupsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/datagoogleworkspacegroupsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/data-sources/group_settings googleworkspace_group_settings}.
type DataGoogleworkspaceGroupSettings interface {
	cdktf.TerraformDataSource
	AllowExternalMembers() cdktf.IResolvable
	AllowWebPosting() cdktf.IResolvable
	ArchiveOnly() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomFooterText() *string
	CustomReplyTo() *string
	CustomRolesEnabledForSettingsToBeMerged() cdktf.IResolvable
	DefaultMessageDenyNotificationText() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	EnableCollaborativeInbox() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludeCustomFooter() cdktf.IResolvable
	IncludeInGlobalAddressList() cdktf.IResolvable
	IsArchived() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MembersCanPostAsTheGroup() cdktf.IResolvable
	MessageModerationLevel() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrimaryLanguage() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReplyTo() *string
	SendMessageDenyNotification() cdktf.IResolvable
	SpamModerationLevel() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WhoCanAssistContent() *string
	WhoCanContactOwner() *string
	WhoCanDiscoverGroup() *string
	WhoCanJoin() *string
	WhoCanLeaveGroup() *string
	WhoCanModerateContent() *string
	WhoCanModerateMembers() *string
	WhoCanPostMessage() *string
	WhoCanViewGroup() *string
	WhoCanViewMembership() *string
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataGoogleworkspaceGroupSettings
type jsiiProxy_DataGoogleworkspaceGroupSettings struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) AllowExternalMembers() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowExternalMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) AllowWebPosting() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowWebPosting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) ArchiveOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"archiveOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) CustomFooterText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFooterText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) CustomReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customReplyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) CustomRolesEnabledForSettingsToBeMerged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"customRolesEnabledForSettingsToBeMerged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) DefaultMessageDenyNotificationText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageDenyNotificationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) EnableCollaborativeInbox() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableCollaborativeInbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) IncludeCustomFooter() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeCustomFooter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) IncludeInGlobalAddressList() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"includeInGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) IsArchived() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isArchived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) MembersCanPostAsTheGroup() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"membersCanPostAsTheGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) MessageModerationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageModerationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) PrimaryLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) SendMessageDenyNotification() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sendMessageDenyNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) SpamModerationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamModerationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanAssistContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanAssistContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanContactOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanContactOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanDiscoverGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanDiscoverGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanJoin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanLeaveGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanLeaveGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanModerateContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanModerateMembers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanPostMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanPostMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanViewGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings) WhoCanViewMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewMembership",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/data-sources/group_settings googleworkspace_group_settings} Data Source.
func NewDataGoogleworkspaceGroupSettings(scope constructs.Construct, id *string, config *DataGoogleworkspaceGroupSettingsConfig) DataGoogleworkspaceGroupSettings {
	_init_.Initialize()

	if err := validateNewDataGoogleworkspaceGroupSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleworkspaceGroupSettings{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/data-sources/group_settings googleworkspace_group_settings} Data Source.
func NewDataGoogleworkspaceGroupSettings_Override(d DataGoogleworkspaceGroupSettings, scope constructs.Construct, id *string, config *DataGoogleworkspaceGroupSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleworkspaceGroupSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleworkspaceGroupSettings resource upon running "cdktf plan <stack-name>".
func DataGoogleworkspaceGroupSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceGroupSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
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
func DataGoogleworkspaceGroupSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceGroupSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleworkspaceGroupSettings_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceGroupSettings_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleworkspaceGroupSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleworkspaceGroupSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleworkspaceGroupSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googleworkspace.dataGoogleworkspaceGroupSettings.DataGoogleworkspaceGroupSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleworkspaceGroupSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

