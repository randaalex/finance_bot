# WebhookAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**WebhookMessageId** | Pointer to **string** | The ID of the webhook message this attempt belongs to. | [optional] 
**StatusCode** | Pointer to **NullableInt32** | The HTTP status code of the error, if any. | [optional] 
**Logs** | Pointer to **NullableString** | Internal log for this attempt. May contain sensitive user data. | [optional] 
**Response** | Pointer to **NullableString** | Webhook receiver response for this attempt, if any. May contain sensitive user data. | [optional] 

## Methods

### NewWebhookAttempt

`func NewWebhookAttempt() *WebhookAttempt`

NewWebhookAttempt instantiates a new WebhookAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAttemptWithDefaults

`func NewWebhookAttemptWithDefaults() *WebhookAttempt`

NewWebhookAttemptWithDefaults instantiates a new WebhookAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WebhookAttempt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookAttempt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookAttempt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookAttempt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookAttempt) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookAttempt) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookAttempt) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookAttempt) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookMessageId

`func (o *WebhookAttempt) GetWebhookMessageId() string`

GetWebhookMessageId returns the WebhookMessageId field if non-nil, zero value otherwise.

### GetWebhookMessageIdOk

`func (o *WebhookAttempt) GetWebhookMessageIdOk() (*string, bool)`

GetWebhookMessageIdOk returns a tuple with the WebhookMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMessageId

`func (o *WebhookAttempt) SetWebhookMessageId(v string)`

SetWebhookMessageId sets WebhookMessageId field to given value.

### HasWebhookMessageId

`func (o *WebhookAttempt) HasWebhookMessageId() bool`

HasWebhookMessageId returns a boolean if a field has been set.

### GetStatusCode

`func (o *WebhookAttempt) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookAttempt) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookAttempt) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WebhookAttempt) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *WebhookAttempt) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *WebhookAttempt) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetLogs

`func (o *WebhookAttempt) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *WebhookAttempt) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *WebhookAttempt) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *WebhookAttempt) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *WebhookAttempt) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *WebhookAttempt) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetResponse

`func (o *WebhookAttempt) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookAttempt) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookAttempt) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebhookAttempt) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *WebhookAttempt) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *WebhookAttempt) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


