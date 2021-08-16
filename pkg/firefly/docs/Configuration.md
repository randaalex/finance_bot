# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | [**ConfigValueFilter**](ConfigValueFilter.md) |  | 
**Value** | [**string**](oneOf&lt;boolean,string,object,array&gt;.md) | Content of the configuration variable. Is very dynamic and can be anything from booleans to arrays. | 
**Editable** | **bool** | If this config variable can be edited by the user | 

## Methods

### NewConfiguration

`func NewConfiguration(title ConfigValueFilter, value string, editable bool, ) *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Configuration) GetTitle() ConfigValueFilter`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Configuration) GetTitleOk() (*ConfigValueFilter, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Configuration) SetTitle(v ConfigValueFilter)`

SetTitle sets Title field to given value.


### GetValue

`func (o *Configuration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Configuration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Configuration) SetValue(v string)`

SetValue sets Value field to given value.


### GetEditable

`func (o *Configuration) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Configuration) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Configuration) SetEditable(v bool)`

SetEditable sets Editable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


