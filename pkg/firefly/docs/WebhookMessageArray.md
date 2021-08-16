# WebhookMessageArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookMessageRead**](WebhookMessageRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewWebhookMessageArray

`func NewWebhookMessageArray(data []WebhookMessageRead, meta Meta, ) *WebhookMessageArray`

NewWebhookMessageArray instantiates a new WebhookMessageArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMessageArrayWithDefaults

`func NewWebhookMessageArrayWithDefaults() *WebhookMessageArray`

NewWebhookMessageArrayWithDefaults instantiates a new WebhookMessageArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookMessageArray) GetData() []WebhookMessageRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookMessageArray) GetDataOk() (*[]WebhookMessageRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookMessageArray) SetData(v []WebhookMessageRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *WebhookMessageArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebhookMessageArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebhookMessageArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


