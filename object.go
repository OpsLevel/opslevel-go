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

// AzureDevopsPermissionError
type AzureDevopsPermissionError struct {
	Name        string   // The name of the object that the error was encountered on (Required)
	Permissions []string // The permissions that are missing (Optional)
	Type        string   // The type of the object that the error was encountered on (Required)
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
	Check        Check        // The check of check result (Required)
	LastUpdated  iso8601.Time // The time the check most recently ran (Required)
	Message      string       // The check message (Required)
	Service      ServiceId    // The service of check result (Optional)
	ServiceAlias string       // The alias for the service (Optional)
	Status       CheckStatus  // The check status (Required)
}

// CheckStats Check stats shows a summary of check results
type CheckStats struct {
	TotalChecks        int // The number of existing checks for the resource (Required)
	TotalPassingChecks int // The number of checks that are passing for the resource (Required)
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
	Id    ID     // The id of the component type.
	Alias string // The human-friendly, unique identifier of the component type.
}

// ComponentType Information about a particular component type
type ComponentType struct {
	ComponentTypeId
	Description string            // The description of the component type (Optional)
	Href        string            // The relative path to link to the component type (Required)
	Icon        ComponentTypeIcon // The icon associated with the component type (Required)
	IsDefault   bool              // Whether or not the component type is the default (Required)
	Name        string            // The name of the component type (Required)
	Timestamps  Timestamps        // When the component type was created and updated (Required)
}

