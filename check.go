package opslevel

import (
	"fmt"

	"github.com/relvacode/iso8601"
	"github.com/shurcooL/graphql"
)

type CheckOwner struct {
	Team TeamId `graphql:"... on Team"`
	// User User `graphql:"... on User"` // TODO: will this be public?
}

type Check struct {
	Category    Category     `graphql:"category"`
	Description string       `graphql:"description"`
	Enabled     bool         `graphql:"enabled"`
	EnableOn    iso8601.Time `graphql:"enableOn"`
	Filter      Filter       `graphql:"filter"`
	Id          graphql.ID   `graphql:"id"`
	Level       Level        `graphql:"level"`
	Name        string       `graphql:"name"`
	Notes       string       `graphql:"notes"`
	Owner       CheckOwner   `graphql:"owner"`
	Type        CheckType    `graphql:"type"`

	AlertSourceUsageCheckFragment `graphql:"... on AlertSourceUsageCheck"`
	CustomEventCheckFragment      `graphql:"... on CustomEventCheck"`
	HasRecentDeployCheckFragment  `graphql:"... on HasRecentDeployCheck"`
	ManualCheckFragment           `graphql:"... on ManualCheck"`
	RepositoryFileCheckFragment   `graphql:"... on RepositoryFileCheck"`
	RepositorySearchCheckFragment `graphql:"... on RepositorySearchCheck"`
	ServiceOwnershipCheckFragment `graphql:"... on ServiceOwnershipCheck"`
	ServicePropertyCheckFragment  `graphql:"... on ServicePropertyCheck"`
	TagDefinedCheckFragment       `graphql:"... on TagDefinedCheck"`
	ToolUsageCheckFragment        `graphql:"... on ToolUsageCheck"`
	HasDocumentationCheckFragment `graphql:"... on HasDocumentationCheck"`
}

type AlertSourceUsageCheckFragment struct {
	AlertSourceNamePredicate Predicate           `graphql:"alertSourceNamePredicate"`
	AlertSourceType          AlertSourceTypeEnum `graphql:"alertSourceType"`
}

type CustomEventCheckFragment struct {
	Integration      Integration `graphql:"integration"`
	PassPending      bool        `graphql:"passPending"`
	ResultMessage    string      `graphql:"resultMessage"`
	ServiceSelector  string      `graphql:"serviceSelector"`
	SuccessCondition string      `graphql:"successCondition"`
}

type ServiceOwnershipCheckFragment struct {
	RequireContactMethod *bool                             `graphql:"requireContactMethod"`
	ContactMethod        *ServiceOwnershipCheckContactType `graphql:"contactMethod"`
	TeamTagKey           string                            `graphql:"tagKey"`
	TeamTagPredicate     *Predicate                        `graphql:"tagPredicate"`
}

type GitBranchProtectionCheckFragment struct {
}

type HasRecentDeployCheckFragment struct {
	Days int `graphql:"days"`
}

type HasDocumentationCheckFragment struct {
	DocumentType    HasDocumentationTypeEnum    `graphql:"documentType"`
	DocumentSubtype HasDocumentationSubtypeEnum `graphql:"documentSubtype"`
}

type ManualCheckFragment struct {
	UpdateFrequency       *ManualCheckFrequency `graphql:"updateFrequency"`
	UpdateRequiresComment bool                  `graphql:"updateRequiresComment"`
}

type RepositoryFileCheckFragment struct {
	DirectorySearch       bool       `graphql:"directorySearch"`
	Filepaths             []string   `graphql:"filePaths"`
	FileContentsPredicate *Predicate `graphql:"fileContentsPredicate"`
	UseAbsoluteRoot       bool       `graphql:"useAbsoluteRoot"`
}

type RepositoryGrepCheckFragment struct {
	DirectorySearch       bool       `graphql:"directorySearch"`
	Filepaths             []string   `graphql:"filePaths"`
	FileContentsPredicate *Predicate `graphql:"fileContentsPredicate"`
}

type RepositorySearchCheckFragment struct {
	FileExtensions        []string  `graphql:"fileExtensions"`
	FileContentsPredicate Predicate `graphql:"fileContentsPredicate"`
}

type ServicePropertyCheckFragment struct {
	Property  ServicePropertyTypeEnum `graphql:"serviceProperty"`
	Predicate *Predicate              `graphql:"propertyValuePredicate"`
}

