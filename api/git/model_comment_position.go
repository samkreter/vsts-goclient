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
type CommentPosition struct {
	// The line number of a thread's position. Starts at 1.
	Line int32 `json:"line,omitempty"`
	// The character offset of a thread's position inside of a line. Starts at 0.
	Offset int32 `json:"offset,omitempty"`
}
