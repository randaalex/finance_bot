# RecurrenceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] 
**FirstDate** | Pointer to **string** | First time the recurring transaction will fire. | [optional] 
**RepeatUntil** | Pointer to **NullableString** | Date until the recurring transaction can fire. After that date, it&#39;s basically inactive. Use either this field or repetitions. | [optional] 
**NrOfRepetitions** | Pointer to **NullableInt32** | Max number of created transactions. Use either this field or repeat_until. | [optional] 
**ApplyRules** | Pointer to **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] 
**Active** | Pointer to **bool** | If the recurrence is even active. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Repetitions** | Pointer to [**[]RecurrenceRepetitionUpdate**](RecurrenceRepetitionUpdate.md) |  | [optional] 
**Transactions** | Pointer to [**[]RecurrenceTransactionUpdate**](RecurrenceTransactionUpdate.md) |  | [optional] 

## Methods

### NewRecurrenceUpdate

`func NewRecurrenceUpdate() *RecurrenceUpdate`

NewRecurrenceUpdate instantiates a new RecurrenceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceUpdateWithDefaults

`func NewRecurrenceUpdateWithDefaults() *RecurrenceUpdate`

NewRecurrenceUpdateWithDefaults instantiates a new RecurrenceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RecurrenceUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecurrenceUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecurrenceUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RecurrenceUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RecurrenceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurrenceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstDate

`func (o *RecurrenceUpdate) GetFirstDate() string`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *RecurrenceUpdate) GetFirstDateOk() (*string, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *RecurrenceUpdate) SetFirstDate(v string)`

SetFirstDate sets FirstDate field to given value.

### HasFirstDate

`func (o *RecurrenceUpdate) HasFirstDate() bool`

HasFirstDate returns a boolean if a field has been set.

### GetRepeatUntil

`func (o *RecurrenceUpdate) GetRepeatUntil() string`

GetRepeatUntil returns the RepeatUntil field if non-nil, zero value otherwise.

### GetRepeatUntilOk

`func (o *RecurrenceUpdate) GetRepeatUntilOk() (*string, bool)`

GetRepeatUntilOk returns a tuple with the RepeatUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatUntil

`func (o *RecurrenceUpdate) SetRepeatUntil(v string)`

SetRepeatUntil sets RepeatUntil field to given value.

### HasRepeatUntil

`func (o *RecurrenceUpdate) HasRepeatUntil() bool`

HasRepeatUntil returns a boolean if a field has been set.

### SetRepeatUntilNil

`func (o *RecurrenceUpdate) SetRepeatUntilNil(b bool)`

 SetRepeatUntilNil sets the value for RepeatUntil to be an explicit nil

### UnsetRepeatUntil
`func (o *RecurrenceUpdate) UnsetRepeatUntil()`

UnsetRepeatUntil ensures that no value is present for RepeatUntil, not even an explicit nil
### GetNrOfRepetitions

`func (o *RecurrenceUpdate) GetNrOfRepetitions() int32`

GetNrOfRepetitions returns the NrOfRepetitions field if non-nil, zero value otherwise.

### GetNrOfRepetitionsOk

`func (o *RecurrenceUpdate) GetNrOfRepetitionsOk() (*int32, bool)`

GetNrOfRepetitionsOk returns a tuple with the NrOfRepetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfRepetitions

`func (o *RecurrenceUpdate) SetNrOfRepetitions(v int32)`

SetNrOfRepetitions sets NrOfRepetitions field to given value.

### HasNrOfRepetitions

`func (o *RecurrenceUpdate) HasNrOfRepetitions() bool`

HasNrOfRepetitions returns a boolean if a field has been set.

### SetNrOfRepetitionsNil

`func (o *RecurrenceUpdate) SetNrOfRepetitionsNil(b bool)`

 SetNrOfRepetitionsNil sets the value for NrOfRepetitions to be an explicit nil

### UnsetNrOfRepetitions
`func (o *RecurrenceUpdate) UnsetNrOfRepetitions()`

UnsetNrOfRepetitions ensures that no value is present for NrOfRepetitions, not even an explicit nil
### GetApplyRules

`func (o *RecurrenceUpdate) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *RecurrenceUpdate) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *RecurrenceUpdate) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *RecurrenceUpdate) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetActive

`func (o *RecurrenceUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RecurrenceUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RecurrenceUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RecurrenceUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *RecurrenceUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RecurrenceUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RecurrenceUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RecurrenceUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *RecurrenceUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *RecurrenceUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRepetitions

`func (o *RecurrenceUpdate) GetRepetitions() []RecurrenceRepetitionUpdate`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *RecurrenceUpdate) GetRepetitionsOk() (*[]RecurrenceRepetitionUpdate, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *RecurrenceUpdate) SetRepetitions(v []RecurrenceRepetitionUpdate)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *RecurrenceUpdate) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetTransactions

`func (o *RecurrenceUpdate) GetTransactions() []RecurrenceTransactionUpdate`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *RecurrenceUpdate) GetTransactionsOk() (*[]RecurrenceTransactionUpdate, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *RecurrenceUpdate) SetTransactions(v []RecurrenceTransactionUpdate)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *RecurrenceUpdate) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


