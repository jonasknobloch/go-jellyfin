# \ActivityLogApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogEntries**](ActivityLogApi.md#GetLogEntries) | **Get** /System/ActivityLog/Entries | Gets activity log entries.



## GetLogEntries

> ActivityLogEntryQueryResult GetLogEntries(ctx, optional)

Gets activity log entries.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLogEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLogEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **minDate** | **optional.Time**| Optional. The minimum date. Format &#x3D; ISO. | 
 **hasUserId** | **optional.Bool**| Optional. Filter log entries if it has user id, or not. | 

### Return type

[**ActivityLogEntryQueryResult**](ActivityLogEntryQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

