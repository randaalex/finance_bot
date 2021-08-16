# RuleGroupStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewRuleGroupStore

`func NewRuleGroupStore(title string, ) *RuleGroupStore`

NewRuleGroupStore instantiates a new RuleGroupStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleGroupStoreWithDefaults

`func NewRuleGroupStoreWithDefaults() *RuleGroupStore`

NewRuleGroupStoreWithDefaults instantiates a new RuleGroupStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RuleGroupStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RuleGroupStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RuleGroupStore) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RuleGroupStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleGroupStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleGroupStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleGroupStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RuleGroupStore) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RuleGroupStore) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *RuleGroupStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleGroupStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleGroupStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleGroupStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *RuleGroupStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RuleGroupStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RuleGroupStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RuleGroupStore) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


