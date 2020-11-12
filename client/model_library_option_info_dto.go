/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// LibraryOptionInfoDto Library option info dto.
type LibraryOptionInfoDto struct {
	// Gets or sets name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets a value indicating whether default enabled.
	DefaultEnabled bool `json:"DefaultEnabled,omitempty"`
}
