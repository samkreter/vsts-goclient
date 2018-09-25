# \PullRequestLabelsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PullRequestLabelsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels | 
[**Delete**](PullRequestLabelsApi.md#Delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels/{labelIdOrName} | 
[**Get**](PullRequestLabelsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels/{labelIdOrName} | 
[**List**](PullRequestLabelsApi.md#List) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullRequests/{pullRequestId}/labels | 


# **Create**
> WebApiTagDefinition Create(ctx, body, repositoryId, pullRequestId, project, apiVersion, optional)


Create a label for a specified pull request. The only required field is the name of the new label.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebApiCreateTagRequestData**](WebApiCreateTagRequestData.md)| Label to assign to the pull request. | 
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **projectId** | **optional.String**| Project ID or project name. | 

### Return type

[**WebApiTagDefinition**](WebApiTagDefinition.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, repositoryId, pullRequestId, labelIdOrName, project, apiVersion, optional)


Removes a label from the set of those assigned to the pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **labelIdOrName** | **string**| The name or ID of the label requested. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **projectId** | **optional.String**| Project ID or project name. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> WebApiTagDefinition Get(ctx, repositoryId, pullRequestId, labelIdOrName, project, apiVersion, optional)


Retrieves a single label that has been assigned to a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **labelIdOrName** | **string**| The name or ID of the label requested. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **projectId** | **optional.String**| Project ID or project name. | 

### Return type

[**WebApiTagDefinition**](WebApiTagDefinition.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []WebApiTagDefinition List(ctx, repositoryId, pullRequestId, project, apiVersion, optional)


Get all the labels assigned to a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request’s target branch. | 
  **pullRequestId** | **int32**| ID of the pull request. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **projectId** | **optional.String**| Project ID or project name. | 

### Return type

[**[]WebApiTagDefinition**](WebApiTagDefinition.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

