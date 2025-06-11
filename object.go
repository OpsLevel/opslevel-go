// Code generated; DO NOT EDIT.
package opslevel

import "github.com/relvacode/iso8601"

// AlertSource An alert source that is currently integrated and belongs to the account
type AlertSource struct {
	Description string              // The description of the alert source (Optional)
	ExternalId  string              // The external id of the alert (Required)
	Id          ID                  // The id of the alert source (Required)
	Integration IntegrationId       // The integration of the alert source (Optional)
	Metadata    string              // The metadata of the alert source (Optional)
	Name        string              // The name of the alert source (Required)
	Type        AlertSourceTypeEnum // The type of the alert (Required)
	Url         string              // The url to the alert source (Optional)
}

// AlertSourceService An alert source that is connected with a service
type AlertSourceService struct {
	AlertSource AlertSource               // The alert source that is mapped to a service (Required)
	Id          ID                        // id of the alert_source_service mapping (Required)
	Service     ServiceId                 // The service the alert source maps to (Required)
	Status      AlertSourceStatusTypeEnum // The status of the alert source (Required)
}

// ApprovalConfig Config for approval
type ApprovalConfig struct {
	ApprovalRequired bool     // Flag indicating approval is required (Required)
	Teams            []TeamId // Teams that can approve (Required)
	Users            []UserId // Users that can approve (Required)
}

// ApprovalResult The result details of an approvable resource being approved or denied
type ApprovalResult struct {
	Actor     *Approver            // The actor that performed the approval/denial (Optional)
	Comment   string               // The comment associated with the approval/denial (Optional)
	DecidedAt iso8601.Time         // The time the resource was approved (Required)
	Status    ApprovalDecisionEnum // The current decision status of the approval process (Required)
}

// AzureDevopsPermissionError
type AzureDevopsPermissionError struct {
	Name        string   // The name of the object that the error was encountered on (Required)
	Permissions []string // The permissions that are missing (Optional)
	Type        string   // The type of the object that the error was encountered on (Required)
}

// Campaign A campaign is a fixed time initiative that allows you to attach checks to and progress through towards completion
type Campaign struct {
	CheckStats      Stats              // A summary of check results on the campaign (Optional)
	EndedDate       iso8601.Time       // The date the campaign ended (Optional)
	Filter          FilterId           // The filter that the campaign belongs to (Optional)
	HtmlUrl         string             // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
	Id              ID                 // The id of the campaign (Required)
	Name            string             // The name of the campaign (Required)
	Owner           TeamId             // The team that owns the campaign (Optional)
	ProjectBrief    string             // The project brief of the campaign (Optional)
	RawProjectBrief string             // The raw unsanitized project brief of the campaign (Optional)
	Reminder        CampaignReminder   // Configuration of an optional campaign reminder (Optional)
	ServiceStats    Stats              // A summary of services that completed the campaign (Optional)
	StartDate       iso8601.Time       // The date the campaign will start (Optional)
	Status          CampaignStatusEnum // The status of the campaign (Required)
	TargetDate      iso8601.Time       // The target date the campaign should end (Optional)
}

// CampaignReminder Configuration of an optional campaign reminder
type CampaignReminder struct {
	Channels            []CampaignReminderChannelEnum     // The communication channels through which the reminder will be delivered (Required)
	DaysOfWeek          []DayOfWeekEnum                   // A list of weekdays on which the reminders will be delivered. Only available with weekly frequency (Optional)
	DefaultSlackChannel string                            // The name of the Slack channel that will be notified if a team doesn't have a default Slack contact (Optional)
	Frequency           int                               // The interval at which reminders will be delivered (Required)
	FrequencyUnit       CampaignReminderFrequencyUnitEnum // The time unit of the value in the 'frequency' field (Required)
	Message             string                            // The message that will be delivered as the reminder (Optional)
	NextOccurrence      iso8601.Time                      // The point in time at which the next reminder will be delivered based on the current configuration (Optional)
	TimeOfDay           string                            // The time of day at which the reminder will be delivered. Format: "HH:MM" (Required)
	Timezone            string                            // The timezone at which the timeOfDay field is evaluated (in IANA format (e.g. "America/Chicago")) (Required)
}

