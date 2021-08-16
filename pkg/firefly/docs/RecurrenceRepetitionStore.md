# RecurrenceRepetitionStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the repetition. ndom means: the n-th weekday of the month, where you can also specify which day of the week. | 
**Moment** | **string** | Information that defined the type of repetition. - For &#39;daily&#39;, this is empty. - For &#39;weekly&#39;, it is day of the week between 1 and 7 (Monday - Sunday). - For &#39;ndom&#39;, it is &#39;1,2&#39; or &#39;4,5&#39; or something else, where the first number is the week in the month, and the second number is the day in the week (between 1 and 7). &#39;2,3&#39; means: the 2nd Wednesday of the month - For &#39;monthly&#39; it is the day of the month (1 - 31) - For yearly, it is a full date, ie &#39;2018-09-17&#39;. The year you use does not matter.  | 
**Skip** | Pointer to **int32** | How many occurrences to skip. 0 means skip nothing. 1 means every other. | [optional] 
**Weekend** | Pointer to **int32** | How to respond when the recurring transaction falls in the weekend. Possible values: 1. Do nothing, just create it 2. Create no transaction. 3. Skip to the previous Friday. 4. Skip to the next Monday.  | [optional] 

## Methods

### NewRecurrenceRepetitionStore

`func NewRecurrenceRepetitionStore(type_ string, moment string, ) *RecurrenceRepetitionStore`

NewRecurrenceRepetitionStore instantiates a new RecurrenceRepetitionStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceRepetitionStoreWithDefaults

`func NewRecurrenceRepetitionStoreWithDefaults() *RecurrenceRepetitionStore`

NewRecurrenceRepetitionStoreWithDefaults instantiates a new RecurrenceRepetitionStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecurrenceRepetitionStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurrenceRepetitionStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurrenceRepetitionStore) SetType(v string)`

SetType sets Type field to given value.


### GetMoment

`func (o *RecurrenceRepetitionStore) GetMoment() string`

GetMoment returns the Moment field if non-nil, zero value otherwise.

### GetMomentOk

`func (o *RecurrenceRepetitionStore) GetMomentOk() (*string, bool)`

GetMomentOk returns a tuple with the Moment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoment

`func (o *RecurrenceRepetitionStore) SetMoment(v string)`

SetMoment sets Moment field to given value.


### GetSkip

`func (o *RecurrenceRepetitionStore) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *RecurrenceRepetitionStore) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *RecurrenceRepetitionStore) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *RecurrenceRepetitionStore) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetWeekend

`func (o *RecurrenceRepetitionStore) GetWeekend() int32`

GetWeekend returns the Weekend field if non-nil, zero value otherwise.

### GetWeekendOk

`func (o *RecurrenceRepetitionStore) GetWeekendOk() (*int32, bool)`

GetWeekendOk returns a tuple with the Weekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekend

`func (o *RecurrenceRepetitionStore) SetWeekend(v int32)`

SetWeekend sets Weekend field to given value.

### HasWeekend

`func (o *RecurrenceRepetitionStore) HasWeekend() bool`

HasWeekend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


