# MediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**EncoderPath** | Pointer to **string** |  | [optional] 
**EncoderProtocol** | [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Type** | [**MediaSourceType**](MediaSourceType.md) |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsRemote** | **bool** | Differentiate internet url vs local network. | [optional] 
**ETag** | Pointer to **string** |  | [optional] 
**RunTimeTicks** | Pointer to **int64** |  | [optional] 
**ReadAtNativeFramerate** | **bool** |  | [optional] 
**IgnoreDts** | **bool** |  | [optional] 
**IgnoreIndex** | **bool** |  | [optional] 
**GenPtsInput** | **bool** |  | [optional] 
**SupportsTranscoding** | **bool** |  | [optional] 
**SupportsDirectStream** | **bool** |  | [optional] 
**SupportsDirectPlay** | **bool** |  | [optional] 
**IsInfiniteStream** | **bool** |  | [optional] 
**RequiresOpening** | **bool** |  | [optional] 
**OpenToken** | Pointer to **string** |  | [optional] 
**RequiresClosing** | **bool** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**BufferMs** | Pointer to **int32** |  | [optional] 
**RequiresLooping** | **bool** |  | [optional] 
**SupportsProbing** | **bool** |  | [optional] 
**VideoType** | [**VideoType**](VideoType.md) |  | [optional] 
**IsoType** | [**IsoType**](IsoType.md) |  | [optional] 
**Video3DFormat** | [**Video3DFormat**](Video3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]MediaStream**](MediaStream.md) |  | [optional] 
**MediaAttachments** | Pointer to [**[]MediaAttachment**](MediaAttachment.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **int32** |  | [optional] 
**Timestamp** | [**TransportStreamTimestamp**](TransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**TranscodingUrl** | Pointer to **string** |  | [optional] 
**TranscodingSubProtocol** | Pointer to **string** |  | [optional] 
**TranscodingContainer** | Pointer to **string** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **int32** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **int32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **int32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