// CampaignSendReminderOutcomeTeams Summarizes list of teams returned from attempt to send reminders for their failed campaigns
type CampaignSendReminderOutcomeTeams struct {
	ReminderType CampaignReminderTypeEnum // The reminder type linked to the attempt at notifying the listed teams (Required)
	TeamIds      []ID                     // List of team_ids in this group of teams (Optional)
	TotalCount   int                      // Count of number of teams listed (Required)
}

// CampaignTeam Team from the campaign context
type CampaignTeam struct {
	Team TeamId // The team impacted by the campaign (Required)
}

// Category A category is used to group related checks in a rubric
type Category struct {
	Description string // The description of the category (Optional)
	Id          ID     // The unique identifier for the category (Required)
	Name        string // The display name of the category (Required)
}

// CategoryLevel The level of a specific category
type CategoryLevel struct {
	Category Category // A category is used to group related checks in a rubric (Required)
	Level    Level    // A performance rating that is used to grade your services against (Optional)
}

// CheckResult The result for a given Check
type CheckResult struct {
	Check        CheckId      // The check of check result (Required)
	LastUpdated  iso8601.Time // The time the check most recently ran (Required)
	Message      string       // The check message (Required)
	Service      ServiceId    // The service of check result (Optional)
	ServiceAlias string       // The alias for the service (Optional)
	Status       CheckStatus  // The check status (Required)
}

// CheckResultsByLevel The check results grouped by level
type CheckResultsByLevel struct {
	Items CheckResultsConnection // A list of check results by level (Optional)
	Level Level                  // The check result level (Optional)
}

// CheckStats Check stats shows a summary of check results
type CheckStats struct {
	PassingChecksPercentage float64 // The percentage of checks that are passing for the resource (Required)
	TotalChecks             int     // The number of existing checks for the resource (Required)
	TotalPassingChecks      int     // The number of checks that are passing for the resource (Required)
}

// CommonVulnerabilityEnumeration A category system for hardware and software weaknesses
type CommonVulnerabilityEnumeration struct {
	Identifier string // The identifer of this item in the CVE system (Required)
	Url        string // The url for this item in the CVE system (Optional)
}

// CommonWeaknessEnumeration A category system for hardware and software weaknesses
type CommonWeaknessEnumeration struct {
	Identifier string // The identifer of this item in the CWE system (Required)
	Url        string // The url for this item in the CWE system (Optional)
}

// ComponentTypeId Information about a particular component type
type ComponentTypeId struct {
	Id      ID       // The id of the component type.
	Aliases []string // A list of human-friendly, unique identifiers of the component type.
}

// ComponentType Information about a particular component type
type ComponentType struct {
	ComponentTypeId
	Description string                        // The description of the component type (Optional)
	Href        string                        // The relative path to link to the component type (Required)
	Icon        ComponentTypeIcon             // The icon associated with the component type (Required)
	IsDefault   bool                          // Whether or not the component type is the default (Required)
	Name        string                        // The name of the component type (Required)
	Timestamps  Timestamps                    // When the component type was created and updated (Required)
	Properties  *PropertyDefinitionConnection `graphql:"-"`
}

// ComponentTypeIcon The icon for a component type
type ComponentTypeIcon struct {
	Color string                // The color, represented as a hexcode, for the icon (Optional)
	Name  ComponentTypeIconEnum // The name of the icon in Phosphor icons for Vue, e.g. `PhBird`. See https://phosphoricons.com/ for a full list (Optional)
}

// ConfigError An error that occurred when syncing an opslevel.yml file
type ConfigError struct {
	Message        string // A description of the error (Optional)
	SourceFilename string // The file name where the error was found (Required)
}

// ConfigFile An OpsLevel config as code definition
type ConfigFile struct {
	OwnerType string // The relation for which the config was returned (Required)
	Yaml      string // The OpsLevel config in yaml format (Required)
}

// Contact A method of contact for a team
type Contact struct {
	Address     string      // The contact address. Examples: support@company.com for type `email`, https://opslevel.com for type `web` (Required)
	DisplayName string      // The name shown in the UI for the contact (Optional)
	DisplayType string      // The type shown in the UI for the contact (Optional)
	ExternalId  string      // The remote identifier of the contact method (Optional)
	Id          ID          // The unique identifier for the contact (Required)
	IsDefault   bool        // Indicates if this address is a team's default for the given type (Optional)
	Type        ContactType // The method of contact [email, slack, slack_handle, web, microsoft_teams] (Required)
}

