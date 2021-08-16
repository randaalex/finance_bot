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

// TransactionLinkArray struct for TransactionLinkArray
type TransactionLinkArray struct {
	Data []TransactionLinkRead `json:"data"`
	Meta Meta `json:"meta"`
	Links PageLink `json:"links"`
}

// NewTransactionLinkArray instantiates a new TransactionLinkArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionLinkArray(data []TransactionLinkRead, meta Meta, links PageLink) *TransactionLinkArray {
	this := TransactionLinkArray{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewTransactionLinkArrayWithDefaults instantiates a new TransactionLinkArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionLinkArrayWithDefaults() *TransactionLinkArray {
	this := TransactionLinkArray{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionLinkArray) GetData() []TransactionLinkRead {
	if o == nil {
		var ret []TransactionLinkRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkArray) GetDataOk() (*[]TransactionLinkRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionLinkArray) SetData(v []TransactionLinkRead) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *TransactionLinkArray) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkArray) GetMetaOk() (*Meta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *TransactionLinkArray) SetMeta(v Meta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *TransactionLinkArray) GetLinks() PageLink {
	if o == nil {
		var ret PageLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkArray) GetLinksOk() (*PageLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TransactionLinkArray) SetLinks(v PageLink) {
	o.Links = v
}

func (o TransactionLinkArray) MarshalJSON() ([]byte, error) {
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

type NullableTransactionLinkArray struct {
	value *TransactionLinkArray
	isSet bool
}

func (v NullableTransactionLinkArray) Get() *TransactionLinkArray {
	return v.value
}

func (v *NullableTransactionLinkArray) Set(val *TransactionLinkArray) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionLinkArray) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionLinkArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionLinkArray(val *TransactionLinkArray) *NullableTransactionLinkArray {
	return &NullableTransactionLinkArray{value: val, isSet: true}
}

func (v NullableTransactionLinkArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionLinkArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


