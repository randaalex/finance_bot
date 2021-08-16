# TransactionLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LinkTypeId** | **string** | The link type ID to use. You can also use the link_type_name field. | 
**LinkTypeName** | Pointer to **string** | The link type name to use. You can also use the link_type_id field. | [optional] 
**InwardId** | **string** | The inward transaction transaction_journal_id for the link. This becomes the &#39;is paid by&#39; transaction of the set. | 
**OutwardId** | **string** | The outward transaction transaction_journal_id for the link. This becomes the &#39;pays for&#39; transaction of the set. | 
**Notes** | Pointer to **NullableString** | Optional. Some notes. | [optional] 

## Methods

### NewTransactionLink

`func NewTransactionLink(linkTypeId string, inwardId string, outwardId string, ) *TransactionLink`

NewTransactionLink instantiates a new TransactionLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLinkWithDefaults

`func NewTransactionLinkWithDefaults() *TransactionLink`

NewTransactionLinkWithDefaults instantiates a new TransactionLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TransactionLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TransactionLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinkTypeId

`func (o *TransactionLink) GetLinkTypeId() string`

GetLinkTypeId returns the LinkTypeId field if non-nil, zero value otherwise.

### GetLinkTypeIdOk

`func (o *TransactionLink) GetLinkTypeIdOk() (*string, bool)`

GetLinkTypeIdOk returns a tuple with the LinkTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeId

`func (o *TransactionLink) SetLinkTypeId(v string)`

SetLinkTypeId sets LinkTypeId field to given value.


### GetLinkTypeName

`func (o *TransactionLink) GetLinkTypeName() string`

GetLinkTypeName returns the LinkTypeName field if non-nil, zero value otherwise.

### GetLinkTypeNameOk

`func (o *TransactionLink) GetLinkTypeNameOk() (*string, bool)`

GetLinkTypeNameOk returns a tuple with the LinkTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeName

`func (o *TransactionLink) SetLinkTypeName(v string)`

SetLinkTypeName sets LinkTypeName field to given value.

### HasLinkTypeName

`func (o *TransactionLink) HasLinkTypeName() bool`

HasLinkTypeName returns a boolean if a field has been set.

### GetInwardId

`func (o *TransactionLink) GetInwardId() string`

GetInwardId returns the InwardId field if non-nil, zero value otherwise.

### GetInwardIdOk

`func (o *TransactionLink) GetInwardIdOk() (*string, bool)`

GetInwardIdOk returns a tuple with the InwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInwardId

`func (o *TransactionLink) SetInwardId(v string)`

SetInwardId sets InwardId field to given value.


### GetOutwardId

`func (o *TransactionLink) GetOutwardId() string`

GetOutwardId returns the OutwardId field if non-nil, zero value otherwise.

### GetOutwardIdOk

`func (o *TransactionLink) GetOutwardIdOk() (*string, bool)`

GetOutwardIdOk returns a tuple with the OutwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutwardId

`func (o *TransactionLink) SetOutwardId(v string)`

SetOutwardId sets OutwardId field to given value.


### GetNotes

`func (o *TransactionLink) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionLink) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionLink) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionLink) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionLink) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionLink) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


