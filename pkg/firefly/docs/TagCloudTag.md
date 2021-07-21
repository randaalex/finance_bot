# TagCloudTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | The tag | [optional] 
**Id** | Pointer to **string** | ID of the tag | [optional] 
**Size** | Pointer to **float64** | The total amount of money related to this tag. There is no currency information available, and this is a basic sum of all amounts added together. | [optional] 
**Relative** | Pointer to **float64** | A number between 0 and 1. 1 is given to the largest tag in the tag cloud, and 0 to the smallest. The rest are given a number between 0 and 1, related to their size in comparison to the largest tag. | [optional] 

## Methods

### NewTagCloudTag

`func NewTagCloudTag() *TagCloudTag`

NewTagCloudTag instantiates a new TagCloudTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCloudTagWithDefaults

`func NewTagCloudTagWithDefaults() *TagCloudTag`

NewTagCloudTagWithDefaults instantiates a new TagCloudTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *TagCloudTag) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagCloudTag) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagCloudTag) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TagCloudTag) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetId

`func (o *TagCloudTag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagCloudTag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagCloudTag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TagCloudTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *TagCloudTag) GetSize() float64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TagCloudTag) GetSizeOk() (*float64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TagCloudTag) SetSize(v float64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TagCloudTag) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetRelative

`func (o *TagCloudTag) GetRelative() float64`

GetRelative returns the Relative field if non-nil, zero value otherwise.

### GetRelativeOk

`func (o *TagCloudTag) GetRelativeOk() (*float64, bool)`

GetRelativeOk returns a tuple with the Relative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelative

`func (o *TagCloudTag) SetRelative(v float64)`

SetRelative sets Relative field to given value.

### HasRelative

`func (o *TagCloudTag) HasRelative() bool`

HasRelative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


