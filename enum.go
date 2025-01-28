// Code generated; DO NOT EDIT.
package opslevel

// AlertSourceStatusTypeEnum The monitor status level
type AlertSourceStatusTypeEnum string

var (
	AlertSourceStatusTypeEnumAlert        AlertSourceStatusTypeEnum = "alert"         // Monitor is reporting an alert
	AlertSourceStatusTypeEnumFetchingData AlertSourceStatusTypeEnum = "fetching_data" // Monitor currently being updated
	AlertSourceStatusTypeEnumNoData       AlertSourceStatusTypeEnum = "no_data"       // No data received yet. Ensure your monitors are configured correctly
	AlertSourceStatusTypeEnumOk           AlertSourceStatusTypeEnum = "ok"            // Monitor is not reporting any warnings or alerts
	AlertSourceStatusTypeEnumWarn         AlertSourceStatusTypeEnum = "warn"          // Monitor is reporting a warning
)

// All AlertSourceStatusTypeEnum as []string
var AllAlertSourceStatusTypeEnum = []string{
	string(AlertSourceStatusTypeEnumAlert),
	string(AlertSourceStatusTypeEnumFetchingData),
	string(AlertSourceStatusTypeEnumNoData),
	string(AlertSourceStatusTypeEnumOk),
	string(AlertSourceStatusTypeEnumWarn),
}

// AlertSourceTypeEnum The type of the alert source
type AlertSourceTypeEnum string

var (
	AlertSourceTypeEnumDatadog     AlertSourceTypeEnum = "datadog"      // A Datadog alert source (aka monitor)
	AlertSourceTypeEnumFireHydrant AlertSourceTypeEnum = "fire_hydrant" // An FireHydrant alert source (aka service)
	AlertSourceTypeEnumIncidentIo  AlertSourceTypeEnum = "incident_io"  // An incident.io alert source (aka service)
	AlertSourceTypeEnumNewRelic    AlertSourceTypeEnum = "new_relic"    // A New Relic alert source (aka service)
	AlertSourceTypeEnumOpsgenie    AlertSourceTypeEnum = "opsgenie"     // An Opsgenie alert source (aka service)
	AlertSourceTypeEnumPagerduty   AlertSourceTypeEnum = "pagerduty"    // A PagerDuty alert source (aka service)
)

// All AlertSourceTypeEnum as []string
var AllAlertSourceTypeEnum = []string{
	string(AlertSourceTypeEnumDatadog),
	string(AlertSourceTypeEnumFireHydrant),
	string(AlertSourceTypeEnumIncidentIo),
	string(AlertSourceTypeEnumNewRelic),
	string(AlertSourceTypeEnumOpsgenie),
	string(AlertSourceTypeEnumPagerduty),
}

// AliasOwnerTypeEnum The owner type an alias is assigned to
type AliasOwnerTypeEnum string

var (
	AliasOwnerTypeEnumDomain                 AliasOwnerTypeEnum = "domain"                  // Aliases that are assigned to domains
	AliasOwnerTypeEnumGroup                  AliasOwnerTypeEnum = "group"                   // Aliases that are assigned to groups
	AliasOwnerTypeEnumInfrastructureResource AliasOwnerTypeEnum = "infrastructure_resource" // Aliases that are assigned to infrastructure resources
	AliasOwnerTypeEnumScorecard              AliasOwnerTypeEnum = "scorecard"               // Aliases that are assigned to scorecards
	AliasOwnerTypeEnumService                AliasOwnerTypeEnum = "service"                 // Aliases that are assigned to services
	AliasOwnerTypeEnumSystem                 AliasOwnerTypeEnum = "system"                  // Aliases that are assigned to systems
	AliasOwnerTypeEnumTeam                   AliasOwnerTypeEnum = "team"                    // Aliases that are assigned to teams
)

// All AliasOwnerTypeEnum as []string
var AllAliasOwnerTypeEnum = []string{
	string(AliasOwnerTypeEnumDomain),
	string(AliasOwnerTypeEnumGroup),
	string(AliasOwnerTypeEnumInfrastructureResource),
	string(AliasOwnerTypeEnumScorecard),
	string(AliasOwnerTypeEnumService),
	string(AliasOwnerTypeEnumSystem),
	string(AliasOwnerTypeEnumTeam),
}

// ApiDocumentSourceEnum The source used to determine the preferred API document
type ApiDocumentSourceEnum string

var (
	ApiDocumentSourceEnumPull ApiDocumentSourceEnum = "PULL" // Use the document that was pulled by OpsLevel via a repo
	ApiDocumentSourceEnumPush ApiDocumentSourceEnum = "PUSH" // Use the document that was pushed to OpsLevel via an API Docs integration
)

// All ApiDocumentSourceEnum as []string
var AllApiDocumentSourceEnum = []string{
	string(ApiDocumentSourceEnumPull),
	string(ApiDocumentSourceEnumPush),
}

// BasicTypeEnum Operations that can be used on filters
type BasicTypeEnum string

var (
	BasicTypeEnumDoesNotEqual BasicTypeEnum = "does_not_equal" // Does not equal a specific value
	BasicTypeEnumEquals       BasicTypeEnum = "equals"         // Equals a specific value
)

// All BasicTypeEnum as []string
var AllBasicTypeEnum = []string{
	string(BasicTypeEnumDoesNotEqual),
	string(BasicTypeEnumEquals),
}

// CampaignFilterEnum Fields that can be used as part of filter for campaigns
type CampaignFilterEnum string

var (
	CampaignFilterEnumID     CampaignFilterEnum = "id"     // Filter by `id` of campaign
	CampaignFilterEnumOwner  CampaignFilterEnum = "owner"  // Filter by campaign owner
	CampaignFilterEnumStatus CampaignFilterEnum = "status" // Filter by campaign status
)

// All CampaignFilterEnum as []string
var AllCampaignFilterEnum = []string{
	string(CampaignFilterEnumID),
	string(CampaignFilterEnumOwner),
	string(CampaignFilterEnumStatus),
}

// CampaignReminderChannelEnum The possible communication channels through which a campaign reminder can be delivered
type CampaignReminderChannelEnum string

var (
	CampaignReminderChannelEnumEmail          CampaignReminderChannelEnum = "email"           // A system for sending messages to one or more recipients via telecommunications links between computers using dedicated software or a web-based service
	CampaignReminderChannelEnumMicrosoftTeams CampaignReminderChannelEnum = "microsoft_teams" // A proprietary business communication platform developed by Microsoft
	CampaignReminderChannelEnumSlack          CampaignReminderChannelEnum = "slack"           // A cloud-based team communication platform developed by Slack Technologies
)

// All CampaignReminderChannelEnum as []string
var AllCampaignReminderChannelEnum = []string{
	string(CampaignReminderChannelEnumEmail),
	string(CampaignReminderChannelEnumMicrosoftTeams),
	string(CampaignReminderChannelEnumSlack),
}

// CampaignReminderFrequencyUnitEnum Possible time units for the frequency at which campaign reminders are delivered
type CampaignReminderFrequencyUnitEnum string

var (
	CampaignReminderFrequencyUnitEnumDay   CampaignReminderFrequencyUnitEnum = "day"   // A period of twenty-four hours as a unit of time, reckoned from one midnight to the next, corresponding to a rotation of the earth on its axis
	CampaignReminderFrequencyUnitEnumMonth CampaignReminderFrequencyUnitEnum = "month" // Each of the twelve named periods into which a year is divided
	CampaignReminderFrequencyUnitEnumWeek  CampaignReminderFrequencyUnitEnum = "week"  // A period of seven days
)

// All CampaignReminderFrequencyUnitEnum as []string
var AllCampaignReminderFrequencyUnitEnum = []string{
	string(CampaignReminderFrequencyUnitEnumDay),
	string(CampaignReminderFrequencyUnitEnumMonth),
	string(CampaignReminderFrequencyUnitEnumWeek),
}

// CampaignReminderTypeEnum Type/Format of the notification
type CampaignReminderTypeEnum string

var (
	CampaignReminderTypeEnumEmail          CampaignReminderTypeEnum = "email"           // Notification will be sent via email
	CampaignReminderTypeEnumMicrosoftTeams CampaignReminderTypeEnum = "microsoft_teams" // Notification will be sent by microsoft teams
	CampaignReminderTypeEnumSlack          CampaignReminderTypeEnum = "slack"           // Notification will be sent by slack
)

// All CampaignReminderTypeEnum as []string
var AllCampaignReminderTypeEnum = []string{
	string(CampaignReminderTypeEnumEmail),
	string(CampaignReminderTypeEnumMicrosoftTeams),
	string(CampaignReminderTypeEnumSlack),
}

// CampaignServiceStatusEnum Status of whether a service is passing all checks for a campaign or not
type CampaignServiceStatusEnum string

