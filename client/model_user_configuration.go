/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// UserConfiguration Class UserConfiguration.
type UserConfiguration struct {
	// Gets or sets the audio language preference.
	AudioLanguagePreference *string `json:"AudioLanguagePreference,omitempty"`
	// Gets or sets a value indicating whether [play default audio track].
	PlayDefaultAudioTrack bool `json:"PlayDefaultAudioTrack,omitempty"`
	// Gets or sets the subtitle language preference.
	SubtitleLanguagePreference *string `json:"SubtitleLanguagePreference,omitempty"`
	DisplayMissingEpisodes bool `json:"DisplayMissingEpisodes,omitempty"`
	GroupedFolders *[]string `json:"GroupedFolders,omitempty"`
	SubtitleMode SubtitlePlaybackMode `json:"SubtitleMode,omitempty"`
	DisplayCollectionsView bool `json:"DisplayCollectionsView,omitempty"`
	EnableLocalPassword bool `json:"EnableLocalPassword,omitempty"`
	OrderedViews *[]string `json:"OrderedViews,omitempty"`
	LatestItemsExcludes *[]string `json:"LatestItemsExcludes,omitempty"`
	MyMediaExcludes *[]string `json:"MyMediaExcludes,omitempty"`
	HidePlayedInLatest bool `json:"HidePlayedInLatest,omitempty"`
	RememberAudioSelections bool `json:"RememberAudioSelections,omitempty"`
	RememberSubtitleSelections bool `json:"RememberSubtitleSelections,omitempty"`
	EnableNextEpisodeAutoPlay bool `json:"EnableNextEpisodeAutoPlay,omitempty"`
}
