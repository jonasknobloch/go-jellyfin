/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ExternalIdInfo Represents the external id information for serialization to the client.
type ExternalIdInfo struct {
	// Gets or sets the display name of the external id provider (IE: IMDB, MusicBrainz, etc).
	Name *string `json:"Name,omitempty"`
	// Gets or sets the unique key for this id. This key should be unique across all providers.
	Key *string `json:"Key,omitempty"`
	Type ExternalIdMediaType `json:"Type,omitempty"`
	// Gets or sets the URL format string.
	UrlFormatString *string `json:"UrlFormatString,omitempty"`
}