var (
	CampaignServiceStatusEnumFailing CampaignServiceStatusEnum = "failing" // Service is failing one or more checks in the campaign
	CampaignServiceStatusEnumPassing CampaignServiceStatusEnum = "passing" // Service is passing all the checks in the campaign
)

// All CampaignServiceStatusEnum as []string
var AllCampaignServiceStatusEnum = []string{
	string(CampaignServiceStatusEnumFailing),
	string(CampaignServiceStatusEnumPassing),
}

// CampaignSortEnum Sort possibilities for campaigns
type CampaignSortEnum string

var (
	CampaignSortEnumChecksPassingAsc     CampaignSortEnum = "checks_passing_ASC"     // Sort by number of `checks passing` ascending
	CampaignSortEnumChecksPassingDesc    CampaignSortEnum = "checks_passing_DESC"    // Sort by number of `checks passing` descending
	CampaignSortEnumEndedDateAsc         CampaignSortEnum = "ended_date_ASC"         // Sort by `endedDate` ascending
	CampaignSortEnumEndedDateDesc        CampaignSortEnum = "ended_date_DESC"        // Sort by `endedDate` descending
	CampaignSortEnumFilterAsc            CampaignSortEnum = "filter_ASC"             // Sort by `filter` ascending
	CampaignSortEnumFilterDesc           CampaignSortEnum = "filter_DESC"            // Sort by `filter` descending
	CampaignSortEnumNameAsc              CampaignSortEnum = "name_ASC"               // Sort by `name` ascending
	CampaignSortEnumNameDesc             CampaignSortEnum = "name_DESC"              // Sort by `name` descending
	CampaignSortEnumOwnerAsc             CampaignSortEnum = "owner_ASC"              // Sort by `owner` ascending
	CampaignSortEnumOwnerDesc            CampaignSortEnum = "owner_DESC"             // Sort by `owner` descending
	CampaignSortEnumServicesCompleteAsc  CampaignSortEnum = "services_complete_ASC"  // Sort by number of `services complete` ascending
	CampaignSortEnumServicesCompleteDesc CampaignSortEnum = "services_complete_DESC" // Sort by number of `services complete` descending
	CampaignSortEnumStartDateAsc         CampaignSortEnum = "start_date_ASC"         // Sort by `startDate` ascending
	CampaignSortEnumStartDateDesc        CampaignSortEnum = "start_date_DESC"        // Sort by `startDate` descending
	CampaignSortEnumStatusAsc            CampaignSortEnum = "status_ASC"             // Sort by `status` ascending
	CampaignSortEnumStatusDesc           CampaignSortEnum = "status_DESC"            // Sort by `status` descending
	CampaignSortEnumTargetDateAsc        CampaignSortEnum = "target_date_ASC"        // Sort by `targetDate` ascending
	CampaignSortEnumTargetDateDesc       CampaignSortEnum = "target_date_DESC"       // Sort by `targetDate` descending
)

// All CampaignSortEnum as []string
var AllCampaignSortEnum = []string{
	string(CampaignSortEnumChecksPassingAsc),
	string(CampaignSortEnumChecksPassingDesc),
	string(CampaignSortEnumEndedDateAsc),
	string(CampaignSortEnumEndedDateDesc),
	string(CampaignSortEnumFilterAsc),
	string(CampaignSortEnumFilterDesc),
	string(CampaignSortEnumNameAsc),
	string(CampaignSortEnumNameDesc),
	string(CampaignSortEnumOwnerAsc),
	string(CampaignSortEnumOwnerDesc),
	string(CampaignSortEnumServicesCompleteAsc),
	string(CampaignSortEnumServicesCompleteDesc),
	string(CampaignSortEnumStartDateAsc),
	string(CampaignSortEnumStartDateDesc),
	string(CampaignSortEnumStatusAsc),
	string(CampaignSortEnumStatusDesc),
	string(CampaignSortEnumTargetDateAsc),
	string(CampaignSortEnumTargetDateDesc),
}

// CampaignStatusEnum The campaign status
type CampaignStatusEnum string

var (
	CampaignStatusEnumDelayed    CampaignStatusEnum = "delayed"     // Campaign is delayed
	CampaignStatusEnumDraft      CampaignStatusEnum = "draft"       // Campaign has been created but is not yet active
	CampaignStatusEnumEnded      CampaignStatusEnum = "ended"       // Campaign ended
	CampaignStatusEnumInProgress CampaignStatusEnum = "in_progress" // Campaign is in progress
	CampaignStatusEnumScheduled  CampaignStatusEnum = "scheduled"   // Campaign has been scheduled to begin in the future
)

// All CampaignStatusEnum as []string
var AllCampaignStatusEnum = []string{
	string(CampaignStatusEnumDelayed),
	string(CampaignStatusEnumDraft),
	string(CampaignStatusEnumEnded),
	string(CampaignStatusEnumInProgress),
	string(CampaignStatusEnumScheduled),
}

// CheckCodeIssueConstraintEnum The values allowed for the constraint type for the code issues check
type CheckCodeIssueConstraintEnum string

var (
	CheckCodeIssueConstraintEnumAny      CheckCodeIssueConstraintEnum = "any"      // The check will look for any code issues regardless of issue name
	CheckCodeIssueConstraintEnumContains CheckCodeIssueConstraintEnum = "contains" // The check will look for any code issues by name containing the issue name
	CheckCodeIssueConstraintEnumExact    CheckCodeIssueConstraintEnum = "exact"    // The check will look for any code issues matching the issue name exactly
)

// All CheckCodeIssueConstraintEnum as []string
var AllCheckCodeIssueConstraintEnum = []string{
	string(CheckCodeIssueConstraintEnumAny),
	string(CheckCodeIssueConstraintEnumContains),
	string(CheckCodeIssueConstraintEnumExact),
}

// CheckResultStatusEnum The status of the check result
type CheckResultStatusEnum string

var (
	CheckResultStatusEnumFailed CheckResultStatusEnum = "failed" // Indicates that the check has failed for the associated service
	CheckResultStatusEnumPassed CheckResultStatusEnum = "passed" // Indicates that the check has passed for the associated service.
)

// All CheckResultStatusEnum as []string
var AllCheckResultStatusEnum = []string{
	string(CheckResultStatusEnumFailed),
	string(CheckResultStatusEnumPassed),
}

// CheckStatus The evaluation status of the check
type CheckStatus string

var (
	CheckStatusFailed  CheckStatus = "failed"  // The check evaluated to a falsy value based on some conditions
	CheckStatusPassed  CheckStatus = "passed"  // The check evaluated to a truthy value based on some conditions
	CheckStatusPending CheckStatus = "pending" // The check has not been evaluated yet.
)

// All CheckStatus as []string
var AllCheckStatus = []string{
	string(CheckStatusFailed),
	string(CheckStatusPassed),
	string(CheckStatusPending),
}

// CheckType The type of check
type CheckType string

var (
	CheckTypeAlertSourceUsage    CheckType = "alert_source_usage"    // Verifies that the service has an alert source of a particular type or name
	CheckTypeCodeIssue           CheckType = "code_issue"            // Verifies that the severity and quantity of code issues does not exceed defined thresholds
	CheckTypeCustom              CheckType = "custom"                // Allows for the creation of programmatic checks that use an API to mark the status as passing or failing
	CheckTypeGeneric             CheckType = "generic"               // Requires a generic integration api call to complete a series of checks for multiple services
	CheckTypeGitBranchProtection CheckType = "git_branch_protection" // Verifies that all the repositories on the service have branch protection enabled
	CheckTypeHasDocumentation    CheckType = "has_documentation"     // Verifies that the service has visible documentation of a particular type and subtype
	CheckTypeHasOwner            CheckType = "has_owner"             // Verifies that the service has an owner defined
	CheckTypeHasRecentDeploy     CheckType = "has_recent_deploy"     // Verifies that the services has received a deploy within a specified number of days
	CheckTypeHasRepository       CheckType = "has_repository"        // Verifies that the service has a repository integrated
	CheckTypeHasServiceConfig    CheckType = "has_service_config"    // Verifies that the service is maintained though the use of an opslevel.yml service config
	CheckTypeManual              CheckType = "manual"                // Requires a service owner to manually complete a check for the service
	CheckTypePackageVersion      CheckType = "package_version"       // Verifies certain aspects of a service using or not using software packages
	CheckTypePayload             CheckType = "payload"               // Requires a payload integration api call to complete a check for the service
	CheckTypeRepoFile            CheckType = "repo_file"             // Quickly scan the service’s repository for the existence or contents of a specific file
	CheckTypeRepoGrep            CheckType = "repo_grep"             // Run a comprehensive search across the service's repository using advanced search parameters
	CheckTypeRepoSearch          CheckType = "repo_search"           // Quickly search the service’s repository for specific contents in any file
	CheckTypeServiceDependency   CheckType = "service_dependency"    // Verifies that the service has either a dependent or dependency
	CheckTypeServiceProperty     CheckType = "service_property"      // Verifies that a service property is set or matches a specified format
	CheckTypeTagDefined          CheckType = "tag_defined"           // Verifies that the service has the specified tag defined
	CheckTypeToolUsage           CheckType = "tool_usage"            // Verifies that the service is using a tool of a particular category or name
)

