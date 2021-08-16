# InsightGroupEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This ID is a reference to the original object. | [optional] 
**Name** | Pointer to **string** | This is the name of the object. | [optional] 
**Difference** | Pointer to **string** | The amount spent or earned between start date and end date, a number defined as a string, for this object and all asset accounts. | [optional] 
**DifferenceFloat** | Pointer to **float64** | The amount spent or earned between start date and end date, a number as a float, for this object and all asset accounts. May have rounding errors. | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the expenses listed for this account. | [optional] 
**CurrencyCode** | Pointer to **string** | The currency code of the expenses listed for this account. | [optional] 

## Methods

### NewInsightGroupEntry

`func NewInsightGroupEntry() *InsightGroupEntry`

NewInsightGroupEntry instantiates a new InsightGroupEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightGroupEntryWithDefaults

`func NewInsightGroupEntryWithDefaults() *InsightGroupEntry`

NewInsightGroupEntryWithDefaults instantiates a new InsightGroupEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InsightGroupEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsightGroupEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsightGroupEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InsightGroupEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InsightGroupEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightGroupEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightGroupEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InsightGroupEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDifference

`func (o *InsightGroupEntry) GetDifference() string`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *InsightGroupEntry) GetDifferenceOk() (*string, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *InsightGroupEntry) SetDifference(v string)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *InsightGroupEntry) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetDifferenceFloat

`func (o *InsightGroupEntry) GetDifferenceFloat() float64`

GetDifferenceFloat returns the DifferenceFloat field if non-nil, zero value otherwise.

### GetDifferenceFloatOk

`func (o *InsightGroupEntry) GetDifferenceFloatOk() (*float64, bool)`

GetDifferenceFloatOk returns a tuple with the DifferenceFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceFloat

`func (o *InsightGroupEntry) SetDifferenceFloat(v float64)`

SetDifferenceFloat sets DifferenceFloat field to given value.

### HasDifferenceFloat

`func (o *InsightGroupEntry) HasDifferenceFloat() bool`

HasDifferenceFloat returns a boolean if a field has been set.

### GetCurrencyId

`func (o *InsightGroupEntry) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *InsightGroupEntry) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *InsightGroupEntry) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *InsightGroupEntry) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InsightGroupEntry) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InsightGroupEntry) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InsightGroupEntry) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InsightGroupEntry) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


