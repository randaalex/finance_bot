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

// Transaction struct for Transaction
type Transaction struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// User ID
	User *int32 `json:"user,omitempty"`
	// Break if the submitted transaction exists already.
	ErrorIfDuplicateHash *bool `json:"error_if_duplicate_hash,omitempty"`
	// Whether or not to apply rules when submitting transaction.
	ApplyRules *bool `json:"apply_rules,omitempty"`
	// Title of the transaction if it has been split in more than one piece. Empty otherwise.
	GroupTitle *string `json:"group_title,omitempty"`
	Transactions []TransactionSplit `json:"transactions"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(transactions []TransactionSplit) *Transaction {
	this := Transaction{}
	this.Transactions = transactions
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Transaction) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Transaction) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Transaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Transaction) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Transaction) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Transaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Transaction) GetUser() int32 {
	if o == nil || o.User == nil {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetUserOk() (*int32, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Transaction) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *Transaction) SetUser(v int32) {
	o.User = &v
}

// GetErrorIfDuplicateHash returns the ErrorIfDuplicateHash field value if set, zero value otherwise.
func (o *Transaction) GetErrorIfDuplicateHash() bool {
	if o == nil || o.ErrorIfDuplicateHash == nil {
		var ret bool
		return ret
	}
	return *o.ErrorIfDuplicateHash
}

// GetErrorIfDuplicateHashOk returns a tuple with the ErrorIfDuplicateHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetErrorIfDuplicateHashOk() (*bool, bool) {
	if o == nil || o.ErrorIfDuplicateHash == nil {
		return nil, false
	}
	return o.ErrorIfDuplicateHash, true
}

// HasErrorIfDuplicateHash returns a boolean if a field has been set.
func (o *Transaction) HasErrorIfDuplicateHash() bool {
	if o != nil && o.ErrorIfDuplicateHash != nil {
		return true
	}

	return false
}

// SetErrorIfDuplicateHash gets a reference to the given bool and assigns it to the ErrorIfDuplicateHash field.
func (o *Transaction) SetErrorIfDuplicateHash(v bool) {
	o.ErrorIfDuplicateHash = &v
}

// GetApplyRules returns the ApplyRules field value if set, zero value otherwise.
func (o *Transaction) GetApplyRules() bool {
	if o == nil || o.ApplyRules == nil {
		var ret bool
		return ret
	}
	return *o.ApplyRules
}

// GetApplyRulesOk returns a tuple with the ApplyRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetApplyRulesOk() (*bool, bool) {
	if o == nil || o.ApplyRules == nil {
		return nil, false
	}
	return o.ApplyRules, true
}

// HasApplyRules returns a boolean if a field has been set.
func (o *Transaction) HasApplyRules() bool {
	if o != nil && o.ApplyRules != nil {
		return true
	}

	return false
}

// SetApplyRules gets a reference to the given bool and assigns it to the ApplyRules field.
func (o *Transaction) SetApplyRules(v bool) {
	o.ApplyRules = &v
}

// GetGroupTitle returns the GroupTitle field value if set, zero value otherwise.
func (o *Transaction) GetGroupTitle() string {
	if o == nil || o.GroupTitle == nil {
		var ret string
		return ret
	}
	return *o.GroupTitle
}

// GetGroupTitleOk returns a tuple with the GroupTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetGroupTitleOk() (*string, bool) {
	if o == nil || o.GroupTitle == nil {
		return nil, false
	}
	return o.GroupTitle, true
}

// HasGroupTitle returns a boolean if a field has been set.
func (o *Transaction) HasGroupTitle() bool {
	if o != nil && o.GroupTitle != nil {
		return true
	}

	return false
}

// SetGroupTitle gets a reference to the given string and assigns it to the GroupTitle field.
func (o *Transaction) SetGroupTitle(v string) {
	o.GroupTitle = &v
}

// GetTransactions returns the Transactions field value
func (o *Transaction) GetTransactions() []TransactionSplit {
	if o == nil {
		var ret []TransactionSplit
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionsOk() (*[]TransactionSplit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *Transaction) SetTransactions(v []TransactionSplit) {
	o.Transactions = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ErrorIfDuplicateHash != nil {
		toSerialize["error_if_duplicate_hash"] = o.ErrorIfDuplicateHash
	}
	if o.ApplyRules != nil {
		toSerialize["apply_rules"] = o.ApplyRules
	}
	if o.GroupTitle != nil {
		toSerialize["group_title"] = o.GroupTitle
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