// All CheckType as []string
var AllCheckType = []string{
	string(CheckTypeAlertSourceUsage),
	string(CheckTypeCodeIssue),
	string(CheckTypeCustom),
	string(CheckTypeGeneric),
	string(CheckTypeGitBranchProtection),
	string(CheckTypeHasDocumentation),
	string(CheckTypeHasOwner),
	string(CheckTypeHasRecentDeploy),
	string(CheckTypeHasRepository),
	string(CheckTypeHasServiceConfig),
	string(CheckTypeManual),
	string(CheckTypePackageVersion),
	string(CheckTypePayload),
	string(CheckTypeRepoFile),
	string(CheckTypeRepoGrep),
	string(CheckTypeRepoSearch),
	string(CheckTypeServiceDependency),
	string(CheckTypeServiceProperty),
	string(CheckTypeTagDefined),
	string(CheckTypeToolUsage),
}

// CodeIssueResolutionTimeUnitEnum The allowed values for duration units for the resolution time
type CodeIssueResolutionTimeUnitEnum string

var (
	CodeIssueResolutionTimeUnitEnumDay   CodeIssueResolutionTimeUnitEnum = "day"   // Day, as a duration
	CodeIssueResolutionTimeUnitEnumMonth CodeIssueResolutionTimeUnitEnum = "month" // Month, as a duration
	CodeIssueResolutionTimeUnitEnumWeek  CodeIssueResolutionTimeUnitEnum = "week"  // Week, as a duration
)

// All CodeIssueResolutionTimeUnitEnum as []string
var AllCodeIssueResolutionTimeUnitEnum = []string{
	string(CodeIssueResolutionTimeUnitEnumDay),
	string(CodeIssueResolutionTimeUnitEnumMonth),
	string(CodeIssueResolutionTimeUnitEnumWeek),
}

// ConnectiveEnum The logical operator to be used in conjunction with multiple filters (requires filters to be supplied)
type ConnectiveEnum string

var (
	ConnectiveEnumAnd ConnectiveEnum = "and" // Used to ensure **all** filters match for a given resource
	ConnectiveEnumOr  ConnectiveEnum = "or"  // Used to ensure **any** filters match for a given resource
)

// All ConnectiveEnum as []string
var AllConnectiveEnum = []string{
	string(ConnectiveEnumAnd),
	string(ConnectiveEnumOr),
}

// ContactType The method of contact
type ContactType string

var (
	ContactTypeEmail          ContactType = "email"           // An email contact method
	ContactTypeGitHub         ContactType = "github"          // A GitHub handle
	ContactTypeMicrosoftTeams ContactType = "microsoft_teams" // A Microsoft Teams channel
	ContactTypeSlack          ContactType = "slack"           // A Slack channel contact method
	ContactTypeSlackHandle    ContactType = "slack_handle"    // A Slack handle contact method
	ContactTypeWeb            ContactType = "web"             // A website contact method
)

// All ContactType as []string
var AllContactType = []string{
	string(ContactTypeEmail),
	string(ContactTypeGitHub),
	string(ContactTypeMicrosoftTeams),
	string(ContactTypeSlack),
	string(ContactTypeSlackHandle),
	string(ContactTypeWeb),
}

// CustomActionsEntityTypeEnum The entity types a custom action can be associated with
type CustomActionsEntityTypeEnum string

var (
	CustomActionsEntityTypeEnumGlobal  CustomActionsEntityTypeEnum = "GLOBAL"  // A custom action associated with the global scope (no particular entity type)
	CustomActionsEntityTypeEnumService CustomActionsEntityTypeEnum = "SERVICE" // A custom action associated with services
)

// All CustomActionsEntityTypeEnum as []string
var AllCustomActionsEntityTypeEnum = []string{
	string(CustomActionsEntityTypeEnumGlobal),
	string(CustomActionsEntityTypeEnumService),
}

// CustomActionsHttpMethodEnum An HTTP request method
type CustomActionsHttpMethodEnum string

var (
	CustomActionsHttpMethodEnumDelete CustomActionsHttpMethodEnum = "DELETE" // An HTTP DELETE request
	CustomActionsHttpMethodEnumGet    CustomActionsHttpMethodEnum = "GET"    // An HTTP GET request
	CustomActionsHttpMethodEnumPatch  CustomActionsHttpMethodEnum = "PATCH"  // An HTTP PATCH request
	CustomActionsHttpMethodEnumPost   CustomActionsHttpMethodEnum = "POST"   // An HTTP POST request
	CustomActionsHttpMethodEnumPut    CustomActionsHttpMethodEnum = "PUT"    // An HTTP PUT request
)

// All CustomActionsHttpMethodEnum as []string
var AllCustomActionsHttpMethodEnum = []string{
	string(CustomActionsHttpMethodEnumDelete),
	string(CustomActionsHttpMethodEnumGet),
	string(CustomActionsHttpMethodEnumPatch),
	string(CustomActionsHttpMethodEnumPost),
	string(CustomActionsHttpMethodEnumPut),
}

// CustomActionsTriggerDefinitionAccessControlEnum Who can see and use the trigger definition
type CustomActionsTriggerDefinitionAccessControlEnum string

var (
	CustomActionsTriggerDefinitionAccessControlEnumAdmins        CustomActionsTriggerDefinitionAccessControlEnum = "admins"         // Admin users
	CustomActionsTriggerDefinitionAccessControlEnumEveryone      CustomActionsTriggerDefinitionAccessControlEnum = "everyone"       // All users of OpsLevel
	CustomActionsTriggerDefinitionAccessControlEnumServiceOwners CustomActionsTriggerDefinitionAccessControlEnum = "service_owners" // The owners of a service
)

// All CustomActionsTriggerDefinitionAccessControlEnum as []string
var AllCustomActionsTriggerDefinitionAccessControlEnum = []string{
	string(CustomActionsTriggerDefinitionAccessControlEnumAdmins),
	string(CustomActionsTriggerDefinitionAccessControlEnumEveryone),
	string(CustomActionsTriggerDefinitionAccessControlEnumServiceOwners),
}

// CustomActionsTriggerEventStatusEnum The status of the custom action trigger event
type CustomActionsTriggerEventStatusEnum string

var (
	CustomActionsTriggerEventStatusEnumFailure CustomActionsTriggerEventStatusEnum = "FAILURE" // The action failed to complete
	CustomActionsTriggerEventStatusEnumPending CustomActionsTriggerEventStatusEnum = "PENDING" // A result has not been determined
	CustomActionsTriggerEventStatusEnumSuccess CustomActionsTriggerEventStatusEnum = "SUCCESS" // The action completed successfully
)

// All CustomActionsTriggerEventStatusEnum as []string
var AllCustomActionsTriggerEventStatusEnum = []string{
	string(CustomActionsTriggerEventStatusEnumFailure),
	string(CustomActionsTriggerEventStatusEnumPending),
	string(CustomActionsTriggerEventStatusEnumSuccess),
}

// DayOfWeekEnum Possible days of the week
type DayOfWeekEnum string

var (
	DayOfWeekEnumFriday    DayOfWeekEnum = "friday"    // Yesterday was Thursday. Tomorrow is Saturday. We so excited
	DayOfWeekEnumMonday    DayOfWeekEnum = "monday"    // Monday is the day of the week that takes place between Sunday and Tuesday
	DayOfWeekEnumSaturday  DayOfWeekEnum = "saturday"  // The day of the week before Sunday and following Friday, and (together with Sunday) forming part of the weekend
	DayOfWeekEnumSunday    DayOfWeekEnum = "sunday"    // The day of the week before Monday and following Saturday, (together with Saturday) forming part of the weekend
	DayOfWeekEnumThursday  DayOfWeekEnum = "thursday"  // The day of the week before Friday and following Wednesday
	DayOfWeekEnumTuesday   DayOfWeekEnum = "tuesday"   // Tuesday is the day of the week between Monday and Wednesday
	DayOfWeekEnumWednesday DayOfWeekEnum = "wednesday" // The day of the week before Thursday and following Tuesday
)

// All DayOfWeekEnum as []string
var AllDayOfWeekEnum = []string{
	string(DayOfWeekEnumFriday),
	string(DayOfWeekEnumMonday),
	string(DayOfWeekEnumSaturday),
	string(DayOfWeekEnumSunday),
	string(DayOfWeekEnumThursday),
	string(DayOfWeekEnumTuesday),
	string(DayOfWeekEnumWednesday),
}

// EventIntegrationEnum The type of event integration
type EventIntegrationEnum string

