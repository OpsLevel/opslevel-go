package opslevel

import (
	"fmt"
	"slices"

	"github.com/gosimple/slug"
)

type FilterId struct {
	Id   ID
	Name string
}

type Filter struct {
	FilterId
	Connective ConnectiveEnum
	HtmlUrl    string
	Predicates []FilterPredicate
}

type FilterPredicate struct {
	Key           PredicateKeyEnum  `json:"key" yaml:"key" default:"repository_ids"`
	KeyData       string            `json:"keyData,omitempty" yaml:"keyData,omitempty" default:"null"`
	Type          PredicateTypeEnum `json:"type" yaml:"type" default:"equals"`
	Value         string            `json:"value,omitempty" yaml:"value,omitempty" default:"1"`
	CaseSensitive *bool             `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty" default:"false"`
}

// Validate the FilterPredicate based on known expectations before sending to API
func (filterPredicate *FilterPredicate) Validate() error {
	// validation common to Predicate and FilterPredicate types
	basicPredicate := Predicate{Type: filterPredicate.Type, Value: filterPredicate.Value}
	if err := basicPredicate.Validate(); err != nil {
		return err
	}

	// validation specific to FilterPredicate types
	if err := filterPredicate.validateKeyHasExpectedType(); err != nil {
		return err
	}
	if err := filterPredicate.validateCaseSensitivity(); err != nil {
		return err
	}
	if err := filterPredicate.validateValue(); err != nil {
		return err
	}

	return nil
}

func (filterPredicate *FilterPredicate) validateCaseSensitivity() error {
	if filterPredicate.CaseSensitive == nil {
		return nil
	}

	caseSensitiveTypes := []PredicateTypeEnum{
		PredicateTypeEnumContains,
		PredicateTypeEnumDoesNotContain,
		PredicateTypeEnumDoesNotEqual,
		PredicateTypeEnumEquals,
		PredicateTypeEnumEndsWith,
		PredicateTypeEnumStartsWith,
	}
	if *filterPredicate.CaseSensitive &&
		!slices.Contains(caseSensitiveTypes, filterPredicate.Type) {
		return fmt.Errorf("FilterPredicate type '%s' cannot have CaseSensitive value set.", filterPredicate.Type)
	}
	return nil
}

func (filterPredicate *FilterPredicate) validateKeyHasExpectedType() error {
	var expectedPredicateTypes []PredicateTypeEnum
	containsTypes := []PredicateTypeEnum{
		PredicateTypeEnumContains,
		PredicateTypeEnumDoesNotContain,
	}
	equalsTypes := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotEqual,
		PredicateTypeEnumEquals,
	}
	lessThanGreaterThanTypes := []PredicateTypeEnum{
		PredicateTypeEnumGreaterThanOrEqualTo,
		PredicateTypeEnumLessThanOrEqualTo,
	}
	regexMatchesTypes := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotMatch,
		PredicateTypeEnumMatches,
	}
	startsOrEndsWithTypes := []PredicateTypeEnum{
		PredicateTypeEnumEndsWith,
		PredicateTypeEnumStartsWith,
	}

	switch filterPredicate.Key {
	case PredicateKeyEnumAliases, PredicateKeyEnumFramework, PredicateKeyEnumLanguage, PredicateKeyEnumName, PredicateKeyEnumProduct:
		expectedPredicateTypes = slices.Concat(
			containsTypes,
			equalsTypes,
			existsTypes,
			regexMatchesTypes,
			startsOrEndsWithTypes,
		)
	case PredicateKeyEnumFilterID:
		expectedPredicateTypes = regexMatchesTypes
	case PredicateKeyEnumLifecycleIndex, PredicateKeyEnumTierIndex:
		expectedPredicateTypes = slices.Concat(equalsTypes, existsTypes, lessThanGreaterThanTypes)
	case PredicateKeyEnumDomainID, PredicateKeyEnumOwnerID, PredicateKeyEnumOwnerIDs, PredicateKeyEnumSystemID:
		expectedPredicateTypes = slices.Concat(equalsTypes, existsTypes)
	case PredicateKeyEnumProperties:
		expectedPredicateTypes = append(existsTypes, PredicateTypeEnumSatisfiesJqExpression)
	case PredicateKeyEnumRepositoryIDs:
		expectedPredicateTypes = existsTypes
	case PredicateKeyEnumTags:
		expectedPredicateTypes = append(slices.Concat(
			containsTypes,
			equalsTypes,
			existsTypes,
			regexMatchesTypes,
			startsOrEndsWithTypes,
		), PredicateTypeEnumSatisfiesVersionConstraint)
	default:
		return nil
	}

	if !slices.Contains(expectedPredicateTypes, filterPredicate.Type) {
		return fmt.Errorf(
			"FilterPredicate key '%s' expected to have one of the following types: %v",
			filterPredicate.Key,
			expectedPredicateTypes,
		)
	}

	return nil
}

// Validates Value requirements expected by some Key types
func (filterPredicate *FilterPredicate) validateValue() error {
	if slices.Contains(existsTypes, filterPredicate.Type) {
		if filterPredicate.Value != "" {
			return fmt.Errorf("FilterPredicate type '%s' cannot have value set.", filterPredicate.Type)
		}
		return nil
	}
	idPredicateKeyTypes := []PredicateKeyEnum{
		PredicateKeyEnumDomainID,
		PredicateKeyEnumFilterID,
		PredicateKeyEnumOwnerID,
		PredicateKeyEnumOwnerIDs,
		PredicateKeyEnumSystemID,
	}
	if slices.Contains(idPredicateKeyTypes, filterPredicate.Key) && !IsID(filterPredicate.Value) {
		return fmt.Errorf("FilterPredicate with key '%s' expects value to be an ID", filterPredicate.Key)
	}
	return nil
}

type FilterConnection struct {
	Nodes      []Filter
	PageInfo   PageInfo
	TotalCount int
}

func (filter *Filter) Alias() string {
	return slug.Make(filter.Name)
}

func (client *Client) CreateFilter(input FilterCreateInput) (*Filter, error) {
	var m struct {
		Payload struct {
			Filter Filter
			Errors []OpsLevelErrors
		} `graphql:"filterCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("FilterCreate"))
	return &m.Payload.Filter, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetFilter(id ID) (*Filter, error) {
	var q struct {
		Account struct {
			Filter Filter `graphql:"filter(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("FilterGet"))
	if q.Account.Filter.Id == "" {
		err = fmt.Errorf("filter with ID '%s' not found", id)
	}
	return &q.Account.Filter, HandleErrors(err, nil)
}

func (client *Client) ListFilters(variables *PayloadVariables) (*FilterConnection, error) {
	var q struct {
		Account struct {
			Filters FilterConnection `graphql:"filters(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("FilterList")); err != nil {
		return nil, err
	}
	for q.Account.Filters.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Filters.PageInfo.End
		resp, err := client.ListFilters(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Filters.Nodes = append(q.Account.Filters.Nodes, resp.Nodes...)
		q.Account.Filters.PageInfo = resp.PageInfo
		q.Account.Filters.TotalCount += resp.TotalCount
	}
	return &q.Account.Filters, nil
}

func (client *Client) UpdateFilter(input FilterUpdateInput) (*Filter, error) {
	var m struct {
		Payload struct {
			Filter Filter
			Errors []OpsLevelErrors
		} `graphql:"filterUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("FilterUpdate"))
	return &m.Payload.Filter, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteFilter(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedId"`
			Errors []OpsLevelErrors
		} `graphql:"filterDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("FilterDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
