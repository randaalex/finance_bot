# InsightTotalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Difference** | Pointer to **string** | The amount spent between start date and end date, defined as a string, for this expense account and all asset accounts. | [optional] 
**DifferenceFloat** | Pointer to **float64** | The amount spent between start date and end date, defined as a string, for this expense account and all asset accounts. This number is a float (double) and may have rounding errors. | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the expenses listed for this expense account. | [optional] 
**CurrencyCode** | Pointer to **string** | The currency code of the expenses listed for this expense account. | [optional] 

## Methods

### NewInsightTotalEntry

`func NewInsightTotalEntry() *InsightTotalEntry`

NewInsightTotalEntry instantiates a new InsightTotalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightTotalEntryWithDefaults

`func NewInsightTotalEntryWithDefaults() *InsightTotalEntry`

NewInsightTotalEntryWithDefaults instantiates a new InsightTotalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDifference

`func (o *InsightTotalEntry) GetDifference() string`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *InsightTotalEntry) GetDifferenceOk() (*string, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *InsightTotalEntry) SetDifference(v string)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *InsightTotalEntry) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetDifferenceFloat

`func (o *InsightTotalEntry) GetDifferenceFloat() float64`

GetDifferenceFloat returns the DifferenceFloat field if non-nil, zero value otherwise.

### GetDifferenceFloatOk

`func (o *InsightTotalEntry) GetDifferenceFloatOk() (*float64, bool)`

GetDifferenceFloatOk returns a tuple with the DifferenceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceFloat

`func (o *InsightTotalEntry) SetDifferenceFloat(v float64)`

SetDifferenceFloat sets DifferenceFloat field to given value.

### HasDifferenceFloat

`func (o *InsightTotalEntry) HasDifferenceFloat() bool`

HasDifferenceFloat returns a boolean if a field has been set.

### GetCurrencyId

`func (o *InsightTotalEntry) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InsightTotalEntry) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InsightTotalEntry) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InsightTotalEntry) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InsightTotalEntry) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InsightTotalEntry) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InsightTotalEntry) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InsightTotalEntry) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


