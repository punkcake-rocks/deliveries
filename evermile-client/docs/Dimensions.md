# Dimensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this set of dimensions | [optional] 
**LengthCm** | **float32** | Length of the parcel in centimeters | 
**WidthCm** | **float32** | Width of the parcel in centimeters | 
**HeightCm** | **float32** | Height of the parcel in centimeters | 

## Methods

### NewDimensions

`func NewDimensions(lengthCm float32, widthCm float32, heightCm float32, ) *Dimensions`

NewDimensions instantiates a new Dimensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionsWithDefaults

`func NewDimensionsWithDefaults() *Dimensions`

NewDimensionsWithDefaults instantiates a new Dimensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Dimensions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dimensions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dimensions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dimensions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLengthCm

`func (o *Dimensions) GetLengthCm() float32`

GetLengthCm returns the LengthCm field if non-nil, zero value otherwise.

### GetLengthCmOk

`func (o *Dimensions) GetLengthCmOk() (*float32, bool)`

GetLengthCmOk returns a tuple with the LengthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthCm

`func (o *Dimensions) SetLengthCm(v float32)`

SetLengthCm sets LengthCm field to given value.


### GetWidthCm

`func (o *Dimensions) GetWidthCm() float32`

GetWidthCm returns the WidthCm field if non-nil, zero value otherwise.

### GetWidthCmOk

`func (o *Dimensions) GetWidthCmOk() (*float32, bool)`

GetWidthCmOk returns a tuple with the WidthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthCm

`func (o *Dimensions) SetWidthCm(v float32)`

SetWidthCm sets WidthCm field to given value.


### GetHeightCm

`func (o *Dimensions) GetHeightCm() float32`

GetHeightCm returns the HeightCm field if non-nil, zero value otherwise.

### GetHeightCmOk

`func (o *Dimensions) GetHeightCmOk() (*float32, bool)`

GetHeightCmOk returns a tuple with the HeightCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightCm

`func (o *Dimensions) SetHeightCm(v float32)`

SetHeightCm sets HeightCm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


