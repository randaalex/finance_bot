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

// TagRead struct for TagRead
type TagRead struct {
	// Immutable value
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes TagModel `json:"attributes"`
	Links ObjectLink `json:"links"`
}

// NewTagRead instantiates a new TagRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRead(type_ string, id string, attributes TagModel, links ObjectLink) *TagRead {
	this := TagRead{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewTagReadWithDefaults instantiates a new TagRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagReadWithDefaults() *TagRead {
	this := TagRead{}
	return &this
}

// GetType returns the Type field value
func (o *TagRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagRead) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TagRead) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *TagRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagRead) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagRead) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TagRead) GetAttributes() TagModel {
	if o == nil {
		var ret TagModel
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TagRead) GetAttributesOk() (*TagModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TagRead) SetAttributes(v TagModel) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *TagRead) GetLinks() ObjectLink {
	if o == nil {
		var ret ObjectLink
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TagRead) GetLinksOk() (*ObjectLink, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TagRead) SetLinks(v ObjectLink) {
	o.Links = v
}

func (o TagRead) MarshalJSON() ([]byte, error) {
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

type NullableTagRead struct {
	value *TagRead
	isSet bool
}

func (v NullableTagRead) Get() *TagRead {
	return v.value
}

func (v *NullableTagRead) Set(val *TagRead) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRead) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRead(val *TagRead) *NullableTagRead {
	return &NullableTagRead{value: val, isSet: true}
}

func (v NullableTagRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


