# GetProgramsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelIds** | Pointer to **string** | Gets or sets the channels to return guide information for. | [optional] 
**UserId** | **string** | Gets or sets optional. Filter by user id. | [optional] 
**MinStartDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the minimum premiere start date.  Optional. | [optional] 
**HasAired** | Pointer to **bool** | Gets or sets filter by programs that have completed airing, or not.  Optional. | [optional] 
**IsAiring** | Pointer to **bool** | Gets or sets filter by programs that are currently airing, or not.  Optional. | [optional] 
**MaxStartDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the maximum premiere start date.  Optional. | [optional] 
**MinEndDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the minimum premiere end date.  Optional. | [optional] 
**MaxEndDate** | Pointer to [**time.Time**](time.Time.md) | Gets or sets the maximum premiere end date.  Optional. | [optional] 
**IsMovie** | Pointer to **bool** | Gets or sets filter for movies.  Optional. | [optional] 
**IsSeries** | Pointer to **bool** | Gets or sets filter for series.  Optional. | [optional] 
**IsNews** | Pointer to **bool** | Gets or sets filter for news.  Optional. | [optional] 
**IsKids** | Pointer to **bool** | Gets or sets filter for kids.  Optional. | [optional] 
**IsSports** | Pointer to **bool** | Gets or sets filter for sports.  Optional. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the record index to start at. All items with a lower index will be dropped from the results.  Optional. | [optional] 
**Limit** | Pointer to **int32** | Gets or sets the maximum number of records to return.  Optional. | [optional] 
**SortBy** | Pointer to **string** | Gets or sets specify one or more sort orders, comma delimited. Options: Name, StartDate.  Optional. | [optional] 
**SortOrder** | Pointer to **string** | Gets or sets sort Order - Ascending,Descending. | [optional] 
**Genres** | Pointer to **string** | Gets or sets the genres to return guide information for. | [optional] 
**GenreIds** | Pointer to **string** | Gets or sets the genre ids to return guide information for. | [optional] 
**EnableImages** | Pointer to **bool** | Gets or sets include image information in output.  Optional. | [optional] 
**EnableTotalRecordCount** | **bool** | Gets or sets a value indicating whether retrieve total record count. | [optional] 
**ImageTypeLimit** | Pointer to **int32** | Gets or sets the max number of images to return, per image type.  Optional. | [optional] 
**EnableImageTypes** | Pointer to **string** | Gets or sets the image types to include in the output.  Optional. | [optional] 
**EnableUserData** | Pointer to **bool** | Gets or sets include user data.  Optional. | [optional] 
**SeriesTimerId** | Pointer to **string** | Gets or sets filter by series timer id.  Optional. | [optional] 
**LibrarySeriesId** | **string** | Gets or sets filter by library series id.  Optional. | [optional] 
**Fields** | Pointer to **string** | Gets or sets specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines.  Optional. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


