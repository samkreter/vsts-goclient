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
type GitBlobRef struct {
	Links *ReferenceLinks `json:"_links,omitempty"`
	// SHA1 hash of git object
	ObjectId string `json:"objectId,omitempty"`
	// Size of blob content (in bytes)
	Size int64 `json:"size,omitempty"`
	Url string `json:"url,omitempty"`
}