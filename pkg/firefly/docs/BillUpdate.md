# BillUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AmountMin** | Pointer to **string** |  | [optional] 
**AmountMax** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**RepeatFreq** | Pointer to **string** | How often the bill must be paid. | [optional] 
**Skip** | Pointer to **int32** | How often the bill must be skipped. 1 means a bi-monthly bill. | [optional] 
**Active** | Pointer to **bool** | If the bill is active. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewBillUpdate

`func NewBillUpdate() *BillUpdate`

NewBillUpdate instantiates a new BillUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillUpdateWithDefaults

`func NewBillUpdateWithDefaults() *BillUpdate`

NewBillUpdateWithDefaults instantiates a new BillUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *BillUpdate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BillUpdate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BillUpdate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BillUpdate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BillUpdate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BillUpdate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BillUpdate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BillUpdate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetName

`func (o *BillUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAmountMin

`func (o *BillUpdate) GetAmountMin() string`

GetAmountMin returns the AmountMin field if non-nil, zero value otherwise.

### GetAmountMinOk

`func (o *BillUpdate) GetAmountMinOk() (*string, bool)`

GetAmountMinOk returns a tuple with the AmountMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMin

`func (o *BillUpdate) SetAmountMin(v string)`

SetAmountMin sets AmountMin field to given value.

### HasAmountMin

`func (o *BillUpdate) HasAmountMin() bool`

HasAmountMin returns a boolean if a field has been set.

### GetAmountMax

`func (o *BillUpdate) GetAmountMax() string`

GetAmountMax returns the AmountMax field if non-nil, zero value otherwise.

### GetAmountMaxOk

`func (o *BillUpdate) GetAmountMaxOk() (*string, bool)`

GetAmountMaxOk returns a tuple with the AmountMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMax

`func (o *BillUpdate) SetAmountMax(v string)`

SetAmountMax sets AmountMax field to given value.

### HasAmountMax

`func (o *BillUpdate) HasAmountMax() bool`

HasAmountMax returns a boolean if a field has been set.

### GetDate

`func (o *BillUpdate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BillUpdate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BillUpdate) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BillUpdate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetRepeatFreq

`func (o *BillUpdate) GetRepeatFreq() string`

GetRepeatFreq returns the RepeatFreq field if non-nil, zero value otherwise.

### GetRepeatFreqOk

`func (o *BillUpdate) GetRepeatFreqOk() (*string, bool)`

GetRepeatFreqOk returns a tuple with the RepeatFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatFreq

`func (o *BillUpdate) SetRepeatFreq(v string)`

SetRepeatFreq sets RepeatFreq field to given value.

### HasRepeatFreq

`func (o *BillUpdate) HasRepeatFreq() bool`

HasRepeatFreq returns a boolean if a field has been set.

### GetSkip

`func (o *BillUpdate) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BillUpdate) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BillUpdate) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BillUpdate) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetActive

`func (o *BillUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BillUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BillUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BillUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *BillUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BillUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BillUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BillUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BillUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BillUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetObjectGroupId

`func (o *BillUpdate) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *BillUpdate) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *BillUpdate) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *BillUpdate) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *BillUpdate) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *BillUpdate) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupTitle

`func (o *BillUpdate) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *BillUpdate) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *BillUpdate) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *BillUpdate) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *BillUpdate) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *BillUpdate) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


