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

// ObjectLink0 struct for ObjectLink0
type ObjectLink0 struct {
	Rel *string `json:"rel,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewObjectLink0 instantiates a new ObjectLink0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLink0() *ObjectLink0 {
	this := ObjectLink0{}
	return &this
}

// NewObjectLink0WithDefaults instantiates a new ObjectLink0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLink0WithDefaults() *ObjectLink0 {
	this := ObjectLink0{}
	return &this
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *ObjectLink0) GetRel() string {
	if o == nil || o.Rel == nil {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLink0) GetRelOk() (*string, bool) {
	if o == nil || o.Rel == nil {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *ObjectLink0) HasRel() bool {
	if o != nil && o.Rel != nil {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *ObjectLink0) SetRel(v string) {
	o.Rel = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ObjectLink0) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLink0) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ObjectLink0) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ObjectLink0) SetUri(v string) {
	o.Uri = &v
}

func (o ObjectLink0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rel != nil {
		toSerialize["rel"] = o.Rel
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableObjectLink0 struct {
	value *ObjectLink0
	isSet bool
}

func (v NullableObjectLink0) Get() *ObjectLink0 {
	return v.value
}

func (v *NullableObjectLink0) Set(val *ObjectLink0) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLink0) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLink0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLink0(val *ObjectLink0) *NullableObjectLink0 {
	return &NullableObjectLink0{value: val, isSet: true}
}

func (v NullableObjectLink0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLink0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


