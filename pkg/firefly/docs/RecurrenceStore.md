# RecurrenceStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] 
**FirstDate** | **string** | First time the recurring transaction will fire. Must be after today. | 
**RepeatUntil** | **NullableString** | Date until the recurring transaction can fire. Use either this field or repetitions. | 
**NrOfRepetitions** | Pointer to **NullableInt32** | Max number of created transactions. Use either this field or repeat_until. | [optional] 
**ApplyRules** | Pointer to **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] 
**Active** | Pointer to **bool** | If the recurrence is even active. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Repetitions** | [**[]RecurrenceRepetitionStore**](RecurrenceRepetitionStore.md) |  | 
**Transactions** | [**[]RecurrenceTransactionStore**](RecurrenceTransactionStore.md) |  | 

## Methods

### NewRecurrenceStore

`func NewRecurrenceStore(type_ string, title string, firstDate string, repeatUntil NullableString, repetitions []RecurrenceRepetitionStore, transactions []RecurrenceTransactionStore, ) *RecurrenceStore`

NewRecurrenceStore instantiates a new RecurrenceStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceStoreWithDefaults

`func NewRecurrenceStoreWithDefaults() *RecurrenceStore`

NewRecurrenceStoreWithDefaults instantiates a new RecurrenceStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecurrenceStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurrenceStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurrenceStore) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *RecurrenceStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecurrenceStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecurrenceStore) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RecurrenceStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurrenceStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstDate

`func (o *RecurrenceStore) GetFirstDate() string`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *RecurrenceStore) GetFirstDateOk() (*string, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *RecurrenceStore) SetFirstDate(v string)`

SetFirstDate sets FirstDate field to given value.


### GetRepeatUntil

`func (o *RecurrenceStore) GetRepeatUntil() string`

GetRepeatUntil returns the RepeatUntil field if non-nil, zero value otherwise.

### GetRepeatUntilOk

`func (o *RecurrenceStore) GetRepeatUntilOk() (*string, bool)`

GetRepeatUntilOk returns a tuple with the RepeatUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatUntil

`func (o *RecurrenceStore) SetRepeatUntil(v string)`

SetRepeatUntil sets RepeatUntil field to given value.


### SetRepeatUntilNil

`func (o *RecurrenceStore) SetRepeatUntilNil(b bool)`

 SetRepeatUntilNil sets the value for RepeatUntil to be an explicit nil

### UnsetRepeatUntil
`func (o *RecurrenceStore) UnsetRepeatUntil()`

UnsetRepeatUntil ensures that no value is present for RepeatUntil, not even an explicit nil
### GetNrOfRepetitions

`func (o *RecurrenceStore) GetNrOfRepetitions() int32`

GetNrOfRepetitions returns the NrOfRepetitions field if non-nil, zero value otherwise.

### GetNrOfRepetitionsOk

`func (o *RecurrenceStore) GetNrOfRepetitionsOk() (*int32, bool)`

GetNrOfRepetitionsOk returns a tuple with the NrOfRepetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfRepetitions

`func (o *RecurrenceStore) SetNrOfRepetitions(v int32)`

SetNrOfRepetitions sets NrOfRepetitions field to given value.

### HasNrOfRepetitions

`func (o *RecurrenceStore) HasNrOfRepetitions() bool`

HasNrOfRepetitions returns a boolean if a field has been set.

### SetNrOfRepetitionsNil

`func (o *RecurrenceStore) SetNrOfRepetitionsNil(b bool)`

 SetNrOfRepetitionsNil sets the value for NrOfRepetitions to be an explicit nil

### UnsetNrOfRepetitions
`func (o *RecurrenceStore) UnsetNrOfRepetitions()`

UnsetNrOfRepetitions ensures that no value is present for NrOfRepetitions, not even an explicit nil
### GetApplyRules

`func (o *RecurrenceStore) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *RecurrenceStore) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *RecurrenceStore) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *RecurrenceStore) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetActive

`func (o *RecurrenceStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RecurrenceStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RecurrenceStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RecurrenceStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *RecurrenceStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RecurrenceStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RecurrenceStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RecurrenceStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *RecurrenceStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *RecurrenceStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRepetitions

`func (o *RecurrenceStore) GetRepetitions() []RecurrenceRepetitionStore`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *RecurrenceStore) GetRepetitionsOk() (*[]RecurrenceRepetitionStore, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *RecurrenceStore) SetRepetitions(v []RecurrenceRepetitionStore)`

SetRepetitions sets Repetitions field to given value.


### GetTransactions

`func (o *RecurrenceStore) GetTransactions() []RecurrenceTransactionStore`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *RecurrenceStore) GetTransactionsOk() (*[]RecurrenceTransactionStore, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *RecurrenceStore) SetTransactions(v []RecurrenceTransactionStore)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


