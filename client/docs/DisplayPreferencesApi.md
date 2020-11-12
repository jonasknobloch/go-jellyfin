# \DisplayPreferencesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisplayPreferences**](DisplayPreferencesApi.md#GetDisplayPreferences) | **Get** /DisplayPreferences/{displayPreferencesId} | Get Display Preferences.
[**UpdateDisplayPreferences**](DisplayPreferencesApi.md#UpdateDisplayPreferences) | **Post** /DisplayPreferences/{displayPreferencesId} | Update Display Preferences.



## GetDisplayPreferences

> DisplayPreferencesDto GetDisplayPreferences(ctx, displayPreferencesId, userId, client)

Get Display Preferences.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string**| Display preferences id. | 
**userId** | [**string**](.md)| User id. | 
**client** | **string**| Client. | 

### Return type

[**DisplayPreferencesDto**](DisplayPreferencesDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDisplayPreferences

> UpdateDisplayPreferences(ctx, displayPreferencesId, userId, client, displayPreferencesDto)

Update Display Preferences.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string**| Display preferences id. | 
**userId** | [**string**](.md)| User Id. | 
**client** | **string**| Client. | 
**displayPreferencesDto** | [**DisplayPreferencesDto**](DisplayPreferencesDto.md)| New Display Preferences object. | 

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

