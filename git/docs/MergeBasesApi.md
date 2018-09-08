# \MergeBasesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](MergeBasesApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryNameOrId}/commits/{commitId}/mergebases | 


# **List**
> []GitCommitRef List(ctx, repositoryNameOrId, commitId, otherCommitId, project, apiVersion, optional)


Find the merge bases of two commits, optionally across forks. If otherRepositoryId is not specified, the merge bases will only be calculated within the context of the local repositoryNameOrId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryNameOrId** | **string**| ID or name of the local repository. | 
  **commitId** | **string**| First commit, usually the tip of the target branch of the potential merge. | 
  **otherCommitId** | **string**| Other commit, usually the tip of the source branch of the potential merge. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **otherCollectionId** | [**optional.Interface of string**](.md)| The collection ID where otherCommitId lives. | 
 **otherRepositoryId** | [**optional.Interface of string**](.md)| The repository ID where otherCommitId lives. | 

### Return type

[**[]GitCommitRef**](GitCommitRef.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

