# \StatusesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](StatusesApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/statuses | 
[**List**](StatusesApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/statuses | 


# **Create**
> GitStatus Create(ctx, body, commitId, repositoryId, project, apiVersion)


Create Git commit status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitStatus**](GitStatus.md)| Git commit status object to create. | 
  **commitId** | **string**| ID of the Git commit. | 
  **repositoryId** | **string**| ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitStatus**](GitStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitStatus List(ctx, commitId, repositoryId, project, apiVersion, optional)


Get statuses associated with the Git commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commitId** | **string**| ID of the Git commit. | 
  **repositoryId** | **string**| ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **top** | **optional.Int32**| Optional. The number of statuses to retrieve. Default is 1000. | 
 **skip** | **optional.Int32**| Optional. The number of statuses to ignore. Default is 0. For example, to retrieve results 101-150, set top to 50 and skip to 100. | 
 **latestOnly** | **optional.Bool**| The flag indicates whether to get only latest statuses grouped by &#x60;Context.Name&#x60; and &#x60;Context.Genre&#x60;. | 

### Return type

[**[]GitStatus**](GitStatus.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

