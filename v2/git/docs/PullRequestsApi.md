# \PullRequestsApi

All URIs are relative to *https://fabrikam-fiber-inc.visualstudio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PullRequestsApi.md#Create) | **Post** /{project}/_apis/git/repositories/{repositoryId}/pullrequests | 
[**GetPullRequest**](PullRequestsApi.md#GetPullRequest) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId} | 
[**GetPullRequestById**](PullRequestsApi.md#GetPullRequestById) | **Get** /{project}/_apis/git/pullrequests/{pullRequestId} | 
[**GetPullRequests**](PullRequestsApi.md#GetPullRequests) | **Get** /{project}/_apis/git/repositories/{repositoryId}/pullrequests | 
[**GetPullRequestsByProject**](PullRequestsApi.md#GetPullRequestsByProject) | **Get** /{project}/_apis/git/pullrequests | 
[**Update**](PullRequestsApi.md#Update) | **Patch** /{project}/_apis/git/repositories/{repositoryId}/pullrequests/{pullRequestId} | 


# **Create**
> GitPullRequest Create(ctx, body, repositoryId, project, apiVersion, optional)


Create a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPullRequest**](GitPullRequest.md)| The pull request to create. | 
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **supportsIterations** | **optional.Bool**| If true, subsequent pushes to the pull request will be individually reviewable. Set this to false for large pull requests for performance reasons if this functionality is not needed. | 

### Return type

[**GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequest**
> GitPullRequest GetPullRequest(ctx, repositoryId, pullRequestId, project, apiVersion, optional)


Retrieve a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| The ID of the pull request to retrieve. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetPullRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPullRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **maxCommentLength** | **optional.Int32**| Not used. | 
 **skip** | **optional.Int32**| Not used. | 
 **top** | **optional.Int32**| Not used. | 
 **includeCommits** | **optional.Bool**| If true, the pull request will be returned with the associated commits. | 
 **includeWorkItemRefs** | **optional.Bool**| If true, the pull request will be returned with the associated work item references. | 

### Return type

[**GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestById**
> GitPullRequest GetPullRequestById(ctx, pullRequestId, project, apiVersion)


Retrieve a pull request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pullRequestId** | **int32**| The ID of the pull request to retrieve. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequests**
> []GitPullRequest GetPullRequests(ctx, repositoryId, project, apiVersion, optional)


Retrieve all pull requests matching a specified criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetPullRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPullRequestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **searchCriteriaIncludeLinks** | **optional.Bool**| Whether to include the _links field on the shallow references | 
 **searchCriteriaSourceRefName** | **optional.String**| If set, search for pull requests from this branch. | 
 **searchCriteriaSourceRepositoryId** | [**optional.Interface of string**](.md)| If set, search for pull requests whose source branch is in this repository. | 
 **searchCriteriaTargetRefName** | **optional.String**| If set, search for pull requests into this branch. | 
 **searchCriteriaStatus** | **optional.String**| If set, search for pull requests that are in this state. Defaults to Active if unset. | 
 **searchCriteriaReviewerId** | [**optional.Interface of string**](.md)| If set, search for pull requests that have this identity as a reviewer. | 
 **searchCriteriaCreatorId** | [**optional.Interface of string**](.md)| If set, search for pull requests that were created by this identity. | 
 **searchCriteriaRepositoryId** | [**optional.Interface of string**](.md)| If set, search for pull requests whose target branch is in this repository. | 
 **maxCommentLength** | **optional.Int32**| Not used. | 
 **skip** | **optional.Int32**| The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100. | 
 **top** | **optional.Int32**| The number of pull requests to retrieve. | 

### Return type

[**[]GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPullRequestsByProject**
> []GitPullRequest GetPullRequestsByProject(ctx, project, apiVersion, optional)


Retrieve all pull requests matching a specified criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 
 **optional** | ***GetPullRequestsByProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPullRequestsByProjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **searchCriteriaIncludeLinks** | **optional.Bool**| Whether to include the _links field on the shallow references | 
 **searchCriteriaSourceRefName** | **optional.String**| If set, search for pull requests from this branch. | 
 **searchCriteriaSourceRepositoryId** | [**optional.Interface of string**](.md)| If set, search for pull requests whose source branch is in this repository. | 
 **searchCriteriaTargetRefName** | **optional.String**| If set, search for pull requests into this branch. | 
 **searchCriteriaStatus** | **optional.String**| If set, search for pull requests that are in this state. Defaults to Active if unset. | 
 **searchCriteriaReviewerId** | [**optional.Interface of string**](.md)| If set, search for pull requests that have this identity as a reviewer. | 
 **searchCriteriaCreatorId** | [**optional.Interface of string**](.md)| If set, search for pull requests that were created by this identity. | 
 **searchCriteriaRepositoryId** | [**optional.Interface of string**](.md)| If set, search for pull requests whose target branch is in this repository. | 
 **maxCommentLength** | **optional.Int32**| Not used. | 
 **skip** | **optional.Int32**| The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100. | 
 **top** | **optional.Int32**| The number of pull requests to retrieve. | 

### Return type

[**[]GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> GitPullRequest Update(ctx, body, repositoryId, pullRequestId, project, apiVersion)


Update a pull request.  These are the properties that can be updated with the API:  - Status  - Title  - Description  - CompletionOptions  - MergeOptions  - AutoCompleteSetBy.Id  - TargetRefName (when the PR retargeting feature is enabled)  Attempting to update other properties outside of this list will either cause the server to throw an `InvalidArgumentValueException`,  or to silently ignore the update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GitPullRequest**](GitPullRequest.md)| The pull request content to update. | 
  **repositoryId** | **string**| The repository ID of the pull request&#39;s target branch. | 
  **pullRequestId** | **int32**| The ID of the pull request to retrieve. | 
  **project** | **string**| Project ID or project name | 
  **apiVersion** | **string**| Version of the API to use.  This should be set to &#39;5.0-preview.1&#39; to use this version of the api. | 

### Return type

[**GitPullRequest**](GitPullRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

