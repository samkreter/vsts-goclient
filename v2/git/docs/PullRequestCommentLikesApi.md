# \PullRequestCommentLikesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PullRequestCommentLikesApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 
[**Delete**](PullRequestCommentLikesApi.md#Delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 
[**List**](PullRequestCommentLikesApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/threads/{threadId}/comments/{commentId}/likes | 


# **Create**
> Create(ctx, repositoryId, pullRequestId, threadId, commentId, project, apiVersion)


Add a like on a comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **threadId** | **int32**| The ID of the thread that contains the comment. | 
  **commentId** | **int32**| The ID of the comment. | 
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

# **Delete**
> Delete(ctx, repositoryId, pullRequestId, threadId, commentId, project, apiVersion)


Delete a like on a comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **threadId** | **int32**| The ID of the thread that contains the comment. | 
  **commentId** | **int32**| The ID of the comment. | 
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

# **List**
> []IdentityRef List(ctx, repositoryId, pullRequestId, threadId, commentId, project, apiVersion)


Get likes for a comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **threadId** | **int32**| The ID of the thread that contains the comment. | 
  **commentId** | **int32**| The ID of the comment. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]IdentityRef**](IdentityRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

