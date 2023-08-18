package opslevel_test

import (
	"testing"

	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateAliases(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation AliasCreate($input:AliasCreateInput!){aliasCreate(input: $input){aliases,ownerId,errors{message,path}}}",
    "variables":{
        "input": {
            "alias": "MyAwesomeAlias",
            "ownerId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2"
        }
    }
}`
	response := `{"data": {
    "aliasCreate": {
        "aliases": [
            "MyCoolService",
            "MyAwesomeAlias"
        ],
        "ownerId": "Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2",
        "errors": []
    }
}}`
	client := ABetterTestClient(t, "aliases/create", request, response)
	// Act
	result, err := client.CreateAliases("Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS80MTc2", []string{"MyAwesomeAlias"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 2, len(result))
	autopilot.Equals(t, "MyAwesomeAlias", result[1])
}

// TODO: Add CreateAliasesOwnerNotFound

func TestDeleteServiceAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}",
    "variables":{
        "input": {
            "alias": "MyAwesomeAlias",
            "ownerType": "service"
        }
    }
}`
	response := `{"data": {
    "aliasDelete": {
        "deletedAlias": "MyAwesomeAlias",
        "errors": []
    }
}}`
	client := ABetterTestClient(t, "aliases/delete_service", request, response)
	// Act
	err := client.DeleteServiceAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation AliasDelete($input:AliasDeleteInput!){aliasDelete(input: $input){deletedAlias,errors{message,path}}}",
    "variables":{
        "input": {
            "alias": "MyAwesomeAlias",
            "ownerType": "team"
        }
    }
}`
	response := `{"data": {
    "aliasDelete": {
        "deletedAlias": "MyAwesomeAlias",
        "errors": []
    }
}}`
	client := ABetterTestClient(t, "aliases/delete_team", request, response)
	// Act
	err := client.DeleteTeamAlias("MyAwesomeAlias")
	// Assert
	autopilot.Ok(t, err)
}

// TODO: Add DeleteAliasNotFound