type TagDefinedCheckFragment struct {
	TagKey       string     `graphql:"tagKey"`
	TagPredicate *Predicate `graphql:"tagPredicate"`
}

type ToolUsageCheckFragment struct {
	ToolCategory         ToolCategory `graphql:"toolCategory"`
	ToolNamePredicate    *Predicate   `graphql:"toolNamePredicate"`
	ToolUrlPredicate     *Predicate   `graphql:"toolUrlPredicate"`
	EnvironmentPredicate *Predicate   `graphql:"environmentPredicate"`
}

type CheckConnection struct {
	Nodes      []Check
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type CheckCreateInputProvider interface {
	GetCheckCreateInput() *CheckCreateInput
}

type CheckCreateInput struct {
	Name     string        `json:"name"`
	Enabled  bool          `json:"enabled"`
	EnableOn *iso8601.Time `json:"enableOn,omitempty"`
	Category graphql.ID    `json:"categoryId"`
	Level    graphql.ID    `json:"levelId"`
	Owner    *graphql.ID   `json:"ownerId,omitempty"`
	Filter   *graphql.ID   `json:"filterId,omitempty"`
	Notes    string        `json:"notes"`
}

func (c *CheckCreateInput) GetCheckCreateInput() *CheckCreateInput {
	return c
}

type CheckUpdateInputProvider interface {
	GetCheckUpdateInput() *CheckUpdateInput
}

type CheckUpdateInput struct {
	Id       graphql.ID    `json:"id"`
	Name     string        `json:"name,omitempty"`
	Enabled  *bool         `json:"enabled,omitempty"`
	EnableOn *iso8601.Time `json:"enableOn,omitempty"`
	Category *graphql.ID   `json:"categoryId"`
	Level    *graphql.ID   `json:"levelId"`
	Owner    *graphql.ID   `json:"ownerId"`
	Filter   *graphql.ID   `json:"filterId"`
	Notes    string        `json:"notes"`
}

func (c *CheckUpdateInput) GetCheckUpdateInput() *CheckUpdateInput {
	return c
}

type CheckAlertSourceUsageCreateInput struct {
	CheckCreateInput

	AlertSourceType          AlertSourceTypeEnum `json:"alertSourceType,omitempty"`
	AlertSourceNamePredicate *PredicateInput     `json:"alertSourceNamePredicate,omitempty"`
}

type CheckAlertSourceUsageUpdateInput struct {
	CheckUpdateInput

	AlertSourceType          AlertSourceTypeEnum   `json:"alertSourceType,omitempty"`
	AlertSourceNamePredicate *PredicateUpdateInput `json:"alertSourceNamePredicate,omitempty"`
}

type CheckCustomEventCreateInput struct {
	CheckCreateInput

	Integration      graphql.ID `json:"integrationId"`
	ServiceSelector  string     `json:"serviceSelector"`
	SuccessCondition string     `json:"successCondition"`
	Message          string     `json:"resultMessage,omitempty"`
	PassPending      *bool      `json:"passPending,omitempty"`
}

type CheckCustomEventUpdateInput struct {
	CheckUpdateInput

	ServiceSelector  string      `json:"serviceSelector,omitempty"`
	SuccessCondition string      `json:"successCondition,omitempty"`
	Message          string      `json:"resultMessage,omitempty"`
	PassPending      *bool       `json:"passPending,omitempty"`
	Integration      *graphql.ID `json:"integrationId,omitempty"`
}

type CheckGitBranchProtectionCreateInput struct {
	CheckCreateInput
}

type CheckGitBranchProtectionUpdateInput struct {
	CheckUpdateInput
}

type CheckHasDocumentationCreateInput struct {
	CheckCreateInput

	DocumentType    HasDocumentationTypeEnum    `json:"documentType"`
	DocumentSubtype HasDocumentationSubtypeEnum `json:"documentSubtype"`
}

type CheckHasDocumentationUpdateInput struct {
	CheckUpdateInput

	DocumentType    HasDocumentationTypeEnum    `json:"documentType"`
	DocumentSubtype HasDocumentationSubtypeEnum `json:"documentSubtype"`
}

type CheckHasRecentDeployCreateInput struct {
	CheckCreateInput

	Days int `json:"days"`
}

type CheckHasRecentDeployUpdateInput struct {
	CheckUpdateInput

	Days *int `json:"days,omitempty"`
}

type ManualCheckFrequency struct {
	StartingDate       iso8601.Time       `graphql:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `graphql:"frequencyTimeScale"`
	FrequencyValue     int                `graphql:"frequencyValue"`
}

type ManualCheckFrequencyInput struct {
	StartingDate       iso8601.Time       `json:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `json:"frequencyTimeScale"`
	FrequencyValue     int                `json:"frequencyValue"`
}

func NewManualCheckFrequencyInput(startingDate string, timeScale FrequencyTimeScale, value int) *ManualCheckFrequencyInput {
	return &ManualCheckFrequencyInput{
		StartingDate:       NewISO8601Date(startingDate),
		FrequencyTimeScale: timeScale,
		FrequencyValue:     value,
	}
}

type CheckManualCreateInput struct {
	CheckCreateInput

	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                       `json:"updateRequiresComment"`
}

type CheckManualUpdateInput struct {
	CheckUpdateInput

	UpdateFrequency       *ManualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                       `json:"updateRequiresComment,omitempty"`
}

type CheckRepositoryFileCreateInput struct {
	CheckCreateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
	UseAbsoluteRoot       *bool           `json:"useAbsoluteRoot,omitempty"`
}

type CheckRepositoryFileUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       *bool           `json:"directorySearch,omitempty"`
	Filepaths             []string        `json:"filePaths,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
	UseAbsoluteRoot       *bool           `json:"useAbsoluteRoot,omitempty"`
}

type CheckRepositoryGrepCreateInput struct {
	CheckCreateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryGrepUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       *bool           `json:"directorySearch,omitempty"`
	Filepaths             []string        `json:"filePaths,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryIntegratedCreateInput struct {
	CheckCreateInput
}

