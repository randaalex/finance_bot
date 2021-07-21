# PreferenceArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PreferenceRead**](PreferenceRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewPreferenceArray

`func NewPreferenceArray(data []PreferenceRead, meta Meta, links PageLink, ) *PreferenceArray`

NewPreferenceArray instantiates a new PreferenceArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferenceArrayWithDefaults

`func NewPreferenceArrayWithDefaults() *PreferenceArray`

NewPreferenceArrayWithDefaults instantiates a new PreferenceArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PreferenceArray) GetData() []PreferenceRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PreferenceArray) GetDataOk() (*[]PreferenceRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PreferenceArray) SetData(v []PreferenceRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *PreferenceArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PreferenceArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PreferenceArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PreferenceArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PreferenceArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PreferenceArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


