# \RemoteImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadRemoteImage**](RemoteImageApi.md#DownloadRemoteImage) | **Post** /Items/{itemId}/RemoteImages/Download | Downloads a remote image for an item.
[**GetRemoteImage**](RemoteImageApi.md#GetRemoteImage) | **Get** /Images/Remote | Gets a remote image.
[**GetRemoteImageProviders**](RemoteImageApi.md#GetRemoteImageProviders) | **Get** /Items/{itemId}/RemoteImages/Providers | Gets available remote image providers for an item.
[**GetRemoteImages**](RemoteImageApi.md#GetRemoteImages) | **Get** /Items/{itemId}/RemoteImages | Gets available remote images for an item.



## DownloadRemoteImage

> DownloadRemoteImage(ctx, itemId, type_, optional)

Downloads a remote image for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item Id. | 
**type_** | [**ImageType**](.md)| The image type. | 
 **optional** | ***DownloadRemoteImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadRemoteImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **imageUrl** | **optional.String**| The image url. | 

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


## GetRemoteImage

> *os.File GetRemoteImage(ctx, imageUrl)

Gets a remote image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageUrl** | **string**| The image url. | 

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


## GetRemoteImageProviders

> []ImageProviderInfo GetRemoteImageProviders(ctx, itemId)

Gets available remote image providers for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item Id. | 

### Return type

[**[]ImageProviderInfo**](ImageProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteImages

> RemoteImageResult GetRemoteImages(ctx, itemId, optional)

Gets available remote images for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item Id. | 
 **optional** | ***GetRemoteImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRemoteImagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of ImageType**](.md)| The image type. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **providerName** | **optional.String**| Optional. The image provider to use. | 
 **includeAllLanguages** | **optional.Bool**| Optional. Include all languages. | [default to false]

### Return type

[**RemoteImageResult**](RemoteImageResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

