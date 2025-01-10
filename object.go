// Code generated; DO NOT EDIT.
package opslevel

import "github.com/relvacode/iso8601"

// AlertSource An alert source that is currently integrated and belongs to the account
type AlertSource struct {
}

// AlertSourceService An alert source that is connected with a service
type AlertSourceService struct {
}

// AlertSourceUsageCheck
type AlertSourceUsageCheck struct {
}

// ApiDocIntegration
type ApiDocIntegration struct {
}

// ArgocdDeployIntegration
type ArgocdDeployIntegration struct {
}

// AwsIntegration
type AwsIntegration struct {
}

// AzureDevopsIntegration
type AzureDevopsIntegration struct {
}

// AzureDevopsPermissionError
type AzureDevopsPermissionError struct {
}

// AzureResourcesIntegration
type AzureResourcesIntegration struct {
}

// BitbucketIntegration
type BitbucketIntegration struct {
}

// Category A category is used to group related checks in a rubric
type Category struct {
}

// CategoryLevel The level of a specific category
type CategoryLevel struct {
}

// CheckIntegration
type CheckIntegration struct {
}

// CheckResult The result for a given Check
type CheckResult struct {
}

// CheckStats Check stats shows a summary of check results
type CheckStats struct {
}

// CircleciDeployIntegration
type CircleciDeployIntegration struct {
}

// CommonVulnerabilityEnumeration A category system for hardware and software weaknesses
type CommonVulnerabilityEnumeration struct {
}

// CommonWeaknessEnumeration A category system for hardware and software weaknesses
type CommonWeaknessEnumeration struct {
}

// ComponentType Information about a particular component type
type ComponentType struct {
}

// ComponentTypeIcon The icon for a component type
type ComponentTypeIcon struct {
}

// ConfigError An error that occurred when syncing an opslevel.yml file
type ConfigError struct {
}

// ConfigFile An OpsLevel config as code definition
type ConfigFile struct {
}

// CustomActionsTemplate Template of a custom action
type CustomActionsTemplate struct {
}

// CustomActionsTemplatesAction The action of a custom action template
type CustomActionsTemplatesAction struct {
}

// CustomActionsTemplatesMetadata The metadata about the custom action template
type CustomActionsTemplatesMetadata struct {
}

// CustomActionsTemplatesTriggerDefinition The definition of a potential trigger for a template custom action
type CustomActionsTemplatesTriggerDefinition struct {
}

// CustomActionsTriggerDefinition The definition of a potential trigger for a custom action
type CustomActionsTriggerDefinition struct {
}

// CustomCheck
type CustomCheck struct {
}

// CustomEventCheck
type CustomEventCheck struct {
}

// DatadogIntegration
type DatadogIntegration struct {
}

// Deploy An event sent via webhook to track deploys
type Deploy struct {
}

// DeployIntegration
type DeployIntegration struct {
}

// Domain A collection of related Systems
type Domain struct {
}

// Filter A filter is used to select which services will have checks applied. It can also be used to filter services in reports
type Filter struct {
}

// FilterPredicate A condition used to select services
type FilterPredicate struct {
}

// FluxIntegration
type FluxIntegration struct {
}

// GenericIntegration
type GenericIntegration struct {
}

// GitBranchProtectionCheck
type GitBranchProtectionCheck struct {
}

// GitLabCIIntegration
type GitLabCIIntegration struct {
}

// GithubActionsIntegration
type GithubActionsIntegration struct {
}

// GithubIntegration
type GithubIntegration struct {
}

// GitlabIntegration
type GitlabIntegration struct {
}

// GoogleCloudIntegration
type GoogleCloudIntegration struct {
}

// GoogleCloudProject
type GoogleCloudProject struct {
}

// HasDocumentationCheck
type HasDocumentationCheck struct {
}

// HasRecentDeployCheck
type HasRecentDeployCheck struct {
}

// IncidentIoIntegration
type IncidentIoIntegration struct {
}

// InfrastructureResource An Infrastructure Resource
type InfrastructureResource struct {
}

// InfrastructureResourceProviderData Data about the provider the infrastructure resource is from
type InfrastructureResourceProviderData struct {
}

// InfrastructureResourceSchema A schema for Infrastructure Resources
type InfrastructureResourceSchema struct {
}

// IssueTrackingIntegration
type IssueTrackingIntegration struct {
}

// JenkinsIntegration
type JenkinsIntegration struct {
}

