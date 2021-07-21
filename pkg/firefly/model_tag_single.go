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

// TagSingle struct for TagSingle
type TagSingle struct {
	Data TagRead `json:"data"`
}

// NewTagSingle instantiates a new TagSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagSingle(data TagRead) *TagSingle {
	this := TagSingle{}
	this.Data = data
	return &this
}

// NewTagSingleWithDefaults instantiates a new TagSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagSingleWithDefaults() *TagSingle {
	this := TagSingle{}
	return &this
}

// GetData returns the Data field value
func (o *TagSingle) GetData() TagRead {
	if o == nil {
		var ret TagRead
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TagSingle) GetDataOk() (*TagRead, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TagSingle) SetData(v TagRead) {
	o.Data = v
}

func (o TagSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTagSingle struct {
	value *TagSingle
	isSet bool
}

func (v NullableTagSingle) Get() *TagSingle {
	return v.value
}

func (v *NullableTagSingle) Set(val *TagSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableTagSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableTagSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagSingle(val *TagSingle) *NullableTagSingle {
	return &NullableTagSingle{value: val, isSet: true}
}

func (v NullableTagSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


