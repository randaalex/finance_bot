# RuleRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Rule**](Rule.md) |  | 
**Links** | [**ObjectLink**](ObjectLink.md) |  | 

## Methods

### NewRuleRead

`func NewRuleRead(type_ string, id string, attributes Rule, links ObjectLink, ) *RuleRead`

NewRuleRead instantiates a new RuleRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleReadWithDefaults

`func NewRuleReadWithDefaults() *RuleRead`

NewRuleReadWithDefaults instantiates a new RuleRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RuleRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RuleRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *RuleRead) GetAttributes() Rule`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RuleRead) GetAttributesOk() (*Rule, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RuleRead) SetAttributes(v Rule)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *RuleRead) GetLinks() ObjectLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RuleRead) GetLinksOk() (*ObjectLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RuleRead) SetLinks(v ObjectLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


