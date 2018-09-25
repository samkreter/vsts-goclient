package vsts

import "time"

// PostBuildResponse holds the api response from posting a new build
type PostBuildResponse struct {
	Links             PostBuildResponseLinks `json:"_links"`
	Tags              []interface{}          `json:"tags"`
	ValidationResults []interface{}          `json:"validationResults"`
	ID                int64                  `json:"id"`
	BuildNumber       string                 `json:"buildNumber"`
	Status            string                 `json:"status"`
	QueueTime         string                 `json:"queueTime"`
	URL               string                 `json:"url"`
	Definition        Definition             `json:"definition"`
	Project           Project                `json:"project"`
	URI               string                 `json:"uri"`
	SourceBranch      string                 `json:"sourceBranch"`
	Priority          string                 `json:"priority"`
	Reason            string                 `json:"reason"`
	RequestedFor      LastChangedBy          `json:"requestedFor"`
	RequestedBy       LastChangedBy          `json:"requestedBy"`
	LastChangedDate   string                 `json:"lastChangedDate"`
	LastChangedBy     LastChangedBy          `json:"lastChangedBy"`
	Parameters        string                 `json:"parameters"`
	KeepForever       bool                   `json:"keepForever"`
	RetainedByRelease bool                   `json:"retainedByRelease"`
	TriggeredByBuild  interface{}            `json:"triggeredByBuild"`
}

// Definition is the build definition
type Definition struct {
	Drafts      []interface{} `json:"drafts"`
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	URL         string        `json:"url"`
	URI         string        `json:"uri"`
	Path        string        `json:"path"`
	Type        string        `json:"type"`
	QueueStatus string        `json:"queueStatus"`
	Revision    int64         `json:"revision"`
	Project     Project       `json:"project"`
}

// Project holds info about the current project
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	State       string `json:"state"`
	Revision    int64  `json:"revision"`
	Visibility  string `json:"visibility"`
}

// LastChangedBy shows the last changed by status
type LastChangedBy struct {
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
	Links       struct {
		Avatar Badge `json:"avatar"`
	} `json:"_links"`
	ID         string `json:"id"`
	UniqueName string `json:"uniqueName"`
	ImageURL   string `json:"imageUrl"`
	Descriptor string `json:"descriptor"`
}

// Badge holds link information
type Badge struct {
	Href string `json:"href"`
}

// PostBuildResponseLinks store all links referenced to a build
type PostBuildResponseLinks struct {
	Self                    Badge `json:"self"`
	Web                     Badge `json:"web"`
	SourceVersionDisplayURI Badge `json:"sourceVersionDisplayUri"`
	Timeline                Badge `json:"timeline"`
	Badge                   Badge `json:"badge"`
}

// Ref holds the object references
type Ref struct {
	Name     string `json:"name"`
	ObjectID string `json:"objectId"`
	URL      string `json:"url"`
}

// Refs object for multiple refs
type Refs struct {
	Value []Ref `json:"value"`
	Count int   `json:"count"`
}

// Branch reference for a branch
type Branch struct {
	Name        string `json:"name"`
	OldObjectID string `json:"oldObjectId"`
	NewObjectID string `json:"newObjectId"`
}

// Commits stores multiple commits and metadata
type Commits struct {
	Count int `json:"count"`
	Value []struct {
		CommitID     string `json:"commitId"`
		Comment      string `json:"comment"`
		ChangeCounts struct {
			Add int `json:"Add"`
		} `json:"changeCounts"`
		URL       string `json:"url"`
		RemoteURL string `json:"remoteUrl"`
	} `json:"value"`
}

type refUpdate struct {
	Name        string `json:"name"`
	OldObjectID string `json:"oldObjectId"`
}

type item struct {
	Path string `json:"path"`
}