type CheckRepositoryIntegratedUpdateInput struct {
	CheckUpdateInput
}

type CheckRepositorySearchCreateInput struct {
	CheckCreateInput

	FileExtensions        []string       `json:"fileExtensions,omitempty"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate"`
}

type CheckRepositorySearchUpdateInput struct {
	CheckUpdateInput

	FileExtensions        []string        `json:"fileExtensions,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckServiceDependencyCreateInput struct {
	CheckCreateInput
}

type CheckServiceDependencyUpdateInput struct {
	CheckUpdateInput
}

type CheckServiceConfigurationCreateInput struct {
	CheckCreateInput
}

type CheckServiceConfigurationUpdateInput struct {
	CheckUpdateInput
}

type CheckServiceOwnershipCreateInput struct {
	CheckCreateInput

	RequireContactMethod *bool                             `json:"requireContactMethod,omitempty"`
	ContactMethod        *ServiceOwnershipCheckContactType `json:"contactMethod,omitempty"`
	TeamTagKey           string                            `json:"tagKey,omitempty"`
	TeamTagPredicate     *PredicateInput                   `json:"tagPredicate,omitempty"`
}

type CheckServiceOwnershipUpdateInput struct {
	CheckUpdateInput

	RequireContactMethod *bool                             `json:"requireContactMethod,omitempty"`
	ContactMethod        *ServiceOwnershipCheckContactType `json:"contactMethod,omitempty"`
	TeamTagKey           string                            `json:"tagKey,omitempty"`
	TeamTagPredicate     *PredicateUpdateInput             `json:"tagPredicate,omitempty"`
}

type CheckServicePropertyCreateInput struct {
	CheckCreateInput

	Property  ServicePropertyTypeEnum `json:"serviceProperty"`
	Predicate *PredicateInput         `json:"propertyValuePredicate,omitempty"`
}

type CheckServicePropertyUpdateInput struct {
	CheckUpdateInput

	Property  ServicePropertyTypeEnum `json:"serviceProperty,omitempty"`
	Predicate *PredicateInput         `json:"propertyValuePredicate,omitempty"`
}

type CheckTagDefinedCreateInput struct {
	CheckCreateInput

	TagKey       string          `json:"tagKey"`
	TagPredicate *PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckTagDefinedUpdateInput struct {
	CheckUpdateInput

	TagKey       string          `json:"tagKey,omitempty"`
	TagPredicate *PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckToolUsageCreateInput struct {
	CheckCreateInput

	ToolCategory         ToolCategory    `json:"toolCategory"`
	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty"`
	ToolUrlPredicate     *PredicateInput `json:"toolUrlPredicate,omitempty"`
	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckToolUsageUpdateInput struct {
	CheckUpdateInput

	ToolCategory         ToolCategory    `json:"toolCategory,omitempty"`
	ToolNamePredicate    *PredicateInput `json:"toolNamePredicate,omitempty"`
	ToolUrlPredicate     *PredicateInput `json:"toolUrlPredicate,omitempty"`
	EnvironmentPredicate *PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckDeleteInput struct {
	Id graphql.ID `json:"id"`
}

// Encompass CheckCreatePayload and CheckUpdatePayload into 1 struct
type CheckResponsePayload struct {
	Check  Check
	Errors []OpsLevelErrors
}

func (conn *CheckConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks(after: $after, first: $first)"`
			}
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Rubric.Checks.PageInfo = conn.PageInfo
	for q.Account.Rubric.Checks.PageInfo.HasNextPage {
		v["after"] = q.Account.Rubric.Checks.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		conn.Nodes = append(conn.Nodes, q.Account.Rubric.Checks.Nodes...)
	}
	return nil
}

func (p *CheckResponsePayload) Mutate(client *Client, m interface{}, v map[string]interface{}) (*Check, error) {
	if err := client.Mutate(m, v); err != nil {
		return nil, err
	}
	return &p.Check, FormatErrors(p.Errors)
}

//#region Create

func (client *Client) CreateCheckAlertSourceUsage(input CheckAlertSourceUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkAlertSourceUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckAlertSourceUsage(input CheckAlertSourceUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkAlertSourceUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckCustomEvent(input CheckCustomEventCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckCustomEvent(input CheckCustomEventUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomEventUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckGitBranchProtection(input CheckGitBranchProtectionCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkGitBranchProtectionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckGitBranchProtection(input CheckGitBranchProtectionUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkGitBranchProtectionUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckHasRecentDeploy(input CheckHasRecentDeployCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasRecentDeployCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckHasRecentDeploy(input CheckHasRecentDeployUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasRecentDeployUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckHasDocumentation(input CheckHasDocumentationCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasDocumentationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckHasDocumentation(input CheckHasDocumentationUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasDocumentationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckManual(input CheckManualCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckManual(input CheckManualUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositoryFile(input CheckRepositoryFileCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckRepositoryFile(input CheckRepositoryFileUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositoryGrep(input CheckRepositoryGrepCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckRepositoryGrep(input CheckRepositoryGrepUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositoryIntegrated(input CheckRepositoryIntegratedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}
func (client *Client) UpdateCheckRepositoryIntegrated(input CheckRepositoryIntegratedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckRepositorySearch(input CheckRepositorySearchCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckRepositorySearch(input CheckRepositorySearchUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceDependency(input CheckServiceDependencyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceDependency(input CheckServiceDependencyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceConfiguration(input CheckServiceConfigurationCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceConfiguration(input CheckServiceConfigurationUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceOwnership(input CheckServiceOwnershipCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceOwnership(input CheckServiceOwnershipUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceOwnershipUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckServiceProperty(input CheckServicePropertyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckServiceProperty(input CheckServicePropertyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckTagDefined(input CheckTagDefinedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckTagDefined(input CheckTagDefinedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) CreateCheckToolUsage(input CheckToolUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckToolUsage(input CheckToolUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

//#endregion

//#region Retrieve

func (client *Client) GetCheck(id graphql.ID) (*Check, error) {
	var q struct {
		Account struct {
			Check Check `graphql:"check(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Check.Id == nil {
		return nil, fmt.Errorf("check with ID '%s' not found", id)
	}
	return &q.Account.Check, nil
}

func (client *Client) ListChecks() ([]Check, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks"`
			}
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Rubric.Checks.Nodes, err
	}
	if err := q.Account.Rubric.Checks.Hydrate(client); err != nil {
		return q.Account.Rubric.Checks.Nodes, err
	}
	return q.Account.Rubric.Checks.Nodes, nil
}

//#endregion

//#region Delete

func (client *Client) DeleteCheck(id graphql.ID) error {
	var m struct {
		Payload IdResponsePayload `graphql:"checkDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CheckDeleteInput{Id: id},
	}
	return m.Payload.Mutate(client, &m, v)
}

//#endregion
