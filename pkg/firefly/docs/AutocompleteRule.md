# AutocompleteRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of the rule found by an auto-complete search. | 
**Description** | Pointer to **string** | Description of the rule found by auto-complete. | [optional] 

## Methods

### NewAutocompleteRule

`func NewAutocompleteRule(id string, name string, ) *AutocompleteRule`

NewAutocompleteRule instantiates a new AutocompleteRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteRuleWithDefaults

`func NewAutocompleteRuleWithDefaults() *AutocompleteRule`

NewAutocompleteRuleWithDefaults instantiates a new AutocompleteRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompleteRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompleteRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompleteRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutocompleteRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompleteRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompleteRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AutocompleteRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutocompleteRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutocompleteRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutocompleteRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


