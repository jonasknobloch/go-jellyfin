# \FilterApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryFilters**](FilterApi.md#GetQueryFilters) | **Get** /Items/Filters2 | Gets query filters.
[**GetQueryFiltersLegacy**](FilterApi.md#GetQueryFiltersLegacy) | **Get** /Items/Filters | Gets legacy query filters.



## GetQueryFilters

> QueryFilters GetQueryFilters(ctx, optional)

Gets query filters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetQueryFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQueryFiltersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. User id. | 
 **parentId** | **optional.String**| Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **includeItemTypes** | **optional.String**| Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **isAiring** | **optional.Bool**| Optional. Is item airing. | 
 **isMovie** | **optional.Bool**| Optional. Is item movie. | 
 **isSports** | **optional.Bool**| Optional. Is item sports. | 
 **isKids** | **optional.Bool**| Optional. Is item kids. | 
 **isNews** | **optional.Bool**| Optional. Is item news. | 
 **isSeries** | **optional.Bool**| Optional. Is item series. | 
 **recursive** | **optional.Bool**| Optional. Search recursive. | 

### Return type

[**QueryFilters**](QueryFilters.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryFiltersLegacy

> QueryFiltersLegacy GetQueryFiltersLegacy(ctx, optional)

Gets legacy query filters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetQueryFiltersLegacyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQueryFiltersLegacyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. User id. | 
 **parentId** | **optional.String**| Optional. Parent id. | 
 **includeItemTypes** | **optional.String**| Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **mediaTypes** | **optional.String**| Optional. Filter by MediaType. Allows multiple, comma delimited. | 

### Return type

[**QueryFiltersLegacy**](QueryFiltersLegacy.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

