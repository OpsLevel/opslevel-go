package opslevel

import (
	"fmt"
	"strings"

	"github.com/shurcooL/graphql"
)

type AliasCreateInput struct {
	Alias   string     `json:"alias"`
	OwnerId graphql.ID `json:"ownerId"`
}

//#region Create
// TODO: make sure duplicate aliases throw an error that we can catch
func (client *Client) CreateAliases(ownerId graphql.ID, aliases []string) ([]string, error) {
	var output []string
	var errors []string
	for _, alias := range aliases {
		input := AliasCreateInput{
			Alias:   alias,
			OwnerId: ownerId,
		}
		result, err := client.CreateAlias(input)
		if err != nil {
			errors = append(errors, err.Error())
		}
		for _, resultAlias := range result {
			output = append(output, string(resultAlias))
		}
	}
	output = removeDuplicates(output)
	if len(errors) > 0 {
		return output, fmt.Errorf(strings.Join(errors, "\n"))
	} else {
		return output, nil
	}
}

func (client *Client) CreateAlias(input AliasCreateInput) ([]string, error) {
	var m struct {
		Payload struct {
			Aliases []graphql.String
			OwnerId graphql.String
			Errors  []OpsLevelErrors
		} `graphql:"aliasCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	if err := client.Mutate(&m, v); err != nil {
		return nil, err
	}
	output := make([]string, len(m.Payload.Aliases))
	for i, item := range m.Payload.Aliases {
		output[i] = string(item)
	}
	return output, FormatErrors(m.Payload.Errors)
}

//#endregion
