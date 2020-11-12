# \PluginsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPluginConfiguration**](PluginsApi.md#GetPluginConfiguration) | **Get** /Plugins/{pluginId}/Configuration | Gets plugin configuration.
[**GetPluginSecurityInfo**](PluginsApi.md#GetPluginSecurityInfo) | **Get** /Plugins/SecurityInfo | Get plugin security info.
[**GetPlugins**](PluginsApi.md#GetPlugins) | **Get** /Plugins | Gets a list of currently installed plugins.
[**GetRegistration**](PluginsApi.md#GetRegistration) | **Get** /Plugins/Registrations/{name} | Gets registration status for a feature.
[**GetRegistrationStatus**](PluginsApi.md#GetRegistrationStatus) | **Post** /Plugins/RegistrationRecords/{name} | Gets registration status for a feature.
[**UninstallPlugin**](PluginsApi.md#UninstallPlugin) | **Delete** /Plugins/{pluginId} | Uninstalls a plugin.
[**UpdatePluginConfiguration**](PluginsApi.md#UpdatePluginConfiguration) | **Post** /Plugins/{pluginId}/Configuration | Updates plugin configuration.
[**UpdatePluginSecurityInfo**](PluginsApi.md#UpdatePluginSecurityInfo) | **Post** /Plugins/SecurityInfo | Updates plugin security info.



## GetPluginConfiguration

> map[string]interface{} GetPluginConfiguration(ctx, pluginId)

Gets plugin configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | [**string**](.md)| Plugin id. | 

### Return type

**map[string]interface{}**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginSecurityInfo

> PluginSecurityInfo GetPluginSecurityInfo(ctx, )

Get plugin security info.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PluginSecurityInfo**](PluginSecurityInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> []PluginInfo GetPlugins(ctx, )

Gets a list of currently installed plugins.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PluginInfo**](PluginInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistration

> GetRegistration(ctx, name)

Gets registration status for a feature.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Feature name. | 

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


## GetRegistrationStatus

> MbRegistrationRecord GetRegistrationStatus(ctx, name)

Gets registration status for a feature.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Feature name. | 

### Return type

[**MbRegistrationRecord**](MBRegistrationRecord.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallPlugin

> UninstallPlugin(ctx, pluginId)

Uninstalls a plugin.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | [**string**](.md)| Plugin id. | 

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


## UpdatePluginConfiguration

> UpdatePluginConfiguration(ctx, pluginId)

Updates plugin configuration.

Accepts plugin configuration as JSON body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | [**string**](.md)| Plugin id. | 

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


## UpdatePluginSecurityInfo

> UpdatePluginSecurityInfo(ctx, pluginSecurityInfo)

Updates plugin security info.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginSecurityInfo** | [**PluginSecurityInfo**](PluginSecurityInfo.md)| Plugin security info. | 

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

