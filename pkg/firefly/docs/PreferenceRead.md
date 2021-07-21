# PreferenceRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Preference**](Preference.md) |  | 

## Methods

### NewPreferenceRead

`func NewPreferenceRead(type_ string, id string, attributes Preference, ) *PreferenceRead`

NewPreferenceRead instantiates a new PreferenceRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferenceReadWithDefaults

`func NewPreferenceReadWithDefaults() *PreferenceRead`

NewPreferenceReadWithDefaults instantiates a new PreferenceRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PreferenceRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreferenceRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreferenceRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PreferenceRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreferenceRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreferenceRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PreferenceRead) GetAttributes() Preference`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PreferenceRead) GetAttributesOk() (*Preference, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PreferenceRead) SetAttributes(v Preference)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


