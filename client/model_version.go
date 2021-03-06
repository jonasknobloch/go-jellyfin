/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Version struct for Version
type Version struct {
	Major int32 `json:"Major,omitempty"`
	Minor int32 `json:"Minor,omitempty"`
	Build int32 `json:"Build,omitempty"`
	Revision int32 `json:"Revision,omitempty"`
	MajorRevision int32 `json:"MajorRevision,omitempty"`
	MinorRevision int32 `json:"MinorRevision,omitempty"`
}
