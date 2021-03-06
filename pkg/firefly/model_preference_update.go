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

// PreferenceUpdate struct for PreferenceUpdate
type PreferenceUpdate struct {
	// The actual preference content.
	Data string `json:"data"`
}

// NewPreferenceUpdate instantiates a new PreferenceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferenceUpdate(data string) *PreferenceUpdate {
	this := PreferenceUpdate{}
	this.Data = data
	return &this
}

// NewPreferenceUpdateWithDefaults instantiates a new PreferenceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferenceUpdateWithDefaults() *PreferenceUpdate {
	this := PreferenceUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *PreferenceUpdate) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PreferenceUpdate) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PreferenceUpdate) SetData(v string) {
	o.Data = v
}

func (o PreferenceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePreferenceUpdate struct {
	value *PreferenceUpdate
	isSet bool
}

func (v NullablePreferenceUpdate) Get() *PreferenceUpdate {
	return v.value
}

func (v *NullablePreferenceUpdate) Set(val *PreferenceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferenceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferenceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferenceUpdate(val *PreferenceUpdate) *NullablePreferenceUpdate {
	return &NullablePreferenceUpdate{value: val, isSet: true}
}

func (v NullablePreferenceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferenceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


