# CurrencyRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Immutable value | 
**Id** | **string** |  | 
**Attributes** | [**Currency**](Currency.md) |  | 

## Methods

### NewCurrencyRead

`func NewCurrencyRead(type_ string, id string, attributes Currency, ) *CurrencyRead`

NewCurrencyRead instantiates a new CurrencyRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyReadWithDefaults

`func NewCurrencyReadWithDefaults() *CurrencyRead`

NewCurrencyReadWithDefaults instantiates a new CurrencyRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CurrencyRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrencyRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrencyRead) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CurrencyRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyRead) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CurrencyRead) GetAttributes() Currency`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CurrencyRead) GetAttributesOk() (*Currency, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CurrencyRead) SetAttributes(v Currency)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


