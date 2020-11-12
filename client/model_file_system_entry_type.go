/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// FileSystemEntryType Enum FileSystemEntryType.
type FileSystemEntryType string

// List of FileSystemEntryType
const (
	FILESYSTEMENTRYTYPE_FILE FileSystemEntryType = "File"
	FILESYSTEMENTRYTYPE_DIRECTORY FileSystemEntryType = "Directory"
	FILESYSTEMENTRYTYPE_NETWORK_COMPUTER FileSystemEntryType = "NetworkComputer"
	FILESYSTEMENTRYTYPE_NETWORK_SHARE FileSystemEntryType = "NetworkShare"
)
