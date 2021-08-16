# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Type** | **string** | Can only be one one these account types. import, initial-balance and reconciliation cannot be set manually. | 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**OpeningBalance** | Pointer to **string** | Represents the opening balance, the initial amount this account holds. | [optional] 
**OpeningBalanceDate** | Pointer to **NullableTime** | Represents the date of the opening balance. | [optional] 
**VirtualBalance** | Pointer to **string** |  | [optional] 
**CurrentBalance** | Pointer to **string** |  | [optional] [readonly] 
**CurrentBalanceDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencySymbol** | Pointer to **string** |  | [optional] [readonly] 
**CurrencyDecimalPlaces** | Pointer to **int32** |  | [optional] [readonly] 
**Active** | Pointer to **bool** | If omitted, defaults to true. | [optional] 
**Order** | Pointer to **NullableInt32** | Order of the account. Is NULL if account is not asset or liability. | [optional] 
**IncludeNetWorth** | Pointer to **bool** | If omitted, defaults to true. | [optional] 
**AccountRole** | Pointer to **NullableString** | Is only mandatory when the type is asset. | [optional] 
**CreditCardType** | Pointer to **NullableString** | Mandatory when the account_role is ccAsset. Can only be monthlyFull or null. | [optional] 
**MonthlyPaymentDate** | Pointer to **NullableTime** | Mandatory when the account_role is ccAsset. Moment at which CC payment installments are asked for by the bank. | [optional] 
**LiabilityType** | Pointer to **NullableString** | Mandatory when type is liability. Specifies the exact type. | [optional] 
**Interest** | Pointer to **NullableString** | Mandatory when type is liability. Interest percentage. | [optional] 
**InterestPeriod** | Pointer to **NullableString** | Mandatory when type is liability. Period over which the interest is calculated. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Latitude of the accounts&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Latitude of the accounts&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**ZoomLevel** | Pointer to **NullableInt32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. | [optional] 

## Methods

### NewAccount

`func NewAccount(name string, type_ string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.


### GetIban

`func (o *Account) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Account) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Account) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Account) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *Account) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Account) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *Account) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Account) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Account) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *Account) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *Account) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *Account) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Account) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *Account) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *Account) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetOpeningBalance

`func (o *Account) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *Account) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *Account) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *Account) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *Account) GetOpeningBalanceDate() time.Time`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *Account) GetOpeningBalanceDateOk() (*time.Time, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *Account) SetOpeningBalanceDate(v time.Time)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *Account) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### SetOpeningBalanceDateNil

`func (o *Account) SetOpeningBalanceDateNil(b bool)`

 SetOpeningBalanceDateNil sets the value for OpeningBalanceDate to be an explicit nil

### UnsetOpeningBalanceDate
`func (o *Account) UnsetOpeningBalanceDate()`

UnsetOpeningBalanceDate ensures that no value is present for OpeningBalanceDate, not even an explicit nil
### GetVirtualBalance

`func (o *Account) GetVirtualBalance() string`

GetVirtualBalance returns the VirtualBalance field if non-nil, zero value otherwise.

### GetVirtualBalanceOk

`func (o *Account) GetVirtualBalanceOk() (*string, bool)`

GetVirtualBalanceOk returns a tuple with the VirtualBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualBalance

`func (o *Account) SetVirtualBalance(v string)`

SetVirtualBalance sets VirtualBalance field to given value.

### HasVirtualBalance

`func (o *Account) HasVirtualBalance() bool`

HasVirtualBalance returns a boolean if a field has been set.

### GetCurrentBalance

`func (o *Account) GetCurrentBalance() string`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *Account) GetCurrentBalanceOk() (*string, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *Account) SetCurrentBalance(v string)`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *Account) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### GetCurrentBalanceDate

`func (o *Account) GetCurrentBalanceDate() time.Time`

GetCurrentBalanceDate returns the CurrentBalanceDate field if non-nil, zero value otherwise.

### GetCurrentBalanceDateOk

`func (o *Account) GetCurrentBalanceDateOk() (*time.Time, bool)`

GetCurrentBalanceDateOk returns a tuple with the CurrentBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalanceDate

`func (o *Account) SetCurrentBalanceDate(v time.Time)`

SetCurrentBalanceDate sets CurrentBalanceDate field to given value.

### HasCurrentBalanceDate

`func (o *Account) HasCurrentBalanceDate() bool`

HasCurrentBalanceDate returns a boolean if a field has been set.

### GetCurrencyId

`func (o *Account) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *Account) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *Account) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *Account) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Account) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Account) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Account) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Account) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencySymbol

`func (o *Account) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *Account) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *Account) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *Account) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetCurrencyDecimalPlaces

`func (o *Account) GetCurrencyDecimalPlaces() int32`

GetCurrencyDecimalPlaces returns the CurrencyDecimalPlaces field if non-nil, zero value otherwise.

### GetCurrencyDecimalPlacesOk

`func (o *Account) GetCurrencyDecimalPlacesOk() (*int32, bool)`

GetCurrencyDecimalPlacesOk returns a tuple with the CurrencyDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDecimalPlaces

`func (o *Account) SetCurrencyDecimalPlaces(v int32)`

SetCurrencyDecimalPlaces sets CurrencyDecimalPlaces field to given value.

### HasCurrencyDecimalPlaces

`func (o *Account) HasCurrencyDecimalPlaces() bool`

HasCurrencyDecimalPlaces returns a boolean if a field has been set.

### GetActive

`func (o *Account) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Account) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Account) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Account) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *Account) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Account) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Account) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Account) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *Account) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *Account) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIncludeNetWorth

`func (o *Account) GetIncludeNetWorth() bool`

GetIncludeNetWorth returns the IncludeNetWorth field if non-nil, zero value otherwise.

### GetIncludeNetWorthOk

`func (o *Account) GetIncludeNetWorthOk() (*bool, bool)`

GetIncludeNetWorthOk returns a tuple with the IncludeNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNetWorth

`func (o *Account) SetIncludeNetWorth(v bool)`

SetIncludeNetWorth sets IncludeNetWorth field to given value.

### HasIncludeNetWorth

`func (o *Account) HasIncludeNetWorth() bool`

HasIncludeNetWorth returns a boolean if a field has been set.

### GetAccountRole

`func (o *Account) GetAccountRole() string`

GetAccountRole returns the AccountRole field if non-nil, zero value otherwise.

### GetAccountRoleOk

`func (o *Account) GetAccountRoleOk() (*string, bool)`

GetAccountRoleOk returns a tuple with the AccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRole

`func (o *Account) SetAccountRole(v string)`

SetAccountRole sets AccountRole field to given value.

### HasAccountRole

`func (o *Account) HasAccountRole() bool`

HasAccountRole returns a boolean if a field has been set.

### SetAccountRoleNil

`func (o *Account) SetAccountRoleNil(b bool)`

 SetAccountRoleNil sets the value for AccountRole to be an explicit nil

### UnsetAccountRole
`func (o *Account) UnsetAccountRole()`

UnsetAccountRole ensures that no value is present for AccountRole, not even an explicit nil
### GetCreditCardType

`func (o *Account) GetCreditCardType() string`

GetCreditCardType returns the CreditCardType field if non-nil, zero value otherwise.

### GetCreditCardTypeOk

`func (o *Account) GetCreditCardTypeOk() (*string, bool)`

GetCreditCardTypeOk returns a tuple with the CreditCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardType

`func (o *Account) SetCreditCardType(v string)`

SetCreditCardType sets CreditCardType field to given value.

### HasCreditCardType

`func (o *Account) HasCreditCardType() bool`

HasCreditCardType returns a boolean if a field has been set.

### SetCreditCardTypeNil

`func (o *Account) SetCreditCardTypeNil(b bool)`

 SetCreditCardTypeNil sets the value for CreditCardType to be an explicit nil

### UnsetCreditCardType
`func (o *Account) UnsetCreditCardType()`

UnsetCreditCardType ensures that no value is present for CreditCardType, not even an explicit nil
### GetMonthlyPaymentDate

`func (o *Account) GetMonthlyPaymentDate() time.Time`

GetMonthlyPaymentDate returns the MonthlyPaymentDate field if non-nil, zero value otherwise.

### GetMonthlyPaymentDateOk

`func (o *Account) GetMonthlyPaymentDateOk() (*time.Time, bool)`

