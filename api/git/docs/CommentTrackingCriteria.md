# CommentTrackingCriteria

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstComparingIteration** | **int32** | The iteration of the file on the left side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0. | [optional] [default to null]
**OrigFilePath** | **string** | Original filepath the thread was created on before tracking. This will be different than the current thread filepath if the file in question was renamed in a later iteration. | [optional] [default to null]
**OrigLeftFileEnd** | [***CommentPosition**](CommentPosition.md) | Original position of last character of the thread&#39;s span in left file. | [optional] [default to null]
**OrigLeftFileStart** | [***CommentPosition**](CommentPosition.md) | Original position of first character of the thread&#39;s span in left file. | [optional] [default to null]
**OrigRightFileEnd** | [***CommentPosition**](CommentPosition.md) | Original position of last character of the thread&#39;s span in right file. | [optional] [default to null]
**OrigRightFileStart** | [***CommentPosition**](CommentPosition.md) | Original position of first character of the thread&#39;s span in right file. | [optional] [default to null]
**SecondComparingIteration** | **int32** | The iteration of the file on the right side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


