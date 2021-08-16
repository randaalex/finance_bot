# TransactionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyRules** | Pointer to **bool** | Whether or not to apply rules when submitting transaction. | [optional] 
**GroupTitle** | Pointer to **NullableString** | Title of the transaction if it has been split in more than one piece. Empty otherwise. | [optional] 
**Transactions** | Pointer to [**[]TransactionSplitUpdate**](TransactionSplitUpdate.md) |  | [optional] 

## Methods

### NewTransactionUpdate

`func NewTransactionUpdate() *TransactionUpdate`

NewTransactionUpdate instantiates a new TransactionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionUpdateWithDefaults

`func NewTransactionUpdateWithDefaults() *TransactionUpdate`

NewTransactionUpdateWithDefaults instantiates a new TransactionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyRules

`func (o *TransactionUpdate) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *TransactionUpdate) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *TransactionUpdate) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *TransactionUpdate) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetGroupTitle

`func (o *TransactionUpdate) GetGroupTitle() string`

GetGroupTitle returns the GroupTitle field if non-nil, zero value otherwise.

### GetGroupTitleOk

`func (o *TransactionUpdate) GetGroupTitleOk() (*string, bool)`

GetGroupTitleOk returns a tuple with the GroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTitle

`func (o *TransactionUpdate) SetGroupTitle(v string)`

SetGroupTitle sets GroupTitle field to given value.

### HasGroupTitle

`func (o *TransactionUpdate) HasGroupTitle() bool`

HasGroupTitle returns a boolean if a field has been set.

### SetGroupTitleNil

`func (o *TransactionUpdate) SetGroupTitleNil(b bool)`

 SetGroupTitleNil sets the value for GroupTitle to be an explicit nil

### UnsetGroupTitle
`func (o *TransactionUpdate) UnsetGroupTitle()`

UnsetGroupTitle ensures that no value is present for GroupTitle, not even an explicit nil
### GetTransactions

`func (o *TransactionUpdate) GetTransactions() []TransactionSplitUpdate`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionUpdate) GetTransactionsOk() (*[]TransactionSplitUpdate, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionUpdate) SetTransactions(v []TransactionSplitUpdate)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *TransactionUpdate) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


