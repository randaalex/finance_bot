# BudgetLimitRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**BudgetLimit**](BudgetLimit.md) |  | 

## Methods

### NewBudgetLimitRead

`func NewBudgetLimitRead(type_ string, id string, attributes BudgetLimit, ) *BudgetLimitRead`

NewBudgetLimitRead instantiates a new BudgetLimitRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetLimitReadWithDefaults

`func NewBudgetLimitReadWithDefaults() *BudgetLimitRead`

NewBudgetLimitReadWithDefaults instantiates a new BudgetLimitRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BudgetLimitRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BudgetLimitRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BudgetLimitRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BudgetLimitRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetLimitRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetLimitRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BudgetLimitRead) GetAttributes() BudgetLimit`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BudgetLimitRead) GetAttributesOk() (*BudgetLimit, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BudgetLimitRead) SetAttributes(v BudgetLimit)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


