package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
	"testing"
)

func TestSystemCreate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemCreate($input:SystemCreateInput!){systemCreate(input:$input){system{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}},errors{message,path}}}",
	"variables":{
		"input": {
			"name": "PlatformSystem3",
			"description": "creating this for testing purposes",
			"ownerId": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "hello world"
		  }
    }}`
	response := `{"data": {
		{{ template "system_create_response" }}
	}}`
	client := ABetterTestClient(t, "system/create", request, response)
	input := ol.SystemCreateInput{
		Name:        "PlatformSystem3",
		Description: "creating this for testing purposes",
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        "hello world",
	}
	// Act
	result, err := client.CreateSystem(input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUz", string(result.Id))
}

func TestSystemAssignService(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "system/assign_service", request, response)
	system := ol.SystemId{
		Id: "",
	}
	// Act
	err := system.AssignService(client, "", "")
	// Assert
	autopilot.Ok(t, err)
}

func TestSystemGetId(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "system/get_id", request, response)
	// Act
	result, err := client.GetSystem("MTIzNDU2Nzg5MTIzNDU2Nzg5")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
}

func TestSystemGetAlias(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "system/get_alias", request, response)
	// Act
	result, err := client.GetSystem("my-system")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
}

func TestSystemGetServices(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "system/get_services", request, response)
	system := ol.SystemId{
		Id: "",
	}
	// Act
	result, err := system.ChildServices(client, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, result.TotalCount)
}

func TestSystemGetTags(t *testing.T) {
	// Arrange
	request := `{
    "query": "",
	"variables":{

    }
}`
	response := `{"data": {

}}`
	client := ABetterTestClient(t, "system/get_tags", request, response)
	system := ol.SystemId{
		Id: "",
	}
	// Act
	result, err := system.Tags(client, nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, result.TotalCount)
}

func TestListSystems(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{"query": "query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_initial_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"systems": {
							"nodes": [
								{{ template "system1_response" }},
								{{ template "system2_response" }}
							],
							{{ template "pagination_initial_pageInfo_response" }},
							"totalCount": 2
						  }}}}`},
		{`{"query": "query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}",
			{{ template "pagination_second_query_variables" }}
			}`,
			`{
				"data": {
					"account": {
						"systems": {
							"nodes": [
								{{ template "system3_response" }}
							],
							{{ template "pagination_second_pageInfo_response" }},
							"totalCount": 1
						  }}}}`},
	}

	client := APaginatedTestClient(t, "system/list", requests...)
	// Act
	response, err := client.ListSystems(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "PlatformSystem1", result[0].Name)
	autopilot.Equals(t, "PlatformSystem2", result[1].Name)
	autopilot.Equals(t, "PlatformSystem3", result[2].Name)
}

func TestSystemUpdate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemUpdate($input:SystemInput!$system:IdentifierInput!){systemUpdate(system:$system,input:$input){system{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}},parent{id,aliases,name,description,htmlUrl,owner{... on Group{alias,id},... on Team{alias,id}}}},errors{message,path}}}",
	"variables":{
		"system":{"id":"Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy"},
		"input":{
			"name": "PlatformSystem4",
			"description":"System entry created for testing",
			"ownerId":"Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU",
			"note": "Please delete me"
		}
	}}`
	response := `{"data": {
		{{ template "systemUpdate_response" }}
	}}`
	client := ABetterTestClient(t, "system/update", request, response)
	input := ol.SystemUpdateInput{
		Name:        "PlatformSystem4",
		Description: "System entry created for testing",
		Owner:       ol.NewID("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzU"),
		Note:        "Please delete me",
	}
	// Act
	result, err := client.UpdateSystem("Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy", input)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvRW50aXR5T2JqZWN0LzUy", string(result.Id))
}

func TestSystemDelete(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation SystemDelete($input:IdentifierInput!){systemDelete(resource: $input){errors{message,path}}}",
	"variables":{"input":{"alias":"PlatformSystem3"}}
}`
	response := `{"data": {
	"systemDelete": {
      "errors": []
    }
}}`
	client := ABetterTestClient(t, "system/delete", request, response)
	// Act
	err := client.DeleteSystem("PlatformSystem3")
	// Assert
	autopilot.Ok(t, err)
}
