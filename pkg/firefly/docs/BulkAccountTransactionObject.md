# BulkAccountTransactionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalAccount** | **float32** | The original account, where you wish to transactions away from. | 
**DestinationAccount** | **float32** | The destination account, where the transactions will be moved to. | 

## Methods

### NewBulkAccountTransactionObject

`func NewBulkAccountTransactionObject(originalAccount float32, destinationAccount float32, ) *BulkAccountTransactionObject`

NewBulkAccountTransactionObject instantiates a new BulkAccountTransactionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAccountTransactionObjectWithDefaults

`func NewBulkAccountTransactionObjectWithDefaults() *BulkAccountTransactionObject`

NewBulkAccountTransactionObjectWithDefaults instantiates a new BulkAccountTransactionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalAccount

`func (o *BulkAccountTransactionObject) GetOriginalAccount() float32`

GetOriginalAccount returns the OriginalAccount field if non-nil, zero value otherwise.

### GetOriginalAccountOk

`func (o *BulkAccountTransactionObject) GetOriginalAccountOk() (*float32, bool)`

GetOriginalAccountOk returns a tuple with the OriginalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAccount

`func (o *BulkAccountTransactionObject) SetOriginalAccount(v float32)`

SetOriginalAccount sets OriginalAccount field to given value.


### GetDestinationAccount

`func (o *BulkAccountTransactionObject) GetDestinationAccount() float32`

GetDestinationAccount returns the DestinationAccount field if non-nil, zero value otherwise.

### GetDestinationAccountOk

`func (o *BulkAccountTransactionObject) GetDestinationAccountOk() (*float32, bool)`

GetDestinationAccountOk returns a tuple with the DestinationAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccount

`func (o *BulkAccountTransactionObject) SetDestinationAccount(v float32)`

SetDestinationAccount sets DestinationAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


