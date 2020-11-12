# \TrailersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrailers**](TrailersApi.md#GetTrailers) | **Get** /Trailers | Finds movies and trailers similar to a given trailer.



## GetTrailers

> BaseItemDtoQueryResult GetTrailers(ctx, optional)

Finds movies and trailers similar to a given trailer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTrailersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrailersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of string**](.md)| The user id. | 
 **maxOfficialRating** | **optional.String**| Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **optional.Bool**| Optional filter by items with theme songs. | 
 **hasThemeVideo** | **optional.Bool**| Optional filter by items with theme videos. | 
 **hasSubtitles** | **optional.Bool**| Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **optional.Bool**| Optional filter by items with special features. | 
 **hasTrailer** | **optional.Bool**| Optional filter by items with trailers. | 
 **adjacentTo** | **optional.String**| Optional. Return items that are siblings of a supplied item. | 
 **parentIndexNumber** | **optional.Int32**| Optional filter by parent index number. | 
 **hasParentalRating** | **optional.Bool**| Optional filter by items that have or do not have a parental rating. | 
 **isHd** | **optional.Bool**| Optional filter by items that are HD or not. | 
 **is4K** | **optional.Bool**| Optional filter by items that are 4K or not. | 
 **locationTypes** | **optional.String**| Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimeted. | 
 **excludeLocationTypes** | [**optional.Interface of []LocationType**](LocationType.md)| Optional. If specified, results will be filtered based on the LocationType. This allows multiple, comma delimeted. | 
 **isMissing** | **optional.Bool**| Optional filter by items that are missing episodes or not. | 
 **isUnaired** | **optional.Bool**| Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **optional.Float64**| Optional filter by minimum community rating. | 
 **minCriticRating** | **optional.Float64**| Optional filter by minimum critic rating. | 
 **minPremiereDate** | **optional.Time**| Optional. The minimum premiere date. Format &#x3D; ISO. | 
 **minDateLastSaved** | **optional.Time**| Optional. The minimum last saved date. Format &#x3D; ISO. | 
 **minDateLastSavedForUser** | **optional.Time**| Optional. The minimum last saved date for the current user. Format &#x3D; ISO. | 
 **maxPremiereDate** | **optional.Time**| Optional. The maximum premiere date. Format &#x3D; ISO. | 
 **hasOverview** | **optional.Bool**| Optional filter by items that have an overview or not. | 
 **hasImdbId** | **optional.Bool**| Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **optional.Bool**| Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **optional.Bool**| Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **optional.String**| Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **optional.Int32**| Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **optional.Int32**| Optional. The maximum number of records to return. | 
 **recursive** | **optional.Bool**| When searching within folders, this determines whether or not the search will be recursive. true/false. | 
 **searchTerm** | **optional.String**| Optional. Filter based on a search term. | 
 **sortOrder** | **optional.String**| Sort Order - Ascending,Descending. | 
 **parentId** | **optional.String**| Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | **optional.String**| Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **excludeItemTypes** | **optional.String**| Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **filters** | [**optional.Interface of []ItemFilter**](ItemFilter.md)| Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes. | 
 **isFavorite** | **optional.Bool**| Optional filter by items that are marked as favorite, or not. | 
 **mediaTypes** | **optional.String**| Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **optional.String**| Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **optional.String**| Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **isPlayed** | **optional.Bool**| Optional filter by items that are played, or not. | 
 **genres** | **optional.String**| Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **optional.String**| Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **optional.String**| Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **optional.String**| Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableUserData** | **optional.Bool**| Optional, include user data. | 
 **imageTypeLimit** | **optional.Int32**| Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | **optional.String**| Optional. The image types to include in the output. | 
 **person** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified person id. | 
 **personTypes** | **optional.String**| Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. | 
 **studios** | **optional.String**| Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **optional.String**| Optional. If specified, results will be filtered based on artists. This allows multiple, pipe delimeted. | 
 **excludeArtistIds** | **optional.String**| Optional. If specified, results will be filtered based on artist id. This allows multiple, pipe delimeted. | 
 **artistIds** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified artist id. | 
 **albumArtistIds** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified album artist id. | 
 **contributingArtistIds** | **optional.String**| Optional. If specified, results will be filtered to include only those containing the specified contributing artist id. | 
 **albums** | **optional.String**| Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **albumIds** | **optional.String**| Optional. If specified, results will be filtered based on album id. This allows multiple, pipe delimeted. | 
 **ids** | **optional.String**| Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **optional.String**| Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **minOfficialRating** | **optional.String**| Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **optional.Bool**| Optional filter by items that are locked. | 
 **isPlaceHolder** | **optional.Bool**| Optional filter by items that are placeholders. | 
 **hasOfficialRating** | **optional.Bool**| Optional filter by items that have official ratings. | 
 **collapseBoxSetItems** | **optional.Bool**| Whether or not to hide items behind their boxsets. | 
 **minWidth** | **optional.Int32**| Optional. Filter by the minimum width of the item. | 
 **minHeight** | **optional.Int32**| Optional. Filter by the minimum height of the item. | 
 **maxWidth** | **optional.Int32**| Optional. Filter by the maximum width of the item. | 
 **maxHeight** | **optional.Int32**| Optional. Filter by the maximum height of the item. | 
 **is3D** | **optional.Bool**| Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **optional.String**| Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **optional.String**| Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **optional.String**| Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **optional.String**| Optional filter by items whose name is equally or lesser than a given input string. | 
 **studioIds** | **optional.String**| Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimeted. | 
 **genreIds** | **optional.String**| Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimeted. | 
 **enableTotalRecordCount** | **optional.Bool**| Optional. Enable the total record count. | [default to true]
 **enableImages** | **optional.Bool**| Optional, include image information in output. | [default to true]

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

