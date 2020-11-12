# SeriesTimerInfoDto

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
**RecordAnyTime** | **bool** | Gets or sets a value indicating whether [record any time]. | [optional] 
**SkipEpisodesInLibrary** | **bool** |  | [optional] 
**RecordAnyChannel** | **bool** | Gets or sets a value indicating whether [record any channel]. | [optional] 
**KeepUpTo** | **int32** |  | [optional] 
**RecordNewOnly** | **bool** | Gets or sets a value indicating whether [record new only]. | [optional] 
**Days** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | Gets or sets the days. | [optional] 
**DayPattern** | [**DayPattern**](DayPattern.md) |  | [optional] 
**ImageTags** | Pointer to [**BaseItemDtoImageTags**](BaseItemDto_ImageTags.md) |  | [optional] 
**ParentThumbItemId** | Pointer to **string** | Gets or sets the parent thumb item id. | [optional] 
**ParentThumbImageTag** | Pointer to **string** | Gets or sets the parent thumb image tag. | [optional] 
**ParentPrimaryImageItemId** | Pointer to **string** | Gets or sets the parent primary image item identifier. | [optional] 
**ParentPrimaryImageTag** | Pointer to **string** | Gets or sets the parent primary image tag. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


