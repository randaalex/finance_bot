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

// RuleGroupUpdate struct for RuleGroupUpdate
type RuleGroupUpdate struct {
	Title *string `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewRuleGroupUpdate instantiates a new RuleGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleGroupUpdate() *RuleGroupUpdate {
	this := RuleGroupUpdate{}
	return &this
}

// NewRuleGroupUpdateWithDefaults instantiates a new RuleGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleGroupUpdateWithDefaults() *RuleGroupUpdate {
	this := RuleGroupUpdate{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RuleGroupUpdate) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroupUpdate) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RuleGroupUpdate) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RuleGroupUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleGroupUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleGroupUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleGroupUpdate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RuleGroupUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RuleGroupUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RuleGroupUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RuleGroupUpdate) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroupUpdate) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RuleGroupUpdate) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *RuleGroupUpdate) SetOrder(v int32) {
	o.Order = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *RuleGroupUpdate) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleGroupUpdate) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *RuleGroupUpdate) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *RuleGroupUpdate) SetActive(v bool) {
	o.Active = &v
}

func (o RuleGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableRuleGroupUpdate struct {
	value *RuleGroupUpdate
	isSet bool
}

func (v NullableRuleGroupUpdate) Get() *RuleGroupUpdate {
	return v.value
}

func (v *NullableRuleGroupUpdate) Set(val *RuleGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleGroupUpdate(val *RuleGroupUpdate) *NullableRuleGroupUpdate {
	return &NullableRuleGroupUpdate{value: val, isSet: true}
}

func (v NullableRuleGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

