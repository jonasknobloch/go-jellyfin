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
// AlbumInfo struct for AlbumInfo
type AlbumInfo struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the path.
	Path *string `json:"Path,omitempty"`
	// Gets or sets the metadata language.
	MetadataLanguage *string `json:"MetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode *string `json:"MetadataCountryCode,omitempty"`
	// Gets or sets the provider ids.
	ProviderIds *map[string]string `json:"ProviderIds,omitempty"`
	// Gets or sets the year.
	Year *int32 `json:"Year,omitempty"`
	IndexNumber *int32 `json:"IndexNumber,omitempty"`
	ParentIndexNumber *int32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate *time.Time `json:"PremiereDate,omitempty"`
	IsAutomated bool `json:"IsAutomated,omitempty"`
	// Gets or sets the album artist.
	AlbumArtists *[]string `json:"AlbumArtists,omitempty"`
	// Gets or sets the artist provider ids.
	ArtistProviderIds *map[string]string `json:"ArtistProviderIds,omitempty"`
	SongInfos *[]SongInfo `json:"SongInfos,omitempty"`
}