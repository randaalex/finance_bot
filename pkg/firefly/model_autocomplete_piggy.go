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

// AutocompletePiggy struct for AutocompletePiggy
type AutocompletePiggy struct {
	Id string `json:"id"`
	// Name of the piggy bank found by an auto-complete search.
	Name string `json:"name"`
	// Currency ID for this piggy bank.
	CurrencyId *int32 `json:"currency_id,omitempty"`
	// Currency code for this piggy bank.
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
}

// NewAutocompletePiggy instantiates a new AutocompletePiggy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutocompletePiggy(id string, name string) *AutocompletePiggy {
	this := AutocompletePiggy{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutocompletePiggyWithDefaults instantiates a new AutocompletePiggy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutocompletePiggyWithDefaults() *AutocompletePiggy {
	this := AutocompletePiggy{}
	return &this
}

// GetId returns the Id field value
func (o *AutocompletePiggy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutocompletePiggy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutocompletePiggy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutocompletePiggy) SetName(v string) {
	o.Name = v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *AutocompletePiggy) GetCurrencyId() int32 {
	if o == nil || o.CurrencyId == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *AutocompletePiggy) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *AutocompletePiggy) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AutocompletePiggy) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AutocompletePiggy) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AutocompletePiggy) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *AutocompletePiggy) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *AutocompletePiggy) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *AutocompletePiggy) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *AutocompletePiggy) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutocompletePiggy) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *AutocompletePiggy) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *AutocompletePiggy) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

func (o AutocompletePiggy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyId != nil {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.CurrencySymbol != nil {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if o.CurrencyDecimalPlaces != nil {
		toSerialize["currency_decimal_places"] = o.CurrencyDecimalPlaces
	}
	return json.Marshal(toSerialize)
}

type NullableAutocompletePiggy struct {
	value *AutocompletePiggy
	isSet bool
}

func (v NullableAutocompletePiggy) Get() *AutocompletePiggy {
	return v.value
}

func (v *NullableAutocompletePiggy) Set(val *AutocompletePiggy) {
	v.value = val
	v.isSet = true
}

func (v NullableAutocompletePiggy) IsSet() bool {
	return v.isSet
}

func (v *NullableAutocompletePiggy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutocompletePiggy(val *AutocompletePiggy) *NullableAutocompletePiggy {
	return &NullableAutocompletePiggy{value: val, isSet: true}
}

func (v NullableAutocompletePiggy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutocompletePiggy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


