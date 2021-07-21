/*
 * Firefly III API
 *
 * This is the official documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. This version of the API is live from version v4.7.9 and onwards. You may use the \"Authorize\" button to try the API below. 
 *
 * API version: 1.4.0
 * Contact: james@firefly-iii.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package firefly

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type TagsApi interface {

	/*
	 * DeleteTag Delete an tag.
	 * Delete an tag.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param tag Either the tag itself or the tag ID.
	 * @return ApiDeleteTagRequest
	 */
	DeleteTag(ctx _context.Context, tag string) ApiDeleteTagRequest

	/*
	 * DeleteTagExecute executes the request
	 */
	DeleteTagExecute(r ApiDeleteTagRequest) (*_nethttp.Response, error)

	/*
	 * GetTag Get a single tag.
	 * Get a single tag.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param tag Either the tag itself or the tag ID.
	 * @return ApiGetTagRequest
	 */
	GetTag(ctx _context.Context, tag string) ApiGetTagRequest

	/*
	 * GetTagExecute executes the request
	 * @return TagSingle
	 */
	GetTagExecute(r ApiGetTagRequest) (TagSingle, *_nethttp.Response, error)

	/*
	 * GetTagCloud Returns a basic tag cloud.
	 * Returns a list of tags, which can be used to draw a basic tag cloud.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiGetTagCloudRequest
	 */
	GetTagCloud(ctx _context.Context) ApiGetTagCloudRequest

	/*
	 * GetTagCloudExecute executes the request
	 * @return TagCloud
	 */
	GetTagCloudExecute(r ApiGetTagCloudRequest) (TagCloud, *_nethttp.Response, error)

	/*
	 * ListAttachmentByTag Lists all attachments.
	 * Lists all attachments.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param tag Either the tag itself or the tag ID.
	 * @return ApiListAttachmentByTagRequest
	 */
	ListAttachmentByTag(ctx _context.Context, tag string) ApiListAttachmentByTagRequest

	/*
	 * ListAttachmentByTagExecute executes the request
	 * @return AttachmentArray
	 */
	ListAttachmentByTagExecute(r ApiListAttachmentByTagRequest) (AttachmentArray, *_nethttp.Response, error)

	/*
	 * ListTag List all tags.
	 * List all of the user's tags.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListTagRequest
	 */
	ListTag(ctx _context.Context) ApiListTagRequest

	/*
	 * ListTagExecute executes the request
	 * @return TagArray
	 */
	ListTagExecute(r ApiListTagRequest) (TagArray, *_nethttp.Response, error)

	/*
	 * ListTransactionByTag List all transactions with this tag.
	 * List all transactions with this tag.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param tag Either the tag itself or the tag ID.
	 * @return ApiListTransactionByTagRequest
	 */
	ListTransactionByTag(ctx _context.Context, tag string) ApiListTransactionByTagRequest

	/*
	 * ListTransactionByTagExecute executes the request
	 * @return TransactionArray
	 */
	ListTransactionByTagExecute(r ApiListTransactionByTagRequest) (TransactionArray, *_nethttp.Response, error)

	/*
	 * StoreTag Store a new tag
	 * Creates a new tag. The data required can be submitted as a JSON body or as a list of parameters.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiStoreTagRequest
	 */
	StoreTag(ctx _context.Context) ApiStoreTagRequest

	/*
	 * StoreTagExecute executes the request
	 * @return TagSingle
	 */
	StoreTagExecute(r ApiStoreTagRequest) (TagSingle, *_nethttp.Response, error)

	/*
	 * UpdateTag Update existing tag.
	 * Update existing tag.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param tag Either the tag itself or the tag ID.
	 * @return ApiUpdateTagRequest
	 */
	UpdateTag(ctx _context.Context, tag string) ApiUpdateTagRequest

	/*
	 * UpdateTagExecute executes the request
	 * @return TagSingle
	 */
	UpdateTagExecute(r ApiUpdateTagRequest) (TagSingle, *_nethttp.Response, error)
}

