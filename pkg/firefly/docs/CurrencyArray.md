# CurrencyArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CurrencyRead**](CurrencyRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewCurrencyArray

`func NewCurrencyArray(data []CurrencyRead, meta Meta, links PageLink, ) *CurrencyArray`

NewCurrencyArray instantiates a new CurrencyArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyArrayWithDefaults

`func NewCurrencyArrayWithDefaults() *CurrencyArray`

NewCurrencyArrayWithDefaults instantiates a new CurrencyArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CurrencyArray) GetData() []CurrencyRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CurrencyArray) GetDataOk() (*[]CurrencyRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CurrencyArray) SetData(v []CurrencyRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *CurrencyArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CurrencyArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CurrencyArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *CurrencyArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CurrencyArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CurrencyArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


