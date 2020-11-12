# \PackageApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPackageInstallation**](PackageApi.md#CancelPackageInstallation) | **Delete** /Packages/Installing/{packageId} | Cancels a package installation.
[**GetPackageInfo**](PackageApi.md#GetPackageInfo) | **Get** /Packages/{name} | Gets a package by name or assembly GUID.
[**GetPackages**](PackageApi.md#GetPackages) | **Get** /Packages | Gets available packages.
[**GetRepositories**](PackageApi.md#GetRepositories) | **Get** /Repositories | Gets all package repositories.
[**InstallPackage**](PackageApi.md#InstallPackage) | **Post** /Packages/Installed/{name} | Installs a package.
[**SetRepositories**](PackageApi.md#SetRepositories) | **Options** /Repositories | Sets the enabled and existing package repositories.



## CancelPackageInstallation

> CancelPackageInstallation(ctx, packageId)

Cancels a package installation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | [**string**](.md)| Installation Id. | 

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


## GetPackageInfo

> PackageInfo GetPackageInfo(ctx, name, optional)

Gets a package by name or assembly GUID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| The name of the package. | 
 **optional** | ***GetPackageInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPackageInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **optional.String**| The GUID of the associated assembly. | 

### Return type

[**PackageInfo**](PackageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackages

> []PackageInfo GetPackages(ctx, )

Gets available packages.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PackageInfo**](PackageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> []RepositoryInfo GetRepositories(ctx, )

Gets all package repositories.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RepositoryInfo**](RepositoryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallPackage

> InstallPackage(ctx, name, optional)

Installs a package.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Package name. | 
 **optional** | ***InstallPackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InstallPackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **optional.String**| GUID of the associated assembly. | 
 **version** | **optional.String**| Optional version. Defaults to latest version. | 
 **repositoryUrl** | **optional.String**| Optional. Specify the repository to install from. | 

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


## SetRepositories

> SetRepositories(ctx, optional)

Sets the enabled and existing package repositories.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetRepositoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryInfo** | [**optional.Interface of []RepositoryInfo**](RepositoryInfo.md)| The list of package repositories. | 

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

