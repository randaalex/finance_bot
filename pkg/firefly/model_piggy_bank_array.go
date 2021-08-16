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

// PiggyBankArray struct for PiggyBankArray
type PiggyBankArray struct {
	Data []PiggyBankRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewPiggyBankArray instantiates a new PiggyBankArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankArray(data []PiggyBankRead, meta Meta, links PageLink) *PiggyBankArray {
	this := PiggyBankArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewPiggyBankArrayWithDefaults instantiates a new PiggyBankArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankArrayWithDefaults() *PiggyBankArray {
	this := PiggyBankArray{}
	return &this
}

// GetData returns the Data field value
func (o *PiggyBankArray) GetData() []PiggyBankRead {
	if o == nil {
		var ret []PiggyBankRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PiggyBankArray) GetDataOk() (*[]PiggyBankRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PiggyBankArray) SetData(v []PiggyBankRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *PiggyBankArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *PiggyBankArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *PiggyBankArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *PiggyBankArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PiggyBankArray) GetLinksOk() (*PageLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PiggyBankArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o PiggyBankArray) MarshalJSON() ([]byte, error) {
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

type NullablePiggyBankArray struct {
	value *PiggyBankArray
	isSet bool
}

func (v NullablePiggyBankArray) Get() *PiggyBankArray {
	return v.value
}

func (v *NullablePiggyBankArray) Set(val *PiggyBankArray) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankArray) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankArray(val *PiggyBankArray) *NullablePiggyBankArray {
	return &NullablePiggyBankArray{value: val, isSet: true}
}

func (v NullablePiggyBankArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


