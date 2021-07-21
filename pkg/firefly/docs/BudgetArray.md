# BudgetArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BudgetRead**](BudgetRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewBudgetArray

`func NewBudgetArray(data []BudgetRead, meta Meta, ) *BudgetArray`

NewBudgetArray instantiates a new BudgetArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetArrayWithDefaults

`func NewBudgetArrayWithDefaults() *BudgetArray`

NewBudgetArrayWithDefaults instantiates a new BudgetArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BudgetArray) GetData() []BudgetRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BudgetArray) GetDataOk() (*[]BudgetRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BudgetArray) SetData(v []BudgetRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *BudgetArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BudgetArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BudgetArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


