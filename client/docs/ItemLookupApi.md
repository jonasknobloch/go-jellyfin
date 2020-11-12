# \ItemLookupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplySearchCriteria**](ItemLookupApi.md#ApplySearchCriteria) | **Post** /Items/RemoteSearch/Apply/{itemId} | Applies search criteria to an item and refreshes metadata.
[**GetBookRemoteSearchResults**](ItemLookupApi.md#GetBookRemoteSearchResults) | **Post** /Items/RemoteSearch/Book | Get book remote search.
[**GetBoxSetRemoteSearchResults**](ItemLookupApi.md#GetBoxSetRemoteSearchResults) | **Post** /Items/RemoteSearch/BoxSet | Get box set remote search.
[**GetExternalIdInfos**](ItemLookupApi.md#GetExternalIdInfos) | **Get** /Items/{itemId}/ExternalIdInfos | Get the item&#39;s external id info.
[**GetMovieRemoteSearchResults**](ItemLookupApi.md#GetMovieRemoteSearchResults) | **Post** /Items/RemoteSearch/Movie | Get movie remote search.
[**GetMusicAlbumRemoteSearchResults**](ItemLookupApi.md#GetMusicAlbumRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicAlbum | Get music album remote search.
[**GetMusicArtistRemoteSearchResults**](ItemLookupApi.md#GetMusicArtistRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicArtist | Get music artist remote search.
[**GetMusicVideoRemoteSearchResults**](ItemLookupApi.md#GetMusicVideoRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicVideo | Get music video remote search.
[**GetPersonRemoteSearchResults**](ItemLookupApi.md#GetPersonRemoteSearchResults) | **Post** /Items/RemoteSearch/Person | Get person remote search.
[**GetRemoteSearchImage**](ItemLookupApi.md#GetRemoteSearchImage) | **Get** /Items/RemoteSearch/Image | Gets a remote image.
[**GetSeriesRemoteSearchResults**](ItemLookupApi.md#GetSeriesRemoteSearchResults) | **Post** /Items/RemoteSearch/Series | Get series remote search.
[**GetTrailerRemoteSearchResults**](ItemLookupApi.md#GetTrailerRemoteSearchResults) | **Post** /Items/RemoteSearch/Trailer | Get trailer remote search.



## ApplySearchCriteria

> ApplySearchCriteria(ctx, itemId, remoteSearchResult, optional)

Applies search criteria to an item and refreshes metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**remoteSearchResult** | [**RemoteSearchResult**](RemoteSearchResult.md)| The remote search result. | 
 **optional** | ***ApplySearchCriteriaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplySearchCriteriaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceAllImages** | **optional.Bool**| Optional. Whether or not to replace all images. Default: True. | [default to true]

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


## GetBookRemoteSearchResults

> []RemoteSearchResult GetBookRemoteSearchResults(ctx, bookInfoRemoteSearchQuery)

Get book remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookInfoRemoteSearchQuery** | [**BookInfoRemoteSearchQuery**](BookInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxSetRemoteSearchResults

> []RemoteSearchResult GetBoxSetRemoteSearchResults(ctx, boxSetInfoRemoteSearchQuery)

Get box set remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxSetInfoRemoteSearchQuery** | [**BoxSetInfoRemoteSearchQuery**](BoxSetInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalIdInfos

> []ExternalIdInfo GetExternalIdInfos(ctx, itemId)

Get the item's external id info.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**[]ExternalIdInfo**](ExternalIdInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieRemoteSearchResults

> []RemoteSearchResult GetMovieRemoteSearchResults(ctx, movieInfoRemoteSearchQuery)

Get movie remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieInfoRemoteSearchQuery** | [**MovieInfoRemoteSearchQuery**](MovieInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicAlbumRemoteSearchResults

> []RemoteSearchResult GetMusicAlbumRemoteSearchResults(ctx, albumInfoRemoteSearchQuery)

Get music album remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumInfoRemoteSearchQuery** | [**AlbumInfoRemoteSearchQuery**](AlbumInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicArtistRemoteSearchResults

> []RemoteSearchResult GetMusicArtistRemoteSearchResults(ctx, artistInfoRemoteSearchQuery)

Get music artist remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artistInfoRemoteSearchQuery** | [**ArtistInfoRemoteSearchQuery**](ArtistInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicVideoRemoteSearchResults

> []RemoteSearchResult GetMusicVideoRemoteSearchResults(ctx, musicVideoInfoRemoteSearchQuery)

Get music video remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**musicVideoInfoRemoteSearchQuery** | [**MusicVideoInfoRemoteSearchQuery**](MusicVideoInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonRemoteSearchResults

> []RemoteSearchResult GetPersonRemoteSearchResults(ctx, personLookupInfoRemoteSearchQuery)

Get person remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personLookupInfoRemoteSearchQuery** | [**PersonLookupInfoRemoteSearchQuery**](PersonLookupInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteSearchImage

> *os.File GetRemoteSearchImage(ctx, imageUrl, providerName)

Gets a remote image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageUrl** | **string**| The image url. | 
**providerName** | **string**| The provider name. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesRemoteSearchResults

> []RemoteSearchResult GetSeriesRemoteSearchResults(ctx, seriesInfoRemoteSearchQuery)

Get series remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesInfoRemoteSearchQuery** | [**SeriesInfoRemoteSearchQuery**](SeriesInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrailerRemoteSearchResults

> []RemoteSearchResult GetTrailerRemoteSearchResults(ctx, trailerInfoRemoteSearchQuery)

Get trailer remote search.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trailerInfoRemoteSearchQuery** | [**TrailerInfoRemoteSearchQuery**](TrailerInfoRemoteSearchQuery.md)| Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

