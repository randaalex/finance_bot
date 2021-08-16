# AutocompletePiggy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of the piggy bank found by an auto-complete search. | 
**CurrencyId** | Pointer to **string** | Currency ID for this piggy bank. | [optional] 
**CurrencyCode** | Pointer to **string** | Currency code for this piggy bank. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewAutocompletePiggy

`func NewAutocompletePiggy(id string, name string, ) *AutocompletePiggy`

NewAutocompletePiggy instantiates a new AutocompletePiggy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompletePiggyWithDefaults

`func NewAutocompletePiggyWithDefaults() *AutocompletePiggy`

NewAutocompletePiggyWithDefaults instantiates a new AutocompletePiggy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompletePiggy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompletePiggy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompletePiggy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutocompletePiggy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompletePiggy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompletePiggy) SetName(v string)`

SetName sets Name field to given value.


### GetCurrencyId

`func (o *AutocompletePiggy) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AutocompletePiggy) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AutocompletePiggy) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AutocompletePiggy) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AutocompletePiggy) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutocompletePiggy) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutocompletePiggy) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AutocompletePiggy) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *AutocompletePiggy) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *AutocompletePiggy) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *AutocompletePiggy) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *AutocompletePiggy) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *AutocompletePiggy) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *AutocompletePiggy) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *AutocompletePiggy) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *AutocompletePiggy) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetObjectGroupId

`func (o *AutocompletePiggy) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *AutocompletePiggy) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *AutocompletePiggy) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *AutocompletePiggy) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *AutocompletePiggy) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *AutocompletePiggy) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupTitle

`func (o *AutocompletePiggy) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *AutocompletePiggy) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *AutocompletePiggy) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *AutocompletePiggy) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *AutocompletePiggy) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *AutocompletePiggy) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


