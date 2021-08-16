# TransactionLinkStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkTypeId** | **string** | The link type ID to use. You can also use the link_type_name field. | 
**LinkTypeName** | Pointer to **string** | The link type name to use. You can also use the link_type_id field. | [optional] 
**InwardId** | **string** | The inward transaction transaction_journal_id for the link. This becomes the &#39;is paid by&#39; transaction of the set. | 
**OutwardId** | **string** | The outward transaction transaction_journal_id for the link. This becomes the &#39;pays for&#39; transaction of the set. | 
**Notes** | Pointer to **NullableString** | Optional. Some notes. | [optional] 

## Methods

### NewTransactionLinkStore

`func NewTransactionLinkStore(linkTypeId string, inwardId string, outwardId string, ) *TransactionLinkStore`

NewTransactionLinkStore instantiates a new TransactionLinkStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLinkStoreWithDefaults

`func NewTransactionLinkStoreWithDefaults() *TransactionLinkStore`

NewTransactionLinkStoreWithDefaults instantiates a new TransactionLinkStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkTypeId

`func (o *TransactionLinkStore) GetLinkTypeId() string`

GetLinkTypeId returns the LinkTypeId field if non-nil, zero value otherwise.

### GetLinkTypeIdOk

`func (o *TransactionLinkStore) GetLinkTypeIdOk() (*string, bool)`

GetLinkTypeIdOk returns a tuple with the LinkTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeId

`func (o *TransactionLinkStore) SetLinkTypeId(v string)`

SetLinkTypeId sets LinkTypeId field to given value.


### GetLinkTypeName

`func (o *TransactionLinkStore) GetLinkTypeName() string`

GetLinkTypeName returns the LinkTypeName field if non-nil, zero value otherwise.

### GetLinkTypeNameOk

`func (o *TransactionLinkStore) GetLinkTypeNameOk() (*string, bool)`

GetLinkTypeNameOk returns a tuple with the LinkTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTypeName

`func (o *TransactionLinkStore) SetLinkTypeName(v string)`

SetLinkTypeName sets LinkTypeName field to given value.

### HasLinkTypeName

`func (o *TransactionLinkStore) HasLinkTypeName() bool`

HasLinkTypeName returns a boolean if a field has been set.

### GetInwardId

`func (o *TransactionLinkStore) GetInwardId() string`

GetInwardId returns the InwardId field if non-nil, zero value otherwise.

### GetInwardIdOk

`func (o *TransactionLinkStore) GetInwardIdOk() (*string, bool)`

GetInwardIdOk returns a tuple with the InwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInwardId

`func (o *TransactionLinkStore) SetInwardId(v string)`

SetInwardId sets InwardId field to given value.


### GetOutwardId

`func (o *TransactionLinkStore) GetOutwardId() string`

GetOutwardId returns the OutwardId field if non-nil, zero value otherwise.

### GetOutwardIdOk

`func (o *TransactionLinkStore) GetOutwardIdOk() (*string, bool)`

GetOutwardIdOk returns a tuple with the OutwardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutwardId

`func (o *TransactionLinkStore) SetOutwardId(v string)`

SetOutwardId sets OutwardId field to given value.


### GetNotes

`func (o *TransactionLinkStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TransactionLinkStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TransactionLinkStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TransactionLinkStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TransactionLinkStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TransactionLinkStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


