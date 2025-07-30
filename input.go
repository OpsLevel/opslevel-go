// Code generated; DO NOT EDIT.
package opslevel

import "github.com/relvacode/iso8601"

// AlertSourceExternalIdentifier Specifies the input needed to find an alert source with external information
type AlertSourceExternalIdentifier struct {
	ExternalId string              `json:"externalId" yaml:"externalId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The external id of the alert (Required)
	Type       AlertSourceTypeEnum `json:"type" yaml:"type" example:"custom"`                                      // The type of the alert (Required)
}

// AlertSourceInput Input fields for the mutations to manage Alert Sources
type AlertSourceInput struct {
	Description *Nullable[string]               `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the alert source (Optional)
	Identifier  ExternalResourceIdentifierInput `json:"identifier" yaml:"identifier"`                                               // The alert source identifier (Required)
	Name        *Nullable[string]               `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The name of the alert source (Optional)
	Url         *Nullable[string]               `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"`                 // The url of the alert source (Optional)
}

// AlertSourceServiceCreateInput Specifies the input used for attaching an alert source to a service
type AlertSourceServiceCreateInput struct {
	AlertSourceExternalIdentifier *AlertSourceExternalIdentifier `json:"alertSourceExternalIdentifier,omitempty" yaml:"alertSourceExternalIdentifier,omitempty"`           // Specifies the input needed to find an alert source with external information (Optional)
	AlertSourceId                 *Nullable[ID]                  `json:"alertSourceId,omitempty" yaml:"alertSourceId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // Specifies the input needed to find an alert source with external information (Optional)
	Service                       IdentifierInput                `json:"service" yaml:"service"`                                                                           // The service that the alert source will be attached to (Required)
}

// AlertSourceServiceDeleteInput Specifies the input fields used in the `alertSourceServiceDelete` mutation
type AlertSourceServiceDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the alert source service to be deleted (Required)
}

// AlertSourceStatusUpdateInput Specifies the input fields used in the `alertSourceStatusUpdate` mutation
type AlertSourceStatusUpdateInput struct {
	AlertSource ExternalResourceIdentifierInput `json:"alertSource" yaml:"alertSource"`       // The alert source to be updated (Required)
	Status      AlertSourceStatusTypeEnum       `json:"status" yaml:"status" example:"alert"` // The new status of the alert source (Required)
}

// AliasCreateInput The input for the `aliasCreate` mutation
type AliasCreateInput struct {
	Alias   string `json:"alias" yaml:"alias" example:"example_value"`     // The alias you wish to create (Required)
	OwnerId ID     `json:"ownerId" yaml:"ownerId" example:"example_value"` // The ID of the resource you want to create the alias on. Services, teams, groups, systems, and domains are supported (Required)
}

// AliasDeleteInput The input for the `aliasDelete` mutation
type AliasDeleteInput struct {
	Alias     string             `json:"alias" yaml:"alias" example:"example_value"`  // The alias you wish to delete (Required)
	OwnerType AliasOwnerTypeEnum `json:"ownerType" yaml:"ownerType" example:"domain"` // The resource the alias you wish to delete belongs to (Required)
}

// ApprovalConfigInput Config for approval
type ApprovalConfigInput struct {
	ApprovalRequired *bool                  `json:"approvalRequired,omitempty" yaml:"approvalRequired,omitempty" example:"false"` // Flag indicating approval is required (Optional)
	Teams            *[]IdentifierInput     `json:"teams,omitempty" yaml:"teams,omitempty" example:"[]"`                          // Teams that can approve (Optional)
	Users            *[]UserIdentifierInput `json:"users,omitempty" yaml:"users,omitempty" example:"[]"`                          // Users that can approve (Optional)
}

// AwsIntegrationInput Specifies the input fields used to create and update an AWS integration
type AwsIntegrationInput struct {
	AwsTagsOverrideOwnership *Nullable[bool]     `json:"awsTagsOverrideOwnership,omitempty" yaml:"awsTagsOverrideOwnership,omitempty" example:"false"`    // Allow tags imported from AWS to override ownership set in OpsLevel directly (Optional)
	ExternalId               *Nullable[string]   `json:"externalId,omitempty" yaml:"externalId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`      // The External ID defined in the trust relationship to ensure OpsLevel is the only third party assuming this role (See https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html for more details) (Optional)
	IamRole                  *Nullable[string]   `json:"iamRole,omitempty" yaml:"iamRole,omitempty" example:"example_value"`                              // The IAM role OpsLevel uses in order to access the AWS account (Optional)
	Name                     *Nullable[string]   `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                    // The name of the integration (Optional)
	OwnershipTagKeys         *Nullable[[]string] `json:"ownershipTagKeys,omitempty" yaml:"ownershipTagKeys,omitempty" example:"['tag_key1', 'tag_key2']"` // An array of tag keys used to associate ownership from an integration. Max 5 (Optional)
	RegionOverride           *Nullable[[]string] `json:"regionOverride,omitempty" yaml:"regionOverride,omitempty" example:"['us-east-1', 'eu-west-1']"`   // Overrides the AWS region(s) that will be synchronized by this integration (Optional)
}

// AzureResourcesIntegrationInput Specifies the input fields used to create and update an Azure resources integration
type AzureResourcesIntegrationInput struct {
	ClientId              *Nullable[string]   `json:"clientId,omitempty" yaml:"clientId,omitempty" example:"example_value"`                            // The client OpsLevel uses to access the Azure account (Optional)
	ClientSecret          *Nullable[string]   `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty" example:"example_value"`                    // The client secret OpsLevel uses to access the Azure account (Optional)
	Name                  *Nullable[string]   `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                    // The name of the integration (Optional)
	OwnershipTagKeys      *Nullable[[]string] `json:"ownershipTagKeys,omitempty" yaml:"ownershipTagKeys,omitempty" example:"['tag_key1', 'tag_key2']"` // An array of tag keys used to associate ownership from an integration. Max 5 (Optional)
	SubscriptionId        *Nullable[string]   `json:"subscriptionId,omitempty" yaml:"subscriptionId,omitempty" example:"example_value"`                // The subscription OpsLevel uses to access the Azure account (Optional)
	TagsOverrideOwnership *Nullable[bool]     `json:"tagsOverrideOwnership,omitempty" yaml:"tagsOverrideOwnership,omitempty" example:"false"`          // Allow tags imported from Azure to override ownership set in OpsLevel directly (Optional)
	TenantId              *Nullable[string]   `json:"tenantId,omitempty" yaml:"tenantId,omitempty" example:"example_value"`                            // The tenant OpsLevel uses to access the Azure account (Optional)
}

// CategoryCreateInput Specifies the input fields used to create a category
type CategoryCreateInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the category (Optional)
	Name        string            `json:"name" yaml:"name" example:"example_value"`                                   // The display name of the category (Required)
}

// CategoryDeleteInput Specifies the input fields used to delete a category
type CategoryDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category to be deleted (Required)
}

// CategoryUpdateInput Specifies the input fields used to update a category
type CategoryUpdateInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the category (Optional)
	Id          ID                `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                     // The id of the category to be updated (Required)
	Name        *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The display name of the category (Optional)
}

