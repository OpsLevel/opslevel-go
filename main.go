package main

import (
	"client_example/opslevel"
	"context"
	"fmt"
	"log"
	"os"
)

const teamAlias = "test_team"
const serviceAlias = "gamma"
const defaultUrl = "https://app.opslevel.com/graphql"

func main() {
	var authToken = os.Getenv("OPSLEVEL_TOKEN")
	var url = os.Getenv("OPSLEVEL_GRAPHQL_URL")
	if url == "" {
		url = defaultUrl
	}
	client := opslevel.NewClient(url, authToken)

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
