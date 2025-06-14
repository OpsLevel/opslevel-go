package opslevel

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"

	"github.com/hasura/go-graphql-client"
	"github.com/relvacode/iso8601"
)

type CheckId struct {
	Id   ID     `json:"id"`
	Name string `json:"name"`
}

type CheckInputConstructor func() any

var CheckCreateConstructors = map[CheckType]CheckInputConstructor{
	CheckTypeAlertSourceUsage:    func() any { return &CheckAlertSourceUsageCreateInput{} },
	CheckTypeCodeIssue:           func() any { return &CheckCodeIssueCreateInput{} },
	CheckTypeCustom:              func() any { return &CheckCreateInput{} },
	CheckTypeGeneric:             func() any { return &CheckCustomEventCreateInput{} },
	CheckTypeGitBranchProtection: func() any { return &CheckGitBranchProtectionCreateInput{} },
	CheckTypeHasDocumentation:    func() any { return &CheckHasDocumentationCreateInput{} },
	CheckTypeHasOwner:            func() any { return &CheckServiceOwnershipCreateInput{} },
	CheckTypeHasRecentDeploy:     func() any { return &CheckHasRecentDeployCreateInput{} },
	CheckTypeHasRepository:       func() any { return &CheckRepositoryIntegratedCreateInput{} },
	CheckTypeHasServiceConfig:    func() any { return &CheckServiceConfigurationCreateInput{} },
	CheckTypeManual:              func() any { return &CheckManualCreateInput{} },
	CheckTypePayload:             func() any { return &CheckCreateInput{} },
	CheckTypeRepoFile:            func() any { return &CheckRepositoryFileCreateInput{} },
	CheckTypeRepoGrep:            func() any { return &CheckRepositoryGrepCreateInput{} },
	CheckTypeRepoSearch:          func() any { return &CheckRepositorySearchCreateInput{} },
	CheckTypeServiceDependency:   func() any { return &CheckServiceDependencyCreateInput{} },
	CheckTypeServiceProperty:     func() any { return &CheckServicePropertyCreateInput{} },
	CheckTypeTagDefined:          func() any { return &CheckTagDefinedCreateInput{} },
	CheckTypeToolUsage:           func() any { return &CheckToolUsageCreateInput{} },
	CheckTypePackageVersion:      func() any { return &CheckPackageVersionCreateInput{} },
}

var CheckUpdateConstructors = map[CheckType]CheckInputConstructor{
	CheckTypeAlertSourceUsage:    func() any { return &CheckAlertSourceUsageUpdateInput{} },
	CheckTypeCodeIssue:           func() any { return &CheckCodeIssueUpdateInput{} },
	CheckTypeCustom:              func() any { return &CheckUpdateInput{} },
	CheckTypeGeneric:             func() any { return &CheckCustomEventUpdateInput{} },
	CheckTypeGitBranchProtection: func() any { return &CheckGitBranchProtectionUpdateInput{} },
	CheckTypeHasDocumentation:    func() any { return &CheckHasDocumentationUpdateInput{} },
	CheckTypeHasOwner:            func() any { return &CheckServiceOwnershipUpdateInput{} },
	CheckTypeHasRecentDeploy:     func() any { return &CheckHasRecentDeployUpdateInput{} },
	CheckTypeHasRepository:       func() any { return &CheckRepositoryIntegratedUpdateInput{} },
	CheckTypeHasServiceConfig:    func() any { return &CheckServiceConfigurationUpdateInput{} },
	CheckTypeManual:              func() any { return &CheckManualUpdateInput{} },
	CheckTypePayload:             func() any { return &CheckUpdateInput{} },
	CheckTypeRepoFile:            func() any { return &CheckRepositoryFileUpdateInput{} },
	CheckTypeRepoGrep:            func() any { return &CheckRepositoryGrepUpdateInput{} },
	CheckTypeRepoSearch:          func() any { return &CheckRepositorySearchUpdateInput{} },
	CheckTypeServiceDependency:   func() any { return &CheckServiceDependencyUpdateInput{} },
	CheckTypeServiceProperty:     func() any { return &CheckServicePropertyUpdateInput{} },
	CheckTypeTagDefined:          func() any { return &CheckTagDefinedUpdateInput{} },
	CheckTypeToolUsage:           func() any { return &CheckToolUsageUpdateInput{} },
	CheckTypePackageVersion:      func() any { return &CheckPackageVersionUpdateInput{} },
}

