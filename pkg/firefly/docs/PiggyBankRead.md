# PiggyBankRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**PiggyBank**](PiggyBank.md) |  | 
**Links** | [**ObjectLink**](ObjectLink.md) |  | 

## Methods

### NewPiggyBankRead

`func NewPiggyBankRead(type_ string, id string, attributes PiggyBank, links ObjectLink, ) *PiggyBankRead`

NewPiggyBankRead instantiates a new PiggyBankRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankReadWithDefaults

`func NewPiggyBankReadWithDefaults() *PiggyBankRead`

NewPiggyBankReadWithDefaults instantiates a new PiggyBankRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PiggyBankRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PiggyBankRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PiggyBankRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PiggyBankRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PiggyBankRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PiggyBankRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PiggyBankRead) GetAttributes() PiggyBank`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PiggyBankRead) GetAttributesOk() (*PiggyBank, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PiggyBankRead) SetAttributes(v PiggyBank)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *PiggyBankRead) GetLinks() ObjectLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PiggyBankRead) GetLinksOk() (*ObjectLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PiggyBankRead) SetLinks(v ObjectLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


