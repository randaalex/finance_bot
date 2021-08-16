# RuleTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | **string** | The type of thing this trigger responds to. A limited set is possible | 
**Value** | **string** | The accompanying value the trigger responds to. This value is often mandatory, but this depends on the trigger. | 
**Order** | Pointer to **int32** | Order of the trigger | [optional] [readonly] 
**Active** | Pointer to **bool** | If the trigger is active. Defaults to true. | [optional] 
**StopProcessing** | Pointer to **bool** | When true, other triggers will not be checked if this trigger was triggered. Defaults to false. | [optional] 

## Methods

### NewRuleTrigger

`func NewRuleTrigger(type_ string, value string, ) *RuleTrigger`

NewRuleTrigger instantiates a new RuleTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTriggerWithDefaults

`func NewRuleTriggerWithDefaults() *RuleTrigger`

NewRuleTriggerWithDefaults instantiates a new RuleTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RuleTrigger) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleTrigger) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleTrigger) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RuleTrigger) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RuleTrigger) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleTrigger) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleTrigger) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RuleTrigger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *RuleTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleTrigger) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *RuleTrigger) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleTrigger) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleTrigger) SetValue(v string)`

SetValue sets Value field to given value.


### GetOrder

`func (o *RuleTrigger) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleTrigger) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleTrigger) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleTrigger) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleTrigger) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleTrigger) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleTrigger) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleTrigger) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleTrigger) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleTrigger) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleTrigger) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleTrigger) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


