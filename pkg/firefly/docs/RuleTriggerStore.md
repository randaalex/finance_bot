# RuleTriggerStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of thing this trigger responds to. A limited set is possible | 
**Value** | **string** | The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger. | 
**Order** | Pointer to **int32** | Order of the trigger | [optional] 
**Active** | Pointer to **bool** | If the trigger is active. Defaults to true. | [optional] 
**StopProcessing** | Pointer to **bool** | When true, other triggers will not be checked if this trigger was triggered. Defaults to false. | [optional] 

## Methods

### NewRuleTriggerStore

`func NewRuleTriggerStore(type_ string, value string, ) *RuleTriggerStore`

NewRuleTriggerStore instantiates a new RuleTriggerStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTriggerStoreWithDefaults

`func NewRuleTriggerStoreWithDefaults() *RuleTriggerStore`

NewRuleTriggerStoreWithDefaults instantiates a new RuleTriggerStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTriggerStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTriggerStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTriggerStore) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *RuleTriggerStore) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleTriggerStore) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleTriggerStore) SetValue(v string)`

SetValue sets Value field to given value.


### GetOrder

`func (o *RuleTriggerStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleTriggerStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleTriggerStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleTriggerStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleTriggerStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleTriggerStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleTriggerStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleTriggerStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleTriggerStore) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleTriggerStore) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleTriggerStore) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleTriggerStore) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


