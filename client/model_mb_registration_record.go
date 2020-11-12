/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
import (
	"time"
)
// MbRegistrationRecord MB Registration Record.
type MbRegistrationRecord struct {
	// Gets or sets expiration date.
	ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
	// Gets or sets a value indicating whether is registered.
	IsRegistered bool `json:"IsRegistered,omitempty"`
	// Gets or sets a value indicating whether reg checked.
	RegChecked bool `json:"RegChecked,omitempty"`
	// Gets or sets a value indicating whether reg error.
	RegError bool `json:"RegError,omitempty"`
	// Gets or sets a value indicating whether trial version.
	TrialVersion bool `json:"TrialVersion,omitempty"`
	// Gets or sets a value indicating whether is valid.
	IsValid bool `json:"IsValid,omitempty"`
}
