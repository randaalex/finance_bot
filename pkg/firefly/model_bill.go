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

// Bill struct for Bill
type Bill struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Use either currency_id or currency_code
	CurrencyId *int32 `json:"currency_id,omitempty"`
	// Use either currency_id or currency_code
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	Name string `json:"name"`
	AmountMin string `json:"amount_min"`
	AmountMax string `json:"amount_max"`
	Date string `json:"date"`
	// How often the bill must be paid.
	RepeatFreq string `json:"repeat_freq"`
	// How often the bill must be skipped. 1 means a bi-monthly bill.
	Skip *int32 `json:"skip,omitempty"`
	// If the bill is active.
	Active *bool `json:"active,omitempty"`
	Notes *string `json:"notes,omitempty"`
	// When the bill is expected to be due.
	NextExpectedMatch *string `json:"next_expected_match,omitempty"`
	// The group ID of the group this object is part of. NULL if no group.
	ObjectGroupId *int32 `json:"object_group_id,omitempty"`
	// The order of the group. At least 1, for the highest sorting.
	ObjectGroupOrder *int32 `json:"object_group_order,omitempty"`
	// The name of the group. NULL if no group.
	ObjectGroupTitle *string `json:"object_group_title,omitempty"`
	// Array of future dates when the bill is expected to be paid. Autogenerated.
	PayDates *[]string `json:"pay_dates,omitempty"`
	// Array of past transactions when the bill was paid.
	PaidDates *[]BillPaidDates `json:"paid_dates,omitempty"`
}

// NewBill instantiates a new Bill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBill(name string, amountMin string, amountMax string, date string, repeatFreq string) *Bill {
	this := Bill{}
	this.Name = name
	this.AmountMin = amountMin
	this.AmountMax = amountMax
	this.Date = date
	this.RepeatFreq = repeatFreq
	return &this
}

// NewBillWithDefaults instantiates a new Bill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillWithDefaults() *Bill {
	this := Bill{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Bill) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Bill) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Bill) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Bill) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Bill) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Bill) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *Bill) GetCurrencyId() int32 {
	if o == nil || o.CurrencyId == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *Bill) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *Bill) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Bill) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Bill) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Bill) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *Bill) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *Bill) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *Bill) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *Bill) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *Bill) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *Bill) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetName returns the Name field value
func (o *Bill) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bill) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bill) SetName(v string) {
	o.Name = v
}

// GetAmountMin returns the AmountMin field value
func (o *Bill) GetAmountMin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountMin
}

// GetAmountMinOk returns a tuple with the AmountMin field value
// and a boolean to check if the value has been set.
func (o *Bill) GetAmountMinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AmountMin, true
}

// SetAmountMin sets field value
func (o *Bill) SetAmountMin(v string) {
	o.AmountMin = v
}

// GetAmountMax returns the AmountMax field value
func (o *Bill) GetAmountMax() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountMax
}

// GetAmountMaxOk returns a tuple with the AmountMax field value
// and a boolean to check if the value has been set.
func (o *Bill) GetAmountMaxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AmountMax, true
}

// SetAmountMax sets field value
func (o *Bill) SetAmountMax(v string) {
	o.AmountMax = v
}

// GetDate returns the Date field value
func (o *Bill) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Bill) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Bill) SetDate(v string) {
	o.Date = v
}

// GetRepeatFreq returns the RepeatFreq field value
func (o *Bill) GetRepeatFreq() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepeatFreq
}

// GetRepeatFreqOk returns a tuple with the RepeatFreq field value
// and a boolean to check if the value has been set.
func (o *Bill) GetRepeatFreqOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepeatFreq, true
}

// SetRepeatFreq sets field value
func (o *Bill) SetRepeatFreq(v string) {
	o.RepeatFreq = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *Bill) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *Bill) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *Bill) SetSkip(v int32) {
	o.Skip = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Bill) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Bill) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Bill) SetActive(v bool) {
	o.Active = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Bill) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Bill) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Bill) SetNotes(v string) {
	o.Notes = &v
}

// GetNextExpectedMatch returns the NextExpectedMatch field value if set, zero value otherwise.
func (o *Bill) GetNextExpectedMatch() string {
	if o == nil || o.NextExpectedMatch == nil {
		var ret string
		return ret
	}
	return *o.NextExpectedMatch
}

// GetNextExpectedMatchOk returns a tuple with the NextExpectedMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetNextExpectedMatchOk() (*string, bool) {
	if o == nil || o.NextExpectedMatch == nil {
		return nil, false
	}
	return o.NextExpectedMatch, true
}

// HasNextExpectedMatch returns a boolean if a field has been set.
func (o *Bill) HasNextExpectedMatch() bool {
	if o != nil && o.NextExpectedMatch != nil {
		return true
	}

	return false
}

// SetNextExpectedMatch gets a reference to the given string and assigns it to the NextExpectedMatch field.
func (o *Bill) SetNextExpectedMatch(v string) {
	o.NextExpectedMatch = &v
}

// GetObjectGroupId returns the ObjectGroupId field value if set, zero value otherwise.
func (o *Bill) GetObjectGroupId() int32 {
	if o == nil || o.ObjectGroupId == nil {
		var ret int32
		return ret
	}
	return *o.ObjectGroupId
}

// GetObjectGroupIdOk returns a tuple with the ObjectGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetObjectGroupIdOk() (*int32, bool) {
	if o == nil || o.ObjectGroupId == nil {
		return nil, false
	}
	return o.ObjectGroupId, true
}

