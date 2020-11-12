# \ImageByNameApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeneralImage**](ImageByNameApi.md#GetGeneralImage) | **Get** /Images/General/{name}/{type} | Get General Image.
[**GetGeneralImages**](ImageByNameApi.md#GetGeneralImages) | **Get** /Images/General | Get all general images.
[**GetMediaInfoImage**](ImageByNameApi.md#GetMediaInfoImage) | **Get** /Images/MediaInfo/{theme}/{name} | Get media info image.
[**GetMediaInfoImages**](ImageByNameApi.md#GetMediaInfoImages) | **Get** /Images/MediaInfo | Get all media info images.
[**GetRatingImage**](ImageByNameApi.md#GetRatingImage) | **Get** /Images/Ratings/{theme}/{name} | Get rating image.
[**GetRatingImages**](ImageByNameApi.md#GetRatingImages) | **Get** /Images/Ratings | Get all general images.



## GetGeneralImage

> *os.File GetGeneralImage(ctx, name, type_)

Get General Image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the image. | 
**type_** | **string**| Image Type (primary, backdrop, logo, etc). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralImages

> []ImageByNameInfo GetGeneralImages(ctx, )

Get all general images.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaInfoImage

> *os.File GetMediaInfoImage(ctx, theme, name)

Get media info image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**theme** | **string**| The theme to get the image from. | 
**name** | **string**| The name of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaInfoImages

> []ImageByNameInfo GetMediaInfoImages(ctx, )

Get all media info images.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRatingImage

> *os.File GetRatingImage(ctx, theme, name)

Get rating image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**theme** | **string**| The theme to get the image from. | 
**name** | **string**| The name of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRatingImages

> []ImageByNameInfo GetRatingImages(ctx, )

Get all general images.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

