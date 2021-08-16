# AccountUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Iban** | Pointer to **NullableString** |  | [optional] 
**Bic** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**OpeningBalance** | Pointer to **string** |  | [optional] 
**OpeningBalanceDate** | Pointer to **NullableString** |  | [optional] 
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
**Latitude** | Pointer to **NullableFloat64** | Latitude of the account&#39;s location, if applicable. Can be used to draw a map. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Latitude of the account&#39;s location, if applicable. Can be used to draw a map. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] 
**ZoomLevel** | Pointer to **NullableInt32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. If omitted, the existing location will be kept. If submitted as NULL, the current location will be removed. | [optional] 

## Methods

### NewAccountUpdate

`func NewAccountUpdate(name string, ) *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetIban

`func (o *AccountUpdate) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountUpdate) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountUpdate) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountUpdate) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *AccountUpdate) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *AccountUpdate) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBic

`func (o *AccountUpdate) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *AccountUpdate) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *AccountUpdate) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *AccountUpdate) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *AccountUpdate) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *AccountUpdate) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetAccountNumber

`func (o *AccountUpdate) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountUpdate) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountUpdate) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountUpdate) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *AccountUpdate) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *AccountUpdate) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetOpeningBalance

`func (o *AccountUpdate) GetOpeningBalance() string`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *AccountUpdate) GetOpeningBalanceOk() (*string, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *AccountUpdate) SetOpeningBalance(v string)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *AccountUpdate) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetOpeningBalanceDate

`func (o *AccountUpdate) GetOpeningBalanceDate() string`

GetOpeningBalanceDate returns the OpeningBalanceDate field if non-nil, zero value otherwise.

### GetOpeningBalanceDateOk

`func (o *AccountUpdate) GetOpeningBalanceDateOk() (*string, bool)`

GetOpeningBalanceDateOk returns a tuple with the OpeningBalanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalanceDate

`func (o *AccountUpdate) SetOpeningBalanceDate(v string)`

SetOpeningBalanceDate sets OpeningBalanceDate field to given value.

### HasOpeningBalanceDate

`func (o *AccountUpdate) HasOpeningBalanceDate() bool`

HasOpeningBalanceDate returns a boolean if a field has been set.

### SetOpeningBalanceDateNil

`func (o *AccountUpdate) SetOpeningBalanceDateNil(b bool)`

 SetOpeningBalanceDateNil sets the value for OpeningBalanceDate to be an explicit nil

### UnsetOpeningBalanceDate
`func (o *AccountUpdate) UnsetOpeningBalanceDate()`

UnsetOpeningBalanceDate ensures that no value is present for OpeningBalanceDate, not even an explicit nil
### GetVirtualBalance

`func (o *AccountUpdate) GetVirtualBalance() string`

GetVirtualBalance returns the VirtualBalance field if non-nil, zero value otherwise.

### GetVirtualBalanceOk

`func (o *AccountUpdate) GetVirtualBalanceOk() (*string, bool)`

GetVirtualBalanceOk returns a tuple with the VirtualBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualBalance

`func (o *AccountUpdate) SetVirtualBalance(v string)`

SetVirtualBalance sets VirtualBalance field to given value.

### HasVirtualBalance

`func (o *AccountUpdate) HasVirtualBalance() bool`

HasVirtualBalance returns a boolean if a field has been set.

### GetCurrencyId

`func (o *AccountUpdate) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *AccountUpdate) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *AccountUpdate) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *AccountUpdate) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountUpdate) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountUpdate) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountUpdate) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AccountUpdate) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetActive

`func (o *AccountUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AccountUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AccountUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AccountUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *AccountUpdate) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AccountUpdate) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AccountUpdate) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AccountUpdate) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIncludeNetWorth

`func (o *AccountUpdate) GetIncludeNetWorth() bool`

GetIncludeNetWorth returns the IncludeNetWorth field if non-nil, zero value otherwise.

### GetIncludeNetWorthOk

`func (o *AccountUpdate) GetIncludeNetWorthOk() (*bool, bool)`

GetIncludeNetWorthOk returns a tuple with the IncludeNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNetWorth

`func (o *AccountUpdate) SetIncludeNetWorth(v bool)`

SetIncludeNetWorth sets IncludeNetWorth field to given value.

### HasIncludeNetWorth

`func (o *AccountUpdate) HasIncludeNetWorth() bool`

HasIncludeNetWorth returns a boolean if a field has been set.

### GetAccountRole

`func (o *AccountUpdate) GetAccountRole() string`

GetAccountRole returns the AccountRole field if non-nil, zero value otherwise.

### GetAccountRoleOk

`func (o *AccountUpdate) GetAccountRoleOk() (*string, bool)`

GetAccountRoleOk returns a tuple with the AccountRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRole

`func (o *AccountUpdate) SetAccountRole(v string)`

SetAccountRole sets AccountRole field to given value.

### HasAccountRole

`func (o *AccountUpdate) HasAccountRole() bool`

HasAccountRole returns a boolean if a field has been set.

### SetAccountRoleNil

`func (o *AccountUpdate) SetAccountRoleNil(b bool)`

 SetAccountRoleNil sets the value for AccountRole to be an explicit nil

### UnsetAccountRole
`func (o *AccountUpdate) UnsetAccountRole()`

UnsetAccountRole ensures that no value is present for AccountRole, not even an explicit nil
### GetCreditCardType

`func (o *AccountUpdate) GetCreditCardType() string`

GetCreditCardType returns the CreditCardType field if non-nil, zero value otherwise.

### GetCreditCardTypeOk

`func (o *AccountUpdate) GetCreditCardTypeOk() (*string, bool)`

GetCreditCardTypeOk returns a tuple with the CreditCardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardType

