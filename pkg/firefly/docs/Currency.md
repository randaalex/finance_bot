# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Defaults to true | [optional] 
**Default** | Pointer to **bool** | Make this currency the default currency. | [optional] 
**Code** | **string** |  | 
**Name** | **string** |  | 
**Symbol** | **string** |  | 
**DecimalPlaces** | Pointer to **int32** | Supports 0-16 decimals. | [optional] 

## Methods

### NewCurrency

`func NewCurrency(code string, name string, symbol string, ) *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Currency) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Currency) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Currency) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Currency) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Currency) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Currency) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Currency) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Currency) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *Currency) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Currency) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Currency) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Currency) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefault

`func (o *Currency) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Currency) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Currency) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Currency) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCode

`func (o *Currency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Currency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Currency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Currency) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *Currency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Currency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Currency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimalPlaces

`func (o *Currency) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *Currency) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *Currency) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *Currency) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


