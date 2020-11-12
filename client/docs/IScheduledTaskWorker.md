# IScheduledTaskWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledTask** | [**IScheduledTask**](IScheduledTask.md) |  | [optional] 
**LastExecutionResult** | [**TaskResult**](TaskResult.md) |  | [optional] 
**Name** | Pointer to **string** | Gets the name. | [optional] [readonly] 
**Description** | Pointer to **string** | Gets the description. | [optional] [readonly] 
**Category** | Pointer to **string** | Gets the category. | [optional] [readonly] 
**State** | [**TaskState**](TaskState.md) |  | [optional] 
**CurrentProgress** | Pointer to **float64** | Gets the current progress. | [optional] [readonly] 
**Triggers** | Pointer to [**[]TaskTriggerInfo**](TaskTriggerInfo.md) | Gets the triggers that define when the task will run. | [optional] 
**Id** | Pointer to **string** | Gets the unique id. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


