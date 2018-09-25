# GitPullRequestQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]GitPullRequestQueryInput**](GitPullRequestQueryInput.md) | The queries to perform. | [optional] [default to null]
**Results** | [**[]map[string][]GitPullRequest**](map.md) | The results of the queries. This matches the QueryInputs list so Results[n] are the results of QueryInputs[n]. Each entry in the list is a dictionary of commit-&gt;pull requests. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


