# DestinationLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**DeliverySlot** | Pointer to [**DeliverySlot**](DeliverySlot.md) |  | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer | [optional] 

## Methods

### NewDestinationLocation

`func NewDestinationLocation() *DestinationLocation`

NewDestinationLocation instantiates a new DestinationLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationLocationWithDefaults

`func NewDestinationLocationWithDefaults() *DestinationLocation`

NewDestinationLocationWithDefaults instantiates a new DestinationLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationLocation) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationLocation) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationLocation) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDeliverySlot

`func (o *DestinationLocation) GetDeliverySlot() DeliverySlot`

GetDeliverySlot returns the DeliverySlot field if non-nil, zero value otherwise.

### GetDeliverySlotOk

`func (o *DestinationLocation) GetDeliverySlotOk() (*DeliverySlot, bool)`

GetDeliverySlotOk returns a tuple with the DeliverySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySlot

`func (o *DestinationLocation) SetDeliverySlot(v DeliverySlot)`

SetDeliverySlot sets DeliverySlot field to given value.

### HasDeliverySlot

`func (o *DestinationLocation) HasDeliverySlot() bool`

HasDeliverySlot returns a boolean if a field has been set.

### GetCustomerId

`func (o *DestinationLocation) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DestinationLocation) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DestinationLocation) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DestinationLocation) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


