# AutocompletePiggyBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** | Name of the piggy bank found by an auto-complete search. | [readonly] 
**NameWithBalance** | Pointer to **string** | Name of the piggy bank found by an auto-complete search with the current balance formatted nicely. | [optional] [readonly] 
**CurrencyId** | Pointer to **int32** | Currency ID for this piggy bank. | [optional] [readonly] 
**CurrencyCode** | Pointer to **string** | Currency code for this piggy bank. | [optional] [readonly] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewAutocompletePiggyBalance

`func NewAutocompletePiggyBalance(id string, name string, ) *AutocompletePiggyBalance`

NewAutocompletePiggyBalance instantiates a new AutocompletePiggyBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompletePiggyBalanceWithDefaults

`func NewAutocompletePiggyBalanceWithDefaults() *AutocompletePiggyBalance`

NewAutocompletePiggyBalanceWithDefaults instantiates a new AutocompletePiggyBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompletePiggyBalance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompletePiggyBalance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompletePiggyBalance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutocompletePiggyBalance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompletePiggyBalance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompletePiggyBalance) SetName(v string)`

SetName sets Name field to given value.


### GetNameWithBalance

`func (o *AutocompletePiggyBalance) GetNameWithBalance() string`

GetNameWithBalance returns the NameWithBalance field if non-nil, zero value otherwise.

### GetNameWithBalanceOk

`func (o *AutocompletePiggyBalance) GetNameWithBalanceOk() (*string, bool)`

GetNameWithBalanceOk returns a tuple with the NameWithBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameWithBalance

`func (o *AutocompletePiggyBalance) SetNameWithBalance(v string)`

SetNameWithBalance sets NameWithBalance field to given value.

### HasNameWithBalance

`func (o *AutocompletePiggyBalance) HasNameWithBalance() bool`

HasNameWithBalance returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AutocompletePiggyBalance) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AutocompletePiggyBalance) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AutocompletePiggyBalance) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AutocompletePiggyBalance) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AutocompletePiggyBalance) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutocompletePiggyBalance) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutocompletePiggyBalance) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AutocompletePiggyBalance) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *AutocompletePiggyBalance) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *AutocompletePiggyBalance) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *AutocompletePiggyBalance) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *AutocompletePiggyBalance) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *AutocompletePiggyBalance) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *AutocompletePiggyBalance) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *AutocompletePiggyBalance) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *AutocompletePiggyBalance) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


