# AvailableBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] [readonly] 
**Amount** | **string** |  | 
**Start** | **time.Time** | Start date of the available budget. | 
**End** | **time.Time** | End date of the available budget. | 
**SpentInBudgets** | Pointer to [**[]BudgetSpent**](BudgetSpent.md) |  | [optional] [readonly] 
**SpentOutsideBudget** | Pointer to [**[]BudgetSpent**](BudgetSpent.md) |  | [optional] [readonly] 

## Methods

### NewAvailableBudget

`func NewAvailableBudget(amount string, start time.Time, end time.Time, ) *AvailableBudget`

NewAvailableBudget instantiates a new AvailableBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableBudgetWithDefaults

`func NewAvailableBudgetWithDefaults() *AvailableBudget`

NewAvailableBudgetWithDefaults instantiates a new AvailableBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AvailableBudget) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AvailableBudget) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AvailableBudget) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AvailableBudget) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AvailableBudget) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AvailableBudget) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AvailableBudget) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AvailableBudget) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AvailableBudget) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AvailableBudget) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AvailableBudget) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AvailableBudget) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AvailableBudget) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AvailableBudget) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AvailableBudget) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AvailableBudget) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *AvailableBudget) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *AvailableBudget) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *AvailableBudget) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *AvailableBudget) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *AvailableBudget) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *AvailableBudget) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *AvailableBudget) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *AvailableBudget) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetAmount

`func (o *AvailableBudget) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AvailableBudget) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AvailableBudget) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetStart

`func (o *AvailableBudget) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AvailableBudget) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AvailableBudget) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *AvailableBudget) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AvailableBudget) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AvailableBudget) SetEnd(v time.Time)`

SetEnd sets End field to given value.


### GetSpentInBudgets

`func (o *AvailableBudget) GetSpentInBudgets() []BudgetSpent`

GetSpentInBudgets returns the SpentInBudgets field if non-nil, zero value otherwise.

### GetSpentInBudgetsOk

`func (o *AvailableBudget) GetSpentInBudgetsOk() (*[]BudgetSpent, bool)`

GetSpentInBudgetsOk returns a tuple with the SpentInBudgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentInBudgets

`func (o *AvailableBudget) SetSpentInBudgets(v []BudgetSpent)`

SetSpentInBudgets sets SpentInBudgets field to given value.

### HasSpentInBudgets

`func (o *AvailableBudget) HasSpentInBudgets() bool`

HasSpentInBudgets returns a boolean if a field has been set.

### GetSpentOutsideBudget

`func (o *AvailableBudget) GetSpentOutsideBudget() []BudgetSpent`

GetSpentOutsideBudget returns the SpentOutsideBudget field if non-nil, zero value otherwise.

### GetSpentOutsideBudgetOk

`func (o *AvailableBudget) GetSpentOutsideBudgetOk() (*[]BudgetSpent, bool)`

GetSpentOutsideBudgetOk returns a tuple with the SpentOutsideBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentOutsideBudget

`func (o *AvailableBudget) SetSpentOutsideBudget(v []BudgetSpent)`

SetSpentOutsideBudget sets SpentOutsideBudget field to given value.

### HasSpentOutsideBudget

`func (o *AvailableBudget) HasSpentOutsideBudget() bool`

HasSpentOutsideBudget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


