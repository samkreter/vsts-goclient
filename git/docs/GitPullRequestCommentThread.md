# GitPullRequestCommentThread

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | Links to other related objects. | [optional] [default to null]
**Comments** | [**[]Comment**](Comment.md) | A list of the comments. | [optional] [default to null]
**Id** | **int32** | The comment thread id. | [optional] [default to null]
**Identities** | [**map[string]IdentityRef**](IdentityRef.md) | Set of identities related to this thread | [optional] [default to null]
**IsDeleted** | **bool** | Specify if the thread is deleted which happens when all comments are deleted. | [optional] [default to null]
**LastUpdatedDate** | [**time.Time**](time.Time.md) | The time this thread was last updated. | [optional] [default to null]
**Properties** | [***PropertiesCollection**](PropertiesCollection.md) | Optional properties associated with the thread as a collection of key-value pairs. | [optional] [default to null]
**PublishedDate** | [**time.Time**](time.Time.md) | The time this thread was published. | [optional] [default to null]
**Status** | **string** | The status of the comment thread. | [optional] [default to null]
**ThreadContext** | [***CommentThreadContext**](CommentThreadContext.md) | Specify thread context such as position in left/right file. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


