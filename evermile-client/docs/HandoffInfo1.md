# HandoffInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | [**HandoffContactInfo1**](HandoffContactInfo1.md) |  | 
**HandoffTime** | [**OptionalTimeWindow**](OptionalTimeWindow.md) |  | 
**WalkingDistanceMeters** | **int64** | walking distance in meters | 

## Methods

### NewHandoffInfo1

`func NewHandoffInfo1(contactInfo HandoffContactInfo1, handoffTime OptionalTimeWindow, walkingDistanceMeters int64, ) *HandoffInfo1`

NewHandoffInfo1 instantiates a new HandoffInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandoffInfo1WithDefaults

`func NewHandoffInfo1WithDefaults() *HandoffInfo1`

NewHandoffInfo1WithDefaults instantiates a new HandoffInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *HandoffInfo1) GetContactInfo() HandoffContactInfo1`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *HandoffInfo1) GetContactInfoOk() (*HandoffContactInfo1, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *HandoffInfo1) SetContactInfo(v HandoffContactInfo1)`

SetContactInfo sets ContactInfo field to given value.


### GetHandoffTime

`func (o *HandoffInfo1) GetHandoffTime() OptionalTimeWindow`

GetHandoffTime returns the HandoffTime field if non-nil, zero value otherwise.

### GetHandoffTimeOk

`func (o *HandoffInfo1) GetHandoffTimeOk() (*OptionalTimeWindow, bool)`

GetHandoffTimeOk returns a tuple with the HandoffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffTime

`func (o *HandoffInfo1) SetHandoffTime(v OptionalTimeWindow)`

SetHandoffTime sets HandoffTime field to given value.


### GetWalkingDistanceMeters

`func (o *HandoffInfo1) GetWalkingDistanceMeters() int64`

GetWalkingDistanceMeters returns the WalkingDistanceMeters field if non-nil, zero value otherwise.

### GetWalkingDistanceMetersOk

`func (o *HandoffInfo1) GetWalkingDistanceMetersOk() (*int64, bool)`

GetWalkingDistanceMetersOk returns a tuple with the WalkingDistanceMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalkingDistanceMeters

`func (o *HandoffInfo1) SetWalkingDistanceMeters(v int64)`

SetWalkingDistanceMeters sets WalkingDistanceMeters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


