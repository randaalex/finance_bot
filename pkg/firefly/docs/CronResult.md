# CronResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurringTransactions** | Pointer to [**CronResultRow**](CronResultRow.md) |  | [optional] 
**AutoBudgets** | Pointer to [**CronResultRow**](CronResultRow.md) |  | [optional] 
**Telemetry** | Pointer to [**CronResultRow**](CronResultRow.md) |  | [optional] 

## Methods

### NewCronResult

`func NewCronResult() *CronResult`

NewCronResult instantiates a new CronResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronResultWithDefaults

`func NewCronResultWithDefaults() *CronResult`

NewCronResultWithDefaults instantiates a new CronResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurringTransactions

`func (o *CronResult) GetRecurringTransactions() CronResultRow`

GetRecurringTransactions returns the RecurringTransactions field if non-nil, zero value otherwise.

### GetRecurringTransactionsOk

`func (o *CronResult) GetRecurringTransactionsOk() (*CronResultRow, bool)`

GetRecurringTransactionsOk returns a tuple with the RecurringTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTransactions

`func (o *CronResult) SetRecurringTransactions(v CronResultRow)`

SetRecurringTransactions sets RecurringTransactions field to given value.

### HasRecurringTransactions

`func (o *CronResult) HasRecurringTransactions() bool`

HasRecurringTransactions returns a boolean if a field has been set.

### GetAutoBudgets

`func (o *CronResult) GetAutoBudgets() CronResultRow`

GetAutoBudgets returns the AutoBudgets field if non-nil, zero value otherwise.

### GetAutoBudgetsOk

`func (o *CronResult) GetAutoBudgetsOk() (*CronResultRow, bool)`

GetAutoBudgetsOk returns a tuple with the AutoBudgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgets

`func (o *CronResult) SetAutoBudgets(v CronResultRow)`

SetAutoBudgets sets AutoBudgets field to given value.

### HasAutoBudgets

`func (o *CronResult) HasAutoBudgets() bool`

HasAutoBudgets returns a boolean if a field has been set.

### GetTelemetry

`func (o *CronResult) GetTelemetry() CronResultRow`

GetTelemetry returns the Telemetry field if non-nil, zero value otherwise.

### GetTelemetryOk

`func (o *CronResult) GetTelemetryOk() (*CronResultRow, bool)`

GetTelemetryOk returns a tuple with the Telemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetry

`func (o *CronResult) SetTelemetry(v CronResultRow)`

SetTelemetry sets Telemetry field to given value.

### HasTelemetry

`func (o *CronResult) HasTelemetry() bool`

HasTelemetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


