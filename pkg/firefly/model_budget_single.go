/*
 * Firefly III API
 *
 * This is the official documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. This version of the API is live from version v4.7.9 and onwards. You may use the \"Authorize\" button to try the API below. 
 *
 * API version: 1.4.0
 * Contact: james@firefly-iii.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package firefly

import (
	"encoding/json"
)

// BudgetSingle struct for BudgetSingle
type BudgetSingle struct {
	Data BudgetRead `json:"data"`
}

// NewBudgetSingle instantiates a new BudgetSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetSingle(data BudgetRead) *BudgetSingle {
	this := BudgetSingle{}
	this.Data = data
	return &this
}

// NewBudgetSingleWithDefaults instantiates a new BudgetSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetSingleWithDefaults() *BudgetSingle {
	this := BudgetSingle{}
	return &this
}

// GetData returns the Data field value
func (o *BudgetSingle) GetData() BudgetRead {
	if o == nil {
		var ret BudgetRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BudgetSingle) GetDataOk() (*BudgetRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BudgetSingle) SetData(v BudgetRead) {
	o.Data = v
}

func (o BudgetSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetSingle struct {
	value *BudgetSingle
	isSet bool
}

func (v NullableBudgetSingle) Get() *BudgetSingle {
	return v.value
}

func (v *NullableBudgetSingle) Set(val *BudgetSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetSingle(val *BudgetSingle) *NullableBudgetSingle {
	return &NullableBudgetSingle{value: val, isSet: true}
}

func (v NullableBudgetSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

