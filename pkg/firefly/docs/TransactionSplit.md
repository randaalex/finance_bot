# TransactionSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | User ID | [optional] [readonly] 
**TransactionJournalId** | Pointer to **int32** | ID of the underlying transaction journal. Each transaction consists of a transaction group (see the top ID) and one or more journals making up the splits of the transaction.  | [optional] [readonly] 
**Type** | **string** | Type of transaction. | 
**Date** | **time.Time** | Date of the transaction | 
**Amount** | **string** | Amount of the transaction. | 
**Description** | **string** | Description of the transaction. | 
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
**SourceId** | **NullableString** | ID of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. | 
**SourceName** | Pointer to **NullableString** | Name of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. Can be used instead of the source_id. If the transaction is a deposit, the source_name can be filled in freely: the account will be created based on the name. | [optional] 
**SourceIban** | Pointer to **NullableString** |  | [optional] [readonly] 
**SourceType** | Pointer to [**AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] 
**DestinationId** | **NullableString** | ID of the destination account. For a deposit or a transfer, this must always be an asset account. For withdrawals this must be an expense account. | 
**DestinationName** | Pointer to **NullableString** | Name of the destination account. You can submit the name instead of the ID. For everything except transfers, the account will be auto-generated if unknown, so submitting a name is enough. | [optional] 
**DestinationIban** | Pointer to **NullableString** |  | [optional] [readonly] 
**DestinationType** | Pointer to [**AccountTypeProperty**](AccountTypeProperty.md) |  | [optional] 
**Reconciled** | Pointer to **bool** | If the transaction has been reconciled already. When you set this, the amount can no longer be edited by the user. | [optional] 
**PiggyBankId** | Pointer to **int32** | Optional. Use either this or the piggy_bank_name | [optional] 
**PiggyBankName** | Pointer to **string** | Optional. Use either this or the piggy_bank_id | [optional] 
**BillId** | Pointer to **NullableString** | Optional. Use either this or the bill_name | [optional] 
**BillName** | Pointer to **NullableString** | Optional. Use either this or the bill_id | [optional] 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**InternalReference** | Pointer to **NullableString** | Reference to internal reference of other systems. | [optional] 
**ExternalId** | Pointer to **NullableString** | Reference to external ID in other systems. | [optional] 
**OriginalSource** | Pointer to **NullableString** | System generated identifier for original creator of transaction. | [optional] [readonly] 
**RecurrenceId** | Pointer to **NullableInt32** | Reference to recurrence that made the transaction. | [optional] [readonly] 
**RecurrenceTotal** | Pointer to **NullableInt32** | Total number of transactions expected to be created by this recurrence repetition. Will be 0 if infinite. | [optional] [readonly] 
**RecurrenceCount** | Pointer to **NullableInt32** | The # of the current transaction created under this recurrence. | [optional] [readonly] 
**BunqPaymentId** | Pointer to **NullableString** | Internal ID of bunq transaction. | [optional] 
**ImportHashV2** | Pointer to **NullableString** | Hash value of original import transaction (for duplicate detection). | [optional] [readonly] 
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

### NewTransactionSplit

`func NewTransactionSplit(type_ string, date time.Time, amount string, description string, sourceId NullableString, destinationId NullableString, ) *TransactionSplit`

NewTransactionSplit instantiates a new TransactionSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSplitWithDefaults

`func NewTransactionSplitWithDefaults() *TransactionSplit`

NewTransactionSplitWithDefaults instantiates a new TransactionSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *TransactionSplit) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionSplit) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionSplit) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *TransactionSplit) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTransactionJournalId

`func (o *TransactionSplit) GetTransactionJournalId() int32`

GetTransactionJournalId returns the TransactionJournalId field if non-nil, zero value otherwise.

### GetTransactionJournalIdOk

`func (o *TransactionSplit) GetTransactionJournalIdOk() (*int32, bool)`

GetTransactionJournalIdOk returns a tuple with the TransactionJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionJournalId

`func (o *TransactionSplit) SetTransactionJournalId(v int32)`

SetTransactionJournalId sets TransactionJournalId field to given value.

### HasTransactionJournalId

`func (o *TransactionSplit) HasTransactionJournalId() bool`

HasTransactionJournalId returns a boolean if a field has been set.

### GetType

`func (o *TransactionSplit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSplit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSplit) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *TransactionSplit) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionSplit) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionSplit) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetAmount

`func (o *TransactionSplit) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSplit) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSplit) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *TransactionSplit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionSplit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionSplit) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOrder

