/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SyncPlayAccess Enum SyncPlayAccess.
type SyncPlayAccess string

// List of SyncPlayAccess
const (
	SYNCPLAYACCESS_CREATE_AND_JOIN_GROUPS SyncPlayAccess = "CreateAndJoinGroups"
	SYNCPLAYACCESS_JOIN_GROUPS SyncPlayAccess = "JoinGroups"
	SYNCPLAYACCESS_NONE SyncPlayAccess = "None"
)