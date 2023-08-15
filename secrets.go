package opslevel

type Secret struct {
	Alias      string     `json:"alias"`
	ID         ID         `json:"id"`
	Owner      *Team      `json:"team,omitempty"`
	Timestamps Timestamps `json:"timestamps"`
}

type SecretInput struct {
	Owner IdentifierInput `json:"owner" yaml:"owner"`
	Value string          `json:"value" yaml:"value"`
}

type SecretsVaultsSecretConnection struct {
	Nodes      []Secret
	PageInfo   PageInfo
	TotalCount int
}

func (client *Client) CreateSecret(alias string, input SecretInput) (*Secret, error) {
	var m struct {
		Payload struct {
			Secret Secret
			Errors []OpsLevelErrors
		} `graphql:"secretsVaultsSecretCreate(alias: $alias, input: $input)"`
	}
	v := PayloadVariables{
		"alias": alias,
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretCreate"))
	return &m.Payload.Secret, HandleErrors(err, m.Payload.Errors)
}

// List all Secrets for your account.
func (client *Client) ListSecretsVaultsSecret(variables *PayloadVariables) (SecretsVaultsSecretConnection, error) {
	var q struct {
		Account struct {
			SecretsVaultsSecret SecretsVaultsSecretConnection `graphql:"secretsVaultsSecrets(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("SecretList")); err != nil {
		return SecretsVaultsSecretConnection{}, err
	}
	for q.Account.SecretsVaultsSecret.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.SecretsVaultsSecret.PageInfo.End
		resp, err := client.ListSecretsVaultsSecret(variables)
		if err != nil {
			return SecretsVaultsSecretConnection{}, err
		}
		q.Account.SecretsVaultsSecret.Nodes = append(q.Account.SecretsVaultsSecret.Nodes, resp.Nodes...)
		q.Account.SecretsVaultsSecret.PageInfo = resp.PageInfo
	}
	return q.Account.SecretsVaultsSecret, nil
}

func (client *Client) UpdateSecret(identifier string, secretInput SecretInput) (*Secret, error) {
	identifierInput := IdentifierInput{
		Id: *NewID(identifier),
	}
	var m struct {
		Payload struct {
			Secret Secret
			Errors []OpsLevelErrors
		} `graphql:"secretsVaultsSecretUpdate(input: $input, secret: $secret)"`
	}
	v := PayloadVariables{
		"input":  secretInput,
		"secret": identifierInput,
	}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretUpdate"))
	return &m.Payload.Secret, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteSecret(identifier string) error {
	var m struct {
		Payload struct {
			Errors []OpsLevelErrors `graphql:"errors"`
		} `graphql:"secretsVaultsSecretDelete(resource: $input)"`
	}
	v := PayloadVariables{"input": *NewIdentifier(identifier)}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetSecret(id ID) (*Secret, error) {
	var q struct {
		Account struct {
			Secret Secret `graphql:"secret(id: $secret)"`
		}
	}
	v := PayloadVariables{"secret": id}
	if err := client.Query(&q, v, WithName("SecretGet")); err != nil {
		return nil, err
	}
	return &q.Account.Secret, nil
}
