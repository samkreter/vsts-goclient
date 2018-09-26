package vsts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	gitclient "github.com/samkreter/vsts-goclient/api/git"
)

type restClient struct {
	username   string
	token      string
	httpClient *http.Client
}

func (c *restClient) Do(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: %s", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return b, nil
}

// Client implementation for the vsts client
type Client struct {
	Instance   string
	Project    string
	Repo       string
	APIVersion string
	restClient *restClient
	apiClient  *gitclient.APIClient
	apiAuth    context.Context
}

// Config stores the configurations for the vsts client
type Config struct {
	Token          string
	Username       string
	APIVersion     string
	RepositoryName string
	Project        string
	Instance       string
}

// NewClient creates a new vsts client
func NewClient(config *Config) (*Client, error) {
	cfg := gitclient.NewConfiguration()
	apiClient := gitclient.NewAPIClient(cfg)
	apiClient.ChangeBasePath(fmt.Sprintf("https://%s", config.Instance))

	auth := context.WithValue(context.Background(), gitclient.ContextBasicAuth, gitclient.BasicAuth{
		UserName: config.Username,
		Password: config.Token,
	})

	// Note: Using the old version rest client until all api routes are updated to use the generated client
	rClient := &restClient{
		username:   config.Username,
		token:      config.Token,
		httpClient: &http.Client{},
	}

	return &Client{
		Instance:   config.Instance,
		Project:    config.Project,
		APIVersion: config.APIVersion,
		Repo:       config.RepositoryName,
		apiClient:  apiClient,
		apiAuth:    auth,
		restClient: rClient,
	}, nil
}

// GetBranches returns branches matching a branch name
func (c *Client) GetBranches(branchName string) (*Refs, error) {
	getBranchURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/refs/heads/{branch}?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{branch}", branchName,
		"{version}", "1.0")

	urlString := r.Replace(getBranchURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	branches := &Refs{}

	err = json.Unmarshal(b, branches)
	if err != nil {
		return nil, fmt.Errorf("json decoding error: %s", err)
	}

	return branches, nil
}

// GetBranch returns the matching branch. If there are multiple matching, returns the first
func (c *Client) GetBranch(branchName string) (*Ref, error) {
	getBranchURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/refs/heads/{branch}?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{branch}", branchName,
		"{version}", "1.0")

	urlString := r.Replace(getBranchURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	branches := Refs{}

	err = json.Unmarshal(b, &branches)
	if err != nil {
		return nil, err
	}

	fmt.Printf("master branches: %v\n", branches.Count)

	if branches.Count == 0 {
		return nil, fmt.Errorf("No %v branch found", branchName)
	}

	branch := branches.Value[0]
	for i := range branches.Value {
		if branches.Value[i].Name == branchName {
			branch = branches.Value[i]
			break
		}
	}
	return &branch, nil
}

// CreateBranch creates a new branch
func (c *Client) CreateBranch(newBranchName string, commitID string) error {
	newBranch := Branch{
		Name:        fmt.Sprintf("%s/%s", "refs/heads", newBranchName),
		OldObjectID: "0000000000000000000000000000000000000000",
		NewObjectID: commitID,
	}

	postBranchURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/refs?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "1.0")

	urlString := r.Replace(postBranchURLTemplate)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode([]Branch{newBranch})

	_, err := c.restClient.Do("POST", urlString, body)
	if err != nil {
		return err
	}

	return nil
}

// GetCommits retrieves commits for a branch in a specified time range
func (c *Client) GetCommits(branch string, startTime time.Time, endTime time.Time, versionPath string) (*Commits, error) {
	toText, _ := endTime.MarshalText()
	fromText, _ := startTime.MarshalText()
	fmt.Printf("Finding commits from %s to %s...\n", string(fromText), string(toText))

	getCommitsURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/commits?api-version={version}&branch={branch}&itemPath={versionPath}&fromDate={fromDateTime}&toDate={toDateTime}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{branch}", branch,
		"{versionPath}", versionPath,
		"{fromDateTime}", string(fromText),
		"{toDateTime}", string(toText),
		"{version}", "1.0")

	urlString := r.Replace(getCommitsURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	commits := &Commits{}

	json.Unmarshal(b, commits)
	return commits, err
}

