# ExchangeRateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**FromCurrencyId** | Pointer to **int32** |  | [optional] 
**FromCurrencyName** | Pointer to **string** |  | [optional] 
**FromCurrencyCode** | Pointer to **string** |  | [optional] 
**FromCurrencySymbol** | Pointer to **string** |  | [optional] 
**FromCurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] 
**ToCurrencyId** | Pointer to **int32** |  | [optional] 
**ToCurrencyCode** | Pointer to **string** |  | [optional] 
**ToCurrencySymbol** | Pointer to **string** |  | [optional] 
**ToCurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **float32** |  | [optional] 
**Amount** | Pointer to **string** | The amount in the \&quot;to\&quot;-currency, if provided in the request. | [optional] 

## Methods

### NewExchangeRateAttributes

`func NewExchangeRateAttributes() *ExchangeRateAttributes`

NewExchangeRateAttributes instantiates a new ExchangeRateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRateAttributesWithDefaults

`func NewExchangeRateAttributesWithDefaults() *ExchangeRateAttributes`

NewExchangeRateAttributesWithDefaults instantiates a new ExchangeRateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ExchangeRateAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExchangeRateAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExchangeRateAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExchangeRateAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExchangeRateAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExchangeRateAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExchangeRateAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExchangeRateAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFromCurrencyId

`func (o *ExchangeRateAttributes) GetFromCurrencyId() int32`

GetFromCurrencyId returns the FromCurrencyId field if non-nil, zero value otherwise.

### GetFromCurrencyIdOk

`func (o *ExchangeRateAttributes) GetFromCurrencyIdOk() (*int32, bool)`

GetFromCurrencyIdOk returns a tuple with the FromCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyId

`func (o *ExchangeRateAttributes) SetFromCurrencyId(v int32)`

SetFromCurrencyId sets FromCurrencyId field to given value.

### HasFromCurrencyId

`func (o *ExchangeRateAttributes) HasFromCurrencyId() bool`

HasFromCurrencyId returns a boolean if a field has been set.

### GetFromCurrencyName

`func (o *ExchangeRateAttributes) GetFromCurrencyName() string`

GetFromCurrencyName returns the FromCurrencyName field if non-nil, zero value otherwise.

### GetFromCurrencyNameOk

`func (o *ExchangeRateAttributes) GetFromCurrencyNameOk() (*string, bool)`

GetFromCurrencyNameOk returns a tuple with the FromCurrencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyName

`func (o *ExchangeRateAttributes) SetFromCurrencyName(v string)`

SetFromCurrencyName sets FromCurrencyName field to given value.

### HasFromCurrencyName

`func (o *ExchangeRateAttributes) HasFromCurrencyName() bool`

HasFromCurrencyName returns a boolean if a field has been set.

### GetFromCurrencyCode

`func (o *ExchangeRateAttributes) GetFromCurrencyCode() string`

GetFromCurrencyCode returns the FromCurrencyCode field if non-nil, zero value otherwise.

### GetFromCurrencyCodeOk

`func (o *ExchangeRateAttributes) GetFromCurrencyCodeOk() (*string, bool)`

GetFromCurrencyCodeOk returns a tuple with the FromCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyCode

`func (o *ExchangeRateAttributes) SetFromCurrencyCode(v string)`

SetFromCurrencyCode sets FromCurrencyCode field to given value.

### HasFromCurrencyCode

`func (o *ExchangeRateAttributes) HasFromCurrencyCode() bool`

HasFromCurrencyCode returns a boolean if a field has been set.

### GetFromCurrencySymbol

`func (o *ExchangeRateAttributes) GetFromCurrencySymbol() string`

GetFromCurrencySymbol returns the FromCurrencySymbol field if non-nil, zero value otherwise.

### GetFromCurrencySymbolOk

`func (o *ExchangeRateAttributes) GetFromCurrencySymbolOk() (*string, bool)`

GetFromCurrencySymbolOk returns a tuple with the FromCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencySymbol

`func (o *ExchangeRateAttributes) SetFromCurrencySymbol(v string)`

SetFromCurrencySymbol sets FromCurrencySymbol field to given value.

### HasFromCurrencySymbol

