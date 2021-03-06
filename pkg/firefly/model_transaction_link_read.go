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

// TransactionLinkRead struct for TransactionLinkRead
type TransactionLinkRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes TransactionLink `json:"attributes"`
	Links ObjectLink `json:"links"`
}

// NewTransactionLinkRead instantiates a new TransactionLinkRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionLinkRead(type_ string, id string, attributes TransactionLink, links ObjectLink) *TransactionLinkRead {
	this := TransactionLinkRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewTransactionLinkReadWithDefaults instantiates a new TransactionLinkRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionLinkReadWithDefaults() *TransactionLinkRead {
	this := TransactionLinkRead{}
	return &this
}

// GetType returns the Type field value
func (o *TransactionLinkRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkRead) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionLinkRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TransactionLinkRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkRead) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionLinkRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TransactionLinkRead) GetAttributes() TransactionLink {
	if o == nil {
		var ret TransactionLink
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkRead) GetAttributesOk() (*TransactionLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TransactionLinkRead) SetAttributes(v TransactionLink) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *TransactionLinkRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TransactionLinkRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TransactionLinkRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o TransactionLinkRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionLinkRead struct {
	value *TransactionLinkRead
	isSet bool
}

func (v NullableTransactionLinkRead) Get() *TransactionLinkRead {
	return v.value
}

func (v *NullableTransactionLinkRead) Set(val *TransactionLinkRead) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionLinkRead) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionLinkRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionLinkRead(val *TransactionLinkRead) *NullableTransactionLinkRead {
	return &NullableTransactionLinkRead{value: val, isSet: true}
}

func (v NullableTransactionLinkRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionLinkRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


