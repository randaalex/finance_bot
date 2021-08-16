# WebhookAttemptArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookAttemptRead**](WebhookAttemptRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewWebhookAttemptArray

`func NewWebhookAttemptArray(data []WebhookAttemptRead, meta Meta, ) *WebhookAttemptArray`

NewWebhookAttemptArray instantiates a new WebhookAttemptArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAttemptArrayWithDefaults

`func NewWebhookAttemptArrayWithDefaults() *WebhookAttemptArray`

NewWebhookAttemptArrayWithDefaults instantiates a new WebhookAttemptArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookAttemptArray) GetData() []WebhookAttemptRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAttemptArray) GetDataOk() (*[]WebhookAttemptRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAttemptArray) SetData(v []WebhookAttemptRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *WebhookAttemptArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebhookAttemptArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebhookAttemptArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


