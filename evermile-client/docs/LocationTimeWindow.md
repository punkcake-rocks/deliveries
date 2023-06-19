# LocationTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOpen** | **bool** | Whether the location is open or not | [default to true]
**OpenWindowStart** | **string** | The time when this location opens in HH:MM format | 
**OpenWindowEnd** | **string** | The time when this location closes in HH:MM format | 

## Methods

### NewLocationTimeWindow

`func NewLocationTimeWindow(isOpen bool, openWindowStart string, openWindowEnd string, ) *LocationTimeWindow`

NewLocationTimeWindow instantiates a new LocationTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationTimeWindowWithDefaults

`func NewLocationTimeWindowWithDefaults() *LocationTimeWindow`

NewLocationTimeWindowWithDefaults instantiates a new LocationTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOpen

`func (o *LocationTimeWindow) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *LocationTimeWindow) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *LocationTimeWindow) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.


### GetOpenWindowStart

`func (o *LocationTimeWindow) GetOpenWindowStart() string`

GetOpenWindowStart returns the OpenWindowStart field if non-nil, zero value otherwise.

### GetOpenWindowStartOk

`func (o *LocationTimeWindow) GetOpenWindowStartOk() (*string, bool)`

GetOpenWindowStartOk returns a tuple with the OpenWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWindowStart

`func (o *LocationTimeWindow) SetOpenWindowStart(v string)`

SetOpenWindowStart sets OpenWindowStart field to given value.


### GetOpenWindowEnd

`func (o *LocationTimeWindow) GetOpenWindowEnd() string`

GetOpenWindowEnd returns the OpenWindowEnd field if non-nil, zero value otherwise.

### GetOpenWindowEndOk

`func (o *LocationTimeWindow) GetOpenWindowEndOk() (*string, bool)`

GetOpenWindowEndOk returns a tuple with the OpenWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWindowEnd

`func (o *LocationTimeWindow) SetOpenWindowEnd(v string)`

SetOpenWindowEnd sets OpenWindowEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