GetMonthlyPaymentDateOk returns a tuple with the MonthlyPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPaymentDate

`func (o *Account) SetMonthlyPaymentDate(v time.Time)`

SetMonthlyPaymentDate sets MonthlyPaymentDate field to given value.

### HasMonthlyPaymentDate

`func (o *Account) HasMonthlyPaymentDate() bool`

HasMonthlyPaymentDate returns a boolean if a field has been set.

### SetMonthlyPaymentDateNil

`func (o *Account) SetMonthlyPaymentDateNil(b bool)`

 SetMonthlyPaymentDateNil sets the value for MonthlyPaymentDate to be an explicit nil

### UnsetMonthlyPaymentDate
`func (o *Account) UnsetMonthlyPaymentDate()`

UnsetMonthlyPaymentDate ensures that no value is present for MonthlyPaymentDate, not even an explicit nil
### GetLiabilityType

`func (o *Account) GetLiabilityType() string`

GetLiabilityType returns the LiabilityType field if non-nil, zero value otherwise.

### GetLiabilityTypeOk

`func (o *Account) GetLiabilityTypeOk() (*string, bool)`

GetLiabilityTypeOk returns a tuple with the LiabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityType

`func (o *Account) SetLiabilityType(v string)`

SetLiabilityType sets LiabilityType field to given value.

### HasLiabilityType

`func (o *Account) HasLiabilityType() bool`

HasLiabilityType returns a boolean if a field has been set.

### SetLiabilityTypeNil

`func (o *Account) SetLiabilityTypeNil(b bool)`

 SetLiabilityTypeNil sets the value for LiabilityType to be an explicit nil

### UnsetLiabilityType
`func (o *Account) UnsetLiabilityType()`

UnsetLiabilityType ensures that no value is present for LiabilityType, not even an explicit nil
### GetInterest

`func (o *Account) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *Account) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *Account) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *Account) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### SetInterestNil

`func (o *Account) SetInterestNil(b bool)`

 SetInterestNil sets the value for Interest to be an explicit nil

### UnsetInterest
`func (o *Account) UnsetInterest()`

UnsetInterest ensures that no value is present for Interest, not even an explicit nil
### GetInterestPeriod

`func (o *Account) GetInterestPeriod() string`

GetInterestPeriod returns the InterestPeriod field if non-nil, zero value otherwise.

### GetInterestPeriodOk

`func (o *Account) GetInterestPeriodOk() (*string, bool)`

GetInterestPeriodOk returns a tuple with the InterestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPeriod

`func (o *Account) SetInterestPeriod(v string)`

SetInterestPeriod sets InterestPeriod field to given value.

### HasInterestPeriod

`func (o *Account) HasInterestPeriod() bool`

HasInterestPeriod returns a boolean if a field has been set.

### SetInterestPeriodNil

`func (o *Account) SetInterestPeriodNil(b bool)`

 SetInterestPeriodNil sets the value for InterestPeriod to be an explicit nil

### UnsetInterestPeriod
`func (o *Account) UnsetInterestPeriod()`

UnsetInterestPeriod ensures that no value is present for InterestPeriod, not even an explicit nil
### GetNotes

`func (o *Account) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Account) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Account) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Account) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Account) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Account) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLatitude

`func (o *Account) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Account) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Account) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Account) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Account) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Account) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Account) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Account) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Account) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Account) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Account) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Account) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetZoomLevel

`func (o *Account) GetZoomLevel() int32`

GetZoomLevel returns the ZoomLevel field if non-nil, zero value otherwise.

### GetZoomLevelOk

`func (o *Account) GetZoomLevelOk() (*int32, bool)`

GetZoomLevelOk returns a tuple with the ZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomLevel

`func (o *Account) SetZoomLevel(v int32)`

SetZoomLevel sets ZoomLevel field to given value.

### HasZoomLevel

`func (o *Account) HasZoomLevel() bool`

HasZoomLevel returns a boolean if a field has been set.

### SetZoomLevelNil

`func (o *Account) SetZoomLevelNil(b bool)`

 SetZoomLevelNil sets the value for ZoomLevel to be an explicit nil

### UnsetZoomLevel
`func (o *Account) UnsetZoomLevel()`

UnsetZoomLevel ensures that no value is present for ZoomLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


