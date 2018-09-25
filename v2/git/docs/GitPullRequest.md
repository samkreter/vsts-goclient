# GitPullRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | Links to other related objects. | [optional] [default to null]
**ArtifactId** | **string** | A string which uniquely identifies this pull request. To generate an artifact ID for a pull request, use this template: &#x60;&#x60;&#x60;vstfs:///Git/PullRequestId/{projectId}/{repositoryId}/{pullRequestId}&#x60;&#x60;&#x60; | [optional] [default to null]
**AutoCompleteSetBy** | [***IdentityRef**](IdentityRef.md) | If set, auto-complete is enabled for this pull request and this is the identity that enabled it. | [optional] [default to null]
**ClosedBy** | [***IdentityRef**](IdentityRef.md) | The user who closed the pull request. | [optional] [default to null]
**ClosedDate** | [**time.Time**](time.Time.md) | The date when the pull request was closed (completed, abandoned, or merged externally). | [optional] [default to null]
**CodeReviewId** | **int32** | The code review ID of the pull request. Used internally. | [optional] [default to null]
**Commits** | [**[]GitCommitRef**](GitCommitRef.md) | The commits contained in the pull request. | [optional] [default to null]
**CompletionOptions** | [***GitPullRequestCompletionOptions**](GitPullRequestCompletionOptions.md) | Options which affect how the pull request will be merged when it is completed. | [optional] [default to null]
**CompletionQueueTime** | [**time.Time**](time.Time.md) | The most recent date at which the pull request entered the queue to be completed. Used internally. | [optional] [default to null]
**CreatedBy** | [***IdentityRef**](IdentityRef.md) | The identity of the user who created the pull request. | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | The date when the pull request was created. | [optional] [default to null]
**Description** | **string** | The description of the pull request. | [optional] [default to null]
**ForkSource** | [***GitForkRef**](GitForkRef.md) | If this is a PR from a fork this will contain information about its source. | [optional] [default to null]
**Labels** | [**[]WebApiTagDefinition**](WebApiTagDefinition.md) | The labels associated with the pull request. | [optional] [default to null]
**LastMergeCommit** | [***GitCommitRef**](GitCommitRef.md) | The commit of the most recent pull request merge. If empty, the most recent merge is in progress or was unsuccessful. | [optional] [default to null]
**LastMergeSourceCommit** | [***GitCommitRef**](GitCommitRef.md) | The commit at the head of the source branch at the time of the last pull request merge. | [optional] [default to null]
**LastMergeTargetCommit** | [***GitCommitRef**](GitCommitRef.md) | The commit at the head of the target branch at the time of the last pull request merge. | [optional] [default to null]
**MergeFailureMessage** | **string** | If set, pull request merge failed for this reason. | [optional] [default to null]
**MergeFailureType** | **string** | The type of failure (if any) of the pull request merge. | [optional] [default to null]
**MergeId** | **string** | The ID of the job used to run the pull request merge. Used internally. | [optional] [default to null]
**MergeOptions** | [***GitPullRequestMergeOptions**](GitPullRequestMergeOptions.md) | Options used when the pull request merge runs. These are separate from completion options since completion happens only once and a new merge will run every time the source branch of the pull request changes. | [optional] [default to null]
**MergeStatus** | **string** | The current status of the pull request merge. | [optional] [default to null]
**PullRequestId** | **int32** | The ID of the pull request. | [optional] [default to null]
**RemoteUrl** | **string** | Used internally. | [optional] [default to null]
**Repository** | [***GitRepository**](GitRepository.md) | The repository containing the target branch of the pull request. | [optional] [default to null]
**Reviewers** | [**[]IdentityRefWithVote**](IdentityRefWithVote.md) | A list of reviewers on the pull request along with the state of their votes. | [optional] [default to null]
**SourceRefName** | **string** | The name of the source branch of the pull request. | [optional] [default to null]
**Status** | **string** | The status of the pull request. | [optional] [default to null]
**SupportsIterations** | **bool** | If true, this pull request supports multiple iterations. Iteration support means individual pushes to the source branch of the pull request can be reviewed and comments left in one iteration will be tracked across future iterations. | [optional] [default to null]
**TargetRefName** | **string** | The name of the target branch of the pull request. | [optional] [default to null]
**Title** | **string** | The title of the pull request. | [optional] [default to null]
**Url** | **string** | Used internally. | [optional] [default to null]
**WorkItemRefs** | [**[]ResourceRef**](ResourceRef.md) | Any work item references associated with this pull request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


