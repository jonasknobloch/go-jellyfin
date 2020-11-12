# QuickConnectResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | **bool** | Gets a value indicating whether this request is authorized. | [optional] [readonly] 
**Secret** | Pointer to **string** | Gets or sets the secret value used to uniquely identify this request. Can be used to retrieve authentication information. | [optional] 
**Code** | Pointer to **string** | Gets or sets the user facing code used so the user can quickly differentiate this request from others. | [optional] 
**Authentication** | Pointer to **string** | Gets or sets the private access token. | [optional] 
**Error** | Pointer to **string** | Gets or sets an error message. | [optional] 
**DateAdded** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the DateTime that this request was created. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


