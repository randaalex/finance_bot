# ObjectGroupArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ObjectGroupRead**](ObjectGroupRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewObjectGroupArray

`func NewObjectGroupArray(data []ObjectGroupRead, meta Meta, ) *ObjectGroupArray`

NewObjectGroupArray instantiates a new ObjectGroupArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupArrayWithDefaults

`func NewObjectGroupArrayWithDefaults() *ObjectGroupArray`

NewObjectGroupArrayWithDefaults instantiates a new ObjectGroupArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ObjectGroupArray) GetData() []ObjectGroupRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectGroupArray) GetDataOk() (*[]ObjectGroupRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectGroupArray) SetData(v []ObjectGroupRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *ObjectGroupArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ObjectGroupArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ObjectGroupArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


