package opslevel_test

import (
	"fmt"
	"testing"

	"github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateSecret(t *testing.T) {
	// Arrange
	request := `{
  "query": "mutation SecretsVaultsSecretCreate($alias:String!$input:SecretInput!){secretsVaultsSecretCreate(alias: $alias, input: $input){secret{alias,id,owner{alias,id},timestamps{createdAt,updatedAt}},errors{message,path}}}",
  "variables": {{ template "secret_create_vars" }}
  }`
	response := `{{ template "secret_create_response" }}`
	id2 := opslevel.NewID("Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5")
	client := ABetterTestClient(t, "secrets/create", request, response)
	fmt.Println(client)
	// Act
	secretInput := opslevel.SecretInput{
		Owner: opslevel.IdentifierInput{Id: *id2},
		Value: "my-secret",
	}
	result, err := client.CreateSecret("alias1", secretInput)

	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, *id2, result.Owner.Id)
}

func TestGetSecret(t *testing.T) {
	// Arrange
	request := `{
	"query": "query SecretsVaultsSecret($input:IdentifierInput!){account{secretsVaultsSecret(input: $input){alias,id,owner{alias,id},timestamps{createdAt,updatedAt}}}}",
  "variables": {{ template "secret_get_vars" }}
  }`
	response := `{{ template "secret_get_response" }}`
	client := ABetterTestClient(t, "secret/get", request, response)
	// Act
	result, err := client.GetSecret("Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5", string(result.ID))
}

func TestListSecrets(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{
			Request: `{"query": "query SecretList($after:String!$first:Int!){account{secretsVaultsSecrets(after: $after, first: $first){nodes{alias,id,owner{alias,id},timestamps{createdAt,updatedAt}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				{{ template "pagination_initial_query_variables" }}
      }`,
			Response: `{{ template "secret_list_response_1" }}`,
		},
		{
			Request: `{"query": "query SecretList($after:String!$first:Int!){account{secretsVaultsSecrets(after: $after, first: $first){nodes{alias,id,owner{alias,id},timestamps{createdAt,updatedAt}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
				{{ template "pagination_second_query_variables" }}
		    }`,
			Response: `{{ template "secret_list_response_2" }}`,
		},
	}

	client := APaginatedTestClient(t, "secrets/list", requests...)
	// Act
	secretsVaultsSecretConnection, err := client.ListSecretsVaultsSecret(nil)
	secretNode := secretsVaultsSecretConnection.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, secretsVaultsSecretConnection.TotalCount)
	autopilot.Equals(t, "example_2", secretNode[1].Alias)
	autopilot.Equals(t, secretNode[1].Alias, secretNode[1].Owner.Alias)
}

func TestUpdateSecret(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SecretsVaultsSecretUpdate($input:SecretInput!$secret:IdentifierInput!){secretsVaultsSecretUpdate(input: $input, secret: $secret){secret{alias,id,owner{alias,id},timestamps{createdAt,updatedAt}},errors{message,path}}}",
    "variables": {{ template "secret_update_vars" }}
  }`
	response := `{{ template "secret_update_response" }}`
	client := ABetterTestClient(t, "secrets/update", request, response)
	// Act
	id2 := opslevel.NewID("Z2lkOi8vOTg3NjU0MzIxMTIzNDU2Nzg5")
	secretInput := opslevel.SecretInput{
		Owner: opslevel.IdentifierInput{Id: *id2},
		Value: "secret_value_2",
	}
	result, err := client.UpdateSecret(string(*id2), secretInput)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, *id2, result.ID)
	autopilot.Equals(t, *id2, result.Owner.Id)
}

func TestDeleteSecrets(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SecretsVaultsSecretDelete($input:IdentifierInput!){secretsVaultsSecretDelete(resource: $input){errors{message,path}}}",
    "variables": {{ template "secret_delete_vars" }}
  }`
	response := `{{ template "secret_delete_response" }}`
	id1 := opslevel.NewID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	client := ABetterTestClient(t, "secrets/delete", request, response)
	// Act
	err := client.DeleteSecret(string(*id1))
	// Assert
	autopilot.Equals(t, nil, err)
}
