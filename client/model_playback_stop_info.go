/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// PlaybackStopInfo Class PlaybackStopInfo.
type PlaybackStopInfo struct {
	Item BaseItemDto `json:"Item,omitempty"`
	// Gets or sets the item identifier.
	ItemId string `json:"ItemId,omitempty"`
	// Gets or sets the session id.
	SessionId *string `json:"SessionId,omitempty"`
	// Gets or sets the media version identifier.
	MediaSourceId *string `json:"MediaSourceId,omitempty"`
	// Gets or sets the position ticks.
	PositionTicks *int64 `json:"PositionTicks,omitempty"`
	// Gets or sets the live stream identifier.
	LiveStreamId *string `json:"LiveStreamId,omitempty"`
	// Gets or sets the play session identifier.
	PlaySessionId *string `json:"PlaySessionId,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Session.PlaybackStopInfo is failed.
	Failed bool `json:"Failed,omitempty"`
	NextMediaType *string `json:"NextMediaType,omitempty"`
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
	NowPlayingQueue *[]QueueItem `json:"NowPlayingQueue,omitempty"`
}
