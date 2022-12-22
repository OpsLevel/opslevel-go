package opslevel

import (
	"fmt"
	"github.com/shurcooL/graphql"
	"sort"
	"strconv"
	"strings"
)

type CustomActionsId struct {
	Aliases []string   `graphql:"aliases"`
	Id      graphql.ID `graphql:"id"`
}

type CustomActionsExternalAction struct {
	CustomActionsId

	Description    string `graphql:"description"`
	LiquidTemplate string `graphql:"liquidTemplate"`
	Name           string `graphql:"name"`

	CustomActionsWebhookAction `graphql:"... on CustomActionsWebhookAction"`
}

type CustomActionsWebhookAction struct {
	BasicAuthUserName string                      `graphql:"basicAuthUserName"`
	BasicAuthPassword string                      `graphql:"basicAuthPassword"`
	HeadersString     string                      `graphql:"headers"`
	HTTPMethod        CustomActionsHttpMethodEnum `graphql:"httpMethod"`
	WebhookURL        string                      `graphql:"webhookUrl"`
}

type CustomActionsTriggerDefinition struct {
	Action      CustomActionsId `graphql: "action"`
	Aliases     []string        `graphql:"aliases"`
	Description string          `graphql:"description"`
	Filter      FilterId        `graphql:"filter"`
	Id          graphql.ID      `graphql:"id"`
	Name        string          `graphql:"name"`
	Owner       TeamId          `graphql:"owner"`
	Timestamps  Timestamps      `graphql:"timestamps"`
}

type CustomActionsExternalActionsConnection struct {
	Nodes      []CustomActionsExternalAction
	PageInfo   PageInfo
	TotalCount int
}

type CustomActionsTriggerDefinitionsConnection struct {
	Nodes      []CustomActionsTriggerDefinition
	PageInfo   PageInfo
	TotalCount int
}

type CustomActionsWebhookActionCreateInput struct {
	Name              string                      `json:"name"`
	Description       *string                     `json:"description,omitempty"`
	LiquidTemplate    string                      `json:"liquidTemplate"`
	WebhookURL        string                      `json:"webhookUrl"`
	HTTPMethod        CustomActionsHttpMethodEnum `json:"httpMethod"`
	BasicAuthUserName *string                     `json:"basicAuthUserName,omitempty"`
	BasicAuthPassword *string                     `json:"basicAuthPassword,omitempty"`
	Headers           map[string]string           `json:"-"`
	HeadersString     *string                     `json:"headers,omitempty"`
}

type CustomActionsTriggerDefinitionCreateInput struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Owner       graphql.ID  `json:"ownerId"`
	Action      *graphql.ID `json:"actionId,omitempty"`
	Filter      *graphql.ID `json:"filterId,omitempty"`
	// This is being explictly left out to reduce the complexity of the implementation
	// action *CustomActionsWebhookActionCreateInput
}

func (s *CustomActionsWebhookAction) Headers() map[string]string {
	return CustomActionsToHeaders(s.HeadersString)
}

func CustomActionsToHeaders(input string) map[string]string {
	output := map[string]string{}
	headers := strings.Split(input, "\n")
	for _, header := range headers {
		parts := strings.Split(header, "=")
		// TODO: handle parts split error / len != 2
		s, _ := strconv.Unquote(parts[1])
		// TODO: handle Unquote error
		output[parts[0]] = s
	}
	return output
}

func CustomActionsFromHeaders(input map[string]string) string {
	// To ensure consistent output we use a sorted list of keys
	// This way the string of headers always appears in alphabetical order
	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s=\"%s\"\n", k, input[k]))
	}
	return strings.TrimSuffix(builder.String(), "\n")
}

func (client *Client) CreateWebhookAction(input CustomActionsWebhookActionCreateInput) (*CustomActionsExternalAction, error) {
	var m struct {
		Payload struct {
			WebhookAction CustomActionsExternalAction
			Errors        []OpsLevelErrors
		} `graphql:"customActionsWebhookActionCreate(input: $input)"`
	}
	marshalledHeaders := CustomActionsFromHeaders(input.Headers)
	input.HeadersString = &marshalledHeaders
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.WebhookAction, FormatErrors(m.Payload.Errors)
}

// TODO: Not implemented in the API yet
//func (client *Client) GetCustomAction(input IdentifierInput) (*CustomActionsTriggerDefinition, error) {

//}

func (client *Client) ListCustomActions(variables *PayloadVariables) (CustomActionsExternalActionsConnection, error) {
	var q struct {
		Account struct {
			Actions CustomActionsExternalActionsConnection `graphql:"customActionsExternalActions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = &PayloadVariables{
			"after": "",
			"first": client.pageSize,
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return CustomActionsExternalActionsConnection{}, err
	}
	for q.Account.Actions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Actions.PageInfo.End
		resp, err := client.ListCustomActions(variables)
		if err != nil {
			return CustomActionsExternalActionsConnection{}, err
		}
		q.Account.Actions.Nodes = append(q.Account.Actions.Nodes, resp.Nodes...)
		q.Account.Actions.PageInfo = resp.PageInfo
	}
	return q.Account.Actions, nil
}

func (client *Client) CreateTriggerDefinition(input CustomActionsTriggerDefinitionCreateInput) (*CustomActionsTriggerDefinition, error) {
	var m struct {
		Payload struct {
			TriggerDefinition CustomActionsTriggerDefinition
			Errors            []OpsLevelErrors
		} `graphql:"customActionsTriggerDefinitionCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.TriggerDefinition, FormatErrors(m.Payload.Errors)
}

func (client *Client) GetTriggerDefinition(input IdentifierInput) (*CustomActionsTriggerDefinition, error) {
	var q struct {
		Account struct {
			Definition CustomActionsTriggerDefinition `graphql:"customActionsTriggerDefinition(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Definition.Id == nil {
		return nil, fmt.Errorf("CustomActionsTriggerDefinition with ID '%s' or Alias '%s' not found", input.Id, input.Alias)
	}
	return &q.Account.Definition, nil
}

func (client *Client) ListTriggerDefinitions(variables *PayloadVariables) (CustomActionsTriggerDefinitionsConnection, error) {
	var q struct {
		Account struct {
			Definitions CustomActionsTriggerDefinitionsConnection `graphql:"customActionsTriggerDefinitions(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = &PayloadVariables{
			"after": "",
			"first": client.pageSize,
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return CustomActionsTriggerDefinitionsConnection{}, err
	}
	for q.Account.Definitions.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Definitions.PageInfo.End
		resp, err := client.ListTriggerDefinitions(variables)
		if err != nil {
			return CustomActionsTriggerDefinitionsConnection{}, err
		}
		q.Account.Definitions.Nodes = append(q.Account.Definitions.Nodes, resp.Nodes...)
		q.Account.Definitions.PageInfo = resp.PageInfo
	}
	return q.Account.Definitions, nil
}
