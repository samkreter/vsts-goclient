# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 5.0-preview
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnnotatedTagsApi* | [**Create**](docs/AnnotatedTagsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/annotatedtags | 
*AnnotatedTagsApi* | [**Get**](docs/AnnotatedTagsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/annotatedtags/{objectId} | 
*BlobsApi* | [**GetBlob**](docs/BlobsApi.md#getblob) | **Get** /{project}/_apis/git/repositories/{repositoryId}/blobs/{sha1} | 
*BlobsApi* | [**GetBlobsZip**](docs/BlobsApi.md#getblobszip) | **Post** /{project}/_apis/git/repositories/{repositoryId}/blobs | 
*CherryPicksApi* | [**Create**](docs/CherryPicksApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks | 
*CherryPicksApi* | [**GetCherryPick**](docs/CherryPicksApi.md#getcherrypick) | **Get** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks/{cherryPickId} | 
*CherryPicksApi* | [**GetCherryPickForRefName**](docs/CherryPicksApi.md#getcherrypickforrefname) | **Get** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks | 
*CommitsApi* | [**Get**](docs/CommitsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId} | 
*CommitsApi* | [**GetChanges**](docs/CommitsApi.md#getchanges) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/changes | 
*CommitsApi* | [**GetCommitsBatch**](docs/CommitsApi.md#getcommitsbatch) | **Post** /{project}/_apis/git/repositories/{repositoryId}/commitsbatch | 
*CommitsApi* | [**GetPushCommits**](docs/CommitsApi.md#getpushcommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits | 
*DiffsApi* | [**Get**](docs/DiffsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/diffs/commits | 
*ForksApi* | [**Create**](docs/ForksApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests | 
*ForksApi* | [**Get**](docs/ForksApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests/{forkSyncOperationId} | 
*ForksApi* | [**GetForkSyncRequests**](docs/ForksApi.md#getforksyncrequests) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests | 
*ForksApi* | [**List**](docs/ForksApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forks/{collectionId} | 
*ImportRequestsApi* | [**Create**](docs/ImportRequestsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/importRequests | 
*ImportRequestsApi* | [**Get**](docs/ImportRequestsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/importRequests/{importRequestId} | 
*ImportRequestsApi* | [**Query**](docs/ImportRequestsApi.md#query) | **Get** /{project}/_apis/git/repositories/{repositoryId}/importRequests | 
*ImportRequestsApi* | [**Update**](docs/ImportRequestsApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/importRequests/{importRequestId} | 
*ItemsApi* | [**GetItemsBatch**](docs/ItemsApi.md#getitemsbatch) | **Post** /{project}/_apis/git/repositories/{repositoryId}/itemsbatch | 
*ItemsApi* | [**List**](docs/ItemsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/items | 
*MergeBasesApi* | [**List**](docs/MergeBasesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/commits/{commitId}/mergebases | 
*PolicyConfigurationsApi* | [**List**](docs/PolicyConfigurationsApi.md#list) | **Get** /{project}/_apis/git/policy/configurations | 
*PullRequestAttachmentsApi* | [**Create**](docs/PullRequestAttachmentsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/attachments/{fileName} | 
*PullRequestAttachmentsApi* | [**Delete**](docs/PullRequestAttachmentsApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/attachments/{fileName} | 
*PullRequestAttachmentsApi* | [**Get**](docs/PullRequestAttachmentsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/attachments/{fileName} | 
*PullRequestAttachmentsApi* | [**List**](docs/PullRequestAttachmentsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/attachments | 
*PullRequestCommentLikesApi* | [**Create**](docs/PullRequestCommentLikesApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 
*PullRequestCommentLikesApi* | [**Delete**](docs/PullRequestCommentLikesApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 
*PullRequestCommentLikesApi* | [**List**](docs/PullRequestCommentLikesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 
*PullRequestCommitsApi* | [**GetPullRequestCommits**](docs/PullRequestCommitsApi.md#getpullrequestcommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/commits | 
*PullRequestCommitsApi* | [**GetPullRequestIterationCommits**](docs/PullRequestCommitsApi.md#getpullrequestiterationcommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/commits | 
*PullRequestIterationChangesApi* | [**Get**](docs/PullRequestIterationChangesApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/changes | 
*PullRequestIterationStatusesApi* | [**Create**](docs/PullRequestIterationStatusesApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 
*PullRequestIterationStatusesApi* | [**Delete**](docs/PullRequestIterationStatusesApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses/{statusId} | 
*PullRequestIterationStatusesApi* | [**Get**](docs/PullRequestIterationStatusesApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses/{statusId} | 
*PullRequestIterationStatusesApi* | [**List**](docs/PullRequestIterationStatusesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 
*PullRequestIterationStatusesApi* | [**Update**](docs/PullRequestIterationStatusesApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 
*PullRequestIterationsApi* | [**Get**](docs/PullRequestIterationsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId} | 
*PullRequestIterationsApi* | [**List**](docs/PullRequestIterationsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations | 
*PullRequestLabelsApi* | [**Create**](docs/PullRequestLabelsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels | 
*PullRequestLabelsApi* | [**Delete**](docs/PullRequestLabelsApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels/{labelIdOrName} | 
*PullRequestLabelsApi* | [**Get**](docs/PullRequestLabelsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels/{labelIdOrName} | 
*PullRequestLabelsApi* | [**List**](docs/PullRequestLabelsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels | 
*PullRequestPropertiesApi* | [**List**](docs/PullRequestPropertiesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/properties | 
*PullRequestPropertiesApi* | [**Update**](docs/PullRequestPropertiesApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/properties | 
*PullRequestQueryApi* | [**Get**](docs/PullRequestQueryApi.md#get) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullrequestquery | 
*PullRequestReviewersApi* | [**CreatePullRequestReviewer**](docs/PullRequestReviewersApi.md#createpullrequestreviewer) | **Put** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
*PullRequestReviewersApi* | [**CreatePullRequestReviewers**](docs/PullRequestReviewersApi.md#createpullrequestreviewers) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 
*PullRequestReviewersApi* | [**Delete**](docs/PullRequestReviewersApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
*PullRequestReviewersApi* | [**Get**](docs/PullRequestReviewersApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
*PullRequestReviewersApi* | [**List**](docs/PullRequestReviewersApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 
*PullRequestReviewersApi* | [**Update**](docs/PullRequestReviewersApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 
*PullRequestShareApi* | [**SharePullRequest**](docs/PullRequestShareApi.md#sharepullrequest) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/share | 
*PullRequestStatusesApi* | [**Create**](docs/PullRequestStatusesApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/statuses | 
*PullRequestStatusesApi* | [**Delete**](docs/PullRequestStatusesApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/statuses/{statusId} | 
*PullRequestStatusesApi* | [**Get**](docs/PullRequestStatusesApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/statuses/{statusId} | 
*PullRequestStatusesApi* | [**List**](docs/PullRequestStatusesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/statuses | 
*PullRequestStatusesApi* | [**Update**](docs/PullRequestStatusesApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/statuses | 
*PullRequestThreadCommentsApi* | [**Create**](docs/PullRequestThreadCommentsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments | 
*PullRequestThreadCommentsApi* | [**Delete**](docs/PullRequestThreadCommentsApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId} | 
*PullRequestThreadCommentsApi* | [**Get**](docs/PullRequestThreadCommentsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId} | 
*PullRequestThreadCommentsApi* | [**List**](docs/PullRequestThreadCommentsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments | 
*PullRequestThreadCommentsApi* | [**Update**](docs/PullRequestThreadCommentsApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId} | 
*PullRequestThreadsApi* | [**Create**](docs/PullRequestThreadsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads | 
*PullRequestThreadsApi* | [**Get**](docs/PullRequestThreadsApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId} | 
*PullRequestThreadsApi* | [**List**](docs/PullRequestThreadsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads | 
*PullRequestThreadsApi* | [**Update**](docs/PullRequestThreadsApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId} | 
*PullRequestWorkItemsApi* | [**List**](docs/PullRequestWorkItemsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/workitems | 
*PullRequestsApi* | [**Create**](docs/PullRequestsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullrequests | 
*PullRequestsApi* | [**GetPullRequest**](docs/PullRequestsApi.md#getpullrequest) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId} | 
*PullRequestsApi* | [**GetPullRequestById**](docs/PullRequestsApi.md#getpullrequestbyid) | **Get** /{project}/_apis/git/pullrequests/{pullRequestId} | 
*PullRequestsApi* | [**GetPullRequests**](docs/PullRequestsApi.md#getpullrequests) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullrequests | 
*PullRequestsApi* | [**GetPullRequestsByProject**](docs/PullRequestsApi.md#getpullrequestsbyproject) | **Get** /{project}/_apis/git/pullrequests | 
*PullRequestsApi* | [**Update**](docs/PullRequestsApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId} | 
*PushesApi* | [**Create**](docs/PushesApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pushes | 
*PushesApi* | [**Get**](docs/PushesApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pushes/{pushId} | 
*PushesApi* | [**List**](docs/PushesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pushes | 
*RefsApi* | [**List**](docs/RefsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/refs | 
*RefsApi* | [**UpdateRef**](docs/RefsApi.md#updateref) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/refs | 
*RefsApi* | [**UpdateRefs**](docs/RefsApi.md#updaterefs) | **Post** /{project}/_apis/git/repositories/{repositoryId}/refs | 
*RefsFavoritesApi* | [**Create**](docs/RefsFavoritesApi.md#create) | **Post** /{project}/_apis/git/favorites/refs | 
*RefsFavoritesApi* | [**Delete**](docs/RefsFavoritesApi.md#delete) | **Delete** /{project}/_apis/git/favorites/refs/{favoriteId} | 
*RefsFavoritesApi* | [**Get**](docs/RefsFavoritesApi.md#get) | **Get** /{project}/_apis/git/favorites/refs/{favoriteId} | 
*RefsFavoritesApi* | [**List**](docs/RefsFavoritesApi.md#list) | **Get** /{project}/_apis/git/favorites/refs | 
*RepositoriesApi* | [**Create**](docs/RepositoriesApi.md#create) | **Post** /{project}/_apis/git/repositories | 
*RepositoriesApi* | [**Delete**](docs/RepositoriesApi.md#delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId} | 
*RepositoriesApi* | [**DeleteRepositoryFromRecycleBin**](docs/RepositoriesApi.md#deleterepositoryfromrecyclebin) | **Delete** /{project}/_apis/git/recycleBin/repositories/{repositoryId} | 
*RepositoriesApi* | [**GetDeletedRepositories**](docs/RepositoriesApi.md#getdeletedrepositories) | **Get** /{project}/_apis/git/deletedrepositories | 
*RepositoriesApi* | [**GetRecycleBinRepositories**](docs/RepositoriesApi.md#getrecyclebinrepositories) | **Get** /{project}/_apis/git/recycleBin/repositories | 
*RepositoriesApi* | [**GetRepository**](docs/RepositoriesApi.md#getrepository) | **Get** /{project}/_apis/git/repositories/{repositoryId} | 
*RepositoriesApi* | [**List**](docs/RepositoriesApi.md#list) | **Get** /{project}/_apis/git/repositories | 
*RepositoriesApi* | [**RestoreRepositoryFromRecycleBin**](docs/RepositoriesApi.md#restorerepositoryfromrecyclebin) | **Patch** /{project}/_apis/git/recycleBin/repositories/{repositoryId} | 
*RepositoriesApi* | [**Update**](docs/RepositoriesApi.md#update) | **Patch** /{project}/_apis/git/repositories/{repositoryId} | 
*RevertsApi* | [**Create**](docs/RevertsApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/reverts | 
*RevertsApi* | [**GetRevert**](docs/RevertsApi.md#getrevert) | **Get** /{project}/_apis/git/repositories/{repositoryId}/reverts/{revertId} | 
*RevertsApi* | [**GetRevertForRefName**](docs/RevertsApi.md#getrevertforrefname) | **Get** /{project}/_apis/git/repositories/{repositoryId}/reverts | 
*StatsApi* | [**List**](docs/StatsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/stats/branches | 
*StatusesApi* | [**Create**](docs/StatusesApi.md#create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/statuses | 
*StatusesApi* | [**List**](docs/StatusesApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/statuses | 
*SuggestionsApi* | [**List**](docs/SuggestionsApi.md#list) | **Get** /{project}/_apis/git/repositories/{repositoryId}/suggestions | 
*TreesApi* | [**Get**](docs/TreesApi.md#get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/trees/{sha1} | 


## Documentation For Models

 - [AssociatedWorkItem](docs/AssociatedWorkItem.md)
 - [AsyncGitOperationNotification](docs/AsyncGitOperationNotification.md)
 - [Attachment](docs/Attachment.md)
 - [Change](docs/Change.md)
 - [ChangeCountDictionary](docs/ChangeCountDictionary.md)
 - [ChangeList](docs/ChangeList.md)
 - [ChangeListSearchCriteria](docs/ChangeListSearchCriteria.md)
 - [CheckinNote](docs/CheckinNote.md)
 - [Comment](docs/Comment.md)
 - [CommentIterationContext](docs/CommentIterationContext.md)
 - [CommentPosition](docs/CommentPosition.md)
 - [CommentThread](docs/CommentThread.md)
 - [CommentThreadContext](docs/CommentThreadContext.md)
 - [CommentTrackingCriteria](docs/CommentTrackingCriteria.md)
 - [FileContentMetadata](docs/FileContentMetadata.md)
 - [GitAnnotatedTag](docs/GitAnnotatedTag.md)
 - [GitAsyncRefOperation](docs/GitAsyncRefOperation.md)
 - [GitAsyncRefOperationDetail](docs/GitAsyncRefOperationDetail.md)
 - [GitAsyncRefOperationParameters](docs/GitAsyncRefOperationParameters.md)
 - [GitAsyncRefOperationSource](docs/GitAsyncRefOperationSource.md)
 - [GitBlobRef](docs/GitBlobRef.md)
 - [GitBranchStats](docs/GitBranchStats.md)
 - [GitCommitChanges](docs/GitCommitChanges.md)
 - [GitCommitDiffs](docs/GitCommitDiffs.md)
 - [GitCommitRef](docs/GitCommitRef.md)
 - [GitCommitToCreate](docs/GitCommitToCreate.md)
 - [GitConflict](docs/GitConflict.md)
 - [GitConflictUpdateResult](docs/GitConflictUpdateResult.md)
 - [GitDeletedRepository](docs/GitDeletedRepository.md)
 - [GitFilePathsCollection](docs/GitFilePathsCollection.md)
 - [GitForkOperationStatusDetail](docs/GitForkOperationStatusDetail.md)
 - [GitForkSyncRequest](docs/GitForkSyncRequest.md)
 - [GitForkSyncRequestParameters](docs/GitForkSyncRequestParameters.md)
 - [GitImportFailedEvent](docs/GitImportFailedEvent.md)
 - [GitImportGitSource](docs/GitImportGitSource.md)
 - [GitImportRequest](docs/GitImportRequest.md)
 - [GitImportRequestParameters](docs/GitImportRequestParameters.md)
 - [GitImportStatusDetail](docs/GitImportStatusDetail.md)
 - [GitImportSucceededEvent](docs/GitImportSucceededEvent.md)
 - [GitImportTfvcSource](docs/GitImportTfvcSource.md)
 - [GitItemDescriptor](docs/GitItemDescriptor.md)
 - [GitItemRequestData](docs/GitItemRequestData.md)
 - [GitLastChangeItem](docs/GitLastChangeItem.md)
 - [GitLastChangeTreeItems](docs/GitLastChangeTreeItems.md)
 - [GitMergeOriginRef](docs/GitMergeOriginRef.md)
 - [GitObject](docs/GitObject.md)
 - [GitPathAction](docs/GitPathAction.md)
 - [GitPathToItemsCollection](docs/GitPathToItemsCollection.md)
 - [GitPullRequest](docs/GitPullRequest.md)
 - [GitPullRequestCommentThreadContext](docs/GitPullRequestCommentThreadContext.md)
 - [GitPullRequestCompletionOptions](docs/GitPullRequestCompletionOptions.md)
 - [GitPullRequestIteration](docs/GitPullRequestIteration.md)
 - [GitPullRequestIterationChanges](docs/GitPullRequestIterationChanges.md)
 - [GitPullRequestMergeOptions](docs/GitPullRequestMergeOptions.md)
 - [GitPullRequestQuery](docs/GitPullRequestQuery.md)
 - [GitPullRequestQueryInput](docs/GitPullRequestQueryInput.md)
 - [GitPullRequestReviewFileContentInfo](docs/GitPullRequestReviewFileContentInfo.md)
 - [GitPullRequestSearchCriteria](docs/GitPullRequestSearchCriteria.md)
 - [GitPushEventData](docs/GitPushEventData.md)
 - [GitPushRef](docs/GitPushRef.md)
 - [GitPushSearchCriteria](docs/GitPushSearchCriteria.md)
 - [GitQueryBranchStatsCriteria](docs/GitQueryBranchStatsCriteria.md)
 - [GitQueryCommitsCriteria](docs/GitQueryCommitsCriteria.md)
 - [GitQueryRefsCriteria](docs/GitQueryRefsCriteria.md)
 - [GitRecycleBinRepositoryDetails](docs/GitRecycleBinRepositoryDetails.md)
 - [GitRef](docs/GitRef.md)
 - [GitRefFavorite](docs/GitRefFavorite.md)
 - [GitRefUpdate](docs/GitRefUpdate.md)
 - [GitRefUpdateResult](docs/GitRefUpdateResult.md)
 - [GitRepository](docs/GitRepository.md)
 - [GitRepositoryCreateOptions](docs/GitRepositoryCreateOptions.md)
 - [GitRepositoryRef](docs/GitRepositoryRef.md)
 - [GitRepositoryStats](docs/GitRepositoryStats.md)
 - [GitResolutionMergeContent](docs/GitResolutionMergeContent.md)
 - [GitResolutionPathConflict](docs/GitResolutionPathConflict.md)
 - [GitResolutionPickOneAction](docs/GitResolutionPickOneAction.md)
 - [GitStatus](docs/GitStatus.md)
 - [GitStatusContext](docs/GitStatusContext.md)
 - [GitSuggestion](docs/GitSuggestion.md)
 - [GitTemplate](docs/GitTemplate.md)
 - [GitTreeDiff](docs/GitTreeDiff.md)
 - [GitTreeDiffEntry](docs/GitTreeDiffEntry.md)
 - [GitTreeDiffResponse](docs/GitTreeDiffResponse.md)
 - [GitTreeEntryRef](docs/GitTreeEntryRef.md)
 - [GitTreeRef](docs/GitTreeRef.md)
 - [GitUserDate](docs/GitUserDate.md)
 - [GitVersionDescriptor](docs/GitVersionDescriptor.md)
 - [GlobalGitRepositoryKey](docs/GlobalGitRepositoryKey.md)
 - [GraphSubjectBase](docs/GraphSubjectBase.md)
 - [HistoryEntry](docs/HistoryEntry.md)
 - [ImportRepositoryValidation](docs/ImportRepositoryValidation.md)
 - [IncludedGitCommit](docs/IncludedGitCommit.md)
 - [ItemContent](docs/ItemContent.md)
 - [ItemDetailsOptions](docs/ItemDetailsOptions.md)
 - [ItemModel](docs/ItemModel.md)
 - [JsonPatchDocument](docs/JsonPatchDocument.md)
 - [JsonPatchOperation](docs/JsonPatchOperation.md)
 - [PolicyConfigurationRef](docs/PolicyConfigurationRef.md)
 - [PolicyTypeRef](docs/PolicyTypeRef.md)
 - [PropertiesCollection](docs/PropertiesCollection.md)
 - [PullRequestTabExtensionConfig](docs/PullRequestTabExtensionConfig.md)
 - [RealTimePullRequestEvent](docs/RealTimePullRequestEvent.md)
 - [ReferenceLinks](docs/ReferenceLinks.md)
 - [ResourceRef](docs/ResourceRef.md)
 - [ShareNotificationContext](docs/ShareNotificationContext.md)
 - [SourceToTargetRef](docs/SourceToTargetRef.md)
 - [SupportedIde](docs/SupportedIde.md)
 - [TeamProjectCollectionReference](docs/TeamProjectCollectionReference.md)
 - [TeamProjectReference](docs/TeamProjectReference.md)
 - [TfvcBranchMapping](docs/TfvcBranchMapping.md)
 - [TfvcChangesetRef](docs/TfvcChangesetRef.md)
 - [TfvcChangesetSearchCriteria](docs/TfvcChangesetSearchCriteria.md)
 - [TfvcChangesetsRequestData](docs/TfvcChangesetsRequestData.md)
 - [TfvcCheckinEventData](docs/TfvcCheckinEventData.md)
 - [TfvcItemDescriptor](docs/TfvcItemDescriptor.md)
 - [TfvcItemRequestData](docs/TfvcItemRequestData.md)
 - [TfvcLabelRef](docs/TfvcLabelRef.md)
 - [TfvcLabelRequestData](docs/TfvcLabelRequestData.md)
 - [TfvcMappingFilter](docs/TfvcMappingFilter.md)
 - [TfvcMergeSource](docs/TfvcMergeSource.md)
 - [TfvcPolicyFailureInfo](docs/TfvcPolicyFailureInfo.md)
 - [TfvcPolicyOverrideInfo](docs/TfvcPolicyOverrideInfo.md)
 - [TfvcShallowBranchRef](docs/TfvcShallowBranchRef.md)
 - [TfvcShelvesetRef](docs/TfvcShelvesetRef.md)
 - [TfvcShelvesetRequestData](docs/TfvcShelvesetRequestData.md)
 - [TfvcStatistics](docs/TfvcStatistics.md)
 - [TfvcVersionDescriptor](docs/TfvcVersionDescriptor.md)
 - [UpdateRefsRequest](docs/UpdateRefsRequest.md)
 - [VersionControlProjectInfo](docs/VersionControlProjectInfo.md)
 - [VssJsonCollectionWrapperBase](docs/VssJsonCollectionWrapperBase.md)
 - [WebApiCreateTagRequestData](docs/WebApiCreateTagRequestData.md)
 - [WebApiTagDefinition](docs/WebApiTagDefinition.md)
 - [AsyncRefOperationCommitLevelEventNotification](docs/AsyncRefOperationCommitLevelEventNotification.md)
 - [AsyncRefOperationCompletedNotification](docs/AsyncRefOperationCompletedNotification.md)
 - [AsyncRefOperationGeneralFailureNotification](docs/AsyncRefOperationGeneralFailureNotification.md)
 - [AsyncRefOperationTimeoutNotification](docs/AsyncRefOperationTimeoutNotification.md)
 - [AutoCompleteUpdatedEvent](docs/AutoCompleteUpdatedEvent.md)
 - [BranchUpdatedEvent](docs/BranchUpdatedEvent.md)
 - [CompletionErrorsEvent](docs/CompletionErrorsEvent.md)
 - [DiscussionsUpdatedEvent](docs/DiscussionsUpdatedEvent.md)
 - [GitBaseVersionDescriptor](docs/GitBaseVersionDescriptor.md)
 - [GitChange](docs/GitChange.md)
 - [GitCherryPick](docs/GitCherryPick.md)
 - [GitCommit](docs/GitCommit.md)
 - [GitConflictAddAdd](docs/GitConflictAddAdd.md)
 - [GitConflictAddRename](docs/GitConflictAddRename.md)
 - [GitConflictDeleteEdit](docs/GitConflictDeleteEdit.md)
 - [GitConflictDeleteRename](docs/GitConflictDeleteRename.md)
 - [GitConflictDirectoryFile](docs/GitConflictDirectoryFile.md)
 - [GitConflictEditDelete](docs/GitConflictEditDelete.md)
 - [GitConflictEditEdit](docs/GitConflictEditEdit.md)
 - [GitConflictFileDirectory](docs/GitConflictFileDirectory.md)
 - [GitConflictRename1to2](docs/GitConflictRename1to2.md)
 - [GitConflictRename2to1](docs/GitConflictRename2to1.md)
 - [GitConflictRenameAdd](docs/GitConflictRenameAdd.md)
 - [GitConflictRenameDelete](docs/GitConflictRenameDelete.md)
 - [GitConflictRenameRename](docs/GitConflictRenameRename.md)
 - [GitForkRef](docs/GitForkRef.md)
 - [GitForkTeamProjectReference](docs/GitForkTeamProjectReference.md)
 - [GitItem](docs/GitItem.md)
 - [GitPullRequestCommentThread](docs/GitPullRequestCommentThread.md)
 - [GitPullRequestStatus](docs/GitPullRequestStatus.md)
 - [GitPush](docs/GitPush.md)
 - [GitResolutionRename1to2](docs/GitResolutionRename1to2.md)
 - [GitRevert](docs/GitRevert.md)
 - [GitTargetVersionDescriptor](docs/GitTargetVersionDescriptor.md)
 - [IdentityRef](docs/IdentityRef.md)
 - [LabelsUpdatedEvent](docs/LabelsUpdatedEvent.md)
 - [MergeCompletedEvent](docs/MergeCompletedEvent.md)
 - [PolicyEvaluationUpdatedEvent](docs/PolicyEvaluationUpdatedEvent.md)
 - [PullRequestCreatedEvent](docs/PullRequestCreatedEvent.md)
 - [RetargetEvent](docs/RetargetEvent.md)
 - [ReviewerVoteUpdatedEvent](docs/ReviewerVoteUpdatedEvent.md)
 - [ReviewersUpdatedEvent](docs/ReviewersUpdatedEvent.md)
 - [ReviewersVotesResetEvent](docs/ReviewersVotesResetEvent.md)
 - [StatusAddedEvent](docs/StatusAddedEvent.md)
 - [StatusUpdatedEvent](docs/StatusUpdatedEvent.md)
 - [StatusesDeletedEvent](docs/StatusesDeletedEvent.md)
 - [TfvcBranchRef](docs/TfvcBranchRef.md)
 - [TfvcChange](docs/TfvcChange.md)
 - [TfvcChangeset](docs/TfvcChangeset.md)
 - [TfvcHistoryEntry](docs/TfvcHistoryEntry.md)
 - [TfvcItem](docs/TfvcItem.md)
 - [TfvcLabel](docs/TfvcLabel.md)
 - [TfvcShelveset](docs/TfvcShelveset.md)
 - [TitleDescriptionUpdatedEvent](docs/TitleDescriptionUpdatedEvent.md)
 - [VersionedPolicyConfigurationRef](docs/VersionedPolicyConfigurationRef.md)
 - [VssJsonCollectionWrapper](docs/VssJsonCollectionWrapper.md)
 - [AsyncRefOperationConflictNotification](docs/AsyncRefOperationConflictNotification.md)
 - [AsyncRefOperationProgressNotification](docs/AsyncRefOperationProgressNotification.md)
 - [GitPullRequestChange](docs/GitPullRequestChange.md)
 - [IdentityRefWithVote](docs/IdentityRefWithVote.md)
 - [PolicyConfiguration](docs/PolicyConfiguration.md)
 - [TfvcBranch](docs/TfvcBranch.md)
 - [TfvcItemPreviousHash](docs/TfvcItemPreviousHash.md)


## Documentation For Authorization

## accessToken
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```
## oauth2
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.vssps.visualstudio.com/oauth2/authorize&response_type=Assertion
- **Scopes**: 
 - **vso.code**: Grants the ability to read source code and metadata about commits, changesets, branches, and other version control artifacts. Also grants the ability to search code and get notified about version control events via service hooks.
 - **vso.code_manage**: Grants the ability to read, update, and delete source code, access metadata about commits, changesets, branches, and other version control artifacts. Also grants the ability to create and manage code repositories, create and manage pull requests and code reviews, and to receive notifications about version control events via service hooks.
 - **vso.code_write**: Grants the ability to read, update, and delete source code, access metadata about commits, changesets, branches, and other version control artifacts. Also grants the ability to create and manage pull requests and code reviews and to receive notifications about version control events via service hooks.
 - **vso.code_status**: Grants the ability to read and write commit and pull request status.

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author

nugetvss@microsoft.com
