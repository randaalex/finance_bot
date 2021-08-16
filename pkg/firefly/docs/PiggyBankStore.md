# PiggyBankStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AccountId** | **string** | The ID of the asset account this piggy bank is connected to. | 
**TargetAmount** | **string** |  | 
**CurrentAmount** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** | The date you started with this piggy bank. | [optional] 
**TargetDate** | Pointer to **NullableString** | The date you intend to finish saving money. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewPiggyBankStore

`func NewPiggyBankStore(name string, accountId string, targetAmount string, ) *PiggyBankStore`

NewPiggyBankStore instantiates a new PiggyBankStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankStoreWithDefaults

`func NewPiggyBankStoreWithDefaults() *PiggyBankStore`

NewPiggyBankStoreWithDefaults instantiates a new PiggyBankStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PiggyBankStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PiggyBankStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PiggyBankStore) SetName(v string)`

SetName sets Name field to given value.


### GetAccountId

`func (o *PiggyBankStore) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PiggyBankStore) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PiggyBankStore) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTargetAmount

`func (o *PiggyBankStore) GetTargetAmount() string`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *PiggyBankStore) GetTargetAmountOk() (*string, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *PiggyBankStore) SetTargetAmount(v string)`

SetTargetAmount sets TargetAmount field to given value.


### GetCurrentAmount

`func (o *PiggyBankStore) GetCurrentAmount() string`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *PiggyBankStore) GetCurrentAmountOk() (*string, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *PiggyBankStore) SetCurrentAmount(v string)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *PiggyBankStore) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetStartDate

`func (o *PiggyBankStore) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PiggyBankStore) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PiggyBankStore) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PiggyBankStore) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTargetDate

`func (o *PiggyBankStore) GetTargetDate() string`

GetTargetDate returns the TargetDate field if non-nil, zero value otherwise.

### GetTargetDateOk

`func (o *PiggyBankStore) GetTargetDateOk() (*string, bool)`

GetTargetDateOk returns a tuple with the TargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDate

`func (o *PiggyBankStore) SetTargetDate(v string)`

SetTargetDate sets TargetDate field to given value.

### HasTargetDate

`func (o *PiggyBankStore) HasTargetDate() bool`

HasTargetDate returns a boolean if a field has been set.

### SetTargetDateNil

`func (o *PiggyBankStore) SetTargetDateNil(b bool)`

 SetTargetDateNil sets the value for TargetDate to be an explicit nil

### UnsetTargetDate
`func (o *PiggyBankStore) UnsetTargetDate()`

UnsetTargetDate ensures that no value is present for TargetDate, not even an explicit nil
### GetOrder

`func (o *PiggyBankStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PiggyBankStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PiggyBankStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PiggyBankStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *PiggyBankStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PiggyBankStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PiggyBankStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PiggyBankStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *PiggyBankStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PiggyBankStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PiggyBankStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PiggyBankStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PiggyBankStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PiggyBankStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetObjectGroupId

`func (o *PiggyBankStore) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *PiggyBankStore) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *PiggyBankStore) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *PiggyBankStore) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *PiggyBankStore) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *PiggyBankStore) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupTitle

`func (o *PiggyBankStore) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *PiggyBankStore) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *PiggyBankStore) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *PiggyBankStore) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *PiggyBankStore) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *PiggyBankStore) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


