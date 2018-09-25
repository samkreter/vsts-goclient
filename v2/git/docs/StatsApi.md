# \StatsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](StatsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/stats/branches | 


# **List**
> []GitBranchStats List(ctx, repositoryId, project, apiVersion, optional)


Retrieve statistics about all branches within a repository.

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



 **baseVersionDescriptorVersionOptions** | **optional.String**| Version options - Specify additional modifiers to version (e.g Previous) | 
 **baseVersionDescriptorVersion** | **optional.String**| Version string identifier (name of tag/branch, SHA1 of commit) | 
 **baseVersionDescriptorVersionType** | **optional.String**| Version type (branch, tag, or commit). Determines how Id is interpreted | 

### Return type

[**[]GitBranchStats**](GitBranchStats.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

