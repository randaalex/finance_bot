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

// PiggyBankRead struct for PiggyBankRead
type PiggyBankRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes PiggyBank `json:"attributes"`
	Links ObjectLink `json:"links"`
}

// NewPiggyBankRead instantiates a new PiggyBankRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPiggyBankRead(type_ string, id string, attributes PiggyBank, links ObjectLink) *PiggyBankRead {
	this := PiggyBankRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPiggyBankReadWithDefaults instantiates a new PiggyBankRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPiggyBankReadWithDefaults() *PiggyBankRead {
	this := PiggyBankRead{}
	return &this
}

// GetType returns the Type field value
func (o *PiggyBankRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PiggyBankRead) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PiggyBankRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PiggyBankRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PiggyBankRead) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PiggyBankRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PiggyBankRead) GetAttributes() PiggyBank {
	if o == nil {
		var ret PiggyBank
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PiggyBankRead) GetAttributesOk() (*PiggyBank, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PiggyBankRead) SetAttributes(v PiggyBank) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *PiggyBankRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PiggyBankRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PiggyBankRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o PiggyBankRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullablePiggyBankRead struct {
	value *PiggyBankRead
	isSet bool
}

func (v NullablePiggyBankRead) Get() *PiggyBankRead {
	return v.value
}

func (v *NullablePiggyBankRead) Set(val *PiggyBankRead) {
	v.value = val
	v.isSet = true
}

func (v NullablePiggyBankRead) IsSet() bool {
	return v.isSet
}

func (v *NullablePiggyBankRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePiggyBankRead(val *PiggyBankRead) *NullablePiggyBankRead {
	return &NullablePiggyBankRead{value: val, isSet: true}
}

func (v NullablePiggyBankRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePiggyBankRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

