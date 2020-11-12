# \EnvironmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultDirectoryBrowser**](EnvironmentApi.md#GetDefaultDirectoryBrowser) | **Get** /Environment/DefaultDirectoryBrowser | Get Default directory browser.
[**GetDirectoryContents**](EnvironmentApi.md#GetDirectoryContents) | **Get** /Environment/DirectoryContents | Gets the contents of a given directory in the file system.
[**GetDrives**](EnvironmentApi.md#GetDrives) | **Get** /Environment/Drives | Gets available drives from the server&#39;s file system.
[**GetNetworkShares**](EnvironmentApi.md#GetNetworkShares) | **Get** /Environment/NetworkShares | Gets network paths.
[**GetParentPath**](EnvironmentApi.md#GetParentPath) | **Get** /Environment/ParentPath | Gets the parent path of a given path.
[**ValidatePath**](EnvironmentApi.md#ValidatePath) | **Post** /Environment/ValidatePath | Validates path.



## GetDefaultDirectoryBrowser

> DefaultDirectoryBrowserInfoDto GetDefaultDirectoryBrowser(ctx, )

Get Default directory browser.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DefaultDirectoryBrowserInfoDto**](DefaultDirectoryBrowserInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectoryContents

> []FileSystemEntryInfo GetDirectoryContents(ctx, path, optional)

Gets the contents of a given directory in the file system.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path. | 
 **optional** | ***GetDirectoryContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDirectoryContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeFiles** | **optional.Bool**| An optional filter to include or exclude files from the results. true/false. | [default to false]
 **includeDirectories** | **optional.Bool**| An optional filter to include or exclude folders from the results. true/false. | [default to false]

### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrives

> []FileSystemEntryInfo GetDrives(ctx, )

Gets available drives from the server's file system.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkShares

> []FileSystemEntryInfo GetNetworkShares(ctx, )

Gets network paths.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentPath

> string GetParentPath(ctx, path)

Gets the parent path of a given path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string**| The path. | 

### Return type

**string**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePath

> ValidatePath(ctx, validatePathDto)

Validates path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**validatePathDto** | [**ValidatePathDto**](ValidatePathDto.md)| Validate request object. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