// GetFileFromRepo returns the file at an empty interface
func (c *Client) GetFileFromRepo(branchName string, filePath string, target interface{}) error {
	getItemURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/items?api-version={version}&versionType={versionType}&version={versionValue}&scopePath={Path}&lastProcessedChange=true"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{versionType}", "branch",
		"{versionValue}", branchName,
		"{Path}", filePath,
		"{version}", "1.0")

	urlString := r.Replace(getItemURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &target); err != nil {
		return err
	}

	return nil
}

// CommitChanges commits the staged commits
func (c *Client) CommitChanges(changes interface{}, branchName, commitID, changePath, commitMessage string) error {

	content, err := json.MarshalIndent(changes, "", "    ")
	if err != nil {
		return err
	}

	postPushURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/pushes?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "2.0-preview")

	urlString := r.Replace(postPushURLTemplate)

	versionResetPush := push{
		RefUpdates: []refUpdate{
			{
				Name:        fmt.Sprintf("%s/%s", "refs/heads", branchName),
				OldObjectID: commitID,
			},
		},
		Commits: []pushCommit{
			{
				Comment: commitMessage,
				Changes: []change{
					{
						ChangeType: "edit",
						Item: item{
							Path: changePath,
						},
						NewContent: newContent{
							ContentType: "rawtext",
							Content:     string(content),
						},
					},
				},
			},
		},
	}

	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(versionResetPush)
	if err != nil {
		return err
	}

	_, err = c.restClient.Do("POST", urlString, body)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: Commit and pushed to branch %s\n", branchName)

	return nil
}

// GetBuildDefinitions returns the matching build definitions
func (c *Client) GetBuildDefinitions(definitionPath, definitionName string) (*Definitions, error) {
	getDefinitionsURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/build/definitions?api-version={version}&path={path}&name={definitionName}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{version}", "3.0-preview.2",
		"{path}", definitionPath,
		"{definitionName}", definitionName)

	urlString := r.Replace(getDefinitionsURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	defs := &Definitions{}
	if err := json.Unmarshal(b, defs); err != nil {
		return nil, err
	}

	return defs, nil
}

// GetBuilds returns the builds for a specific build definition
func (c *Client) GetBuilds(buildDefinitionID int) (*Builds, error) {
	getBuildsURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/build/builds?api-version={version}&definitions={defID}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{version}", "2.0",
		"{defID}", strconv.Itoa(buildDefinitionID))

	urlString := r.Replace(getBuildsURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	builds := &Builds{}
	err = json.Unmarshal(b, builds)

	return builds, err
}

