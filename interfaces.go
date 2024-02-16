// Code generated by gen.go; DO NOT EDIT.

package opslevel

import "github.com/relvacode/iso8601"

// Check represents checks allow you to monitor how your services are built and operated.
type Check struct {
	Category    Category     `graphql:"category"`        // The category that the check belongs to.
	Description string       `graphql:"description"`     // Description of the check type's purpose.
	EnableOn    iso8601.Time `graphql:"enableOn"`        // The date when the check will be automatically enabled.
	Enabled     bool         `graphql:"enabled"`         // If the check is enabled or not.
	Filter      Filter       `graphql:"filter"`          // The filter that the check belongs to.
	Id          ID           `graphql:"id"`              // The id of the check.
	Level       Level        `graphql:"level"`           // The level that the check belongs to.
	Name        string       `graphql:"name"`            // The display name of the check.
	Notes       string       `graphql:"notes: rawNotes"` // Additional information about the check.
	Owner       CheckOwner   `graphql:"owner"`           // The owner of the check.
	Type        CheckType    `graphql:"type"`            // The type of check.

	AlertSourceUsageCheckFragment `graphql:"... on AlertSourceUsageCheck"`
	CustomEventCheckFragment      `graphql:"... on CustomEventCheck"`
	HasRecentDeployCheckFragment  `graphql:"... on HasRecentDeployCheck"`
	ManualCheckFragment           `graphql:"... on ManualCheck"`
	RepositoryFileCheckFragment   `graphql:"... on RepositoryFileCheck"`
	RepositoryGrepCheckFragment   `graphql:"... on RepositoryGrepCheck"`
	RepositorySearchCheckFragment `graphql:"... on RepositorySearchCheck"`
	ServiceOwnershipCheckFragment `graphql:"... on ServiceOwnershipCheck"`
	ServicePropertyCheckFragment  `graphql:"... on ServicePropertyCheck"`
	TagDefinedCheckFragment       `graphql:"... on TagDefinedCheck"`
	ToolUsageCheckFragment        `graphql:"... on ToolUsageCheck"`
	HasDocumentationCheckFragment `graphql:"... on HasDocumentationCheck"`
}

// CustomActionsExternalAction represents an external action to be triggered by a custom action.
type CustomActionsExternalAction struct {
	CustomActionsId

	Description    string `graphql:"description"`    // A description of what the action should accomplish.
	LiquidTemplate string `graphql:"liquidTemplate"` // The liquid template used to generate the data sent to the external action.
	Name           string `graphql:"name"`           // The name of the external action.

	CustomActionsWebhookAction `graphql:"... on CustomActionsWebhookAction"`
}

// CustomActionsTriggerDefinitionBase represents .
type CustomActionsTriggerDefinitionBase struct {
	AccessControl          string `graphql:"accessControl"`          // The set of users that should be able to use the trigger definition.
	Description            string `graphql:"description"`            // The description of what the trigger definition will do, supports Markdown.
	ManualInputsDefinition string `graphql:"manualInputsDefinition"` // The YAML definition of any custom inputs for this trigger definition.
	Name                   string `graphql:"name"`                   // The name of the trigger definition.
	Published              bool   `graphql:"published"`              // The published state of the action; true if the definition is ready for use; false if it is a draft.
	ResponseTemplate       string `graphql:"responseTemplate"`       // The liquid template used to parse the response from the External Action.
}

// HasProperties represents an entity type that can have custom properties.
type HasProperties struct {
	Properties string `graphql:"properties"` // Custom properties assigned to this entity.
	Property   string `graphql:"property"`   // A custom property value assigned to this entity.
}

// Integration represents an integration is a way of extending OpsLevel functionality.
type Integration struct {
	IntegrationId

	CreatedAt   iso8601.Time `graphql:"createdAt"`   // The time this integration was created.
	InstalledAt iso8601.Time `graphql:"installedAt"` // The time that this integration was successfully installed, if null, this indicates the integration was not completed installed.

	AWSIntegrationFragment      `graphql:"... on AwsIntegration"`
	NewRelicIntegrationFragment `graphql:"... on NewRelicIntegration"`
}

// ManualAlertSourceSync represents .
type ManualAlertSourceSync struct {
	AllowManualSyncAlertSources string `graphql:"allowManualSyncAlertSources"` // Indicates if manual alert source synchronization can be triggered.
	LastManualSyncAlertSources  string `graphql:"lastManualSyncAlertSources"`  // The time that alert sources were last manually synchronized at.
}
