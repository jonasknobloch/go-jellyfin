/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// NotificationTypeInfo struct for NotificationTypeInfo
type NotificationTypeInfo struct {
	Type *string `json:"Type,omitempty"`
	Name *string `json:"Name,omitempty"`
	Enabled bool `json:"Enabled,omitempty"`
	Category *string `json:"Category,omitempty"`
	IsBasedOnUserEvent bool `json:"IsBasedOnUserEvent,omitempty"`
}