# \SessionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToSession**](SessionApi.md#AddUserToSession) | **Post** /Sessions/{sessionId}/User/{userId} | Adds an additional user to a session.
[**DisplayContent**](SessionApi.md#DisplayContent) | **Post** /Sessions/{sessionId}/Viewing | Instructs a session to browse to an item or view.
[**GetAuthProviders**](SessionApi.md#GetAuthProviders) | **Get** /Auth/Providers | Get all auth providers.
[**GetPasswordResetProviders**](SessionApi.md#GetPasswordResetProviders) | **Get** /Auth/PasswordResetProviders | Get all password reset providers.
[**GetSessions**](SessionApi.md#GetSessions) | **Get** /Sessions | Gets a list of sessions.
[**Play**](SessionApi.md#Play) | **Post** /Sessions/{sessionId}/Playing | Instructs a session to play an item.
[**PostCapabilities**](SessionApi.md#PostCapabilities) | **Post** /Sessions/Capabilities | Updates capabilities for a device.
[**PostFullCapabilities**](SessionApi.md#PostFullCapabilities) | **Post** /Sessions/Capabilities/Full | Updates capabilities for a device.
[**RemoveUserFromSession**](SessionApi.md#RemoveUserFromSession) | **Delete** /Sessions/{sessionId}/User/{userId} | Removes an additional user from a session.
[**ReportSessionEnded**](SessionApi.md#ReportSessionEnded) | **Post** /Sessions/Logout | Reports that a session has ended.
[**ReportViewing**](SessionApi.md#ReportViewing) | **Post** /Sessions/Viewing | Reports that a session is viewing an item.
[**SendFullGeneralCommand**](SessionApi.md#SendFullGeneralCommand) | **Post** /Sessions/{sessionId}/Command | Issues a full general command to a client.
[**SendGeneralCommand**](SessionApi.md#SendGeneralCommand) | **Post** /Sessions/{sessionId}/Command/{command} | Issues a general command to a client.
[**SendMessageCommand**](SessionApi.md#SendMessageCommand) | **Post** /Sessions/{sessionId}/Message | Issues a command to a client to display a message to the user.
[**SendPlaystateCommand**](SessionApi.md#SendPlaystateCommand) | **Post** /Sessions/{sessionId}/Playing/{command} | Issues a playstate command to a client.
[**SendSystemCommand**](SessionApi.md#SendSystemCommand) | **Post** /Sessions/{sessionId}/System/{command} | Issues a system command to a client.



## AddUserToSession

> AddUserToSession(ctx, sessionId, userId)

Adds an additional user to a session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**userId** | [**string**](.md)| The user id. | 

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


## DisplayContent

> DisplayContent(ctx, sessionId, itemType, itemId, itemName)

Instructs a session to browse to an item or view.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session Id. | 
**itemType** | **string**| The type of item to browse to. | 
**itemId** | **string**| The Id of the item. | 
**itemName** | **string**| The name of the item. | 

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


## GetAuthProviders

> []NameIdPair GetAuthProviders(ctx, )

Get all auth providers.

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


## GetPasswordResetProviders

> []NameIdPair GetPasswordResetProviders(ctx, )

Get all password reset providers.

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


## GetSessions

> []SessionInfo GetSessions(ctx, optional)

Gets a list of sessions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controllableByUserId** | [**optional.Interface of string**](.md)| Filter by sessions that a given user is allowed to remote control. | 
 **deviceId** | **optional.String**| Filter by device Id. | 
 **activeWithinSeconds** | **optional.Int32**| Optional. Filter by sessions that were active in the last n seconds. | 

### Return type

[**[]SessionInfo**](SessionInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Play

> Play(ctx, sessionId, playCommand, itemIds, optional)

Instructs a session to play an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**playCommand** | [**PlayCommand**](.md)| The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now. | 
**itemIds** | **string**| The ids of the items to play, comma delimited. | 
 **optional** | ***PlayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startPositionTicks** | **optional.Int64**| The starting position of the first item. | 

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


## PostCapabilities

> PostCapabilities(ctx, optional)

Updates capabilities for a device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostCapabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostCapabilitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| The session id. | 
 **playableMediaTypes** | **optional.String**| A list of playable media types, comma delimited. Audio, Video, Book, Photo. | 
 **supportedCommands** | [**optional.Interface of []GeneralCommandType**](GeneralCommandType.md)| A list of supported remote control commands, comma delimited. | 
 **supportsMediaControl** | **optional.Bool**| Determines whether media can be played remotely.. | [default to false]
 **supportsSync** | **optional.Bool**| Determines whether sync is supported. | [default to false]
 **supportsPersistentIdentifier** | **optional.Bool**| Determines whether the device supports a unique identifier. | [default to true]

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


## PostFullCapabilities

> PostFullCapabilities(ctx, clientCapabilities, optional)

Updates capabilities for a device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientCapabilities** | [**ClientCapabilities**](ClientCapabilities.md)| The MediaBrowser.Model.Session.ClientCapabilities. | 
 **optional** | ***PostFullCapabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostFullCapabilitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| The session id. | 

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


## RemoveUserFromSession

> RemoveUserFromSession(ctx, sessionId, userId)

Removes an additional user from a session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**userId** | [**string**](.md)| The user id. | 

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


## ReportSessionEnded

> ReportSessionEnded(ctx, )

Reports that a session has ended.

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


## ReportViewing

> ReportViewing(ctx, itemId, optional)

Reports that a session is viewing an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| The item id. | 
 **optional** | ***ReportViewingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReportViewingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionId** | **optional.String**| The session id. | 

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


## SendFullGeneralCommand

> SendFullGeneralCommand(ctx, sessionId, generalCommand)

Issues a full general command to a client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**generalCommand** | [**GeneralCommand**](GeneralCommand.md)| The MediaBrowser.Model.Session.GeneralCommand. | 

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


## SendGeneralCommand

> SendGeneralCommand(ctx, sessionId, command)

Issues a general command to a client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**command** | [**GeneralCommandType**](.md)| The command to send. | 

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


## SendMessageCommand

> SendMessageCommand(ctx, sessionId, text, optional)

Issues a command to a client to display a message to the user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**text** | **string**| The message test. | 
 **optional** | ***SendMessageCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendMessageCommandOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **header** | **optional.String**| The message header. | 
 **timeoutMs** | **optional.Int64**| The message timeout. If omitted the user will have to confirm viewing the message. | 

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


## SendPlaystateCommand

> SendPlaystateCommand(ctx, sessionId, command, optional)

Issues a playstate command to a client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**command** | [**PlaystateCommand**](.md)| The MediaBrowser.Model.Session.PlaystateCommand. | 
 **optional** | ***SendPlaystateCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendPlaystateCommandOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seekPositionTicks** | **optional.Int64**| The optional position ticks. | 
 **controllingUserId** | **optional.String**| The optional controlling user id. | 

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


## SendSystemCommand

> SendSystemCommand(ctx, sessionId, command)

Issues a system command to a client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string**| The session id. | 
**command** | [**GeneralCommandType**](.md)| The command to send. | 

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

