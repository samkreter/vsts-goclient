# GitQueryCommitsCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | **int32** | Number of entries to skip | [optional] [default to null]
**Top** | **int32** | Maximum number of entries to retrieve | [optional] [default to null]
**Author** | **string** | Alias or display name of the author | [optional] [default to null]
**CompareVersion** | [***GitVersionDescriptor**](GitVersionDescriptor.md) | If provided, the earliest commit in the graph to search | [optional] [default to null]
**ExcludeDeletes** | **bool** | Only applies when an itemPath is specified. This determines whether to exclude delete entries of the specified path. | [optional] [default to null]
**FromCommitId** | **string** | If provided, a lower bound for filtering commits alphabetically | [optional] [default to null]
**FromDate** | **string** | If provided, only include history entries created after this date (string) | [optional] [default to null]
**HistoryMode** | **string** | What Git history mode should be used. This only applies to the search criteria when Ids &#x3D; null and an itemPath is specified. | [optional] [default to null]
**Ids** | **[]string** | If provided, specifies the exact commit ids of the commits to fetch. May not be combined with other parameters. | [optional] [default to null]
**IncludeLinks** | **bool** | Whether to include the _links field on the shallow references | [optional] [default to null]
**IncludePushData** | **bool** | Whether to include the push information | [optional] [default to null]
**IncludeUserImageUrl** | **bool** | Whether to include the image Url for committers and authors | [optional] [default to null]
**IncludeWorkItems** | **bool** | Whether to include linked work items | [optional] [default to null]
**ItemPath** | **string** | Path of item to search under | [optional] [default to null]
**ItemVersion** | [***GitVersionDescriptor**](GitVersionDescriptor.md) | If provided, identifies the commit or branch to search | [optional] [default to null]
**ToCommitId** | **string** | If provided, an upper bound for filtering commits alphabetically | [optional] [default to null]
**ToDate** | **string** | If provided, only include history entries created before this date (string) | [optional] [default to null]
**User** | **string** | Alias or display name of the committer | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


