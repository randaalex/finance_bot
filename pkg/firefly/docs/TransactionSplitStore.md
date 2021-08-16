# TransactionSplitStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of transaction. | 
**Date** | **time.Time** | Date of the transaction | 
**Amount** | **string** | Amount of the transaction. | 
**Description** | **string** | Description of the transaction. | 
**Order** | Pointer to **NullableInt32** | Order of this entry in the list of transactions. | [optional] 
**CurrencyId** | Pointer to **NullableString** | Currency ID. Default is the source account&#39;s currency, or the user&#39;s default currency. The value you submit may be overruled by the source or destination account. | [optional] 
**CurrencyCode** | Pointer to **NullableString** | Currency code. Default is the source account&#39;s currency, or the user&#39;s default currency. The value you submit may be overruled by the source or destination account. | [optional] 
**ForeignAmount** | Pointer to **NullableString** | The amount in a foreign currency. | [optional] 
**ForeignCurrencyId** | Pointer to **NullableString** | Currency ID of the foreign currency. Default is null. Is required when you submit a foreign amount. | [optional] 
**ForeignCurrencyCode** | Pointer to **NullableString** | Currency code of the foreign currency. Default is NULL. Can be used instead of the foreign_currency_id, but this or the ID is required when submitting a foreign amount. | [optional] 
**BudgetId** | Pointer to **NullableString** | The budget ID for this transaction. | [optional] 
**BudgetName** | Pointer to **NullableString** | The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored. | [optional] [readonly] 
**CategoryId** | Pointer to **NullableString** | The category ID for this transaction. | [optional] 
**CategoryName** | Pointer to **NullableString** | The name of the category to be used. If the category is unknown, it will be created. If the ID and the name point to different categories, the ID overrules the name. | [optional] 
**SourceId** | **NullableString** | ID of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. | 
**SourceName** | Pointer to **NullableString** | Name of the source account. For a withdrawal or a transfer, this must always be an asset account. For deposits, this must be a revenue account. Can be used instead of the source_id. If the transaction is a deposit, the source_name can be filled in freely: the account will be created based on the name. | [optional] 
**DestinationId** | **NullableString** | ID of the destination account. For a deposit or a transfer, this must always be an asset account. For withdrawals this must be an expense account. | 
**DestinationName** | Pointer to **NullableString** | Name of the destination account. You can submit the name instead of the ID. For everything except transfers, the account will be auto-generated if unknown, so submitting a name is enough. | [optional] 
**Reconciled** | Pointer to **bool** | If the transaction has been reconciled already. When you set this, the amount can no longer be edited by the user. | [optional] 
**PiggyBankId** | Pointer to **int32** | Optional. Use either this or the piggy_bank_name | [optional] 
**PiggyBankName** | Pointer to **string** | Optional. Use either this or the piggy_bank_id | [optional] 
**BillId** | Pointer to **NullableString** | Optional. Use either this or the bill_name | [optional] 
**BillName** | Pointer to **NullableString** | Optional. Use either this or the bill_id | [optional] 
**Tags** | Pointer to **[]string** | Array of tags. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**InternalReference** | Pointer to **NullableString** | Reference to internal reference of other systems. | [optional] 
**ExternalId** | Pointer to **NullableString** | Reference to external ID in other systems. | [optional] 
**BunqPaymentId** | Pointer to **NullableString** | Internal ID of bunq transaction. Field is no longer used but still works. | [optional] 
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

### NewTransactionSplitStore

`func NewTransactionSplitStore(type_ string, date time.Time, amount string, description string, sourceId NullableString, destinationId NullableString, ) *TransactionSplitStore`

NewTransactionSplitStore instantiates a new TransactionSplitStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSplitStoreWithDefaults

`func NewTransactionSplitStoreWithDefaults() *TransactionSplitStore`

NewTransactionSplitStoreWithDefaults instantiates a new TransactionSplitStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionSplitStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSplitStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSplitStore) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *TransactionSplitStore) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionSplitStore) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionSplitStore) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetAmount

