# \MediaInfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseLiveStream**](MediaInfoApi.md#CloseLiveStream) | **Post** /LiveStreams/Close | Closes a media source.
[**GetBitrateTestBytes**](MediaInfoApi.md#GetBitrateTestBytes) | **Get** /Playback/BitrateTest | Tests the network with a request with the size of the bitrate.
[**GetPlaybackInfo**](MediaInfoApi.md#GetPlaybackInfo) | **Get** /Items/{itemId}/PlaybackInfo | Gets live playback media info for an item.
[**GetPostedPlaybackInfo**](MediaInfoApi.md#GetPostedPlaybackInfo) | **Post** /Items/{itemId}/PlaybackInfo | Gets live playback media info for an item.
[**OpenLiveStream**](MediaInfoApi.md#OpenLiveStream) | **Post** /LiveStreams/Open | Opens a media source.



## CloseLiveStream

> CloseLiveStream(ctx, liveStreamId)

Closes a media source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveStreamId** | **string**| The livestream id. | 

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


## GetBitrateTestBytes

> *os.File GetBitrateTestBytes(ctx, optional)

Tests the network with a request with the size of the bitrate.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBitrateTestBytesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBitrateTestBytesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **optional.Int32**| The bitrate. Defaults to 102400. | [default to 102400]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaybackInfo

> PlaybackInfoResponse GetPlaybackInfo(ctx, itemId, userId)

Gets live playback media info for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**userId** | [**string**](.md)| The user id. | 

### Return type

[**PlaybackInfoResponse**](PlaybackInfoResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostedPlaybackInfo

> PlaybackInfoResponse GetPostedPlaybackInfo(ctx, itemId, optional)

Gets live playback media info for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetPostedPlaybackInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPostedPlaybackInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| The user id. | 
 **maxStreamingBitrate** | **optional.Int64**| The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| The start time in ticks. | 
 **audioStreamIndex** | **optional.Int32**| The audio stream index. | 
 **subtitleStreamIndex** | **optional.Int32**| The subtitle stream index. | 
 **maxAudioChannels** | **optional.Int32**| The maximum number of audio channels. | 
 **mediaSourceId** | **optional.String**| The media source id. | 
 **liveStreamId** | **optional.String**| The livestream id. | 
 **autoOpenLiveStream** | **optional.Bool**| Whether to auto open the livestream. | [default to false]
 **enableDirectPlay** | **optional.Bool**| Whether to enable direct play. Default: true. | [default to true]
 **enableDirectStream** | **optional.Bool**| Whether to enable direct stream. Default: true. | [default to true]
 **enableTranscoding** | **optional.Bool**| Whether to enable transcoding. Default: true. | [default to true]
 **allowVideoStreamCopy** | **optional.Bool**| Whether to allow to copy the video stream. Default: true. | [default to true]
 **allowAudioStreamCopy** | **optional.Bool**| Whether to allow to copy the audio stream. Default: true. | [default to true]
 **deviceProfileDto** | [**optional.Interface of DeviceProfileDto**](DeviceProfileDto.md)| The device profile. | 

### Return type

[**PlaybackInfoResponse**](PlaybackInfoResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenLiveStream

> LiveStreamResponse OpenLiveStream(ctx, optional)

Opens a media source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenLiveStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OpenLiveStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openToken** | **optional.String**| The open token. | 
 **userId** | [**optional.Interface of string**](.md)| The user id. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **maxStreamingBitrate** | **optional.Int64**| The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| The start time in ticks. | 
 **audioStreamIndex** | **optional.Int32**| The audio stream index. | 
 **subtitleStreamIndex** | **optional.Int32**| The subtitle stream index. | 
 **maxAudioChannels** | **optional.Int32**| The maximum number of audio channels. | 
 **itemId** | [**optional.Interface of string**](.md)| The item id. | 
 **enableDirectPlay** | **optional.Bool**| Whether to enable direct play. Default: true. | [default to true]
 **enableDirectStream** | **optional.Bool**| Whether to enable direct stream. Default: true. | [default to true]
 **openLiveStreamDto** | [**optional.Interface of OpenLiveStreamDto**](OpenLiveStreamDto.md)| The open live stream dto. | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

