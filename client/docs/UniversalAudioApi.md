# \UniversalAudioApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniversalAudioStream**](UniversalAudioApi.md#GetUniversalAudioStream) | **Get** /Audio/{itemId}/universal | Gets an audio stream.
[**GetUniversalAudioStream2**](UniversalAudioApi.md#GetUniversalAudioStream2) | **Get** /Audio/{itemId}/universal.{container} | Gets an audio stream.
[**HeadUniversalAudioStream**](UniversalAudioApi.md#HeadUniversalAudioStream) | **Head** /Audio/{itemId}/universal | Gets an audio stream.
[**HeadUniversalAudioStream2**](UniversalAudioApi.md#HeadUniversalAudioStream2) | **Head** /Audio/{itemId}/universal.{container} | Gets an audio stream.



## GetUniversalAudioStream

> *os.File GetUniversalAudioStream(ctx, itemId, container, optional)

Gets an audio stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| Optional. The audio container. | 
 **optional** | ***GetUniversalAudioStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUniversalAudioStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. The user id. | 
 **audioCodec** | **optional.String**| Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **optional.Int32**| Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **optional.Int64**| Optional. The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **optional.String**| Optional. The container to transcode to. | 
 **transcodingProtocol** | **optional.String**| Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **optional.Int32**| Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **optional.Bool**| Optional. Whether to enable remote media. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **enableRedirection** | **optional.Bool**| Whether to enable redirection. Defaults to true. | [default to true]

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


## GetUniversalAudioStream2

> *os.File GetUniversalAudioStream2(ctx, itemId, container, optional)

Gets an audio stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| Optional. The audio container. | 
 **optional** | ***GetUniversalAudioStream2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUniversalAudioStream2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. The user id. | 
 **audioCodec** | **optional.String**| Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **optional.Int32**| Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **optional.Int64**| Optional. The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **optional.String**| Optional. The container to transcode to. | 
 **transcodingProtocol** | **optional.String**| Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **optional.Int32**| Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **optional.Bool**| Optional. Whether to enable remote media. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **enableRedirection** | **optional.Bool**| Whether to enable redirection. Defaults to true. | [default to true]

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


## HeadUniversalAudioStream

> *os.File HeadUniversalAudioStream(ctx, itemId, container, optional)

Gets an audio stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| Optional. The audio container. | 
 **optional** | ***HeadUniversalAudioStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadUniversalAudioStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. The user id. | 
 **audioCodec** | **optional.String**| Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **optional.Int32**| Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **optional.Int64**| Optional. The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **optional.String**| Optional. The container to transcode to. | 
 **transcodingProtocol** | **optional.String**| Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **optional.Int32**| Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **optional.Bool**| Optional. Whether to enable remote media. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **enableRedirection** | **optional.Bool**| Whether to enable redirection. Defaults to true. | [default to true]

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


## HeadUniversalAudioStream2

> *os.File HeadUniversalAudioStream2(ctx, itemId, container, optional)

Gets an audio stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| Optional. The audio container. | 
 **optional** | ***HeadUniversalAudioStream2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadUniversalAudioStream2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. The user id. | 
 **audioCodec** | **optional.String**| Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **optional.Int32**| Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **optional.Int64**| Optional. The maximum streaming bitrate. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **optional.String**| Optional. The container to transcode to. | 
 **transcodingProtocol** | **optional.String**| Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **optional.Int32**| Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **optional.Bool**| Optional. Whether to enable remote media. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **enableRedirection** | **optional.Bool**| Whether to enable redirection. Defaults to true. | [default to true]

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

