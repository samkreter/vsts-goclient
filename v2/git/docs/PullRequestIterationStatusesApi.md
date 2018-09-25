# \PullRequestIterationStatusesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PullRequestIterationStatusesApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 
[**Delete**](PullRequestIterationStatusesApi.md#Delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses/{statusId} | 
[**Get**](PullRequestIterationStatusesApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses/{statusId} | 
[**List**](PullRequestIterationStatusesApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 
[**Update**](PullRequestIterationStatusesApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/iterations/{iterationId}/statuses | 


# **Create**
> GitPullRequestStatus Create(ctx, body, repositoryId, pullRequestId, iterationId, project, apiVersion)


Create a pull request status on the iteration. This operation will have the same result as Create status on pull request with specified iteration ID in the request body.  The only required field for the status is `Context.Name` that uniquely identifies the status. Note that `iterationId` in the request body is optional since `iterationId` can be specified in the URL. A conflict between `iterationId` in the URL and `iterationId` in the request body will result in status code 400.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPullRequestStatus**](GitPullRequestStatus.md)| Pull request status to create. | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequestStatus**](GitPullRequestStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, repositoryId, pullRequestId, iterationId, statusId, project, apiVersion)


Delete pull request iteration status.  You can remove multiple statuses in one call by using Update operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. | 
  **statusId** | **int32**| ID of the pull request status. | 
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
> GitPullRequestStatus Get(ctx, repositoryId, pullRequestId, iterationId, statusId, project, apiVersion)


Get the specific pull request iteration status by ID. The status ID is unique within the pull request across all iterations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. | 
  **statusId** | **int32**| ID of the pull request status. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequestStatus**](GitPullRequestStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitPullRequestStatus List(ctx, repositoryId, pullRequestId, iterationId, project, apiVersion)


Get all the statuses associated with a pull request iteration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]GitPullRequestStatus**](GitPullRequestStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, body, repositoryId, pullRequestId, iterationId, project, apiVersion)


Update pull request iteration statuses collection. The only supported operation type is `remove`.  This operation allows to delete multiple statuses in one call. The path of the `remove` operation should refer to the ID of the pull request status. For example `path=\"/1\"` refers to the pull request status with ID 1.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonPatchDocument**](JsonPatchDocument.md)| Operations to apply to the pull request statuses in JSON Patch format. | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **iterationId** | **int32**| ID of the pull request iteration. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

