# TransactionLinkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkTypeId** | Pointer to **string** | The link type ID to use. Use this field OR use the link_type_name field. | [optional] 
**LinkTypeName** | Pointer to **string** | The link type name to use. Use this field OR use the link_type_id field. | [optional] 
**InwardId** | Pointer to **string** | The inward transaction transaction_journal_id for the link. This becomes the &#39;is paid by&#39; transaction of the set. | [optional] 
**OutwardId** | Pointer to **string** | The outward transaction transaction_journal_id for the link. This becomes the &#39;pays for&#39; transaction of the set. | [optional] 
**Notes** | Pointer to **NullableString** | Optional. Some notes. If you submit an empty string the current notes will be removed | [optional] 

## Methods

### NewTransactionLinkUpdate

`func NewTransactionLinkUpdate() *TransactionLinkUpdate`

NewTransactionLinkUpdate instantiates a new TransactionLinkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLinkUpdateWithDefaults

`func NewTransactionLinkUpdateWithDefaults() *TransactionLinkUpdate`

NewTransactionLinkUpdateWithDefaults instantiates a new TransactionLinkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkTypeId

`func (o *TransactionLinkUpdate) GetLinkTypeId() string`

GetLinkTypeId returns the LinkTypeId field if non-nil, zero value otherwise.

### GetLinkTypeIdOk

`func (o *TransactionLinkUpdate) GetLinkTypeIdOk() (*string, bool)`

GetLinkTypeIdOk returns a tuple with the LinkTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeId

`func (o *TransactionLinkUpdate) SetLinkTypeId(v string)`

SetLinkTypeId sets LinkTypeId field to given value.

### HasLinkTypeId

`func (o *TransactionLinkUpdate) HasLinkTypeId() bool`

HasLinkTypeId returns a boolean if a field has been set.

### GetLinkTypeName

`func (o *TransactionLinkUpdate) GetLinkTypeName() string`

GetLinkTypeName returns the LinkTypeName field if non-nil, zero value otherwise.

### GetLinkTypeNameOk

`func (o *TransactionLinkUpdate) GetLinkTypeNameOk() (*string, bool)`

GetLinkTypeNameOk returns a tuple with the LinkTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeName

`func (o *TransactionLinkUpdate) SetLinkTypeName(v string)`

SetLinkTypeName sets LinkTypeName field to given value.

### HasLinkTypeName

`func (o *TransactionLinkUpdate) HasLinkTypeName() bool`

HasLinkTypeName returns a boolean if a field has been set.

### GetInwardId

`func (o *TransactionLinkUpdate) GetInwardId() string`

GetInwardId returns the InwardId field if non-nil, zero value otherwise.

### GetInwardIdOk

`func (o *TransactionLinkUpdate) GetInwardIdOk() (*string, bool)`

GetInwardIdOk returns a tuple with the InwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInwardId

`func (o *TransactionLinkUpdate) SetInwardId(v string)`

SetInwardId sets InwardId field to given value.

### HasInwardId

`func (o *TransactionLinkUpdate) HasInwardId() bool`

HasInwardId returns a boolean if a field has been set.

### GetOutwardId

`func (o *TransactionLinkUpdate) GetOutwardId() string`

GetOutwardId returns the OutwardId field if non-nil, zero value otherwise.

### GetOutwardIdOk

`func (o *TransactionLinkUpdate) GetOutwardIdOk() (*string, bool)`

GetOutwardIdOk returns a tuple with the OutwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutwardId

`func (o *TransactionLinkUpdate) SetOutwardId(v string)`

SetOutwardId sets OutwardId field to given value.

### HasOutwardId

`func (o *TransactionLinkUpdate) HasOutwardId() bool`

HasOutwardId returns a boolean if a field has been set.

### GetNotes

`func (o *TransactionLinkUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionLinkUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionLinkUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionLinkUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionLinkUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionLinkUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