// CustomActionsTemplate Template of a custom action
type CustomActionsTemplate struct {
	Action            CustomActionsTemplatesAction            // The template's action (Required)
	Metadata          CustomActionsTemplatesMetadata          // The template's metadata (Required)
	TriggerDefinition CustomActionsTemplatesTriggerDefinition // The template's trigger definition (Required)
}

// CustomActionsTemplatesAction The action of a custom action template
type CustomActionsTemplatesAction struct {
	Description    string                      // A description of what the action should accomplish (Optional)
	Headers        JSON                        `scalar:"true"` // The headers sent along with the webhook, if any (Optional)
	HttpMethod     CustomActionsHttpMethodEnum // The HTTP Method used to call the webhook action (Required)
	LiquidTemplate string                      // The liquid template used to generate the data sent to the external action (Optional)
	Name           string                      // The name of the external action (Required)
	Url            string                      // The URL of the webhook action (Required)
}

// CustomActionsTemplatesMetadata The metadata about the custom action template
type CustomActionsTemplatesMetadata struct {
	Categories  []string // The categories for the custom action template (Required)
	Description string   // The description of the custom action template (Optional)
	Icon        string   // The icon for the custom action template (Optional)
	Name        string   // The name of the custom action template (Required)
}

// CustomActionsTemplatesTriggerDefinition The definition of a potential trigger for a template custom action
type CustomActionsTemplatesTriggerDefinition struct {
	AccessControl          CustomActionsTriggerDefinitionAccessControlEnum // The set of users that should be able to use the trigger definition (Required)
	Description            string                                          // The description of what the trigger definition will do, supports Markdown (Optional)
	ManualInputsDefinition string                                          // The YAML definition of any custom inputs for this trigger definition (Optional)
	Name                   string                                          // The name of the trigger definition (Required)
	Published              bool                                            // The published state of the action; true if the definition is ready for use; false if it is a draft (Required)
	ResponseTemplate       string                                          // The liquid template used to parse the response from the External Action (Optional)
}

// CustomActionsTriggerDefinition The definition of a potential trigger for a custom action
type CustomActionsTriggerDefinition struct {
	AccessControl          CustomActionsTriggerDefinitionAccessControlEnum // The set of users that should be able to use the trigger definition (Required)
	Action                 CustomActionsId                                 // The action that would be triggered (Required)
	Aliases                []string                                        // Any aliases for this trigger definition (Required)
	ApprovalConfig         ApprovalConfig                                  // Configuration defining conditions of approval if it is required (Required)
	Description            string                                          // The description of what the trigger definition will do, supports Markdown (Optional)
	EntityType             CustomActionsEntityTypeEnum                     // The entity type associated with this trigger definition (Required)
	Filter                 FilterId                                        // A filter defining which services this trigger definition applies to, if present (Optional)
	Id                     ID                                              // The ID of the trigger definition (Required)
	ManualInputsDefinition string                                          // The YAML definition of any custom inputs for this trigger definition (Optional)
	Name                   string                                          // The name of the trigger definition (Required)
	Owner                  TeamId                                          // The owner of the trigger definition (Optional)
	Published              bool                                            // The published state of the action; true if the definition is ready for use; false if it is a draft (Required)
	ResponseTemplate       string                                          // The liquid template used to parse the response from the External Action (Optional)
	Timestamps             Timestamps                                      // Relevant timestamps (Required)
}

// CustomActionsWebhookAction An external webhook action to be triggered by a custom action
type CustomActionsWebhookAction struct {
	Aliases        []string                    // Any aliases for this external action (Required)
	Async          bool                        // Whether the action expects an additional, asynchronous response upon completion (Required)
	Description    string                      // A description of what the action should accomplish (Optional)
	Headers        JSON                        `scalar:"true"` // The headers sent along with the webhook, if any (Optional)
	HttpMethod     CustomActionsHttpMethodEnum // The HTTP Method used to call the webhook action (Required)
	Id             ID                          // The ID of the external action (Required)
	LiquidTemplate string                      // The liquid template used to generate the data sent to the external action (Optional)
	Name           string                      // The name of the external action (Required)
	WebhookUrl     string                      // The URL of the webhook action (Required)
}

