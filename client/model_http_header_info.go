/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// HttpHeaderInfo struct for HttpHeaderInfo
type HttpHeaderInfo struct {
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
	Match HeaderMatchType `json:"Match,omitempty"`
}
