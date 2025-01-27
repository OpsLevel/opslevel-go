package opslevel

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/hasura/go-graphql-client"
)

type CategoryConnection struct {
	Nodes      []Category
	PageInfo   PageInfo
	TotalCount int
}

func (category *Category) Alias() string {
	return slug.Make(category.Name)
}

func (client *Client) CreateCategory(input CategoryCreateInput) (*Category, error) {
	var m struct {
		Payload struct {
			Category Category
			Errors   []Error
		} `graphql:"categoryCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CategoryCreate"))
	return &m.Payload.Category, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetCategory(id ID) (*Category, error) {
	var q struct {
		Account struct {
			Category Category `graphql:"category(id: $id)"`
		}
	}
	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("CategoryGet"))
	if q.Account.Category.Id == "" {
		err = graphql.Errors{graphql.Error{
			Message: fmt.Sprintf("category with ID '%s' not found", id),
			Path:    []any{"account", "category"},
		}}
	}
	return &q.Account.Category, HandleErrors(err, nil)
}

func (client *Client) ListCategories(variables *PayloadVariables) (*CategoryConnection, error) {
	var q struct {
		Account struct {
			Rubric struct {
				Categories CategoryConnection `graphql:"categories(after: $after, first: $first)"`
			}
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("CategoryList")); err != nil {
		return nil, err
	}
	for q.Account.Rubric.Categories.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Rubric.Categories.PageInfo.End
		resp, err := client.ListCategories(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Rubric.Categories.Nodes = append(q.Account.Rubric.Categories.Nodes, resp.Nodes...)
		q.Account.Rubric.Categories.PageInfo = resp.PageInfo
		q.Account.Rubric.Categories.TotalCount += resp.TotalCount
	}
	return &q.Account.Rubric.Categories, nil
}

func (client *Client) UpdateCategory(input CategoryUpdateInput) (*Category, error) {
	var m struct {
		Payload struct {
			Category Category
			Errors   []Error
		} `graphql:"categoryUpdate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("CategoryUpdate"))
	return &m.Payload.Category, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteCategory(id ID) error {
	var m struct {
		Payload struct {
			Id     ID `graphql:"deletedCategoryId"`
			Errors []Error
		} `graphql:"categoryDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": CategoryDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("CategoryDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
