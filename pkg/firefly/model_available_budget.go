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
	"time"
)

// AvailableBudget struct for AvailableBudget
type AvailableBudget struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Use either currency_id or currency_code.
	CurrencyId *int32 `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code.
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	Amount string `json:"amount"`
	// Start date of the available budget.
	Start string `json:"start"`
	// End date of the available budget.
	End string `json:"end"`
	SpentInBudgets *[]BudgetSpent `json:"spent_in_budgets,omitempty"`
	SpentOutsideBudget *[]BudgetSpent `json:"spent_outside_budget,omitempty"`
}

// NewAvailableBudget instantiates a new AvailableBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableBudget(amount string, start string, end string) *AvailableBudget {
	this := AvailableBudget{}
	this.Amount = amount
	this.Start = start
	this.End = end
	return &this
}

// NewAvailableBudgetWithDefaults instantiates a new AvailableBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableBudgetWithDefaults() *AvailableBudget {
	this := AvailableBudget{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AvailableBudget) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AvailableBudget) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AvailableBudget) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AvailableBudget) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AvailableBudget) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AvailableBudget) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *AvailableBudget) GetCurrencyId() int32 {
	if o == nil || o.CurrencyId == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *AvailableBudget) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *AvailableBudget) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AvailableBudget) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AvailableBudget) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AvailableBudget) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *AvailableBudget) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *AvailableBudget) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *AvailableBudget) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *AvailableBudget) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *AvailableBudget) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *AvailableBudget) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetAmount returns the Amount field value
func (o *AvailableBudget) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AvailableBudget) SetAmount(v string) {
	o.Amount = v
}

// GetStart returns the Start field value
func (o *AvailableBudget) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetStartOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *AvailableBudget) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *AvailableBudget) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetEndOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *AvailableBudget) SetEnd(v string) {
	o.End = v
}

// GetSpentInBudgets returns the SpentInBudgets field value if set, zero value otherwise.
func (o *AvailableBudget) GetSpentInBudgets() []BudgetSpent {
	if o == nil || o.SpentInBudgets == nil {
		var ret []BudgetSpent
		return ret
	}
	return *o.SpentInBudgets
}

// GetSpentInBudgetsOk returns a tuple with the SpentInBudgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetSpentInBudgetsOk() (*[]BudgetSpent, bool) {
	if o == nil || o.SpentInBudgets == nil {
		return nil, false
	}
	return o.SpentInBudgets, true
}

// HasSpentInBudgets returns a boolean if a field has been set.
func (o *AvailableBudget) HasSpentInBudgets() bool {
	if o != nil && o.SpentInBudgets != nil {
		return true
	}

	return false
}

// SetSpentInBudgets gets a reference to the given []BudgetSpent and assigns it to the SpentInBudgets field.
func (o *AvailableBudget) SetSpentInBudgets(v []BudgetSpent) {
	o.SpentInBudgets = &v
}

// GetSpentOutsideBudget returns the SpentOutsideBudget field value if set, zero value otherwise.
func (o *AvailableBudget) GetSpentOutsideBudget() []BudgetSpent {
	if o == nil || o.SpentOutsideBudget == nil {
		var ret []BudgetSpent
		return ret
	}
	return *o.SpentOutsideBudget
}

// GetSpentOutsideBudgetOk returns a tuple with the SpentOutsideBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableBudget) GetSpentOutsideBudgetOk() (*[]BudgetSpent, bool) {
	if o == nil || o.SpentOutsideBudget == nil {
		return nil, false
	}
	return o.SpentOutsideBudget, true
}

// HasSpentOutsideBudget returns a boolean if a field has been set.
func (o *AvailableBudget) HasSpentOutsideBudget() bool {
	if o != nil && o.SpentOutsideBudget != nil {
		return true
	}

	return false
}

// SetSpentOutsideBudget gets a reference to the given []BudgetSpent and assigns it to the SpentOutsideBudget field.
func (o *AvailableBudget) SetSpentOutsideBudget(v []BudgetSpent) {
	o.SpentOutsideBudget = &v
}

func (o AvailableBudget) MarshalJSON() ([]byte, error) {
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
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if o.SpentInBudgets != nil {
		toSerialize["spent_in_budgets"] = o.SpentInBudgets
	}
	if o.SpentOutsideBudget != nil {
		toSerialize["spent_outside_budget"] = o.SpentOutsideBudget
	}
	return json.Marshal(toSerialize)
}

type NullableAvailableBudget struct {
	value *AvailableBudget
	isSet bool
}

func (v NullableAvailableBudget) Get() *AvailableBudget {
	return v.value
}

func (v *NullableAvailableBudget) Set(val *AvailableBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableBudget(val *AvailableBudget) *NullableAvailableBudget {
	return &NullableAvailableBudget{value: val, isSet: true}
}

func (v NullableAvailableBudget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


