# TransactionLinkArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TransactionLinkRead**](TransactionLinkRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewTransactionLinkArray

`func NewTransactionLinkArray(data []TransactionLinkRead, meta Meta, links PageLink, ) *TransactionLinkArray`

NewTransactionLinkArray instantiates a new TransactionLinkArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLinkArrayWithDefaults

`func NewTransactionLinkArrayWithDefaults() *TransactionLinkArray`

NewTransactionLinkArrayWithDefaults instantiates a new TransactionLinkArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransactionLinkArray) GetData() []TransactionLinkRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionLinkArray) GetDataOk() (*[]TransactionLinkRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionLinkArray) SetData(v []TransactionLinkRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *TransactionLinkArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TransactionLinkArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TransactionLinkArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *TransactionLinkArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionLinkArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionLinkArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


