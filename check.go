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
