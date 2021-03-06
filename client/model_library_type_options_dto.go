/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// LibraryTypeOptionsDto Library type options dto.
type LibraryTypeOptionsDto struct {
	// Gets or sets the type.
	Type *string `json:"Type,omitempty"`
	// Gets or sets the metadata fetchers.
	MetadataFetchers *[]LibraryOptionInfoDto `json:"MetadataFetchers,omitempty"`
	// Gets or sets the image fetchers.
	ImageFetchers *[]LibraryOptionInfoDto `json:"ImageFetchers,omitempty"`
	// Gets or sets the supported image types.
	SupportedImageTypes *[]ImageType `json:"SupportedImageTypes,omitempty"`
	// Gets or sets the default image options.
	DefaultImageOptions *[]ImageOption `json:"DefaultImageOptions,omitempty"`
}
