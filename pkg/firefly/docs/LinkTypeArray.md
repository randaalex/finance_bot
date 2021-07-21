# LinkTypeArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]LinkTypeRead**](LinkTypeRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewLinkTypeArray

`func NewLinkTypeArray(data []LinkTypeRead, meta Meta, links PageLink, ) *LinkTypeArray`

NewLinkTypeArray instantiates a new LinkTypeArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTypeArrayWithDefaults

`func NewLinkTypeArrayWithDefaults() *LinkTypeArray`

NewLinkTypeArrayWithDefaults instantiates a new LinkTypeArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LinkTypeArray) GetData() []LinkTypeRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkTypeArray) GetDataOk() (*[]LinkTypeRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkTypeArray) SetData(v []LinkTypeRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *LinkTypeArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LinkTypeArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LinkTypeArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *LinkTypeArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkTypeArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkTypeArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