var (
	EventIntegrationEnumApidoc        EventIntegrationEnum = "apiDoc"        // API Documentation integration
	EventIntegrationEnumAquasecurity  EventIntegrationEnum = "aquaSecurity"  // Aqua Security Custom Event Check integration
	EventIntegrationEnumArgocd        EventIntegrationEnum = "argocd"        // ArgoCD deploy integration
	EventIntegrationEnumAwsecr        EventIntegrationEnum = "awsEcr"        // AWS ECR Custom Event Check integration
	EventIntegrationEnumBugsnag       EventIntegrationEnum = "bugsnag"       // Bugsnag Custom Event Check integration
	EventIntegrationEnumCircleci      EventIntegrationEnum = "circleci"      // CircleCI deploy integration
	EventIntegrationEnumCodacy        EventIntegrationEnum = "codacy"        // Codacy Custom Event Check integration
	EventIntegrationEnumCoveralls     EventIntegrationEnum = "coveralls"     // Coveralls Custom Event Check integration
	EventIntegrationEnumCustomevent   EventIntegrationEnum = "customEvent"   // Custom Event integration
	EventIntegrationEnumDatadogcheck  EventIntegrationEnum = "datadogCheck"  // Datadog Check integration
	EventIntegrationEnumDeploy        EventIntegrationEnum = "deploy"        // Deploy integration
	EventIntegrationEnumDynatrace     EventIntegrationEnum = "dynatrace"     // Dynatrace Custom Event Check integration
	EventIntegrationEnumFlux          EventIntegrationEnum = "flux"          // Flux deploy integration
	EventIntegrationEnumGithubactions EventIntegrationEnum = "githubActions" // Github Actions deploy integration
	EventIntegrationEnumGitlabci      EventIntegrationEnum = "gitlabCi"      // Gitlab CI deploy integration
	EventIntegrationEnumGrafana       EventIntegrationEnum = "grafana"       // Grafana Custom Event Check integration
	EventIntegrationEnumGrype         EventIntegrationEnum = "grype"         // Grype Custom Event Check integration
	EventIntegrationEnumJenkins       EventIntegrationEnum = "jenkins"       // Jenkins deploy integration
	EventIntegrationEnumJfrogxray     EventIntegrationEnum = "jfrogXray"     // JFrog Xray Custom Event Check integration
	EventIntegrationEnumLacework      EventIntegrationEnum = "lacework"      // Lacework Custom Event Check integration
	EventIntegrationEnumNewreliccheck EventIntegrationEnum = "newRelicCheck" // New Relic Check integration
	EventIntegrationEnumOctopus       EventIntegrationEnum = "octopus"       // Octopus deploy integration
	EventIntegrationEnumPrismacloud   EventIntegrationEnum = "prismaCloud"   // Prisma Cloud Custom Event Check integration
	EventIntegrationEnumPrometheus    EventIntegrationEnum = "prometheus"    // Prometheus Custom Event Check integration
	EventIntegrationEnumRollbar       EventIntegrationEnum = "rollbar"       // Rollbar Custom Event Check integration
	EventIntegrationEnumSentry        EventIntegrationEnum = "sentry"        // Sentry Custom Event Check integration
	EventIntegrationEnumSnyk          EventIntegrationEnum = "snyk"          // Snyk Custom Event Check integration
	EventIntegrationEnumSonarqube     EventIntegrationEnum = "sonarqube"     // SonarQube Custom Event Check integration
	EventIntegrationEnumStackhawk     EventIntegrationEnum = "stackhawk"     // StackHawk Custom Event Check integration
	EventIntegrationEnumSumologic     EventIntegrationEnum = "sumoLogic"     // Sumo Logic Custom Event Check integration
	EventIntegrationEnumVeracode      EventIntegrationEnum = "veracode"      // Veracode Custom Event Check integration
)

// All EventIntegrationEnum as []string
var AllEventIntegrationEnum = []string{
	string(EventIntegrationEnumApidoc),
	string(EventIntegrationEnumAquasecurity),
	string(EventIntegrationEnumArgocd),
	string(EventIntegrationEnumAwsecr),
	string(EventIntegrationEnumBugsnag),
	string(EventIntegrationEnumCircleci),
	string(EventIntegrationEnumCodacy),
	string(EventIntegrationEnumCoveralls),
	string(EventIntegrationEnumCustomevent),
	string(EventIntegrationEnumDatadogcheck),
	string(EventIntegrationEnumDeploy),
	string(EventIntegrationEnumDynatrace),
	string(EventIntegrationEnumFlux),
	string(EventIntegrationEnumGithubactions),
	string(EventIntegrationEnumGitlabci),
	string(EventIntegrationEnumGrafana),
	string(EventIntegrationEnumGrype),
	string(EventIntegrationEnumJenkins),
	string(EventIntegrationEnumJfrogxray),
	string(EventIntegrationEnumLacework),
	string(EventIntegrationEnumNewreliccheck),
	string(EventIntegrationEnumOctopus),
	string(EventIntegrationEnumPrismacloud),
	string(EventIntegrationEnumPrometheus),
	string(EventIntegrationEnumRollbar),
	string(EventIntegrationEnumSentry),
	string(EventIntegrationEnumSnyk),
	string(EventIntegrationEnumSonarqube),
	string(EventIntegrationEnumStackhawk),
	string(EventIntegrationEnumSumologic),
	string(EventIntegrationEnumVeracode),
}

// FrequencyTimeScale The time scale type for the frequency
type FrequencyTimeScale string

var (
	FrequencyTimeScaleDay   FrequencyTimeScale = "day"   // Consider the time scale of days
	FrequencyTimeScaleMonth FrequencyTimeScale = "month" // Consider the time scale of months
	FrequencyTimeScaleWeek  FrequencyTimeScale = "week"  // Consider the time scale of weeks
	FrequencyTimeScaleYear  FrequencyTimeScale = "year"  // Consider the time scale of years
)

// All FrequencyTimeScale as []string
var AllFrequencyTimeScale = []string{
	string(FrequencyTimeScaleDay),
	string(FrequencyTimeScaleMonth),
	string(FrequencyTimeScaleWeek),
	string(FrequencyTimeScaleYear),
}

// HasDocumentationSubtypeEnum The subtype of the document
type HasDocumentationSubtypeEnum string

var (
	HasDocumentationSubtypeEnumOpenapi HasDocumentationSubtypeEnum = "openapi" // Document is an OpenAPI document
)

// All HasDocumentationSubtypeEnum as []string
var AllHasDocumentationSubtypeEnum = []string{
	string(HasDocumentationSubtypeEnumOpenapi),
}

// HasDocumentationTypeEnum The type of the document
type HasDocumentationTypeEnum string

var (
	HasDocumentationTypeEnumAPI  HasDocumentationTypeEnum = "api"  // Document is an API document
	HasDocumentationTypeEnumTech HasDocumentationTypeEnum = "tech" // Document is a Tech document
)

// All HasDocumentationTypeEnum as []string
var AllHasDocumentationTypeEnum = []string{
	string(HasDocumentationTypeEnumAPI),
	string(HasDocumentationTypeEnumTech),
}

// PackageConstraintEnum Possible values of a package version check constraint
type PackageConstraintEnum string

var (
	PackageConstraintEnumDoesNotExist   PackageConstraintEnum = "does_not_exist"  // The package must not be used by a service
	PackageConstraintEnumExists         PackageConstraintEnum = "exists"          // The package must be used by a service
	PackageConstraintEnumMatchesVersion PackageConstraintEnum = "matches_version" // The package usage by a service must match certain specified version constraints
)

// All PackageConstraintEnum as []string
var AllPackageConstraintEnum = []string{
	string(PackageConstraintEnumDoesNotExist),
	string(PackageConstraintEnumExists),
	string(PackageConstraintEnumMatchesVersion),
}

// PackageManagerEnum Supported software package manager types
type PackageManagerEnum string

