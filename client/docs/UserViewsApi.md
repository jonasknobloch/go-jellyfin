# \UserViewsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupingOptions**](UserViewsApi.md#GetGroupingOptions) | **Get** /Users/{userId}/GroupingOptions | Get user view grouping options.
[**GetUserViews**](UserViewsApi.md#GetUserViews) | **Get** /Users/{userId}/Views | Get user views.



## GetGroupingOptions

> []SpecialViewOptionDto GetGroupingOptions(ctx, userId)

Get user view grouping options.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 

### Return type

[**[]SpecialViewOptionDto**](SpecialViewOptionDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserViews

> BaseItemDtoQueryResult GetUserViews(ctx, userId, optional)

Get user views.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
 **optional** | ***GetUserViewsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserViewsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeExternalContent** | **optional.Bool**| Whether or not to include external views such as channels or live tv. | 
 **presetViews** | **optional.String**| Preset views. | 
 **includeHidden** | **optional.Bool**| Whether or not to include hidden content. | [default to false]

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

