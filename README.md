<p align="center">
    <a href="https://github.com/OpsLevel/opslevel-go/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/OpsLevel/opslevel-go.svg" alt="License"/></a>
    <a href="https://go.dev">
        <img src="https://img.shields.io/github/go-mod/go-version/OpsLevel/opslevel-go" alt="Made With Go"/></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-go/releases/">
        <img src="https://img.shields.io/github/v/release/OpsLevel/opslevel-go?include_prereleases" alt="Release"/></a>
    <a href="https://GitHub.com/OpsLevel/opslevel-go/issues/">
        <img src="https://img.shields.io/github/issues/OpsLevel/opslevel-go.svg" alt="Issues"/></a>
    <a href="https://github.com/OpsLevel/opslevel-go/graphs/contributors">
        <img src="https://img.shields.io/github/contributors/OpsLevel/opslevel-go" alt="Contributors"/></a>
    <a href="https://github.com/OpsLevel/opslevel-go/pulse">
        <img src="https://img.shields.io/github/commit-activity/m/OpsLevel/opslevel-go" alt="Activity"/></a>
	<a href="https://codecov.io/gh/OpsLevel/opslevel-go">
        <img src="https://codecov.io/gh/OpsLevel/opslevel-go/branch/main/graph/badge.svg?token=GHQHRIJ9UW" alt="CodeCov"/></a>
    <a href="https://dependabot.com/">
        <img src="https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot" alt="Dependabot"/></a>
    <a href="https://pkg.go.dev/github.com/opslevel/opslevel-go/v2026">
        <img src="https://pkg.go.dev/badge/github.com/opslevel/opslevel.svg" alt="Go Reference"/></a>
</p>

<img align="right" src="logo.png" width="320" height="160" alt="The OpsLevel Gopher">

[![Overall](https://img.shields.io/endpoint?style=flat&url=https%3A%2F%2Fapp.opslevel.com%2Fapi%2Fservice_level%2FOrfRqpiglK-WdxPAHJrUWzwYaweF_gDsmkSKWFYw9LU)](https://app.opslevel.com/services/opslevel_api_clients/maturity-report)

# opslevel-go

Package `opslevel` provides an OpsLevel API client implementation.


# Installation

`opslevel` requires Go version 1.8 or later.

```bash
go get -u github.com/opslevel/opslevel-go/v2026
```

# Usage

Construct a client, specifying the [API token](https://app.opslevel.com/api_tokens). Then, you can use it to make GraphQL queries and mutations.

```go
client := opslevel.NewGQLClient(opslevel.SetAPIToken("XXX_API_TOKEN_XXX"))
// Use client...
```

You can validate the client can successfully talk to the OpsLevel API.

```go
client := opslevel.NewGQLClient(opslevel.SetAPIToken("XXX_API_TOKEN_XXX"))
if err := client.Validate(); err != nil {
	panic(err)
}
```

Every resource (IE: service, lifecycle, tier, etc.) in OpsLevel API has a corresponding data structure in go as well as the graphql query & mutation inputs.  Additionally, there are also some helper functions that use native go types like `string` and `[]string` to make it easier to work with.  The following are a handful of examples:

Find a service given an alias and print the owning team name:

```go
foundService, foundServiceErr := client.GetService("MyCoolService")
if foundServiceErr != nil {
	panic(foundServiceErr)
}
fmt.Println(foundService.Owner.Name)
```

Create a new service in OpsLevel and print the ID:

```go
serviceCreateInput := opslevel.ServiceCreateInput{
	Name:        "MyCoolService",
	Description: opslevel.RefOf("The Coolest Service"),
	Language:    opslevel.RefOf("go"),
	OwnerAlias:  opslevel.RefOf("team-platform"),
}
newService, newServiceErr := client.CreateService(serviceCreateInput)
if newServiceErr != nil {
	panic(newServiceErr)
}
fmt.Println(newService.Id)
```

Assign the tag `{"hello": "world"}` to our newly created service and print all the tags currently on it:

```go
allTagsOnThisService, assignTagsErr := client.AssignTags(string(newService.Id), map[string]string{"hello": "world"})
if assignTagsErr != nil {
	panic(assignTagsErr)
}
for _, tagOnService := range allTagsOnThisService {
	fmt.Printf("Tag '{%s : %s}'", tagOnService.Id, tagOnService.Value)
}
```

List all the tags for a service:

```go
myService, foundServiceErr := client.GetService("MyCoolService")
if foundServiceErr != nil {
	panic(foundServiceErr)
}
tags, getTagsErr := myService.GetTags(client, nil)
if getTagsErr != nil {
	panic(getTagsErr)
}
for _, tag := range tags {
	fmt.Printf("Tag '{%s : %s}'\n", tag.Key, tag.Value)
}
```

Build a lookup table of teams:

```go
func GetTeams(client *opslevel.Client) (map[string]opslevel.Team, error) {
	teams := make(map[string]opslevel.Team)
	data, dataErr := client.ListTeams(nil)
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

The client also exposes functions `Query` and `Mutate` for doing custom query or mutations.  We are running on top of this [go graphql library](https://github.com/hasura/go-graphql-client).
