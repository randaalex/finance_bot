# BudgetLimitArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BudgetLimitRead**](BudgetLimitRead.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewBudgetLimitArray

`func NewBudgetLimitArray(data []BudgetLimitRead, meta Meta, ) *BudgetLimitArray`

NewBudgetLimitArray instantiates a new BudgetLimitArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetLimitArrayWithDefaults

`func NewBudgetLimitArrayWithDefaults() *BudgetLimitArray`

NewBudgetLimitArrayWithDefaults instantiates a new BudgetLimitArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BudgetLimitArray) GetData() []BudgetLimitRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BudgetLimitArray) GetDataOk() (*[]BudgetLimitRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BudgetLimitArray) SetData(v []BudgetLimitRead)`

SetData sets Data field to given value.


### GetMeta

`func (o *BudgetLimitArray) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BudgetLimitArray) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BudgetLimitArray) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


