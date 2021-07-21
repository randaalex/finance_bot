# ImportJobAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TagId** | Pointer to **int32** | ID of the tag related to the import job, if present. | [optional] 
**TagTag** | Pointer to **string** | Tag related to the import job, if present. | [optional] 
**Key** | Pointer to **string** | Import job unique identifier. | [optional] 
**FileType** | Pointer to **string** | File type, if relevant. | [optional] 
**Provider** | Pointer to **string** | Import provider that did the import. | [optional] 
**Status** | Pointer to **string** | Status of import job. | [optional] 
**Stage** | Pointer to **string** | Current stage. | [optional] 
**Configuration** | Pointer to **string** | JSON string with job-specific configuration. | [optional] 
**ExtendedStatus** | Pointer to **string** | JSON string with job-specific status. | [optional] 
**Transactions** | Pointer to **string** | JSON string with a count of transactions in the job. | [optional] 
**Errors** | Pointer to **string** | JSON string with a list of errors. | [optional] 

## Methods

### NewImportJobAttributes

`func NewImportJobAttributes() *ImportJobAttributes`

NewImportJobAttributes instantiates a new ImportJobAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportJobAttributesWithDefaults

`func NewImportJobAttributesWithDefaults() *ImportJobAttributes`

NewImportJobAttributesWithDefaults instantiates a new ImportJobAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ImportJobAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImportJobAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImportJobAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImportJobAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ImportJobAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ImportJobAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ImportJobAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ImportJobAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTagId

`func (o *ImportJobAttributes) GetTagId() int32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *ImportJobAttributes) GetTagIdOk() (*int32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *ImportJobAttributes) SetTagId(v int32)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *ImportJobAttributes) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetTagTag

`func (o *ImportJobAttributes) GetTagTag() string`

GetTagTag returns the TagTag field if non-nil, zero value otherwise.

### GetTagTagOk

`func (o *ImportJobAttributes) GetTagTagOk() (*string, bool)`

GetTagTagOk returns a tuple with the TagTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTag

`func (o *ImportJobAttributes) SetTagTag(v string)`

SetTagTag sets TagTag field to given value.

### HasTagTag

`func (o *ImportJobAttributes) HasTagTag() bool`

HasTagTag returns a boolean if a field has been set.

### GetKey

`func (o *ImportJobAttributes) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ImportJobAttributes) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ImportJobAttributes) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ImportJobAttributes) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFileType

`func (o *ImportJobAttributes) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ImportJobAttributes) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ImportJobAttributes) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *ImportJobAttributes) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetProvider

`func (o *ImportJobAttributes) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ImportJobAttributes) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ImportJobAttributes) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ImportJobAttributes) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ImportJobAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportJobAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportJobAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportJobAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStage

`func (o *ImportJobAttributes) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ImportJobAttributes) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ImportJobAttributes) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ImportJobAttributes) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetConfiguration

`func (o *ImportJobAttributes) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ImportJobAttributes) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ImportJobAttributes) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ImportJobAttributes) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *ImportJobAttributes) GetExtendedStatus() string`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *ImportJobAttributes) GetExtendedStatusOk() (*string, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *ImportJobAttributes) SetExtendedStatus(v string)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *ImportJobAttributes) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetTransactions

`func (o *ImportJobAttributes) GetTransactions() string`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ImportJobAttributes) GetTransactionsOk() (*string, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ImportJobAttributes) SetTransactions(v string)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ImportJobAttributes) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetErrors

`func (o *ImportJobAttributes) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImportJobAttributes) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImportJobAttributes) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ImportJobAttributes) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


