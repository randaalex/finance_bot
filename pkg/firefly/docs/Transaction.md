# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**User** | Pointer to **int32** | User ID | [optional] [readonly] 
**ErrorIfDuplicateHash** | Pointer to **bool** | Break if the submitted transaction exists already. | [optional] 
**ApplyRules** | Pointer to **bool** | Whether or not to apply rules when submitting transaction. | [optional] 
**GroupTitle** | Pointer to **string** | Title of the transaction if it has been split in more than one piece. Empty otherwise. | [optional] 
**Transactions** | [**[]TransactionSplit**](TransactionSplit.md) |  | 

## Methods

### NewTransaction

`func NewTransaction(transactions []TransactionSplit, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Transaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Transaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *Transaction) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transaction) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transaction) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *Transaction) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetErrorIfDuplicateHash

`func (o *Transaction) GetErrorIfDuplicateHash() bool`

GetErrorIfDuplicateHash returns the ErrorIfDuplicateHash field if non-nil, zero value otherwise.

### GetErrorIfDuplicateHashOk

`func (o *Transaction) GetErrorIfDuplicateHashOk() (*bool, bool)`

GetErrorIfDuplicateHashOk returns a tuple with the ErrorIfDuplicateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIfDuplicateHash

`func (o *Transaction) SetErrorIfDuplicateHash(v bool)`

SetErrorIfDuplicateHash sets ErrorIfDuplicateHash field to given value.

### HasErrorIfDuplicateHash

`func (o *Transaction) HasErrorIfDuplicateHash() bool`

HasErrorIfDuplicateHash returns a boolean if a field has been set.

### GetApplyRules

`func (o *Transaction) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *Transaction) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *Transaction) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *Transaction) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetGroupTitle

`func (o *Transaction) GetGroupTitle() string`

GetGroupTitle returns the GroupTitle field if non-nil, zero value otherwise.

### GetGroupTitleOk

`func (o *Transaction) GetGroupTitleOk() (*string, bool)`

GetGroupTitleOk returns a tuple with the GroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTitle

`func (o *Transaction) SetGroupTitle(v string)`

SetGroupTitle sets GroupTitle field to given value.

### HasGroupTitle

`func (o *Transaction) HasGroupTitle() bool`

HasGroupTitle returns a boolean if a field has been set.

### GetTransactions

`func (o *Transaction) GetTransactions() []TransactionSplit`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Transaction) GetTransactionsOk() (*[]TransactionSplit, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Transaction) SetTransactions(v []TransactionSplit)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


