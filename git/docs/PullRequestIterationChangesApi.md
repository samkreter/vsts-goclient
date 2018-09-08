# \PullRequestIterationChangesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](PullRequestIterationChangesApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/changes | 


# **Get**
> GitPullRequestIterationChanges Get(ctx, repositoryId, pullRequestId, iterationId, project, apiVersion, optional)


Retrieve the changes made in a pull request between two iterations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. &lt;br /&gt; Iteration IDs are zero-based with zero indicating the common commit between the source and target branches. Iteration one is the head of the source branch at the time the pull request is created and subsequent iterations are created when there are pushes to the source branch. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **top** | **optional.Int32**| Optional. The number of changes to retrieve.  The default value is 100 and the maximum value is 2000. | 
 **skip** | **optional.Int32**| Optional. The number of changes to ignore.  For example, to retrieve changes 101-150, set top 50 and skip to 100. | 
 **compareTo** | **optional.Int32**| ID of the pull request iteration to compare against.  The default value is zero which indicates the comparison is made against the common commit between the source and target branches | 

### Return type

[**GitPullRequestIterationChanges**](GitPullRequestIterationChanges.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

