/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// DeviceProfile struct for DeviceProfile
type DeviceProfile struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	Identification DeviceIdentification `json:"Identification,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty"`
	ManufacturerUrl *string `json:"ManufacturerUrl,omitempty"`
	ModelName *string `json:"ModelName,omitempty"`
	ModelDescription *string `json:"ModelDescription,omitempty"`
	ModelNumber *string `json:"ModelNumber,omitempty"`
	ModelUrl *string `json:"ModelUrl,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty"`
	EnableAlbumArtInDidl bool `json:"EnableAlbumArtInDidl,omitempty"`
	EnableSingleAlbumArtLimit bool `json:"EnableSingleAlbumArtLimit,omitempty"`
	EnableSingleSubtitleLimit bool `json:"EnableSingleSubtitleLimit,omitempty"`
	SupportedMediaTypes *string `json:"SupportedMediaTypes,omitempty"`
	UserId *string `json:"UserId,omitempty"`
	AlbumArtPn *string `json:"AlbumArtPn,omitempty"`
	MaxAlbumArtWidth int32 `json:"MaxAlbumArtWidth,omitempty"`
	MaxAlbumArtHeight int32 `json:"MaxAlbumArtHeight,omitempty"`
	MaxIconWidth *int32 `json:"MaxIconWidth,omitempty"`
	MaxIconHeight *int32 `json:"MaxIconHeight,omitempty"`
	MaxStreamingBitrate *int64 `json:"MaxStreamingBitrate,omitempty"`
	MaxStaticBitrate *int64 `json:"MaxStaticBitrate,omitempty"`
	MusicStreamingTranscodingBitrate *int32 `json:"MusicStreamingTranscodingBitrate,omitempty"`
	MaxStaticMusicBitrate *int32 `json:"MaxStaticMusicBitrate,omitempty"`
	// Controls the content of the aggregationFlags element in the urn:schemas-sonycom:av namespace.
	SonyAggregationFlags *string `json:"SonyAggregationFlags,omitempty"`
	ProtocolInfo *string `json:"ProtocolInfo,omitempty"`
	TimelineOffsetSeconds int32 `json:"TimelineOffsetSeconds,omitempty"`
	RequiresPlainVideoItems bool `json:"RequiresPlainVideoItems,omitempty"`
	RequiresPlainFolders bool `json:"RequiresPlainFolders,omitempty"`
	EnableMSMediaReceiverRegistrar bool `json:"EnableMSMediaReceiverRegistrar,omitempty"`
	IgnoreTranscodeByteRangeRequests bool `json:"IgnoreTranscodeByteRangeRequests,omitempty"`
	XmlRootAttributes *[]XmlAttribute `json:"XmlRootAttributes,omitempty"`
	// Gets or sets the direct play profiles.
	DirectPlayProfiles *[]DirectPlayProfile `json:"DirectPlayProfiles,omitempty"`
	// Gets or sets the transcoding profiles.
	TranscodingProfiles *[]TranscodingProfile `json:"TranscodingProfiles,omitempty"`
	ContainerProfiles *[]ContainerProfile `json:"ContainerProfiles,omitempty"`
	CodecProfiles *[]CodecProfile `json:"CodecProfiles,omitempty"`
	ResponseProfiles *[]ResponseProfile `json:"ResponseProfiles,omitempty"`
	SubtitleProfiles *[]SubtitleProfile `json:"SubtitleProfiles,omitempty"`
}
