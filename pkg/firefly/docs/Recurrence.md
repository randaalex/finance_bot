# Recurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **string** | Not to be confused with the description of the actual transaction(s) being created. | [optional] 
**FirstDate** | **string** | First time the recurring transaction will fire. Must be after today. | 
**LatestDate** | Pointer to **string** | First time the recurring transaction will fire. Must be after today. | [optional] 
**RepeatUntil** | Pointer to **string** | Date until the recurring transaction can fire. Use either this field or repetitions. | [optional] 
**NrOfRepetitions** | Pointer to **int32** | Max number of created transactions. Use either this field or repeat_until. | [optional] 
**ApplyRules** | Pointer to **bool** | Whether or not to fire the rules after the creation of a transaction. | [optional] 
**Active** | Pointer to **bool** | If the recurrence is even active. | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Repetitions** | Pointer to [**[]RecurrenceRepetition**](RecurrenceRepetition.md) |  | [optional] 
**Transactions** | Pointer to [**[]RecurrenceTransaction**](RecurrenceTransaction.md) |  | [optional] 

## Methods

### NewRecurrence

`func NewRecurrence(type_ string, title string, firstDate string, ) *Recurrence`

NewRecurrence instantiates a new Recurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceWithDefaults

`func NewRecurrenceWithDefaults() *Recurrence`

NewRecurrenceWithDefaults instantiates a new Recurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Recurrence) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Recurrence) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Recurrence) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Recurrence) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Recurrence) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Recurrence) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Recurrence) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Recurrence) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *Recurrence) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Recurrence) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Recurrence) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *Recurrence) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Recurrence) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Recurrence) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Recurrence) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Recurrence) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Recurrence) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Recurrence) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstDate

`func (o *Recurrence) GetFirstDate() string`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *Recurrence) GetFirstDateOk() (*string, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *Recurrence) SetFirstDate(v string)`

SetFirstDate sets FirstDate field to given value.


### GetLatestDate

`func (o *Recurrence) GetLatestDate() string`

GetLatestDate returns the LatestDate field if non-nil, zero value otherwise.

### GetLatestDateOk

`func (o *Recurrence) GetLatestDateOk() (*string, bool)`

GetLatestDateOk returns a tuple with the LatestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDate

`func (o *Recurrence) SetLatestDate(v string)`

SetLatestDate sets LatestDate field to given value.

### HasLatestDate

`func (o *Recurrence) HasLatestDate() bool`

HasLatestDate returns a boolean if a field has been set.

### GetRepeatUntil

`func (o *Recurrence) GetRepeatUntil() string`

GetRepeatUntil returns the RepeatUntil field if non-nil, zero value otherwise.

### GetRepeatUntilOk

`func (o *Recurrence) GetRepeatUntilOk() (*string, bool)`

GetRepeatUntilOk returns a tuple with the RepeatUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatUntil

`func (o *Recurrence) SetRepeatUntil(v string)`

SetRepeatUntil sets RepeatUntil field to given value.

### HasRepeatUntil

`func (o *Recurrence) HasRepeatUntil() bool`

HasRepeatUntil returns a boolean if a field has been set.

### GetNrOfRepetitions

`func (o *Recurrence) GetNrOfRepetitions() int32`

GetNrOfRepetitions returns the NrOfRepetitions field if non-nil, zero value otherwise.

### GetNrOfRepetitionsOk

`func (o *Recurrence) GetNrOfRepetitionsOk() (*int32, bool)`

GetNrOfRepetitionsOk returns a tuple with the NrOfRepetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfRepetitions

`func (o *Recurrence) SetNrOfRepetitions(v int32)`

SetNrOfRepetitions sets NrOfRepetitions field to given value.

### HasNrOfRepetitions

`func (o *Recurrence) HasNrOfRepetitions() bool`

HasNrOfRepetitions returns a boolean if a field has been set.

### GetApplyRules

`func (o *Recurrence) GetApplyRules() bool`

GetApplyRules returns the ApplyRules field if non-nil, zero value otherwise.

### GetApplyRulesOk

`func (o *Recurrence) GetApplyRulesOk() (*bool, bool)`

GetApplyRulesOk returns a tuple with the ApplyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRules

`func (o *Recurrence) SetApplyRules(v bool)`

SetApplyRules sets ApplyRules field to given value.

### HasApplyRules

`func (o *Recurrence) HasApplyRules() bool`

HasApplyRules returns a boolean if a field has been set.

### GetActive

`func (o *Recurrence) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Recurrence) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Recurrence) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Recurrence) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNotes

`func (o *Recurrence) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Recurrence) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Recurrence) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Recurrence) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRepetitions

`func (o *Recurrence) GetRepetitions() []RecurrenceRepetition`

GetRepetitions returns the Repetitions field if non-nil, zero value otherwise.

### GetRepetitionsOk

`func (o *Recurrence) GetRepetitionsOk() (*[]RecurrenceRepetition, bool)`

GetRepetitionsOk returns a tuple with the Repetitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitions

`func (o *Recurrence) SetRepetitions(v []RecurrenceRepetition)`

SetRepetitions sets Repetitions field to given value.

### HasRepetitions

`func (o *Recurrence) HasRepetitions() bool`

HasRepetitions returns a boolean if a field has been set.

### GetTransactions

`func (o *Recurrence) GetTransactions() []RecurrenceTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Recurrence) GetTransactionsOk() (*[]RecurrenceTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Recurrence) SetTransactions(v []RecurrenceTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Recurrence) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


