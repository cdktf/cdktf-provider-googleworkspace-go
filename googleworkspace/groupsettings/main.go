// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupsettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettings",
		reflect.TypeOf((*GroupSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowExternalMembers", GoGetter: "AllowExternalMembers"},
			_jsii_.MemberProperty{JsiiProperty: "allowExternalMembersInput", GoGetter: "AllowExternalMembersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowWebPosting", GoGetter: "AllowWebPosting"},
			_jsii_.MemberProperty{JsiiProperty: "allowWebPostingInput", GoGetter: "AllowWebPostingInput"},
			_jsii_.MemberProperty{JsiiProperty: "archiveOnly", GoGetter: "ArchiveOnly"},
			_jsii_.MemberProperty{JsiiProperty: "archiveOnlyInput", GoGetter: "ArchiveOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customFooterText", GoGetter: "CustomFooterText"},
			_jsii_.MemberProperty{JsiiProperty: "customFooterTextInput", GoGetter: "CustomFooterTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "customReplyTo", GoGetter: "CustomReplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "customReplyToInput", GoGetter: "CustomReplyToInput"},
			_jsii_.MemberProperty{JsiiProperty: "customRolesEnabledForSettingsToBeMerged", GoGetter: "CustomRolesEnabledForSettingsToBeMerged"},
			_jsii_.MemberProperty{JsiiProperty: "defaultMessageDenyNotificationText", GoGetter: "DefaultMessageDenyNotificationText"},
			_jsii_.MemberProperty{JsiiProperty: "defaultMessageDenyNotificationTextInput", GoGetter: "DefaultMessageDenyNotificationTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableCollaborativeInbox", GoGetter: "EnableCollaborativeInbox"},
			_jsii_.MemberProperty{JsiiProperty: "enableCollaborativeInboxInput", GoGetter: "EnableCollaborativeInboxInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "includeCustomFooter", GoGetter: "IncludeCustomFooter"},
			_jsii_.MemberProperty{JsiiProperty: "includeCustomFooterInput", GoGetter: "IncludeCustomFooterInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeInGlobalAddressList", GoGetter: "IncludeInGlobalAddressList"},
			_jsii_.MemberProperty{JsiiProperty: "includeInGlobalAddressListInput", GoGetter: "IncludeInGlobalAddressListInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isArchived", GoGetter: "IsArchived"},
			_jsii_.MemberProperty{JsiiProperty: "isArchivedInput", GoGetter: "IsArchivedInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanPostAsTheGroup", GoGetter: "MembersCanPostAsTheGroup"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanPostAsTheGroupInput", GoGetter: "MembersCanPostAsTheGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "messageModerationLevel", GoGetter: "MessageModerationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "messageModerationLevelInput", GoGetter: "MessageModerationLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLanguage", GoGetter: "PrimaryLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLanguageInput", GoGetter: "PrimaryLanguageInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "replyTo", GoGetter: "ReplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "replyToInput", GoGetter: "ReplyToInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowExternalMembers", GoMethod: "ResetAllowExternalMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowWebPosting", GoMethod: "ResetAllowWebPosting"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveOnly", GoMethod: "ResetArchiveOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomFooterText", GoMethod: "ResetCustomFooterText"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomReplyTo", GoMethod: "ResetCustomReplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultMessageDenyNotificationText", GoMethod: "ResetDefaultMessageDenyNotificationText"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableCollaborativeInbox", GoMethod: "ResetEnableCollaborativeInbox"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCustomFooter", GoMethod: "ResetIncludeCustomFooter"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeInGlobalAddressList", GoMethod: "ResetIncludeInGlobalAddressList"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsArchived", GoMethod: "ResetIsArchived"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanPostAsTheGroup", GoMethod: "ResetMembersCanPostAsTheGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageModerationLevel", GoMethod: "ResetMessageModerationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryLanguage", GoMethod: "ResetPrimaryLanguage"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplyTo", GoMethod: "ResetReplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendMessageDenyNotification", GoMethod: "ResetSendMessageDenyNotification"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpamModerationLevel", GoMethod: "ResetSpamModerationLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanAssistContent", GoMethod: "ResetWhoCanAssistContent"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanContactOwner", GoMethod: "ResetWhoCanContactOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanDiscoverGroup", GoMethod: "ResetWhoCanDiscoverGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanJoin", GoMethod: "ResetWhoCanJoin"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanLeaveGroup", GoMethod: "ResetWhoCanLeaveGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanModerateContent", GoMethod: "ResetWhoCanModerateContent"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanModerateMembers", GoMethod: "ResetWhoCanModerateMembers"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanPostMessage", GoMethod: "ResetWhoCanPostMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanViewGroup", GoMethod: "ResetWhoCanViewGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhoCanViewMembership", GoMethod: "ResetWhoCanViewMembership"},
			_jsii_.MemberProperty{JsiiProperty: "sendMessageDenyNotification", GoGetter: "SendMessageDenyNotification"},
			_jsii_.MemberProperty{JsiiProperty: "sendMessageDenyNotificationInput", GoGetter: "SendMessageDenyNotificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "spamModerationLevel", GoGetter: "SpamModerationLevel"},
			_jsii_.MemberProperty{JsiiProperty: "spamModerationLevelInput", GoGetter: "SpamModerationLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanAssistContent", GoGetter: "WhoCanAssistContent"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanAssistContentInput", GoGetter: "WhoCanAssistContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanContactOwner", GoGetter: "WhoCanContactOwner"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanContactOwnerInput", GoGetter: "WhoCanContactOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanDiscoverGroup", GoGetter: "WhoCanDiscoverGroup"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanDiscoverGroupInput", GoGetter: "WhoCanDiscoverGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanJoin", GoGetter: "WhoCanJoin"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanJoinInput", GoGetter: "WhoCanJoinInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanLeaveGroup", GoGetter: "WhoCanLeaveGroup"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanLeaveGroupInput", GoGetter: "WhoCanLeaveGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanModerateContent", GoGetter: "WhoCanModerateContent"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanModerateContentInput", GoGetter: "WhoCanModerateContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanModerateMembers", GoGetter: "WhoCanModerateMembers"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanModerateMembersInput", GoGetter: "WhoCanModerateMembersInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanPostMessage", GoGetter: "WhoCanPostMessage"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanPostMessageInput", GoGetter: "WhoCanPostMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanViewGroup", GoGetter: "WhoCanViewGroup"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanViewGroupInput", GoGetter: "WhoCanViewGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanViewMembership", GoGetter: "WhoCanViewMembership"},
			_jsii_.MemberProperty{JsiiProperty: "whoCanViewMembershipInput", GoGetter: "WhoCanViewMembershipInput"},
		},
		func() interface{} {
			j := jsiiProxy_GroupSettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettingsConfig",
		reflect.TypeOf((*GroupSettingsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettingsTimeouts",
		reflect.TypeOf((*GroupSettingsTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-googleworkspace.groupSettings.GroupSettingsTimeoutsOutputReference",
		reflect.TypeOf((*GroupSettingsTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_GroupSettingsTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
