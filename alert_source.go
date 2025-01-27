package opslevel

type AlertSourceDeleteInput struct {
	Id ID `json:"id"`
}

func NewAlertSource(kind AlertSourceTypeEnum, id string) *AlertSourceExternalIdentifier {
	output := AlertSourceExternalIdentifier{
		Type:       kind,
		ExternalId: id,
	}
	return &output
}

func (client *Client) CreateAlertSourceService(input AlertSourceServiceCreateInput) (*AlertSourceService, error) {
	var m struct {
		Payload struct {
			AlertSourceService AlertSourceService
			Errors             []Error
		} `graphql:"alertSourceServiceCreate(input: $input)"`
	}
	v := PayloadVariables{
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("AlertSourceServiceCreate"))
	return &m.Payload.AlertSourceService, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetAlertSourceWithExternalIdentifier(input AlertSourceExternalIdentifier) (*AlertSource, error) {
	var q struct {
		Account struct {
			AlertSource AlertSource `graphql:"alertSource(externalIdentifier: $externalIdentifier)"`
		}
	}

	v := PayloadVariables{
		"externalIdentifier": input,
	}
	err := client.Query(&q, v, WithName("AlertSourceGet"))
	return &q.Account.AlertSource, HandleErrors(err, nil)
}

func (client *Client) GetAlertSource(id ID) (*AlertSource, error) {
	var q struct {
		Account struct {
			AlertSource AlertSource `graphql:"alertSource(id: $id)"`
		}
	}

	v := PayloadVariables{
		"id": id,
	}
	err := client.Query(&q, v, WithName("AlertSourceGet"))
	return &q.Account.AlertSource, HandleErrors(err, nil)
}

func (client *Client) DeleteAlertSourceService(id ID) error {
	var m struct {
		Payload struct {
			Errors []Error
		} `graphql:"alertSourceServiceDelete(input: $input)"`
	}
	v := PayloadVariables{
		"input": AlertSourceDeleteInput{Id: id},
	}
	err := client.Mutate(&m, v, WithName("AlertSourceServiceDelete"))
	return HandleErrors(err, m.Payload.Errors)
}
