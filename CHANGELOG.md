<a name="unreleased"></a>
## [Unreleased]


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


[Unreleased]: https://github.com/OpsLevel/opslevel-go/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.3...v0.2.0
[v0.1.3]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.0...v0.1.1
