/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// FFmpegLocation Enum describing the location of the FFmpeg tool.
type FFmpegLocation string

// List of FFmpegLocation
const (
	FFMPEGLOCATION_NOT_FOUND FFmpegLocation = "NotFound"
	FFMPEGLOCATION_SET_BY_ARGUMENT FFmpegLocation = "SetByArgument"
	FFMPEGLOCATION_CUSTOM FFmpegLocation = "Custom"
	FFMPEGLOCATION_SYSTEM FFmpegLocation = "System"
)