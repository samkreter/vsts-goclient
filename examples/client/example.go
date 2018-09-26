package main

import (
	"fmt"
	"log"

	vsts "github.com/samkreter/vsts-goclient/client"
)

func main() {
	vstsConfig := &vsts.Config{
		Token:          "personal-access-token",
		Username:       "username",
		APIVersion:     "3.0",
		RepositoryName: "repo-name",
		Project:        "project-name",
		Instance:       "url-instance",
	}

	vstsClient, err := vsts.NewClient(vstsConfig)
	if err != nil {
		log.Fatal(err)
	}

	prs, err := vstsClient.GetPullRequests()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prs)
}
