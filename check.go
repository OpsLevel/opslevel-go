package opslevel

import (
	"encoding/json"
	"fmt"

	"github.com/relvacode/iso8601"
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
	Id          ID           `graphql:"id"`
	Level       Level        `graphql:"level"`
	Name        string       `graphql:"name"`
	Notes       string       `graphql:"notes: rawNotes"`
	Owner       CheckOwner   `graphql:"owner"`
	Type        CheckType    `graphql:"type"`

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

type CheckInputConstructor func() any

var CheckCreateConstructors = map[CheckType]CheckInputConstructor{
	CheckTypeAlertSourceUsage:    func() any { return &CheckAlertSourceUsageCreateInput{} },
	CheckTypeCustom:              func() any { return &CheckCreateInput{} },
	CheckTypeGeneric:             func() any { return &CheckCustomEventCreateInput{} },
	CheckTypeGitBranchProtection: func() any { return &CheckGitBranchProtectionCreateInput{} },
	CheckTypeHasDocumentation:    func() any { return &CheckHasDocumentationCreateInput{} },
	CheckTypeHasOwner:            func() any { return &CheckServiceOwnershipCreateInput{} },
	CheckTypeHasRecentDeploy:     func() any { return &CheckHasRecentDeployCreateInput{} },
	CheckTypeHasRepository:       func() any { return &CheckRepositoryIntegratedCreateInput{} },
	CheckTypeHasServiceConfig:    func() any { return &CheckServiceConfigurationCreateInput{} },
	CheckTypeManual:              func() any { return &CheckManualCreateInput{} },
	CheckTypePayload:             func() any { return &CheckCreateInput{} },
	CheckTypeRepoFile:            func() any { return &CheckRepositoryFileCreateInput{} },
	CheckTypeRepoGrep:            func() any { return &CheckRepositoryGrepCreateInput{} },
	CheckTypeRepoSearch:          func() any { return &CheckRepositorySearchCreateInput{} },
	CheckTypeServiceDependency:   func() any { return &CheckServiceDependencyCreateInput{} },
	CheckTypeServiceProperty:     func() any { return &CheckServicePropertyCreateInput{} },
	CheckTypeTagDefined:          func() any { return &CheckTagDefinedCreateInput{} },
	CheckTypeToolUsage:           func() any { return &CheckToolUsageCreateInput{} },
}

var CheckUpdateConstructors = map[CheckType]CheckInputConstructor{
	CheckTypeAlertSourceUsage:    func() any { return &CheckAlertSourceUsageUpdateInput{} },
	CheckTypeCustom:              func() any { return &CheckUpdateInput{} },
	CheckTypeGeneric:             func() any { return &CheckCustomEventUpdateInput{} },
	CheckTypeGitBranchProtection: func() any { return &CheckGitBranchProtectionUpdateInput{} },
	CheckTypeHasDocumentation:    func() any { return &CheckHasDocumentationUpdateInput{} },
	CheckTypeHasOwner:            func() any { return &CheckServiceOwnershipUpdateInput{} },
	CheckTypeHasRecentDeploy:     func() any { return &CheckHasRecentDeployUpdateInput{} },
	CheckTypeHasRepository:       func() any { return &CheckRepositoryIntegratedUpdateInput{} },
	CheckTypeHasServiceConfig:    func() any { return &CheckServiceConfigurationUpdateInput{} },
	CheckTypeManual:              func() any { return &CheckManualUpdateInput{} },
	CheckTypePayload:             func() any { return &CheckUpdateInput{} },
	CheckTypeRepoFile:            func() any { return &CheckRepositoryFileUpdateInput{} },
	CheckTypeRepoGrep:            func() any { return &CheckRepositoryGrepUpdateInput{} },
	CheckTypeRepoSearch:          func() any { return &CheckRepositorySearchUpdateInput{} },
	CheckTypeServiceDependency:   func() any { return &CheckServiceDependencyUpdateInput{} },
	CheckTypeServiceProperty:     func() any { return &CheckServicePropertyUpdateInput{} },
	CheckTypeTagDefined:          func() any { return &CheckTagDefinedUpdateInput{} },
	CheckTypeToolUsage:           func() any { return &CheckToolUsageUpdateInput{} },
}

func UnmarshalCheckCreateInput(checkType CheckType, data []byte) (any, error) {
	output := CheckCreateConstructors[checkType]()
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output, nil
}

func UnmarshalCheckUpdateInput(checkType CheckType, data []byte) (any, error) {
	output := CheckUpdateConstructors[checkType]()
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output, nil
}

type HasRecentDeployCheckFragment struct {
	Days int `graphql:"days"`
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
	TotalCount int
}

type CheckCreateInputProvider interface {
	GetCheckCreateInput() *CheckCreateInput
}

type CheckCreateInput struct {
	Name     string        `json:"name"`
	Enabled  bool          `json:"enabled"`
	EnableOn *iso8601.Time `json:"enableOn,omitempty"`
	Category ID            `json:"categoryId"`
	Level    ID            `json:"levelId"`
	Owner    *ID           `json:"ownerId,omitempty"`
	Filter   *ID           `json:"filterId,omitempty"`
	Notes    string        `json:"notes"`
}

func (c *CheckCreateInput) GetCheckCreateInput() *CheckCreateInput {
	return c
}

type CheckUpdateInputProvider interface {
	GetCheckUpdateInput() *CheckUpdateInput
}

type CheckUpdateInput struct {
	Id       ID            `json:"id"`
	Name     string        `json:"name,omitempty"`
	Enabled  *bool         `json:"enabled,omitempty"`
	EnableOn *iso8601.Time `json:"enableOn,omitempty"`
	Category ID            `json:"categoryId,omitempty"`
	Level    ID            `json:"levelId,omitempty"`
	Owner    *ID           `json:"ownerId,omitempty"`
	Filter   *ID           `json:"filterId,omitempty"`
	Notes    *string       `json:"notes,omitempty"`
}

func (c *CheckUpdateInput) GetCheckUpdateInput() *CheckUpdateInput {
	return c
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
	UseAbsoluteRoot       bool            `json:"useAbsoluteRoot"`
}

type CheckRepositoryFileUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths,omitempty"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
	UseAbsoluteRoot       bool            `json:"useAbsoluteRoot"`
}

