# \PullRequestReviewersApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePullRequestReviewer**](PullRequestReviewersApi.md#CreatePullRequestReviewer) | **Put** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
[**CreatePullRequestReviewers**](PullRequestReviewersApi.md#CreatePullRequestReviewers) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 
[**Delete**](PullRequestReviewersApi.md#Delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
[**Get**](PullRequestReviewersApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers/{reviewerId} | 
[**List**](PullRequestReviewersApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 
[**Update**](PullRequestReviewersApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/reviewers | 


# **CreatePullRequestReviewer**
> IdentityRefWithVote CreatePullRequestReviewer(ctx, body, repositoryId, pullRequestId, reviewerId, project, apiVersion)


Add a reviewer to a pull request or cast a vote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdentityRefWithVote**](IdentityRefWithVote.md)| Reviewer&#39;s vote.&lt;br /&gt;If the reviewer&#39;s ID is included here, it must match the reviewerID parameter.&lt;br /&gt;Reviewers can set their own vote with this method.  When adding other reviewers, vote must be set to zero. | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **reviewerId** | **string**| ID of the reviewer. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**IdentityRefWithVote**](IdentityRefWithVote.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePullRequestReviewers**
> []IdentityRefWithVote CreatePullRequestReviewers(ctx, body, repositoryId, pullRequestId, project, apiVersion)


Add reviewers to a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]IdentityRef**](IdentityRef.md)| Reviewers to add to the pull request. | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]IdentityRefWithVote**](IdentityRefWithVote.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, repositoryId, pullRequestId, reviewerId, project, apiVersion)


Remove a reviewer from a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **reviewerId** | **string**| ID of the reviewer to remove. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> IdentityRefWithVote Get(ctx, repositoryId, pullRequestId, reviewerId, project, apiVersion)


Retrieve information about a particular reviewer on a pull request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **reviewerId** | **string**| ID of the reviewer. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**IdentityRefWithVote**](IdentityRefWithVote.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []IdentityRefWithVote List(ctx, repositoryId, pullRequestId, project, apiVersion)


Retrieve the reviewers for a pull request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]IdentityRefWithVote**](IdentityRefWithVote.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, body, repositoryId, pullRequestId, project, apiVersion)


Reset the votes of multiple reviewers on a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]IdentityRefWithVote**](IdentityRefWithVote.md)| IDs of the reviewers whose votes will be reset to zero | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

