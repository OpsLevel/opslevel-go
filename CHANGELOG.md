<a name="unreleased"></a>
## [Unreleased]


<a name="v0.3.3"></a>
## [v0.3.3] - 2021-09-14
### Bugfix
- add newly added PredicateType `PredicateTypeSatisfiesJqExpression` to the GetPredicateTypes return list


<a name="v0.3.2"></a>
## [v0.3.2] - 2021-09-14
### Feature
- add PredicateType “SatisfiesJqExpression”


<a name="v0.3.1"></a>
## [v0.3.1] - 2021-09-11
### Dependency
- bump github.com/rs/zerolog from 1.21.0 to 1.23.0 ([#13](https://github.com/OpsLevel/opslevel-go/issues/13))

### Feature
- expose defaultAlias on Repository struct as this is used for lookup by alias and is needed in downstream tools


<a name="v0.3.0"></a>
## [v0.3.0] - 2021-08-25
### Bugfix
- `GetTeamWithAlias` did not use correct graphql argument type
- fix issue when using TagUpdateInput where either key or value is blank
- fields in Contact and User do not need to be graphql.String - converted to just regular string
- make Tier.Id be a graphql ID rather then a string
- fix typo of field Description in `ServiceUpdateInput`

### Feature
- implement graphql fragments on Check struct so endusers can get at the custom fields per check type
- add ability to specify the index of the rubric level at creation time
- add validation of tag key names before sending to API
- add ability to list repositories by tier
- add ability to list teams with a manager email
- add more specialized listing methods for services to list by lifecycle, product and tier
- add ability to update or delete service tools
- add ability to update and delete service repository
- add “Team” as check owner to mutation repsonse
- add default 10 second time out to http client used by graphql client
- add list for checks
- add get for rubric category and level
- add get and list of integrations
- add create and update for manual check
- add create and update for custom and custom event check
- add create and update for payload check
- add create and update for tag defined and tool usage check
- add update repository check for file, integrated and search
- add create repository check for file, integrated and search
- add create service check for ownership, property and configuration
- add get and delete check by ID
- implement get filter by id
- implement CRUD for filters ([#9](https://github.com/OpsLevel/opslevel-go/issues/9))
- add dependabot ([#11](https://github.com/OpsLevel/opslevel-go/issues/11))

### Refactor
- remove field description from rubic category mutations
- CreateAliases should return an error that is aggregated from any errors
- check tests to use a map of testcases
- port service queries that were in terraform provider to core library for reuse


<a name="v0.2.2"></a>
## [v0.2.2] - 2021-07-21
### Feature
- implement CRUD for rubric categories and levels ([#8](https://github.com/OpsLevel/opslevel-go/issues/8))


<a name="v0.2.1"></a>
## [v0.2.1] - 2021-06-13
### Feature
- add ability to update a ServiceRepository
- add ability to connect service and repositories together
- add ability to query for repositories


<a name="v0.2.0"></a>
## [v0.2.0] - 2021-05-22
### Chore
- set MIT license

### Docs
- add badges to readme

### Feature
- support hydration of nested resources that are paginated
- improve error output during 401 unauthorized
- add method to check if a service has a tool

### Fix
- fix bug with assign tag to a service with alias or id


<a name="v0.1.3"></a>
## [v0.1.3] - 2021-05-16

<a name="v0.1.2"></a>
## [v0.1.2] - 2021-05-16
### Bugfix
- fix problem with listing teams due to invalid graphql types and recursion


<a name="v0.1.1"></a>
## [v0.1.1] - 2021-05-16
### Bugfix
- dedupe the aliases in the response of CreateAliases method

### Docs
- Add logo

### Feature
- Add tags for service paginated retrieval

### Refactor
- tests can now validate graphql query and variables
- remove usages of graphql types in favor of plain go types in public structures
- mark all methods that were named WithId as deprecated to adhere to better naming convetion
- change naming convetion for getting service data to not use the suffix WithId
- use graphql.ID instead of string in client methods to adhere to standards


<a name="v0.1.0"></a>
## v0.1.0 - 2021-05-12
### Refactor
- client configuration to allow for settings api visibility header


[Unreleased]: https://github.com/OpsLevel/opslevel-go/compare/v0.3.3...HEAD
[v0.3.3]: https://github.com/OpsLevel/opslevel-go/compare/v0.3.2...v0.3.3
[v0.3.2]: https://github.com/OpsLevel/opslevel-go/compare/v0.3.1...v0.3.2
[v0.3.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/OpsLevel/opslevel-go/compare/v0.2.2...v0.3.0
[v0.2.2]: https://github.com/OpsLevel/opslevel-go/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.3...v0.2.0
[v0.1.3]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.0...v0.1.1
