# WebhookMessageRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**WebhookMessage**](WebhookMessage.md) |  | 

## Methods

### NewWebhookMessageRead

`func NewWebhookMessageRead(type_ string, id string, attributes WebhookMessage, ) *WebhookMessageRead`

NewWebhookMessageRead instantiates a new WebhookMessageRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMessageReadWithDefaults

`func NewWebhookMessageReadWithDefaults() *WebhookMessageRead`

NewWebhookMessageReadWithDefaults instantiates a new WebhookMessageRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookMessageRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookMessageRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookMessageRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *WebhookMessageRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookMessageRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookMessageRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *WebhookMessageRead) GetAttributes() WebhookMessage`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WebhookMessageRead) GetAttributesOk() (*WebhookMessage, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WebhookMessageRead) SetAttributes(v WebhookMessage)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


