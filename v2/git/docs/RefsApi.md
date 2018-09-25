# \RefsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](RefsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/refs | 
[**UpdateRef**](RefsApi.md#UpdateRef) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/refs | 
[**UpdateRefs**](RefsApi.md#UpdateRefs) | **Post** /{project}/_apis/git/repositories/{repositoryId}/refs | 


# **List**
> []GitRef List(ctx, repositoryId, project, apiVersion, optional)


Queries the provided repository for its refs and returns them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **optional.String**| [optional] A filter to apply to the refs. | 
 **includeLinks** | **optional.Bool**| [optional] Specifies if referenceLinks should be included in the result. default is false. | 
 **includeStatuses** | **optional.Bool**| [optional] Includes up to the first 1000 commit statuses for each ref. The default value is false. | 
 **includeMyBranches** | **optional.Bool**| [optional] Includes only branches that the user owns, the branches the user favorites, and the default branch. The default value is false. Cannot be combined with the filter parameter. | 
 **latestStatusesOnly** | **optional.Bool**| [optional] True to include only the tip commit status for each ref. This option requires &#x60;includeStatuses&#x60; to be true. The default value is false. | 
 **peelTags** | **optional.Bool**| [optional] Annotated tags will populate the PeeledObjectId property. default is false. | 

### Return type

[**[]GitRef**](GitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRef**
> GitRef UpdateRef(ctx, body, repositoryId, filter, project, apiVersion, optional)


Lock or Unlock a branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitRefUpdate**](GitRefUpdate.md)| The ref update action (lock/unlock) to perform | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **filter** | **string**| The name of the branch to lock/unlock | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***UpdateRefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateRefOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **projectId** | **optional.String**| ID or name of the team project. Optional if specifying an ID for repository. | 

### Return type

[**GitRef**](GitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRefs**
> []GitRefUpdateResult UpdateRefs(ctx, body, repositoryId, project, apiVersion, optional)


Creating, updating, or deleting refs(branches).  Updating a ref means making it point at a different commit than it used to. You must specify both the old and new commit to avoid race conditions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]GitRefUpdate**](GitRefUpdate.md)| List of ref updates to attempt to perform | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***UpdateRefsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateRefsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **projectId** | **optional.String**| ID or name of the team project. Optional if specifying an ID for repository. | 

### Return type

[**[]GitRefUpdateResult**](GitRefUpdateResult.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

