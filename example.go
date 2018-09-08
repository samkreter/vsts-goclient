package main

import (
	"context"
	"fmt"
	"log"

	vsts "github.com/samkreter/vsts-goclient/git"
)

func main() {

	cfg := vsts.NewConfiguration()
	cfg.BasePath = "basepath"
	client := vsts.NewAPIClient(cfg)

	auth := context.WithValue(context.Background(), vsts.ContextBasicAuth, vsts.BasicAuth{
		UserName: "username",
		Password: "passwordtoken",
	})
	prs, resp, err := client.PullRequestsApi.GetPullRequests(auth, "repo", "project", "4.0", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prs)
	fmt.Println(resp.StatusCode)
}
