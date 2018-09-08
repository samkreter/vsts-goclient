# \PullRequestCommitsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullRequestCommits**](PullRequestCommitsApi.md#GetPullRequestCommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/commits | 
[**GetPullRequestIterationCommits**](PullRequestCommitsApi.md#GetPullRequestIterationCommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/commits | 


# **GetPullRequestCommits**
> []GitCommitRef GetPullRequestCommits(ctx, repositoryId, pullRequestId, project, apiVersion)


Get the commits for the specified pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| ID or name of the repository. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]GitCommitRef**](GitCommitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestIterationCommits**
> []GitCommitRef GetPullRequestIterationCommits(ctx, repositoryId, pullRequestId, iterationId, project, apiVersion)


Get the commits for the specified iteration of a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| ID or name of the repository. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the iteration from which to get the commits. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]GitCommitRef**](GitCommitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

