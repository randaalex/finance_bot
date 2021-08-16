# PiggyBank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**AccountId** | **string** | The ID of the asset account this piggy bank is connected to. | 
**AccountName** | Pointer to **string** | The name of the asset account this piggy bank is connected to. | [optional] [readonly] 
**CurrencyId** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyCode** | Pointer to **string** |  | [optional] [readonly] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** | Number of decimals supported by the currency | [optional] [readonly] 
**TargetAmount** | **string** |  | 
**Percentage** | Pointer to **float32** |  | [optional] [readonly] 
**CurrentAmount** | Pointer to **string** |  | [optional] 
**LeftToSave** | Pointer to **string** |  | [optional] [readonly] 
**SavePerMonth** | Pointer to **string** |  | [optional] [readonly] 
**StartDate** | Pointer to **time.Time** | The date you started with this piggy bank. | [optional] 
**TargetDate** | Pointer to **NullableTime** | The date you intend to finish saving money. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**ObjectGroupId** | Pointer to **NullableString** | The group ID of the group this object is part of. NULL if no group. | [optional] 
**ObjectGroupOrder** | Pointer to **NullableInt32** | The order of the group. At least 1, for the highest sorting. | [optional] [readonly] 
**ObjectGroupTitle** | Pointer to **NullableString** | The name of the group. NULL if no group. | [optional] 

## Methods

### NewPiggyBank

`func NewPiggyBank(name string, accountId string, targetAmount string, ) *PiggyBank`

NewPiggyBank instantiates a new PiggyBank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPiggyBankWithDefaults

`func NewPiggyBankWithDefaults() *PiggyBank`

NewPiggyBankWithDefaults instantiates a new PiggyBank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PiggyBank) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PiggyBank) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PiggyBank) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PiggyBank) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PiggyBank) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PiggyBank) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PiggyBank) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PiggyBank) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PiggyBank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PiggyBank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PiggyBank) SetName(v string)`

SetName sets Name field to given value.


### GetAccountId

`func (o *PiggyBank) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PiggyBank) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PiggyBank) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *PiggyBank) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PiggyBank) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PiggyBank) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PiggyBank) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetCurrencyId

`func (o *PiggyBank) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *PiggyBank) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *PiggyBank) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *PiggyBank) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PiggyBank) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PiggyBank) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PiggyBank) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PiggyBank) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *PiggyBank) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *PiggyBank) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *PiggyBank) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *PiggyBank) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *PiggyBank) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *PiggyBank) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *PiggyBank) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *PiggyBank) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetTargetAmount

`func (o *PiggyBank) GetTargetAmount() string`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *PiggyBank) GetTargetAmountOk() (*string, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *PiggyBank) SetTargetAmount(v string)`

SetTargetAmount sets TargetAmount field to given value.


### GetPercentage

`func (o *PiggyBank) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PiggyBank) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PiggyBank) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PiggyBank) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCurrentAmount

`func (o *PiggyBank) GetCurrentAmount() string`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *PiggyBank) GetCurrentAmountOk() (*string, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *PiggyBank) SetCurrentAmount(v string)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *PiggyBank) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetLeftToSave

`func (o *PiggyBank) GetLeftToSave() string`

GetLeftToSave returns the LeftToSave field if non-nil, zero value otherwise.

### GetLeftToSaveOk

`func (o *PiggyBank) GetLeftToSaveOk() (*string, bool)`

GetLeftToSaveOk returns a tuple with the LeftToSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftToSave

`func (o *PiggyBank) SetLeftToSave(v string)`

SetLeftToSave sets LeftToSave field to given value.

### HasLeftToSave

`func (o *PiggyBank) HasLeftToSave() bool`

HasLeftToSave returns a boolean if a field has been set.

### GetSavePerMonth

`func (o *PiggyBank) GetSavePerMonth() string`

GetSavePerMonth returns the SavePerMonth field if non-nil, zero value otherwise.

### GetSavePerMonthOk

`func (o *PiggyBank) GetSavePerMonthOk() (*string, bool)`

GetSavePerMonthOk returns a tuple with the SavePerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePerMonth

`func (o *PiggyBank) SetSavePerMonth(v string)`

SetSavePerMonth sets SavePerMonth field to given value.

### HasSavePerMonth

`func (o *PiggyBank) HasSavePerMonth() bool`

HasSavePerMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *PiggyBank) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PiggyBank) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PiggyBank) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PiggyBank) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTargetDate

`func (o *PiggyBank) GetTargetDate() time.Time`

GetTargetDate returns the TargetDate field if non-nil, zero value otherwise.

### GetTargetDateOk

`func (o *PiggyBank) GetTargetDateOk() (*time.Time, bool)`

GetTargetDateOk returns a tuple with the TargetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDate

`func (o *PiggyBank) SetTargetDate(v time.Time)`

SetTargetDate sets TargetDate field to given value.

### HasTargetDate

`func (o *PiggyBank) HasTargetDate() bool`

HasTargetDate returns a boolean if a field has been set.

### SetTargetDateNil

`func (o *PiggyBank) SetTargetDateNil(b bool)`

 SetTargetDateNil sets the value for TargetDate to be an explicit nil

### UnsetTargetDate
`func (o *PiggyBank) UnsetTargetDate()`

UnsetTargetDate ensures that no value is present for TargetDate, not even an explicit nil
### GetOrder

`func (o *PiggyBank) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PiggyBank) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PiggyBank) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PiggyBank) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetActive

