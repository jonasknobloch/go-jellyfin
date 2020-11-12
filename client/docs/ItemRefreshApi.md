# \ItemRefreshApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Post**](ItemRefreshApi.md#Post) | **Post** /Items/{itemId}/Refresh | Refreshes metadata for an item.



## Post

> Post(ctx, itemId, optional)

Refreshes metadata for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***PostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataRefreshMode** | [**optional.Interface of MetadataRefreshMode**](.md)| (Optional) Specifies the metadata refresh mode. | 
 **imageRefreshMode** | [**optional.Interface of MetadataRefreshMode**](.md)| (Optional) Specifies the image refresh mode. | 
 **replaceAllMetadata** | **optional.Bool**| (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh. | [default to false]
 **replaceAllImages** | **optional.Bool**| (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh. | [default to false]

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

