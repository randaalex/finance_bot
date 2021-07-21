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

// BillPaidDates struct for BillPaidDates
type BillPaidDates struct {
	// Transaction group ID of the paid bill.
	TransactionGroupId *int32 `json:"transaction_group_id,omitempty"`
	// Transaction journal ID of the paid bill.
	TransactionJournalId *int32 `json:"transaction_journal_id,omitempty"`
	// Date the bill was paid.
	Date *string `json:"date,omitempty"`
}

// NewBillPaidDates instantiates a new BillPaidDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillPaidDates() *BillPaidDates {
	this := BillPaidDates{}
	return &this
}

// NewBillPaidDatesWithDefaults instantiates a new BillPaidDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillPaidDatesWithDefaults() *BillPaidDates {
	this := BillPaidDates{}
	return &this
}

// GetTransactionGroupId returns the TransactionGroupId field value if set, zero value otherwise.
func (o *BillPaidDates) GetTransactionGroupId() int32 {
	if o == nil || o.TransactionGroupId == nil {
		var ret int32
		return ret
	}
	return *o.TransactionGroupId
}

// GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillPaidDates) GetTransactionGroupIdOk() (*int32, bool) {
	if o == nil || o.TransactionGroupId == nil {
		return nil, false
	}
	return o.TransactionGroupId, true
}

// HasTransactionGroupId returns a boolean if a field has been set.
func (o *BillPaidDates) HasTransactionGroupId() bool {
	if o != nil && o.TransactionGroupId != nil {
		return true
	}

	return false
}

// SetTransactionGroupId gets a reference to the given int32 and assigns it to the TransactionGroupId field.
func (o *BillPaidDates) SetTransactionGroupId(v int32) {
	o.TransactionGroupId = &v
}

// GetTransactionJournalId returns the TransactionJournalId field value if set, zero value otherwise.
func (o *BillPaidDates) GetTransactionJournalId() int32 {
	if o == nil || o.TransactionJournalId == nil {
		var ret int32
		return ret
	}
	return *o.TransactionJournalId
}

// GetTransactionJournalIdOk returns a tuple with the TransactionJournalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillPaidDates) GetTransactionJournalIdOk() (*int32, bool) {
	if o == nil || o.TransactionJournalId == nil {
		return nil, false
	}
	return o.TransactionJournalId, true
}

// HasTransactionJournalId returns a boolean if a field has been set.
func (o *BillPaidDates) HasTransactionJournalId() bool {
	if o != nil && o.TransactionJournalId != nil {
		return true
	}

	return false
}

// SetTransactionJournalId gets a reference to the given int32 and assigns it to the TransactionJournalId field.
func (o *BillPaidDates) SetTransactionJournalId(v int32) {
	o.TransactionJournalId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BillPaidDates) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillPaidDates) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BillPaidDates) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *BillPaidDates) SetDate(v string) {
	o.Date = &v
}

func (o BillPaidDates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionGroupId != nil {
		toSerialize["transaction_group_id"] = o.TransactionGroupId
	}
	if o.TransactionJournalId != nil {
		toSerialize["transaction_journal_id"] = o.TransactionJournalId
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableBillPaidDates struct {
	value *BillPaidDates
	isSet bool
}

func (v NullableBillPaidDates) Get() *BillPaidDates {
	return v.value
}

func (v *NullableBillPaidDates) Set(val *BillPaidDates) {
	v.value = val
	v.isSet = true
}

func (v NullableBillPaidDates) IsSet() bool {
	return v.isSet
}

func (v *NullableBillPaidDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillPaidDates(val *BillPaidDates) *NullableBillPaidDates {
	return &NullableBillPaidDates{value: val, isSet: true}
}

func (v NullableBillPaidDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillPaidDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

