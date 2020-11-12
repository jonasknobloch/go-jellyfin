/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// TypeOptions struct for TypeOptions
type TypeOptions struct {
	Type *string `json:"Type,omitempty"`
	MetadataFetchers *[]string `json:"MetadataFetchers,omitempty"`
	MetadataFetcherOrder *[]string `json:"MetadataFetcherOrder,omitempty"`
	ImageFetchers *[]string `json:"ImageFetchers,omitempty"`
	ImageFetcherOrder *[]string `json:"ImageFetcherOrder,omitempty"`
	ImageOptions *[]ImageOption `json:"ImageOptions,omitempty"`
}