var (
	PackageManagerEnumAlpm      PackageManagerEnum = "alpm"      //
	PackageManagerEnumApk       PackageManagerEnum = "apk"       //
	PackageManagerEnumBitbucket PackageManagerEnum = "bitbucket" //
	PackageManagerEnumBitnami   PackageManagerEnum = "bitnami"   //
	PackageManagerEnumCargo     PackageManagerEnum = "cargo"     //
	PackageManagerEnumCocoapods PackageManagerEnum = "cocoapods" //
	PackageManagerEnumComposer  PackageManagerEnum = "composer"  //
	PackageManagerEnumConan     PackageManagerEnum = "conan"     //
	PackageManagerEnumConda     PackageManagerEnum = "conda"     //
	PackageManagerEnumCpan      PackageManagerEnum = "cpan"      //
	PackageManagerEnumCran      PackageManagerEnum = "cran"      //
	PackageManagerEnumDeb       PackageManagerEnum = "deb"       //
	PackageManagerEnumDocker    PackageManagerEnum = "docker"    //
	PackageManagerEnumGem       PackageManagerEnum = "gem"       //
	PackageManagerEnumGeneric   PackageManagerEnum = "generic"   //
	PackageManagerEnumGitHub    PackageManagerEnum = "github"    //
	PackageManagerEnumGolang    PackageManagerEnum = "golang"    //
	PackageManagerEnumGradle    PackageManagerEnum = "gradle"    //
	PackageManagerEnumHackage   PackageManagerEnum = "hackage"   //
	PackageManagerEnumHelm      PackageManagerEnum = "helm"      //
	PackageManagerEnumHex       PackageManagerEnum = "hex"       //
	PackageManagerEnumMaven     PackageManagerEnum = "maven"     //
	PackageManagerEnumMlflow    PackageManagerEnum = "mlflow"    //
	PackageManagerEnumNpm       PackageManagerEnum = "npm"       //
	PackageManagerEnumNuget     PackageManagerEnum = "nuget"     //
	PackageManagerEnumOci       PackageManagerEnum = "oci"       //
	PackageManagerEnumPub       PackageManagerEnum = "pub"       //
	PackageManagerEnumPypi      PackageManagerEnum = "pypi"      //
	PackageManagerEnumQpkg      PackageManagerEnum = "qpkg"      //
	PackageManagerEnumRpm       PackageManagerEnum = "rpm"       //
	PackageManagerEnumSwid      PackageManagerEnum = "swid"      //
	PackageManagerEnumSwift     PackageManagerEnum = "swift"     //
)

// All PackageManagerEnum as []string
var AllPackageManagerEnum = []string{
	string(PackageManagerEnumAlpm),
	string(PackageManagerEnumApk),
	string(PackageManagerEnumBitbucket),
	string(PackageManagerEnumBitnami),
	string(PackageManagerEnumCargo),
	string(PackageManagerEnumCocoapods),
	string(PackageManagerEnumComposer),
	string(PackageManagerEnumConan),
	string(PackageManagerEnumConda),
	string(PackageManagerEnumCpan),
	string(PackageManagerEnumCran),
	string(PackageManagerEnumDeb),
	string(PackageManagerEnumDocker),
	string(PackageManagerEnumGem),
	string(PackageManagerEnumGeneric),
	string(PackageManagerEnumGitHub),
	string(PackageManagerEnumGolang),
	string(PackageManagerEnumGradle),
	string(PackageManagerEnumHackage),
	string(PackageManagerEnumHelm),
	string(PackageManagerEnumHex),
	string(PackageManagerEnumMaven),
	string(PackageManagerEnumMlflow),
	string(PackageManagerEnumNpm),
	string(PackageManagerEnumNuget),
	string(PackageManagerEnumOci),
	string(PackageManagerEnumPub),
	string(PackageManagerEnumPypi),
	string(PackageManagerEnumQpkg),
	string(PackageManagerEnumRpm),
	string(PackageManagerEnumSwid),
	string(PackageManagerEnumSwift),
}

// PayloadFilterEnum Fields that can be used as part of filters for payloads
type PayloadFilterEnum string

var (
	PayloadFilterEnumIntegrationID PayloadFilterEnum = "integration_id" // Filter by `integration` field. Note that this is an internal id, ex. "123"
)

// All PayloadFilterEnum as []string
var AllPayloadFilterEnum = []string{
	string(PayloadFilterEnumIntegrationID),
}

// PayloadSortEnum Sort possibilities for payloads
type PayloadSortEnum string

var (
	PayloadSortEnumCreatedAtAsc    PayloadSortEnum = "created_at_ASC"    // Order by `created_at` ascending
	PayloadSortEnumCreatedAtDesc   PayloadSortEnum = "created_at_DESC"   // Order by `created_at` descending
	PayloadSortEnumProcessedAtAsc  PayloadSortEnum = "processed_at_ASC"  // Order by `processed_at` ascending
	PayloadSortEnumProcessedAtDesc PayloadSortEnum = "processed_at_DESC" // Order by `processed_at` descending
)

// All PayloadSortEnum as []string
var AllPayloadSortEnum = []string{
	string(PayloadSortEnumCreatedAtAsc),
	string(PayloadSortEnumCreatedAtDesc),
	string(PayloadSortEnumProcessedAtAsc),
	string(PayloadSortEnumProcessedAtDesc),
}

// PredicateKeyEnum Fields that can be used as part of filter for services
type PredicateKeyEnum string

var (
	PredicateKeyEnumAliases        PredicateKeyEnum = "aliases"         // Filter by Alias attached to this service, if any
	PredicateKeyEnumCreationSource PredicateKeyEnum = "creation_source" // Filter by the creation source
	PredicateKeyEnumDomainID       PredicateKeyEnum = "domain_id"       // Filter by Domain that includes the System this service is assigned to, if any
	PredicateKeyEnumFilterID       PredicateKeyEnum = "filter_id"       // Filter by another filter
	PredicateKeyEnumFramework      PredicateKeyEnum = "framework"       // Filter by `framework` field
	PredicateKeyEnumGroupIDs       PredicateKeyEnum = "group_ids"       // Filter by group hierarchy. Will return resources who's owner is in the group ancestry chain
	PredicateKeyEnumLanguage       PredicateKeyEnum = "language"        // Filter by `language` field
	PredicateKeyEnumLifecycleIndex PredicateKeyEnum = "lifecycle_index" // Filter by `lifecycle` field
	PredicateKeyEnumName           PredicateKeyEnum = "name"            // Filter by `name` field
	PredicateKeyEnumOwnerID        PredicateKeyEnum = "owner_id"        // Filter by `owner` field
	PredicateKeyEnumOwnerIDs       PredicateKeyEnum = "owner_ids"       // Filter by `owner` hierarchy. Will return resources who's owner is in the team ancestry chain
	PredicateKeyEnumProduct        PredicateKeyEnum = "product"         // Filter by `product` field
	PredicateKeyEnumProperties     PredicateKeyEnum = "properties"      // Filter by custom-defined properties
	PredicateKeyEnumRepositoryIDs  PredicateKeyEnum = "repository_ids"  // Filter by Repository that this service is attached to, if any
	PredicateKeyEnumSystemID       PredicateKeyEnum = "system_id"       // Filter by System that this service is assigned to, if any
	PredicateKeyEnumTags           PredicateKeyEnum = "tags"            // Filter by `tags` field
	PredicateKeyEnumTierIndex      PredicateKeyEnum = "tier_index"      // Filter by `tier` field
)

// All PredicateKeyEnum as []string
var AllPredicateKeyEnum = []string{
	string(PredicateKeyEnumAliases),
	string(PredicateKeyEnumCreationSource),
	string(PredicateKeyEnumDomainID),
	string(PredicateKeyEnumFilterID),
	string(PredicateKeyEnumFramework),
	string(PredicateKeyEnumGroupIDs),
	string(PredicateKeyEnumLanguage),
	string(PredicateKeyEnumLifecycleIndex),
	string(PredicateKeyEnumName),
	string(PredicateKeyEnumOwnerID),
	string(PredicateKeyEnumOwnerIDs),
	string(PredicateKeyEnumProduct),
	string(PredicateKeyEnumProperties),
	string(PredicateKeyEnumRepositoryIDs),
	string(PredicateKeyEnumSystemID),
	string(PredicateKeyEnumTags),
	string(PredicateKeyEnumTierIndex),
}

// PredicateTypeEnum Operations that can be used on predicates
type PredicateTypeEnum string

var (
	PredicateTypeEnumBelongsTo                  PredicateTypeEnum = "belongs_to"                   // Belongs to a group's hierarchy
	PredicateTypeEnumContains                   PredicateTypeEnum = "contains"                     // Contains a specific value
	PredicateTypeEnumDoesNotContain             PredicateTypeEnum = "does_not_contain"             // Does not contain a specific value
	PredicateTypeEnumDoesNotEqual               PredicateTypeEnum = "does_not_equal"               // Does not equal a specific value
	PredicateTypeEnumDoesNotExist               PredicateTypeEnum = "does_not_exist"               // Specific attribute does not exist
	PredicateTypeEnumDoesNotMatch               PredicateTypeEnum = "does_not_match"               // A certain filter is not matched
	PredicateTypeEnumDoesNotMatchRegex          PredicateTypeEnum = "does_not_match_regex"         // Does not match a value using a regular expression
	PredicateTypeEnumEndsWith                   PredicateTypeEnum = "ends_with"                    // Ends with a specific value
	PredicateTypeEnumEquals                     PredicateTypeEnum = "equals"                       // Equals a specific value
	PredicateTypeEnumExists                     PredicateTypeEnum = "exists"                       // Specific attribute exists
	PredicateTypeEnumGreaterThanOrEqualTo       PredicateTypeEnum = "greater_than_or_equal_to"     // Greater than or equal to a specific value (numeric only)
	PredicateTypeEnumLessThanOrEqualTo          PredicateTypeEnum = "less_than_or_equal_to"        // Less than or equal to a specific value (numeric only)
	PredicateTypeEnumMatches                    PredicateTypeEnum = "matches"                      // A certain filter is matched
	PredicateTypeEnumMatchesRegex               PredicateTypeEnum = "matches_regex"                // Matches a value using a regular expression
	PredicateTypeEnumSatisfiesJqExpression      PredicateTypeEnum = "satisfies_jq_expression"      // Satisfies an expression defined in jq
	PredicateTypeEnumSatisfiesVersionConstraint PredicateTypeEnum = "satisfies_version_constraint" // Satisfies version constraint (tag value only)
	PredicateTypeEnumStartsWith                 PredicateTypeEnum = "starts_with"                  // Starts with a specific value
)

