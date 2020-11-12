# \SuggestionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSuggestions**](SuggestionsApi.md#GetSuggestions) | **Get** /Users/{userId}/Suggestions | Gets suggestions.



## GetSuggestions

> BaseItemDtoQueryResult GetSuggestions(ctx, userId, optional)

Gets suggestions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The user id. | 
 **optional** | ***GetSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSuggestionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaType** | **optional.String**| The media types. | 
 **type_** | **optional.String**| The type. | 
 **startIndex** | **optional.Int32**| Optional. The start index. | 
 **limit** | **optional.Int32**| Optional. The limit. | 
 **enableTotalRecordCount** | **optional.Bool**| Whether to enable the total record count. | [default to false]

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

