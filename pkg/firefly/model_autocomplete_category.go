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

// AutocompleteCategory struct for AutocompleteCategory
type AutocompleteCategory struct {
	Id string `json:"id"`
	// Name of the category found by an auto-complete search.
	Name string `json:"name"`
}

// NewAutocompleteCategory instantiates a new AutocompleteCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteCategory(id string, name string) *AutocompleteCategory {
	this := AutocompleteCategory{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutocompleteCategoryWithDefaults instantiates a new AutocompleteCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteCategoryWithDefaults() *AutocompleteCategory {
	this := AutocompleteCategory{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteCategory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCategory) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteCategory) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteCategory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteCategory) SetName(v string) {
	o.Name = v
}

func (o AutocompleteCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAutocompleteCategory struct {
	value *AutocompleteCategory
	isSet bool
}

func (v NullableAutocompleteCategory) Get() *AutocompleteCategory {
	return v.value
}

func (v *NullableAutocompleteCategory) Set(val *AutocompleteCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteCategory(val *AutocompleteCategory) *NullableAutocompleteCategory {
	return &NullableAutocompleteCategory{value: val, isSet: true}
}

func (v NullableAutocompleteCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