// All PredicateTypeEnum as []string
var AllPredicateTypeEnum = []string{
	string(PredicateTypeEnumBelongsTo),
	string(PredicateTypeEnumContains),
	string(PredicateTypeEnumDoesNotContain),
	string(PredicateTypeEnumDoesNotEqual),
	string(PredicateTypeEnumDoesNotExist),
	string(PredicateTypeEnumDoesNotMatch),
	string(PredicateTypeEnumDoesNotMatchRegex),
	string(PredicateTypeEnumEndsWith),
	string(PredicateTypeEnumEquals),
	string(PredicateTypeEnumExists),
	string(PredicateTypeEnumGreaterThanOrEqualTo),
	string(PredicateTypeEnumLessThanOrEqualTo),
	string(PredicateTypeEnumMatches),
	string(PredicateTypeEnumMatchesRegex),
	string(PredicateTypeEnumSatisfiesJqExpression),
	string(PredicateTypeEnumSatisfiesVersionConstraint),
	string(PredicateTypeEnumStartsWith),
}

// PropertyDefinitionDisplayTypeEnum The set of possible display types of a property definition schema
type PropertyDefinitionDisplayTypeEnum string

var (
	PropertyDefinitionDisplayTypeEnumArray    PropertyDefinitionDisplayTypeEnum = "ARRAY"    // An array
	PropertyDefinitionDisplayTypeEnumBoolean  PropertyDefinitionDisplayTypeEnum = "BOOLEAN"  // A boolean
	PropertyDefinitionDisplayTypeEnumDropdown PropertyDefinitionDisplayTypeEnum = "DROPDOWN" // A dropdown
	PropertyDefinitionDisplayTypeEnumNumber   PropertyDefinitionDisplayTypeEnum = "NUMBER"   // A number
	PropertyDefinitionDisplayTypeEnumObject   PropertyDefinitionDisplayTypeEnum = "OBJECT"   // An object
	PropertyDefinitionDisplayTypeEnumText     PropertyDefinitionDisplayTypeEnum = "TEXT"     // A text string
)

// All PropertyDefinitionDisplayTypeEnum as []string
var AllPropertyDefinitionDisplayTypeEnum = []string{
	string(PropertyDefinitionDisplayTypeEnumArray),
	string(PropertyDefinitionDisplayTypeEnumBoolean),
	string(PropertyDefinitionDisplayTypeEnumDropdown),
	string(PropertyDefinitionDisplayTypeEnumNumber),
	string(PropertyDefinitionDisplayTypeEnumObject),
	string(PropertyDefinitionDisplayTypeEnumText),
}

// PropertyDisplayStatusEnum The display status of a custom property on service pages
type PropertyDisplayStatusEnum string

var (
	PropertyDisplayStatusEnumHidden  PropertyDisplayStatusEnum = "hidden"  // The property is not shown on the service page
	PropertyDisplayStatusEnumVisible PropertyDisplayStatusEnum = "visible" // The property is shown on the service page
)

// All PropertyDisplayStatusEnum as []string
var AllPropertyDisplayStatusEnum = []string{
	string(PropertyDisplayStatusEnumHidden),
	string(PropertyDisplayStatusEnumVisible),
}

// PropertyLockedStatusEnum Values for which lock is assigned to a property definition to restrict what sources can assign values to it
type PropertyLockedStatusEnum string

var (
	PropertyLockedStatusEnumUILocked PropertyLockedStatusEnum = "ui_locked" // Value assignments on the property cannot be changed through the UI
	PropertyLockedStatusEnumUnlocked PropertyLockedStatusEnum = "unlocked"  // There are no restrictions on what sources can assign values to the property
)

// All PropertyLockedStatusEnum as []string
var AllPropertyLockedStatusEnum = []string{
	string(PropertyLockedStatusEnumUILocked),
	string(PropertyLockedStatusEnumUnlocked),
}

// PropertyOwnerTypeEnum The possible entity types that a property can be assigned to
type PropertyOwnerTypeEnum string

var (
	PropertyOwnerTypeEnumComponent PropertyOwnerTypeEnum = "COMPONENT" // A component
	PropertyOwnerTypeEnumTeam      PropertyOwnerTypeEnum = "TEAM"      // A team
)

// All PropertyOwnerTypeEnum as []string
var AllPropertyOwnerTypeEnum = []string{
	string(PropertyOwnerTypeEnumComponent),
	string(PropertyOwnerTypeEnumTeam),
}

// RelatedResourceRelationshipTypeEnum The type of the relationship between two resources
type RelatedResourceRelationshipTypeEnum string

var (
	RelatedResourceRelationshipTypeEnumBelongsTo    RelatedResourceRelationshipTypeEnum = "belongs_to"    // The resource belongs to the node on the edge
	RelatedResourceRelationshipTypeEnumContains     RelatedResourceRelationshipTypeEnum = "contains"      // The resource contains the node on the edge
	RelatedResourceRelationshipTypeEnumDependencyOf RelatedResourceRelationshipTypeEnum = "dependency_of" // The resource is a dependency of the node on the edge
	RelatedResourceRelationshipTypeEnumDependsOn    RelatedResourceRelationshipTypeEnum = "depends_on"    // The resource depends on the node on the edge
	RelatedResourceRelationshipTypeEnumMemberOf     RelatedResourceRelationshipTypeEnum = "member_of"     // The resource is a member of the node on the edge
)

// All RelatedResourceRelationshipTypeEnum as []string
var AllRelatedResourceRelationshipTypeEnum = []string{
	string(RelatedResourceRelationshipTypeEnumBelongsTo),
	string(RelatedResourceRelationshipTypeEnumContains),
	string(RelatedResourceRelationshipTypeEnumDependencyOf),
	string(RelatedResourceRelationshipTypeEnumDependsOn),
	string(RelatedResourceRelationshipTypeEnumMemberOf),
}

// RelationshipTypeEnum The type of relationship between two resources
type RelationshipTypeEnum string

var (
	RelationshipTypeEnumBelongsTo RelationshipTypeEnum = "belongs_to" // The source resource belongs to the target resource
	RelationshipTypeEnumDependsOn RelationshipTypeEnum = "depends_on" // The source resource depends on the target resource
)

// All RelationshipTypeEnum as []string
var AllRelationshipTypeEnum = []string{
	string(RelationshipTypeEnumBelongsTo),
	string(RelationshipTypeEnumDependsOn),
}

// RepositoryVisibilityEnum Possible visibility levels for repositories
type RepositoryVisibilityEnum string

var (
	RepositoryVisibilityEnumInternal     RepositoryVisibilityEnum = "INTERNAL"     // Repositories that are only accessible to organization users (Github, Gitlab)
	RepositoryVisibilityEnumOrganization RepositoryVisibilityEnum = "ORGANIZATION" // Repositories that are only accessible to organization users (ADO)
	RepositoryVisibilityEnumPrivate      RepositoryVisibilityEnum = "PRIVATE"      // Repositories that are private to the user
	RepositoryVisibilityEnumPublic       RepositoryVisibilityEnum = "PUBLIC"       // Repositories that are publicly accessible
)

// All RepositoryVisibilityEnum as []string
var AllRepositoryVisibilityEnum = []string{
	string(RepositoryVisibilityEnumInternal),
	string(RepositoryVisibilityEnumOrganization),
	string(RepositoryVisibilityEnumPrivate),
	string(RepositoryVisibilityEnumPublic),
}

// ResourceDocumentStatusTypeEnum Status of a document on a resource
type ResourceDocumentStatusTypeEnum string

var (
	ResourceDocumentStatusTypeEnumHidden  ResourceDocumentStatusTypeEnum = "hidden"  // Document is hidden
	ResourceDocumentStatusTypeEnumPinned  ResourceDocumentStatusTypeEnum = "pinned"  // Document is pinned
	ResourceDocumentStatusTypeEnumVisible ResourceDocumentStatusTypeEnum = "visible" // Document is visible
)

