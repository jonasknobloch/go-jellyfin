/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// LibraryOptions struct for LibraryOptions
type LibraryOptions struct {
	EnablePhotos bool `json:"EnablePhotos,omitempty"`
	EnableRealtimeMonitor bool `json:"EnableRealtimeMonitor,omitempty"`
	EnableChapterImageExtraction bool `json:"EnableChapterImageExtraction,omitempty"`
	ExtractChapterImagesDuringLibraryScan bool `json:"ExtractChapterImagesDuringLibraryScan,omitempty"`
	DownloadImagesInAdvance bool `json:"DownloadImagesInAdvance,omitempty"`
	PathInfos *[]MediaPathInfo `json:"PathInfos,omitempty"`
	SaveLocalMetadata bool `json:"SaveLocalMetadata,omitempty"`
	EnableInternetProviders bool `json:"EnableInternetProviders,omitempty"`
	EnableAutomaticSeriesGrouping bool `json:"EnableAutomaticSeriesGrouping,omitempty"`
	EnableEmbeddedTitles bool `json:"EnableEmbeddedTitles,omitempty"`
	EnableEmbeddedEpisodeInfos bool `json:"EnableEmbeddedEpisodeInfos,omitempty"`
	AutomaticRefreshIntervalDays int32 `json:"AutomaticRefreshIntervalDays,omitempty"`
	// Gets or sets the preferred metadata language.
	PreferredMetadataLanguage *string `json:"PreferredMetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode *string `json:"MetadataCountryCode,omitempty"`
	SeasonZeroDisplayName *string `json:"SeasonZeroDisplayName,omitempty"`
	MetadataSavers *[]string `json:"MetadataSavers,omitempty"`
	DisabledLocalMetadataReaders *[]string `json:"DisabledLocalMetadataReaders,omitempty"`
	LocalMetadataReaderOrder *[]string `json:"LocalMetadataReaderOrder,omitempty"`
	DisabledSubtitleFetchers *[]string `json:"DisabledSubtitleFetchers,omitempty"`
	SubtitleFetcherOrder *[]string `json:"SubtitleFetcherOrder,omitempty"`
	SkipSubtitlesIfEmbeddedSubtitlesPresent bool `json:"SkipSubtitlesIfEmbeddedSubtitlesPresent,omitempty"`
	SkipSubtitlesIfAudioTrackMatches bool `json:"SkipSubtitlesIfAudioTrackMatches,omitempty"`
	SubtitleDownloadLanguages *[]string `json:"SubtitleDownloadLanguages,omitempty"`
	RequirePerfectSubtitleMatch bool `json:"RequirePerfectSubtitleMatch,omitempty"`
	SaveSubtitlesWithMedia bool `json:"SaveSubtitlesWithMedia,omitempty"`
	TypeOptions *[]TypeOptions `json:"TypeOptions,omitempty"`
}
