# InsightTransferEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This ID is a reference to the original object. | [optional] 
**Name** | Pointer to **string** | This is the name of the object. | [optional] 
**Difference** | Pointer to **string** | The total amount transferred between start date and end date, a number defined as a string, for this asset account. | [optional] 
**DifferenceFloat** | Pointer to **float64** | The total amount transferred between start date and end date, a number as a float, for this asset account. May have rounding errors. | [optional] 
**In** | Pointer to **string** | The total amount transferred TO this account between start date and end date, a number defined as a string, for this asset account. | [optional] 
**InFloat** | Pointer to **float64** | The total amount transferred FROM this account between start date and end date, a number as a float, for this asset account. May have rounding errors. | [optional] 
**Out** | Pointer to **string** | The total amount transferred FROM this account between start date and end date, a number defined as a string, for this asset account. | [optional] 
**OutFloat** | Pointer to **float64** | The total amount transferred TO this account between start date and end date, a number as a float, for this asset account. May have rounding errors. | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the expenses listed for this account. | [optional] 
**CurrencyCode** | Pointer to **string** | The currency code of the expenses listed for this account. | [optional] 

## Methods

### NewInsightTransferEntry

`func NewInsightTransferEntry() *InsightTransferEntry`

NewInsightTransferEntry instantiates a new InsightTransferEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightTransferEntryWithDefaults

`func NewInsightTransferEntryWithDefaults() *InsightTransferEntry`

NewInsightTransferEntryWithDefaults instantiates a new InsightTransferEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightTransferEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightTransferEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightTransferEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InsightTransferEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InsightTransferEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightTransferEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightTransferEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InsightTransferEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDifference

`func (o *InsightTransferEntry) GetDifference() string`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *InsightTransferEntry) GetDifferenceOk() (*string, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *InsightTransferEntry) SetDifference(v string)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *InsightTransferEntry) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetDifferenceFloat

`func (o *InsightTransferEntry) GetDifferenceFloat() float64`

GetDifferenceFloat returns the DifferenceFloat field if non-nil, zero value otherwise.

### GetDifferenceFloatOk

`func (o *InsightTransferEntry) GetDifferenceFloatOk() (*float64, bool)`

GetDifferenceFloatOk returns a tuple with the DifferenceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceFloat

`func (o *InsightTransferEntry) SetDifferenceFloat(v float64)`

SetDifferenceFloat sets DifferenceFloat field to given value.

### HasDifferenceFloat

`func (o *InsightTransferEntry) HasDifferenceFloat() bool`

HasDifferenceFloat returns a boolean if a field has been set.

### GetIn

`func (o *InsightTransferEntry) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *InsightTransferEntry) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *InsightTransferEntry) SetIn(v string)`

SetIn sets In field to given value.

### HasIn

`func (o *InsightTransferEntry) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetInFloat

`func (o *InsightTransferEntry) GetInFloat() float64`

GetInFloat returns the InFloat field if non-nil, zero value otherwise.

### GetInFloatOk

`func (o *InsightTransferEntry) GetInFloatOk() (*float64, bool)`

GetInFloatOk returns a tuple with the InFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFloat

`func (o *InsightTransferEntry) SetInFloat(v float64)`

SetInFloat sets InFloat field to given value.

### HasInFloat

`func (o *InsightTransferEntry) HasInFloat() bool`

HasInFloat returns a boolean if a field has been set.

### GetOut

`func (o *InsightTransferEntry) GetOut() string`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *InsightTransferEntry) GetOutOk() (*string, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *InsightTransferEntry) SetOut(v string)`

SetOut sets Out field to given value.

### HasOut

`func (o *InsightTransferEntry) HasOut() bool`

HasOut returns a boolean if a field has been set.

### GetOutFloat

`func (o *InsightTransferEntry) GetOutFloat() float64`

GetOutFloat returns the OutFloat field if non-nil, zero value otherwise.

### GetOutFloatOk

`func (o *InsightTransferEntry) GetOutFloatOk() (*float64, bool)`

GetOutFloatOk returns a tuple with the OutFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutFloat

`func (o *InsightTransferEntry) SetOutFloat(v float64)`

SetOutFloat sets OutFloat field to given value.

### HasOutFloat

`func (o *InsightTransferEntry) HasOutFloat() bool`

HasOutFloat returns a boolean if a field has been set.

### GetCurrencyId

`func (o *InsightTransferEntry) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InsightTransferEntry) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InsightTransferEntry) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InsightTransferEntry) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InsightTransferEntry) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InsightTransferEntry) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InsightTransferEntry) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InsightTransferEntry) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


