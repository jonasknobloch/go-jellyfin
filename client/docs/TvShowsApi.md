# \TvShowsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEpisodes**](TvShowsApi.md#GetEpisodes) | **Get** /Shows/{seriesId}/Episodes | Gets episodes for a tv season.
[**GetNextUp**](TvShowsApi.md#GetNextUp) | **Get** /Shows/NextUp | Gets a list of next up episodes.
[**GetSeasons**](TvShowsApi.md#GetSeasons) | **Get** /Shows/{seriesId}/Seasons | Gets seasons for a tv series.
[**GetUpcomingEpisodes**](TvShowsApi.md#GetUpcomingEpisodes) | **Get** /Shows/Upcoming | Gets a list of upcoming episodes.



## GetEpisodes

> BaseItemDtoQueryResult GetEpisodes(ctx, seriesId, optional)

Gets episodes for a tv season.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **string**| The series id. | 
 **optional** | ***GetEpisodesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEpisodesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| The user id. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **season** | **optional.Int32**| Optional filter by season number. | 
 **seasonId** | **optional.String**| Optional. Filter by season id. | 
 **isMissing** | **optional.Bool**| Optional. Filter by items that are missing episodes or not. | 
 **adjacentTo** | **optional.String**| Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **optional.String**| Optional. Skip through the list until a given item is found. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **enableImages** | **optional.Bool**| Optional, include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **sortBy** | **optional.String**| Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 

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


## GetNextUp

> BaseItemDtoQueryResult GetNextUp(ctx, optional)

Gets a list of next up episodes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetNextUpOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNextUpOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| The user id of the user to get the next up episodes for. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **seriesId** | **optional.String**| Optional. Filter by series id. | 
 **parentId** | **optional.String**| Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **enableImges** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **enableTotalRecordCount** | **optional.Bool**| Whether to enable the total records count. Defaults to true. | [default to true]

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


## GetSeasons

> BaseItemDtoQueryResult GetSeasons(ctx, seriesId, optional)

Gets seasons for a tv series.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **string**| The series id. | 
 **optional** | ***GetSeasonsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSeasonsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| The user id. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **isSpecialSeason** | **optional.Bool**| Optional. Filter by special season. | 
 **isMissing** | **optional.Bool**| Optional. Filter by items that are missing episodes or not. | 
 **adjacentTo** | **optional.String**| Optional. Return items that are siblings of a supplied item. | 
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 

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


## GetUpcomingEpisodes

> BaseItemDtoQueryResult GetUpcomingEpisodes(ctx, optional)

Gets a list of upcoming episodes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUpcomingEpisodesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUpcomingEpisodesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| The user id of the user to get the upcoming episodes for. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **parentId** | **optional.String**| Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **enableImges** | **optional.Bool**| Optional. Include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 

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

