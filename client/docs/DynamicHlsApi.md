# \DynamicHlsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHlsAudioSegment**](DynamicHlsApi.md#GetHlsAudioSegment) | **Get** /Audio/{itemId}/hls1/{playlistId}/{segmentId}.{container} | Gets a video stream using HTTP live streaming.
[**GetHlsVideoSegment**](DynamicHlsApi.md#GetHlsVideoSegment) | **Get** /Videos/{itemId}/hls1/{playlistId}/{segmentId}.{container} | Gets a video stream using HTTP live streaming.
[**GetMasterHlsAudioPlaylist**](DynamicHlsApi.md#GetMasterHlsAudioPlaylist) | **Get** /Audio/{itemId}/master.m3u8 | Gets an audio hls playlist stream.
[**GetMasterHlsVideoPlaylist**](DynamicHlsApi.md#GetMasterHlsVideoPlaylist) | **Get** /Videos/{itemId}/master.m3u8 | Gets a video hls playlist stream.
[**GetVariantHlsAudioPlaylist**](DynamicHlsApi.md#GetVariantHlsAudioPlaylist) | **Get** /Audio/{itemId}/main.m3u8 | Gets an audio stream using HTTP live streaming.
[**GetVariantHlsVideoPlaylist**](DynamicHlsApi.md#GetVariantHlsVideoPlaylist) | **Get** /Videos/{itemId}/main.m3u8 | Gets a video stream using HTTP live streaming.
[**HeadMasterHlsAudioPlaylist**](DynamicHlsApi.md#HeadMasterHlsAudioPlaylist) | **Head** /Audio/{itemId}/master.m3u8 | Gets an audio hls playlist stream.
[**HeadMasterHlsVideoPlaylist**](DynamicHlsApi.md#HeadMasterHlsVideoPlaylist) | **Head** /Videos/{itemId}/master.m3u8 | Gets a video hls playlist stream.



## GetHlsAudioSegment

> *os.File GetHlsAudioSegment(ctx, itemId, playlistId, segmentId, container, optional)

Gets a video stream using HTTP live streaming.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**playlistId** | **string**| The playlist id. | 
**segmentId** | **int32**| The segment id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
 **optional** | ***GetHlsAudioSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHlsAudioSegmentOpts struct


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
- **Accept**: audio/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsVideoSegment

> *os.File GetHlsVideoSegment(ctx, itemId, playlistId, segmentId, container, optional)

Gets a video stream using HTTP live streaming.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**playlistId** | **string**| The playlist id. | 
**segmentId** | **int32**| The segment id. | 
**container** | **string**| The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
 **optional** | ***GetHlsVideoSegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHlsVideoSegmentOpts struct


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


## GetMasterHlsAudioPlaylist

> *os.File GetMasterHlsAudioPlaylist(ctx, itemId, mediaSourceId, optional)

Gets an audio hls playlist stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media version id, if playing an alternate version. | 
 **optional** | ***GetMasterHlsAudioPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMasterHlsAudioPlaylistOpts struct


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
 **enableAdaptiveBitrateStreaming** | **optional.Bool**| Enable adaptive bitrate streaming. | [default to true]

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


## GetMasterHlsVideoPlaylist

> *os.File GetMasterHlsVideoPlaylist(ctx, itemId, mediaSourceId, optional)

Gets a video hls playlist stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media version id, if playing an alternate version. | 
 **optional** | ***GetMasterHlsVideoPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMasterHlsVideoPlaylistOpts struct


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
 **enableAdaptiveBitrateStreaming** | **optional.Bool**| Enable adaptive bitrate streaming. | [default to true]

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


## GetVariantHlsAudioPlaylist

> *os.File GetVariantHlsAudioPlaylist(ctx, itemId, optional)

Gets an audio stream using HTTP live streaming.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetVariantHlsAudioPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVariantHlsAudioPlaylistOpts struct


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
- **Accept**: application/x-mpegURL

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantHlsVideoPlaylist

> *os.File GetVariantHlsVideoPlaylist(ctx, itemId, optional)

Gets a video stream using HTTP live streaming.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetVariantHlsVideoPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVariantHlsVideoPlaylistOpts struct


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
- **Accept**: application/x-mpegURL

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMasterHlsAudioPlaylist

> *os.File HeadMasterHlsAudioPlaylist(ctx, itemId, mediaSourceId, optional)

Gets an audio hls playlist stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media version id, if playing an alternate version. | 
 **optional** | ***HeadMasterHlsAudioPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadMasterHlsAudioPlaylistOpts struct


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
 **enableAdaptiveBitrateStreaming** | **optional.Bool**| Enable adaptive bitrate streaming. | [default to true]

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


## HeadMasterHlsVideoPlaylist

> *os.File HeadMasterHlsVideoPlaylist(ctx, itemId, mediaSourceId, optional)

Gets a video hls playlist stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**mediaSourceId** | **string**| The media version id, if playing an alternate version. | 
 **optional** | ***HeadMasterHlsVideoPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadMasterHlsVideoPlaylistOpts struct


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
 **enableAdaptiveBitrateStreaming** | **optional.Bool**| Enable adaptive bitrate streaming. | [default to true]

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

