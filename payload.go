// Code generated; DO NOT EDIT.
package opslevel

type BasePayload struct {
	Errors []Error // List of errors that occurred while executing the mutation (Required)
}

// AlertSourceServiceCreatePayload Return type for the `alertSourceServiceCreate` mutation
type AlertSourceServiceCreatePayload struct {
	AlertSourceService AlertSourceService // An alert source service representing a connection between a service and an alert source (Optional)
	BasePayload
}

// AliasCreatePayload Return type for the `aliasCreate` mutation
type AliasCreatePayload struct {
	Aliases []string // All of the aliases attached to the resource (Optional)
	OwnerId string   // The ID of the resource that had an alias attached (Optional)
	BasePayload
}

// CategoryCreatePayload The return type of the `categoryCreate` mutation
type CategoryCreatePayload struct {
	Category Category // A category is used to group related checks in a rubric (Optional)
	BasePayload
}

// CategoryUpdatePayload The return type of the `categoryUpdate` mutation
type CategoryUpdatePayload struct {
	Category Category // A category is used to group related checks in a rubric (Optional)
	BasePayload
}

// CheckCopyPayload The result of a check copying operation
type CheckCopyPayload struct {
	TargetCategory Category // The category to which the checks have been copied (Optional)
	BasePayload
}

// CheckResponsePayload The return type of a `checkCreate` mutation and `checkUpdate` mutation
type CheckResponsePayload struct {
	Check Check // The newly created check (Optional)
	BasePayload
}

// ComponentTypePayload Return type for the `componentTypeCreate` mutation
type ComponentTypePayload struct {
	ComponentType ComponentType // The created component type (Optional)
	BasePayload
}

// ContactCreatePayload The return type of a `contactCreate` mutation
type ContactCreatePayload struct {
	Contact Contact // A method of contact for a team (Optional)
	BasePayload
}

// ContactUpdatePayload The return type of a `contactUpdate` mutation
type ContactUpdatePayload struct {
	Contact Contact // A method of contact for a team (Optional)
	BasePayload
}

// CustomActionsTriggerDefinitionCreatePayload Return type for the `customActionsTriggerDefinitionCreate` mutation
type CustomActionsTriggerDefinitionCreatePayload struct {
	TriggerDefinition CustomActionsTriggerDefinition // The definition of a potential trigger for a custom action (Optional)
	BasePayload
}

// CustomActionsTriggerDefinitionUpdatePayload Return type for the `customActionsTriggerDefinitionUpdate` mutation
type CustomActionsTriggerDefinitionUpdatePayload struct {
	TriggerDefinition CustomActionsTriggerDefinition // The definition of a potential trigger for a custom action (Optional)
	BasePayload
}

// CustomActionsWebhookActionCreatePayload Return type for the `customActionsWebhookActionCreate` mutation
type CustomActionsWebhookActionCreatePayload struct {
	WebhookAction CustomActionsWebhookAction // An external webhook action to be triggered by a custom action (Optional)
	BasePayload
}

// CustomActionsWebhookActionUpdatePayload The response returned after updating a Webhook Action
type CustomActionsWebhookActionUpdatePayload struct {
	WebhookAction CustomActionsWebhookAction // An external webhook action to be triggered by a custom action (Optional)
	BasePayload
}

// DomainChildAssignPayload Return type for the `domainChildAssign` mutation
type DomainChildAssignPayload struct {
	Domain Domain // The domain after children have been assigned (Optional)
	BasePayload
}

// DomainChildRemovePayload Return type for the `domainChildRemove` mutation
type DomainChildRemovePayload struct {
	Domain Domain // The domain after children have been removed (Optional)
	BasePayload
}

// DomainPayload Return type for `domainCreate` and `domainUpdate` mutations
type DomainPayload struct {
	Domain Domain // A collection of related Systems (Optional)
	BasePayload
}

// ExportConfigFilePayload The result of exporting an object as YAML
type ExportConfigFilePayload struct {
	Kind string // The GraphQL type that represents the exported object (Optional)
	Yaml string // The YAML representation of the object (Optional)
	BasePayload
}

// ExternalUuidMutationPayload Return type for the external UUID mutations
type ExternalUuidMutationPayload struct {
	ExternalUuid string // The updated external UUID of the resource (Optional)
	BasePayload
}

