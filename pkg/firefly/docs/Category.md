# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Spent** | Pointer to [**[]CategorySpent**](CategorySpent.md) |  | [optional] [readonly] 
**Earned** | Pointer to [**[]CategoryEarned**](CategoryEarned.md) |  | [optional] [readonly] 

## Methods

### NewCategory

`func NewCategory(name string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Category) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Category) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Category) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Category) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Category) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Category) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Category) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Category) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *Category) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Category) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Category) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Category) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Category) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Category) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSpent

`func (o *Category) GetSpent() []CategorySpent`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *Category) GetSpentOk() (*[]CategorySpent, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *Category) SetSpent(v []CategorySpent)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *Category) HasSpent() bool`

HasSpent returns a boolean if a field has been set.

### GetEarned

`func (o *Category) GetEarned() []CategoryEarned`

GetEarned returns the Earned field if non-nil, zero value otherwise.

### GetEarnedOk

`func (o *Category) GetEarnedOk() (*[]CategoryEarned, bool)`

GetEarnedOk returns a tuple with the Earned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarned

`func (o *Category) SetEarned(v []CategoryEarned)`

SetEarned sets Earned field to given value.

### HasEarned

`func (o *Category) HasEarned() bool`

HasEarned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


