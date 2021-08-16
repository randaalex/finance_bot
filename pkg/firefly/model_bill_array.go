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

// BillArray struct for BillArray
type BillArray struct {
	Data []BillRead `json:"data"`
	Meta Meta `json:"meta"`
}

// NewBillArray instantiates a new BillArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillArray(data []BillRead, meta Meta) *BillArray {
	this := BillArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewBillArrayWithDefaults instantiates a new BillArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillArrayWithDefaults() *BillArray {
	this := BillArray{}
	return &this
}

// GetData returns the Data field value
func (o *BillArray) GetData() []BillRead {
	if o == nil {
		var ret []BillRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BillArray) GetDataOk() (*[]BillRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BillArray) SetData(v []BillRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *BillArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BillArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BillArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o BillArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableBillArray struct {
	value *BillArray
	isSet bool
}

func (v NullableBillArray) Get() *BillArray {
	return v.value
}

func (v *NullableBillArray) Set(val *BillArray) {
	v.value = val
	v.isSet = true
}

func (v NullableBillArray) IsSet() bool {
	return v.isSet
}

func (v *NullableBillArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillArray(val *BillArray) *NullableBillArray {
	return &NullableBillArray{value: val, isSet: true}
}

func (v NullableBillArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


