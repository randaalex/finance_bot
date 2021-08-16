# CronResultRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobFired** | Pointer to **NullableBool** | This value tells you if this specific cron job actually fired. It may not fire. Some cron jobs only fire every 24 hours, for example.  | [optional] 
**JobSucceeded** | Pointer to **NullableBool** | This value tells you if this specific cron job actually did something. The job may fire but not change anything.  | [optional] 
**JobErrored** | Pointer to **NullableBool** | If the cron job ran into some kind of an error, this value will be true. | [optional] 
**Message** | Pointer to **NullableString** | If the cron job ran into some kind of an error, this value will be the error message. The success message if the job actually ran OK.  | [optional] 

## Methods

### NewCronResultRow

`func NewCronResultRow() *CronResultRow`

NewCronResultRow instantiates a new CronResultRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronResultRowWithDefaults

`func NewCronResultRowWithDefaults() *CronResultRow`

NewCronResultRowWithDefaults instantiates a new CronResultRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobFired

`func (o *CronResultRow) GetJobFired() bool`

GetJobFired returns the JobFired field if non-nil, zero value otherwise.

### GetJobFiredOk

`func (o *CronResultRow) GetJobFiredOk() (*bool, bool)`

GetJobFiredOk returns a tuple with the JobFired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFired

`func (o *CronResultRow) SetJobFired(v bool)`

SetJobFired sets JobFired field to given value.

### HasJobFired

`func (o *CronResultRow) HasJobFired() bool`

HasJobFired returns a boolean if a field has been set.

### SetJobFiredNil

`func (o *CronResultRow) SetJobFiredNil(b bool)`

 SetJobFiredNil sets the value for JobFired to be an explicit nil

### UnsetJobFired
`func (o *CronResultRow) UnsetJobFired()`

UnsetJobFired ensures that no value is present for JobFired, not even an explicit nil
### GetJobSucceeded

`func (o *CronResultRow) GetJobSucceeded() bool`

GetJobSucceeded returns the JobSucceeded field if non-nil, zero value otherwise.

### GetJobSucceededOk

`func (o *CronResultRow) GetJobSucceededOk() (*bool, bool)`

GetJobSucceededOk returns a tuple with the JobSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSucceeded

`func (o *CronResultRow) SetJobSucceeded(v bool)`

SetJobSucceeded sets JobSucceeded field to given value.

### HasJobSucceeded

`func (o *CronResultRow) HasJobSucceeded() bool`

HasJobSucceeded returns a boolean if a field has been set.

### SetJobSucceededNil

`func (o *CronResultRow) SetJobSucceededNil(b bool)`

 SetJobSucceededNil sets the value for JobSucceeded to be an explicit nil

### UnsetJobSucceeded
`func (o *CronResultRow) UnsetJobSucceeded()`

UnsetJobSucceeded ensures that no value is present for JobSucceeded, not even an explicit nil
### GetJobErrored

`func (o *CronResultRow) GetJobErrored() bool`

GetJobErrored returns the JobErrored field if non-nil, zero value otherwise.

### GetJobErroredOk

`func (o *CronResultRow) GetJobErroredOk() (*bool, bool)`

GetJobErroredOk returns a tuple with the JobErrored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobErrored

`func (o *CronResultRow) SetJobErrored(v bool)`

SetJobErrored sets JobErrored field to given value.

### HasJobErrored

`func (o *CronResultRow) HasJobErrored() bool`

HasJobErrored returns a boolean if a field has been set.

### SetJobErroredNil

`func (o *CronResultRow) SetJobErroredNil(b bool)`

 SetJobErroredNil sets the value for JobErrored to be an explicit nil

### UnsetJobErrored
`func (o *CronResultRow) UnsetJobErrored()`

UnsetJobErrored ensures that no value is present for JobErrored, not even an explicit nil
### GetMessage

`func (o *CronResultRow) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CronResultRow) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CronResultRow) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CronResultRow) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CronResultRow) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CronResultRow) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