type CheckRepositoryGrepCreateInput struct {
	CheckCreateInput

	DirectorySearch       bool            `json:"directorySearch"`
	Filepaths             []string        `json:"filePaths"`
	FileContentsPredicate *PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryGrepUpdateInput struct {
	CheckUpdateInput

	DirectorySearch       bool            `json:"directorySearch"`
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
	Id ID `json:"id"`
}

// Encompass CheckCreatePayload and CheckUpdatePayload into 1 struct
type CheckResponsePayload struct {
	Check  Check
	Errors []OpsLevelErrors
}

//#region Create

func (client *Client) CreateCheckHasRecentDeploy(input CheckHasRecentDeployCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasRecentDeployCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckHasRecentDeployCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckHasRecentDeploy(input CheckHasRecentDeployUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkHasRecentDeployUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckHasRecentDeployUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckManual(input CheckManualCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckManualCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckManual(input CheckManualUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkManualUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckManualUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckRepositoryFile(input CheckRepositoryFileCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryFileCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositoryFile(input CheckRepositoryFileUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryFileUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckRepositoryGrep(input CheckRepositoryGrepCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryGrepCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositoryGrep(input CheckRepositoryGrepUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryGrepUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryGrepUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckRepositoryIntegrated(input CheckRepositoryIntegratedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryIntegratedCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositoryIntegrated(input CheckRepositoryIntegratedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryIntegratedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositoryIntegratedUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckRepositorySearch(input CheckRepositorySearchCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositorySearchCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckRepositorySearch(input CheckRepositorySearchUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckRepositorySearchUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckServiceDependency(input CheckServiceDependencyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceDependencyCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceDependency(input CheckServiceDependencyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceDependencyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceDependencyUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckServiceConfiguration(input CheckServiceConfigurationCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceConfigurationCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceConfiguration(input CheckServiceConfigurationUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServiceConfigurationUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServiceConfigurationUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckServiceProperty(input CheckServicePropertyCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServicePropertyCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckServiceProperty(input CheckServicePropertyUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkServicePropertyUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckServicePropertyUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckTagDefined(input CheckTagDefinedCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckTagDefinedCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckTagDefined(input CheckTagDefinedUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkTagDefinedUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckTagDefinedUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) CreateCheckToolUsage(input CheckToolUsageCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckToolUsageCreate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) UpdateCheckToolUsage(input CheckToolUsageUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkToolUsageUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CheckToolUsageUpdate"))
	return &m.Payload.Check, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetCheck(id ID) (*Check, error) {
	var q struct {
		Account struct {
			Check Check `graphql:"check(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("CheckGet"))
	if q.Account.Check.Id == "" {
		err = fmt.Errorf("check with ID '%s' not found", id)
	}
	return &q.Account.Check, HandleErrors(err, nil)
}

func (client *Client) ListChecks(variables *PayloadVariables) (CheckConnection, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks(after: $after, first: $first)"`
			}
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("CheckList")); err != nil {
		return CheckConnection{}, err
	}
	for q.Account.Rubric.Checks.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Rubric.Checks.PageInfo.End
		resp, err := client.ListChecks(variables)
		if err != nil {
			return CheckConnection{}, err
		}
		q.Account.Rubric.Checks.Nodes = append(q.Account.Rubric.Checks.Nodes, resp.Nodes...)
		q.Account.Rubric.Checks.PageInfo = resp.PageInfo
		q.Account.Rubric.Checks.TotalCount += resp.TotalCount
	}
	return q.Account.Rubric.Checks, nil
}

//#endregion

//#region Delete

func (client *Client) DeleteCheck(id ID) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors
		} `graphql:"checkDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CheckDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("CheckDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
