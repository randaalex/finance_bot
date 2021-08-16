# RuleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RuleGroupId** | Pointer to **string** | ID of the rule group under which the rule must be stored. Either this field or rule_group_title is mandatory. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Trigger** | Pointer to **string** | Which action is necessary for the rule to fire? Use either store-journal or update-journal. | [optional] 
**Active** | Pointer to **bool** | Whether or not the rule is even active. Default is true. | [optional] 
**Strict** | Pointer to **bool** | If the rule is set to be strict, ALL triggers must hit in order for the rule to fire. Otherwise, just one is enough. Default value is true. | [optional] 
**StopProcessing** | Pointer to **bool** | If this value is true and the rule is triggered, other rules  after this one in the group will be skipped. Default value is false. | [optional] 
**Triggers** | Pointer to [**[]RuleTriggerUpdate**](RuleTriggerUpdate.md) |  | [optional] 
**Actions** | Pointer to [**[]RuleActionUpdate**](RuleActionUpdate.md) |  | [optional] 

## Methods

### NewRuleUpdate

`func NewRuleUpdate() *RuleUpdate`

NewRuleUpdate instantiates a new RuleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleUpdateWithDefaults

`func NewRuleUpdateWithDefaults() *RuleUpdate`

NewRuleUpdateWithDefaults instantiates a new RuleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RuleUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RuleUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleGroupId

`func (o *RuleUpdate) GetRuleGroupId() string`

GetRuleGroupId returns the RuleGroupId field if non-nil, zero value otherwise.

### GetRuleGroupIdOk

`func (o *RuleUpdate) GetRuleGroupIdOk() (*string, bool)`

GetRuleGroupIdOk returns a tuple with the RuleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupId

`func (o *RuleUpdate) SetRuleGroupId(v string)`

SetRuleGroupId sets RuleGroupId field to given value.

### HasRuleGroupId

`func (o *RuleUpdate) HasRuleGroupId() bool`

HasRuleGroupId returns a boolean if a field has been set.

### GetOrder

`func (o *RuleUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTrigger

`func (o *RuleUpdate) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RuleUpdate) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RuleUpdate) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *RuleUpdate) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetActive

`func (o *RuleUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrict

`func (o *RuleUpdate) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *RuleUpdate) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *RuleUpdate) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *RuleUpdate) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleUpdate) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleUpdate) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleUpdate) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleUpdate) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.

### GetTriggers

`func (o *RuleUpdate) GetTriggers() []RuleTriggerUpdate`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *RuleUpdate) GetTriggersOk() (*[]RuleTriggerUpdate, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *RuleUpdate) SetTriggers(v []RuleTriggerUpdate)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *RuleUpdate) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetActions

`func (o *RuleUpdate) GetActions() []RuleActionUpdate`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RuleUpdate) GetActionsOk() (*[]RuleActionUpdate, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RuleUpdate) SetActions(v []RuleActionUpdate)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RuleUpdate) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


