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

// ValidationError struct for ValidationError
type ValidationError struct {
	Message *string `json:"message,omitempty"`
	Errors *ValidationErrorErrors `json:"errors,omitempty"`
}

// NewValidationError instantiates a new ValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationError() *ValidationError {
	this := ValidationError{}
	return &this
}

// NewValidationErrorWithDefaults instantiates a new ValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorWithDefaults() *ValidationError {
	this := ValidationError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationError) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidationError) GetErrors() ValidationErrorErrors {
	if o == nil || o.Errors == nil {
		var ret ValidationErrorErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetErrorsOk() (*ValidationErrorErrors, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidationError) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given ValidationErrorErrors and assigns it to the Errors field.
func (o *ValidationError) SetErrors(v ValidationErrorErrors) {
	o.Errors = &v
}

func (o ValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableValidationError struct {
	value *ValidationError
	isSet bool
}

func (v NullableValidationError) Get() *ValidationError {
	return v.value
}

func (v *NullableValidationError) Set(val *ValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationError(val *ValidationError) *NullableValidationError {
	return &NullableValidationError{value: val, isSet: true}
}

func (v NullableValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


