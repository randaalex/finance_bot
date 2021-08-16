# RuleActionStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of thing this action will do. A limited set is possible. | 
**Value** | **NullableString** | The accompanying value the action will set, change or update. Can be empty, but for some types this value is mandatory. | 
**Order** | Pointer to **int32** | Order of the action | [optional] 
**Active** | Pointer to **bool** | If the action is active. Defaults to true. | [optional] 
**StopProcessing** | Pointer to **bool** | When true, other actions will not be fired after this action has fired. Defaults to false. | [optional] 

## Methods

### NewRuleActionStore

`func NewRuleActionStore(type_ string, value NullableString, ) *RuleActionStore`

NewRuleActionStore instantiates a new RuleActionStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleActionStoreWithDefaults

`func NewRuleActionStoreWithDefaults() *RuleActionStore`

NewRuleActionStoreWithDefaults instantiates a new RuleActionStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleActionStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleActionStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleActionStore) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *RuleActionStore) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleActionStore) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleActionStore) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *RuleActionStore) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RuleActionStore) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetOrder

`func (o *RuleActionStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleActionStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleActionStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleActionStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleActionStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleActionStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleActionStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleActionStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleActionStore) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleActionStore) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleActionStore) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleActionStore) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


