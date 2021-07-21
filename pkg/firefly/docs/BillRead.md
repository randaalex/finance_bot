# BillRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Bill**](Bill.md) |  | 

## Methods

### NewBillRead

`func NewBillRead(type_ string, id string, attributes Bill, ) *BillRead`

NewBillRead instantiates a new BillRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillReadWithDefaults

`func NewBillReadWithDefaults() *BillRead`

NewBillReadWithDefaults instantiates a new BillRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BillRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BillRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BillRead) GetAttributes() Bill`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillRead) GetAttributesOk() (*Bill, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillRead) SetAttributes(v Bill)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


