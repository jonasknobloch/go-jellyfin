# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](SearchApi.md#Get) | **Get** /Search/Hints | Gets the search hint result.



## Get

> SearchHintResult Get(ctx, searchTerm, optional)

Gets the search hint result.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchTerm** | **string**| The search term to filter on. | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **userId** | [**optional.Interface of string**](.md)| Optional. Supply a user id to search within a user&#39;s library or omit to search all. | 
 **includeItemTypes** | **optional.String**| If specified, only results with the specified item types are returned. This allows multiple, comma delimeted. | 
 **excludeItemTypes** | **optional.String**| If specified, results with these item types are filtered out. This allows multiple, comma delimeted. | 
 **mediaTypes** | **optional.String**| If specified, only results with the specified media types are returned. This allows multiple, comma delimeted. | 
 **parentId** | **optional.String**| If specified, only children of the parent are returned. | 
 **isMovie** | **optional.Bool**| Optional filter for movies. | 
 **isSeries** | **optional.Bool**| Optional filter for series. | 
 **isNews** | **optional.Bool**| Optional filter for news. | 
 **isKids** | **optional.Bool**| Optional filter for kids. | 
 **isSports** | **optional.Bool**| Optional filter for sports. | 
 **includePeople** | **optional.Bool**| Optional filter whether to include people. | [default to true]
 **includeMedia** | **optional.Bool**| Optional filter whether to include media. | [default to true]
 **includeGenres** | **optional.Bool**| Optional filter whether to include genres. | [default to true]
 **includeStudios** | **optional.Bool**| Optional filter whether to include studios. | [default to true]
 **includeArtists** | **optional.Bool**| Optional filter whether to include artists. | [default to true]

### Return type

[**SearchHintResult**](SearchHintResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