// KubernetesIntegration
type KubernetesIntegration struct {
}

// Language A language that can be assigned to a repository
type Language struct {
}

// Level A performance rating that is used to grade your services against
type Level struct {
}

// LevelCount The total number of services in each level
type LevelCount struct {
}

// Lifecycle A lifecycle represents the current development stage of a service
type Lifecycle struct {
}

// ManualCheck
type ManualCheck struct {
}

// ManualCheckFrequency
type ManualCheckFrequency struct {
}

// MicrosoftTeamsIntegration
type MicrosoftTeamsIntegration struct {
}

// NewRelicIntegration
type NewRelicIntegration struct {
}

// OctopusDeployIntegration
type OctopusDeployIntegration struct {
}

// OnPremGitlabIntegration
type OnPremGitlabIntegration struct {
}

// OpsgenieIntegration
type OpsgenieIntegration struct {
}

// PackageVersionCheck
type PackageVersionCheck struct {
}

// PageInfo Information about pagination in a connection
type PageInfo struct {
}

// PagerdutyIntegration
type PagerdutyIntegration struct {
}

// PayloadCheck
type PayloadCheck struct {
}

// PayloadIntegration
type PayloadIntegration struct {
}

// Predicate A condition used to select services
type Predicate struct {
}

// Property A custom property value assigned to an entity
type Property struct {
}

// PropertyDefinition The definition of a property
type PropertyDefinition struct {
}

// RelationshipType The type specifying a relationship between two resources
type RelationshipType struct {
}

// Repository A repository contains code that pertains to a service
type Repository struct {
}

// RepositoryFileCheck
type RepositoryFileCheck struct {
}

// RepositoryGrepCheck
type RepositoryGrepCheck struct {
}

// RepositoryIntegratedCheck
type RepositoryIntegratedCheck struct {
}

// RepositoryPath The repository path used for this service
type RepositoryPath struct {
}

// RepositorySearchCheck
type RepositorySearchCheck struct {
}

// Rubric Rubrics allow you to score your services against different categories and levels
type Rubric struct {
}

// ScimIntegration
type ScimIntegration struct {
}

// Scorecard A scorecard
type Scorecard struct {
}

// ScorecardServicesReport Service stats regarding this scorecard
type ScorecardServicesReport struct {
}

// Secret A sensitive value
type Secret struct {
}

// Service A service represents software deployed in your production infrastructure
type Service struct {
}

// ServiceConfigurationCheck
type ServiceConfigurationCheck struct {
}

// ServiceDependency A service dependency edge
type ServiceDependency struct {
}

// ServiceDependencyCheck
type ServiceDependencyCheck struct {
}

// ServiceDocument A document that is attached to resource(s) in OpsLevel
type ServiceDocument struct {
}

// ServiceLevelNotifications
type ServiceLevelNotifications struct {
}

// ServiceMaturityReport The health report for this service in terms of its levels and checks
type ServiceMaturityReport struct {
}

// ServiceOwnershipCheck
type ServiceOwnershipCheck struct {
}

// ServicePropertyCheck
type ServicePropertyCheck struct {
}

// ServiceRepository A record of the connection between a service and a repository
type ServiceRepository struct {
}

// SlackIntegration
type SlackIntegration struct {
}

// SnykIntegration
type SnykIntegration struct {
}

// SonarqubeCloudIntegration
type SonarqubeCloudIntegration struct {
}

// SonarqubeIntegration
type SonarqubeIntegration struct {
}

// Stats An object that contains statistics
type Stats struct {
}

// System A collection of related Services
type System struct {
}

// Tag An arbitrary key-value pair associated with a resource
type Tag struct {
}

// TagDefinedCheck
type TagDefinedCheck struct {
}

// TagRelationshipKeys Returns the keys that set relationships when imported from AWS
type TagRelationshipKeys struct {
}

// Team A team belongs to your organization. Teams can own multiple services
type Team struct {
}

// TeamMembership
type TeamMembership struct {
}

// TerraformIntegration
type TerraformIntegration struct {
}

// Tier A tier measures how critical or important a service is to your business
type Tier struct {
}

// Timestamps Relevant timestamps
type Timestamps struct {
}

// Tool A tool is used to support the operations of a service
type Tool struct {
}

// ToolUsageCheck
type ToolUsageCheck struct {
}

// User A user is someone who belongs to an organization
type User struct {
}
