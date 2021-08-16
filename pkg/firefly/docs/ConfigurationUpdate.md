# ConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**string**](oneOf&lt;boolean,string,object,array&gt;.md) | Can be a number, a string, boolean or object. This depends on the actual configuration value. | 

## Methods

### NewConfigurationUpdate

`func NewConfigurationUpdate(value string, ) *ConfigurationUpdate`

NewConfigurationUpdate instantiates a new ConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationUpdateWithDefaults

`func NewConfigurationUpdateWithDefaults() *ConfigurationUpdate`

NewConfigurationUpdateWithDefaults instantiates a new ConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ConfigurationUpdate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigurationUpdate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigurationUpdate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


