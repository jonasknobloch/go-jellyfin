/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// IScheduledTask Interface IScheduledTaskWorker.
type IScheduledTask struct {
	// Gets the name of the task.
	Name *string `json:"Name,omitempty"`
	// Gets the key of the task.
	Key *string `json:"Key,omitempty"`
	// Gets the description.
	Description *string `json:"Description,omitempty"`
	// Gets the category.
	Category *string `json:"Category,omitempty"`
}
