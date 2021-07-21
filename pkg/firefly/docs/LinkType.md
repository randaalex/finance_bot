# LinkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Inward** | **string** |  | 
**Outward** | **string** |  | 
**Editable** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewLinkType

`func NewLinkType(name string, inward string, outward string, ) *LinkType`

NewLinkType instantiates a new LinkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTypeWithDefaults

`func NewLinkTypeWithDefaults() *LinkType`

NewLinkTypeWithDefaults instantiates a new LinkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinkType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkType) SetName(v string)`

SetName sets Name field to given value.


### GetInward

`func (o *LinkType) GetInward() string`

GetInward returns the Inward field if non-nil, zero value otherwise.

### GetInwardOk

`func (o *LinkType) GetInwardOk() (*string, bool)`

GetInwardOk returns a tuple with the Inward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInward

`func (o *LinkType) SetInward(v string)`

SetInward sets Inward field to given value.


### GetOutward

`func (o *LinkType) GetOutward() string`

GetOutward returns the Outward field if non-nil, zero value otherwise.

### GetOutwardOk

`func (o *LinkType) GetOutwardOk() (*string, bool)`

GetOutwardOk returns a tuple with the Outward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutward

`func (o *LinkType) SetOutward(v string)`

SetOutward sets Outward field to given value.


### GetEditable

`func (o *LinkType) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *LinkType) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *LinkType) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *LinkType) HasEditable() bool`

HasEditable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


