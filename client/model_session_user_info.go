/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SessionUserInfo Class SessionUserInfo.
type SessionUserInfo struct {
	// Gets or sets the user identifier.
	UserId string `json:"UserId,omitempty"`
	// Gets or sets the name of the user.
	UserName *string `json:"UserName,omitempty"`
}
