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
	"fmt"
)

// TransactionTypeFilter the model 'TransactionTypeFilter'
type TransactionTypeFilter string

// List of TransactionTypeFilter
const (
	TRANSACTIONTYPEFILTER_ALL TransactionTypeFilter = "all"
	TRANSACTIONTYPEFILTER_WITHDRAWAL TransactionTypeFilter = "withdrawal"
	TRANSACTIONTYPEFILTER_WITHDRAWALS TransactionTypeFilter = "withdrawals"
	TRANSACTIONTYPEFILTER_EXPENSE TransactionTypeFilter = "expense"
	TRANSACTIONTYPEFILTER_DEPOSIT TransactionTypeFilter = "deposit"
	TRANSACTIONTYPEFILTER_DEPOSITS TransactionTypeFilter = "deposits"
	TRANSACTIONTYPEFILTER_INCOME TransactionTypeFilter = "income"
	TRANSACTIONTYPEFILTER_TRANSFER TransactionTypeFilter = "transfer"
	TRANSACTIONTYPEFILTER_TRANSFERS TransactionTypeFilter = "transfers"
	TRANSACTIONTYPEFILTER_OPENING_BALANCE TransactionTypeFilter = "opening_balance"
	TRANSACTIONTYPEFILTER_RECONCILIATION TransactionTypeFilter = "reconciliation"
	TRANSACTIONTYPEFILTER_SPECIAL TransactionTypeFilter = "special"
	TRANSACTIONTYPEFILTER_SPECIALS TransactionTypeFilter = "specials"
	TRANSACTIONTYPEFILTER_DEFAULT TransactionTypeFilter = "default"
)

var allowedTransactionTypeFilterEnumValues = []TransactionTypeFilter{
	"all",
	"withdrawal",
	"withdrawals",
	"expense",
	"deposit",
	"deposits",
	"income",
	"transfer",
	"transfers",
	"opening_balance",
	"reconciliation",
	"special",
	"specials",
	"default",
}

func (v *TransactionTypeFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionTypeFilter(value)
	for _, existing := range allowedTransactionTypeFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionTypeFilter", value)
}

// NewTransactionTypeFilterFromValue returns a pointer to a valid TransactionTypeFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionTypeFilterFromValue(v string) (*TransactionTypeFilter, error) {
	ev := TransactionTypeFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionTypeFilter: valid values are %v", v, allowedTransactionTypeFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionTypeFilter) IsValid() bool {
	for _, existing := range allowedTransactionTypeFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionTypeFilter value
func (v TransactionTypeFilter) Ptr() *TransactionTypeFilter {
	return &v
}

type NullableTransactionTypeFilter struct {
	value *TransactionTypeFilter
	isSet bool
}

func (v NullableTransactionTypeFilter) Get() *TransactionTypeFilter {
	return v.value
}

func (v *NullableTransactionTypeFilter) Set(val *TransactionTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionTypeFilter(val *TransactionTypeFilter) *NullableTransactionTypeFilter {
	return &NullableTransactionTypeFilter{value: val, isSet: true}
}

func (v NullableTransactionTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

