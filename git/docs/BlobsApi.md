# \BlobsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlob**](BlobsApi.md#GetBlob) | **Get** /{project}/_apis/git/repositories/{repositoryId}/blobs/{sha1} | 
[**GetBlobsZip**](BlobsApi.md#GetBlobsZip) | **Post** /{project}/_apis/git/repositories/{repositoryId}/blobs | 


# **GetBlob**
> GitBlobRef GetBlob(ctx, repositoryId, sha1, project, apiVersion, optional)


Get a single blob.  Repositories have both a name and an identifier. Identifiers are globally unique, but several projects may contain a repository of the same name. You don't need to include the project if you specify a repository by ID. However, if you specify a repository by name, you must also specify the project (by name or ID).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The name or ID of the repository. | 
  **sha1** | **string**| SHA1 hash of the file. You can get the SHA1 of a file using the \&quot;Git/Items/Get Item\&quot; endpoint. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetBlobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBlobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **download** | **optional.Bool**| If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip | 
 **fileName** | **optional.String**| Provide a fileName to use for a download. | 
 **format** | **optional.String**| Options: json, zip, text, octetstream. If not set, defaults to the MIME type set in the Accept header. | 
 **resolveLfs** | **optional.Bool**| If true, try to resolve a blob to its LFS contents, if it&#39;s an LFS pointer file. Only compatible with octet-stream Accept headers or $format types | 

### Return type

[**GitBlobRef**](GitBlobRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlobsZip**
> string GetBlobsZip(ctx, body, repositoryId, project, apiVersion, optional)


Gets one or more blobs in a zip file download.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | **[]string**| Blob IDs (SHA1 hashes) to be returned in the zip file. | 
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetBlobsZipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBlobsZipOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filename** | **optional.String**|  | 

### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

