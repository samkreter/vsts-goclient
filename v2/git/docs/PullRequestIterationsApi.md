# \PullRequestIterationsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](PullRequestIterationsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId} | 
[**List**](PullRequestIterationsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations | 


# **Get**
> GitPullRequestIteration Get(ctx, repositoryId, pullRequestId, iterationId, project, apiVersion)


Get the specified iteration for a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| ID or name of the repository. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration to return. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequestIteration**](GitPullRequestIteration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitPullRequestIteration List(ctx, repositoryId, pullRequestId, project, apiVersion, optional)


Get the list of iterations for the specified pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| ID or name of the repository. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeCommits** | **optional.Bool**| If true, include the commits associated with each iteration in the response. | 

### Return type

[**[]GitPullRequestIteration**](GitPullRequestIteration.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

