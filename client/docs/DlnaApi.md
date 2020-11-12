# \DlnaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](DlnaApi.md#CreateProfile) | **Post** /Dlna/Profiles | Creates a profile.
[**DeleteProfile**](DlnaApi.md#DeleteProfile) | **Delete** /Dlna/Profiles/{profileId} | Deletes a profile.
[**GetDefaultProfile**](DlnaApi.md#GetDefaultProfile) | **Get** /Dlna/Profiles/Default | Gets the default profile.
[**GetProfile**](DlnaApi.md#GetProfile) | **Get** /Dlna/Profiles/{profileId} | Gets a single profile.
[**GetProfileInfos**](DlnaApi.md#GetProfileInfos) | **Get** /Dlna/ProfileInfos | Get profile infos.
[**UpdateProfile**](DlnaApi.md#UpdateProfile) | **Post** /Dlna/Profiles/{profileId} | Updates a profile.



## CreateProfile

> CreateProfile(ctx, optional)

Creates a profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceProfile** | [**optional.Interface of DeviceProfile**](DeviceProfile.md)| Device profile. | 

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


## DeleteProfile

> DeleteProfile(ctx, profileId)

Deletes a profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string**| Profile id. | 

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


## GetDefaultProfile

> DeviceProfile GetDefaultProfile(ctx, )

Gets the default profile.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> DeviceProfile GetProfile(ctx, profileId)

Gets a single profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string**| Profile Id. | 

### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileInfos

> []DeviceProfileInfo GetProfileInfos(ctx, )

Get profile infos.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DeviceProfileInfo**](DeviceProfileInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile(ctx, profileId, optional)

Updates a profile.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string**| Profile id. | 
 **optional** | ***UpdateProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProfileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceProfile** | [**optional.Interface of DeviceProfile**](DeviceProfile.md)| Device profile. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

