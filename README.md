opslevel-go
===========

Package `opslevel` provides an OpsLevel API client implementation.

*NOTE: this library is still a WIP and does not match the API 100% yet*

# Installation

`opslevel` requires Go version 1.8 or later.

```bash
go get -u github.com/opslevel/opslevel-go
```

# Usage

Construct a client, specifying the API token. Then, you can use it to make GraphQL queries and mutations.

```Go
client := opslevel.NewClient("XXX_API_TOKEN_XXX")
// Use client...
```

You can validate the client can successfully talk to the OpsLevel API.

```go
client := opslevel.NewClient("XXX_API_TOKEN_XXX")
if err := client.Validate() {
	panic(err)
}
```

Every resource (IE: service, lifecycle, tier, etc) in OpsLevel API has a corresponding data structure in go as well as the graphql query & mutation inputs.  Additionally there is also some helper functions that use native go types like `string` and `[]string` to make it easier to work with.  The following are a handful of examples:

Find a service given an alias and print the owning team name:

```go
foundService, foundServiceErr := client.GetServiceWithAlias("MyCoolService")
if foundServiceErr != nil {
	panic(foundServiceErr)
}
fmt.Println(foundService.Owner.Name)
```

Create a new service in OpsLevel and print the ID:

```go
serviceCreateInput := opslevel.ServiceCreateInput{
	Name:        "MyCoolService",
	Product:     "MyProduct",
	Description: "The Coolest Service",
	Language:    "go",
}
newService, newServiceErr := client.CreateService(serviceCreateInput)
if newServiceErr != nil {
	panic(newServiceErr)
}
fmt.Println(newService.Id)
```

Assign the tag `{"hello": "world"}` to our newly created service and print all the tags currently on it:

```go
allTagsOnThisService, err := client.AssignTagForId(newService.Id, "Hello", "World")
for tagKey, tagValue := range allTagsOnThisService {
	fmt.Printf("Tag '{%s : %s}'", tagKey, tagValue)
}
```

Build a lookup table of teams:

```go
func GetTeams(client *opslevel.Client) (map[string]opslevel.Team, error) {
	teams := make(map[string]opslevel.Team)
	data, dataErr := client.ListTeams()
	if dataErr != nil {
		return teams, dataErr
	}
	for _, team := range data {
		teams[string(team.Alias)] = team
	}
	return teams, nil
}
```

# Advanced Usage

The client also exposes functions to directly inject custom structures for doing custom query and mutations.  We are running ontop of this [go graphql library](https://github.com/shurcooL/graphql) so you can read up on how to define structures that represent a query or mutation there but here is an example of each:

### Query

```go
var query struct {
	Account struct {
		Tiers []Tier
	}
}
if err := client.Query(&query, nil); err != nil {
	panic(err)
}
for _, tier := range m.Account.Tiers {
	fmt.Println(tier.Name)
}
```

### Mutation

```go
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
```