// Deploy An event sent via webhook to track deploys
type Deploy struct {
	AssociatedUser      UserId           // The associated OpsLevel user for the deploy (Optional)
	Author              string           // The author of the deploy (Optional)
	CommitAuthorEmail   string           // The email of the commit (Optional)
	CommitAuthorName    string           // The author of the commit (Optional)
	CommitAuthoringDate iso8601.Time     // The time the commit was authored (Optional)
	CommitBranch        string           // The branch the commit took place on (Optional)
	CommitMessage       string           // The commit message associated with the deploy (Optional)
	CommitSha           string           // The sha associated with the commit of the deploy (Optional)
	CommittedAt         iso8601.Time     // The time the commit happened (Optional)
	CommitterEmail      string           // The email of the person who created the commit (Optional)
	CommitterName       string           // The name of the person who created the commit (Optional)
	DedupId             string           // The deduplication ID provided to prevent duplicate deploys (Optional)
	DeployNumber        string           // An identifier to keep track of the version of the deploy (Optional)
	DeployStatus        DeployStatusEnum // The normalized status of the deploy. This is derived from the status field (Optional)
	DeployUrl           string           // The url the where the deployment can be found (Optional)
	DeployedAt          iso8601.Time     // The time the deployment happened (Optional)
	DeployerEmail       string           // The email of who is responsible for the deployment (Optional)
	DeployerId          string           // An external id of who deployed (Optional)
	DeployerName        string           // The name of who is responsible for the deployment (Optional)
	Description         string           // The given description of the deploy (Required)
	Duration            int              // The duration of the deploy (Optional)
	Environment         string           // The environment in which the deployment happened in (Optional)
	Id                  ID               // The id of the deploy (Required)
	ProviderName        string           // The integration name of the deploy (Optional)
	ProviderType        string           // The integration type used the deploy (Optional)
	ProviderUrl         string           // The url to the deploy integration (Optional)
	Service             ServiceId        // The service object the deploy is attached to (Optional)
	ServiceAlias        string           // The alias used to associated this deploy to its service (Required)
	ServiceId           string           // The id the deploy is associated to (Optional)
	StartedAt           iso8601.Time     // The time the deploy started (Optional)
	Status              string           // The deployment status (Optional)
}

// DomainId A collection of related Systems
type DomainId struct {
	Id      ID       // The identifier of the object.
	Aliases []string // All of the aliases attached to the resource.
}

// Domain A collection of related Systems
type Domain struct {
	DomainId
	Description    string      // The description of the Domain (Optional)
	HtmlUrl        string      // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
	ManagedAliases []string    // A list of aliases that can be set by users. The unique identifier for the resource is omitted (Required)
	Name           string      // The name of the object (Required)
	Note           string      // Additional information about the domain (Optional)
	Owner          EntityOwner // The owner of the object (Optional)
}

// Error The input error of a mutation
type Error struct {
	Message string   // The error message (Required)
	Path    []string // The path to the input field with an error (Required)
}

// FilterId A filter is used to select which services will have checks applied. It can also be used to filter services in reports
type FilterId struct {
	Id   ID     // The unique identifier for the filter.
	Name string // The display name of the filter.
}

// Filter A filter is used to select which services will have checks applied. It can also be used to filter services in reports
type Filter struct {
	FilterId
	Connective ConnectiveEnum    // The logical operator to be used in conjunction with predicates (Optional)
	HtmlUrl    string            // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
	Predicates []FilterPredicate // The predicates used to select services (Required)
}

// FilterPredicate A condition used to select services
type FilterPredicate struct {
	CaseSensitive *bool             // Option for determining whether to compare strings case-sensitively (Optional)
	Key           PredicateKeyEnum  // The key of the condition (Required)
	KeyData       string            // Additional data used in the condition (Optional)
	Type          PredicateTypeEnum // Type of operation to be used in the condition (Required)
	Value         string            // The value of the condition (Optional)
}

