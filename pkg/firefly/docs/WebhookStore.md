# WebhookStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Boolean to indicate if the webhook is active | [optional] 
**Title** | **string** | A title for the webhook for easy recognition. | 
**Trigger** | **string** | The trigger for the webhook. | 
**Response** | **string** | Indicator for what Firefly III will deliver to the webhook URL. | 
**Delivery** | **string** | Format of the delivered response. | 
**Url** | **string** | The URL of the webhook. Has to start with &#x60;https&#x60;. | 

## Methods

### NewWebhookStore

`func NewWebhookStore(title string, trigger string, response string, delivery string, url string, ) *WebhookStore`

NewWebhookStore instantiates a new WebhookStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookStoreWithDefaults

`func NewWebhookStoreWithDefaults() *WebhookStore`

NewWebhookStoreWithDefaults instantiates a new WebhookStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTitle

`func (o *WebhookStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookStore) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTrigger

`func (o *WebhookStore) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *WebhookStore) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *WebhookStore) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetResponse

`func (o *WebhookStore) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookStore) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookStore) SetResponse(v string)`

SetResponse sets Response field to given value.


### GetDelivery

`func (o *WebhookStore) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *WebhookStore) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *WebhookStore) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.


### GetUrl

`func (o *WebhookStore) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookStore) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookStore) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


