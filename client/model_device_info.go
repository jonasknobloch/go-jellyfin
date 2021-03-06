/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// DeviceInfo struct for DeviceInfo
type DeviceInfo struct {
	Name *string `json:"Name,omitempty"`
	// Gets or sets the identifier.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the last name of the user.
	LastUserName *string `json:"LastUserName,omitempty"`
	// Gets or sets the name of the application.
	AppName *string `json:"AppName,omitempty"`
	// Gets or sets the application version.
	AppVersion *string `json:"AppVersion,omitempty"`
	// Gets or sets the last user identifier.
	LastUserId string `json:"LastUserId,omitempty"`
	// Gets or sets the date last modified.
	DateLastActivity time.Time `json:"DateLastActivity,omitempty"`
	Capabilities ClientCapabilities `json:"Capabilities,omitempty"`
	IconUrl *string `json:"IconUrl,omitempty"`
}
