# TransactionStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorIfDuplicateHash** | Pointer to **bool** | Break if the submitted transaction exists already. | [optional] 
**ApplyRules** | Pointer to **bool** | Whether or not to apply rules when submitting transaction. | [optional] 
**GroupTitle** | Pointer to **NullableString** | Title of the transaction if it has been split in more than one piece. Empty otherwise. | [optional] 
**Transactions** | [**[]TransactionSplitStore**](TransactionSplitStore.md) |  | 

## Methods

### NewTransactionStore

`func NewTransactionStore(transactions []TransactionSplitStore, ) *TransactionStore`

NewTransactionStore instantiates a new TransactionStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStoreWithDefaults

`func NewTransactionStoreWithDefaults() *TransactionStore`

NewTransactionStoreWithDefaults instantiates a new TransactionStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorIfDuplicateHash

`func (o *TransactionStore) GetErrorIfDuplicateHash() bool`

GetErrorIfDuplicateHash returns the ErrorIfDuplicateHash field if non-nil, zero value otherwise.

### GetErrorIfDuplicateHashOk

`func (o *TransactionStore) GetErrorIfDuplicateHashOk() (*bool, bool)`

GetErrorIfDuplicateHashOk returns a tuple with the ErrorIfDuplicateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIfDuplicateHash

`func (o *TransactionStore) SetErrorIfDuplicateHash(v bool)`

SetErrorIfDuplicateHash sets ErrorIfDuplicateHash field to given value.

### HasErrorIfDuplicateHash

`func (o *TransactionStore) HasErrorIfDuplicateHash() bool`

HasErrorIfDuplicateHash returns a boolean if a field has been set.

### GetApplyRules

`func (o *TransactionStore) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *TransactionStore) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *TransactionStore) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *TransactionStore) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetGroupTitle

`func (o *TransactionStore) GetGroupTitle() string`

GetGroupTitle returns the GroupTitle field if non-nil, zero value otherwise.

### GetGroupTitleOk

`func (o *TransactionStore) GetGroupTitleOk() (*string, bool)`

GetGroupTitleOk returns a tuple with the GroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTitle

`func (o *TransactionStore) SetGroupTitle(v string)`

SetGroupTitle sets GroupTitle field to given value.

### HasGroupTitle

`func (o *TransactionStore) HasGroupTitle() bool`

HasGroupTitle returns a boolean if a field has been set.

### SetGroupTitleNil

`func (o *TransactionStore) SetGroupTitleNil(b bool)`

 SetGroupTitleNil sets the value for GroupTitle to be an explicit nil

### UnsetGroupTitle
`func (o *TransactionStore) UnsetGroupTitle()`

UnsetGroupTitle ensures that no value is present for GroupTitle, not even an explicit nil
### GetTransactions

`func (o *TransactionStore) GetTransactions() []TransactionSplitStore`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionStore) GetTransactionsOk() (*[]TransactionSplitStore, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionStore) SetTransactions(v []TransactionSplitStore)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