`func (o *ExchangeRateAttributes) HasFromCurrencySymbol() bool`

HasFromCurrencySymbol returns a boolean if a field has been set.

### GetFromCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) GetFromCurrencyDecimalPlaces() int32`

GetFromCurrencyDecimalPlaces returns the FromCurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetFromCurrencyDecimalPlacesOk

`func (o *ExchangeRateAttributes) GetFromCurrencyDecimalPlacesOk() (*int32, bool)`

GetFromCurrencyDecimalPlacesOk returns a tuple with the FromCurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) SetFromCurrencyDecimalPlaces(v int32)`

SetFromCurrencyDecimalPlaces sets FromCurrencyDecimalPlaces field to given value.

### HasFromCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) HasFromCurrencyDecimalPlaces() bool`

HasFromCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetToCurrencyId

`func (o *ExchangeRateAttributes) GetToCurrencyId() int32`

GetToCurrencyId returns the ToCurrencyId field if non-nil, zero value otherwise.

### GetToCurrencyIdOk

`func (o *ExchangeRateAttributes) GetToCurrencyIdOk() (*int32, bool)`

GetToCurrencyIdOk returns a tuple with the ToCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyId

`func (o *ExchangeRateAttributes) SetToCurrencyId(v int32)`

SetToCurrencyId sets ToCurrencyId field to given value.

### HasToCurrencyId

`func (o *ExchangeRateAttributes) HasToCurrencyId() bool`

HasToCurrencyId returns a boolean if a field has been set.

### GetToCurrencyCode

`func (o *ExchangeRateAttributes) GetToCurrencyCode() string`

GetToCurrencyCode returns the ToCurrencyCode field if non-nil, zero value otherwise.

### GetToCurrencyCodeOk

`func (o *ExchangeRateAttributes) GetToCurrencyCodeOk() (*string, bool)`

GetToCurrencyCodeOk returns a tuple with the ToCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyCode

`func (o *ExchangeRateAttributes) SetToCurrencyCode(v string)`

SetToCurrencyCode sets ToCurrencyCode field to given value.

### HasToCurrencyCode

`func (o *ExchangeRateAttributes) HasToCurrencyCode() bool`

HasToCurrencyCode returns a boolean if a field has been set.

### GetToCurrencySymbol

`func (o *ExchangeRateAttributes) GetToCurrencySymbol() string`

GetToCurrencySymbol returns the ToCurrencySymbol field if non-nil, zero value otherwise.

### GetToCurrencySymbolOk

`func (o *ExchangeRateAttributes) GetToCurrencySymbolOk() (*string, bool)`

GetToCurrencySymbolOk returns a tuple with the ToCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencySymbol

`func (o *ExchangeRateAttributes) SetToCurrencySymbol(v string)`

SetToCurrencySymbol sets ToCurrencySymbol field to given value.

### HasToCurrencySymbol

`func (o *ExchangeRateAttributes) HasToCurrencySymbol() bool`

HasToCurrencySymbol returns a boolean if a field has been set.

### GetToCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) GetToCurrencyDecimalPlaces() int32`

GetToCurrencyDecimalPlaces returns the ToCurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetToCurrencyDecimalPlacesOk

`func (o *ExchangeRateAttributes) GetToCurrencyDecimalPlacesOk() (*int32, bool)`

GetToCurrencyDecimalPlacesOk returns a tuple with the ToCurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) SetToCurrencyDecimalPlaces(v int32)`

SetToCurrencyDecimalPlaces sets ToCurrencyDecimalPlaces field to given value.

### HasToCurrencyDecimalPlaces

`func (o *ExchangeRateAttributes) HasToCurrencyDecimalPlaces() bool`

HasToCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetDate

`func (o *ExchangeRateAttributes) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExchangeRateAttributes) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExchangeRateAttributes) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ExchangeRateAttributes) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetRate

`func (o *ExchangeRateAttributes) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExchangeRateAttributes) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExchangeRateAttributes) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ExchangeRateAttributes) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetAmount

`func (o *ExchangeRateAttributes) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExchangeRateAttributes) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExchangeRateAttributes) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExchangeRateAttributes) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


