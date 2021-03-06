/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// ProfileConditionType the model 'ProfileConditionType'
type ProfileConditionType string

// List of ProfileConditionType
const (
	PROFILECONDITIONTYPE_EQUALS ProfileConditionType = "Equals"
	PROFILECONDITIONTYPE_NOT_EQUALS ProfileConditionType = "NotEquals"
	PROFILECONDITIONTYPE_LESS_THAN_EQUAL ProfileConditionType = "LessThanEqual"
	PROFILECONDITIONTYPE_GREATER_THAN_EQUAL ProfileConditionType = "GreaterThanEqual"
	PROFILECONDITIONTYPE_EQUALS_ANY ProfileConditionType = "EqualsAny"
)