// FilterCreatePayload The return type of a `filterCreatePayload` mutation
type FilterCreatePayload struct {
	Filter Filter // The newly created filter (Optional)
	BasePayload
}

// FilterUpdatePayload The return type of the `filterUpdate` mutation
type FilterUpdatePayload struct {
	Filter Filter // The updated filter (Optional)
	BasePayload
}

// ImportEntityFromBackstagePayload Results of importing an Entity from Backstage into OpsLevel
type ImportEntityFromBackstagePayload struct {
	ActionMessage string // The action taken by OpsLevel (ie: service created) (Required)
	HtmlUrl       string // A link to the created or updated object in OpsLevel, if any (Optional)
	BasePayload
}

// InfrastructureResourcePayload Return type for the `infrastructureResourceUpdate` mutation
type InfrastructureResourcePayload struct {
	InfrastructureResource InfrastructureResource // An Infrastructure Resource (Optional)
	Warnings               []Warning              // The warnings of the mutation (Required)
	BasePayload
}

// IntegrationCreatePayload The result of creating an integration
type IntegrationCreatePayload struct {
	Integration Integration // The newly created integration (Optional)
	BasePayload
}

// IntegrationReactivatePayload The return type of a 'integrationReactivate' mutation
type IntegrationReactivatePayload struct {
	Integration Integration // The newly reactivated integration (Optional)
	BasePayload
}

// IntegrationSourceObjectUpsertPayload The return type of a 'integrationSourceObjectUpsert' mutation
type IntegrationSourceObjectUpsertPayload struct {
	Integration Integration // The integration that the source object was upserted to (Optional)
	BasePayload
}

// IntegrationUpdatePayload The return type of a 'integrationUpdate' mutation
type IntegrationUpdatePayload struct {
	Integration Integration // The newly updated integration (Optional)
	BasePayload
}

// LevelCreatePayload The return type of the `levelCreate` mutation
type LevelCreatePayload struct {
	Level Level // A performance rating that is used to grade your services against (Optional)
	BasePayload
}

// LevelUpdatePayload The return type of the `levelUpdate` mutation
type LevelUpdatePayload struct {
	Level Level // A performance rating that is used to grade your services against (Optional)
	BasePayload
}

// NewRelicAccountsPayload
type NewRelicAccountsPayload struct {
	BasePayload
}

// PropertyDefinitionPayload The return type for property definition mutations
type PropertyDefinitionPayload struct {
	Definition PropertyDefinition // The property that was defined (Optional)
	BasePayload
}

// PropertyPayload The payload for setting a property
type PropertyPayload struct {
	Property Property // The property that was set (Optional)
	BasePayload
}

// PropertyUnassignPayload The payload for unassigning a property
type PropertyUnassignPayload struct {
	Definition PropertyDefinition // The definition of the property that was unassigned (Optional)
	Owner      EntityOwnerService // The entity that the property was unassigned from (Optional)
	BasePayload
}

// RepositoriesUpdatePayload Return type for the `repositoriesUpdate` mutation
type RepositoriesUpdatePayload struct {
	NotUpdatedRepositories []RepositoryOperationErrorPayload // The repository objects that were not updated along with the error that happened when attempting to update the repository (Optional)
	UpdatedRepositories    []Repository                      // The identifiers of the updated repositories (Optional)
	BasePayload
}

// RepositoryOperationErrorPayload Specifies the repository and error after attempting and failing to perform a CRUD operation on a repository
type RepositoryOperationErrorPayload struct {
	Error      string     // The error message after an operation was attempted (Optional)
	Repository Repository // The repository on which an operation was attempted (Required)
	BasePayload
}

// RepositoryUpdatePayload The return type of a `repositoryUpdate` mutation
type RepositoryUpdatePayload struct {
	Repository Repository // A repository contains code that pertains to a service (Optional)
	BasePayload
}

// ScorecardPayload The type returned when creating a scorecard
type ScorecardPayload struct {
	Scorecard Scorecard // The created scorecard (Optional)
	BasePayload
}

// SecretPayload Return type for secret operations
type SecretPayload struct {
	Secret Secret // A sensitive value (Optional)
	BasePayload
}