// ComponentTypeIcon The icon for a component type
type ComponentTypeIcon struct {
	Color string // The color, represented as a hexcode, for the icon (Optional)
	Name  string // The name of the icon in Phosphor icons for Vue, e.g. `PhBird`. See https://phosphoricons.com/ for a full list (Optional)
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
	AssociatedUser      UserId       // The associated OpsLevel user for the deploy (Optional)
	Author              string       // The author of the deploy (Optional)
	CommitAuthorEmail   string       // The email of the commit (Optional)
	CommitAuthorName    string       // The author of the commit (Optional)
	CommitAuthoringDate iso8601.Time // The time the commit was authored (Optional)
	CommitBranch        string       // The branch the commit took place on (Optional)
	CommitMessage       string       // The commit message associated with the deploy (Optional)
	CommitSha           string       // The sha associated with the commit of the deploy (Optional)
	CommittedAt         iso8601.Time // The time the commit happened (Optional)
	CommitterEmail      string       // The email of the person who created the commit (Optional)
	CommitterName       string       // The name of the person who created the commit (Optional)
	DedupId             string       // The deduplication ID provided to prevent duplicate deploys (Optional)
	DeployNumber        string       // An identifier to keep track of the version of the deploy (Optional)
	DeployUrl           string       // The url the where the deployment can be found (Optional)
	DeployedAt          iso8601.Time // The time the deployment happened (Optional)
	DeployerEmail       string       // The email of who is responsible for the deployment (Optional)
	DeployerId          string       // An external id of who deployed (Optional)
	DeployerName        string       // The name of who is responsible for the deployment (Optional)
	Description         string       // The given description of the deploy (Required)
	Environment         string       // The environment in which the deployment happened in (Optional)
	Id                  ID           // The id of the deploy (Required)
	ProviderName        string       // The integration name of the deploy (Optional)
	ProviderType        string       // The integration type used the deploy (Optional)
	ProviderUrl         string       // The url to the deploy integration (Optional)
	Service             ServiceId    // The service object the deploy is attached to (Optional)
	ServiceAlias        string       // The alias used to associated this deploy to its service (Required)
	ServiceId           string       // The id the deploy is associated to (Optional)
	Status              string       // The deployment status (Optional)
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
	Alias       string // The human-friendly, unique identifier for the level (Optional)
	Description string // A brief description of the level (Optional)
	Id          ID     // The unique identifier for the level (Required)
	Index       int    // The numerical representation of the level (highest is better) (Optional)
	Name        string // The display name of the level (Optional)
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

//
//// RepositoryId A repository contains code that pertains to a service
//type RepositoryId struct {
//	Id      ID       // The unique identifier for the repository.
//	Aliases []string //
//}
//
//// Repository A repository contains code that pertains to a service
//type Repository struct {
//	RepositoryId
//	ArchivedAt         iso8601.Time             // The date the repository was archived (Optional)
//	ConfigErrors       []ConfigError            // List of errors that occurred while syncing an opslevel.yml file (Optional)
//	CreatedOn          iso8601.Time             // The date the repository was created (Optional)
//	DefaultAlias       string                   // The default human-friendly identifier assigned to the repository (Optional)
//	DefaultBranch      string                   // The default branch from the repository's settings (Optional)
//	Description        string                   // A brief description of the repository (Optional)
//	Forked             bool                     // Indicates if the repository is forked (Optional)
//	HtmlUrl            string                   // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
//	Languages          []Language               // A list of languages used in the repository (Required)
//	LastOwnerChangedAt iso8601.Time             // The date and time of the most recent ownership change (Optional)
//	Locked             bool                     // Indicates if the repository is locked by an opslevel.yml (Required)
//	Name               string                   // The name of the repository (Required)
//	Organization       string                   // The organization to which the repository belongs to (Optional)
//	Owner              TeamId                   // The team that owns the repository (Optional)
//	OwnerAlias         string                   // The human-friendly identifier for the owning team (whether such a team exists or not) (Optional)
//	Private            bool                     // Indicates if the repository is private (Optional)
//	RepoKey            string                   // The repository's unique key from its management platform (Required)
//	Tier               Tier                     // The software tier that the repository belongs to (Optional)
//	Type               string                   // The management platform of the repository (Required)
//	Url                string                   // The URL of the repository (Optional)
//	Visibility         RepositoryVisibilityEnum // The level of visibility of the repository (Optional)
//	Visible            bool                     // Indicates if the repository is visible (Optional)
//}
//
//// RepositoryPath The repository path used for this service
//type RepositoryPath struct {
//	Href string // The deep link to the repository path where the linked service's code exists (Required)
//	Path string // The path where the linked service's code exists, relative to the root of the repository (Required)
//}
//
//// Rubric Rubrics allow you to score your services against different categories and levels
//type Rubric struct {
//}
//
//// Scorecard A scorecard
//type Scorecard struct {
//	AffectsOverallServiceLevels bool                    // Specifies whether the checks on this scorecard affect services' overall maturity level (Required)
//	Aliases                     []string                // Aliases of the scorecard (Required)
//	Description                 string                  // Description of the scorecard (Optional)
//	Filter                      FilterId                // Filter used by the scorecard to restrict services (Optional)
//	Href                        string                  // The hypertext reference (link) to the UI showing this scorecard (Required)
//	Id                          ID                      // A reference to the scorecard (Required)
//	Name                        string                  // Name of the scorecard (Required)
//	Owner                       EntityOwner             // The owner of this scorecard. Can currently either be a team or a group (Optional)
//	PassingChecks               int                     // The number of checks that are passing on this scorecard. A check executed against two services counts as two (Required)
//	ServiceCount                int                     // The number of services covered by this scorecard (Required)
//	ServicesReport              ScorecardServicesReport // Service stats regarding this scorecard (Optional)
//	Slug                        string                  // Slug of the scorecard (Required)
//	TotalChecks                 int                     // The number of checks that are performed on this scorecard. A check executed against two services counts as two (Required)
//}
//
//// ScorecardServicesReport Service stats regarding this scorecard
//type ScorecardServicesReport struct {
//	LevelCounts []LevelCount // Services per level regarding this scorecard (Required)
//}
//
//// Secret A sensitive value
//type Secret struct {
//	Alias      string     // A human reference for the secret (Required)
//	Id         ID         // A reference for the secret (Required)
//	Owner      TeamId     // The owner of this secret (Optional)
//	Timestamps Timestamps // Relevant timestamps (Required)
//}
//
//// ServiceId A service represents software deployed in your production infrastructure
//type ServiceId struct {
//	Id      ID       // The unique identifier for the service.
//	Aliases []string // A list of human-friendly, unique identifiers for the service.
//}
//
//// Service A service represents software deployed in your production infrastructure
//type Service struct {
//	ServiceId
//	Alias                      string                // A human-friendly, unique identifier for the service (Optional)
//	ApiDocumentPath            string                // The path, relative to the service repository's base directory, from which to fetch the API document. If null, the API document is fetched from the path in the account's apiDocsDefaultPath field (Optional)
//	CheckStats                 CheckStats            // A summary of check results on the service (Optional)
//	ConfigErrors               []ConfigError         // List of errors that occurred while syncing an opslevel.yml file (Optional)
//	DefaultServiceRepository   ServiceRepository     // Either the primary service repository, or the first valid non-archived repo, if any (Optional)
//	Description                string                // A brief description of the service (Optional)
//	ExternalUuid               string                // The service's external UUID (Optional)
//	Framework                  string                // The primary software development framework that the service uses (Optional)
//	HealthReport               ServiceMaturityReport // A health report on the service (Optional)
//	HtmlUrl                    string                // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
//	Language                   string                // The primary programming language that the service is written in (Optional)
//	LastDeploy                 Deploy                // The most recent production `Deploy` for this service. A deploy is considered production if its environment contains 'prod' or is empty (Optional)
//	Lifecycle                  Lifecycle             // The lifecycle stage of the service (Optional)
//	Locked                     bool                  // Indicates if the service is locked by an opslevel.yml (Required)
//	ManagedAliases             []string              // A list of aliases that can be set by users. The unique identifier for the resource is omitted (Required)
//	MaturityReport             ServiceMaturityReport // A health report on the service (Optional)
//	Name                       string                // The display name of the service (Required)
//	Note                       string                // Additional information about the service (Optional)
//	Owner                      TeamId                // The team that owns the service (Optional)
//	Parent                     SystemId              // Parent System of the Service (Optional)
//	PreferredApiDocument       ServiceDocument       // The API document selected for display on the API docs tab on the service's page (Optional)
//	PreferredApiDocumentSource ApiDocumentSourceEnum // The API document source (push or pull) used to determine the preferred API document. If null, we try the pushed doc and then the pulled doc (in that order) (Optional)
//	Product                    string                // A product is an application that your end user interacts with. Multiple services can work together to power a single product (Optional)
//	Property                   Property              // A custom property value assigned to this entity (Optional)
//	RawNote                    string                // The raw unsanitized additional information about the service (Optional)
//	Tier                       Tier                  // The software tier that the service belongs to (Optional)
//	Timestamps                 Timestamps            // Relevant timestamps (Required)
//	Type                       ComponentTypeId       // The type of the component (Required)
//}
//
//// ServiceDependency A service dependency edge
//type ServiceDependency struct {
//	DestinationService ServiceId // The service that was depended upon (Required)
//	Id                 ID        // ID of the service dependency edge (Required)
//	Notes              string    // Notes about the dependency edge (Optional)
//	SourceService      ServiceId // The service that had the dependency (Required)
//}
//
//// ServiceDocument A document that is attached to resource(s) in OpsLevel
//type ServiceDocument struct {
//	Content    string                // The contents of the document (Optional)
//	HtmlUrl    string                // The URL of the document, if any (Optional)
//	Id         ID                    // The ID of the Document (Required)
//	Source     ServiceDocumentSource // The source of the document (Required)
//	Timestamps Timestamps            // When the document was created and updated (Required)
//}
//
//// ServiceLevelNotifications
//type ServiceLevelNotifications struct {
//	SlackNotificationEnabled bool // Whether slack notifications on service level changes are enabled on your account (Required)
//}
//
//// ServiceMaturityReport The health report for this service in terms of its levels and checks
//type ServiceMaturityReport struct {
//	CategoryBreakdown  []CategoryLevel // The level of each category for this service (Required)
//	LatestCheckResults []CheckResult   // The latest check results for this service across the given checks (Optional)
//	OverallLevel       Level           // The overall level for this service (Required)
//}
//
//// ServiceRepository A record of the connection between a service and a repository
//type ServiceRepository struct {
//	BaseDirectory string       // The directory in the repository where service information exists, including the opslevel.yml file. This path is always returned without leading and trailing slashes (Optional)
//	DisplayName   string       // The name displayed in the UI for the service repository (Optional)
//	Id            ID           // ID of the service repository (Required)
//	Repository    RepositoryId // The repository that is part of this connection (Required)
//	Service       ServiceId    // The service that is part of this connection (Required)
//}
//
//// Stats An object that contains statistics
//type Stats struct {
//	Total           int // How many there are (Required)
//	TotalSuccessful int // How many are successfully passing (Required)
//}
//
//// SystemId A collection of related Services
//type SystemId struct {
//	Id      ID       // The identifier of the object.
//	Aliases []string // All of the aliases attached to the resource.
//}
//
//// System A collection of related Services
//type System struct {
//	SystemId
//	Description    string      // The description of the System (Optional)
//	HtmlUrl        string      // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
//	ManagedAliases []string    // A list of aliases that can be set by users. The unique identifier for the resource is omitted (Required)
//	Name           string      // The name of the object (Required)
//	Note           string      // Additional information about the system (Optional)
//	Owner          EntityOwner // The owner of the object (Optional)
//	Parent         DomainId    // Parent domain of the System (Optional)
//}
//
//// Tag An arbitrary key-value pair associated with a resource
//type Tag struct {
//	Id    ID       // The unique identifier for the tag (Required)
//	Key   string   // The tag's key (Required)
//	Owner TagOwner // The resource that the tag applies to (Required)
//	Value string   // The tag's value (Required)
//}
//
//// TagRelationshipKeys Returns the keys that set relationships when imported from AWS
//type TagRelationshipKeys struct {
//	BelongsTo    string   // The tag key that will create `belongs_to` relationships (Required)
//	DependencyOf []string // The tag keys that will create `dependency_of` relationships (Required)
//	DependsOn    []string // The tag keys that will create `depends_on` relationships (Required)
//}
//
//// TeamId A team belongs to your organization. Teams can own multiple services
//type TeamId struct {
//	Id      ID       // The unique identifier for the team.
//	Aliases []string // A list of human-friendly, unique identifiers for the team.
//}
//
//// Team A team belongs to your organization. Teams can own multiple services
//type Team struct {
//	TeamId
//	Alias            string    // The human-friendly, unique identifier for the team (Optional)
//	Contacts         []Contact // The contacts for the team (Optional)
//	HtmlUrl          string    // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
//	ManagedAliases   []string  // A list of aliases that can be set by users. The unique identifier for the resource is omitted (Required)
//	Manager          UserId    // The user who manages the team (Optional)
//	Name             string    // The team's display name (Required)
//	ParentTeam       TeamId    // The parent team (Optional)
//	Responsibilities string    // A description of what the team is responsible for (Optional)
//}
//
//// TeamMembership
//type TeamMembership struct {
//	Role string // Role of the user on the Team (Optional)
//	Team TeamId // Team for the membership (Required)
//	User UserId // User for the membership (Required)
//}
//
//// Tier A tier measures how critical or important a service is to your business
//type Tier struct {
//	Alias       string // The human-friendly, unique identifier for the tier (Optional)
//	Description string // A brief description of the tier (Optional)
//	Id          ID     // The unique identifier for the tier (Required)
//	Index       int    // The numerical representation of the tier (Optional)
//	Name        string // The display name of the tier (Optional)
//}
//
//// Timestamps Relevant timestamps
//type Timestamps struct {
//	CreatedAt iso8601.Time // The time at which the entity was created (Required)
//	UpdatedAt iso8601.Time // The time at which the entity was most recently updated (Required)
//}
//
//// Tool A tool is used to support the operations of a service
//type Tool struct {
//	Category      ToolCategory // The category that the tool belongs to (Optional)
//	CategoryAlias string       // The human-friendly, unique identifier for the tool's category (Optional)
//	DisplayName   string       // The display name of the tool (Optional)
//	Environment   string       // The environment that the tool belongs to (Optional)
//	Id            ID           // The unique identifier for the tool (Required)
//	Service       ServiceId    // The service that is associated to the tool (Required)
//	Url           string       // The URL of the tool (Required)
//}
//
//// UserId A user is someone who belongs to an organization
//type UserId struct {
//	Id    ID     // The unique identifier for the user.
//	Email string // The user's email.
//}
//
//// User A user is someone who belongs to an organization
//type User struct {
//	UserId
//	HtmlUrl string   // A link to the HTML page for the resource. Ex. https://app.opslevel.com/services/shopping_cart (Required)
//	Name    string   // The user's full name (Required)
//	Role    UserRole // The user's assigned role (Optional)
//}
//
//// Warning The warnings of the mutation
//type Warning struct {
//	Message string // The warning message (Required)
//}
