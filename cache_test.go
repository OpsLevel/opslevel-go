package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2024"
	"github.com/rocktavious/autopilot/v2023"
)

func TestCache(t *testing.T) {
	// Arrange
	testRequestOne := autopilot.NewTestRequest(
		`query TierList{account{tiers{alias,description,id,index,name}}}`,
		`{}`,
		`{"data":{"account":{ "tiers": [ {{ template "tier_1" }} ] }}}`,
	)
	testRequestTwo := autopilot.NewTestRequest(
		`query LifecycleList{account{lifecycles{alias,description,id,index,name}}}`,
		`{}`,
		`{"data":{"account":{ "lifecycles":[{{ template "lifecycle_1" }}] }}}`,
	)
	testRequestThree := autopilot.NewTestRequest(
		`query SystemsList($after:String!$first:Int!){account{systems(after: $after, first: $first){nodes{id,aliases,managedAliases,name,description,htmlUrl,owner{... on Team{teamAlias:alias,id}},parent{id,aliases,description,htmlUrl,managedAliases,name,note,owner{... on Team{teamAlias:alias,id}}},note},{{ template "pagination_request" }}}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "systems":{ "nodes":[{{ template "system1_response" }}] } }}}`,
	)
	testRequestFour := autopilot.NewTestRequest(
		`query TeamList($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,managedAliases,contacts{address,displayName,displayType,externalId,id,isDefault,type},htmlUrl,manager{id,email,htmlUrl,name,role},memberships{nodes{role,team{alias,id},user{id,email}},{{ template "pagination_request" }},totalCount},name,parentTeam{alias,id},responsibilities,tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount}},{{ template "pagination_request" }},totalCount}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "teams":{ "nodes":[{{ template "team_1" }}] } }}}`,
	)
	testRequestFive := autopilot.NewTestRequest(
		`query CategoryList($after:String!$first:Int!){account{rubric{categories(after: $after, first: $first){nodes{description,id,name},{{ template "pagination_request" }},totalCount}}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{"rubric":{ "categories":{ "nodes":[{{ template "category_1" }}] } }}}}`,
	)
	testRequestSix := autopilot.NewTestRequest(
		`{account{rubric{levels{nodes{alias,description,id,index,name},{{ template "pagination_request" }},totalCount}}}}`,
		`{}`,
		`{"data":{"account":{"rubric":{ "levels":{ "nodes":[{{ template "level_1" }}] } }}}}`,
	)
	testRequestSeven := autopilot.NewTestRequest(
		`query FilterList($after:String!$first:Int!){account{filters(after: $after, first: $first){nodes{id,name,connective,htmlUrl,predicates{caseSensitive,key,keyData,type,value}},{{ template "pagination_request" }},totalCount}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "filters":{ "nodes":[{{ template "filter_1" }}] } }}}`,
	)
	testRequestEight := autopilot.NewTestRequest(
		`query IntegrationList($after:String!$first:Int!){account{integrations(after: $after, first: $first){nodes{{ template "integration_request" }},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "integrations":{ "nodes":[{{ template "integration_1" }}] } }}}`,
	)
	testRequestNine := autopilot.NewTestRequest(
		`query RepositoryList($after:String!$first:Int!$visible:Boolean!){account{repositories(after: $after, first: $first, visible: $visible){hiddenCount,nodes{archivedAt,createdOn,defaultAlias,defaultBranch,description,forked,htmlUrl,id,languages{name,usage},lastOwnerChangedAt,locked,name,organization,owner{alias,id},private,repoKey,services{edges{atRoot,node{id,aliases},paths{href,path},serviceRepositories{baseDirectory,displayName,id,repository{id,defaultAlias},service{id,aliases}}},{{ template "pagination_request" }},totalCount},tags{nodes{id,key,value},{{ template "pagination_request" }},totalCount},tier{alias,description,id,index,name},type,url,visible},organizationCount,ownedCount,{{ template "pagination_request" }},totalCount,visibleCount}}}`,
		`{ "after": "", "first": 100, "visible": true }`,
		`{"data":{"account":{ "repositories":{ "hiddenCount": 0, "nodes":[{{ template "repository_1" }}] } }}}`,
	)
	testRequestTen := autopilot.NewTestRequest(
		`query InfrastructureResourceSchemaList($after:String!$first:Int!){account{infrastructureResourceSchemas(after: $after, first: $first){nodes{type,schema},{{ template "pagination_request" }}}}}`,
		`{ "after": "", "first": 100 }`,
		`{"data":{"account":{ "infrastructureResourceSchemas":{ "nodes":[ {{ template "infra_schema_1" }} ] }}}}`,
	)
	testRequestEleven := autopilot.TestRequest{}

	requests := []autopilot.TestRequest{
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
		testRequestEleven,
	}

	client1 := BestTestClient(t, "cache1", requests...)
	client2 := BestTestClient(t, "cache2", requests...)

	// Act
	ol.Cache.CacheAll(client1)

	ol.Cache.CacheTiers(client2)
	ol.Cache.CacheLifecycles(client2)
	ol.Cache.CacheSystems(client2)
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

	system1, system1Ok := ol.Cache.TryGetSystems("platformsystem1")
	system2, system2Ok := ol.Cache.TryGetSystems("does_not_exist")

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
	autopilot.Equals(t, id1, tier1.Id)
	autopilot.Equals(t, false, tier2Ok)
	autopilot.Equals(t, true, tier2 == nil)

	autopilot.Equals(t, true, lifecycle1Ok)
	autopilot.Equals(t, id1, lifecycle1.Id)
	autopilot.Equals(t, false, lifecycle2Ok)
	autopilot.Equals(t, true, lifecycle2 == nil)

	autopilot.Equals(t, true, system1Ok)
	autopilot.Equals(t, id1, system1.Id)
	autopilot.Equals(t, false, system2Ok)
	autopilot.Equals(t, true, system2 == nil)

	autopilot.Equals(t, true, team1Ok)
	autopilot.Equals(t, id1, team1.Id)
	autopilot.Equals(t, false, team2Ok)
	autopilot.Equals(t, true, team2 == nil)

	autopilot.Equals(t, true, category1Ok)
	autopilot.Equals(t, id1, category1.Id)
	autopilot.Equals(t, false, category2Ok)
	autopilot.Equals(t, true, category2 == nil)

	autopilot.Equals(t, true, level1Ok)
	autopilot.Equals(t, id1, level1.Id)
	autopilot.Equals(t, false, level2Ok)
	autopilot.Equals(t, true, level2 == nil)

	autopilot.Equals(t, true, filter1Ok)
	autopilot.Equals(t, id1, filter1.Id)
	autopilot.Equals(t, false, filter2Ok)
	autopilot.Equals(t, true, filter2 == nil)

	autopilot.Equals(t, true, integration1Ok)
	autopilot.Equals(t, id1, integration1.Id)
	autopilot.Equals(t, false, integration2Ok)
	autopilot.Equals(t, true, integration2 == nil)

	autopilot.Equals(t, true, repository1Ok)
	autopilot.Equals(t, id1, repository1.Id)
	autopilot.Equals(t, false, repository2Ok)
	autopilot.Equals(t, true, repository2 == nil)

	autopilot.Equals(t, true, infraSchema1OK)
	autopilot.Equals(t, "Database", infraSchema1.Type)
	autopilot.Equals(t, false, infraSchema2Ok)
	autopilot.Equals(t, true, infraSchema2 == nil)
}
