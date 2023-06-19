# LocationWithId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The location&#39;s evermile ID. Used for updating a location (if provided) | 
**MerchantId** | **string** | The customer ID to which this location belongs | 
**Name** | **string** | An informative name for this location | 
**Address** | [**Address**](Address.md) |  | 
**Phone** | Pointer to **string** | A phone number for this location (relevant when location is used as sender_info) | [optional] 
**Email** | Pointer to **string** | An email of the contact person for this location (relevant when location is used as sender_info) | [optional] 
**Notes** | Pointer to **string** | Additional notes to describe this location | [optional] 
**IsDefault** | Pointer to **bool** | True if this is the default location | [optional] 
**StoreId** | Pointer to **string** | A store ID associated with this location | [optional] 
**Schedule** | Pointer to [**LocationWeeklySchedule**](LocationWeeklySchedule.md) |  | [optional] 

## Methods

### NewLocationWithId

`func NewLocationWithId(id string, merchantId string, name string, address Address, ) *LocationWithId`

NewLocationWithId instantiates a new LocationWithId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithIdWithDefaults

`func NewLocationWithIdWithDefaults() *LocationWithId`

NewLocationWithIdWithDefaults instantiates a new LocationWithId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationWithId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationWithId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationWithId) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *LocationWithId) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *LocationWithId) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *LocationWithId) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetName

`func (o *LocationWithId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationWithId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationWithId) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *LocationWithId) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LocationWithId) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LocationWithId) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *LocationWithId) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LocationWithId) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LocationWithId) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LocationWithId) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *LocationWithId) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LocationWithId) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LocationWithId) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LocationWithId) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNotes

`func (o *LocationWithId) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *LocationWithId) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *LocationWithId) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *LocationWithId) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsDefault

`func (o *LocationWithId) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *LocationWithId) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *LocationWithId) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *LocationWithId) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetStoreId

`func (o *LocationWithId) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *LocationWithId) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *LocationWithId) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *LocationWithId) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetSchedule

`func (o *LocationWithId) GetSchedule() LocationWeeklySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LocationWithId) GetScheduleOk() (*LocationWeeklySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LocationWithId) SetSchedule(v LocationWeeklySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *LocationWithId) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


