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

import (
	"time"
)

// 
type TfvcLabel struct {
	Links *ReferenceLinks `json:"_links,omitempty"`
	Description string `json:"description,omitempty"`
	Id int32 `json:"id,omitempty"`
	LabelScope string `json:"labelScope,omitempty"`
	ModifiedDate time.Time `json:"modifiedDate,omitempty"`
	Name string `json:"name,omitempty"`
	Owner *IdentityRef `json:"owner,omitempty"`
	Url string `json:"url,omitempty"`
}
