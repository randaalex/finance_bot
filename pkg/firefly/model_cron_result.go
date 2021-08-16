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

// CronResult struct for CronResult
type CronResult struct {
	RecurringTransactions *CronResultRow `json:"recurring_transactions,omitempty"`
	AutoBudgets *CronResultRow `json:"auto_budgets,omitempty"`
	Telemetry *CronResultRow `json:"telemetry,omitempty"`
}

// NewCronResult instantiates a new CronResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronResult() *CronResult {
	this := CronResult{}
	return &this
}

// NewCronResultWithDefaults instantiates a new CronResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronResultWithDefaults() *CronResult {
	this := CronResult{}
	return &this
}

// GetRecurringTransactions returns the RecurringTransactions field value if set, zero value otherwise.
func (o *CronResult) GetRecurringTransactions() CronResultRow {
	if o == nil || o.RecurringTransactions == nil {
		var ret CronResultRow
		return ret
	}
	return *o.RecurringTransactions
}

// GetRecurringTransactionsOk returns a tuple with the RecurringTransactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronResult) GetRecurringTransactionsOk() (*CronResultRow, bool) {
	if o == nil || o.RecurringTransactions == nil {
		return nil, false
	}
	return o.RecurringTransactions, true
}

// HasRecurringTransactions returns a boolean if a field has been set.
func (o *CronResult) HasRecurringTransactions() bool {
	if o != nil && o.RecurringTransactions != nil {
		return true
	}

	return false
}

// SetRecurringTransactions gets a reference to the given CronResultRow and assigns it to the RecurringTransactions field.
func (o *CronResult) SetRecurringTransactions(v CronResultRow) {
	o.RecurringTransactions = &v
}

// GetAutoBudgets returns the AutoBudgets field value if set, zero value otherwise.
func (o *CronResult) GetAutoBudgets() CronResultRow {
	if o == nil || o.AutoBudgets == nil {
		var ret CronResultRow
		return ret
	}
	return *o.AutoBudgets
}

// GetAutoBudgetsOk returns a tuple with the AutoBudgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronResult) GetAutoBudgetsOk() (*CronResultRow, bool) {
	if o == nil || o.AutoBudgets == nil {
		return nil, false
	}
	return o.AutoBudgets, true
}

// HasAutoBudgets returns a boolean if a field has been set.
func (o *CronResult) HasAutoBudgets() bool {
	if o != nil && o.AutoBudgets != nil {
		return true
	}

	return false
}

// SetAutoBudgets gets a reference to the given CronResultRow and assigns it to the AutoBudgets field.
func (o *CronResult) SetAutoBudgets(v CronResultRow) {
	o.AutoBudgets = &v
}

// GetTelemetry returns the Telemetry field value if set, zero value otherwise.
func (o *CronResult) GetTelemetry() CronResultRow {
	if o == nil || o.Telemetry == nil {
		var ret CronResultRow
		return ret
	}
	return *o.Telemetry
}

// GetTelemetryOk returns a tuple with the Telemetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronResult) GetTelemetryOk() (*CronResultRow, bool) {
	if o == nil || o.Telemetry == nil {
		return nil, false
	}
	return o.Telemetry, true
}

// HasTelemetry returns a boolean if a field has been set.
func (o *CronResult) HasTelemetry() bool {
	if o != nil && o.Telemetry != nil {
		return true
	}

	return false
}

// SetTelemetry gets a reference to the given CronResultRow and assigns it to the Telemetry field.
func (o *CronResult) SetTelemetry(v CronResultRow) {
	o.Telemetry = &v
}

func (o CronResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecurringTransactions != nil {
		toSerialize["recurring_transactions"] = o.RecurringTransactions
	}
	if o.AutoBudgets != nil {
		toSerialize["auto_budgets"] = o.AutoBudgets
	}
	if o.Telemetry != nil {
		toSerialize["telemetry"] = o.Telemetry
	}
	return json.Marshal(toSerialize)
}

type NullableCronResult struct {
	value *CronResult
	isSet bool
}

func (v NullableCronResult) Get() *CronResult {
	return v.value
}

func (v *NullableCronResult) Set(val *CronResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCronResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCronResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronResult(val *CronResult) *NullableCronResult {
	return &NullableCronResult{value: val, isSet: true}
}

func (v NullableCronResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


