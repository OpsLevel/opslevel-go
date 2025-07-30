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
	AlertSourceTypeEnumCustom      AlertSourceTypeEnum = "custom"       // A custom alert source (aka service)
	AlertSourceTypeEnumDatadog     AlertSourceTypeEnum = "datadog"      // A Datadog alert source (aka monitor)
	AlertSourceTypeEnumFireHydrant AlertSourceTypeEnum = "fire_hydrant" // An FireHydrant alert source (aka service)
	AlertSourceTypeEnumIncidentIo  AlertSourceTypeEnum = "incident_io"  // An incident.io alert source (aka service)
	AlertSourceTypeEnumNewRelic    AlertSourceTypeEnum = "new_relic"    // A New Relic alert source (aka service)
	AlertSourceTypeEnumOpsgenie    AlertSourceTypeEnum = "opsgenie"     // An Opsgenie alert source (aka service)
	AlertSourceTypeEnumPagerduty   AlertSourceTypeEnum = "pagerduty"    // A PagerDuty alert source (aka service)
)

// All AlertSourceTypeEnum as []string
var AllAlertSourceTypeEnum = []string{
	string(AlertSourceTypeEnumCustom),
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

// ApprovalDecisionEnum The set of possible outcomes for an approval decision
type ApprovalDecisionEnum string

var (
	ApprovalDecisionEnumApproved ApprovalDecisionEnum = "APPROVED" //
	ApprovalDecisionEnumDenied   ApprovalDecisionEnum = "DENIED"   //
)

// All ApprovalDecisionEnum as []string
var AllApprovalDecisionEnum = []string{
	string(ApprovalDecisionEnumApproved),
	string(ApprovalDecisionEnumDenied),
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
	CheckTypeRelationship        CheckType = "relationship"          // Verifies that the component has a specific number of relationship items defined for a specific relationship definition, with support for minimum, maximum, or exact count requirements
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
	string(CheckTypeRelationship),
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

// ComponentTypeIconEnum The possible icon names for a component type, provided by Phosphor icons for Vue: https://phosphoricons.com/
type ComponentTypeIconEnum string

var (
	ComponentTypeIconEnumPhactivity                    ComponentTypeIconEnum = "PhActivity"                    //
	ComponentTypeIconEnumPhaddressbook                 ComponentTypeIconEnum = "PhAddressBook"                 //
	ComponentTypeIconEnumPhairplane                    ComponentTypeIconEnum = "PhAirplane"                    //
	ComponentTypeIconEnumPhairplaneinflight            ComponentTypeIconEnum = "PhAirplaneInFlight"            //
	ComponentTypeIconEnumPhairplanelanding             ComponentTypeIconEnum = "PhAirplaneLanding"             //
	ComponentTypeIconEnumPhairplanetakeoff             ComponentTypeIconEnum = "PhAirplaneTakeoff"             //
	ComponentTypeIconEnumPhairplanetilt                ComponentTypeIconEnum = "PhAirplaneTilt"                //
	ComponentTypeIconEnumPhairplay                     ComponentTypeIconEnum = "PhAirplay"                     //
	ComponentTypeIconEnumPhalarm                       ComponentTypeIconEnum = "PhAlarm"                       //
	ComponentTypeIconEnumPhalien                       ComponentTypeIconEnum = "PhAlien"                       //
	ComponentTypeIconEnumPhalignbottom                 ComponentTypeIconEnum = "PhAlignBottom"                 //
	ComponentTypeIconEnumPhalignbottomsimple           ComponentTypeIconEnum = "PhAlignBottomSimple"           //
	ComponentTypeIconEnumPhaligncenterhorizontal       ComponentTypeIconEnum = "PhAlignCenterHorizontal"       //
	ComponentTypeIconEnumPhaligncenterhorizontalsimple ComponentTypeIconEnum = "PhAlignCenterHorizontalSimple" //
	ComponentTypeIconEnumPhaligncentervertical         ComponentTypeIconEnum = "PhAlignCenterVertical"         //
	ComponentTypeIconEnumPhaligncenterverticalsimple   ComponentTypeIconEnum = "PhAlignCenterVerticalSimple"   //
	ComponentTypeIconEnumPhalignleft                   ComponentTypeIconEnum = "PhAlignLeft"                   //
	ComponentTypeIconEnumPhalignleftsimple             ComponentTypeIconEnum = "PhAlignLeftSimple"             //
	ComponentTypeIconEnumPhalignright                  ComponentTypeIconEnum = "PhAlignRight"                  //
	ComponentTypeIconEnumPhalignrightsimple            ComponentTypeIconEnum = "PhAlignRightSimple"            //
	ComponentTypeIconEnumPhaligntop                    ComponentTypeIconEnum = "PhAlignTop"                    //
	ComponentTypeIconEnumPhaligntopsimple              ComponentTypeIconEnum = "PhAlignTopSimple"              //
	ComponentTypeIconEnumPhanchor                      ComponentTypeIconEnum = "PhAnchor"                      //
	ComponentTypeIconEnumPhanchorsimple                ComponentTypeIconEnum = "PhAnchorSimple"                //
	ComponentTypeIconEnumPhandroidlogo                 ComponentTypeIconEnum = "PhAndroidLogo"                 //
	ComponentTypeIconEnumPhangularlogo                 ComponentTypeIconEnum = "PhAngularLogo"                 //
	ComponentTypeIconEnumPhaperture                    ComponentTypeIconEnum = "PhAperture"                    //
	ComponentTypeIconEnumPhappstorelogo                ComponentTypeIconEnum = "PhAppStoreLogo"                //
	ComponentTypeIconEnumPhappwindow                   ComponentTypeIconEnum = "PhAppWindow"                   //
	ComponentTypeIconEnumPhapplelogo                   ComponentTypeIconEnum = "PhAppleLogo"                   //
	ComponentTypeIconEnumPhapplepodcastslogo           ComponentTypeIconEnum = "PhApplePodcastsLogo"           //
	ComponentTypeIconEnumPharchive                     ComponentTypeIconEnum = "PhArchive"                     //
	ComponentTypeIconEnumPharchivebox                  ComponentTypeIconEnum = "PhArchiveBox"                  //
	ComponentTypeIconEnumPharchivetray                 ComponentTypeIconEnum = "PhArchiveTray"                 //
	ComponentTypeIconEnumPharmchair                    ComponentTypeIconEnum = "PhArmchair"                    //
	ComponentTypeIconEnumPharrowarcleft                ComponentTypeIconEnum = "PhArrowArcLeft"                //
	ComponentTypeIconEnumPharrowarcright               ComponentTypeIconEnum = "PhArrowArcRight"               //
	ComponentTypeIconEnumPharrowbenddoubleupleft       ComponentTypeIconEnum = "PhArrowBendDoubleUpLeft"       //
	ComponentTypeIconEnumPharrowbenddoubleupright      ComponentTypeIconEnum = "PhArrowBendDoubleUpRight"      //
	ComponentTypeIconEnumPharrowbenddownleft           ComponentTypeIconEnum = "PhArrowBendDownLeft"           //
	ComponentTypeIconEnumPharrowbenddownright          ComponentTypeIconEnum = "PhArrowBendDownRight"          //
	ComponentTypeIconEnumPharrowbendleftdown           ComponentTypeIconEnum = "PhArrowBendLeftDown"           //
	ComponentTypeIconEnumPharrowbendleftup             ComponentTypeIconEnum = "PhArrowBendLeftUp"             //
	ComponentTypeIconEnumPharrowbendrightdown          ComponentTypeIconEnum = "PhArrowBendRightDown"          //
	ComponentTypeIconEnumPharrowbendrightup            ComponentTypeIconEnum = "PhArrowBendRightUp"            //
	ComponentTypeIconEnumPharrowbendupleft             ComponentTypeIconEnum = "PhArrowBendUpLeft"             //
	ComponentTypeIconEnumPharrowbendupright            ComponentTypeIconEnum = "PhArrowBendUpRight"            //
	ComponentTypeIconEnumPharrowcircledown             ComponentTypeIconEnum = "PhArrowCircleDown"             //
	ComponentTypeIconEnumPharrowcircledownleft         ComponentTypeIconEnum = "PhArrowCircleDownLeft"         //
	ComponentTypeIconEnumPharrowcircledownright        ComponentTypeIconEnum = "PhArrowCircleDownRight"        //
	ComponentTypeIconEnumPharrowcircleleft             ComponentTypeIconEnum = "PhArrowCircleLeft"             //
	ComponentTypeIconEnumPharrowcircleright            ComponentTypeIconEnum = "PhArrowCircleRight"            //
	ComponentTypeIconEnumPharrowcircleup               ComponentTypeIconEnum = "PhArrowCircleUp"               //
	ComponentTypeIconEnumPharrowcircleupleft           ComponentTypeIconEnum = "PhArrowCircleUpLeft"           //
	ComponentTypeIconEnumPharrowcircleupright          ComponentTypeIconEnum = "PhArrowCircleUpRight"          //
	ComponentTypeIconEnumPharrowclockwise              ComponentTypeIconEnum = "PhArrowClockwise"              //
	ComponentTypeIconEnumPharrowcounterclockwise       ComponentTypeIconEnum = "PhArrowCounterClockwise"       //
	ComponentTypeIconEnumPharrowdown                   ComponentTypeIconEnum = "PhArrowDown"                   //
	ComponentTypeIconEnumPharrowdownleft               ComponentTypeIconEnum = "PhArrowDownLeft"               //
	ComponentTypeIconEnumPharrowdownright              ComponentTypeIconEnum = "PhArrowDownRight"              //
	ComponentTypeIconEnumPharrowelbowdownleft          ComponentTypeIconEnum = "PhArrowElbowDownLeft"          //
	ComponentTypeIconEnumPharrowelbowdownright         ComponentTypeIconEnum = "PhArrowElbowDownRight"         //
	ComponentTypeIconEnumPharrowelbowleft              ComponentTypeIconEnum = "PhArrowElbowLeft"              //
	ComponentTypeIconEnumPharrowelbowleftdown          ComponentTypeIconEnum = "PhArrowElbowLeftDown"          //
	ComponentTypeIconEnumPharrowelbowleftup            ComponentTypeIconEnum = "PhArrowElbowLeftUp"            //
	ComponentTypeIconEnumPharrowelbowright             ComponentTypeIconEnum = "PhArrowElbowRight"             //
	ComponentTypeIconEnumPharrowelbowrightdown         ComponentTypeIconEnum = "PhArrowElbowRightDown"         //
	ComponentTypeIconEnumPharrowelbowrightup           ComponentTypeIconEnum = "PhArrowElbowRightUp"           //
	ComponentTypeIconEnumPharrowelbowupleft            ComponentTypeIconEnum = "PhArrowElbowUpLeft"            //
	ComponentTypeIconEnumPharrowelbowupright           ComponentTypeIconEnum = "PhArrowElbowUpRight"           //
	ComponentTypeIconEnumPharrowfatdown                ComponentTypeIconEnum = "PhArrowFatDown"                //
	ComponentTypeIconEnumPharrowfatleft                ComponentTypeIconEnum = "PhArrowFatLeft"                //
	ComponentTypeIconEnumPharrowfatlinedown            ComponentTypeIconEnum = "PhArrowFatLineDown"            //
	ComponentTypeIconEnumPharrowfatlineleft            ComponentTypeIconEnum = "PhArrowFatLineLeft"            //
	ComponentTypeIconEnumPharrowfatlineright           ComponentTypeIconEnum = "PhArrowFatLineRight"           //
	ComponentTypeIconEnumPharrowfatlineup              ComponentTypeIconEnum = "PhArrowFatLineUp"              //
	ComponentTypeIconEnumPharrowfatlinesdown           ComponentTypeIconEnum = "PhArrowFatLinesDown"           //
	ComponentTypeIconEnumPharrowfatlinesleft           ComponentTypeIconEnum = "PhArrowFatLinesLeft"           //
	ComponentTypeIconEnumPharrowfatlinesright          ComponentTypeIconEnum = "PhArrowFatLinesRight"          //
	ComponentTypeIconEnumPharrowfatlinesup             ComponentTypeIconEnum = "PhArrowFatLinesUp"             //
	ComponentTypeIconEnumPharrowfatright               ComponentTypeIconEnum = "PhArrowFatRight"               //
	ComponentTypeIconEnumPharrowfatup                  ComponentTypeIconEnum = "PhArrowFatUp"                  //
	ComponentTypeIconEnumPharrowleft                   ComponentTypeIconEnum = "PhArrowLeft"                   //
	ComponentTypeIconEnumPharrowlinedown               ComponentTypeIconEnum = "PhArrowLineDown"               //
	ComponentTypeIconEnumPharrowlinedownleft           ComponentTypeIconEnum = "PhArrowLineDownLeft"           //
	ComponentTypeIconEnumPharrowlinedownright          ComponentTypeIconEnum = "PhArrowLineDownRight"          //
	ComponentTypeIconEnumPharrowlineleft               ComponentTypeIconEnum = "PhArrowLineLeft"               //
	ComponentTypeIconEnumPharrowlineright              ComponentTypeIconEnum = "PhArrowLineRight"              //
	ComponentTypeIconEnumPharrowlineup                 ComponentTypeIconEnum = "PhArrowLineUp"                 //
	ComponentTypeIconEnumPharrowlineupleft             ComponentTypeIconEnum = "PhArrowLineUpLeft"             //
	ComponentTypeIconEnumPharrowlineupright            ComponentTypeIconEnum = "PhArrowLineUpRight"            //
	ComponentTypeIconEnumPharrowright                  ComponentTypeIconEnum = "PhArrowRight"                  //
	ComponentTypeIconEnumPharrowsquaredown             ComponentTypeIconEnum = "PhArrowSquareDown"             //
	ComponentTypeIconEnumPharrowsquaredownleft         ComponentTypeIconEnum = "PhArrowSquareDownLeft"         //
	ComponentTypeIconEnumPharrowsquaredownright        ComponentTypeIconEnum = "PhArrowSquareDownRight"        //
	ComponentTypeIconEnumPharrowsquarein               ComponentTypeIconEnum = "PhArrowSquareIn"               //
	ComponentTypeIconEnumPharrowsquareleft             ComponentTypeIconEnum = "PhArrowSquareLeft"             //
	ComponentTypeIconEnumPharrowsquareout              ComponentTypeIconEnum = "PhArrowSquareOut"              //
	ComponentTypeIconEnumPharrowsquareright            ComponentTypeIconEnum = "PhArrowSquareRight"            //
	ComponentTypeIconEnumPharrowsquareup               ComponentTypeIconEnum = "PhArrowSquareUp"               //
	ComponentTypeIconEnumPharrowsquareupleft           ComponentTypeIconEnum = "PhArrowSquareUpLeft"           //
	ComponentTypeIconEnumPharrowsquareupright          ComponentTypeIconEnum = "PhArrowSquareUpRight"          //
	ComponentTypeIconEnumPharrowudownleft              ComponentTypeIconEnum = "PhArrowUDownLeft"              //
	ComponentTypeIconEnumPharrowudownright             ComponentTypeIconEnum = "PhArrowUDownRight"             //
	ComponentTypeIconEnumPharrowuleftdown              ComponentTypeIconEnum = "PhArrowULeftDown"              //
	ComponentTypeIconEnumPharrowuleftup                ComponentTypeIconEnum = "PhArrowULeftUp"                //
	ComponentTypeIconEnumPharrowurightdown             ComponentTypeIconEnum = "PhArrowURightDown"             //
	ComponentTypeIconEnumPharrowurightup               ComponentTypeIconEnum = "PhArrowURightUp"               //
	ComponentTypeIconEnumPharrowuupleft                ComponentTypeIconEnum = "PhArrowUUpLeft"                //
	ComponentTypeIconEnumPharrowuupright               ComponentTypeIconEnum = "PhArrowUUpRight"               //
	ComponentTypeIconEnumPharrowup                     ComponentTypeIconEnum = "PhArrowUp"                     //
	ComponentTypeIconEnumPharrowupleft                 ComponentTypeIconEnum = "PhArrowUpLeft"                 //
	ComponentTypeIconEnumPharrowupright                ComponentTypeIconEnum = "PhArrowUpRight"                //
	ComponentTypeIconEnumPharrowsclockwise             ComponentTypeIconEnum = "PhArrowsClockwise"             //
	ComponentTypeIconEnumPharrowscounterclockwise      ComponentTypeIconEnum = "PhArrowsCounterClockwise"      //
	ComponentTypeIconEnumPharrowsdownup                ComponentTypeIconEnum = "PhArrowsDownUp"                //
	ComponentTypeIconEnumPharrowshorizontal            ComponentTypeIconEnum = "PhArrowsHorizontal"            //
	ComponentTypeIconEnumPharrowsin                    ComponentTypeIconEnum = "PhArrowsIn"                    //
	ComponentTypeIconEnumPharrowsincardinal            ComponentTypeIconEnum = "PhArrowsInCardinal"            //
	ComponentTypeIconEnumPharrowsinlinehorizontal      ComponentTypeIconEnum = "PhArrowsInLineHorizontal"      //
	ComponentTypeIconEnumPharrowsinlinevertical        ComponentTypeIconEnum = "PhArrowsInLineVertical"        //
	ComponentTypeIconEnumPharrowsinsimple              ComponentTypeIconEnum = "PhArrowsInSimple"              //
	ComponentTypeIconEnumPharrowsleftright             ComponentTypeIconEnum = "PhArrowsLeftRight"             //
	ComponentTypeIconEnumPharrowsout                   ComponentTypeIconEnum = "PhArrowsOut"                   //
	ComponentTypeIconEnumPharrowsoutcardinal           ComponentTypeIconEnum = "PhArrowsOutCardinal"           //
	ComponentTypeIconEnumPharrowsoutlinehorizontal     ComponentTypeIconEnum = "PhArrowsOutLineHorizontal"     //
	ComponentTypeIconEnumPharrowsoutlinevertical       ComponentTypeIconEnum = "PhArrowsOutLineVertical"       //
	ComponentTypeIconEnumPharrowsoutsimple             ComponentTypeIconEnum = "PhArrowsOutSimple"             //
	ComponentTypeIconEnumPharrowsvertical              ComponentTypeIconEnum = "PhArrowsVertical"              //
	ComponentTypeIconEnumPharticle                     ComponentTypeIconEnum = "PhArticle"                     //
	ComponentTypeIconEnumPharticlemedium               ComponentTypeIconEnum = "PhArticleMedium"               //
	ComponentTypeIconEnumPharticlenytimes              ComponentTypeIconEnum = "PhArticleNyTimes"              //
	ComponentTypeIconEnumPhasterisk                    ComponentTypeIconEnum = "PhAsterisk"                    //
	ComponentTypeIconEnumPhasterisksimple              ComponentTypeIconEnum = "PhAsteriskSimple"              //
	ComponentTypeIconEnumPhat                          ComponentTypeIconEnum = "PhAt"                          //
	ComponentTypeIconEnumPhatom                        ComponentTypeIconEnum = "PhAtom"                        //
	ComponentTypeIconEnumPhbaby                        ComponentTypeIconEnum = "PhBaby"                        //
	ComponentTypeIconEnumPhbackpack                    ComponentTypeIconEnum = "PhBackpack"                    //
	ComponentTypeIconEnumPhbackspace                   ComponentTypeIconEnum = "PhBackspace"                   //
	ComponentTypeIconEnumPhbag                         ComponentTypeIconEnum = "PhBag"                         //
	ComponentTypeIconEnumPhbagsimple                   ComponentTypeIconEnum = "PhBagSimple"                   //
	ComponentTypeIconEnumPhballoon                     ComponentTypeIconEnum = "PhBalloon"                     //
	ComponentTypeIconEnumPhbandaids                    ComponentTypeIconEnum = "PhBandaids"                    //
	ComponentTypeIconEnumPhbank                        ComponentTypeIconEnum = "PhBank"                        //
	ComponentTypeIconEnumPhbarbell                     ComponentTypeIconEnum = "PhBarbell"                     //
	ComponentTypeIconEnumPhbarcode                     ComponentTypeIconEnum = "PhBarcode"                     //
	ComponentTypeIconEnumPhbarricade                   ComponentTypeIconEnum = "PhBarricade"                   //
	ComponentTypeIconEnumPhbaseball                    ComponentTypeIconEnum = "PhBaseball"                    //
	ComponentTypeIconEnumPhbasketball                  ComponentTypeIconEnum = "PhBasketball"                  //
	ComponentTypeIconEnumPhbathtub                     ComponentTypeIconEnum = "PhBathtub"                     //
	ComponentTypeIconEnumPhbatterycharging             ComponentTypeIconEnum = "PhBatteryCharging"             //
	ComponentTypeIconEnumPhbatterychargingvertical     ComponentTypeIconEnum = "PhBatteryChargingVertical"     //
	ComponentTypeIconEnumPhbatteryempty                ComponentTypeIconEnum = "PhBatteryEmpty"                //
	ComponentTypeIconEnumPhbatteryfull                 ComponentTypeIconEnum = "PhBatteryFull"                 //
	ComponentTypeIconEnumPhbatteryhigh                 ComponentTypeIconEnum = "PhBatteryHigh"                 //
	ComponentTypeIconEnumPhbatterylow                  ComponentTypeIconEnum = "PhBatteryLow"                  //
	ComponentTypeIconEnumPhbatterymedium               ComponentTypeIconEnum = "PhBatteryMedium"               //
	ComponentTypeIconEnumPhbatteryplus                 ComponentTypeIconEnum = "PhBatteryPlus"                 //
	ComponentTypeIconEnumPhbatterywarning              ComponentTypeIconEnum = "PhBatteryWarning"              //
	ComponentTypeIconEnumPhbatterywarningvertical      ComponentTypeIconEnum = "PhBatteryWarningVertical"      //
	ComponentTypeIconEnumPhbed                         ComponentTypeIconEnum = "PhBed"                         //
	ComponentTypeIconEnumPhbeerbottle                  ComponentTypeIconEnum = "PhBeerBottle"                  //
	ComponentTypeIconEnumPhbehancelogo                 ComponentTypeIconEnum = "PhBehanceLogo"                 //
	ComponentTypeIconEnumPhbell                        ComponentTypeIconEnum = "PhBell"                        //
	ComponentTypeIconEnumPhbellringing                 ComponentTypeIconEnum = "PhBellRinging"                 //
	ComponentTypeIconEnumPhbellsimple                  ComponentTypeIconEnum = "PhBellSimple"                  //
	ComponentTypeIconEnumPhbellsimpleringing           ComponentTypeIconEnum = "PhBellSimpleRinging"           //
	ComponentTypeIconEnumPhbellsimpleslash             ComponentTypeIconEnum = "PhBellSimpleSlash"             //
	ComponentTypeIconEnumPhbellsimplez                 ComponentTypeIconEnum = "PhBellSimpleZ"                 //
	ComponentTypeIconEnumPhbellslash                   ComponentTypeIconEnum = "PhBellSlash"                   //
	ComponentTypeIconEnumPhbellz                       ComponentTypeIconEnum = "PhBellZ"                       //
	ComponentTypeIconEnumPhbeziercurve                 ComponentTypeIconEnum = "PhBezierCurve"                 //
	ComponentTypeIconEnumPhbicycle                     ComponentTypeIconEnum = "PhBicycle"                     //
	ComponentTypeIconEnumPhbinoculars                  ComponentTypeIconEnum = "PhBinoculars"                  //
	ComponentTypeIconEnumPhbird                        ComponentTypeIconEnum = "PhBird"                        //
	ComponentTypeIconEnumPhbluetooth                   ComponentTypeIconEnum = "PhBluetooth"                   //
	ComponentTypeIconEnumPhbluetoothconnected          ComponentTypeIconEnum = "PhBluetoothConnected"          //
	ComponentTypeIconEnumPhbluetoothslash              ComponentTypeIconEnum = "PhBluetoothSlash"              //
	ComponentTypeIconEnumPhbluetoothx                  ComponentTypeIconEnum = "PhBluetoothX"                  //
	ComponentTypeIconEnumPhboat                        ComponentTypeIconEnum = "PhBoat"                        //
	ComponentTypeIconEnumPhbook                        ComponentTypeIconEnum = "PhBook"                        //
	ComponentTypeIconEnumPhbookbookmark                ComponentTypeIconEnum = "PhBookBookmark"                //
	ComponentTypeIconEnumPhbookopen                    ComponentTypeIconEnum = "PhBookOpen"                    //
	ComponentTypeIconEnumPhbookmark                    ComponentTypeIconEnum = "PhBookmark"                    //
	ComponentTypeIconEnumPhbookmarksimple              ComponentTypeIconEnum = "PhBookmarkSimple"              //
	ComponentTypeIconEnumPhbookmarks                   ComponentTypeIconEnum = "PhBookmarks"                   //
	ComponentTypeIconEnumPhbookmarkssimple             ComponentTypeIconEnum = "PhBookmarksSimple"             //
	ComponentTypeIconEnumPhbooks                       ComponentTypeIconEnum = "PhBooks"                       //
	ComponentTypeIconEnumPhboundingbox                 ComponentTypeIconEnum = "PhBoundingBox"                 //
	ComponentTypeIconEnumPhbracketsangle               ComponentTypeIconEnum = "PhBracketsAngle"               //
	ComponentTypeIconEnumPhbracketscurly               ComponentTypeIconEnum = "PhBracketsCurly"               //
	ComponentTypeIconEnumPhbracketsround               ComponentTypeIconEnum = "PhBracketsRound"               //
	ComponentTypeIconEnumPhbracketssquare              ComponentTypeIconEnum = "PhBracketsSquare"              //
	ComponentTypeIconEnumPhbrain                       ComponentTypeIconEnum = "PhBrain"                       //
	ComponentTypeIconEnumPhbrandy                      ComponentTypeIconEnum = "PhBrandy"                      //
	ComponentTypeIconEnumPhbriefcase                   ComponentTypeIconEnum = "PhBriefcase"                   //
	ComponentTypeIconEnumPhbriefcasemetal              ComponentTypeIconEnum = "PhBriefcaseMetal"              //
	ComponentTypeIconEnumPhbroadcast                   ComponentTypeIconEnum = "PhBroadcast"                   //
	ComponentTypeIconEnumPhbrowser                     ComponentTypeIconEnum = "PhBrowser"                     //
	ComponentTypeIconEnumPhbrowsers                    ComponentTypeIconEnum = "PhBrowsers"                    //
	ComponentTypeIconEnumPhbug                         ComponentTypeIconEnum = "PhBug"                         //
	ComponentTypeIconEnumPhbugbeetle                   ComponentTypeIconEnum = "PhBugBeetle"                   //
	ComponentTypeIconEnumPhbugdroid                    ComponentTypeIconEnum = "PhBugDroid"                    //
	ComponentTypeIconEnumPhbuildings                   ComponentTypeIconEnum = "PhBuildings"                   //
	ComponentTypeIconEnumPhbus                         ComponentTypeIconEnum = "PhBus"                         //
	ComponentTypeIconEnumPhbutterfly                   ComponentTypeIconEnum = "PhButterfly"                   //
	ComponentTypeIconEnumPhcactus                      ComponentTypeIconEnum = "PhCactus"                      //
	ComponentTypeIconEnumPhcake                        ComponentTypeIconEnum = "PhCake"                        //
	ComponentTypeIconEnumPhcalculator                  ComponentTypeIconEnum = "PhCalculator"                  //
	ComponentTypeIconEnumPhcalendar                    ComponentTypeIconEnum = "PhCalendar"                    //
	ComponentTypeIconEnumPhcalendarblank               ComponentTypeIconEnum = "PhCalendarBlank"               //
	ComponentTypeIconEnumPhcalendarcheck               ComponentTypeIconEnum = "PhCalendarCheck"               //
	ComponentTypeIconEnumPhcalendarplus                ComponentTypeIconEnum = "PhCalendarPlus"                //
	ComponentTypeIconEnumPhcalendarx                   ComponentTypeIconEnum = "PhCalendarX"                   //
	ComponentTypeIconEnumPhcamera                      ComponentTypeIconEnum = "PhCamera"                      //
	ComponentTypeIconEnumPhcamerarotate                ComponentTypeIconEnum = "PhCameraRotate"                //
	ComponentTypeIconEnumPhcameraslash                 ComponentTypeIconEnum = "PhCameraSlash"                 //
	ComponentTypeIconEnumPhcampfire                    ComponentTypeIconEnum = "PhCampfire"                    //
	ComponentTypeIconEnumPhcar                         ComponentTypeIconEnum = "PhCar"                         //
	ComponentTypeIconEnumPhcarsimple                   ComponentTypeIconEnum = "PhCarSimple"                   //
	ComponentTypeIconEnumPhcardholder                  ComponentTypeIconEnum = "PhCardholder"                  //
	ComponentTypeIconEnumPhcards                       ComponentTypeIconEnum = "PhCards"                       //
	ComponentTypeIconEnumPhcaretcircledoubledown       ComponentTypeIconEnum = "PhCaretCircleDoubleDown"       //
	ComponentTypeIconEnumPhcaretcircledoubleleft       ComponentTypeIconEnum = "PhCaretCircleDoubleLeft"       //
	ComponentTypeIconEnumPhcaretcircledoubleright      ComponentTypeIconEnum = "PhCaretCircleDoubleRight"      //
	ComponentTypeIconEnumPhcaretcircledoubleup         ComponentTypeIconEnum = "PhCaretCircleDoubleUp"         //
	ComponentTypeIconEnumPhcaretcircledown             ComponentTypeIconEnum = "PhCaretCircleDown"             //
	ComponentTypeIconEnumPhcaretcircleleft             ComponentTypeIconEnum = "PhCaretCircleLeft"             //
	ComponentTypeIconEnumPhcaretcircleright            ComponentTypeIconEnum = "PhCaretCircleRight"            //
	ComponentTypeIconEnumPhcaretcircleup               ComponentTypeIconEnum = "PhCaretCircleUp"               //
	ComponentTypeIconEnumPhcaretdoubledown             ComponentTypeIconEnum = "PhCaretDoubleDown"             //
	ComponentTypeIconEnumPhcaretdoubleleft             ComponentTypeIconEnum = "PhCaretDoubleLeft"             //
	ComponentTypeIconEnumPhcaretdoubleright            ComponentTypeIconEnum = "PhCaretDoubleRight"            //
	ComponentTypeIconEnumPhcaretdoubleup               ComponentTypeIconEnum = "PhCaretDoubleUp"               //
	ComponentTypeIconEnumPhcaretdown                   ComponentTypeIconEnum = "PhCaretDown"                   //
	ComponentTypeIconEnumPhcaretleft                   ComponentTypeIconEnum = "PhCaretLeft"                   //
	ComponentTypeIconEnumPhcaretright                  ComponentTypeIconEnum = "PhCaretRight"                  //
	ComponentTypeIconEnumPhcaretup                     ComponentTypeIconEnum = "PhCaretUp"                     //
	ComponentTypeIconEnumPhcat                         ComponentTypeIconEnum = "PhCat"                         //
	ComponentTypeIconEnumPhcellsignalfull              ComponentTypeIconEnum = "PhCellSignalFull"              //
	ComponentTypeIconEnumPhcellsignalhigh              ComponentTypeIconEnum = "PhCellSignalHigh"              //
	ComponentTypeIconEnumPhcellsignallow               ComponentTypeIconEnum = "PhCellSignalLow"               //
	ComponentTypeIconEnumPhcellsignalmedium            ComponentTypeIconEnum = "PhCellSignalMedium"            //
	ComponentTypeIconEnumPhcellsignalnone              ComponentTypeIconEnum = "PhCellSignalNone"              //
	ComponentTypeIconEnumPhcellsignalslash             ComponentTypeIconEnum = "PhCellSignalSlash"             //
	ComponentTypeIconEnumPhcellsignalx                 ComponentTypeIconEnum = "PhCellSignalX"                 //
	ComponentTypeIconEnumPhchalkboard                  ComponentTypeIconEnum = "PhChalkboard"                  //
	ComponentTypeIconEnumPhchalkboardsimple            ComponentTypeIconEnum = "PhChalkboardSimple"            //
	ComponentTypeIconEnumPhchalkboardteacher           ComponentTypeIconEnum = "PhChalkboardTeacher"           //
	ComponentTypeIconEnumPhchartbar                    ComponentTypeIconEnum = "PhChartBar"                    //
	ComponentTypeIconEnumPhchartbarhorizontal          ComponentTypeIconEnum = "PhChartBarHorizontal"          //
	ComponentTypeIconEnumPhchartline                   ComponentTypeIconEnum = "PhChartLine"                   //
	ComponentTypeIconEnumPhchartlineup                 ComponentTypeIconEnum = "PhChartLineUp"                 //
	ComponentTypeIconEnumPhchartpie                    ComponentTypeIconEnum = "PhChartPie"                    //
	ComponentTypeIconEnumPhchartpieslice               ComponentTypeIconEnum = "PhChartPieSlice"               //
	ComponentTypeIconEnumPhchat                        ComponentTypeIconEnum = "PhChat"                        //
	ComponentTypeIconEnumPhchatcentered                ComponentTypeIconEnum = "PhChatCentered"                //
	ComponentTypeIconEnumPhchatcentereddots            ComponentTypeIconEnum = "PhChatCenteredDots"            //
	ComponentTypeIconEnumPhchatcenteredtext            ComponentTypeIconEnum = "PhChatCenteredText"            //
	ComponentTypeIconEnumPhchatcircle                  ComponentTypeIconEnum = "PhChatCircle"                  //
	ComponentTypeIconEnumPhchatcircledots              ComponentTypeIconEnum = "PhChatCircleDots"              //
	ComponentTypeIconEnumPhchatcircletext              ComponentTypeIconEnum = "PhChatCircleText"              //
	ComponentTypeIconEnumPhchatdots                    ComponentTypeIconEnum = "PhChatDots"                    //
	ComponentTypeIconEnumPhchatteardrop                ComponentTypeIconEnum = "PhChatTeardrop"                //
	ComponentTypeIconEnumPhchatteardropdots            ComponentTypeIconEnum = "PhChatTeardropDots"            //
	ComponentTypeIconEnumPhchatteardroptext            ComponentTypeIconEnum = "PhChatTeardropText"            //
	ComponentTypeIconEnumPhchattext                    ComponentTypeIconEnum = "PhChatText"                    //
	ComponentTypeIconEnumPhchats                       ComponentTypeIconEnum = "PhChats"                       //
	ComponentTypeIconEnumPhchatscircle                 ComponentTypeIconEnum = "PhChatsCircle"                 //
	ComponentTypeIconEnumPhchatsteardrop               ComponentTypeIconEnum = "PhChatsTeardrop"               //
	ComponentTypeIconEnumPhcheck                       ComponentTypeIconEnum = "PhCheck"                       //
	ComponentTypeIconEnumPhcheckcircle                 ComponentTypeIconEnum = "PhCheckCircle"                 //
	ComponentTypeIconEnumPhchecksquare                 ComponentTypeIconEnum = "PhCheckSquare"                 //
	ComponentTypeIconEnumPhchecksquareoffset           ComponentTypeIconEnum = "PhCheckSquareOffset"           //
	ComponentTypeIconEnumPhchecks                      ComponentTypeIconEnum = "PhChecks"                      //
	ComponentTypeIconEnumPhcircle                      ComponentTypeIconEnum = "PhCircle"                      //
	ComponentTypeIconEnumPhcircledashed                ComponentTypeIconEnum = "PhCircleDashed"                //
	ComponentTypeIconEnumPhcirclehalf                  ComponentTypeIconEnum = "PhCircleHalf"                  //
	ComponentTypeIconEnumPhcirclehalftilt              ComponentTypeIconEnum = "PhCircleHalfTilt"              //
	ComponentTypeIconEnumPhcirclenotch                 ComponentTypeIconEnum = "PhCircleNotch"                 //
	ComponentTypeIconEnumPhcirclewavy                  ComponentTypeIconEnum = "PhCircleWavy"                  //
	ComponentTypeIconEnumPhcirclewavycheck             ComponentTypeIconEnum = "PhCircleWavyCheck"             //
	ComponentTypeIconEnumPhcirclewavyquestion          ComponentTypeIconEnum = "PhCircleWavyQuestion"          //
	ComponentTypeIconEnumPhcirclewavywarning           ComponentTypeIconEnum = "PhCircleWavyWarning"           //
	ComponentTypeIconEnumPhcirclesfour                 ComponentTypeIconEnum = "PhCirclesFour"                 //
	ComponentTypeIconEnumPhcirclesthree                ComponentTypeIconEnum = "PhCirclesThree"                //
	ComponentTypeIconEnumPhcirclesthreeplus            ComponentTypeIconEnum = "PhCirclesThreePlus"            //
	ComponentTypeIconEnumPhclipboard                   ComponentTypeIconEnum = "PhClipboard"                   //
	ComponentTypeIconEnumPhclipboardtext               ComponentTypeIconEnum = "PhClipboardText"               //
	ComponentTypeIconEnumPhclock                       ComponentTypeIconEnum = "PhClock"                       //
	ComponentTypeIconEnumPhclockafternoon              ComponentTypeIconEnum = "PhClockAfternoon"              //
	ComponentTypeIconEnumPhclockclockwise              ComponentTypeIconEnum = "PhClockClockwise"              //
	ComponentTypeIconEnumPhclockcounterclockwise       ComponentTypeIconEnum = "PhClockCounterClockwise"       //
	ComponentTypeIconEnumPhclosedcaptioning            ComponentTypeIconEnum = "PhClosedCaptioning"            //
	ComponentTypeIconEnumPhcloud                       ComponentTypeIconEnum = "PhCloud"                       //
	ComponentTypeIconEnumPhcloudarrowdown              ComponentTypeIconEnum = "PhCloudArrowDown"              //
	ComponentTypeIconEnumPhcloudarrowup                ComponentTypeIconEnum = "PhCloudArrowUp"                //
	ComponentTypeIconEnumPhcloudcheck                  ComponentTypeIconEnum = "PhCloudCheck"                  //
	ComponentTypeIconEnumPhcloudfog                    ComponentTypeIconEnum = "PhCloudFog"                    //
	ComponentTypeIconEnumPhcloudlightning              ComponentTypeIconEnum = "PhCloudLightning"              //
	ComponentTypeIconEnumPhcloudmoon                   ComponentTypeIconEnum = "PhCloudMoon"                   //
	ComponentTypeIconEnumPhcloudrain                   ComponentTypeIconEnum = "PhCloudRain"                   //
	ComponentTypeIconEnumPhcloudslash                  ComponentTypeIconEnum = "PhCloudSlash"                  //
	ComponentTypeIconEnumPhcloudsnow                   ComponentTypeIconEnum = "PhCloudSnow"                   //
	ComponentTypeIconEnumPhcloudsun                    ComponentTypeIconEnum = "PhCloudSun"                    //
	ComponentTypeIconEnumPhclub                        ComponentTypeIconEnum = "PhClub"                        //
	ComponentTypeIconEnumPhcoathanger                  ComponentTypeIconEnum = "PhCoatHanger"                  //
	ComponentTypeIconEnumPhcode                        ComponentTypeIconEnum = "PhCode"                        //
	ComponentTypeIconEnumPhcodesimple                  ComponentTypeIconEnum = "PhCodeSimple"                  //
	ComponentTypeIconEnumPhcodepenlogo                 ComponentTypeIconEnum = "PhCodepenLogo"                 //
	ComponentTypeIconEnumPhcodesandboxlogo             ComponentTypeIconEnum = "PhCodesandboxLogo"             //
	ComponentTypeIconEnumPhcoffee                      ComponentTypeIconEnum = "PhCoffee"                      //
	ComponentTypeIconEnumPhcoin                        ComponentTypeIconEnum = "PhCoin"                        //
	ComponentTypeIconEnumPhcoinvertical                ComponentTypeIconEnum = "PhCoinVertical"                //
	ComponentTypeIconEnumPhcoins                       ComponentTypeIconEnum = "PhCoins"                       //
	ComponentTypeIconEnumPhcolumns                     ComponentTypeIconEnum = "PhColumns"                     //
	ComponentTypeIconEnumPhcommand                     ComponentTypeIconEnum = "PhCommand"                     //
	ComponentTypeIconEnumPhcompass                     ComponentTypeIconEnum = "PhCompass"                     //
	ComponentTypeIconEnumPhcomputertower               ComponentTypeIconEnum = "PhComputerTower"               //
	ComponentTypeIconEnumPhconfetti                    ComponentTypeIconEnum = "PhConfetti"                    //
	ComponentTypeIconEnumPhcookie                      ComponentTypeIconEnum = "PhCookie"                      //
	ComponentTypeIconEnumPhcookingpot                  ComponentTypeIconEnum = "PhCookingPot"                  //
	ComponentTypeIconEnumPhcopy                        ComponentTypeIconEnum = "PhCopy"                        //
	ComponentTypeIconEnumPhcopysimple                  ComponentTypeIconEnum = "PhCopySimple"                  //
	ComponentTypeIconEnumPhcopyleft                    ComponentTypeIconEnum = "PhCopyleft"                    //
	ComponentTypeIconEnumPhcopyright                   ComponentTypeIconEnum = "PhCopyright"                   //
	ComponentTypeIconEnumPhcornersin                   ComponentTypeIconEnum = "PhCornersIn"                   //
	ComponentTypeIconEnumPhcornersout                  ComponentTypeIconEnum = "PhCornersOut"                  //
	ComponentTypeIconEnumPhcpu                         ComponentTypeIconEnum = "PhCpu"                         //
	ComponentTypeIconEnumPhcreditcard                  ComponentTypeIconEnum = "PhCreditCard"                  //
	ComponentTypeIconEnumPhcrop                        ComponentTypeIconEnum = "PhCrop"                        //
	ComponentTypeIconEnumPhcrosshair                   ComponentTypeIconEnum = "PhCrosshair"                   //
	ComponentTypeIconEnumPhcrosshairsimple             ComponentTypeIconEnum = "PhCrosshairSimple"             //
	ComponentTypeIconEnumPhcrown                       ComponentTypeIconEnum = "PhCrown"                       //
	ComponentTypeIconEnumPhcrownsimple                 ComponentTypeIconEnum = "PhCrownSimple"                 //
	ComponentTypeIconEnumPhcube                        ComponentTypeIconEnum = "PhCube"                        //
	ComponentTypeIconEnumPhcurrencybtc                 ComponentTypeIconEnum = "PhCurrencyBtc"                 //
	ComponentTypeIconEnumPhcurrencycircledollar        ComponentTypeIconEnum = "PhCurrencyCircleDollar"        //
	ComponentTypeIconEnumPhcurrencycny                 ComponentTypeIconEnum = "PhCurrencyCny"                 //
	ComponentTypeIconEnumPhcurrencydollar              ComponentTypeIconEnum = "PhCurrencyDollar"              //
	ComponentTypeIconEnumPhcurrencydollarsimple        ComponentTypeIconEnum = "PhCurrencyDollarSimple"        //
	ComponentTypeIconEnumPhcurrencyeth                 ComponentTypeIconEnum = "PhCurrencyEth"                 //
	ComponentTypeIconEnumPhcurrencyeur                 ComponentTypeIconEnum = "PhCurrencyEur"                 //
	ComponentTypeIconEnumPhcurrencygbp                 ComponentTypeIconEnum = "PhCurrencyGbp"                 //
	ComponentTypeIconEnumPhcurrencyinr                 ComponentTypeIconEnum = "PhCurrencyInr"                 //
	ComponentTypeIconEnumPhcurrencyjpy                 ComponentTypeIconEnum = "PhCurrencyJpy"                 //
	ComponentTypeIconEnumPhcurrencykrw                 ComponentTypeIconEnum = "PhCurrencyKrw"                 //
	ComponentTypeIconEnumPhcurrencykzt                 ComponentTypeIconEnum = "PhCurrencyKzt"                 //
	ComponentTypeIconEnumPhcurrencyngn                 ComponentTypeIconEnum = "PhCurrencyNgn"                 //
	ComponentTypeIconEnumPhcurrencyrub                 ComponentTypeIconEnum = "PhCurrencyRub"                 //
	ComponentTypeIconEnumPhcursor                      ComponentTypeIconEnum = "PhCursor"                      //
	ComponentTypeIconEnumPhcursortext                  ComponentTypeIconEnum = "PhCursorText"                  //
	ComponentTypeIconEnumPhcylinder                    ComponentTypeIconEnum = "PhCylinder"                    //
	ComponentTypeIconEnumPhdatabase                    ComponentTypeIconEnum = "PhDatabase"                    //
	ComponentTypeIconEnumPhdesktop                     ComponentTypeIconEnum = "PhDesktop"                     //
	ComponentTypeIconEnumPhdesktoptower                ComponentTypeIconEnum = "PhDesktopTower"                //
	ComponentTypeIconEnumPhdetective                   ComponentTypeIconEnum = "PhDetective"                   //
	ComponentTypeIconEnumPhdevicemobile                ComponentTypeIconEnum = "PhDeviceMobile"                //
	ComponentTypeIconEnumPhdevicemobilecamera          ComponentTypeIconEnum = "PhDeviceMobileCamera"          //
	ComponentTypeIconEnumPhdevicemobilespeaker         ComponentTypeIconEnum = "PhDeviceMobileSpeaker"         //
	ComponentTypeIconEnumPhdevicetablet                ComponentTypeIconEnum = "PhDeviceTablet"                //
	ComponentTypeIconEnumPhdevicetabletcamera          ComponentTypeIconEnum = "PhDeviceTabletCamera"          //
	ComponentTypeIconEnumPhdevicetabletspeaker         ComponentTypeIconEnum = "PhDeviceTabletSpeaker"         //
	ComponentTypeIconEnumPhdiamond                     ComponentTypeIconEnum = "PhDiamond"                     //
	ComponentTypeIconEnumPhdiamondsfour                ComponentTypeIconEnum = "PhDiamondsFour"                //
	ComponentTypeIconEnumPhdicefive                    ComponentTypeIconEnum = "PhDiceFive"                    //
	ComponentTypeIconEnumPhdicefour                    ComponentTypeIconEnum = "PhDiceFour"                    //
	ComponentTypeIconEnumPhdiceone                     ComponentTypeIconEnum = "PhDiceOne"                     //
	ComponentTypeIconEnumPhdicesix                     ComponentTypeIconEnum = "PhDiceSix"                     //
	ComponentTypeIconEnumPhdicethree                   ComponentTypeIconEnum = "PhDiceThree"                   //
	ComponentTypeIconEnumPhdicetwo                     ComponentTypeIconEnum = "PhDiceTwo"                     //
	ComponentTypeIconEnumPhdisc                        ComponentTypeIconEnum = "PhDisc"                        //
	ComponentTypeIconEnumPhdiscordlogo                 ComponentTypeIconEnum = "PhDiscordLogo"                 //
	ComponentTypeIconEnumPhdivide                      ComponentTypeIconEnum = "PhDivide"                      //
	ComponentTypeIconEnumPhdog                         ComponentTypeIconEnum = "PhDog"                         //
	ComponentTypeIconEnumPhdoor                        ComponentTypeIconEnum = "PhDoor"                        //
	ComponentTypeIconEnumPhdotsnine                    ComponentTypeIconEnum = "PhDotsNine"                    //
	ComponentTypeIconEnumPhdotssix                     ComponentTypeIconEnum = "PhDotsSix"                     //
	ComponentTypeIconEnumPhdotssixvertical             ComponentTypeIconEnum = "PhDotsSixVertical"             //
	ComponentTypeIconEnumPhdotsthree                   ComponentTypeIconEnum = "PhDotsThree"                   //
	ComponentTypeIconEnumPhdotsthreecircle             ComponentTypeIconEnum = "PhDotsThreeCircle"             //
	ComponentTypeIconEnumPhdotsthreecirclevertical     ComponentTypeIconEnum = "PhDotsThreeCircleVertical"     //
	ComponentTypeIconEnumPhdotsthreeoutline            ComponentTypeIconEnum = "PhDotsThreeOutline"            //
	ComponentTypeIconEnumPhdotsthreeoutlinevertical    ComponentTypeIconEnum = "PhDotsThreeOutlineVertical"    //
	ComponentTypeIconEnumPhdotsthreevertical           ComponentTypeIconEnum = "PhDotsThreeVertical"           //
	ComponentTypeIconEnumPhdownload                    ComponentTypeIconEnum = "PhDownload"                    //
	ComponentTypeIconEnumPhdownloadsimple              ComponentTypeIconEnum = "PhDownloadSimple"              //
	ComponentTypeIconEnumPhdribbblelogo                ComponentTypeIconEnum = "PhDribbbleLogo"                //
	ComponentTypeIconEnumPhdrop                        ComponentTypeIconEnum = "PhDrop"                        //
	ComponentTypeIconEnumPhdrophalf                    ComponentTypeIconEnum = "PhDropHalf"                    //
	ComponentTypeIconEnumPhdrophalfbottom              ComponentTypeIconEnum = "PhDropHalfBottom"              //
	ComponentTypeIconEnumPhear                         ComponentTypeIconEnum = "PhEar"                         //
	ComponentTypeIconEnumPhearslash                    ComponentTypeIconEnum = "PhEarSlash"                    //
	ComponentTypeIconEnumPhegg                         ComponentTypeIconEnum = "PhEgg"                         //
	ComponentTypeIconEnumPheggcrack                    ComponentTypeIconEnum = "PhEggCrack"                    //
	ComponentTypeIconEnumPheject                       ComponentTypeIconEnum = "PhEject"                       //
	ComponentTypeIconEnumPhejectsimple                 ComponentTypeIconEnum = "PhEjectSimple"                 //
	ComponentTypeIconEnumPhenvelope                    ComponentTypeIconEnum = "PhEnvelope"                    //
	ComponentTypeIconEnumPhenvelopeopen                ComponentTypeIconEnum = "PhEnvelopeOpen"                //
	ComponentTypeIconEnumPhenvelopesimple              ComponentTypeIconEnum = "PhEnvelopeSimple"              //
	ComponentTypeIconEnumPhenvelopesimpleopen          ComponentTypeIconEnum = "PhEnvelopeSimpleOpen"          //
	ComponentTypeIconEnumPhequalizer                   ComponentTypeIconEnum = "PhEqualizer"                   //
	ComponentTypeIconEnumPhequals                      ComponentTypeIconEnum = "PhEquals"                      //
	ComponentTypeIconEnumPheraser                      ComponentTypeIconEnum = "PhEraser"                      //
	ComponentTypeIconEnumPhexam                        ComponentTypeIconEnum = "PhExam"                        //
	ComponentTypeIconEnumPhexport                      ComponentTypeIconEnum = "PhExport"                      //
	ComponentTypeIconEnumPheye                         ComponentTypeIconEnum = "PhEye"                         //
	ComponentTypeIconEnumPheyeclosed                   ComponentTypeIconEnum = "PhEyeClosed"                   //
	ComponentTypeIconEnumPheyeslash                    ComponentTypeIconEnum = "PhEyeSlash"                    //
	ComponentTypeIconEnumPheyedropper                  ComponentTypeIconEnum = "PhEyedropper"                  //
	ComponentTypeIconEnumPheyedroppersample            ComponentTypeIconEnum = "PhEyedropperSample"            //
	ComponentTypeIconEnumPheyeglasses                  ComponentTypeIconEnum = "PhEyeglasses"                  //
	ComponentTypeIconEnumPhfacemask                    ComponentTypeIconEnum = "PhFaceMask"                    //
	ComponentTypeIconEnumPhfacebooklogo                ComponentTypeIconEnum = "PhFacebookLogo"                //
	ComponentTypeIconEnumPhfactory                     ComponentTypeIconEnum = "PhFactory"                     //
	ComponentTypeIconEnumPhfaders                      ComponentTypeIconEnum = "PhFaders"                      //
	ComponentTypeIconEnumPhfadershorizontal            ComponentTypeIconEnum = "PhFadersHorizontal"            //
	ComponentTypeIconEnumPhfastforward                 ComponentTypeIconEnum = "PhFastForward"                 //
	ComponentTypeIconEnumPhfastforwardcircle           ComponentTypeIconEnum = "PhFastForwardCircle"           //
	ComponentTypeIconEnumPhfigmalogo                   ComponentTypeIconEnum = "PhFigmaLogo"                   //
	ComponentTypeIconEnumPhfile                        ComponentTypeIconEnum = "PhFile"                        //
	ComponentTypeIconEnumPhfilearrowdown               ComponentTypeIconEnum = "PhFileArrowDown"               //
	ComponentTypeIconEnumPhfilearrowup                 ComponentTypeIconEnum = "PhFileArrowUp"                 //
	ComponentTypeIconEnumPhfileaudio                   ComponentTypeIconEnum = "PhFileAudio"                   //
	ComponentTypeIconEnumPhfilecloud                   ComponentTypeIconEnum = "PhFileCloud"                   //
	ComponentTypeIconEnumPhfilecode                    ComponentTypeIconEnum = "PhFileCode"                    //
	ComponentTypeIconEnumPhfilecss                     ComponentTypeIconEnum = "PhFileCss"                     //
	ComponentTypeIconEnumPhfilecsv                     ComponentTypeIconEnum = "PhFileCsv"                     //
	ComponentTypeIconEnumPhfiledoc                     ComponentTypeIconEnum = "PhFileDoc"                     //
	ComponentTypeIconEnumPhfiledotted                  ComponentTypeIconEnum = "PhFileDotted"                  //
	ComponentTypeIconEnumPhfilehtml                    ComponentTypeIconEnum = "PhFileHtml"                    //
	ComponentTypeIconEnumPhfileimage                   ComponentTypeIconEnum = "PhFileImage"                   //
	ComponentTypeIconEnumPhfilejpg                     ComponentTypeIconEnum = "PhFileJpg"                     //
	ComponentTypeIconEnumPhfilejs                      ComponentTypeIconEnum = "PhFileJs"                      //
	ComponentTypeIconEnumPhfilejsx                     ComponentTypeIconEnum = "PhFileJsx"                     //
	ComponentTypeIconEnumPhfilelock                    ComponentTypeIconEnum = "PhFileLock"                    //
	ComponentTypeIconEnumPhfileminus                   ComponentTypeIconEnum = "PhFileMinus"                   //
	ComponentTypeIconEnumPhfilepdf                     ComponentTypeIconEnum = "PhFilePdf"                     //
	ComponentTypeIconEnumPhfileplus                    ComponentTypeIconEnum = "PhFilePlus"                    //
	ComponentTypeIconEnumPhfilepng                     ComponentTypeIconEnum = "PhFilePng"                     //
	ComponentTypeIconEnumPhfileppt                     ComponentTypeIconEnum = "PhFilePpt"                     //
	ComponentTypeIconEnumPhfilers                      ComponentTypeIconEnum = "PhFileRs"                      //
	ComponentTypeIconEnumPhfilesearch                  ComponentTypeIconEnum = "PhFileSearch"                  //
	ComponentTypeIconEnumPhfiletext                    ComponentTypeIconEnum = "PhFileText"                    //
	ComponentTypeIconEnumPhfilets                      ComponentTypeIconEnum = "PhFileTs"                      //
	ComponentTypeIconEnumPhfiletsx                     ComponentTypeIconEnum = "PhFileTsx"                     //
	ComponentTypeIconEnumPhfilevideo                   ComponentTypeIconEnum = "PhFileVideo"                   //
	ComponentTypeIconEnumPhfilevue                     ComponentTypeIconEnum = "PhFileVue"                     //
	ComponentTypeIconEnumPhfilex                       ComponentTypeIconEnum = "PhFileX"                       //
	ComponentTypeIconEnumPhfilexls                     ComponentTypeIconEnum = "PhFileXls"                     //
	ComponentTypeIconEnumPhfilezip                     ComponentTypeIconEnum = "PhFileZip"                     //
	ComponentTypeIconEnumPhfiles                       ComponentTypeIconEnum = "PhFiles"                       //
	ComponentTypeIconEnumPhfilmscript                  ComponentTypeIconEnum = "PhFilmScript"                  //
	ComponentTypeIconEnumPhfilmslate                   ComponentTypeIconEnum = "PhFilmSlate"                   //
	ComponentTypeIconEnumPhfilmstrip                   ComponentTypeIconEnum = "PhFilmStrip"                   //
	ComponentTypeIconEnumPhfingerprint                 ComponentTypeIconEnum = "PhFingerprint"                 //
	ComponentTypeIconEnumPhfingerprintsimple           ComponentTypeIconEnum = "PhFingerprintSimple"           //
	ComponentTypeIconEnumPhfinnthehuman                ComponentTypeIconEnum = "PhFinnTheHuman"                //
	ComponentTypeIconEnumPhfire                        ComponentTypeIconEnum = "PhFire"                        //
	ComponentTypeIconEnumPhfiresimple                  ComponentTypeIconEnum = "PhFireSimple"                  //
	ComponentTypeIconEnumPhfirstaid                    ComponentTypeIconEnum = "PhFirstAid"                    //
	ComponentTypeIconEnumPhfirstaidkit                 ComponentTypeIconEnum = "PhFirstAidKit"                 //
	ComponentTypeIconEnumPhfish                        ComponentTypeIconEnum = "PhFish"                        //
	ComponentTypeIconEnumPhfishsimple                  ComponentTypeIconEnum = "PhFishSimple"                  //
	ComponentTypeIconEnumPhflag                        ComponentTypeIconEnum = "PhFlag"                        //
	ComponentTypeIconEnumPhflagbanner                  ComponentTypeIconEnum = "PhFlagBanner"                  //
	ComponentTypeIconEnumPhflagcheckered               ComponentTypeIconEnum = "PhFlagCheckered"               //
	ComponentTypeIconEnumPhflame                       ComponentTypeIconEnum = "PhFlame"                       //
	ComponentTypeIconEnumPhflashlight                  ComponentTypeIconEnum = "PhFlashlight"                  //
	ComponentTypeIconEnumPhflask                       ComponentTypeIconEnum = "PhFlask"                       //
	ComponentTypeIconEnumPhfloppydisk                  ComponentTypeIconEnum = "PhFloppyDisk"                  //
	ComponentTypeIconEnumPhfloppydiskback              ComponentTypeIconEnum = "PhFloppyDiskBack"              //
	ComponentTypeIconEnumPhflowarrow                   ComponentTypeIconEnum = "PhFlowArrow"                   //
	ComponentTypeIconEnumPhflower                      ComponentTypeIconEnum = "PhFlower"                      //
	ComponentTypeIconEnumPhflowerlotus                 ComponentTypeIconEnum = "PhFlowerLotus"                 //
	ComponentTypeIconEnumPhflyingsaucer                ComponentTypeIconEnum = "PhFlyingSaucer"                //
	ComponentTypeIconEnumPhfolder                      ComponentTypeIconEnum = "PhFolder"                      //
	ComponentTypeIconEnumPhfolderdotted                ComponentTypeIconEnum = "PhFolderDotted"                //
	ComponentTypeIconEnumPhfolderlock                  ComponentTypeIconEnum = "PhFolderLock"                  //
	ComponentTypeIconEnumPhfolderminus                 ComponentTypeIconEnum = "PhFolderMinus"                 //
	ComponentTypeIconEnumPhfoldernotch                 ComponentTypeIconEnum = "PhFolderNotch"                 //
	ComponentTypeIconEnumPhfoldernotchminus            ComponentTypeIconEnum = "PhFolderNotchMinus"            //
	ComponentTypeIconEnumPhfoldernotchopen             ComponentTypeIconEnum = "PhFolderNotchOpen"             //
	ComponentTypeIconEnumPhfoldernotchplus             ComponentTypeIconEnum = "PhFolderNotchPlus"             //
	ComponentTypeIconEnumPhfolderopen                  ComponentTypeIconEnum = "PhFolderOpen"                  //
	ComponentTypeIconEnumPhfolderplus                  ComponentTypeIconEnum = "PhFolderPlus"                  //
	ComponentTypeIconEnumPhfoldersimple                ComponentTypeIconEnum = "PhFolderSimple"                //
	ComponentTypeIconEnumPhfoldersimpledotted          ComponentTypeIconEnum = "PhFolderSimpleDotted"          //
	ComponentTypeIconEnumPhfoldersimplelock            ComponentTypeIconEnum = "PhFolderSimpleLock"            //
	ComponentTypeIconEnumPhfoldersimpleminus           ComponentTypeIconEnum = "PhFolderSimpleMinus"           //
	ComponentTypeIconEnumPhfoldersimpleplus            ComponentTypeIconEnum = "PhFolderSimplePlus"            //
	ComponentTypeIconEnumPhfoldersimplestar            ComponentTypeIconEnum = "PhFolderSimpleStar"            //
	ComponentTypeIconEnumPhfoldersimpleuser            ComponentTypeIconEnum = "PhFolderSimpleUser"            //
	ComponentTypeIconEnumPhfolderstar                  ComponentTypeIconEnum = "PhFolderStar"                  //
	ComponentTypeIconEnumPhfolderuser                  ComponentTypeIconEnum = "PhFolderUser"                  //
	ComponentTypeIconEnumPhfolders                     ComponentTypeIconEnum = "PhFolders"                     //
	ComponentTypeIconEnumPhfootball                    ComponentTypeIconEnum = "PhFootball"                    //
	ComponentTypeIconEnumPhforkknife                   ComponentTypeIconEnum = "PhForkKnife"                   //
	ComponentTypeIconEnumPhframecorners                ComponentTypeIconEnum = "PhFrameCorners"                //
	ComponentTypeIconEnumPhframerlogo                  ComponentTypeIconEnum = "PhFramerLogo"                  //
	ComponentTypeIconEnumPhfunction                    ComponentTypeIconEnum = "PhFunction"                    //
	ComponentTypeIconEnumPhfunnel                      ComponentTypeIconEnum = "PhFunnel"                      //
	ComponentTypeIconEnumPhfunnelsimple                ComponentTypeIconEnum = "PhFunnelSimple"                //
	ComponentTypeIconEnumPhgamecontroller              ComponentTypeIconEnum = "PhGameController"              //
	ComponentTypeIconEnumPhgaspump                     ComponentTypeIconEnum = "PhGasPump"                     //
	ComponentTypeIconEnumPhgauge                       ComponentTypeIconEnum = "PhGauge"                       //
	ComponentTypeIconEnumPhgear                        ComponentTypeIconEnum = "PhGear"                        //
	ComponentTypeIconEnumPhgearsix                     ComponentTypeIconEnum = "PhGearSix"                     //
	ComponentTypeIconEnumPhgenderfemale                ComponentTypeIconEnum = "PhGenderFemale"                //
	ComponentTypeIconEnumPhgenderintersex              ComponentTypeIconEnum = "PhGenderIntersex"              //
	ComponentTypeIconEnumPhgendermale                  ComponentTypeIconEnum = "PhGenderMale"                  //
	ComponentTypeIconEnumPhgenderneuter                ComponentTypeIconEnum = "PhGenderNeuter"                //
	ComponentTypeIconEnumPhgendernonbinary             ComponentTypeIconEnum = "PhGenderNonbinary"             //
	ComponentTypeIconEnumPhgendertransgender           ComponentTypeIconEnum = "PhGenderTransgender"           //
	ComponentTypeIconEnumPhghost                       ComponentTypeIconEnum = "PhGhost"                       //
	ComponentTypeIconEnumPhgif                         ComponentTypeIconEnum = "PhGif"                         //
	ComponentTypeIconEnumPhgift                        ComponentTypeIconEnum = "PhGift"                        //
	ComponentTypeIconEnumPhgitbranch                   ComponentTypeIconEnum = "PhGitBranch"                   //
	ComponentTypeIconEnumPhgitcommit                   ComponentTypeIconEnum = "PhGitCommit"                   //
	ComponentTypeIconEnumPhgitdiff                     ComponentTypeIconEnum = "PhGitDiff"                     //
	ComponentTypeIconEnumPhgitfork                     ComponentTypeIconEnum = "PhGitFork"                     //
	ComponentTypeIconEnumPhgitmerge                    ComponentTypeIconEnum = "PhGitMerge"                    //
	ComponentTypeIconEnumPhgitpullrequest              ComponentTypeIconEnum = "PhGitPullRequest"              //
	ComponentTypeIconEnumPhgithublogo                  ComponentTypeIconEnum = "PhGithubLogo"                  //
	ComponentTypeIconEnumPhgitlablogo                  ComponentTypeIconEnum = "PhGitlabLogo"                  //
	ComponentTypeIconEnumPhgitlablogosimple            ComponentTypeIconEnum = "PhGitlabLogoSimple"            //
	ComponentTypeIconEnumPhglobe                       ComponentTypeIconEnum = "PhGlobe"                       //
	ComponentTypeIconEnumPhglobehemisphereeast         ComponentTypeIconEnum = "PhGlobeHemisphereEast"         //
	ComponentTypeIconEnumPhglobehemispherewest         ComponentTypeIconEnum = "PhGlobeHemisphereWest"         //
	ComponentTypeIconEnumPhglobesimple                 ComponentTypeIconEnum = "PhGlobeSimple"                 //
	ComponentTypeIconEnumPhglobestand                  ComponentTypeIconEnum = "PhGlobeStand"                  //
	ComponentTypeIconEnumPhgooglechromelogo            ComponentTypeIconEnum = "PhGoogleChromeLogo"            //
	ComponentTypeIconEnumPhgooglelogo                  ComponentTypeIconEnum = "PhGoogleLogo"                  //
	ComponentTypeIconEnumPhgooglephotoslogo            ComponentTypeIconEnum = "PhGooglePhotosLogo"            //
	ComponentTypeIconEnumPhgoogleplaylogo              ComponentTypeIconEnum = "PhGooglePlayLogo"              //
	ComponentTypeIconEnumPhgooglepodcastslogo          ComponentTypeIconEnum = "PhGooglePodcastsLogo"          //
	ComponentTypeIconEnumPhgradient                    ComponentTypeIconEnum = "PhGradient"                    //
	ComponentTypeIconEnumPhgraduationcap               ComponentTypeIconEnum = "PhGraduationCap"               //
	ComponentTypeIconEnumPhgraph                       ComponentTypeIconEnum = "PhGraph"                       //
	ComponentTypeIconEnumPhgridfour                    ComponentTypeIconEnum = "PhGridFour"                    //
	ComponentTypeIconEnumPhhamburger                   ComponentTypeIconEnum = "PhHamburger"                   //
	ComponentTypeIconEnumPhhand                        ComponentTypeIconEnum = "PhHand"                        //
	ComponentTypeIconEnumPhhandeye                     ComponentTypeIconEnum = "PhHandEye"                     //
	ComponentTypeIconEnumPhhandfist                    ComponentTypeIconEnum = "PhHandFist"                    //
	ComponentTypeIconEnumPhhandgrabbing                ComponentTypeIconEnum = "PhHandGrabbing"                //
	ComponentTypeIconEnumPhhandpalm                    ComponentTypeIconEnum = "PhHandPalm"                    //
	ComponentTypeIconEnumPhhandpointing                ComponentTypeIconEnum = "PhHandPointing"                //
	ComponentTypeIconEnumPhhandsoap                    ComponentTypeIconEnum = "PhHandSoap"                    //
	ComponentTypeIconEnumPhhandwaving                  ComponentTypeIconEnum = "PhHandWaving"                  //
	ComponentTypeIconEnumPhhandbag                     ComponentTypeIconEnum = "PhHandbag"                     //
	ComponentTypeIconEnumPhhandbagsimple               ComponentTypeIconEnum = "PhHandbagSimple"               //
	ComponentTypeIconEnumPhhandsclapping               ComponentTypeIconEnum = "PhHandsClapping"               //
	ComponentTypeIconEnumPhhandshake                   ComponentTypeIconEnum = "PhHandshake"                   //
	ComponentTypeIconEnumPhharddrive                   ComponentTypeIconEnum = "PhHardDrive"                   //
	ComponentTypeIconEnumPhharddrives                  ComponentTypeIconEnum = "PhHardDrives"                  //
	ComponentTypeIconEnumPhhash                        ComponentTypeIconEnum = "PhHash"                        //
	ComponentTypeIconEnumPhhashstraight                ComponentTypeIconEnum = "PhHashStraight"                //
	ComponentTypeIconEnumPhheadlights                  ComponentTypeIconEnum = "PhHeadlights"                  //
	ComponentTypeIconEnumPhheadphones                  ComponentTypeIconEnum = "PhHeadphones"                  //
	ComponentTypeIconEnumPhheadset                     ComponentTypeIconEnum = "PhHeadset"                     //
	ComponentTypeIconEnumPhheart                       ComponentTypeIconEnum = "PhHeart"                       //
	ComponentTypeIconEnumPhheartbreak                  ComponentTypeIconEnum = "PhHeartBreak"                  //
	ComponentTypeIconEnumPhheartstraight               ComponentTypeIconEnum = "PhHeartStraight"               //
	ComponentTypeIconEnumPhheartstraightbreak          ComponentTypeIconEnum = "PhHeartStraightBreak"          //
	ComponentTypeIconEnumPhheartbeat                   ComponentTypeIconEnum = "PhHeartbeat"                   //
	ComponentTypeIconEnumPhhexagon                     ComponentTypeIconEnum = "PhHexagon"                     //
	ComponentTypeIconEnumPhhighlightercircle           ComponentTypeIconEnum = "PhHighlighterCircle"           //
	ComponentTypeIconEnumPhhorse                       ComponentTypeIconEnum = "PhHorse"                       //
	ComponentTypeIconEnumPhhourglass                   ComponentTypeIconEnum = "PhHourglass"                   //
	ComponentTypeIconEnumPhhourglasshigh               ComponentTypeIconEnum = "PhHourglassHigh"               //
	ComponentTypeIconEnumPhhourglasslow                ComponentTypeIconEnum = "PhHourglassLow"                //
	ComponentTypeIconEnumPhhourglassmedium             ComponentTypeIconEnum = "PhHourglassMedium"             //
	ComponentTypeIconEnumPhhourglasssimple             ComponentTypeIconEnum = "PhHourglassSimple"             //
	ComponentTypeIconEnumPhhourglasssimplehigh         ComponentTypeIconEnum = "PhHourglassSimpleHigh"         //
	ComponentTypeIconEnumPhhourglasssimplelow          ComponentTypeIconEnum = "PhHourglassSimpleLow"          //
	ComponentTypeIconEnumPhhourglasssimplemedium       ComponentTypeIconEnum = "PhHourglassSimpleMedium"       //
	ComponentTypeIconEnumPhhouse                       ComponentTypeIconEnum = "PhHouse"                       //
	ComponentTypeIconEnumPhhouseline                   ComponentTypeIconEnum = "PhHouseLine"                   //
	ComponentTypeIconEnumPhhousesimple                 ComponentTypeIconEnum = "PhHouseSimple"                 //
	ComponentTypeIconEnumPhidentificationbadge         ComponentTypeIconEnum = "PhIdentificationBadge"         //
	ComponentTypeIconEnumPhidentificationcard          ComponentTypeIconEnum = "PhIdentificationCard"          //
	ComponentTypeIconEnumPhimage                       ComponentTypeIconEnum = "PhImage"                       //
	ComponentTypeIconEnumPhimagesquare                 ComponentTypeIconEnum = "PhImageSquare"                 //
	ComponentTypeIconEnumPhinfinity                    ComponentTypeIconEnum = "PhInfinity"                    //
	ComponentTypeIconEnumPhinfo                        ComponentTypeIconEnum = "PhInfo"                        //
	ComponentTypeIconEnumPhinstagramlogo               ComponentTypeIconEnum = "PhInstagramLogo"               //
	ComponentTypeIconEnumPhintersect                   ComponentTypeIconEnum = "PhIntersect"                   //
	ComponentTypeIconEnumPhjeep                        ComponentTypeIconEnum = "PhJeep"                        //
	ComponentTypeIconEnumPhkanban                      ComponentTypeIconEnum = "PhKanban"                      //
	ComponentTypeIconEnumPhkey                         ComponentTypeIconEnum = "PhKey"                         //
	ComponentTypeIconEnumPhkeyreturn                   ComponentTypeIconEnum = "PhKeyReturn"                   //
	ComponentTypeIconEnumPhkeyboard                    ComponentTypeIconEnum = "PhKeyboard"                    //
	ComponentTypeIconEnumPhkeyhole                     ComponentTypeIconEnum = "PhKeyhole"                     //
	ComponentTypeIconEnumPhknife                       ComponentTypeIconEnum = "PhKnife"                       //
	ComponentTypeIconEnumPhladder                      ComponentTypeIconEnum = "PhLadder"                      //
	ComponentTypeIconEnumPhladdersimple                ComponentTypeIconEnum = "PhLadderSimple"                //
	ComponentTypeIconEnumPhlamp                        ComponentTypeIconEnum = "PhLamp"                        //
	ComponentTypeIconEnumPhlaptop                      ComponentTypeIconEnum = "PhLaptop"                      //
	ComponentTypeIconEnumPhlayout                      ComponentTypeIconEnum = "PhLayout"                      //
	ComponentTypeIconEnumPhleaf                        ComponentTypeIconEnum = "PhLeaf"                        //
	ComponentTypeIconEnumPhlifebuoy                    ComponentTypeIconEnum = "PhLifebuoy"                    //
	ComponentTypeIconEnumPhlightbulb                   ComponentTypeIconEnum = "PhLightbulb"                   //
	ComponentTypeIconEnumPhlightbulbfilament           ComponentTypeIconEnum = "PhLightbulbFilament"           //
	ComponentTypeIconEnumPhlightning                   ComponentTypeIconEnum = "PhLightning"                   //
	ComponentTypeIconEnumPhlightningslash              ComponentTypeIconEnum = "PhLightningSlash"              //
	ComponentTypeIconEnumPhlinesegment                 ComponentTypeIconEnum = "PhLineSegment"                 //
	ComponentTypeIconEnumPhlinesegments                ComponentTypeIconEnum = "PhLineSegments"                //
	ComponentTypeIconEnumPhlink                        ComponentTypeIconEnum = "PhLink"                        //
	ComponentTypeIconEnumPhlinkbreak                   ComponentTypeIconEnum = "PhLinkBreak"                   //
	ComponentTypeIconEnumPhlinksimple                  ComponentTypeIconEnum = "PhLinkSimple"                  //
	ComponentTypeIconEnumPhlinksimplebreak             ComponentTypeIconEnum = "PhLinkSimpleBreak"             //
	ComponentTypeIconEnumPhlinksimplehorizontal        ComponentTypeIconEnum = "PhLinkSimpleHorizontal"        //
	ComponentTypeIconEnumPhlinksimplehorizontalbreak   ComponentTypeIconEnum = "PhLinkSimpleHorizontalBreak"   //
	ComponentTypeIconEnumPhlinkedinlogo                ComponentTypeIconEnum = "PhLinkedinLogo"                //
	ComponentTypeIconEnumPhlinuxlogo                   ComponentTypeIconEnum = "PhLinuxLogo"                   //
	ComponentTypeIconEnumPhlist                        ComponentTypeIconEnum = "PhList"                        //
	ComponentTypeIconEnumPhlistbullets                 ComponentTypeIconEnum = "PhListBullets"                 //
	ComponentTypeIconEnumPhlistchecks                  ComponentTypeIconEnum = "PhListChecks"                  //
	ComponentTypeIconEnumPhlistdashes                  ComponentTypeIconEnum = "PhListDashes"                  //
	ComponentTypeIconEnumPhlistnumbers                 ComponentTypeIconEnum = "PhListNumbers"                 //
	ComponentTypeIconEnumPhlistplus                    ComponentTypeIconEnum = "PhListPlus"                    //
	ComponentTypeIconEnumPhlock                        ComponentTypeIconEnum = "PhLock"                        //
	ComponentTypeIconEnumPhlockkey                     ComponentTypeIconEnum = "PhLockKey"                     //
	ComponentTypeIconEnumPhlockkeyopen                 ComponentTypeIconEnum = "PhLockKeyOpen"                 //
	ComponentTypeIconEnumPhlocklaminated               ComponentTypeIconEnum = "PhLockLaminated"               //
	ComponentTypeIconEnumPhlocklaminatedopen           ComponentTypeIconEnum = "PhLockLaminatedOpen"           //
	ComponentTypeIconEnumPhlockopen                    ComponentTypeIconEnum = "PhLockOpen"                    //
	ComponentTypeIconEnumPhlocksimple                  ComponentTypeIconEnum = "PhLockSimple"                  //
	ComponentTypeIconEnumPhlocksimpleopen              ComponentTypeIconEnum = "PhLockSimpleOpen"              //
	ComponentTypeIconEnumPhmagicwand                   ComponentTypeIconEnum = "PhMagicWand"                   //
	ComponentTypeIconEnumPhmagnet                      ComponentTypeIconEnum = "PhMagnet"                      //
	ComponentTypeIconEnumPhmagnetstraight              ComponentTypeIconEnum = "PhMagnetStraight"              //
	ComponentTypeIconEnumPhmagnifyingglass             ComponentTypeIconEnum = "PhMagnifyingGlass"             //
	ComponentTypeIconEnumPhmagnifyingglassminus        ComponentTypeIconEnum = "PhMagnifyingGlassMinus"        //
	ComponentTypeIconEnumPhmagnifyingglassplus         ComponentTypeIconEnum = "PhMagnifyingGlassPlus"         //
	ComponentTypeIconEnumPhmappin                      ComponentTypeIconEnum = "PhMapPin"                      //
	ComponentTypeIconEnumPhmappinline                  ComponentTypeIconEnum = "PhMapPinLine"                  //
	ComponentTypeIconEnumPhmaptrifold                  ComponentTypeIconEnum = "PhMapTrifold"                  //
	ComponentTypeIconEnumPhmarkercircle                ComponentTypeIconEnum = "PhMarkerCircle"                //
	ComponentTypeIconEnumPhmartini                     ComponentTypeIconEnum = "PhMartini"                     //
	ComponentTypeIconEnumPhmaskhappy                   ComponentTypeIconEnum = "PhMaskHappy"                   //
	ComponentTypeIconEnumPhmasksad                     ComponentTypeIconEnum = "PhMaskSad"                     //
	ComponentTypeIconEnumPhmathoperations              ComponentTypeIconEnum = "PhMathOperations"              //
	ComponentTypeIconEnumPhmedal                       ComponentTypeIconEnum = "PhMedal"                       //
	ComponentTypeIconEnumPhmediumlogo                  ComponentTypeIconEnum = "PhMediumLogo"                  //
	ComponentTypeIconEnumPhmegaphone                   ComponentTypeIconEnum = "PhMegaphone"                   //
	ComponentTypeIconEnumPhmegaphonesimple             ComponentTypeIconEnum = "PhMegaphoneSimple"             //
	ComponentTypeIconEnumPhmessengerlogo               ComponentTypeIconEnum = "PhMessengerLogo"               //
	ComponentTypeIconEnumPhmicrophone                  ComponentTypeIconEnum = "PhMicrophone"                  //
	ComponentTypeIconEnumPhmicrophoneslash             ComponentTypeIconEnum = "PhMicrophoneSlash"             //
	ComponentTypeIconEnumPhmicrophonestage             ComponentTypeIconEnum = "PhMicrophoneStage"             //
	ComponentTypeIconEnumPhmicrosoftexcellogo          ComponentTypeIconEnum = "PhMicrosoftExcelLogo"          //
	ComponentTypeIconEnumPhmicrosoftpowerpointlogo     ComponentTypeIconEnum = "PhMicrosoftPowerpointLogo"     //
	ComponentTypeIconEnumPhmicrosoftteamslogo          ComponentTypeIconEnum = "PhMicrosoftTeamsLogo"          //
	ComponentTypeIconEnumPhmicrosoftwordlogo           ComponentTypeIconEnum = "PhMicrosoftWordLogo"           //
	ComponentTypeIconEnumPhminus                       ComponentTypeIconEnum = "PhMinus"                       //
	ComponentTypeIconEnumPhminuscircle                 ComponentTypeIconEnum = "PhMinusCircle"                 //
	ComponentTypeIconEnumPhmoney                       ComponentTypeIconEnum = "PhMoney"                       //
	ComponentTypeIconEnumPhmonitor                     ComponentTypeIconEnum = "PhMonitor"                     //
	ComponentTypeIconEnumPhmonitorplay                 ComponentTypeIconEnum = "PhMonitorPlay"                 //
	ComponentTypeIconEnumPhmoon                        ComponentTypeIconEnum = "PhMoon"                        //
	ComponentTypeIconEnumPhmoonstars                   ComponentTypeIconEnum = "PhMoonStars"                   //
	ComponentTypeIconEnumPhmountains                   ComponentTypeIconEnum = "PhMountains"                   //
	ComponentTypeIconEnumPhmouse                       ComponentTypeIconEnum = "PhMouse"                       //
	ComponentTypeIconEnumPhmousesimple                 ComponentTypeIconEnum = "PhMouseSimple"                 //
	ComponentTypeIconEnumPhmusicnote                   ComponentTypeIconEnum = "PhMusicNote"                   //
	ComponentTypeIconEnumPhmusicnotesimple             ComponentTypeIconEnum = "PhMusicNoteSimple"             //
	ComponentTypeIconEnumPhmusicnotes                  ComponentTypeIconEnum = "PhMusicNotes"                  //
	ComponentTypeIconEnumPhmusicnotesplus              ComponentTypeIconEnum = "PhMusicNotesPlus"              //
	ComponentTypeIconEnumPhmusicnotessimple            ComponentTypeIconEnum = "PhMusicNotesSimple"            //
	ComponentTypeIconEnumPhnavigationarrow             ComponentTypeIconEnum = "PhNavigationArrow"             //
	ComponentTypeIconEnumPhneedle                      ComponentTypeIconEnum = "PhNeedle"                      //
	ComponentTypeIconEnumPhnewspaper                   ComponentTypeIconEnum = "PhNewspaper"                   //
	ComponentTypeIconEnumPhnewspaperclipping           ComponentTypeIconEnum = "PhNewspaperClipping"           //
	ComponentTypeIconEnumPhnote                        ComponentTypeIconEnum = "PhNote"                        //
	ComponentTypeIconEnumPhnoteblank                   ComponentTypeIconEnum = "PhNoteBlank"                   //
	ComponentTypeIconEnumPhnotepencil                  ComponentTypeIconEnum = "PhNotePencil"                  //
	ComponentTypeIconEnumPhnotebook                    ComponentTypeIconEnum = "PhNotebook"                    //
	ComponentTypeIconEnumPhnotepad                     ComponentTypeIconEnum = "PhNotepad"                     //
	ComponentTypeIconEnumPhnotification                ComponentTypeIconEnum = "PhNotification"                //
	ComponentTypeIconEnumPhnumbercircleeight           ComponentTypeIconEnum = "PhNumberCircleEight"           //
	ComponentTypeIconEnumPhnumbercirclefive            ComponentTypeIconEnum = "PhNumberCircleFive"            //
	ComponentTypeIconEnumPhnumbercirclefour            ComponentTypeIconEnum = "PhNumberCircleFour"            //
	ComponentTypeIconEnumPhnumbercirclenine            ComponentTypeIconEnum = "PhNumberCircleNine"            //
	ComponentTypeIconEnumPhnumbercircleone             ComponentTypeIconEnum = "PhNumberCircleOne"             //
	ComponentTypeIconEnumPhnumbercircleseven           ComponentTypeIconEnum = "PhNumberCircleSeven"           //
	ComponentTypeIconEnumPhnumbercirclesix             ComponentTypeIconEnum = "PhNumberCircleSix"             //
	ComponentTypeIconEnumPhnumbercirclethree           ComponentTypeIconEnum = "PhNumberCircleThree"           //
	ComponentTypeIconEnumPhnumbercircletwo             ComponentTypeIconEnum = "PhNumberCircleTwo"             //
	ComponentTypeIconEnumPhnumbercirclezero            ComponentTypeIconEnum = "PhNumberCircleZero"            //
	ComponentTypeIconEnumPhnumbereight                 ComponentTypeIconEnum = "PhNumberEight"                 //
	ComponentTypeIconEnumPhnumberfive                  ComponentTypeIconEnum = "PhNumberFive"                  //
	ComponentTypeIconEnumPhnumberfour                  ComponentTypeIconEnum = "PhNumberFour"                  //
	ComponentTypeIconEnumPhnumbernine                  ComponentTypeIconEnum = "PhNumberNine"                  //
	ComponentTypeIconEnumPhnumberone                   ComponentTypeIconEnum = "PhNumberOne"                   //
	ComponentTypeIconEnumPhnumberseven                 ComponentTypeIconEnum = "PhNumberSeven"                 //
	ComponentTypeIconEnumPhnumbersix                   ComponentTypeIconEnum = "PhNumberSix"                   //
	ComponentTypeIconEnumPhnumbersquareeight           ComponentTypeIconEnum = "PhNumberSquareEight"           //
	ComponentTypeIconEnumPhnumbersquarefive            ComponentTypeIconEnum = "PhNumberSquareFive"            //
	ComponentTypeIconEnumPhnumbersquarefour            ComponentTypeIconEnum = "PhNumberSquareFour"            //
	ComponentTypeIconEnumPhnumbersquarenine            ComponentTypeIconEnum = "PhNumberSquareNine"            //
	ComponentTypeIconEnumPhnumbersquareone             ComponentTypeIconEnum = "PhNumberSquareOne"             //
	ComponentTypeIconEnumPhnumbersquareseven           ComponentTypeIconEnum = "PhNumberSquareSeven"           //
	ComponentTypeIconEnumPhnumbersquaresix             ComponentTypeIconEnum = "PhNumberSquareSix"             //
	ComponentTypeIconEnumPhnumbersquarethree           ComponentTypeIconEnum = "PhNumberSquareThree"           //
	ComponentTypeIconEnumPhnumbersquaretwo             ComponentTypeIconEnum = "PhNumberSquareTwo"             //
	ComponentTypeIconEnumPhnumbersquarezero            ComponentTypeIconEnum = "PhNumberSquareZero"            //
	ComponentTypeIconEnumPhnumberthree                 ComponentTypeIconEnum = "PhNumberThree"                 //
	ComponentTypeIconEnumPhnumbertwo                   ComponentTypeIconEnum = "PhNumberTwo"                   //
	ComponentTypeIconEnumPhnumberzero                  ComponentTypeIconEnum = "PhNumberZero"                  //
	ComponentTypeIconEnumPhnut                         ComponentTypeIconEnum = "PhNut"                         //
	ComponentTypeIconEnumPhnytimeslogo                 ComponentTypeIconEnum = "PhNyTimesLogo"                 //
	ComponentTypeIconEnumPhoctagon                     ComponentTypeIconEnum = "PhOctagon"                     //
	ComponentTypeIconEnumPhoption                      ComponentTypeIconEnum = "PhOption"                      //
	ComponentTypeIconEnumPhpackage                     ComponentTypeIconEnum = "PhPackage"                     //
	ComponentTypeIconEnumPhpaintbrush                  ComponentTypeIconEnum = "PhPaintBrush"                  //
	ComponentTypeIconEnumPhpaintbrushbroad             ComponentTypeIconEnum = "PhPaintBrushBroad"             //
	ComponentTypeIconEnumPhpaintbrushhousehold         ComponentTypeIconEnum = "PhPaintBrushHousehold"         //
	ComponentTypeIconEnumPhpaintbucket                 ComponentTypeIconEnum = "PhPaintBucket"                 //
	ComponentTypeIconEnumPhpaintroller                 ComponentTypeIconEnum = "PhPaintRoller"                 //
	ComponentTypeIconEnumPhpalette                     ComponentTypeIconEnum = "PhPalette"                     //
	ComponentTypeIconEnumPhpaperplane                  ComponentTypeIconEnum = "PhPaperPlane"                  //
	ComponentTypeIconEnumPhpaperplaneright             ComponentTypeIconEnum = "PhPaperPlaneRight"             //
	ComponentTypeIconEnumPhpaperplanetilt              ComponentTypeIconEnum = "PhPaperPlaneTilt"              //
	ComponentTypeIconEnumPhpaperclip                   ComponentTypeIconEnum = "PhPaperclip"                   //
	ComponentTypeIconEnumPhpapercliphorizontal         ComponentTypeIconEnum = "PhPaperclipHorizontal"         //
	ComponentTypeIconEnumPhparachute                   ComponentTypeIconEnum = "PhParachute"                   //
	ComponentTypeIconEnumPhpassword                    ComponentTypeIconEnum = "PhPassword"                    //
	ComponentTypeIconEnumPhpath                        ComponentTypeIconEnum = "PhPath"                        //
	ComponentTypeIconEnumPhpause                       ComponentTypeIconEnum = "PhPause"                       //
	ComponentTypeIconEnumPhpausecircle                 ComponentTypeIconEnum = "PhPauseCircle"                 //
	ComponentTypeIconEnumPhpawprint                    ComponentTypeIconEnum = "PhPawPrint"                    //
	ComponentTypeIconEnumPhpeace                       ComponentTypeIconEnum = "PhPeace"                       //
	ComponentTypeIconEnumPhpen                         ComponentTypeIconEnum = "PhPen"                         //
	ComponentTypeIconEnumPhpennib                      ComponentTypeIconEnum = "PhPenNib"                      //
	ComponentTypeIconEnumPhpennibstraight              ComponentTypeIconEnum = "PhPenNibStraight"              //
	ComponentTypeIconEnumPhpencil                      ComponentTypeIconEnum = "PhPencil"                      //
	ComponentTypeIconEnumPhpencilcircle                ComponentTypeIconEnum = "PhPencilCircle"                //
	ComponentTypeIconEnumPhpencilline                  ComponentTypeIconEnum = "PhPencilLine"                  //
	ComponentTypeIconEnumPhpencilsimple                ComponentTypeIconEnum = "PhPencilSimple"                //
	ComponentTypeIconEnumPhpencilsimpleline            ComponentTypeIconEnum = "PhPencilSimpleLine"            //
	ComponentTypeIconEnumPhpercent                     ComponentTypeIconEnum = "PhPercent"                     //
	ComponentTypeIconEnumPhperson                      ComponentTypeIconEnum = "PhPerson"                      //
	ComponentTypeIconEnumPhpersonsimple                ComponentTypeIconEnum = "PhPersonSimple"                //
	ComponentTypeIconEnumPhpersonsimplerun             ComponentTypeIconEnum = "PhPersonSimpleRun"             //
	ComponentTypeIconEnumPhpersonsimplewalk            ComponentTypeIconEnum = "PhPersonSimpleWalk"            //
	ComponentTypeIconEnumPhperspective                 ComponentTypeIconEnum = "PhPerspective"                 //
	ComponentTypeIconEnumPhphone                       ComponentTypeIconEnum = "PhPhone"                       //
	ComponentTypeIconEnumPhphonecall                   ComponentTypeIconEnum = "PhPhoneCall"                   //
	ComponentTypeIconEnumPhphonedisconnect             ComponentTypeIconEnum = "PhPhoneDisconnect"             //
	ComponentTypeIconEnumPhphoneincoming               ComponentTypeIconEnum = "PhPhoneIncoming"               //
	ComponentTypeIconEnumPhphoneoutgoing               ComponentTypeIconEnum = "PhPhoneOutgoing"               //
	ComponentTypeIconEnumPhphoneslash                  ComponentTypeIconEnum = "PhPhoneSlash"                  //
	ComponentTypeIconEnumPhphonex                      ComponentTypeIconEnum = "PhPhoneX"                      //
	ComponentTypeIconEnumPhphosphorlogo                ComponentTypeIconEnum = "PhPhosphorLogo"                //
	ComponentTypeIconEnumPhpianokeys                   ComponentTypeIconEnum = "PhPianoKeys"                   //
	ComponentTypeIconEnumPhpictureinpicture            ComponentTypeIconEnum = "PhPictureInPicture"            //
	ComponentTypeIconEnumPhpill                        ComponentTypeIconEnum = "PhPill"                        //
	ComponentTypeIconEnumPhpinterestlogo               ComponentTypeIconEnum = "PhPinterestLogo"               //
	ComponentTypeIconEnumPhpinwheel                    ComponentTypeIconEnum = "PhPinwheel"                    //
	ComponentTypeIconEnumPhpizza                       ComponentTypeIconEnum = "PhPizza"                       //
	ComponentTypeIconEnumPhplaceholder                 ComponentTypeIconEnum = "PhPlaceholder"                 //
	ComponentTypeIconEnumPhplanet                      ComponentTypeIconEnum = "PhPlanet"                      //
	ComponentTypeIconEnumPhplay                        ComponentTypeIconEnum = "PhPlay"                        //
	ComponentTypeIconEnumPhplaycircle                  ComponentTypeIconEnum = "PhPlayCircle"                  //
	ComponentTypeIconEnumPhplaylist                    ComponentTypeIconEnum = "PhPlaylist"                    //
	ComponentTypeIconEnumPhplug                        ComponentTypeIconEnum = "PhPlug"                        //
	ComponentTypeIconEnumPhplugs                       ComponentTypeIconEnum = "PhPlugs"                       //
	ComponentTypeIconEnumPhplugsconnected              ComponentTypeIconEnum = "PhPlugsConnected"              //
	ComponentTypeIconEnumPhplus                        ComponentTypeIconEnum = "PhPlus"                        //
	ComponentTypeIconEnumPhpluscircle                  ComponentTypeIconEnum = "PhPlusCircle"                  //
	ComponentTypeIconEnumPhplusminus                   ComponentTypeIconEnum = "PhPlusMinus"                   //
	ComponentTypeIconEnumPhpokerchip                   ComponentTypeIconEnum = "PhPokerChip"                   //
	ComponentTypeIconEnumPhpolicecar                   ComponentTypeIconEnum = "PhPoliceCar"                   //
	ComponentTypeIconEnumPhpolygon                     ComponentTypeIconEnum = "PhPolygon"                     //
	ComponentTypeIconEnumPhpopcorn                     ComponentTypeIconEnum = "PhPopcorn"                     //
	ComponentTypeIconEnumPhpower                       ComponentTypeIconEnum = "PhPower"                       //
	ComponentTypeIconEnumPhprescription                ComponentTypeIconEnum = "PhPrescription"                //
	ComponentTypeIconEnumPhpresentation                ComponentTypeIconEnum = "PhPresentation"                //
	ComponentTypeIconEnumPhpresentationchart           ComponentTypeIconEnum = "PhPresentationChart"           //
	ComponentTypeIconEnumPhprinter                     ComponentTypeIconEnum = "PhPrinter"                     //
	ComponentTypeIconEnumPhprohibit                    ComponentTypeIconEnum = "PhProhibit"                    //
	ComponentTypeIconEnumPhprohibitinset               ComponentTypeIconEnum = "PhProhibitInset"               //
	ComponentTypeIconEnumPhprojectorscreen             ComponentTypeIconEnum = "PhProjectorScreen"             //
	ComponentTypeIconEnumPhprojectorscreenchart        ComponentTypeIconEnum = "PhProjectorScreenChart"        //
	ComponentTypeIconEnumPhpushpin                     ComponentTypeIconEnum = "PhPushPin"                     //
	ComponentTypeIconEnumPhpushpinsimple               ComponentTypeIconEnum = "PhPushPinSimple"               //
	ComponentTypeIconEnumPhpushpinsimpleslash          ComponentTypeIconEnum = "PhPushPinSimpleSlash"          //
	ComponentTypeIconEnumPhpushpinslash                ComponentTypeIconEnum = "PhPushPinSlash"                //
	ComponentTypeIconEnumPhpuzzlepiece                 ComponentTypeIconEnum = "PhPuzzlePiece"                 //
	ComponentTypeIconEnumPhqrcode                      ComponentTypeIconEnum = "PhQrCode"                      //
	ComponentTypeIconEnumPhquestion                    ComponentTypeIconEnum = "PhQuestion"                    //
	ComponentTypeIconEnumPhqueue                       ComponentTypeIconEnum = "PhQueue"                       //
	ComponentTypeIconEnumPhquotes                      ComponentTypeIconEnum = "PhQuotes"                      //
	ComponentTypeIconEnumPhradical                     ComponentTypeIconEnum = "PhRadical"                     //
	ComponentTypeIconEnumPhradio                       ComponentTypeIconEnum = "PhRadio"                       //
	ComponentTypeIconEnumPhradiobutton                 ComponentTypeIconEnum = "PhRadioButton"                 //
	ComponentTypeIconEnumPhrainbow                     ComponentTypeIconEnum = "PhRainbow"                     //
	ComponentTypeIconEnumPhrainbowcloud                ComponentTypeIconEnum = "PhRainbowCloud"                //
	ComponentTypeIconEnumPhreceipt                     ComponentTypeIconEnum = "PhReceipt"                     //
	ComponentTypeIconEnumPhrecord                      ComponentTypeIconEnum = "PhRecord"                      //
	ComponentTypeIconEnumPhrectangle                   ComponentTypeIconEnum = "PhRectangle"                   //
	ComponentTypeIconEnumPhrecycle                     ComponentTypeIconEnum = "PhRecycle"                     //
	ComponentTypeIconEnumPhredditlogo                  ComponentTypeIconEnum = "PhRedditLogo"                  //
	ComponentTypeIconEnumPhrepeat                      ComponentTypeIconEnum = "PhRepeat"                      //
	ComponentTypeIconEnumPhrepeatonce                  ComponentTypeIconEnum = "PhRepeatOnce"                  //
	ComponentTypeIconEnumPhrewind                      ComponentTypeIconEnum = "PhRewind"                      //
	ComponentTypeIconEnumPhrewindcircle                ComponentTypeIconEnum = "PhRewindCircle"                //
	ComponentTypeIconEnumPhrobot                       ComponentTypeIconEnum = "PhRobot"                       //
	ComponentTypeIconEnumPhrocket                      ComponentTypeIconEnum = "PhRocket"                      //
	ComponentTypeIconEnumPhrocketlaunch                ComponentTypeIconEnum = "PhRocketLaunch"                //
	ComponentTypeIconEnumPhrows                        ComponentTypeIconEnum = "PhRows"                        //
	ComponentTypeIconEnumPhrss                         ComponentTypeIconEnum = "PhRss"                         //
	ComponentTypeIconEnumPhrsssimple                   ComponentTypeIconEnum = "PhRssSimple"                   //
	ComponentTypeIconEnumPhrug                         ComponentTypeIconEnum = "PhRug"                         //
	ComponentTypeIconEnumPhruler                       ComponentTypeIconEnum = "PhRuler"                       //
	ComponentTypeIconEnumPhscales                      ComponentTypeIconEnum = "PhScales"                      //
	ComponentTypeIconEnumPhscan                        ComponentTypeIconEnum = "PhScan"                        //
	ComponentTypeIconEnumPhscissors                    ComponentTypeIconEnum = "PhScissors"                    //
	ComponentTypeIconEnumPhscreencast                  ComponentTypeIconEnum = "PhScreencast"                  //
	ComponentTypeIconEnumPhscribbleloop                ComponentTypeIconEnum = "PhScribbleLoop"                //
	ComponentTypeIconEnumPhscroll                      ComponentTypeIconEnum = "PhScroll"                      //
	ComponentTypeIconEnumPhselection                   ComponentTypeIconEnum = "PhSelection"                   //
	ComponentTypeIconEnumPhselectionall                ComponentTypeIconEnum = "PhSelectionAll"                //
	ComponentTypeIconEnumPhselectionbackground         ComponentTypeIconEnum = "PhSelectionBackground"         //
	ComponentTypeIconEnumPhselectionforeground         ComponentTypeIconEnum = "PhSelectionForeground"         //
	ComponentTypeIconEnumPhselectioninverse            ComponentTypeIconEnum = "PhSelectionInverse"            //
	ComponentTypeIconEnumPhselectionplus               ComponentTypeIconEnum = "PhSelectionPlus"               //
	ComponentTypeIconEnumPhselectionslash              ComponentTypeIconEnum = "PhSelectionSlash"              //
	ComponentTypeIconEnumPhshare                       ComponentTypeIconEnum = "PhShare"                       //
	ComponentTypeIconEnumPhsharenetwork                ComponentTypeIconEnum = "PhShareNetwork"                //
	ComponentTypeIconEnumPhshield                      ComponentTypeIconEnum = "PhShield"                      //
	ComponentTypeIconEnumPhshieldcheck                 ComponentTypeIconEnum = "PhShieldCheck"                 //
	ComponentTypeIconEnumPhshieldcheckered             ComponentTypeIconEnum = "PhShieldCheckered"             //
	ComponentTypeIconEnumPhshieldchevron               ComponentTypeIconEnum = "PhShieldChevron"               //
	ComponentTypeIconEnumPhshieldplus                  ComponentTypeIconEnum = "PhShieldPlus"                  //
	ComponentTypeIconEnumPhshieldslash                 ComponentTypeIconEnum = "PhShieldSlash"                 //
	ComponentTypeIconEnumPhshieldstar                  ComponentTypeIconEnum = "PhShieldStar"                  //
	ComponentTypeIconEnumPhshieldwarning               ComponentTypeIconEnum = "PhShieldWarning"               //
	ComponentTypeIconEnumPhshoppingbag                 ComponentTypeIconEnum = "PhShoppingBag"                 //
	ComponentTypeIconEnumPhshoppingbagopen             ComponentTypeIconEnum = "PhShoppingBagOpen"             //
	ComponentTypeIconEnumPhshoppingcart                ComponentTypeIconEnum = "PhShoppingCart"                //
	ComponentTypeIconEnumPhshoppingcartsimple          ComponentTypeIconEnum = "PhShoppingCartSimple"          //
	ComponentTypeIconEnumPhshower                      ComponentTypeIconEnum = "PhShower"                      //
	ComponentTypeIconEnumPhshuffle                     ComponentTypeIconEnum = "PhShuffle"                     //
	ComponentTypeIconEnumPhshuffleangular              ComponentTypeIconEnum = "PhShuffleAngular"              //
	ComponentTypeIconEnumPhshufflesimple               ComponentTypeIconEnum = "PhShuffleSimple"               //
	ComponentTypeIconEnumPhsidebar                     ComponentTypeIconEnum = "PhSidebar"                     //
	ComponentTypeIconEnumPhsidebarsimple               ComponentTypeIconEnum = "PhSidebarSimple"               //
	ComponentTypeIconEnumPhsignin                      ComponentTypeIconEnum = "PhSignIn"                      //
	ComponentTypeIconEnumPhsignout                     ComponentTypeIconEnum = "PhSignOut"                     //
	ComponentTypeIconEnumPhsignpost                    ComponentTypeIconEnum = "PhSignpost"                    //
	ComponentTypeIconEnumPhsimcard                     ComponentTypeIconEnum = "PhSimCard"                     //
	ComponentTypeIconEnumPhsketchlogo                  ComponentTypeIconEnum = "PhSketchLogo"                  //
	ComponentTypeIconEnumPhskipback                    ComponentTypeIconEnum = "PhSkipBack"                    //
	ComponentTypeIconEnumPhskipbackcircle              ComponentTypeIconEnum = "PhSkipBackCircle"              //
	ComponentTypeIconEnumPhskipforward                 ComponentTypeIconEnum = "PhSkipForward"                 //
	ComponentTypeIconEnumPhskipforwardcircle           ComponentTypeIconEnum = "PhSkipForwardCircle"           //
	ComponentTypeIconEnumPhskull                       ComponentTypeIconEnum = "PhSkull"                       //
	ComponentTypeIconEnumPhslacklogo                   ComponentTypeIconEnum = "PhSlackLogo"                   //
	ComponentTypeIconEnumPhsliders                     ComponentTypeIconEnum = "PhSliders"                     //
	ComponentTypeIconEnumPhslidershorizontal           ComponentTypeIconEnum = "PhSlidersHorizontal"           //
	ComponentTypeIconEnumPhsmiley                      ComponentTypeIconEnum = "PhSmiley"                      //
	ComponentTypeIconEnumPhsmileyblank                 ComponentTypeIconEnum = "PhSmileyBlank"                 //
	ComponentTypeIconEnumPhsmileymeh                   ComponentTypeIconEnum = "PhSmileyMeh"                   //
	ComponentTypeIconEnumPhsmileynervous               ComponentTypeIconEnum = "PhSmileyNervous"               //
	ComponentTypeIconEnumPhsmileysad                   ComponentTypeIconEnum = "PhSmileySad"                   //
	ComponentTypeIconEnumPhsmileysticker               ComponentTypeIconEnum = "PhSmileySticker"               //
	ComponentTypeIconEnumPhsmileywink                  ComponentTypeIconEnum = "PhSmileyWink"                  //
	ComponentTypeIconEnumPhsmileyxeyes                 ComponentTypeIconEnum = "PhSmileyXEyes"                 //
	ComponentTypeIconEnumPhsnapchatlogo                ComponentTypeIconEnum = "PhSnapchatLogo"                //
	ComponentTypeIconEnumPhsnowflake                   ComponentTypeIconEnum = "PhSnowflake"                   //
	ComponentTypeIconEnumPhsoccerball                  ComponentTypeIconEnum = "PhSoccerBall"                  //
	ComponentTypeIconEnumPhsortascending               ComponentTypeIconEnum = "PhSortAscending"               //
	ComponentTypeIconEnumPhsortdescending              ComponentTypeIconEnum = "PhSortDescending"              //
	ComponentTypeIconEnumPhspade                       ComponentTypeIconEnum = "PhSpade"                       //
	ComponentTypeIconEnumPhsparkle                     ComponentTypeIconEnum = "PhSparkle"                     //
	ComponentTypeIconEnumPhspeakerhigh                 ComponentTypeIconEnum = "PhSpeakerHigh"                 //
	ComponentTypeIconEnumPhspeakerlow                  ComponentTypeIconEnum = "PhSpeakerLow"                  //
	ComponentTypeIconEnumPhspeakernone                 ComponentTypeIconEnum = "PhSpeakerNone"                 //
	ComponentTypeIconEnumPhspeakersimplehigh           ComponentTypeIconEnum = "PhSpeakerSimpleHigh"           //
	ComponentTypeIconEnumPhspeakersimplelow            ComponentTypeIconEnum = "PhSpeakerSimpleLow"            //
	ComponentTypeIconEnumPhspeakersimplenone           ComponentTypeIconEnum = "PhSpeakerSimpleNone"           //
	ComponentTypeIconEnumPhspeakersimpleslash          ComponentTypeIconEnum = "PhSpeakerSimpleSlash"          //
	ComponentTypeIconEnumPhspeakersimplex              ComponentTypeIconEnum = "PhSpeakerSimpleX"              //
	ComponentTypeIconEnumPhspeakerslash                ComponentTypeIconEnum = "PhSpeakerSlash"                //
	ComponentTypeIconEnumPhspeakerx                    ComponentTypeIconEnum = "PhSpeakerX"                    //
	ComponentTypeIconEnumPhspinner                     ComponentTypeIconEnum = "PhSpinner"                     //
	ComponentTypeIconEnumPhspinnergap                  ComponentTypeIconEnum = "PhSpinnerGap"                  //
	ComponentTypeIconEnumPhspiral                      ComponentTypeIconEnum = "PhSpiral"                      //
	ComponentTypeIconEnumPhspotifylogo                 ComponentTypeIconEnum = "PhSpotifyLogo"                 //
	ComponentTypeIconEnumPhsquare                      ComponentTypeIconEnum = "PhSquare"                      //
	ComponentTypeIconEnumPhsquarehalf                  ComponentTypeIconEnum = "PhSquareHalf"                  //
	ComponentTypeIconEnumPhsquarehalfbottom            ComponentTypeIconEnum = "PhSquareHalfBottom"            //
	ComponentTypeIconEnumPhsquarelogo                  ComponentTypeIconEnum = "PhSquareLogo"                  //
	ComponentTypeIconEnumPhsquaresfour                 ComponentTypeIconEnum = "PhSquaresFour"                 //
	ComponentTypeIconEnumPhstack                       ComponentTypeIconEnum = "PhStack"                       //
	ComponentTypeIconEnumPhstackoverflowlogo           ComponentTypeIconEnum = "PhStackOverflowLogo"           //
	ComponentTypeIconEnumPhstacksimple                 ComponentTypeIconEnum = "PhStackSimple"                 //
	ComponentTypeIconEnumPhstamp                       ComponentTypeIconEnum = "PhStamp"                       //
	ComponentTypeIconEnumPhstar                        ComponentTypeIconEnum = "PhStar"                        //
	ComponentTypeIconEnumPhstarfour                    ComponentTypeIconEnum = "PhStarFour"                    //
	ComponentTypeIconEnumPhstarhalf                    ComponentTypeIconEnum = "PhStarHalf"                    //
	ComponentTypeIconEnumPhsticker                     ComponentTypeIconEnum = "PhSticker"                     //
	ComponentTypeIconEnumPhstop                        ComponentTypeIconEnum = "PhStop"                        //
	ComponentTypeIconEnumPhstopcircle                  ComponentTypeIconEnum = "PhStopCircle"                  //
	ComponentTypeIconEnumPhstorefront                  ComponentTypeIconEnum = "PhStorefront"                  //
	ComponentTypeIconEnumPhstrategy                    ComponentTypeIconEnum = "PhStrategy"                    //
	ComponentTypeIconEnumPhstripelogo                  ComponentTypeIconEnum = "PhStripeLogo"                  //
	ComponentTypeIconEnumPhstudent                     ComponentTypeIconEnum = "PhStudent"                     //
	ComponentTypeIconEnumPhsuitcase                    ComponentTypeIconEnum = "PhSuitcase"                    //
	ComponentTypeIconEnumPhsuitcasesimple              ComponentTypeIconEnum = "PhSuitcaseSimple"              //
	ComponentTypeIconEnumPhsun                         ComponentTypeIconEnum = "PhSun"                         //
	ComponentTypeIconEnumPhsundim                      ComponentTypeIconEnum = "PhSunDim"                      //
	ComponentTypeIconEnumPhsunhorizon                  ComponentTypeIconEnum = "PhSunHorizon"                  //
	ComponentTypeIconEnumPhsunglasses                  ComponentTypeIconEnum = "PhSunglasses"                  //
	ComponentTypeIconEnumPhswap                        ComponentTypeIconEnum = "PhSwap"                        //
	ComponentTypeIconEnumPhswatches                    ComponentTypeIconEnum = "PhSwatches"                    //
	ComponentTypeIconEnumPhsword                       ComponentTypeIconEnum = "PhSword"                       //
	ComponentTypeIconEnumPhsyringe                     ComponentTypeIconEnum = "PhSyringe"                     //
	ComponentTypeIconEnumPhtshirt                      ComponentTypeIconEnum = "PhTShirt"                      //
	ComponentTypeIconEnumPhtable                       ComponentTypeIconEnum = "PhTable"                       //
	ComponentTypeIconEnumPhtabs                        ComponentTypeIconEnum = "PhTabs"                        //
	ComponentTypeIconEnumPhtag                         ComponentTypeIconEnum = "PhTag"                         //
	ComponentTypeIconEnumPhtagchevron                  ComponentTypeIconEnum = "PhTagChevron"                  //
	ComponentTypeIconEnumPhtagsimple                   ComponentTypeIconEnum = "PhTagSimple"                   //
	ComponentTypeIconEnumPhtarget                      ComponentTypeIconEnum = "PhTarget"                      //
	ComponentTypeIconEnumPhtaxi                        ComponentTypeIconEnum = "PhTaxi"                        //
	ComponentTypeIconEnumPhtelegramlogo                ComponentTypeIconEnum = "PhTelegramLogo"                //
	ComponentTypeIconEnumPhtelevision                  ComponentTypeIconEnum = "PhTelevision"                  //
	ComponentTypeIconEnumPhtelevisionsimple            ComponentTypeIconEnum = "PhTelevisionSimple"            //
	ComponentTypeIconEnumPhtennisball                  ComponentTypeIconEnum = "PhTennisBall"                  //
	ComponentTypeIconEnumPhterminal                    ComponentTypeIconEnum = "PhTerminal"                    //
	ComponentTypeIconEnumPhterminalwindow              ComponentTypeIconEnum = "PhTerminalWindow"              //
	ComponentTypeIconEnumPhtesttube                    ComponentTypeIconEnum = "PhTestTube"                    //
	ComponentTypeIconEnumPhtextaa                      ComponentTypeIconEnum = "PhTextAa"                      //
	ComponentTypeIconEnumPhtextaligncenter             ComponentTypeIconEnum = "PhTextAlignCenter"             //
	ComponentTypeIconEnumPhtextalignjustify            ComponentTypeIconEnum = "PhTextAlignJustify"            //
	ComponentTypeIconEnumPhtextalignleft               ComponentTypeIconEnum = "PhTextAlignLeft"               //
	ComponentTypeIconEnumPhtextalignright              ComponentTypeIconEnum = "PhTextAlignRight"              //
	ComponentTypeIconEnumPhtextbolder                  ComponentTypeIconEnum = "PhTextBolder"                  //
	ComponentTypeIconEnumPhtexth                       ComponentTypeIconEnum = "PhTextH"                       //
	ComponentTypeIconEnumPhtexthfive                   ComponentTypeIconEnum = "PhTextHFive"                   //
	ComponentTypeIconEnumPhtexthfour                   ComponentTypeIconEnum = "PhTextHFour"                   //
	ComponentTypeIconEnumPhtexthone                    ComponentTypeIconEnum = "PhTextHOne"                    //
	ComponentTypeIconEnumPhtexthsix                    ComponentTypeIconEnum = "PhTextHSix"                    //
	ComponentTypeIconEnumPhtexththree                  ComponentTypeIconEnum = "PhTextHThree"                  //
	ComponentTypeIconEnumPhtexthtwo                    ComponentTypeIconEnum = "PhTextHTwo"                    //
	ComponentTypeIconEnumPhtextindent                  ComponentTypeIconEnum = "PhTextIndent"                  //
	ComponentTypeIconEnumPhtextitalic                  ComponentTypeIconEnum = "PhTextItalic"                  //
	ComponentTypeIconEnumPhtextoutdent                 ComponentTypeIconEnum = "PhTextOutdent"                 //
	ComponentTypeIconEnumPhtextstrikethrough           ComponentTypeIconEnum = "PhTextStrikethrough"           //
	ComponentTypeIconEnumPhtextt                       ComponentTypeIconEnum = "PhTextT"                       //
	ComponentTypeIconEnumPhtextunderline               ComponentTypeIconEnum = "PhTextUnderline"               //
	ComponentTypeIconEnumPhtextbox                     ComponentTypeIconEnum = "PhTextbox"                     //
	ComponentTypeIconEnumPhthermometer                 ComponentTypeIconEnum = "PhThermometer"                 //
	ComponentTypeIconEnumPhthermometercold             ComponentTypeIconEnum = "PhThermometerCold"             //
	ComponentTypeIconEnumPhthermometerhot              ComponentTypeIconEnum = "PhThermometerHot"              //
	ComponentTypeIconEnumPhthermometersimple           ComponentTypeIconEnum = "PhThermometerSimple"           //
	ComponentTypeIconEnumPhthumbsdown                  ComponentTypeIconEnum = "PhThumbsDown"                  //
	ComponentTypeIconEnumPhthumbsup                    ComponentTypeIconEnum = "PhThumbsUp"                    //
	ComponentTypeIconEnumPhticket                      ComponentTypeIconEnum = "PhTicket"                      //
	ComponentTypeIconEnumPhtiktoklogo                  ComponentTypeIconEnum = "PhTiktokLogo"                  //
	ComponentTypeIconEnumPhtimer                       ComponentTypeIconEnum = "PhTimer"                       //
	ComponentTypeIconEnumPhtoggleleft                  ComponentTypeIconEnum = "PhToggleLeft"                  //
	ComponentTypeIconEnumPhtoggleright                 ComponentTypeIconEnum = "PhToggleRight"                 //
	ComponentTypeIconEnumPhtoilet                      ComponentTypeIconEnum = "PhToilet"                      //
	ComponentTypeIconEnumPhtoiletpaper                 ComponentTypeIconEnum = "PhToiletPaper"                 //
	ComponentTypeIconEnumPhtote                        ComponentTypeIconEnum = "PhTote"                        //
	ComponentTypeIconEnumPhtotesimple                  ComponentTypeIconEnum = "PhToteSimple"                  //
	ComponentTypeIconEnumPhtrademarkregistered         ComponentTypeIconEnum = "PhTrademarkRegistered"         //
	ComponentTypeIconEnumPhtrafficcone                 ComponentTypeIconEnum = "PhTrafficCone"                 //
	ComponentTypeIconEnumPhtrafficsign                 ComponentTypeIconEnum = "PhTrafficSign"                 //
	ComponentTypeIconEnumPhtrafficsignal               ComponentTypeIconEnum = "PhTrafficSignal"               //
	ComponentTypeIconEnumPhtrain                       ComponentTypeIconEnum = "PhTrain"                       //
	ComponentTypeIconEnumPhtrainregional               ComponentTypeIconEnum = "PhTrainRegional"               //
	ComponentTypeIconEnumPhtrainsimple                 ComponentTypeIconEnum = "PhTrainSimple"                 //
	ComponentTypeIconEnumPhtranslate                   ComponentTypeIconEnum = "PhTranslate"                   //
	ComponentTypeIconEnumPhtrash                       ComponentTypeIconEnum = "PhTrash"                       //
	ComponentTypeIconEnumPhtrashsimple                 ComponentTypeIconEnum = "PhTrashSimple"                 //
	ComponentTypeIconEnumPhtray                        ComponentTypeIconEnum = "PhTray"                        //
	ComponentTypeIconEnumPhtree                        ComponentTypeIconEnum = "PhTree"                        //
	ComponentTypeIconEnumPhtreeevergreen               ComponentTypeIconEnum = "PhTreeEvergreen"               //
	ComponentTypeIconEnumPhtreestructure               ComponentTypeIconEnum = "PhTreeStructure"               //
	ComponentTypeIconEnumPhtrenddown                   ComponentTypeIconEnum = "PhTrendDown"                   //
	ComponentTypeIconEnumPhtrendup                     ComponentTypeIconEnum = "PhTrendUp"                     //
	ComponentTypeIconEnumPhtriangle                    ComponentTypeIconEnum = "PhTriangle"                    //
	ComponentTypeIconEnumPhtrophy                      ComponentTypeIconEnum = "PhTrophy"                      //
	ComponentTypeIconEnumPhtruck                       ComponentTypeIconEnum = "PhTruck"                       //
	ComponentTypeIconEnumPhtwitchlogo                  ComponentTypeIconEnum = "PhTwitchLogo"                  //
	ComponentTypeIconEnumPhtwitterlogo                 ComponentTypeIconEnum = "PhTwitterLogo"                 //
	ComponentTypeIconEnumPhumbrella                    ComponentTypeIconEnum = "PhUmbrella"                    //
	ComponentTypeIconEnumPhumbrellasimple              ComponentTypeIconEnum = "PhUmbrellaSimple"              //
	ComponentTypeIconEnumPhupload                      ComponentTypeIconEnum = "PhUpload"                      //
	ComponentTypeIconEnumPhuploadsimple                ComponentTypeIconEnum = "PhUploadSimple"                //
	ComponentTypeIconEnumPhuser                        ComponentTypeIconEnum = "PhUser"                        //
	ComponentTypeIconEnumPhusercircle                  ComponentTypeIconEnum = "PhUserCircle"                  //
	ComponentTypeIconEnumPhusercirclegear              ComponentTypeIconEnum = "PhUserCircleGear"              //
	ComponentTypeIconEnumPhusercircleminus             ComponentTypeIconEnum = "PhUserCircleMinus"             //
	ComponentTypeIconEnumPhusercircleplus              ComponentTypeIconEnum = "PhUserCirclePlus"              //
	ComponentTypeIconEnumPhuserfocus                   ComponentTypeIconEnum = "PhUserFocus"                   //
	ComponentTypeIconEnumPhusergear                    ComponentTypeIconEnum = "PhUserGear"                    //
	ComponentTypeIconEnumPhuserlist                    ComponentTypeIconEnum = "PhUserList"                    //
	ComponentTypeIconEnumPhuserminus                   ComponentTypeIconEnum = "PhUserMinus"                   //
	ComponentTypeIconEnumPhuserplus                    ComponentTypeIconEnum = "PhUserPlus"                    //
	ComponentTypeIconEnumPhuserrectangle               ComponentTypeIconEnum = "PhUserRectangle"               //
	ComponentTypeIconEnumPhusersquare                  ComponentTypeIconEnum = "PhUserSquare"                  //
	ComponentTypeIconEnumPhuserswitch                  ComponentTypeIconEnum = "PhUserSwitch"                  //
	ComponentTypeIconEnumPhusers                       ComponentTypeIconEnum = "PhUsers"                       //
	ComponentTypeIconEnumPhusersfour                   ComponentTypeIconEnum = "PhUsersFour"                   //
	ComponentTypeIconEnumPhusersthree                  ComponentTypeIconEnum = "PhUsersThree"                  //
	ComponentTypeIconEnumPhvault                       ComponentTypeIconEnum = "PhVault"                       //
	ComponentTypeIconEnumPhvibrate                     ComponentTypeIconEnum = "PhVibrate"                     //
	ComponentTypeIconEnumPhvideocamera                 ComponentTypeIconEnum = "PhVideoCamera"                 //
	ComponentTypeIconEnumPhvideocameraslash            ComponentTypeIconEnum = "PhVideoCameraSlash"            //
	ComponentTypeIconEnumPhvignette                    ComponentTypeIconEnum = "PhVignette"                    //
	ComponentTypeIconEnumPhvoicemail                   ComponentTypeIconEnum = "PhVoicemail"                   //
	ComponentTypeIconEnumPhvolleyball                  ComponentTypeIconEnum = "PhVolleyball"                  //
	ComponentTypeIconEnumPhwall                        ComponentTypeIconEnum = "PhWall"                        //
	ComponentTypeIconEnumPhwallet                      ComponentTypeIconEnum = "PhWallet"                      //
	ComponentTypeIconEnumPhwarning                     ComponentTypeIconEnum = "PhWarning"                     //
	ComponentTypeIconEnumPhwarningcircle               ComponentTypeIconEnum = "PhWarningCircle"               //
	ComponentTypeIconEnumPhwarningoctagon              ComponentTypeIconEnum = "PhWarningOctagon"              //
	ComponentTypeIconEnumPhwatch                       ComponentTypeIconEnum = "PhWatch"                       //
	ComponentTypeIconEnumPhwavesawtooth                ComponentTypeIconEnum = "PhWaveSawtooth"                //
	ComponentTypeIconEnumPhwavesine                    ComponentTypeIconEnum = "PhWaveSine"                    //
	ComponentTypeIconEnumPhwavesquare                  ComponentTypeIconEnum = "PhWaveSquare"                  //
	ComponentTypeIconEnumPhwavetriangle                ComponentTypeIconEnum = "PhWaveTriangle"                //
	ComponentTypeIconEnumPhwaves                       ComponentTypeIconEnum = "PhWaves"                       //
	ComponentTypeIconEnumPhwebcam                      ComponentTypeIconEnum = "PhWebcam"                      //
	ComponentTypeIconEnumPhwhatsapplogo                ComponentTypeIconEnum = "PhWhatsappLogo"                //
	ComponentTypeIconEnumPhwheelchair                  ComponentTypeIconEnum = "PhWheelchair"                  //
	ComponentTypeIconEnumPhwifihigh                    ComponentTypeIconEnum = "PhWifiHigh"                    //
	ComponentTypeIconEnumPhwifilow                     ComponentTypeIconEnum = "PhWifiLow"                     //
	ComponentTypeIconEnumPhwifimedium                  ComponentTypeIconEnum = "PhWifiMedium"                  //
	ComponentTypeIconEnumPhwifinone                    ComponentTypeIconEnum = "PhWifiNone"                    //
	ComponentTypeIconEnumPhwifislash                   ComponentTypeIconEnum = "PhWifiSlash"                   //
	ComponentTypeIconEnumPhwifix                       ComponentTypeIconEnum = "PhWifiX"                       //
	ComponentTypeIconEnumPhwind                        ComponentTypeIconEnum = "PhWind"                        //
	ComponentTypeIconEnumPhwindowslogo                 ComponentTypeIconEnum = "PhWindowsLogo"                 //
	ComponentTypeIconEnumPhwine                        ComponentTypeIconEnum = "PhWine"                        //
	ComponentTypeIconEnumPhwrench                      ComponentTypeIconEnum = "PhWrench"                      //
	ComponentTypeIconEnumPhx                           ComponentTypeIconEnum = "PhX"                           //
	ComponentTypeIconEnumPhxcircle                     ComponentTypeIconEnum = "PhXCircle"                     //
	ComponentTypeIconEnumPhxsquare                     ComponentTypeIconEnum = "PhXSquare"                     //
	ComponentTypeIconEnumPhyinyang                     ComponentTypeIconEnum = "PhYinYang"                     //
	ComponentTypeIconEnumPhyoutubelogo                 ComponentTypeIconEnum = "PhYoutubeLogo"                 //
)

// All ComponentTypeIconEnum as []string
var AllComponentTypeIconEnum = []string{
	string(ComponentTypeIconEnumPhactivity),
	string(ComponentTypeIconEnumPhaddressbook),
	string(ComponentTypeIconEnumPhairplane),
	string(ComponentTypeIconEnumPhairplaneinflight),
	string(ComponentTypeIconEnumPhairplanelanding),
	string(ComponentTypeIconEnumPhairplanetakeoff),
	string(ComponentTypeIconEnumPhairplanetilt),
	string(ComponentTypeIconEnumPhairplay),
	string(ComponentTypeIconEnumPhalarm),
	string(ComponentTypeIconEnumPhalien),
	string(ComponentTypeIconEnumPhalignbottom),
	string(ComponentTypeIconEnumPhalignbottomsimple),
	string(ComponentTypeIconEnumPhaligncenterhorizontal),
	string(ComponentTypeIconEnumPhaligncenterhorizontalsimple),
	string(ComponentTypeIconEnumPhaligncentervertical),
	string(ComponentTypeIconEnumPhaligncenterverticalsimple),
	string(ComponentTypeIconEnumPhalignleft),
	string(ComponentTypeIconEnumPhalignleftsimple),
	string(ComponentTypeIconEnumPhalignright),
	string(ComponentTypeIconEnumPhalignrightsimple),
	string(ComponentTypeIconEnumPhaligntop),
	string(ComponentTypeIconEnumPhaligntopsimple),
	string(ComponentTypeIconEnumPhanchor),
	string(ComponentTypeIconEnumPhanchorsimple),
	string(ComponentTypeIconEnumPhandroidlogo),
	string(ComponentTypeIconEnumPhangularlogo),
	string(ComponentTypeIconEnumPhaperture),
	string(ComponentTypeIconEnumPhappstorelogo),
	string(ComponentTypeIconEnumPhappwindow),
	string(ComponentTypeIconEnumPhapplelogo),
	string(ComponentTypeIconEnumPhapplepodcastslogo),
	string(ComponentTypeIconEnumPharchive),
	string(ComponentTypeIconEnumPharchivebox),
	string(ComponentTypeIconEnumPharchivetray),
	string(ComponentTypeIconEnumPharmchair),
	string(ComponentTypeIconEnumPharrowarcleft),
	string(ComponentTypeIconEnumPharrowarcright),
	string(ComponentTypeIconEnumPharrowbenddoubleupleft),
	string(ComponentTypeIconEnumPharrowbenddoubleupright),
	string(ComponentTypeIconEnumPharrowbenddownleft),
	string(ComponentTypeIconEnumPharrowbenddownright),
	string(ComponentTypeIconEnumPharrowbendleftdown),
	string(ComponentTypeIconEnumPharrowbendleftup),
	string(ComponentTypeIconEnumPharrowbendrightdown),
	string(ComponentTypeIconEnumPharrowbendrightup),
	string(ComponentTypeIconEnumPharrowbendupleft),
	string(ComponentTypeIconEnumPharrowbendupright),
	string(ComponentTypeIconEnumPharrowcircledown),
	string(ComponentTypeIconEnumPharrowcircledownleft),
	string(ComponentTypeIconEnumPharrowcircledownright),
	string(ComponentTypeIconEnumPharrowcircleleft),
	string(ComponentTypeIconEnumPharrowcircleright),
	string(ComponentTypeIconEnumPharrowcircleup),
	string(ComponentTypeIconEnumPharrowcircleupleft),
	string(ComponentTypeIconEnumPharrowcircleupright),
	string(ComponentTypeIconEnumPharrowclockwise),
	string(ComponentTypeIconEnumPharrowcounterclockwise),
	string(ComponentTypeIconEnumPharrowdown),
	string(ComponentTypeIconEnumPharrowdownleft),
	string(ComponentTypeIconEnumPharrowdownright),
	string(ComponentTypeIconEnumPharrowelbowdownleft),
	string(ComponentTypeIconEnumPharrowelbowdownright),
	string(ComponentTypeIconEnumPharrowelbowleft),
	string(ComponentTypeIconEnumPharrowelbowleftdown),
	string(ComponentTypeIconEnumPharrowelbowleftup),
	string(ComponentTypeIconEnumPharrowelbowright),
	string(ComponentTypeIconEnumPharrowelbowrightdown),
	string(ComponentTypeIconEnumPharrowelbowrightup),
	string(ComponentTypeIconEnumPharrowelbowupleft),
	string(ComponentTypeIconEnumPharrowelbowupright),
	string(ComponentTypeIconEnumPharrowfatdown),
	string(ComponentTypeIconEnumPharrowfatleft),
	string(ComponentTypeIconEnumPharrowfatlinedown),
	string(ComponentTypeIconEnumPharrowfatlineleft),
	string(ComponentTypeIconEnumPharrowfatlineright),
	string(ComponentTypeIconEnumPharrowfatlineup),
	string(ComponentTypeIconEnumPharrowfatlinesdown),
	string(ComponentTypeIconEnumPharrowfatlinesleft),
	string(ComponentTypeIconEnumPharrowfatlinesright),
	string(ComponentTypeIconEnumPharrowfatlinesup),
	string(ComponentTypeIconEnumPharrowfatright),
	string(ComponentTypeIconEnumPharrowfatup),
	string(ComponentTypeIconEnumPharrowleft),
	string(ComponentTypeIconEnumPharrowlinedown),
	string(ComponentTypeIconEnumPharrowlinedownleft),
	string(ComponentTypeIconEnumPharrowlinedownright),
	string(ComponentTypeIconEnumPharrowlineleft),
	string(ComponentTypeIconEnumPharrowlineright),
	string(ComponentTypeIconEnumPharrowlineup),
	string(ComponentTypeIconEnumPharrowlineupleft),
	string(ComponentTypeIconEnumPharrowlineupright),
	string(ComponentTypeIconEnumPharrowright),
	string(ComponentTypeIconEnumPharrowsquaredown),
	string(ComponentTypeIconEnumPharrowsquaredownleft),
	string(ComponentTypeIconEnumPharrowsquaredownright),
	string(ComponentTypeIconEnumPharrowsquarein),
	string(ComponentTypeIconEnumPharrowsquareleft),
	string(ComponentTypeIconEnumPharrowsquareout),
	string(ComponentTypeIconEnumPharrowsquareright),
	string(ComponentTypeIconEnumPharrowsquareup),
	string(ComponentTypeIconEnumPharrowsquareupleft),
	string(ComponentTypeIconEnumPharrowsquareupright),
	string(ComponentTypeIconEnumPharrowudownleft),
	string(ComponentTypeIconEnumPharrowudownright),
	string(ComponentTypeIconEnumPharrowuleftdown),
	string(ComponentTypeIconEnumPharrowuleftup),
	string(ComponentTypeIconEnumPharrowurightdown),
	string(ComponentTypeIconEnumPharrowurightup),
	string(ComponentTypeIconEnumPharrowuupleft),
	string(ComponentTypeIconEnumPharrowuupright),
	string(ComponentTypeIconEnumPharrowup),
	string(ComponentTypeIconEnumPharrowupleft),
	string(ComponentTypeIconEnumPharrowupright),
	string(ComponentTypeIconEnumPharrowsclockwise),
	string(ComponentTypeIconEnumPharrowscounterclockwise),
	string(ComponentTypeIconEnumPharrowsdownup),
	string(ComponentTypeIconEnumPharrowshorizontal),
	string(ComponentTypeIconEnumPharrowsin),
	string(ComponentTypeIconEnumPharrowsincardinal),
	string(ComponentTypeIconEnumPharrowsinlinehorizontal),
	string(ComponentTypeIconEnumPharrowsinlinevertical),
	string(ComponentTypeIconEnumPharrowsinsimple),
	string(ComponentTypeIconEnumPharrowsleftright),
	string(ComponentTypeIconEnumPharrowsout),
	string(ComponentTypeIconEnumPharrowsoutcardinal),
	string(ComponentTypeIconEnumPharrowsoutlinehorizontal),
	string(ComponentTypeIconEnumPharrowsoutlinevertical),
	string(ComponentTypeIconEnumPharrowsoutsimple),
	string(ComponentTypeIconEnumPharrowsvertical),
	string(ComponentTypeIconEnumPharticle),
	string(ComponentTypeIconEnumPharticlemedium),
	string(ComponentTypeIconEnumPharticlenytimes),
	string(ComponentTypeIconEnumPhasterisk),
	string(ComponentTypeIconEnumPhasterisksimple),
	string(ComponentTypeIconEnumPhat),
	string(ComponentTypeIconEnumPhatom),
	string(ComponentTypeIconEnumPhbaby),
	string(ComponentTypeIconEnumPhbackpack),
	string(ComponentTypeIconEnumPhbackspace),
	string(ComponentTypeIconEnumPhbag),
	string(ComponentTypeIconEnumPhbagsimple),
	string(ComponentTypeIconEnumPhballoon),
	string(ComponentTypeIconEnumPhbandaids),
	string(ComponentTypeIconEnumPhbank),
	string(ComponentTypeIconEnumPhbarbell),
	string(ComponentTypeIconEnumPhbarcode),
	string(ComponentTypeIconEnumPhbarricade),
	string(ComponentTypeIconEnumPhbaseball),
	string(ComponentTypeIconEnumPhbasketball),
	string(ComponentTypeIconEnumPhbathtub),
	string(ComponentTypeIconEnumPhbatterycharging),
	string(ComponentTypeIconEnumPhbatterychargingvertical),
	string(ComponentTypeIconEnumPhbatteryempty),
	string(ComponentTypeIconEnumPhbatteryfull),
	string(ComponentTypeIconEnumPhbatteryhigh),
	string(ComponentTypeIconEnumPhbatterylow),
	string(ComponentTypeIconEnumPhbatterymedium),
	string(ComponentTypeIconEnumPhbatteryplus),
	string(ComponentTypeIconEnumPhbatterywarning),
	string(ComponentTypeIconEnumPhbatterywarningvertical),
	string(ComponentTypeIconEnumPhbed),
	string(ComponentTypeIconEnumPhbeerbottle),
	string(ComponentTypeIconEnumPhbehancelogo),
	string(ComponentTypeIconEnumPhbell),
	string(ComponentTypeIconEnumPhbellringing),
	string(ComponentTypeIconEnumPhbellsimple),
	string(ComponentTypeIconEnumPhbellsimpleringing),
	string(ComponentTypeIconEnumPhbellsimpleslash),
	string(ComponentTypeIconEnumPhbellsimplez),
	string(ComponentTypeIconEnumPhbellslash),
	string(ComponentTypeIconEnumPhbellz),
	string(ComponentTypeIconEnumPhbeziercurve),
	string(ComponentTypeIconEnumPhbicycle),
	string(ComponentTypeIconEnumPhbinoculars),
	string(ComponentTypeIconEnumPhbird),
	string(ComponentTypeIconEnumPhbluetooth),
	string(ComponentTypeIconEnumPhbluetoothconnected),
	string(ComponentTypeIconEnumPhbluetoothslash),
	string(ComponentTypeIconEnumPhbluetoothx),
	string(ComponentTypeIconEnumPhboat),
	string(ComponentTypeIconEnumPhbook),
	string(ComponentTypeIconEnumPhbookbookmark),
	string(ComponentTypeIconEnumPhbookopen),
	string(ComponentTypeIconEnumPhbookmark),
	string(ComponentTypeIconEnumPhbookmarksimple),
	string(ComponentTypeIconEnumPhbookmarks),
	string(ComponentTypeIconEnumPhbookmarkssimple),
	string(ComponentTypeIconEnumPhbooks),
	string(ComponentTypeIconEnumPhboundingbox),
	string(ComponentTypeIconEnumPhbracketsangle),
	string(ComponentTypeIconEnumPhbracketscurly),
	string(ComponentTypeIconEnumPhbracketsround),
	string(ComponentTypeIconEnumPhbracketssquare),
	string(ComponentTypeIconEnumPhbrain),
	string(ComponentTypeIconEnumPhbrandy),
	string(ComponentTypeIconEnumPhbriefcase),
	string(ComponentTypeIconEnumPhbriefcasemetal),
	string(ComponentTypeIconEnumPhbroadcast),
	string(ComponentTypeIconEnumPhbrowser),
	string(ComponentTypeIconEnumPhbrowsers),
	string(ComponentTypeIconEnumPhbug),
	string(ComponentTypeIconEnumPhbugbeetle),
	string(ComponentTypeIconEnumPhbugdroid),
	string(ComponentTypeIconEnumPhbuildings),
	string(ComponentTypeIconEnumPhbus),
	string(ComponentTypeIconEnumPhbutterfly),
	string(ComponentTypeIconEnumPhcactus),
	string(ComponentTypeIconEnumPhcake),
	string(ComponentTypeIconEnumPhcalculator),
	string(ComponentTypeIconEnumPhcalendar),
	string(ComponentTypeIconEnumPhcalendarblank),
	string(ComponentTypeIconEnumPhcalendarcheck),
	string(ComponentTypeIconEnumPhcalendarplus),
	string(ComponentTypeIconEnumPhcalendarx),
	string(ComponentTypeIconEnumPhcamera),
	string(ComponentTypeIconEnumPhcamerarotate),
	string(ComponentTypeIconEnumPhcameraslash),
	string(ComponentTypeIconEnumPhcampfire),
	string(ComponentTypeIconEnumPhcar),
	string(ComponentTypeIconEnumPhcarsimple),
	string(ComponentTypeIconEnumPhcardholder),
	string(ComponentTypeIconEnumPhcards),
	string(ComponentTypeIconEnumPhcaretcircledoubledown),
	string(ComponentTypeIconEnumPhcaretcircledoubleleft),
	string(ComponentTypeIconEnumPhcaretcircledoubleright),
	string(ComponentTypeIconEnumPhcaretcircledoubleup),
	string(ComponentTypeIconEnumPhcaretcircledown),
	string(ComponentTypeIconEnumPhcaretcircleleft),
	string(ComponentTypeIconEnumPhcaretcircleright),
	string(ComponentTypeIconEnumPhcaretcircleup),
	string(ComponentTypeIconEnumPhcaretdoubledown),
	string(ComponentTypeIconEnumPhcaretdoubleleft),
	string(ComponentTypeIconEnumPhcaretdoubleright),
	string(ComponentTypeIconEnumPhcaretdoubleup),
	string(ComponentTypeIconEnumPhcaretdown),
	string(ComponentTypeIconEnumPhcaretleft),
	string(ComponentTypeIconEnumPhcaretright),
	string(ComponentTypeIconEnumPhcaretup),
	string(ComponentTypeIconEnumPhcat),
	string(ComponentTypeIconEnumPhcellsignalfull),
	string(ComponentTypeIconEnumPhcellsignalhigh),
	string(ComponentTypeIconEnumPhcellsignallow),
	string(ComponentTypeIconEnumPhcellsignalmedium),
	string(ComponentTypeIconEnumPhcellsignalnone),
	string(ComponentTypeIconEnumPhcellsignalslash),
	string(ComponentTypeIconEnumPhcellsignalx),
	string(ComponentTypeIconEnumPhchalkboard),
	string(ComponentTypeIconEnumPhchalkboardsimple),
	string(ComponentTypeIconEnumPhchalkboardteacher),
	string(ComponentTypeIconEnumPhchartbar),
	string(ComponentTypeIconEnumPhchartbarhorizontal),
	string(ComponentTypeIconEnumPhchartline),
	string(ComponentTypeIconEnumPhchartlineup),
	string(ComponentTypeIconEnumPhchartpie),
	string(ComponentTypeIconEnumPhchartpieslice),
	string(ComponentTypeIconEnumPhchat),
	string(ComponentTypeIconEnumPhchatcentered),
	string(ComponentTypeIconEnumPhchatcentereddots),
	string(ComponentTypeIconEnumPhchatcenteredtext),
	string(ComponentTypeIconEnumPhchatcircle),
	string(ComponentTypeIconEnumPhchatcircledots),
	string(ComponentTypeIconEnumPhchatcircletext),
	string(ComponentTypeIconEnumPhchatdots),
	string(ComponentTypeIconEnumPhchatteardrop),
	string(ComponentTypeIconEnumPhchatteardropdots),
	string(ComponentTypeIconEnumPhchatteardroptext),
	string(ComponentTypeIconEnumPhchattext),
	string(ComponentTypeIconEnumPhchats),
	string(ComponentTypeIconEnumPhchatscircle),
	string(ComponentTypeIconEnumPhchatsteardrop),
	string(ComponentTypeIconEnumPhcheck),
	string(ComponentTypeIconEnumPhcheckcircle),
	string(ComponentTypeIconEnumPhchecksquare),
	string(ComponentTypeIconEnumPhchecksquareoffset),
	string(ComponentTypeIconEnumPhchecks),
	string(ComponentTypeIconEnumPhcircle),
	string(ComponentTypeIconEnumPhcircledashed),
	string(ComponentTypeIconEnumPhcirclehalf),
	string(ComponentTypeIconEnumPhcirclehalftilt),
	string(ComponentTypeIconEnumPhcirclenotch),
	string(ComponentTypeIconEnumPhcirclewavy),
	string(ComponentTypeIconEnumPhcirclewavycheck),
	string(ComponentTypeIconEnumPhcirclewavyquestion),
	string(ComponentTypeIconEnumPhcirclewavywarning),
	string(ComponentTypeIconEnumPhcirclesfour),
	string(ComponentTypeIconEnumPhcirclesthree),
	string(ComponentTypeIconEnumPhcirclesthreeplus),
	string(ComponentTypeIconEnumPhclipboard),
	string(ComponentTypeIconEnumPhclipboardtext),
	string(ComponentTypeIconEnumPhclock),
	string(ComponentTypeIconEnumPhclockafternoon),
	string(ComponentTypeIconEnumPhclockclockwise),
	string(ComponentTypeIconEnumPhclockcounterclockwise),
	string(ComponentTypeIconEnumPhclosedcaptioning),
	string(ComponentTypeIconEnumPhcloud),
	string(ComponentTypeIconEnumPhcloudarrowdown),
	string(ComponentTypeIconEnumPhcloudarrowup),
	string(ComponentTypeIconEnumPhcloudcheck),
	string(ComponentTypeIconEnumPhcloudfog),
	string(ComponentTypeIconEnumPhcloudlightning),
	string(ComponentTypeIconEnumPhcloudmoon),
	string(ComponentTypeIconEnumPhcloudrain),
	string(ComponentTypeIconEnumPhcloudslash),
	string(ComponentTypeIconEnumPhcloudsnow),
	string(ComponentTypeIconEnumPhcloudsun),
	string(ComponentTypeIconEnumPhclub),
	string(ComponentTypeIconEnumPhcoathanger),
	string(ComponentTypeIconEnumPhcode),
	string(ComponentTypeIconEnumPhcodesimple),
	string(ComponentTypeIconEnumPhcodepenlogo),
	string(ComponentTypeIconEnumPhcodesandboxlogo),
	string(ComponentTypeIconEnumPhcoffee),
	string(ComponentTypeIconEnumPhcoin),
	string(ComponentTypeIconEnumPhcoinvertical),
	string(ComponentTypeIconEnumPhcoins),
	string(ComponentTypeIconEnumPhcolumns),
	string(ComponentTypeIconEnumPhcommand),
	string(ComponentTypeIconEnumPhcompass),
	string(ComponentTypeIconEnumPhcomputertower),
	string(ComponentTypeIconEnumPhconfetti),
	string(ComponentTypeIconEnumPhcookie),
	string(ComponentTypeIconEnumPhcookingpot),
	string(ComponentTypeIconEnumPhcopy),
	string(ComponentTypeIconEnumPhcopysimple),
	string(ComponentTypeIconEnumPhcopyleft),
	string(ComponentTypeIconEnumPhcopyright),
	string(ComponentTypeIconEnumPhcornersin),
	string(ComponentTypeIconEnumPhcornersout),
	string(ComponentTypeIconEnumPhcpu),
	string(ComponentTypeIconEnumPhcreditcard),
	string(ComponentTypeIconEnumPhcrop),
	string(ComponentTypeIconEnumPhcrosshair),
	string(ComponentTypeIconEnumPhcrosshairsimple),
	string(ComponentTypeIconEnumPhcrown),
	string(ComponentTypeIconEnumPhcrownsimple),
	string(ComponentTypeIconEnumPhcube),
	string(ComponentTypeIconEnumPhcurrencybtc),
	string(ComponentTypeIconEnumPhcurrencycircledollar),
	string(ComponentTypeIconEnumPhcurrencycny),
	string(ComponentTypeIconEnumPhcurrencydollar),
	string(ComponentTypeIconEnumPhcurrencydollarsimple),
	string(ComponentTypeIconEnumPhcurrencyeth),
	string(ComponentTypeIconEnumPhcurrencyeur),
	string(ComponentTypeIconEnumPhcurrencygbp),
	string(ComponentTypeIconEnumPhcurrencyinr),
	string(ComponentTypeIconEnumPhcurrencyjpy),
	string(ComponentTypeIconEnumPhcurrencykrw),
	string(ComponentTypeIconEnumPhcurrencykzt),
	string(ComponentTypeIconEnumPhcurrencyngn),
	string(ComponentTypeIconEnumPhcurrencyrub),
	string(ComponentTypeIconEnumPhcursor),
	string(ComponentTypeIconEnumPhcursortext),
	string(ComponentTypeIconEnumPhcylinder),
	string(ComponentTypeIconEnumPhdatabase),
	string(ComponentTypeIconEnumPhdesktop),
	string(ComponentTypeIconEnumPhdesktoptower),
	string(ComponentTypeIconEnumPhdetective),
	string(ComponentTypeIconEnumPhdevicemobile),
	string(ComponentTypeIconEnumPhdevicemobilecamera),
	string(ComponentTypeIconEnumPhdevicemobilespeaker),
	string(ComponentTypeIconEnumPhdevicetablet),
	string(ComponentTypeIconEnumPhdevicetabletcamera),
	string(ComponentTypeIconEnumPhdevicetabletspeaker),
	string(ComponentTypeIconEnumPhdiamond),
	string(ComponentTypeIconEnumPhdiamondsfour),
	string(ComponentTypeIconEnumPhdicefive),
	string(ComponentTypeIconEnumPhdicefour),
	string(ComponentTypeIconEnumPhdiceone),
	string(ComponentTypeIconEnumPhdicesix),
	string(ComponentTypeIconEnumPhdicethree),
	string(ComponentTypeIconEnumPhdicetwo),
	string(ComponentTypeIconEnumPhdisc),
	string(ComponentTypeIconEnumPhdiscordlogo),
	string(ComponentTypeIconEnumPhdivide),
	string(ComponentTypeIconEnumPhdog),
	string(ComponentTypeIconEnumPhdoor),
	string(ComponentTypeIconEnumPhdotsnine),
	string(ComponentTypeIconEnumPhdotssix),
	string(ComponentTypeIconEnumPhdotssixvertical),
	string(ComponentTypeIconEnumPhdotsthree),
	string(ComponentTypeIconEnumPhdotsthreecircle),
	string(ComponentTypeIconEnumPhdotsthreecirclevertical),
	string(ComponentTypeIconEnumPhdotsthreeoutline),
	string(ComponentTypeIconEnumPhdotsthreeoutlinevertical),
	string(ComponentTypeIconEnumPhdotsthreevertical),
	string(ComponentTypeIconEnumPhdownload),
	string(ComponentTypeIconEnumPhdownloadsimple),
	string(ComponentTypeIconEnumPhdribbblelogo),
	string(ComponentTypeIconEnumPhdrop),
	string(ComponentTypeIconEnumPhdrophalf),
	string(ComponentTypeIconEnumPhdrophalfbottom),
	string(ComponentTypeIconEnumPhear),
	string(ComponentTypeIconEnumPhearslash),
	string(ComponentTypeIconEnumPhegg),
	string(ComponentTypeIconEnumPheggcrack),
	string(ComponentTypeIconEnumPheject),
	string(ComponentTypeIconEnumPhejectsimple),
	string(ComponentTypeIconEnumPhenvelope),
	string(ComponentTypeIconEnumPhenvelopeopen),
	string(ComponentTypeIconEnumPhenvelopesimple),
	string(ComponentTypeIconEnumPhenvelopesimpleopen),
	string(ComponentTypeIconEnumPhequalizer),
	string(ComponentTypeIconEnumPhequals),
	string(ComponentTypeIconEnumPheraser),
	string(ComponentTypeIconEnumPhexam),
	string(ComponentTypeIconEnumPhexport),
	string(ComponentTypeIconEnumPheye),
	string(ComponentTypeIconEnumPheyeclosed),
	string(ComponentTypeIconEnumPheyeslash),
	string(ComponentTypeIconEnumPheyedropper),
	string(ComponentTypeIconEnumPheyedroppersample),
	string(ComponentTypeIconEnumPheyeglasses),
	string(ComponentTypeIconEnumPhfacemask),
	string(ComponentTypeIconEnumPhfacebooklogo),
	string(ComponentTypeIconEnumPhfactory),
	string(ComponentTypeIconEnumPhfaders),
	string(ComponentTypeIconEnumPhfadershorizontal),
	string(ComponentTypeIconEnumPhfastforward),
	string(ComponentTypeIconEnumPhfastforwardcircle),
	string(ComponentTypeIconEnumPhfigmalogo),
	string(ComponentTypeIconEnumPhfile),
	string(ComponentTypeIconEnumPhfilearrowdown),
	string(ComponentTypeIconEnumPhfilearrowup),
	string(ComponentTypeIconEnumPhfileaudio),
	string(ComponentTypeIconEnumPhfilecloud),
	string(ComponentTypeIconEnumPhfilecode),
	string(ComponentTypeIconEnumPhfilecss),
	string(ComponentTypeIconEnumPhfilecsv),
	string(ComponentTypeIconEnumPhfiledoc),
	string(ComponentTypeIconEnumPhfiledotted),
	string(ComponentTypeIconEnumPhfilehtml),
	string(ComponentTypeIconEnumPhfileimage),
	string(ComponentTypeIconEnumPhfilejpg),
	string(ComponentTypeIconEnumPhfilejs),
	string(ComponentTypeIconEnumPhfilejsx),
	string(ComponentTypeIconEnumPhfilelock),
	string(ComponentTypeIconEnumPhfileminus),
	string(ComponentTypeIconEnumPhfilepdf),
	string(ComponentTypeIconEnumPhfileplus),
	string(ComponentTypeIconEnumPhfilepng),
	string(ComponentTypeIconEnumPhfileppt),
	string(ComponentTypeIconEnumPhfilers),
	string(ComponentTypeIconEnumPhfilesearch),
	string(ComponentTypeIconEnumPhfiletext),
	string(ComponentTypeIconEnumPhfilets),
	string(ComponentTypeIconEnumPhfiletsx),
	string(ComponentTypeIconEnumPhfilevideo),
	string(ComponentTypeIconEnumPhfilevue),
	string(ComponentTypeIconEnumPhfilex),
	string(ComponentTypeIconEnumPhfilexls),
	string(ComponentTypeIconEnumPhfilezip),
	string(ComponentTypeIconEnumPhfiles),
	string(ComponentTypeIconEnumPhfilmscript),
	string(ComponentTypeIconEnumPhfilmslate),
	string(ComponentTypeIconEnumPhfilmstrip),
	string(ComponentTypeIconEnumPhfingerprint),
	string(ComponentTypeIconEnumPhfingerprintsimple),
	string(ComponentTypeIconEnumPhfinnthehuman),
	string(ComponentTypeIconEnumPhfire),
	string(ComponentTypeIconEnumPhfiresimple),
	string(ComponentTypeIconEnumPhfirstaid),
	string(ComponentTypeIconEnumPhfirstaidkit),
	string(ComponentTypeIconEnumPhfish),
	string(ComponentTypeIconEnumPhfishsimple),
	string(ComponentTypeIconEnumPhflag),
	string(ComponentTypeIconEnumPhflagbanner),
	string(ComponentTypeIconEnumPhflagcheckered),
	string(ComponentTypeIconEnumPhflame),
	string(ComponentTypeIconEnumPhflashlight),
	string(ComponentTypeIconEnumPhflask),
	string(ComponentTypeIconEnumPhfloppydisk),
	string(ComponentTypeIconEnumPhfloppydiskback),
	string(ComponentTypeIconEnumPhflowarrow),
	string(ComponentTypeIconEnumPhflower),
	string(ComponentTypeIconEnumPhflowerlotus),
	string(ComponentTypeIconEnumPhflyingsaucer),
	string(ComponentTypeIconEnumPhfolder),
	string(ComponentTypeIconEnumPhfolderdotted),
	string(ComponentTypeIconEnumPhfolderlock),
	string(ComponentTypeIconEnumPhfolderminus),
	string(ComponentTypeIconEnumPhfoldernotch),
	string(ComponentTypeIconEnumPhfoldernotchminus),
	string(ComponentTypeIconEnumPhfoldernotchopen),
	string(ComponentTypeIconEnumPhfoldernotchplus),
	string(ComponentTypeIconEnumPhfolderopen),
	string(ComponentTypeIconEnumPhfolderplus),
	string(ComponentTypeIconEnumPhfoldersimple),
	string(ComponentTypeIconEnumPhfoldersimpledotted),
	string(ComponentTypeIconEnumPhfoldersimplelock),
	string(ComponentTypeIconEnumPhfoldersimpleminus),
	string(ComponentTypeIconEnumPhfoldersimpleplus),
	string(ComponentTypeIconEnumPhfoldersimplestar),
	string(ComponentTypeIconEnumPhfoldersimpleuser),
	string(ComponentTypeIconEnumPhfolderstar),
	string(ComponentTypeIconEnumPhfolderuser),
	string(ComponentTypeIconEnumPhfolders),
	string(ComponentTypeIconEnumPhfootball),
	string(ComponentTypeIconEnumPhforkknife),
	string(ComponentTypeIconEnumPhframecorners),
	string(ComponentTypeIconEnumPhframerlogo),
	string(ComponentTypeIconEnumPhfunction),
	string(ComponentTypeIconEnumPhfunnel),
	string(ComponentTypeIconEnumPhfunnelsimple),
	string(ComponentTypeIconEnumPhgamecontroller),
	string(ComponentTypeIconEnumPhgaspump),
	string(ComponentTypeIconEnumPhgauge),
	string(ComponentTypeIconEnumPhgear),
	string(ComponentTypeIconEnumPhgearsix),
	string(ComponentTypeIconEnumPhgenderfemale),
	string(ComponentTypeIconEnumPhgenderintersex),
	string(ComponentTypeIconEnumPhgendermale),
	string(ComponentTypeIconEnumPhgenderneuter),
	string(ComponentTypeIconEnumPhgendernonbinary),
	string(ComponentTypeIconEnumPhgendertransgender),
	string(ComponentTypeIconEnumPhghost),
	string(ComponentTypeIconEnumPhgif),
	string(ComponentTypeIconEnumPhgift),
	string(ComponentTypeIconEnumPhgitbranch),
	string(ComponentTypeIconEnumPhgitcommit),
	string(ComponentTypeIconEnumPhgitdiff),
	string(ComponentTypeIconEnumPhgitfork),
	string(ComponentTypeIconEnumPhgitmerge),
	string(ComponentTypeIconEnumPhgitpullrequest),
	string(ComponentTypeIconEnumPhgithublogo),
	string(ComponentTypeIconEnumPhgitlablogo),
	string(ComponentTypeIconEnumPhgitlablogosimple),
	string(ComponentTypeIconEnumPhglobe),
	string(ComponentTypeIconEnumPhglobehemisphereeast),
	string(ComponentTypeIconEnumPhglobehemispherewest),
	string(ComponentTypeIconEnumPhglobesimple),
	string(ComponentTypeIconEnumPhglobestand),
	string(ComponentTypeIconEnumPhgooglechromelogo),
	string(ComponentTypeIconEnumPhgooglelogo),
	string(ComponentTypeIconEnumPhgooglephotoslogo),
	string(ComponentTypeIconEnumPhgoogleplaylogo),
	string(ComponentTypeIconEnumPhgooglepodcastslogo),
	string(ComponentTypeIconEnumPhgradient),
	string(ComponentTypeIconEnumPhgraduationcap),
	string(ComponentTypeIconEnumPhgraph),
	string(ComponentTypeIconEnumPhgridfour),
	string(ComponentTypeIconEnumPhhamburger),
	string(ComponentTypeIconEnumPhhand),
	string(ComponentTypeIconEnumPhhandeye),
	string(ComponentTypeIconEnumPhhandfist),
	string(ComponentTypeIconEnumPhhandgrabbing),
	string(ComponentTypeIconEnumPhhandpalm),
	string(ComponentTypeIconEnumPhhandpointing),
	string(ComponentTypeIconEnumPhhandsoap),
	string(ComponentTypeIconEnumPhhandwaving),
	string(ComponentTypeIconEnumPhhandbag),
	string(ComponentTypeIconEnumPhhandbagsimple),
	string(ComponentTypeIconEnumPhhandsclapping),
	string(ComponentTypeIconEnumPhhandshake),
	string(ComponentTypeIconEnumPhharddrive),
	string(ComponentTypeIconEnumPhharddrives),
	string(ComponentTypeIconEnumPhhash),
	string(ComponentTypeIconEnumPhhashstraight),
	string(ComponentTypeIconEnumPhheadlights),
	string(ComponentTypeIconEnumPhheadphones),
	string(ComponentTypeIconEnumPhheadset),
	string(ComponentTypeIconEnumPhheart),
	string(ComponentTypeIconEnumPhheartbreak),
	string(ComponentTypeIconEnumPhheartstraight),
	string(ComponentTypeIconEnumPhheartstraightbreak),
	string(ComponentTypeIconEnumPhheartbeat),
	string(ComponentTypeIconEnumPhhexagon),
	string(ComponentTypeIconEnumPhhighlightercircle),
	string(ComponentTypeIconEnumPhhorse),
	string(ComponentTypeIconEnumPhhourglass),
	string(ComponentTypeIconEnumPhhourglasshigh),
	string(ComponentTypeIconEnumPhhourglasslow),
	string(ComponentTypeIconEnumPhhourglassmedium),
	string(ComponentTypeIconEnumPhhourglasssimple),
	string(ComponentTypeIconEnumPhhourglasssimplehigh),
	string(ComponentTypeIconEnumPhhourglasssimplelow),
	string(ComponentTypeIconEnumPhhourglasssimplemedium),
	string(ComponentTypeIconEnumPhhouse),
	string(ComponentTypeIconEnumPhhouseline),
	string(ComponentTypeIconEnumPhhousesimple),
	string(ComponentTypeIconEnumPhidentificationbadge),
	string(ComponentTypeIconEnumPhidentificationcard),
	string(ComponentTypeIconEnumPhimage),
	string(ComponentTypeIconEnumPhimagesquare),
	string(ComponentTypeIconEnumPhinfinity),
	string(ComponentTypeIconEnumPhinfo),
	string(ComponentTypeIconEnumPhinstagramlogo),
	string(ComponentTypeIconEnumPhintersect),
	string(ComponentTypeIconEnumPhjeep),
	string(ComponentTypeIconEnumPhkanban),
	string(ComponentTypeIconEnumPhkey),
	string(ComponentTypeIconEnumPhkeyreturn),
	string(ComponentTypeIconEnumPhkeyboard),
	string(ComponentTypeIconEnumPhkeyhole),
	string(ComponentTypeIconEnumPhknife),
	string(ComponentTypeIconEnumPhladder),
	string(ComponentTypeIconEnumPhladdersimple),
	string(ComponentTypeIconEnumPhlamp),
	string(ComponentTypeIconEnumPhlaptop),
	string(ComponentTypeIconEnumPhlayout),
	string(ComponentTypeIconEnumPhleaf),
	string(ComponentTypeIconEnumPhlifebuoy),
	string(ComponentTypeIconEnumPhlightbulb),
	string(ComponentTypeIconEnumPhlightbulbfilament),
	string(ComponentTypeIconEnumPhlightning),
	string(ComponentTypeIconEnumPhlightningslash),
	string(ComponentTypeIconEnumPhlinesegment),
	string(ComponentTypeIconEnumPhlinesegments),
	string(ComponentTypeIconEnumPhlink),
	string(ComponentTypeIconEnumPhlinkbreak),
	string(ComponentTypeIconEnumPhlinksimple),
	string(ComponentTypeIconEnumPhlinksimplebreak),
	string(ComponentTypeIconEnumPhlinksimplehorizontal),
	string(ComponentTypeIconEnumPhlinksimplehorizontalbreak),
	string(ComponentTypeIconEnumPhlinkedinlogo),
	string(ComponentTypeIconEnumPhlinuxlogo),
	string(ComponentTypeIconEnumPhlist),
	string(ComponentTypeIconEnumPhlistbullets),
	string(ComponentTypeIconEnumPhlistchecks),
	string(ComponentTypeIconEnumPhlistdashes),
	string(ComponentTypeIconEnumPhlistnumbers),
	string(ComponentTypeIconEnumPhlistplus),
	string(ComponentTypeIconEnumPhlock),
	string(ComponentTypeIconEnumPhlockkey),
	string(ComponentTypeIconEnumPhlockkeyopen),
	string(ComponentTypeIconEnumPhlocklaminated),
	string(ComponentTypeIconEnumPhlocklaminatedopen),
	string(ComponentTypeIconEnumPhlockopen),
	string(ComponentTypeIconEnumPhlocksimple),
	string(ComponentTypeIconEnumPhlocksimpleopen),
	string(ComponentTypeIconEnumPhmagicwand),
	string(ComponentTypeIconEnumPhmagnet),
	string(ComponentTypeIconEnumPhmagnetstraight),
	string(ComponentTypeIconEnumPhmagnifyingglass),
	string(ComponentTypeIconEnumPhmagnifyingglassminus),
	string(ComponentTypeIconEnumPhmagnifyingglassplus),
	string(ComponentTypeIconEnumPhmappin),
	string(ComponentTypeIconEnumPhmappinline),
	string(ComponentTypeIconEnumPhmaptrifold),
	string(ComponentTypeIconEnumPhmarkercircle),
	string(ComponentTypeIconEnumPhmartini),
	string(ComponentTypeIconEnumPhmaskhappy),
	string(ComponentTypeIconEnumPhmasksad),
	string(ComponentTypeIconEnumPhmathoperations),
	string(ComponentTypeIconEnumPhmedal),
	string(ComponentTypeIconEnumPhmediumlogo),
	string(ComponentTypeIconEnumPhmegaphone),
	string(ComponentTypeIconEnumPhmegaphonesimple),
	string(ComponentTypeIconEnumPhmessengerlogo),
	string(ComponentTypeIconEnumPhmicrophone),
	string(ComponentTypeIconEnumPhmicrophoneslash),
	string(ComponentTypeIconEnumPhmicrophonestage),
	string(ComponentTypeIconEnumPhmicrosoftexcellogo),
	string(ComponentTypeIconEnumPhmicrosoftpowerpointlogo),
	string(ComponentTypeIconEnumPhmicrosoftteamslogo),
	string(ComponentTypeIconEnumPhmicrosoftwordlogo),
	string(ComponentTypeIconEnumPhminus),
	string(ComponentTypeIconEnumPhminuscircle),
	string(ComponentTypeIconEnumPhmoney),
	string(ComponentTypeIconEnumPhmonitor),
	string(ComponentTypeIconEnumPhmonitorplay),
	string(ComponentTypeIconEnumPhmoon),
	string(ComponentTypeIconEnumPhmoonstars),
	string(ComponentTypeIconEnumPhmountains),
	string(ComponentTypeIconEnumPhmouse),
	string(ComponentTypeIconEnumPhmousesimple),
	string(ComponentTypeIconEnumPhmusicnote),
	string(ComponentTypeIconEnumPhmusicnotesimple),
	string(ComponentTypeIconEnumPhmusicnotes),
	string(ComponentTypeIconEnumPhmusicnotesplus),
	string(ComponentTypeIconEnumPhmusicnotessimple),
	string(ComponentTypeIconEnumPhnavigationarrow),
	string(ComponentTypeIconEnumPhneedle),
	string(ComponentTypeIconEnumPhnewspaper),
	string(ComponentTypeIconEnumPhnewspaperclipping),
	string(ComponentTypeIconEnumPhnote),
	string(ComponentTypeIconEnumPhnoteblank),
	string(ComponentTypeIconEnumPhnotepencil),
	string(ComponentTypeIconEnumPhnotebook),
	string(ComponentTypeIconEnumPhnotepad),
	string(ComponentTypeIconEnumPhnotification),
	string(ComponentTypeIconEnumPhnumbercircleeight),
	string(ComponentTypeIconEnumPhnumbercirclefive),
	string(ComponentTypeIconEnumPhnumbercirclefour),
	string(ComponentTypeIconEnumPhnumbercirclenine),
	string(ComponentTypeIconEnumPhnumbercircleone),
	string(ComponentTypeIconEnumPhnumbercircleseven),
	string(ComponentTypeIconEnumPhnumbercirclesix),
	string(ComponentTypeIconEnumPhnumbercirclethree),
	string(ComponentTypeIconEnumPhnumbercircletwo),
	string(ComponentTypeIconEnumPhnumbercirclezero),
	string(ComponentTypeIconEnumPhnumbereight),
	string(ComponentTypeIconEnumPhnumberfive),
	string(ComponentTypeIconEnumPhnumberfour),
	string(ComponentTypeIconEnumPhnumbernine),
	string(ComponentTypeIconEnumPhnumberone),
	string(ComponentTypeIconEnumPhnumberseven),
	string(ComponentTypeIconEnumPhnumbersix),
	string(ComponentTypeIconEnumPhnumbersquareeight),
	string(ComponentTypeIconEnumPhnumbersquarefive),
	string(ComponentTypeIconEnumPhnumbersquarefour),
	string(ComponentTypeIconEnumPhnumbersquarenine),
	string(ComponentTypeIconEnumPhnumbersquareone),
	string(ComponentTypeIconEnumPhnumbersquareseven),
	string(ComponentTypeIconEnumPhnumbersquaresix),
	string(ComponentTypeIconEnumPhnumbersquarethree),
	string(ComponentTypeIconEnumPhnumbersquaretwo),
	string(ComponentTypeIconEnumPhnumbersquarezero),
	string(ComponentTypeIconEnumPhnumberthree),
	string(ComponentTypeIconEnumPhnumbertwo),
	string(ComponentTypeIconEnumPhnumberzero),
	string(ComponentTypeIconEnumPhnut),
	string(ComponentTypeIconEnumPhnytimeslogo),
	string(ComponentTypeIconEnumPhoctagon),
	string(ComponentTypeIconEnumPhoption),
	string(ComponentTypeIconEnumPhpackage),
	string(ComponentTypeIconEnumPhpaintbrush),
	string(ComponentTypeIconEnumPhpaintbrushbroad),
	string(ComponentTypeIconEnumPhpaintbrushhousehold),
	string(ComponentTypeIconEnumPhpaintbucket),
	string(ComponentTypeIconEnumPhpaintroller),
	string(ComponentTypeIconEnumPhpalette),
	string(ComponentTypeIconEnumPhpaperplane),
	string(ComponentTypeIconEnumPhpaperplaneright),
	string(ComponentTypeIconEnumPhpaperplanetilt),
	string(ComponentTypeIconEnumPhpaperclip),
	string(ComponentTypeIconEnumPhpapercliphorizontal),
	string(ComponentTypeIconEnumPhparachute),
	string(ComponentTypeIconEnumPhpassword),
	string(ComponentTypeIconEnumPhpath),
	string(ComponentTypeIconEnumPhpause),
	string(ComponentTypeIconEnumPhpausecircle),
	string(ComponentTypeIconEnumPhpawprint),
	string(ComponentTypeIconEnumPhpeace),
	string(ComponentTypeIconEnumPhpen),
	string(ComponentTypeIconEnumPhpennib),
	string(ComponentTypeIconEnumPhpennibstraight),
	string(ComponentTypeIconEnumPhpencil),
	string(ComponentTypeIconEnumPhpencilcircle),
	string(ComponentTypeIconEnumPhpencilline),
	string(ComponentTypeIconEnumPhpencilsimple),
	string(ComponentTypeIconEnumPhpencilsimpleline),
	string(ComponentTypeIconEnumPhpercent),
	string(ComponentTypeIconEnumPhperson),
	string(ComponentTypeIconEnumPhpersonsimple),
	string(ComponentTypeIconEnumPhpersonsimplerun),
	string(ComponentTypeIconEnumPhpersonsimplewalk),
	string(ComponentTypeIconEnumPhperspective),
	string(ComponentTypeIconEnumPhphone),
	string(ComponentTypeIconEnumPhphonecall),
	string(ComponentTypeIconEnumPhphonedisconnect),
	string(ComponentTypeIconEnumPhphoneincoming),
	string(ComponentTypeIconEnumPhphoneoutgoing),
	string(ComponentTypeIconEnumPhphoneslash),
	string(ComponentTypeIconEnumPhphonex),
	string(ComponentTypeIconEnumPhphosphorlogo),
	string(ComponentTypeIconEnumPhpianokeys),
	string(ComponentTypeIconEnumPhpictureinpicture),
	string(ComponentTypeIconEnumPhpill),
	string(ComponentTypeIconEnumPhpinterestlogo),
	string(ComponentTypeIconEnumPhpinwheel),
	string(ComponentTypeIconEnumPhpizza),
	string(ComponentTypeIconEnumPhplaceholder),
	string(ComponentTypeIconEnumPhplanet),
	string(ComponentTypeIconEnumPhplay),
	string(ComponentTypeIconEnumPhplaycircle),
	string(ComponentTypeIconEnumPhplaylist),
	string(ComponentTypeIconEnumPhplug),
	string(ComponentTypeIconEnumPhplugs),
	string(ComponentTypeIconEnumPhplugsconnected),
	string(ComponentTypeIconEnumPhplus),
	string(ComponentTypeIconEnumPhpluscircle),
	string(ComponentTypeIconEnumPhplusminus),
	string(ComponentTypeIconEnumPhpokerchip),
	string(ComponentTypeIconEnumPhpolicecar),
	string(ComponentTypeIconEnumPhpolygon),
	string(ComponentTypeIconEnumPhpopcorn),
	string(ComponentTypeIconEnumPhpower),
	string(ComponentTypeIconEnumPhprescription),
	string(ComponentTypeIconEnumPhpresentation),
	string(ComponentTypeIconEnumPhpresentationchart),
	string(ComponentTypeIconEnumPhprinter),
	string(ComponentTypeIconEnumPhprohibit),
	string(ComponentTypeIconEnumPhprohibitinset),
	string(ComponentTypeIconEnumPhprojectorscreen),
	string(ComponentTypeIconEnumPhprojectorscreenchart),
	string(ComponentTypeIconEnumPhpushpin),
	string(ComponentTypeIconEnumPhpushpinsimple),
	string(ComponentTypeIconEnumPhpushpinsimpleslash),
	string(ComponentTypeIconEnumPhpushpinslash),
	string(ComponentTypeIconEnumPhpuzzlepiece),
	string(ComponentTypeIconEnumPhqrcode),
	string(ComponentTypeIconEnumPhquestion),
	string(ComponentTypeIconEnumPhqueue),
	string(ComponentTypeIconEnumPhquotes),
	string(ComponentTypeIconEnumPhradical),
	string(ComponentTypeIconEnumPhradio),
	string(ComponentTypeIconEnumPhradiobutton),
	string(ComponentTypeIconEnumPhrainbow),
	string(ComponentTypeIconEnumPhrainbowcloud),
	string(ComponentTypeIconEnumPhreceipt),
	string(ComponentTypeIconEnumPhrecord),
	string(ComponentTypeIconEnumPhrectangle),
	string(ComponentTypeIconEnumPhrecycle),
	string(ComponentTypeIconEnumPhredditlogo),
	string(ComponentTypeIconEnumPhrepeat),
	string(ComponentTypeIconEnumPhrepeatonce),
	string(ComponentTypeIconEnumPhrewind),
	string(ComponentTypeIconEnumPhrewindcircle),
	string(ComponentTypeIconEnumPhrobot),
	string(ComponentTypeIconEnumPhrocket),
	string(ComponentTypeIconEnumPhrocketlaunch),
	string(ComponentTypeIconEnumPhrows),
	string(ComponentTypeIconEnumPhrss),
	string(ComponentTypeIconEnumPhrsssimple),
	string(ComponentTypeIconEnumPhrug),
	string(ComponentTypeIconEnumPhruler),
	string(ComponentTypeIconEnumPhscales),
	string(ComponentTypeIconEnumPhscan),
	string(ComponentTypeIconEnumPhscissors),
	string(ComponentTypeIconEnumPhscreencast),
	string(ComponentTypeIconEnumPhscribbleloop),
	string(ComponentTypeIconEnumPhscroll),
	string(ComponentTypeIconEnumPhselection),
	string(ComponentTypeIconEnumPhselectionall),
	string(ComponentTypeIconEnumPhselectionbackground),
	string(ComponentTypeIconEnumPhselectionforeground),
	string(ComponentTypeIconEnumPhselectioninverse),
	string(ComponentTypeIconEnumPhselectionplus),
	string(ComponentTypeIconEnumPhselectionslash),
	string(ComponentTypeIconEnumPhshare),
	string(ComponentTypeIconEnumPhsharenetwork),
	string(ComponentTypeIconEnumPhshield),
	string(ComponentTypeIconEnumPhshieldcheck),
	string(ComponentTypeIconEnumPhshieldcheckered),
	string(ComponentTypeIconEnumPhshieldchevron),
	string(ComponentTypeIconEnumPhshieldplus),
	string(ComponentTypeIconEnumPhshieldslash),
	string(ComponentTypeIconEnumPhshieldstar),
	string(ComponentTypeIconEnumPhshieldwarning),
	string(ComponentTypeIconEnumPhshoppingbag),
	string(ComponentTypeIconEnumPhshoppingbagopen),
	string(ComponentTypeIconEnumPhshoppingcart),
	string(ComponentTypeIconEnumPhshoppingcartsimple),
	string(ComponentTypeIconEnumPhshower),
	string(ComponentTypeIconEnumPhshuffle),
	string(ComponentTypeIconEnumPhshuffleangular),
	string(ComponentTypeIconEnumPhshufflesimple),
	string(ComponentTypeIconEnumPhsidebar),
	string(ComponentTypeIconEnumPhsidebarsimple),
	string(ComponentTypeIconEnumPhsignin),
	string(ComponentTypeIconEnumPhsignout),
	string(ComponentTypeIconEnumPhsignpost),
	string(ComponentTypeIconEnumPhsimcard),
	string(ComponentTypeIconEnumPhsketchlogo),
	string(ComponentTypeIconEnumPhskipback),
	string(ComponentTypeIconEnumPhskipbackcircle),
	string(ComponentTypeIconEnumPhskipforward),
	string(ComponentTypeIconEnumPhskipforwardcircle),
	string(ComponentTypeIconEnumPhskull),
	string(ComponentTypeIconEnumPhslacklogo),
	string(ComponentTypeIconEnumPhsliders),
	string(ComponentTypeIconEnumPhslidershorizontal),
	string(ComponentTypeIconEnumPhsmiley),
	string(ComponentTypeIconEnumPhsmileyblank),
	string(ComponentTypeIconEnumPhsmileymeh),
	string(ComponentTypeIconEnumPhsmileynervous),
	string(ComponentTypeIconEnumPhsmileysad),
	string(ComponentTypeIconEnumPhsmileysticker),
	string(ComponentTypeIconEnumPhsmileywink),
	string(ComponentTypeIconEnumPhsmileyxeyes),
	string(ComponentTypeIconEnumPhsnapchatlogo),
	string(ComponentTypeIconEnumPhsnowflake),
	string(ComponentTypeIconEnumPhsoccerball),
	string(ComponentTypeIconEnumPhsortascending),
	string(ComponentTypeIconEnumPhsortdescending),
	string(ComponentTypeIconEnumPhspade),
	string(ComponentTypeIconEnumPhsparkle),
	string(ComponentTypeIconEnumPhspeakerhigh),
	string(ComponentTypeIconEnumPhspeakerlow),
	string(ComponentTypeIconEnumPhspeakernone),
	string(ComponentTypeIconEnumPhspeakersimplehigh),
	string(ComponentTypeIconEnumPhspeakersimplelow),
	string(ComponentTypeIconEnumPhspeakersimplenone),
	string(ComponentTypeIconEnumPhspeakersimpleslash),
	string(ComponentTypeIconEnumPhspeakersimplex),
	string(ComponentTypeIconEnumPhspeakerslash),
	string(ComponentTypeIconEnumPhspeakerx),
	string(ComponentTypeIconEnumPhspinner),
	string(ComponentTypeIconEnumPhspinnergap),
	string(ComponentTypeIconEnumPhspiral),
	string(ComponentTypeIconEnumPhspotifylogo),
	string(ComponentTypeIconEnumPhsquare),
	string(ComponentTypeIconEnumPhsquarehalf),
	string(ComponentTypeIconEnumPhsquarehalfbottom),
	string(ComponentTypeIconEnumPhsquarelogo),
	string(ComponentTypeIconEnumPhsquaresfour),
	string(ComponentTypeIconEnumPhstack),
	string(ComponentTypeIconEnumPhstackoverflowlogo),
	string(ComponentTypeIconEnumPhstacksimple),
	string(ComponentTypeIconEnumPhstamp),
	string(ComponentTypeIconEnumPhstar),
	string(ComponentTypeIconEnumPhstarfour),
	string(ComponentTypeIconEnumPhstarhalf),
	string(ComponentTypeIconEnumPhsticker),
	string(ComponentTypeIconEnumPhstop),
	string(ComponentTypeIconEnumPhstopcircle),
	string(ComponentTypeIconEnumPhstorefront),
	string(ComponentTypeIconEnumPhstrategy),
	string(ComponentTypeIconEnumPhstripelogo),
	string(ComponentTypeIconEnumPhstudent),
	string(ComponentTypeIconEnumPhsuitcase),
	string(ComponentTypeIconEnumPhsuitcasesimple),
	string(ComponentTypeIconEnumPhsun),
	string(ComponentTypeIconEnumPhsundim),
	string(ComponentTypeIconEnumPhsunhorizon),
	string(ComponentTypeIconEnumPhsunglasses),
	string(ComponentTypeIconEnumPhswap),
	string(ComponentTypeIconEnumPhswatches),
	string(ComponentTypeIconEnumPhsword),
	string(ComponentTypeIconEnumPhsyringe),
	string(ComponentTypeIconEnumPhtshirt),
	string(ComponentTypeIconEnumPhtable),
	string(ComponentTypeIconEnumPhtabs),
	string(ComponentTypeIconEnumPhtag),
	string(ComponentTypeIconEnumPhtagchevron),
	string(ComponentTypeIconEnumPhtagsimple),
	string(ComponentTypeIconEnumPhtarget),
	string(ComponentTypeIconEnumPhtaxi),
	string(ComponentTypeIconEnumPhtelegramlogo),
	string(ComponentTypeIconEnumPhtelevision),
	string(ComponentTypeIconEnumPhtelevisionsimple),
	string(ComponentTypeIconEnumPhtennisball),
	string(ComponentTypeIconEnumPhterminal),
	string(ComponentTypeIconEnumPhterminalwindow),
	string(ComponentTypeIconEnumPhtesttube),
	string(ComponentTypeIconEnumPhtextaa),
	string(ComponentTypeIconEnumPhtextaligncenter),
	string(ComponentTypeIconEnumPhtextalignjustify),
	string(ComponentTypeIconEnumPhtextalignleft),
	string(ComponentTypeIconEnumPhtextalignright),
	string(ComponentTypeIconEnumPhtextbolder),
	string(ComponentTypeIconEnumPhtexth),
	string(ComponentTypeIconEnumPhtexthfive),
	string(ComponentTypeIconEnumPhtexthfour),
	string(ComponentTypeIconEnumPhtexthone),
	string(ComponentTypeIconEnumPhtexthsix),
	string(ComponentTypeIconEnumPhtexththree),
	string(ComponentTypeIconEnumPhtexthtwo),
	string(ComponentTypeIconEnumPhtextindent),
	string(ComponentTypeIconEnumPhtextitalic),
	string(ComponentTypeIconEnumPhtextoutdent),
	string(ComponentTypeIconEnumPhtextstrikethrough),
	string(ComponentTypeIconEnumPhtextt),
	string(ComponentTypeIconEnumPhtextunderline),
	string(ComponentTypeIconEnumPhtextbox),
	string(ComponentTypeIconEnumPhthermometer),
	string(ComponentTypeIconEnumPhthermometercold),
	string(ComponentTypeIconEnumPhthermometerhot),
	string(ComponentTypeIconEnumPhthermometersimple),
	string(ComponentTypeIconEnumPhthumbsdown),
	string(ComponentTypeIconEnumPhthumbsup),
	string(ComponentTypeIconEnumPhticket),
	string(ComponentTypeIconEnumPhtiktoklogo),
	string(ComponentTypeIconEnumPhtimer),
	string(ComponentTypeIconEnumPhtoggleleft),
	string(ComponentTypeIconEnumPhtoggleright),
	string(ComponentTypeIconEnumPhtoilet),
	string(ComponentTypeIconEnumPhtoiletpaper),
	string(ComponentTypeIconEnumPhtote),
	string(ComponentTypeIconEnumPhtotesimple),
	string(ComponentTypeIconEnumPhtrademarkregistered),
	string(ComponentTypeIconEnumPhtrafficcone),
	string(ComponentTypeIconEnumPhtrafficsign),
	string(ComponentTypeIconEnumPhtrafficsignal),
	string(ComponentTypeIconEnumPhtrain),
	string(ComponentTypeIconEnumPhtrainregional),
	string(ComponentTypeIconEnumPhtrainsimple),
	string(ComponentTypeIconEnumPhtranslate),
	string(ComponentTypeIconEnumPhtrash),
	string(ComponentTypeIconEnumPhtrashsimple),
	string(ComponentTypeIconEnumPhtray),
	string(ComponentTypeIconEnumPhtree),
	string(ComponentTypeIconEnumPhtreeevergreen),
	string(ComponentTypeIconEnumPhtreestructure),
	string(ComponentTypeIconEnumPhtrenddown),
	string(ComponentTypeIconEnumPhtrendup),
	string(ComponentTypeIconEnumPhtriangle),
	string(ComponentTypeIconEnumPhtrophy),
	string(ComponentTypeIconEnumPhtruck),
	string(ComponentTypeIconEnumPhtwitchlogo),
	string(ComponentTypeIconEnumPhtwitterlogo),
	string(ComponentTypeIconEnumPhumbrella),
	string(ComponentTypeIconEnumPhumbrellasimple),
	string(ComponentTypeIconEnumPhupload),
	string(ComponentTypeIconEnumPhuploadsimple),
	string(ComponentTypeIconEnumPhuser),
	string(ComponentTypeIconEnumPhusercircle),
	string(ComponentTypeIconEnumPhusercirclegear),
	string(ComponentTypeIconEnumPhusercircleminus),
	string(ComponentTypeIconEnumPhusercircleplus),
	string(ComponentTypeIconEnumPhuserfocus),
	string(ComponentTypeIconEnumPhusergear),
	string(ComponentTypeIconEnumPhuserlist),
	string(ComponentTypeIconEnumPhuserminus),
	string(ComponentTypeIconEnumPhuserplus),
	string(ComponentTypeIconEnumPhuserrectangle),
	string(ComponentTypeIconEnumPhusersquare),
	string(ComponentTypeIconEnumPhuserswitch),
	string(ComponentTypeIconEnumPhusers),
	string(ComponentTypeIconEnumPhusersfour),
	string(ComponentTypeIconEnumPhusersthree),
	string(ComponentTypeIconEnumPhvault),
	string(ComponentTypeIconEnumPhvibrate),
	string(ComponentTypeIconEnumPhvideocamera),
	string(ComponentTypeIconEnumPhvideocameraslash),
	string(ComponentTypeIconEnumPhvignette),
	string(ComponentTypeIconEnumPhvoicemail),
	string(ComponentTypeIconEnumPhvolleyball),
	string(ComponentTypeIconEnumPhwall),
	string(ComponentTypeIconEnumPhwallet),
	string(ComponentTypeIconEnumPhwarning),
	string(ComponentTypeIconEnumPhwarningcircle),
	string(ComponentTypeIconEnumPhwarningoctagon),
	string(ComponentTypeIconEnumPhwatch),
	string(ComponentTypeIconEnumPhwavesawtooth),
	string(ComponentTypeIconEnumPhwavesine),
	string(ComponentTypeIconEnumPhwavesquare),
	string(ComponentTypeIconEnumPhwavetriangle),
	string(ComponentTypeIconEnumPhwaves),
	string(ComponentTypeIconEnumPhwebcam),
	string(ComponentTypeIconEnumPhwhatsapplogo),
	string(ComponentTypeIconEnumPhwheelchair),
	string(ComponentTypeIconEnumPhwifihigh),
	string(ComponentTypeIconEnumPhwifilow),
	string(ComponentTypeIconEnumPhwifimedium),
	string(ComponentTypeIconEnumPhwifinone),
	string(ComponentTypeIconEnumPhwifislash),
	string(ComponentTypeIconEnumPhwifix),
	string(ComponentTypeIconEnumPhwind),
	string(ComponentTypeIconEnumPhwindowslogo),
	string(ComponentTypeIconEnumPhwine),
	string(ComponentTypeIconEnumPhwrench),
	string(ComponentTypeIconEnumPhx),
	string(ComponentTypeIconEnumPhxcircle),
	string(ComponentTypeIconEnumPhxsquare),
	string(ComponentTypeIconEnumPhyinyang),
	string(ComponentTypeIconEnumPhyoutubelogo),
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
	CustomActionsTriggerEventStatusEnumFailure         CustomActionsTriggerEventStatusEnum = "FAILURE"          // The action failed to complete
	CustomActionsTriggerEventStatusEnumPending         CustomActionsTriggerEventStatusEnum = "PENDING"          // A result has not been determined
	CustomActionsTriggerEventStatusEnumPendingApproval CustomActionsTriggerEventStatusEnum = "PENDING_APPROVAL" // The action is waiting for an approval before it executes
	CustomActionsTriggerEventStatusEnumSuccess         CustomActionsTriggerEventStatusEnum = "SUCCESS"          // The action completed successfully
)

// All CustomActionsTriggerEventStatusEnum as []string
var AllCustomActionsTriggerEventStatusEnum = []string{
	string(CustomActionsTriggerEventStatusEnumFailure),
	string(CustomActionsTriggerEventStatusEnumPending),
	string(CustomActionsTriggerEventStatusEnumPendingApproval),
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

// DeployStatusEnum The possible statuses of a deploy
type DeployStatusEnum string

var (
	DeployStatusEnumCanceled DeployStatusEnum = "canceled"  // The deploy was canceled
	DeployStatusEnumFailure  DeployStatusEnum = "failure"   // The deploy failed
	DeployStatusEnumNoStatus DeployStatusEnum = "no_status" // The deploy has no recognized status
	DeployStatusEnumQueued   DeployStatusEnum = "queued"    // The deploy is queued
	DeployStatusEnumRunning  DeployStatusEnum = "running"   // The deploy is currently running
	DeployStatusEnumSuccess  DeployStatusEnum = "success"   // The deploy was successful
)

// All DeployStatusEnum as []string
var AllDeployStatusEnum = []string{
	string(DeployStatusEnumCanceled),
	string(DeployStatusEnumFailure),
	string(DeployStatusEnumNoStatus),
	string(DeployStatusEnumQueued),
	string(DeployStatusEnumRunning),
	string(DeployStatusEnumSuccess),
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

var HasDocumentationSubtypeEnumOpenapi HasDocumentationSubtypeEnum = "openapi" // Document is an OpenAPI document

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

var PayloadFilterEnumIntegrationID PayloadFilterEnum = "integration_id" // Filter by `integration` field. Note that this is an internal id, ex. "123"

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
	PredicateKeyEnumAliases         PredicateKeyEnum = "aliases"           // Filter by Alias attached to this service, if any
	PredicateKeyEnumComponentTypeID PredicateKeyEnum = "component_type_id" // Filter by the `component_type` field
	PredicateKeyEnumCreationSource  PredicateKeyEnum = "creation_source"   // Filter by the creation source
	PredicateKeyEnumDomainID        PredicateKeyEnum = "domain_id"         // Filter by Domain that includes the System this service is assigned to, if any
	PredicateKeyEnumFilterID        PredicateKeyEnum = "filter_id"         // Filter by another filter
	PredicateKeyEnumFramework       PredicateKeyEnum = "framework"         // Filter by `framework` field
	PredicateKeyEnumGroupIDs        PredicateKeyEnum = "group_ids"         // Filter by group hierarchy. Will return resources who's owner is in the group ancestry chain
	PredicateKeyEnumLanguage        PredicateKeyEnum = "language"          // Filter by `language` field
	PredicateKeyEnumLifecycleIndex  PredicateKeyEnum = "lifecycle_index"   // Filter by `lifecycle` field
	PredicateKeyEnumName            PredicateKeyEnum = "name"              // Filter by `name` field
	PredicateKeyEnumOwnerID         PredicateKeyEnum = "owner_id"          // Filter by `owner` field
	PredicateKeyEnumOwnerIDs        PredicateKeyEnum = "owner_ids"         // Filter by `owner` hierarchy. Will return resources who's owner is in the team ancestry chain
	PredicateKeyEnumProduct         PredicateKeyEnum = "product"           // Filter by `product` field
	PredicateKeyEnumProperties      PredicateKeyEnum = "properties"        // Filter by custom-defined properties
	PredicateKeyEnumRelationships   PredicateKeyEnum = "relationships"     // Filter by `relationships`
	PredicateKeyEnumRepositoryIDs   PredicateKeyEnum = "repository_ids"    // Filter by Repository that this service is attached to, if any
	PredicateKeyEnumSystemID        PredicateKeyEnum = "system_id"         // Filter by System that this service is assigned to, if any
	PredicateKeyEnumTags            PredicateKeyEnum = "tags"              // Filter by `tags` field
	PredicateKeyEnumTierIndex       PredicateKeyEnum = "tier_index"        // Filter by `tier` field
)

// All PredicateKeyEnum as []string
var AllPredicateKeyEnum = []string{
	string(PredicateKeyEnumAliases),
	string(PredicateKeyEnumComponentTypeID),
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
	string(PredicateKeyEnumRelationships),
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

// PropertyDisplayStatusEnum The UI display status of a custom property
type PropertyDisplayStatusEnum string

var (
	PropertyDisplayStatusEnumHidden  PropertyDisplayStatusEnum = "hidden"  // The property is not shown on resource pages
	PropertyDisplayStatusEnumVisible PropertyDisplayStatusEnum = "visible" // The property is shown on resource pages
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

// ProvisionedByEnum
type ProvisionedByEnum string

var (
	ProvisionedByEnumAPICli          ProvisionedByEnum = "api_cli"          //
	ProvisionedByEnumAPIOther        ProvisionedByEnum = "api_other"        //
	ProvisionedByEnumAPITerraform    ProvisionedByEnum = "api_terraform"    //
	ProvisionedByEnumBackstage       ProvisionedByEnum = "backstage"        //
	ProvisionedByEnumIntegrationScim ProvisionedByEnum = "integration_scim" //
	ProvisionedByEnumSsoOkta         ProvisionedByEnum = "sso_okta"         //
	ProvisionedByEnumSsoOther        ProvisionedByEnum = "sso_other"        //
	ProvisionedByEnumUnknown         ProvisionedByEnum = "unknown"          //
	ProvisionedByEnumUser            ProvisionedByEnum = "user"             //
)

// All ProvisionedByEnum as []string
var AllProvisionedByEnum = []string{
	string(ProvisionedByEnumAPICli),
	string(ProvisionedByEnumAPIOther),
	string(ProvisionedByEnumAPITerraform),
	string(ProvisionedByEnumBackstage),
	string(ProvisionedByEnumIntegrationScim),
	string(ProvisionedByEnumSsoOkta),
	string(ProvisionedByEnumSsoOther),
	string(ProvisionedByEnumUnknown),
	string(ProvisionedByEnumUser),
}

// RelatedResourceRelationshipTypeEnum The type of the relationship between two resources
type RelatedResourceRelationshipTypeEnum string

var (
	RelatedResourceRelationshipTypeEnumBelongsTo    RelatedResourceRelationshipTypeEnum = "belongs_to"    // The resource belongs to the node on the edge
	RelatedResourceRelationshipTypeEnumContains     RelatedResourceRelationshipTypeEnum = "contains"      // The resource contains the node on the edge
	RelatedResourceRelationshipTypeEnumDependencyOf RelatedResourceRelationshipTypeEnum = "dependency_of" // The resource is a dependency of the node on the edge
	RelatedResourceRelationshipTypeEnumDependsOn    RelatedResourceRelationshipTypeEnum = "depends_on"    // The resource depends on the node on the edge
	RelatedResourceRelationshipTypeEnumIsRelatedTo  RelatedResourceRelationshipTypeEnum = "is_related_to" // The resource is part of a specialized relationship defined on another node
	RelatedResourceRelationshipTypeEnumMemberOf     RelatedResourceRelationshipTypeEnum = "member_of"     // The resource is a member of the node on the edge
	RelatedResourceRelationshipTypeEnumRelatedTo    RelatedResourceRelationshipTypeEnum = "related_to"    // The resource has a specialized relationship to another node
)

// All RelatedResourceRelationshipTypeEnum as []string
var AllRelatedResourceRelationshipTypeEnum = []string{
	string(RelatedResourceRelationshipTypeEnumBelongsTo),
	string(RelatedResourceRelationshipTypeEnumContains),
	string(RelatedResourceRelationshipTypeEnumDependencyOf),
	string(RelatedResourceRelationshipTypeEnumDependsOn),
	string(RelatedResourceRelationshipTypeEnumIsRelatedTo),
	string(RelatedResourceRelationshipTypeEnumMemberOf),
	string(RelatedResourceRelationshipTypeEnumRelatedTo),
}

// RelationshipDefinitionManagementRuleOperator The operator used in a relationship definition management rule
type RelationshipDefinitionManagementRuleOperator string

var (
	RelationshipDefinitionManagementRuleOperatorArrayContains RelationshipDefinitionManagementRuleOperator = "ARRAY_CONTAINS" //
	RelationshipDefinitionManagementRuleOperatorEquals        RelationshipDefinitionManagementRuleOperator = "EQUALS"         //
)

// All RelationshipDefinitionManagementRuleOperator as []string
var AllRelationshipDefinitionManagementRuleOperator = []string{
	string(RelationshipDefinitionManagementRuleOperatorArrayContains),
	string(RelationshipDefinitionManagementRuleOperatorEquals),
}

// RelationshipTypeEnum The type of relationship between two resources
type RelationshipTypeEnum string

var (
	RelationshipTypeEnumBelongsTo RelationshipTypeEnum = "belongs_to" // The source resource belongs to the target resource. Can be used to allow Components to belong to Systems and Domains, or for Infrastructure to belong to Components
	RelationshipTypeEnumDependsOn RelationshipTypeEnum = "depends_on" // The source resource depends on the target resource. Can be used to specify that a Component depends on some Infrastructure, or that a System depends on a Component
	RelationshipTypeEnumRelatedTo RelationshipTypeEnum = "related_to" // The source resource is related to the target resource through a custom relationship definition. These are dynamic and can be used to extend our out-of-the-box relationships
)

// All RelationshipTypeEnum as []string
var AllRelationshipTypeEnum = []string{
	string(RelationshipTypeEnumBelongsTo),
	string(RelationshipTypeEnumDependsOn),
	string(RelationshipTypeEnumRelatedTo),
}

// RepositorySBOMGenerationConfigEnum The enumerated list of configuration values for SBOM generation at the repository level
type RepositorySBOMGenerationConfigEnum string

var (
	RepositorySBOMGenerationConfigEnumOptIn  RepositorySBOMGenerationConfigEnum = "opt_in"  // Indicates that the repository will opt in to automated SBOM generation if it would be otherwise enabled at an integration or account level
	RepositorySBOMGenerationConfigEnumOptOut RepositorySBOMGenerationConfigEnum = "opt_out" // Indicates that the repository will opt out of automated SBOM generation if it would be otherwise enabled at an integration or account level
)

// All RepositorySBOMGenerationConfigEnum as []string
var AllRepositorySBOMGenerationConfigEnum = []string{
	string(RepositorySBOMGenerationConfigEnumOptIn),
	string(RepositorySBOMGenerationConfigEnumOptOut),
}

// RepositorySBOMGenerationDisabledReasonEnum The set of values that explain why SBOM autogeneration is disabled
type RepositorySBOMGenerationDisabledReasonEnum string

var (
	RepositorySBOMGenerationDisabledReasonEnumAccount     RepositorySBOMGenerationDisabledReasonEnum = "account"     // SBOM autogeneration is disabled at the account level
	RepositorySBOMGenerationDisabledReasonEnumIntegration RepositorySBOMGenerationDisabledReasonEnum = "integration" // SBOM autogeneration is disabled at the integration level
	RepositorySBOMGenerationDisabledReasonEnumRepository  RepositorySBOMGenerationDisabledReasonEnum = "repository"  // SBOM autogeneration is disabled at the repository level
)

// All RepositorySBOMGenerationDisabledReasonEnum as []string
var AllRepositorySBOMGenerationDisabledReasonEnum = []string{
	string(RepositorySBOMGenerationDisabledReasonEnumAccount),
	string(RepositorySBOMGenerationDisabledReasonEnumIntegration),
	string(RepositorySBOMGenerationDisabledReasonEnumRepository),
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

// ServiceFilterEnum Fields that can be used as part of filter for services
type ServiceFilterEnum string

var (
	ServiceFilterEnumAlertStatus       ServiceFilterEnum = "alert_status"       // Filter by `alert status` field
	ServiceFilterEnumAliases           ServiceFilterEnum = "aliases"            // Filter by Alias attached to this service, if any
	ServiceFilterEnumComponentTypeID   ServiceFilterEnum = "component_type_id"  // Filter by the type of service
	ServiceFilterEnumCreationSource    ServiceFilterEnum = "creation_source"    // Filter by the creation source
	ServiceFilterEnumDeployEnvironment ServiceFilterEnum = "deploy_environment" // Filter by the existence of a deploy to an environment
	ServiceFilterEnumDomainID          ServiceFilterEnum = "domain_id"          // Filter by Domain that includes the System this service is assigned to, if any
	ServiceFilterEnumFilterID          ServiceFilterEnum = "filter_id"          // Filter by another filter
	ServiceFilterEnumFramework         ServiceFilterEnum = "framework"          // Filter by `framework` field
	ServiceFilterEnumGroupIDs          ServiceFilterEnum = "group_ids"          // Filter by group hierarchy. Will return resources who's owner is in the group ancestry chain
	ServiceFilterEnumLanguage          ServiceFilterEnum = "language"           // Filter by `language` field
	ServiceFilterEnumLevelIndex        ServiceFilterEnum = "level_index"        // Filter by `level` field
	ServiceFilterEnumLifecycleIndex    ServiceFilterEnum = "lifecycle_index"    // Filter by `lifecycle` field
	ServiceFilterEnumName              ServiceFilterEnum = "name"               // Filter by `name` field
	ServiceFilterEnumOwnerID           ServiceFilterEnum = "owner_id"           // Filter by `owner` field
	ServiceFilterEnumOwnerIDs          ServiceFilterEnum = "owner_ids"          // Filter by `owner` hierarchy. Will return resources who's owner is in the team ancestry chain
	ServiceFilterEnumProduct           ServiceFilterEnum = "product"            // Filter by `product` field
	ServiceFilterEnumProperty          ServiceFilterEnum = "property"           // Filter by a custom-defined property value
	ServiceFilterEnumRelationship      ServiceFilterEnum = "relationship"       // Filter by the existence of a relationship to another catalog component
	ServiceFilterEnumRepositoryIDs     ServiceFilterEnum = "repository_ids"     // Filter by Repository that this service is attached to, if any
	ServiceFilterEnumSystemID          ServiceFilterEnum = "system_id"          // Filter by System that this service is assigned to, if any
	ServiceFilterEnumTag               ServiceFilterEnum = "tag"                // Filter by `tag` field
	ServiceFilterEnumTierIndex         ServiceFilterEnum = "tier_index"         // Filter by `tier` field
)

// All ServiceFilterEnum as []string
var AllServiceFilterEnum = []string{
	string(ServiceFilterEnumAlertStatus),
	string(ServiceFilterEnumAliases),
	string(ServiceFilterEnumComponentTypeID),
	string(ServiceFilterEnumCreationSource),
	string(ServiceFilterEnumDeployEnvironment),
	string(ServiceFilterEnumDomainID),
	string(ServiceFilterEnumFilterID),
	string(ServiceFilterEnumFramework),
	string(ServiceFilterEnumGroupIDs),
	string(ServiceFilterEnumLanguage),
	string(ServiceFilterEnumLevelIndex),
	string(ServiceFilterEnumLifecycleIndex),
	string(ServiceFilterEnumName),
	string(ServiceFilterEnumOwnerID),
	string(ServiceFilterEnumOwnerIDs),
	string(ServiceFilterEnumProduct),
	string(ServiceFilterEnumProperty),
	string(ServiceFilterEnumRelationship),
	string(ServiceFilterEnumRepositoryIDs),
	string(ServiceFilterEnumSystemID),
	string(ServiceFilterEnumTag),
	string(ServiceFilterEnumTierIndex),
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

// TypeEnum Operations that can be used on filters
type TypeEnum string

var (
	TypeEnumBelongsTo                  TypeEnum = "belongs_to"                   // Belongs to a group's hierarchy
	TypeEnumContains                   TypeEnum = "contains"                     // Contains a specific value
	TypeEnumDoesNotContain             TypeEnum = "does_not_contain"             // Does not contain a specific value
	TypeEnumDoesNotEqual               TypeEnum = "does_not_equal"               // Does not equal a specific value
	TypeEnumDoesNotExist               TypeEnum = "does_not_exist"               // Specific attribute does not exist
	TypeEnumDoesNotMatch               TypeEnum = "does_not_match"               // A certain filter is not matched
	TypeEnumDoesNotMatchRegex          TypeEnum = "does_not_match_regex"         // Does not match a value using a regular expression
	TypeEnumEndsWith                   TypeEnum = "ends_with"                    // Ends with a specific value
	TypeEnumEquals                     TypeEnum = "equals"                       // Equals a specific value
	TypeEnumExists                     TypeEnum = "exists"                       // Specific attribute exists
	TypeEnumGreaterThanOrEqualTo       TypeEnum = "greater_than_or_equal_to"     // Greater than or equal to a specific value (numeric only)
	TypeEnumLessThanOrEqualTo          TypeEnum = "less_than_or_equal_to"        // Less than or equal to a specific value (numeric only)
	TypeEnumMatches                    TypeEnum = "matches"                      // A certain filter is matched
	TypeEnumMatchesRegex               TypeEnum = "matches_regex"                // Matches a value using a regular expression
	TypeEnumSatisfiesJqExpression      TypeEnum = "satisfies_jq_expression"      // Satisfies an expression defined in jq (property value only)
	TypeEnumSatisfiesVersionConstraint TypeEnum = "satisfies_version_constraint" // Satisfies version constraint (tag value only)
	TypeEnumStartsWith                 TypeEnum = "starts_with"                  // Starts with a specific value
)

// All TypeEnum as []string
var AllTypeEnum = []string{
	string(TypeEnumBelongsTo),
	string(TypeEnumContains),
	string(TypeEnumDoesNotContain),
	string(TypeEnumDoesNotEqual),
	string(TypeEnumDoesNotExist),
	string(TypeEnumDoesNotMatch),
	string(TypeEnumDoesNotMatchRegex),
	string(TypeEnumEndsWith),
	string(TypeEnumEquals),
	string(TypeEnumExists),
	string(TypeEnumGreaterThanOrEqualTo),
	string(TypeEnumLessThanOrEqualTo),
	string(TypeEnumMatches),
	string(TypeEnumMatchesRegex),
	string(TypeEnumSatisfiesJqExpression),
	string(TypeEnumSatisfiesVersionConstraint),
	string(TypeEnumStartsWith),
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

var UsersInviteScopeEnumPending UsersInviteScopeEnum = "pending" // All users who have yet to log in to OpsLevel for the first time

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