// GoogleCloudProject
type GoogleCloudProject struct {
	Id   string // The ID of the Google Cloud project (Required)
	Name string // The name of the Google Cloud project (Required)
	Url  string // The URL to the Google Cloud project (Required)
}

// InfrastructureResourceProviderData Data about the provider the infrastructure resource is from
type InfrastructureResourceProviderData struct {
	AccountName  string // The account name of the provider (Required)
	ExternalUrl  string // The external URL of the infrastructure resource in its provider (Optional)
	ProviderName string // The name of the provider (e.g. AWS, GCP, Azure) (Optional)
}

// Language A language that can be assigned to a repository
type Language struct {
	Name  string  // The name of the language (Required)
	Usage float64 // The percentage of the code written in that language (Required)
}

// Level A performance rating that is used to grade your services against
type Level struct {
	Alias       string    // The human-friendly, unique identifier for the level (Optional)
	Checks      []CheckId // The checks that belong to the level (Optional)
	Description string    // A brief description of the level (Optional)
	Id          ID        // The unique identifier for the level (Required)
	Index       int       // The numerical representation of the level (highest is better) (Optional)
	Name        string    // The display name of the level (Optional)
}

// LevelCount The total number of services in each level
type LevelCount struct {
	Level        Level // A performance rating that is used to grade your services against (Required)
	ServiceCount int   // The number of services (Required)
}

// Lifecycle A lifecycle represents the current development stage of a service
type Lifecycle struct {
	Alias       string // The human-friendly, unique identifier for the lifecycle (Optional)
	Description string // The lifecycle's description (Optional)
	Id          ID     // The unique identifier for the lifecycle (Required)
	Index       int    // The numerical representation of the lifecycle (Optional)
	Name        string // The lifecycle's display name (Optional)
}

// ManualCheckFrequency
type ManualCheckFrequency struct {
	FrequencyTimeScale FrequencyTimeScale // The time scale type for the frequency (Required)
	FrequencyValue     int                // The value to be used together with the frequency scale (Required)
	StartingDate       iso8601.Time       // The date that the check will start to evaluate (Required)
}

// Predicate A condition used to select services
type Predicate struct {
	Type  PredicateTypeEnum // Type of operation to be used in the condition (Required)
	Value string            // The value of the condition (Optional)
}

// RelationshipDefinitionMetadata The metadata of the relationship
type RelationshipDefinitionMetadata struct {
	AllowedTypes []string // The aliases of which types this relationship can target. Valid values include any component type alias on your account, or `team` (Required)
	MaxItems     int      // The maximum number of records this relationship can associate to the component type. Defaults to null (no maximum) (Optional)
	MinItems     int      // The minimum number of records this relationship must associate to the component type. Defaults to 0 (optional) (Optional)
}

// RelationshipDefinitionType A dynamic definition for a relationship between one catalog entity to another
type RelationshipDefinitionType struct {
	Alias         string                         // The programmatic alias that can be used to reference the relationship in OpsLevel tooling (Required)
	ComponentType ComponentTypeId                // The component type that the relationship belongs to (Required)
	Description   string                         // The long-form descripion of what the relationship represents (Optional)
	Id            ID                             // The ID of the relationship definition (Required)
	Metadata      RelationshipDefinitionMetadata // JSON data that defines rules for how the relationship should be validated internally (Required)
	Name          string                         // The human-readable name for a relationship (Required)
}

// RelationshipNode The relationship between two resources. A pair of source and destination resources
type RelationshipNode struct {
	Destination RelationshipResource // The catalog item that a relationship points to (Required)
	Id          ID                   // The ID of the relationship (Required)
	Source      RelationshipResource // The catalog item that a relationship stems from (Required)
}

// RelationshipType The type specifying a relationship between two resources
type RelationshipType struct {
	Id     ID                   // The ID of the relationship (Required)
	Source RelationshipResource // The resource that is the source of the relationship (Required)
	Target RelationshipResource // The resource that is the target of the relationship (Required)
	Type   RelationshipTypeEnum // The type of the relationship between source and target (Required)
}

// RepositoryPath The repository path used for this service
type RepositoryPath struct {
	Href string // The deep link to the repository path where the linked service's code exists (Required)
	Path string // The path where the linked service's code exists, relative to the root of the repository (Required)
}

