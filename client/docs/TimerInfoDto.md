# TimerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the recording. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** | Gets or sets the server identifier. | [optional] 
**ExternalId** | Pointer to **string** | Gets or sets the external identifier. | [optional] 
**ChannelId** | **string** | ChannelId of the recording. | [optional] 
**ExternalChannelId** | Pointer to **string** | Gets or sets the external channel identifier. | [optional] 
**ChannelName** | Pointer to **string** | ChannelName of the recording. | [optional] 
**ChannelPrimaryImageTag** | Pointer to **string** |  | [optional] 
**ProgramId** | Pointer to **string** | Gets or sets the program identifier. | [optional] 
**ExternalProgramId** | Pointer to **string** | Gets or sets the external program identifier. | [optional] 
**Name** | Pointer to **string** | Name of the recording. | [optional] 
**Overview** | Pointer to **string** | Description of the recording. | [optional] 
**StartDate** | [**time.Time**](time.Time.md) | The start date of the recording, in UTC. | [optional] 
**EndDate** | [**time.Time**](time.Time.md) | The end date of the recording, in UTC. | [optional] 
**ServiceName** | Pointer to **string** | Gets or sets the name of the service. | [optional] 
**Priority** | **int32** | Gets or sets the priority. | [optional] 
**PrePaddingSeconds** | **int32** | Gets or sets the pre padding seconds. | [optional] 
**PostPaddingSeconds** | **int32** | Gets or sets the post padding seconds. | [optional] 
**IsPrePaddingRequired** | **bool** | Gets or sets a value indicating whether this instance is pre padding required. | [optional] 
**ParentBackdropItemId** | Pointer to **string** | If the item does not have any backdrops, this will hold the Id of the Parent that has one. | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** | Gets or sets the parent backdrop image tags. | [optional] 
**IsPostPaddingRequired** | **bool** | Gets or sets a value indicating whether this instance is post padding required. | [optional] 
**KeepUntil** | [**KeepUntil**](KeepUntil.md) |  | [optional] 
**Status** | [**RecordingStatus**](RecordingStatus.md) |  | [optional] 
**SeriesTimerId** | Pointer to **string** | Gets or sets the series timer identifier. | [optional] 
**ExternalSeriesTimerId** | Pointer to **string** | Gets or sets the external series timer identifier. | [optional] 
**RunTimeTicks** | Pointer to **int64** | Gets or sets the run time ticks. | [optional] 
**ProgramInfo** | [**BaseItemDto**](BaseItemDto.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


