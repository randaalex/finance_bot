# RuleGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewRuleGroupUpdate

`func NewRuleGroupUpdate() *RuleGroupUpdate`

NewRuleGroupUpdate instantiates a new RuleGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleGroupUpdateWithDefaults

`func NewRuleGroupUpdateWithDefaults() *RuleGroupUpdate`

NewRuleGroupUpdateWithDefaults instantiates a new RuleGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleGroupUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleGroupUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleGroupUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RuleGroupUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RuleGroupUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleGroupUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleGroupUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleGroupUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RuleGroupUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RuleGroupUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *RuleGroupUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleGroupUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleGroupUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleGroupUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleGroupUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleGroupUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleGroupUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleGroupUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


