# \DashboardApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationPages**](DashboardApi.md#GetConfigurationPages) | **Get** /web/ConfigurationPages | Gets the configuration pages.
[**GetDashboardConfigurationPage**](DashboardApi.md#GetDashboardConfigurationPage) | **Get** /web/ConfigurationPage | Gets a dashboard configuration page.



## GetConfigurationPages

> []ConfigurationPageInfo GetConfigurationPages(ctx, optional)

Gets the configuration pages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetConfigurationPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetConfigurationPagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableInMainMenu** | **optional.Bool**| Whether to enable in the main menu. | 
 **pageType** | [**optional.Interface of ConfigurationPageType**](.md)| The Jellyfin.Api.Models.ConfigurationPageInfo. | 

### Return type

[**[]ConfigurationPageInfo**](ConfigurationPageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardConfigurationPage

> *os.File GetDashboardConfigurationPage(ctx, optional)

Gets a dashboard configuration page.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDashboardConfigurationPageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDashboardConfigurationPageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the page. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/x-javascript, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