// CheckAlertSourceUsageCreateInput Specifies the input fields used to create an alert source usage check
type CheckAlertSourceUsageCreateInput struct {
	AlertSourceNamePredicate *PredicateInput         `json:"alertSourceNamePredicate,omitempty" yaml:"alertSourceNamePredicate,omitempty"`           // The condition that the alert source name should satisfy to be evaluated (Optional)
	AlertSourceType          *AlertSourceTypeEnum    `json:"alertSourceType,omitempty" yaml:"alertSourceType,omitempty" example:"custom"`            // The type of the alert source (Optional)
	CategoryId               ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn                 *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled                  *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId                 *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId                  ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                     string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                    *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId                  *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckAlertSourceUsageUpdateInput Specifies the input fields used to update an alert source usage check
type CheckAlertSourceUsageUpdateInput struct {
	AlertSourceNamePredicate *PredicateUpdateInput   `json:"alertSourceNamePredicate,omitempty" yaml:"alertSourceNamePredicate,omitempty"`               // The condition that the alert source name should satisfy to be evaluated (Optional)
	AlertSourceType          *AlertSourceTypeEnum    `json:"alertSourceType,omitempty" yaml:"alertSourceType,omitempty" example:"custom"`                // The type of the alert source (Optional)
	CategoryId               *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn                 *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled                  *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId                 *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                       ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId                  *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                     *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                    *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId                  *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckCodeIssueCreateInput Specifies the input fields used to create a code issue check
type CheckCodeIssueCreateInput struct {
	CategoryId     ID                            `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	Constraint     CheckCodeIssueConstraintEnum  `json:"constraint" yaml:"constraint" example:"any"`                                             // The type of constraint used in evaluation the code issues check (Required)
	EnableOn       *Nullable[iso8601.Time]       `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled        *Nullable[bool]               `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId       *Nullable[ID]                 `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	IssueName      *Nullable[string]             `json:"issueName,omitempty" yaml:"issueName,omitempty" example:"example_value"`                 // The issue name used for code issue lookup (Optional)
	IssueType      *Nullable[[]string]           `json:"issueType,omitempty" yaml:"issueType,omitempty" example:"['bug', 'error']"`              // The type of code issue to consider (Optional)
	LevelId        ID                            `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	MaxAllowed     *int                          `json:"maxAllowed,omitempty" yaml:"maxAllowed,omitempty" example:"3"`                           // The threshold count of code issues beyond which the check starts failing (Optional)
	Name           string                        `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes          *string                       `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId        *Nullable[ID]                 `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	ResolutionTime *CodeIssueResolutionTimeInput `json:"resolutionTime,omitempty" yaml:"resolutionTime,omitempty"`                               // The resolution time recommended by the reporting source of the code issue (Optional)
	Severity       *Nullable[[]string]           `json:"severity,omitempty" yaml:"severity,omitempty" example:"['sev1', 'sev2']"`                // The severity levels of the issue (Optional)
}

// CheckCodeIssueUpdateInput Specifies the input fields used to update an exasting code issue check
type CheckCodeIssueUpdateInput struct {
	CategoryId     *Nullable[ID]                 `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	Constraint     CheckCodeIssueConstraintEnum  `json:"constraint" yaml:"constraint" example:"any"`                                                 // The type of constraint used in evaluation the code issues check (Required)
	EnableOn       *Nullable[iso8601.Time]       `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled        *Nullable[bool]               `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId       *Nullable[ID]                 `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id             ID                            `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	IssueName      *Nullable[string]             `json:"issueName,omitempty" yaml:"issueName,omitempty" example:"example_value"`                     // The issue name used for code issue lookup (Optional)
	IssueType      *Nullable[[]string]           `json:"issueType,omitempty" yaml:"issueType,omitempty" example:"['bug', 'error']"`                  // The type of code issue to consider (Optional)
	LevelId        *Nullable[ID]                 `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	MaxAllowed     *int                          `json:"maxAllowed,omitempty" yaml:"maxAllowed,omitempty" example:"3"`                               // The threshold count of code issues beyond which the check starts failing (Optional)
	Name           *Nullable[string]             `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes          *string                       `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId        *Nullable[ID]                 `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	ResolutionTime *CodeIssueResolutionTimeInput `json:"resolutionTime,omitempty" yaml:"resolutionTime,omitempty"`                                   // The resolution time recommended by the reporting source of the code issue (Optional)
	Severity       *Nullable[[]string]           `json:"severity,omitempty" yaml:"severity,omitempty" example:"['sev1', 'sev2']"`                    // The severity levels of the issue (Optional)
}

// CheckCopyInput Information about the check(s) that are to be copied
type CheckCopyInput struct {
	CheckIds         []ID            `json:"checkIds" yaml:"checkIds" example:"['Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk', 'Z2lkOi8vc2VydmljZS85ODc2NTQzMjE']"` // The IDs of the checks to be copied (Required)
	Move             *Nullable[bool] `json:"move,omitempty" yaml:"move,omitempty" example:"false"`                                                      // If set to true, the original checks will be deleted after being successfully copied (Optional)
	TargetCategoryId ID              `json:"targetCategoryId" yaml:"targetCategoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                        // The ID of the category to which the checks are copied. Belongs to either the rubric or a scorecard (Required)
	TargetLevelId    *Nullable[ID]   `json:"targetLevelId,omitempty" yaml:"targetLevelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`          // The ID of the level which the copied checks are associated with (Optional)
}

// CheckCustomEventCreateInput Creates a custom event check
type CheckCustomEventCreateInput struct {
	CategoryId       ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn         *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled          *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId         *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	IntegrationId    ID                      `json:"integrationId" yaml:"integrationId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The integration id this check will use (Required)
	LevelId          ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name             string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes            *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId          *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	PassPending      *Nullable[bool]         `json:"passPending,omitempty" yaml:"passPending,omitempty" example:"false"`                     // True if this check should pass by default. Otherwise the default 'pending' state counts as a failure (Optional)
	ResultMessage    *Nullable[string]       `json:"resultMessage,omitempty" yaml:"resultMessage,omitempty" example:"example_value"`         // The check result message template. It is compiled with Liquid and formatted in Markdown. [More info about liquid templates](https://docs.opslevel.com/docs/checks/payload-checks/#liquid-templating) (Optional)
	ServiceSelector  string                  `json:"serviceSelector" yaml:"serviceSelector" example:"example_value"`                         // A jq expression that will be ran against your payload. This will parse out the service identifier. [More info about jq](https://jqplay.org/) (Required)
	SuccessCondition string                  `json:"successCondition" yaml:"successCondition" example:"example_value"`                       // A jq expression that will be ran against your payload. A truthy value will result in the check passing. [More info about jq](https://jqplay.org/) (Required)
}

// CheckCustomEventUpdateInput Specifies the input fields used to update a custom event check
type CheckCustomEventUpdateInput struct {
	CategoryId       *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the category the check belongs to (Optional)
	EnableOn         *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`                  // The date when the check will be automatically enabled (Optional)
	Enabled          *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                       // Whether the check is enabled or not (Optional)
	FilterId         *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The id of the filter the check belongs to (Optional)
	Id               ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                           // The id of the check to be updated (Required)
	IntegrationId    *Nullable[ID]           `json:"integrationId,omitempty" yaml:"integrationId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The integration id this check will use (Optional)
	LevelId          *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The id of the level the check belongs to (Optional)
	Name             *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                     // The display name of the check (Optional)
	Notes            *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                                   // Additional information about the check (Optional)
	OwnerId          *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The id of the owner of the check (Optional)
	PassPending      *Nullable[bool]         `json:"passPending,omitempty" yaml:"passPending,omitempty" example:"false"`                               // True if this check should pass by default. Otherwise the default 'pending' state counts as a failure (Optional)
	ResultMessage    *Nullable[string]       `json:"resultMessage,omitempty" yaml:"resultMessage,omitempty" example:"example_value"`                   // The check result message template. It is compiled with Liquid and formatted in Markdown. [More info about liquid templates](https://docs.opslevel.com/docs/checks/payload-checks/#liquid-templating) (Optional)
	ServiceSelector  *Nullable[string]       `json:"serviceSelector,omitempty" yaml:"serviceSelector,omitempty" example:"example_value"`               // A jq expression that will be ran against your payload. This will parse out the service identifier. [More info about jq](https://jqplay.org/) (Optional)
	SuccessCondition *Nullable[string]       `json:"successCondition,omitempty" yaml:"successCondition,omitempty" example:"example_value"`             // A jq expression that will be ran against your payload. A truthy value will result in the check passing. [More info about jq](https://jqplay.org/) (Optional)
}

// CheckDeleteInput Specifies the input fields used to delete a check
type CheckDeleteInput struct {
	Id *Nullable[ID] `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the check to be deleted (Optional)
}

// CheckGitBranchProtectionCreateInput Specifies the input fields used to create a branch protection check
type CheckGitBranchProtectionCreateInput struct {
	CategoryId ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckGitBranchProtectionUpdateInput Specifies the input fields used to update a branch protection check
type CheckGitBranchProtectionUpdateInput struct {
	CategoryId *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckHasDocumentationCreateInput Specifies the input fields used to create a documentation check
type CheckHasDocumentationCreateInput struct {
	CategoryId      ID                          `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	DocumentSubtype HasDocumentationSubtypeEnum `json:"documentSubtype" yaml:"documentSubtype" example:"openapi"`                               // The subtype of the document (Required)
	DocumentType    HasDocumentationTypeEnum    `json:"documentType" yaml:"documentType" example:"api"`                                         // The type of the document (Required)
	EnableOn        *Nullable[iso8601.Time]     `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled         *Nullable[bool]             `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId        *Nullable[ID]               `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId         ID                          `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name            string                      `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes           *string                     `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId         *Nullable[ID]               `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckHasDocumentationUpdateInput Specifies the input fields used to update a documentation check
type CheckHasDocumentationUpdateInput struct {
	CategoryId      *Nullable[ID]                `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	DocumentSubtype *HasDocumentationSubtypeEnum `json:"documentSubtype,omitempty" yaml:"documentSubtype,omitempty" example:"openapi"`               // The subtype of the document (Optional)
	DocumentType    *HasDocumentationTypeEnum    `json:"documentType,omitempty" yaml:"documentType,omitempty" example:"api"`                         // The type of the document (Optional)
	EnableOn        *Nullable[iso8601.Time]      `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled         *Nullable[bool]              `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId        *Nullable[ID]                `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id              ID                           `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId         *Nullable[ID]                `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name            *Nullable[string]            `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes           *string                      `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId         *Nullable[ID]                `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckHasRecentDeployCreateInput Specifies the input fields used to create a recent deploys check
type CheckHasRecentDeployCreateInput struct {
	CategoryId ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	Days       int                     `json:"days" yaml:"days" example:"3"`                                                           // The number of days to check since the last deploy (Required)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckHasRecentDeployUpdateInput Specifies the input fields used to update a has recent deploy check
type CheckHasRecentDeployUpdateInput struct {
	CategoryId *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	Days       *int                    `json:"days,omitempty" yaml:"days,omitempty" example:"3"`                                           // The number of days to check since the last deploy (Optional)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckManualCreateInput Specifies the input fields used to create a manual check
type CheckManualCreateInput struct {
	CategoryId            ID                         `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn              *Nullable[iso8601.Time]    `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]            `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId              *Nullable[ID]              `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId               ID                         `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                  string                     `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                 *string                    `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]              `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty" yaml:"updateFrequency,omitempty"`                             // Defines the minimum frequency of the updates (Optional)
	UpdateRequiresComment bool                       `json:"updateRequiresComment" yaml:"updateRequiresComment" example:"false"`                     // Whether the check requires a comment or not (Required)
}

// CheckManualUpdateInput Specifies the input fields used to update a manual check
type CheckManualUpdateInput struct {
	CategoryId            *Nullable[ID]                    `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn              *Nullable[iso8601.Time]          `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]                  `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId              *Nullable[ID]                    `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                    ID                               `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId               *Nullable[ID]                    `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                  *Nullable[string]                `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                 *string                          `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]                    `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	UpdateFrequency       *ManualCheckFrequencyUpdateInput `json:"updateFrequency,omitempty" yaml:"updateFrequency,omitempty"`                                 // Defines the minimum frequency of the updates (Optional)
	UpdateRequiresComment *Nullable[bool]                  `json:"updateRequiresComment,omitempty" yaml:"updateRequiresComment,omitempty" example:"false"`     // Whether the check requires a comment or not (Optional)
}

// CheckPackageVersionCreateInput Information about the package version check to be created
type CheckPackageVersionCreateInput struct {
	CategoryId                 ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn                   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled                    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId                   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId                    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	MissingPackageResult       *CheckResultStatusEnum  `json:"missingPackageResult,omitempty" yaml:"missingPackageResult,omitempty" example:"failed"`  // The check result if the package isn't being used by a service (Optional)
	Name                       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId                    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	PackageConstraint          PackageConstraintEnum   `json:"packageConstraint" yaml:"packageConstraint" example:"does_not_exist"`                    // The package constraint the service is to be checked for (Required)
	PackageManager             PackageManagerEnum      `json:"packageManager" yaml:"packageManager" example:"alpm"`                                    // The package manager (ecosystem) this package relates to (Required)
	PackageName                string                  `json:"packageName" yaml:"packageName" example:"example_value"`                                 // The name of the package to be checked (Required)
	PackageNameIsRegex         *Nullable[bool]         `json:"packageNameIsRegex,omitempty" yaml:"packageNameIsRegex,omitempty" example:"false"`       // Whether or not the value in the package name field is a regular expression (Optional)
	VersionConstraintPredicate *PredicateInput         `json:"versionConstraintPredicate,omitempty" yaml:"versionConstraintPredicate,omitempty"`       // The predicate that describes the version constraint the package must satisfy (Optional)
}

// CheckPackageVersionUpdateInput Information about the package version check to be updated
type CheckPackageVersionUpdateInput struct {
	CategoryId                 *Nullable[ID]                    `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn                   *Nullable[iso8601.Time]          `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled                    *Nullable[bool]                  `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId                   *Nullable[ID]                    `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                         ID                               `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId                    *Nullable[ID]                    `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	MissingPackageResult       *Nullable[CheckResultStatusEnum] `json:"missingPackageResult,omitempty" yaml:"missingPackageResult,omitempty" example:"failed"`      // The check result if the package isn't being used by a service (Optional)
	Name                       *Nullable[string]                `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                      *string                          `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId                    *Nullable[ID]                    `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	PackageConstraint          *Nullable[PackageConstraintEnum] `json:"packageConstraint,omitempty" yaml:"packageConstraint,omitempty" example:"does_not_exist"`    // The package constraint the service is to be checked for (Optional)
	PackageManager             *Nullable[PackageManagerEnum]    `json:"packageManager,omitempty" yaml:"packageManager,omitempty" example:"alpm"`                    // The package manager (ecosystem) this package relates to (Optional)
	PackageName                *Nullable[string]                `json:"packageName,omitempty" yaml:"packageName,omitempty" example:"example_value"`                 // The name of the package to be checked (Optional)
	PackageNameIsRegex         *Nullable[bool]                  `json:"packageNameIsRegex,omitempty" yaml:"packageNameIsRegex,omitempty" example:"false"`           // Whether or not the value in the package name field is a regular expression (Optional)
	VersionConstraintPredicate *Nullable[PredicateUpdateInput]  `json:"versionConstraintPredicate,omitempty" yaml:"versionConstraintPredicate,omitempty"`           // The predicate that describes the version constraint the package must satisfy (Optional)
}

// CheckRelationshipCreateInput Specifies the input fields used to create a relationships check
type CheckRelationshipCreateInput struct {
	CategoryId                 ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                             // The id of the category the check belongs to (Required)
	EnableOn                   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`                    // The date when the check will be automatically enabled (Optional)
	Enabled                    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                         // Whether the check is enabled or not (Optional Default: false)
	FilterId                   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The id of the filter of the check (Optional)
	LevelId                    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                   // The id of the level the check belongs to (Required)
	Name                       string                  `json:"name" yaml:"name" example:"example_value"`                                                           // The display name of the check (Required)
	Notes                      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                                     // Additional information about the check (Optional)
	OwnerId                    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`               // The id of the team that owns the check (Optional)
	RelationshipCountPredicate PredicateInput          `json:"relationshipCountPredicate" yaml:"relationshipCountPredicate"`                                       // The condition that should be satisfied by the number of RelatedTo relationships (Required)
	RelationshipDefinitionId   ID                      `json:"relationshipDefinitionId" yaml:"relationshipDefinitionId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // Count relationships of a specific relationship definition (Required)
}

// CheckRelationshipUpdateInput Specifies the input fields used to update a relationships check
type CheckRelationshipUpdateInput struct {
	CategoryId                 *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                             // The id of the category the check belongs to (Optional)
	EnableOn                   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`                                        // The date when the check will be automatically enabled (Optional)
	Enabled                    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                                             // Whether the check is enabled or not (Optional)
	FilterId                   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                 // The id of the filter the check belongs to (Optional)
	Id                         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                                                 // The id of the check to be updated (Required)
	LevelId                    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                   // The id of the level the check belongs to (Optional)
	Name                       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                                           // The display name of the check (Optional)
	Notes                      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                                                         // Additional information about the check (Optional)
	OwnerId                    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                   // The id of the owner of the check (Optional)
	RelationshipCountPredicate *PredicateInput         `json:"relationshipCountPredicate,omitempty" yaml:"relationshipCountPredicate,omitempty"`                                       // The condition that should be satisfied by the number of RelatedTo relationships (Optional)
	RelationshipDefinitionId   *Nullable[ID]           `json:"relationshipDefinitionId,omitempty" yaml:"relationshipDefinitionId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // Count relationships of a specific relationship definition (Optional)
}

// CheckRepositoryFileCreateInput Specifies the input fields used to create a repo file check
type CheckRepositoryFileCreateInput struct {
	CategoryId            ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	DirectorySearch       *Nullable[bool]         `json:"directorySearch,omitempty" yaml:"directorySearch,omitempty" example:"false"`             // Whether the check looks for the existence of a directory instead of a file (Optional Default: false)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FileContentsPredicate *PredicateInput         `json:"fileContentsPredicate,omitempty" yaml:"fileContentsPredicate,omitempty"`                 // Condition to match the file content (Optional)
	FilePaths             []string                `json:"filePaths" yaml:"filePaths" example:"['/usr/local/bin', '/home/opslevel']"`              // Restrict the search to certain file paths (Required)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId               ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                  string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	UseAbsoluteRoot       *Nullable[bool]         `json:"useAbsoluteRoot,omitempty" yaml:"useAbsoluteRoot,omitempty" example:"false"`             // Whether the checks looks at the absolute root of a repo or the relative root (the directory specified when attached a repo to a service) (Optional Default: false)
}

// CheckRepositoryFileUpdateInput Specifies the input fields used to update a repo file check
type CheckRepositoryFileUpdateInput struct {
	CategoryId            *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`    // The id of the category the check belongs to (Optional)
	DirectorySearch       *Nullable[bool]         `json:"directorySearch,omitempty" yaml:"directorySearch,omitempty" example:"false"`                    // Whether the check looks for the existence of a directory instead of a file (Optional Default: false)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`               // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                    // Whether the check is enabled or not (Optional)
	FileContentsPredicate *PredicateUpdateInput   `json:"fileContentsPredicate,omitempty" yaml:"fileContentsPredicate,omitempty"`                        // Condition to match the file content (Optional)
	FilePaths             *Nullable[[]string]     `json:"filePaths,omitempty" yaml:"filePaths,omitempty" example:"['/usr/local/bin', '/home/opslevel']"` // Restrict the search to certain file paths (Optional)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`        // The id of the filter the check belongs to (Optional)
	Id                    ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                        // The id of the check to be updated (Required)
	LevelId               *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`          // The id of the level the check belongs to (Optional)
	Name                  *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                  // The display name of the check (Optional)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                                // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`          // The id of the owner of the check (Optional)
	UseAbsoluteRoot       *Nullable[bool]         `json:"useAbsoluteRoot,omitempty" yaml:"useAbsoluteRoot,omitempty" example:"false"`                    // Whether the checks looks at the absolute root of a repo or the relative root (the directory specified when attached a repo to a service) (Optional Default: false)
}

// CheckRepositoryGrepCreateInput Specifies the input fields used to create a repo grep check
type CheckRepositoryGrepCreateInput struct {
	CategoryId            ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	DirectorySearch       *Nullable[bool]         `json:"directorySearch,omitempty" yaml:"directorySearch,omitempty" example:"false"`             // Whether the check looks for the existence of a directory instead of a file (Optional Default: false)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FileContentsPredicate PredicateInput          `json:"fileContentsPredicate" yaml:"fileContentsPredicate"`                                     // Condition to match the file content (Required)
	FilePaths             []string                `json:"filePaths" yaml:"filePaths" example:"['/usr/local/bin', '/home/opslevel']"`              // Restrict the search to certain file paths (Required)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId               ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                  string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckRepositoryGrepUpdateInput Specifies the input fields used to update a repo file check
type CheckRepositoryGrepUpdateInput struct {
	CategoryId            *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`    // The id of the category the check belongs to (Optional)
	DirectorySearch       *Nullable[bool]         `json:"directorySearch,omitempty" yaml:"directorySearch,omitempty" example:"false"`                    // Whether the check looks for the existence of a directory instead of a file (Optional Default: false)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`               // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                    // Whether the check is enabled or not (Optional)
	FileContentsPredicate *PredicateUpdateInput   `json:"fileContentsPredicate,omitempty" yaml:"fileContentsPredicate,omitempty"`                        // Condition to match the file content (Optional)
	FilePaths             *Nullable[[]string]     `json:"filePaths,omitempty" yaml:"filePaths,omitempty" example:"['/usr/local/bin', '/home/opslevel']"` // Restrict the search to certain file paths (Optional)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`        // The id of the filter the check belongs to (Optional)
	Id                    ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                        // The id of the check to be updated (Required)
	LevelId               *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`          // The id of the level the check belongs to (Optional)
	Name                  *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                  // The display name of the check (Optional)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                                // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`          // The id of the owner of the check (Optional)
}

// CheckRepositoryIntegratedCreateInput Specifies the input fields used to create a repository integrated check
type CheckRepositoryIntegratedCreateInput struct {
	CategoryId ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckRepositoryIntegratedUpdateInput Specifies the input fields used to update a repository integrated check
type CheckRepositoryIntegratedUpdateInput struct {
	CategoryId *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckRepositorySearchCreateInput Specifies the input fields used to create a repo search check
type CheckRepositorySearchCreateInput struct {
	CategoryId            ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FileContentsPredicate PredicateInput          `json:"fileContentsPredicate" yaml:"fileContentsPredicate"`                                     // Condition to match the text content (Required)
	FileExtensions        *Nullable[[]string]     `json:"fileExtensions,omitempty" yaml:"fileExtensions,omitempty" example:"['go', 'py', 'rb']"`  // Restrict the search to files of given extensions. Extensions should contain only letters and numbers. For example: `['py', 'rb']` (Optional)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId               ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                  string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckRepositorySearchUpdateInput Specifies the input fields used to update a repo search check
type CheckRepositorySearchUpdateInput struct {
	CategoryId            *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn              *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled               *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FileContentsPredicate *PredicateUpdateInput   `json:"fileContentsPredicate,omitempty" yaml:"fileContentsPredicate,omitempty"`                     // Condition to match the text content (Optional)
	FileExtensions        *Nullable[[]string]     `json:"fileExtensions,omitempty" yaml:"fileExtensions,omitempty" example:"['go', 'py', 'rb']"`      // Restrict the search to files of given extensions. Extensions should contain only letters and numbers. For example: `['py', 'rb']` (Optional)
	FilterId              *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                    ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId               *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                  *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                 *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId               *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckServiceConfigurationCreateInput Specifies the input fields used to create a configuration check
type CheckServiceConfigurationCreateInput struct {
	CategoryId ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckServiceConfigurationUpdateInput Specifies the input fields used to update a configuration check
type CheckServiceConfigurationUpdateInput struct {
	CategoryId *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckServiceDependencyCreateInput Specifies the input fields used to create a service dependency check
type CheckServiceDependencyCreateInput struct {
	CategoryId ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId    ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
}

// CheckServiceDependencyUpdateInput Specifies the input fields used to update a service dependency check
type CheckServiceDependencyUpdateInput struct {
	CategoryId *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn   *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled    *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId    *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes      *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
}

// CheckServiceOwnershipCreateInput Specifies the input fields used to create an ownership check
type CheckServiceOwnershipCreateInput struct {
	CategoryId           ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	ContactMethod        *Nullable[string]       `json:"contactMethod,omitempty" yaml:"contactMethod,omitempty" example:"example_value"`         // The type of contact method that an owner should provide (Optional)
	EnableOn             *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled              *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId             *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId              ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                 string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId              *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	RequireContactMethod *Nullable[bool]         `json:"requireContactMethod,omitempty" yaml:"requireContactMethod,omitempty" example:"false"`   // Whether to require a contact method for a service owner or not (Optional)
	TagKey               *Nullable[string]       `json:"tagKey,omitempty" yaml:"tagKey,omitempty" example:"example_value"`                       // The tag key that should exist for a service owner (Optional)
	TagPredicate         *PredicateInput         `json:"tagPredicate,omitempty" yaml:"tagPredicate,omitempty"`                                   // The condition that should be satisfied by the tag value (Optional)
}

// CheckServiceOwnershipUpdateInput Specifies the input fields used to update an ownership check
type CheckServiceOwnershipUpdateInput struct {
	CategoryId           *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	ContactMethod        *Nullable[string]       `json:"contactMethod,omitempty" yaml:"contactMethod,omitempty" example:"example_value"`             // The type of contact method that an owner should provide (Optional)
	EnableOn             *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled              *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId             *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                   ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId              *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                 *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId              *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	RequireContactMethod *Nullable[bool]         `json:"requireContactMethod,omitempty" yaml:"requireContactMethod,omitempty" example:"false"`       // Whether to require a contact method for a service owner or not (Optional)
	TagKey               *Nullable[string]       `json:"tagKey,omitempty" yaml:"tagKey,omitempty" example:"example_value"`                           // The tag key that should exist for a service owner (Optional)
	TagPredicate         *PredicateUpdateInput   `json:"tagPredicate,omitempty" yaml:"tagPredicate,omitempty"`                                       // The condition that should be satisfied by the tag value (Optional)
}

// CheckServicePropertyCreateInput Specifies the input fields used to create a service property check
type CheckServicePropertyCreateInput struct {
	CategoryId             ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	ComponentType          *IdentifierInput        `json:"componentType,omitempty" yaml:"componentType,omitempty"`                                 // The Component Type that a custom property belongs to. Defaults to Service properties if not provided (Optional)
	EnableOn               *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled                *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId               *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId                ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                   string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                  *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId                *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	PropertyDefinition     *IdentifierInput        `json:"propertyDefinition,omitempty" yaml:"propertyDefinition,omitempty"`                       // The secondary key of the property that the check will verify (e.g. the specific custom property) (Optional)
	PropertyValuePredicate *PredicateInput         `json:"propertyValuePredicate,omitempty" yaml:"propertyValuePredicate,omitempty"`               // The condition that should be satisfied by the service property value (Optional)
	ServiceProperty        ServicePropertyTypeEnum `json:"serviceProperty" yaml:"serviceProperty" example:"custom_property"`                       // The property of the service that the check will verify (Required)
}

// CheckServicePropertyUpdateInput Specifies the input fields used to update a service property check
type CheckServicePropertyUpdateInput struct {
	CategoryId             *Nullable[ID]            `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	ComponentType          *IdentifierInput         `json:"componentType,omitempty" yaml:"componentType,omitempty"`                                     // The Component Type that a custom property belongs to. Defaults to Service properties if not provided (Optional)
	EnableOn               *Nullable[iso8601.Time]  `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled                *Nullable[bool]          `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId               *Nullable[ID]            `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                     ID                       `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId                *Nullable[ID]            `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                   *Nullable[string]        `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                  *string                  `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId                *Nullable[ID]            `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	PropertyDefinition     *IdentifierInput         `json:"propertyDefinition,omitempty" yaml:"propertyDefinition,omitempty"`                           // The secondary key of the property that the check will verify (e.g. the specific custom property) (Optional)
	PropertyValuePredicate *PredicateUpdateInput    `json:"propertyValuePredicate,omitempty" yaml:"propertyValuePredicate,omitempty"`                   // The condition that should be satisfied by the service property value (Optional)
	ServiceProperty        *ServicePropertyTypeEnum `json:"serviceProperty,omitempty" yaml:"serviceProperty,omitempty" example:"custom_property"`       // The property of the service that the check will verify (Optional)
}

// CheckTagDefinedCreateInput Specifies the input fields used to create a tag check
type CheckTagDefinedCreateInput struct {
	CategoryId   ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn     *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled      *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	FilterId     *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId      ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name         string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes        *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId      *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	TagKey       string                  `json:"tagKey" yaml:"tagKey" example:"example_value"`                                           // The tag key where the tag predicate should be applied (Required)
	TagPredicate *PredicateInput         `json:"tagPredicate,omitempty" yaml:"tagPredicate,omitempty"`                                   // The condition that should be satisfied by the tag value (Optional)
}

// CheckTagDefinedUpdateInput Specifies the input fields used to update a tag defined check
type CheckTagDefinedUpdateInput struct {
	CategoryId   *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn     *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled      *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	FilterId     *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id           ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId      *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name         *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes        *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId      *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	TagKey       *Nullable[string]       `json:"tagKey,omitempty" yaml:"tagKey,omitempty" example:"example_value"`                           // The tag key where the tag predicate should be applied (Optional)
	TagPredicate *PredicateUpdateInput   `json:"tagPredicate,omitempty" yaml:"tagPredicate,omitempty"`                                       // The condition that should be satisfied by the tag value (Optional)
}

// CheckToPromoteInput Specifies the input fields used to promote a campaign check to the rubric
type CheckToPromoteInput struct {
	CategoryId ID `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the category that the promoted check will be linked to (Required)
	CheckId    ID `json:"checkId" yaml:"checkId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The ID of the check to be promoted to the rubric (Required)
	LevelId    ID `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The ID of the level that the promoted check will be linked to (Required)
}

// CheckToolUsageCreateInput Specifies the input fields used to create a tool usage check
type CheckToolUsageCreateInput struct {
	CategoryId           ID                      `json:"categoryId" yaml:"categoryId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // The id of the category the check belongs to (Required)
	EnableOn             *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`        // The date when the check will be automatically enabled (Optional)
	Enabled              *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                             // Whether the check is enabled or not (Optional Default: false)
	EnvironmentPredicate *PredicateInput         `json:"environmentPredicate,omitempty" yaml:"environmentPredicate,omitempty"`                   // The condition that the environment should satisfy to be evaluated (Optional)
	FilterId             *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the filter of the check (Optional)
	LevelId              ID                      `json:"levelId" yaml:"levelId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                       // The id of the level the check belongs to (Required)
	Name                 string                  `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the check (Required)
	Notes                *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                         // Additional information about the check (Optional)
	OwnerId              *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`   // The id of the team that owns the check (Optional)
	ToolCategory         ToolCategory            `json:"toolCategory" yaml:"toolCategory" example:"admin"`                                       // The category that the tool belongs to (Required)
	ToolNamePredicate    *PredicateInput         `json:"toolNamePredicate,omitempty" yaml:"toolNamePredicate,omitempty"`                         // The condition that the tool name should satisfy to be evaluated (Optional)
	ToolUrlPredicate     *PredicateInput         `json:"toolUrlPredicate,omitempty" yaml:"toolUrlPredicate,omitempty"`                           // The condition that the tool url should satisfy to be evaluated (Optional)
}

// CheckToolUsageUpdateInput Specifies the input fields used to update a tool usage check
type CheckToolUsageUpdateInput struct {
	CategoryId           *Nullable[ID]           `json:"categoryId,omitempty" yaml:"categoryId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the category the check belongs to (Optional)
	EnableOn             *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" example:"2025-01-05T01:00:00.000Z"`            // The date when the check will be automatically enabled (Optional)
	Enabled              *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" example:"false"`                                 // Whether the check is enabled or not (Optional)
	EnvironmentPredicate *PredicateUpdateInput   `json:"environmentPredicate,omitempty" yaml:"environmentPredicate,omitempty"`                       // The condition that the environment should satisfy to be evaluated (Optional)
	FilterId             *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`     // The id of the filter the check belongs to (Optional)
	Id                   ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The id of the check to be updated (Required)
	LevelId              *Nullable[ID]           `json:"levelId,omitempty" yaml:"levelId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the level the check belongs to (Optional)
	Name                 *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                               // The display name of the check (Optional)
	Notes                *string                 `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                             // Additional information about the check (Optional)
	OwnerId              *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of the check (Optional)
	ToolCategory         *ToolCategory           `json:"toolCategory,omitempty" yaml:"toolCategory,omitempty" example:"admin"`                       // The category that the tool belongs to (Optional)
	ToolNamePredicate    *PredicateUpdateInput   `json:"toolNamePredicate,omitempty" yaml:"toolNamePredicate,omitempty"`                             // The condition that the tool name should satisfy to be evaluated (Optional)
	ToolUrlPredicate     *PredicateUpdateInput   `json:"toolUrlPredicate,omitempty" yaml:"toolUrlPredicate,omitempty"`                               // The condition that the tool url should satisfy to be evaluated (Optional)
}

// CodeIssueIdentifierInput Input for identifying a code issue
type CodeIssueIdentifierInput struct {
	CodeIssueProject CodeIssueProjectIdentifierInput `json:"codeIssueProject" yaml:"codeIssueProject"`                               // Identifier of the code issue project to associate with this issue (Required)
	ExternalId       string                          `json:"externalId" yaml:"externalId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // External identifier for this issue (Required)
}

// CodeIssueInput Input for creating a new code issue
type CodeIssueInput struct {
	Cves          []CommonVulnerabilityEnumerationInput `json:"cves,omitempty" yaml:"cves,omitempty" example:"LIST_TODO"`                                // List of CVE identifiers related to this issue (Optional)
	Cwes          []CommonWeaknessEnumerationInput      `json:"cwes,omitempty" yaml:"cwes,omitempty" example:"LIST_TODO"`                                // List of CWE identifiers related to this issue (Optional)
	Identifier    CodeIssueIdentifierInput              `json:"identifier" yaml:"identifier"`                                                            // Identifier of the code issue project to associate with this issue (Required)
	IntroducedAt  *Nullable[iso8601.Time]               `json:"introducedAt,omitempty" yaml:"introducedAt,omitempty" example:"2025-01-05T01:00:00.000Z"` // Timestamp when this issue was introduced (Optional)
	IssueCategory *Nullable[string]                     `json:"issueCategory,omitempty" yaml:"issueCategory,omitempty" example:"example_value"`          // Category of this code issue. Required to create code issue (Optional)
	Name          *Nullable[string]                     `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                            // Name of the code issue. Required to create code issue (Optional)
	Severity      *Nullable[string]                     `json:"severity,omitempty" yaml:"severity,omitempty" example:"example_value"`                    // Severity level of this issue (Optional)
	Url           *Nullable[string]                     `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"`                              // URL with more information about this issue (Optional)
}

// CodeIssueProjectIdentifierInput Input for upserting a code issue project by external ID
type CodeIssueProjectIdentifierInput struct {
	ExternalId  string          `json:"externalId" yaml:"externalId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // External ID of the code issue project (Required)
	Integration IdentifierInput `json:"integration" yaml:"integration"`                                         // Integration Identifier (Required)
}

// CodeIssueProjectInput Input for upserting a code issue project by external ID
type CodeIssueProjectInput struct {
	Identifier CodeIssueProjectIdentifierInput `json:"identifier" yaml:"identifier"`                                 // Code Issue Project Identifier (Required)
	Name       *Nullable[string]               `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"` // Name of code issue project (Optional)
	Url        *Nullable[string]               `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"`   // URL of code issue project (Optional)
}

// CodeIssueProjectResourceConnectInput Input for connecting a code issue project to a service / repository using their IDs
type CodeIssueProjectResourceConnectInput struct {
	CodeIssueProjectIds []ID `json:"codeIssueProjectIds" yaml:"codeIssueProjectIds" example:"LIST_TODO"`     // IDs of the code issue project to connect (Required)
	ResourceId          ID   `json:"resourceId" yaml:"resourceId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // ID of the service or repository to connect to the code issue project (Required)
}

// CodeIssueProjectResourceDisconnectInput Input for disconnecting a code issue project from a service using their IDs
type CodeIssueProjectResourceDisconnectInput struct {
	CodeIssueProjectId ID `json:"codeIssueProjectId" yaml:"codeIssueProjectId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // ID of the code issue project to disconnect (Required)
	ResourceId         ID `json:"resourceId" yaml:"resourceId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                 // ID of the service to disconnect from the code issue project (Required)
}

// CodeIssueResolutionTimeInput The allowed threshold for how long an issue has been detected before the check starts failing
type CodeIssueResolutionTimeInput struct {
	Unit  CodeIssueResolutionTimeUnitEnum `json:"unit" yaml:"unit" example:"day"` //  (Required)
	Value int                             `json:"value" yaml:"value" example:"3"` //  (Required)
}

// CommonVulnerabilityEnumerationInput Input for a Common Vulnerability Enumeration
type CommonVulnerabilityEnumerationInput struct {
	Identifier string            `json:"identifier" yaml:"identifier" example:"example_value"`       //  (Required)
	Url        *Nullable[string] `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"` //  (Optional)
}

// CommonWeaknessEnumerationInput Input for a Common Weakness Enumeration
type CommonWeaknessEnumerationInput struct {
	Identifier string            `json:"identifier" yaml:"identifier" example:"example_value"`       //  (Required)
	Url        *Nullable[string] `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"` //  (Optional)
}

// ComponentTypeIconInput The input for defining a component type's icon
type ComponentTypeIconInput struct {
	Color string                `json:"color" yaml:"color" example:"example_value"` // The color, represented as a hexcode, for the icon (Required)
	Name  ComponentTypeIconEnum `json:"name" yaml:"name" example:"PhActivity"`      // The name of the icon in Phosphor icons for Vue, e.g. `PhBird`. See https://phosphoricons.com/ for a full list (Required)
}

// ComponentTypeInput Specifies the input fields used to create a component type
type ComponentTypeInput struct {
	Alias       *Nullable[string]                       `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The unique alias of the component type (Optional)
	Description *Nullable[string]                       `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the component type (Optional)
	Icon        *ComponentTypeIconInput                 `json:"icon,omitempty" yaml:"icon,omitempty"`                                       // The icon associated with the component type (Optional)
	Name        *Nullable[string]                       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The unique name of the component type (Optional)
	Properties  *[]ComponentTypePropertyDefinitionInput `json:"properties,omitempty" yaml:"properties,omitempty" example:"[]"`              // A list of property definitions for the component type (Optional)
}

// ComponentTypePropertyDefinitionInput The input for defining a property
type ComponentTypePropertyDefinitionInput struct {
	Alias                 string                    `json:"alias" yaml:"alias" example:"example_value"`                               // The human-friendly, unique identifier for the resource (Required)
	AllowedInConfigFiles  bool                      `json:"allowedInConfigFiles" yaml:"allowedInConfigFiles" example:"false"`         // Whether or not the property is allowed to be set in opslevel.yml config files (Required)
	Description           string                    `json:"description" yaml:"description" example:"example_value"`                   // The description of the property definition (Required)
	LockedStatus          *PropertyLockedStatusEnum `json:"lockedStatus,omitempty" yaml:"lockedStatus,omitempty" example:"ui_locked"` // Restricts what sources are able to assign values to this property (Optional)
	Name                  string                    `json:"name" yaml:"name" example:"example_value"`                                 // The name of the property definition (Required)
	PropertyDisplayStatus PropertyDisplayStatusEnum `json:"propertyDisplayStatus" yaml:"propertyDisplayStatus" example:"hidden"`      // The UI display status of the custom property (Required)
	Schema                JSONSchema                `json:"schema" yaml:"schema" example:"SCHEMA_TBD"`                                // The schema of the property definition (Required)
}

// ContactCreateInput Specifies the input fields used to create a contact
type ContactCreateInput struct {
	Address     string            `json:"address" yaml:"address" example:"example_value"`                                             // The contact address. Examples: support@company.com for type `email`, https://opslevel.com for type `web` (Required)
	DisplayName *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"`                 // The name shown in the UI for the contact (Optional)
	DisplayType *Nullable[string] `json:"displayType,omitempty" yaml:"displayType,omitempty" example:"example_value"`                 // The type shown in the UI for the contact (Optional)
	ExternalId  *Nullable[string] `json:"externalId,omitempty" yaml:"externalId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The remote identifier of the contact method (Optional)
	OwnerId     *Nullable[ID]     `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`       // The id of the owner of this contact (Optional)
	TeamAlias   *Nullable[string] `json:"teamAlias,omitempty" yaml:"teamAlias,omitempty" example:"example_value"`                     // The alias of the team the contact belongs to (Optional)
	TeamId      *Nullable[ID]     `json:"teamId,omitempty" yaml:"teamId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`         // The id of the team the contact belongs to (Optional)
	Type        ContactType       `json:"type" yaml:"type" example:"email"`                                                           // The method of contact [email, slack, slack_handle, web, microsoft_teams] (Required)
}

// ContactDeleteInput Specifies the input fields used to delete a contact
type ContactDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The `id` of the contact you wish to delete (Required)
}

// ContactInput Specifies the input fields used to create a contact
type ContactInput struct {
	Address     string            `json:"address" yaml:"address" example:"example_value"`                             // The contact address. Examples: support@company.com for type `email`, https://opslevel.com for type `web` (Required)
	DisplayName *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"` // The name shown in the UI for the contact (Optional)
	Type        ContactType       `json:"type" yaml:"type" example:"email"`                                           // The method of contact [email, slack, slack_handle, web, microsoft_teams] (Required)
}

// ContactUpdateInput Specifies the input fields used to update a contact
type ContactUpdateInput struct {
	Address     *Nullable[string] `json:"address,omitempty" yaml:"address,omitempty" example:"example_value"`                         // The contact address. Examples: support@company.com for type `email`, https://opslevel.com for type `web` (Optional)
	DisplayName *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"`                 // The name shown in the UI for the contact (Optional)
	DisplayType *Nullable[string] `json:"displayType,omitempty" yaml:"displayType,omitempty" example:"example_value"`                 // The type shown in the UI for the contact (Optional)
	ExternalId  *Nullable[string] `json:"externalId,omitempty" yaml:"externalId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The remote identifier of the contact method (Optional)
	Id          ID                `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                     // The unique identifier for the contact (Required)
	MakeDefault *Nullable[bool]   `json:"makeDefault,omitempty" yaml:"makeDefault,omitempty" example:"false"`                         // Makes the contact the default for the given type. Only available for team contacts (Optional)
	Type        *ContactType      `json:"type,omitempty" yaml:"type,omitempty" example:"email"`                                       // The method of contact [email, slack, slack_handle, web, microsoft_teams] (Optional)
}

// CustomActionsTriggerDefinitionCreateInput Specifies the input fields used in the `customActionsTriggerDefinitionCreate` mutation
type CustomActionsTriggerDefinitionCreateInput struct {
	AccessControl          *CustomActionsTriggerDefinitionAccessControlEnum `json:"accessControl,omitempty" yaml:"accessControl,omitempty" example:"admins"`                          // The set of users that should be able to use the trigger definition (Optional)
	ActionId               *Nullable[ID]                                    `json:"actionId,omitempty" yaml:"actionId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The action that will be triggered by the Trigger Definition (Optional)
	ApprovalConfig         *ApprovalConfigInput                             `json:"approvalConfig,omitempty" yaml:"approvalConfig,omitempty"`                                         // Config for approval of action (Optional)
	Description            *Nullable[string]                                `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`                       // The description of what the Trigger Definition will do, supports Markdown (Optional)
	ExtendedTeamAccess     *[]IdentifierInput                               `json:"extendedTeamAccess,omitempty" yaml:"extendedTeamAccess,omitempty" example:"[]"`                    // The set of additional teams who can invoke this Trigger Definition (Optional)
	EntityType             *CustomActionsEntityTypeEnum                     `json:"entityType,omitempty" yaml:"entityType,omitempty" example:"GLOBAL"`                                // The entity type to associate with the Trigger Definition (Optional Default: SERVICE)
	FilterId               *Nullable[ID]                                    `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The filter that will determine which services apply to the Trigger Definition (Optional)
	ManualInputsDefinition *Nullable[string]                                `json:"manualInputsDefinition,omitempty" yaml:"manualInputsDefinition,omitempty" example:"example_value"` // The YAML definition of custom inputs for the Trigger Definition (Optional)
	Name                   string                                           `json:"name" yaml:"name" example:"example_value"`                                                         // The name of the Trigger Definition (Required)
	OwnerId                ID                                               `json:"ownerId" yaml:"ownerId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                 // The owner of the Trigger Definition (Required)
	Published              *Nullable[bool]                                  `json:"published,omitempty" yaml:"published,omitempty" example:"false"`                                   // The published state of the action; true if the definition is ready for use; false if it is a draft (Optional)
	ResponseTemplate       *Nullable[string]                                `json:"responseTemplate,omitempty" yaml:"responseTemplate,omitempty" example:"example_value"`             // The liquid template used to parse the response from the External Action (Optional)
}

// CustomActionsTriggerDefinitionUpdateInput Specifies the input fields used in the `customActionsTriggerDefinitionUpdate` mutation
type CustomActionsTriggerDefinitionUpdateInput struct {
	AccessControl          *CustomActionsTriggerDefinitionAccessControlEnum `json:"accessControl,omitempty" yaml:"accessControl,omitempty" example:"admins"`                          // The set of users that should be able to use the trigger definition (Optional)
	Action                 *CustomActionsWebhookActionUpdateInput           `json:"action,omitempty" yaml:"action,omitempty"`                                                         // The details for the action to update for the Trigger Definition (Optional)
	ActionId               *Nullable[ID]                                    `json:"actionId,omitempty" yaml:"actionId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The action that will be triggered by the Trigger Definition (Optional)
	ApprovalConfig         *ApprovalConfigInput                             `json:"approvalConfig,omitempty" yaml:"approvalConfig,omitempty"`                                         // Config for approval of action (Optional)
	Description            *Nullable[string]                                `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`                       // The description of what the Trigger Definition will do, support Markdown (Optional)
	ExtendedTeamAccess     *[]IdentifierInput                               `json:"extendedTeamAccess,omitempty" yaml:"extendedTeamAccess,omitempty" example:"[]"`                    // The set of additional teams who can invoke this Trigger Definition (Optional)
	EntityType             *CustomActionsEntityTypeEnum                     `json:"entityType,omitempty" yaml:"entityType,omitempty" example:"GLOBAL"`                                // The entity type to associate with the Trigger Definition (Optional Default: SERVICE)
	FilterId               *Nullable[ID]                                    `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The filter that will determine which services apply to the Trigger Definition (Optional)
	Id                     ID                                               `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                           // The ID of the trigger definition (Required)
	ManualInputsDefinition *Nullable[string]                                `json:"manualInputsDefinition,omitempty" yaml:"manualInputsDefinition,omitempty" example:"example_value"` // The YAML definition of custom inputs for the Trigger Definition (Optional)
	Name                   *Nullable[string]                                `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                     // The name of the Trigger Definition (Optional)
	OwnerId                *Nullable[ID]                                    `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The owner of the Trigger Definition (Optional)
	Published              *Nullable[bool]                                  `json:"published,omitempty" yaml:"published,omitempty" example:"false"`                                   // The published state of the action; true if the definition is ready for use; false if it is a draft (Optional)
	ResponseTemplate       *Nullable[string]                                `json:"responseTemplate,omitempty" yaml:"responseTemplate,omitempty" example:"example_value"`             // The liquid template used to parse the response from the External Action (Optional)
}

// CustomActionsTriggerInvokeInput Inputs that specify the trigger definition to invoke, the user that invoked it, and what object it is invoked on
type CustomActionsTriggerInvokeInput struct {
	ManualInputs      JSON             `json:"manualInputs,omitempty" yaml:"manualInputs,omitempty" example:"{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}"` // Additional details provided for a specific invocation of this Custom Action (Optional Default: "{}")
	TargetObject      *IdentifierInput `json:"targetObject,omitempty" yaml:"targetObject,omitempty"`                                                                                                                     // The identifier of the object to perform the custom action on (Optional)
	TriggerDefinition IdentifierInput  `json:"triggerDefinition" yaml:"triggerDefinition"`                                                                                                                               // The trigger definition to invoke (Required)
}

// CustomActionsWebhookActionCreateInput Specifies the input fields used in the `customActionsWebhookActionCreate` mutation
type CustomActionsWebhookActionCreateInput struct {
	Async          *bool                       `json:"async,omitempty" yaml:"async,omitempty" example:"false"`                                                                                                         // Whether the action expects an additional, asynchronous response upon completion (Required Default: false)
	Description    *Nullable[string]           `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`                                                                                     // The description that gets assigned to the Webhook Action you're creating (Optional)
	Headers        *JSON                       `json:"headers,omitempty" yaml:"headers,omitempty" example:"{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}"` // HTTP headers be passed along with your Webhook when triggered (Optional)
	HttpMethod     CustomActionsHttpMethodEnum `json:"httpMethod" yaml:"httpMethod" example:"DELETE"`                                                                                                                  // HTTP used when the Webhook is triggered. Either POST or PUT (Required)
	LiquidTemplate *Nullable[string]           `json:"liquidTemplate,omitempty" yaml:"liquidTemplate,omitempty" example:"example_value"`                                                                               // Template that can be used to generate a Webhook payload (Optional)
	Name           string                      `json:"name" yaml:"name" example:"example_value"`                                                                                                                       // The name that gets assigned to the Webhook Action you're creating (Required)
	WebhookUrl     string                      `json:"webhookUrl" yaml:"webhookUrl" example:"example_value"`                                                                                                           // The URL that you wish to send the Webhook to when triggered (Required)
}

// CustomActionsWebhookActionUpdateInput Inputs that specify the details of a Webhook Action you wish to update
type CustomActionsWebhookActionUpdateInput struct {
	Async          *bool                        `json:"async,omitempty" yaml:"async,omitempty" example:"false"`                                                                                                         // Whether the action expects an additional, asynchronous response upon completion (Optional)
	Description    *Nullable[string]            `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`                                                                                     // The description that gets assigned to the Webhook Action you're creating (Optional)
	Headers        *JSON                        `json:"headers,omitempty" yaml:"headers,omitempty" example:"{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}"` // HTTP headers be passed along with your Webhook when triggered (Optional)
	HttpMethod     *CustomActionsHttpMethodEnum `json:"httpMethod,omitempty" yaml:"httpMethod,omitempty" example:"DELETE"`                                                                                              // HTTP used when the Webhook is triggered. Either POST or PUT (Optional)
	Id             ID                           `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                                                                                         // The ID of the Webhook Action you wish to update (Required)
	LiquidTemplate *Nullable[string]            `json:"liquidTemplate,omitempty" yaml:"liquidTemplate,omitempty" example:"example_value"`                                                                               // Template that can be used to generate a Webhook payload (Optional)
	Name           *Nullable[string]            `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                                                                                   // The name that gets assigned to the Webhook Action you're creating (Optional)
	WebhookUrl     *Nullable[string]            `json:"webhookUrl,omitempty" yaml:"webhookUrl,omitempty" example:"example_value"`                                                                                       // The URL that you wish to send the Webhook too when triggered (Optional)
}

// CustomIntegrationInput Input for upserting a custom integration
type CustomIntegrationInput struct {
	Name *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"` // Name of the custom integration type (Optional)
}

// DeleteInput Specifies the input fields used to delete an entity
type DeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the entity to be deleted (Required)
}

// DomainInput Specifies the input fields for a domain
type DomainInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`           // The description for the domain (Optional)
	Name        *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                         // The name for the domain (Optional)
	Note        *Nullable[string] `json:"note,omitempty" yaml:"note,omitempty" example:"example_value"`                         // Additional information about the domain (Optional)
	OwnerId     *Nullable[ID]     `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the owner for the domain (Optional)
}

// EventIntegrationInput
type EventIntegrationInput struct {
	Name *Nullable[string]    `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"` // The name of the event integration (Optional)
	Type EventIntegrationEnum `json:"type" yaml:"type" example:"apiDoc"`                            // The type of event integration to create (Required)
}

// EventIntegrationUpdateInput
type EventIntegrationUpdateInput struct {
	Id   ID     `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the event integration to update (Required)
	Name string `json:"name" yaml:"name" example:"example_value"`               // The name of the event integration (Required)
}

// ExternalResourceIdentifierInput Specifies the input fields to locate resouce created via API in OpsLevel
type ExternalResourceIdentifierInput struct {
	ExternalId  string          `json:"externalId" yaml:"externalId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the resource in your system (Required)
	Integration IdentifierInput `json:"integration" yaml:"integration"`                                         // The integration identifier (Required)
}

// ExternalUuidMutationInput Specifies the input used for modifying a resource's external UUID
type ExternalUuidMutationInput struct {
	ResourceId ID `json:"resourceId" yaml:"resourceId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the resource (Required)
}

// FilterCreateInput Specifies the input fields used to create a filter
type FilterCreateInput struct {
	Connective *ConnectiveEnum         `json:"connective,omitempty" yaml:"connective,omitempty" example:"and"` // The logical operator to be used in conjunction with predicates (Optional)
	Name       string                  `json:"name" yaml:"name" example:"example_value"`                       // The display name of the filter (Required)
	Predicates *[]FilterPredicateInput `json:"predicates,omitempty" yaml:"predicates,omitempty" example:"[]"`  // The list of predicates used to select which services apply to the filter (Optional)
}

// FilterPredicateInput A condition that should be satisfied
type FilterPredicateInput struct {
	CaseSensitive *Nullable[bool]   `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty" example:"false"` //  (Optional)
	Key           PredicateKeyEnum  `json:"key" yaml:"key" example:"aliases"`                                       // The condition key used by the predicate (Required)
	KeyData       *Nullable[string] `json:"keyData,omitempty" yaml:"keyData,omitempty" example:"example_value"`     // Additional data used by the predicate. This field is used by predicates with key = 'tags' to specify the tag key. For example, to create a predicate for services containing the tag 'db:mysql', set keyData = 'db' and value = 'mysql' (Optional)
	Type          PredicateTypeEnum `json:"type" yaml:"type" example:"belongs_to"`                                  // The condition type used by the predicate (Required)
	Value         *Nullable[string] `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"`         // The condition value used by the predicate (Optional)
}

// FilterUpdateInput Specifies the input fields used to update a filter
type FilterUpdateInput struct {
	Connective *ConnectiveEnum         `json:"connective,omitempty" yaml:"connective,omitempty" example:"and"` // The logical operator to be used in conjunction with predicates (Optional)
	Id         ID                      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`         // The id of the filter (Required)
	Name       *Nullable[string]       `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`   // The display name of the filter (Optional)
	Predicates *[]FilterPredicateInput `json:"predicates,omitempty" yaml:"predicates,omitempty" example:"[]"`  // The list of predicates used to select which services apply to the filter. All existing predicates will be replaced by these predicates (Optional)
}

// FireHydrantIntegrationInput A FireHydrant integration input
type FireHydrantIntegrationInput struct {
	ApiKey *Nullable[string] `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"` // The API Key for the FireHydrant API (Optional)
	Name   *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`     // The name for the FireHydrant integration (Optional)
}

// GoogleCloudIntegrationInput Specifies the input fields used to create and update a Google Cloud integration
type GoogleCloudIntegrationInput struct {
	ClientEmail           *Nullable[string]   `json:"clientEmail,omitempty" yaml:"clientEmail,omitempty" example:"example_value"`                      // The service account email OpsLevel uses to access the Google Cloud account (Optional)
	Name                  *Nullable[string]   `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                                    // The name of the integration (Optional)
	OwnershipTagKeys      *Nullable[[]string] `json:"ownershipTagKeys,omitempty" yaml:"ownershipTagKeys,omitempty" example:"['tag_key1', 'tag_key2']"` // An array of tag keys used to associate ownership from an integration. Max 5 (Optional)
	PrivateKey            *Nullable[string]   `json:"privateKey,omitempty" yaml:"privateKey,omitempty" example:"example_value"`                        // The private key for the service account that OpsLevel uses to access the Google Cloud account (Optional)
	TagsOverrideOwnership *Nullable[bool]     `json:"tagsOverrideOwnership,omitempty" yaml:"tagsOverrideOwnership,omitempty" example:"false"`          // Allow tags imported from Google Cloud to override ownership set in OpsLevel directly (Optional)
}

// IdentifierInput Specifies the input fields used to identify a resource
type IdentifierInput struct {
	Alias *string `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The human-friendly, unique identifier for the resource (Optional)
	Id    *ID     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the resource (Optional)
}

// InfrastructureResourceInput Specifies the input fields for a infrastructure resource
type InfrastructureResourceInput struct {
	Data                 *JSON                                    `json:"data,omitempty" yaml:"data,omitempty" example:"{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}"` // The data for the infrastructure_resource (Optional)
	OwnerId              *Nullable[ID]                            `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                                                     // The id of the owner for the infrastructure_resource (Optional)
	ProviderData         *InfrastructureResourceProviderDataInput `json:"providerData,omitempty" yaml:"providerData,omitempty"`                                                                                                     // Data about the provider of the infrastructure resource (Optional)
	ProviderResourceType *Nullable[string]                        `json:"providerResourceType,omitempty" yaml:"providerResourceType,omitempty" example:"example_value"`                                                             // The type of the infrastructure resource in its provider (Optional)
	Schema               *InfrastructureResourceSchemaInput       `json:"schema,omitempty" yaml:"schema,omitempty"`                                                                                                                 // The schema for the infrastructure_resource that determines its type (Optional)
}

// InfrastructureResourceProviderDataInput Specifies the input fields for data about an infrastructure resource's provider
type InfrastructureResourceProviderDataInput struct {
	AccountName  string            `json:"accountName" yaml:"accountName" example:"example_value"`                       // The account name of the provider (Required)
	ExternalUrl  *Nullable[string] `json:"externalUrl,omitempty" yaml:"externalUrl,omitempty" example:"example_value"`   // The external URL of the infrastructure resource in its provider (Optional)
	ProviderName *Nullable[string] `json:"providerName,omitempty" yaml:"providerName,omitempty" example:"example_value"` // The name of the provider (e.g. AWS, GCP, Azure) (Optional)
}

// InfrastructureResourceSchemaInput Specifies the schema for an infrastructure resource
type InfrastructureResourceSchemaInput struct {
	Type string `json:"type" yaml:"type" example:"example_value"` // The type of the infrastructure resource (Required)
}

// LevelCreateInput Specifies the input fields used to create a level. The new level will be added as the highest level (greatest level index)
type LevelCreateInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the level (Optional)
	Index       *int              `json:"index,omitempty" yaml:"index,omitempty" example:"3"`                         // an integer allowing this level to be inserted between others. Must be unique per Rubric (Optional)
	Name        string            `json:"name" yaml:"name" example:"example_value"`                                   // The display name of the level (Required)
}

// LevelDeleteInput Specifies the input fields used to delete a level
type LevelDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the level to be deleted (Required)
}

// LevelUpdateInput Specifies the input fields used to update a level
type LevelUpdateInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the level (Optional)
	Id          ID                `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                     // The id of the level to be updated (Required)
	Name        *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The display name of the level (Optional)
}

// ManualCheckFrequencyInput Defines a frequency for the check update
type ManualCheckFrequencyInput struct {
	FrequencyTimeScale FrequencyTimeScale `json:"frequencyTimeScale" yaml:"frequencyTimeScale" example:"day"`          // The time scale type for the frequency (Required)
	FrequencyValue     int                `json:"frequencyValue" yaml:"frequencyValue" example:"3"`                    // The value to be used together with the frequency scale (Required)
	StartingDate       iso8601.Time       `json:"startingDate" yaml:"startingDate" example:"2025-01-05T01:00:00.000Z"` // The date that the check will start to evaluate (Required)
}

// ManualCheckFrequencyUpdateInput Defines a frequency for the check update
type ManualCheckFrequencyUpdateInput struct {
	FrequencyTimeScale *FrequencyTimeScale     `json:"frequencyTimeScale,omitempty" yaml:"frequencyTimeScale,omitempty" example:"day"`          // The time scale type for the frequency (Optional)
	FrequencyValue     *Nullable[int]          `json:"frequencyValue,omitempty" yaml:"frequencyValue,omitempty" example:"3"`                    // The value to be used together with the frequency scale (Optional)
	StartingDate       *Nullable[iso8601.Time] `json:"startingDate,omitempty" yaml:"startingDate,omitempty" example:"2025-01-05T01:00:00.000Z"` // The date that the check will start to evaluate (Optional)
}

// MemberInput Input for specifying members on a group
type MemberInput struct {
	Email string `json:"email" yaml:"email" example:"example_value"` // The user's email (Required)
}

// NewRelicIntegrationAccountsInput
type NewRelicIntegrationAccountsInput struct {
	ApiKey  string `json:"apiKey" yaml:"apiKey" example:"example_value"`   // The API Key for the New Relic API (Required)
	BaseUrl string `json:"baseUrl" yaml:"baseUrl" example:"example_value"` // The API URL for New Relic API (Required)
}

// NewRelicIntegrationInput
type NewRelicIntegrationInput struct {
	ApiKey  *Nullable[string] `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"`   // The API Key for the New Relic API (Optional)
	BaseUrl *Nullable[string] `json:"baseUrl,omitempty" yaml:"baseUrl,omitempty" example:"example_value"` // The API URL for New Relic API (Optional)
}

// OctopusDeployIntegrationInput Specifies the input fields used to create and update an Octopus Deploy integration
type OctopusDeployIntegrationInput struct {
	ApiKey      *Nullable[string] `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"`           // The API Key for the Octopus Deploy API (Optional)
	InstanceUrl *Nullable[string] `json:"instanceUrl,omitempty" yaml:"instanceUrl,omitempty" example:"example_value"` // The URL the Octopus Deploy instance if hosted on (Optional)
	Name        *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The name of the integration (Optional)
}

// PayloadFilterInput Input to be used to filter types
type PayloadFilterInput struct {
	Arg  *Nullable[string] `json:"arg,omitempty" yaml:"arg,omitempty" example:"example_value"`    // Value to be filtered (Optional)
	Key  PayloadFilterEnum `json:"key" yaml:"key" example:"integration_id"`                       // Field to be filtered (Required)
	Type *BasicTypeEnum    `json:"type,omitempty" yaml:"type,omitempty" example:"does_not_equal"` // Type of operation to be applied to value on the field (Optional Default: equals)
}

// PredicateInput A condition that should be satisfied
type PredicateInput struct {
	Type  PredicateTypeEnum `json:"type" yaml:"type" example:"belongs_to"`                          // The condition type used by the predicate (Required)
	Value *Nullable[string] `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"` // The condition value used by the predicate (Optional)
}

// PredicateUpdateInput A condition that should be satisfied
type PredicateUpdateInput struct {
	Type  *PredicateTypeEnum `json:"type,omitempty" yaml:"type,omitempty" example:"belongs_to"`      // The condition type used by the predicate (Optional)
	Value *Nullable[string]  `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"` // The condition value used by the predicate (Optional)
}

// PropertyDefinitionInput The input for defining a property
type PropertyDefinitionInput struct {
	AllowedInConfigFiles  *Nullable[bool]            `json:"allowedInConfigFiles,omitempty" yaml:"allowedInConfigFiles,omitempty" example:"false"`    // Whether or not the property is allowed to be set in opslevel.yml config files (Optional Default: true)
	Description           *Nullable[string]          `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`              // The description of the property definition (Optional)
	LockedStatus          *PropertyLockedStatusEnum  `json:"lockedStatus,omitempty" yaml:"lockedStatus,omitempty" example:"ui_locked"`                // Restricts what sources are able to assign values to this property (Optional)
	Name                  *Nullable[string]          `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                            // The name of the property definition (Optional)
	PropertyDisplayStatus *PropertyDisplayStatusEnum `json:"propertyDisplayStatus,omitempty" yaml:"propertyDisplayStatus,omitempty" example:"hidden"` // The UI display status of the custom property (Optional)
	Schema                *JSONSchema                `json:"schema,omitempty" yaml:"schema,omitempty" example:"SCHEMA_TBD"`                           // The schema of the property definition (Optional)
}

// PropertyInput The input for setting a property
type PropertyInput struct {
	Definition    IdentifierInput        `json:"definition" yaml:"definition"`                                           // The definition of the property (Required)
	Owner         IdentifierInput        `json:"owner" yaml:"owner"`                                                     // The entity that the property has been assigned to (Required)
	OwnerType     *PropertyOwnerTypeEnum `json:"ownerType,omitempty" yaml:"ownerType,omitempty" example:"COMPONENT"`     // The type of the entity that the property has been assigned to. Defaults to `COMPONENT` if alias is provided for `owner` and `definition` (Optional)
	RunValidation *Nullable[bool]        `json:"runValidation,omitempty" yaml:"runValidation,omitempty" example:"false"` // Validate the property value against the schema. On by default (Optional Default: true)
	Value         JsonString             `json:"value" yaml:"value" example:"JSON_TBD"`                                  // The value of the property (Required)
}

// RelationshipDefinition A source, target and relationship type specifying a relationship between two resources
type RelationshipDefinition struct {
	RelationshipDefinition *IdentifierInput     `json:"relationshipDefinition,omitempty" yaml:"relationshipDefinition,omitempty"` // A dynamic definition that specifies how the source and target are related (Optional)
	Source                 IdentifierInput      `json:"source" yaml:"source"`                                                     // The resource that is the source of the relationship (Required)
	Target                 IdentifierInput      `json:"target" yaml:"target"`                                                     // The resource that is the target of the relationship (Required)
	Type                   RelationshipTypeEnum `json:"type" yaml:"type" example:"belongs_to"`                                    // The type of the relationship between source and target (Required)
}

// RelationshipDefinitionInput The input for defining a relationship on a component type
type RelationshipDefinitionInput struct {
	Alias         *string                              `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The unique identifier of the relationship (Optional)
	ComponentType *IdentifierInput                     `json:"componentType,omitempty" yaml:"componentType,omitempty"`                     // The component type to create the relationship on (Optional)
	Description   *Nullable[string]                    `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"` // The description of the relationship (Optional)
	Metadata      *RelationshipDefinitionMetadataInput `json:"metadata,omitempty" yaml:"metadata,omitempty"`                               // The metadata of the relationship (Optional)
	Name          *string                              `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`               // The name of the relationship (Optional)
}

// RelationshipDefinitionMetadataInput The metadata of the relationship
type RelationshipDefinitionMetadataInput struct {
	AllowedTypes []string `json:"allowedTypes,omitempty" yaml:"allowedTypes,omitempty" example:"LIST_TODO"` // The aliases of which types this relationship can target. Valid values include any component type alias on your account, `team`, or `user` (Optional)
	MaxItems     *int     `json:"maxItems,omitempty" yaml:"maxItems,omitempty" example:"3"`                 // The maximum number of records this relationship can associate to the component type. Defaults to null (no maximum) (Optional)
	MinItems     *int     `json:"minItems,omitempty" yaml:"minItems,omitempty" example:"3"`                 // The minimum number of records this relationship must associate to the component type. Defaults to 0 (optional) (Optional)
}

// RepositoryUpdateInput Specifies the input fields used to update a repository
type RepositoryUpdateInput struct {
	Id             ID                                  `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                               // The id of the repository to be updated (Required)
	OwnerId        *Nullable[ID]                       `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The team that owns the repository (Optional)
	SbomGeneration *RepositorySBOMGenerationConfigEnum `json:"sbomGeneration,omitempty" yaml:"sbomGeneration,omitempty" example:"opt_in"`            // The desired configuration state at the repository level for SBOM generation (Optional)
	Visible        *Nullable[bool]                     `json:"visible,omitempty" yaml:"visible,omitempty" example:"false"`                           // Indicates if the repository is visible (Optional)
}

// ScorecardInput Input used to create scorecards
type ScorecardInput struct {
	AffectsOverallServiceLevels *Nullable[bool]   `json:"affectsOverallServiceLevels,omitempty" yaml:"affectsOverallServiceLevels,omitempty" example:"false"` //  (Optional)
	Description                 *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`                         // Description of the scorecard (Optional)
	FilterId                    *Nullable[ID]     `json:"filterId,omitempty" yaml:"filterId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // Filter used by the scorecard to restrict services (Optional)
	Name                        string            `json:"name" yaml:"name" example:"example_value"`                                                           // Name of the scorecard (Required)
	OwnerId                     ID                `json:"ownerId" yaml:"ownerId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                                   // Owner of the scorecard. Can currently be a team or a group (Required)
}

// SecretInput Arguments for secret operations
type SecretInput struct {
	Owner *IdentifierInput  `json:"owner,omitempty" yaml:"owner,omitempty"`                         // The owner of this secret (Optional)
	Value *Nullable[string] `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"` // A sensitive value (Optional)
}

// ServiceCreateInput Specifies the input fields used in the `serviceCreate` mutation
type ServiceCreateInput struct {
	Description           *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`             // A brief description of the service (Optional)
	Framework             *Nullable[string] `json:"framework,omitempty" yaml:"framework,omitempty" example:"example_value"`                 // The primary software development framework that the service uses (Optional)
	Language              *Nullable[string] `json:"language,omitempty" yaml:"language,omitempty" example:"example_value"`                   // The primary programming language that the service is written in (Optional)
	LifecycleAlias        *Nullable[string] `json:"lifecycleAlias,omitempty" yaml:"lifecycleAlias,omitempty" example:"example_value"`       // The lifecycle stage of the service (Optional)
	Name                  string            `json:"name" yaml:"name" example:"example_value"`                                               // The display name of the service (Required)
	OwnerAlias            *Nullable[string] `json:"ownerAlias,omitempty" yaml:"ownerAlias,omitempty" example:"example_value"`               // The team that owns the service (Optional)
	OwnerInput            *IdentifierInput  `json:"ownerInput,omitempty" yaml:"ownerInput,omitempty"`                                       // The owner for this service (Optional)
	Parent                *IdentifierInput  `json:"parent,omitempty" yaml:"parent,omitempty"`                                               // The parent system for the service (Optional)
	Product               *Nullable[string] `json:"product,omitempty" yaml:"product,omitempty" example:"example_value"`                     // A product is an application that your end user interacts with. Multiple services can work together to power a single product (Optional)
	SkipAliasesValidation *Nullable[bool]   `json:"skipAliasesValidation,omitempty" yaml:"skipAliasesValidation,omitempty" example:"false"` // Allows for the creation of a service with invalid aliases (Optional Default: false)
	TierAlias             *Nullable[string] `json:"tierAlias,omitempty" yaml:"tierAlias,omitempty" example:"example_value"`                 // The software tier that the service belongs to (Optional)
	Type                  *IdentifierInput  `json:"type,omitempty" yaml:"type,omitempty"`                                                   // The type of the component (Optional)
}

// ServiceDeleteInput Specifies the input fields used in the `serviceDelete` mutation
type ServiceDeleteInput struct {
	Alias *Nullable[string] `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The alias of the service to be deleted (Optional)
	Id    *Nullable[ID]     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the service to be deleted (Optional)
}

// ServiceDependencyCreateInput Specifies the input fields used for creating a service dependency
type ServiceDependencyCreateInput struct {
	DependencyKey ServiceDependencyKey `json:"dependencyKey" yaml:"dependencyKey"`                             // A source, destination pair specifying a dependency between services (Required)
	Notes         *Nullable[string]    `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"` // Notes for service dependency (Optional)
}

// ServiceDependencyKey A source, destination pair specifying a dependency between services
type ServiceDependencyKey struct {
	Destination           *Nullable[ID]     `json:"destination,omitempty" yaml:"destination,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the service that is depended upon (Optional)
	DestinationIdentifier *IdentifierInput  `json:"destinationIdentifier,omitempty" yaml:"destinationIdentifier,omitempty"`                       // The ID or alias identifier of the service that is depended upon (Optional)
	Notes                 *Nullable[string] `json:"notes,omitempty" yaml:"notes,omitempty" example:"example_value"`                               // Notes about the dependency edge (Optional)
	Source                *Nullable[ID]     `json:"source,omitempty" yaml:"source,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The ID of the service with the dependency (Optional)
	SourceIdentifier      *IdentifierInput  `json:"sourceIdentifier,omitempty" yaml:"sourceIdentifier,omitempty"`                                 // The ID or alias identifier of the service with the dependency (Optional)
}

// ServiceFilterInput Input to be used to filter types
type ServiceFilterInput struct {
	Arg           string               `json:"arg,omitempty" yaml:"arg,omitempty" example:"example_value"`     // Value to be filtered (Optional)
	CaseSensitive bool                 `json:"caseSensitive" yaml:"caseSensitive"`                             // Whether or not the filter should be case sensitive (Optional)
	Connective    *ConnectiveEnum      `json:"connective,omitempty" yaml:"connective,omitempty" example:"and"` // The logical operator to be used in conjunction with multiple filters (requires predicates to be supplied) (Optional Default: or)
	Key           *ServiceFilterEnum   `json:"key,omitempty" yaml:"key,omitempty" example:"alert_status"`      // Field to be filtered (Optional)
	Predicates    []ServiceFilterInput `json:"predicates,omitempty" yaml:"predicates,omitempty" example:"[]"`  // A list of service filter input (Optional)
	Type          *TypeEnum            `json:"type,omitempty" yaml:"type,omitempty" example:"belongs_to"`      // Type of operation to be applied to value on the field (Optional Default: equals)
}

// ServiceLevelNotificationsUpdateInput Specifies the input fields used to update service level notification settings
type ServiceLevelNotificationsUpdateInput struct {
	EnableSlackNotifications *Nullable[bool] `json:"enableSlackNotifications,omitempty" yaml:"enableSlackNotifications,omitempty" example:"false"` // Whether or not to enable receiving slack notifications on service level changes (Optional)
}

// ServiceNoteUpdateInput Specifies the input fields used in the `serviceNoteUpdate` mutation
type ServiceNoteUpdateInput struct {
	Note    *Nullable[string] `json:"note,omitempty" yaml:"note,omitempty" example:"example_value"` // Note about the service (Optional)
	Service IdentifierInput   `json:"service" yaml:"service"`                                       // The identifier for the service (Required)
}

// ServiceRepositoryCreateInput Specifies the input fields used in the `serviceRepositoryCreate` mutation
type ServiceRepositoryCreateInput struct {
	BaseDirectory *Nullable[string] `json:"baseDirectory,omitempty" yaml:"baseDirectory,omitempty" example:"example_value"` // The directory in the repository where service information exists, including the opslevel.yml file. This path is always returned without leading and trailing slashes (Optional)
	DisplayName   *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"`     // The name displayed in the UI for the service repository (Optional)
	Repository    IdentifierInput   `json:"repository" yaml:"repository"`                                                   // The identifier for the repository (Required)
	Service       IdentifierInput   `json:"service" yaml:"service"`                                                         // The identifier for the service (Required)
}

// ServiceRepositoryUpdateInput Specifies the input fields used to update a service repository
type ServiceRepositoryUpdateInput struct {
	BaseDirectory *Nullable[string] `json:"baseDirectory,omitempty" yaml:"baseDirectory,omitempty" example:"example_value"` // The directory in the repository where service information exists, including the opslevel.yml file. This path is always returned without leading and trailing slashes (Optional)
	DisplayName   *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"`     // The name displayed in the UI for the service repository (Optional)
	Id            ID                `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                         // The ID of the service repository to be updated (Required)
}

// ServiceUpdateInput Specifies the input fields used in the `serviceUpdate` mutation
type ServiceUpdateInput struct {
	Alias                 *Nullable[string] `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`                         // The alias of the service to be updated (Optional)
	Description           *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`             // A brief description of the service (Optional)
	Framework             *Nullable[string] `json:"framework,omitempty" yaml:"framework,omitempty" example:"example_value"`                 // The primary software development framework that the service uses (Optional)
	Id                    *Nullable[ID]     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`             // The id of the service to be updated (Optional)
	Language              *Nullable[string] `json:"language,omitempty" yaml:"language,omitempty" example:"example_value"`                   // The primary programming language that the service is written in (Optional)
	LifecycleAlias        *Nullable[string] `json:"lifecycleAlias,omitempty" yaml:"lifecycleAlias,omitempty" example:"example_value"`       // The lifecycle stage of the service (Optional)
	Name                  *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                           // The display name of the service (Optional)
	OwnerAlias            *Nullable[string] `json:"ownerAlias,omitempty" yaml:"ownerAlias,omitempty" example:"example_value"`               // The team that owns the service (Optional)
	OwnerInput            *IdentifierInput  `json:"ownerInput,omitempty" yaml:"ownerInput,omitempty"`                                       // The owner for the service (Optional)
	Parent                *IdentifierInput  `json:"parent,omitempty" yaml:"parent,omitempty"`                                               // The parent system for the service (Optional)
	Product               *Nullable[string] `json:"product,omitempty" yaml:"product,omitempty" example:"example_value"`                     // A product is an application that your end user interacts with. Multiple services can work together to power a single product (Optional)
	SkipAliasesValidation *Nullable[bool]   `json:"skipAliasesValidation,omitempty" yaml:"skipAliasesValidation,omitempty" example:"false"` // Allows updating a service with invalid aliases (Optional Default: false)
	TierAlias             *Nullable[string] `json:"tierAlias,omitempty" yaml:"tierAlias,omitempty" example:"example_value"`                 // The software tier that the service belongs to (Optional)
	Type                  *IdentifierInput  `json:"type,omitempty" yaml:"type,omitempty"`                                                   // The type of the component (Optional)
}

// SnykIntegrationInput Specifies the input fields used to create and update a Snyk integration
type SnykIntegrationInput struct {
	ApiKey  *Nullable[string]          `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"`   // The API Key for the Snyk API (Optional)
	BaseUrl *Nullable[string]          `json:"baseUrl,omitempty" yaml:"baseUrl,omitempty" example:"example_value"` // The base url for your Snyk installation (Optional)
	GroupId *Nullable[string]          `json:"groupId,omitempty" yaml:"groupId,omitempty" example:"example_value"` // The group ID for the Snyk API (Optional)
	Name    *Nullable[string]          `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`       // The name of the integration (Optional)
	Region  *SnykIntegrationRegionEnum `json:"region,omitempty" yaml:"region,omitempty" example:"AU"`              // The region in which your data is hosted (Optional)
}

// SonarqubeCloudIntegrationInput Specifies the input fields used to create and update a SonarQube Cloud integration
type SonarqubeCloudIntegrationInput struct {
	ApiKey          *Nullable[string] `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"`                   // The API Key for the SonarQube Cloud API (Optional)
	Name            *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                       // The name of the integration (Optional)
	OrganizationKey *Nullable[string] `json:"organizationKey,omitempty" yaml:"organizationKey,omitempty" example:"example_value"` // The Organization Key for the SonarQube Cloud organization (Optional)
}

// SonarqubeIntegrationInput Specifies the input fields used to create and update a SonarQube integration
type SonarqubeIntegrationInput struct {
	ApiKey  *Nullable[string] `json:"apiKey,omitempty" yaml:"apiKey,omitempty" example:"example_value"`   // The API Key for the SonarQube API (Optional)
	BaseUrl *Nullable[string] `json:"baseUrl,omitempty" yaml:"baseUrl,omitempty" example:"example_value"` // The base URL for the SonarQube instance (Optional)
	Name    *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`       // The name of the integration (Optional)
}

// SystemInput Specifies the input fields for a system
type SystemInput struct {
	Description *Nullable[string] `json:"description,omitempty" yaml:"description,omitempty" example:"example_value"`           // The description for the system (Optional)
	Name        *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                         // The name for the system (Optional)
	Note        *Nullable[string] `json:"note,omitempty" yaml:"note,omitempty" example:"example_value"`                         // Additional information about the system (Optional)
	OwnerId     *Nullable[ID]     `json:"ownerId,omitempty" yaml:"ownerId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the owner for the system (Optional)
	Parent      *IdentifierInput  `json:"parent,omitempty" yaml:"parent,omitempty"`                                             // The parent domain for the system (Optional)
}

// TagArgs Arguments used to query with a certain tag
type TagArgs struct {
	Key   *Nullable[string] `json:"key,omitempty" yaml:"key,omitempty" example:"example_value"`     // The key of a tag (Optional)
	Value *Nullable[string] `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"` // The value of a tag (Optional)
}

// TagAssignInput Specifies the input fields used to assign tags
type TagAssignInput struct {
	Alias *Nullable[string] `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The alias of the resource that tags will be added to (Optional)
	Id    *Nullable[ID]     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the resource that the tags will be assigned to (Optional)
	Tags  []TagInput        `json:"tags" yaml:"tags" example:"[]"`                                              // The desired tags to assign to the resource (Required)
	Type  *TaggableResource `json:"type,omitempty" yaml:"type,omitempty" example:"Domain"`                      // The type of resource `alias` refers to, if `alias` is provided (Optional Default: Service)
}

// TagCreateInput Specifies the input fields used to create a tag
type TagCreateInput struct {
	Alias *string           `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The alias of the resource that this tag will be added to (Optional)
	Id    *ID               `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the resource that this tag will be added to (Optional)
	Key   string            `json:"key" yaml:"key" example:"example_value"`                                     // The tag's key (Required)
	Type  *TaggableResource `json:"type,omitempty" yaml:"type,omitempty" example:"Domain"`                      // The type of resource `alias` refers to, if `alias` is provided (Optional Default: Service)
	Value string            `json:"value" yaml:"value" example:"example_value"`                                 // The tag's value (Required)
}

// TagDeleteInput Specifies the input fields used to delete a tag
type TagDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the tag to be deleted (Required)
}

// TagInput Specifies the basic input fields used to construct a tag
type TagInput struct {
	Key   string `json:"key" yaml:"key" example:"example_value"`     // The tag's key (Required)
	Value string `json:"value" yaml:"value" example:"example_value"` // The tag's value (Required)
}

// TagRelationshipKeysAssignInput The input for the `tagRelationshipKeysAssign` mutation
type TagRelationshipKeysAssignInput struct {
	BelongsTo    *Nullable[string]   `json:"belongsTo,omitempty" yaml:"belongsTo,omitempty" example:"example_value"` //  (Optional)
	DependencyOf *Nullable[[]string] `json:"dependencyOf,omitempty" yaml:"dependencyOf,omitempty" example:"[]"`      //  (Optional)
	DependsOn    *Nullable[[]string] `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty" example:"[]"`            //  (Optional)
}

// TagUpdateInput Specifies the input fields used to update a tag
type TagUpdateInput struct {
	Id    ID      `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`         // The id of the tag to be updated (Required)
	Key   *string `json:"key,omitempty" yaml:"key,omitempty" example:"example_value"`     // The tag's key (Optional)
	Value *string `json:"value,omitempty" yaml:"value,omitempty" example:"example_value"` // The tag's value (Optional)
}

// TeamCreateInput Specifies the input fields used to create a team
type TeamCreateInput struct {
	Contacts         *[]ContactInput            `json:"contacts,omitempty" yaml:"contacts,omitempty" example:"[]"`                            // The contacts for the team (Optional)
	Group            *IdentifierInput           `json:"group,omitempty" yaml:"group,omitempty"`                                               // The group this team belongs to (Optional)
	ManagerEmail     *Nullable[string]          `json:"managerEmail,omitempty" yaml:"managerEmail,omitempty" example:"example_value"`         // The email of the user who manages the team (Optional)
	Members          *[]TeamMembershipUserInput `json:"members,omitempty" yaml:"members,omitempty" example:"[]"`                              // A set of emails that identify users in OpsLevel (Optional)
	Name             string                     `json:"name" yaml:"name" example:"example_value"`                                             // The team's display name (Required)
	ParentTeam       *IdentifierInput           `json:"parentTeam,omitempty" yaml:"parentTeam,omitempty"`                                     // The parent team (Optional)
	Responsibilities *Nullable[string]          `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty" example:"example_value"` // A description of what the team is responsible for (Optional)
}

// TeamDeleteInput Specifies the input fields used to delete a team
type TeamDeleteInput struct {
	Alias *Nullable[string] `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`             // The alias of the team to be deleted (Optional)
	Id    *Nullable[ID]     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the team to be deleted (Optional)
}

// TeamMembershipCreateInput Input for adding members to a team
type TeamMembershipCreateInput struct {
	Members []TeamMembershipUserInput `json:"members" yaml:"members" example:"[]"`                            // A set of emails that identify users in OpsLevel (Required)
	TeamId  ID                        `json:"teamId" yaml:"teamId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the team to add members (Required)
}

// TeamMembershipDeleteInput Input for removing members from a team
type TeamMembershipDeleteInput struct {
	Members []TeamMembershipUserInput `json:"members" yaml:"members" example:"[]"`                            // A set of emails that identify users in OpsLevel (Required)
	TeamId  ID                        `json:"teamId" yaml:"teamId" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the team to remove members from (Required)
}

// TeamMembershipUserInput Input for specifying members on a team
type TeamMembershipUserInput struct {
	Email *Nullable[string]    `json:"email,omitempty" yaml:"email,omitempty" example:"example_value"` // The user's email (Optional)
	Role  *Nullable[string]    `json:"role,omitempty" yaml:"role,omitempty" example:"example_value"`   // The type of relationship this membership implies (Optional)
	User  *UserIdentifierInput `json:"user,omitempty" yaml:"user,omitempty"`                           // The email address or ID of the user to add to a team (Optional)
}

// TeamPropertyDefinitionInput The input for defining a property
type TeamPropertyDefinitionInput struct {
	Alias        string                    `json:"alias" yaml:"alias" example:"example_value"`                               // The human-friendly, unique identifier for the resource (Required)
	Description  string                    `json:"description" yaml:"description" example:"example_value"`                   // The description of the property definition (Required)
	LockedStatus *PropertyLockedStatusEnum `json:"lockedStatus,omitempty" yaml:"lockedStatus,omitempty" example:"ui_locked"` // Restricts what sources are able to assign values to this property (Optional)
	Name         string                    `json:"name" yaml:"name" example:"example_value"`                                 // The name of the property definition (Required)
	Schema       JSONSchema                `json:"schema" yaml:"schema" example:"SCHEMA_TBD"`                                // The schema of the property definition (Required)
}

// TeamPropertyDefinitionsAssignInput Specifies the input fields used to define properties that apply to teams
type TeamPropertyDefinitionsAssignInput struct {
	Properties []TeamPropertyDefinitionInput `json:"properties" yaml:"properties" example:"[]"` // A list of property definitions (Required)
}

// TeamUpdateInput Specifies the input fields used to update a team
type TeamUpdateInput struct {
	Alias            *string                    `json:"alias,omitempty" yaml:"alias,omitempty" example:"example_value"`                       // The alias of the team to be updated (Optional)
	Group            *IdentifierInput           `json:"group,omitempty" yaml:"group,omitempty"`                                               // The group this team belongs to (Optional)
	Id               *ID                        `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`           // The id of the team to be updated (Optional)
	ManagerEmail     *Nullable[string]          `json:"managerEmail,omitempty" yaml:"managerEmail,omitempty" example:"example_value"`         // The email of the user who manages the team (Optional)
	Members          *[]TeamMembershipUserInput `json:"members,omitempty" yaml:"members,omitempty" example:"[]"`                              // A set of emails that identify users in OpsLevel (Optional)
	Name             *Nullable[string]          `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                         // The team's display name (Optional)
	ParentTeam       *IdentifierInput           `json:"parentTeam,omitempty" yaml:"parentTeam,omitempty"`                                     // The parent team (Optional)
	Responsibilities *Nullable[string]          `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty" example:"example_value"` // A description of what the team is responsible for (Optional)
}

// ToolCreateInput Specifies the input fields used to create a tool
type ToolCreateInput struct {
	Category     ToolCategory      `json:"category" yaml:"category" example:"admin"`                                                 // The category that the tool belongs to (Required)
	DisplayName  string            `json:"displayName" yaml:"displayName" example:"example_value"`                                   // The display name of the tool (Required)
	Environment  *Nullable[string] `json:"environment,omitempty" yaml:"environment,omitempty" example:"example_value"`               // The environment that the tool belongs to (Optional)
	ServiceAlias *string           `json:"serviceAlias,omitempty" yaml:"serviceAlias,omitempty" example:"example_value"`             // The alias of the service the tool will be assigned to (Optional)
	ServiceId    *ID               `json:"serviceId,omitempty" yaml:"serviceId,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the service the tool will be assigned to (Optional)
	Url          string            `json:"url" yaml:"url" example:"example_value"`                                                   // The URL of the tool (Required)
}

// ToolDeleteInput Specifies the input fields used to delete a tool
type ToolDeleteInput struct {
	Id ID `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The id of the tool to be deleted (Required)
}

// ToolUpdateInput Specifies the input fields used to update a tool
type ToolUpdateInput struct {
	Category    *ToolCategory     `json:"category,omitempty" yaml:"category,omitempty" example:"admin"`               // The category that the tool belongs to (Optional)
	DisplayName *Nullable[string] `json:"displayName,omitempty" yaml:"displayName,omitempty" example:"example_value"` // The display name of the tool (Optional)
	Environment *Nullable[string] `json:"environment,omitempty" yaml:"environment,omitempty" example:"example_value"` // The environment that the tool belongs to (Optional)
	Id          ID                `json:"id" yaml:"id" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"`                     // The id of the tool to be updated (Required)
	Url         *Nullable[string] `json:"url,omitempty" yaml:"url,omitempty" example:"example_value"`                 // The URL of the tool (Optional)
}

// UserIdentifierInput Specifies the input fields used to identify a user. Exactly one field should be provided
type UserIdentifierInput struct {
	Email *Nullable[string] `json:"email,omitempty" yaml:"email,omitempty" example:"example_value"`             // The email address of the user (Optional)
	Id    *Nullable[ID]     `json:"id,omitempty" yaml:"id,omitempty" example:"Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk"` // The ID of the user (Optional)
}

// UserInput Specifies the input fields used to create and update a user
type UserInput struct {
	Name             *Nullable[string] `json:"name,omitempty" yaml:"name,omitempty" example:"example_value"`                 // The name of the user (Optional)
	Role             *UserRole         `json:"role,omitempty" yaml:"role,omitempty" example:"admin"`                         // The access role (e.g. user vs admin) of the user (Optional)
	SkipWelcomeEmail *Nullable[bool]   `json:"skipWelcomeEmail,omitempty" yaml:"skipWelcomeEmail,omitempty" example:"false"` // Don't send an email welcoming the user to OpsLevel (Optional)
}

// UsersFilterInput The input for filtering users
type UsersFilterInput struct {
	Arg  *Nullable[string] `json:"arg,omitempty" yaml:"arg,omitempty" example:"example_value"`    // Value to be filtered (Optional)
	Key  UsersFilterEnum   `json:"key" yaml:"key" example:"deactivated_at"`                       // Field to be filtered (Required)
	Type *BasicTypeEnum    `json:"type,omitempty" yaml:"type,omitempty" example:"does_not_equal"` // The operation applied to value on the field (Optional Default: equals)
}

// UsersInviteInput Specifies the input fields used in the `usersInvite` mutation
type UsersInviteInput struct {
	Scope *UsersInviteScopeEnum  `json:"scope,omitempty" yaml:"scope,omitempty" example:"pending"` // A classification of users to invite (Optional)
	Users *[]UserIdentifierInput `json:"users,omitempty" yaml:"users,omitempty" example:"[]"`      // A list of individual users to invite (Optional)
}
