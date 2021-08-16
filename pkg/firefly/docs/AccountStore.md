# AccountStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** | Can only be one one these account types. import, initial-balance and reconciliation cannot be set manually. | 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**OpeningBalance** | Pointer to **string** | Represents the opening balance, the initial amount this account holds. | [optional] 
**OpeningBalanceDate** | Pointer to **NullableString** | Represents the date of the opening balance. | [optional] 
**VirtualBalance** | Pointer to **string** |  | [optional] 
**CurrencyId** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**CurrencyCode** | Pointer to **string** | Use either currency_id or currency_code. Defaults to the user&#39;s default currency. | [optional] 
**Active** | Pointer to **bool** | If omitted, defaults to true. | [optional] 
**Order** | Pointer to **int32** | Order of the account | [optional] 
**IncludeNetWorth** | Pointer to **bool** | If omitted, defaults to true. | [optional] 
**AccountRole** | Pointer to **NullableString** | Is only mandatory when the type is asset. | [optional] 
**CreditCardType** | Pointer to **NullableString** | Mandatory when the account_role is ccAsset. Can only be monthlyFull or null. | [optional] 
**MonthlyPaymentDate** | Pointer to **NullableString** | Mandatory when the account_role is ccAsset. Moment at which CC payment installments are asked for by the bank. | [optional] 
**LiabilityType** | Pointer to **NullableString** | Mandatory when type is liability. Specifies the exact type. | [optional] 
**Interest** | Pointer to **NullableString** | Mandatory when type is liability. Interest percentage. | [optional] 
**InterestPeriod** | Pointer to **NullableString** | Mandatory when type is liability. Period over which the interest is calculated. | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Latitude of the accounts&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Latitude of the accounts&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**ZoomLevel** | Pointer to **NullableInt32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. | [optional] 

## Methods

### NewAccountStore

`func NewAccountStore(name string, type_ string, ) *AccountStore`

NewAccountStore instantiates a new AccountStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStoreWithDefaults

`func NewAccountStoreWithDefaults() *AccountStore`

NewAccountStoreWithDefaults instantiates a new AccountStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountStore) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AccountStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountStore) SetType(v string)`

SetType sets Type field to given value.


### GetIban

`func (o *AccountStore) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountStore) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountStore) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountStore) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *AccountStore) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *AccountStore) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *AccountStore) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *AccountStore) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *AccountStore) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *AccountStore) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *AccountStore) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *AccountStore) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetAccountNumber

`func (o *AccountStore) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountStore) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountStore) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountStore) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *AccountStore) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *AccountStore) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetOpeningBalance

`func (o *AccountStore) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *AccountStore) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *AccountStore) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *AccountStore) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *AccountStore) GetOpeningBalanceDate() string`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *AccountStore) GetOpeningBalanceDateOk() (*string, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *AccountStore) SetOpeningBalanceDate(v string)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *AccountStore) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### SetOpeningBalanceDateNil

`func (o *AccountStore) SetOpeningBalanceDateNil(b bool)`

 SetOpeningBalanceDateNil sets the value for OpeningBalanceDate to be an explicit nil

### UnsetOpeningBalanceDate
`func (o *AccountStore) UnsetOpeningBalanceDate()`

UnsetOpeningBalanceDate ensures that no value is present for OpeningBalanceDate, not even an explicit nil
### GetVirtualBalance

`func (o *AccountStore) GetVirtualBalance() string`

GetVirtualBalance returns the VirtualBalance field if non-nil, zero value otherwise.

### GetVirtualBalanceOk

`func (o *AccountStore) GetVirtualBalanceOk() (*string, bool)`

GetVirtualBalanceOk returns a tuple with the VirtualBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualBalance

`func (o *AccountStore) SetVirtualBalance(v string)`

SetVirtualBalance sets VirtualBalance field to given value.

### HasVirtualBalance

`func (o *AccountStore) HasVirtualBalance() bool`

HasVirtualBalance returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AccountStore) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountStore) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountStore) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AccountStore) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountStore) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountStore) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountStore) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountStore) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetActive

`func (o *AccountStore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AccountStore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AccountStore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AccountStore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *AccountStore) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AccountStore) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AccountStore) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AccountStore) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIncludeNetWorth

`func (o *AccountStore) GetIncludeNetWorth() bool`

GetIncludeNetWorth returns the IncludeNetWorth field if non-nil, zero value otherwise.

### GetIncludeNetWorthOk

`func (o *AccountStore) GetIncludeNetWorthOk() (*bool, bool)`

GetIncludeNetWorthOk returns a tuple with the IncludeNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNetWorth

`func (o *AccountStore) SetIncludeNetWorth(v bool)`

SetIncludeNetWorth sets IncludeNetWorth field to given value.

### HasIncludeNetWorth

`func (o *AccountStore) HasIncludeNetWorth() bool`

HasIncludeNetWorth returns a boolean if a field has been set.

### GetAccountRole

`func (o *AccountStore) GetAccountRole() string`

GetAccountRole returns the AccountRole field if non-nil, zero value otherwise.

### GetAccountRoleOk

`func (o *AccountStore) GetAccountRoleOk() (*string, bool)`

GetAccountRoleOk returns a tuple with the AccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRole

`func (o *AccountStore) SetAccountRole(v string)`

SetAccountRole sets AccountRole field to given value.

### HasAccountRole

`func (o *AccountStore) HasAccountRole() bool`

HasAccountRole returns a boolean if a field has been set.

### SetAccountRoleNil

`func (o *AccountStore) SetAccountRoleNil(b bool)`

 SetAccountRoleNil sets the value for AccountRole to be an explicit nil

### UnsetAccountRole
`func (o *AccountStore) UnsetAccountRole()`

UnsetAccountRole ensures that no value is present for AccountRole, not even an explicit nil
### GetCreditCardType

`func (o *AccountStore) GetCreditCardType() string`

GetCreditCardType returns the CreditCardType field if non-nil, zero value otherwise.

### GetCreditCardTypeOk

`func (o *AccountStore) GetCreditCardTypeOk() (*string, bool)`

GetCreditCardTypeOk returns a tuple with the CreditCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardType

`func (o *AccountStore) SetCreditCardType(v string)`

SetCreditCardType sets CreditCardType field to given value.

### HasCreditCardType

`func (o *AccountStore) HasCreditCardType() bool`

HasCreditCardType returns a boolean if a field has been set.

### SetCreditCardTypeNil

`func (o *AccountStore) SetCreditCardTypeNil(b bool)`

 SetCreditCardTypeNil sets the value for CreditCardType to be an explicit nil

### UnsetCreditCardType
`func (o *AccountStore) UnsetCreditCardType()`

UnsetCreditCardType ensures that no value is present for CreditCardType, not even an explicit nil
### GetMonthlyPaymentDate

`func (o *AccountStore) GetMonthlyPaymentDate() string`

GetMonthlyPaymentDate returns the MonthlyPaymentDate field if non-nil, zero value otherwise.

### GetMonthlyPaymentDateOk

`func (o *AccountStore) GetMonthlyPaymentDateOk() (*string, bool)`

GetMonthlyPaymentDateOk returns a tuple with the MonthlyPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPaymentDate

`func (o *AccountStore) SetMonthlyPaymentDate(v string)`

SetMonthlyPaymentDate sets MonthlyPaymentDate field to given value.

### HasMonthlyPaymentDate

`func (o *AccountStore) HasMonthlyPaymentDate() bool`

HasMonthlyPaymentDate returns a boolean if a field has been set.

### SetMonthlyPaymentDateNil

`func (o *AccountStore) SetMonthlyPaymentDateNil(b bool)`

 SetMonthlyPaymentDateNil sets the value for MonthlyPaymentDate to be an explicit nil