`func (o *TransactionSplit) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TransactionSplit) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TransactionSplit) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TransactionSplit) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *TransactionSplit) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *TransactionSplit) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetCurrencyId

`func (o *TransactionSplit) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TransactionSplit) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TransactionSplit) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TransactionSplit) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TransactionSplit) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TransactionSplit) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionSplit) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionSplit) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionSplit) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionSplit) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionSplit) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionSplit) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetCurrencySymbol

`func (o *TransactionSplit) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *TransactionSplit) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *TransactionSplit) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *TransactionSplit) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyName

`func (o *TransactionSplit) GetCurrencyName() string`

GetCurrencyName returns the CurrencyName field if non-nil, zero value otherwise.

### GetCurrencyNameOk

`func (o *TransactionSplit) GetCurrencyNameOk() (*string, bool)`

GetCurrencyNameOk returns a tuple with the CurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyName

`func (o *TransactionSplit) SetCurrencyName(v string)`

SetCurrencyName sets CurrencyName field to given value.

### HasCurrencyName

`func (o *TransactionSplit) HasCurrencyName() bool`

HasCurrencyName returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *TransactionSplit) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *TransactionSplit) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *TransactionSplit) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *TransactionSplit) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetForeignAmount

`func (o *TransactionSplit) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *TransactionSplit) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *TransactionSplit) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *TransactionSplit) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *TransactionSplit) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *TransactionSplit) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetForeignCurrencyId

`func (o *TransactionSplit) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *TransactionSplit) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *TransactionSplit) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *TransactionSplit) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *TransactionSplit) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *TransactionSplit) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetForeignCurrencyCode

`func (o *TransactionSplit) GetForeignCurrencyCode() string`

GetForeignCurrencyCode returns the ForeignCurrencyCode field if non-nil, zero value otherwise.

### GetForeignCurrencyCodeOk

`func (o *TransactionSplit) GetForeignCurrencyCodeOk() (*string, bool)`

GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyCode

`func (o *TransactionSplit) SetForeignCurrencyCode(v string)`

SetForeignCurrencyCode sets ForeignCurrencyCode field to given value.

### HasForeignCurrencyCode

`func (o *TransactionSplit) HasForeignCurrencyCode() bool`

HasForeignCurrencyCode returns a boolean if a field has been set.

### SetForeignCurrencyCodeNil

`func (o *TransactionSplit) SetForeignCurrencyCodeNil(b bool)`

 SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil

### UnsetForeignCurrencyCode
`func (o *TransactionSplit) UnsetForeignCurrencyCode()`

UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
### GetForeignCurrencySymbol

`func (o *TransactionSplit) GetForeignCurrencySymbol() string`

GetForeignCurrencySymbol returns the ForeignCurrencySymbol field if non-nil, zero value otherwise.

### GetForeignCurrencySymbolOk

`func (o *TransactionSplit) GetForeignCurrencySymbolOk() (*string, bool)`

GetForeignCurrencySymbolOk returns a tuple with the ForeignCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencySymbol

`func (o *TransactionSplit) SetForeignCurrencySymbol(v string)`

SetForeignCurrencySymbol sets ForeignCurrencySymbol field to given value.

### HasForeignCurrencySymbol

`func (o *TransactionSplit) HasForeignCurrencySymbol() bool`

HasForeignCurrencySymbol returns a boolean if a field has been set.

### SetForeignCurrencySymbolNil

`func (o *TransactionSplit) SetForeignCurrencySymbolNil(b bool)`

 SetForeignCurrencySymbolNil sets the value for ForeignCurrencySymbol to be an explicit nil

### UnsetForeignCurrencySymbol
`func (o *TransactionSplit) UnsetForeignCurrencySymbol()`

UnsetForeignCurrencySymbol ensures that no value is present for ForeignCurrencySymbol, not even an explicit nil
### GetForeignCurrencyDecimalPlaces

`func (o *TransactionSplit) GetForeignCurrencyDecimalPlaces() int32`

GetForeignCurrencyDecimalPlaces returns the ForeignCurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetForeignCurrencyDecimalPlacesOk

`func (o *TransactionSplit) GetForeignCurrencyDecimalPlacesOk() (*int32, bool)`

GetForeignCurrencyDecimalPlacesOk returns a tuple with the ForeignCurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyDecimalPlaces

`func (o *TransactionSplit) SetForeignCurrencyDecimalPlaces(v int32)`

SetForeignCurrencyDecimalPlaces sets ForeignCurrencyDecimalPlaces field to given value.

### HasForeignCurrencyDecimalPlaces

`func (o *TransactionSplit) HasForeignCurrencyDecimalPlaces() bool`

HasForeignCurrencyDecimalPlaces returns a boolean if a field has been set.

### SetForeignCurrencyDecimalPlacesNil

`func (o *TransactionSplit) SetForeignCurrencyDecimalPlacesNil(b bool)`

 SetForeignCurrencyDecimalPlacesNil sets the value for ForeignCurrencyDecimalPlaces to be an explicit nil

### UnsetForeignCurrencyDecimalPlaces
`func (o *TransactionSplit) UnsetForeignCurrencyDecimalPlaces()`

UnsetForeignCurrencyDecimalPlaces ensures that no value is present for ForeignCurrencyDecimalPlaces, not even an explicit nil
### GetBudgetId

`func (o *TransactionSplit) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *TransactionSplit) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *TransactionSplit) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *TransactionSplit) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *TransactionSplit) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *TransactionSplit) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetBudgetName

`func (o *TransactionSplit) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *TransactionSplit) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *TransactionSplit) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.

### HasBudgetName

`func (o *TransactionSplit) HasBudgetName() bool`

HasBudgetName returns a boolean if a field has been set.

### SetBudgetNameNil

`func (o *TransactionSplit) SetBudgetNameNil(b bool)`

 SetBudgetNameNil sets the value for BudgetName to be an explicit nil

### UnsetBudgetName
`func (o *TransactionSplit) UnsetBudgetName()`

UnsetBudgetName ensures that no value is present for BudgetName, not even an explicit nil
### GetCategoryId

`func (o *TransactionSplit) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransactionSplit) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransactionSplit) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *TransactionSplit) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *TransactionSplit) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *TransactionSplit) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryName

`func (o *TransactionSplit) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *TransactionSplit) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *TransactionSplit) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *TransactionSplit) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### SetCategoryNameNil

`func (o *TransactionSplit) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *TransactionSplit) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
### GetSourceId

`func (o *TransactionSplit) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TransactionSplit) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TransactionSplit) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *TransactionSplit) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *TransactionSplit) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *TransactionSplit) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TransactionSplit) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TransactionSplit) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TransactionSplit) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TransactionSplit) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TransactionSplit) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceIban

`func (o *TransactionSplit) GetSourceIban() string`

GetSourceIban returns the SourceIban field if non-nil, zero value otherwise.

### GetSourceIbanOk

`func (o *TransactionSplit) GetSourceIbanOk() (*string, bool)`

GetSourceIbanOk returns a tuple with the SourceIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIban

`func (o *TransactionSplit) SetSourceIban(v string)`

SetSourceIban sets SourceIban field to given value.

### HasSourceIban

`func (o *TransactionSplit) HasSourceIban() bool`

HasSourceIban returns a boolean if a field has been set.

### SetSourceIbanNil

`func (o *TransactionSplit) SetSourceIbanNil(b bool)`

 SetSourceIbanNil sets the value for SourceIban to be an explicit nil

### UnsetSourceIban
`func (o *TransactionSplit) UnsetSourceIban()`

UnsetSourceIban ensures that no value is present for SourceIban, not even an explicit nil
### GetSourceType

`func (o *TransactionSplit) GetSourceType() AccountTypeProperty`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TransactionSplit) GetSourceTypeOk() (*AccountTypeProperty, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TransactionSplit) SetSourceType(v AccountTypeProperty)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *TransactionSplit) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationId

`func (o *TransactionSplit) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *TransactionSplit) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *TransactionSplit) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### SetDestinationIdNil

`func (o *TransactionSplit) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *TransactionSplit) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetDestinationName

`func (o *TransactionSplit) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *TransactionSplit) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *TransactionSplit) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *TransactionSplit) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### SetDestinationNameNil

`func (o *TransactionSplit) SetDestinationNameNil(b bool)`

 SetDestinationNameNil sets the value for DestinationName to be an explicit nil

### UnsetDestinationName
`func (o *TransactionSplit) UnsetDestinationName()`

UnsetDestinationName ensures that no value is present for DestinationName, not even an explicit nil
### GetDestinationIban

`func (o *TransactionSplit) GetDestinationIban() string`

GetDestinationIban returns the DestinationIban field if non-nil, zero value otherwise.

### GetDestinationIbanOk

