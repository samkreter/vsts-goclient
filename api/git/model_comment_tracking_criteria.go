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

// Comment tracking criteria is used to identify which iteration context the thread has been tracked to (if any) along with some detail about the original position and filename.
type CommentTrackingCriteria struct {
	// The iteration of the file on the left side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0.
	FirstComparingIteration int32 `json:"firstComparingIteration,omitempty"`
	// Original filepath the thread was created on before tracking. This will be different than the current thread filepath if the file in question was renamed in a later iteration.
	OrigFilePath string `json:"origFilePath,omitempty"`
	// Original position of last character of the thread's span in left file.
	OrigLeftFileEnd *CommentPosition `json:"origLeftFileEnd,omitempty"`
	// Original position of first character of the thread's span in left file.
	OrigLeftFileStart *CommentPosition `json:"origLeftFileStart,omitempty"`
	// Original position of last character of the thread's span in right file.
	OrigRightFileEnd *CommentPosition `json:"origRightFileEnd,omitempty"`
	// Original position of first character of the thread's span in right file.
	OrigRightFileStart *CommentPosition `json:"origRightFileStart,omitempty"`
	// The iteration of the file on the right side of the diff that the thread will be tracked to. Threads were tracked if this is greater than 0.
	SecondComparingIteration int32 `json:"secondComparingIteration,omitempty"`
}