`func (o *TransactionSplitStore) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSplitStore) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSplitStore) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *TransactionSplitStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionSplitStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionSplitStore) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOrder

`func (o *TransactionSplitStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TransactionSplitStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TransactionSplitStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TransactionSplitStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *TransactionSplitStore) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *TransactionSplitStore) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetCurrencyId

`func (o *TransactionSplitStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TransactionSplitStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TransactionSplitStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *TransactionSplitStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *TransactionSplitStore) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *TransactionSplitStore) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCurrencyCode

`func (o *TransactionSplitStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionSplitStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionSplitStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionSplitStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *TransactionSplitStore) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *TransactionSplitStore) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetForeignAmount

`func (o *TransactionSplitStore) GetForeignAmount() string`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *TransactionSplitStore) GetForeignAmountOk() (*string, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *TransactionSplitStore) SetForeignAmount(v string)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *TransactionSplitStore) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### SetForeignAmountNil

`func (o *TransactionSplitStore) SetForeignAmountNil(b bool)`

 SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil

### UnsetForeignAmount
`func (o *TransactionSplitStore) UnsetForeignAmount()`

UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
### GetForeignCurrencyId

`func (o *TransactionSplitStore) GetForeignCurrencyId() string`

GetForeignCurrencyId returns the ForeignCurrencyId field if non-nil, zero value otherwise.

### GetForeignCurrencyIdOk

`func (o *TransactionSplitStore) GetForeignCurrencyIdOk() (*string, bool)`

GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyId

`func (o *TransactionSplitStore) SetForeignCurrencyId(v string)`

SetForeignCurrencyId sets ForeignCurrencyId field to given value.

### HasForeignCurrencyId

`func (o *TransactionSplitStore) HasForeignCurrencyId() bool`

HasForeignCurrencyId returns a boolean if a field has been set.

### SetForeignCurrencyIdNil

`func (o *TransactionSplitStore) SetForeignCurrencyIdNil(b bool)`

 SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil

### UnsetForeignCurrencyId
`func (o *TransactionSplitStore) UnsetForeignCurrencyId()`

UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
### GetForeignCurrencyCode

`func (o *TransactionSplitStore) GetForeignCurrencyCode() string`

GetForeignCurrencyCode returns the ForeignCurrencyCode field if non-nil, zero value otherwise.

### GetForeignCurrencyCodeOk

`func (o *TransactionSplitStore) GetForeignCurrencyCodeOk() (*string, bool)`

GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignCurrencyCode

`func (o *TransactionSplitStore) SetForeignCurrencyCode(v string)`

SetForeignCurrencyCode sets ForeignCurrencyCode field to given value.

### HasForeignCurrencyCode

`func (o *TransactionSplitStore) HasForeignCurrencyCode() bool`

HasForeignCurrencyCode returns a boolean if a field has been set.

### SetForeignCurrencyCodeNil

`func (o *TransactionSplitStore) SetForeignCurrencyCodeNil(b bool)`

 SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil

### UnsetForeignCurrencyCode
`func (o *TransactionSplitStore) UnsetForeignCurrencyCode()`

UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
### GetBudgetId