`func (o *TransactionSplit) GetDestinationIbanOk() (*string, bool)`

GetDestinationIbanOk returns a tuple with the DestinationIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIban

`func (o *TransactionSplit) SetDestinationIban(v string)`

SetDestinationIban sets DestinationIban field to given value.

### HasDestinationIban

`func (o *TransactionSplit) HasDestinationIban() bool`

HasDestinationIban returns a boolean if a field has been set.

### SetDestinationIbanNil

`func (o *TransactionSplit) SetDestinationIbanNil(b bool)`

 SetDestinationIbanNil sets the value for DestinationIban to be an explicit nil

### UnsetDestinationIban
`func (o *TransactionSplit) UnsetDestinationIban()`

UnsetDestinationIban ensures that no value is present for DestinationIban, not even an explicit nil
### GetDestinationType

`func (o *TransactionSplit) GetDestinationType() AccountTypeProperty`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransactionSplit) GetDestinationTypeOk() (*AccountTypeProperty, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransactionSplit) SetDestinationType(v AccountTypeProperty)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *TransactionSplit) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetReconciled

`func (o *TransactionSplit) GetReconciled() bool`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *TransactionSplit) GetReconciledOk() (*bool, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *TransactionSplit) SetReconciled(v bool)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *TransactionSplit) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetPiggyBankId

`func (o *TransactionSplit) GetPiggyBankId() int32`

GetPiggyBankId returns the PiggyBankId field if non-nil, zero value otherwise.

### GetPiggyBankIdOk

`func (o *TransactionSplit) GetPiggyBankIdOk() (*int32, bool)`

GetPiggyBankIdOk returns a tuple with the PiggyBankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankId

`func (o *TransactionSplit) SetPiggyBankId(v int32)`

SetPiggyBankId sets PiggyBankId field to given value.

### HasPiggyBankId

`func (o *TransactionSplit) HasPiggyBankId() bool`

HasPiggyBankId returns a boolean if a field has been set.

### GetPiggyBankName

`func (o *TransactionSplit) GetPiggyBankName() string`

GetPiggyBankName returns the PiggyBankName field if non-nil, zero value otherwise.

### GetPiggyBankNameOk

`func (o *TransactionSplit) GetPiggyBankNameOk() (*string, bool)`

GetPiggyBankNameOk returns a tuple with the PiggyBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankName

`func (o *TransactionSplit) SetPiggyBankName(v string)`

SetPiggyBankName sets PiggyBankName field to given value.

### HasPiggyBankName

`func (o *TransactionSplit) HasPiggyBankName() bool`

HasPiggyBankName returns a boolean if a field has been set.

### GetBillId

`func (o *TransactionSplit) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *TransactionSplit) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *TransactionSplit) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *TransactionSplit) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### SetBillIdNil

`func (o *TransactionSplit) SetBillIdNil(b bool)`

 SetBillIdNil sets the value for BillId to be an explicit nil

### UnsetBillId
`func (o *TransactionSplit) UnsetBillId()`

UnsetBillId ensures that no value is present for BillId, not even an explicit nil
### GetBillName

`func (o *TransactionSplit) GetBillName() string`

GetBillName returns the BillName field if non-nil, zero value otherwise.

### GetBillNameOk

`func (o *TransactionSplit) GetBillNameOk() (*string, bool)`

GetBillNameOk returns a tuple with the BillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillName

`func (o *TransactionSplit) SetBillName(v string)`

SetBillName sets BillName field to given value.

### HasBillName

`func (o *TransactionSplit) HasBillName() bool`

HasBillName returns a boolean if a field has been set.

### SetBillNameNil

`func (o *TransactionSplit) SetBillNameNil(b bool)`

 SetBillNameNil sets the value for BillName to be an explicit nil

### UnsetBillName
`func (o *TransactionSplit) UnsetBillName()`

UnsetBillName ensures that no value is present for BillName, not even an explicit nil
### GetTags

`func (o *TransactionSplit) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransactionSplit) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransactionSplit) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TransactionSplit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TransactionSplit) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TransactionSplit) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetNotes

