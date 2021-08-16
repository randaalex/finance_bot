# ObjectGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int32** | Order of the object group | [optional] 

## Methods

### NewObjectGroupUpdate

`func NewObjectGroupUpdate() *ObjectGroupUpdate`

NewObjectGroupUpdate instantiates a new ObjectGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupUpdateWithDefaults

`func NewObjectGroupUpdateWithDefaults() *ObjectGroupUpdate`

NewObjectGroupUpdateWithDefaults instantiates a new ObjectGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ObjectGroupUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ObjectGroupUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ObjectGroupUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ObjectGroupUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrder

`func (o *ObjectGroupUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ObjectGroupUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ObjectGroupUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ObjectGroupUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


