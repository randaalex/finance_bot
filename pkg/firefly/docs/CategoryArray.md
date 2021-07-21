# CategoryArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CategoryRead**](CategoryRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewCategoryArray

`func NewCategoryArray(data []CategoryRead, meta Meta, ) *CategoryArray`

NewCategoryArray instantiates a new CategoryArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryArrayWithDefaults

`func NewCategoryArrayWithDefaults() *CategoryArray`

NewCategoryArrayWithDefaults instantiates a new CategoryArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CategoryArray) GetData() []CategoryRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CategoryArray) GetDataOk() (*[]CategoryRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CategoryArray) SetData(v []CategoryRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *CategoryArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CategoryArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CategoryArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


