# GitPullRequestSearchCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorId** | **string** | If set, search for pull requests that were created by this identity. | [optional] [default to null]
**IncludeLinks** | **bool** | Whether to include the _links field on the shallow references | [optional] [default to null]
**RepositoryId** | **string** | If set, search for pull requests whose target branch is in this repository. | [optional] [default to null]
**ReviewerId** | **string** | If set, search for pull requests that have this identity as a reviewer. | [optional] [default to null]
**SourceRefName** | **string** | If set, search for pull requests from this branch. | [optional] [default to null]
**SourceRepositoryId** | **string** | If set, search for pull requests whose source branch is in this repository. | [optional] [default to null]
**Status** | **string** | If set, search for pull requests that are in this state. Defaults to Active if unset. | [optional] [default to null]
**TargetRefName** | **string** | If set, search for pull requests into this branch. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


