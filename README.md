<img align="right" src="https://user-images.githubusercontent.com/735015/118144171-c2a40d00-b3d1-11eb-9b8c-a1cdd687cb36.png" width="320" height="160">


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

Construct a client, specifying the [API token](https://app.opslevel.com/api_tokens). Then, you can use it to make GraphQL queries and mutations.

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
if foundServiceErr != nil {![OpsLevel Golang Gopher NoBG](https://user-images.githubusercontent.com/735015/118144162-c041b300-b3d1-11eb-8eb2-03d01e7a3fc7.png)

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

List all the tags for a service:

```go
tags, tagsErr := client.GetTagsForServiceWithAlias("MyCoolService")
for _, tag := range tags {
	fmt.Printf("Tag '{%s : %s}'\n", tag.Key, tag.Value)
}
// OR
service, serviceErr := client.GetServiceWithAlias("MyCoolService")
tags, tagsErr := client.GetTagsForServiceWithId(service.Id)
for _, tag := range tags {
	fmt.Printf("Tag '{%s : %s}'\n", tag.Key, tag.Value)
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

The client also exposes functions `Query` and `Mutate` for doing custom query or mutations.  We are running ontop of this [go graphql library](https://github.com/shurcooL/graphql) so you can read up on how to define go structures that represent a query or mutation there but here is an example of each:

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
