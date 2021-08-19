package opslevel

import (
	"fmt"

	"github.com/shurcooL/graphql"
)

type ConnectiveType string

const (
	ConnectiveTypeAnd ConnectiveType = "and"
	ConnectiveTypeOr  ConnectiveType = "or"
)

func GetConnectiveTypes() []string {
	return []string{
		string(ConnectiveTypeAnd),
		string(ConnectiveTypeOr),
	}
}

type PredicateType string

const (
	PredicateTypeContains                   PredicateType = "contains"
	PredicateTypeDoesNotContain             PredicateType = "does_not_contain"
	PredicateTypeDoesNotEqual               PredicateType = "does_not_equal"
	PredicateTypeDoesNotExist               PredicateType = "does_not_exist"
	PredicateTypeEndsWith                   PredicateType = "ends_with"
	PredicateTypeEquals                     PredicateType = "equals"
	PredicateTypeExists                     PredicateType = "exists"
	PredicateTypeGreaterThanOrEqualTo       PredicateType = "greater_than_or_equal_to"
	PredicateTypeLessThanOrEqualTo          PredicateType = "less_than_or_equal_to"
	PredicateTypeStartsWith                 PredicateType = "starts_with"
	PredicateTypeSatisfiesVersionConstraint PredicateType = "satisfies_version_constraint"
	PredicateTypeMatchesRegex               PredicateType = "matches_regex"
)

func GetPredicateTypes() []string {
	return []string{
		string(PredicateTypeContains),
		string(PredicateTypeDoesNotContain),
		string(PredicateTypeDoesNotEqual),
		string(PredicateTypeDoesNotExist),
		string(PredicateTypeEndsWith),
		string(PredicateTypeEquals),
		string(PredicateTypeExists),
		string(PredicateTypeGreaterThanOrEqualTo),
		string(PredicateTypeLessThanOrEqualTo),
		string(PredicateTypeStartsWith),
		string(PredicateTypeSatisfiesVersionConstraint),
		string(PredicateTypeMatchesRegex),
	}
}

type PredicateInput struct {
	Type  PredicateType `json:"type"`
	Value string        `json:"value,omitempty"`
}

type Filter struct {
	Connective ConnectiveType
	HtmlURL    string
	Id         graphql.ID
	Name       string
	Predicates []FilterPredicate
}

type FilterPredicate struct {
	Key     string        `json:"key"`
	KeyData string        `json:"keyData,omitempty"`
	Type    PredicateType `json:"type"`
	Value   string        `json:"value,omitempty"`
}

type FilterConnection struct {
	Nodes      []Filter
	PageInfo   PageInfo
	TotalCount graphql.Int
}

type FilterCreateInput struct {
	Name       string            `json:"name"`
	Predicates []FilterPredicate `json:"predicates"`
	Connective ConnectiveType    `json:"connective,omitempty"`
}

type FilterUpdateInput struct {
	Id         graphql.ID        `json:"id"`
	Name       string            `json:"name,omitempty"`
	Predicates []FilterPredicate `json:"predicates"` //The list of predicates used to select which services apply to the filter. All existing predicates will be replaced by these predicates.
	Connective ConnectiveType    `json:"connective,omitempty"`
}

type FilterDeleteInput struct {
	Id graphql.ID `json:"id"`
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
	if q.Account.Filter.Id == nil {
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
