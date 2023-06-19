# Location1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | An informative name for this location | 
**Address** | [**Address1**](Address1.md) |  | 
**Phone** | Pointer to **string** | A phone number for this location (relevant when location is used as sender_info) | [optional] 
**Email** | Pointer to **string** | An email of the contact person for this location (relevant when location is used as sender_info) | [optional] 
**Notes** | Pointer to **string** | Additional notes to describe this location | [optional] 
**IsDefault** | Pointer to **bool** | True if this is the default location | [optional] 
**Id** | Pointer to **string** | The location&#39;s evermile ID. Used for updating a location (if provided) | [optional] 
**StoreId** | Pointer to **string** | A store ID associated with this location | [optional] 
**Schedule** | Pointer to [**LocationWeeklySchedule1**](LocationWeeklySchedule1.md) |  | [optional] 

## Methods

### NewLocation1

`func NewLocation1(name string, address Address1, ) *Location1`

NewLocation1 instantiates a new Location1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocation1WithDefaults

`func NewLocation1WithDefaults() *Location1`

NewLocation1WithDefaults instantiates a new Location1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Location1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location1) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *Location1) GetAddress() Address1`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Location1) GetAddressOk() (*Address1, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Location1) SetAddress(v Address1)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *Location1) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Location1) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Location1) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Location1) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Location1) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Location1) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Location1) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Location1) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *Location1) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Location1) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Location1) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Location1) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsDefault

`func (o *Location1) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Location1) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Location1) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Location1) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetId

`func (o *Location1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Location1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStoreId

`func (o *Location1) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *Location1) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *Location1) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *Location1) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetSchedule

`func (o *Location1) GetSchedule() LocationWeeklySchedule1`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Location1) GetScheduleOk() (*LocationWeeklySchedule1, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Location1) SetSchedule(v LocationWeeklySchedule1)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Location1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