`func (o *TransactionSplitStore) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *TransactionSplitStore) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *TransactionSplitStore) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *TransactionSplitStore) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### SetBudgetIdNil

`func (o *TransactionSplitStore) SetBudgetIdNil(b bool)`

 SetBudgetIdNil sets the value for BudgetId to be an explicit nil

### UnsetBudgetId
`func (o *TransactionSplitStore) UnsetBudgetId()`

UnsetBudgetId ensures that no value is present for BudgetId, not even an explicit nil
### GetBudgetName

`func (o *TransactionSplitStore) GetBudgetName() string`

GetBudgetName returns the BudgetName field if non-nil, zero value otherwise.

### GetBudgetNameOk

`func (o *TransactionSplitStore) GetBudgetNameOk() (*string, bool)`

GetBudgetNameOk returns a tuple with the BudgetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetName

`func (o *TransactionSplitStore) SetBudgetName(v string)`

SetBudgetName sets BudgetName field to given value.

### HasBudgetName

`func (o *TransactionSplitStore) HasBudgetName() bool`

HasBudgetName returns a boolean if a field has been set.

### SetBudgetNameNil

`func (o *TransactionSplitStore) SetBudgetNameNil(b bool)`

 SetBudgetNameNil sets the value for BudgetName to be an explicit nil

### UnsetBudgetName
`func (o *TransactionSplitStore) UnsetBudgetName()`

UnsetBudgetName ensures that no value is present for BudgetName, not even an explicit nil
### GetCategoryId

`func (o *TransactionSplitStore) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransactionSplitStore) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransactionSplitStore) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *TransactionSplitStore) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *TransactionSplitStore) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *TransactionSplitStore) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryName

`func (o *TransactionSplitStore) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *TransactionSplitStore) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *TransactionSplitStore) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *TransactionSplitStore) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### SetCategoryNameNil

`func (o *TransactionSplitStore) SetCategoryNameNil(b bool)`

 SetCategoryNameNil sets the value for CategoryName to be an explicit nil

### UnsetCategoryName
`func (o *TransactionSplitStore) UnsetCategoryName()`

UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
### GetSourceId

`func (o *TransactionSplitStore) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TransactionSplitStore) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TransactionSplitStore) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *TransactionSplitStore) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *TransactionSplitStore) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *TransactionSplitStore) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TransactionSplitStore) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TransactionSplitStore) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TransactionSplitStore) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TransactionSplitStore) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TransactionSplitStore) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetDestinationId

`func (o *TransactionSplitStore) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *TransactionSplitStore) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *TransactionSplitStore) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### SetDestinationIdNil

`func (o *TransactionSplitStore) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *TransactionSplitStore) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetDestinationName

`func (o *TransactionSplitStore) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *TransactionSplitStore) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *TransactionSplitStore) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *TransactionSplitStore) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### SetDestinationNameNil

`func (o *TransactionSplitStore) SetDestinationNameNil(b bool)`

 SetDestinationNameNil sets the value for DestinationName to be an explicit nil

### UnsetDestinationName
`func (o *TransactionSplitStore) UnsetDestinationName()`

UnsetDestinationName ensures that no value is present for DestinationName, not even an explicit nil
### GetReconciled

`func (o *TransactionSplitStore) GetReconciled() bool`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *TransactionSplitStore) GetReconciledOk() (*bool, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *TransactionSplitStore) SetReconciled(v bool)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *TransactionSplitStore) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetPiggyBankId

`func (o *TransactionSplitStore) GetPiggyBankId() int32`

GetPiggyBankId returns the PiggyBankId field if non-nil, zero value otherwise.

### GetPiggyBankIdOk

`func (o *TransactionSplitStore) GetPiggyBankIdOk() (*int32, bool)`

GetPiggyBankIdOk returns a tuple with the PiggyBankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankId

`func (o *TransactionSplitStore) SetPiggyBankId(v int32)`

SetPiggyBankId sets PiggyBankId field to given value.

### HasPiggyBankId

`func (o *TransactionSplitStore) HasPiggyBankId() bool`

HasPiggyBankId returns a boolean if a field has been set.

### GetPiggyBankName

`func (o *TransactionSplitStore) GetPiggyBankName() string`

GetPiggyBankName returns the PiggyBankName field if non-nil, zero value otherwise.

### GetPiggyBankNameOk

`func (o *TransactionSplitStore) GetPiggyBankNameOk() (*string, bool)`

GetPiggyBankNameOk returns a tuple with the PiggyBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiggyBankName

`func (o *TransactionSplitStore) SetPiggyBankName(v string)`

SetPiggyBankName sets PiggyBankName field to given value.

### HasPiggyBankName

`func (o *TransactionSplitStore) HasPiggyBankName() bool`

HasPiggyBankName returns a boolean if a field has been set.

### GetBillId

`func (o *TransactionSplitStore) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *TransactionSplitStore) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *TransactionSplitStore) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *TransactionSplitStore) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### SetBillIdNil

