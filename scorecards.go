package opslevel

import (
	"errors"
	"fmt"
	"slices"
)

type ScorecardId struct {
	Aliases []string `graphql:"aliases"`
	Id      ID       `graphql:"id"`
}

type Scorecard struct {
	ScorecardId

	AffectsOverallServiceLevels bool        `graphql:"affectsOverallServiceLevels"`
	Description                 string      `graphql:"description"` // optional
	Filter                      Filter      `graphql:"filter"`      // optional
	Name                        string      `graphql:"name"`
	Owner                       EntityOwner `graphql:"owner"`
	PassingChecks               int         `graphql:"passingChecks"`
	ServiceCount                int         `graphql:"serviceCount"`
	ChecksCount                 int         `graphql:"totalChecks"`
}

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

func (scorecard *ScorecardId) ReconcileAliases(client *Client, aliasesWanted []string) error {
	var allErrors, err error

	aliasesToCreate, aliasesToDelete := ExtractAliases(scorecard.Aliases, aliasesWanted)
	for _, alias := range aliasesToDelete {
		err := client.DeleteAlias(AliasDeleteInput{
			Alias:     alias,
			OwnerType: AliasOwnerTypeEnumScorecard,
		})
		allErrors = errors.Join(allErrors, err)
	}

	if len(aliasesToCreate) > 0 {
		// CreateAliases returns current list of aliases from owned by Scorecard
		scorecard.Aliases, err = client.CreateAliases(scorecard.Id, aliasesToCreate)
		allErrors = errors.Join(allErrors, err)
	} else {
		scorecard.Aliases = slices.DeleteFunc(scorecard.Aliases, func(alias string) bool {
			return slices.Contains(aliasesToDelete, alias)
		})
	}

	return allErrors
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
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	(*variables)["scorecard"] = *NewIdentifier(string(scorecard.Id))
	if err := client.Query(&q, *variables, WithName("ScorecardCategoryList")); err != nil {
		return nil, err
	}

	for q.Account.Scorecard.Categories.PageInfo.HasNextPage {
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
		Payload struct {
			Scorecard Scorecard        `graphql:"scorecard"`
			Errors    []OpsLevelErrors `graphql:"errors"`
		} `graphql:"scorecardCreate(input: $input)"`
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
		err = fmt.Errorf("scorecard with ID or Alias matching '%s' not found", input)
	}
	return &q.Account.Scorecard, HandleErrors(err, nil)
}

func (client *Client) ListScorecards(variables *PayloadVariables) (*ScorecardConnection, error) {
	var q struct {
		Account struct {
			Scorecards ScorecardConnection `graphql:"scorecards(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("ScorecardsList")); err != nil {
		return nil, err
	}
	for q.Account.Scorecards.PageInfo.HasNextPage {
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
		Payload struct {
			Scorecard Scorecard        `graphql:"scorecard"`
			Errors    []OpsLevelErrors `graphql:"errors"`
		} `graphql:"scorecardUpdate(scorecard: $scorecard, input: $input)"`
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
		Payload struct {
			DeletedScorecardId ID               `graphql:"deletedScorecardId"`
			Errors             []OpsLevelErrors `graphql:"errors"`
		} `graphql:"scorecardDelete(input: $input)"`
	}
	input := *NewIdentifier(identifier)
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("ScorecardDelete"))
	return &m.Payload.DeletedScorecardId, HandleErrors(err, m.Payload.Errors)
}
