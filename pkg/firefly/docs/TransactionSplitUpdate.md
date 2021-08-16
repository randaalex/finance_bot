# TransactionSplitUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | Date of the transaction | [optional] 
**Amount** | Pointer to **string** | Amount of the transaction. | [optional] 
**Description** | Pointer to **string** | Description of the transaction. | [optional] 
**Order** | Pointer to **NullableInt32** | Order of this entry in the list of transactions. | [optional] 
**CurrencyId** | Pointer to **NullableString** | Currency ID. Default is the source account&#39;s currency, or the user&#39;s default currency. Can be used instead of currency_code. | [optional] 
**CurrencyCode** | Pointer to **NullableString** | Currency code. Default is the source account&#39;s currency, or the user&#39;s default currency. Can be used instead of currency_id. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyName** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** | Number of decimals used in this currency. | [optional] [readonly] 
**ForeignAmount** | Pointer to **NullableString** | The amount in a foreign currency. | [optional] 
**ForeignCurrencyId** | Pointer to **NullableString** | Currency ID of the foreign currency. Default is null. Is required when you submit a foreign amount. | [optional] 
**ForeignCurrencyCode** | Pointer to **NullableString** | Currency code of the foreign currency. Default is NULL. Can be used instead of the foreign_currency_id, but this or the ID is required when submitting a foreign amount. | [optional] 
**ForeignCurrencySymbol** | Pointer to **NullableString** |  | [optional] [readonly] 
**ForeignCurrencyDecimalPlaces** | Pointer to **NullableInt32** | Number of decimals in the currency | [optional] [readonly] 
**BudgetId** | Pointer to **NullableString** | The budget ID for this transaction. | [optional] 
**BudgetName** | Pointer to **NullableString** | The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored. | [optional] [readonly] 
**CategoryId** | Pointer to **NullableString** | The category ID for this transaction. | [optional] 
**CategoryName** | Pointer to **NullableString** | The name of the category to be used. If the category is unknown, it will be created. If the ID and the name point to different categories, the ID overrules the name. | [optional] 
**SourceId** | Pointer to **NullableString** | ID of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. | [optional] 
**SourceName** | Pointer to **NullableString** | Name of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. Can be used instead of the source_id. If the transaction is a deposit, the source_name can be filled in freely: the account will be created based on the name. | [optional] 
**SourceIban** | Pointer to **NullableString** |  | [optional] 
**DestinationId** | Pointer to **NullableString** | ID of the destination account. For a deposit or a transfer, this must always be an asset account. For withdrawals this must be an expense account. | [optional] 
**DestinationName** | Pointer to **NullableString** | Name of the destination account. You can submit the name instead of the ID. For everything except transfers, the account will be auto-generated if unknown, so submitting a name is enough. | [optional] 
**DestinationIban** | Pointer to **NullableString** |  | [optional] 
**Reconciled** | Pointer to **bool** | If the transaction has been reconciled already. When you set this, the amount can no longer be edited by the user. | [optional] 
**BillId** | Pointer to **NullableString** | Optional. Use either this or the bill_name | [optional] 
**BillName** | Pointer to **NullableString** | Optional. Use either this or the bill_id | [optional] 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**InternalReference** | Pointer to **NullableString** | Reference to internal reference of other systems. | [optional] 
**ExternalId** | Pointer to **NullableString** | Reference to external ID in other systems. | [optional] 
**BunqPaymentId** | Pointer to **NullableString** | Internal ID of bunq transaction. | [optional] 
**SepaCc** | Pointer to **NullableString** | SEPA Clearing Code | [optional] 
**SepaCtOp** | Pointer to **NullableString** | SEPA Opposing Account Identifier | [optional] 
**SepaCtId** | Pointer to **NullableString** | SEPA end-to-end Identifier | [optional] 
**SepaDb** | Pointer to **NullableString** | SEPA mandate identifier | [optional] 
**SepaCountry** | Pointer to **NullableString** | SEPA Country | [optional] 
**SepaEp** | Pointer to **NullableString** | SEPA External Purpose indicator | [optional] 
**SepaCi** | Pointer to **NullableString** | SEPA Creditor Identifier | [optional] 
**SepaBatchId** | Pointer to **NullableString** | SEPA Batch ID | [optional] 
**InterestDate** | Pointer to **NullableTime** |  | [optional] 
**BookDate** | Pointer to **NullableTime** |  | [optional] 
**ProcessDate** | Pointer to **NullableTime** |  | [optional] 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**PaymentDate** | Pointer to **NullableTime** |  | [optional] 
**InvoiceDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTransactionSplitUpdate

`func NewTransactionSplitUpdate() *TransactionSplitUpdate`

NewTransactionSplitUpdate instantiates a new TransactionSplitUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSplitUpdateWithDefaults

`func NewTransactionSplitUpdateWithDefaults() *TransactionSplitUpdate`

NewTransactionSplitUpdateWithDefaults instantiates a new TransactionSplitUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TransactionSplitUpdate) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionSplitUpdate) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionSplitUpdate) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionSplitUpdate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionSplitUpdate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSplitUpdate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSplitUpdate) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionSplitUpdate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionSplitUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionSplitUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionSplitUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionSplitUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *TransactionSplitUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TransactionSplitUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TransactionSplitUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TransactionSplitUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *TransactionSplitUpdate) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *TransactionSplitUpdate) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetCurrencyId

`func (o *TransactionSplitUpdate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TransactionSplitUpdate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TransactionSplitUpdate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TransactionSplitUpdate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TransactionSplitUpdate) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TransactionSplitUpdate) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionSplitUpdate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionSplitUpdate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionSplitUpdate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionSplitUpdate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionSplitUpdate) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionSplitUpdate) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCurrencySymbol

`func (o *TransactionSplitUpdate) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *TransactionSplitUpdate) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *TransactionSplitUpdate) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *TransactionSplitUpdate) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyName

`func (o *TransactionSplitUpdate) GetCurrencyName() string`

GetCurrencyName returns the CurrencyName field if non-nil, zero value otherwise.

### GetCurrencyNameOk

`func (o *TransactionSplitUpdate) GetCurrencyNameOk() (*string, bool)`

GetCurrencyNameOk returns a tuple with the CurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyName

`func (o *TransactionSplitUpdate) SetCurrencyName(v string)`

SetCurrencyName sets CurrencyName field to given value.

### HasCurrencyName

`func (o *TransactionSplitUpdate) HasCurrencyName() bool`

HasCurrencyName returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *TransactionSplitUpdate) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetForeignAmount

`func (o *TransactionSplitUpdate) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *TransactionSplitUpdate) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *TransactionSplitUpdate) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *TransactionSplitUpdate) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *TransactionSplitUpdate) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *TransactionSplitUpdate) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetForeignCurrencyId

`func (o *TransactionSplitUpdate) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *TransactionSplitUpdate) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *TransactionSplitUpdate) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *TransactionSplitUpdate) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *TransactionSplitUpdate) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *TransactionSplitUpdate) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetForeignCurrencyCode

`func (o *TransactionSplitUpdate) GetForeignCurrencyCode() string`

GetForeignCurrencyCode returns the ForeignCurrencyCode field if non-nil, zero value otherwise.

### GetForeignCurrencyCodeOk

`func (o *TransactionSplitUpdate) GetForeignCurrencyCodeOk() (*string, bool)`

GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyCode

`func (o *TransactionSplitUpdate) SetForeignCurrencyCode(v string)`

SetForeignCurrencyCode sets ForeignCurrencyCode field to given value.

### HasForeignCurrencyCode

`func (o *TransactionSplitUpdate) HasForeignCurrencyCode() bool`

HasForeignCurrencyCode returns a boolean if a field has been set.

### SetForeignCurrencyCodeNil

`func (o *TransactionSplitUpdate) SetForeignCurrencyCodeNil(b bool)`

 SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil

### UnsetForeignCurrencyCode
`func (o *TransactionSplitUpdate) UnsetForeignCurrencyCode()`

UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
### GetForeignCurrencySymbol

`func (o *TransactionSplitUpdate) GetForeignCurrencySymbol() string`

GetForeignCurrencySymbol returns the ForeignCurrencySymbol field if non-nil, zero value otherwise.

### GetForeignCurrencySymbolOk

`func (o *TransactionSplitUpdate) GetForeignCurrencySymbolOk() (*string, bool)`

GetForeignCurrencySymbolOk returns a tuple with the ForeignCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencySymbol

`func (o *TransactionSplitUpdate) SetForeignCurrencySymbol(v string)`

SetForeignCurrencySymbol sets ForeignCurrencySymbol field to given value.

### HasForeignCurrencySymbol

`func (o *TransactionSplitUpdate) HasForeignCurrencySymbol() bool`

HasForeignCurrencySymbol returns a boolean if a field has been set.

### SetForeignCurrencySymbolNil

`func (o *TransactionSplitUpdate) SetForeignCurrencySymbolNil(b bool)`

 SetForeignCurrencySymbolNil sets the value for ForeignCurrencySymbol to be an explicit nil

### UnsetForeignCurrencySymbol
`func (o *TransactionSplitUpdate) UnsetForeignCurrencySymbol()`

UnsetForeignCurrencySymbol ensures that no value is present for ForeignCurrencySymbol, not even an explicit nil
### GetForeignCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) GetForeignCurrencyDecimalPlaces() int32`

GetForeignCurrencyDecimalPlaces returns the ForeignCurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetForeignCurrencyDecimalPlacesOk

`func (o *TransactionSplitUpdate) GetForeignCurrencyDecimalPlacesOk() (*int32, bool)`

GetForeignCurrencyDecimalPlacesOk returns a tuple with the ForeignCurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) SetForeignCurrencyDecimalPlaces(v int32)`

SetForeignCurrencyDecimalPlaces sets ForeignCurrencyDecimalPlaces field to given value.

### HasForeignCurrencyDecimalPlaces

`func (o *TransactionSplitUpdate) HasForeignCurrencyDecimalPlaces() bool`

HasForeignCurrencyDecimalPlaces returns a boolean if a field has been set.

### SetForeignCurrencyDecimalPlacesNil

`func (o *TransactionSplitUpdate) SetForeignCurrencyDecimalPlacesNil(b bool)`

 SetForeignCurrencyDecimalPlacesNil sets the value for ForeignCurrencyDecimalPlaces to be an explicit nil

### UnsetForeignCurrencyDecimalPlaces
`func (o *TransactionSplitUpdate) UnsetForeignCurrencyDecimalPlaces()`

UnsetForeignCurrencyDecimalPlaces ensures that no value is present for ForeignCurrencyDecimalPlaces, not even an explicit nil
### GetBudgetId

`func (o *TransactionSplitUpdate) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *TransactionSplitUpdate) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *TransactionSplitUpdate) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *TransactionSplitUpdate) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *TransactionSplitUpdate) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *TransactionSplitUpdate) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetBudgetName

`func (o *TransactionSplitUpdate) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *TransactionSplitUpdate) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *TransactionSplitUpdate) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.

### HasBudgetName

`func (o *TransactionSplitUpdate) HasBudgetName() bool`

HasBudgetName returns a boolean if a field has been set.

### SetBudgetNameNil

`func (o *TransactionSplitUpdate) SetBudgetNameNil(b bool)`

 SetBudgetNameNil sets the value for BudgetName to be an explicit nil

### UnsetBudgetName
`func (o *TransactionSplitUpdate) UnsetBudgetName()`

UnsetBudgetName ensures that no value is present for BudgetName, not even an explicit nil
### GetCategoryId

`func (o *TransactionSplitUpdate) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransactionSplitUpdate) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransactionSplitUpdate) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *TransactionSplitUpdate) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *TransactionSplitUpdate) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *TransactionSplitUpdate) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryName

`func (o *TransactionSplitUpdate) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *TransactionSplitUpdate) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *TransactionSplitUpdate) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *TransactionSplitUpdate) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### SetCategoryNameNil

`func (o *TransactionSplitUpdate) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *TransactionSplitUpdate) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
### GetSourceId

`func (o *TransactionSplitUpdate) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TransactionSplitUpdate) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TransactionSplitUpdate) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TransactionSplitUpdate) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *TransactionSplitUpdate) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *TransactionSplitUpdate) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *TransactionSplitUpdate) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TransactionSplitUpdate) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TransactionSplitUpdate) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TransactionSplitUpdate) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TransactionSplitUpdate) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TransactionSplitUpdate) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceIban

`func (o *TransactionSplitUpdate) GetSourceIban() string`

GetSourceIban returns the SourceIban field if non-nil, zero value otherwise.

### GetSourceIbanOk

`func (o *TransactionSplitUpdate) GetSourceIbanOk() (*string, bool)`

GetSourceIbanOk returns a tuple with the SourceIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIban

`func (o *TransactionSplitUpdate) SetSourceIban(v string)`

SetSourceIban sets SourceIban field to given value.

### HasSourceIban

`func (o *TransactionSplitUpdate) HasSourceIban() bool`

HasSourceIban returns a boolean if a field has been set.

### SetSourceIbanNil

`func (o *TransactionSplitUpdate) SetSourceIbanNil(b bool)`

 SetSourceIbanNil sets the value for SourceIban to be an explicit nil

### UnsetSourceIban
`func (o *TransactionSplitUpdate) UnsetSourceIban()`

UnsetSourceIban ensures that no value is present for SourceIban, not even an explicit nil
### GetDestinationId

`func (o *TransactionSplitUpdate) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *TransactionSplitUpdate) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *TransactionSplitUpdate) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *TransactionSplitUpdate) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### SetDestinationIdNil

`func (o *TransactionSplitUpdate) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *TransactionSplitUpdate) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetDestinationName

`func (o *TransactionSplitUpdate) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *TransactionSplitUpdate) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *TransactionSplitUpdate) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *TransactionSplitUpdate) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### SetDestinationNameNil

`func (o *TransactionSplitUpdate) SetDestinationNameNil(b bool)`

 SetDestinationNameNil sets the value for DestinationName to be an explicit nil

### UnsetDestinationName
`func (o *TransactionSplitUpdate) UnsetDestinationName()`

UnsetDestinationName ensures that no value is present for DestinationName, not even an explicit nil
### GetDestinationIban

`func (o *TransactionSplitUpdate) GetDestinationIban() string`

GetDestinationIban returns the DestinationIban field if non-nil, zero value otherwise.

### GetDestinationIbanOk

`func (o *TransactionSplitUpdate) GetDestinationIbanOk() (*string, bool)`

GetDestinationIbanOk returns a tuple with the DestinationIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIban

`func (o *TransactionSplitUpdate) SetDestinationIban(v string)`

SetDestinationIban sets DestinationIban field to given value.

### HasDestinationIban

`func (o *TransactionSplitUpdate) HasDestinationIban() bool`

HasDestinationIban returns a boolean if a field has been set.

### SetDestinationIbanNil

`func (o *TransactionSplitUpdate) SetDestinationIbanNil(b bool)`

 SetDestinationIbanNil sets the value for DestinationIban to be an explicit nil

### UnsetDestinationIban
`func (o *TransactionSplitUpdate) UnsetDestinationIban()`

UnsetDestinationIban ensures that no value is present for DestinationIban, not even an explicit nil
### GetReconciled

`func (o *TransactionSplitUpdate) GetReconciled() bool`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *TransactionSplitUpdate) GetReconciledOk() (*bool, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *TransactionSplitUpdate) SetReconciled(v bool)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *TransactionSplitUpdate) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetBillId

`func (o *TransactionSplitUpdate) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *TransactionSplitUpdate) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *TransactionSplitUpdate) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *TransactionSplitUpdate) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### SetBillIdNil

`func (o *TransactionSplitUpdate) SetBillIdNil(b bool)`

 SetBillIdNil sets the value for BillId to be an explicit nil

### UnsetBillId
`func (o *TransactionSplitUpdate) UnsetBillId()`

UnsetBillId ensures that no value is present for BillId, not even an explicit nil
### GetBillName

`func (o *TransactionSplitUpdate) GetBillName() string`

GetBillName returns the BillName field if non-nil, zero value otherwise.

### GetBillNameOk

`func (o *TransactionSplitUpdate) GetBillNameOk() (*string, bool)`

GetBillNameOk returns a tuple with the BillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillName

`func (o *TransactionSplitUpdate) SetBillName(v string)`

SetBillName sets BillName field to given value.

### HasBillName

`func (o *TransactionSplitUpdate) HasBillName() bool`

HasBillName returns a boolean if a field has been set.

### SetBillNameNil

`func (o *TransactionSplitUpdate) SetBillNameNil(b bool)`

 SetBillNameNil sets the value for BillName to be an explicit nil

### UnsetBillName
`func (o *TransactionSplitUpdate) UnsetBillName()`

UnsetBillName ensures that no value is present for BillName, not even an explicit nil
### GetTags

`func (o *TransactionSplitUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransactionSplitUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransactionSplitUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TransactionSplitUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TransactionSplitUpdate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TransactionSplitUpdate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetNotes

`func (o *TransactionSplitUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionSplitUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionSplitUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionSplitUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionSplitUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionSplitUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInternalReference

`func (o *TransactionSplitUpdate) GetInternalReference() string`

GetInternalReference returns the InternalReference field if non-nil, zero value otherwise.

### GetInternalReferenceOk

`func (o *TransactionSplitUpdate) GetInternalReferenceOk() (*string, bool)`

GetInternalReferenceOk returns a tuple with the InternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalReference

`func (o *TransactionSplitUpdate) SetInternalReference(v string)`

SetInternalReference sets InternalReference field to given value.

### HasInternalReference

`func (o *TransactionSplitUpdate) HasInternalReference() bool`

HasInternalReference returns a boolean if a field has been set.

### SetInternalReferenceNil

`func (o *TransactionSplitUpdate) SetInternalReferenceNil(b bool)`

 SetInternalReferenceNil sets the value for InternalReference to be an explicit nil

### UnsetInternalReference
`func (o *TransactionSplitUpdate) UnsetInternalReference()`

UnsetInternalReference ensures that no value is present for InternalReference, not even an explicit nil
### GetExternalId

`func (o *TransactionSplitUpdate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TransactionSplitUpdate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TransactionSplitUpdate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TransactionSplitUpdate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TransactionSplitUpdate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TransactionSplitUpdate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBunqPaymentId

`func (o *TransactionSplitUpdate) GetBunqPaymentId() string`

GetBunqPaymentId returns the BunqPaymentId field if non-nil, zero value otherwise.

### GetBunqPaymentIdOk

`func (o *TransactionSplitUpdate) GetBunqPaymentIdOk() (*string, bool)`

GetBunqPaymentIdOk returns a tuple with the BunqPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBunqPaymentId

`func (o *TransactionSplitUpdate) SetBunqPaymentId(v string)`

SetBunqPaymentId sets BunqPaymentId field to given value.

### HasBunqPaymentId

`func (o *TransactionSplitUpdate) HasBunqPaymentId() bool`

HasBunqPaymentId returns a boolean if a field has been set.

### SetBunqPaymentIdNil

`func (o *TransactionSplitUpdate) SetBunqPaymentIdNil(b bool)`

 SetBunqPaymentIdNil sets the value for BunqPaymentId to be an explicit nil

### UnsetBunqPaymentId
`func (o *TransactionSplitUpdate) UnsetBunqPaymentId()`

UnsetBunqPaymentId ensures that no value is present for BunqPaymentId, not even an explicit nil
### GetSepaCc

`func (o *TransactionSplitUpdate) GetSepaCc() string`

GetSepaCc returns the SepaCc field if non-nil, zero value otherwise.

### GetSepaCcOk

`func (o *TransactionSplitUpdate) GetSepaCcOk() (*string, bool)`

GetSepaCcOk returns a tuple with the SepaCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCc

`func (o *TransactionSplitUpdate) SetSepaCc(v string)`

SetSepaCc sets SepaCc field to given value.

### HasSepaCc

`func (o *TransactionSplitUpdate) HasSepaCc() bool`

HasSepaCc returns a boolean if a field has been set.

### SetSepaCcNil

`func (o *TransactionSplitUpdate) SetSepaCcNil(b bool)`

 SetSepaCcNil sets the value for SepaCc to be an explicit nil

### UnsetSepaCc
`func (o *TransactionSplitUpdate) UnsetSepaCc()`

UnsetSepaCc ensures that no value is present for SepaCc, not even an explicit nil
### GetSepaCtOp

`func (o *TransactionSplitUpdate) GetSepaCtOp() string`

GetSepaCtOp returns the SepaCtOp field if non-nil, zero value otherwise.

### GetSepaCtOpOk

`func (o *TransactionSplitUpdate) GetSepaCtOpOk() (*string, bool)`

GetSepaCtOpOk returns a tuple with the SepaCtOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtOp

`func (o *TransactionSplitUpdate) SetSepaCtOp(v string)`

SetSepaCtOp sets SepaCtOp field to given value.

### HasSepaCtOp

`func (o *TransactionSplitUpdate) HasSepaCtOp() bool`

HasSepaCtOp returns a boolean if a field has been set.

### SetSepaCtOpNil

`func (o *TransactionSplitUpdate) SetSepaCtOpNil(b bool)`

 SetSepaCtOpNil sets the value for SepaCtOp to be an explicit nil

### UnsetSepaCtOp
`func (o *TransactionSplitUpdate) UnsetSepaCtOp()`

UnsetSepaCtOp ensures that no value is present for SepaCtOp, not even an explicit nil
### GetSepaCtId

`func (o *TransactionSplitUpdate) GetSepaCtId() string`

GetSepaCtId returns the SepaCtId field if non-nil, zero value otherwise.

### GetSepaCtIdOk

`func (o *TransactionSplitUpdate) GetSepaCtIdOk() (*string, bool)`

GetSepaCtIdOk returns a tuple with the SepaCtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtId

`func (o *TransactionSplitUpdate) SetSepaCtId(v string)`

SetSepaCtId sets SepaCtId field to given value.

### HasSepaCtId

`func (o *TransactionSplitUpdate) HasSepaCtId() bool`

HasSepaCtId returns a boolean if a field has been set.

### SetSepaCtIdNil

`func (o *TransactionSplitUpdate) SetSepaCtIdNil(b bool)`

 SetSepaCtIdNil sets the value for SepaCtId to be an explicit nil

### UnsetSepaCtId
`func (o *TransactionSplitUpdate) UnsetSepaCtId()`

UnsetSepaCtId ensures that no value is present for SepaCtId, not even an explicit nil
### GetSepaDb

`func (o *TransactionSplitUpdate) GetSepaDb() string`

GetSepaDb returns the SepaDb field if non-nil, zero value otherwise.

### GetSepaDbOk

`func (o *TransactionSplitUpdate) GetSepaDbOk() (*string, bool)`

GetSepaDbOk returns a tuple with the SepaDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaDb

`func (o *TransactionSplitUpdate) SetSepaDb(v string)`

SetSepaDb sets SepaDb field to given value.

### HasSepaDb

`func (o *TransactionSplitUpdate) HasSepaDb() bool`

HasSepaDb returns a boolean if a field has been set.

### SetSepaDbNil

`func (o *TransactionSplitUpdate) SetSepaDbNil(b bool)`

 SetSepaDbNil sets the value for SepaDb to be an explicit nil

### UnsetSepaDb
`func (o *TransactionSplitUpdate) UnsetSepaDb()`

UnsetSepaDb ensures that no value is present for SepaDb, not even an explicit nil
### GetSepaCountry

`func (o *TransactionSplitUpdate) GetSepaCountry() string`

GetSepaCountry returns the SepaCountry field if non-nil, zero value otherwise.

### GetSepaCountryOk

`func (o *TransactionSplitUpdate) GetSepaCountryOk() (*string, bool)`

GetSepaCountryOk returns a tuple with the SepaCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCountry

`func (o *TransactionSplitUpdate) SetSepaCountry(v string)`

SetSepaCountry sets SepaCountry field to given value.

### HasSepaCountry

`func (o *TransactionSplitUpdate) HasSepaCountry() bool`

HasSepaCountry returns a boolean if a field has been set.

### SetSepaCountryNil

`func (o *TransactionSplitUpdate) SetSepaCountryNil(b bool)`

 SetSepaCountryNil sets the value for SepaCountry to be an explicit nil

### UnsetSepaCountry
`func (o *TransactionSplitUpdate) UnsetSepaCountry()`

UnsetSepaCountry ensures that no value is present for SepaCountry, not even an explicit nil
### GetSepaEp

`func (o *TransactionSplitUpdate) GetSepaEp() string`

GetSepaEp returns the SepaEp field if non-nil, zero value otherwise.

### GetSepaEpOk

`func (o *TransactionSplitUpdate) GetSepaEpOk() (*string, bool)`

GetSepaEpOk returns a tuple with the SepaEp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaEp

`func (o *TransactionSplitUpdate) SetSepaEp(v string)`

SetSepaEp sets SepaEp field to given value.

### HasSepaEp

`func (o *TransactionSplitUpdate) HasSepaEp() bool`

HasSepaEp returns a boolean if a field has been set.

### SetSepaEpNil

`func (o *TransactionSplitUpdate) SetSepaEpNil(b bool)`

 SetSepaEpNil sets the value for SepaEp to be an explicit nil

### UnsetSepaEp
`func (o *TransactionSplitUpdate) UnsetSepaEp()`

UnsetSepaEp ensures that no value is present for SepaEp, not even an explicit nil
### GetSepaCi

`func (o *TransactionSplitUpdate) GetSepaCi() string`

GetSepaCi returns the SepaCi field if non-nil, zero value otherwise.

### GetSepaCiOk

`func (o *TransactionSplitUpdate) GetSepaCiOk() (*string, bool)`

GetSepaCiOk returns a tuple with the SepaCi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCi

`func (o *TransactionSplitUpdate) SetSepaCi(v string)`

SetSepaCi sets SepaCi field to given value.

### HasSepaCi

`func (o *TransactionSplitUpdate) HasSepaCi() bool`

HasSepaCi returns a boolean if a field has been set.

### SetSepaCiNil

`func (o *TransactionSplitUpdate) SetSepaCiNil(b bool)`

 SetSepaCiNil sets the value for SepaCi to be an explicit nil

### UnsetSepaCi
`func (o *TransactionSplitUpdate) UnsetSepaCi()`

UnsetSepaCi ensures that no value is present for SepaCi, not even an explicit nil
### GetSepaBatchId

`func (o *TransactionSplitUpdate) GetSepaBatchId() string`

GetSepaBatchId returns the SepaBatchId field if non-nil, zero value otherwise.

### GetSepaBatchIdOk

`func (o *TransactionSplitUpdate) GetSepaBatchIdOk() (*string, bool)`

GetSepaBatchIdOk returns a tuple with the SepaBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaBatchId

`func (o *TransactionSplitUpdate) SetSepaBatchId(v string)`

SetSepaBatchId sets SepaBatchId field to given value.

### HasSepaBatchId

`func (o *TransactionSplitUpdate) HasSepaBatchId() bool`

HasSepaBatchId returns a boolean if a field has been set.

### SetSepaBatchIdNil

`func (o *TransactionSplitUpdate) SetSepaBatchIdNil(b bool)`

 SetSepaBatchIdNil sets the value for SepaBatchId to be an explicit nil

### UnsetSepaBatchId
`func (o *TransactionSplitUpdate) UnsetSepaBatchId()`

UnsetSepaBatchId ensures that no value is present for SepaBatchId, not even an explicit nil
### GetInterestDate

`func (o *TransactionSplitUpdate) GetInterestDate() time.Time`

GetInterestDate returns the InterestDate field if non-nil, zero value otherwise.

### GetInterestDateOk

`func (o *TransactionSplitUpdate) GetInterestDateOk() (*time.Time, bool)`

GetInterestDateOk returns a tuple with the InterestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestDate

`func (o *TransactionSplitUpdate) SetInterestDate(v time.Time)`

SetInterestDate sets InterestDate field to given value.

### HasInterestDate

`func (o *TransactionSplitUpdate) HasInterestDate() bool`

HasInterestDate returns a boolean if a field has been set.

### SetInterestDateNil

`func (o *TransactionSplitUpdate) SetInterestDateNil(b bool)`

 SetInterestDateNil sets the value for InterestDate to be an explicit nil

### UnsetInterestDate
`func (o *TransactionSplitUpdate) UnsetInterestDate()`

UnsetInterestDate ensures that no value is present for InterestDate, not even an explicit nil
### GetBookDate

`func (o *TransactionSplitUpdate) GetBookDate() time.Time`

GetBookDate returns the BookDate field if non-nil, zero value otherwise.

### GetBookDateOk

`func (o *TransactionSplitUpdate) GetBookDateOk() (*time.Time, bool)`

GetBookDateOk returns a tuple with the BookDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookDate

`func (o *TransactionSplitUpdate) SetBookDate(v time.Time)`

SetBookDate sets BookDate field to given value.

### HasBookDate

`func (o *TransactionSplitUpdate) HasBookDate() bool`

HasBookDate returns a boolean if a field has been set.

### SetBookDateNil

`func (o *TransactionSplitUpdate) SetBookDateNil(b bool)`

 SetBookDateNil sets the value for BookDate to be an explicit nil

### UnsetBookDate
`func (o *TransactionSplitUpdate) UnsetBookDate()`

UnsetBookDate ensures that no value is present for BookDate, not even an explicit nil
### GetProcessDate

`func (o *TransactionSplitUpdate) GetProcessDate() time.Time`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *TransactionSplitUpdate) GetProcessDateOk() (*time.Time, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *TransactionSplitUpdate) SetProcessDate(v time.Time)`

SetProcessDate sets ProcessDate field to given value.

### HasProcessDate

`func (o *TransactionSplitUpdate) HasProcessDate() bool`

HasProcessDate returns a boolean if a field has been set.

### SetProcessDateNil

`func (o *TransactionSplitUpdate) SetProcessDateNil(b bool)`

 SetProcessDateNil sets the value for ProcessDate to be an explicit nil

### UnsetProcessDate
`func (o *TransactionSplitUpdate) UnsetProcessDate()`

UnsetProcessDate ensures that no value is present for ProcessDate, not even an explicit nil
### GetDueDate

`func (o *TransactionSplitUpdate) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TransactionSplitUpdate) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TransactionSplitUpdate) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TransactionSplitUpdate) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TransactionSplitUpdate) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TransactionSplitUpdate) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentDate

`func (o *TransactionSplitUpdate) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *TransactionSplitUpdate) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *TransactionSplitUpdate) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *TransactionSplitUpdate) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *TransactionSplitUpdate) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *TransactionSplitUpdate) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetInvoiceDate

`func (o *TransactionSplitUpdate) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *TransactionSplitUpdate) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *TransactionSplitUpdate) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *TransactionSplitUpdate) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### SetInvoiceDateNil

`func (o *TransactionSplitUpdate) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *TransactionSplitUpdate) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


