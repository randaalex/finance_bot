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

// TagCloudTag struct for TagCloudTag
type TagCloudTag struct {
	// The tag
	Tag *string `json:"tag,omitempty"`
	// ID of the tag
	Id *string `json:"id,omitempty"`
	// The total amount of money related to this tag. There is no currency information available, and this is a basic sum of all amounts added together.
	Size *float64 `json:"size,omitempty"`
	// A number between 0 and 1. 1 is given to the largest tag in the tag cloud, and 0 to the smallest. The rest are given a number between 0 and 1, related to their size in comparison to the largest tag.
	Relative *float64 `json:"relative,omitempty"`
}

// NewTagCloudTag instantiates a new TagCloudTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCloudTag() *TagCloudTag {
	this := TagCloudTag{}
	return &this
}

// NewTagCloudTagWithDefaults instantiates a new TagCloudTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCloudTagWithDefaults() *TagCloudTag {
	this := TagCloudTag{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TagCloudTag) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloudTag) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TagCloudTag) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *TagCloudTag) SetTag(v string) {
	o.Tag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TagCloudTag) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloudTag) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TagCloudTag) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TagCloudTag) SetId(v string) {
	o.Id = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TagCloudTag) GetSize() float64 {
	if o == nil || o.Size == nil {
		var ret float64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloudTag) GetSizeOk() (*float64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TagCloudTag) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given float64 and assigns it to the Size field.
func (o *TagCloudTag) SetSize(v float64) {
	o.Size = &v
}

// GetRelative returns the Relative field value if set, zero value otherwise.
func (o *TagCloudTag) GetRelative() float64 {
	if o == nil || o.Relative == nil {
		var ret float64
		return ret
	}
	return *o.Relative
}

// GetRelativeOk returns a tuple with the Relative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCloudTag) GetRelativeOk() (*float64, bool) {
	if o == nil || o.Relative == nil {
		return nil, false
	}
	return o.Relative, true
}

// HasRelative returns a boolean if a field has been set.
func (o *TagCloudTag) HasRelative() bool {
	if o != nil && o.Relative != nil {
		return true
	}

	return false
}

// SetRelative gets a reference to the given float64 and assigns it to the Relative field.
func (o *TagCloudTag) SetRelative(v float64) {
	o.Relative = &v
}

func (o TagCloudTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Relative != nil {
		toSerialize["relative"] = o.Relative
	}
	return json.Marshal(toSerialize)
}

type NullableTagCloudTag struct {
	value *TagCloudTag
	isSet bool
}

func (v NullableTagCloudTag) Get() *TagCloudTag {
	return v.value
}

func (v *NullableTagCloudTag) Set(val *TagCloudTag) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCloudTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCloudTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCloudTag(val *TagCloudTag) *NullableTagCloudTag {
	return &NullableTagCloudTag{value: val, isSet: true}
}

func (v NullableTagCloudTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCloudTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

