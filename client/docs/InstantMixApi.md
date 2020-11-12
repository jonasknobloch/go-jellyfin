# \InstantMixApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstantMixFromAlbum**](InstantMixApi.md#GetInstantMixFromAlbum) | **Get** /Albums/{id}/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromArtists**](InstantMixApi.md#GetInstantMixFromArtists) | **Get** /Artists/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromItem**](InstantMixApi.md#GetInstantMixFromItem) | **Get** /Items/{id}/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromMusicGenre**](InstantMixApi.md#GetInstantMixFromMusicGenre) | **Get** /MusicGenres/{name}/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromMusicGenres**](InstantMixApi.md#GetInstantMixFromMusicGenres) | **Get** /MusicGenres/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromPlaylist**](InstantMixApi.md#GetInstantMixFromPlaylist) | **Get** /Playlists/{id}/InstantMix | Creates an instant playlist based on a given song.
[**GetInstantMixFromSong**](InstantMixApi.md#GetInstantMixFromSong) | **Get** /Songs/{id}/InstantMix | Creates an instant playlist based on a given song.



## GetInstantMixFromAlbum

> BaseItemDtoQueryResult GetInstantMixFromAlbum(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromAlbumOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromAlbumOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromArtists

> BaseItemDtoQueryResult GetInstantMixFromArtists(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromArtistsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromArtistsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromItem

> BaseItemDtoQueryResult GetInstantMixFromItem(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromMusicGenre

> BaseItemDtoQueryResult GetInstantMixFromMusicGenre(ctx, name, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The genre name. | 
 **optional** | ***GetInstantMixFromMusicGenreOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromMusicGenreOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromMusicGenres

> BaseItemDtoQueryResult GetInstantMixFromMusicGenres(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromMusicGenresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromMusicGenresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromPlaylist

> BaseItemDtoQueryResult GetInstantMixFromPlaylist(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromPlaylistOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromPlaylistOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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


## GetInstantMixFromSong

> BaseItemDtoQueryResult GetInstantMixFromSong(ctx, id, optional)

Creates an instant playlist based on a given song.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| The item id. | 
 **optional** | ***GetInstantMixFromSongOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInstantMixFromSongOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
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

