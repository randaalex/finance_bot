# WebhookMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Sent** | Pointer to **bool** | If this message is sent yet. | [optional] 
**Errored** | Pointer to **bool** | If this message has errored out. | [optional] 
**WebhookId** | Pointer to **string** | The ID of the webhook this message belongs to. | [optional] 
**Uuid** | Pointer to **string** | Long UUID string for identification of this webhook message. | [optional] 
**String** | Pointer to **NullableString** | The actual message that is sent or will be sent as JSON string. | [optional] 

## Methods

### NewWebhookMessage

`func NewWebhookMessage() *WebhookMessage`

NewWebhookMessage instantiates a new WebhookMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookMessageWithDefaults

`func NewWebhookMessageWithDefaults() *WebhookMessage`

NewWebhookMessageWithDefaults instantiates a new WebhookMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WebhookMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookMessage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSent

`func (o *WebhookMessage) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *WebhookMessage) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *WebhookMessage) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *WebhookMessage) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetErrored

`func (o *WebhookMessage) GetErrored() bool`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *WebhookMessage) GetErroredOk() (*bool, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *WebhookMessage) SetErrored(v bool)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *WebhookMessage) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookMessage) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookMessage) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookMessage) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookMessage) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetUuid

`func (o *WebhookMessage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *WebhookMessage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *WebhookMessage) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *WebhookMessage) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetString

`func (o *WebhookMessage) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *WebhookMessage) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *WebhookMessage) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *WebhookMessage) HasString() bool`

HasString returns a boolean if a field has been set.

### SetStringNil

`func (o *WebhookMessage) SetStringNil(b bool)`

 SetStringNil sets the value for String to be an explicit nil

### UnsetString
`func (o *WebhookMessage) UnsetString()`

UnsetString ensures that no value is present for String, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


