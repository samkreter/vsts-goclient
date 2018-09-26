# Comment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | Links to other related objects. | [optional] [default to null]
**Author** | [***IdentityRef**](IdentityRef.md) | The author of the comment. | [optional] [default to null]
**CommentType** | **string** | The comment type at the time of creation. | [optional] [default to null]
**Content** | **string** | The comment content. | [optional] [default to null]
**Id** | **int32** | The comment ID. IDs start at 1 and are unique to a pull request. | [optional] [default to null]
**IsDeleted** | **bool** | Whether or not this comment was soft-deleted. | [optional] [default to null]
**LastContentUpdatedDate** | [**time.Time**](time.Time.md) | The date the comment&#39;s content was last updated. | [optional] [default to null]
**LastUpdatedDate** | [**time.Time**](time.Time.md) | The date the comment was last updated. | [optional] [default to null]
**ParentCommentId** | **int32** | The ID of the parent comment. This is used for replies. | [optional] [default to null]
**PublishedDate** | [**time.Time**](time.Time.md) | The date the comment was first published. | [optional] [default to null]
**UsersLiked** | [**[]IdentityRef**](IdentityRef.md) | A list of the users who have liked this comment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


