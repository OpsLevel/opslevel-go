package opslevel

func (client *Client) CreateSecret(alias string, input SecretInput) (*Secret, error) {
	var m struct {
		Payload SecretPayload `graphql:"secretsVaultsSecretCreate(alias: $alias, input: $input)"`
	}
	v := PayloadVariables{
		"alias": alias,
		"input": input,
	}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretCreate"))
	return &m.Payload.Secret, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) ListSecretsVaultsSecret(variables *PayloadVariables) (*SecretsVaultsSecretsConnection, error) {
	var q struct {
		Account struct {
			SecretsVaultsSecrets SecretsVaultsSecretsConnection `graphql:"secretsVaultsSecrets(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("SecretList")); err != nil {
		return nil, err
	}
	if q.Account.SecretsVaultsSecrets.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.SecretsVaultsSecrets.PageInfo.End
		resp, err := client.ListSecretsVaultsSecret(variables)
		if err != nil {
			return nil, err
		}
		q.Account.SecretsVaultsSecrets.Nodes = append(q.Account.SecretsVaultsSecrets.Nodes, resp.Nodes...)
		q.Account.SecretsVaultsSecrets.PageInfo = resp.PageInfo
	}
	q.Account.SecretsVaultsSecrets.TotalCount = len(q.Account.SecretsVaultsSecrets.Nodes)
	return &q.Account.SecretsVaultsSecrets, nil
}

func (client *Client) UpdateSecret(identifier string, secretInput SecretInput) (*Secret, error) {
	var m struct {
		Payload SecretPayload `graphql:"secretsVaultsSecretUpdate(input: $input, secret: $secret)"`
	}
	v := PayloadVariables{
		"input":  secretInput,
		"secret": *NewIdentifier(identifier),
	}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretUpdate"))
	return &m.Payload.Secret, HandleErrors(err, m.Payload.Errors)
}

func (client *Client) DeleteSecret(identifier string) error {
	var m struct {
		Payload BasePayload `graphql:"secretsVaultsSecretDelete(resource: $input)"`
	}
	v := PayloadVariables{"input": *NewIdentifier(identifier)}
	err := client.Mutate(&m, v, WithName("SecretsVaultsSecretDelete"))
	return HandleErrors(err, m.Payload.Errors)
}

func (client *Client) GetSecret(identifier string) (*Secret, error) {
	var q struct {
		Account struct {
			Secret Secret `graphql:"secretsVaultsSecret(input: $input)"`
		}
	}
	v := PayloadVariables{"input": *NewIdentifier(identifier)}
	if err := client.Query(&q, v, WithName("SecretsVaultsSecret")); err != nil {
		return nil, err
	}
	return &q.Account.Secret, nil
}
