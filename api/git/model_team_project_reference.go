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

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
	// Project abbreviation.
	Abbreviation string `json:"abbreviation,omitempty"`
	// The project's description (if any).
	Description string `json:"description,omitempty"`
	// Project identifier.
	Id string `json:"id,omitempty"`
	// Project name.
	Name string `json:"name,omitempty"`
	// Project revision.
	Revision int64 `json:"revision,omitempty"`
	// Project state.
	State string `json:"state,omitempty"`
	// Url to the full version of the object.
	Url string `json:"url,omitempty"`
	// Project visibility.
	Visibility string `json:"visibility,omitempty"`
}
