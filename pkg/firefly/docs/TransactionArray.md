# TransactionArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TransactionRead**](TransactionRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewTransactionArray

`func NewTransactionArray(data []TransactionRead, meta Meta, links PageLink, ) *TransactionArray`

NewTransactionArray instantiates a new TransactionArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionArrayWithDefaults

`func NewTransactionArrayWithDefaults() *TransactionArray`

NewTransactionArrayWithDefaults instantiates a new TransactionArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionArray) GetData() []TransactionRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionArray) GetDataOk() (*[]TransactionRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionArray) SetData(v []TransactionRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *TransactionArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TransactionArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TransactionArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *TransactionArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


