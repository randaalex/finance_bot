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

// LinkType struct for LinkType
type LinkType struct {
	Name string `json:"name"`
	Inward string `json:"inward"`
	Outward string `json:"outward"`
	Editable *bool `json:"editable,omitempty"`
}

// NewLinkType instantiates a new LinkType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkType(name string, inward string, outward string) *LinkType {
	this := LinkType{}
	this.Name = name
	this.Inward = inward
	this.Outward = outward
	return &this
}

// NewLinkTypeWithDefaults instantiates a new LinkType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTypeWithDefaults() *LinkType {
	this := LinkType{}
	return &this
}

// GetName returns the Name field value
func (o *LinkType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LinkType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LinkType) SetName(v string) {
	o.Name = v
}

// GetInward returns the Inward field value
func (o *LinkType) GetInward() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Inward
}

// GetInwardOk returns a tuple with the Inward field value
// and a boolean to check if the value has been set.
func (o *LinkType) GetInwardOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Inward, true
}

// SetInward sets field value
func (o *LinkType) SetInward(v string) {
	o.Inward = v
}

// GetOutward returns the Outward field value
func (o *LinkType) GetOutward() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outward
}

// GetOutwardOk returns a tuple with the Outward field value
// and a boolean to check if the value has been set.
func (o *LinkType) GetOutwardOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Outward, true
}

// SetOutward sets field value
func (o *LinkType) SetOutward(v string) {
	o.Outward = v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *LinkType) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkType) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *LinkType) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *LinkType) SetEditable(v bool) {
	o.Editable = &v
}

func (o LinkType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["inward"] = o.Inward
	}
	if true {
		toSerialize["outward"] = o.Outward
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	return json.Marshal(toSerialize)
}

type NullableLinkType struct {
	value *LinkType
	isSet bool
}

func (v NullableLinkType) Get() *LinkType {
	return v.value
}

func (v *NullableLinkType) Set(val *LinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkType(val *LinkType) *NullableLinkType {
	return &NullableLinkType{value: val, isSet: true}
}

func (v NullableLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


