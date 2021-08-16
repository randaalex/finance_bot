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
	"time"
)

// BudgetLimit struct for BudgetLimit
type BudgetLimit struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyId *string `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code. Defaults to the user's default currency.
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	// The budget ID of the associated budget.
	BudgetId string `json:"budget_id"`
	// Start date of the budget limit.
	Start time.Time `json:"start"`
	// Period of the budget limit. Only used when auto-generated by auto-budget.
	Period NullableString `json:"period,omitempty"`
	// End date of the budget limit.
	End time.Time `json:"end"`
	Amount string `json:"amount"`
	Spent NullableString `json:"spent,omitempty"`
}

// NewBudgetLimit instantiates a new BudgetLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetLimit(budgetId string, start time.Time, end time.Time, amount string) *BudgetLimit {
	this := BudgetLimit{}
	this.BudgetId = budgetId
	this.Start = start
	this.End = end
	this.Amount = amount
	return &this
}

// NewBudgetLimitWithDefaults instantiates a new BudgetLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetLimitWithDefaults() *BudgetLimit {
	this := BudgetLimit{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BudgetLimit) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BudgetLimit) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BudgetLimit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BudgetLimit) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BudgetLimit) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BudgetLimit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *BudgetLimit) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *BudgetLimit) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *BudgetLimit) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BudgetLimit) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BudgetLimit) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BudgetLimit) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *BudgetLimit) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *BudgetLimit) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *BudgetLimit) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *BudgetLimit) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *BudgetLimit) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *BudgetLimit) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetBudgetId returns the BudgetId field value
func (o *BudgetLimit) GetBudgetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetBudgetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BudgetId, true
}

// SetBudgetId sets field value
func (o *BudgetLimit) SetBudgetId(v string) {
	o.BudgetId = v
}

// GetStart returns the Start field value
func (o *BudgetLimit) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetStartOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *BudgetLimit) SetStart(v time.Time) {
	o.Start = v
}

// GetPeriod returns the Period field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetLimit) GetPeriod() string {
	if o == nil || o.Period.Get() == nil {
		var ret string
		return ret
	}
	return *o.Period.Get()
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetLimit) GetPeriodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Period.Get(), o.Period.IsSet()
}

// HasPeriod returns a boolean if a field has been set.
func (o *BudgetLimit) HasPeriod() bool {
	if o != nil && o.Period.IsSet() {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given NullableString and assigns it to the Period field.
func (o *BudgetLimit) SetPeriod(v string) {
	o.Period.Set(&v)
}
// SetPeriodNil sets the value for Period to be an explicit nil
func (o *BudgetLimit) SetPeriodNil() {
	o.Period.Set(nil)
}

// UnsetPeriod ensures that no value is present for Period, not even an explicit nil
func (o *BudgetLimit) UnsetPeriod() {
	o.Period.Unset()
}

// GetEnd returns the End field value
func (o *BudgetLimit) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetEndOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *BudgetLimit) SetEnd(v time.Time) {
	o.End = v
}

// GetAmount returns the Amount field value
func (o *BudgetLimit) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BudgetLimit) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BudgetLimit) SetAmount(v string) {
	o.Amount = v
}

// GetSpent returns the Spent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BudgetLimit) GetSpent() string {
	if o == nil || o.Spent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Spent.Get()
}

// GetSpentOk returns a tuple with the Spent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BudgetLimit) GetSpentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Spent.Get(), o.Spent.IsSet()
}

// HasSpent returns a boolean if a field has been set.
func (o *BudgetLimit) HasSpent() bool {
	if o != nil && o.Spent.IsSet() {
		return true
	}

	return false
}

// SetSpent gets a reference to the given NullableString and assigns it to the Spent field.
func (o *BudgetLimit) SetSpent(v string) {
	o.Spent.Set(&v)
}
// SetSpentNil sets the value for Spent to be an explicit nil
func (o *BudgetLimit) SetSpentNil() {
	o.Spent.Set(nil)
}

// UnsetSpent ensures that no value is present for Spent, not even an explicit nil
func (o *BudgetLimit) UnsetSpent() {
	o.Spent.Unset()
}

func (o BudgetLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CurrencyId != nil {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.CurrencySymbol != nil {
		toSerialize["currency_symbol"] = o.CurrencySymbol
	}
	if o.CurrencyDecimalPlaces != nil {
		toSerialize["currency_decimal_places"] = o.CurrencyDecimalPlaces
	}
	if true {
		toSerialize["budget_id"] = o.BudgetId
	}
	if true {
		toSerialize["start"] = o.Start
	}
	if o.Period.IsSet() {
		toSerialize["period"] = o.Period.Get()
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Spent.IsSet() {
		toSerialize["spent"] = o.Spent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBudgetLimit struct {
	value *BudgetLimit
	isSet bool
}

func (v NullableBudgetLimit) Get() *BudgetLimit {
	return v.value
}

func (v *NullableBudgetLimit) Set(val *BudgetLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetLimit(val *BudgetLimit) *NullableBudgetLimit {
	return &NullableBudgetLimit{value: val, isSet: true}
}

func (v NullableBudgetLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


