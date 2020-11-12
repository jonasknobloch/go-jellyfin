# \VideosApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlternateSources**](VideosApi.md#DeleteAlternateSources) | **Delete** /Videos/{itemId}/AlternateSources | Removes alternate video sources.
[**GetAdditionalPart**](VideosApi.md#GetAdditionalPart) | **Get** /Videos/{itemId}/AdditionalParts | Gets additional parts for a video.
[**GetVideoStream**](VideosApi.md#GetVideoStream) | **Get** /Videos/{itemId}/stream | Gets a video stream.
[**GetVideoStreamWithExt**](VideosApi.md#GetVideoStreamWithExt) | **Get** /Videos/{itemId}/{stream}.{container} | Gets a video stream.
[**HeadVideoStream**](VideosApi.md#HeadVideoStream) | **Head** /Videos/{itemId}/stream | Gets a video stream.
[**HeadVideoStreamWithExt**](VideosApi.md#HeadVideoStreamWithExt) | **Head** /Videos/{itemId}/{stream}.{container} | Gets a video stream.
[**MergeVersions**](VideosApi.md#MergeVersions) | **Post** /Videos/MergeVersions | Merges videos into a single record.



## DeleteAlternateSources

> DeleteAlternateSources(ctx, itemId)

Removes alternate video sources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 

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


## GetAdditionalPart

> BaseItemDtoQueryResult GetAdditionalPart(ctx, itemId, optional)

Gets additional parts for a video.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetAdditionalPartOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAdditionalPartOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStream

> *os.File GetVideoStream(ctx, itemId, container, optional)

Gets a video stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
 **optional** | ***GetVideoStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVideoStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **static** | **optional.Bool**| Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **optional.String**| The streaming parameters. | 
 **tag** | **optional.String**| The tag. | 
 **deviceProfileId** | **optional.String**| Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **segmentContainer** | **optional.String**| The segment container. | 
 **segmentLength** | **optional.Int32**| The segment lenght. | 
 **minSegments** | **optional.Int32**| The minimum number of segments. | 
 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **optional.String**| Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **optional.Bool**| Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **optional.Bool**| Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **optional.Bool**| Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **optional.Int32**| Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **audioBitRate** | **optional.Int32**| Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **optional.Int32**| Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **optional.Int32**| Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **optional.String**| Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **optional.String**| Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **optional.Float32**| Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **optional.Float32**| Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **optional.Bool**| Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **optional.Int32**| Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **optional.Int32**| Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **optional.Int32**| Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **optional.Int32**| Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**optional.Interface of SubtitleDeliveryMethod**](.md)| Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **optional.Int32**| Optional. | 
 **maxVideoBitDepth** | **optional.Int32**| Optional. The maximum video bit depth. | 
 **requireAvc** | **optional.Bool**| Optional. Whether to require avc. | 
 **deInterlace** | **optional.Bool**| Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **optional.Bool**| Optional. Whether to require a non anamporphic stream. | 
 **transcodingMaxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **optional.Int32**| Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **enableMpegtsM2TsMode** | **optional.Bool**| Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **optional.String**| Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vpx, wmv. | 
 **subtitleCodec** | **optional.String**| Optional. Specify a subtitle codec to encode to. | 
 **transcodingReasons** | **optional.String**| Optional. The transcoding reason. | 
 **audioStreamIndex** | **optional.Int32**| Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **optional.Int32**| Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | [**optional.Interface of EncodingContext**](.md)| Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | [**optional.Interface of map[string]string**](string.md)| Optional. The streaming options. | 

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


## GetVideoStreamWithExt

> *os.File GetVideoStreamWithExt(ctx, itemId, container, stream, optional)

Gets a video stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
**stream** | **string**|  | 
 **optional** | ***GetVideoStreamWithExtOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVideoStreamWithExtOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **static** | **optional.Bool**| Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **optional.String**| The streaming parameters. | 
 **tag** | **optional.String**| The tag. | 
 **deviceProfileId** | **optional.String**| Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **segmentContainer** | **optional.String**| The segment container. | 
 **segmentLength** | **optional.Int32**| The segment lenght. | 
 **minSegments** | **optional.Int32**| The minimum number of segments. | 
 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **optional.String**| Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **optional.Bool**| Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **optional.Bool**| Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **optional.Bool**| Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **optional.Int32**| Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **audioBitRate** | **optional.Int32**| Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **optional.Int32**| Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **optional.Int32**| Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **optional.String**| Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **optional.String**| Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **optional.Float32**| Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **optional.Float32**| Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **optional.Bool**| Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **optional.Int32**| Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **optional.Int32**| Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **optional.Int32**| Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **optional.Int32**| Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**optional.Interface of SubtitleDeliveryMethod**](.md)| Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **optional.Int32**| Optional. | 
 **maxVideoBitDepth** | **optional.Int32**| Optional. The maximum video bit depth. | 
 **requireAvc** | **optional.Bool**| Optional. Whether to require avc. | 
 **deInterlace** | **optional.Bool**| Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **optional.Bool**| Optional. Whether to require a non anamporphic stream. | 
 **transcodingMaxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **optional.Int32**| Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **enableMpegtsM2TsMode** | **optional.Bool**| Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **optional.String**| Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vpx, wmv. | 
 **subtitleCodec** | **optional.String**| Optional. Specify a subtitle codec to encode to. | 
 **transcodingReasons** | **optional.String**| Optional. The transcoding reason. | 
 **audioStreamIndex** | **optional.Int32**| Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **optional.Int32**| Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | [**optional.Interface of EncodingContext**](.md)| Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | [**optional.Interface of map[string]string**](string.md)| Optional. The streaming options. | 

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


## HeadVideoStream

> *os.File HeadVideoStream(ctx, itemId, container, optional)

Gets a video stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
 **optional** | ***HeadVideoStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadVideoStreamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **static** | **optional.Bool**| Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **optional.String**| The streaming parameters. | 
 **tag** | **optional.String**| The tag. | 
 **deviceProfileId** | **optional.String**| Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **segmentContainer** | **optional.String**| The segment container. | 
 **segmentLength** | **optional.Int32**| The segment lenght. | 
 **minSegments** | **optional.Int32**| The minimum number of segments. | 
 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **optional.String**| Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **optional.Bool**| Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **optional.Bool**| Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **optional.Bool**| Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **optional.Int32**| Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **audioBitRate** | **optional.Int32**| Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **optional.Int32**| Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **optional.Int32**| Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **optional.String**| Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **optional.String**| Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **optional.Float32**| Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **optional.Float32**| Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **optional.Bool**| Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **optional.Int32**| Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **optional.Int32**| Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **optional.Int32**| Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **optional.Int32**| Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**optional.Interface of SubtitleDeliveryMethod**](.md)| Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **optional.Int32**| Optional. | 
 **maxVideoBitDepth** | **optional.Int32**| Optional. The maximum video bit depth. | 
 **requireAvc** | **optional.Bool**| Optional. Whether to require avc. | 
 **deInterlace** | **optional.Bool**| Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **optional.Bool**| Optional. Whether to require a non anamporphic stream. | 
 **transcodingMaxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **optional.Int32**| Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **enableMpegtsM2TsMode** | **optional.Bool**| Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **optional.String**| Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vpx, wmv. | 
 **subtitleCodec** | **optional.String**| Optional. Specify a subtitle codec to encode to. | 
 **transcodingReasons** | **optional.String**| Optional. The transcoding reason. | 
 **audioStreamIndex** | **optional.Int32**| Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **optional.Int32**| Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | [**optional.Interface of EncodingContext**](.md)| Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | [**optional.Interface of map[string]string**](string.md)| Optional. The streaming options. | 

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


## HeadVideoStreamWithExt

> *os.File HeadVideoStreamWithExt(ctx, itemId, container, stream, optional)

Gets a video stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
**stream** | **string**|  | 
 **optional** | ***HeadVideoStreamWithExtOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadVideoStreamWithExtOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **static** | **optional.Bool**| Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **optional.String**| The streaming parameters. | 
 **tag** | **optional.String**| The tag. | 
 **deviceProfileId** | **optional.String**| Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **optional.String**| The play session id. | 
 **segmentContainer** | **optional.String**| The segment container. | 
 **segmentLength** | **optional.Int32**| The segment lenght. | 
 **minSegments** | **optional.Int32**| The minimum number of segments. | 
 **mediaSourceId** | **optional.String**| The media version id, if playing an alternate version. | 
 **deviceId** | **optional.String**| The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **optional.String**| Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **optional.Bool**| Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **optional.Bool**| Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **optional.Bool**| Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **optional.Bool**| Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **optional.Int32**| Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **optional.Int32**| Optional. The maximum audio bit depth. | 
 **audioBitRate** | **optional.Int32**| Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **optional.Int32**| Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **optional.Int32**| Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **optional.String**| Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **optional.String**| Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **optional.Float32**| Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **optional.Float32**| Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **optional.Bool**| Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **optional.Int64**| Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **optional.Int32**| Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **optional.Int32**| Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **optional.Int32**| Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **optional.Int32**| Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**optional.Interface of SubtitleDeliveryMethod**](.md)| Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **optional.Int32**| Optional. | 
 **maxVideoBitDepth** | **optional.Int32**| Optional. The maximum video bit depth. | 
 **requireAvc** | **optional.Bool**| Optional. Whether to require avc. | 
 **deInterlace** | **optional.Bool**| Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **optional.Bool**| Optional. Whether to require a non anamporphic stream. | 
 **transcodingMaxAudioChannels** | **optional.Int32**| Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **optional.Int32**| Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **optional.String**| The live stream id. | 
 **enableMpegtsM2TsMode** | **optional.Bool**| Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **optional.String**| Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vpx, wmv. | 
 **subtitleCodec** | **optional.String**| Optional. Specify a subtitle codec to encode to. | 
 **transcodingReasons** | **optional.String**| Optional. The transcoding reason. | 
 **audioStreamIndex** | **optional.Int32**| Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **optional.Int32**| Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | [**optional.Interface of EncodingContext**](.md)| Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | [**optional.Interface of map[string]string**](string.md)| Optional. The streaming options. | 

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


## MergeVersions

> MergeVersions(ctx, itemIds)

Merges videos into a single record.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemIds** | **string**| Item id list. This allows multiple, comma delimited. | 

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

