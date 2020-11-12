# \AlbumsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSimilarAlbums**](AlbumsApi.md#GetSimilarAlbums) | **Get** /Albums/{albumId}/Similar | Finds albums similar to a given album.
[**GetSimilarArtists**](AlbumsApi.md#GetSimilarArtists) | **Get** /Artists/{artistId}/Similar | Finds artists similar to a given artist.



## GetSimilarAlbums

> BaseItemDtoQueryResult GetSimilarAlbums(ctx, albumId, optional)

Finds albums similar to a given album.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**albumId** | **string**| The album id. | 
 **optional** | ***GetSimilarAlbumsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarAlbumsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **excludeArtistIds** | **optional.String**| Optional. Ids of artists to exclude. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 

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


## GetSimilarArtists

> BaseItemDtoQueryResult GetSimilarArtists(ctx, artistId, optional)

Finds artists similar to a given artist.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artistId** | **string**| The artist id. | 
 **optional** | ***GetSimilarArtistsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSimilarArtistsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **excludeArtistIds** | **optional.String**| Optional. Ids of artists to exclude. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 

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

