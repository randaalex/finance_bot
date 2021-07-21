# RecurrenceRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Recurrence**](Recurrence.md) |  | 
**Links** | [**ObjectLink**](ObjectLink.md) |  | 

## Methods

### NewRecurrenceRead

`func NewRecurrenceRead(type_ string, id string, attributes Recurrence, links ObjectLink, ) *RecurrenceRead`

NewRecurrenceRead instantiates a new RecurrenceRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceReadWithDefaults

`func NewRecurrenceReadWithDefaults() *RecurrenceRead`

NewRecurrenceReadWithDefaults instantiates a new RecurrenceRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecurrenceRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecurrenceRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecurrenceRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RecurrenceRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurrenceRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurrenceRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *RecurrenceRead) GetAttributes() Recurrence`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RecurrenceRead) GetAttributesOk() (*Recurrence, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RecurrenceRead) SetAttributes(v Recurrence)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *RecurrenceRead) GetLinks() ObjectLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RecurrenceRead) GetLinksOk() (*ObjectLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RecurrenceRead) SetLinks(v ObjectLink)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


