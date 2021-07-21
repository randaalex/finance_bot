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

// AutocompleteTransactionType struct for AutocompleteTransactionType
type AutocompleteTransactionType struct {
	Id string `json:"id"`
	// Type of the object found by an auto-complete search.
	Name string `json:"name"`
	// Name of the object found by an auto-complete search.
	Type string `json:"type"`
}

// NewAutocompleteTransactionType instantiates a new AutocompleteTransactionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompleteTransactionType(id string, name string, type_ string) *AutocompleteTransactionType {
	this := AutocompleteTransactionType{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewAutocompleteTransactionTypeWithDefaults instantiates a new AutocompleteTransactionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompleteTransactionTypeWithDefaults() *AutocompleteTransactionType {
	this := AutocompleteTransactionType{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompleteTransactionType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionType) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompleteTransactionType) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompleteTransactionType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionType) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompleteTransactionType) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AutocompleteTransactionType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AutocompleteTransactionType) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AutocompleteTransactionType) SetType(v string) {
	o.Type = v
}

func (o AutocompleteTransactionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAutocompleteTransactionType struct {
	value *AutocompleteTransactionType
	isSet bool
}

func (v NullableAutocompleteTransactionType) Get() *AutocompleteTransactionType {
	return v.value
}

func (v *NullableAutocompleteTransactionType) Set(val *AutocompleteTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompleteTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompleteTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompleteTransactionType(val *AutocompleteTransactionType) *NullableAutocompleteTransactionType {
	return &NullableAutocompleteTransactionType{value: val, isSet: true}
}

func (v NullableAutocompleteTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompleteTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


