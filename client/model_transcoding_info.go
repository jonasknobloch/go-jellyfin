/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// TranscodingInfo struct for TranscodingInfo
type TranscodingInfo struct {
	AudioCodec *string `json:"AudioCodec,omitempty"`
	VideoCodec *string `json:"VideoCodec,omitempty"`
	Container *string `json:"Container,omitempty"`
	IsVideoDirect bool `json:"IsVideoDirect,omitempty"`
	IsAudioDirect bool `json:"IsAudioDirect,omitempty"`
	Bitrate *int32 `json:"Bitrate,omitempty"`
	Framerate *float32 `json:"Framerate,omitempty"`
	CompletionPercentage *float64 `json:"CompletionPercentage,omitempty"`
	Width *int32 `json:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty"`
	AudioChannels *int32 `json:"AudioChannels,omitempty"`
	TranscodeReasons *[]TranscodeReason `json:"TranscodeReasons,omitempty"`
}
