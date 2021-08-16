# AvailableBudgetStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. | [optional] 
**Amount** | **string** |  | 
**Start** | **string** | Start date of the available budget. | 
**End** | **string** | End date of the available budget. | 

## Methods

### NewAvailableBudgetStore

`func NewAvailableBudgetStore(amount string, start string, end string, ) *AvailableBudgetStore`

NewAvailableBudgetStore instantiates a new AvailableBudgetStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableBudgetStoreWithDefaults

`func NewAvailableBudgetStoreWithDefaults() *AvailableBudgetStore`

NewAvailableBudgetStoreWithDefaults instantiates a new AvailableBudgetStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *AvailableBudgetStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AvailableBudgetStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AvailableBudgetStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AvailableBudgetStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AvailableBudgetStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AvailableBudgetStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AvailableBudgetStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AvailableBudgetStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmount

`func (o *AvailableBudgetStore) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AvailableBudgetStore) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AvailableBudgetStore) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetStart

`func (o *AvailableBudgetStore) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AvailableBudgetStore) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AvailableBudgetStore) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *AvailableBudgetStore) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AvailableBudgetStore) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AvailableBudgetStore) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


