# \LocalizationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](LocalizationApi.md#GetCountries) | **Get** /Localization/Countries | Gets known countries.
[**GetCultures**](LocalizationApi.md#GetCultures) | **Get** /Localization/Cultures | Gets known cultures.
[**GetLocalizationOptions**](LocalizationApi.md#GetLocalizationOptions) | **Get** /Localization/Options | Gets localization options.
[**GetParentalRatings**](LocalizationApi.md#GetParentalRatings) | **Get** /Localization/ParentalRatings | Gets known parental ratings.



## GetCountries

> []CountryInfo GetCountries(ctx, )

Gets known countries.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CountryInfo**](CountryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCultures

> []CultureDto GetCultures(ctx, )

Gets known cultures.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CultureDto**](CultureDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationOptions

> []LocalizationOption GetLocalizationOptions(ctx, )

Gets localization options.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]LocalizationOption**](LocalizationOption.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentalRatings

> []ParentalRating GetParentalRatings(ctx, )

Gets known parental ratings.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ParentalRating**](ParentalRating.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

