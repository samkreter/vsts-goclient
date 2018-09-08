# \CherryPicksApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CherryPicksApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks | 
[**GetCherryPick**](CherryPicksApi.md#GetCherryPick) | **Get** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks/{cherryPickId} | 
[**GetCherryPickForRefName**](CherryPicksApi.md#GetCherryPickForRefName) | **Get** /{project}/_apis/git/repositories/{repositoryId}/cherryPicks | 


# **Create**
> GitCherryPick Create(ctx, body, project, repositoryId, apiVersion)


Cherry pick a specific commit or commits that are associated to a pull request into a new branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitAsyncRefOperationParameters**](GitAsyncRefOperationParameters.md)|  | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitCherryPick**](GitCherryPick.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCherryPick**
> GitCherryPick GetCherryPick(ctx, project, cherryPickId, repositoryId, apiVersion)


Retrieve information about a cherry pick by cherry pick Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **cherryPickId** | **int32**| ID of the cherry pick. | 
  **repositoryId** | **string**| ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitCherryPick**](GitCherryPick.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCherryPickForRefName**
> GitCherryPick GetCherryPickForRefName(ctx, project, repositoryId, refName, apiVersion)


Retrieve information about a cherry pick for a specific branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID of the repository. | 
  **refName** | **string**| The GitAsyncRefOperationParameters generatedRefName used for the cherry pick operation. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitCherryPick**](GitCherryPick.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

