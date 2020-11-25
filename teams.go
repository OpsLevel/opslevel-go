package opslevel

import (
	"context"
	"fmt"
)

func (c *Client) GetTeam(ctx context.Context, alias string) (*Team, error) {
	params := map[string]interface{}{
		"teamAlias": alias,
	}
	var res teamResponse
	if err := c.Do(ctx, teamQuery, params, &res); err != nil {
		return nil, fmt.Errorf("could not find team: %w", err)
	}
	if res.Account.Team == nil {
		return nil, fmt.Errorf("no team was found by alias: %s", alias)
	}
	return res.Account.Team, nil
}

type Contact struct {
	DisplayName string
	Address string
}

type Team struct {
	Id               string
	Name             string
	Responsibilities string
	Manager          User
	Contacts         []Contact
}

type User struct {
	Name string
	Email string
}


type teamResponse struct {
	Account struct {
		Team *Team
	}
}

const teamQuery = `
query($teamAlias: String) {
  account {
    team(alias: $teamAlias){
      id
      name
      responsibilities
      manager {
        name
        email
      }
      contacts {
        displayName
        address
      }
    }
  }
}
`
