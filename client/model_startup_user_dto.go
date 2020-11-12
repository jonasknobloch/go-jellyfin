/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// StartupUserDto The startup user DTO.
type StartupUserDto struct {
	// Gets or sets the username.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the user's password.
	Password *string `json:"Password,omitempty"`
}