func UnmarshalCheckCreateInput(checkType CheckType, data []byte) (any, error) {
	output := CheckCreateConstructors[checkType]()
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output, nil
}

func UnmarshalCheckUpdateInput(checkType CheckType, data []byte) (any, error) {
	output := CheckUpdateConstructors[checkType]()
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return output, nil
}

type CheckCreateInputProvider interface {
	GetCheckCreateInput() *CheckCreateInput
}

type CheckCreateInput struct {
	Category ID                      `json:"categoryId" yaml:"categoryId" mapstructure:"categoryId"`
	EnableOn *Nullable[iso8601.Time] `json:"enableOn,omitempty" yaml:"enableOn,omitempty" mapstructure:"enabledOn,omitempty"`
	Enabled  *Nullable[bool]         `json:"enabled,omitempty" yaml:"enabled,omitempty" mapstructure:"enabled"`
	Filter   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" mapstructure:"filterId,omitempty"`
	Level    ID                      `json:"levelId" yaml:"levelId" mapstructure:"levelId"`
	Name     string                  `json:"name" yaml:"name" mapstructure:"name"`
	Notes    *string                 `json:"notes,omitempty" yaml:"notes,omitempty" mapstructure:"notes,omitempty"`
	Owner    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" mapstructure:"ownerId,omitempty"`
}

func NewCheckCreateInputTypeOf[T any](checkCreateInput CheckCreateInput) *T {
	newCheck := new(T)
	if err := mapstructure.Decode(checkCreateInput, newCheck); err != nil {
		panic(err)
	}
	return newCheck
}

type CheckUpdateInputProvider interface {
	GetCheckUpdateInput() *CheckUpdateInput
}

type CheckUpdateInput struct {
	Category *Nullable[ID]           `json:"categoryId,omitempty" mapstructure:"categoryId,omitempty"`
	EnableOn *Nullable[iso8601.Time] `json:"enableOn,omitempty" mapstructure:"enabledOn,omitempty"`
	Enabled  *Nullable[bool]         `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Filter   *Nullable[ID]           `json:"filterId,omitempty" yaml:"filterId,omitempty" mapstructure:"filterId,omitempty"`
	Id       ID                      `json:"id" mapstructure:"id"`
	Level    *Nullable[ID]           `json:"levelId,omitempty" mapstructure:"levelId,omitempty"`
	Name     *Nullable[string]       `json:"name,omitempty" mapstructure:"name,omitempty"`
	Notes    *string                 `json:"notes,omitempty" mapstructure:"notes,omitempty"`
	Owner    *Nullable[ID]           `json:"ownerId,omitempty" yaml:"ownerId,omitempty" mapstructure:"ownerId,omitempty"`
}

func NewCheckUpdateInputTypeOf[T any](checkUpdateInput CheckUpdateInput) *T {
	newCheck := new(T)
	if err := mapstructure.Decode(checkUpdateInput, newCheck); err != nil {
		panic(err)
	}
	return newCheck
}

func (client *Client) CreateCheck(input any) (*Check, error) {
	switch v := input.(type) {
	case *CheckAlertSourceUsageCreateInput:
		return client.CreateCheckAlertSourceUsage(*v)
	case *CheckCodeIssueCreateInput:
		return client.CreateCheckCodeIssue(*v)
	case *CheckCustomEventCreateInput:
		return client.CreateCheckCustomEvent(*v)
	case *CheckGitBranchProtectionCreateInput:
		return client.CreateCheckGitBranchProtection(*v)
	case *CheckHasDocumentationCreateInput:
		return client.CreateCheckHasDocumentation(*v)
	case *CheckServiceOwnershipCreateInput:
		return client.CreateCheckServiceOwnership(*v)
	case *CheckHasRecentDeployCreateInput:
		return client.CreateCheckHasRecentDeploy(*v)
	case *CheckRepositoryIntegratedCreateInput:
		return client.CreateCheckRepositoryIntegrated(*v)
	case *CheckServiceConfigurationCreateInput:
		return client.CreateCheckServiceConfiguration(*v)
	case *CheckManualCreateInput:
		return client.CreateCheckManual(*v)
	case *CheckRepositoryFileCreateInput:
		return client.CreateCheckRepositoryFile(*v)
	case *CheckRepositoryGrepCreateInput:
		return client.CreateCheckRepositoryGrep(*v)
	case *CheckRepositorySearchCreateInput:
		return client.CreateCheckRepositorySearch(*v)
	case *CheckServiceDependencyCreateInput:
		return client.CreateCheckServiceDependency(*v)
	case *CheckServicePropertyCreateInput:
		return client.CreateCheckServiceProperty(*v)
	case *CheckTagDefinedCreateInput:
		return client.CreateCheckTagDefined(*v)
	case *CheckToolUsageCreateInput:
		return client.CreateCheckToolUsage(*v)
	case *CheckPackageVersionCreateInput:
		return client.CreateCheckPackageVersion(*v)
	}
	return nil, fmt.Errorf("unknown input type %T", input)
}

func (client *Client) GetCheck(id ID) (*Check, error) {
	var q struct {
		Account struct {
			Check Check `graphql:"check(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("CheckGet"))
	if q.Account.Check.Id == "" {
		err = graphql.Errors{graphql.Error{
			Message: fmt.Sprintf("check with ID '%s' not found", id),
			Path:    []any{"account", "check"},
		}}
	}
	return &q.Account.Check, HandleErrors(err, nil)
}

