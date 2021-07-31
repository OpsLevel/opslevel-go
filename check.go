package opslevel

import (
	"fmt"

	"github.com/relvacode/iso8601"
	"github.com/shurcooL/graphql"
)

type CheckType string

const (
	CheckTypeHasOwner         CheckType = "has_owner"
	CheckTypeServiceProperty  CheckType = "service_property"
	CheckTypeHasServiceConfig CheckType = "has_service_config"
	CheckTypeHasRepository    CheckType = "has_repository"
	CheckTypeToolUsage        CheckType = "tool_usage"
	CheckTypeTagDefined       CheckType = "tag_defined"
	CheckTypeRepoFile         CheckType = "repo_file"
	CheckTypeRepoSearch       CheckType = "repo_search"
	CheckTypeCustom           CheckType = "custom"
	CheckTypePayload          CheckType = "payload"
	CheckTypeManual           CheckType = "manual"
	CheckTypeGeneric          CheckType = "generic"
)

type Check struct {
	Category    Category    `json:"category"`
	Checklist   Checklist   `json:"checklist"`
	Description string      `json:"description"`
	Enabled     bool        `json:"enabled"`
	Filter      Filter      `json:"filter"`
	Id          graphql.ID  `json:"id"`
	Integration Integration `json:"integration"`
	Level       Level       `json:"level"`
	Name        string      `json:"name"`
	Notes       string      `json:"notes"`
	//Owner   Team or User - need to look into Fragments
}

type Checklist struct {
	Id   graphql.ID `json:"id"`
	Name string     `json:"name"`
	Path string     `json:"path"`
}

type CheckCustomCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckCustomUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckCustomEventCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter           graphql.ID `json:"filterId,omitempty"`
	ServiceSelector  string     `json:"serviceSelector"`
	SuccessCondition string     `json:"successCondition"`
	Message          string     `json:"resultMessage,omitempty"`
	Integration      graphql.ID `json:"integrationId"`
}

type CheckCustomEventUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter           graphql.ID `json:"filterId,omitempty"`
	ServiceSelector  string     `json:"serviceSelector"`
	SuccessCondition string     `json:"successCondition"`
	Message          string     `json:"resultMessage,omitempty"`
	Integration      graphql.ID `json:"integrationId"`
}

// FrequencyTimeScale represents the time scale type for the frequency.
type FrequencyTimeScale string

// The time scale type for the frequency.
const (
	FrequencyTimeScaleDay   FrequencyTimeScale = "day"   // Consider the time scale of days.
	FrequencyTimeScaleWeek  FrequencyTimeScale = "week"  // Consider the time scale of weeks.
	FrequencyTimeScaleMonth FrequencyTimeScale = "month" // Consider the time scale of months.
	FrequencyTimeScaleYear  FrequencyTimeScale = "year"  // Consider the time scale of years.
)

type manualCheckFrequencyInput struct {
	StartingDate       iso8601.Time       `json:"startingDate"`
	FrequencyTimeScale FrequencyTimeScale `json:"frequencyTimeScale"`
	FrequencyValue     int                `json:"frequencyValue"`
}

func NewManualCheckFrequencyInput(startingDate string, timeScale FrequencyTimeScale, value int) manualCheckFrequencyInput {
	return manualCheckFrequencyInput{
		StartingDate:       NewISO8601Date(startingDate),
		FrequencyTimeScale: timeScale,
		FrequencyValue:     value,
	}
}

type CheckManualCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID                `json:"filterId,omitempty"`
	UpdateFrequency       manualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                      `json:"updateRequiresComment"`
}

type CheckManualUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID                `json:"filterId,omitempty"`
	UpdateFrequency       manualCheckFrequencyInput `json:"updateFrequency,omitempty"`
	UpdateRequiresComment bool                      `json:"updateRequiresComment"`
}

type CheckPayloadCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter       graphql.ID `json:"filterId,omitempty"`
	JQExpression string     `json:"jqExpression"`
	Message      string     `json:"resultMessage,omitempty"`
}

type CheckPayloadUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter       graphql.ID `json:"filterId,omitempty"`
	JQExpression string     `json:"jqExpression"`
	Message      string     `json:"resultMessage,omitempty"`
}

type CheckRepositoryFileCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID     `json:"filterId,omitempty"`
	DirectorySearch       bool           `json:"directorySearch"`
	Filepaths             []string       `json:"filePaths"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryFileUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID     `json:"filterId,omitempty"`
	DirectorySearch       bool           `json:"directorySearch"`
	Filepaths             []string       `json:"filePaths"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate,omitempty"`
}

type CheckRepositoryIntegratedCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckRepositoryIntegratedUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckRepositorySearchCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID     `json:"filterId,omitempty"`
	FileExtensions        []string       `json:"fileExtensions,omitempty"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate"`
}

type CheckRepositorySearchUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter                graphql.ID     `json:"filterId,omitempty"`
	FileExtensions        []string       `json:"fileExtensions,omitempty"`
	FileContentsPredicate PredicateInput `json:"fileContentsPredicate"`
}

type CheckServiceConfigurationCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckServiceConfigurationUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckServiceOwnershipCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type CheckServiceOwnershipUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter graphql.ID `json:"filterId,omitempty"`
}

type ServiceProperty string

const (
	ServicePropertyDescription ServiceProperty = "description"
	ServicePropertyName        ServiceProperty = "name"
	ServicePropertyLanguage    ServiceProperty = "language"
	ServicePropertyFramework   ServiceProperty = "framework"
	ServicePropertyProduct     ServiceProperty = "product"
	ServicePropertyLifecycle   ServiceProperty = "lifecycle_index"
	ServicePropertyTier        ServiceProperty = "tier_index"
)

type CheckServicePropertyCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter    graphql.ID      `json:"filterId,omitempty"`
	Property  ServiceProperty `json:"serviceProperty"`
	Predicate PredicateInput  `json:"propertyValuePredicate,omitempty"`
}

type CheckServicePropertyUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter    graphql.ID      `json:"filterId,omitempty"`
	Property  ServiceProperty `json:"serviceProperty"`
	Predicate PredicateInput  `json:"propertyValuePredicate,omitempty"`
}

type CheckTagDefinedCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter       graphql.ID     `json:"filterId,omitempty"`
	TagKey       string         `json:"tagKey"`
	TagPredicate PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckTagDefinedUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter       graphql.ID     `json:"filterId,omitempty"`
	TagKey       string         `json:"tagKey"`
	TagPredicate PredicateInput `json:"tagPredicate,omitempty"`
}

type CheckToolUsageCreateInput struct {
	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter               graphql.ID     `json:"filterId,omitempty"`
	ToolCategory         ToolCategory   `json:"toolCategory"`
	ToolNamePredicate    PredicateInput `json:"toolNamePredicate,omitempty"`
	EnvironmentPredicate PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckToolUsageUpdateInput struct {
	// ID
	Id graphql.ID `json:"id"`

	// Base
	Name     string     `json:"name"`
	Enabled  bool       `json:"enabled"`
	Category graphql.ID `json:"categoryId"`
	Level    graphql.ID `json:"levelId"`
	Owner    graphql.ID `json:"ownerId,omitempty"`
	Notes    string     `json:"notes,omitempty"`

	// Specific
	Filter               graphql.ID     `json:"filterId,omitempty"`
	ToolCategory         ToolCategory   `json:"toolCategory"`
	ToolNamePredicate    PredicateInput `json:"toolNamePredicate,omitempty"`
	EnvironmentPredicate PredicateInput `json:"environmentPredicate,omitempty"`
}

type CheckDeleteInput struct {
	Id graphql.ID `json:"id"`
}

type CheckResponsePayload struct {
	Check  Check
	Errors []OpsLevelErrors
}

func (p *CheckResponsePayload) Mutate(client *Client, m interface{}, v map[string]interface{}) (*Check, error) {
	if err := client.Mutate(m, v); err != nil {
		return nil, err
	}
	return &p.Check, FormatErrors(p.Errors)
}

//#region Create

func (client *Client) CreateCheckCustom(input CheckCustomCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckCustom(input CheckCustomUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkCustomUpdate(input: $input)"`
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

func (client *Client) CreateCheckPayload(input CheckPayloadCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkPayloadCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	return m.Payload.Mutate(client, &m, v)
}

func (client *Client) UpdateCheckPayload(input CheckPayloadUpdateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkPayloadUpdate(input: $input)"`
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
		return nil, fmt.Errorf("Check with ID '%s' not found!", id)
	}
	return &q.Account.Check, nil
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
