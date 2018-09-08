# GitItemRequestData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeContentMetadata** | **bool** | Whether to include metadata for all items | [optional] [default to null]
**IncludeLinks** | **bool** | Whether to include the _links field on the shallow references | [optional] [default to null]
**ItemDescriptors** | [**[]GitItemDescriptor**](GitItemDescriptor.md) | Collection of items to fetch, including path, version, and recursion level | [optional] [default to null]
**LatestProcessedChange** | **bool** | Whether to include shallow ref to commit that last changed each item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


