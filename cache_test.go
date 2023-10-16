package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCache(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query TierList{account{tiers{alias,description,id,index,name}}}"`,
		`{}`,
		`{"data":{"account":{ "tiers": [ {{ template "tier_1" }} ] }}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query LifecycleList{account{lifecycles{alias,description,id,index,name}}}"`,
		`{}`,
		`{"data":{"account":{ "lifecycles":[{{ template "lifecycle_1" }}] }}}`,
	)
	testRequestThree := NewTestRequest(
		`"query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},name,parentTeam{id,alias},responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "teams":{ "nodes":[{{ template "team_1" }}] } }}}`,
	)
	testRequestFour := NewTestRequest(
		`"query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{id,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{"rubric":{ "categories":{ "nodes":[{{ template "category_1" }}] } }}}}`,
	)
	testRequestFive := NewTestRequest(
		`"{account{rubric{levels{nodes{alias,description,id,index,name},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}}"`,
		`{}`,
		`{"data":{"account":{"rubric":{ "levels":{ "nodes":[{{ template "level_1" }}] } }}}}`,
	)
	testRequestSix := NewTestRequest(
		`"query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "filters":{ "nodes":[{{ template "filter_1" }}] } }}}`,
	)
	testRequestSeven := NewTestRequest(
		`"query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{id,name,type,createdAt,installedAt,... on AwsIntegration{iamRole,externalId,awsTagsOverrideOwnership,ownershipTagKeys},... on NewRelicIntegration{baseUrl,accountKey}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "integrations":{ "nodes":[{{ template "integration_1" }}] } }}}`,
	)
	testRequestEight := NewTestRequest(
		`"query RepositoryList($after:String!$first:Int!){account{repositories(after: $after, first: $first){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount,visibleCount}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "repositories":{ "hiddenCount": 0, "nodes":[{{ template "repository_1" }}] } }}}`,
	)
	testRequestNine := NewTestRequest(
		`"query IntegrationList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{type,schema},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}"`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "infrastructureResourceSchemas":{ "nodes":[ {{ template "infra_schema_1" }} ] }}}}`,
	)
	testRequestTen := TestRequest{}

	requests := []TestRequest{
		testRequestOne,
		testRequestTwo,
		testRequestThree,
		testRequestFour,
		testRequestFive,
		testRequestSix,
		testRequestSeven,
		testRequestEight,
		testRequestNine,
		testRequestTen,
	}

	id := ol.ID("Z2lkOi8vMTIzNDU2Nzg5OTg3NjU0MzIx")
	client1 := BestTestClient(t, "cache1", requests...)
	client2 := BestTestClient(t, "cache2", requests...)

	// Act
	ol.Cache.CacheAll(client1)

	ol.Cache.CacheTiers(client2)
	ol.Cache.CacheLifecycles(client2)
	ol.Cache.CacheTeams(client2)
	ol.Cache.CacheCategories(client2)
	ol.Cache.CacheLevels(client2)
	ol.Cache.CacheFilters(client2)
	ol.Cache.CacheIntegrations(client2)
	ol.Cache.CacheRepositories(client2)
	ol.Cache.CacheInfraSchemas(client2)

	tier1, tier1Ok := ol.Cache.TryGetTier("example")
	tier2, tier2Ok := ol.Cache.TryGetTier("does_not_exist")

	lifecycle1, lifecycle1Ok := ol.Cache.TryGetLifecycle("example")
	lifecycle2, lifecycle2Ok := ol.Cache.TryGetLifecycle("does_not_exist")

	team1, team1Ok := ol.Cache.TryGetTeam("example")
	team2, team2Ok := ol.Cache.TryGetTeam("does_not_exist")

	category1, category1Ok := ol.Cache.TryGetCategory("example")
	category2, category2Ok := ol.Cache.TryGetCategory("does_not_exist")

	level1, level1Ok := ol.Cache.TryGetLevel("example")
	level2, level2Ok := ol.Cache.TryGetLevel("does_not_exist")

	filter1, filter1Ok := ol.Cache.TryGetFilter("example")
	filter2, filter2Ok := ol.Cache.TryGetFilter("does_not_exist")

	integration1, integration1Ok := ol.Cache.TryGetIntegration("deploy-example")
	integration2, integration2Ok := ol.Cache.TryGetIntegration("does_not_exist")

	repository1, repository1Ok := ol.Cache.TryGetRepository("github.com:rocktavious/autopilot")
	repository2, repository2Ok := ol.Cache.TryGetRepository("does_not_exist")

	infraSchema1, infraSchema1OK := ol.Cache.TryGetInfrastructureSchema("Database")
	infraSchema2, infraSchema2Ok := ol.Cache.TryGetInfrastructureSchema("does_not_exist")

	// Assert
	autopilot.Equals(t, true, tier1Ok)
	autopilot.Equals(t, id, tier1.Id)
	autopilot.Equals(t, false, tier2Ok)
	autopilot.Equals(t, true, tier2 == nil)

	autopilot.Equals(t, true, lifecycle1Ok)
	autopilot.Equals(t, id, lifecycle1.Id)
	autopilot.Equals(t, false, lifecycle2Ok)
	autopilot.Equals(t, true, lifecycle2 == nil)

	autopilot.Equals(t, true, team1Ok)
	autopilot.Equals(t, id, team1.Id)
	autopilot.Equals(t, false, team2Ok)
	autopilot.Equals(t, true, team2 == nil)

	autopilot.Equals(t, true, category1Ok)
	autopilot.Equals(t, id, category1.Id)
	autopilot.Equals(t, false, category2Ok)
	autopilot.Equals(t, true, category2 == nil)

	autopilot.Equals(t, true, level1Ok)
	autopilot.Equals(t, id, level1.Id)
	autopilot.Equals(t, false, level2Ok)
	autopilot.Equals(t, true, level2 == nil)

	autopilot.Equals(t, true, filter1Ok)
	autopilot.Equals(t, id, filter1.Id)
	autopilot.Equals(t, false, filter2Ok)
	autopilot.Equals(t, true, filter2 == nil)

	autopilot.Equals(t, true, integration1Ok)
	autopilot.Equals(t, id, integration1.Id)
	autopilot.Equals(t, false, integration2Ok)
	autopilot.Equals(t, true, integration2 == nil)

	autopilot.Equals(t, true, repository1Ok)
	autopilot.Equals(t, id, repository1.Id)
	autopilot.Equals(t, false, repository2Ok)
	autopilot.Equals(t, true, repository2 == nil)

	autopilot.Equals(t, true, infraSchema1OK)
	autopilot.Equals(t, "Database", infraSchema1.Type)
	autopilot.Equals(t, false, infraSchema2Ok)
	autopilot.Equals(t, true, infraSchema2 == nil)
}
