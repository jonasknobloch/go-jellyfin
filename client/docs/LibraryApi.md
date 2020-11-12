# \LibraryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItem**](LibraryApi.md#DeleteItem) | **Delete** /Items/{itemId} | Deletes an item from the library and filesystem.
[**DeleteItems**](LibraryApi.md#DeleteItems) | **Delete** /Items | Deletes items from the library and filesystem.
[**GetAncestors**](LibraryApi.md#GetAncestors) | **Get** /Items/{itemId}/Ancestors | Gets all parents of an item.
[**GetCriticReviews**](LibraryApi.md#GetCriticReviews) | **Get** /Items/{itemId}/CriticReviews | Gets critic review for an item.
[**GetDownload**](LibraryApi.md#GetDownload) | **Get** /Items/{itemId}/Download | Downloads item media.
[**GetFile**](LibraryApi.md#GetFile) | **Get** /Items/{itemId}/File | Get the original file of an item.
[**GetItemCounts**](LibraryApi.md#GetItemCounts) | **Get** /Items/Counts | Get item counts.
[**GetLibraryOptionsInfo**](LibraryApi.md#GetLibraryOptionsInfo) | **Get** /Libraries/AvailableOptions | Gets the library options info.
[**GetMediaFolders**](LibraryApi.md#GetMediaFolders) | **Get** /Library/MediaFolders | Gets all user media folders.
[**GetPhysicalPaths**](LibraryApi.md#GetPhysicalPaths) | **Get** /Library/PhysicalPaths | Gets a list of physical paths from virtual folders.
[**GetSimilarAlbums2**](LibraryApi.md#GetSimilarAlbums2) | **Get** /Albums/{itemId}/Similar | Gets similar items.
[**GetSimilarArtists2**](LibraryApi.md#GetSimilarArtists2) | **Get** /Artists/{itemId}/Similar | Gets similar items.
[**GetSimilarItems**](LibraryApi.md#GetSimilarItems) | **Get** /Items/{itemId}/Similar | Gets similar items.
[**GetSimilarMovies2**](LibraryApi.md#GetSimilarMovies2) | **Get** /Movies/{itemId}/Similar | Gets similar items.
[**GetSimilarShows2**](LibraryApi.md#GetSimilarShows2) | **Get** /Shows/{itemId}/Similar | Gets similar items.
[**GetSimilarTrailers2**](LibraryApi.md#GetSimilarTrailers2) | **Get** /Trailers/{itemId}/Similar | Gets similar items.
[**GetThemeMedia**](LibraryApi.md#GetThemeMedia) | **Get** /Items/{itemId}/ThemeMedia | Get theme songs and videos for an item.
[**GetThemeSongs**](LibraryApi.md#GetThemeSongs) | **Get** /Items/{itemId}/ThemeSongs | Get theme songs for an item.
[**GetThemeVideos**](LibraryApi.md#GetThemeVideos) | **Get** /Items/{itemId}/ThemeVideos | Get theme videos for an item.
[**PostAddedMovies**](LibraryApi.md#PostAddedMovies) | **Post** /Library/Movies/Added | Reports that new movies have been added by an external source.
[**PostAddedSeries**](LibraryApi.md#PostAddedSeries) | **Post** /Library/Series/Added | Reports that new episodes of a series have been added by an external source.
[**PostUpdatedMedia**](LibraryApi.md#PostUpdatedMedia) | **Post** /Library/Media/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedMovies**](LibraryApi.md#PostUpdatedMovies) | **Post** /Library/Movies/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedSeries**](LibraryApi.md#PostUpdatedSeries) | **Post** /Library/Series/Updated | Reports that new episodes of a series have been added by an external source.
[**RefreshLibrary**](LibraryApi.md#RefreshLibrary) | **Get** /Library/Refresh | Starts a library scan.



## DeleteItem

> DeleteItem(ctx, itemId)

Deletes an item from the library and filesystem.

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


## DeleteItems

> DeleteItems(ctx, optional)

Deletes items from the library and filesystem.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **optional.String**| The item ids. | 

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


## GetAncestors

> []BaseItemDto GetAncestors(ctx, itemId, optional)

Gets all parents of an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetAncestorsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAncestorsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCriticReviews

> BaseItemDtoQueryResult GetCriticReviews(ctx, itemId)

Gets critic review for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**|  | 

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


## GetDownload

> *os.File GetDownload(ctx, itemId)

Downloads item media.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/_*, audio/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> *os.File GetFile(ctx, itemId)

Get the original file of an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/_*, audio/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemCounts

> ItemCounts GetItemCounts(ctx, optional)

Get item counts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetItemCountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemCountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. Get counts from a specific user&#39;s library. | 
 **isFavorite** | **optional.Bool**| Optional. Get counts of favorite items. | 

### Return type

[**ItemCounts**](ItemCounts.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryOptionsInfo

> LibraryOptionsResultDto GetLibraryOptionsInfo(ctx, optional)

Gets the library options info.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLibraryOptionsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLibraryOptionsInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryContentType** | **optional.String**| Library content type. | 
 **isNewLibrary** | **optional.Bool**| Whether this is a new library. | 

### Return type

[**LibraryOptionsResultDto**](LibraryOptionsResultDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaFolders

> BaseItemDtoQueryResult GetMediaFolders(ctx, optional)

Gets all user media folders.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMediaFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMediaFoldersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **optional.Bool**| Optional. Filter by folders that are marked hidden, or not. | 

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


## GetPhysicalPaths

> []string GetPhysicalPaths(ctx, )

Gets a list of physical paths from virtual folders.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarAlbums2

> BaseItemDtoQueryResult GetSimilarAlbums2(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarAlbums2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarAlbums2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarArtists2

> BaseItemDtoQueryResult GetSimilarArtists2(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarArtists2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarArtists2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarItems

> BaseItemDtoQueryResult GetSimilarItems(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarMovies2

> BaseItemDtoQueryResult GetSimilarMovies2(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarMovies2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarMovies2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarShows2

> BaseItemDtoQueryResult GetSimilarShows2(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarShows2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarShows2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarTrailers2

> BaseItemDtoQueryResult GetSimilarTrailers2(ctx, itemId, optional)

Gets similar items.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetSimilarTrailers2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarTrailers2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **optional.String**| Exclude artist ids. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetThemeMedia

> AllThemeMediaResult GetThemeMedia(ctx, itemId, optional)

Get theme songs and videos for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetThemeMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetThemeMediaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **optional.Bool**| Optional. Determines whether or not parent items should be searched for theme media. | [default to false]

### Return type

[**AllThemeMediaResult**](AllThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThemeSongs

> ThemeMediaResult GetThemeSongs(ctx, itemId, optional)

Get theme songs for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetThemeSongsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetThemeSongsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **optional.Bool**| Optional. Determines whether or not parent items should be searched for theme media. | [default to false]

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThemeVideos

> ThemeMediaResult GetThemeVideos(ctx, itemId, optional)

Get theme videos for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***GetThemeVideosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetThemeVideosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **optional.Bool**| Optional. Determines whether or not parent items should be searched for theme media. | [default to false]

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddedMovies

> PostAddedMovies(ctx, optional)

Reports that new movies have been added by an external source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAddedMoviesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAddedMoviesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **optional.String**| The tmdbId. | 
 **imdbId** | **optional.String**| The imdbId. | 

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


## PostAddedSeries

> PostAddedSeries(ctx, optional)

Reports that new episodes of a series have been added by an external source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostAddedSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostAddedSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **optional.String**| The tvdbId. | 

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


## PostUpdatedMedia

> PostUpdatedMedia(ctx, mediaUpdateInfoDto)

Reports that new movies have been added by an external source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaUpdateInfoDto** | [**[]MediaUpdateInfoDto**](MediaUpdateInfoDto.md)| A list of updated media paths. | 

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


## PostUpdatedMovies

> PostUpdatedMovies(ctx, optional)

Reports that new movies have been added by an external source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostUpdatedMoviesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUpdatedMoviesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **optional.String**| The tmdbId. | 
 **imdbId** | **optional.String**| The imdbId. | 

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


## PostUpdatedSeries

> PostUpdatedSeries(ctx, optional)

Reports that new episodes of a series have been added by an external source.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostUpdatedSeriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUpdatedSeriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **optional.String**| The tvdbId. | 

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


## RefreshLibrary

> RefreshLibrary(ctx, )

Starts a library scan.

### Required Parameters

This endpoint does not need any parameter.

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

