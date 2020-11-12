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
// NotificationDto The notification DTO.
type NotificationDto struct {
	// Gets or sets the notification ID. Defaults to an empty string.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the notification's user ID. Defaults to an empty string.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets the notification date.
	Date time.Time `json:"Date,omitempty"`
	// Gets or sets a value indicating whether the notification has been read. Defaults to false.
	IsRead bool `json:"IsRead,omitempty"`
	// Gets or sets the notification's name. Defaults to an empty string.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the notification's description. Defaults to an empty string.
	Description *string `json:"Description,omitempty"`
	// Gets or sets the notification's URL. Defaults to an empty string.
	Url *string `json:"Url,omitempty"`
	Level NotificationLevel `json:"Level,omitempty"`
}
