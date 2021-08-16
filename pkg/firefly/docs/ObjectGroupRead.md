# ObjectGroupRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**ObjectGroup**](ObjectGroup.md) |  | 

## Methods

### NewObjectGroupRead

`func NewObjectGroupRead(type_ string, id string, attributes ObjectGroup, ) *ObjectGroupRead`

NewObjectGroupRead instantiates a new ObjectGroupRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupReadWithDefaults

`func NewObjectGroupReadWithDefaults() *ObjectGroupRead`

NewObjectGroupReadWithDefaults instantiates a new ObjectGroupRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ObjectGroupRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectGroupRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectGroupRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ObjectGroupRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectGroupRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectGroupRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ObjectGroupRead) GetAttributes() ObjectGroup`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ObjectGroupRead) GetAttributesOk() (*ObjectGroup, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ObjectGroupRead) SetAttributes(v ObjectGroup)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


