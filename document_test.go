package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2025"

	"github.com/rocktavious/autopilot/v2023"
)

func TestServiceApiDocSettingsUpdate(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{"docPath":"/src/swagger.json", "docSource":"PULL", "service": {"alias":"service_alias" }}`,
		`{"data": {"serviceApiDocSettingsUpdate": {"service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/api_doc_settings_update", testRequest)
	// Act
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "/src/swagger.json", &ol.ApiDocumentSourceEnumPull)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, true, result.Locked)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}

func TestServiceApiDocSettingsUpdateDocSourceNull(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{"docPath":"/src/swagger.json", "docSource": null, "service": {"alias":"service_alias" }}`,
		`{"data": { "serviceApiDocSettingsUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/api_doc_settings_update_doc_source_null", testRequest)
	// Act
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "/src/swagger.json", nil)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}

func TestServiceApiDocSettingsUpdateDocPathNull(t *testing.T) {
	// Arrange
	testRequest := autopilot.NewTestRequest(
		`mutation ServiceApiDocSettingsUpdate($docPath:String$docSource:ApiDocumentSourceEnum$service:IdentifierInput!){serviceApiDocSettingsUpdate(service: $service, apiDocumentPath: $docPath, preferredApiDocumentSource: $docSource){service{apiDocumentPath,description,framework,htmlUrl,id,aliases,language,lifecycle{alias,description,id,index,name},locked,managedAliases,maturityReport{overallLevel{alias,checks{id,name},description,id,index,name}},name,note,owner{alias,id},parent{id,aliases},preferredApiDocument{id,htmlUrl,source{... on ApiDocIntegration{id,name,type},... on ServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},timestamps{createdAt,updatedAt}},preferredApiDocumentSource,product,repos{edges{node{id,defaultAlias},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }}},defaultServiceRepository{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}},tags{nodes{id,key,value},{{ template "pagination_request" }}},tier{alias,description,id,index,name},timestamps{createdAt,updatedAt},tools{nodes{category,categoryAlias,displayName,environment,id,service{id,aliases},url},{{ template "pagination_request" }}},type{id,aliases}},errors{message,path}}}`,
		`{"docPath":null, "docSource":"PULL", "service": {"alias":"service_alias" }}`,
		`{"data": { "serviceApiDocSettingsUpdate": { "service": {{ template "service_1" }}, "errors": [] }}}`,
	)

	client := BestTestClient(t, "service/api_doc_settings_update_doc_path_null", testRequest)
	// Act
	result, err := client.ServiceApiDocSettingsUpdate("service_alias", "", &ol.ApiDocumentSourceEnumPull)
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, id1, result.Id)
	autopilot.Equals(t, ol.ApiDocumentSourceEnumPull, *result.PreferredApiDocumentSource)
	autopilot.Equals(t, "/src/swagger.json", result.ApiDocumentPath)
}
