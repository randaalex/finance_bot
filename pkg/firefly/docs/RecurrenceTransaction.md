# RecurrenceTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Amount** | **string** | Amount of the transaction. | 
**ForeignAmount** | Pointer to **NullableString** | Foreign amount of the transaction. | [optional] 
**CurrencyId** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**CurrencyCode** | Pointer to **string** | Submit either a currency_id or a currency_code. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** | Number of decimals in the currency | [optional] [readonly] 
**ForeignCurrencyId** | Pointer to **NullableString** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] 
**ForeignCurrencyCode** | Pointer to **NullableString** | Submit either a foreign_currency_id or a foreign_currency_code, or neither. | [optional] 
**ForeignCurrencySymbol** | Pointer to **NullableString** |  | [optional] [readonly] 
**ForeignCurrencyDecimalPlaces** | Pointer to **NullableInt32** | Number of decimals in the currency | [optional] [readonly] 
**BudgetId** | Pointer to **string** | The budget ID for this transaction. | [optional] 
**BudgetName** | Pointer to **NullableString** | The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored. | [optional] [readonly] 
**CategoryId** | Pointer to **string** | Category ID for this transaction. | [optional] 
**CategoryName** | Pointer to **string** | Category name for this transaction. | [optional] 
**SourceId** | Pointer to **string** | ID of the source account. Submit either this or source_name. | [optional] 
**SourceName** | Pointer to **string** | Name of the source account. Submit either this or source_id. | [optional] 
**SourceIban** | Pointer to **NullableString** |  | [optional] [readonly] 
**SourceType** | Pointer to [**AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] 
**DestinationId** | Pointer to **string** | ID of the destination account. Submit either this or destination_name. | [optional] 
**DestinationName** | Pointer to **string** | Name of the destination account. Submit either this or destination_id. | [optional] 
**DestinationIban** | Pointer to **NullableString** |  | [optional] [readonly] 
**DestinationType** | Pointer to [**AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**PiggyBankId** | Pointer to **NullableString** | Optional. Use either this or the piggy_bank_name | [optional] 
**PiggyBankName** | Pointer to **NullableString** | Optional. Use either this or the piggy_bank_id | [optional] 

## Methods

### NewRecurrenceTransaction

`func NewRecurrenceTransaction(description string, amount string, ) *RecurrenceTransaction`

NewRecurrenceTransaction instantiates a new RecurrenceTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceTransactionWithDefaults

`func NewRecurrenceTransactionWithDefaults() *RecurrenceTransaction`

NewRecurrenceTransactionWithDefaults instantiates a new RecurrenceTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RecurrenceTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *RecurrenceTransaction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecurrenceTransaction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecurrenceTransaction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetForeignAmount

`func (o *RecurrenceTransaction) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *RecurrenceTransaction) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *RecurrenceTransaction) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *RecurrenceTransaction) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *RecurrenceTransaction) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *RecurrenceTransaction) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetCurrencyId

`func (o *RecurrenceTransaction) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *RecurrenceTransaction) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *RecurrenceTransaction) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *RecurrenceTransaction) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *RecurrenceTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *RecurrenceTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *RecurrenceTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *RecurrenceTransaction) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *RecurrenceTransaction) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *RecurrenceTransaction) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *RecurrenceTransaction) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *RecurrenceTransaction) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *RecurrenceTransaction) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetForeignCurrencyId

`func (o *RecurrenceTransaction) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *RecurrenceTransaction) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *RecurrenceTransaction) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *RecurrenceTransaction) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *RecurrenceTransaction) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *RecurrenceTransaction) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetForeignCurrencyCode

`func (o *RecurrenceTransaction) GetForeignCurrencyCode() string`

GetForeignCurrencyCode returns the ForeignCurrencyCode field if non-nil, zero value otherwise.

### GetForeignCurrencyCodeOk

`func (o *RecurrenceTransaction) GetForeignCurrencyCodeOk() (*string, bool)`

GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyCode

`func (o *RecurrenceTransaction) SetForeignCurrencyCode(v string)`

SetForeignCurrencyCode sets ForeignCurrencyCode field to given value.

### HasForeignCurrencyCode

`func (o *RecurrenceTransaction) HasForeignCurrencyCode() bool`

HasForeignCurrencyCode returns a boolean if a field has been set.

### SetForeignCurrencyCodeNil

`func (o *RecurrenceTransaction) SetForeignCurrencyCodeNil(b bool)`

 SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil

### UnsetForeignCurrencyCode
`func (o *RecurrenceTransaction) UnsetForeignCurrencyCode()`

UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
### GetForeignCurrencySymbol

`func (o *RecurrenceTransaction) GetForeignCurrencySymbol() string`

GetForeignCurrencySymbol returns the ForeignCurrencySymbol field if non-nil, zero value otherwise.

### GetForeignCurrencySymbolOk

`func (o *RecurrenceTransaction) GetForeignCurrencySymbolOk() (*string, bool)`

GetForeignCurrencySymbolOk returns a tuple with the ForeignCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencySymbol

`func (o *RecurrenceTransaction) SetForeignCurrencySymbol(v string)`

SetForeignCurrencySymbol sets ForeignCurrencySymbol field to given value.

### HasForeignCurrencySymbol

`func (o *RecurrenceTransaction) HasForeignCurrencySymbol() bool`

HasForeignCurrencySymbol returns a boolean if a field has been set.

### SetForeignCurrencySymbolNil

`func (o *RecurrenceTransaction) SetForeignCurrencySymbolNil(b bool)`

 SetForeignCurrencySymbolNil sets the value for ForeignCurrencySymbol to be an explicit nil

### UnsetForeignCurrencySymbol
`func (o *RecurrenceTransaction) UnsetForeignCurrencySymbol()`

UnsetForeignCurrencySymbol ensures that no value is present for ForeignCurrencySymbol, not even an explicit nil
### GetForeignCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) GetForeignCurrencyDecimalPlaces() int32`

