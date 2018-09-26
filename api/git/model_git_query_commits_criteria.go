/*
 * Git
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.0-preview
 * Contact: nugetvss@microsoft.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 
type GitQueryCommitsCriteria struct {
	// Number of entries to skip
	Skip int32 `json:"$skip,omitempty"`
	// Maximum number of entries to retrieve
	Top int32 `json:"$top,omitempty"`
	// Alias or display name of the author
	Author string `json:"author,omitempty"`
	// If provided, the earliest commit in the graph to search
	CompareVersion *GitVersionDescriptor `json:"compareVersion,omitempty"`
	// Only applies when an itemPath is specified. This determines whether to exclude delete entries of the specified path.
	ExcludeDeletes bool `json:"excludeDeletes,omitempty"`
	// If provided, a lower bound for filtering commits alphabetically
	FromCommitId string `json:"fromCommitId,omitempty"`
	// If provided, only include history entries created after this date (string)
	FromDate string `json:"fromDate,omitempty"`
	// What Git history mode should be used. This only applies to the search criteria when Ids = null and an itemPath is specified.
	HistoryMode string `json:"historyMode,omitempty"`
	// If provided, specifies the exact commit ids of the commits to fetch. May not be combined with other parameters.
	Ids []string `json:"ids,omitempty"`
	// Whether to include the _links field on the shallow references
	IncludeLinks bool `json:"includeLinks,omitempty"`
	// Whether to include the push information
	IncludePushData bool `json:"includePushData,omitempty"`
	// Whether to include the image Url for committers and authors
	IncludeUserImageUrl bool `json:"includeUserImageUrl,omitempty"`
	// Whether to include linked work items
	IncludeWorkItems bool `json:"includeWorkItems,omitempty"`
	// Path of item to search under
	ItemPath string `json:"itemPath,omitempty"`
	// If provided, identifies the commit or branch to search
	ItemVersion *GitVersionDescriptor `json:"itemVersion,omitempty"`
	// If provided, an upper bound for filtering commits alphabetically
	ToCommitId string `json:"toCommitId,omitempty"`
	// If provided, only include history entries created before this date (string)
	ToDate string `json:"toDate,omitempty"`
	// Alias or display name of the committer
	User string `json:"user,omitempty"`
}