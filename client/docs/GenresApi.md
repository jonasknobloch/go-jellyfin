# \GenresApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenre**](GenresApi.md#GetGenre) | **Get** /Genres/{genreName} | Gets a genre, by name.
[**GetGenres**](GenresApi.md#GetGenres) | **Get** /Genres | Gets all genres from a given item, folder, or the entire library.



## GetGenre

> BaseItemDto GetGenre(ctx, genreName, optional)

Gets a genre, by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreName** | **string**| The genre name. | 
 **optional** | ***GetGenreOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGenreOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| The user id. | 

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


## GetGenres

> BaseItemDtoQueryResult GetGenres(ctx, optional)

Gets all genres from a given item, folder, or the entire library.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGenresOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGenresOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minCommunityRating** | **optional.Float64**| Optional filter by minimum community rating. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **searchTerm** | **optional.String**| The search term. | 
 **parentId** | **optional.String**| Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **excludeItemTypes** | **optional.String**| Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | **optional.String**| Optional. If specified, results will be filtered in based on item type. This allows multiple, comma delimited. | 
 **filters** | [**optional.Interface of []ItemFilter**](ItemFilter.md)| Optional. Specify additional filters to apply. | 
 **isFavorite** | **optional.Bool**| Optional filter by items that are marked as favorite, or not. | 
 **mediaTypes** | **optional.String**| Optional filter by MediaType. Allows multiple, comma delimited. | 
 **genres** | **optional.String**| Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. | 
 **genreIds** | **optional.String**| Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. | 
 **officialRatings** | **optional.String**| Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. | 
 **tags** | **optional.String**| Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. | 
 **years** | **optional.String**| Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. | 
 **enableUserData** | **optional.Bool**| Optional, include user data. | 
 **imageTypeLimit** | **optional.Int32**| Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **person** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified person id. | 
 **personTypes** | **optional.String**| Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. | 
 **studios** | **optional.String**| Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. | 
 **studioIds** | **optional.String**| Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. | 
 **userId** | [**optional.Interface of string**](.md)| User id. | 
 **nameStartsWithOrGreater** | **optional.String**| Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **optional.String**| Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **optional.String**| Optional filter by items whose name is equally or lesser than a given input string. | 
 **enableImages** | **optional.Bool**| Optional, include image information in output. | [default to true]
 **enableTotalRecordCount** | **optional.Bool**| Optional. Include total record count. | [default to true]

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

