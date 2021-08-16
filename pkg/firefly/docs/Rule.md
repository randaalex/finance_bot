# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**RuleGroupId** | **string** | ID of the rule group under which the rule must be stored. Either this field or rule_group_title is mandatory. | 
**RuleGroupTitle** | Pointer to **string** | Title of the rule group under which the rule must be stored. Either this field or rule_group_id is mandatory. | [optional] 
**Order** | Pointer to **int32** |  | [optional] [readonly] 
**Trigger** | **string** | Which action is necessary for the rule to fire? Use either store-journal or update-journal. | 
**Active** | Pointer to **bool** | Whether or not the rule is even active. Default is true. | [optional] 
**Strict** | Pointer to **bool** | If the rule is set to be strict, ALL triggers must hit in order for the rule to fire. Otherwise, just one is enough. Default value is true. | [optional] 
**StopProcessing** | Pointer to **bool** | If this value is true and the rule is triggered, other rules  after this one in the group will be skipped. Default value is false. | [optional] 
**Triggers** | [**[]RuleTrigger**](RuleTrigger.md) |  | 
**Actions** | [**[]RuleAction**](RuleAction.md) |  | 

## Methods

### NewRule

`func NewRule(title string, ruleGroupId string, trigger string, triggers []RuleTrigger, actions []RuleAction, ) *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Rule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Rule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Rule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Rule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Rule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Rule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Rule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Rule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *Rule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Rule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Rule) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Rule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Rule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Rule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRuleGroupId

`func (o *Rule) GetRuleGroupId() string`

GetRuleGroupId returns the RuleGroupId field if non-nil, zero value otherwise.

### GetRuleGroupIdOk

`func (o *Rule) GetRuleGroupIdOk() (*string, bool)`

GetRuleGroupIdOk returns a tuple with the RuleGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupId

`func (o *Rule) SetRuleGroupId(v string)`

SetRuleGroupId sets RuleGroupId field to given value.


### GetRuleGroupTitle

`func (o *Rule) GetRuleGroupTitle() string`

GetRuleGroupTitle returns the RuleGroupTitle field if non-nil, zero value otherwise.

### GetRuleGroupTitleOk

`func (o *Rule) GetRuleGroupTitleOk() (*string, bool)`

GetRuleGroupTitleOk returns a tuple with the RuleGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupTitle

`func (o *Rule) SetRuleGroupTitle(v string)`

SetRuleGroupTitle sets RuleGroupTitle field to given value.

### HasRuleGroupTitle

`func (o *Rule) HasRuleGroupTitle() bool`

HasRuleGroupTitle returns a boolean if a field has been set.

### GetOrder

`func (o *Rule) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Rule) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Rule) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Rule) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetTrigger

`func (o *Rule) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Rule) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Rule) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.


### GetActive

`func (o *Rule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Rule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Rule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Rule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStrict

`func (o *Rule) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *Rule) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *Rule) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *Rule) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetStopProcessing

`func (o *Rule) GetStopProcessing() bool`

GetStopProcessing returns the StopProcessing field if non-nil, zero value otherwise.

### GetStopProcessingOk

`func (o *Rule) GetStopProcessingOk() (*bool, bool)`

GetStopProcessingOk returns a tuple with the StopProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProcessing

`func (o *Rule) SetStopProcessing(v bool)`

SetStopProcessing sets StopProcessing field to given value.

### HasStopProcessing

`func (o *Rule) HasStopProcessing() bool`

HasStopProcessing returns a boolean if a field has been set.

### GetTriggers

`func (o *Rule) GetTriggers() []RuleTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *Rule) GetTriggersOk() (*[]RuleTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *Rule) SetTriggers(v []RuleTrigger)`

SetTriggers sets Triggers field to given value.


### GetActions

`func (o *Rule) GetActions() []RuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Rule) GetActionsOk() (*[]RuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Rule) SetActions(v []RuleAction)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


