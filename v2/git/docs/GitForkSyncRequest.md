# GitForkSyncRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | Collection of related links | [optional] [default to null]
**DetailedStatus** | [***GitForkOperationStatusDetail**](GitForkOperationStatusDetail.md) |  | [optional] [default to null]
**OperationId** | **int32** | Unique identifier for the operation. | [optional] [default to null]
**Source** | [***GlobalGitRepositoryKey**](GlobalGitRepositoryKey.md) | Fully-qualified identifier for the source repository. | [optional] [default to null]
**SourceToTargetRefs** | [**[]SourceToTargetRef**](SourceToTargetRef.md) | If supplied, the set of ref mappings to use when performing a \&quot;sync\&quot; or create. If missing, all refs will be synchronized. | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


