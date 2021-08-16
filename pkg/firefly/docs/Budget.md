# Budget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] [readonly] 
**AutoBudgetType** | Pointer to **NullableString** | The type of auto-budget that Firefly III must create. | [optional] 
**AutoBudgetCurrencyId** | Pointer to **NullableString** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**AutoBudgetCurrencyCode** | Pointer to **NullableString** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**AutoBudgetAmount** | Pointer to **NullableString** |  | [optional] 
**AutoBudgetPeriod** | Pointer to **NullableString** | Period for the auto budget | [optional] 
**Spent** | Pointer to [**[]BudgetSpent**](BudgetSpent.md) | Information on how much was spent in this budget. Is only filled in when the start and end date are submitted. | [optional] [readonly] 

## Methods

### NewBudget

`func NewBudget(name string, ) *Budget`

NewBudget instantiates a new Budget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetWithDefaults

`func NewBudgetWithDefaults() *Budget`

NewBudgetWithDefaults instantiates a new Budget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Budget) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Budget) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Budget) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Budget) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Budget) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Budget) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Budget) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Budget) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Budget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Budget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Budget) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *Budget) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Budget) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Budget) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Budget) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *Budget) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Budget) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Budget) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Budget) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAutoBudgetType

`func (o *Budget) GetAutoBudgetType() string`

GetAutoBudgetType returns the AutoBudgetType field if non-nil, zero value otherwise.

### GetAutoBudgetTypeOk

`func (o *Budget) GetAutoBudgetTypeOk() (*string, bool)`

GetAutoBudgetTypeOk returns a tuple with the AutoBudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetType

`func (o *Budget) SetAutoBudgetType(v string)`

SetAutoBudgetType sets AutoBudgetType field to given value.

### HasAutoBudgetType

`func (o *Budget) HasAutoBudgetType() bool`

HasAutoBudgetType returns a boolean if a field has been set.

### SetAutoBudgetTypeNil

`func (o *Budget) SetAutoBudgetTypeNil(b bool)`

 SetAutoBudgetTypeNil sets the value for AutoBudgetType to be an explicit nil

### UnsetAutoBudgetType
`func (o *Budget) UnsetAutoBudgetType()`

UnsetAutoBudgetType ensures that no value is present for AutoBudgetType, not even an explicit nil
### GetAutoBudgetCurrencyId

`func (o *Budget) GetAutoBudgetCurrencyId() string`

GetAutoBudgetCurrencyId returns the AutoBudgetCurrencyId field if non-nil, zero value otherwise.

### GetAutoBudgetCurrencyIdOk

`func (o *Budget) GetAutoBudgetCurrencyIdOk() (*string, bool)`

GetAutoBudgetCurrencyIdOk returns a tuple with the AutoBudgetCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetCurrencyId

`func (o *Budget) SetAutoBudgetCurrencyId(v string)`

SetAutoBudgetCurrencyId sets AutoBudgetCurrencyId field to given value.

### HasAutoBudgetCurrencyId

`func (o *Budget) HasAutoBudgetCurrencyId() bool`

HasAutoBudgetCurrencyId returns a boolean if a field has been set.

### SetAutoBudgetCurrencyIdNil

`func (o *Budget) SetAutoBudgetCurrencyIdNil(b bool)`

 SetAutoBudgetCurrencyIdNil sets the value for AutoBudgetCurrencyId to be an explicit nil

### UnsetAutoBudgetCurrencyId
`func (o *Budget) UnsetAutoBudgetCurrencyId()`

UnsetAutoBudgetCurrencyId ensures that no value is present for AutoBudgetCurrencyId, not even an explicit nil
### GetAutoBudgetCurrencyCode

`func (o *Budget) GetAutoBudgetCurrencyCode() string`

GetAutoBudgetCurrencyCode returns the AutoBudgetCurrencyCode field if non-nil, zero value otherwise.

### GetAutoBudgetCurrencyCodeOk

