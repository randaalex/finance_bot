# AttachmentArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AttachmentRead**](AttachmentRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewAttachmentArray

`func NewAttachmentArray(data []AttachmentRead, meta Meta, ) *AttachmentArray`

NewAttachmentArray instantiates a new AttachmentArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentArrayWithDefaults

`func NewAttachmentArrayWithDefaults() *AttachmentArray`

NewAttachmentArrayWithDefaults instantiates a new AttachmentArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AttachmentArray) GetData() []AttachmentRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AttachmentArray) GetDataOk() (*[]AttachmentRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AttachmentArray) SetData(v []AttachmentRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *AttachmentArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AttachmentArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AttachmentArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


