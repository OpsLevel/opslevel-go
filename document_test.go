package opslevel_test

import (
	ol "github.com/opslevel/opslevel-go/v2023"
	"testing"

	"github.com/rocktavious/autopilot/v2022"
)

func TestServiceApiDocSettingsUpdate(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},errors{message,path}}}",
	"variables":{
		"docPath":"/src/swagger.json",
		"docSource":"PULL",
		"service":{
			"alias":"service_alias"
		}
    }
}`
	response := `{"data": {
	"serviceApiDocSettingsUpdate": {
		"service": {{ template "service_1" }},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "service/api_doc_settings_update", request, response)
	// Act
	docSource := ol.ApiDocumentSourceEnumPull
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "/src/swagger.json", &docSource)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}

func TestServiceApiDocSettingsUpdateDocSourceNull(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},errors{message,path}}}",
	"variables":{
		"docPath":"/src/swagger.json",
		"docSource": null,
		"service":{
			"alias":"service_alias"
		}
    }
}`
	response := `{"data": {
	"serviceApiDocSettingsUpdate": {
		"service": {{ template "service_1" }},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "service/api_doc_settings_update_doc_source_null", request, response)
	// Act
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "/src/swagger.json", nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}

func TestServiceApiDocSettingsUpdateDocPathNull(t *testing.T) {
	// Arrange
	request := `{
    "query": "mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},name,owner{alias,id},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,url,service{id,aliases}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},errors{message,path}}}",
	"variables":{
		"docPath":null,
		"docSource":"PULL",
		"service":{
			"alias":"service_alias"
		}
    }
}`
	response := `{"data": {
	"serviceApiDocSettingsUpdate": {
		"service": {{ template "service_1" }},
		"errors": []
	}
}}`
	client := ABetterTestClient(t, "service/api_doc_settings_update_doc_path_null", request, response)
	// Act
	docSource := ol.ApiDocumentSourceEnumPull
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "", &docSource)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "MTIzNDU2Nzg5MTIzNDU2Nzg5", string(result.Id))
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}
