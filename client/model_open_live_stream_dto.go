/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// OpenLiveStreamDto Open live stream dto.
type OpenLiveStreamDto struct {
	DeviceProfile DeviceProfile `json:"DeviceProfile,omitempty"`
	// Gets or sets the device play protocols.
	DirectPlayProtocols *[]MediaProtocol `json:"DirectPlayProtocols,omitempty"`
}
