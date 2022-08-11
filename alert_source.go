package opslevel

import (
	"github.com/shurcooL/graphql"
)

type AlertSourceExternalIdentifier struct {
	Type       AlertSourceTypeEnum  `json:"type"`
	ExternalId string               `json:"externalId"`
}

type AlertSource struct {
	Name           string                 `graphql:"name"`
	Description    string                 `graphql:"description"`
	Id             graphql.ID             `graphql:"id"`
	Type           AlertSourceTypeEnum    `graphql:"type"`
	ExternalId     string                 `graphql:"externalId"`
	Integration    Integration            `graphql:"integration"`
	Url            string                 `graphql:"url"`
}

//#region Retrieve

func (client *Client) GetAlertSourceWithExternalIdentifier(input AlertSourceExternalIdentifier) (*AlertSource, error) {
	var q struct {
		Account struct {
			AlertSource AlertSource `graphql:"alertSource(externalIdentifier: $externalIdentifier)"`
		}
	}

	v := PayloadVariables{
		"externalIdentifier": input,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.AlertSource, nil
}

func (client *Client) GetAlertSource(id graphql.ID) (*AlertSource, error) {
	var q struct {
		Account struct {
			AlertSource AlertSource `graphql:"alertSource(id: $id)"`
		}
	}

	v := PayloadVariables{
		"id": id,
	}
	if err := client.Query(&q, v); err != nil {
		return nil, err
	}
	return &q.Account.AlertSource, nil
}

//#endregion