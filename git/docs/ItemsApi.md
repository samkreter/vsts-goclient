# \ItemsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsBatch**](ItemsApi.md#GetItemsBatch) | **Post** /{project}/_apis/git/repositories/{repositoryId}/itemsbatch | 
[**List**](ItemsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/items | 


# **GetItemsBatch**
> [][]ErrorUnknown GetItemsBatch(ctx, body, repositoryId, project, apiVersion)


Post for retrieving a creating a batch out of a set of items in a repo / project given a list of paths or a long path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitItemRequestData**](GitItemRequestData.md)| Request data attributes: ItemDescriptors, IncludeContentMetadata, LatestProcessedChange, IncludeLinks. ItemDescriptors: Collection of items to fetch, including path, version, and recursion level. IncludeContentMetadata: Whether to include metadata for all items LatestProcessedChange: Whether to include shallow ref to commit that last changed each item. IncludeLinks: Whether to include the _links field on the shallow references. | 
  **repositoryId** | **string**| The name or ID of the repository | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[][]ErrorUnknown**](array.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitItem List(ctx, repositoryId, project, apiVersion, optional)


Get Item Metadata and/or Content for a collection of items. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.

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



 **scopePath** | **optional.String**| The path scope.  The default is null. | 
 **recursionLevel** | **optional.String**| The recursion level of this request. The default is &#39;none&#39;, no recursion. | 
 **includeContentMetadata** | **optional.Bool**| Set to true to include content metadata.  Default is false. | 
 **latestProcessedChange** | **optional.Bool**| Set to true to include the lastest changes.  Default is false. | 
 **download** | **optional.Bool**| Set to true to download the response as a file.  Default is false. | 
 **includeLinks** | **optional.Bool**| Set to true to include links to items.  Default is false. | 
 **format** | **optional.String**| If specified, this overrides the HTTP Accept request header to return either &#39;json&#39; or &#39;zip&#39;. If $format is specified, then api-version should also be specified as a query parameter. | 
 **versionDescriptorVersionOptions** | **optional.String**| Version options - Specify additional modifiers to version (e.g Previous) | 
 **versionDescriptorVersion** | **optional.String**| Version string identifier (name of tag/branch, SHA1 of commit) | 
 **versionDescriptorVersionType** | **optional.String**| Version type (branch, tag, or commit). Determines how Id is interpreted | 

### Return type

[**[]GitItem**](GitItem.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

