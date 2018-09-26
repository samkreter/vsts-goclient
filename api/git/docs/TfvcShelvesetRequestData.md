# TfvcShelvesetRequestData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeDetails** | **bool** | Whether to include policyOverride and notes Only applies when requesting a single deep shelveset | [optional] [default to null]
**IncludeLinks** | **bool** | Whether to include the _links field on the shallow references. Does not apply when requesting a single deep shelveset object. Links will always be included in the deep shelveset. | [optional] [default to null]
**IncludeWorkItems** | **bool** | Whether to include workItems | [optional] [default to null]
**MaxChangeCount** | **int32** | Max number of changes to include | [optional] [default to null]
**MaxCommentLength** | **int32** | Max length of comment | [optional] [default to null]
**Name** | **string** | Shelveset&#39;s name | [optional] [default to null]
**Owner** | **string** | Owner&#39;s ID. Could be a name or a guid. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


