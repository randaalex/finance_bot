# CategoryRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Category**](Category.md) |  | 

## Methods

### NewCategoryRead

`func NewCategoryRead(type_ string, id string, attributes Category, ) *CategoryRead`

NewCategoryRead instantiates a new CategoryRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryReadWithDefaults

`func NewCategoryReadWithDefaults() *CategoryRead`

NewCategoryReadWithDefaults instantiates a new CategoryRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CategoryRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoryRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoryRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CategoryRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CategoryRead) GetAttributes() Category`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CategoryRead) GetAttributesOk() (*Category, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CategoryRead) SetAttributes(v Category)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


