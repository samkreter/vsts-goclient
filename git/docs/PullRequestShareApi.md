# \PullRequestShareApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharePullRequest**](PullRequestShareApi.md#SharePullRequest) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/share | 


# **SharePullRequest**
> SharePullRequest(ctx, body, repositoryId, pullRequestId, project, apiVersion)


Sends an e-mail notification about a specific pull request to a set of recipients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ShareNotificationContext**](ShareNotificationContext.md)|  | 
  **repositoryId** | **string**| ID of the git repository. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
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