func (client *Client) ListChecks(variables *PayloadVariables) (*CheckConnection, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Checks CheckConnection `graphql:"checks(after: $after, first: $first)"`
			}
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("CheckList")); err != nil {
		return nil, err
	}
	if q.Account.Rubric.Checks.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Rubric.Checks.PageInfo.End
		resp, err := client.ListChecks(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Rubric.Checks.Nodes = append(q.Account.Rubric.Checks.Nodes, resp.Nodes...)
		q.Account.Rubric.Checks.PageInfo = resp.PageInfo
	}
	q.Account.Rubric.Checks.TotalCount = len(q.Account.Rubric.Checks.Nodes)
	return &q.Account.Rubric.Checks, nil
}

func (client *Client) UpdateCheck(input any) (*Check, error) {
	switch v := input.(type) {
	case *CheckAlertSourceUsageUpdateInput:
		return client.UpdateCheckAlertSourceUsage(*v)
	case *CheckCodeIssueUpdateInput:
		return client.UpdateCheckCodeIssue(*v)
	case *CheckCustomEventUpdateInput:
		return client.UpdateCheckCustomEvent(*v)
	case *CheckGitBranchProtectionUpdateInput:
		return client.UpdateCheckGitBranchProtection(*v)
	case *CheckHasDocumentationUpdateInput:
		return client.UpdateCheckHasDocumentation(*v)
	case *CheckServiceOwnershipUpdateInput:
		return client.UpdateCheckServiceOwnership(*v)
	case *CheckHasRecentDeployUpdateInput:
		return client.UpdateCheckHasRecentDeploy(*v)
	case *CheckRepositoryIntegratedUpdateInput:
		return client.UpdateCheckRepositoryIntegrated(*v)
	case *CheckServiceConfigurationUpdateInput:
		return client.UpdateCheckServiceConfiguration(*v)
	case *CheckManualUpdateInput:
		return client.UpdateCheckManual(*v)
	case *CheckRepositoryFileUpdateInput:
		return client.UpdateCheckRepositoryFile(*v)
	case *CheckRepositoryGrepUpdateInput:
		return client.UpdateCheckRepositoryGrep(*v)
	case *CheckRepositorySearchUpdateInput:
		return client.UpdateCheckRepositorySearch(*v)
	case *CheckServiceDependencyUpdateInput:
		return client.UpdateCheckServiceDependency(*v)
	case *CheckServicePropertyUpdateInput:
		return client.UpdateCheckServiceProperty(*v)
	case *CheckTagDefinedUpdateInput:
		return client.UpdateCheckTagDefined(*v)
	case *CheckToolUsageUpdateInput:
		return client.UpdateCheckToolUsage(*v)
	case *CheckPackageVersionUpdateInput:
		return client.UpdateCheckPackageVersion(*v)
	}
	return nil, fmt.Errorf("unknown input type %T", input)
}

func (client *Client) DeleteCheck(id ID) error {
	var m struct {
		Payload BasePayload `graphql:"checkDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CheckDeleteInput{Id: RefOf(id)},
	}
	err := client.Mutate(&m, v, WithName("CheckDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
