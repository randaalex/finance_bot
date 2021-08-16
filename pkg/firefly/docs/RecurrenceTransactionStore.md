# RecurrenceTransactionStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Amount** | **string** | Amount of the transaction. | 
**ForeignAmount** | Pointer to **NullableString** | Foreign amount of the transaction. | [optional] 
**CurrencyId** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**CurrencyCode** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**ForeignCurrencyId** | Pointer to **NullableString** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] 
**ForeignCurrencyCode** | Pointer to **NullableString** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] 
**BudgetId** | Pointer to **string** | The budget ID for this transaction. | [optional] 
**CategoryId** | Pointer to **string** | Category ID for this transaction. | [optional] 
**SourceId** | **string** | ID of the source account. | 
**DestinationId** | **string** | ID of the destination account. | 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**PiggyBankId** | Pointer to **NullableString** | Optional. | [optional] 

## Methods

### NewRecurrenceTransactionStore

`func NewRecurrenceTransactionStore(description string, amount string, sourceId string, destinationId string, ) *RecurrenceTransactionStore`

NewRecurrenceTransactionStore instantiates a new RecurrenceTransactionStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceTransactionStoreWithDefaults

`func NewRecurrenceTransactionStoreWithDefaults() *RecurrenceTransactionStore`

NewRecurrenceTransactionStoreWithDefaults instantiates a new RecurrenceTransactionStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RecurrenceTransactionStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceTransactionStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceTransactionStore) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *RecurrenceTransactionStore) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecurrenceTransactionStore) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecurrenceTransactionStore) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetForeignAmount

`func (o *RecurrenceTransactionStore) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *RecurrenceTransactionStore) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *RecurrenceTransactionStore) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *RecurrenceTransactionStore) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *RecurrenceTransactionStore) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *RecurrenceTransactionStore) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetCurrencyId

`func (o *RecurrenceTransactionStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *RecurrenceTransactionStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *RecurrenceTransactionStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *RecurrenceTransactionStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RecurrenceTransactionStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RecurrenceTransactionStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RecurrenceTransactionStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RecurrenceTransactionStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetForeignCurrencyId

`func (o *RecurrenceTransactionStore) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *RecurrenceTransactionStore) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *RecurrenceTransactionStore) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *RecurrenceTransactionStore) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *RecurrenceTransactionStore) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *RecurrenceTransactionStore) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetForeignCurrencyCode

`func (o *RecurrenceTransactionStore) GetForeignCurrencyCode() string`

GetForeignCurrencyCode returns the ForeignCurrencyCode field if non-nil, zero value otherwise.

### GetForeignCurrencyCodeOk

`func (o *RecurrenceTransactionStore) GetForeignCurrencyCodeOk() (*string, bool)`

GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyCode

`func (o *RecurrenceTransactionStore) SetForeignCurrencyCode(v string)`

SetForeignCurrencyCode sets ForeignCurrencyCode field to given value.

### HasForeignCurrencyCode

`func (o *RecurrenceTransactionStore) HasForeignCurrencyCode() bool`

HasForeignCurrencyCode returns a boolean if a field has been set.

### SetForeignCurrencyCodeNil

`func (o *RecurrenceTransactionStore) SetForeignCurrencyCodeNil(b bool)`

 SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil

### UnsetForeignCurrencyCode
`func (o *RecurrenceTransactionStore) UnsetForeignCurrencyCode()`

UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
### GetBudgetId

`func (o *RecurrenceTransactionStore) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *RecurrenceTransactionStore) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *RecurrenceTransactionStore) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *RecurrenceTransactionStore) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### GetCategoryId

`func (o *RecurrenceTransactionStore) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *RecurrenceTransactionStore) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *RecurrenceTransactionStore) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *RecurrenceTransactionStore) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSourceId

`func (o *RecurrenceTransactionStore) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RecurrenceTransactionStore) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RecurrenceTransactionStore) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *RecurrenceTransactionStore) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RecurrenceTransactionStore) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RecurrenceTransactionStore) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetTags

`func (o *RecurrenceTransactionStore) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecurrenceTransactionStore) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecurrenceTransactionStore) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecurrenceTransactionStore) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RecurrenceTransactionStore) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RecurrenceTransactionStore) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPiggyBankId

`func (o *RecurrenceTransactionStore) GetPiggyBankId() string`

GetPiggyBankId returns the PiggyBankId field if non-nil, zero value otherwise.

### GetPiggyBankIdOk

`func (o *RecurrenceTransactionStore) GetPiggyBankIdOk() (*string, bool)`

GetPiggyBankIdOk returns a tuple with the PiggyBankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankId

`func (o *RecurrenceTransactionStore) SetPiggyBankId(v string)`

SetPiggyBankId sets PiggyBankId field to given value.

### HasPiggyBankId

`func (o *RecurrenceTransactionStore) HasPiggyBankId() bool`

HasPiggyBankId returns a boolean if a field has been set.

### SetPiggyBankIdNil

`func (o *RecurrenceTransactionStore) SetPiggyBankIdNil(b bool)`

 SetPiggyBankIdNil sets the value for PiggyBankId to be an explicit nil

### UnsetPiggyBankId
`func (o *RecurrenceTransactionStore) UnsetPiggyBankId()`

UnsetPiggyBankId ensures that no value is present for PiggyBankId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


