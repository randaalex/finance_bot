# PiggyBankEventArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PiggyBankEventRead**](PiggyBankEventRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewPiggyBankEventArray

`func NewPiggyBankEventArray(data []PiggyBankEventRead, meta Meta, links PageLink, ) *PiggyBankEventArray`

NewPiggyBankEventArray instantiates a new PiggyBankEventArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankEventArrayWithDefaults

`func NewPiggyBankEventArrayWithDefaults() *PiggyBankEventArray`

NewPiggyBankEventArrayWithDefaults instantiates a new PiggyBankEventArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PiggyBankEventArray) GetData() []PiggyBankEventRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PiggyBankEventArray) GetDataOk() (*[]PiggyBankEventRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PiggyBankEventArray) SetData(v []PiggyBankEventRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *PiggyBankEventArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PiggyBankEventArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PiggyBankEventArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PiggyBankEventArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PiggyBankEventArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PiggyBankEventArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


