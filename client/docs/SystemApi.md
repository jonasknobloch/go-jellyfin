# \SystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEndpointInfo**](SystemApi.md#GetEndpointInfo) | **Get** /System/Endpoint | Gets information about the request endpoint.
[**GetLogFile**](SystemApi.md#GetLogFile) | **Get** /System/Logs/Log | Gets a log file.
[**GetPingSystem**](SystemApi.md#GetPingSystem) | **Get** /System/Ping | Pings the system.
[**GetPublicSystemInfo**](SystemApi.md#GetPublicSystemInfo) | **Get** /System/Info/Public | Gets public information about the server.
[**GetServerLogs**](SystemApi.md#GetServerLogs) | **Get** /System/Logs | Gets a list of available server log files.
[**GetSystemInfo**](SystemApi.md#GetSystemInfo) | **Get** /System/Info | Gets information about the server.
[**GetWakeOnLanInfo**](SystemApi.md#GetWakeOnLanInfo) | **Get** /System/WakeOnLanInfo | Gets wake on lan information.
[**PostPingSystem**](SystemApi.md#PostPingSystem) | **Post** /System/Ping | Pings the system.
[**RestartApplication**](SystemApi.md#RestartApplication) | **Post** /System/Restart | Restarts the application.
[**ShutdownApplication**](SystemApi.md#ShutdownApplication) | **Post** /System/Shutdown | Shuts down the application.



## GetEndpointInfo

> EndPointInfo GetEndpointInfo(ctx, )

Gets information about the request endpoint.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EndPointInfo**](EndPointInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogFile

> *os.File GetLogFile(ctx, name)

Gets a log file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the log file to get. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingSystem

> string GetPingSystem(ctx, )

Pings the system.

### Required Parameters

This endpoint does not need any parameter.

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


## GetPublicSystemInfo

> PublicSystemInfo GetPublicSystemInfo(ctx, )

Gets public information about the server.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PublicSystemInfo**](PublicSystemInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerLogs

> []LogFile GetServerLogs(ctx, )

Gets a list of available server log files.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]LogFile**](LogFile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInfo

> SystemInfo GetSystemInfo(ctx, )

Gets information about the server.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWakeOnLanInfo

> []WakeOnLanInfo GetWakeOnLanInfo(ctx, )

Gets wake on lan information.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]WakeOnLanInfo**](WakeOnLanInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPingSystem

> string PostPingSystem(ctx, )

Pings the system.

### Required Parameters

This endpoint does not need any parameter.

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


## RestartApplication

> RestartApplication(ctx, )

Restarts the application.

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


## ShutdownApplication

> ShutdownApplication(ctx, )

Shuts down the application.

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

