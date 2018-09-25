# \PullRequestThreadsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PullRequestThreadsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads | 
[**Get**](PullRequestThreadsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId} | 
[**List**](PullRequestThreadsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads | 
[**Update**](PullRequestThreadsApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId} | 


# **Create**
> GitPullRequestCommentThread Create(ctx, body, repositoryId, pullRequestId, project, apiVersion)


Create a thread in a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPullRequestCommentThread**](GitPullRequestCommentThread.md)| The thread to create. Thread must contain at least one comment. | 
  **repositoryId** | **string**| Repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequestCommentThread**](GitPullRequestCommentThread.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> GitPullRequestCommentThread Get(ctx, repositoryId, pullRequestId, threadId, project, apiVersion, optional)


Retrieve a thread in a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **threadId** | **int32**| ID of the thread. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **iteration** | **optional.Int32**| If specified, thread position will be tracked using this iteration as the right side of the diff. | 
 **baseIteration** | **optional.Int32**| If specified, thread position will be tracked using this iteration as the left side of the diff. | 

### Return type

[**GitPullRequestCommentThread**](GitPullRequestCommentThread.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitPullRequestCommentThread List(ctx, repositoryId, pullRequestId, project, apiVersion, optional)


Retrieve all threads in a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **iteration** | **optional.Int32**| If specified, thread positions will be tracked using this iteration as the right side of the diff. | 
 **baseIteration** | **optional.Int32**| If specified, thread positions will be tracked using this iteration as the left side of the diff. | 

### Return type

[**[]GitPullRequestCommentThread**](GitPullRequestCommentThread.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> GitPullRequestCommentThread Update(ctx, body, repositoryId, pullRequestId, threadId, project, apiVersion)


Update a thread in a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPullRequestCommentThread**](GitPullRequestCommentThread.md)| The thread content that should be updated. | 
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **threadId** | **int32**| ID of the thread to update. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequestCommentThread**](GitPullRequestCommentThread.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

