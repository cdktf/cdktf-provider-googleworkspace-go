// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googleworkspace-go/googleworkspace/v8/groupsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings googleworkspace_group_settings}.
type GroupSettings interface {
	cdktf.TerraformResource
	AllowExternalMembers() interface{}
	SetAllowExternalMembers(val interface{})
	AllowExternalMembersInput() interface{}
	AllowWebPosting() interface{}
	SetAllowWebPosting(val interface{})
	AllowWebPostingInput() interface{}
	ArchiveOnly() interface{}
	SetArchiveOnly(val interface{})
	ArchiveOnlyInput() interface{}
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
	CustomFooterText() *string
	SetCustomFooterText(val *string)
	CustomFooterTextInput() *string
	CustomReplyTo() *string
	SetCustomReplyTo(val *string)
	CustomReplyToInput() *string
	CustomRolesEnabledForSettingsToBeMerged() cdktf.IResolvable
	DefaultMessageDenyNotificationText() *string
	SetDefaultMessageDenyNotificationText(val *string)
	DefaultMessageDenyNotificationTextInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	EnableCollaborativeInbox() interface{}
	SetEnableCollaborativeInbox(val interface{})
	EnableCollaborativeInboxInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IncludeCustomFooter() interface{}
	SetIncludeCustomFooter(val interface{})
	IncludeCustomFooterInput() interface{}
	IncludeInGlobalAddressList() interface{}
	SetIncludeInGlobalAddressList(val interface{})
	IncludeInGlobalAddressListInput() interface{}
	IsArchived() interface{}
	SetIsArchived(val interface{})
	IsArchivedInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MembersCanPostAsTheGroup() interface{}
	SetMembersCanPostAsTheGroup(val interface{})
	MembersCanPostAsTheGroupInput() interface{}
	MessageModerationLevel() *string
	SetMessageModerationLevel(val *string)
	MessageModerationLevelInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	PrimaryLanguage() *string
	SetPrimaryLanguage(val *string)
	PrimaryLanguageInput() *string
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
	ReplyTo() *string
	SetReplyTo(val *string)
	ReplyToInput() *string
	SendMessageDenyNotification() interface{}
	SetSendMessageDenyNotification(val interface{})
	SendMessageDenyNotificationInput() interface{}
	SpamModerationLevel() *string
	SetSpamModerationLevel(val *string)
	SpamModerationLevelInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GroupSettingsTimeoutsOutputReference
	TimeoutsInput() interface{}
	WhoCanAssistContent() *string
	SetWhoCanAssistContent(val *string)
	WhoCanAssistContentInput() *string
	WhoCanContactOwner() *string
	SetWhoCanContactOwner(val *string)
	WhoCanContactOwnerInput() *string
	WhoCanDiscoverGroup() *string
	SetWhoCanDiscoverGroup(val *string)
	WhoCanDiscoverGroupInput() *string
	WhoCanJoin() *string
	SetWhoCanJoin(val *string)
	WhoCanJoinInput() *string
	WhoCanLeaveGroup() *string
	SetWhoCanLeaveGroup(val *string)
	WhoCanLeaveGroupInput() *string
	WhoCanModerateContent() *string
	SetWhoCanModerateContent(val *string)
	WhoCanModerateContentInput() *string
	WhoCanModerateMembers() *string
	SetWhoCanModerateMembers(val *string)
	WhoCanModerateMembersInput() *string
	WhoCanPostMessage() *string
	SetWhoCanPostMessage(val *string)
	WhoCanPostMessageInput() *string
	WhoCanViewGroup() *string
	SetWhoCanViewGroup(val *string)
	WhoCanViewGroupInput() *string
	WhoCanViewMembership() *string
	SetWhoCanViewMembership(val *string)
	WhoCanViewMembershipInput() *string
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
	PutTimeouts(value *GroupSettingsTimeouts)
	ResetAllowExternalMembers()
	ResetAllowWebPosting()
	ResetArchiveOnly()
	ResetCustomFooterText()
	ResetCustomReplyTo()
	ResetDefaultMessageDenyNotificationText()
	ResetEnableCollaborativeInbox()
	ResetIncludeCustomFooter()
	ResetIncludeInGlobalAddressList()
	ResetIsArchived()
	ResetMembersCanPostAsTheGroup()
	ResetMessageModerationLevel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryLanguage()
	ResetReplyTo()
	ResetSendMessageDenyNotification()
	ResetSpamModerationLevel()
	ResetTimeouts()
	ResetWhoCanAssistContent()
	ResetWhoCanContactOwner()
	ResetWhoCanDiscoverGroup()
	ResetWhoCanJoin()
	ResetWhoCanLeaveGroup()
	ResetWhoCanModerateContent()
	ResetWhoCanModerateMembers()
	ResetWhoCanPostMessage()
	ResetWhoCanViewGroup()
	ResetWhoCanViewMembership()
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

// The jsii proxy struct for GroupSettings
type jsiiProxy_GroupSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GroupSettings) AllowExternalMembers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) AllowExternalMembersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) AllowWebPosting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWebPosting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) AllowWebPostingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowWebPostingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ArchiveOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ArchiveOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CustomFooterText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFooterText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CustomFooterTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFooterTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CustomReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customReplyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CustomReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customReplyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) CustomRolesEnabledForSettingsToBeMerged() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"customRolesEnabledForSettingsToBeMerged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) DefaultMessageDenyNotificationText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageDenyNotificationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) DefaultMessageDenyNotificationTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageDenyNotificationTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) EnableCollaborativeInbox() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCollaborativeInbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) EnableCollaborativeInboxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCollaborativeInboxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IncludeCustomFooter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCustomFooter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IncludeCustomFooterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCustomFooterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IncludeInGlobalAddressList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInGlobalAddressList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IncludeInGlobalAddressListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInGlobalAddressListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IsArchived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isArchived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) IsArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isArchivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) MembersCanPostAsTheGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanPostAsTheGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) MembersCanPostAsTheGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanPostAsTheGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) MessageModerationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageModerationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) MessageModerationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageModerationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) PrimaryLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) PrimaryLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ReplyTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) ReplyToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replyToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) SendMessageDenyNotification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendMessageDenyNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) SendMessageDenyNotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendMessageDenyNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) SpamModerationLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamModerationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) SpamModerationLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spamModerationLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) Timeouts() GroupSettingsTimeoutsOutputReference {
	var returns GroupSettingsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanAssistContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanAssistContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanAssistContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanAssistContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanContactOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanContactOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanContactOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanContactOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanDiscoverGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanDiscoverGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanDiscoverGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanDiscoverGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanJoin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanJoin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanJoinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanJoinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanLeaveGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanLeaveGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanLeaveGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanLeaveGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanModerateContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanModerateContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanModerateMembers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanModerateMembersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanModerateMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanPostMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanPostMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanPostMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanPostMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanViewGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanViewGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanViewMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GroupSettings) WhoCanViewMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whoCanViewMembershipInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings googleworkspace_group_settings} Resource.
