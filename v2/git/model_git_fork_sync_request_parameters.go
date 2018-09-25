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

// Parameters for creating a fork request
type GitForkSyncRequestParameters struct {
	// Fully-qualified identifier for the source repository.
	Source *GlobalGitRepositoryKey `json:"source,omitempty"`
	// If supplied, the set of ref mappings to use when performing a \"sync\" or create. If missing, all refs will be synchronized.
	SourceToTargetRefs []SourceToTargetRef `json:"sourceToTargetRefs,omitempty"`
}