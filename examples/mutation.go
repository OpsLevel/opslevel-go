package opslevel_example

import (
	"github.com/shurcooL/graphql"
)

var mutation struct {
	Payload struct {
		Aliases []graphql.String
		OwnerId graphql.String
		Errors  []opslevel.OpsLevelErrors
	} `graphql:"aliasCreate(input: $input)"`
}
variables := PayloadVariables{
	"input": opslevel.AliasCreateInput{
		Alias:   "MyNewAlias",
		OwnerId: "XXXXXXXXXXX",
	},
}
if err := client.Mutate(&mutation, variables); err != nil {
	panic(err)
}
for _, alias := range m.Payload.Aliases {
	fmt.Println(alias)
}