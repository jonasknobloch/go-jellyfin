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
// RemoteSubtitleInfo struct for RemoteSubtitleInfo
type RemoteSubtitleInfo struct {
	ThreeLetterISOLanguageName *string `json:"ThreeLetterISOLanguageName,omitempty"`
	Id *string `json:"Id,omitempty"`
	ProviderName *string `json:"ProviderName,omitempty"`
	Name *string `json:"Name,omitempty"`
	Format *string `json:"Format,omitempty"`
	Author *string `json:"Author,omitempty"`
	Comment *string `json:"Comment,omitempty"`
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	CommunityRating *float32 `json:"CommunityRating,omitempty"`
	DownloadCount *int32 `json:"DownloadCount,omitempty"`
	IsHashMatch *bool `json:"IsHashMatch,omitempty"`
}