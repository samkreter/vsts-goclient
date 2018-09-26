# \TreesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](TreesApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/trees/{sha1} | 


# **Get**
> GitTreeRef Get(ctx, repositoryId, sha1, project, apiVersion, optional)


The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.  Repositories have both a name and an identifier. Identifiers are globally unique, but several projects may contain a repository of the same name. You don't need to include the project if you specify a repository by ID. However, if you specify a repository by name, you must also specify the project (by name or ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| Repository Id. | 
  **sha1** | **string**| SHA1 hash of the tree object. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **projectId** | **optional.String**| Project Id. | 
 **recursive** | **optional.Bool**| Search recursively. Include trees underneath this tree. Default is false. | 
 **fileName** | **optional.String**| Name to use if a .zip file is returned. Default is the object ID. | 
 **format** | **optional.String**| Use \&quot;zip\&quot;. Defaults to the MIME type set in the Accept header. | 

### Return type

[**GitTreeRef**](GitTreeRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

