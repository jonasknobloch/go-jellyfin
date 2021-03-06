/*
 * Jellyfin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// ArtistsApiService ArtistsApi service
type ArtistsApiService service

// GetAlbumArtistsOpts Optional parameters for the method 'GetAlbumArtists'
type GetAlbumArtistsOpts struct {
    MinCommunityRating optional.Float64
    StartIndex optional.Int32
    Limit optional.Int32
    SearchTerm optional.String
    ParentId optional.String
    Fields optional.String
    ExcludeItemTypes optional.String
    IncludeItemTypes optional.String
    Filters optional.Interface
    IsFavorite optional.Bool
    MediaTypes optional.String
    Genres optional.String
    GenreIds optional.String
    OfficialRatings optional.String
    Tags optional.String
    Years optional.String
    EnableUserData optional.Bool
    ImageTypeLimit optional.Int32
    EnableImageTypes optional.String
    Person optional.String
    PersonIds optional.String
    PersonTypes optional.String
    Studios optional.String
    StudioIds optional.String
    UserId optional.Interface
    NameStartsWithOrGreater optional.String
    NameStartsWith optional.String
    NameLessThan optional.String
    EnableImages optional.Bool
    EnableTotalRecordCount optional.Bool
}

/*
GetAlbumArtists Gets all album artists from a given item, folder, or the entire library.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetAlbumArtistsOpts - Optional Parameters:
 * @param "MinCommunityRating" (optional.Float64) -  Optional filter by minimum community rating.
 * @param "StartIndex" (optional.Int32) -  Optional. The record index to start at. All items with a lower index will be dropped from the results.
 * @param "Limit" (optional.Int32) -  Optional. The maximum number of records to return.
 * @param "SearchTerm" (optional.String) -  Optional. Search term.
 * @param "ParentId" (optional.String) -  Specify this to localize the search to a specific item or folder. Omit to use the root.
 * @param "Fields" (optional.String) -  Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines.
 * @param "ExcludeItemTypes" (optional.String) -  Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited.
 * @param "IncludeItemTypes" (optional.String) -  Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited.
 * @param "Filters" (optional.Interface of []ItemFilter) -  Optional. Specify additional filters to apply.
 * @param "IsFavorite" (optional.Bool) -  Optional filter by items that are marked as favorite, or not.
 * @param "MediaTypes" (optional.String) -  Optional filter by MediaType. Allows multiple, comma delimited.
 * @param "Genres" (optional.String) -  Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited.
 * @param "GenreIds" (optional.String) -  Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited.
 * @param "OfficialRatings" (optional.String) -  Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited.
 * @param "Tags" (optional.String) -  Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited.
 * @param "Years" (optional.String) -  Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited.
 * @param "EnableUserData" (optional.Bool) -  Optional, include user data.
 * @param "ImageTypeLimit" (optional.Int32) -  Optional, the max number of images to return, per image type.
 * @param "EnableImageTypes" (optional.String) -  Optional. The image types to include in the output.
 * @param "Person" (optional.String) -  Optional. If specified, results will be filtered to include only those containing the specified person.
 * @param "PersonIds" (optional.String) -  Optional. If specified, results will be filtered to include only those containing the specified person ids.
 * @param "PersonTypes" (optional.String) -  Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited.
 * @param "Studios" (optional.String) -  Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited.
 * @param "StudioIds" (optional.String) -  Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited.
 * @param "UserId" (optional.Interface of string) -  User id.
 * @param "NameStartsWithOrGreater" (optional.String) -  Optional filter by items whose name is sorted equally or greater than a given input string.
 * @param "NameStartsWith" (optional.String) -  Optional filter by items whose name is sorted equally than a given input string.
 * @param "NameLessThan" (optional.String) -  Optional filter by items whose name is equally or lesser than a given input string.
 * @param "EnableImages" (optional.Bool) -  Optional, include image information in output.
 * @param "EnableTotalRecordCount" (optional.Bool) -  Total record count.
@return BaseItemDtoQueryResult
*/
func (a *ArtistsApiService) GetAlbumArtists(ctx _context.Context, localVarOptionals *GetAlbumArtistsOpts) (BaseItemDtoQueryResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BaseItemDtoQueryResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/Artists/AlbumArtists"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.MinCommunityRating.IsSet() {
		localVarQueryParams.Add("minCommunityRating", parameterToString(localVarOptionals.MinCommunityRating.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartIndex.IsSet() {
		localVarQueryParams.Add("startIndex", parameterToString(localVarOptionals.StartIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentId.IsSet() {
		localVarQueryParams.Add("parentId", parameterToString(localVarOptionals.ParentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeItemTypes.IsSet() {
		localVarQueryParams.Add("excludeItemTypes", parameterToString(localVarOptionals.ExcludeItemTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeItemTypes.IsSet() {
		localVarQueryParams.Add("includeItemTypes", parameterToString(localVarOptionals.IncludeItemTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filters.IsSet() {
		t:=localVarOptionals.Filters.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.IsFavorite.IsSet() {
		localVarQueryParams.Add("isFavorite", parameterToString(localVarOptionals.IsFavorite.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MediaTypes.IsSet() {
		localVarQueryParams.Add("mediaTypes", parameterToString(localVarOptionals.MediaTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Genres.IsSet() {
		localVarQueryParams.Add("genres", parameterToString(localVarOptionals.Genres.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GenreIds.IsSet() {
		localVarQueryParams.Add("genreIds", parameterToString(localVarOptionals.GenreIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OfficialRatings.IsSet() {
		localVarQueryParams.Add("officialRatings", parameterToString(localVarOptionals.OfficialRatings.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		localVarQueryParams.Add("tags", parameterToString(localVarOptionals.Tags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Years.IsSet() {
		localVarQueryParams.Add("years", parameterToString(localVarOptionals.Years.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableUserData.IsSet() {
		localVarQueryParams.Add("enableUserData", parameterToString(localVarOptionals.EnableUserData.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ImageTypeLimit.IsSet() {
		localVarQueryParams.Add("imageTypeLimit", parameterToString(localVarOptionals.ImageTypeLimit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableImageTypes.IsSet() {
		localVarQueryParams.Add("enableImageTypes", parameterToString(localVarOptionals.EnableImageTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Person.IsSet() {
		localVarQueryParams.Add("person", parameterToString(localVarOptionals.Person.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PersonIds.IsSet() {
		localVarQueryParams.Add("personIds", parameterToString(localVarOptionals.PersonIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PersonTypes.IsSet() {
		localVarQueryParams.Add("personTypes", parameterToString(localVarOptionals.PersonTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Studios.IsSet() {
		localVarQueryParams.Add("studios", parameterToString(localVarOptionals.Studios.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StudioIds.IsSet() {
		localVarQueryParams.Add("studioIds", parameterToString(localVarOptionals.StudioIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("userId", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameStartsWithOrGreater.IsSet() {
		localVarQueryParams.Add("nameStartsWithOrGreater", parameterToString(localVarOptionals.NameStartsWithOrGreater.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameStartsWith.IsSet() {
		localVarQueryParams.Add("nameStartsWith", parameterToString(localVarOptionals.NameStartsWith.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameLessThan.IsSet() {
		localVarQueryParams.Add("nameLessThan", parameterToString(localVarOptionals.NameLessThan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableImages.IsSet() {
		localVarQueryParams.Add("enableImages", parameterToString(localVarOptionals.EnableImages.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableTotalRecordCount.IsSet() {
		localVarQueryParams.Add("enableTotalRecordCount", parameterToString(localVarOptionals.EnableTotalRecordCount.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Emby-Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetArtistByNameOpts Optional parameters for the method 'GetArtistByName'
type GetArtistByNameOpts struct {
    UserId optional.Interface
}

/*
GetArtistByName Gets an artist by name.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Studio name.
 * @param optional nil or *GetArtistByNameOpts - Optional Parameters:
 * @param "UserId" (optional.Interface of string) -  Optional. Filter by user id, and attach user data.
@return BaseItemDto
*/
func (a *ArtistsApiService) GetArtistByName(ctx _context.Context, name string, localVarOptionals *GetArtistByNameOpts) (BaseItemDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BaseItemDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/Artists/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.QueryEscape(parameterToString(name, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("userId", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Emby-Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetArtistsOpts Optional parameters for the method 'GetArtists'
type GetArtistsOpts struct {
    MinCommunityRating optional.Float64
    StartIndex optional.Int32
    Limit optional.Int32
    SearchTerm optional.String
    ParentId optional.String
    Fields optional.String
    ExcludeItemTypes optional.String
    IncludeItemTypes optional.String
    Filters optional.Interface
    IsFavorite optional.Bool
    MediaTypes optional.String
    Genres optional.String
    GenreIds optional.String
    OfficialRatings optional.String
    Tags optional.String
    Years optional.String
    EnableUserData optional.Bool
    ImageTypeLimit optional.Int32
    EnableImageTypes optional.String
    Person optional.String
    PersonIds optional.String
    PersonTypes optional.String
    Studios optional.String
    StudioIds optional.String
    UserId optional.Interface
    NameStartsWithOrGreater optional.String
    NameStartsWith optional.String
    NameLessThan optional.String
    EnableImages optional.Bool
    EnableTotalRecordCount optional.Bool
}

/*
GetArtists Gets all artists from a given item, folder, or the entire library.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetArtistsOpts - Optional Parameters:
 * @param "MinCommunityRating" (optional.Float64) -  Optional filter by minimum community rating.
 * @param "StartIndex" (optional.Int32) -  Optional. The record index to start at. All items with a lower index will be dropped from the results.
 * @param "Limit" (optional.Int32) -  Optional. The maximum number of records to return.
 * @param "SearchTerm" (optional.String) -  Optional. Search term.
 * @param "ParentId" (optional.String) -  Specify this to localize the search to a specific item or folder. Omit to use the root.
 * @param "Fields" (optional.String) -  Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines.
 * @param "ExcludeItemTypes" (optional.String) -  Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited.
 * @param "IncludeItemTypes" (optional.String) -  Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited.
 * @param "Filters" (optional.Interface of []ItemFilter) -  Optional. Specify additional filters to apply.
 * @param "IsFavorite" (optional.Bool) -  Optional filter by items that are marked as favorite, or not.
 * @param "MediaTypes" (optional.String) -  Optional filter by MediaType. Allows multiple, comma delimited.
 * @param "Genres" (optional.String) -  Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited.
 * @param "GenreIds" (optional.String) -  Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited.
 * @param "OfficialRatings" (optional.String) -  Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited.
 * @param "Tags" (optional.String) -  Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited.
 * @param "Years" (optional.String) -  Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited.
 * @param "EnableUserData" (optional.Bool) -  Optional, include user data.
 * @param "ImageTypeLimit" (optional.Int32) -  Optional, the max number of images to return, per image type.
 * @param "EnableImageTypes" (optional.String) -  Optional. The image types to include in the output.
 * @param "Person" (optional.String) -  Optional. If specified, results will be filtered to include only those containing the specified person.
 * @param "PersonIds" (optional.String) -  Optional. If specified, results will be filtered to include only those containing the specified person ids.
 * @param "PersonTypes" (optional.String) -  Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited.
 * @param "Studios" (optional.String) -  Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited.
 * @param "StudioIds" (optional.String) -  Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited.
 * @param "UserId" (optional.Interface of string) -  User id.
 * @param "NameStartsWithOrGreater" (optional.String) -  Optional filter by items whose name is sorted equally or greater than a given input string.
 * @param "NameStartsWith" (optional.String) -  Optional filter by items whose name is sorted equally than a given input string.
 * @param "NameLessThan" (optional.String) -  Optional filter by items whose name is equally or lesser than a given input string.
 * @param "EnableImages" (optional.Bool) -  Optional, include image information in output.
 * @param "EnableTotalRecordCount" (optional.Bool) -  Total record count.
@return BaseItemDtoQueryResult
*/
func (a *ArtistsApiService) GetArtists(ctx _context.Context, localVarOptionals *GetArtistsOpts) (BaseItemDtoQueryResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BaseItemDtoQueryResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/Artists"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.MinCommunityRating.IsSet() {
		localVarQueryParams.Add("minCommunityRating", parameterToString(localVarOptionals.MinCommunityRating.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartIndex.IsSet() {
		localVarQueryParams.Add("startIndex", parameterToString(localVarOptionals.StartIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SearchTerm.IsSet() {
		localVarQueryParams.Add("searchTerm", parameterToString(localVarOptionals.SearchTerm.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ParentId.IsSet() {
		localVarQueryParams.Add("parentId", parameterToString(localVarOptionals.ParentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeItemTypes.IsSet() {
		localVarQueryParams.Add("excludeItemTypes", parameterToString(localVarOptionals.ExcludeItemTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeItemTypes.IsSet() {
		localVarQueryParams.Add("includeItemTypes", parameterToString(localVarOptionals.IncludeItemTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filters.IsSet() {
		t:=localVarOptionals.Filters.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filters", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filters", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.IsFavorite.IsSet() {
		localVarQueryParams.Add("isFavorite", parameterToString(localVarOptionals.IsFavorite.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MediaTypes.IsSet() {
		localVarQueryParams.Add("mediaTypes", parameterToString(localVarOptionals.MediaTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Genres.IsSet() {
		localVarQueryParams.Add("genres", parameterToString(localVarOptionals.Genres.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GenreIds.IsSet() {
		localVarQueryParams.Add("genreIds", parameterToString(localVarOptionals.GenreIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OfficialRatings.IsSet() {
		localVarQueryParams.Add("officialRatings", parameterToString(localVarOptionals.OfficialRatings.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		localVarQueryParams.Add("tags", parameterToString(localVarOptionals.Tags.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Years.IsSet() {
		localVarQueryParams.Add("years", parameterToString(localVarOptionals.Years.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableUserData.IsSet() {
		localVarQueryParams.Add("enableUserData", parameterToString(localVarOptionals.EnableUserData.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ImageTypeLimit.IsSet() {
		localVarQueryParams.Add("imageTypeLimit", parameterToString(localVarOptionals.ImageTypeLimit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableImageTypes.IsSet() {
		localVarQueryParams.Add("enableImageTypes", parameterToString(localVarOptionals.EnableImageTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Person.IsSet() {
		localVarQueryParams.Add("person", parameterToString(localVarOptionals.Person.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PersonIds.IsSet() {
		localVarQueryParams.Add("personIds", parameterToString(localVarOptionals.PersonIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PersonTypes.IsSet() {
		localVarQueryParams.Add("personTypes", parameterToString(localVarOptionals.PersonTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Studios.IsSet() {
		localVarQueryParams.Add("studios", parameterToString(localVarOptionals.Studios.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StudioIds.IsSet() {
		localVarQueryParams.Add("studioIds", parameterToString(localVarOptionals.StudioIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarQueryParams.Add("userId", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameStartsWithOrGreater.IsSet() {
		localVarQueryParams.Add("nameStartsWithOrGreater", parameterToString(localVarOptionals.NameStartsWithOrGreater.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameStartsWith.IsSet() {
		localVarQueryParams.Add("nameStartsWith", parameterToString(localVarOptionals.NameStartsWith.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NameLessThan.IsSet() {
		localVarQueryParams.Add("nameLessThan", parameterToString(localVarOptionals.NameLessThan.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableImages.IsSet() {
		localVarQueryParams.Add("enableImages", parameterToString(localVarOptionals.EnableImages.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnableTotalRecordCount.IsSet() {
		localVarQueryParams.Add("enableTotalRecordCount", parameterToString(localVarOptionals.EnableTotalRecordCount.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Emby-Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