// All ResourceDocumentStatusTypeEnum as []string
var AllResourceDocumentStatusTypeEnum = []string{
	string(ResourceDocumentStatusTypeEnumHidden),
	string(ResourceDocumentStatusTypeEnumPinned),
	string(ResourceDocumentStatusTypeEnumVisible),
}

// ScorecardSortEnum The possible options to sort the resulting list of scorecards
type ScorecardSortEnum string

var (
	ScorecardSortEnumAffectsoverallservicelevelsAsc  ScorecardSortEnum = "affectsOverallServiceLevels_ASC"  // Order by whether or not the checks on the scorecard affect the overall maturity, in ascending order
	ScorecardSortEnumAffectsoverallservicelevelsDesc ScorecardSortEnum = "affectsOverallServiceLevels_DESC" // Order by whether or not the checks on the scorecard affect the overall maturity, in descending order
	ScorecardSortEnumFilterAsc                       ScorecardSortEnum = "filter_ASC"                       // Order by the associated filter's name, in ascending order
	ScorecardSortEnumFilterDesc                      ScorecardSortEnum = "filter_DESC"                      // Order by the associated filter's name, in descending order
	ScorecardSortEnumNameAsc                         ScorecardSortEnum = "name_ASC"                         // Order by the scorecard's name, in ascending order
	ScorecardSortEnumNameDesc                        ScorecardSortEnum = "name_DESC"                        // Order by the scorecard's name, in descending order
	ScorecardSortEnumOwnerAsc                        ScorecardSortEnum = "owner_ASC"                        // Order by the scorecard owner's name, in ascending order
	ScorecardSortEnumOwnerDesc                       ScorecardSortEnum = "owner_DESC"                       // Order by the scorecard owner's name, in descending order
	ScorecardSortEnumPassingcheckfractionAsc         ScorecardSortEnum = "passingCheckFraction_ASC"         // Order by the fraction of passing checks on the scorecard, in ascending order
	ScorecardSortEnumPassingcheckfractionDesc        ScorecardSortEnum = "passingCheckFraction_DESC"        // Order by the fraction of passing checks on the scorecard, in descending order
	ScorecardSortEnumServicecountAsc                 ScorecardSortEnum = "serviceCount_ASC"                 // Order by the number of services covered by the scorecard, in ascending order
	ScorecardSortEnumServicecountDesc                ScorecardSortEnum = "serviceCount_DESC"                // Order by the number of services covered by the scorecard, in descending order
)

// All ScorecardSortEnum as []string
var AllScorecardSortEnum = []string{
	string(ScorecardSortEnumAffectsoverallservicelevelsAsc),
	string(ScorecardSortEnumAffectsoverallservicelevelsDesc),
	string(ScorecardSortEnumFilterAsc),
	string(ScorecardSortEnumFilterDesc),
	string(ScorecardSortEnumNameAsc),
	string(ScorecardSortEnumNameDesc),
	string(ScorecardSortEnumOwnerAsc),
	string(ScorecardSortEnumOwnerDesc),
	string(ScorecardSortEnumPassingcheckfractionAsc),
	string(ScorecardSortEnumPassingcheckfractionDesc),
	string(ScorecardSortEnumServicecountAsc),
	string(ScorecardSortEnumServicecountDesc),
}

// ServicePropertyTypeEnum Properties of services that can be validated
type ServicePropertyTypeEnum string

var (
	ServicePropertyTypeEnumCustomProperty ServicePropertyTypeEnum = "custom_property" // A custom property that is associated with the service
	ServicePropertyTypeEnumDescription    ServicePropertyTypeEnum = "description"     // The description of a service
	ServicePropertyTypeEnumFramework      ServicePropertyTypeEnum = "framework"       // The primary software development framework of a service
	ServicePropertyTypeEnumLanguage       ServicePropertyTypeEnum = "language"        // The primary programming language of a service
	ServicePropertyTypeEnumLifecycleIndex ServicePropertyTypeEnum = "lifecycle_index" // The index of the lifecycle a service belongs to
	ServicePropertyTypeEnumName           ServicePropertyTypeEnum = "name"            // The name of a service
	ServicePropertyTypeEnumNote           ServicePropertyTypeEnum = "note"            // Additional information about the service
	ServicePropertyTypeEnumProduct        ServicePropertyTypeEnum = "product"         // The product that is associated with a service
	ServicePropertyTypeEnumSystem         ServicePropertyTypeEnum = "system"          // The system that the service belongs to
	ServicePropertyTypeEnumTierIndex      ServicePropertyTypeEnum = "tier_index"      // The index of the tier a service belongs to
)

// All ServicePropertyTypeEnum as []string
var AllServicePropertyTypeEnum = []string{
	string(ServicePropertyTypeEnumCustomProperty),
	string(ServicePropertyTypeEnumDescription),
	string(ServicePropertyTypeEnumFramework),
	string(ServicePropertyTypeEnumLanguage),
	string(ServicePropertyTypeEnumLifecycleIndex),
	string(ServicePropertyTypeEnumName),
	string(ServicePropertyTypeEnumNote),
	string(ServicePropertyTypeEnumProduct),
	string(ServicePropertyTypeEnumSystem),
	string(ServicePropertyTypeEnumTierIndex),
}

// ServiceSortEnum Sort possibilities for services
type ServiceSortEnum string

var (
	ServiceSortEnumAlertStatusAsc    ServiceSortEnum = "alert_status_ASC"    // Sort by alert status ascending
	ServiceSortEnumAlertStatusDesc   ServiceSortEnum = "alert_status_DESC"   // Sort by alert status descending
	ServiceSortEnumChecksPassingAsc  ServiceSortEnum = "checks_passing_ASC"  // Sort by `checks_passing` ascending
	ServiceSortEnumChecksPassingDesc ServiceSortEnum = "checks_passing_DESC" // Sort by `checks_passing` descending
	ServiceSortEnumComponentTypeAsc  ServiceSortEnum = "component_type_ASC"  // Sort by component type ascending
	ServiceSortEnumComponentTypeDesc ServiceSortEnum = "component_type_DESC" // Sort by component type descending
	ServiceSortEnumLastDeployAsc     ServiceSortEnum = "last_deploy_ASC"     // Sort by last deploy time ascending
	ServiceSortEnumLastDeployDesc    ServiceSortEnum = "last_deploy_DESC"    // Sort by last deploy time descending
	ServiceSortEnumLevelIndexAsc     ServiceSortEnum = "level_index_ASC"     // Sort by level ascending
	ServiceSortEnumLevelIndexDesc    ServiceSortEnum = "level_index_DESC"    // Sort by level descending
	ServiceSortEnumLifecycleAsc      ServiceSortEnum = "lifecycle_ASC"       // Sort by lifecycle ascending
	ServiceSortEnumLifecycleDesc     ServiceSortEnum = "lifecycle_DESC"      // Sort by lifecycle descending
	ServiceSortEnumNameAsc           ServiceSortEnum = "name_ASC"            // Sort by `name` ascending
	ServiceSortEnumNameDesc          ServiceSortEnum = "name_DESC"           // Sort by `name` descending
	ServiceSortEnumOwnerAsc          ServiceSortEnum = "owner_ASC"           // Sort by `owner` ascending
	ServiceSortEnumOwnerDesc         ServiceSortEnum = "owner_DESC"          // Sort by `owner` descending
	ServiceSortEnumProductAsc        ServiceSortEnum = "product_ASC"         // Sort by `product` ascending
	ServiceSortEnumProductDesc       ServiceSortEnum = "product_DESC"        // Sort by `product` descending
	ServiceSortEnumServiceStatAsc    ServiceSortEnum = "service_stat_ASC"    // Alias to sort by `checks_passing` ascending
	ServiceSortEnumServiceStatDesc   ServiceSortEnum = "service_stat_DESC"   // Alias to sort by `checks_passing` descending
	ServiceSortEnumTierAsc           ServiceSortEnum = "tier_ASC"            // Sort by `tier` ascending
	ServiceSortEnumTierDesc          ServiceSortEnum = "tier_DESC"           // Sort by `tier` descending
)

// All ServiceSortEnum as []string
var AllServiceSortEnum = []string{
	string(ServiceSortEnumAlertStatusAsc),
	string(ServiceSortEnumAlertStatusDesc),
	string(ServiceSortEnumChecksPassingAsc),
	string(ServiceSortEnumChecksPassingDesc),
	string(ServiceSortEnumComponentTypeAsc),
	string(ServiceSortEnumComponentTypeDesc),
	string(ServiceSortEnumLastDeployAsc),
	string(ServiceSortEnumLastDeployDesc),
	string(ServiceSortEnumLevelIndexAsc),
	string(ServiceSortEnumLevelIndexDesc),
	string(ServiceSortEnumLifecycleAsc),
	string(ServiceSortEnumLifecycleDesc),
	string(ServiceSortEnumNameAsc),
	string(ServiceSortEnumNameDesc),
	string(ServiceSortEnumOwnerAsc),
	string(ServiceSortEnumOwnerDesc),
	string(ServiceSortEnumProductAsc),
	string(ServiceSortEnumProductDesc),
	string(ServiceSortEnumServiceStatAsc),
	string(ServiceSortEnumServiceStatDesc),
	string(ServiceSortEnumTierAsc),
	string(ServiceSortEnumTierDesc),
}

