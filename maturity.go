package opslevel

type CategoryBreakdown struct {
	Category Category
	Level    Level
}

type MaturityReport struct {
	CategoryBreakdown []CategoryBreakdown
	OverallLevel      Level
}

type ServiceMaturity struct {
	Name           string
	MaturityReport MaturityReport
}

// Get Given a 'category' name returns the 'Level'
func (maturityReport *MaturityReport) Get(category string) *Level {
	for _, breakdown := range maturityReport.CategoryBreakdown {
		if category == breakdown.Category.Name {
			return &breakdown.Level
		}
	}
	return nil
}

func (client *Client) GetServiceMaturityWithAlias(alias string) (*ServiceMaturity, error) {
	var q struct {
		Account struct {
			Service ServiceMaturity `graphql:"service(alias:$service)"`
		}
	}
	v := PayloadVariables{
		"service": alias,
	}
	err := client.Query(&q, v)
	return &q.Account.Service, HandleErrors(err, nil)
}

// NOTE: ServiceMaturityConnection is not part of GraphQL API schema
type ServiceMaturityConnection struct {
	Nodes      []ServiceMaturity
	PageInfo   PageInfo
	TotalCount int `graphql:"-"`
}

func (client *Client) ListServicesMaturity(variables *PayloadVariables) (*ServiceMaturityConnection, error) {
	var q struct {
		Account struct {
			Services ServiceMaturityConnection `graphql:"services(after: $after, first: $first)"`
		}
	}
	if variables == nil {
		variables = client.InitialPageVariablesPointer()
	}
	if err := client.Query(&q, *variables, WithName("ServiceMaturityList")); err != nil {
		return nil, err
	}

	for q.Account.Services.PageInfo.HasNextPage {
		(*variables)["after"] = q.Account.Services.PageInfo.End
		resp, err := client.ListServicesMaturity(variables)
		if err != nil {
			return nil, err
		}
		q.Account.Services.Nodes = append(q.Account.Services.Nodes, resp.Nodes...)
		q.Account.Services.PageInfo = resp.PageInfo
	}
	q.Account.Services.TotalCount = len(q.Account.Services.Nodes)
	return &q.Account.Services, nil
}
