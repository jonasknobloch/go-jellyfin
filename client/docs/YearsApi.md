# \YearsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetYear**](YearsApi.md#GetYear) | **Get** /Years/{year} | Gets a year.
[**GetYears**](YearsApi.md#GetYears) | **Get** /Years | Get years.



## GetYear

> BaseItemDto GetYear(ctx, year, optional)

Gets a year.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32**| The year. | 
 **optional** | ***GetYearOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetYearOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 

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


## GetYears

> BaseItemDtoQueryResult GetYears(ctx, optional)

Get years.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetYearsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetYearsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **optional.Int32**| Skips over a given number of items within the results. Use for paging. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **sortOrder** | **optional.String**| Sort Order - Ascending,Descending. | 
 **parentId** | **optional.String**| Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **excludeItemTypes** | **optional.String**| Optional. If specified, results will be excluded based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | **optional.String**| Optional. If specified, results will be included based on item type. This allows multiple, comma delimited. | 
 **mediaTypes** | **optional.String**| Optional. Filter by MediaType. Allows multiple, comma delimited. | 
 **sortBy** | **optional.String**| Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **enableUserData** | **optional.Bool**| Optional. Include user data. | 
 **imageTypeLimit** | **optional.Int32**| Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **userId** | [**optional.Interface of string**](.md)| User Id. | 
 **recursive** | **optional.Bool**| Search recursively. | [default to true]
 **enableImages** | **optional.Bool**| Optional. Include image information in output. | [default to true]

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

