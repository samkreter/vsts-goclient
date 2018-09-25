package main

import (
	"context"
	"fmt"
	"log"

	vsts "github.com/samkreter/vsts-goclient/v2/git"
)

func main() {

	cfg := vsts.NewConfiguration()
	client := vsts.NewAPIClient(cfg)
	client.ChangeBasePath("https://youraccount.visualstudio.com")

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