// HasObjectGroupId returns a boolean if a field has been set.
func (o *Bill) HasObjectGroupId() bool {
	if o != nil && o.ObjectGroupId != nil {
		return true
	}

	return false
}

// SetObjectGroupId gets a reference to the given int32 and assigns it to the ObjectGroupId field.
func (o *Bill) SetObjectGroupId(v int32) {
	o.ObjectGroupId = &v
}

// GetObjectGroupOrder returns the ObjectGroupOrder field value if set, zero value otherwise.
func (o *Bill) GetObjectGroupOrder() int32 {
	if o == nil || o.ObjectGroupOrder == nil {
		var ret int32
		return ret
	}
	return *o.ObjectGroupOrder
}

// GetObjectGroupOrderOk returns a tuple with the ObjectGroupOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetObjectGroupOrderOk() (*int32, bool) {
	if o == nil || o.ObjectGroupOrder == nil {
		return nil, false
	}
	return o.ObjectGroupOrder, true
}

// HasObjectGroupOrder returns a boolean if a field has been set.
func (o *Bill) HasObjectGroupOrder() bool {
	if o != nil && o.ObjectGroupOrder != nil {
		return true
	}

	return false
}

// SetObjectGroupOrder gets a reference to the given int32 and assigns it to the ObjectGroupOrder field.
func (o *Bill) SetObjectGroupOrder(v int32) {
	o.ObjectGroupOrder = &v
}

// GetObjectGroupTitle returns the ObjectGroupTitle field value if set, zero value otherwise.
func (o *Bill) GetObjectGroupTitle() string {
	if o == nil || o.ObjectGroupTitle == nil {
		var ret string
		return ret
	}
	return *o.ObjectGroupTitle
}

// GetObjectGroupTitleOk returns a tuple with the ObjectGroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetObjectGroupTitleOk() (*string, bool) {
	if o == nil || o.ObjectGroupTitle == nil {
		return nil, false
	}
	return o.ObjectGroupTitle, true
}

// HasObjectGroupTitle returns a boolean if a field has been set.
func (o *Bill) HasObjectGroupTitle() bool {
	if o != nil && o.ObjectGroupTitle != nil {
		return true
	}

	return false
}

// SetObjectGroupTitle gets a reference to the given string and assigns it to the ObjectGroupTitle field.
func (o *Bill) SetObjectGroupTitle(v string) {
	o.ObjectGroupTitle = &v
}

// GetPayDates returns the PayDates field value if set, zero value otherwise.
func (o *Bill) GetPayDates() []string {
	if o == nil || o.PayDates == nil {
		var ret []string
		return ret
	}
	return *o.PayDates
}

// GetPayDatesOk returns a tuple with the PayDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetPayDatesOk() (*[]string, bool) {
	if o == nil || o.PayDates == nil {
		return nil, false
	}
	return o.PayDates, true
}

// HasPayDates returns a boolean if a field has been set.
func (o *Bill) HasPayDates() bool {
	if o != nil && o.PayDates != nil {
		return true
	}

	return false
}

// SetPayDates gets a reference to the given []string and assigns it to the PayDates field.
func (o *Bill) SetPayDates(v []string) {
	o.PayDates = &v
}

// GetPaidDates returns the PaidDates field value if set, zero value otherwise.
func (o *Bill) GetPaidDates() []BillPaidDates {
	if o == nil || o.PaidDates == nil {
		var ret []BillPaidDates
		return ret
	}
	return *o.PaidDates
}

// GetPaidDatesOk returns a tuple with the PaidDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bill) GetPaidDatesOk() (*[]BillPaidDates, bool) {
	if o == nil || o.PaidDates == nil {
		return nil, false
	}
	return o.PaidDates, true
}

// HasPaidDates returns a boolean if a field has been set.
func (o *Bill) HasPaidDates() bool {
	if o != nil && o.PaidDates != nil {
		return true
	}

	return false
}

// SetPaidDates gets a reference to the given []BillPaidDates and assigns it to the PaidDates field.
func (o *Bill) SetPaidDates(v []BillPaidDates) {
	o.PaidDates = &v
}

func (o Bill) MarshalJSON() ([]byte, error) {
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
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["amount_min"] = o.AmountMin
	}
	if true {
		toSerialize["amount_max"] = o.AmountMax
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["repeat_freq"] = o.RepeatFreq
	}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.NextExpectedMatch != nil {
		toSerialize["next_expected_match"] = o.NextExpectedMatch
	}
	if o.ObjectGroupId != nil {
		toSerialize["object_group_id"] = o.ObjectGroupId
	}
	if o.ObjectGroupOrder != nil {
		toSerialize["object_group_order"] = o.ObjectGroupOrder
	}
	if o.ObjectGroupTitle != nil {
		toSerialize["object_group_title"] = o.ObjectGroupTitle
	}
	if o.PayDates != nil {
		toSerialize["pay_dates"] = o.PayDates
	}
	if o.PaidDates != nil {
		toSerialize["paid_dates"] = o.PaidDates
	}
	return json.Marshal(toSerialize)
}

type NullableBill struct {
	value *Bill
	isSet bool
}

func (v NullableBill) Get() *Bill {
	return v.value
}

func (v *NullableBill) Set(val *Bill) {
	v.value = val
	v.isSet = true
}

func (v NullableBill) IsSet() bool {
	return v.isSet
}

func (v *NullableBill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBill(val *Bill) *NullableBill {
	return &NullableBill{value: val, isSet: true}
}

func (v NullableBill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