`func (o *TransactionSplit) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionSplit) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionSplit) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionSplit) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionSplit) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionSplit) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInternalReference

`func (o *TransactionSplit) GetInternalReference() string`

GetInternalReference returns the InternalReference field if non-nil, zero value otherwise.

### GetInternalReferenceOk

`func (o *TransactionSplit) GetInternalReferenceOk() (*string, bool)`

GetInternalReferenceOk returns a tuple with the InternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalReference

`func (o *TransactionSplit) SetInternalReference(v string)`

SetInternalReference sets InternalReference field to given value.

### HasInternalReference

`func (o *TransactionSplit) HasInternalReference() bool`

HasInternalReference returns a boolean if a field has been set.

### SetInternalReferenceNil

`func (o *TransactionSplit) SetInternalReferenceNil(b bool)`

 SetInternalReferenceNil sets the value for InternalReference to be an explicit nil

### UnsetInternalReference
`func (o *TransactionSplit) UnsetInternalReference()`

UnsetInternalReference ensures that no value is present for InternalReference, not even an explicit nil
### GetExternalId

`func (o *TransactionSplit) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TransactionSplit) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TransactionSplit) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TransactionSplit) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TransactionSplit) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TransactionSplit) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetOriginalSource

`func (o *TransactionSplit) GetOriginalSource() string`

GetOriginalSource returns the OriginalSource field if non-nil, zero value otherwise.

### GetOriginalSourceOk

`func (o *TransactionSplit) GetOriginalSourceOk() (*string, bool)`

GetOriginalSourceOk returns a tuple with the OriginalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSource

`func (o *TransactionSplit) SetOriginalSource(v string)`

SetOriginalSource sets OriginalSource field to given value.

### HasOriginalSource

`func (o *TransactionSplit) HasOriginalSource() bool`

HasOriginalSource returns a boolean if a field has been set.

### SetOriginalSourceNil

`func (o *TransactionSplit) SetOriginalSourceNil(b bool)`

 SetOriginalSourceNil sets the value for OriginalSource to be an explicit nil

### UnsetOriginalSource
`func (o *TransactionSplit) UnsetOriginalSource()`

UnsetOriginalSource ensures that no value is present for OriginalSource, not even an explicit nil
### GetRecurrenceId

