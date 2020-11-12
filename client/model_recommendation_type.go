/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// RecommendationType the model 'RecommendationType'
type RecommendationType string

// List of RecommendationType
const (
	RECOMMENDATIONTYPE_SIMILAR_TO_RECENTLY_PLAYED RecommendationType = "SimilarToRecentlyPlayed"
	RECOMMENDATIONTYPE_SIMILAR_TO_LIKED_ITEM RecommendationType = "SimilarToLikedItem"
	RECOMMENDATIONTYPE_HAS_DIRECTOR_FROM_RECENTLY_PLAYED RecommendationType = "HasDirectorFromRecentlyPlayed"
	RECOMMENDATIONTYPE_HAS_ACTOR_FROM_RECENTLY_PLAYED RecommendationType = "HasActorFromRecentlyPlayed"
	RECOMMENDATIONTYPE_HAS_LIKED_DIRECTOR RecommendationType = "HasLikedDirector"
	RECOMMENDATIONTYPE_HAS_LIKED_ACTOR RecommendationType = "HasLikedActor"
)