// PostBuild queues a new build
func (c *Client) PostBuild(buildDefinitionID int, branch string, parameters interface{}) (*PostBuildResponse, error) {
	postBuildURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/build/builds?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{version}", "2.0")

	urlString := r.Replace(postBuildURLTemplate)

	params, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	relBuild := BuildReq{
		Definition: Definition{
			ID: buildDefinitionID,
		},
		SourceBranch: fmt.Sprintf("%s/%s", "refs/heads", branch),
		Parameters:   string(params),
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(relBuild)

	b, err := c.restClient.Do("POST", urlString, body)
	if err != nil {
		return nil, err
	}

	var buildResp PostBuildResponse
	if err := json.Unmarshal(b, &buildResp); err != nil {
		return nil, err
	}

	return &buildResp, nil
}

// SubmitPullRequest creates a pull request
func (c *Client) SubmitPullRequest(targetBranch string, sourceBranch string, title string, description string) error {
	postPullRequestURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/pullRequests?api-version={version}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "3.0-preview")

	urlString := r.Replace(postPullRequestURLTemplate)

	pullRequest := PullRequest{
		SourceRefName: fmt.Sprintf("%s/%s", "refs/heads", sourceBranch),
		TargetRefName: fmt.Sprintf("%s/%s", "refs/heads", targetBranch),
		Title:         title,
		Description:   description,
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(pullRequest)

	_, err := c.restClient.Do("POST", urlString, body)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: Starting PR from %s to %s...\n", sourceBranch, targetBranch)
	return nil
}

// GetDiffsBetweenBranches returns the diffs of two branches
func (c *Client) GetDiffsBetweenBranches(baseBranch string, targetBranch string) (*Diffs, error) {
	getDiffsURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/diffs/commits?api-version={version}&targetVersionType=branch&targetVersion={targetBranch}&baseVersionType=branch&baseVersion={baseBranch}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "1.0",
		"{baseBranch}", baseBranch,
		"{targetBranch}", targetBranch)

	urlString := r.Replace(getDiffsURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	diffs := &Diffs{}

	json.Unmarshal(b, &diffs)

	return diffs, nil
}

// GetCommit returns a specific commit
func (c *Client) GetCommit(commitID string) (*Commit, error) {
	getCommitURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/commits/{commitId}?api-version={version}&changeCount={changeCount}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "1.0",
		"{commitId}", commitID,
		"changeCount", "100")

	urlString := r.Replace(getCommitURLTemplate)

	b, err := c.restClient.Do("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	commit := &Commit{}
	if err := json.Unmarshal(b, commit); err != nil {
		return nil, err
	}

	return commit, err
}

// CompletePullRequest completes the specified pull request
func (c *Client) CompletePullRequest(pullRequestID int, commitID string, mergeMessage string, bypassPolicy bool, deleteSourceBranch bool) error {
	patchPullRequestURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/pullRequests/{pullRequest}?api-version={version}"

	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{pullRequest}", strconv.Itoa(pullRequestID),
		"{version}", "3.0-preview")

	urlString := r.Replace(patchPullRequestURLTemplate)

	patchPullRequest := patchPullRequest{
		Status: "completed",
		LastMergeSourceCommit: LastMergeSourceCommit{
			CommitID: commitID,
		},
		CompletionOptions: CompletionOptions{
			MergeCommitMessage: mergeMessage,
			SquashMerge:        strconv.FormatBool(true),
			DeleteSourceBranch: strconv.FormatBool(deleteSourceBranch),
			BypassPolicy:       strconv.FormatBool(bypassPolicy),
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(patchPullRequest)

	_, err := c.restClient.Do("PATCH", urlString, body)
	if err != nil {
		return err
	}

	fmt.Printf("INFO: Complete PR %v...\n", pullRequestID)
	return nil
}

// GetPullRequests returns the pull requests for la specific repository
func (c *Client) GetPullRequests() ([]gitclient.GitPullRequest, error) {
	prs, resp, err := c.apiClient.PullRequestsApi.GetPullRequests(c.apiAuth, c.Repo, c.Project, c.APIVersion, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return prs, nil
}

// AddThread Create a comment on the pull request
func (c *Client) AddThread(pullRequestID int32, comment string) error {
	body := gitclient.GitPullRequestCommentThread{
		Comments: []gitclient.Comment{
			{
				Content: comment,
			},
		},
	}

	_, resp, err := c.apiClient.PullRequestThreadsApi.Create(c.apiAuth, body, c.Repo, pullRequestID, c.Project, c.APIVersion)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return nil
}

// GetThreads lists all the threads on a pull request
func (c *Client) GetThreads(pullRequestID int32) ([]gitclient.GitPullRequestCommentThread, error) {
	threads, resp, err := c.apiClient.PullRequestThreadsApi.List(c.apiAuth, c.Repo, pullRequestID, c.Project, c.APIVersion, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return threads, nil
}

// AddReviewer Create a comment on the pull request
func (c *Client) AddReviewer(pullRequestID int32, reviewer gitclient.IdentityRefWithVote) error {

	_, resp, err := c.apiClient.PullRequestReviewersApi.CreatePullRequestReviewer(c.apiAuth, reviewer, c.Repo, pullRequestID, reviewer.ID, c.Project, c.APIVersion)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return nil
}

// AddReviewers Create a comment on the pull request
func (c *Client) AddReviewers(pullRequestID int32, reviewers []gitclient.IdentityRef) error {

	_, resp, err := c.apiClient.PullRequestReviewersApi.CreatePullRequestReviewers(c.apiAuth, reviewers, c.Repo, pullRequestID, c.Project, c.APIVersion)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	return nil
}
