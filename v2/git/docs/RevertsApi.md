# \RevertsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RevertsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/reverts | 
[**GetRevert**](RevertsApi.md#GetRevert) | **Get** /{project}/_apis/git/repositories/{repositoryId}/reverts/{revertId} | 
[**GetRevertForRefName**](RevertsApi.md#GetRevertForRefName) | **Get** /{project}/_apis/git/repositories/{repositoryId}/reverts | 


# **Create**
> GitRevert Create(ctx, body, project, repositoryId, apiVersion)


Starts the operation to create a new branch which reverts changes introduced by either a specific commit or commits that are associated to a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitAsyncRefOperationParameters**](GitAsyncRefOperationParameters.md)|  | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRevert**](GitRevert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRevert**
> GitRevert GetRevert(ctx, project, revertId, repositoryId, apiVersion)


Retrieve information about a revert operation by revert Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **revertId** | **int32**| ID of the revert operation. | 
  **repositoryId** | **string**| ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRevert**](GitRevert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRevertForRefName**
> GitRevert GetRevertForRefName(ctx, project, repositoryId, refName, apiVersion)


Retrieve information about a revert operation for a specific branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID of the repository. | 
  **refName** | **string**| The GitAsyncRefOperationParameters generatedRefName used for the revert operation. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRevert**](GitRevert.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

