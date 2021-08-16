# BillStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code | [optional] 
**Name** | **string** |  | 
**AmountMin** | **string** |  | 
**AmountMax** | **string** |  | 
**Date** | **string** |  | 
**RepeatFreq** | **string** | How often the bill must be paid. | 
**Skip** | Pointer to **int32** | How often the bill must be skipped. 1 means a bi-monthly bill. | [optional] 
**Active** | Pointer to **bool** | If the bill is active. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewBillStore

`func NewBillStore(name string, amountMin string, amountMax string, date string, repeatFreq string, ) *BillStore`

NewBillStore instantiates a new BillStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillStoreWithDefaults

`func NewBillStoreWithDefaults() *BillStore`

NewBillStoreWithDefaults instantiates a new BillStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *BillStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BillStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BillStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BillStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BillStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BillStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BillStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BillStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetName

`func (o *BillStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillStore) SetName(v string)`

SetName sets Name field to given value.


### GetAmountMin

`func (o *BillStore) GetAmountMin() string`

GetAmountMin returns the AmountMin field if non-nil, zero value otherwise.

### GetAmountMinOk

`func (o *BillStore) GetAmountMinOk() (*string, bool)`

GetAmountMinOk returns a tuple with the AmountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMin

`func (o *BillStore) SetAmountMin(v string)`

SetAmountMin sets AmountMin field to given value.


### GetAmountMax

`func (o *BillStore) GetAmountMax() string`

GetAmountMax returns the AmountMax field if non-nil, zero value otherwise.

### GetAmountMaxOk

`func (o *BillStore) GetAmountMaxOk() (*string, bool)`

GetAmountMaxOk returns a tuple with the AmountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMax

`func (o *BillStore) SetAmountMax(v string)`

SetAmountMax sets AmountMax field to given value.


### GetDate

`func (o *BillStore) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BillStore) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BillStore) SetDate(v string)`

SetDate sets Date field to given value.


### GetRepeatFreq

`func (o *BillStore) GetRepeatFreq() string`

GetRepeatFreq returns the RepeatFreq field if non-nil, zero value otherwise.

### GetRepeatFreqOk

`func (o *BillStore) GetRepeatFreqOk() (*string, bool)`

GetRepeatFreqOk returns a tuple with the RepeatFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatFreq

`func (o *BillStore) SetRepeatFreq(v string)`

SetRepeatFreq sets RepeatFreq field to given value.


### GetSkip

`func (o *BillStore) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BillStore) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BillStore) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BillStore) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetActive

`func (o *BillStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BillStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BillStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BillStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *BillStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BillStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BillStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BillStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BillStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BillStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetObjectGroupId

`func (o *BillStore) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *BillStore) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *BillStore) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *BillStore) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *BillStore) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *BillStore) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupTitle

`func (o *BillStore) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *BillStore) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *BillStore) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *BillStore) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *BillStore) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *BillStore) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


