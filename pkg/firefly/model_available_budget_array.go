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

// AvailableBudgetArray struct for AvailableBudgetArray
type AvailableBudgetArray struct {
	Data []AvailableBudgetRead `json:"data"`
	Meta Meta `json:"meta"`
}

// NewAvailableBudgetArray instantiates a new AvailableBudgetArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableBudgetArray(data []AvailableBudgetRead, meta Meta) *AvailableBudgetArray {
	this := AvailableBudgetArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewAvailableBudgetArrayWithDefaults instantiates a new AvailableBudgetArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableBudgetArrayWithDefaults() *AvailableBudgetArray {
	this := AvailableBudgetArray{}
	return &this
}

// GetData returns the Data field value
func (o *AvailableBudgetArray) GetData() []AvailableBudgetRead {
	if o == nil {
		var ret []AvailableBudgetRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvailableBudgetArray) GetDataOk() (*[]AvailableBudgetRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvailableBudgetArray) SetData(v []AvailableBudgetRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *AvailableBudgetArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *AvailableBudgetArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *AvailableBudgetArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o AvailableBudgetArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableBudgetArray struct {
	value *AvailableBudgetArray
	isSet bool
}

func (v NullableAvailableBudgetArray) Get() *AvailableBudgetArray {
	return v.value
}

func (v *NullableAvailableBudgetArray) Set(val *AvailableBudgetArray) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableBudgetArray) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableBudgetArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableBudgetArray(val *AvailableBudgetArray) *NullableAvailableBudgetArray {
	return &NullableAvailableBudgetArray{value: val, isSet: true}
}

func (v NullableAvailableBudgetArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableBudgetArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


