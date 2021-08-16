# RuleStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**RuleGroupId** | **string** | ID of the rule group under which the rule must be stored. Either this field or rule_group_title is mandatory. | 
**RuleGroupTitle** | Pointer to **string** | Title of the rule group under which the rule must be stored. Either this field or rule_group_id is mandatory. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Trigger** | **string** | Which action is necessary for the rule to fire? Use either store-journal or update-journal. | 
**Active** | Pointer to **bool** | Whether or not the rule is even active. Default is true. | [optional] 
**Strict** | Pointer to **bool** | If the rule is set to be strict, ALL triggers must hit in order for the rule to fire. Otherwise, just one is enough. Default value is true. | [optional] 
**StopProcessing** | Pointer to **bool** | If this value is true and the rule is triggered, other rules  after this one in the group will be skipped. Default value is false. | [optional] 
**Triggers** | [**[]RuleTriggerStore**](RuleTriggerStore.md) |  | 
**Actions** | [**[]RuleActionStore**](RuleActionStore.md) |  | 

## Methods

### NewRuleStore

`func NewRuleStore(title string, ruleGroupId string, trigger string, triggers []RuleTriggerStore, actions []RuleActionStore, ) *RuleStore`

NewRuleStore instantiates a new RuleStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleStoreWithDefaults

`func NewRuleStoreWithDefaults() *RuleStore`

NewRuleStoreWithDefaults instantiates a new RuleStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleStore) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RuleStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleGroupId

`func (o *RuleStore) GetRuleGroupId() string`

GetRuleGroupId returns the RuleGroupId field if non-nil, zero value otherwise.

### GetRuleGroupIdOk

`func (o *RuleStore) GetRuleGroupIdOk() (*string, bool)`

GetRuleGroupIdOk returns a tuple with the RuleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupId

`func (o *RuleStore) SetRuleGroupId(v string)`

SetRuleGroupId sets RuleGroupId field to given value.


### GetRuleGroupTitle

`func (o *RuleStore) GetRuleGroupTitle() string`

GetRuleGroupTitle returns the RuleGroupTitle field if non-nil, zero value otherwise.

### GetRuleGroupTitleOk

`func (o *RuleStore) GetRuleGroupTitleOk() (*string, bool)`

GetRuleGroupTitleOk returns a tuple with the RuleGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupTitle

`func (o *RuleStore) SetRuleGroupTitle(v string)`

SetRuleGroupTitle sets RuleGroupTitle field to given value.

### HasRuleGroupTitle

`func (o *RuleStore) HasRuleGroupTitle() bool`

HasRuleGroupTitle returns a boolean if a field has been set.

### GetOrder

`func (o *RuleStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTrigger

`func (o *RuleStore) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RuleStore) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RuleStore) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetActive

`func (o *RuleStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrict

`func (o *RuleStore) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *RuleStore) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *RuleStore) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *RuleStore) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetStopProcessing

`func (o *RuleStore) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *RuleStore) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *RuleStore) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *RuleStore) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.

### GetTriggers

`func (o *RuleStore) GetTriggers() []RuleTriggerStore`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *RuleStore) GetTriggersOk() (*[]RuleTriggerStore, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *RuleStore) SetTriggers(v []RuleTriggerStore)`

SetTriggers sets Triggers field to given value.


### GetActions

`func (o *RuleStore) GetActions() []RuleActionStore`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RuleStore) GetActionsOk() (*[]RuleActionStore, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RuleStore) SetActions(v []RuleActionStore)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


