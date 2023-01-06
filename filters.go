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
	Id   graphql.ID
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
	Id         graphql.ID        `json:"id"`
	Name       string            `json:"name,omitempty"`
	Predicates []FilterPredicate `json:"predicates"` //The list of predicates used to select which services apply to the filter. All existing predicates will be replaced by these predicates.
	Connective ConnectiveEnum    `json:"connective,omitempty"`
}

func (self *Filter) Alias() string {
	return slug.Make(self.Name)
}

func (conn *FilterConnection) Hydrate(client *Client) error {
	var q struct {
		Account struct {
			Filters FilterConnection `graphql:"filters(after: $after, first: $first)"`
		}
	}
	v := PayloadVariables{
		"first": client.pageSize,
	}
	q.Account.Filters.PageInfo = conn.PageInfo
	for q.Account.Filters.PageInfo.HasNextPage {
		v["after"] = q.Account.Filters.PageInfo.End
		if err := client.Query(&q, v); err != nil {
			return err
		}
		for _, item := range q.Account.Filters.Nodes {
			conn.Nodes = append(conn.Nodes, item)
		}
	}
	return nil
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
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Filter, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Retrieve

func (client *Client) GetFilter(id graphql.ID) (*Filter, error) {
	var q struct {
		Account struct {
			Filter Filter `graphql:"filter(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	if q.Account.Filter.Id == "" {
		return nil, fmt.Errorf("Filter with ID '%s' not found!", id)
	}
	return &q.Account.Filter, nil
}

func (client *Client) ListFilters() ([]Filter, error) {
	var q struct {
		Account struct {
			Filters FilterConnection `graphql:"filters"`
		}
	}
	if err := client.Query(&q, nil); err != nil {
		return q.Account.Filters.Nodes, err
	}
	if err := q.Account.Filters.Hydrate(client); err != nil {
		return q.Account.Filters.Nodes, err
	}
	return q.Account.Filters.Nodes, nil
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
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	return &m.Payload.Filter, FormatErrors(m.Payload.Errors)
}

//#endregion

//#region Delete

func (client *Client) DeleteFilter(id graphql.ID) error {
	var m struct {
		Payload struct {
			Id     graphql.ID `graphql:"deletedId"`
			Errors []OpsLevelErrors
		} `graphql:"filterDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: graphql.ID(id)},
	}
	if err := client.Mutate(&m, v); err != nil {
		return err
	}
	return FormatErrors(m.Payload.Errors)
}

//#endregion
