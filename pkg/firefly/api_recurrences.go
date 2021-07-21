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

type RecurrencesApi interface {

	/*
	 * DeleteRecurrence Delete a recurring transaction.
	 * Delete a recurring transaction. Transactions created will not be deleted.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the recurring transaction.
	 * @return ApiDeleteRecurrenceRequest
	 */
	DeleteRecurrence(ctx _context.Context, id int32) ApiDeleteRecurrenceRequest

	/*
	 * DeleteRecurrenceExecute executes the request
	 */
	DeleteRecurrenceExecute(r ApiDeleteRecurrenceRequest) (*_nethttp.Response, error)

	/*
	 * GetRecurrence Get a single recurring transaction.
	 * Get a single recurring transaction.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the recurring transaction.
	 * @return ApiGetRecurrenceRequest
	 */
	GetRecurrence(ctx _context.Context, id int32) ApiGetRecurrenceRequest

	/*
	 * GetRecurrenceExecute executes the request
	 * @return RecurrenceSingle
	 */
	GetRecurrenceExecute(r ApiGetRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error)

	/*
	 * ListRecurrence List all recurring transactions.
	 * List all recurring transactions.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiListRecurrenceRequest
	 */
	ListRecurrence(ctx _context.Context) ApiListRecurrenceRequest

	/*
	 * ListRecurrenceExecute executes the request
	 * @return RecurrenceArray
	 */
	ListRecurrenceExecute(r ApiListRecurrenceRequest) (RecurrenceArray, *_nethttp.Response, error)

	/*
	 * ListTransactionByRecurrence List all transactions created by a recurring transaction.
	 * List all transactions created by a recurring transaction, optionally limited to the date ranges specified.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the recurring transaction.
	 * @return ApiListTransactionByRecurrenceRequest
	 */
	ListTransactionByRecurrence(ctx _context.Context, id int32) ApiListTransactionByRecurrenceRequest

	/*
	 * ListTransactionByRecurrenceExecute executes the request
	 * @return TransactionArray
	 */
	ListTransactionByRecurrenceExecute(r ApiListTransactionByRecurrenceRequest) (TransactionArray, *_nethttp.Response, error)

	/*
	 * StoreRecurrence Store a new recurring transaction
	 * Creates a new recurring transaction. The data required can be submitted as a JSON body or as a list of parameters.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiStoreRecurrenceRequest
	 */
	StoreRecurrence(ctx _context.Context) ApiStoreRecurrenceRequest

	/*
	 * StoreRecurrenceExecute executes the request
	 * @return RecurrenceSingle
	 */
	StoreRecurrenceExecute(r ApiStoreRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error)

	/*
	 * TriggerRecurrence Trigger the creation of recurring transactions (like a cron job).
	 * Triggers the recurring transactions, like a cron job would. If the schedule does not call for a new transaction to be created, nothing will happen.

	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiTriggerRecurrenceRequest
	 */
	TriggerRecurrence(ctx _context.Context) ApiTriggerRecurrenceRequest

	/*
	 * TriggerRecurrenceExecute executes the request
	 */
	TriggerRecurrenceExecute(r ApiTriggerRecurrenceRequest) (*_nethttp.Response, error)

	/*
	 * UpdateRecurrence Update existing recurring transaction.
	 * Update existing recurring transaction.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id The ID of the recurring transaction.
	 * @return ApiUpdateRecurrenceRequest
	 */
	UpdateRecurrence(ctx _context.Context, id int32) ApiUpdateRecurrenceRequest

	/*
	 * UpdateRecurrenceExecute executes the request
	 * @return RecurrenceSingle
	 */
	UpdateRecurrenceExecute(r ApiUpdateRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error)
}

// RecurrencesApiService RecurrencesApi service
type RecurrencesApiService service

type ApiDeleteRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	id int32
}


func (r ApiDeleteRecurrenceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteRecurrenceExecute(r)
}

/*
 * DeleteRecurrence Delete a recurring transaction.
 * Delete a recurring transaction. Transactions created will not be deleted.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the recurring transaction.
 * @return ApiDeleteRecurrenceRequest
 */
func (a *RecurrencesApiService) DeleteRecurrence(ctx _context.Context, id int32) ApiDeleteRecurrenceRequest {
	return ApiDeleteRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *RecurrencesApiService) DeleteRecurrenceExecute(r ApiDeleteRecurrenceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.DeleteRecurrence")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences/{id}"
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

type ApiGetRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	id int32
}


func (r ApiGetRecurrenceRequest) Execute() (RecurrenceSingle, *_nethttp.Response, error) {
	return r.ApiService.GetRecurrenceExecute(r)
}

/*
 * GetRecurrence Get a single recurring transaction.
 * Get a single recurring transaction.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the recurring transaction.
 * @return ApiGetRecurrenceRequest
 */
