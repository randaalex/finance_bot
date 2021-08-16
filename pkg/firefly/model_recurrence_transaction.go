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

// RecurrenceTransaction struct for RecurrenceTransaction
type RecurrenceTransaction struct {
	Description string `json:"description"`
	// Amount of the transaction.
	Amount string `json:"amount"`
	// Foreign amount of the transaction.
	ForeignAmount NullableString `json:"foreign_amount,omitempty"`
	// Submit either a currency_id or a currency_code.
	CurrencyId *string `json:"currency_id,omitempty"`
	// Submit either a currency_id or a currency_code.
	CurrencyCode *string `json:"currency_code,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
	// Number of decimals in the currency
	CurrencyDecimalPlaces *int32 `json:"currency_decimal_places,omitempty"`
	// Submit either a foreign_currency_id or a foreign_currency_code, or neither.
	ForeignCurrencyId NullableString `json:"foreign_currency_id,omitempty"`
	// Submit either a foreign_currency_id or a foreign_currency_code, or neither.
	ForeignCurrencyCode NullableString `json:"foreign_currency_code,omitempty"`
	ForeignCurrencySymbol NullableString `json:"foreign_currency_symbol,omitempty"`
	// Number of decimals in the currency
	ForeignCurrencyDecimalPlaces NullableInt32 `json:"foreign_currency_decimal_places,omitempty"`
	// The budget ID for this transaction.
	BudgetId *string `json:"budget_id,omitempty"`
	// The name of the budget to be used. If the budget name is unknown, the ID will be used or the value will be ignored.
	BudgetName NullableString `json:"budget_name,omitempty"`
	// Category ID for this transaction.
	CategoryId *string `json:"category_id,omitempty"`
	// Category name for this transaction.
	CategoryName *string `json:"category_name,omitempty"`
	// ID of the source account. Submit either this or source_name.
	SourceId *string `json:"source_id,omitempty"`
	// Name of the source account. Submit either this or source_id.
	SourceName *string `json:"source_name,omitempty"`
	SourceIban NullableString `json:"source_iban,omitempty"`
	SourceType *AccountTypeProperty `json:"source_type,omitempty"`
	// ID of the destination account. Submit either this or destination_name.
	DestinationId *string `json:"destination_id,omitempty"`
	// Name of the destination account. Submit either this or destination_id.
	DestinationName *string `json:"destination_name,omitempty"`
	DestinationIban NullableString `json:"destination_iban,omitempty"`
	DestinationType *AccountTypeProperty `json:"destination_type,omitempty"`
	// Array of tags.
	Tags []string `json:"tags,omitempty"`
	// Optional. Use either this or the piggy_bank_name
	PiggyBankId NullableString `json:"piggy_bank_id,omitempty"`
	// Optional. Use either this or the piggy_bank_id
	PiggyBankName NullableString `json:"piggy_bank_name,omitempty"`
}

// NewRecurrenceTransaction instantiates a new RecurrenceTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrenceTransaction(description string, amount string) *RecurrenceTransaction {
	this := RecurrenceTransaction{}
	this.Description = description
	this.Amount = amount
	return &this
}

// NewRecurrenceTransactionWithDefaults instantiates a new RecurrenceTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceTransactionWithDefaults() *RecurrenceTransaction {
	this := RecurrenceTransaction{}
	return &this
}

// GetDescription returns the Description field value
func (o *RecurrenceTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RecurrenceTransaction) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *RecurrenceTransaction) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RecurrenceTransaction) SetAmount(v string) {
	o.Amount = v
}

// GetForeignAmount returns the ForeignAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetForeignAmount() string {
	if o == nil || o.ForeignAmount.Get() == nil {
		var ret string
		return ret
	}
	return *o.ForeignAmount.Get()
}

// GetForeignAmountOk returns a tuple with the ForeignAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetForeignAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForeignAmount.Get(), o.ForeignAmount.IsSet()
}

// HasForeignAmount returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasForeignAmount() bool {
	if o != nil && o.ForeignAmount.IsSet() {
		return true
	}

	return false
}

// SetForeignAmount gets a reference to the given NullableString and assigns it to the ForeignAmount field.
func (o *RecurrenceTransaction) SetForeignAmount(v string) {
	o.ForeignAmount.Set(&v)
}
// SetForeignAmountNil sets the value for ForeignAmount to be an explicit nil
func (o *RecurrenceTransaction) SetForeignAmountNil() {
	o.ForeignAmount.Set(nil)
}

// UnsetForeignAmount ensures that no value is present for ForeignAmount, not even an explicit nil
func (o *RecurrenceTransaction) UnsetForeignAmount() {
	o.ForeignAmount.Unset()
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCurrencyId() string {
	if o == nil || o.CurrencyId == nil {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCurrencyIdOk() (*string, bool) {
	if o == nil || o.CurrencyId == nil {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCurrencyId() bool {
	if o != nil && o.CurrencyId != nil {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *RecurrenceTransaction) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *RecurrenceTransaction) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *RecurrenceTransaction) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCurrencyDecimalPlaces() int32 {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.CurrencyDecimalPlaces
}

// GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.CurrencyDecimalPlaces == nil {
		return nil, false
	}
	return o.CurrencyDecimalPlaces, true
}

// HasCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCurrencyDecimalPlaces() bool {
	if o != nil && o.CurrencyDecimalPlaces != nil {
		return true
	}

	return false
}

// SetCurrencyDecimalPlaces gets a reference to the given int32 and assigns it to the CurrencyDecimalPlaces field.
func (o *RecurrenceTransaction) SetCurrencyDecimalPlaces(v int32) {
	o.CurrencyDecimalPlaces = &v
}

// GetForeignCurrencyId returns the ForeignCurrencyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetForeignCurrencyId() string {
	if o == nil || o.ForeignCurrencyId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ForeignCurrencyId.Get()
}

// GetForeignCurrencyIdOk returns a tuple with the ForeignCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetForeignCurrencyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForeignCurrencyId.Get(), o.ForeignCurrencyId.IsSet()
}

// HasForeignCurrencyId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasForeignCurrencyId() bool {
	if o != nil && o.ForeignCurrencyId.IsSet() {
		return true
	}

	return false
}

// SetForeignCurrencyId gets a reference to the given NullableString and assigns it to the ForeignCurrencyId field.
func (o *RecurrenceTransaction) SetForeignCurrencyId(v string) {
	o.ForeignCurrencyId.Set(&v)
}
// SetForeignCurrencyIdNil sets the value for ForeignCurrencyId to be an explicit nil
func (o *RecurrenceTransaction) SetForeignCurrencyIdNil() {
	o.ForeignCurrencyId.Set(nil)
}

// UnsetForeignCurrencyId ensures that no value is present for ForeignCurrencyId, not even an explicit nil
func (o *RecurrenceTransaction) UnsetForeignCurrencyId() {
	o.ForeignCurrencyId.Unset()
}

// GetForeignCurrencyCode returns the ForeignCurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetForeignCurrencyCode() string {
	if o == nil || o.ForeignCurrencyCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ForeignCurrencyCode.Get()
}

// GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetForeignCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForeignCurrencyCode.Get(), o.ForeignCurrencyCode.IsSet()
}

// HasForeignCurrencyCode returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasForeignCurrencyCode() bool {
	if o != nil && o.ForeignCurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetForeignCurrencyCode gets a reference to the given NullableString and assigns it to the ForeignCurrencyCode field.
func (o *RecurrenceTransaction) SetForeignCurrencyCode(v string) {
	o.ForeignCurrencyCode.Set(&v)
}
// SetForeignCurrencyCodeNil sets the value for ForeignCurrencyCode to be an explicit nil
func (o *RecurrenceTransaction) SetForeignCurrencyCodeNil() {
	o.ForeignCurrencyCode.Set(nil)
}

// UnsetForeignCurrencyCode ensures that no value is present for ForeignCurrencyCode, not even an explicit nil
func (o *RecurrenceTransaction) UnsetForeignCurrencyCode() {
	o.ForeignCurrencyCode.Unset()
}

// GetForeignCurrencySymbol returns the ForeignCurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetForeignCurrencySymbol() string {
	if o == nil || o.ForeignCurrencySymbol.Get() == nil {
		var ret string
		return ret
	}
	return *o.ForeignCurrencySymbol.Get()
}

// GetForeignCurrencySymbolOk returns a tuple with the ForeignCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetForeignCurrencySymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForeignCurrencySymbol.Get(), o.ForeignCurrencySymbol.IsSet()
}

// HasForeignCurrencySymbol returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasForeignCurrencySymbol() bool {
	if o != nil && o.ForeignCurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetForeignCurrencySymbol gets a reference to the given NullableString and assigns it to the ForeignCurrencySymbol field.
func (o *RecurrenceTransaction) SetForeignCurrencySymbol(v string) {
	o.ForeignCurrencySymbol.Set(&v)
}
// SetForeignCurrencySymbolNil sets the value for ForeignCurrencySymbol to be an explicit nil
func (o *RecurrenceTransaction) SetForeignCurrencySymbolNil() {
	o.ForeignCurrencySymbol.Set(nil)
}

// UnsetForeignCurrencySymbol ensures that no value is present for ForeignCurrencySymbol, not even an explicit nil
func (o *RecurrenceTransaction) UnsetForeignCurrencySymbol() {
	o.ForeignCurrencySymbol.Unset()
}

// GetForeignCurrencyDecimalPlaces returns the ForeignCurrencyDecimalPlaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetForeignCurrencyDecimalPlaces() int32 {
	if o == nil || o.ForeignCurrencyDecimalPlaces.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ForeignCurrencyDecimalPlaces.Get()
}

// GetForeignCurrencyDecimalPlacesOk returns a tuple with the ForeignCurrencyDecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetForeignCurrencyDecimalPlacesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ForeignCurrencyDecimalPlaces.Get(), o.ForeignCurrencyDecimalPlaces.IsSet()
}

// HasForeignCurrencyDecimalPlaces returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasForeignCurrencyDecimalPlaces() bool {
	if o != nil && o.ForeignCurrencyDecimalPlaces.IsSet() {
		return true
	}

	return false
}

// SetForeignCurrencyDecimalPlaces gets a reference to the given NullableInt32 and assigns it to the ForeignCurrencyDecimalPlaces field.
func (o *RecurrenceTransaction) SetForeignCurrencyDecimalPlaces(v int32) {
	o.ForeignCurrencyDecimalPlaces.Set(&v)
}
// SetForeignCurrencyDecimalPlacesNil sets the value for ForeignCurrencyDecimalPlaces to be an explicit nil
func (o *RecurrenceTransaction) SetForeignCurrencyDecimalPlacesNil() {
	o.ForeignCurrencyDecimalPlaces.Set(nil)
}

// UnsetForeignCurrencyDecimalPlaces ensures that no value is present for ForeignCurrencyDecimalPlaces, not even an explicit nil
func (o *RecurrenceTransaction) UnsetForeignCurrencyDecimalPlaces() {
	o.ForeignCurrencyDecimalPlaces.Unset()
}

// GetBudgetId returns the BudgetId field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetBudgetId() string {
	if o == nil || o.BudgetId == nil {
		var ret string
		return ret
	}
	return *o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetBudgetIdOk() (*string, bool) {
	if o == nil || o.BudgetId == nil {
		return nil, false
	}
	return o.BudgetId, true
}

// HasBudgetId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasBudgetId() bool {
	if o != nil && o.BudgetId != nil {
		return true
	}

	return false
}

// SetBudgetId gets a reference to the given string and assigns it to the BudgetId field.
func (o *RecurrenceTransaction) SetBudgetId(v string) {
	o.BudgetId = &v
}

// GetBudgetName returns the BudgetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetBudgetName() string {
	if o == nil || o.BudgetName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BudgetName.Get()
}

// GetBudgetNameOk returns a tuple with the BudgetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetBudgetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BudgetName.Get(), o.BudgetName.IsSet()
}

// HasBudgetName returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasBudgetName() bool {
	if o != nil && o.BudgetName.IsSet() {
		return true
	}

	return false
}

// SetBudgetName gets a reference to the given NullableString and assigns it to the BudgetName field.
func (o *RecurrenceTransaction) SetBudgetName(v string) {
	o.BudgetName.Set(&v)
}
// SetBudgetNameNil sets the value for BudgetName to be an explicit nil
func (o *RecurrenceTransaction) SetBudgetNameNil() {
	o.BudgetName.Set(nil)
}

// UnsetBudgetName ensures that no value is present for BudgetName, not even an explicit nil
func (o *RecurrenceTransaction) UnsetBudgetName() {
	o.BudgetName.Unset()
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCategoryId() string {
	if o == nil || o.CategoryId == nil {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCategoryIdOk() (*string, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *RecurrenceTransaction) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetCategoryName() string {
	if o == nil || o.CategoryName == nil {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetCategoryNameOk() (*string, bool) {
	if o == nil || o.CategoryName == nil {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasCategoryName() bool {
	if o != nil && o.CategoryName != nil {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *RecurrenceTransaction) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *RecurrenceTransaction) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetSourceName() string {
	if o == nil || o.SourceName == nil {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetSourceNameOk() (*string, bool) {
	if o == nil || o.SourceName == nil {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasSourceName() bool {
	if o != nil && o.SourceName != nil {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *RecurrenceTransaction) SetSourceName(v string) {
	o.SourceName = &v
}

// GetSourceIban returns the SourceIban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetSourceIban() string {
	if o == nil || o.SourceIban.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceIban.Get()
}

// GetSourceIbanOk returns a tuple with the SourceIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetSourceIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceIban.Get(), o.SourceIban.IsSet()
}

// HasSourceIban returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasSourceIban() bool {
	if o != nil && o.SourceIban.IsSet() {
		return true
	}

	return false
}

// SetSourceIban gets a reference to the given NullableString and assigns it to the SourceIban field.
func (o *RecurrenceTransaction) SetSourceIban(v string) {
	o.SourceIban.Set(&v)
}
// SetSourceIbanNil sets the value for SourceIban to be an explicit nil
func (o *RecurrenceTransaction) SetSourceIbanNil() {
	o.SourceIban.Set(nil)
}

// UnsetSourceIban ensures that no value is present for SourceIban, not even an explicit nil
func (o *RecurrenceTransaction) UnsetSourceIban() {
	o.SourceIban.Unset()
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetSourceType() AccountTypeProperty {
	if o == nil || o.SourceType == nil {
		var ret AccountTypeProperty
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetSourceTypeOk() (*AccountTypeProperty, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given AccountTypeProperty and assigns it to the SourceType field.
func (o *RecurrenceTransaction) SetSourceType(v AccountTypeProperty) {
	o.SourceType = &v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *RecurrenceTransaction) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetDestinationName returns the DestinationName field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetDestinationName() string {
	if o == nil || o.DestinationName == nil {
		var ret string
		return ret
	}
	return *o.DestinationName
}

// GetDestinationNameOk returns a tuple with the DestinationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetDestinationNameOk() (*string, bool) {
	if o == nil || o.DestinationName == nil {
		return nil, false
	}
	return o.DestinationName, true
}

// HasDestinationName returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasDestinationName() bool {
	if o != nil && o.DestinationName != nil {
		return true
	}

	return false
}

// SetDestinationName gets a reference to the given string and assigns it to the DestinationName field.
func (o *RecurrenceTransaction) SetDestinationName(v string) {
	o.DestinationName = &v
}

// GetDestinationIban returns the DestinationIban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetDestinationIban() string {
	if o == nil || o.DestinationIban.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationIban.Get()
}

// GetDestinationIbanOk returns a tuple with the DestinationIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetDestinationIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DestinationIban.Get(), o.DestinationIban.IsSet()
}

// HasDestinationIban returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasDestinationIban() bool {
	if o != nil && o.DestinationIban.IsSet() {
		return true
	}

	return false
}

// SetDestinationIban gets a reference to the given NullableString and assigns it to the DestinationIban field.
func (o *RecurrenceTransaction) SetDestinationIban(v string) {
	o.DestinationIban.Set(&v)
}
// SetDestinationIbanNil sets the value for DestinationIban to be an explicit nil
func (o *RecurrenceTransaction) SetDestinationIbanNil() {
	o.DestinationIban.Set(nil)
}

// UnsetDestinationIban ensures that no value is present for DestinationIban, not even an explicit nil
func (o *RecurrenceTransaction) UnsetDestinationIban() {
	o.DestinationIban.Unset()
}

// GetDestinationType returns the DestinationType field value if set, zero value otherwise.
func (o *RecurrenceTransaction) GetDestinationType() AccountTypeProperty {
	if o == nil || o.DestinationType == nil {
		var ret AccountTypeProperty
		return ret
	}
	return *o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceTransaction) GetDestinationTypeOk() (*AccountTypeProperty, bool) {
	if o == nil || o.DestinationType == nil {
		return nil, false
	}
	return o.DestinationType, true
}

// HasDestinationType returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasDestinationType() bool {
	if o != nil && o.DestinationType != nil {
		return true
	}

	return false
}

// SetDestinationType gets a reference to the given AccountTypeProperty and assigns it to the DestinationType field.
func (o *RecurrenceTransaction) SetDestinationType(v AccountTypeProperty) {
	o.DestinationType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetTags() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RecurrenceTransaction) SetTags(v []string) {
	o.Tags = v
}

// GetPiggyBankId returns the PiggyBankId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetPiggyBankId() string {
	if o == nil || o.PiggyBankId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PiggyBankId.Get()
}

// GetPiggyBankIdOk returns a tuple with the PiggyBankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetPiggyBankIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PiggyBankId.Get(), o.PiggyBankId.IsSet()
}

// HasPiggyBankId returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasPiggyBankId() bool {
	if o != nil && o.PiggyBankId.IsSet() {
		return true
	}

	return false
}

// SetPiggyBankId gets a reference to the given NullableString and assigns it to the PiggyBankId field.
func (o *RecurrenceTransaction) SetPiggyBankId(v string) {
	o.PiggyBankId.Set(&v)
}
// SetPiggyBankIdNil sets the value for PiggyBankId to be an explicit nil
func (o *RecurrenceTransaction) SetPiggyBankIdNil() {
	o.PiggyBankId.Set(nil)
}

// UnsetPiggyBankId ensures that no value is present for PiggyBankId, not even an explicit nil
func (o *RecurrenceTransaction) UnsetPiggyBankId() {
	o.PiggyBankId.Unset()
}

// GetPiggyBankName returns the PiggyBankName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecurrenceTransaction) GetPiggyBankName() string {
	if o == nil || o.PiggyBankName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PiggyBankName.Get()
}

// GetPiggyBankNameOk returns a tuple with the PiggyBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecurrenceTransaction) GetPiggyBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PiggyBankName.Get(), o.PiggyBankName.IsSet()
}

// HasPiggyBankName returns a boolean if a field has been set.
func (o *RecurrenceTransaction) HasPiggyBankName() bool {
	if o != nil && o.PiggyBankName.IsSet() {
		return true
	}

	return false
}

// SetPiggyBankName gets a reference to the given NullableString and assigns it to the PiggyBankName field.
func (o *RecurrenceTransaction) SetPiggyBankName(v string) {
	o.PiggyBankName.Set(&v)
}
// SetPiggyBankNameNil sets the value for PiggyBankName to be an explicit nil
func (o *RecurrenceTransaction) SetPiggyBankNameNil() {
	o.PiggyBankName.Set(nil)
}

// UnsetPiggyBankName ensures that no value is present for PiggyBankName, not even an explicit nil
func (o *RecurrenceTransaction) UnsetPiggyBankName() {
	o.PiggyBankName.Unset()
}

func (o RecurrenceTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.ForeignAmount.IsSet() {
		toSerialize["foreign_amount"] = o.ForeignAmount.Get()
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
	if o.ForeignCurrencyId.IsSet() {
		toSerialize["foreign_currency_id"] = o.ForeignCurrencyId.Get()
	}
	if o.ForeignCurrencyCode.IsSet() {
		toSerialize["foreign_currency_code"] = o.ForeignCurrencyCode.Get()
	}
	if o.ForeignCurrencySymbol.IsSet() {
		toSerialize["foreign_currency_symbol"] = o.ForeignCurrencySymbol.Get()
	}
	if o.ForeignCurrencyDecimalPlaces.IsSet() {
		toSerialize["foreign_currency_decimal_places"] = o.ForeignCurrencyDecimalPlaces.Get()
	}
	if o.BudgetId != nil {
		toSerialize["budget_id"] = o.BudgetId
	}
	if o.BudgetName.IsSet() {
		toSerialize["budget_name"] = o.BudgetName.Get()
	}
	if o.CategoryId != nil {
		toSerialize["category_id"] = o.CategoryId
	}
	if o.CategoryName != nil {
		toSerialize["category_name"] = o.CategoryName
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.SourceName != nil {
		toSerialize["source_name"] = o.SourceName
	}
	if o.SourceIban.IsSet() {
		toSerialize["source_iban"] = o.SourceIban.Get()
	}
	if o.SourceType != nil {
		toSerialize["source_type"] = o.SourceType
	}
	if o.DestinationId != nil {
		toSerialize["destination_id"] = o.DestinationId
	}
	if o.DestinationName != nil {
		toSerialize["destination_name"] = o.DestinationName
	}
	if o.DestinationIban.IsSet() {
		toSerialize["destination_iban"] = o.DestinationIban.Get()
	}
	if o.DestinationType != nil {
		toSerialize["destination_type"] = o.DestinationType
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.PiggyBankId.IsSet() {
		toSerialize["piggy_bank_id"] = o.PiggyBankId.Get()
	}
	if o.PiggyBankName.IsSet() {
		toSerialize["piggy_bank_name"] = o.PiggyBankName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRecurrenceTransaction struct {
	value *RecurrenceTransaction
	isSet bool
}

func (v NullableRecurrenceTransaction) Get() *RecurrenceTransaction {
	return v.value
}

func (v *NullableRecurrenceTransaction) Set(val *RecurrenceTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceTransaction(val *RecurrenceTransaction) *NullableRecurrenceTransaction {
	return &NullableRecurrenceTransaction{value: val, isSet: true}
}

func (v NullableRecurrenceTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


