/*
 * Firefly III API v1.5.2
 *
 * This is the documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. You may use the \"Authorize\" button to try the API below. This file was last generated on 2021-05-14T15:49:56+00:00 
 *
 * API version: 1.5.2
 * Contact: james@firefly-iii.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package firefly

import (
	"encoding/json"
)

// InsightTotalEntry struct for InsightTotalEntry
type InsightTotalEntry struct {
	// The amount spent between start date and end date, defined as a string, for this expense account and all asset accounts.
	Difference *string `json:"difference,omitempty"`
	// The amount spent between start date and end date, defined as a string, for this expense account and all asset accounts. This number is a float (double) and may have rounding errors.
	DifferenceFloat *float64 `json:"difference_float,omitempty"`
	// The currency ID of the expenses listed for this expense account.
	CurrencyId *string `json:"currency_id,omitempty"`
	// The currency code of the expenses listed for this expense account.
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// NewInsightTotalEntry instantiates a new InsightTotalEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightTotalEntry() *InsightTotalEntry {
	this := InsightTotalEntry{}
	return &this
}

// NewInsightTotalEntryWithDefaults instantiates a new InsightTotalEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightTotalEntryWithDefaults() *InsightTotalEntry {
	this := InsightTotalEntry{}
	return &this
}

// GetDifference returns the Difference field value if set, zero value otherwise.
func (o *InsightTotalEntry) GetDifference() string {
	if o == nil || o.Difference == nil {
		var ret string
		return ret
	}
	return *o.Difference
}

// GetDifferenceOk returns a tuple with the Difference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightTotalEntry) GetDifferenceOk() (*string, bool) {
	if o == nil || o.Difference == nil {
		return nil, false
	}
	return o.Difference, true
}

// HasDifference returns a boolean if a field has been set.
func (o *InsightTotalEntry) HasDifference() bool {
	if o != nil && o.Difference != nil {
		return true
	}

	return false
}

// SetDifference gets a reference to the given string and assigns it to the Difference field.
func (o *InsightTotalEntry) SetDifference(v string) {
	o.Difference = &v
}

// GetDifferenceFloat returns the DifferenceFloat field value if set, zero value otherwise.
func (o *InsightTotalEntry) GetDifferenceFloat() float64 {
	if o == nil || o.DifferenceFloat == nil {
		var ret float64
		return ret
	}
	return *o.DifferenceFloat
}

// GetDifferenceFloatOk returns a tuple with the DifferenceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightTotalEntry) GetDifferenceFloatOk() (*float64, bool) {
	if o == nil || o.DifferenceFloat == nil {
		return nil, false
	}
	return o.DifferenceFloat, true
}

// HasDifferenceFloat returns a boolean if a field has been set.
func (o *InsightTotalEntry) HasDifferenceFloat() bool {
	if o != nil && o.DifferenceFloat != nil {
		return true
	}

	return false
}

// SetDifferenceFloat gets a reference to the given float64 and assigns it to the DifferenceFloat field.
func (o *InsightTotalEntry) SetDifferenceFloat(v float64) {
	o.DifferenceFloat = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *InsightTotalEntry) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightTotalEntry) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *InsightTotalEntry) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *InsightTotalEntry) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *InsightTotalEntry) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsightTotalEntry) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *InsightTotalEntry) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *InsightTotalEntry) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

func (o InsightTotalEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Difference != nil {
		toSerialize["difference"] = o.Difference
	}
	if o.DifferenceFloat != nil {
		toSerialize["difference_float"] = o.DifferenceFloat
	}
	if o.CurrencyId != nil {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	return json.Marshal(toSerialize)
}

type NullableInsightTotalEntry struct {
	value *InsightTotalEntry
	isSet bool
}

func (v NullableInsightTotalEntry) Get() *InsightTotalEntry {
	return v.value
}

func (v *NullableInsightTotalEntry) Set(val *InsightTotalEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightTotalEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightTotalEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightTotalEntry(val *InsightTotalEntry) *NullableInsightTotalEntry {
	return &NullableInsightTotalEntry{value: val, isSet: true}
}

func (v NullableInsightTotalEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightTotalEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


