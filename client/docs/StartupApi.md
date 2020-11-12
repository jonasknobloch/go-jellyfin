# \StartupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteWizard**](StartupApi.md#CompleteWizard) | **Post** /Startup/Complete | Completes the startup wizard.
[**GetFirstUser**](StartupApi.md#GetFirstUser) | **Get** /Startup/User | Gets the first user.
[**GetFirstUser2**](StartupApi.md#GetFirstUser2) | **Get** /Startup/FirstUser | Gets the first user.
[**GetStartupConfiguration**](StartupApi.md#GetStartupConfiguration) | **Get** /Startup/Configuration | Gets the initial startup wizard configuration.
[**SetRemoteAccess**](StartupApi.md#SetRemoteAccess) | **Post** /Startup/RemoteAccess | Sets remote access and UPnP.
[**UpdateInitialConfiguration**](StartupApi.md#UpdateInitialConfiguration) | **Post** /Startup/Configuration | Sets the initial startup wizard configuration.
[**UpdateStartupUser**](StartupApi.md#UpdateStartupUser) | **Post** /Startup/User | Sets the user name and password.



## CompleteWizard

> CompleteWizard(ctx, )

Completes the startup wizard.

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


## GetFirstUser

> StartupUserDto GetFirstUser(ctx, )

Gets the first user.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StartupUserDto**](StartupUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirstUser2

> StartupUserDto GetFirstUser2(ctx, )

Gets the first user.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StartupUserDto**](StartupUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartupConfiguration

> StartupConfigurationDto GetStartupConfiguration(ctx, )

Gets the initial startup wizard configuration.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StartupConfigurationDto**](StartupConfigurationDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRemoteAccess

> SetRemoteAccess(ctx, startupRemoteAccessDto)

Sets remote access and UPnP.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startupRemoteAccessDto** | [**StartupRemoteAccessDto**](StartupRemoteAccessDto.md)| The startup remote access dto. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInitialConfiguration

> UpdateInitialConfiguration(ctx, startupConfigurationDto)

Sets the initial startup wizard configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**startupConfigurationDto** | [**StartupConfigurationDto**](StartupConfigurationDto.md)| The updated startup configuration. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStartupUser

> UpdateStartupUser(ctx, optional)

Sets the user name and password.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateStartupUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateStartupUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startupUserDto** | [**optional.Interface of StartupUserDto**](StartupUserDto.md)| The DTO containing username and password. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

