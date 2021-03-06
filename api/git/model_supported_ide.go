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

// Represents a Supported IDE entity.
type SupportedIde struct {
	// The download URL for the IDE.
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// The type of the IDE.
	IdeType string `json:"ideType,omitempty"`
	// The name of the IDE.
	Name string `json:"name,omitempty"`
	// The URL to open the protocol handler for the IDE.
	ProtocolHandlerUrl string `json:"protocolHandlerUrl,omitempty"`
	// A list of SupportedPlatforms.
	SupportedPlatforms []string `json:"supportedPlatforms,omitempty"`
}
