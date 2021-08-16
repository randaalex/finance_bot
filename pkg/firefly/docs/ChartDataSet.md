# ChartDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | This is the title of the current set. It can refer to an account, a budget or another object (by name). | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the currency associated to the data in the entries. | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**CurrencyDecimalPlaces** | Pointer to **int32** | Number of decimals for the currency associated to the data in the entries. | [optional] 
**Type** | Pointer to **string** | Indicated the type of chart that is expected to be rendered. You can safely ignore this if you want. | [optional] 
**YAxisID** | Pointer to **int32** | Used to indicate the Y axis for this data set. Is usually between 0 and 1 (left and right side of the chart). | [optional] 
**Entries** | Pointer to [**[]ChartDataPoint**](ChartDataPoint.md) | The actual entries for this data set. They &#39;key&#39; value is the label for the data point. The value is the actual (numerical) value. | [optional] 

## Methods

### NewChartDataSet

`func NewChartDataSet() *ChartDataSet`

NewChartDataSet instantiates a new ChartDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartDataSetWithDefaults

`func NewChartDataSetWithDefaults() *ChartDataSet`

NewChartDataSetWithDefaults instantiates a new ChartDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ChartDataSet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ChartDataSet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ChartDataSet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ChartDataSet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCurrencyId

`func (o *ChartDataSet) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ChartDataSet) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ChartDataSet) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ChartDataSet) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ChartDataSet) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ChartDataSet) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ChartDataSet) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ChartDataSet) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *ChartDataSet) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *ChartDataSet) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *ChartDataSet) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *ChartDataSet) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *ChartDataSet) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *ChartDataSet) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *ChartDataSet) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *ChartDataSet) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetType

`func (o *ChartDataSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChartDataSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChartDataSet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChartDataSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYAxisID

`func (o *ChartDataSet) GetYAxisID() int32`

GetYAxisID returns the YAxisID field if non-nil, zero value otherwise.

### GetYAxisIDOk

`func (o *ChartDataSet) GetYAxisIDOk() (*int32, bool)`

GetYAxisIDOk returns a tuple with the YAxisID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisID

`func (o *ChartDataSet) SetYAxisID(v int32)`

SetYAxisID sets YAxisID field to given value.

### HasYAxisID

`func (o *ChartDataSet) HasYAxisID() bool`

HasYAxisID returns a boolean if a field has been set.

### GetEntries

`func (o *ChartDataSet) GetEntries() []ChartDataPoint`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ChartDataSet) GetEntriesOk() (*[]ChartDataPoint, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ChartDataSet) SetEntries(v []ChartDataPoint)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ChartDataSet) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


