# \ImageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItemImage**](ImageApi.md#DeleteItemImage) | **Delete** /Items/{itemId}/Images/{imageType} | Delete an item&#39;s image.
[**DeleteItemImage2**](ImageApi.md#DeleteItemImage2) | **Delete** /Items/{itemId}/Images/{imageType}/{imageIndex} | Delete an item&#39;s image.
[**DeleteUserImage**](ImageApi.md#DeleteUserImage) | **Delete** /Users/{userId}/Images/{itemType} | Delete the user&#39;s image.
[**DeleteUserImage2**](ImageApi.md#DeleteUserImage2) | **Delete** /Users/{userId}/Images/{itemType}/{index} | Delete the user&#39;s image.
[**GetArtistImage**](ImageApi.md#GetArtistImage) | **Get** /Artists/{name}/Images/{imageType}/{imageIndex} | Get artist image by name.
[**GetGenreImage**](ImageApi.md#GetGenreImage) | **Get** /Genres/{name}/Images/{imageType}/{imageIndex} | Get genre image by name.
[**GetItemImage**](ImageApi.md#GetItemImage) | **Get** /Items/{itemId}/Images/{imageType} | Gets the item&#39;s image.
[**GetItemImage2**](ImageApi.md#GetItemImage2) | **Get** /Items/{itemId}/Images/{imageType}/{imageIndex} | Gets the item&#39;s image.
[**GetItemImage2_0**](ImageApi.md#GetItemImage2_0) | **Get** /Items/{itemId}/Images/{imageType}/{imageIndex}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount} | Gets the item&#39;s image.
[**GetItemImageInfos**](ImageApi.md#GetItemImageInfos) | **Get** /Items/{itemId}/Images | Get item image infos.
[**GetMusicGenreImage**](ImageApi.md#GetMusicGenreImage) | **Get** /MusicGenres/{name}/Images/{imageType}/{imageIndex} | Get music genre image by name.
[**GetPersonImage**](ImageApi.md#GetPersonImage) | **Get** /Persons/{name}/Images/{imageType}/{imageIndex} | Get person image by name.
[**GetStudioImage**](ImageApi.md#GetStudioImage) | **Get** /Studios/{name}/Images/{imageType}/{imageIndex} | Get studio image by name.
[**GetUserImage**](ImageApi.md#GetUserImage) | **Get** /Users/{userId}/Images/{imageType}/{imageIndex} | Get user profile image.
[**HeadArtistImage**](ImageApi.md#HeadArtistImage) | **Head** /Artists/{name}/Images/{imageType}/{imageIndex} | Get artist image by name.
[**HeadGenreImage**](ImageApi.md#HeadGenreImage) | **Head** /Genres/{name}/Images/{imageType}/{imageIndex} | Get genre image by name.
[**HeadItemImage**](ImageApi.md#HeadItemImage) | **Head** /Items/{itemId}/Images/{imageType} | Gets the item&#39;s image.
[**HeadItemImage2**](ImageApi.md#HeadItemImage2) | **Head** /Items/{itemId}/Images/{imageType}/{imageIndex} | Gets the item&#39;s image.
[**HeadItemImage2_0**](ImageApi.md#HeadItemImage2_0) | **Head** /Items/{itemId}/Images/{imageType}/{imageIndex}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount} | Gets the item&#39;s image.
[**HeadMusicGenreImage**](ImageApi.md#HeadMusicGenreImage) | **Head** /MusicGenres/{name}/Images/{imageType}/{imageIndex} | Get music genre image by name.
[**HeadPersonImage**](ImageApi.md#HeadPersonImage) | **Head** /Persons/{name}/Images/{imageType}/{imageIndex} | Get person image by name.
[**HeadStudioImage**](ImageApi.md#HeadStudioImage) | **Head** /Studios/{name}/Images/{imageType}/{imageIndex} | Get studio image by name.
[**HeadUserImage**](ImageApi.md#HeadUserImage) | **Head** /Users/{userId}/Images/{imageType}/{imageIndex} | Get user profile image.
[**PostUserImage**](ImageApi.md#PostUserImage) | **Post** /Users/{userId}/Images/{imageType} | Sets the user image.
[**PostUserImage2**](ImageApi.md#PostUserImage2) | **Post** /Users/{userId}/Images/{imageType}/{index} | Sets the user image.
[**SetItemImage**](ImageApi.md#SetItemImage) | **Post** /Items/{itemId}/Images/{imageType} | Set item image.
[**SetItemImage2**](ImageApi.md#SetItemImage2) | **Post** /Items/{itemId}/Images/{imageType}/{imageIndex} | Set item image.
[**UpdateItemImageIndex**](ImageApi.md#UpdateItemImageIndex) | **Post** /Items/{itemId}/Images/{imageType}/{imageIndex}/Index | Updates the index for an item image.



## DeleteItemImage

> DeleteItemImage(ctx, itemId, imageType, imageIndex)

Delete an item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| The image index. | 

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


## DeleteItemImage2

> DeleteItemImage2(ctx, itemId, imageType, imageIndex)

Delete an item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| The image index. | 

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


## DeleteUserImage

> DeleteUserImage(ctx, userId, imageType, index, itemType)

Delete the user's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User Id. | 
**imageType** | [**ImageType**](.md)| (Unused) Image type. | 
**index** | **int32**| (Unused) Image index. | 
**itemType** | **string**|  | 

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


## DeleteUserImage2

> DeleteUserImage2(ctx, userId, imageType, index, itemType)

Delete the user's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User Id. | 
**imageType** | [**ImageType**](.md)| (Unused) Image type. | 
**index** | **int32**| (Unused) Image index. | 
**itemType** | **string**|  | 

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


## GetArtistImage

> *os.File GetArtistImage(ctx, name, imageType, imageIndex, optional)

Get artist image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Artist name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetArtistImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetArtistImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenreImage

> *os.File GetGenreImage(ctx, name, imageType, imageIndex, optional)

Get genre image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Genre name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetGenreImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetGenreImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImage

> *os.File GetItemImage(ctx, itemId, imageType, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetItemImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImage2

> *os.File GetItemImage2(ctx, itemId, imageType, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetItemImage2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemImage2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImage2_0

> *os.File GetItemImage2_0(ctx, itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**maxWidth** | **int32**| The maximum image width to return. | 
**maxHeight** | **int32**| The maximum image height to return. | 
**tag** | **string**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | [**ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
**percentPlayed** | **float64**| Optional. Percent to render for the percent played overlay. | 
**unplayedCount** | **int32**| Optional. Unplayed count overlay to render. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetItemImage2_1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemImage2_1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImageInfos

> []ImageInfo GetItemImageInfos(ctx, itemId)

Get item image infos.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 

### Return type

[**[]ImageInfo**](ImageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicGenreImage

> *os.File GetMusicGenreImage(ctx, name, imageType, imageIndex, optional)

Get music genre image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Music genre name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetMusicGenreImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMusicGenreImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonImage

> *os.File GetPersonImage(ctx, name, imageType, imageIndex, optional)

Get person image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Person name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetPersonImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPersonImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioImage

> *os.File GetStudioImage(ctx, name, imageType, tag, format, imageIndex, optional)

Get studio image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Studio name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**tag** | **string**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | [**ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetStudioImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStudioImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserImage

> *os.File GetUserImage(ctx, userId, imageType, imageIndex, optional)

Get user profile image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***GetUserImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadArtistImage

> *os.File HeadArtistImage(ctx, name, imageType, imageIndex, optional)

Get artist image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Artist name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadArtistImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadArtistImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadGenreImage

> *os.File HeadGenreImage(ctx, name, imageType, imageIndex, optional)

Get genre image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Genre name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadGenreImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadGenreImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImage

> *os.File HeadItemImage(ctx, itemId, imageType, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadItemImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadItemImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImage2

> *os.File HeadItemImage2(ctx, itemId, imageType, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadItemImage2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadItemImage2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImage2_0

> *os.File HeadItemImage2_0(ctx, itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex, optional)

Gets the item's image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**maxWidth** | **int32**| The maximum image width to return. | 
**maxHeight** | **int32**| The maximum image height to return. | 
**tag** | **string**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | [**ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
**percentPlayed** | **float64**| Optional. Percent to render for the percent played overlay. | 
**unplayedCount** | **int32**| Optional. Unplayed count overlay to render. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadItemImage2_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadItemImage2_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMusicGenreImage

> *os.File HeadMusicGenreImage(ctx, name, imageType, imageIndex, optional)

Get music genre image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Music genre name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadMusicGenreImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadMusicGenreImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadPersonImage

> *os.File HeadPersonImage(ctx, name, imageType, imageIndex, optional)

Get person image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Person name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadPersonImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadPersonImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadStudioImage

> *os.File HeadStudioImage(ctx, name, imageType, tag, format, imageIndex, optional)

Get studio image by name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Studio name. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**tag** | **string**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | [**ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadStudioImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadStudioImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadUserImage

> *os.File HeadUserImage(ctx, userId, imageType, imageIndex, optional)

Get user profile image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Image index. | 
 **optional** | ***HeadUserImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HeadUserImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **optional.String**| Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | [**optional.Interface of ImageFormat**](.md)| Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **optional.Int32**| The maximum image width to return. | 
 **maxHeight** | **optional.Int32**| The maximum image height to return. | 
 **percentPlayed** | **optional.Float64**| Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **optional.Int32**| Optional. Unplayed count overlay to render. | 
 **width** | **optional.Int32**| The fixed image width to return. | 
 **height** | **optional.Int32**| The fixed image height to return. | 
 **quality** | **optional.Int32**| Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **cropWhitespace** | **optional.Bool**| Optional. Specify if whitespace should be cropped out of the image. True/False. If unspecified, whitespace will be cropped from logos and clear art. | 
 **addPlayedIndicator** | **optional.Bool**| Optional. Add a played indicator. | 
 **blur** | **optional.Int32**| Optional. Blur image. | 
 **backgroundColor** | **optional.String**| Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **optional.String**| Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserImage

> PostUserImage(ctx, userId, imageType, index)

Sets the user image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User Id. | 
**imageType** | [**ImageType**](.md)| (Unused) Image type. | 
**index** | **int32**| (Unused) Image index. | 

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


## PostUserImage2

> PostUserImage2(ctx, userId, imageType, index)

Sets the user image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| User Id. | 
**imageType** | [**ImageType**](.md)| (Unused) Image type. | 
**index** | **int32**| (Unused) Image index. | 

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


## SetItemImage

> SetItemImage(ctx, itemId, imageType, imageIndex)

Set item image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| (Unused) Image index. | 

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


## SetItemImage2

> SetItemImage2(ctx, itemId, imageType, imageIndex)

Set item image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| (Unused) Image index. | 

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


## UpdateItemImageIndex

> UpdateItemImageIndex(ctx, itemId, imageType, imageIndex, optional)

Updates the index for an item image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | [**string**](.md)| Item id. | 
**imageType** | [**ImageType**](.md)| Image type. | 
**imageIndex** | **int32**| Old image index. | 
 **optional** | ***UpdateItemImageIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateItemImageIndexOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **newIndex** | **optional.Int32**| New image index. | 

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

