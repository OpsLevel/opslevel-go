package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2022"
)

// TODO: not sure if there is a better way to handle reusing a client
// Probably should be a feature of autopilot
var getWithAliasClient *ol.Client

func getWithAliasTestClient(t *testing.T) *ol.Client {
	if getWithAliasClient == nil {
		getWithAliasClient = ATestClientAlt(t, "team/get", "team/get_with_alias")
	}
	return getWithAliasClient
}

func TestCreateTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/create")
	// Act
	contacts := []ol.ContactInput{
		ol.CreateContactSlack("#general", ""),
		ol.CreateContactWeb("https://example.com", "Homepage"),
	}
	result, err := client.CreateTeam(ol.TeamCreateInput{
		Name:             "Example",
		ManagerEmail:     "john@example.com",
		Responsibilities: "Foo & bar",
		Contacts:         contacts,
		Group:            ol.NewIdentifier("test_group"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "test_group", result.Group.Alias)
}

func TestGetTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/get")
	// Act
	result, err := client.GetTeam("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
}

func TestGetTeamWithAlias(t *testing.T) {
	// Arrange
	client := getWithAliasTestClient(t)
	// Act
	result, err := client.GetTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "john@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, 3, result.Tags.TotalCount)
}

func TestListTeams(t *testing.T) {
	// Arrange
	requests := []TestRequest{
		{`{
    "query": "query ($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
    "variables": {
      "after": "",
      "first": 100
    }
  }`,
			`{
    "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "devops",
              "aliases": [
                "devops"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/devops",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
              "manager": {
                "name": "Kyle Rockman",
                "email": "kyle.rockman@mac.com"
              },
              "members": {
                "nodes": [
                  {
                    "name": "Andrew Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Ashley Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Drew Example",
                    "email": "example@opslevel.com"
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "DevOps",
              "responsibilities": "Own Infra & Tools."
            },
            {
              "alias": "developers",
              "aliases": [
                "developers"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/developers",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NDk",
              "manager": {
                "name": "Kyle Rockman",
                "email": "kyle.rockman@mac.com"
              },
              "members": {
                "nodes": [
                  {
                    "name": "Andrew Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Ashley Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Drew Example",
                    "email": "example@opslevel.com"
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "Developers",
              "responsibilities": null
            },
            {
              "alias": "marketing",
              "aliases": [
                "marketing"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/marketing",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTA",
              "manager": {
                "name": "Kyle Rockman",
                "email": "kyle.rockman@mac.com"
              },
              "members": {
                "nodes": [
                  {
                    "name": "Andrew Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Ashley Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Drew Example",
                    "email": "example@opslevel.com"
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "Marketing",
              "responsibilities": null
            }
          ],
          "pageInfo": {
            "hasNextPage": true,
            "hasPreviousPage": false,
            "startCursor": "MQ",
            "endCursor": "NDc"
          }
        }
      }
    }
  }`},
		{`{
    "query": "query ($after:String!$first:Int!){account{teams(after: $after, first: $first){nodes{alias,id,aliases,contacts{address,displayName,id,type},group{alias,id},htmlUrl,manager{id,email,htmlUrl,name,role},members{nodes{id,email,htmlUrl,name,role},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}},name,responsibilities,tags{nodes{id,key,value},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor},totalCount}},pageInfo{hasNextPage,hasPreviousPage,startCursor,endCursor}}}}",
    "variables": {
      "after": "NDc",
      "first": 100
    }
  }`,
			`{
    "data": {
      "account": {
        "teams": {
          "nodes": [
            {
              "alias": "security",
              "aliases": [
                "security"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/security",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTE",
              "manager": {
                "name": "Kyle Rockman",
                "email": "kyle.rockman@mac.com"
              },
              "members": {
                "nodes": [
                  {
                    "name": "Andrew Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Ashley Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Drew Example",
                    "email": "example@opslevel.com"
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "Security",
              "responsibilities": null
            },
            {
              "alias": "vps",
              "aliases": [
                "vps"
              ],
              "contacts": [],
              "htmlUrl": "https://app.opslevel.com/teams/vps",
              "id": "Z2lkOi8vb3BzbGV2ZWwvVGVhbS84NTI",
              "manager": {
                "name": "Kyle Rockman",
                "email": "kyle.rockman@mac.com"
              },
              "members": {
                "nodes": [
                  {
                    "name": "Andrew Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Ashley Example",
                    "email": "example@opslevel.com"
                  },
                  {
                    "name": "Drew Example",
                    "email": "example@opslevel.com"
                  }
                ],
                "pageInfo": {
                  "hasNextPage": false,
                  "hasPreviousPage": false,
                  "startCursor": "MQ",
                  "endCursor": "Mw"
                }
              },
              "name": "VPs",
              "responsibilities": null
            }
          ],
          "pageInfo": {
            "hasNextPage": false,
            "hasPreviousPage": true,
            "startCursor": "NDc",
            "endCursor": "EOf"
          }
        }
      }
    }
  }`},
	}
	client := APaginatedTestClient(t, "team/list", requests...)
	// Act
	response, err := client.ListTeams(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 5, len(result))
	autopilot.Equals(t, "devops", result[0].Alias)
	autopilot.Equals(t, "developers", result[1].Alias)
	autopilot.Equals(t, "security", result[3].Alias)
	autopilot.Equals(t, "vps", result[4].Alias)
	autopilot.Equals(t, "Own Infra & Tools.", result[0].Responsibilities)
}

func TestUpdateTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/update")
	// Act
	result, err := client.UpdateTeam(ol.TeamUpdateInput{
		Id:               "Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ",
		ManagerEmail:     "ken@example.com",
		Responsibilities: "Foo & bar",
		Group:            ol.NewIdentifier("test_group"),
	})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Example", result.Name)
	autopilot.Equals(t, "ken@example.com", result.Manager.Email)
	autopilot.Equals(t, "Foo & bar", result.Responsibilities)
	autopilot.Equals(t, "test_group", result.Group.Alias)
}

