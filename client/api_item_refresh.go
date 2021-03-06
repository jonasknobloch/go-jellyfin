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
)

// Linger please
var (
	_ _context.Context
)

// ItemRefreshApiService ItemRefreshApi service
type ItemRefreshApiService service

// PostOpts Optional parameters for the method 'Post'
type PostOpts struct {
    MetadataRefreshMode optional.Interface
    ImageRefreshMode optional.Interface
    ReplaceAllMetadata optional.Bool
    ReplaceAllImages optional.Bool
}

/*
Post Refreshes metadata for an item.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param itemId Item id.
 * @param optional nil or *PostOpts - Optional Parameters:
 * @param "MetadataRefreshMode" (optional.Interface of MetadataRefreshMode) -  (Optional) Specifies the metadata refresh mode.
 * @param "ImageRefreshMode" (optional.Interface of MetadataRefreshMode) -  (Optional) Specifies the image refresh mode.
 * @param "ReplaceAllMetadata" (optional.Bool) -  (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh.
 * @param "ReplaceAllImages" (optional.Bool) -  (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh.
*/
func (a *ItemRefreshApiService) Post(ctx _context.Context, itemId string, localVarOptionals *PostOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/Items/{itemId}/Refresh"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", _neturl.QueryEscape(parameterToString(itemId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.MetadataRefreshMode.IsSet() {
		localVarQueryParams.Add("metadataRefreshMode", parameterToString(localVarOptionals.MetadataRefreshMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ImageRefreshMode.IsSet() {
		localVarQueryParams.Add("imageRefreshMode", parameterToString(localVarOptionals.ImageRefreshMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReplaceAllMetadata.IsSet() {
		localVarQueryParams.Add("replaceAllMetadata", parameterToString(localVarOptionals.ReplaceAllMetadata.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ReplaceAllImages.IsSet() {
		localVarQueryParams.Add("replaceAllImages", parameterToString(localVarOptionals.ReplaceAllImages.Value(), ""))
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