// ServiceCreatePayload Return type for the `serviceCreate` mutation
type ServiceCreatePayload struct {
	Service Service // The newly created service (Optional)
	BasePayload
}

// ServiceDependencyPayload Return type for the requested `serviceDependency`
type ServiceDependencyPayload struct {
	ServiceDependency ServiceDependency // A service dependency edge (Optional)
	BasePayload
}

// ServiceLevelNotificationsPayload The return type of the service level notifications update mutation
type ServiceLevelNotificationsPayload struct {
	ServiceLevelNotifications ServiceLevelNotifications // The updated service level notification settings (Optional)
	BasePayload
}

// ServiceNoteUpdatePayload Return type for the `serviceNoteUpdate` mutation
type ServiceNoteUpdatePayload struct {
	Service Service // A service represents software deployed in your production infrastructure (Optional)
	BasePayload
}

// ServiceRepositoryCreatePayload Return type for the `serviceRepositoryCreate` mutation
type ServiceRepositoryCreatePayload struct {
	ServiceRepository ServiceRepository // A record of the connection between a service and a repository (Optional)
	BasePayload
}

// ServiceRepositoryUpdatePayload The return type of the `serviceRepositoryUpdate` mutation
type ServiceRepositoryUpdatePayload struct {
	ServiceRepository ServiceRepository // The updated service repository (Optional)
	BasePayload
}

// ServiceUpdatePayload Return type for the `serviceUpdate` mutation
type ServiceUpdatePayload struct {
	Service Service // The updated service (Optional)
	BasePayload
}

// SystemChildAssignPayload Return type for the `systemChildAssign` mutation
type SystemChildAssignPayload struct {
	System System // The system after children have been assigned (Optional)
	BasePayload
}

// SystemChildRemovePayload Return type for the `systemChildRemove` mutation
type SystemChildRemovePayload struct {
	System System // The system after children have been removed (Optional)
	BasePayload
}

// SystemPayload Return type for the `systemCreate` and `systemUpdate` mutations
type SystemPayload struct {
	System System // A collection of related Services (Optional)
	BasePayload
}

// TagAssignPayload The return type of a `tagAssign` mutation
type TagAssignPayload struct {
	Tags []Tag // The new tags that have been assigned to the resource (Optional)
	BasePayload
}

// TagCreatePayload The return type of a `tagCreate` mutation
type TagCreatePayload struct {
	Tag Tag // The newly created tag (Optional)
	BasePayload
}

// TagUpdatePayload The return type of a `tagUpdate` mutation
type TagUpdatePayload struct {
	Tag Tag // The newly updated tag (Optional)
	BasePayload
}

// TeamCreatePayload The return type of a `teamCreate` mutation
type TeamCreatePayload struct {
	Team Team // A team belongs to your organization. Teams can own multiple services (Optional)
	BasePayload
}

// TeamMembershipCreatePayload The response returned when creating memberships on teams
type TeamMembershipCreatePayload struct {
	Members     []User           // A list of users that are a member of the team (Optional)
	Memberships []TeamMembership // A list of memberships on the team (Optional)
	BasePayload
}

// TeamPropertyDefinitionPayload The return type for team property definition mutations
type TeamPropertyDefinitionPayload struct {
	Definition TeamPropertyDefinition // The team property that was defined (Optional)
	BasePayload
}

// TeamUpdatePayload The return type of a `teamUpdate` mutation
type TeamUpdatePayload struct {
	Team Team // A team belongs to your organization. Teams can own multiple services (Optional)
	BasePayload
}

// ToolCreatePayload The return type of a `toolCreate` mutation
type ToolCreatePayload struct {
	Tool Tool // A tool is used to support the operations of a service (Optional)
	BasePayload
}

// ToolUpdatePayload The return type of a `toolUpdate` payload
type ToolUpdatePayload struct {
	Tool Tool // A tool is used to support the operations of a service (Optional)
	BasePayload
}

// UserPayload The return type of user management mutations
type UserPayload struct {
	User User // A user is someone who belongs to an organization (Optional)
	BasePayload
}

// UsersInvitePayload The return type of the users invite mutation
type UsersInvitePayload struct {
	Failed []string // The user identifiers which failed to successfully send an invite (Optional)
	Users  []User   // The users that were successfully invited (Optional)
	BasePayload
}
