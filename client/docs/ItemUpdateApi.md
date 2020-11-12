# \ItemUpdateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetadataEditorInfo**](ItemUpdateApi.md#GetMetadataEditorInfo) | **Get** /Items/{itemId}/MetadataEditor | Gets metadata editor info for an item.
[**UpdateItem**](ItemUpdateApi.md#UpdateItem) | **Post** /Items/{itemId} | Updates an item.
[**UpdateItemContentType**](ItemUpdateApi.md#UpdateItemContentType) | **Post** /Items/{itemId}/ContentType | Updates an item&#39;s content type.



## GetMetadataEditorInfo

> MetadataEditorInfo GetMetadataEditorInfo(ctx, itemId)

Gets metadata editor info for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 

### Return type

[**MetadataEditorInfo**](MetadataEditorInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> UpdateItem(ctx, itemId, baseItemDto)

Updates an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
**baseItemDto** | [**BaseItemDto**](BaseItemDto.md)| The new item properties. | 

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


## UpdateItemContentType

> UpdateItemContentType(ctx, itemId, optional)

Updates an item's content type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| The item id. | 
 **optional** | ***UpdateItemContentTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateItemContentTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **optional.String**| The content type of the item. | 

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

