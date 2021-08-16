# BillPaidDates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionGroupId** | Pointer to **string** | Transaction group ID of the paid bill. | [optional] [readonly] 
**TransactionJournalId** | Pointer to **string** | Transaction journal ID of the paid bill. | [optional] [readonly] 
**Date** | Pointer to **time.Time** | Date the bill was paid. | [optional] [readonly] 

## Methods

### NewBillPaidDates

`func NewBillPaidDates() *BillPaidDates`

NewBillPaidDates instantiates a new BillPaidDates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillPaidDatesWithDefaults

`func NewBillPaidDatesWithDefaults() *BillPaidDates`

NewBillPaidDatesWithDefaults instantiates a new BillPaidDates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionGroupId

`func (o *BillPaidDates) GetTransactionGroupId() string`

GetTransactionGroupId returns the TransactionGroupId field if non-nil, zero value otherwise.

### GetTransactionGroupIdOk

`func (o *BillPaidDates) GetTransactionGroupIdOk() (*string, bool)`

GetTransactionGroupIdOk returns a tuple with the TransactionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionGroupId

`func (o *BillPaidDates) SetTransactionGroupId(v string)`

SetTransactionGroupId sets TransactionGroupId field to given value.

### HasTransactionGroupId

`func (o *BillPaidDates) HasTransactionGroupId() bool`

HasTransactionGroupId returns a boolean if a field has been set.

### GetTransactionJournalId

`func (o *BillPaidDates) GetTransactionJournalId() string`

GetTransactionJournalId returns the TransactionJournalId field if non-nil, zero value otherwise.

### GetTransactionJournalIdOk

`func (o *BillPaidDates) GetTransactionJournalIdOk() (*string, bool)`

GetTransactionJournalIdOk returns a tuple with the TransactionJournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionJournalId

`func (o *BillPaidDates) SetTransactionJournalId(v string)`

SetTransactionJournalId sets TransactionJournalId field to given value.

### HasTransactionJournalId

`func (o *BillPaidDates) HasTransactionJournalId() bool`

HasTransactionJournalId returns a boolean if a field has been set.

### GetDate

`func (o *BillPaidDates) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BillPaidDates) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BillPaidDates) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BillPaidDates) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