`func (o *Budget) GetAutoBudgetCurrencyCodeOk() (*string, bool)`

GetAutoBudgetCurrencyCodeOk returns a tuple with the AutoBudgetCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetCurrencyCode

`func (o *Budget) SetAutoBudgetCurrencyCode(v string)`

SetAutoBudgetCurrencyCode sets AutoBudgetCurrencyCode field to given value.

### HasAutoBudgetCurrencyCode

`func (o *Budget) HasAutoBudgetCurrencyCode() bool`

HasAutoBudgetCurrencyCode returns a boolean if a field has been set.

### SetAutoBudgetCurrencyCodeNil

`func (o *Budget) SetAutoBudgetCurrencyCodeNil(b bool)`

 SetAutoBudgetCurrencyCodeNil sets the value for AutoBudgetCurrencyCode to be an explicit nil

### UnsetAutoBudgetCurrencyCode
`func (o *Budget) UnsetAutoBudgetCurrencyCode()`

UnsetAutoBudgetCurrencyCode ensures that no value is present for AutoBudgetCurrencyCode, not even an explicit nil
### GetAutoBudgetAmount

`func (o *Budget) GetAutoBudgetAmount() string`

GetAutoBudgetAmount returns the AutoBudgetAmount field if non-nil, zero value otherwise.

### GetAutoBudgetAmountOk

`func (o *Budget) GetAutoBudgetAmountOk() (*string, bool)`

GetAutoBudgetAmountOk returns a tuple with the AutoBudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetAmount

`func (o *Budget) SetAutoBudgetAmount(v string)`

SetAutoBudgetAmount sets AutoBudgetAmount field to given value.

### HasAutoBudgetAmount

`func (o *Budget) HasAutoBudgetAmount() bool`

HasAutoBudgetAmount returns a boolean if a field has been set.

### SetAutoBudgetAmountNil

`func (o *Budget) SetAutoBudgetAmountNil(b bool)`

 SetAutoBudgetAmountNil sets the value for AutoBudgetAmount to be an explicit nil

### UnsetAutoBudgetAmount
`func (o *Budget) UnsetAutoBudgetAmount()`

UnsetAutoBudgetAmount ensures that no value is present for AutoBudgetAmount, not even an explicit nil
### GetAutoBudgetPeriod

`func (o *Budget) GetAutoBudgetPeriod() string`

GetAutoBudgetPeriod returns the AutoBudgetPeriod field if non-nil, zero value otherwise.

### GetAutoBudgetPeriodOk

`func (o *Budget) GetAutoBudgetPeriodOk() (*string, bool)`

GetAutoBudgetPeriodOk returns a tuple with the AutoBudgetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetPeriod

`func (o *Budget) SetAutoBudgetPeriod(v string)`

SetAutoBudgetPeriod sets AutoBudgetPeriod field to given value.

### HasAutoBudgetPeriod

`func (o *Budget) HasAutoBudgetPeriod() bool`

HasAutoBudgetPeriod returns a boolean if a field has been set.

### SetAutoBudgetPeriodNil

`func (o *Budget) SetAutoBudgetPeriodNil(b bool)`

 SetAutoBudgetPeriodNil sets the value for AutoBudgetPeriod to be an explicit nil

### UnsetAutoBudgetPeriod
`func (o *Budget) UnsetAutoBudgetPeriod()`

UnsetAutoBudgetPeriod ensures that no value is present for AutoBudgetPeriod, not even an explicit nil
### GetSpent

`func (o *Budget) GetSpent() []BudgetSpent`

GetSpent returns the Spent field if non-nil, zero value otherwise.

### GetSpentOk

`func (o *Budget) GetSpentOk() (*[]BudgetSpent, bool)`

GetSpentOk returns a tuple with the Spent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpent

`func (o *Budget) SetSpent(v []BudgetSpent)`

SetSpent sets Spent field to given value.

### HasSpent

`func (o *Budget) HasSpent() bool`

HasSpent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


