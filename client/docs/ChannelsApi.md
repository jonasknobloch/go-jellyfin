# \ChannelsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllChannelFeatures**](ChannelsApi.md#GetAllChannelFeatures) | **Get** /Channels/Features | Get all channel features.
[**GetChannelFeatures**](ChannelsApi.md#GetChannelFeatures) | **Get** /Channels/{channelId}/Features | Get channel features.
[**GetChannelItems**](ChannelsApi.md#GetChannelItems) | **Get** /Channels/{channelId}/Items | Get channel items.
[**GetChannels**](ChannelsApi.md#GetChannels) | **Get** /Channels | Gets available channels.
[**GetLatestChannelItems**](ChannelsApi.md#GetLatestChannelItems) | **Get** /Channels/Items/Latest | Gets latest channel items.



## GetAllChannelFeatures

> []ChannelFeatures GetAllChannelFeatures(ctx, )

Get all channel features.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ChannelFeatures**](ChannelFeatures.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelFeatures

> ChannelFeatures GetChannelFeatures(ctx, channelId)

Get channel features.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| Channel id. | 

### Return type

[**ChannelFeatures**](ChannelFeatures.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelItems

> BaseItemDtoQueryResult GetChannelItems(ctx, channelId, optional)

Get channel items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | [**string**](.md)| Channel Id. | 
 **optional** | ***GetChannelItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderId** | [**optional.Interface of string**](.md)| Optional. Folder Id. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. User Id. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **sortOrder** | **optional.String**| Optional. Sort Order - Ascending,Descending. | 
 **filters** | [**optional.Interface of []ItemFilter**](ItemFilter.md)| Optional. Specify additional filters to apply. | 
 **sortBy** | **optional.String**| Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 

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


## GetChannels

> BaseItemDtoQueryResult GetChannels(ctx, optional)

Gets available channels.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| User Id to filter by. Use System.Guid.Empty to not filter by user. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **supportsLatestItems** | **optional.Bool**| Optional. Filter by channels that support getting latest items. | 
 **supportsMediaDeletion** | **optional.Bool**| Optional. Filter by channels that support media deletion. | 
 **isFavorite** | **optional.Bool**| Optional. Filter by channels that are favorite. | 

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


## GetLatestChannelItems

> BaseItemDtoQueryResult GetLatestChannelItems(ctx, optional)

Gets latest channel items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestChannelItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestChannelItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. User Id. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **filters** | [**optional.Interface of []ItemFilter**](ItemFilter.md)| Optional. Specify additional filters to apply. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **channelIds** | **optional.String**| Optional. Specify one or more channel id&#39;s, comma delimited. | 

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

