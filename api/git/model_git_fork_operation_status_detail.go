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

// Status information about a requested fork operation.
type GitForkOperationStatusDetail struct {
	// All valid steps for the forking process
	AllSteps []string `json:"allSteps,omitempty"`
	// Index into AllSteps for the current step
	CurrentStep int32 `json:"currentStep,omitempty"`
	// Error message if the operation failed.
	ErrorMessage string `json:"errorMessage,omitempty"`
}
