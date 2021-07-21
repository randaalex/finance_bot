# BudgetLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CurrencyId** | Pointer to **int32** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] [readonly] 
**BudgetId** | **int32** | The budget ID of the associated budget. | 
**Start** | **string** | Start date of the budget limit. | 
**End** | **string** | End date of the budget limit. | 
**Amount** | **string** |  | 
**Spent** | Pointer to [**[]BudgetSpent**](BudgetSpent.md) |  | [optional] [readonly] 

## Methods

### NewBudgetLimit

`func NewBudgetLimit(budgetId int32, start string, end string, amount string, ) *BudgetLimit`

NewBudgetLimit instantiates a new BudgetLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetLimitWithDefaults

`func NewBudgetLimitWithDefaults() *BudgetLimit`

NewBudgetLimitWithDefaults instantiates a new BudgetLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BudgetLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BudgetLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BudgetLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BudgetLimit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BudgetLimit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BudgetLimit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BudgetLimit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BudgetLimit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCurrencyId

`func (o *BudgetLimit) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BudgetLimit) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BudgetLimit) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BudgetLimit) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BudgetLimit) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BudgetLimit) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BudgetLimit) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BudgetLimit) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *BudgetLimit) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *BudgetLimit) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *BudgetLimit) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *BudgetLimit) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *BudgetLimit) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *BudgetLimit) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *BudgetLimit) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *BudgetLimit) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetBudgetId

`func (o *BudgetLimit) GetBudgetId() int32`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *BudgetLimit) GetBudgetIdOk() (*int32, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *BudgetLimit) SetBudgetId(v int32)`

SetBudgetId sets BudgetId field to given value.


### GetStart

`func (o *BudgetLimit) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *BudgetLimit) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *BudgetLimit) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *BudgetLimit) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *BudgetLimit) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *BudgetLimit) SetEnd(v string)`

SetEnd sets End field to given value.


### GetAmount

`func (o *BudgetLimit) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BudgetLimit) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BudgetLimit) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSpent

`func (o *BudgetLimit) GetSpent() []BudgetSpent`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *BudgetLimit) GetSpentOk() (*[]BudgetSpent, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *BudgetLimit) SetSpent(v []BudgetSpent)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *BudgetLimit) HasSpent() bool`

HasSpent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


