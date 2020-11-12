# \LiveTvApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddListingProvider**](LiveTvApi.md#AddListingProvider) | **Post** /LiveTv/ListingProviders | Adds a listings provider.
[**AddTunerHost**](LiveTvApi.md#AddTunerHost) | **Post** /LiveTv/TunerHosts | Adds a tuner host.
[**CancelSeriesTimer**](LiveTvApi.md#CancelSeriesTimer) | **Delete** /LiveTv/SeriesTimers/{timerId} | Cancels a live tv series timer.
[**CancelTimer**](LiveTvApi.md#CancelTimer) | **Delete** /LiveTv/Timers/{timerId} | Cancels a live tv timer.
[**CreateSeriesTimer**](LiveTvApi.md#CreateSeriesTimer) | **Post** /LiveTv/SeriesTimers | Creates a live tv series timer.
[**CreateTimer**](LiveTvApi.md#CreateTimer) | **Post** /LiveTv/Timers | Creates a live tv timer.
[**DeleteListingProvider**](LiveTvApi.md#DeleteListingProvider) | **Delete** /LiveTv/ListingProviders | Delete listing provider.
[**DeleteRecording**](LiveTvApi.md#DeleteRecording) | **Delete** /LiveTv/Recordings/{recordingId} | Deletes a live tv recording.
[**DeleteTunerHost**](LiveTvApi.md#DeleteTunerHost) | **Delete** /LiveTv/TunerHosts | Deletes a tuner host.
[**DiscoverTuners**](LiveTvApi.md#DiscoverTuners) | **Get** /LiveTv/Tuners/Discvover | Discover tuners.
[**GetChannel**](LiveTvApi.md#GetChannel) | **Get** /LiveTv/Channels/{channelId} | Gets a live tv channel.
[**GetChannelMappingOptions**](LiveTvApi.md#GetChannelMappingOptions) | **Get** /LiveTv/ChannelMappingOptions | Get channel mapping options.
[**GetDefaultListingProvider**](LiveTvApi.md#GetDefaultListingProvider) | **Get** /LiveTv/ListingProviders/Default | Gets default listings provider info.
[**GetDefaultTimer**](LiveTvApi.md#GetDefaultTimer) | **Get** /LiveTv/Timers/Defaults | Gets the default values for a new timer.
[**GetGuideInfo**](LiveTvApi.md#GetGuideInfo) | **Get** /LiveTv/GuideInfo | Get guid info.
[**GetLineups**](LiveTvApi.md#GetLineups) | **Get** /LiveTv/ListingProviders/Lineups | Gets available lineups.
[**GetLiveRecordingFile**](LiveTvApi.md#GetLiveRecordingFile) | **Get** /LiveTv/LiveRecordings/{recordingId}/stream | Gets a live tv recording stream.
[**GetLiveStreamFile**](LiveTvApi.md#GetLiveStreamFile) | **Get** /LiveTv/LiveStreamFiles/{streamId}/stream.{container} | Gets a live tv channel stream.
[**GetLiveTvChannels**](LiveTvApi.md#GetLiveTvChannels) | **Get** /LiveTv/Channels | Gets available live tv channels.
[**GetLiveTvInfo**](LiveTvApi.md#GetLiveTvInfo) | **Get** /LiveTv/Info | Gets available live tv services.
[**GetLiveTvPrograms**](LiveTvApi.md#GetLiveTvPrograms) | **Get** /LiveTv/Programs | Gets available live tv epgs.
[**GetProgram**](LiveTvApi.md#GetProgram) | **Get** /LiveTv/Programs/{programId} | Gets a live tv program.
[**GetPrograms**](LiveTvApi.md#GetPrograms) | **Post** /LiveTv/Programs | Gets available live tv epgs.
[**GetRecommendedPrograms**](LiveTvApi.md#GetRecommendedPrograms) | **Get** /LiveTv/Programs/Recommended | Gets recommended live tv epgs.
[**GetRecording**](LiveTvApi.md#GetRecording) | **Get** /LiveTv/Recordings/{recordingId} | Gets a live tv recording.
[**GetRecordingFolders**](LiveTvApi.md#GetRecordingFolders) | **Get** /LiveTv/Recordings/Folders | Gets recording folders.
[**GetRecordingGroup**](LiveTvApi.md#GetRecordingGroup) | **Get** /LiveTv/Recordings/Groups/{groupId} | Get recording group.
[**GetRecordingGroups**](LiveTvApi.md#GetRecordingGroups) | **Get** /LiveTv/Recordings/Groups | Gets live tv recording groups.
[**GetRecordings**](LiveTvApi.md#GetRecordings) | **Get** /LiveTv/Recordings | Gets live tv recordings.
[**GetRecordingsSeries**](LiveTvApi.md#GetRecordingsSeries) | **Get** /LiveTv/Recordings/Series | Gets live tv recording series.
[**GetSchedulesDirectCountries**](LiveTvApi.md#GetSchedulesDirectCountries) | **Get** /LiveTv/ListingProviders/SchedulesDirect/Countries | Gets available countries.
[**GetSeriesTimer**](LiveTvApi.md#GetSeriesTimer) | **Get** /LiveTv/SeriesTimers/{timerId} | Gets a live tv series timer.
[**GetSeriesTimers**](LiveTvApi.md#GetSeriesTimers) | **Get** /LiveTv/SeriesTimers | Gets live tv series timers.
[**GetTimer**](LiveTvApi.md#GetTimer) | **Get** /LiveTv/Timers/{timerId} | Gets a timer.
[**GetTimers**](LiveTvApi.md#GetTimers) | **Get** /LiveTv/Timers | Gets the live tv timers.
[**GetTunerHostTypes**](LiveTvApi.md#GetTunerHostTypes) | **Get** /LiveTv/TunerHosts/Types | Get tuner host types.
[**ResetTuner**](LiveTvApi.md#ResetTuner) | **Post** /LiveTv/Tuners/{tunerId}/Reset | Resets a tv tuner.
[**SetChannelMapping**](LiveTvApi.md#SetChannelMapping) | **Post** /LiveTv/ChannelMappings | Set channel mappings.
[**UpdateSeriesTimer**](LiveTvApi.md#UpdateSeriesTimer) | **Post** /LiveTv/SeriesTimers/{timerId} | Updates a live tv series timer.
[**UpdateTimer**](LiveTvApi.md#UpdateTimer) | **Post** /LiveTv/Timers/{timerId} | Updates a live tv timer.



## AddListingProvider

> ListingsProviderInfo AddListingProvider(ctx, optional)

Adds a listings provider.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddListingProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddListingProviderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pw** | **optional.String**| Password. | 
 **validateListings** | **optional.Bool**| Validate listings. | [default to false]
 **validateLogin** | **optional.Bool**| Validate login. | [default to false]
 **listingsProviderInfo** | [**optional.Interface of ListingsProviderInfo**](ListingsProviderInfo.md)| New listings info. | 

### Return type

[**ListingsProviderInfo**](ListingsProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTunerHost

> TunerHostInfo AddTunerHost(ctx, optional)

Adds a tuner host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddTunerHostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddTunerHostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tunerHostInfo** | [**optional.Interface of TunerHostInfo**](TunerHostInfo.md)| New tuner host. | 

### Return type

[**TunerHostInfo**](TunerHostInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSeriesTimer

> CancelSeriesTimer(ctx, timerId)

Cancels a live tv series timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 

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


## CancelTimer

> CancelTimer(ctx, timerId)

Cancels a live tv timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 

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


## CreateSeriesTimer

> CreateSeriesTimer(ctx, optional)

Creates a live tv series timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSeriesTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSeriesTimerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesTimerInfoDto** | [**optional.Interface of SeriesTimerInfoDto**](SeriesTimerInfoDto.md)| New series timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTimer

> CreateTimer(ctx, optional)

Creates a live tv timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTimerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timerInfoDto** | [**optional.Interface of TimerInfoDto**](TimerInfoDto.md)| New timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListingProvider

> DeleteListingProvider(ctx, optional)

Delete listing provider.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteListingProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteListingProviderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Listing provider id. | 

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


## DeleteRecording

> DeleteRecording(ctx, recordingId)

Deletes a live tv recording.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | [**string**](.md)| Recording id. | 

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


## DeleteTunerHost

> DeleteTunerHost(ctx, optional)

Deletes a tuner host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteTunerHostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTunerHostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Tuner host id. | 

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


## DiscoverTuners

> []TunerHostInfo DiscoverTuners(ctx, optional)

Discover tuners.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiscoverTunersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DiscoverTunersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDevicesOnly** | **optional.Bool**| Only discover new tuners. | [default to false]

### Return type

[**[]TunerHostInfo**](TunerHostInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannel

> BaseItemDto GetChannel(ctx, channelId, optional)

Gets a live tv channel.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | [**string**](.md)| Channel id. | 
 **optional** | ***GetChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Attach user data. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelMappingOptions

> ChannelMappingOptionsDto GetChannelMappingOptions(ctx, optional)

Get channel mapping options.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetChannelMappingOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelMappingOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **optional.String**| Provider id. | 

### Return type

[**ChannelMappingOptionsDto**](ChannelMappingOptionsDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultListingProvider

> ListingsProviderInfo GetDefaultListingProvider(ctx, )

Gets default listings provider info.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ListingsProviderInfo**](ListingsProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTimer

> SeriesTimerInfoDto GetDefaultTimer(ctx, optional)

Gets the default values for a new timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDefaultTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDefaultTimerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **optional.String**| Optional. To attach default values based on a program. | 

### Return type

[**SeriesTimerInfoDto**](SeriesTimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideInfo

> GuideInfo GetGuideInfo(ctx, )

Get guid info.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GuideInfo**](GuideInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLineups

> []NameIdPair GetLineups(ctx, optional)

Gets available lineups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLineupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLineupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Provider id. | 
 **type_** | **optional.String**| Provider type. | 
 **location** | **optional.String**| Location. | 
 **country** | **optional.String**| Country. | 

### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveRecordingFile

> *os.File GetLiveRecordingFile(ctx, recordingId)

Gets a live tv recording stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | **string**| Recording id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveStreamFile

> *os.File GetLiveStreamFile(ctx, streamId, container)

Gets a live tv channel stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string**| Stream id. | 
**container** | **string**| Container type. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveTvChannels

> BaseItemDtoQueryResult GetLiveTvChannels(ctx, optional)

Gets available live tv channels.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLiveTvChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLiveTvChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ChannelType**](.md)| Optional. Filter by channel type. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user and attach user data. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **isMovie** | **optional.Bool**| Optional. Filter for movies. | 
 **isSeries** | **optional.Bool**| Optional. Filter for series. | 
 **isNews** | **optional.Bool**| Optional. Filter for news. | 
 **isKids** | **optional.Bool**| Optional. Filter for kids. | 
 **isSports** | **optional.Bool**| Optional. Filter for sports. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **isFavorite** | **optional.Bool**| Optional. Filter by channels that are favorites, or not. | 
 **isLiked** | **optional.Bool**| Optional. Filter by channels that are liked, or not. | 
 **isDisliked** | **optional.Bool**| Optional. Filter by channels that are disliked, or not. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| \&quot;Optional. The image types to include in the output. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **sortBy** | **optional.String**| Optional. Key to sort by. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| Optional. Sort order. | 
 **enableFavoriteSorting** | **optional.Bool**| Optional. Incorporate favorite and like status into channel sorting. | [default to false]
 **addCurrentProgram** | **optional.Bool**| Optional. Adds current program info to each channel. | [default to true]

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


## GetLiveTvInfo

> LiveTvInfo GetLiveTvInfo(ctx, )

Gets available live tv services.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LiveTvInfo**](LiveTvInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveTvPrograms

> BaseItemDtoQueryResult GetLiveTvPrograms(ctx, optional)

Gets available live tv epgs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLiveTvProgramsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLiveTvProgramsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **optional.String**| The channels to return guide information for. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id. | 
 **minStartDate** | **optional.Time**| Optional. The minimum premiere start date. | 
 **hasAired** | **optional.Bool**| Optional. Filter by programs that have completed airing, or not. | 
 **isAiring** | **optional.Bool**| Optional. Filter by programs that are currently airing, or not. | 
 **maxStartDate** | **optional.Time**| Optional. The maximum premiere start date. | 
 **minEndDate** | **optional.Time**| Optional. The minimum premiere end date. | 
 **maxEndDate** | **optional.Time**| Optional. The maximum premiere end date. | 
 **isMovie** | **optional.Bool**| Optional. Filter for movies. | 
 **isSeries** | **optional.Bool**| Optional. Filter for series. | 
 **isNews** | **optional.Bool**| Optional. Filter for news. | 
 **isKids** | **optional.Bool**| Optional. Filter for kids. | 
 **isSports** | **optional.Bool**| Optional. Filter for sports. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **sortBy** | **optional.String**| Optional. Specify one or more sort orders, comma delimited. Options: Name, StartDate. | 
 **sortOrder** | **optional.String**| Sort Order - Ascending,Descending. | 
 **genres** | **optional.String**| The genres to return guide information for. | 
 **genreIds** | **optional.String**| The genre ids to return guide information for. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **seriesTimerId** | **optional.String**| Optional. Filter by series timer id. | 
 **librarySeriesId** | [**optional.Interface of string**](.md)| Optional. Filter by library series id. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableTotalRecordCount** | **optional.Bool**| Retrieve total record count. | [default to true]

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


## GetProgram

> BaseItemDto GetProgram(ctx, programId, optional)

Gets a live tv program.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programId** | **string**| Program id. | 
 **optional** | ***GetProgramOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProgramOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Attach user data. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrograms

> BaseItemDtoQueryResult GetPrograms(ctx, optional)

Gets available live tv epgs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetProgramsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetProgramsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProgramsDto** | [**optional.Interface of GetProgramsDto**](GetProgramsDto.md)| Request body. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedPrograms

> BaseItemDtoQueryResult GetRecommendedPrograms(ctx, optional)

Gets recommended live tv epgs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRecommendedProgramsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecommendedProgramsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. filter by user id. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **isAiring** | **optional.Bool**| Optional. Filter by programs that are currently airing, or not. | 
 **hasAired** | **optional.Bool**| Optional. Filter by programs that have completed airing, or not. | 
 **isSeries** | **optional.Bool**| Optional. Filter for series. | 
 **isMovie** | **optional.Bool**| Optional. Filter for movies. | 
 **isNews** | **optional.Bool**| Optional. Filter for news. | 
 **isKids** | **optional.Bool**| Optional. Filter for kids. | 
 **isSports** | **optional.Bool**| Optional. Filter for sports. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **genreIds** | **optional.String**| The genres to return guide information for. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableUserData** | **optional.Bool**| Optional. include user data. | 
 **enableTotalRecordCount** | **optional.Bool**| Retrieve total record count. | [default to true]

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


## GetRecording

> BaseItemDto GetRecording(ctx, recordingId, optional)

Gets a live tv recording.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordingId** | [**string**](.md)| Recording id. | 
 **optional** | ***GetRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecordingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Attach user data. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordingFolders

> BaseItemDtoQueryResult GetRecordingFolders(ctx, optional)

Gets recording folders.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRecordingFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecordingFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user and attach user data. | 

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


## GetRecordingGroup

> GetRecordingGroup(ctx, groupId)

Get recording group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md)| Group id. | 

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


## GetRecordingGroups

> BaseItemDtoQueryResult GetRecordingGroups(ctx, optional)

Gets live tv recording groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRecordingGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecordingGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user and attach user data. | 

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


## GetRecordings

> BaseItemDtoQueryResult GetRecordings(ctx, optional)

Gets live tv recordings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRecordingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecordingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**| Optional. Filter by channel id. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user and attach user data. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **status** | [**optional.Interface of RecordingStatus**](.md)| Optional. Filter by recording status. | 
 **isInProgress** | **optional.Bool**| Optional. Filter by recordings that are in progress, or not. | 
 **seriesTimerId** | **optional.String**| Optional. Filter by recordings belonging to a series timer. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **isMovie** | **optional.Bool**| Optional. Filter for movies. | 
 **isSeries** | **optional.Bool**| Optional. Filter for series. | 
 **isKids** | **optional.Bool**| Optional. Filter for kids. | 
 **isSports** | **optional.Bool**| Optional. Filter for sports. | 
 **isNews** | **optional.Bool**| Optional. Filter for news. | 
 **isLibraryItem** | **optional.Bool**| Optional. Filter for is library item. | 
 **enableTotalRecordCount** | **optional.Bool**| Optional. Return total record count. | [default to true]

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


## GetRecordingsSeries

> BaseItemDtoQueryResult GetRecordingsSeries(ctx, optional)

Gets live tv recording series.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetRecordingsSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRecordingsSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**| Optional. Filter by channel id. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user and attach user data. | 
 **groupId** | **optional.String**| Optional. Filter by recording group. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **status** | [**optional.Interface of RecordingStatus**](.md)| Optional. Filter by recording status. | 
 **isInProgress** | **optional.Bool**| Optional. Filter by recordings that are in progress, or not. | 
 **seriesTimerId** | **optional.String**| Optional. Filter by recordings belonging to a series timer. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **enableTotalRecordCount** | **optional.Bool**| Optional. Return total record count. | [default to true]

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


## GetSchedulesDirectCountries

> *os.File GetSchedulesDirectCountries(ctx, )

Gets available countries.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesTimer

> SeriesTimerInfoDto GetSeriesTimer(ctx, timerId)

Gets a live tv series timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 

### Return type

[**SeriesTimerInfoDto**](SeriesTimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesTimers

> SeriesTimerInfoDtoQueryResult GetSeriesTimers(ctx, optional)

Gets live tv series timers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSeriesTimersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSeriesTimersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **optional.String**| Optional. Sort by SortName or Priority. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| Optional. Sort in Ascending or Descending order. | 

### Return type

[**SeriesTimerInfoDtoQueryResult**](SeriesTimerInfoDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimer

> TimerInfoDto GetTimer(ctx, timerId)

Gets a timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 

### Return type

[**TimerInfoDto**](TimerInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimers

> TimerInfoDtoQueryResult GetTimers(ctx, optional)

Gets the live tv timers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTimersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTimersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**| Optional. Filter by channel id. | 
 **seriesTimerId** | **optional.String**| Optional. Filter by timers belonging to a series timer. | 
 **isActive** | **optional.Bool**| Optional. Filter by timers that are active. | 
 **isScheduled** | **optional.Bool**| Optional. Filter by timers that are scheduled. | 

### Return type

[**TimerInfoDtoQueryResult**](TimerInfoDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTunerHostTypes

> []NameIdPair GetTunerHostTypes(ctx, )

Get tuner host types.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetTuner

> ResetTuner(ctx, tunerId)

Resets a tv tuner.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tunerId** | **string**| Tuner id. | 

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


## SetChannelMapping

> TunerChannelMapping SetChannelMapping(ctx, optional)

Set channel mappings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetChannelMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetChannelMappingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **optional.String**| Provider id. | 
 **tunerChannelId** | **optional.String**| Tuner channel id. | 
 **providerChannelId** | **optional.String**| Provider channel id. | 

### Return type

[**TunerChannelMapping**](TunerChannelMapping.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSeriesTimer

> UpdateSeriesTimer(ctx, timerId, optional)

Updates a live tv series timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 
 **optional** | ***UpdateSeriesTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSeriesTimerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesTimerInfoDto** | [**optional.Interface of SeriesTimerInfoDto**](SeriesTimerInfoDto.md)| New series timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTimer

> UpdateTimer(ctx, timerId, optional)

Updates a live tv timer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timerId** | **string**| Timer id. | 
 **optional** | ***UpdateTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTimerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timerInfoDto** | [**optional.Interface of TimerInfoDto**](TimerInfoDto.md)| New timer info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

