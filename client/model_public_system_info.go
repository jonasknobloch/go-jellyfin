/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// PublicSystemInfo struct for PublicSystemInfo
type PublicSystemInfo struct {
	// Gets or sets the local address.
	LocalAddress *string `json:"LocalAddress,omitempty"`
	// Gets or sets the name of the server.
	ServerName *string `json:"ServerName,omitempty"`
	// Gets or sets the server version.
	Version *string `json:"Version,omitempty"`
	// Gets or sets the product name. This is the AssemblyProduct name.
	ProductName *string `json:"ProductName,omitempty"`
	// Gets or sets the operating system.
	OperatingSystem *string `json:"OperatingSystem,omitempty"`
	// Gets or sets the id.
	Id *string `json:"Id,omitempty"`
	// Gets or sets a value indicating whether the startup wizard is completed.
	StartupWizardCompleted *bool `json:"StartupWizardCompleted,omitempty"`
}
