/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// InstallationInfo Class InstallationInfo.
type InstallationInfo struct {
	// Gets or sets the guid.
	Guid string `json:"Guid,omitempty"`
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	Version Version `json:"Version,omitempty"`
	// Gets or sets the changelog for this version.
	Changelog *string `json:"Changelog,omitempty"`
	// Gets or sets the source URL.
	SourceUrl *string `json:"SourceUrl,omitempty"`
	// Gets or sets a checksum for the binary.
	Checksum *string `json:"Checksum,omitempty"`
}