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
type GitResolutionMergeContent struct {
	MergeType string `json:"mergeType,omitempty"`
	UserMergedBlob *GitBlobRef `json:"userMergedBlob,omitempty"`
	UserMergedContent []string `json:"userMergedContent,omitempty"`
}
