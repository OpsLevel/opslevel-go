# OpsLevel Go Client Example
An example program that demonstrates interacting with the OpsLevel GraphQL API with an example Go client.

To use this you will need to have an OpsLevel account and create an API token through the UI.

Currently, this library is a work in a progress and only covers a small section of the API, but can be extended.

- teams.go provides an example query
- tags.go provides an example mutation

More information about using the GraphQL API can be found [here](https://www.opslevel.com/docs/api/graphql/).

# Example Usage
```go
import (
	"context"
	"fmt"
	opslevel "github.com/opslevel/opslevel-go"
	"log"
	"os"
)

const teamAlias = "dev_team"
const serviceAlias = "coffee_service"
const defaultUrl = "https://api.opslevel.com/graphql"

func main() {
	var authToken = os.Getenv("OPSLEVEL_TOKEN")
	client := opslevel.NewClient(authToken)

	team, err := client.GetTeam(context.TODO(), teamAlias)
	if err != nil {
		log.Fatal(err)
	}

	tag, err := client.CreateTag(context.TODO(), "team", team.Name, serviceAlias, "Service")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tag created", tag)
}
```