func (a *RecurrencesApiService) GetRecurrence(ctx _context.Context, id int32) ApiGetRecurrenceRequest {
	return ApiGetRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return RecurrenceSingle
 */
func (a *RecurrencesApiService) GetRecurrenceExecute(r ApiGetRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RecurrenceSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.GetRecurrence")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences/{id}"
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

type ApiListRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	page *int32
}

func (r ApiListRecurrenceRequest) Page(page int32) ApiListRecurrenceRequest {
	r.page = &page
	return r
}

func (r ApiListRecurrenceRequest) Execute() (RecurrenceArray, *_nethttp.Response, error) {
	return r.ApiService.ListRecurrenceExecute(r)
}

/*
 * ListRecurrence List all recurring transactions.
 * List all recurring transactions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListRecurrenceRequest
 */
func (a *RecurrencesApiService) ListRecurrence(ctx _context.Context) ApiListRecurrenceRequest {
	return ApiListRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RecurrenceArray
 */
func (a *RecurrencesApiService) ListRecurrenceExecute(r ApiListRecurrenceRequest) (RecurrenceArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RecurrenceArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.ListRecurrence")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences"

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

type ApiListTransactionByRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	id int32
	page *int32
	start *string
	end *string
	type_ *TransactionTypeFilter
}

func (r ApiListTransactionByRecurrenceRequest) Page(page int32) ApiListTransactionByRecurrenceRequest {
	r.page = &page
	return r
}
func (r ApiListTransactionByRecurrenceRequest) Start(start string) ApiListTransactionByRecurrenceRequest {
	r.start = &start
	return r
}
func (r ApiListTransactionByRecurrenceRequest) End(end string) ApiListTransactionByRecurrenceRequest {
	r.end = &end
	return r
}
func (r ApiListTransactionByRecurrenceRequest) Type_(type_ TransactionTypeFilter) ApiListTransactionByRecurrenceRequest {
	r.type_ = &type_
	return r
}

func (r ApiListTransactionByRecurrenceRequest) Execute() (TransactionArray, *_nethttp.Response, error) {
	return r.ApiService.ListTransactionByRecurrenceExecute(r)
}

/*
 * ListTransactionByRecurrence List all transactions created by a recurring transaction.
 * List all transactions created by a recurring transaction, optionally limited to the date ranges specified.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the recurring transaction.
 * @return ApiListTransactionByRecurrenceRequest
 */
func (a *RecurrencesApiService) ListTransactionByRecurrence(ctx _context.Context, id int32) ApiListTransactionByRecurrenceRequest {
	return ApiListTransactionByRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return TransactionArray
 */
func (a *RecurrencesApiService) ListTransactionByRecurrenceExecute(r ApiListTransactionByRecurrenceRequest) (TransactionArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TransactionArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.ListTransactionByRecurrence")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences/{id}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiStoreRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	recurrence *Recurrence
}

func (r ApiStoreRecurrenceRequest) Recurrence(recurrence Recurrence) ApiStoreRecurrenceRequest {
	r.recurrence = &recurrence
	return r
}

func (r ApiStoreRecurrenceRequest) Execute() (RecurrenceSingle, *_nethttp.Response, error) {
	return r.ApiService.StoreRecurrenceExecute(r)
}

/*
 * StoreRecurrence Store a new recurring transaction
 * Creates a new recurring transaction. The data required can be submitted as a JSON body or as a list of parameters.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiStoreRecurrenceRequest
 */
func (a *RecurrencesApiService) StoreRecurrence(ctx _context.Context) ApiStoreRecurrenceRequest {
	return ApiStoreRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RecurrenceSingle
 */
func (a *RecurrencesApiService) StoreRecurrenceExecute(r ApiStoreRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RecurrenceSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.StoreRecurrence")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.recurrence == nil {
		return localVarReturnValue, nil, reportError("recurrence is required and must be specified")
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
	localVarPostBody = r.recurrence
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

type ApiTriggerRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
}


func (r ApiTriggerRecurrenceRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TriggerRecurrenceExecute(r)
}

/*
 * TriggerRecurrence Trigger the creation of recurring transactions (like a cron job).
 * Triggers the recurring transactions, like a cron job would. If the schedule does not call for a new transaction to be created, nothing will happen.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTriggerRecurrenceRequest
 */
func (a *RecurrencesApiService) TriggerRecurrence(ctx _context.Context) ApiTriggerRecurrenceRequest {
	return ApiTriggerRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *RecurrencesApiService) TriggerRecurrenceExecute(r ApiTriggerRecurrenceRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.TriggerRecurrence")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences/trigger"

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

type ApiUpdateRecurrenceRequest struct {
	ctx _context.Context
	ApiService RecurrencesApi
	id int32
	recurrence *Recurrence
}

func (r ApiUpdateRecurrenceRequest) Recurrence(recurrence Recurrence) ApiUpdateRecurrenceRequest {
	r.recurrence = &recurrence
	return r
}

func (r ApiUpdateRecurrenceRequest) Execute() (RecurrenceSingle, *_nethttp.Response, error) {
	return r.ApiService.UpdateRecurrenceExecute(r)
}

/*
 * UpdateRecurrence Update existing recurring transaction.
 * Update existing recurring transaction.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the recurring transaction.
 * @return ApiUpdateRecurrenceRequest
 */
func (a *RecurrencesApiService) UpdateRecurrence(ctx _context.Context, id int32) ApiUpdateRecurrenceRequest {
	return ApiUpdateRecurrenceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return RecurrenceSingle
 */
func (a *RecurrencesApiService) UpdateRecurrenceExecute(r ApiUpdateRecurrenceRequest) (RecurrenceSingle, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RecurrenceSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurrencesApiService.UpdateRecurrence")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurrences/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.recurrence == nil {
		return localVarReturnValue, nil, reportError("recurrence is required and must be specified")
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
	localVarPostBody = r.recurrence
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