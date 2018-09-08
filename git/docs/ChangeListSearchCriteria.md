# ChangeListSearchCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareVersion** | **string** | If provided, a version descriptor to compare against base | [optional] [default to null]
**ExcludeDeletes** | **bool** | If true, don&#39;t include delete history entries | [optional] [default to null]
**FollowRenames** | **bool** | Whether or not to follow renames for the given item being queried | [optional] [default to null]
**FromDate** | **string** | If provided, only include history entries created after this date (string) | [optional] [default to null]
**FromVersion** | **string** | If provided, a version descriptor for the earliest change list to include | [optional] [default to null]
**ItemPath** | **string** | Path of item to search under. If the itemPaths memebr is used then it will take precedence over this. | [optional] [default to null]
**ItemPaths** | **[]string** | List of item paths to search under. If this member is used then itemPath will be ignored. | [optional] [default to null]
**ItemVersion** | **string** | Version of the items to search | [optional] [default to null]
**Skip** | **int32** | Number of results to skip (used when clicking more...) | [optional] [default to null]
**ToDate** | **string** | If provided, only include history entries created before this date (string) | [optional] [default to null]
**Top** | **int32** | If provided, the maximum number of history entries to return | [optional] [default to null]
**ToVersion** | **string** | If provided, a version descriptor for the latest change list to include | [optional] [default to null]
**User** | **string** | Alias or display name of user who made the changes | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


