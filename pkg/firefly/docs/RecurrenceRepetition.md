# RecurrenceRepetition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | **string** | The type of the repetition. ndom means: the n-th weekday of the month, where you can also specify which day of the week. | 
**Moment** | **string** | Information that defined the type of repetition. - For &#39;daily&#39;, this is empty. - For &#39;weekly&#39;, it is day of the week between 1 and 7 (Monday - Sunday). - For &#39;ndom&#39;, it is &#39;1,2&#39; or &#39;4,5&#39; or something else, where the first number is the week in the month, and the second number is the day in the week (between 1 and 7). &#39;2,3&#39; means: the 2nd Wednesday of the month - For &#39;monthly&#39; it is the day of the month (1 - 31) - For yearly, it is a full date, ie &#39;2018-09-17&#39;. The year you use does not matter.  | 
**Skip** | Pointer to **int32** | How many occurrences to skip. 0 means skip nothing. 1 means every other. | [optional] 
**Weekend** | Pointer to **int32** | How to respond when the recurring transaction falls in the weekend. Possible values: 1. Do nothing, just create it 2. Create no transaction. 3. Skip to the previous Friday. 4. Skip to the next Monday.  | [optional] 
**Description** | Pointer to **string** | Auto-generated repetition description. | [optional] [readonly] 
**Occurrences** | Pointer to [**[]time.Time**](time.Time.md) | Array of future dates when the repetition will apply to. Auto generated. | [optional] [readonly] 

## Methods

### NewRecurrenceRepetition

`func NewRecurrenceRepetition(type_ string, moment string, ) *RecurrenceRepetition`

NewRecurrenceRepetition instantiates a new RecurrenceRepetition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceRepetitionWithDefaults

`func NewRecurrenceRepetitionWithDefaults() *RecurrenceRepetition`

NewRecurrenceRepetitionWithDefaults instantiates a new RecurrenceRepetition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurrenceRepetition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurrenceRepetition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurrenceRepetition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecurrenceRepetition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RecurrenceRepetition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecurrenceRepetition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecurrenceRepetition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RecurrenceRepetition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RecurrenceRepetition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecurrenceRepetition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecurrenceRepetition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecurrenceRepetition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *RecurrenceRepetition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurrenceRepetition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurrenceRepetition) SetType(v string)`

SetType sets Type field to given value.


### GetMoment

`func (o *RecurrenceRepetition) GetMoment() string`

GetMoment returns the Moment field if non-nil, zero value otherwise.

### GetMomentOk

`func (o *RecurrenceRepetition) GetMomentOk() (*string, bool)`

GetMomentOk returns a tuple with the Moment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoment

`func (o *RecurrenceRepetition) SetMoment(v string)`

SetMoment sets Moment field to given value.


### GetSkip

`func (o *RecurrenceRepetition) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *RecurrenceRepetition) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *RecurrenceRepetition) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *RecurrenceRepetition) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetWeekend

`func (o *RecurrenceRepetition) GetWeekend() int32`

GetWeekend returns the Weekend field if non-nil, zero value otherwise.

### GetWeekendOk

`func (o *RecurrenceRepetition) GetWeekendOk() (*int32, bool)`

GetWeekendOk returns a tuple with the Weekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekend

`func (o *RecurrenceRepetition) SetWeekend(v int32)`

SetWeekend sets Weekend field to given value.

### HasWeekend

`func (o *RecurrenceRepetition) HasWeekend() bool`

HasWeekend returns a boolean if a field has been set.

### GetDescription

`func (o *RecurrenceRepetition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecurrenceRepetition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecurrenceRepetition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecurrenceRepetition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOccurrences

`func (o *RecurrenceRepetition) GetOccurrences() []time.Time`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *RecurrenceRepetition) GetOccurrencesOk() (*[]time.Time, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *RecurrenceRepetition) SetOccurrences(v []time.Time)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *RecurrenceRepetition) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


