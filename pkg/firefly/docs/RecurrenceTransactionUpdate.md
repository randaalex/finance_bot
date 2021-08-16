# RecurrenceTransactionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** | Amount of the transaction. | [optional] 
**ForeignAmount** | Pointer to **NullableString** | Foreign amount of the transaction. | [optional] 
**CurrencyId** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**CurrencyCode** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**ForeignCurrencyId** | Pointer to **NullableString** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] 
**BudgetId** | Pointer to **string** | The budget ID for this transaction. | [optional] 
**CategoryId** | Pointer to **string** | Category ID for this transaction. | [optional] 
**SourceId** | Pointer to **string** | ID of the source account. Submit either this or source_name. | [optional] 
**DestinationId** | Pointer to **string** | ID of the destination account. Submit either this or destination_name. | [optional] 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**PiggyBankId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRecurrenceTransactionUpdate

`func NewRecurrenceTransactionUpdate() *RecurrenceTransactionUpdate`

NewRecurrenceTransactionUpdate instantiates a new RecurrenceTransactionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceTransactionUpdateWithDefaults

`func NewRecurrenceTransactionUpdateWithDefaults() *RecurrenceTransactionUpdate`

NewRecurrenceTransactionUpdateWithDefaults instantiates a new RecurrenceTransactionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RecurrenceTransactionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceTransactionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceTransactionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurrenceTransactionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *RecurrenceTransactionUpdate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecurrenceTransactionUpdate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecurrenceTransactionUpdate) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RecurrenceTransactionUpdate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetForeignAmount

`func (o *RecurrenceTransactionUpdate) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *RecurrenceTransactionUpdate) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *RecurrenceTransactionUpdate) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *RecurrenceTransactionUpdate) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *RecurrenceTransactionUpdate) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *RecurrenceTransactionUpdate) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetCurrencyId

`func (o *RecurrenceTransactionUpdate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *RecurrenceTransactionUpdate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *RecurrenceTransactionUpdate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *RecurrenceTransactionUpdate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RecurrenceTransactionUpdate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RecurrenceTransactionUpdate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RecurrenceTransactionUpdate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RecurrenceTransactionUpdate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetForeignCurrencyId

`func (o *RecurrenceTransactionUpdate) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *RecurrenceTransactionUpdate) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *RecurrenceTransactionUpdate) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *RecurrenceTransactionUpdate) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *RecurrenceTransactionUpdate) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *RecurrenceTransactionUpdate) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetBudgetId

`func (o *RecurrenceTransactionUpdate) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *RecurrenceTransactionUpdate) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *RecurrenceTransactionUpdate) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *RecurrenceTransactionUpdate) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### GetCategoryId

`func (o *RecurrenceTransactionUpdate) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *RecurrenceTransactionUpdate) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *RecurrenceTransactionUpdate) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *RecurrenceTransactionUpdate) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetSourceId

`func (o *RecurrenceTransactionUpdate) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RecurrenceTransactionUpdate) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RecurrenceTransactionUpdate) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *RecurrenceTransactionUpdate) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *RecurrenceTransactionUpdate) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RecurrenceTransactionUpdate) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RecurrenceTransactionUpdate) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *RecurrenceTransactionUpdate) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetTags

`func (o *RecurrenceTransactionUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecurrenceTransactionUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecurrenceTransactionUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecurrenceTransactionUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RecurrenceTransactionUpdate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RecurrenceTransactionUpdate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPiggyBankId

`func (o *RecurrenceTransactionUpdate) GetPiggyBankId() string`

GetPiggyBankId returns the PiggyBankId field if non-nil, zero value otherwise.

### GetPiggyBankIdOk

`func (o *RecurrenceTransactionUpdate) GetPiggyBankIdOk() (*string, bool)`

GetPiggyBankIdOk returns a tuple with the PiggyBankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankId

`func (o *RecurrenceTransactionUpdate) SetPiggyBankId(v string)`

SetPiggyBankId sets PiggyBankId field to given value.

### HasPiggyBankId

`func (o *RecurrenceTransactionUpdate) HasPiggyBankId() bool`

HasPiggyBankId returns a boolean if a field has been set.

### SetPiggyBankIdNil

`func (o *RecurrenceTransactionUpdate) SetPiggyBankIdNil(b bool)`

 SetPiggyBankIdNil sets the value for PiggyBankId to be an explicit nil

### UnsetPiggyBankId
`func (o *RecurrenceTransactionUpdate) UnsetPiggyBankId()`

UnsetPiggyBankId ensures that no value is present for PiggyBankId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


