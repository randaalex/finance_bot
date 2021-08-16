# AutocompleteAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of the account found by an auto-complete search. | 
**NameWithBalance** | **string** | Asset accounts and liabilities have a second field with the given date&#39;s account balance. | 
**Type** | **string** | Account type of the account found by the auto-complete search. | 
**CurrencyId** | **string** | Currency ID for this account. | 
**CurrencyCode** | **string** | Currency code for this account. | 
**CurrencySymbol** | **string** |  | 
**CurrencyDecimalPlaces** | **int32** |  | 

## Methods

### NewAutocompleteAccount

`func NewAutocompleteAccount(id string, name string, nameWithBalance string, type_ string, currencyId string, currencyCode string, currencySymbol string, currencyDecimalPlaces int32, ) *AutocompleteAccount`

NewAutocompleteAccount instantiates a new AutocompleteAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteAccountWithDefaults

`func NewAutocompleteAccountWithDefaults() *AutocompleteAccount`

NewAutocompleteAccountWithDefaults instantiates a new AutocompleteAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutocompleteAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutocompleteAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutocompleteAccount) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutocompleteAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutocompleteAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutocompleteAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNameWithBalance

`func (o *AutocompleteAccount) GetNameWithBalance() string`

GetNameWithBalance returns the NameWithBalance field if non-nil, zero value otherwise.

### GetNameWithBalanceOk

`func (o *AutocompleteAccount) GetNameWithBalanceOk() (*string, bool)`

GetNameWithBalanceOk returns a tuple with the NameWithBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameWithBalance

`func (o *AutocompleteAccount) SetNameWithBalance(v string)`

SetNameWithBalance sets NameWithBalance field to given value.


### GetType

`func (o *AutocompleteAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutocompleteAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutocompleteAccount) SetType(v string)`

SetType sets Type field to given value.


### GetCurrencyId

`func (o *AutocompleteAccount) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AutocompleteAccount) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AutocompleteAccount) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetCurrencyCode

`func (o *AutocompleteAccount) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutocompleteAccount) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutocompleteAccount) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCurrencySymbol

`func (o *AutocompleteAccount) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *AutocompleteAccount) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *AutocompleteAccount) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.


### GetCurrencyDecimalPlaces

`func (o *AutocompleteAccount) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *AutocompleteAccount) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *AutocompleteAccount) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


