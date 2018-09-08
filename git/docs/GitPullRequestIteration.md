# GitPullRequestIteration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | A collection of related REST reference links. | [optional] [default to null]
**Author** | [***IdentityRef**](IdentityRef.md) | Author of the pull request iteration. | [optional] [default to null]
**ChangeList** | [**[]GitPullRequestChange**](GitPullRequestChange.md) | Changes included with the pull request iteration. | [optional] [default to null]
**Commits** | [**[]GitCommitRef**](GitCommitRef.md) | The commits included with the pull request iteration. | [optional] [default to null]
**CommonRefCommit** | [***GitCommitRef**](GitCommitRef.md) | The first common Git commit of the source and target refs. | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The creation date of the pull request iteration. | [optional] [default to null]
**Description** | **string** | Description of the pull request iteration. | [optional] [default to null]
**HasMoreCommits** | **bool** | Indicates if the Commits property contains a truncated list of commits in this pull request iteration. | [optional] [default to null]
**Id** | **int32** | ID of the pull request iteration. Iterations are created as a result of creating and pushing updates to a pull request. | [optional] [default to null]
**NewTargetRefName** | **string** | If the iteration reason is Retarget, this is the refName of the new target | [optional] [default to null]
**Push** | [***GitPushRef**](GitPushRef.md) | The Git push information associated with this pull request iteration. | [optional] [default to null]
**Reason** | **string** | The reason for which the pull request iteration was created. | [optional] [default to null]
**SourceRefCommit** | [***GitCommitRef**](GitCommitRef.md) | The source Git commit of this iteration. | [optional] [default to null]
**TargetRefCommit** | [***GitCommitRef**](GitCommitRef.md) | The target Git commit of this iteration. | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | The updated date of the pull request iteration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