// RubricReport The check result information for the service's rubric
type RubricReport struct {
	CategoryLevel Level               // The level of a specific category (Optional)
	CheckResults  ServiceCheckResults // The service check results (Optional)
	Level         Level               // The overall level of the service (Optional)
}

// SBOMGenerationConfiguration The configuration details that explain whether SBOM generation is allowed for the repository
type SBOMGenerationConfiguration struct {
	DisabledReason   RepositorySBOMGenerationDisabledReasonEnum // A brief explanation of why SBOM autogeneration is disabled (Optional)
	Enabled          bool                                       // Whether SBOM autogeneration is enabled through all associated configuration objects (Required)
	NextGenerationAt iso8601.Time                               // The approximate time at which a new software bill of material will be generated for this repository (Optional)
	State            RepositorySBOMGenerationConfigEnum         // The configuration option set by the current object (Optional)
}

// ScorecardId A scorecard
type ScorecardId struct {
	Id      ID       // A reference to the scorecard.
	Aliases []string // Aliases of the scorecard.
}

// Scorecard A scorecard
type Scorecard struct {
	ScorecardId
	AffectsOverallServiceLevels bool                    // Specifies whether the checks on this scorecard affect services' overall maturity level (Required)
	Description                 string                  // Description of the scorecard (Optional)
	Filter                      Filter                  // Filter used by the scorecard to restrict services (Optional)
	Href                        string                  // The hypertext reference (link) to the UI showing this scorecard (Required)
	Name                        string                  // Name of the scorecard (Required)
	Owner                       EntityOwner             // The owner of this scorecard. Can currently either be a team or a group (Optional)
	PassingChecks               int                     // The number of checks that are passing on this scorecard. A check executed against two services counts as two (Required)
	PassingChecksPercentage     float64                 // The percentage of checks that are passing on this scorecard. A check executed against two services counts as two (Required)
	ServiceCount                int                     // The number of services covered by this scorecard (Required)
	ServicesReport              ScorecardServicesReport // Service stats regarding this scorecard (Optional)
	Slug                        string                  // Slug of the scorecard (Required)
	TotalChecks                 int                     // The number of checks that are performed on this scorecard. A check executed against two services counts as two (Required)
}

// ScorecardServicesReport Service stats regarding this scorecard
type ScorecardServicesReport struct {
	LevelCounts []LevelCount // Services per level regarding this scorecard (Required)
}

// ScorecardStats Service maturity information about a scorecard
type ScorecardStats struct {
	CheckResults ServiceCheckResults // The service check results (Optional)
	Scorecard    Scorecard           // The scorecard (Optional)
}

// Secret A sensitive value
type Secret struct {
	Alias      string     // A human reference for the secret (Required)
	Id         ID         // A reference for the secret (Required)
	Owner      TeamId     // The owner of this secret (Optional)
	Timestamps Timestamps // Relevant timestamps (Required)
}

// ServiceCheckResults The service check results
type ServiceCheckResults struct {
	ByLevel   CheckResultsByLevelConnection // The list of service check results grouped by level (Optional)
	NextLevel ServiceNextLevel              // The next level for this service to achieve and the associated checks (Optional)
}

// ServiceLevelNotifications
type ServiceLevelNotifications struct {
	SlackNotificationEnabled bool // Whether slack notifications on service level changes are enabled on your account (Required)
}

// ServiceMaturityReport The health report for this service in terms of its levels and checks
type ServiceMaturityReport struct {
	CategoryBreakdown  []CategoryLevel // The level of each category for this service (Required)
	LatestCheckResults []CheckResult   // The latest check results for this service across the given checks (Optional)
	OverallLevel       Level           // The overall level for this service (Required)
}

// ServiceNextLevel The next level for a service to achieve and the associated checks
type ServiceNextLevel struct {
	Level Level // The next level for a service to achieve (Optional)
}

// ServiceRepository A record of the connection between a service and a repository
type ServiceRepository struct {
	BaseDirectory string       // The directory in the repository where service information exists, including the opslevel.yml file. This path is always returned without leading and trailing slashes (Optional)
	DisplayName   string       // The name displayed in the UI for the service repository (Optional)
	Id            ID           // ID of the service repository (Required)
	Repository    RepositoryId // The repository that is part of this connection (Required)
	Service       ServiceId    // The service that is part of this connection (Required)
}

