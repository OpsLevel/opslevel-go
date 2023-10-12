package opslevel_test

import (
	"testing"

	"github.com/opslevel/opslevel-go/v2023"

	"github.com/rocktavious/autopilot/v2023"
)

func TestCreateAWSIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation AWSIntegrationCreate($input:AwsIntegrationInput!){awsIntegrationCreate(input: $input){integration{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},errors{message,path}}}"`,
		`"variables": {"input": { "iamRole": "arn:aws:iam::XXXX:role/aws-integration-role", "externalId": "123456789", "ownershipTagKeys": ["owner"] }}`,
		`{"data": {
      "awsIntegrationCreate": {
        "integration": {
          "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
          "name": "AWS - XXXX",
          "type": "aws",
          "createdAt": "2023-04-26T16:25:29.574450Z",
          "installedAt": "2023-04-26T16:25:28.541124Z",
          "iamRole": "arn:aws:iam::XXXX:role/aws-integration-role",
          "externalId": "123456789",
          "awsTagsOverrideOwnership": true,
          "ownershipTagKeys": [
            "owner"
          ]
        },
        "errors": []
      }}}`,
	)
	client := BestTestClient(t, "integration/create_aws", testRequest)
	// Act
	result, err := client.CreateIntegrationAWS(opslevel.AWSIntegrationInput{
		IAMRole:    opslevel.NewString("arn:aws:iam::XXXX:role/aws-integration-role"),
		ExternalID: opslevel.NewString("123456789"),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "AWS - XXXX", result.Name)
}

func TestCreateNewRelicIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation NewRelicIntegrationCreate($input:NewRelicIntegrationInput!){newRelicIntegrationCreate(input: $input){integration{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},errors{message,path}}}"`,
		`"variables": { "input": { "apiKey": "123456789", "baseUrl": "https://api.newrelic.com/graphql", "accountKey": "XXXX" }}`,
		`{"data": {
      "newRelicIntegrationCreate": {
        "integration": {
          "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
          "name": "New Relic - XXXX",
          "type": "new_relic",
          "createdAt": "2023-04-26T16:25:29.574450Z",
          "installedAt": "2023-04-26T16:25:28.541124Z",
          "accountKey": "XXXX",
          "baseUrl": "https://api.newrelic.com/graphql"
        },
        "errors": []
      }}}`,
	)
	client := BestTestClient(t, "integration/create_new_relic", testRequest)
	// Act
	result, err := client.CreateIntegrationNewRelic(opslevel.NewRelicIntegrationInput{
		ApiKey:     opslevel.NewString("123456789"),
		BaseUrl:    opslevel.NewString("https://api.newrelic.com/graphql"),
		AccountKey: opslevel.NewString("XXXX"),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "New Relic - XXXX", result.Name)
}

func TestGetIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "query IntegrationGet($id:ID!){account{integration(id: $id){id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}}}}"`,
		`"variables": {"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf" }`,
		`{"data": {
      "account": {
        "integration": {
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf",
          "name": "Deploy",
          "type": "deploy"
        }
      }}}`,
	)
	client := BestTestClient(t, "integration/get", testRequest)
	// Act
	result, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf", string(result.Id))
	autopilot.Equals(t, "Deploy", result.Name)
}

func TestGetMissingIntegraion(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "query IntegrationGet($id:ID!){account{integration(id: $id){id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}}}}"`,
		`"variables":{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf" }`,
		`{"data": { "account": { "integration": null }}}`,
	)
	client := BestTestClient(t, "integration/get_missing", testRequest)
	// Act
	_, err := client.GetIntegration("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestListIntegrations(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query": "query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "integrations": { "nodes": [ { {{ template "deploy_integration_response" }} }, { {{ template "payload_integration_response" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query": "query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "integrations": { "nodes": [ { {{ template "kubernetes_integration_response" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := APaginatedTestClient(t, "integration/list", requests...)
	// Act
	response, err := client.ListIntegrations(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Payload", result[1].Name)
	autopilot.Equals(t, "Kubernetes", result[2].Name)
}

func TestUpdateAWSIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation AWSIntegrationUpdate($input:AwsIntegrationInput!$integration:IdentifierInput!){awsIntegrationUpdate(integration: $integration input: $input){integration{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},errors{message,path}}}"`,
		`"variables": {"integration": { "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx" }, "input": { "name": "Dev2", "externalId": "123456789", "ownershipTagKeys": null }}`,
		`{"data": {
      "awsIntegrationUpdate": {
        "integration": {
        "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
        "name": "Dev2",
        "type": "aws",
        "createdAt": "2023-04-26T16:25:29.574450Z",
        "installedAt": "2023-04-26T16:25:28.541124Z",
        "iamRole": "arn:aws:iam::XXXX:role/aws-integration-role",
        "externalId": "123456789",
        "awsTagsOverrideOwnership": true,
        "ownershipTagKeys": [
          "owner"
        ]
      },
      "errors": []
    }}}`,
	)
	client := BestTestClient(t, "integration/update_aws", testRequest)
	// Act
	result, err := client.UpdateIntegrationAWS("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", opslevel.AWSIntegrationInput{
		Name:       opslevel.NewString("Dev2"),
		ExternalID: opslevel.NewString("123456789"),
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "Dev2", result.Name)
}

func TestUpdateNewRelicIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation NewRelicIntegrationUpdate($input:NewRelicIntegrationInput!$resource:IdentifierInput!){newRelicIntegrationUpdate(input: $input resource: $resource){integration{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},errors{message,path}}}"`,
		`"variables": {"resource": { "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx" }, "input": { "baseUrl": "https://api-test.newrelic.com/graphql" }}`,
		`{"data": {
      "newRelicIntegrationUpdate": {
        "integration": {
        "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
        "name": "New Relic - XXXX",
        "type": "new_relic",
        "createdAt": "2023-04-26T16:25:29.574450Z",
        "installedAt": "2023-04-26T16:25:28.541124Z",
        "accountKey": "XXXX",
        "baseUrl": "https://api-test.newrelic.com/graphql"
      },
      "errors": []
    }}}`,
	)

	client := BestTestClient(t, "integration/update_new_relic", testRequest)
	// Act
	result, err := client.UpdateIntegrationNewRelic(
		"Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx",
		opslevel.NewRelicIntegrationInput{
			BaseUrl: opslevel.NewString("https://api-test.newrelic.com/graphql"),
		},
	)
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "https://api-test.newrelic.com/graphql", result.BaseUrl)
}

func TestDeleteIntegration(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query": "mutation IntegrationDelete($input:IdentifierInput!){integrationDelete(resource: $input){errors{message,path}}}"`,
		`"variables": {"input": { "id": "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx" }}`,
		`{"data": { "integrationDelete": { "errors": [] }}}`,
	)
	client := BestTestClient(t, "integration/delete", testRequest)
	// Act
	err := client.DeleteIntegration("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Equals(t, nil, err)
}
