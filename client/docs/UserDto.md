# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**ServerId** | Pointer to **string** | Gets or sets the server identifier. | [optional] 
**ServerName** | Pointer to **string** | Gets or sets the name of the server.  This is not used by the server and is for client-side usage only. | [optional] 
**Id** | **string** | Gets or sets the id. | [optional] 
**PrimaryImageTag** | Pointer to **string** | Gets or sets the primary image tag. | [optional] 
**HasPassword** | **bool** | Gets or sets a value indicating whether this instance has password. | [optional] 
**HasConfiguredPassword** | **bool** | Gets or sets a value indicating whether this instance has configured password. | [optional] 
**HasConfiguredEasyPassword** | **bool** | Gets or sets a value indicating whether this instance has configured easy password. | [optional] 
**EnableAutoLogin** | Pointer to **bool** | Gets or sets whether async login is enabled or not. | [optional] 
**LastLoginDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the last login date. | [optional] 
**LastActivityDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the last activity date. | [optional] 
**Configuration** | [**UserConfiguration**](UserConfiguration.md) |  | [optional] 
**Policy** | [**UserPolicy**](UserPolicy.md) |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **float64** | Gets or sets the primary image aspect ratio. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


