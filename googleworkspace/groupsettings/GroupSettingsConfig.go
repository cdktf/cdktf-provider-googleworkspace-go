// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groupsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GroupSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The group's email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#email GroupSettings#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Defaults to `false`.
	//
	// Identifies whether members external to your organization can join the group. If true, Google Workspace users external to your organization can become members of this group. If false, users not belonging to the organization are not allowed to become members of this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#allow_external_members GroupSettings#allow_external_members}
	AllowExternalMembers interface{} `field:"optional" json:"allowExternalMembers" yaml:"allowExternalMembers"`
	// Defaults to `true`.
	//
	// Allows posting from web. If true, allows any member to post to the group forum. If false, Members only use Gmail to communicate with the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#allow_web_posting GroupSettings#allow_web_posting}
	AllowWebPosting interface{} `field:"optional" json:"allowWebPosting" yaml:"allowWebPosting"`
	// Defaults to `false`.
	//
	// Allows the group to be archived only. If true, Group is archived and the group is inactive. New messages to this group are rejected. The older archived messages are browsable and searchable. If true, the `who_can_post_message` property is set to `NONE_CAN_POST`. If reverted from true to false, `who_can_post_message` is set to `ALL_MANAGERS_CAN_POST`. If false, The group is active and can receive messages. When false, updating `who_can_post_message` to `NONE_CAN_POST`, results in an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#archive_only GroupSettings#archive_only}
	ArchiveOnly interface{} `field:"optional" json:"archiveOnly" yaml:"archiveOnly"`
	// Set the content of custom footer text. The maximum number of characters is 1,000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#custom_footer_text GroupSettings#custom_footer_text}
	CustomFooterText *string `field:"optional" json:"customFooterText" yaml:"customFooterText"`
	// An email address used when replying to a message if the `reply_to` property is set to `REPLY_TO_CUSTOM`.
	//
	// This address is defined by an account administrator. When the group's `reply_to` property is set to `REPLY_TO_CUSTOM`, the `custom_reply_to` property holds a custom email address used when replying to a message, the `custom_reply_to` property must have a text value or an error is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#custom_reply_to GroupSettings#custom_reply_to}
	CustomReplyTo *string `field:"optional" json:"customReplyTo" yaml:"customReplyTo"`
	// When a message is rejected, this is text for the rejection notification sent to the message's author.
	//
	// By default, this property is empty and has no value in the API's response body. The maximum notification text size is 10,000 characters. Requires `send_message_deny_notification` property to be true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#default_message_deny_notification_text GroupSettings#default_message_deny_notification_text}
	DefaultMessageDenyNotificationText *string `field:"optional" json:"defaultMessageDenyNotificationText" yaml:"defaultMessageDenyNotificationText"`
	// Defaults to `false`. Specifies whether a collaborative inbox will remain turned on for the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#enable_collaborative_inbox GroupSettings#enable_collaborative_inbox}
	EnableCollaborativeInbox interface{} `field:"optional" json:"enableCollaborativeInbox" yaml:"enableCollaborativeInbox"`
	// Defaults to `false`. Whether to include custom footer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#include_custom_footer GroupSettings#include_custom_footer}
	IncludeCustomFooter interface{} `field:"optional" json:"includeCustomFooter" yaml:"includeCustomFooter"`
	// Defaults to `true`.
	//
	// Enables the group to be included in the Global Address List. If true, the group is included in the Global Address List. If false, it is not included in the Global Address List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#include_in_global_address_list GroupSettings#include_in_global_address_list}
	IncludeInGlobalAddressList interface{} `field:"optional" json:"includeInGlobalAddressList" yaml:"includeInGlobalAddressList"`
	// Defaults to `false`.
	//
	// Allows the Group contents to be archived. If true, archive messages sent to the group. If false, Do not keep an archive of messages sent to this group. If false, previously archived messages remain in the archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#is_archived GroupSettings#is_archived}
	IsArchived interface{} `field:"optional" json:"isArchived" yaml:"isArchived"`
	// Defaults to `false`.
	//
	// Enables members to post messages as the group. If true, group member can post messages using the group's email address instead of their own email address. Message appear to originate from the group itself. Any message moderation settings on individual users or new members do not apply to posts made on behalf of the group. If false, members can not post in behalf of the group's email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#members_can_post_as_the_group GroupSettings#members_can_post_as_the_group}
	MembersCanPostAsTheGroup interface{} `field:"optional" json:"membersCanPostAsTheGroup" yaml:"membersCanPostAsTheGroup"`
	// Defaults to `MODERATE_NONE`.
	//
	// Moderation level of incoming messages. Possible values are:
	// 	- `MODERATE_ALL_MESSAGES`: All messages are sent to the group owner's email address for approval. If approved, the message is sent to the group.
	// 	- `MODERATE_NON_MEMBERS`: All messages from non group members are sent to the group owner's email address for approval. If approved, the message is sent to the group.
	// 	- `MODERATE_NEW_MEMBERS`: All messages from new members are sent to the group owner's email address for approval. If approved, the message is sent to the group.
	// 	- `MODERATE_NONE`: No moderator approval is required. Messages are delivered directly to the group.
	// 	Note: When the `who_can_post_message` is set to `ANYONE_CAN_POST`, we recommend the `message_moderation_level` be set to `MODERATE_NON_MEMBERS` to protect the group from possible spam.When `member_can_post_as_the_group` is true, any message moderation settings on individual users or new members will not apply to posts made on behalf of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#message_moderation_level GroupSettings#message_moderation_level}
	MessageModerationLevel *string `field:"optional" json:"messageModerationLevel" yaml:"messageModerationLevel"`
	// The primary language for group.
	//
	// For a group's primary language use the language tags from the Google Workspace languages found at Google Workspace Email Settings API Email Language Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#primary_language GroupSettings#primary_language}
	PrimaryLanguage *string `field:"optional" json:"primaryLanguage" yaml:"primaryLanguage"`
	// Defaults to `REPLY_TO_IGNORE`.
	//
	// Specifies who receives the default reply. Possible values are:
	// 	- `REPLY_TO_CUSTOM`: For replies to messages, use the group's custom email address. When set to `REPLY_TO_CUSTOM`, the `custom_reply_to` property holds the custom email address used when replying to a message, the customReplyTo property must have a value. Otherwise an error is returned.
	// 	- `REPLY_TO_SENDER`: The reply sent to author of message.
	// 	- `REPLY_TO_LIST`: This reply message is sent to the group.
	// 	- `REPLY_TO_OWNER`: The reply is sent to the owner(s) of the group. This does not include the group's managers.
	// 	- `REPLY_TO_IGNORE`: Group users individually decide where the message reply is sent.
	// 	- `REPLY_TO_MANAGERS`: This reply message is sent to the group's managers, which includes all managers and the group owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#reply_to GroupSettings#reply_to}
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// Defaults to `false`.
	//
	// Allows a member to be notified if the member's message to the group is denied by the group owner. If true, when a message is rejected, send the deny message notification to the message author. The `default_message_deny_notification_text` property is dependent on the `send_message_deny_notification` property being true. If false, when a message is rejected, no notification is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#send_message_deny_notification GroupSettings#send_message_deny_notification}
	SendMessageDenyNotification interface{} `field:"optional" json:"sendMessageDenyNotification" yaml:"sendMessageDenyNotification"`
	// Defaults to `MODERATE`.
	//
	// Specifies moderation levels for messages detected as spam. Possible values are:
	// 	- `ALLOW`: Post the message to the group.
	// 	- `MODERATE`: Send the message to the moderation queue. This is the default.
	// 	- `SILENTLY_MODERATE`: Send the message to the moderation queue, but do not send notification to moderators.
	// 	- `REJECT`: Immediately reject the message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#spam_moderation_level GroupSettings#spam_moderation_level}
	SpamModerationLevel *string `field:"optional" json:"spamModerationLevel" yaml:"spamModerationLevel"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#timeouts GroupSettings#timeouts}
	Timeouts *GroupSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Defaults to `NONE`. Specifies who can moderate metadata. Possible values are:  	- `ALL_MEMBERS` 	- `OWNERS_AND_MANAGERS` 	- `MANAGERS_ONLY` 	- `OWNERS_ONLY` 	- `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_assist_content GroupSettings#who_can_assist_content}
	WhoCanAssistContent *string `field:"optional" json:"whoCanAssistContent" yaml:"whoCanAssistContent"`
	// Defaults to `ANYONE_CAN_CONTACT`.
	//
	// Permission to contact owner of the group via web UI. Possible values are:
	// 	- `ALL_IN_DOMAIN_CAN_CONTACT`
	// 	- `ALL_MANAGERS_CAN_CONTACT`
	// 	- `ALL_MEMBERS_CAN_CONTACT`
	// 	- `ANYONE_CAN_CONTACT`
	// 	- `ALL_OWNERS_CAN_CONTACT`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_contact_owner GroupSettings#who_can_contact_owner}
	WhoCanContactOwner *string `field:"optional" json:"whoCanContactOwner" yaml:"whoCanContactOwner"`
	// Defaults to `ALL_IN_DOMAIN_CAN_DISCOVER`.
	//
	// Specifies the set of users for whom this group is discoverable. Possible values are:
	// 	- `ANYONE_CAN_DISCOVER`
	// 	- `ALL_IN_DOMAIN_CAN_DISCOVER`
	// 	- `ALL_MEMBERS_CAN_DISCOVER`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_discover_group GroupSettings#who_can_discover_group}
	WhoCanDiscoverGroup *string `field:"optional" json:"whoCanDiscoverGroup" yaml:"whoCanDiscoverGroup"`
	// Defaults to `CAN_REQUEST_TO_JOIN`.
	//
	// Permission to join group. Possible values are:
	// 	- `ANYONE_CAN_JOIN`: Any Internet user, both inside and outside your domain, can join the group.
	// 	- `ALL_IN_DOMAIN_CAN_JOIN`: Anyone in the account domain can join. This includes accounts with multiple domains.
	// 	- `INVITED_CAN_JOIN`: Candidates for membership can be invited to join.
	// 	- `CAN_REQUEST_TO_JOIN`: Non members can request an invitation to join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_join GroupSettings#who_can_join}
	WhoCanJoin *string `field:"optional" json:"whoCanJoin" yaml:"whoCanJoin"`
	// Defaults to `ALL_MEMBERS_CAN_LEAVE`. Permission to leave the group. Possible values are: 	- `ALL_MANAGERS_CAN_LEAVE` 	- `ALL_MEMBERS_CAN_LEAVE` 	- `NONE_CAN_LEAVE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_leave_group GroupSettings#who_can_leave_group}
	WhoCanLeaveGroup *string `field:"optional" json:"whoCanLeaveGroup" yaml:"whoCanLeaveGroup"`
	// Defaults to `OWNERS_AND_MANAGERS`. Specifies who can moderate content. Possible values are:  	- `ALL_MEMBERS` 	- `OWNERS_AND_MANAGERS` 	- `OWNERS_ONLY` 	- `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_moderate_content GroupSettings#who_can_moderate_content}
	WhoCanModerateContent *string `field:"optional" json:"whoCanModerateContent" yaml:"whoCanModerateContent"`
	// Defaults to `OWNERS_AND_MANAGERS`. Specifies who can manage members. Possible values are:  	- `ALL_MEMBERS` 	- `OWNERS_AND_MANAGERS` 	- `OWNERS_ONLY` 	- `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_moderate_members GroupSettings#who_can_moderate_members}
	WhoCanModerateMembers *string `field:"optional" json:"whoCanModerateMembers" yaml:"whoCanModerateMembers"`
	// Permissions to post messages.
	//
	// Possible values are:
	// 	- `NONE_CAN_POST`: The group is disabled and archived. No one can post a message to this group. * When archiveOnly is false, updating whoCanPostMessage to NONE_CAN_POST, results in an error. * If archiveOnly is reverted from true to false, whoCanPostMessages is set to ALL_MANAGERS_CAN_POST.
	// 	- `ALL_MANAGERS_CAN_POST`: Managers, including group owners, can post messages.
	// 	- `ALL_MEMBERS_CAN_POST`: Any group member can post a message.
	// 	- `ALL_OWNERS_CAN_POST`: Only group owners can post a message.
	// 	- `ALL_IN_DOMAIN_CAN_POST`: Anyone in the account can post a message.
	// 	- `ANYONE_CAN_POST`: Any Internet user who outside your account can access your Google Groups service and post a message.
	// 	*Note: When `who_can_post_message` is set to `ANYONE_CAN_POST`, we recommend the`message_moderation_level` be set to `MODERATE_NON_MEMBERS` to protect the group from possible spam. Users not belonging to the organization are not allowed to become members of this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_post_message GroupSettings#who_can_post_message}
	WhoCanPostMessage *string `field:"optional" json:"whoCanPostMessage" yaml:"whoCanPostMessage"`
	// Defaults to `ALL_MEMBERS_CAN_VIEW`.
	//
	// Permissions to view group messages. Possible values are:
	// 	- `ANYONE_CAN_VIEW`: Any Internet user can view the group's messages.
	// 	- `ALL_IN_DOMAIN_CAN_VIEW`: Anyone in your account can view this group's messages.
	// 	- `ALL_MEMBERS_CAN_VIEW`: All group members can view the group's messages.
	// 	- `ALL_MANAGERS_CAN_VIEW`: Any group manager can view this group's messages.
	// 	- `ALL_OWNERS_CAN_VIEW`: The group owners can view this group's messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_view_group GroupSettings#who_can_view_group}
	WhoCanViewGroup *string `field:"optional" json:"whoCanViewGroup" yaml:"whoCanViewGroup"`
	// Defaults to `ALL_MEMBERS_CAN_VIEW`.
	//
	// Permissions to view membership. Possible values are:
	// 	- `ALL_IN_DOMAIN_CAN_VIEW`: Anyone in the account can view the group members list. If a group already has external members, those members can still send email to this group.
	// 	- `ALL_MEMBERS_CAN_VIEW`: The group members can view the group members list.
	// 	- `ALL_MANAGERS_CAN_VIEW`: The group managers can view group members list.
	// 	- `ALL_OWNERS_CAN_VIEW`: The group owners can view group members list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/googleworkspace/0.7.0/docs/resources/group_settings#who_can_view_membership GroupSettings#who_can_view_membership}
	WhoCanViewMembership *string `field:"optional" json:"whoCanViewMembership" yaml:"whoCanViewMembership"`
}

