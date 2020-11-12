/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ControlResponse struct for ControlResponse
type ControlResponse struct {
	Headers *map[string]string `json:"Headers,omitempty"`
	Xml *string `json:"Xml,omitempty"`
	IsSuccessful bool `json:"IsSuccessful,omitempty"`
}
