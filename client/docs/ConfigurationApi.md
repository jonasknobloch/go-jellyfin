# \ConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationApi.md#GetConfiguration) | **Get** /System/Configuration | Gets application configuration.
[**GetDefaultMetadataOptions**](ConfigurationApi.md#GetDefaultMetadataOptions) | **Get** /System/Configuration/MetadataOptions/Default | Gets a default MetadataOptions object.
[**GetNamedConfiguration**](ConfigurationApi.md#GetNamedConfiguration) | **Get** /System/Configuration/{key} | Gets a named configuration.
[**UpdateConfiguration**](ConfigurationApi.md#UpdateConfiguration) | **Post** /System/Configuration | Updates application configuration.
[**UpdateMediaEncoderPath**](ConfigurationApi.md#UpdateMediaEncoderPath) | **Post** /System/MediaEncoder/Path | Updates the path to the media encoder.
[**UpdateNamedConfiguration**](ConfigurationApi.md#UpdateNamedConfiguration) | **Post** /System/Configuration/{key} | Updates named configuration.



## GetConfiguration

> ServerConfiguration GetConfiguration(ctx, )

Gets application configuration.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServerConfiguration**](ServerConfiguration.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultMetadataOptions

> MetadataOptions GetDefaultMetadataOptions(ctx, )

Gets a default MetadataOptions object.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**MetadataOptions**](MetadataOptions.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedConfiguration

> *os.File GetNamedConfiguration(ctx, key)

Gets a named configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Configuration key. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> UpdateConfiguration(ctx, serverConfiguration)

Updates application configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverConfiguration** | [**ServerConfiguration**](ServerConfiguration.md)| Configuration. | 

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


## UpdateMediaEncoderPath

> UpdateMediaEncoderPath(ctx, mediaEncoderPathDto)

Updates the path to the media encoder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaEncoderPathDto** | [**MediaEncoderPathDto**](MediaEncoderPathDto.md)| Media encoder path form body. | 

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


## UpdateNamedConfiguration

> UpdateNamedConfiguration(ctx, key)

Updates named configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string**| Configuration key. | 

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

