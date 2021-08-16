# PiggyBankEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrencyId** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**TransactionJournalId** | Pointer to **string** | The journal associated with the event. | [optional] 
**TransactionGroupId** | Pointer to **string** | The transaction group associated with the event. | [optional] 

## Methods

### NewPiggyBankEvent

`func NewPiggyBankEvent() *PiggyBankEvent`

NewPiggyBankEvent instantiates a new PiggyBankEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankEventWithDefaults

`func NewPiggyBankEventWithDefaults() *PiggyBankEvent`

NewPiggyBankEventWithDefaults instantiates a new PiggyBankEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PiggyBankEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PiggyBankEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PiggyBankEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PiggyBankEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PiggyBankEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PiggyBankEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PiggyBankEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PiggyBankEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PiggyBankEvent) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PiggyBankEvent) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PiggyBankEvent) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PiggyBankEvent) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PiggyBankEvent) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PiggyBankEvent) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PiggyBankEvent) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PiggyBankEvent) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *PiggyBankEvent) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *PiggyBankEvent) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *PiggyBankEvent) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *PiggyBankEvent) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *PiggyBankEvent) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *PiggyBankEvent) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *PiggyBankEvent) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *PiggyBankEvent) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetAmount

`func (o *PiggyBankEvent) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PiggyBankEvent) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PiggyBankEvent) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PiggyBankEvent) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionJournalId

`func (o *PiggyBankEvent) GetTransactionJournalId() string`

GetTransactionJournalId returns the TransactionJournalId field if non-nil, zero value otherwise.

### GetTransactionJournalIdOk

`func (o *PiggyBankEvent) GetTransactionJournalIdOk() (*string, bool)`

GetTransactionJournalIdOk returns a tuple with the TransactionJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionJournalId

`func (o *PiggyBankEvent) SetTransactionJournalId(v string)`

SetTransactionJournalId sets TransactionJournalId field to given value.

### HasTransactionJournalId

`func (o *PiggyBankEvent) HasTransactionJournalId() bool`

HasTransactionJournalId returns a boolean if a field has been set.

### GetTransactionGroupId

`func (o *PiggyBankEvent) GetTransactionGroupId() string`

GetTransactionGroupId returns the TransactionGroupId field if non-nil, zero value otherwise.

### GetTransactionGroupIdOk

`func (o *PiggyBankEvent) GetTransactionGroupIdOk() (*string, bool)`

GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroupId

`func (o *PiggyBankEvent) SetTransactionGroupId(v string)`

SetTransactionGroupId sets TransactionGroupId field to given value.

### HasTransactionGroupId

`func (o *PiggyBankEvent) HasTransactionGroupId() bool`

HasTransactionGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


