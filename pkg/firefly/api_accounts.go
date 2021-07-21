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

type AccountsApi interface {

	/*
	 * DeleteAccount Permanently delete account.
	 * Will permanently delete an account. Any associated transactions and piggy banks are ALSO deleted. Cannot be recovered from.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiDeleteAccountRequest
	 */
	DeleteAccount(ctx _context.Context, id int32) ApiDeleteAccountRequest

	/*
	 * DeleteAccountExecute executes the request
	 */
	DeleteAccountExecute(r ApiDeleteAccountRequest) (*_nethttp.Response, error)

	/*
	 * GetAccount Get single account.
	 * Returns a single account by its ID.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiGetAccountRequest
	 */
	GetAccount(ctx _context.Context, id int32) ApiGetAccountRequest

	/*
	 * GetAccountExecute executes the request
	 * @return AccountSingle
	 */
	GetAccountExecute(r ApiGetAccountRequest) (AccountSingle, *_nethttp.Response, error)

	/*
	 * ListAccount List all accounts.
	 * This endpoint returns a list of all the accounts owned by the authenticated user.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListAccountRequest
	 */
	ListAccount(ctx _context.Context) ApiListAccountRequest

	/*
	 * ListAccountExecute executes the request
	 * @return AccountArray
	 */
	ListAccountExecute(r ApiListAccountRequest) (AccountArray, *_nethttp.Response, error)

	/*
	 * ListAttachmentByAccount Lists all attachments.
	 * Lists all attachments.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiListAttachmentByAccountRequest
	 */
	ListAttachmentByAccount(ctx _context.Context, id int32) ApiListAttachmentByAccountRequest

	/*
	 * ListAttachmentByAccountExecute executes the request
	 * @return AttachmentArray
	 */
	ListAttachmentByAccountExecute(r ApiListAttachmentByAccountRequest) (AttachmentArray, *_nethttp.Response, error)

	/*
	 * ListPiggyBankByAccount List all piggy banks related to the account.
	 * This endpoint returns a list of all the piggy banks connected to the account.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiListPiggyBankByAccountRequest
	 */
	ListPiggyBankByAccount(ctx _context.Context, id int32) ApiListPiggyBankByAccountRequest

	/*
	 * ListPiggyBankByAccountExecute executes the request
	 * @return PiggyBankArray
	 */
	ListPiggyBankByAccountExecute(r ApiListPiggyBankByAccountRequest) (PiggyBankArray, *_nethttp.Response, error)

	/*
	 * ListTransactionByAccount List all transactions related to the account.
	 * This endpoint returns a list of all the transactions connected to the account.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiListTransactionByAccountRequest
	 */
	ListTransactionByAccount(ctx _context.Context, id int32) ApiListTransactionByAccountRequest

	/*
	 * ListTransactionByAccountExecute executes the request
	 * @return TransactionArray
	 */
	ListTransactionByAccountExecute(r ApiListTransactionByAccountRequest) (TransactionArray, *_nethttp.Response, error)

	/*
	 * StoreAccount Create new account.
	 * Creates a new account. The data required can be submitted as a JSON body or as a list of parameters (in key=value pairs, like a webform).
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiStoreAccountRequest
	 */
	StoreAccount(ctx _context.Context) ApiStoreAccountRequest

	/*
	 * StoreAccountExecute executes the request
	 * @return AccountSingle
	 */
	StoreAccountExecute(r ApiStoreAccountRequest) (AccountSingle, *_nethttp.Response, error)

	/*
	 * UpdateAccount Update existing account.
	 * Used to update a single account. All fields that are not submitted will be cleared (set to NULL). The model will tell you which fields are mandatory.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the account.
	 * @return ApiUpdateAccountRequest
	 */
	UpdateAccount(ctx _context.Context, id int32) ApiUpdateAccountRequest

	/*
	 * UpdateAccountExecute executes the request
	 * @return AccountSingle
	 */
	UpdateAccountExecute(r ApiUpdateAccountRequest) (AccountSingle, *_nethttp.Response, error)
}

