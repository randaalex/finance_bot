# BasicSummaryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | This is a reference to the type of info shared, not influenced by translations or user preferences. The EUR value is a reference to the currency code. Possibilities are: balance-in-ABC, spent-in-ABC, earned-in-ABC, bills-paid-in-ABC, bills-unpaid-in-ABC, left-to-spend-in-ABC and net-worth-in-ABC. | [optional] 
**Title** | Pointer to **string** | A translated title for the information shared. | [optional] 
**MonetaryValue** | Pointer to **float64** | The amount as a float. | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the associated currency. | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**CurrencyDecimalPlaces** | Pointer to **int32** | Number of decimals for the associated currency. | [optional] 
**ValueParsed** | Pointer to **string** | The amount formatted according to the users locale | [optional] 
**LocalIcon** | Pointer to **string** | Reference to a font-awesome icon without the fa- part. | [optional] 
**SubTitle** | Pointer to **string** | A short explanation of the amounts origin. Already formatted according to the locale of the user or translated, if relevant. | [optional] 

## Methods

### NewBasicSummaryEntry

`func NewBasicSummaryEntry() *BasicSummaryEntry`

NewBasicSummaryEntry instantiates a new BasicSummaryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicSummaryEntryWithDefaults

`func NewBasicSummaryEntryWithDefaults() *BasicSummaryEntry`

NewBasicSummaryEntryWithDefaults instantiates a new BasicSummaryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BasicSummaryEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BasicSummaryEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BasicSummaryEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BasicSummaryEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTitle

`func (o *BasicSummaryEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BasicSummaryEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BasicSummaryEntry) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BasicSummaryEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMonetaryValue

`func (o *BasicSummaryEntry) GetMonetaryValue() float64`

GetMonetaryValue returns the MonetaryValue field if non-nil, zero value otherwise.

### GetMonetaryValueOk

`func (o *BasicSummaryEntry) GetMonetaryValueOk() (*float64, bool)`

GetMonetaryValueOk returns a tuple with the MonetaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonetaryValue

`func (o *BasicSummaryEntry) SetMonetaryValue(v float64)`

SetMonetaryValue sets MonetaryValue field to given value.

### HasMonetaryValue

`func (o *BasicSummaryEntry) HasMonetaryValue() bool`

HasMonetaryValue returns a boolean if a field has been set.

### GetCurrencyId

`func (o *BasicSummaryEntry) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BasicSummaryEntry) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BasicSummaryEntry) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BasicSummaryEntry) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *BasicSummaryEntry) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *BasicSummaryEntry) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *BasicSummaryEntry) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *BasicSummaryEntry) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *BasicSummaryEntry) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *BasicSummaryEntry) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *BasicSummaryEntry) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *BasicSummaryEntry) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *BasicSummaryEntry) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *BasicSummaryEntry) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *BasicSummaryEntry) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *BasicSummaryEntry) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetValueParsed

`func (o *BasicSummaryEntry) GetValueParsed() string`

GetValueParsed returns the ValueParsed field if non-nil, zero value otherwise.

### GetValueParsedOk

`func (o *BasicSummaryEntry) GetValueParsedOk() (*string, bool)`

GetValueParsedOk returns a tuple with the ValueParsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueParsed

`func (o *BasicSummaryEntry) SetValueParsed(v string)`

SetValueParsed sets ValueParsed field to given value.

### HasValueParsed

`func (o *BasicSummaryEntry) HasValueParsed() bool`

HasValueParsed returns a boolean if a field has been set.

### GetLocalIcon

`func (o *BasicSummaryEntry) GetLocalIcon() string`

GetLocalIcon returns the LocalIcon field if non-nil, zero value otherwise.

### GetLocalIconOk

`func (o *BasicSummaryEntry) GetLocalIconOk() (*string, bool)`

GetLocalIconOk returns a tuple with the LocalIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIcon

`func (o *BasicSummaryEntry) SetLocalIcon(v string)`

SetLocalIcon sets LocalIcon field to given value.

### HasLocalIcon

`func (o *BasicSummaryEntry) HasLocalIcon() bool`

HasLocalIcon returns a boolean if a field has been set.

### GetSubTitle

`func (o *BasicSummaryEntry) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *BasicSummaryEntry) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *BasicSummaryEntry) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.

### HasSubTitle

`func (o *BasicSummaryEntry) HasSubTitle() bool`

HasSubTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


