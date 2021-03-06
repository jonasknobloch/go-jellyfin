/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ImageInfo Class ImageInfo.
type ImageInfo struct {
	ImageType ImageType `json:"ImageType,omitempty"`
	// Gets or sets the index of the image.
	ImageIndex *int32 `json:"ImageIndex,omitempty"`
	// Gets or sets the image tag.
	ImageTag *string `json:"ImageTag,omitempty"`
	// Gets or sets the path.
	Path *string `json:"Path,omitempty"`
	// Gets or sets the blurhash.
	BlurHash *string `json:"BlurHash,omitempty"`
	// Gets or sets the height.
	Height *int32 `json:"Height,omitempty"`
	// Gets or sets the width.
	Width *int32 `json:"Width,omitempty"`
	// Gets or sets the size.
	Size int64 `json:"Size,omitempty"`
}
