# AccountArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccountRead**](AccountRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewAccountArray

`func NewAccountArray(data []AccountRead, meta Meta, ) *AccountArray`

NewAccountArray instantiates a new AccountArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountArrayWithDefaults

`func NewAccountArrayWithDefaults() *AccountArray`

NewAccountArrayWithDefaults instantiates a new AccountArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccountArray) GetData() []AccountRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountArray) GetDataOk() (*[]AccountRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountArray) SetData(v []AccountRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *AccountArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AccountArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AccountArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


