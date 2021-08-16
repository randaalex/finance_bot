# AttachmentStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**AttachableType** | **string** | The object class to which the attachment must be linked. | 
**AttachableId** | **string** | ID of the model this attachment is linked to. | 
**Title** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAttachmentStore

`func NewAttachmentStore(filename string, attachableType string, attachableId string, ) *AttachmentStore`

NewAttachmentStore instantiates a new AttachmentStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentStoreWithDefaults

`func NewAttachmentStoreWithDefaults() *AttachmentStore`

NewAttachmentStoreWithDefaults instantiates a new AttachmentStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *AttachmentStore) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AttachmentStore) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AttachmentStore) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetAttachableType

`func (o *AttachmentStore) GetAttachableType() string`

GetAttachableType returns the AttachableType field if non-nil, zero value otherwise.

### GetAttachableTypeOk

`func (o *AttachmentStore) GetAttachableTypeOk() (*string, bool)`

GetAttachableTypeOk returns a tuple with the AttachableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableType

`func (o *AttachmentStore) SetAttachableType(v string)`

SetAttachableType sets AttachableType field to given value.


### GetAttachableId

`func (o *AttachmentStore) GetAttachableId() string`

GetAttachableId returns the AttachableId field if non-nil, zero value otherwise.

### GetAttachableIdOk

`func (o *AttachmentStore) GetAttachableIdOk() (*string, bool)`

GetAttachableIdOk returns a tuple with the AttachableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachableId

`func (o *AttachmentStore) SetAttachableId(v string)`

SetAttachableId sets AttachableId field to given value.


### GetTitle

`func (o *AttachmentStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttachmentStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttachmentStore) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttachmentStore) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetNotes

`func (o *AttachmentStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AttachmentStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AttachmentStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AttachmentStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AttachmentStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AttachmentStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


