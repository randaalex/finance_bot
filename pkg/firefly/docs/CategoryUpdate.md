# CategoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCategoryUpdate

`func NewCategoryUpdate() *CategoryUpdate`

NewCategoryUpdate instantiates a new CategoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryUpdateWithDefaults

`func NewCategoryUpdateWithDefaults() *CategoryUpdate`

NewCategoryUpdateWithDefaults instantiates a new CategoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoryUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *CategoryUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CategoryUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CategoryUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CategoryUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CategoryUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CategoryUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


