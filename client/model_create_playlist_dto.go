/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// CreatePlaylistDto Create new playlist dto.
type CreatePlaylistDto struct {
	// Gets or sets the name of the new playlist.
	Name *string `json:"Name,omitempty"`
	// Gets or sets item ids to add to the playlist.
	Ids *string `json:"Ids,omitempty"`
	// Gets or sets the user id.
	UserId string `json:"UserId,omitempty"`
	// Gets or sets the media type.
	MediaType *string `json:"MediaType,omitempty"`
}