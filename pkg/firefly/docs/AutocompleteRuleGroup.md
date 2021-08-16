# AutocompleteRuleGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of the rule group found by an auto-complete search. | 
**Description** | Pointer to **string** | Description of the rule group found by auto-complete. | [optional] 

## Methods

### NewAutocompleteRuleGroup

`func NewAutocompleteRuleGroup(id string, name string, ) *AutocompleteRuleGroup`

NewAutocompleteRuleGroup instantiates a new AutocompleteRuleGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteRuleGroupWithDefaults

`func NewAutocompleteRuleGroupWithDefaults() *AutocompleteRuleGroup`

NewAutocompleteRuleGroupWithDefaults instantiates a new AutocompleteRuleGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompleteRuleGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompleteRuleGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompleteRuleGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutocompleteRuleGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompleteRuleGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompleteRuleGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AutocompleteRuleGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutocompleteRuleGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutocompleteRuleGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutocompleteRuleGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


