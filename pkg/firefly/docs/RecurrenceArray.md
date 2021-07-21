# RecurrenceArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RecurrenceRead**](RecurrenceRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | [**PageLink**](PageLink.md) |  | 

## Methods

### NewRecurrenceArray

`func NewRecurrenceArray(data []RecurrenceRead, meta Meta, links PageLink, ) *RecurrenceArray`

NewRecurrenceArray instantiates a new RecurrenceArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceArrayWithDefaults

`func NewRecurrenceArrayWithDefaults() *RecurrenceArray`

NewRecurrenceArrayWithDefaults instantiates a new RecurrenceArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RecurrenceArray) GetData() []RecurrenceRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RecurrenceArray) GetDataOk() (*[]RecurrenceRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RecurrenceArray) SetData(v []RecurrenceRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *RecurrenceArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RecurrenceArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RecurrenceArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *RecurrenceArray) GetLinks() PageLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RecurrenceArray) GetLinksOk() (*PageLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RecurrenceArray) SetLinks(v PageLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


