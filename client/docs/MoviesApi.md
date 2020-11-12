# \MoviesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovieRecommendations**](MoviesApi.md#GetMovieRecommendations) | **Get** /Movies/Recommendations | Gets movie recommendations.



## GetMovieRecommendations

> []RecommendationDto GetMovieRecommendations(ctx, optional)

Gets movie recommendations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMovieRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMovieRecommendationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| Optional. Filter by user id, and attach user data. | 
 **parentId** | **optional.String**| Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | **optional.String**| Optional. The fields to return. | 
 **categoryLimit** | **optional.Int32**| The max number of categories to return. | [default to 5]
 **itemLimit** | **optional.Int32**| The max number of items to return per category. | [default to 8]

### Return type

[**[]RecommendationDto**](RecommendationDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

