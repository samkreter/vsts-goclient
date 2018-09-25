# \RepositoriesApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RepositoriesApi.md#Create) | **Post** /{project}/_apis/git/repositories | 
[**Delete**](RepositoriesApi.md#Delete) | **Delete** /{project}/_apis/git/repositories/{repositoryId} | 
[**DeleteRepositoryFromRecycleBin**](RepositoriesApi.md#DeleteRepositoryFromRecycleBin) | **Delete** /{project}/_apis/git/recycleBin/repositories/{repositoryId} | 
[**GetDeletedRepositories**](RepositoriesApi.md#GetDeletedRepositories) | **Get** /{project}/_apis/git/deletedrepositories | 
[**GetRecycleBinRepositories**](RepositoriesApi.md#GetRecycleBinRepositories) | **Get** /{project}/_apis/git/recycleBin/repositories | 
[**GetRepository**](RepositoriesApi.md#GetRepository) | **Get** /{project}/_apis/git/repositories/{repositoryId} | 
[**List**](RepositoriesApi.md#List) | **Get** /{project}/_apis/git/repositories | 
[**RestoreRepositoryFromRecycleBin**](RepositoriesApi.md#RestoreRepositoryFromRecycleBin) | **Patch** /{project}/_apis/git/recycleBin/repositories/{repositoryId} | 
[**Update**](RepositoriesApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId} | 


# **Create**
> GitRepository Create(ctx, body, project, apiVersion, optional)


Create a git repository in a team project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitRepositoryCreateOptions**](GitRepositoryCreateOptions.md)| Specify the repo name, team project and/or parent repository. Team project information can be ommitted from gitRepositoryToCreate if the request is project-scoped (i.e., includes project Id). | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sourceRef** | **optional.String**| [optional] Specify the source refs to use while creating a fork repo | 

### Return type

[**GitRepository**](GitRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, repositoryId, project, apiVersion)


Delete a git repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | [**string**](.md)| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepositoryFromRecycleBin**
> DeleteRepositoryFromRecycleBin(ctx, project, repositoryId, apiVersion)


Destroy (hard delete) a soft-deleted Git repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **repositoryId** | [**string**](.md)| The ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeletedRepositories**
> []GitDeletedRepository GetDeletedRepositories(ctx, project, apiVersion)


Retrieve deleted git repositories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]GitDeletedRepository**](GitDeletedRepository.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecycleBinRepositories**
> []GitDeletedRepository GetRecycleBinRepositories(ctx, project, apiVersion)


Retrieve soft-deleted git repositories from the recycle bin.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**[]GitDeletedRepository**](GitDeletedRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository**
> GitRepository GetRepository(ctx, repositoryId, project, apiVersion)


Retrieve a git repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRepository**](GitRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []GitRepository List(ctx, project, apiVersion, optional)


Retrieve git repositories.

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


 **includeLinks** | **optional.Bool**| [optional] True to include reference links. The default value is false. | 
 **includeAllUrls** | **optional.Bool**| [optional] True to include all remote URLs. The default value is false. | 
 **includeHidden** | **optional.Bool**| [optional] True to include hidden repositories. The default value is false. | 

### Return type

[**[]GitRepository**](GitRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreRepositoryFromRecycleBin**
> GitRepository RestoreRepositoryFromRecycleBin(ctx, body, project, repositoryId, apiVersion)


Recover a soft-deleted Git repository. Recently deleted repositories go into a soft-delete state for a period of time before they are hard deleted and become unrecoverable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitRecycleBinRepositoryDetails**](GitRecycleBinRepositoryDetails.md)|  | 
  **project** | **string**| Project ID or project name | 
  **repositoryId** | [**string**](.md)| The ID of the repository. | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRepository**](GitRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> GitRepository Update(ctx, body, repositoryId, project, apiVersion)


Updates the Git repository with either a new repo name or a new default branch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitRepository**](GitRepository.md)| Specify a new repo name or a new default branch of the repository | 
  **repositoryId** | [**string**](.md)| The name or ID of the repository. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitRepository**](GitRepository.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

