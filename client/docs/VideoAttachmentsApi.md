# \VideoAttachmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachment**](VideoAttachmentsApi.md#GetAttachment) | **Get** /Videos/{videoId}/{mediaSourceId}/Attachments/{index} | Get video attachment.



## GetAttachment

> *os.File GetAttachment(ctx, videoId, mediaSourceId, index)

Get video attachment.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | [**string**](.md)| Video ID. | 
**mediaSourceId** | **string**| Media Source ID. | 
**index** | **int32**| Attachment Index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

