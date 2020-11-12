/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// DeviceIdentification struct for DeviceIdentification
type DeviceIdentification struct {
	// Gets or sets the name of the friendly.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Gets or sets the model number.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// Gets or sets the serial number.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Gets or sets the name of the model.
	ModelName *string `json:"ModelName,omitempty"`
	// Gets or sets the model description.
	ModelDescription *string `json:"ModelDescription,omitempty"`
	// Gets or sets the model URL.
	ModelUrl *string `json:"ModelUrl,omitempty"`
	// Gets or sets the manufacturer.
	Manufacturer *string `json:"Manufacturer,omitempty"`
	// Gets or sets the manufacturer URL.
	ManufacturerUrl *string `json:"ManufacturerUrl,omitempty"`
	// Gets or sets the headers.
	Headers *[]HttpHeaderInfo `json:"Headers,omitempty"`
}
