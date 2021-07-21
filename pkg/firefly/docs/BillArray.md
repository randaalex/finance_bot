# BillArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillRead**](BillRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewBillArray

`func NewBillArray(data []BillRead, meta Meta, ) *BillArray`

NewBillArray instantiates a new BillArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillArrayWithDefaults

`func NewBillArrayWithDefaults() *BillArray`

NewBillArrayWithDefaults instantiates a new BillArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BillArray) GetData() []BillRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillArray) GetDataOk() (*[]BillRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillArray) SetData(v []BillRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *BillArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BillArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BillArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


