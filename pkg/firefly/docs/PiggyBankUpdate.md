# PiggyBankUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** | The ID of the asset account this piggy bank is connected to. | [optional] 
**CurrencyId** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyCode** | Pointer to **string** |  | [optional] [readonly] 
**TargetAmount** | Pointer to **string** |  | [optional] 
**CurrentAmount** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** | The date you started with this piggy bank. | [optional] 
**TargetDate** | Pointer to **NullableString** | The date you intend to finish saving money. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewPiggyBankUpdate

`func NewPiggyBankUpdate() *PiggyBankUpdate`

NewPiggyBankUpdate instantiates a new PiggyBankUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankUpdateWithDefaults

`func NewPiggyBankUpdateWithDefaults() *PiggyBankUpdate`

NewPiggyBankUpdateWithDefaults instantiates a new PiggyBankUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PiggyBankUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PiggyBankUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PiggyBankUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PiggyBankUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *PiggyBankUpdate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PiggyBankUpdate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PiggyBankUpdate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PiggyBankUpdate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PiggyBankUpdate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PiggyBankUpdate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PiggyBankUpdate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PiggyBankUpdate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PiggyBankUpdate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PiggyBankUpdate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PiggyBankUpdate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PiggyBankUpdate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetTargetAmount

`func (o *PiggyBankUpdate) GetTargetAmount() string`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *PiggyBankUpdate) GetTargetAmountOk() (*string, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *PiggyBankUpdate) SetTargetAmount(v string)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *PiggyBankUpdate) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.

### GetCurrentAmount

`func (o *PiggyBankUpdate) GetCurrentAmount() string`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *PiggyBankUpdate) GetCurrentAmountOk() (*string, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *PiggyBankUpdate) SetCurrentAmount(v string)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *PiggyBankUpdate) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetStartDate

`func (o *PiggyBankUpdate) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PiggyBankUpdate) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PiggyBankUpdate) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PiggyBankUpdate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTargetDate

`func (o *PiggyBankUpdate) GetTargetDate() string`

GetTargetDate returns the TargetDate field if non-nil, zero value otherwise.

### GetTargetDateOk

`func (o *PiggyBankUpdate) GetTargetDateOk() (*string, bool)`

GetTargetDateOk returns a tuple with the TargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDate

`func (o *PiggyBankUpdate) SetTargetDate(v string)`

SetTargetDate sets TargetDate field to given value.

### HasTargetDate

`func (o *PiggyBankUpdate) HasTargetDate() bool`

HasTargetDate returns a boolean if a field has been set.

### SetTargetDateNil

`func (o *PiggyBankUpdate) SetTargetDateNil(b bool)`

 SetTargetDateNil sets the value for TargetDate to be an explicit nil

### UnsetTargetDate
`func (o *PiggyBankUpdate) UnsetTargetDate()`

UnsetTargetDate ensures that no value is present for TargetDate, not even an explicit nil
### GetOrder

`func (o *PiggyBankUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PiggyBankUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PiggyBankUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PiggyBankUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *PiggyBankUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PiggyBankUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PiggyBankUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PiggyBankUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *PiggyBankUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PiggyBankUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PiggyBankUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PiggyBankUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PiggyBankUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PiggyBankUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetObjectGroupId

`func (o *PiggyBankUpdate) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *PiggyBankUpdate) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *PiggyBankUpdate) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *PiggyBankUpdate) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *PiggyBankUpdate) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *PiggyBankUpdate) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupTitle

`func (o *PiggyBankUpdate) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *PiggyBankUpdate) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *PiggyBankUpdate) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *PiggyBankUpdate) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *PiggyBankUpdate) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *PiggyBankUpdate) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