// SnykIntegrationRegionEnum The data residency regions offered by Snyk
type SnykIntegrationRegionEnum string

var (
	SnykIntegrationRegionEnumAu SnykIntegrationRegionEnum = "AU" // Australia (https://api.au.snyk.io)
	SnykIntegrationRegionEnumEu SnykIntegrationRegionEnum = "EU" // Europe (https://api.eu.snyk.io)
	SnykIntegrationRegionEnumUs SnykIntegrationRegionEnum = "US" // USA (https://app.snyk.io)
)

// All SnykIntegrationRegionEnum as []string
var AllSnykIntegrationRegionEnum = []string{
	string(SnykIntegrationRegionEnumAu),
	string(SnykIntegrationRegionEnumEu),
	string(SnykIntegrationRegionEnumUs),
}

// TaggableResource Possible types to apply tags to
type TaggableResource string

var (
	TaggableResourceDomain                 TaggableResource = "Domain"                 // Used to identify a Domain
	TaggableResourceInfrastructureresource TaggableResource = "InfrastructureResource" // Used to identify an Infrastructure Resource
	TaggableResourceRepository             TaggableResource = "Repository"             // Used to identify a Repository
	TaggableResourceService                TaggableResource = "Service"                // Used to identify a Service
	TaggableResourceSystem                 TaggableResource = "System"                 // Used to identify a System
	TaggableResourceTeam                   TaggableResource = "Team"                   // Used to identify a Team
	TaggableResourceUser                   TaggableResource = "User"                   // Used to identify a User
)

// All TaggableResource as []string
var AllTaggableResource = []string{
	string(TaggableResourceDomain),
	string(TaggableResourceInfrastructureresource),
	string(TaggableResourceRepository),
	string(TaggableResourceService),
	string(TaggableResourceSystem),
	string(TaggableResourceTeam),
	string(TaggableResourceUser),
}

// ToolCategory The specific categories that a tool can belong to
type ToolCategory string

var (
	ToolCategoryAdmin                 ToolCategory = "admin"                  // Tools used for administrative purposes
	ToolCategoryAPIDocumentation      ToolCategory = "api_documentation"      // Tools used as API documentation for this service
	ToolCategoryArchitectureDiagram   ToolCategory = "architecture_diagram"   // Tools used for diagramming architecture
	ToolCategoryBacklog               ToolCategory = "backlog"                // Tools used for tracking issues
	ToolCategoryCode                  ToolCategory = "code"                   // Tools used for source code
	ToolCategoryContinuousIntegration ToolCategory = "continuous_integration" // Tools used for building/unit testing a service
	ToolCategoryDeployment            ToolCategory = "deployment"             // Tools used for deploying changes to a service
	ToolCategoryDesignDocumentation   ToolCategory = "design_documentation"   // Tools used for documenting design
	ToolCategoryErrors                ToolCategory = "errors"                 // Tools used for tracking/reporting errors
	ToolCategoryFeatureFlag           ToolCategory = "feature_flag"           // Tools used for managing feature flags
	ToolCategoryHealthChecks          ToolCategory = "health_checks"          // Tools used for tracking/reporting the health of a service
	ToolCategoryIncidents             ToolCategory = "incidents"              // Tools used to surface incidents on a service
	ToolCategoryIssueTracking         ToolCategory = "issue_tracking"         // Tools used for tracking issues
	ToolCategoryLogs                  ToolCategory = "logs"                   // Tools used for displaying logs from services
	ToolCategoryMetrics               ToolCategory = "metrics"                // Tools used for tracking/reporting service metrics
	ToolCategoryObservability         ToolCategory = "observability"          // Tools used for observability
	ToolCategoryOrchestrator          ToolCategory = "orchestrator"           // Tools used for orchestrating a service
	ToolCategoryOther                 ToolCategory = "other"                  // Tools that do not fit into the available categories
	ToolCategoryResiliency            ToolCategory = "resiliency"             // Tools used for testing the resiliency of a service
	ToolCategoryRunbooks              ToolCategory = "runbooks"               // Tools used for managing runbooks for a service
	ToolCategorySecurityScans         ToolCategory = "security_scans"         // Tools used for performing security scans
	ToolCategoryStatusPage            ToolCategory = "status_page"            // Tools used for reporting the status of a service
	ToolCategoryWiki                  ToolCategory = "wiki"                   // Tools used as a wiki for this service
)

// All ToolCategory as []string
var AllToolCategory = []string{
	string(ToolCategoryAdmin),
	string(ToolCategoryAPIDocumentation),
	string(ToolCategoryArchitectureDiagram),
	string(ToolCategoryBacklog),
	string(ToolCategoryCode),
	string(ToolCategoryContinuousIntegration),
	string(ToolCategoryDeployment),
	string(ToolCategoryDesignDocumentation),
	string(ToolCategoryErrors),
	string(ToolCategoryFeatureFlag),
	string(ToolCategoryHealthChecks),
	string(ToolCategoryIncidents),
	string(ToolCategoryIssueTracking),
	string(ToolCategoryLogs),
	string(ToolCategoryMetrics),
	string(ToolCategoryObservability),
	string(ToolCategoryOrchestrator),
	string(ToolCategoryOther),
	string(ToolCategoryResiliency),
	string(ToolCategoryRunbooks),
	string(ToolCategorySecurityScans),
	string(ToolCategoryStatusPage),
	string(ToolCategoryWiki),
}

// UserRole A role that can be assigned to a user
type UserRole string

var (
	UserRoleAdmin          UserRole = "admin"           // An administrator on the account
	UserRoleStandardsAdmin UserRole = "standards_admin" // Full write access to Standards resources, including rubric, campaigns, and checks. User-level access to all other entities
	UserRoleTeamMember     UserRole = "team_member"     // Read access to all resources. Write access based on team membership
	UserRoleUser           UserRole = "user"            // A regular user on the account
)

// All UserRole as []string
var AllUserRole = []string{
	string(UserRoleAdmin),
	string(UserRoleStandardsAdmin),
	string(UserRoleTeamMember),
	string(UserRoleUser),
}

// UsersFilterEnum Fields that can be used as part of filter for users
type UsersFilterEnum string

var (
	UsersFilterEnumDeactivatedAt UsersFilterEnum = "deactivated_at"  // Filter by the `deactivated_at` field
	UsersFilterEnumEmail         UsersFilterEnum = "email"           // Filter by `email` field
	UsersFilterEnumLastSignInAt  UsersFilterEnum = "last_sign_in_at" // Filter by the `last_sign_in_at` field
	UsersFilterEnumName          UsersFilterEnum = "name"            // Filter by `name` field
	UsersFilterEnumRole          UsersFilterEnum = "role"            // Filter by `role` field. (user or admin)
	UsersFilterEnumTag           UsersFilterEnum = "tag"             // Filter by `tags` belonging to user
)

// All UsersFilterEnum as []string
var AllUsersFilterEnum = []string{
	string(UsersFilterEnumDeactivatedAt),
	string(UsersFilterEnumEmail),
	string(UsersFilterEnumLastSignInAt),
	string(UsersFilterEnumName),
	string(UsersFilterEnumRole),
	string(UsersFilterEnumTag),
}

// UsersInviteScopeEnum A classification of users to invite
type UsersInviteScopeEnum string

var (
	UsersInviteScopeEnumPending UsersInviteScopeEnum = "pending" // All users who have yet to log in to OpsLevel for the first time
)

// All UsersInviteScopeEnum as []string
var AllUsersInviteScopeEnum = []string{
	string(UsersInviteScopeEnumPending),
}

// VaultSecretsSortEnum Sort possibilities for secrets
type VaultSecretsSortEnum string

var (
	VaultSecretsSortEnumSlugAsc       VaultSecretsSortEnum = "slug_ASC"        // Sort by slug ascending
	VaultSecretsSortEnumSlugDesc      VaultSecretsSortEnum = "slug_DESC"       // Sort by slug descending
	VaultSecretsSortEnumUpdatedAtAsc  VaultSecretsSortEnum = "updated_at_ASC"  // Sort by updated_at ascending
	VaultSecretsSortEnumUpdatedAtDesc VaultSecretsSortEnum = "updated_at_DESC" // Sort by updated_at descending
)

// All VaultSecretsSortEnum as []string
var AllVaultSecretsSortEnum = []string{
	string(VaultSecretsSortEnumSlugAsc),
	string(VaultSecretsSortEnumSlugDesc),
	string(VaultSecretsSortEnumUpdatedAtAsc),
	string(VaultSecretsSortEnumUpdatedAtDesc),
}
