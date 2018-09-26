# \ForksApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ForksApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests | 
[**Get**](ForksApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests/{forkSyncOperationId} | 
[**GetForkSyncRequests**](ForksApi.md#GetForkSyncRequests) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forkSyncRequests | 
[**List**](ForksApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/forks/{collectionId} | 


# **Create**
> GitForkSyncRequest Create(ctx, body, repositoryNameOrId, project, apiVersion, optional)


Request that another repository's refs be fetched into this one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitForkSyncRequestParameters**](GitForkSyncRequestParameters.md)| Source repository and ref mapping. | 
  **repositoryNameOrId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeLinks** | **optional.Bool**| True to include links | 

### Return type

[**GitForkSyncRequest**](GitForkSyncRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> GitForkSyncRequest Get(ctx, repositoryNameOrId, forkSyncOperationId, project, apiVersion, optional)


Get a specific fork sync operation's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryNameOrId** | **string**| The name or ID of the repository. | 
  **forkSyncOperationId** | **int32**| OperationId of the sync request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeLinks** | **optional.Bool**| True to include links. | 

### Return type

[**GitForkSyncRequest**](GitForkSyncRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForkSyncRequests**
> []GitForkSyncRequest GetForkSyncRequests(ctx, repositoryNameOrId, project, apiVersion, optional)


Retrieve all requested fork sync operations on this repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryNameOrId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetForkSyncRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetForkSyncRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeAbandoned** | **optional.Bool**| True to include abandoned requests. | 
 **includeLinks** | **optional.Bool**| True to include links. | 

### Return type

[**[]GitForkSyncRequest**](GitForkSyncRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitRepositoryRef List(ctx, repositoryNameOrId, collectionId, project, apiVersion, optional)


Retrieve all forks of a repository in the collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryNameOrId** | **string**| The name or ID of the repository. | 
  **collectionId** | [**string**](.md)| Team project collection ID. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeLinks** | **optional.Bool**| True to include links. | 

### Return type

[**[]GitRepositoryRef**](GitRepositoryRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

