// Code generated; DO NOT EDIT.
package opslevel

// TODO: should probably rename this
type Connection interface {
	GetNodes() any
}

type ConnectionBase[T any] struct {
	Nodes      []T      // A list of nodes
	PageInfo   PageInfo // Information to aid in pagination
	TotalCount int      `graphql:"-"` // The number of returned nodes
}

func (s *ConnectionBase[T]) GetNodes() any {
	return s.Nodes
}

// AlertSourceConnection The connection type for AlertSource
type AlertSourceConnection ConnectionBase[AlertSource]

// AlertSourceServiceConnection The connection type for AlertSource
type AlertSourceServiceConnection ConnectionBase[AlertSource]

// AlertSourceServiceV2Connection The connection type for AlertSourceService
type AlertSourceServiceV2Connection ConnectionBase[AlertSourceService]

// CampaignConnection The connection type for Campaign
type CampaignConnection ConnectionBase[Campaign]

// CategoryConnection The connection type for Category
type CategoryConnection ConnectionBase[Category]

// CheckConnection The connection type for Check
type CheckConnection ConnectionBase[Check]

// CheckResultsByLevelConnection The connection type for CheckResultsByLevel
type CheckResultsByLevelConnection ConnectionBase[CheckResultsByLevel]

// CheckResultsConnection The connection type for CheckResult
type CheckResultsConnection ConnectionBase[CheckResult]

// ComponentTypeConnection The connection type for ComponentType
type ComponentTypeConnection ConnectionBase[ComponentType]

// CustomActionsExternalActionsConnection The connection type for CustomActionsExternalAction
type CustomActionsExternalActionsConnection ConnectionBase[CustomActionsExternalAction]

// CustomActionsTriggerDefinitionConnection The connection type for CustomActionsTriggerDefinition
type CustomActionsTriggerDefinitionConnection ConnectionBase[CustomActionsTriggerDefinition]

// DeployConnection The connection type for Deploy
type DeployConnection ConnectionBase[Deploy]

// DomainConnection The connection type for Domain
type DomainConnection ConnectionBase[Domain]

// FilterConnection The connection type for Filter
type FilterConnection ConnectionBase[Filter]

// InfrastructureResourceConnection The connection type for InfrastructureResource
type InfrastructureResourceConnection ConnectionBase[InfrastructureResource]

// InfrastructureResourceSchemaConnection The connection type for InfrastructureResourceSchema
type InfrastructureResourceSchemaConnection ConnectionBase[InfrastructureResourceSchema]

// IntegrationConnection The connection type for Integration
type IntegrationConnection ConnectionBase[Integration]

// LevelConnection The connection type for Level
type LevelConnection ConnectionBase[Level]

// PropertyConnection The connection type for Property
type PropertyConnection ConnectionBase[Property]

// PropertyDefinitionConnection The connection type for PropertyDefinition
type PropertyDefinitionConnection ConnectionBase[PropertyDefinition]

// RelatedResourceConnection The connection type for RelationshipResource
type RelatedResourceConnection ConnectionBase[RelationshipResource]

// RelationshipConnection The connection type for RelationshipNode
type RelationshipConnection ConnectionBase[RelationshipNode]

// RelationshipDefinitionConnection The connection type for RelationshipDefinitionType
type RelationshipDefinitionConnection ConnectionBase[RelationshipDefinitionType]

// ScorecardCategoryConnection The connection type for Category
type ScorecardCategoryConnection ConnectionBase[Category]

// ScorecardCheckConnection The connection type for Check
type ScorecardCheckConnection ConnectionBase[Check]

// ScorecardConnection The connection type for Scorecard
type ScorecardConnection ConnectionBase[Scorecard]

// ScorecardStatsConnection The connection type for ScorecardStats
type ScorecardStatsConnection ConnectionBase[ScorecardStats]

// SecretsVaultsSecretsConnection The connection type for Secret
type SecretsVaultsSecretsConnection ConnectionBase[Secret]

// ServiceCategoryConnection The connection type for Category
type ServiceCategoryConnection ConnectionBase[Category]

// ServiceConnection The connection type for Service
type ServiceConnection ConnectionBase[Service]

// ServiceDocumentConnection The connection type for ServiceDocument
type ServiceDocumentConnection ConnectionBase[ServiceDocument]

// SystemConnection The connection type for System
type SystemConnection ConnectionBase[System]

// TagConnection The connection type for Tag
type TagConnection ConnectionBase[Tag]

// TagRepositoryConnection The connection type for Tag
type TagRepositoryConnection ConnectionBase[Tag]

// TeamConnection The connection type for Team
type TeamConnection ConnectionBase[Team]

// TeamMembershipConnection The connection type for TeamMembership
type TeamMembershipConnection ConnectionBase[TeamMembership]

// TeamPropertyDefinitionConnection The connection type for TeamPropertyDefinition
type TeamPropertyDefinitionConnection ConnectionBase[TeamPropertyDefinition]

// ToolConnection The connection type for Tool
type ToolConnection ConnectionBase[Tool]

// UserConnection The connection type for User
type UserConnection ConnectionBase[User]
