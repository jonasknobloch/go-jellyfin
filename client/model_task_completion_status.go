/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// TaskCompletionStatus Enum TaskCompletionStatus.
type TaskCompletionStatus string

// List of TaskCompletionStatus
const (
	TASKCOMPLETIONSTATUS_COMPLETED TaskCompletionStatus = "Completed"
	TASKCOMPLETIONSTATUS_FAILED TaskCompletionStatus = "Failed"
	TASKCOMPLETIONSTATUS_CANCELLED TaskCompletionStatus = "Cancelled"
	TASKCOMPLETIONSTATUS_ABORTED TaskCompletionStatus = "Aborted"
)
