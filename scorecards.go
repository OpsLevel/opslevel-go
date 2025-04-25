package opslevel

import (
	"errors"
	"fmt"

	"github.com/hasura/go-graphql-client"
)

type ScorecardConnection struct {
	Nodes      []Scorecard `graphql:"nodes"`
	PageInfo   PageInfo    `graphql:"pageInfo"`
	TotalCount int         `graphql:"totalCount"`
}

type ScorecardCategoryConnection struct {
	Nodes      []Category `graphql:"nodes"`
	PageInfo   PageInfo   `graphql:"pageInfo"`
	TotalCount int        `graphql:"totalCount"`
}

func (scorecard *ScorecardId) ResourceId() ID {
	return scorecard.Id
}

func (scorecard *ScorecardId) AliasableType() AliasOwnerTypeEnum {
	return AliasOwnerTypeEnumScorecard
}

func (scorecard *ScorecardId) GetAliases() []string {
	return scorecard.Aliases
}

func (scorecard *ScorecardId) ReconcileAliases(client *Client, aliasesWanted []string) error {
	aliasesToCreate, aliasesToDelete := extractAliases(scorecard.Aliases, aliasesWanted)

	// reconcile wanted aliases with actual aliases
	deleteErr := client.DeleteAliases(AliasOwnerTypeEnumScorecard, aliasesToDelete)
	_, createErr := client.CreateAliases(scorecard.Id, aliasesToCreate)

	// update scorecard to reflect API updates
	updatedScorecard, getErr := client.GetScorecard(string(scorecard.Id))
	if updatedScorecard != nil {
		scorecard.Aliases = updatedScorecard.Aliases
	}

	return errors.Join(deleteErr, createErr, getErr)
}

func (scorecard *Scorecard) ListCategories(client *Client, variables *PayloadVariables) (*ScorecardCategoryConnection, error) {
	if scorecard.Id == "" {
		return nil, fmt.Errorf("unable to get categories, invalid scorecard id: '%s'", scorecard.Id)
	}
	var q struct {
		Account struct {
			Scorecard struct {
				Categories ScorecardCategoryConnection `graphql:"categories(after: $after, first: $first)"`
			} `graphql:"scorecard(input: $scorecard)"`
		}
	}

	variables = client.PopulatePaginationParams(variables)
	(*variables)["scorecard"] = *NewIdentifier(string(scorecard.Id))
	if err := client.Query(&q, *variables, WithName("ScorecardCategoryList")); err != nil {
		return nil, err
	}

	if q.Account.Scorecard.Categories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Scorecard.Categories.PageInfo.End
		resp, err := scorecard.ListCategories(client, variables)
		if err != nil {
			return nil, err
		}
		q.Account.Scorecard.Categories.Nodes = append(q.Account.Scorecard.Categories.Nodes, resp.Nodes...)
		q.Account.Scorecard.Categories.PageInfo = resp.PageInfo
		q.Account.Scorecard.Categories.TotalCount += resp.TotalCount
	}

	return &q.Account.Scorecard.Categories, nil
}

func (client *Client) CreateScorecard(input ScorecardInput) (*Scorecard, error) {
	var m struct {
		Payload ScorecardPayload `graphql:"scorecardCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ScorecardCreate"))
	return &m.Payload.Scorecard, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetScorecard(input string) (*Scorecard, error) {
	var q struct {
		Account struct {
			Scorecard Scorecard `graphql:"scorecard(input: $input)"`
		}
	}
	v := PayloadVariables{
		"input": *NewIdentifier(input),
	}
	err := client.Query(&q, v, WithName("ScorecardGet"))
	if q.Account.Scorecard.Id == "" {
		err = graphql.Errors{graphql.Error{
			Message: fmt.Sprintf("scorecard with ID or Alias matching '%s' not found", input),
			Path:    []any{"account", "scorecard"},
		}}
	}
	return &q.Account.Scorecard, HandleErrors(err, nil)
}

func (client *Client) ListScorecards(variables *PayloadVariables) (*ScorecardConnection, error) {
	var q struct {
		Account struct {
			Scorecards ScorecardConnection `graphql:"scorecards(after: $after, first: $first)"`
		}
	}

	variables = client.PopulatePaginationParams(variables)
	if err := client.Query(&q, *variables, WithName("ScorecardsList")); err != nil {
		return nil, err
	}
	if q.Account.Scorecards.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Scorecards.PageInfo.End
		resp, err := client.ListScorecards(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Scorecards.Nodes = append(q.Account.Scorecards.Nodes, resp.Nodes...)
		q.Account.Scorecards.PageInfo = resp.PageInfo
		q.Account.Scorecards.TotalCount = len(q.Account.Scorecards.Nodes)
	}
	return &q.Account.Scorecards, nil
}

func (client *Client) UpdateScorecard(identifier string, input ScorecardInput) (*Scorecard, error) {
	var m struct {
		Payload ScorecardPayload `graphql:"scorecardUpdate(scorecard: $scorecard, input: $input)"`
	}
	scorecard := *NewIdentifier(identifier)
	v := PayloadVariables{
		"scorecard": scorecard,
		"input":     input,
	}
	err := client.Mutate(&m, v, WithName("ScorecardUpdate"))
	return &m.Payload.Scorecard, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteScorecard(identifier string) (*ID, error) {
	var m struct {
		Payload struct { // TODO: need to fix this
			DeletedScorecardId ID      `graphql:"deletedScorecardId"`
			Errors             []Error `graphql:"errors"`
		} `graphql:"scorecardDelete(input: $input)"`
	}
	input := *NewIdentifier(identifier)
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ScorecardDelete"))
	return &m.Payload.DeletedScorecardId, HandleErrors(err, m.Payload.Errors)
}
