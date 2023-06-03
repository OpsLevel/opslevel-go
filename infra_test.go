package opslevel_test

import (
	"github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestCreateInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "infra/create", request, response)
	// Act
	result, err := client.CreateInfrastructure(opslevel.InfrastructureResourceInput{})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "", result.Name)
}

func TestGetInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "",
	"variables":{
		"id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf"
    }
}`
	response := `{"data": {
	"account": {
	}
}}`
	client := ABetterTestClient(t, "infra/get", request, response)
	// Act
	result, err := client.GetInfrastructure("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf", string(result.Id))
	autopilot.Equals(t, "", result.Name)
}

func TestListInfraSchemas(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructures": {
							"nodes": [
								{
									{{ template "" }}
								},
								{
									{{ template "" }} 
								}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
					}}}}`},
		{`{"query": "",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructures": {
							"nodes": [
								{
									{{ template "" }}
								}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
					}}}}`},
	}

	client := APaginatedTestClient(t, "infra/list_schemas", requests...)
	// Act
	response, err := client.ListInfrastructureSchemas(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "", result[1].Type)
	autopilot.Equals(t, "", result[2].Type)
}

func TestListInfra(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructures": {
							"nodes": [
								{
									{{ template "" }}
								},
								{
									{{ template "" }} 
								}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
					}}}}`},
		{`{"query": "",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"infrastructures": {
							"nodes": [
								{
									{{ template "" }}
								}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
					}}}}`},
	}

	client := APaginatedTestClient(t, "infra/list", requests...)
	// Act
	response, err := client.ListInfrastructure(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "", result[1].Name)
	autopilot.Equals(t, "", result[2].Name)
}

func TestUpdateInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "infra/update", request, response)
	// Act
	result, err := client.UpdateInfrastructure("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", opslevel.InfrastructureResourceInput{})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx", string(result.Id))
	autopilot.Equals(t, "", result.Name)
}

func TestDeleteInfra(t *testing.T) {
	// Arrange
	request := `{
	"query": "",
	"variables":{

   }
}`
	response := `{"data": {
	"infrastructureDelete": {
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "infra/delete", request, response)
	// Act
	err := client.DeleteInfrastructure("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	// Assert
	autopilot.Equals(t, nil, err)
}
