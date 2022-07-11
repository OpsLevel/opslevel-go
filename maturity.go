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
func (s *MaturityReport) Get(category string) *Level {
	for _, breakdown := range s.CategoryBreakdown {
		if category == breakdown.Category.Name {
			return &breakdown.Level
		}
	}
	return nil
}

func (c *Client) ListServicesMaturity() ([]ServiceMaturity, error) {
	var q struct {
		Account struct {
			Services struct {
				Nodes    []ServiceMaturity
				PageInfo PageInfo
			} `graphql:"services(after: $after, first: $first)"`
		}
	}
	v := c.InitialPageVariables()

	var output []ServiceMaturity
	if err := c.Query(&q, v); err != nil {
		return nil, err
	}
	output = append(output, q.Account.Services.Nodes...)
	for q.Account.Services.PageInfo.HasNextPage {
		v["after"] = q.Account.Services.PageInfo.End
		if err := c.Query(&q, v); err != nil {
			return nil, err
		}
		output = append(output, q.Account.Services.Nodes...)
	}

	return output, nil
}
