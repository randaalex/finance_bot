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

// BulkAccountTransactionObject struct for BulkAccountTransactionObject
type BulkAccountTransactionObject struct {
	// The original account, where you wish to transactions away from.
	OriginalAccount float32 `json:"original_account"`
	// The destination account, where the transactions will be moved to.
	DestinationAccount float32 `json:"destination_account"`
}

// NewBulkAccountTransactionObject instantiates a new BulkAccountTransactionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAccountTransactionObject(originalAccount float32, destinationAccount float32) *BulkAccountTransactionObject {
	this := BulkAccountTransactionObject{}
	this.OriginalAccount = originalAccount
	this.DestinationAccount = destinationAccount
	return &this
}

// NewBulkAccountTransactionObjectWithDefaults instantiates a new BulkAccountTransactionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAccountTransactionObjectWithDefaults() *BulkAccountTransactionObject {
	this := BulkAccountTransactionObject{}
	return &this
}

// GetOriginalAccount returns the OriginalAccount field value
func (o *BulkAccountTransactionObject) GetOriginalAccount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OriginalAccount
}

// GetOriginalAccountOk returns a tuple with the OriginalAccount field value
// and a boolean to check if the value has been set.
func (o *BulkAccountTransactionObject) GetOriginalAccountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalAccount, true
}

// SetOriginalAccount sets field value
func (o *BulkAccountTransactionObject) SetOriginalAccount(v float32) {
	o.OriginalAccount = v
}

// GetDestinationAccount returns the DestinationAccount field value
func (o *BulkAccountTransactionObject) GetDestinationAccount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DestinationAccount
}

// GetDestinationAccountOk returns a tuple with the DestinationAccount field value
// and a boolean to check if the value has been set.
func (o *BulkAccountTransactionObject) GetDestinationAccountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationAccount, true
}

// SetDestinationAccount sets field value
func (o *BulkAccountTransactionObject) SetDestinationAccount(v float32) {
	o.DestinationAccount = v
}

func (o BulkAccountTransactionObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["original_account"] = o.OriginalAccount
	}
	if true {
		toSerialize["destination_account"] = o.DestinationAccount
	}
	return json.Marshal(toSerialize)
}

type NullableBulkAccountTransactionObject struct {
	value *BulkAccountTransactionObject
	isSet bool
}

func (v NullableBulkAccountTransactionObject) Get() *BulkAccountTransactionObject {
	return v.value
}

func (v *NullableBulkAccountTransactionObject) Set(val *BulkAccountTransactionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAccountTransactionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAccountTransactionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAccountTransactionObject(val *BulkAccountTransactionObject) *NullableBulkAccountTransactionObject {
	return &NullableBulkAccountTransactionObject{value: val, isSet: true}
}

func (v NullableBulkAccountTransactionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAccountTransactionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


