# \CommitsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](CommitsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId} | 
[**GetChanges**](CommitsApi.md#GetChanges) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits/{commitId}/changes | 
[**GetCommitsBatch**](CommitsApi.md#GetCommitsBatch) | **Post** /{project}/_apis/git/repositories/{repositoryId}/commitsbatch | 
[**GetPushCommits**](CommitsApi.md#GetPushCommits) | **Get** /{project}/_apis/git/repositories/{repositoryId}/commits | 


# **Get**
> GitCommit Get(ctx, commitId, repositoryId, project, apiVersion, optional)


Retrieve a particular commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commitId** | **string**| The id of the commit. | 
  **repositoryId** | **string**| The id or friendly name of the repository. To use the friendly name, projectId must also be specified. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **changeCount** | **optional.Int32**| The number of changes to include in the result. | 

### Return type

[**GitCommit**](GitCommit.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChanges**
> GitCommitChanges GetChanges(ctx, commitId, repositoryId, project, apiVersion, optional)


Retrieve changes for a particular commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commitId** | **string**| The id of the commit. | 
  **repositoryId** | **string**| The id or friendly name of the repository. To use the friendly name, projectId must also be specified. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChangesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **top** | **optional.Int32**| The maximum number of changes to return. | 
 **skip** | **optional.Int32**| The number of changes to skip. | 

### Return type

[**GitCommitChanges**](GitCommitChanges.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommitsBatch**
> []GitCommitRef GetCommitsBatch(ctx, body, repositoryId, project, apiVersion, optional)


Retrieve git commits for a project matching the search criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitQueryCommitsCriteria**](GitQueryCommitsCriteria.md)| Search options | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetCommitsBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCommitsBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **skip** | **optional.Int32**| Number of commits to skip. | 
 **top** | **optional.Int32**| Maximum number of commits to return. | 
 **includeStatuses** | **optional.Bool**| True to include additional commit status information. | 

### Return type

[**[]GitCommitRef**](GitCommitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPushCommits**
> []GitCommitRef GetPushCommits(ctx, repositoryId, pushId, project, apiVersion, optional)


Retrieve a list of commits associated with a particular push.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The id or friendly name of the repository. To use the friendly name, projectId must also be specified. | 
  **pushId** | **int32**| The id of the push. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetPushCommitsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPushCommitsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **top** | **optional.Int32**| The maximum number of commits to return (\&quot;get the top x commits\&quot;). | 
 **skip** | **optional.Int32**| The number of commits to skip. | 
 **includeLinks** | **optional.Bool**| Set to false to avoid including REST Url links for resources. Defaults to true. | 

### Return type

[**[]GitCommitRef**](GitCommitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

