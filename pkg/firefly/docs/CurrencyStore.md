# CurrencyStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Defaults to true | [optional] 
**Default** | Pointer to **bool** | Make this currency the default currency. | [optional] 
**Code** | **string** |  | 
**Name** | **string** |  | 
**Symbol** | **string** |  | 
**DecimalPlaces** | Pointer to **int32** | Supports 0-16 decimals. | [optional] 

## Methods

### NewCurrencyStore

`func NewCurrencyStore(code string, name string, symbol string, ) *CurrencyStore`

NewCurrencyStore instantiates a new CurrencyStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyStoreWithDefaults

`func NewCurrencyStoreWithDefaults() *CurrencyStore`

NewCurrencyStoreWithDefaults instantiates a new CurrencyStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CurrencyStore) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CurrencyStore) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CurrencyStore) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CurrencyStore) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefault

`func (o *CurrencyStore) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CurrencyStore) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CurrencyStore) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CurrencyStore) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetCode

`func (o *CurrencyStore) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyStore) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyStore) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *CurrencyStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrencyStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrencyStore) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *CurrencyStore) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyStore) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyStore) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimalPlaces

`func (o *CurrencyStore) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *CurrencyStore) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *CurrencyStore) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *CurrencyStore) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


