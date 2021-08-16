# CurrencyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If the currency is enabled | [optional] 
**Default** | Pointer to **bool** | If the currency must be the default for the user. You can only submit TRUE. | [optional] 
**Code** | Pointer to **string** | The currency code | [optional] 
**Name** | Pointer to **string** | The currency name | [optional] 
**Symbol** | Pointer to **string** | The currency symbol | [optional] 
**DecimalPlaces** | Pointer to **int32** | How many decimals to use when displaying this currency. Between 0 and 16. | [optional] 

## Methods

### NewCurrencyUpdate

`func NewCurrencyUpdate() *CurrencyUpdate`

NewCurrencyUpdate instantiates a new CurrencyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyUpdateWithDefaults

`func NewCurrencyUpdateWithDefaults() *CurrencyUpdate`

NewCurrencyUpdateWithDefaults instantiates a new CurrencyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CurrencyUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CurrencyUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CurrencyUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CurrencyUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefault

`func (o *CurrencyUpdate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CurrencyUpdate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CurrencyUpdate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CurrencyUpdate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCode

`func (o *CurrencyUpdate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyUpdate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyUpdate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CurrencyUpdate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CurrencyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrencyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrencyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrencyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *CurrencyUpdate) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyUpdate) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyUpdate) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrencyUpdate) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *CurrencyUpdate) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *CurrencyUpdate) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *CurrencyUpdate) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *CurrencyUpdate) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


