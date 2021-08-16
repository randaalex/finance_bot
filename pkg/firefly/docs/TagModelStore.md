# TagModelStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The tag | 
**Date** | Pointer to **NullableString** | The date to which the tag is applicable. | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Latitude of the tag&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Latitude of the tag&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**ZoomLevel** | Pointer to **NullableInt32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. | [optional] 

## Methods

### NewTagModelStore

`func NewTagModelStore(tag string, ) *TagModelStore`

NewTagModelStore instantiates a new TagModelStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelStoreWithDefaults

`func NewTagModelStoreWithDefaults() *TagModelStore`

NewTagModelStoreWithDefaults instantiates a new TagModelStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *TagModelStore) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagModelStore) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagModelStore) SetTag(v string)`

SetTag sets Tag field to given value.


### GetDate

`func (o *TagModelStore) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TagModelStore) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TagModelStore) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TagModelStore) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *TagModelStore) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *TagModelStore) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *TagModelStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagModelStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagModelStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagModelStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TagModelStore) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TagModelStore) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLatitude

`func (o *TagModelStore) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TagModelStore) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TagModelStore) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TagModelStore) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TagModelStore) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TagModelStore) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *TagModelStore) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TagModelStore) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TagModelStore) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TagModelStore) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TagModelStore) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TagModelStore) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetZoomLevel

`func (o *TagModelStore) GetZoomLevel() int32`

GetZoomLevel returns the ZoomLevel field if non-nil, zero value otherwise.

### GetZoomLevelOk

`func (o *TagModelStore) GetZoomLevelOk() (*int32, bool)`

GetZoomLevelOk returns a tuple with the ZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomLevel

`func (o *TagModelStore) SetZoomLevel(v int32)`

SetZoomLevel sets ZoomLevel field to given value.

### HasZoomLevel

`func (o *TagModelStore) HasZoomLevel() bool`

HasZoomLevel returns a boolean if a field has been set.

### SetZoomLevelNil

`func (o *TagModelStore) SetZoomLevelNil(b bool)`

 SetZoomLevelNil sets the value for ZoomLevel to be an explicit nil

### UnsetZoomLevel
`func (o *TagModelStore) UnsetZoomLevel()`

UnsetZoomLevel ensures that no value is present for ZoomLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