GetForeignCurrencyDecimalPlaces returns the ForeignCurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetForeignCurrencyDecimalPlacesOk

`func (o *RecurrenceTransaction) GetForeignCurrencyDecimalPlacesOk() (*int32, bool)`

GetForeignCurrencyDecimalPlacesOk returns a tuple with the ForeignCurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) SetForeignCurrencyDecimalPlaces(v int32)`

SetForeignCurrencyDecimalPlaces sets ForeignCurrencyDecimalPlaces field to given value.

### HasForeignCurrencyDecimalPlaces

`func (o *RecurrenceTransaction) HasForeignCurrencyDecimalPlaces() bool`

HasForeignCurrencyDecimalPlaces returns a boolean if a field has been set.

### SetForeignCurrencyDecimalPlacesNil

`func (o *RecurrenceTransaction) SetForeignCurrencyDecimalPlacesNil(b bool)`

 SetForeignCurrencyDecimalPlacesNil sets the value for ForeignCurrencyDecimalPlaces to be an explicit nil

### UnsetForeignCurrencyDecimalPlaces
`func (o *RecurrenceTransaction) UnsetForeignCurrencyDecimalPlaces()`

UnsetForeignCurrencyDecimalPlaces ensures that no value is present for ForeignCurrencyDecimalPlaces, not even an explicit nil
### GetBudgetId

`func (o *RecurrenceTransaction) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *RecurrenceTransaction) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *RecurrenceTransaction) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *RecurrenceTransaction) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### GetBudgetName

`func (o *RecurrenceTransaction) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *RecurrenceTransaction) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *RecurrenceTransaction) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.

### HasBudgetName

`func (o *RecurrenceTransaction) HasBudgetName() bool`

HasBudgetName returns a boolean if a field has been set.

### SetBudgetNameNil

`func (o *RecurrenceTransaction) SetBudgetNameNil(b bool)`

 SetBudgetNameNil sets the value for BudgetName to be an explicit nil

### UnsetBudgetName
`func (o *RecurrenceTransaction) UnsetBudgetName()`

UnsetBudgetName ensures that no value is present for BudgetName, not even an explicit nil
### GetCategoryId

`func (o *RecurrenceTransaction) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *RecurrenceTransaction) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *RecurrenceTransaction) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *RecurrenceTransaction) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryName

`func (o *RecurrenceTransaction) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *RecurrenceTransaction) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *RecurrenceTransaction) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *RecurrenceTransaction) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetSourceId

`func (o *RecurrenceTransaction) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RecurrenceTransaction) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RecurrenceTransaction) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *RecurrenceTransaction) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *RecurrenceTransaction) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *RecurrenceTransaction) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *RecurrenceTransaction) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *RecurrenceTransaction) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceIban

`func (o *RecurrenceTransaction) GetSourceIban() string`

GetSourceIban returns the SourceIban field if non-nil, zero value otherwise.

### GetSourceIbanOk

`func (o *RecurrenceTransaction) GetSourceIbanOk() (*string, bool)`

GetSourceIbanOk returns a tuple with the SourceIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIban

`func (o *RecurrenceTransaction) SetSourceIban(v string)`

SetSourceIban sets SourceIban field to given value.

### HasSourceIban

`func (o *RecurrenceTransaction) HasSourceIban() bool`

HasSourceIban returns a boolean if a field has been set.

### SetSourceIbanNil

