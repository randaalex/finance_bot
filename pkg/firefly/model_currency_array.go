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

// CurrencyArray struct for CurrencyArray
type CurrencyArray struct {
	Data []CurrencyRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewCurrencyArray instantiates a new CurrencyArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyArray(data []CurrencyRead, meta Meta, links PageLink) *CurrencyArray {
	this := CurrencyArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewCurrencyArrayWithDefaults instantiates a new CurrencyArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyArrayWithDefaults() *CurrencyArray {
	this := CurrencyArray{}
	return &this
}

// GetData returns the Data field value
func (o *CurrencyArray) GetData() []CurrencyRead {
	if o == nil {
		var ret []CurrencyRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CurrencyArray) GetDataOk() (*[]CurrencyRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CurrencyArray) SetData(v []CurrencyRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *CurrencyArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CurrencyArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CurrencyArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *CurrencyArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CurrencyArray) GetLinksOk() (*PageLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CurrencyArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o CurrencyArray) MarshalJSON() ([]byte, error) {
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

type NullableCurrencyArray struct {
	value *CurrencyArray
	isSet bool
}

func (v NullableCurrencyArray) Get() *CurrencyArray {
	return v.value
}

func (v *NullableCurrencyArray) Set(val *CurrencyArray) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyArray) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyArray(val *CurrencyArray) *NullableCurrencyArray {
	return &NullableCurrencyArray{value: val, isSet: true}
}

func (v NullableCurrencyArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


