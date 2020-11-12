# \PlaystateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkPlayedItem**](PlaystateApi.md#MarkPlayedItem) | **Post** /Users/{userId}/PlayedItems/{itemId} | Marks an item as played for user.
[**MarkUnplayedItem**](PlaystateApi.md#MarkUnplayedItem) | **Delete** /Users/{userId}/PlayedItems/{itemId} | Marks an item as unplayed for user.
[**OnPlaybackProgress**](PlaystateApi.md#OnPlaybackProgress) | **Post** /Users/{userId}/PlayingItems/{itemId}/Progress | Reports a user&#39;s playback progress.
[**OnPlaybackStart**](PlaystateApi.md#OnPlaybackStart) | **Post** /Users/{userId}/PlayingItems/{itemId} | Reports that a user has begun playing an item.
[**OnPlaybackStopped**](PlaystateApi.md#OnPlaybackStopped) | **Delete** /Users/{userId}/PlayingItems/{itemId} | Reports that a user has stopped playing an item.
[**PingPlaybackSession**](PlaystateApi.md#PingPlaybackSession) | **Post** /Sessions/Playing/Ping | Pings a playback session.
[**ReportPlaybackProgress**](PlaystateApi.md#ReportPlaybackProgress) | **Post** /Sessions/Playing/Progress | Reports playback progress within a session.
[**ReportPlaybackStart**](PlaystateApi.md#ReportPlaybackStart) | **Post** /Sessions/Playing | Reports playback has started within a session.
[**ReportPlaybackStopped**](PlaystateApi.md#ReportPlaybackStopped) | **Post** /Sessions/Playing/Stopped | Reports playback has stopped within a session.



## MarkPlayedItem

> UserItemDataDto MarkPlayedItem(ctx, userId, itemId, optional)

Marks an item as played for user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***MarkPlayedItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MarkPlayedItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datePlayed** | **optional.Time**| Optional. The date the item was played. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkUnplayedItem

> UserItemDataDto MarkUnplayedItem(ctx, userId, itemId)

Marks an item as unplayed for user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPlaybackProgress

> OnPlaybackProgress(ctx, userId, itemId, optional)

Reports a user's playback progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***OnPlaybackProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OnPlaybackProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The id of the MediaSource. | 
 **positionTicks** | **optional.Int64**| Optional. The current position, in ticks. 1 tick &#x3D; 10000 ms. | 
 **audioStreamIndex** | **optional.Int32**| The audio stream index. | 
 **subtitleStreamIndex** | **optional.Int32**| The subtitle stream index. | 
 **volumeLevel** | **optional.Int32**| Scale of 0-100. | 
 **playMethod** | [**optional.Interface of PlayMethod**](.md)| The play method. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **repeatMode** | [**optional.Interface of RepeatMode**](.md)| The repeat mode. | 
 **isPaused** | **optional.Bool**| Indicates if the player is paused. | [default to false]
 **isMuted** | **optional.Bool**| Indicates if the player is muted. | [default to false]

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPlaybackStart

> OnPlaybackStart(ctx, userId, itemId, optional)

Reports that a user has begun playing an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***OnPlaybackStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OnPlaybackStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The id of the MediaSource. | 
 **audioStreamIndex** | **optional.Int32**| The audio stream index. | 
 **subtitleStreamIndex** | **optional.Int32**| The subtitle stream index. | 
 **playMethod** | [**optional.Interface of PlayMethod**](.md)| The play method. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **canSeek** | **optional.Bool**| Indicates if the client can seek. | [default to false]

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPlaybackStopped

> OnPlaybackStopped(ctx, userId, itemId, optional)

Reports that a user has stopped playing an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***OnPlaybackStoppedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OnPlaybackStoppedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The id of the MediaSource. | 
 **nextMediaType** | **optional.String**| The next media type that will play. | 
 **positionTicks** | **optional.Int64**| Optional. The position, in ticks, where playback stopped. 1 tick &#x3D; 10000 ms. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **playSessionId** | **optional.String**| The play session id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingPlaybackSession

> PingPlaybackSession(ctx, optional)

Pings a playback session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PingPlaybackSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PingPlaybackSessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playSessionId** | **optional.String**| Playback session id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportPlaybackProgress

> ReportPlaybackProgress(ctx, optional)

Reports playback progress within a session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportPlaybackProgressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportPlaybackProgressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackProgressInfo** | [**optional.Interface of PlaybackProgressInfo**](PlaybackProgressInfo.md)| The playback progress info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportPlaybackStart

> ReportPlaybackStart(ctx, optional)

Reports playback has started within a session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportPlaybackStartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportPlaybackStartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStartInfo** | [**optional.Interface of PlaybackStartInfo**](PlaybackStartInfo.md)| The playback start info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportPlaybackStopped

> ReportPlaybackStopped(ctx, optional)

Reports playback has stopped within a session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportPlaybackStoppedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportPlaybackStoppedOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStopInfo** | [**optional.Interface of PlaybackStopInfo**](PlaybackStopInfo.md)| The playback stop info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

