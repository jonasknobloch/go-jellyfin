# \NotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdminNotification**](NotificationsApi.md#CreateAdminNotification) | **Post** /Notifications/Admin | Sends a notification to all admins.
[**GetNotificationServices**](NotificationsApi.md#GetNotificationServices) | **Get** /Notifications/Services | Gets notification services.
[**GetNotificationTypes**](NotificationsApi.md#GetNotificationTypes) | **Get** /Notifications/Types | Gets notification types.
[**GetNotifications**](NotificationsApi.md#GetNotifications) | **Get** /Notifications/{userId} | Gets a user&#39;s notifications.
[**GetNotificationsSummary**](NotificationsApi.md#GetNotificationsSummary) | **Get** /Notifications/{userId}/Summary | Gets a user&#39;s notification summary.
[**SetRead**](NotificationsApi.md#SetRead) | **Post** /Notifications/{userId}/Read | Sets notifications as read.
[**SetUnread**](NotificationsApi.md#SetUnread) | **Post** /Notifications/{userId}/Unread | Sets notifications as unread.



## CreateAdminNotification

> CreateAdminNotification(ctx, optional)

Sends a notification to all admins.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAdminNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAdminNotificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**| The URL of the notification. | 
 **level** | [**optional.Interface of NotificationLevel**](.md)| The level of the notification. | 
 **name** | **optional.String**| The name of the notification. | [default to ]
 **description** | **optional.String**| The description of the notification. | [default to ]

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


## GetNotificationServices

> []NameIdPair GetNotificationServices(ctx, )

Gets notification services.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationTypes

> []NotificationTypeInfo GetNotificationTypes(ctx, )

Gets notification types.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]NotificationTypeInfo**](NotificationTypeInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> NotificationResultDto GetNotifications(ctx, userId)

Gets a user's notifications.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**NotificationResultDto**](NotificationResultDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsSummary

> NotificationsSummaryDto GetNotificationsSummary(ctx, userId)

Gets a user's notification summary.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**NotificationsSummaryDto**](NotificationsSummaryDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRead

> SetRead(ctx, userId)

Sets notifications as read.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

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


## SetUnread

> SetUnread(ctx, userId)

Sets notifications as unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

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

