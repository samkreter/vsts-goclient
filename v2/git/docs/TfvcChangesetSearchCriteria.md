# TfvcChangesetSearchCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** | Alias or display name of user who made the changes | [optional] [default to null]
**FollowRenames** | **bool** | Whether or not to follow renames for the given item being queried | [optional] [default to null]
**FromDate** | **string** | If provided, only include changesets created after this date (string) Think of a better name for this. | [optional] [default to null]
**FromId** | **int32** | If provided, only include changesets after this changesetID | [optional] [default to null]
**IncludeLinks** | **bool** | Whether to include the _links field on the shallow references | [optional] [default to null]
**ItemPath** | **string** | Path of item to search under | [optional] [default to null]
**Mappings** | [**[]TfvcMappingFilter**](TfvcMappingFilter.md) |  | [optional] [default to null]
**ToDate** | **string** | If provided, only include changesets created before this date (string) Think of a better name for this. | [optional] [default to null]
**ToId** | **int32** | If provided, a version descriptor for the latest change list to include | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


