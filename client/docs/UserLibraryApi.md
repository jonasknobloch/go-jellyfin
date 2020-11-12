# \UserLibraryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserItemRating**](UserLibraryApi.md#DeleteUserItemRating) | **Delete** /Users/{userId}/Items/{itemId}/Rating | Deletes a user&#39;s saved personal rating for an item.
[**GetIntros**](UserLibraryApi.md#GetIntros) | **Get** /Users/{userId}/Items/{itemId}/Intros | Gets intros to play before the main media item plays.
[**GetItem**](UserLibraryApi.md#GetItem) | **Get** /Users/{userId}/Items/{itemId} | Gets an item from a user&#39;s library.
[**GetLatestMedia**](UserLibraryApi.md#GetLatestMedia) | **Get** /Users/{userId}/Items/Latest | Gets latest media.
[**GetLocalTrailers**](UserLibraryApi.md#GetLocalTrailers) | **Get** /Users/{userId}/Items/{itemId}/LocalTrailers | Gets local trailers for an item.
[**GetRootFolder**](UserLibraryApi.md#GetRootFolder) | **Get** /Users/{userId}/Items/Root | Gets the root folder from a user&#39;s library.
[**GetSpecialFeatures**](UserLibraryApi.md#GetSpecialFeatures) | **Get** /Users/{userId}/Items/{itemId}/SpecialFeatures | Gets special features for an item.
[**MarkFavoriteItem**](UserLibraryApi.md#MarkFavoriteItem) | **Post** /Users/{userId}/FavoriteItems/{itemId} | Marks an item as a favorite.
[**UnmarkFavoriteItem**](UserLibraryApi.md#UnmarkFavoriteItem) | **Delete** /Users/{userId}/FavoriteItems/{itemId} | Unmarks item as a favorite.
[**UpdateUserItemRating**](UserLibraryApi.md#UpdateUserItemRating) | **Post** /Users/{userId}/Items/{itemId}/Rating | Updates a user&#39;s rating for an item.



## DeleteUserItemRating

> UserItemDataDto DeleteUserItemRating(ctx, userId, itemId)

Deletes a user's saved personal rating for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntros

> BaseItemDtoQueryResult GetIntros(ctx, userId, itemId)

Gets intros to play before the main media item plays.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> BaseItemDto GetItem(ctx, userId, itemId)

Gets an item from a user's library.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestMedia

> []BaseItemDto GetLatestMedia(ctx, userId, optional)

Gets latest media.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
 **optional** | ***GetLatestMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestMediaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | [**optional.Interface of string**](.md)| Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, SortName, Studios, Taglines. | 
 **includeItemTypes** | **optional.String**| Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **isPlayed** | **optional.Bool**| Filter by items that are played, or not. | 
 **enableImages** | **optional.Bool**| Optional. include image information in output. | 
 **imageTypeLimit** | **optional.Int32**| Optional. the max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **enableUserData** | **optional.Bool**| Optional. include user data. | 
 **limit** | **optional.Int32**| Return item limit. | [default to 20]
 **groupItems** | **optional.Bool**| Whether or not to group items into a parent container. | [default to true]

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalTrailers

> []BaseItemDto GetLocalTrailers(ctx, userId, itemId)

Gets local trailers for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootFolder

> BaseItemDto GetRootFolder(ctx, userId)

Gets the root folder from a user's library.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecialFeatures

> []BaseItemDto GetSpecialFeatures(ctx, userId, itemId)

Gets special features for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkFavoriteItem

> UserItemDataDto MarkFavoriteItem(ctx, userId, itemId)

Marks an item as a favorite.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmarkFavoriteItem

> UserItemDataDto UnmarkFavoriteItem(ctx, userId, itemId)

Unmarks item as a favorite.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserItemRating

> UserItemDataDto UpdateUserItemRating(ctx, userId, itemId, optional)

Updates a user's rating for an item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**itemId** | [**string**](.md)| Item id. | 
 **optional** | ***UpdateUserItemRatingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserItemRatingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **likes** | **optional.Bool**| Whether this M:Jellyfin.Api.Controllers.UserLibraryController.UpdateUserItemRating(System.Guid,System.Guid,System.Nullable{System.Boolean}) is likes. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