// ServiceStats The summary of check results for this service
type ServiceStats struct {
	Rubric RubricReport // The check result information for the service's rubric (Required)
}

// Stats An object that contains statistics
type Stats struct {
	Total           int // How many there are (Required)
	TotalSuccessful int // How many are successfully passing (Required)
}

// SystemId A collection of related Services
type SystemId struct {
	Id      ID       // The identifier of the object.
	Aliases []string // All of the aliases attached to the resource.
}

// System A collection of related Services
type System struct {
	SystemId
	Description    string      // The description of the System (Optional)
	HtmlUrl        string      // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
	ManagedAliases []string    // A list of aliases that can be set by users. The unique identifier for the resource is omitted (Required)
	Name           string      // The name of the object (Required)
	Note           string      // Additional information about the system (Optional)
	Owner          EntityOwner // The owner of the object (Optional)
	Parent         Domain      // Parent domain of the System (Optional)
}

// Tag An arbitrary key-value pair associated with a resource
type Tag struct {
	Id    ID     // The unique identifier for the tag (Required)
	Key   string // The tag's key (Required)
	Value string // The tag's value (Required)
}

// TagRelationshipKeys Returns the keys that set relationships when imported from AWS
type TagRelationshipKeys struct {
	BelongsTo    string   // The tag key that will create `belongs_to` relationships (Required)
	DependencyOf []string // The tag keys that will create `dependency_of` relationships (Required)
	DependsOn    []string // The tag keys that will create `depends_on` relationships (Required)
}

// TeamMembership
type TeamMembership struct {
	Role string // Role of the user on the Team (Optional)
	Team TeamId // Team for the membership (Required)
	User UserId // User for the membership (Required)
}

// TeamPropertyDefinition The definition of a property
type TeamPropertyDefinition struct {
	Alias          string                            // The human-friendly, unique identifier of the property definition (Required)
	Description    string                            // The description of the property definition (Required)
	DisplaySubtype PropertyDefinitionDisplayTypeEnum // The secondary inferred type of the schema (Optional)
	DisplayType    PropertyDefinitionDisplayTypeEnum // The primary inferred type of the schema (Required)
	Id             ID                                // The id of the property definition (Required)
	LockedStatus   PropertyLockedStatusEnum          // Restricts what sources are able to assign values to this property (Required)
	Name           string                            // The name of the property definition (Required)
	Schema         JSONSchema                        `scalar:"true"` // The schema of the property definition (Required)
}

// Tier A tier measures how critical or important a service is to your business
type Tier struct {
	Alias       string // The human-friendly, unique identifier for the tier (Optional)
	Description string // A brief description of the tier (Optional)
	Id          ID     // The unique identifier for the tier (Required)
	Index       int    // The numerical representation of the tier (Optional)
	Name        string // The display name of the tier (Optional)
}

// Timestamps Relevant timestamps
type Timestamps struct {
	CreatedAt iso8601.Time // The time at which the entity was created (Required)
	UpdatedAt iso8601.Time // The time at which the entity was most recently updated (Required)
}

// Tool A tool is used to support the operations of a service
type Tool struct {
	Category      ToolCategory // The category that the tool belongs to (Optional)
	CategoryAlias string       // The human-friendly, unique identifier for the tool's category (Optional)
	DisplayName   string       // The display name of the tool (Optional)
	Environment   string       // The environment that the tool belongs to (Optional)
	Id            ID           // The unique identifier for the tool (Required)
	Service       ServiceId    // The service that is associated to the tool (Required)
	Url           string       // The URL of the tool (Required)
}

// UserId A user is someone who belongs to an organization
type UserId struct {
	Id    ID     // The unique identifier for the user.
	Email string // The user's email.
	Name  string // The user's full name (Required)
}

// User A user is someone who belongs to an organization
type User struct {
	UserId
	HtmlUrl       string            // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
	ProvisionedBy ProvisionedByEnum // What provisioned this user (Optional)
	Role          UserRole          // The user's assigned role (Optional)
}

// Warning The warnings of the mutation
type Warning struct {
	Message string // The warning message (Required)
}
