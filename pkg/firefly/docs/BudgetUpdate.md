# BudgetUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**AutoBudgetType** | Pointer to **NullableString** | The type of auto-budget that Firefly III must create. | [optional] 
**AutoBudgetCurrencyId** | Pointer to **NullableString** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**AutoBudgetCurrencyCode** | Pointer to **NullableString** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**AutoBudgetAmount** | Pointer to **NullableString** |  | [optional] 
**AutoBudgetPeriod** | Pointer to **NullableString** | Period for the auto budget | [optional] 

## Methods

### NewBudgetUpdate

`func NewBudgetUpdate() *BudgetUpdate`

NewBudgetUpdate instantiates a new BudgetUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetUpdateWithDefaults

`func NewBudgetUpdateWithDefaults() *BudgetUpdate`

NewBudgetUpdateWithDefaults instantiates a new BudgetUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BudgetUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BudgetUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BudgetUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BudgetUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *BudgetUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BudgetUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BudgetUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BudgetUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *BudgetUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *BudgetUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *BudgetUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *BudgetUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAutoBudgetType

`func (o *BudgetUpdate) GetAutoBudgetType() string`

GetAutoBudgetType returns the AutoBudgetType field if non-nil, zero value otherwise.

### GetAutoBudgetTypeOk

`func (o *BudgetUpdate) GetAutoBudgetTypeOk() (*string, bool)`

GetAutoBudgetTypeOk returns a tuple with the AutoBudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetType

`func (o *BudgetUpdate) SetAutoBudgetType(v string)`

SetAutoBudgetType sets AutoBudgetType field to given value.

### HasAutoBudgetType

`func (o *BudgetUpdate) HasAutoBudgetType() bool`

HasAutoBudgetType returns a boolean if a field has been set.

### SetAutoBudgetTypeNil

`func (o *BudgetUpdate) SetAutoBudgetTypeNil(b bool)`

 SetAutoBudgetTypeNil sets the value for AutoBudgetType to be an explicit nil

### UnsetAutoBudgetType
`func (o *BudgetUpdate) UnsetAutoBudgetType()`

UnsetAutoBudgetType ensures that no value is present for AutoBudgetType, not even an explicit nil
### GetAutoBudgetCurrencyId

`func (o *BudgetUpdate) GetAutoBudgetCurrencyId() string`

GetAutoBudgetCurrencyId returns the AutoBudgetCurrencyId field if non-nil, zero value otherwise.

### GetAutoBudgetCurrencyIdOk

`func (o *BudgetUpdate) GetAutoBudgetCurrencyIdOk() (*string, bool)`

GetAutoBudgetCurrencyIdOk returns a tuple with the AutoBudgetCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetCurrencyId

`func (o *BudgetUpdate) SetAutoBudgetCurrencyId(v string)`

SetAutoBudgetCurrencyId sets AutoBudgetCurrencyId field to given value.

### HasAutoBudgetCurrencyId

`func (o *BudgetUpdate) HasAutoBudgetCurrencyId() bool`

HasAutoBudgetCurrencyId returns a boolean if a field has been set.

### SetAutoBudgetCurrencyIdNil

`func (o *BudgetUpdate) SetAutoBudgetCurrencyIdNil(b bool)`

 SetAutoBudgetCurrencyIdNil sets the value for AutoBudgetCurrencyId to be an explicit nil

### UnsetAutoBudgetCurrencyId
`func (o *BudgetUpdate) UnsetAutoBudgetCurrencyId()`

UnsetAutoBudgetCurrencyId ensures that no value is present for AutoBudgetCurrencyId, not even an explicit nil
### GetAutoBudgetCurrencyCode

`func (o *BudgetUpdate) GetAutoBudgetCurrencyCode() string`

GetAutoBudgetCurrencyCode returns the AutoBudgetCurrencyCode field if non-nil, zero value otherwise.

### GetAutoBudgetCurrencyCodeOk

`func (o *BudgetUpdate) GetAutoBudgetCurrencyCodeOk() (*string, bool)`

GetAutoBudgetCurrencyCodeOk returns a tuple with the AutoBudgetCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetCurrencyCode

`func (o *BudgetUpdate) SetAutoBudgetCurrencyCode(v string)`

SetAutoBudgetCurrencyCode sets AutoBudgetCurrencyCode field to given value.

### HasAutoBudgetCurrencyCode

`func (o *BudgetUpdate) HasAutoBudgetCurrencyCode() bool`

HasAutoBudgetCurrencyCode returns a boolean if a field has been set.

### SetAutoBudgetCurrencyCodeNil

`func (o *BudgetUpdate) SetAutoBudgetCurrencyCodeNil(b bool)`

 SetAutoBudgetCurrencyCodeNil sets the value for AutoBudgetCurrencyCode to be an explicit nil

### UnsetAutoBudgetCurrencyCode
`func (o *BudgetUpdate) UnsetAutoBudgetCurrencyCode()`

UnsetAutoBudgetCurrencyCode ensures that no value is present for AutoBudgetCurrencyCode, not even an explicit nil
### GetAutoBudgetAmount

`func (o *BudgetUpdate) GetAutoBudgetAmount() string`

GetAutoBudgetAmount returns the AutoBudgetAmount field if non-nil, zero value otherwise.

### GetAutoBudgetAmountOk

`func (o *BudgetUpdate) GetAutoBudgetAmountOk() (*string, bool)`

GetAutoBudgetAmountOk returns a tuple with the AutoBudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetAmount

`func (o *BudgetUpdate) SetAutoBudgetAmount(v string)`

SetAutoBudgetAmount sets AutoBudgetAmount field to given value.

### HasAutoBudgetAmount

`func (o *BudgetUpdate) HasAutoBudgetAmount() bool`

HasAutoBudgetAmount returns a boolean if a field has been set.

### SetAutoBudgetAmountNil

`func (o *BudgetUpdate) SetAutoBudgetAmountNil(b bool)`

 SetAutoBudgetAmountNil sets the value for AutoBudgetAmount to be an explicit nil

### UnsetAutoBudgetAmount
`func (o *BudgetUpdate) UnsetAutoBudgetAmount()`

UnsetAutoBudgetAmount ensures that no value is present for AutoBudgetAmount, not even an explicit nil
### GetAutoBudgetPeriod

`func (o *BudgetUpdate) GetAutoBudgetPeriod() string`

GetAutoBudgetPeriod returns the AutoBudgetPeriod field if non-nil, zero value otherwise.

### GetAutoBudgetPeriodOk

`func (o *BudgetUpdate) GetAutoBudgetPeriodOk() (*string, bool)`

GetAutoBudgetPeriodOk returns a tuple with the AutoBudgetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBudgetPeriod

`func (o *BudgetUpdate) SetAutoBudgetPeriod(v string)`

SetAutoBudgetPeriod sets AutoBudgetPeriod field to given value.

### HasAutoBudgetPeriod

`func (o *BudgetUpdate) HasAutoBudgetPeriod() bool`

HasAutoBudgetPeriod returns a boolean if a field has been set.

### SetAutoBudgetPeriodNil

`func (o *BudgetUpdate) SetAutoBudgetPeriodNil(b bool)`

 SetAutoBudgetPeriodNil sets the value for AutoBudgetPeriod to be an explicit nil

### UnsetAutoBudgetPeriod
`func (o *BudgetUpdate) UnsetAutoBudgetPeriod()`

UnsetAutoBudgetPeriod ensures that no value is present for AutoBudgetPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


