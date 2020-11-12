# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | [**PlayerStateInfo**](PlayerStateInfo.md) |  | [optional] 
**AdditionalUsers** | Pointer to [**[]SessionUserInfo**](SessionUserInfo.md) |  | [optional] 
**Capabilities** | [**ClientCapabilities**](ClientCapabilities.md) |  | [optional] 
**RemoteEndPoint** | Pointer to **string** | Gets or sets the remote end point. | [optional] 
**PlayableMediaTypes** | Pointer to **[]string** | Gets or sets the playable media types. | [optional] [readonly] 
**Id** | Pointer to **string** | Gets or sets the id. | [optional] 
**UserId** | **string** | Gets or sets the user id. | [optional] 
**UserName** | Pointer to **string** | Gets or sets the username. | [optional] 
**Client** | Pointer to **string** | Gets or sets the type of the client. | [optional] 
**LastActivityDate** | [**time.Time**](time.Time.md) | Gets or sets the last activity date. | [optional] 
**LastPlaybackCheckIn** | [**time.Time**](time.Time.md) | Gets or sets the last playback check in. | [optional] 
**DeviceName** | Pointer to **string** | Gets or sets the name of the device. | [optional] 
**DeviceType** | Pointer to **string** | Gets or sets the type of the device. | [optional] 
**NowPlayingItem** | [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**FullNowPlayingItem** | [**BaseItem**](BaseItem.md) |  | [optional] 
**NowViewingItem** | [**BaseItemDto**](BaseItemDto.md) |  | [optional] 
**DeviceId** | Pointer to **string** | Gets or sets the device id. | [optional] 
**ApplicationVersion** | Pointer to **string** | Gets or sets the application version. | [optional] 
**TranscodingInfo** | [**TranscodingInfo**](TranscodingInfo.md) |  | [optional] 
**IsActive** | **bool** | Gets a value indicating whether this instance is active. | [optional] [readonly] 
**SupportsMediaControl** | **bool** |  | [optional] [readonly] 
**SupportsRemoteControl** | **bool** |  | [optional] [readonly] 
**NowPlayingQueue** | Pointer to [**[]QueueItem**](QueueItem.md) |  | [optional] 
**HasCustomDeviceName** | **bool** |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) | Gets or sets the supported commands. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