`func (o *TransactionSplitStore) SetBillIdNil(b bool)`

 SetBillIdNil sets the value for BillId to be an explicit nil

### UnsetBillId
`func (o *TransactionSplitStore) UnsetBillId()`

UnsetBillId ensures that no value is present for BillId, not even an explicit nil
### GetBillName

`func (o *TransactionSplitStore) GetBillName() string`

GetBillName returns the BillName field if non-nil, zero value otherwise.

### GetBillNameOk

`func (o *TransactionSplitStore) GetBillNameOk() (*string, bool)`

GetBillNameOk returns a tuple with the BillName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillName

`func (o *TransactionSplitStore) SetBillName(v string)`

SetBillName sets BillName field to given value.

### HasBillName

`func (o *TransactionSplitStore) HasBillName() bool`

HasBillName returns a boolean if a field has been set.

### SetBillNameNil

`func (o *TransactionSplitStore) SetBillNameNil(b bool)`

 SetBillNameNil sets the value for BillName to be an explicit nil

### UnsetBillName
`func (o *TransactionSplitStore) UnsetBillName()`

UnsetBillName ensures that no value is present for BillName, not even an explicit nil
### GetTags

`func (o *TransactionSplitStore) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TransactionSplitStore) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TransactionSplitStore) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TransactionSplitStore) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TransactionSplitStore) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TransactionSplitStore) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetNotes

`func (o *TransactionSplitStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionSplitStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionSplitStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionSplitStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionSplitStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionSplitStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInternalReference

`func (o *TransactionSplitStore) GetInternalReference() string`

GetInternalReference returns the InternalReference field if non-nil, zero value otherwise.

### GetInternalReferenceOk

`func (o *TransactionSplitStore) GetInternalReferenceOk() (*string, bool)`

GetInternalReferenceOk returns a tuple with the InternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalReference

`func (o *TransactionSplitStore) SetInternalReference(v string)`

SetInternalReference sets InternalReference field to given value.

### HasInternalReference

`func (o *TransactionSplitStore) HasInternalReference() bool`

HasInternalReference returns a boolean if a field has been set.

### SetInternalReferenceNil

`func (o *TransactionSplitStore) SetInternalReferenceNil(b bool)`

 SetInternalReferenceNil sets the value for InternalReference to be an explicit nil

### UnsetInternalReference
`func (o *TransactionSplitStore) UnsetInternalReference()`

UnsetInternalReference ensures that no value is present for InternalReference, not even an explicit nil
### GetExternalId

`func (o *TransactionSplitStore) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TransactionSplitStore) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TransactionSplitStore) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TransactionSplitStore) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TransactionSplitStore) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TransactionSplitStore) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBunqPaymentId

`func (o *TransactionSplitStore) GetBunqPaymentId() string`

GetBunqPaymentId returns the BunqPaymentId field if non-nil, zero value otherwise.

### GetBunqPaymentIdOk

`func (o *TransactionSplitStore) GetBunqPaymentIdOk() (*string, bool)`

GetBunqPaymentIdOk returns a tuple with the BunqPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBunqPaymentId

`func (o *TransactionSplitStore) SetBunqPaymentId(v string)`

SetBunqPaymentId sets BunqPaymentId field to given value.

### HasBunqPaymentId

`func (o *TransactionSplitStore) HasBunqPaymentId() bool`

HasBunqPaymentId returns a boolean if a field has been set.

### SetBunqPaymentIdNil

`func (o *TransactionSplitStore) SetBunqPaymentIdNil(b bool)`

 SetBunqPaymentIdNil sets the value for BunqPaymentId to be an explicit nil

### UnsetBunqPaymentId
`func (o *TransactionSplitStore) UnsetBunqPaymentId()`

UnsetBunqPaymentId ensures that no value is present for BunqPaymentId, not even an explicit nil
### GetSepaCc

`func (o *TransactionSplitStore) GetSepaCc() string`

GetSepaCc returns the SepaCc field if non-nil, zero value otherwise.

### GetSepaCcOk

`func (o *TransactionSplitStore) GetSepaCcOk() (*string, bool)`

GetSepaCcOk returns a tuple with the SepaCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCc

`func (o *TransactionSplitStore) SetSepaCc(v string)`

SetSepaCc sets SepaCc field to given value.

### HasSepaCc

`func (o *TransactionSplitStore) HasSepaCc() bool`

HasSepaCc returns a boolean if a field has been set.

### SetSepaCcNil

`func (o *TransactionSplitStore) SetSepaCcNil(b bool)`

 SetSepaCcNil sets the value for SepaCc to be an explicit nil

### UnsetSepaCc
`func (o *TransactionSplitStore) UnsetSepaCc()`

UnsetSepaCc ensures that no value is present for SepaCc, not even an explicit nil
### GetSepaCtOp

`func (o *TransactionSplitStore) GetSepaCtOp() string`

GetSepaCtOp returns the SepaCtOp field if non-nil, zero value otherwise.

### GetSepaCtOpOk

`func (o *TransactionSplitStore) GetSepaCtOpOk() (*string, bool)`

GetSepaCtOpOk returns a tuple with the SepaCtOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtOp

`func (o *TransactionSplitStore) SetSepaCtOp(v string)`

SetSepaCtOp sets SepaCtOp field to given value.

### HasSepaCtOp

`func (o *TransactionSplitStore) HasSepaCtOp() bool`

HasSepaCtOp returns a boolean if a field has been set.

### SetSepaCtOpNil

`func (o *TransactionSplitStore) SetSepaCtOpNil(b bool)`

 SetSepaCtOpNil sets the value for SepaCtOp to be an explicit nil

### UnsetSepaCtOp
`func (o *TransactionSplitStore) UnsetSepaCtOp()`

UnsetSepaCtOp ensures that no value is present for SepaCtOp, not even an explicit nil
### GetSepaCtId

`func (o *TransactionSplitStore) GetSepaCtId() string`

GetSepaCtId returns the SepaCtId field if non-nil, zero value otherwise.

### GetSepaCtIdOk

`func (o *TransactionSplitStore) GetSepaCtIdOk() (*string, bool)`

GetSepaCtIdOk returns a tuple with the SepaCtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCtId

`func (o *TransactionSplitStore) SetSepaCtId(v string)`

SetSepaCtId sets SepaCtId field to given value.

### HasSepaCtId

`func (o *TransactionSplitStore) HasSepaCtId() bool`

HasSepaCtId returns a boolean if a field has been set.

### SetSepaCtIdNil

`func (o *TransactionSplitStore) SetSepaCtIdNil(b bool)`

 SetSepaCtIdNil sets the value for SepaCtId to be an explicit nil

### UnsetSepaCtId
`func (o *TransactionSplitStore) UnsetSepaCtId()`

UnsetSepaCtId ensures that no value is present for SepaCtId, not even an explicit nil
### GetSepaDb

`func (o *TransactionSplitStore) GetSepaDb() string`

GetSepaDb returns the SepaDb field if non-nil, zero value otherwise.

### GetSepaDbOk

`func (o *TransactionSplitStore) GetSepaDbOk() (*string, bool)`

GetSepaDbOk returns a tuple with the SepaDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaDb

`func (o *TransactionSplitStore) SetSepaDb(v string)`

SetSepaDb sets SepaDb field to given value.

### HasSepaDb

`func (o *TransactionSplitStore) HasSepaDb() bool`

HasSepaDb returns a boolean if a field has been set.

### SetSepaDbNil

`func (o *TransactionSplitStore) SetSepaDbNil(b bool)`

 SetSepaDbNil sets the value for SepaDb to be an explicit nil

### UnsetSepaDb
`func (o *TransactionSplitStore) UnsetSepaDb()`

UnsetSepaDb ensures that no value is present for SepaDb, not even an explicit nil
### GetSepaCountry

`func (o *TransactionSplitStore) GetSepaCountry() string`

GetSepaCountry returns the SepaCountry field if non-nil, zero value otherwise.

### GetSepaCountryOk

`func (o *TransactionSplitStore) GetSepaCountryOk() (*string, bool)`

GetSepaCountryOk returns a tuple with the SepaCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCountry

`func (o *TransactionSplitStore) SetSepaCountry(v string)`

SetSepaCountry sets SepaCountry field to given value.

### HasSepaCountry

`func (o *TransactionSplitStore) HasSepaCountry() bool`

HasSepaCountry returns a boolean if a field has been set.

### SetSepaCountryNil

`func (o *TransactionSplitStore) SetSepaCountryNil(b bool)`

 SetSepaCountryNil sets the value for SepaCountry to be an explicit nil

### UnsetSepaCountry
`func (o *TransactionSplitStore) UnsetSepaCountry()`

UnsetSepaCountry ensures that no value is present for SepaCountry, not even an explicit nil
### GetSepaEp

`func (o *TransactionSplitStore) GetSepaEp() string`

GetSepaEp returns the SepaEp field if non-nil, zero value otherwise.

### GetSepaEpOk

`func (o *TransactionSplitStore) GetSepaEpOk() (*string, bool)`

GetSepaEpOk returns a tuple with the SepaEp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaEp

`func (o *TransactionSplitStore) SetSepaEp(v string)`

SetSepaEp sets SepaEp field to given value.

### HasSepaEp

`func (o *TransactionSplitStore) HasSepaEp() bool`

HasSepaEp returns a boolean if a field has been set.

### SetSepaEpNil

`func (o *TransactionSplitStore) SetSepaEpNil(b bool)`

 SetSepaEpNil sets the value for SepaEp to be an explicit nil

### UnsetSepaEp
`func (o *TransactionSplitStore) UnsetSepaEp()`

UnsetSepaEp ensures that no value is present for SepaEp, not even an explicit nil
### GetSepaCi

`func (o *TransactionSplitStore) GetSepaCi() string`

GetSepaCi returns the SepaCi field if non-nil, zero value otherwise.

### GetSepaCiOk

`func (o *TransactionSplitStore) GetSepaCiOk() (*string, bool)`

GetSepaCiOk returns a tuple with the SepaCi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCi

`func (o *TransactionSplitStore) SetSepaCi(v string)`

SetSepaCi sets SepaCi field to given value.

### HasSepaCi

`func (o *TransactionSplitStore) HasSepaCi() bool`

HasSepaCi returns a boolean if a field has been set.

### SetSepaCiNil

`func (o *TransactionSplitStore) SetSepaCiNil(b bool)`

 SetSepaCiNil sets the value for SepaCi to be an explicit nil

### UnsetSepaCi
`func (o *TransactionSplitStore) UnsetSepaCi()`

UnsetSepaCi ensures that no value is present for SepaCi, not even an explicit nil
### GetSepaBatchId

`func (o *TransactionSplitStore) GetSepaBatchId() string`

GetSepaBatchId returns the SepaBatchId field if non-nil, zero value otherwise.

### GetSepaBatchIdOk

`func (o *TransactionSplitStore) GetSepaBatchIdOk() (*string, bool)`

GetSepaBatchIdOk returns a tuple with the SepaBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaBatchId

`func (o *TransactionSplitStore) SetSepaBatchId(v string)`

SetSepaBatchId sets SepaBatchId field to given value.

### HasSepaBatchId

`func (o *TransactionSplitStore) HasSepaBatchId() bool`

HasSepaBatchId returns a boolean if a field has been set.

### SetSepaBatchIdNil

`func (o *TransactionSplitStore) SetSepaBatchIdNil(b bool)`

 SetSepaBatchIdNil sets the value for SepaBatchId to be an explicit nil

### UnsetSepaBatchId
`func (o *TransactionSplitStore) UnsetSepaBatchId()`

UnsetSepaBatchId ensures that no value is present for SepaBatchId, not even an explicit nil
### GetInterestDate

`func (o *TransactionSplitStore) GetInterestDate() time.Time`

GetInterestDate returns the InterestDate field if non-nil, zero value otherwise.

### GetInterestDateOk

`func (o *TransactionSplitStore) GetInterestDateOk() (*time.Time, bool)`

GetInterestDateOk returns a tuple with the InterestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestDate

`func (o *TransactionSplitStore) SetInterestDate(v time.Time)`

SetInterestDate sets InterestDate field to given value.

### HasInterestDate

`func (o *TransactionSplitStore) HasInterestDate() bool`

HasInterestDate returns a boolean if a field has been set.

### SetInterestDateNil

`func (o *TransactionSplitStore) SetInterestDateNil(b bool)`

 SetInterestDateNil sets the value for InterestDate to be an explicit nil

### UnsetInterestDate
`func (o *TransactionSplitStore) UnsetInterestDate()`

UnsetInterestDate ensures that no value is present for InterestDate, not even an explicit nil
### GetBookDate

`func (o *TransactionSplitStore) GetBookDate() time.Time`

GetBookDate returns the BookDate field if non-nil, zero value otherwise.

### GetBookDateOk

`func (o *TransactionSplitStore) GetBookDateOk() (*time.Time, bool)`

GetBookDateOk returns a tuple with the BookDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookDate

`func (o *TransactionSplitStore) SetBookDate(v time.Time)`

SetBookDate sets BookDate field to given value.

### HasBookDate

`func (o *TransactionSplitStore) HasBookDate() bool`

HasBookDate returns a boolean if a field has been set.

### SetBookDateNil

`func (o *TransactionSplitStore) SetBookDateNil(b bool)`

 SetBookDateNil sets the value for BookDate to be an explicit nil

### UnsetBookDate
`func (o *TransactionSplitStore) UnsetBookDate()`

UnsetBookDate ensures that no value is present for BookDate, not even an explicit nil
### GetProcessDate

`func (o *TransactionSplitStore) GetProcessDate() time.Time`

GetProcessDate returns the ProcessDate field if non-nil, zero value otherwise.

### GetProcessDateOk

`func (o *TransactionSplitStore) GetProcessDateOk() (*time.Time, bool)`

GetProcessDateOk returns a tuple with the ProcessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessDate

`func (o *TransactionSplitStore) SetProcessDate(v time.Time)`

SetProcessDate sets ProcessDate field to given value.

### HasProcessDate

`func (o *TransactionSplitStore) HasProcessDate() bool`

HasProcessDate returns a boolean if a field has been set.

### SetProcessDateNil

`func (o *TransactionSplitStore) SetProcessDateNil(b bool)`

 SetProcessDateNil sets the value for ProcessDate to be an explicit nil

### UnsetProcessDate
`func (o *TransactionSplitStore) UnsetProcessDate()`

UnsetProcessDate ensures that no value is present for ProcessDate, not even an explicit nil
### GetDueDate

`func (o *TransactionSplitStore) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *TransactionSplitStore) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *TransactionSplitStore) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *TransactionSplitStore) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *TransactionSplitStore) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *TransactionSplitStore) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentDate

`func (o *TransactionSplitStore) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *TransactionSplitStore) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *TransactionSplitStore) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *TransactionSplitStore) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *TransactionSplitStore) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *TransactionSplitStore) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetInvoiceDate

`func (o *TransactionSplitStore) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *TransactionSplitStore) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *TransactionSplitStore) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *TransactionSplitStore) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### SetInvoiceDateNil

`func (o *TransactionSplitStore) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *TransactionSplitStore) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


