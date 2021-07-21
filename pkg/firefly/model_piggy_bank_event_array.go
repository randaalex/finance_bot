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

// PiggyBankEventArray struct for PiggyBankEventArray
type PiggyBankEventArray struct {
	Data []PiggyBankEventRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewPiggyBankEventArray instantiates a new PiggyBankEventArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankEventArray(data []PiggyBankEventRead, meta Meta, links PageLink) *PiggyBankEventArray {
	this := PiggyBankEventArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewPiggyBankEventArrayWithDefaults instantiates a new PiggyBankEventArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankEventArrayWithDefaults() *PiggyBankEventArray {
	this := PiggyBankEventArray{}
	return &this
}

// GetData returns the Data field value
func (o *PiggyBankEventArray) GetData() []PiggyBankEventRead {
	if o == nil {
		var ret []PiggyBankEventRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetDataOk() (*[]PiggyBankEventRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PiggyBankEventArray) SetData(v []PiggyBankEventRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *PiggyBankEventArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *PiggyBankEventArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *PiggyBankEventArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PiggyBankEventArray) GetLinksOk() (*PageLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PiggyBankEventArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o PiggyBankEventArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullablePiggyBankEventArray struct {
	value *PiggyBankEventArray
	isSet bool
}

func (v NullablePiggyBankEventArray) Get() *PiggyBankEventArray {
	return v.value
}

func (v *NullablePiggyBankEventArray) Set(val *PiggyBankEventArray) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankEventArray) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankEventArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankEventArray(val *PiggyBankEventArray) *NullablePiggyBankEventArray {
	return &NullablePiggyBankEventArray{value: val, isSet: true}
}

func (v NullablePiggyBankEventArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankEventArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


