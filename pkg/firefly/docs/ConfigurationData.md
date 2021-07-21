# ConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDemoSite** | Pointer to **bool** |  | [optional] 
**PermissionUpdateCheck** | Pointer to **NullableInt32** | If the user has given permission to check for updates. - null &#x3D; never asked. - -1 &#x3D; never asked. - 0 &#x3D; no permission. - 1 &#x3D; permission  | [optional] 
**LastUpdateCheck** | Pointer to **time.Time** |  | [optional] 
**SingleUserMode** | Pointer to **bool** | Whether other users can register. | [optional] 

## Methods

### NewConfigurationData

`func NewConfigurationData() *ConfigurationData`

NewConfigurationData instantiates a new ConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationDataWithDefaults

`func NewConfigurationDataWithDefaults() *ConfigurationData`

NewConfigurationDataWithDefaults instantiates a new ConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDemoSite

`func (o *ConfigurationData) GetIsDemoSite() bool`

GetIsDemoSite returns the IsDemoSite field if non-nil, zero value otherwise.

### GetIsDemoSiteOk

`func (o *ConfigurationData) GetIsDemoSiteOk() (*bool, bool)`

GetIsDemoSiteOk returns a tuple with the IsDemoSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemoSite

`func (o *ConfigurationData) SetIsDemoSite(v bool)`

SetIsDemoSite sets IsDemoSite field to given value.

### HasIsDemoSite

`func (o *ConfigurationData) HasIsDemoSite() bool`

HasIsDemoSite returns a boolean if a field has been set.

### GetPermissionUpdateCheck

`func (o *ConfigurationData) GetPermissionUpdateCheck() int32`

GetPermissionUpdateCheck returns the PermissionUpdateCheck field if non-nil, zero value otherwise.

### GetPermissionUpdateCheckOk

`func (o *ConfigurationData) GetPermissionUpdateCheckOk() (*int32, bool)`

GetPermissionUpdateCheckOk returns a tuple with the PermissionUpdateCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionUpdateCheck

`func (o *ConfigurationData) SetPermissionUpdateCheck(v int32)`

SetPermissionUpdateCheck sets PermissionUpdateCheck field to given value.

### HasPermissionUpdateCheck

`func (o *ConfigurationData) HasPermissionUpdateCheck() bool`

HasPermissionUpdateCheck returns a boolean if a field has been set.

### SetPermissionUpdateCheckNil

`func (o *ConfigurationData) SetPermissionUpdateCheckNil(b bool)`

 SetPermissionUpdateCheckNil sets the value for PermissionUpdateCheck to be an explicit nil

### UnsetPermissionUpdateCheck
`func (o *ConfigurationData) UnsetPermissionUpdateCheck()`

UnsetPermissionUpdateCheck ensures that no value is present for PermissionUpdateCheck, not even an explicit nil
### GetLastUpdateCheck

`func (o *ConfigurationData) GetLastUpdateCheck() time.Time`

GetLastUpdateCheck returns the LastUpdateCheck field if non-nil, zero value otherwise.

### GetLastUpdateCheckOk

`func (o *ConfigurationData) GetLastUpdateCheckOk() (*time.Time, bool)`

GetLastUpdateCheckOk returns a tuple with the LastUpdateCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateCheck

`func (o *ConfigurationData) SetLastUpdateCheck(v time.Time)`

SetLastUpdateCheck sets LastUpdateCheck field to given value.

### HasLastUpdateCheck

`func (o *ConfigurationData) HasLastUpdateCheck() bool`

HasLastUpdateCheck returns a boolean if a field has been set.

### GetSingleUserMode

`func (o *ConfigurationData) GetSingleUserMode() bool`

GetSingleUserMode returns the SingleUserMode field if non-nil, zero value otherwise.

### GetSingleUserModeOk

`func (o *ConfigurationData) GetSingleUserModeOk() (*bool, bool)`

GetSingleUserModeOk returns a tuple with the SingleUserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUserMode

`func (o *ConfigurationData) SetSingleUserMode(v bool)`

SetSingleUserMode sets SingleUserMode field to given value.

### HasSingleUserMode

`func (o *ConfigurationData) HasSingleUserMode() bool`

HasSingleUserMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


