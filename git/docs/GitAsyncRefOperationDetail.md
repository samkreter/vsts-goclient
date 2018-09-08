# GitAsyncRefOperationDetail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conflict** | **bool** | Indicates if there was a conflict generated when trying to cherry pick or revert the changes. | [optional] [default to null]
**CurrentCommitId** | **string** | The current commit from the list of commits that are being cherry picked or reverted. | [optional] [default to null]
**FailureMessage** | **string** | Detailed information about why the cherry pick or revert failed to complete. | [optional] [default to null]
**Progress** | **float64** | A number between 0 and 1 indicating the percent complete of the operation. | [optional] [default to null]
**Status** | **string** | Provides a status code that indicates the reason the cherry pick or revert failed. | [optional] [default to null]
**Timedout** | **bool** | Indicates if the operation went beyond the maximum time allowed for a cherry pick or revert operation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


