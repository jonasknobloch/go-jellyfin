# \LibraryStructureApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaPath**](LibraryStructureApi.md#AddMediaPath) | **Post** /Library/VirtualFolders/Paths | Add a media path to a library.
[**AddVirtualFolder**](LibraryStructureApi.md#AddVirtualFolder) | **Post** /Library/VirtualFolders | Adds a virtual folder.
[**GetVirtualFolders**](LibraryStructureApi.md#GetVirtualFolders) | **Get** /Library/VirtualFolders | Gets all virtual folders.
[**RemoveMediaPath**](LibraryStructureApi.md#RemoveMediaPath) | **Delete** /Library/VirtualFolders/Paths | Remove a media path.
[**RemoveVirtualFolder**](LibraryStructureApi.md#RemoveVirtualFolder) | **Delete** /Library/VirtualFolders | Removes a virtual folder.
[**RenameVirtualFolder**](LibraryStructureApi.md#RenameVirtualFolder) | **Post** /Library/VirtualFolders/Name | Renames a virtual folder.
[**UpdateLibraryOptions**](LibraryStructureApi.md#UpdateLibraryOptions) | **Post** /Library/VirtualFolders/LibraryOptions | Update library options.
[**UpdateMediaPath**](LibraryStructureApi.md#UpdateMediaPath) | **Post** /Library/VirtualFolders/Paths/Update | Updates a media path.



## AddMediaPath

> AddMediaPath(ctx, mediaPathDto, optional)

Add a media path to a library.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaPathDto** | [**MediaPathDto**](MediaPathDto.md)| The media path dto. | 
 **optional** | ***AddMediaPathOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddMediaPathOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshLibrary** | **optional.Bool**| Whether to refresh the library. | [default to false]

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


## AddVirtualFolder

> AddVirtualFolder(ctx, optional)

Adds a virtual folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddVirtualFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddVirtualFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the virtual folder. | 
 **collectionType** | **optional.String**| The type of the collection. | 
 **paths** | [**optional.Interface of []string**](string.md)| The paths of the virtual folder. | 
 **refreshLibrary** | **optional.Bool**| Whether to refresh the library. | [default to false]
 **addVirtualFolderDto** | [**optional.Interface of AddVirtualFolderDto**](AddVirtualFolderDto.md)| The library options. | 

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


## GetVirtualFolders

> []VirtualFolderInfo GetVirtualFolders(ctx, )

Gets all virtual folders.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]VirtualFolderInfo**](VirtualFolderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMediaPath

> RemoveMediaPath(ctx, optional)

Remove a media path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemoveMediaPathOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveMediaPathOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the library. | 
 **path** | **optional.String**| The path to remove. | 
 **refreshLibrary** | **optional.Bool**| Whether to refresh the library. | [default to false]

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


## RemoveVirtualFolder

> RemoveVirtualFolder(ctx, optional)

Removes a virtual folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemoveVirtualFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveVirtualFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the folder. | 
 **refreshLibrary** | **optional.Bool**| Whether to refresh the library. | [default to false]

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


## RenameVirtualFolder

> RenameVirtualFolder(ctx, optional)

Renames a virtual folder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RenameVirtualFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RenameVirtualFolderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the virtual folder. | 
 **newName** | **optional.String**| The new name. | 
 **refreshLibrary** | **optional.Bool**| Whether to refresh the library. | [default to false]

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


## UpdateLibraryOptions

> UpdateLibraryOptions(ctx, optional)

Update library options.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateLibraryOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateLibraryOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLibraryOptionsDto** | [**optional.Interface of UpdateLibraryOptionsDto**](UpdateLibraryOptionsDto.md)| The library name and options. | 

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


## UpdateMediaPath

> UpdateMediaPath(ctx, optional)

Updates a media path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateMediaPathOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMediaPathOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the library. | 
 **mediaPathInfo** | [**optional.Interface of MediaPathInfo**](MediaPathInfo.md)| The path info. | 

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

