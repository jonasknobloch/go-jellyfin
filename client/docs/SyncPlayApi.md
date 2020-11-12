# \SyncPlayApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncPlayBuffering**](SyncPlayApi.md#SyncPlayBuffering) | **Post** /SyncPlay/Buffering | Request group wait in SyncPlay group while buffering.
[**SyncPlayCreateGroup**](SyncPlayApi.md#SyncPlayCreateGroup) | **Post** /SyncPlay/New | Create a new SyncPlay group.
[**SyncPlayGetGroups**](SyncPlayApi.md#SyncPlayGetGroups) | **Get** /SyncPlay/List | Gets all SyncPlay groups.
[**SyncPlayJoinGroup**](SyncPlayApi.md#SyncPlayJoinGroup) | **Post** /SyncPlay/Join | Join an existing SyncPlay group.
[**SyncPlayLeaveGroup**](SyncPlayApi.md#SyncPlayLeaveGroup) | **Post** /SyncPlay/Leave | Leave the joined SyncPlay group.
[**SyncPlayPause**](SyncPlayApi.md#SyncPlayPause) | **Post** /SyncPlay/Pause | Request pause in SyncPlay group.
[**SyncPlayPing**](SyncPlayApi.md#SyncPlayPing) | **Post** /SyncPlay/Ping | Update session ping.
[**SyncPlayPlay**](SyncPlayApi.md#SyncPlayPlay) | **Post** /SyncPlay/Play | Request play in SyncPlay group.
[**SyncPlaySeek**](SyncPlayApi.md#SyncPlaySeek) | **Post** /SyncPlay/Seek | Request seek in SyncPlay group.



## SyncPlayBuffering

> SyncPlayBuffering(ctx, optional)

Request group wait in SyncPlay group while buffering.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyncPlayBufferingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncPlayBufferingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **when** | **optional.Time**| When the request has been made by the client. | 
 **positionTicks** | **optional.Int64**| The playback position in ticks. | 
 **bufferingDone** | **optional.Bool**| Whether the buffering is done. | 

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


## SyncPlayCreateGroup

> SyncPlayCreateGroup(ctx, )

Create a new SyncPlay group.

### Required Parameters

This endpoint does not need any parameter.

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


## SyncPlayGetGroups

> []GroupInfoView SyncPlayGetGroups(ctx, optional)

Gets all SyncPlay groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyncPlayGetGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncPlayGetGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterItemId** | [**optional.Interface of string**](.md)| Optional. Filter by item id. | 

### Return type

[**[]GroupInfoView**](GroupInfoView.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPlayJoinGroup

> SyncPlayJoinGroup(ctx, groupId)

Join an existing SyncPlay group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**string**](.md)| The sync play group id. | 

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


## SyncPlayLeaveGroup

> SyncPlayLeaveGroup(ctx, )

Leave the joined SyncPlay group.

### Required Parameters

This endpoint does not need any parameter.

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


## SyncPlayPause

> SyncPlayPause(ctx, )

Request pause in SyncPlay group.

### Required Parameters

This endpoint does not need any parameter.

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


## SyncPlayPing

> SyncPlayPing(ctx, optional)

Update session ping.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyncPlayPingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncPlayPingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ping** | **optional.Float64**| The ping. | 

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


## SyncPlayPlay

> SyncPlayPlay(ctx, )

Request play in SyncPlay group.

### Required Parameters

This endpoint does not need any parameter.

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


## SyncPlaySeek

> SyncPlaySeek(ctx, optional)

Request seek in SyncPlay group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SyncPlaySeekOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncPlaySeekOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **positionTicks** | **optional.Int64**| The playback position in ticks. | 

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