func NewGroupSettings(scope constructs.Construct, id *string, config *GroupSettingsConfig) GroupSettings {
	_init_.Initialize()

	if err := validateNewGroupSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GroupSettings{}

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings googleworkspace_group_settings} Resource.
func NewGroupSettings_Override(g GroupSettings, scope constructs.Construct, id *string, config *GroupSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GroupSettings)SetAllowExternalMembers(val interface{}) {
	if err := j.validateSetAllowExternalMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExternalMembers",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetAllowWebPosting(val interface{}) {
	if err := j.validateSetAllowWebPostingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowWebPosting",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetArchiveOnly(val interface{}) {
	if err := j.validateSetArchiveOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveOnly",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetCustomFooterText(val *string) {
	if err := j.validateSetCustomFooterTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFooterText",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetCustomReplyTo(val *string) {
	if err := j.validateSetCustomReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customReplyTo",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetDefaultMessageDenyNotificationText(val *string) {
	if err := j.validateSetDefaultMessageDenyNotificationTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMessageDenyNotificationText",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetEnableCollaborativeInbox(val interface{}) {
	if err := j.validateSetEnableCollaborativeInboxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCollaborativeInbox",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetIncludeCustomFooter(val interface{}) {
	if err := j.validateSetIncludeCustomFooterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeCustomFooter",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetIncludeInGlobalAddressList(val interface{}) {
	if err := j.validateSetIncludeInGlobalAddressListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeInGlobalAddressList",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetIsArchived(val interface{}) {
	if err := j.validateSetIsArchivedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isArchived",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetMembersCanPostAsTheGroup(val interface{}) {
	if err := j.validateSetMembersCanPostAsTheGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanPostAsTheGroup",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetMessageModerationLevel(val *string) {
	if err := j.validateSetMessageModerationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageModerationLevel",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetPrimaryLanguage(val *string) {
	if err := j.validateSetPrimaryLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryLanguage",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetReplyTo(val *string) {
	if err := j.validateSetReplyToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replyTo",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetSendMessageDenyNotification(val interface{}) {
	if err := j.validateSetSendMessageDenyNotificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendMessageDenyNotification",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetSpamModerationLevel(val *string) {
	if err := j.validateSetSpamModerationLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spamModerationLevel",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanAssistContent(val *string) {
	if err := j.validateSetWhoCanAssistContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanAssistContent",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanContactOwner(val *string) {
	if err := j.validateSetWhoCanContactOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanContactOwner",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanDiscoverGroup(val *string) {
	if err := j.validateSetWhoCanDiscoverGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanDiscoverGroup",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanJoin(val *string) {
	if err := j.validateSetWhoCanJoinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanJoin",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanLeaveGroup(val *string) {
	if err := j.validateSetWhoCanLeaveGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanLeaveGroup",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanModerateContent(val *string) {
	if err := j.validateSetWhoCanModerateContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanModerateContent",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanModerateMembers(val *string) {
	if err := j.validateSetWhoCanModerateMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanModerateMembers",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanPostMessage(val *string) {
	if err := j.validateSetWhoCanPostMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanPostMessage",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanViewGroup(val *string) {
	if err := j.validateSetWhoCanViewGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanViewGroup",
		val,
	)
}

