# vsts-goclient

A Golang client for Visual Studios Team Services

### Install

Go get it:

    $ go get github.com/samkreter/vsts-goclient

Then import the specific client:

    import vstsgit "github.com/samkreter/vsts-goclient/git"

### Making a request

1. To authenticate, Get a VSTS personal access token. This can be generated on your account page under the security tab.

2. Create an auth context using your username and the personal access token as the password:

   ```go
       auth := context.WithValue(context.Background(), vstsgit.ContextBasicAuth, vstsgit.BasicAuth{
           UserName: "username",
           Password: "personalAccessToken",
       })
   ```

3. Create a client and change the basepath to your vsts account:

   ```go
       cfg := vsts.NewConfiguration()
       client := vsts.NewAPIClient(cfg)
       client.ChangeBasePath("https://youraccount.visualstudio.com")
   ```

4. Use the auth to make calls using the client

   ```go
       prs, _, err := client.PullRequestsApi.GetPullRequests(auth, "repositoryName", "ProjectName", "4.0", nil)
       if err != nil {
           log.Fatal(err)
       }
       fmt.Println(prs)
   ```

### Contributing

Please feel free to contribute! Bug fixes are more than welcome any time.
If you'd like to change an existing implementation or see a new feature,
open an issue first so we can discuss it.
