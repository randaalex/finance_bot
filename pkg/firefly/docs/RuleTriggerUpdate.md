# RuleTriggerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of thing this trigger responds to. A limited set is possible | [optional] 
**Value** | Pointer to **string** | The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger. | [optional] 
**Order** | Pointer to **int32** | Order of the trigger | [optional] 
**Active** | Pointer to **bool** | If the trigger is active. | [optional] 
**StopProcessing** | Pointer to **bool** | When true, other triggers will not be checked if this trigger was triggered. | [optional] 

## Methods

### NewRuleTriggerUpdate

`func NewRuleTriggerUpdate() *RuleTriggerUpdate`

NewRuleTriggerUpdate instantiates a new RuleTriggerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTriggerUpdateWithDefaults

`func NewRuleTriggerUpdateWithDefaults() *RuleTriggerUpdate`

NewRuleTriggerUpdateWithDefaults instantiates a new RuleTriggerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleTriggerUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTriggerUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTriggerUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RuleTriggerUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *RuleTriggerUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleTriggerUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleTriggerUpdate) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RuleTriggerUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOrder

`func (o *RuleTriggerUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleTriggerUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleTriggerUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleTriggerUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleTriggerUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleTriggerUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleTriggerUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleTriggerUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleTriggerUpdate) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleTriggerUpdate) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleTriggerUpdate) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleTriggerUpdate) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