### UnsetMonthlyPaymentDate
`func (o *AccountStore) UnsetMonthlyPaymentDate()`

UnsetMonthlyPaymentDate ensures that no value is present for MonthlyPaymentDate, not even an explicit nil
### GetLiabilityType

`func (o *AccountStore) GetLiabilityType() string`

GetLiabilityType returns the LiabilityType field if non-nil, zero value otherwise.

### GetLiabilityTypeOk

`func (o *AccountStore) GetLiabilityTypeOk() (*string, bool)`

GetLiabilityTypeOk returns a tuple with the LiabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityType

`func (o *AccountStore) SetLiabilityType(v string)`

SetLiabilityType sets LiabilityType field to given value.

### HasLiabilityType

`func (o *AccountStore) HasLiabilityType() bool`

HasLiabilityType returns a boolean if a field has been set.

### SetLiabilityTypeNil

`func (o *AccountStore) SetLiabilityTypeNil(b bool)`

 SetLiabilityTypeNil sets the value for LiabilityType to be an explicit nil

### UnsetLiabilityType
`func (o *AccountStore) UnsetLiabilityType()`

UnsetLiabilityType ensures that no value is present for LiabilityType, not even an explicit nil
### GetInterest

`func (o *AccountStore) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *AccountStore) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *AccountStore) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *AccountStore) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### SetInterestNil

`func (o *AccountStore) SetInterestNil(b bool)`

 SetInterestNil sets the value for Interest to be an explicit nil

### UnsetInterest
`func (o *AccountStore) UnsetInterest()`

UnsetInterest ensures that no value is present for Interest, not even an explicit nil
### GetInterestPeriod

`func (o *AccountStore) GetInterestPeriod() string`

GetInterestPeriod returns the InterestPeriod field if non-nil, zero value otherwise.

### GetInterestPeriodOk

`func (o *AccountStore) GetInterestPeriodOk() (*string, bool)`

GetInterestPeriodOk returns a tuple with the InterestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPeriod

`func (o *AccountStore) SetInterestPeriod(v string)`

SetInterestPeriod sets InterestPeriod field to given value.

### HasInterestPeriod

`func (o *AccountStore) HasInterestPeriod() bool`

HasInterestPeriod returns a boolean if a field has been set.

### SetInterestPeriodNil

`func (o *AccountStore) SetInterestPeriodNil(b bool)`

 SetInterestPeriodNil sets the value for InterestPeriod to be an explicit nil

### UnsetInterestPeriod
`func (o *AccountStore) UnsetInterestPeriod()`

UnsetInterestPeriod ensures that no value is present for InterestPeriod, not even an explicit nil
### GetNotes

`func (o *AccountStore) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AccountStore) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AccountStore) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AccountStore) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AccountStore) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AccountStore) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLatitude

`func (o *AccountStore) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AccountStore) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AccountStore) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AccountStore) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *AccountStore) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *AccountStore) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *AccountStore) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AccountStore) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AccountStore) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AccountStore) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *AccountStore) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *AccountStore) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetZoomLevel

`func (o *AccountStore) GetZoomLevel() int32`

GetZoomLevel returns the ZoomLevel field if non-nil, zero value otherwise.

### GetZoomLevelOk

`func (o *AccountStore) GetZoomLevelOk() (*int32, bool)`

GetZoomLevelOk returns a tuple with the ZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomLevel

`func (o *AccountStore) SetZoomLevel(v int32)`

SetZoomLevel sets ZoomLevel field to given value.

### HasZoomLevel

`func (o *AccountStore) HasZoomLevel() bool`

HasZoomLevel returns a boolean if a field has been set.

### SetZoomLevelNil

`func (o *AccountStore) SetZoomLevelNil(b bool)`

 SetZoomLevelNil sets the value for ZoomLevel to be an explicit nil

### UnsetZoomLevel
`func (o *AccountStore) UnsetZoomLevel()`

UnsetZoomLevel ensures that no value is present for ZoomLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


