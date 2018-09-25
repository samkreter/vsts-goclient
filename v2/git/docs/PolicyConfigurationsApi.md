# \PolicyConfigurationsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](PolicyConfigurationsApi.md#List) | **Get** /{project}/_apis/git/policy/configurations | 


# **List**
> []PolicyConfiguration List(ctx, project, apiVersion, optional)


Retrieve a list of policy configurations by a given set of scope/filtering criteria.  Below is a short description of how all of the query parameters interact with each other: - repositoryId set, refName set: returns all policy configurations that *apply* to a particular branch in a repository - repositoryId set, refName unset: returns all policy configurations that *apply* to a particular repository - repositoryId unset, refName unset: returns all policy configurations that are *defined* at the project level - repositoryId unset, refName set: returns all project-level branch policies, plus the project level configurations For all of the examples above, when policyType is set, it'll restrict results to the given policy type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repositoryId** | [**optional.Interface of string**](.md)| The repository id. | 
 **refName** | **optional.String**| The fully-qualified Git ref name (e.g. refs/heads/master). | 
 **policyType** | [**optional.Interface of string**](.md)| The policy type filter. | 

### Return type

[**[]PolicyConfiguration**](PolicyConfiguration.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