`func (o *PiggyBank) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PiggyBank) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PiggyBank) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PiggyBank) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *PiggyBank) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PiggyBank) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PiggyBank) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PiggyBank) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PiggyBank) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PiggyBank) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetObjectGroupId

`func (o *PiggyBank) GetObjectGroupId() string`

GetObjectGroupId returns the ObjectGroupId field if non-nil, zero value otherwise.

### GetObjectGroupIdOk

`func (o *PiggyBank) GetObjectGroupIdOk() (*string, bool)`

GetObjectGroupIdOk returns a tuple with the ObjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupId

`func (o *PiggyBank) SetObjectGroupId(v string)`

SetObjectGroupId sets ObjectGroupId field to given value.

### HasObjectGroupId

`func (o *PiggyBank) HasObjectGroupId() bool`

HasObjectGroupId returns a boolean if a field has been set.

### SetObjectGroupIdNil

`func (o *PiggyBank) SetObjectGroupIdNil(b bool)`

 SetObjectGroupIdNil sets the value for ObjectGroupId to be an explicit nil

### UnsetObjectGroupId
`func (o *PiggyBank) UnsetObjectGroupId()`

UnsetObjectGroupId ensures that no value is present for ObjectGroupId, not even an explicit nil
### GetObjectGroupOrder

`func (o *PiggyBank) GetObjectGroupOrder() int32`

GetObjectGroupOrder returns the ObjectGroupOrder field if non-nil, zero value otherwise.

### GetObjectGroupOrderOk

`func (o *PiggyBank) GetObjectGroupOrderOk() (*int32, bool)`

GetObjectGroupOrderOk returns a tuple with the ObjectGroupOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupOrder

`func (o *PiggyBank) SetObjectGroupOrder(v int32)`

SetObjectGroupOrder sets ObjectGroupOrder field to given value.

### HasObjectGroupOrder

`func (o *PiggyBank) HasObjectGroupOrder() bool`

HasObjectGroupOrder returns a boolean if a field has been set.

### SetObjectGroupOrderNil

`func (o *PiggyBank) SetObjectGroupOrderNil(b bool)`

 SetObjectGroupOrderNil sets the value for ObjectGroupOrder to be an explicit nil

### UnsetObjectGroupOrder
`func (o *PiggyBank) UnsetObjectGroupOrder()`

UnsetObjectGroupOrder ensures that no value is present for ObjectGroupOrder, not even an explicit nil
### GetObjectGroupTitle

`func (o *PiggyBank) GetObjectGroupTitle() string`

GetObjectGroupTitle returns the ObjectGroupTitle field if non-nil, zero value otherwise.

### GetObjectGroupTitleOk

`func (o *PiggyBank) GetObjectGroupTitleOk() (*string, bool)`

GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroupTitle

`func (o *PiggyBank) SetObjectGroupTitle(v string)`

SetObjectGroupTitle sets ObjectGroupTitle field to given value.

### HasObjectGroupTitle

`func (o *PiggyBank) HasObjectGroupTitle() bool`

HasObjectGroupTitle returns a boolean if a field has been set.

### SetObjectGroupTitleNil

`func (o *PiggyBank) SetObjectGroupTitleNil(b bool)`

 SetObjectGroupTitleNil sets the value for ObjectGroupTitle to be an explicit nil

### UnsetObjectGroupTitle
`func (o *PiggyBank) UnsetObjectGroupTitle()`

UnsetObjectGroupTitle ensures that no value is present for ObjectGroupTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


