/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// LiveTvServiceInfo Class ServiceInfo.
type LiveTvServiceInfo struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the home page URL.
	HomePageUrl *string `json:"HomePageUrl,omitempty"`
	Status LiveTvServiceStatus `json:"Status,omitempty"`
	// Gets or sets the status message.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// Gets or sets the version.
	Version *string `json:"Version,omitempty"`
	// Gets or sets a value indicating whether this instance has update available.
	HasUpdateAvailable bool `json:"HasUpdateAvailable,omitempty"`
	// Gets or sets a value indicating whether this instance is visible.
	IsVisible bool `json:"IsVisible,omitempty"`
	Tuners *[]string `json:"Tuners,omitempty"`
}