`func (o *TransactionSplit) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *TransactionSplit) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *TransactionSplit) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *TransactionSplit) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### SetRecurrenceIdNil

`func (o *TransactionSplit) SetRecurrenceIdNil(b bool)`

 SetRecurrenceIdNil sets the value for RecurrenceId to be an explicit nil

### UnsetRecurrenceId
`func (o *TransactionSplit) UnsetRecurrenceId()`

UnsetRecurrenceId ensures that no value is present for RecurrenceId, not even an explicit nil
### GetRecurrenceTotal

`func (o *TransactionSplit) GetRecurrenceTotal() int32`

GetRecurrenceTotal returns the RecurrenceTotal field if non-nil, zero value otherwise.

### GetRecurrenceTotalOk

`func (o *TransactionSplit) GetRecurrenceTotalOk() (*int32, bool)`

GetRecurrenceTotalOk returns a tuple with the RecurrenceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceTotal

`func (o *TransactionSplit) SetRecurrenceTotal(v int32)`

SetRecurrenceTotal sets RecurrenceTotal field to given value.

### HasRecurrenceTotal

`func (o *TransactionSplit) HasRecurrenceTotal() bool`

HasRecurrenceTotal returns a boolean if a field has been set.

### SetRecurrenceTotalNil

`func (o *TransactionSplit) SetRecurrenceTotalNil(b bool)`

 SetRecurrenceTotalNil sets the value for RecurrenceTotal to be an explicit nil

### UnsetRecurrenceTotal
`func (o *TransactionSplit) UnsetRecurrenceTotal()`

UnsetRecurrenceTotal ensures that no value is present for RecurrenceTotal, not even an explicit nil
### GetRecurrenceCount

`func (o *TransactionSplit) GetRecurrenceCount() int32`

GetRecurrenceCount returns the RecurrenceCount field if non-nil, zero value otherwise.

### GetRecurrenceCountOk

`func (o *TransactionSplit) GetRecurrenceCountOk() (*int32, bool)`

GetRecurrenceCountOk returns a tuple with the RecurrenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceCount

`func (o *TransactionSplit) SetRecurrenceCount(v int32)`

SetRecurrenceCount sets RecurrenceCount field to given value.

### HasRecurrenceCount

`func (o *TransactionSplit) HasRecurrenceCount() bool`

HasRecurrenceCount returns a boolean if a field has been set.

### SetRecurrenceCountNil

`func (o *TransactionSplit) SetRecurrenceCountNil(b bool)`

 SetRecurrenceCountNil sets the value for RecurrenceCount to be an explicit nil

### UnsetRecurrenceCount
`func (o *TransactionSplit) UnsetRecurrenceCount()`

UnsetRecurrenceCount ensures that no value is present for RecurrenceCount, not even an explicit nil
### GetBunqPaymentId

`func (o *TransactionSplit) GetBunqPaymentId() string`

GetBunqPaymentId returns the BunqPaymentId field if non-nil, zero value otherwise.

### GetBunqPaymentIdOk

`func (o *TransactionSplit) GetBunqPaymentIdOk() (*string, bool)`

GetBunqPaymentIdOk returns a tuple with the BunqPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBunqPaymentId

`func (o *TransactionSplit) SetBunqPaymentId(v string)`

SetBunqPaymentId sets BunqPaymentId field to given value.

### HasBunqPaymentId

`func (o *TransactionSplit) HasBunqPaymentId() bool`

HasBunqPaymentId returns a boolean if a field has been set.

### SetBunqPaymentIdNil

`func (o *TransactionSplit) SetBunqPaymentIdNil(b bool)`

 SetBunqPaymentIdNil sets the value for BunqPaymentId to be an explicit nil

### UnsetBunqPaymentId
`func (o *TransactionSplit) UnsetBunqPaymentId()`

UnsetBunqPaymentId ensures that no value is present for BunqPaymentId, not even an explicit nil
### GetImportHashV2

`func (o *TransactionSplit) GetImportHashV2() string`

GetImportHashV2 returns the ImportHashV2 field if non-nil, zero value otherwise.

### GetImportHashV2Ok

`func (o *TransactionSplit) GetImportHashV2Ok() (*string, bool)`

GetImportHashV2Ok returns a tuple with the ImportHashV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportHashV2

`func (o *TransactionSplit) SetImportHashV2(v string)`

SetImportHashV2 sets ImportHashV2 field to given value.

### HasImportHashV2

`func (o *TransactionSplit) HasImportHashV2() bool`

HasImportHashV2 returns a boolean if a field has been set.

### SetImportHashV2Nil

`func (o *TransactionSplit) SetImportHashV2Nil(b bool)`

 SetImportHashV2Nil sets the value for ImportHashV2 to be an explicit nil

### UnsetImportHashV2
`func (o *TransactionSplit) UnsetImportHashV2()`

UnsetImportHashV2 ensures that no value is present for ImportHashV2, not even an explicit nil
### GetSepaCc

`func (o *TransactionSplit) GetSepaCc() string`

GetSepaCc returns the SepaCc field if non-nil, zero value otherwise.

### GetSepaCcOk

`func (o *TransactionSplit) GetSepaCcOk() (*string, bool)`

GetSepaCcOk returns a tuple with the SepaCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCc

`func (o *TransactionSplit) SetSepaCc(v string)`

SetSepaCc sets SepaCc field to given value.

### HasSepaCc

`func (o *TransactionSplit) HasSepaCc() bool`

HasSepaCc returns a boolean if a field has been set.

### SetSepaCcNil

`func (o *TransactionSplit) SetSepaCcNil(b bool)`

 SetSepaCcNil sets the value for SepaCc to be an explicit nil

### UnsetSepaCc
`func (o *TransactionSplit) UnsetSepaCc()`

UnsetSepaCc ensures that no value is present for SepaCc, not even an explicit nil
### GetSepaCtOp

`func (o *TransactionSplit) GetSepaCtOp() string`

GetSepaCtOp returns the SepaCtOp field if non-nil, zero value otherwise.

### GetSepaCtOpOk

`func (o *TransactionSplit) GetSepaCtOpOk() (*string, bool)`

GetSepaCtOpOk returns a tuple with the SepaCtOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtOp

`func (o *TransactionSplit) SetSepaCtOp(v string)`

SetSepaCtOp sets SepaCtOp field to given value.

### HasSepaCtOp

`func (o *TransactionSplit) HasSepaCtOp() bool`

HasSepaCtOp returns a boolean if a field has been set.

### SetSepaCtOpNil

`func (o *TransactionSplit) SetSepaCtOpNil(b bool)`

 SetSepaCtOpNil sets the value for SepaCtOp to be an explicit nil

### UnsetSepaCtOp
`func (o *TransactionSplit) UnsetSepaCtOp()`

UnsetSepaCtOp ensures that no value is present for SepaCtOp, not even an explicit nil
### GetSepaCtId

`func (o *TransactionSplit) GetSepaCtId() string`

GetSepaCtId returns the SepaCtId field if non-nil, zero value otherwise.

### GetSepaCtIdOk

`func (o *TransactionSplit) GetSepaCtIdOk() (*string, bool)`

GetSepaCtIdOk returns a tuple with the SepaCtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtId

`func (o *TransactionSplit) SetSepaCtId(v string)`

SetSepaCtId sets SepaCtId field to given value.

### HasSepaCtId

`func (o *TransactionSplit) HasSepaCtId() bool`

HasSepaCtId returns a boolean if a field has been set.

### SetSepaCtIdNil

`func (o *TransactionSplit) SetSepaCtIdNil(b bool)`

 SetSepaCtIdNil sets the value for SepaCtId to be an explicit nil

### UnsetSepaCtId
`func (o *TransactionSplit) UnsetSepaCtId()`

UnsetSepaCtId ensures that no value is present for SepaCtId, not even an explicit nil
### GetSepaDb

`func (o *TransactionSplit) GetSepaDb() string`

GetSepaDb returns the SepaDb field if non-nil, zero value otherwise.

### GetSepaDbOk

`func (o *TransactionSplit) GetSepaDbOk() (*string, bool)`

GetSepaDbOk returns a tuple with the SepaDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaDb

`func (o *TransactionSplit) SetSepaDb(v string)`

SetSepaDb sets SepaDb field to given value.

### HasSepaDb

`func (o *TransactionSplit) HasSepaDb() bool`

HasSepaDb returns a boolean if a field has been set.

### SetSepaDbNil

`func (o *TransactionSplit) SetSepaDbNil(b bool)`

 SetSepaDbNil sets the value for SepaDb to be an explicit nil

### UnsetSepaDb
`func (o *TransactionSplit) UnsetSepaDb()`

UnsetSepaDb ensures that no value is present for SepaDb, not even an explicit nil
### GetSepaCountry

`func (o *TransactionSplit) GetSepaCountry() string`

GetSepaCountry returns the SepaCountry field if non-nil, zero value otherwise.

### GetSepaCountryOk

`func (o *TransactionSplit) GetSepaCountryOk() (*string, bool)`

GetSepaCountryOk returns a tuple with the SepaCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCountry

`func (o *TransactionSplit) SetSepaCountry(v string)`

SetSepaCountry sets SepaCountry field to given value.

### HasSepaCountry

`func (o *TransactionSplit) HasSepaCountry() bool`

HasSepaCountry returns a boolean if a field has been set.

### SetSepaCountryNil

`func (o *TransactionSplit) SetSepaCountryNil(b bool)`

 SetSepaCountryNil sets the value for SepaCountry to be an explicit nil

### UnsetSepaCountry
`func (o *TransactionSplit) UnsetSepaCountry()`

UnsetSepaCountry ensures that no value is present for SepaCountry, not even an explicit nil
### GetSepaEp

`func (o *TransactionSplit) GetSepaEp() string`

GetSepaEp returns the SepaEp field if non-nil, zero value otherwise.

### GetSepaEpOk

`func (o *TransactionSplit) GetSepaEpOk() (*string, bool)`

GetSepaEpOk returns a tuple with the SepaEp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaEp

`func (o *TransactionSplit) SetSepaEp(v string)`

SetSepaEp sets SepaEp field to given value.

### HasSepaEp

`func (o *TransactionSplit) HasSepaEp() bool`

HasSepaEp returns a boolean if a field has been set.

### SetSepaEpNil

`func (o *TransactionSplit) SetSepaEpNil(b bool)`

 SetSepaEpNil sets the value for SepaEp to be an explicit nil

### UnsetSepaEp
`func (o *TransactionSplit) UnsetSepaEp()`

UnsetSepaEp ensures that no value is present for SepaEp, not even an explicit nil
### GetSepaCi

`func (o *TransactionSplit) GetSepaCi() string`

GetSepaCi returns the SepaCi field if non-nil, zero value otherwise.

### GetSepaCiOk

`func (o *TransactionSplit) GetSepaCiOk() (*string, bool)`

GetSepaCiOk returns a tuple with the SepaCi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCi

`func (o *TransactionSplit) SetSepaCi(v string)`

SetSepaCi sets SepaCi field to given value.

### HasSepaCi

`func (o *TransactionSplit) HasSepaCi() bool`

HasSepaCi returns a boolean if a field has been set.

### SetSepaCiNil

`func (o *TransactionSplit) SetSepaCiNil(b bool)`

 SetSepaCiNil sets the value for SepaCi to be an explicit nil

### UnsetSepaCi
`func (o *TransactionSplit) UnsetSepaCi()`

UnsetSepaCi ensures that no value is present for SepaCi, not even an explicit nil
### GetSepaBatchId

`func (o *TransactionSplit) GetSepaBatchId() string`

GetSepaBatchId returns the SepaBatchId field if non-nil, zero value otherwise.

### GetSepaBatchIdOk

`func (o *TransactionSplit) GetSepaBatchIdOk() (*string, bool)`

GetSepaBatchIdOk returns a tuple with the SepaBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaBatchId

`func (o *TransactionSplit) SetSepaBatchId(v string)`

SetSepaBatchId sets SepaBatchId field to given value.

### HasSepaBatchId

`func (o *TransactionSplit) HasSepaBatchId() bool`

HasSepaBatchId returns a boolean if a field has been set.

### SetSepaBatchIdNil

`func (o *TransactionSplit) SetSepaBatchIdNil(b bool)`

 SetSepaBatchIdNil sets the value for SepaBatchId to be an explicit nil

### UnsetSepaBatchId
`func (o *TransactionSplit) UnsetSepaBatchId()`

UnsetSepaBatchId ensures that no value is present for SepaBatchId, not even an explicit nil
### GetInterestDate

`func (o *TransactionSplit) GetInterestDate() time.Time`

GetInterestDate returns the InterestDate field if non-nil, zero value otherwise.

### GetInterestDateOk

`func (o *TransactionSplit) GetInterestDateOk() (*time.Time, bool)`

GetInterestDateOk returns a tuple with the InterestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestDate

`func (o *TransactionSplit) SetInterestDate(v time.Time)`

SetInterestDate sets InterestDate field to given value.

### HasInterestDate

`func (o *TransactionSplit) HasInterestDate() bool`

HasInterestDate returns a boolean if a field has been set.

### SetInterestDateNil

`func (o *TransactionSplit) SetInterestDateNil(b bool)`

 SetInterestDateNil sets the value for InterestDate to be an explicit nil

### UnsetInterestDate
`func (o *TransactionSplit) UnsetInterestDate()`

UnsetInterestDate ensures that no value is present for InterestDate, not even an explicit nil
### GetBookDate

`func (o *TransactionSplit) GetBookDate() time.Time`

GetBookDate returns the BookDate field if non-nil, zero value otherwise.

### GetBookDateOk

`func (o *TransactionSplit) GetBookDateOk() (*time.Time, bool)`

GetBookDateOk returns a tuple with the BookDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookDate

`func (o *TransactionSplit) SetBookDate(v time.Time)`

SetBookDate sets BookDate field to given value.

### HasBookDate

`func (o *TransactionSplit) HasBookDate() bool`

HasBookDate returns a boolean if a field has been set.

### SetBookDateNil

`func (o *TransactionSplit) SetBookDateNil(b bool)`

 SetBookDateNil sets the value for BookDate to be an explicit nil

### UnsetBookDate
`func (o *TransactionSplit) UnsetBookDate()`

UnsetBookDate ensures that no value is present for BookDate, not even an explicit nil
### GetProcessDate

`func (o *TransactionSplit) GetProcessDate() time.Time`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *TransactionSplit) GetProcessDateOk() (*time.Time, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *TransactionSplit) SetProcessDate(v time.Time)`

SetProcessDate sets ProcessDate field to given value.

### HasProcessDate

`func (o *TransactionSplit) HasProcessDate() bool`

HasProcessDate returns a boolean if a field has been set.

### SetProcessDateNil

`func (o *TransactionSplit) SetProcessDateNil(b bool)`

 SetProcessDateNil sets the value for ProcessDate to be an explicit nil

### UnsetProcessDate
`func (o *TransactionSplit) UnsetProcessDate()`

UnsetProcessDate ensures that no value is present for ProcessDate, not even an explicit nil
### GetDueDate

`func (o *TransactionSplit) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TransactionSplit) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TransactionSplit) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TransactionSplit) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TransactionSplit) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TransactionSplit) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentDate

`func (o *TransactionSplit) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *TransactionSplit) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *TransactionSplit) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *TransactionSplit) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *TransactionSplit) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *TransactionSplit) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetInvoiceDate

`func (o *TransactionSplit) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *TransactionSplit) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *TransactionSplit) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *TransactionSplit) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### SetInvoiceDateNil

`func (o *TransactionSplit) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *TransactionSplit) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


