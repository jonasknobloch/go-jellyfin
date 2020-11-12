# \QuickConnectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activate**](QuickConnectApi.md#Activate) | **Post** /QuickConnect/Activate | Temporarily activates quick connect for five minutes.
[**Authorize**](QuickConnectApi.md#Authorize) | **Post** /QuickConnect/Authorize | Authorizes a pending quick connect request.
[**Available**](QuickConnectApi.md#Available) | **Post** /QuickConnect/Available | Enables or disables quick connect.
[**Connect**](QuickConnectApi.md#Connect) | **Get** /QuickConnect/Connect | Attempts to retrieve authentication information.
[**Deauthorize**](QuickConnectApi.md#Deauthorize) | **Post** /QuickConnect/Deauthorize | Deauthorize all quick connect devices for the current user.
[**GetStatus**](QuickConnectApi.md#GetStatus) | **Get** /QuickConnect/Status | Gets the current quick connect state.
[**Initiate**](QuickConnectApi.md#Initiate) | **Get** /QuickConnect/Initiate | Initiate a new quick connect request.



## Activate

> Activate(ctx, )

Temporarily activates quick connect for five minutes.

### Required Parameters

This endpoint does not need any parameter.

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


## Authorize

> bool Authorize(ctx, code)

Authorizes a pending quick connect request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**| Quick connect code to authorize. | 

### Return type

**bool**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Available

> Available(ctx, optional)

Enables or disables quick connect.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AvailableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AvailableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**optional.Interface of QuickConnectState**](.md)| New MediaBrowser.Model.QuickConnect.QuickConnectState. | 

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


## Connect

> QuickConnectResult Connect(ctx, secret)

Attempts to retrieve authentication information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secret** | **string**| Secret previously returned from the Initiate endpoint. | 

### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deauthorize

> int32 Deauthorize(ctx, )

Deauthorize all quick connect devices for the current user.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**int32**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> QuickConnectState GetStatus(ctx, )

Gets the current quick connect state.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**QuickConnectState**](QuickConnectState.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Initiate

> QuickConnectResult Initiate(ctx, )

Initiate a new quick connect request.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