func TestDeleteTeam(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/delete")
	// Act
	err := client.DeleteTeam("Z2lkOi8vb3BzbGV2ZWwvVGVhbS83NzQ")
	// Assert
	autopilot.Ok(t, err)
}

func TestDeleteTeamWithAlias(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "team/delete", "team/delete_with_alias")
	// Act
	err := client.DeleteTeamWithAlias("example")
	// Assert
	autopilot.Ok(t, err)
}

func TestTeamAddMember(t *testing.T) {
	// Arrange
	client1 := getWithAliasTestClient(t)
	client2 := ATestClient(t, "team/add_member")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddMember(&team.TeamId, "john@example.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamRemoveMember(t *testing.T) {
	// Arrange
	client1 := getWithAliasTestClient(t)
	client2 := ATestClient(t, "team/remove_member")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.RemoveMember(&team.TeamId, "john@example.com")
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 1, len(result))
}

func TestTeamAddContact(t *testing.T) {
	// Arrange
	client1 := getWithAliasTestClient(t)
	client2 := ATestClient(t, "team/add_contact")
	// Act
	team, _ := client1.GetTeamWithAlias("example")
	result, err := client2.AddContact(string(team.TeamId.Id), ol.CreateContactSlack("#general", ""))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "#general", result.Address)
}

func TestTeamUpdateContact(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/update_contact")
	// Act
	result, err := client.UpdateContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", ol.CreateContactSlack("#general", "Main Channel"))
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamUpdateContactWithTypeNil(t *testing.T) {
	// Arrange
	client := ATestClientAlt(t, "team/update_contact", "team/update_contact_nil_type")
	// Act
	result, err := client.UpdateContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4", ol.ContactInput{Address: "#general", DisplayName: "Main Channel"})
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, "Main Channel", result.DisplayName)
}

func TestTeamRemoveContact(t *testing.T) {
	// Arrange
	client := ATestClient(t, "team/remove_contact")
	// Act
	err := client.RemoveContact("Z2lkOi8vb3BzbGV2ZWwvVXNlci8zMDY4")
	// Assert
	autopilot.Ok(t, err)
}
