/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ImageProviderInfo Class ImageProviderInfo.
type ImageProviderInfo struct {
	// Gets the name.
	Name *string `json:"Name,omitempty"`
	// Gets the supported image types.
	SupportedImages *[]ImageType `json:"SupportedImages,omitempty"`
}
