# \CollectionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToCollection**](CollectionApi.md#AddToCollection) | **Post** /Collections/{collectionId}/Items | Adds items to a collection.
[**CreateCollection**](CollectionApi.md#CreateCollection) | **Post** /Collections | Creates a new collection.
[**RemoveFromCollection**](CollectionApi.md#RemoveFromCollection) | **Delete** /Collections/{collectionId}/Items | Removes items from a collection.



## AddToCollection

> AddToCollection(ctx, collectionId, itemIds)

Adds items to a collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | [**string**](.md)| The collection id. | 
**itemIds** | **string**| Item ids, comma delimited. | 

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


## CreateCollection

> CollectionCreationResult CreateCollection(ctx, optional)

Creates a new collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCollectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of the collection. | 
 **ids** | **optional.String**| Item Ids to add to the collection. | 
 **parentId** | [**optional.Interface of string**](.md)| Optional. Create the collection within a specific folder. | 
 **isLocked** | **optional.Bool**| Whether or not to lock the new collection. | [default to false]

### Return type

[**CollectionCreationResult**](CollectionCreationResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFromCollection

> RemoveFromCollection(ctx, collectionId, itemIds)

Removes items from a collection.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | [**string**](.md)| The collection id. | 
**itemIds** | **string**| Item ids, comma delimited. | 

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

