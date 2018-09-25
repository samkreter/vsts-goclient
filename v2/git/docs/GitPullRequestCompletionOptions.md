# GitPullRequestCompletionOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BypassPolicy** | **bool** | If true, policies will be explicitly bypassed while the pull request is completed. | [optional] [default to null]
**BypassReason** | **string** | If policies are bypassed, this reason is stored as to why bypass was used. | [optional] [default to null]
**DeleteSourceBranch** | **bool** | If true, the source branch of the pull request will be deleted after completion. | [optional] [default to null]
**MergeCommitMessage** | **string** | If set, this will be used as the commit message of the merge commit. | [optional] [default to null]
**SquashMerge** | **bool** | If true, the commits in the pull request will be squash-merged into the specified target branch on completion. | [optional] [default to null]
**TransitionWorkItems** | **bool** | If true, we will attempt to transition any work items linked to the pull request into the next logical state (i.e. Active -&gt; Resolved) | [optional] [default to null]
**TriggeredByAutoComplete** | **bool** | If true, the current completion attempt was triggered via auto-complete. Used internally. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