func (j *jsiiProxy_GroupSettings)SetWhoCanViewMembership(val *string) {
	if err := j.validateSetWhoCanViewMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"whoCanViewMembership",
		val,
	)
}

// Generates CDKTF code for importing a GroupSettings resource upon running "cdktf plan <stack-name>".
func GroupSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGroupSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
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
func GroupSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GroupSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGroupSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GroupSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GroupSettings) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GroupSettings) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GroupSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GroupSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GroupSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GroupSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GroupSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GroupSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GroupSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GroupSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GroupSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GroupSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GroupSettings) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GroupSettings) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GroupSettings) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GroupSettings) PutTimeouts(value *GroupSettingsTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GroupSettings) ResetAllowExternalMembers() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowExternalMembers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetAllowWebPosting() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowWebPosting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetArchiveOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetArchiveOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetCustomFooterText() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomFooterText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetCustomReplyTo() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomReplyTo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetDefaultMessageDenyNotificationText() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultMessageDenyNotificationText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetEnableCollaborativeInbox() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCollaborativeInbox",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetIncludeCustomFooter() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeCustomFooter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetIncludeInGlobalAddressList() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeInGlobalAddressList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetIsArchived() {
	_jsii_.InvokeVoid(
		g,
		"resetIsArchived",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetMembersCanPostAsTheGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetMembersCanPostAsTheGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetMessageModerationLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetMessageModerationLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetPrimaryLanguage() {
	_jsii_.InvokeVoid(
		g,
		"resetPrimaryLanguage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetReplyTo() {
	_jsii_.InvokeVoid(
		g,
		"resetReplyTo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetSendMessageDenyNotification() {
	_jsii_.InvokeVoid(
		g,
		"resetSendMessageDenyNotification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetSpamModerationLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetSpamModerationLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanAssistContent() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanAssistContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanContactOwner() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanContactOwner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanDiscoverGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanDiscoverGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanJoin() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanJoin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanLeaveGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanLeaveGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanModerateContent() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanModerateContent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanModerateMembers() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanModerateMembers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanPostMessage() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanPostMessage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanViewGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanViewGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) ResetWhoCanViewMembership() {
	_jsii_.InvokeVoid(
		g,
		"resetWhoCanViewMembership",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GroupSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GroupSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

