# Attachment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***ReferenceLinks**](ReferenceLinks.md) | Links to other related objects. | [optional] [default to null]
**Author** | [***IdentityRef**](IdentityRef.md) | The person that uploaded this attachment. | [optional] [default to null]
**ContentHash** | **string** | Content hash of on-disk representation of file content. Its calculated by the server by using SHA1 hash function. | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | The time the attachment was uploaded. | [optional] [default to null]
**Description** | **string** | The description of the attachment. | [optional] [default to null]
**DisplayName** | **string** | The display name of the attachment. Can&#39;t be null or empty. | [optional] [default to null]
**Id** | **int32** | Id of the attachment. | [optional] [default to null]
**Properties** | [***PropertiesCollection**](PropertiesCollection.md) | Extended properties. | [optional] [default to null]
**Url** | **string** | The url to download the content of the attachment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


