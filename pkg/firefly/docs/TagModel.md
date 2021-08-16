# TagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tag** | **string** | The tag | 
**Date** | Pointer to **NullableTime** | The date to which the tag is applicable. | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Latitude of the tag&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Latitude of the tag&#39;s location, if applicable. Can be used to draw a map. | [optional] 
**ZoomLevel** | Pointer to **NullableInt32** | Zoom level for the map, if drawn. This to set the box right. Unfortunately this is a proprietary value because each map provider has different zoom levels. | [optional] 

## Methods

### NewTagModel

`func NewTagModel(tag string, ) *TagModel`

NewTagModel instantiates a new TagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelWithDefaults

`func NewTagModelWithDefaults() *TagModel`

NewTagModelWithDefaults instantiates a new TagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TagModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TagModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TagModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TagModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTag

`func (o *TagModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagModel) SetTag(v string)`

SetTag sets Tag field to given value.


### GetDate

`func (o *TagModel) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TagModel) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TagModel) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *TagModel) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *TagModel) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *TagModel) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetDescription

`func (o *TagModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TagModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TagModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TagModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TagModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TagModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLatitude

`func (o *TagModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TagModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TagModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TagModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TagModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TagModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *TagModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TagModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TagModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TagModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TagModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TagModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetZoomLevel

`func (o *TagModel) GetZoomLevel() int32`

GetZoomLevel returns the ZoomLevel field if non-nil, zero value otherwise.

### GetZoomLevelOk

`func (o *TagModel) GetZoomLevelOk() (*int32, bool)`

GetZoomLevelOk returns a tuple with the ZoomLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomLevel

`func (o *TagModel) SetZoomLevel(v int32)`

SetZoomLevel sets ZoomLevel field to given value.

### HasZoomLevel

`func (o *TagModel) HasZoomLevel() bool`

HasZoomLevel returns a boolean if a field has been set.

### SetZoomLevelNil

`func (o *TagModel) SetZoomLevelNil(b bool)`

 SetZoomLevelNil sets the value for ZoomLevel to be an explicit nil

### UnsetZoomLevel
`func (o *TagModel) UnsetZoomLevel()`

UnsetZoomLevel ensures that no value is present for ZoomLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


