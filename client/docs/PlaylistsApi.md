# \PlaylistsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPlaylist**](PlaylistsApi.md#AddToPlaylist) | **Post** /Playlists/{playlistId}/Items | Adds items to a playlist.
[**CreatePlaylist**](PlaylistsApi.md#CreatePlaylist) | **Post** /Playlists | Creates a new playlist.
[**GetPlaylistItems**](PlaylistsApi.md#GetPlaylistItems) | **Get** /Playlists/{playlistId}/Items | Gets the original items of a playlist.
[**MoveItem**](PlaylistsApi.md#MoveItem) | **Post** /Playlists/{playlistId}/Items/{itemId}/Move/{newIndex} | Moves a playlist item.
[**RemoveFromPlaylist**](PlaylistsApi.md#RemoveFromPlaylist) | **Delete** /Playlists/{playlistId}/Items | Removes items from a playlist.



## AddToPlaylist

> AddToPlaylist(ctx, playlistId, optional)

Adds items to a playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | [**string**](.md)| The playlist id. | 
 **optional** | ***AddToPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddToPlaylistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **optional.String**| Item id, comma delimited. | 
 **userId** | [**optional.Interface of string**](.md)| The userId. | 

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


## CreatePlaylist

> PlaylistCreationResult CreatePlaylist(ctx, createPlaylistDto)

Creates a new playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPlaylistDto** | [**CreatePlaylistDto**](CreatePlaylistDto.md)| The create playlist payload. | 

### Return type

[**PlaylistCreationResult**](PlaylistCreationResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistItems

> BaseItemDtoQueryResult GetPlaylistItems(ctx, playlistId, userId, optional)

Gets the original items of a playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | [**string**](.md)| The playlist id. | 
**userId** | [**string**](.md)| User id. | 
 **optional** | ***GetPlaylistItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPlaylistItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 

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


## MoveItem

> MoveItem(ctx, playlistId, itemId, newIndex)

Moves a playlist item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string**| The playlist id. | 
**itemId** | **string**| The item id. | 
**newIndex** | **int32**| The new index. | 

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


## RemoveFromPlaylist

> RemoveFromPlaylist(ctx, playlistId, optional)

Removes items from a playlist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string**| The playlist id. | 
 **optional** | ***RemoveFromPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveFromPlaylistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | **optional.String**| The item ids, comma delimited. | 

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

