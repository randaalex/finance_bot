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
	"fmt"
)

// DataDestroyObject the model 'DataDestroyObject'
type DataDestroyObject string

// List of DataDestroyObject
const (
	DATADESTROYOBJECT_BUDGETS DataDestroyObject = "budgets"
	DATADESTROYOBJECT_BILLS DataDestroyObject = "bills"
	DATADESTROYOBJECT_PIGGY_BANKS DataDestroyObject = "piggy_banks"
	DATADESTROYOBJECT_RULES DataDestroyObject = "rules"
	DATADESTROYOBJECT_RECURRING DataDestroyObject = "recurring"
	DATADESTROYOBJECT_CATEGORIES DataDestroyObject = "categories"
	DATADESTROYOBJECT_TAGS DataDestroyObject = "tags"
	DATADESTROYOBJECT_OBJECT_GROUPS DataDestroyObject = "object_groups"
	DATADESTROYOBJECT_ACCOUNTS DataDestroyObject = "accounts"
	DATADESTROYOBJECT_ASSET_ACCOUNTS DataDestroyObject = "asset_accounts"
	DATADESTROYOBJECT_EXPENSE_ACCOUNTS DataDestroyObject = "expense_accounts"
	DATADESTROYOBJECT_REVENUE_ACCOUNTS DataDestroyObject = "revenue_accounts"
	DATADESTROYOBJECT_LIABILITIES DataDestroyObject = "liabilities"
	DATADESTROYOBJECT_TRANSACTIONS DataDestroyObject = "transactions"
	DATADESTROYOBJECT_WITHDRAWALS DataDestroyObject = "withdrawals"
	DATADESTROYOBJECT_DEPOSITS DataDestroyObject = "deposits"
	DATADESTROYOBJECT_TRANSFERS DataDestroyObject = "transfers"
)

var allowedDataDestroyObjectEnumValues = []DataDestroyObject{
	"budgets",
	"bills",
	"piggy_banks",
	"rules",
	"recurring",
	"categories",
	"tags",
	"object_groups",
	"accounts",
	"asset_accounts",
	"expense_accounts",
	"revenue_accounts",
	"liabilities",
	"transactions",
	"withdrawals",
	"deposits",
	"transfers",
}

func (v *DataDestroyObject) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataDestroyObject(value)
	for _, existing := range allowedDataDestroyObjectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataDestroyObject", value)
}

// NewDataDestroyObjectFromValue returns a pointer to a valid DataDestroyObject
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataDestroyObjectFromValue(v string) (*DataDestroyObject, error) {
	ev := DataDestroyObject(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataDestroyObject: valid values are %v", v, allowedDataDestroyObjectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataDestroyObject) IsValid() bool {
	for _, existing := range allowedDataDestroyObjectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataDestroyObject value
func (v DataDestroyObject) Ptr() *DataDestroyObject {
	return &v
}

type NullableDataDestroyObject struct {
	value *DataDestroyObject
	isSet bool
}

func (v NullableDataDestroyObject) Get() *DataDestroyObject {
	return v.value
}

func (v *NullableDataDestroyObject) Set(val *DataDestroyObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDataDestroyObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDataDestroyObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataDestroyObject(val *DataDestroyObject) *NullableDataDestroyObject {
	return &NullableDataDestroyObject{value: val, isSet: true}
}

func (v NullableDataDestroyObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataDestroyObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

