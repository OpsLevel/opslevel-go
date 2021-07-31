package opslevel

import (
	"fmt"

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

type Integration struct {
	AccountKey  string     `json:"accountKey"`
	AccountName string     `json:"accountName"`
	AccountURL  string     `json:"accountUrl"`
	Id          graphql.ID `json:"id"`
	Name        string     `json:"name"`
	WebhookURL  string     `json:"webhookUrl"`
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

func (client *Client) CreateCheckRepositoryFile(input CheckRepositoryFileCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositoryFileCreate(input: $input)"`
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

func (client *Client) CreateCheckRepositorySearch(input CheckRepositorySearchCreateInput) (*Check, error) {
	var m struct {
		Payload CheckResponsePayload `graphql:"checkRepositorySearchCreate(input: $input)"`
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
