# GitCommit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | A collection of related REST reference links. | [optional] [default to null]
**Author** | [***GitUserDate**](GitUserDate.md) | Author of the commit. | [optional] [default to null]
**ChangeCounts** | [***ChangeCountDictionary**](ChangeCountDictionary.md) | Counts of the types of changes (edits, deletes, etc.) included with the commit. | [optional] [default to null]
**Changes** | [**[]GitChange**](GitChange.md) | An enumeration of the changes included with the commit. | [optional] [default to null]
**Comment** | **string** | Comment or message of the commit. | [optional] [default to null]
**CommentTruncated** | **bool** | Indicates if the comment is truncated from the full Git commit comment message. | [optional] [default to null]
**CommitId** | **string** | ID (SHA-1) of the commit. | [optional] [default to null]
**Committer** | [***GitUserDate**](GitUserDate.md) | Committer of the commit. | [optional] [default to null]
**Parents** | **[]string** | An enumeration of the parent commit IDs for this commit. | [optional] [default to null]
**Push** | [***GitPushRef**](GitPushRef.md) | The push associated with this commit. | [optional] [default to null]
**RemoteUrl** | **string** | Remote URL path to the commit. | [optional] [default to null]
**Statuses** | [**[]GitStatus**](GitStatus.md) | A list of status metadata from services and extensions that may associate additional information to the commit. | [optional] [default to null]
**Url** | **string** | REST URL for this resource. | [optional] [default to null]
**WorkItems** | [**[]ResourceRef**](ResourceRef.md) | A list of workitems associated with this commit. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


