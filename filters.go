package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type Predicate struct {
	Type  PredicateTypeEnum `graphql:"type"`
	Value string            `graphql:"value"`
}

type PredicateInput struct {
	Type  PredicateTypeEnum `json:"type"`
	Value string            `json:"value,omitempty"`
}

type PredicateUpdateInput struct {
	Type  PredicateTypeEnum `json:"type,omitempty"`
	Value string            `json:"value,omitempty"`
}

type FilterId struct {
	Id   ID
	Name string
}

type Filter struct {
	Connective ConnectiveEnum
	HtmlURL    string
	FilterId
	Predicates []FilterPredicate
}

type FilterPredicate struct {
	Key     PredicateKeyEnum  `json:"key"`
	KeyData string            `json:"keyData,omitempty"`
	Type    PredicateTypeEnum `json:"type"`
	Value   string            `json:"value,omitempty"`
}

type FilterConnection struct {
	Nodes      []Filter
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type FilterCreateInput struct {
	Name       string            `json:"name"`
	Predicates []FilterPredicate `json:"predicates"`
	Connective ConnectiveEnum    `json:"connective,omitempty"`
}

type FilterUpdateInput struct {
	Id         ID                `json:"id"`
	Name       string            `json:"name,omitempty"`
	Predicates []FilterPredicate `json:"predicates"` //The list of predicates used to select which services apply to the filter. All existing predicates will be replaced by these predicates.
	Connective ConnectiveEnum    `json:"connective,omitempty"`
}

func (self *Filter) Alias() string {
	return slug.Make(self.Name)
}

//#region Create

func (client *Client) CreateFilter(input FilterCreateInput) (*Filter, error) {
	var m struct {
		Payload struct {
			Filter Filter
			Errors []OpsLevelErrors
		} `graphql:"filterCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Filter, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetFilter(id ID) (*Filter, error) {
	var q struct {
		Account struct {
			Filter Filter `graphql:"filter(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v)
	if q.Account.Filter.Id == "" {
		err = fmt.Errorf("Filter with ID '%s' not found!", id)
	}
	return &q.Account.Filter, HandleErrors(err, nil)
}

func (client *Client) ListFilters(variables *PayloadVariables) (FilterConnection, error) {
	var q struct {
		Account struct {
			Filters FilterConnection `graphql:"filters"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables); err != nil {
		return FilterConnection{}, err
	}
	for q.Account.Filters.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Filters.PageInfo.End
		resp, err := client.ListFilters(variables)
		if err != nil {
			return FilterConnection{}, err
		}
		q.Account.Filters.Nodes = append(q.Account.Filters.Nodes, resp.Nodes...)
		q.Account.Filters.PageInfo = resp.PageInfo
	}
	return q.Account.Filters, nil
}

//#endregion

//#region Update

func (client *Client) UpdateFilter(input FilterUpdateInput) (*Filter, error) {
	var m struct {
		Payload struct {
			Filter Filter
			Errors []OpsLevelErrors
		} `graphql:"filterUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v)
	return &m.Payload.Filter, HandleErrors(err, m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteFilter(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedId"`
			Errors []OpsLevelErrors
		} `graphql:"filterDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v)
	return HandleErrors(err, m.Payload.Errors)
}

//#endregion