// TagsApiService TagsApi service
type TagsApiService service

type ApiDeleteTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tag string
}


func (r ApiDeleteTagRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteTagExecute(r)
}

/*
 * DeleteTag Delete an tag.
 * Delete an tag.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tag Either the tag itself or the tag ID.
 * @return ApiDeleteTagRequest
 */
func (a *TagsApiService) DeleteTag(ctx _context.Context, tag string) ApiDeleteTagRequest {
	return ApiDeleteTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

/*
 * Execute executes the request
 */
func (a *TagsApiService) DeleteTagExecute(r ApiDeleteTagRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.DeleteTag")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/{tag}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", _neturl.PathEscape(parameterToString(r.tag, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tag string
	page *int32
}

func (r ApiGetTagRequest) Page(page int32) ApiGetTagRequest {
	r.page = &page
	return r
}

func (r ApiGetTagRequest) Execute() (TagSingle, *_nethttp.Response, error) {
	return r.ApiService.GetTagExecute(r)
}

/*
 * GetTag Get a single tag.
 * Get a single tag.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tag Either the tag itself or the tag ID.
 * @return ApiGetTagRequest
 */
func (a *TagsApiService) GetTag(ctx _context.Context, tag string) ApiGetTagRequest {
	return ApiGetTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

/*
 * Execute executes the request
 * @return TagSingle
 */
func (a *TagsApiService) GetTagExecute(r ApiGetTagRequest) (TagSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TagSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.GetTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/{tag}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", _neturl.PathEscape(parameterToString(r.tag, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetTagCloudRequest struct {
	ctx _context.Context
	ApiService TagsApi
	start *string
	end *string
}

func (r ApiGetTagCloudRequest) Start(start string) ApiGetTagCloudRequest {
	r.start = &start
	return r
}
func (r ApiGetTagCloudRequest) End(end string) ApiGetTagCloudRequest {
	r.end = &end
	return r
}

func (r ApiGetTagCloudRequest) Execute() (TagCloud, *_nethttp.Response, error) {
	return r.ApiService.GetTagCloudExecute(r)
}

/*
 * GetTagCloud Returns a basic tag cloud.
 * Returns a list of tags, which can be used to draw a basic tag cloud.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetTagCloudRequest
 */
func (a *TagsApiService) GetTagCloud(ctx _context.Context) ApiGetTagCloudRequest {
	return ApiGetTagCloudRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TagCloud
 */
func (a *TagsApiService) GetTagCloudExecute(r ApiGetTagCloudRequest) (TagCloud, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TagCloud
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.GetTagCloud")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tag-cloud"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}
	if r.end == nil {
		return localVarReturnValue, nil, reportError("end is required and must be specified")
	}

	localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListAttachmentByTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tag string
	page *int32
}

func (r ApiListAttachmentByTagRequest) Page(page int32) ApiListAttachmentByTagRequest {
	r.page = &page
	return r
}

func (r ApiListAttachmentByTagRequest) Execute() (AttachmentArray, *_nethttp.Response, error) {
	return r.ApiService.ListAttachmentByTagExecute(r)
}

/*
 * ListAttachmentByTag Lists all attachments.
 * Lists all attachments.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tag Either the tag itself or the tag ID.
 * @return ApiListAttachmentByTagRequest
 */
func (a *TagsApiService) ListAttachmentByTag(ctx _context.Context, tag string) ApiListAttachmentByTagRequest {
	return ApiListAttachmentByTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

/*
 * Execute executes the request
 * @return AttachmentArray
 */
func (a *TagsApiService) ListAttachmentByTagExecute(r ApiListAttachmentByTagRequest) (AttachmentArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AttachmentArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.ListAttachmentByTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/{tag}/attachments"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", _neturl.PathEscape(parameterToString(r.tag, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	page *int32
}

func (r ApiListTagRequest) Page(page int32) ApiListTagRequest {
	r.page = &page
	return r
}

func (r ApiListTagRequest) Execute() (TagArray, *_nethttp.Response, error) {
	return r.ApiService.ListTagExecute(r)
}

/*
 * ListTag List all tags.
 * List all of the user's tags.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListTagRequest
 */
func (a *TagsApiService) ListTag(ctx _context.Context) ApiListTagRequest {
	return ApiListTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TagArray
 */
func (a *TagsApiService) ListTagExecute(r ApiListTagRequest) (TagArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TagArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.ListTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListTransactionByTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tag string
	page *int32
	start *string
	end *string
	type_ *TransactionTypeFilter
}

func (r ApiListTransactionByTagRequest) Page(page int32) ApiListTransactionByTagRequest {
	r.page = &page
	return r
}
func (r ApiListTransactionByTagRequest) Start(start string) ApiListTransactionByTagRequest {
	r.start = &start
	return r
}
func (r ApiListTransactionByTagRequest) End(end string) ApiListTransactionByTagRequest {
	r.end = &end
	return r
}
func (r ApiListTransactionByTagRequest) Type_(type_ TransactionTypeFilter) ApiListTransactionByTagRequest {
	r.type_ = &type_
	return r
}

func (r ApiListTransactionByTagRequest) Execute() (TransactionArray, *_nethttp.Response, error) {
	return r.ApiService.ListTransactionByTagExecute(r)
}

/*
 * ListTransactionByTag List all transactions with this tag.
 * List all transactions with this tag.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tag Either the tag itself or the tag ID.
 * @return ApiListTransactionByTagRequest
 */
func (a *TagsApiService) ListTransactionByTag(ctx _context.Context, tag string) ApiListTransactionByTagRequest {
	return ApiListTransactionByTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

/*
 * Execute executes the request
 * @return TransactionArray
 */
func (a *TagsApiService) ListTransactionByTagExecute(r ApiListTransactionByTagRequest) (TransactionArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TransactionArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.ListTransactionByTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/{tag}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", _neturl.PathEscape(parameterToString(r.tag, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiStoreTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tagModel *TagModel
}

func (r ApiStoreTagRequest) TagModel(tagModel TagModel) ApiStoreTagRequest {
	r.tagModel = &tagModel
	return r
}

func (r ApiStoreTagRequest) Execute() (TagSingle, *_nethttp.Response, error) {
	return r.ApiService.StoreTagExecute(r)
}

/*
 * StoreTag Store a new tag
 * Creates a new tag. The data required can be submitted as a JSON body or as a list of parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiStoreTagRequest
 */
func (a *TagsApiService) StoreTag(ctx _context.Context) ApiStoreTagRequest {
	return ApiStoreTagRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TagSingle
 */
func (a *TagsApiService) StoreTagExecute(r ApiStoreTagRequest) (TagSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TagSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.StoreTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tagModel == nil {
		return localVarReturnValue, nil, reportError("tagModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tagModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiUpdateTagRequest struct {
	ctx _context.Context
	ApiService TagsApi
	tag string
	tagModel *TagModel
}

func (r ApiUpdateTagRequest) TagModel(tagModel TagModel) ApiUpdateTagRequest {
	r.tagModel = &tagModel
	return r
}

func (r ApiUpdateTagRequest) Execute() (TagSingle, *_nethttp.Response, error) {
	return r.ApiService.UpdateTagExecute(r)
}

/*
 * UpdateTag Update existing tag.
 * Update existing tag.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tag Either the tag itself or the tag ID.
 * @return ApiUpdateTagRequest
 */
func (a *TagsApiService) UpdateTag(ctx _context.Context, tag string) ApiUpdateTagRequest {
	return ApiUpdateTagRequest{
		ApiService: a,
		ctx: ctx,
		tag: tag,
	}
}

/*
 * Execute executes the request
 * @return TagSingle
 */
func (a *TagsApiService) UpdateTagExecute(r ApiUpdateTagRequest) (TagSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TagSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsApiService.UpdateTag")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tags/{tag}"
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", _neturl.PathEscape(parameterToString(r.tag, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.tagModel == nil {
		return localVarReturnValue, nil, reportError("tagModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tagModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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