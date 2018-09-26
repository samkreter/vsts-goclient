# \DiffsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DiffsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/diffs/commits | 


# **Get**
> GitCommitDiffs Get(ctx, repositoryId, project, apiVersion, optional)


Find the closest common commit (the merge base) between base and target commits, and get the diff between either the base and target commits or common and target commits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **diffCommonCommit** | **optional.Bool**| If true, diff between common and target commits. If false, diff between base and target commits. | 
 **top** | **optional.Int32**| Maximum number of changes to return. Defaults to 100. | 
 **skip** | **optional.Int32**| Number of changes to skip | 
 **baseVersionOptions** | **optional.String**| Version options - Specify additional modifiers to version (e.g Previous) | 
 **baseVersion** | **optional.String**| Version string identifier (name of tag/branch, SHA1 of commit) | 
 **baseVersionType** | **optional.String**| Version type (branch, tag, or commit). Determines how Id is interpreted | 
 **targetVersionOptions** | **optional.String**| Version options - Specify additional modifiers to version (e.g Previous) | 
 **targetVersion** | **optional.String**| Version string identifier (name of tag/branch, SHA1 of commit) | 
 **targetVersionType** | **optional.String**| Version type (branch, tag, or commit). Determines how Id is interpreted | 

### Return type

[**GitCommitDiffs**](GitCommitDiffs.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

