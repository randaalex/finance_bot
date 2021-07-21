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

// CategoryArray struct for CategoryArray
type CategoryArray struct {
	Data []CategoryRead `json:"data"`
	Meta Meta `json:"meta"`
}

// NewCategoryArray instantiates a new CategoryArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryArray(data []CategoryRead, meta Meta) *CategoryArray {
	this := CategoryArray{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewCategoryArrayWithDefaults instantiates a new CategoryArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryArrayWithDefaults() *CategoryArray {
	this := CategoryArray{}
	return &this
}

// GetData returns the Data field value
func (o *CategoryArray) GetData() []CategoryRead {
	if o == nil {
		var ret []CategoryRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CategoryArray) GetDataOk() (*[]CategoryRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CategoryArray) SetData(v []CategoryRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *CategoryArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CategoryArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CategoryArray) SetMeta(v Meta) {
	o.Meta = v
}

func (o CategoryArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryArray struct {
	value *CategoryArray
	isSet bool
}

func (v NullableCategoryArray) Get() *CategoryArray {
	return v.value
}

func (v *NullableCategoryArray) Set(val *CategoryArray) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryArray) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryArray(val *CategoryArray) *NullableCategoryArray {
	return &NullableCategoryArray{value: val, isSet: true}
}

func (v NullableCategoryArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


