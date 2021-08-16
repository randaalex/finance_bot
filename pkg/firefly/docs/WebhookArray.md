# WebhookArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookRead**](WebhookRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewWebhookArray

`func NewWebhookArray(data []WebhookRead, meta Meta, links PageLink, ) *WebhookArray`

NewWebhookArray instantiates a new WebhookArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookArrayWithDefaults

`func NewWebhookArrayWithDefaults() *WebhookArray`

NewWebhookArrayWithDefaults instantiates a new WebhookArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookArray) GetData() []WebhookRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookArray) GetDataOk() (*[]WebhookRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookArray) SetData(v []WebhookRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *WebhookArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebhookArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebhookArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *WebhookArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebhookArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebhookArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


