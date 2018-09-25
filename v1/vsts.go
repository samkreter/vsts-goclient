package vsts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Client implementation for the vsts client
type Client struct {
	Username                 string `json:"username"`
	Token                    string `json:"token"`
	Instance                 string `json:"instance"`
	Project                  string `json:"project"`
	Repo                     string `json:"repo"`
	OnboardBuildDefinitionID int    `json:"onboardBuildDefinitionId"`
	ReleaseBranchPrefix      string `json:"releaseBranchPrefix"`
	httpClient               *http.Client
}

// NewClient creates a new vsts client
func NewClient(username, token, instance, project, repo string, httpClient *http.Client) (*Client, error) {
	return &Client{
		Username:   username,
		Token:      token,
		Instance:   instance,
		Project:    project,
		Repo:       repo,
		httpClient: httpClient,
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	branches := &Refs{}

	err = json.NewDecoder(resp.Body).Decode(branches)
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
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

	branches := Refs{}

	err = json.NewDecoder(resp.Body).Decode(&branches)
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

	req, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
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

	commits := &Commits{}

	json.NewDecoder(resp.Body).Decode(commits)
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
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

	bodyText, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyText, &target)

	return err
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

	req, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.httpClient.Do(req)
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

	defs := &Definitions{}
	err = json.NewDecoder(resp.Body).Decode(defs)

	return defs, err
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
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
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

	builds := &Builds{}
	err = json.NewDecoder(resp.Body).Decode(builds)

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
	req, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved non 200 status code of %d and body %s", resp.StatusCode, string(b))
	}

	var buildResp PostBuildResponse
	err = json.Unmarshal(b, &buildResp)
	if err != nil {
		return nil, err
	}

	return &buildResp, nil
}

// GetPullRequests retrieves all pull request from a source branch to a target branch
func (c *Client) GetPullRequests(targetBranch string, sourceBranch string) (*PullRequests, error) {
	getPullRequestsURLTemplate := "https://{instance}/DefaultCollection/{project}/_apis/git/repositories/{repository}/pullRequests?api-version={version}&status={status}&sourceRefName={sourceBranch}&targetRefName={targetBranch}"
	r := strings.NewReplacer(
		"{instance}", c.Instance,
		"{project}", c.Project,
		"{repository}", c.Repo,
		"{version}", "3.0-preview",
		"{status}", "Active",
		"{sourceBranch}", fmt.Sprintf("%s/%s", "refs/heads", sourceBranch),
		"{targetBranch}", fmt.Sprintf("%s/%s", "refs/heads", targetBranch))

	urlString := r.Replace(getPullRequestsURLTemplate)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
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

	pullRequests := &PullRequests{}

	err = json.NewDecoder(resp.Body).Decode(pullRequests)
	if err != nil {
		return nil, err
	}

	return pullRequests, nil
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
	req, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
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

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
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

	diffs := &Diffs{}

	json.NewDecoder(resp.Body).Decode(&diffs)

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
	commit := &Commit{}

	err := c.getJSONResponse("GET", urlString, commit)

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
	req, err := http.NewRequest("PATCH", urlString, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
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

	fmt.Printf("INFO: Complete PR %v...\n", pullRequestID)
	return nil
}

func (c *Client) getJSONResponse(method, url string, target interface{}) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	return err
}
