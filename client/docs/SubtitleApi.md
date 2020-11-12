# \SubtitleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubtitle**](SubtitleApi.md#DeleteSubtitle) | **Delete** /Videos/{itemId}/Subtitles/{index} | Deletes an external subtitle file.
[**DownloadRemoteSubtitles**](SubtitleApi.md#DownloadRemoteSubtitles) | **Post** /Items/{itemId}/RemoteSearch/Subtitles/{subtitleId} | Downloads a remote subtitle.
[**GetRemoteSubtitles**](SubtitleApi.md#GetRemoteSubtitles) | **Get** /Providers/Subtitles/Subtitles/{id} | Gets the remote subtitles.
[**GetSubtitle**](SubtitleApi.md#GetSubtitle) | **Get** /Videos/{itemId}/{mediaSourceId}/Subtitles/{index}/Stream.{format} | Gets subtitles in a specified format.
[**GetSubtitle2**](SubtitleApi.md#GetSubtitle2) | **Get** /Videos/{itemId}/{mediaSourceId}/Subtitles/{index}/{startPositionTicks}/Stream.{format} | Gets subtitles in a specified format.
[**GetSubtitlePlaylist**](SubtitleApi.md#GetSubtitlePlaylist) | **Get** /Videos/{itemId}/{mediaSourceId}/Subtitles/{index}/subtitles.m3u8 | Gets an HLS subtitle playlist.
[**SearchRemoteSubtitles**](SubtitleApi.md#SearchRemoteSubtitles) | **Get** /Items/{itemId}/RemoteSearch/Subtitles/{language} | Search remote subtitles.



## DeleteSubtitle

> DeleteSubtitle(ctx, itemId, index)

Deletes an external subtitle file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**index** | **int32**| The index of the subtitle file. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRemoteSubtitles

> DownloadRemoteSubtitles(ctx, itemId, subtitleId)

Downloads a remote subtitle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**subtitleId** | **string**| The subtitle id. | 

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


## GetRemoteSubtitles

> *os.File GetRemoteSubtitles(ctx, id)

Gets the remote subtitles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The item id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtitle

> *os.File GetSubtitle(ctx, itemId, mediaSourceId, index, format, startPositionTicks, optional)

Gets subtitles in a specified format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media source id. | 
**index** | **int32**| The subtitle stream index. | 
**format** | **string**| The format of the returned subtitle. | 
**startPositionTicks** | **int64**| Optional. The start position of the subtitle in ticks. | [default to 0]
 **optional** | ***GetSubtitleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubtitleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **endPositionTicks** | **optional.Int64**| Optional. The end position of the subtitle in ticks. | 
 **copyTimestamps** | **optional.Bool**| Optional. Whether to copy the timestamps. | [default to false]
 **addVttTimeMap** | **optional.Bool**| Optional. Whether to add a VTT time map. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtitle2

> *os.File GetSubtitle2(ctx, itemId, mediaSourceId, index, format, startPositionTicks, optional)

Gets subtitles in a specified format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media source id. | 
**index** | **int32**| The subtitle stream index. | 
**format** | **string**| The format of the returned subtitle. | 
**startPositionTicks** | **int64**| Optional. The start position of the subtitle in ticks. | [default to 0]
 **optional** | ***GetSubtitle2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubtitle2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **endPositionTicks** | **optional.Int64**| Optional. The end position of the subtitle in ticks. | 
 **copyTimestamps** | **optional.Bool**| Optional. Whether to copy the timestamps. | [default to false]
 **addVttTimeMap** | **optional.Bool**| Optional. Whether to add a VTT time map. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtitlePlaylist

> *os.File GetSubtitlePlaylist(ctx, itemId, index, mediaSourceId, segmentLength)

Gets an HLS subtitle playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**index** | **int32**| The subtitle stream index. | 
**mediaSourceId** | **string**| The media source id. | 
**segmentLength** | **int32**| The subtitle segment length. | 

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


## SearchRemoteSubtitles

> []RemoteSubtitleInfo SearchRemoteSubtitles(ctx, itemId, language, optional)

Search remote subtitles.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**language** | **string**| The language of the subtitles. | 
 **optional** | ***SearchRemoteSubtitlesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchRemoteSubtitlesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isPerfectMatch** | **optional.Bool**| Optional. Only show subtitles which are a perfect match. | 

### Return type

[**[]RemoteSubtitleInfo**](RemoteSubtitleInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

