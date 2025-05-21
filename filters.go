package opslevel

import (
	"fmt"
	"slices"

	"github.com/gosimple/slug"
)

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
	if err := filterPredicate.validateValue(); err != nil {
		return err
	}
	if err := filterPredicate.validateCaseSensitivity(); err != nil {
		return err
	}
	if err := filterPredicate.validateKeyData(); err != nil {
		return err
	}

	return nil
}

func (filterPredicate *FilterPredicate) validateCaseSensitivity() error {
	if filterPredicate.CaseSensitive == nil {
		return nil
	}

	knownNotCaseSensitiveTypes := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotExist,
		PredicateTypeEnumExists,
		PredicateTypeEnumGreaterThanOrEqualTo,
		PredicateTypeEnumLessThanOrEqualTo,
		PredicateTypeEnumSatisfiesVersionConstraint,
		PredicateTypeEnumMatchesRegex,
		PredicateTypeEnumDoesNotMatchRegex,
		PredicateTypeEnumBelongsTo,
		PredicateTypeEnumMatches,
		PredicateTypeEnumDoesNotMatch,
		PredicateTypeEnumSatisfiesJqExpression,
	}
	if slices.Contains(knownNotCaseSensitiveTypes, filterPredicate.Type) && *filterPredicate.CaseSensitive {
		return fmt.Errorf("FilterPredicate type '%s' cannot have CaseSensitive value set", filterPredicate.Type)
	}
	return nil
}

func (filterPredicate *FilterPredicate) validateKeyData() error {
	keyDataExpectedTypes := []PredicateKeyEnum{
		PredicateKeyEnumProperties,
		PredicateKeyEnumTags,
	}
	knownNoKeyDataSetTypes := []PredicateKeyEnum{
		PredicateKeyEnumAliases,
		PredicateKeyEnumDomainID,
		PredicateKeyEnumFilterID,
		PredicateKeyEnumFramework,
		PredicateKeyEnumLifecycleIndex,
		PredicateKeyEnumLanguage,
		PredicateKeyEnumName,
		PredicateKeyEnumOwnerID,
		PredicateKeyEnumOwnerIDs,
		PredicateKeyEnumProduct,
		PredicateKeyEnumRepositoryIDs,
		PredicateKeyEnumSystemID,
		PredicateKeyEnumTierIndex,
	}

	if slices.Contains(keyDataExpectedTypes, filterPredicate.Key) && filterPredicate.KeyData == "" {
		return fmt.Errorf("FilterPredicate key '%s' expects a value for 'key_data'", filterPredicate.Key)
	}
	if slices.Contains(knownNoKeyDataSetTypes, filterPredicate.Key) && filterPredicate.KeyData != "" {
		return fmt.Errorf("FilterPredicate key '%s' cannot have a value set for 'key_data'", filterPredicate.Key)
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
	filterMatchTypes := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotMatch,
		PredicateTypeEnumMatches,
	}
	regexMatchesTypes := []PredicateTypeEnum{
		PredicateTypeEnumDoesNotMatchRegex,
		PredicateTypeEnumMatchesRegex,
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
		expectedPredicateTypes = filterMatchTypes
	case PredicateKeyEnumLifecycleIndex, PredicateKeyEnumTierIndex:
		expectedPredicateTypes = slices.Concat(equalsTypes, existsTypes, lessThanGreaterThanTypes)
	case PredicateKeyEnumOwnerIDs:
		expectedPredicateTypes = []PredicateTypeEnum{PredicateTypeEnumEquals}
	case PredicateKeyEnumDomainID, PredicateKeyEnumOwnerID, PredicateKeyEnumSystemID:
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
			return fmt.Errorf("FilterPredicate type '%s' cannot have value set", filterPredicate.Type)
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

func (filter *Filter) Alias() string {
	return slug.Make(filter.Name)
}

func (client *Client) CreateFilter(input FilterCreateInput) (*Filter, error) {
	var m struct {
		Payload FilterCreatePayload `graphql:"filterCreate(input: $input)"`
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
	if q.Account.Filters.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Filters.PageInfo.End
		resp, err := client.ListFilters(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Filters.Nodes = append(q.Account.Filters.Nodes, resp.Nodes...)
		q.Account.Filters.PageInfo = resp.PageInfo
	}
	q.Account.Filters.TotalCount = len(q.Account.Filters.Nodes)
	return &q.Account.Filters, nil
}

func (client *Client) UpdateFilter(input FilterUpdateInput) (*Filter, error) {
	var m struct {
		Payload FilterUpdatePayload `graphql:"filterUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("FilterUpdate"))
	return &m.Payload.Filter, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteFilter(id ID) error {
	var m struct {
		Payload struct { // TODO: fix this
			Id     ID `graphql:"deletedId"`
			Errors []Error
		} `graphql:"filterDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": DeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("FilterDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
