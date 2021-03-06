/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// QueryFiltersLegacy struct for QueryFiltersLegacy
type QueryFiltersLegacy struct {
	Genres *[]string `json:"Genres,omitempty"`
	Tags *[]string `json:"Tags,omitempty"`
	OfficialRatings *[]string `json:"OfficialRatings,omitempty"`
	Years *[]int32 `json:"Years,omitempty"`
}
