# \HlsSegmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHlsAudioSegmentLegacyAac**](HlsSegmentApi.md#GetHlsAudioSegmentLegacyAac) | **Get** /Audio/{itemId}/hls/{segmentId}/stream.aac | Gets the specified audio segment for an audio item.
[**GetHlsAudioSegmentLegacyMp3**](HlsSegmentApi.md#GetHlsAudioSegmentLegacyMp3) | **Get** /Audio/{itemId}/hls/{segmentId}/stream.mp3 | Gets the specified audio segment for an audio item.
[**GetHlsPlaylistLegacy**](HlsSegmentApi.md#GetHlsPlaylistLegacy) | **Get** /Videos/{itemId}/hls/{playlistId}/stream.m3u8 | Gets a hls video playlist.
[**GetHlsVideoSegmentLegacy**](HlsSegmentApi.md#GetHlsVideoSegmentLegacy) | **Get** /Videos/{itemId}/hls/{playlistId}/{segmentId}.{segmentContainer} | Gets a hls video segment.
[**StopEncodingProcess**](HlsSegmentApi.md#StopEncodingProcess) | **Delete** /Videos/ActiveEncodings | Stops an active encoding.



## GetHlsAudioSegmentLegacyAac

> *os.File GetHlsAudioSegmentLegacyAac(ctx, itemId, segmentId)

Gets the specified audio segment for an audio item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| The item id. | 
**segmentId** | **string**| The segment id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsAudioSegmentLegacyMp3

> *os.File GetHlsAudioSegmentLegacyMp3(ctx, itemId, segmentId)

Gets the specified audio segment for an audio item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| The item id. | 
**segmentId** | **string**| The segment id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsPlaylistLegacy

> *os.File GetHlsPlaylistLegacy(ctx, itemId, playlistId)

Gets a hls video playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| The video id. | 
**playlistId** | **string**| The playlist id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsVideoSegmentLegacy

> *os.File GetHlsVideoSegmentLegacy(ctx, itemId, playlistId, segmentId, segmentContainer)

Gets a hls video segment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| The item id. | 
**playlistId** | **string**| The playlist id. | 
**segmentId** | **string**| The segment id. | 
**segmentContainer** | **string**| The segment container. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopEncodingProcess

> StopEncodingProcess(ctx, optional)

Stops an active encoding.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StopEncodingProcessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StopEncodingProcessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
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

