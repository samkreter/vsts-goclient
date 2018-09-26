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
type GitRefUpdateResult struct {
	// Custom message for the result object For instance, Reason for failing.
	CustomMessage string `json:"customMessage,omitempty"`
	// Whether the ref is locked or not
	IsLocked bool `json:"isLocked,omitempty"`
	// Ref name
	Name string `json:"name,omitempty"`
	// New object ID
	NewObjectId string `json:"newObjectId,omitempty"`
	// Old object ID
	OldObjectId string `json:"oldObjectId,omitempty"`
	// Name of the plugin that rejected the updated.
	RejectedBy string `json:"rejectedBy,omitempty"`
	// Repository ID
	RepositoryId string `json:"repositoryId,omitempty"`
	// True if the ref update succeeded, false otherwise
	Success bool `json:"success,omitempty"`
	// Status of the update from the TFS server.
	UpdateStatus string `json:"updateStatus,omitempty"`
}