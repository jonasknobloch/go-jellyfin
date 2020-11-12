/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// MetadataEditorInfo struct for MetadataEditorInfo
type MetadataEditorInfo struct {
	ParentalRatingOptions *[]ParentalRating `json:"ParentalRatingOptions,omitempty"`
	Countries *[]CountryInfo `json:"Countries,omitempty"`
	Cultures *[]CultureDto `json:"Cultures,omitempty"`
	ExternalIdInfos *[]ExternalIdInfo `json:"ExternalIdInfos,omitempty"`
	ContentType *string `json:"ContentType,omitempty"`
	ContentTypeOptions *[]NameValuePair `json:"ContentTypeOptions,omitempty"`
}