`func (o *AccountUpdate) SetCreditCardType(v string)`

SetCreditCardType sets CreditCardType field to given value.

### HasCreditCardType

`func (o *AccountUpdate) HasCreditCardType() bool`

HasCreditCardType returns a boolean if a field has been set.

### SetCreditCardTypeNil

`func (o *AccountUpdate) SetCreditCardTypeNil(b bool)`

 SetCreditCardTypeNil sets the value for CreditCardType to be an explicit nil

### UnsetCreditCardType
`func (o *AccountUpdate) UnsetCreditCardType()`

UnsetCreditCardType ensures that no value is present for CreditCardType, not even an explicit nil
### GetMonthlyPaymentDate

`func (o *AccountUpdate) GetMonthlyPaymentDate() string`

GetMonthlyPaymentDate returns the MonthlyPaymentDate field if non-nil, zero value otherwise.

### GetMonthlyPaymentDateOk

`func (o *AccountUpdate) GetMonthlyPaymentDateOk() (*string, bool)`

GetMonthlyPaymentDateOk returns a tuple with the MonthlyPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPaymentDate

`func (o *AccountUpdate) SetMonthlyPaymentDate(v string)`

SetMonthlyPaymentDate sets MonthlyPaymentDate field to given value.

### HasMonthlyPaymentDate

`func (o *AccountUpdate) HasMonthlyPaymentDate() bool`

HasMonthlyPaymentDate returns a boolean if a field has been set.

### SetMonthlyPaymentDateNil

`func (o *AccountUpdate) SetMonthlyPaymentDateNil(b bool)`

 SetMonthlyPaymentDateNil sets the value for MonthlyPaymentDate to be an explicit nil

### UnsetMonthlyPaymentDate
`func (o *AccountUpdate) UnsetMonthlyPaymentDate()`

UnsetMonthlyPaymentDate ensures that no value is present for MonthlyPaymentDate, not even an explicit nil
### GetLiabilityType

`func (o *AccountUpdate) GetLiabilityType() string`

GetLiabilityType returns the LiabilityType field if non-nil, zero value otherwise.

### GetLiabilityTypeOk

`func (o *AccountUpdate) GetLiabilityTypeOk() (*string, bool)`

GetLiabilityTypeOk returns a tuple with the LiabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityType

`func (o *AccountUpdate) SetLiabilityType(v string)`

SetLiabilityType sets LiabilityType field to given value.

### HasLiabilityType

`func (o *AccountUpdate) HasLiabilityType() bool`

HasLiabilityType returns a boolean if a field has been set.

### SetLiabilityTypeNil

`func (o *AccountUpdate) SetLiabilityTypeNil(b bool)`

 SetLiabilityTypeNil sets the value for LiabilityType to be an explicit nil

### UnsetLiabilityType
`func (o *AccountUpdate) UnsetLiabilityType()`

UnsetLiabilityType ensures that no value is present for LiabilityType, not even an explicit nil
### GetInterest

`func (o *AccountUpdate) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *AccountUpdate) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *AccountUpdate) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *AccountUpdate) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### SetInterestNil

`func (o *AccountUpdate) SetInterestNil(b bool)`

 SetInterestNil sets the value for Interest to be an explicit nil

### UnsetInterest
`func (o *AccountUpdate) UnsetInterest()`

UnsetInterest ensures that no value is present for Interest, not even an explicit nil
### GetInterestPeriod

`func (o *AccountUpdate) GetInterestPeriod() string`

GetInterestPeriod returns the InterestPeriod field if non-nil, zero value otherwise.

### GetInterestPeriodOk

`func (o *AccountUpdate) GetInterestPeriodOk() (*string, bool)`

GetInterestPeriodOk returns a tuple with the InterestPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPeriod

`func (o *AccountUpdate) SetInterestPeriod(v string)`

SetInterestPeriod sets InterestPeriod field to given value.

### HasInterestPeriod

`func (o *AccountUpdate) HasInterestPeriod() bool`

HasInterestPeriod returns a boolean if a field has been set.

### SetInterestPeriodNil

`func (o *AccountUpdate) SetInterestPeriodNil(b bool)`

 SetInterestPeriodNil sets the value for InterestPeriod to be an explicit nil

### UnsetInterestPeriod
`func (o *AccountUpdate) UnsetInterestPeriod()`

UnsetInterestPeriod ensures that no value is present for InterestPeriod, not even an explicit nil
### GetNotes

`func (o *AccountUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AccountUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AccountUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AccountUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AccountUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AccountUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetLatitude

`func (o *AccountUpdate) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AccountUpdate) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AccountUpdate) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AccountUpdate) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *AccountUpdate) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *AccountUpdate) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *AccountUpdate) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AccountUpdate) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AccountUpdate) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AccountUpdate) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *AccountUpdate) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *AccountUpdate) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetZoomLevel

`func (o *AccountUpdate) GetZoomLevel() int32`

GetZoomLevel returns the ZoomLevel field if non-nil, zero value otherwise.

### GetZoomLevelOk

`func (o *AccountUpdate) GetZoomLevelOk() (*int32, bool)`

GetZoomLevelOk returns a tuple with the ZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomLevel

`func (o *AccountUpdate) SetZoomLevel(v int32)`

SetZoomLevel sets ZoomLevel field to given value.

### HasZoomLevel

`func (o *AccountUpdate) HasZoomLevel() bool`

HasZoomLevel returns a boolean if a field has been set.

### SetZoomLevelNil

`func (o *AccountUpdate) SetZoomLevelNil(b bool)`

 SetZoomLevelNil sets the value for ZoomLevel to be an explicit nil

### UnsetZoomLevel
`func (o *AccountUpdate) UnsetZoomLevel()`

UnsetZoomLevel ensures that no value is present for ZoomLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


