# \ImportRequestsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ImportRequestsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/importRequests | 
[**Get**](ImportRequestsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/importRequests/{importRequestId} | 
[**Query**](ImportRequestsApi.md#Query) | **Get** /{project}/_apis/git/repositories/{repositoryId}/importRequests | 
[**Update**](ImportRequestsApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/importRequests/{importRequestId} | 


# **Create**
> GitImportRequest Create(ctx, body, project, repositoryId, apiVersion)


Create an import request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitImportRequest**](GitImportRequest.md)| The import request to create. | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitImportRequest**](GitImportRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> GitImportRequest Get(ctx, project, repositoryId, importRequestId, apiVersion)


Retrieve a particular import request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **importRequestId** | **int32**| The unique identifier for the import request. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitImportRequest**](GitImportRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Query**
> []GitImportRequest Query(ctx, project, repositoryId, apiVersion, optional)


Retrieve import requests for a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***QueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeAbandoned** | **optional.Bool**| True to include abandoned import requests in the results. | 

### Return type

[**[]GitImportRequest**](GitImportRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> GitImportRequest Update(ctx, body, project, repositoryId, importRequestId, apiVersion)


Retry or abandon a failed import request.  There can only be one active import request associated with a repository. Marking a failed import request abandoned makes it inactive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitImportRequest**](GitImportRequest.md)| The updated version of the import request. Currently, the only change allowed is setting the Status to Queued or Abandoned. | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **importRequestId** | **int32**| The unique identifier for the import request to update. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitImportRequest**](GitImportRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