// AccountsApiService AccountsApi service
type AccountsApiService service

type ApiDeleteAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
}


func (r ApiDeleteAccountRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteAccountExecute(r)
}

/*
 * DeleteAccount Permanently delete account.
 * Will permanently delete an account. Any associated transactions and piggy banks are ALSO deleted. Cannot be recovered from.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiDeleteAccountRequest
 */
func (a *AccountsApiService) DeleteAccount(ctx _context.Context, id int32) ApiDeleteAccountRequest {
	return ApiDeleteAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *AccountsApiService) DeleteAccountExecute(r ApiDeleteAccountRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.DeleteAccount")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiGetAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
	date *string
}

func (r ApiGetAccountRequest) Date(date string) ApiGetAccountRequest {
	r.date = &date
	return r
}

func (r ApiGetAccountRequest) Execute() (AccountSingle, *_nethttp.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
 * GetAccount Get single account.
 * Returns a single account by its ID.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiGetAccountRequest
 */
func (a *AccountsApiService) GetAccount(ctx _context.Context, id int32) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AccountSingle
 */
func (a *AccountsApiService) GetAccountExecute(r ApiGetAccountRequest) (AccountSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.date != nil {
		localVarQueryParams.Add("date", parameterToString(*r.date, ""))
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

type ApiListAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	page *int32
	date *string
	type_ *AccountTypeFilter
}

func (r ApiListAccountRequest) Page(page int32) ApiListAccountRequest {
	r.page = &page
	return r
}
func (r ApiListAccountRequest) Date(date string) ApiListAccountRequest {
	r.date = &date
	return r
}
func (r ApiListAccountRequest) Type_(type_ AccountTypeFilter) ApiListAccountRequest {
	r.type_ = &type_
	return r
}

func (r ApiListAccountRequest) Execute() (AccountArray, *_nethttp.Response, error) {
	return r.ApiService.ListAccountExecute(r)
}

/*
 * ListAccount List all accounts.
 * This endpoint returns a list of all the accounts owned by the authenticated user.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListAccountRequest
 */
func (a *AccountsApiService) ListAccount(ctx _context.Context) ApiListAccountRequest {
	return ApiListAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AccountArray
 */
func (a *AccountsApiService) ListAccountExecute(r ApiListAccountRequest) (AccountArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.date != nil {
		localVarQueryParams.Add("date", parameterToString(*r.date, ""))
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

type ApiListAttachmentByAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
	page *int32
}

func (r ApiListAttachmentByAccountRequest) Page(page int32) ApiListAttachmentByAccountRequest {
	r.page = &page
	return r
}

func (r ApiListAttachmentByAccountRequest) Execute() (AttachmentArray, *_nethttp.Response, error) {
	return r.ApiService.ListAttachmentByAccountExecute(r)
}

/*
 * ListAttachmentByAccount Lists all attachments.
 * Lists all attachments.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiListAttachmentByAccountRequest
 */
func (a *AccountsApiService) ListAttachmentByAccount(ctx _context.Context, id int32) ApiListAttachmentByAccountRequest {
	return ApiListAttachmentByAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AttachmentArray
 */
func (a *AccountsApiService) ListAttachmentByAccountExecute(r ApiListAttachmentByAccountRequest) (AttachmentArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AttachmentArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListAttachmentByAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}/attachments"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiListPiggyBankByAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
	page *int32
}

func (r ApiListPiggyBankByAccountRequest) Page(page int32) ApiListPiggyBankByAccountRequest {
	r.page = &page
	return r
}

func (r ApiListPiggyBankByAccountRequest) Execute() (PiggyBankArray, *_nethttp.Response, error) {
	return r.ApiService.ListPiggyBankByAccountExecute(r)
}

/*
 * ListPiggyBankByAccount List all piggy banks related to the account.
 * This endpoint returns a list of all the piggy banks connected to the account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiListPiggyBankByAccountRequest
 */
func (a *AccountsApiService) ListPiggyBankByAccount(ctx _context.Context, id int32) ApiListPiggyBankByAccountRequest {
	return ApiListPiggyBankByAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return PiggyBankArray
 */
func (a *AccountsApiService) ListPiggyBankByAccountExecute(r ApiListPiggyBankByAccountRequest) (PiggyBankArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PiggyBankArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListPiggyBankByAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}/piggy_banks"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiListTransactionByAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
	page *int32
	limit *int32
	start *string
	end *string
	type_ *TransactionTypeFilter
}

func (r ApiListTransactionByAccountRequest) Page(page int32) ApiListTransactionByAccountRequest {
	r.page = &page
	return r
}
func (r ApiListTransactionByAccountRequest) Limit(limit int32) ApiListTransactionByAccountRequest {
	r.limit = &limit
	return r
}
func (r ApiListTransactionByAccountRequest) Start(start string) ApiListTransactionByAccountRequest {
	r.start = &start
	return r
}
func (r ApiListTransactionByAccountRequest) End(end string) ApiListTransactionByAccountRequest {
	r.end = &end
	return r
}
func (r ApiListTransactionByAccountRequest) Type_(type_ TransactionTypeFilter) ApiListTransactionByAccountRequest {
	r.type_ = &type_
	return r
}

func (r ApiListTransactionByAccountRequest) Execute() (TransactionArray, *_nethttp.Response, error) {
	return r.ApiService.ListTransactionByAccountExecute(r)
}

/*
 * ListTransactionByAccount List all transactions related to the account.
 * This endpoint returns a list of all the transactions connected to the account.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiListTransactionByAccountRequest
 */
func (a *AccountsApiService) ListTransactionByAccount(ctx _context.Context, id int32) ApiListTransactionByAccountRequest {
	return ApiListTransactionByAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return TransactionArray
 */
func (a *AccountsApiService) ListTransactionByAccountExecute(r ApiListTransactionByAccountRequest) (TransactionArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TransactionArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListTransactionByAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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

type ApiStoreAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	account *Account
}

func (r ApiStoreAccountRequest) Account(account Account) ApiStoreAccountRequest {
	r.account = &account
	return r
}

func (r ApiStoreAccountRequest) Execute() (AccountSingle, *_nethttp.Response, error) {
	return r.ApiService.StoreAccountExecute(r)
}

/*
 * StoreAccount Create new account.
 * Creates a new account. The data required can be submitted as a JSON body or as a list of parameters (in key=value pairs, like a webform).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiStoreAccountRequest
 */
func (a *AccountsApiService) StoreAccount(ctx _context.Context) ApiStoreAccountRequest {
	return ApiStoreAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AccountSingle
 */
func (a *AccountsApiService) StoreAccountExecute(r ApiStoreAccountRequest) (AccountSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.StoreAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.account == nil {
		return localVarReturnValue, nil, reportError("account is required and must be specified")
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
	localVarPostBody = r.account
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

type ApiUpdateAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	id int32
	account *Account
}

func (r ApiUpdateAccountRequest) Account(account Account) ApiUpdateAccountRequest {
	r.account = &account
	return r
}

func (r ApiUpdateAccountRequest) Execute() (AccountSingle, *_nethttp.Response, error) {
	return r.ApiService.UpdateAccountExecute(r)
}

/*
 * UpdateAccount Update existing account.
 * Used to update a single account. All fields that are not submitted will be cleared (set to NULL). The model will tell you which fields are mandatory.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the account.
 * @return ApiUpdateAccountRequest
 */
func (a *AccountsApiService) UpdateAccount(ctx _context.Context, id int32) ApiUpdateAccountRequest {
	return ApiUpdateAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return AccountSingle
 */
func (a *AccountsApiService) UpdateAccountExecute(r ApiUpdateAccountRequest) (AccountSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.UpdateAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.account == nil {
		return localVarReturnValue, nil, reportError("account is required and must be specified")
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
	localVarPostBody = r.account
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
