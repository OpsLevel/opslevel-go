<a name="unreleased"></a>
## [Unreleased]


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


[Unreleased]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.1...HEAD
[v0.1.1]: https://github.com/OpsLevel/opslevel-go/compare/v0.1.0...v0.1.1