`func (o *RecurrenceTransaction) SetSourceIbanNil(b bool)`

 SetSourceIbanNil sets the value for SourceIban to be an explicit nil

### UnsetSourceIban
`func (o *RecurrenceTransaction) UnsetSourceIban()`

UnsetSourceIban ensures that no value is present for SourceIban, not even an explicit nil
### GetSourceType

`func (o *RecurrenceTransaction) GetSourceType() AccountTypeProperty`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RecurrenceTransaction) GetSourceTypeOk() (*AccountTypeProperty, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RecurrenceTransaction) SetSourceType(v AccountTypeProperty)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *RecurrenceTransaction) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationId

`func (o *RecurrenceTransaction) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *RecurrenceTransaction) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *RecurrenceTransaction) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *RecurrenceTransaction) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationName

`func (o *RecurrenceTransaction) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *RecurrenceTransaction) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *RecurrenceTransaction) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *RecurrenceTransaction) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDestinationIban

`func (o *RecurrenceTransaction) GetDestinationIban() string`

GetDestinationIban returns the DestinationIban field if non-nil, zero value otherwise.

### GetDestinationIbanOk

`func (o *RecurrenceTransaction) GetDestinationIbanOk() (*string, bool)`

GetDestinationIbanOk returns a tuple with the DestinationIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIban

`func (o *RecurrenceTransaction) SetDestinationIban(v string)`

SetDestinationIban sets DestinationIban field to given value.

### HasDestinationIban

`func (o *RecurrenceTransaction) HasDestinationIban() bool`

HasDestinationIban returns a boolean if a field has been set.

### SetDestinationIbanNil

`func (o *RecurrenceTransaction) SetDestinationIbanNil(b bool)`

 SetDestinationIbanNil sets the value for DestinationIban to be an explicit nil

### UnsetDestinationIban
`func (o *RecurrenceTransaction) UnsetDestinationIban()`

UnsetDestinationIban ensures that no value is present for DestinationIban, not even an explicit nil
### GetDestinationType

`func (o *RecurrenceTransaction) GetDestinationType() AccountTypeProperty`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RecurrenceTransaction) GetDestinationTypeOk() (*AccountTypeProperty, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RecurrenceTransaction) SetDestinationType(v AccountTypeProperty)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *RecurrenceTransaction) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetTags

`func (o *RecurrenceTransaction) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecurrenceTransaction) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecurrenceTransaction) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecurrenceTransaction) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RecurrenceTransaction) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RecurrenceTransaction) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPiggyBankId

`func (o *RecurrenceTransaction) GetPiggyBankId() string`

GetPiggyBankId returns the PiggyBankId field if non-nil, zero value otherwise.

### GetPiggyBankIdOk

`func (o *RecurrenceTransaction) GetPiggyBankIdOk() (*string, bool)`

GetPiggyBankIdOk returns a tuple with the PiggyBankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankId

`func (o *RecurrenceTransaction) SetPiggyBankId(v string)`

SetPiggyBankId sets PiggyBankId field to given value.

### HasPiggyBankId

`func (o *RecurrenceTransaction) HasPiggyBankId() bool`

HasPiggyBankId returns a boolean if a field has been set.

### SetPiggyBankIdNil

`func (o *RecurrenceTransaction) SetPiggyBankIdNil(b bool)`

 SetPiggyBankIdNil sets the value for PiggyBankId to be an explicit nil

### UnsetPiggyBankId
`func (o *RecurrenceTransaction) UnsetPiggyBankId()`

UnsetPiggyBankId ensures that no value is present for PiggyBankId, not even an explicit nil
### GetPiggyBankName

`func (o *RecurrenceTransaction) GetPiggyBankName() string`

GetPiggyBankName returns the PiggyBankName field if non-nil, zero value otherwise.

### GetPiggyBankNameOk

`func (o *RecurrenceTransaction) GetPiggyBankNameOk() (*string, bool)`

GetPiggyBankNameOk returns a tuple with the PiggyBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankName

`func (o *RecurrenceTransaction) SetPiggyBankName(v string)`

SetPiggyBankName sets PiggyBankName field to given value.

### HasPiggyBankName

`func (o *RecurrenceTransaction) HasPiggyBankName() bool`

HasPiggyBankName returns a boolean if a field has been set.

### SetPiggyBankNameNil

`func (o *RecurrenceTransaction) SetPiggyBankNameNil(b bool)`

 SetPiggyBankNameNil sets the value for PiggyBankName to be an explicit nil

### UnsetPiggyBankName
`func (o *RecurrenceTransaction) UnsetPiggyBankName()`

UnsetPiggyBankName ensures that no value is present for PiggyBankName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