type newContent struct {
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

type change struct {
	ChangeType string     `json:"changeType"`
	Item       item       `json:"item"`
	NewContent newContent `json:"newContent"`
}

type pushCommit struct {
	Comment string   `json:"comment"`
	Changes []change `json:"changes"`
}

type push struct {
	RefUpdates []refUpdate  `json:"refUpdates"`
	Commits    []pushCommit `json:"commits"`
}

type version struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type root struct {
	Versions []version `xml:"versions>version"`
}

// BuildReq stores is a request for a build
type BuildReq struct {
	Definition   Definition `json:"definition"`
	SourceBranch string     `json:"sourceBranch"`
	Parameters   string     `json:"parameters"`
}

// Definitions stores multiple definitions
type Definitions struct {
	Count int          `json:"count"`
	Value []Definition `json:"value"`
}

// Builds hold multiple builds
type Builds struct {
	Count int `json:"count"`
	Value []struct {
		ID            int       `json:"id"`
		URL           string    `json:"url"`
		BuildNumber   string    `json:"buildNumber"`
		URI           string    `json:"uri"`
		SourceBranch  string    `json:"sourceBranch"`
		SourceVersion string    `json:"sourceVersion"`
		Status        string    `json:"status"`
		QueueTime     time.Time `json:"queueTime"`
		Priority      string    `json:"priority"`
		StartTime     time.Time `json:"startTime"`
		FinishTime    time.Time `json:"finishTime"`
		Reason        string    `json:"reason"`
		Result        string    `json:"result"`
		Parameters    string    `json:"parameters"`
		KeepForever   bool      `json:"keepForever"`
	} `json:"value"`
}

// PullRequests holds multiple pull requests
type PullRequests struct {
	Value []PullRequest `json:"value"`
	Count int           `json:"count"`
}

// PullRequest reference to a pull request
type PullRequest struct {
	PullRequestID      int       `json:"pullRequestId"`
	CodeReviewID       int       `json:"codeReviewId"`
	Status             string    `json:"status"`
	CreationDate       time.Time `json:"creationDate"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	SourceRefName      string    `json:"sourceRefName"`
	TargetRefName      string    `json:"targetRefName"`
	MergeStatus        string    `json:"mergeStatus"`
	MergeID            string    `json:"mergeId"`
	URL                string    `json:"url"`
	SupportsIterations bool      `json:"supportsIterations"`
}

// LastMergeSourceCommit commit id of the last merged commit
type LastMergeSourceCommit struct {
	CommitID string `json:"commitId"`
}

// CompletionOptions options for completing a pull request
type CompletionOptions struct {
	DeleteSourceBranch string `json:"deleteSourceBranch"`
	MergeCommitMessage string `json:"mergeCommitMessage"`
	SquashMerge        string `json:"squashMerge"`
	BypassPolicy       string `json:"bypassPolicy"`
}

type patchPullRequest struct {
	Status                string                `json:"status"`
	LastMergeSourceCommit LastMergeSourceCommit `json:"lastMergeSourceCommit"`
	CompletionOptions     CompletionOptions     `json:"completionOptions"`
}

// Diffs stores the diffs of two branches
type Diffs struct {
	AllChangesIncluded bool `json:"allChangesIncluded"`
	ChangeCounts       struct {
		Edit int `json:"Edit"`
	} `json:"changeCounts"`
	Changes []struct {
		Item struct {
			ObjectID         string `json:"objectId"`
			OriginalObjectID string `json:"originalObjectId"`
			GitObjectType    string `json:"gitObjectType"`
			CommitID         string `json:"commitId"`
			Path             string `json:"path"`
			IsFolder         bool   `json:"isFolder"`
			URL              string `json:"url"`
		} `json:"item"`
		ChangeType string `json:"changeType"`
	} `json:"changes"`
	CommonCommit string `json:"commonCommit"`
	BaseCommit   string `json:"baseCommit"`
	TargetCommit string `json:"targetCommit"`
	AheadCount   int    `json:"aheadCount"`
	BehindCount  int    `json:"behindCount"`
}

// Commit holds data about a commit
type Commit struct {
	CommitID     string `json:"commitId"`
	ChangeCounts struct {
		Edit int `json:"Edit"`
	} `json:"changeCounts"`
	Changes []struct {
		Item struct {
			ObjectID         string `json:"objectId"`
			OriginalObjectID string `json:"originalObjectId"`
			GitObjectType    string `json:"gitObjectType"`
			CommitID         string `json:"commitId"`
			Path             string `json:"path"`
			IsFolder         bool   `json:"isFolder"`
			URL              string `json:"url"`
		} `json:"item"`
		ChangeType string `json:"changeType"`
	} `json:"changes"`
}
