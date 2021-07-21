# BudgetRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Budget**](Budget.md) |  | 

## Methods

### NewBudgetRead

`func NewBudgetRead(type_ string, id string, attributes Budget, ) *BudgetRead`

NewBudgetRead instantiates a new BudgetRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetReadWithDefaults

`func NewBudgetReadWithDefaults() *BudgetRead`

NewBudgetReadWithDefaults instantiates a new BudgetRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BudgetRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BudgetRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BudgetRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BudgetRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BudgetRead) GetAttributes() Budget`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BudgetRead) GetAttributesOk() (*Budget, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BudgetRead) SetAttributes(v Budget)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


