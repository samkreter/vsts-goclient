# \AnnotatedTagsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](AnnotatedTagsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/annotatedtags | 
[**Get**](AnnotatedTagsApi.md#Get) | **Get** /{project}/_apis/git/repositories/{repositoryId}/annotatedtags/{objectId} | 


# **Create**
> GitAnnotatedTag Create(ctx, body, project, repositoryId, apiVersion)


Create an annotated tag.  Repositories have both a name and an identifier. Identifiers are globally unique, but several projects may contain a repository of the same name. You don't need to include the project if you specify a repository by ID. However, if you specify a repository by name, you must also specify the project (by name or ID).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitAnnotatedTag**](GitAnnotatedTag.md)| Object containing details of tag to be created. | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID or name of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitAnnotatedTag**](GitAnnotatedTag.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> GitAnnotatedTag Get(ctx, project, repositoryId, objectId, apiVersion)


Get an annotated tag.  Repositories have both a name and an identifier. Identifiers are globally unique, but several projects may contain a repository of the same name. You don't need to include the project if you specify a repository by ID. However, if you specify a repository by name, you must also specify the project (by name or ID).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | **string**| ID or name of the repository. | 
  **objectId** | **string**| ObjectId (Sha1Id) of tag to get. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitAnnotatedTag**](GitAnnotatedTag.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

