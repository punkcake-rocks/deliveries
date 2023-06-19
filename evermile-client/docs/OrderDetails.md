# OrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the order | 
**ExternalOrderId** | **string** | An external order ID attached to this order | 
**StoreId** | Pointer to **string** | The ID of the e-commerce store for this order (if exists) | [optional] 
**StoreName** | Pointer to **string** | The name of the store for this order (if exists) | [optional] 
**Status** | [**OrderExtendedStatus**](OrderExtendedStatus.md) |  | 
**CommunicatedPickup** | [**DeliverySlot**](DeliverySlot.md) |  | 
**CommunicatedDropoff** | [**DeliverySlot**](DeliverySlot.md) |  | 
**ActualPickup** | Pointer to **time.Time** | The actual pickup time in ISO8601 format (for packages picked up) | [optional] 
**ActualDropoff** | Pointer to **time.Time** | The actual dropoff time in ISO8601 format (for packages dropped off) | [optional] 
**PickupEta** | Pointer to **time.Time** | The expected pickup time (if known) | [optional] 
**DropoffEta** | Pointer to **time.Time** | The expected dropoff time (if known) | [optional] 
**CourierLocation** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 
**Origin** | [**Address**](Address.md) |  | 
**Destination** | [**Address**](Address.md) |  | 
**Customer** | [**Customer**](Customer.md) |  | 
**DeliveryCompany** | Pointer to **string** | The company that will perform the delivery (if known) | [optional] 
**Parcels** | [**[]Parcel**](Parcel.md) | The list of parcels in this order | 
**Value** | [**Price**](Price.md) |  | 
**Instructions** | Pointer to **string** | Additional handling instructions for the packages | [optional] 
**Handling** | Pointer to [**[]HandlingInstruction**](HandlingInstruction.md) | Handling instructions for a parcel | [optional] 
**ProofOfDeliveryRequirement** | Pointer to [**[]ProofOfDeliveryRequirement**](ProofOfDeliveryRequirement.md) | Proof of delivery requirements for a parcel | [optional] [default to ["none"]]
**TrackingUrl** | Pointer to **string** | The carrier tracking URL for this order (if exists) | [optional] 
**CustomerTrackingUrl** | Pointer to **string** | The Evermile customer tracking URL for this order (if exists) | [optional] 
**OrderTrackingInfo** | Pointer to [**OrderTrackingInfo**](OrderTrackingInfo.md) |  | [optional] 
**HandoffType** | [**HandoffType**](HandoffType.md) |  | 
**HandoffInfo** | Pointer to [**HandoffInfo**](HandoffInfo.md) |  | [optional] 
**UsedCredits** | **bool** | Whether credits were used to pay for this order | 
**ProofOfDelivery** | Pointer to [**[]ProofOfDelivery**](ProofOfDelivery.md) | Proof of delivery for an order | [optional] 
**Notes** | Pointer to **[]string** | Courier notes from each delivery in this order | [optional] 
**MerchantName** | **string** | The name of the merchant for this order | 
**LabelRequired** | **bool** | Whether label is required | 

## Methods

### NewOrderDetails

`func NewOrderDetails(id string, externalOrderId string, status OrderExtendedStatus, communicatedPickup DeliverySlot, communicatedDropoff DeliverySlot, origin Address, destination Address, customer Customer, parcels []Parcel, value Price, handoffType HandoffType, usedCredits bool, merchantName string, labelRequired bool, ) *OrderDetails`

NewOrderDetails instantiates a new OrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailsWithDefaults

`func NewOrderDetailsWithDefaults() *OrderDetails`

NewOrderDetailsWithDefaults instantiates a new OrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderDetails) SetId(v string)`

SetId sets Id field to given value.


### GetExternalOrderId

`func (o *OrderDetails) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *OrderDetails) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *OrderDetails) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.


### GetStoreId

`func (o *OrderDetails) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderDetails) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderDetails) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderDetails) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetStoreName

`func (o *OrderDetails) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *OrderDetails) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *OrderDetails) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *OrderDetails) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetStatus

`func (o *OrderDetails) GetStatus() OrderExtendedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetails) GetStatusOk() (*OrderExtendedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetails) SetStatus(v OrderExtendedStatus)`

SetStatus sets Status field to given value.


### GetCommunicatedPickup

`func (o *OrderDetails) GetCommunicatedPickup() DeliverySlot`

GetCommunicatedPickup returns the CommunicatedPickup field if non-nil, zero value otherwise.

### GetCommunicatedPickupOk

`func (o *OrderDetails) GetCommunicatedPickupOk() (*DeliverySlot, bool)`

GetCommunicatedPickupOk returns a tuple with the CommunicatedPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicatedPickup

`func (o *OrderDetails) SetCommunicatedPickup(v DeliverySlot)`

SetCommunicatedPickup sets CommunicatedPickup field to given value.


### GetCommunicatedDropoff

`func (o *OrderDetails) GetCommunicatedDropoff() DeliverySlot`

GetCommunicatedDropoff returns the CommunicatedDropoff field if non-nil, zero value otherwise.

### GetCommunicatedDropoffOk

`func (o *OrderDetails) GetCommunicatedDropoffOk() (*DeliverySlot, bool)`

GetCommunicatedDropoffOk returns a tuple with the CommunicatedDropoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicatedDropoff

`func (o *OrderDetails) SetCommunicatedDropoff(v DeliverySlot)`

SetCommunicatedDropoff sets CommunicatedDropoff field to given value.


### GetActualPickup

`func (o *OrderDetails) GetActualPickup() time.Time`

GetActualPickup returns the ActualPickup field if non-nil, zero value otherwise.

### GetActualPickupOk

`func (o *OrderDetails) GetActualPickupOk() (*time.Time, bool)`

GetActualPickupOk returns a tuple with the ActualPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickup

`func (o *OrderDetails) SetActualPickup(v time.Time)`

SetActualPickup sets ActualPickup field to given value.

### HasActualPickup

`func (o *OrderDetails) HasActualPickup() bool`

HasActualPickup returns a boolean if a field has been set.

### GetActualDropoff

`func (o *OrderDetails) GetActualDropoff() time.Time`

GetActualDropoff returns the ActualDropoff field if non-nil, zero value otherwise.

### GetActualDropoffOk

`func (o *OrderDetails) GetActualDropoffOk() (*time.Time, bool)`

GetActualDropoffOk returns a tuple with the ActualDropoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDropoff

`func (o *OrderDetails) SetActualDropoff(v time.Time)`

SetActualDropoff sets ActualDropoff field to given value.

### HasActualDropoff

`func (o *OrderDetails) HasActualDropoff() bool`

HasActualDropoff returns a boolean if a field has been set.

### GetPickupEta

`func (o *OrderDetails) GetPickupEta() time.Time`

GetPickupEta returns the PickupEta field if non-nil, zero value otherwise.

### GetPickupEtaOk

`func (o *OrderDetails) GetPickupEtaOk() (*time.Time, bool)`

GetPickupEtaOk returns a tuple with the PickupEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEta

`func (o *OrderDetails) SetPickupEta(v time.Time)`

SetPickupEta sets PickupEta field to given value.

### HasPickupEta

`func (o *OrderDetails) HasPickupEta() bool`

HasPickupEta returns a boolean if a field has been set.

### GetDropoffEta

`func (o *OrderDetails) GetDropoffEta() time.Time`

GetDropoffEta returns the DropoffEta field if non-nil, zero value otherwise.

### GetDropoffEtaOk

`func (o *OrderDetails) GetDropoffEtaOk() (*time.Time, bool)`

GetDropoffEtaOk returns a tuple with the DropoffEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffEta

`func (o *OrderDetails) SetDropoffEta(v time.Time)`

SetDropoffEta sets DropoffEta field to given value.

### HasDropoffEta

`func (o *OrderDetails) HasDropoffEta() bool`

HasDropoffEta returns a boolean if a field has been set.

### GetCourierLocation

`func (o *OrderDetails) GetCourierLocation() GeoLocation`

GetCourierLocation returns the CourierLocation field if non-nil, zero value otherwise.

### GetCourierLocationOk

`func (o *OrderDetails) GetCourierLocationOk() (*GeoLocation, bool)`

GetCourierLocationOk returns a tuple with the CourierLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierLocation

`func (o *OrderDetails) SetCourierLocation(v GeoLocation)`

SetCourierLocation sets CourierLocation field to given value.

### HasCourierLocation

`func (o *OrderDetails) HasCourierLocation() bool`

HasCourierLocation returns a boolean if a field has been set.

### GetOrigin

`func (o *OrderDetails) GetOrigin() Address`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *OrderDetails) GetOriginOk() (*Address, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *OrderDetails) SetOrigin(v Address)`

SetOrigin sets Origin field to given value.


### GetDestination

`func (o *OrderDetails) GetDestination() Address`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *OrderDetails) GetDestinationOk() (*Address, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *OrderDetails) SetDestination(v Address)`

SetDestination sets Destination field to given value.


### GetCustomer

`func (o *OrderDetails) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderDetails) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderDetails) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetDeliveryCompany

`func (o *OrderDetails) GetDeliveryCompany() string`

GetDeliveryCompany returns the DeliveryCompany field if non-nil, zero value otherwise.

### GetDeliveryCompanyOk

`func (o *OrderDetails) GetDeliveryCompanyOk() (*string, bool)`

GetDeliveryCompanyOk returns a tuple with the DeliveryCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCompany

`func (o *OrderDetails) SetDeliveryCompany(v string)`

SetDeliveryCompany sets DeliveryCompany field to given value.

### HasDeliveryCompany

`func (o *OrderDetails) HasDeliveryCompany() bool`

HasDeliveryCompany returns a boolean if a field has been set.

### GetParcels

`func (o *OrderDetails) GetParcels() []Parcel`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *OrderDetails) GetParcelsOk() (*[]Parcel, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *OrderDetails) SetParcels(v []Parcel)`

SetParcels sets Parcels field to given value.


### GetValue

`func (o *OrderDetails) GetValue() Price`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrderDetails) GetValueOk() (*Price, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrderDetails) SetValue(v Price)`

SetValue sets Value field to given value.


### GetInstructions

`func (o *OrderDetails) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *OrderDetails) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *OrderDetails) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *OrderDetails) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetHandling

`func (o *OrderDetails) GetHandling() []HandlingInstruction`

GetHandling returns the Handling field if non-nil, zero value otherwise.

### GetHandlingOk

`func (o *OrderDetails) GetHandlingOk() (*[]HandlingInstruction, bool)`

GetHandlingOk returns a tuple with the Handling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandling

`func (o *OrderDetails) SetHandling(v []HandlingInstruction)`

SetHandling sets Handling field to given value.

### HasHandling

`func (o *OrderDetails) HasHandling() bool`

HasHandling returns a boolean if a field has been set.

### GetProofOfDeliveryRequirement

`func (o *OrderDetails) GetProofOfDeliveryRequirement() []ProofOfDeliveryRequirement`

GetProofOfDeliveryRequirement returns the ProofOfDeliveryRequirement field if non-nil, zero value otherwise.

### GetProofOfDeliveryRequirementOk

`func (o *OrderDetails) GetProofOfDeliveryRequirementOk() (*[]ProofOfDeliveryRequirement, bool)`

GetProofOfDeliveryRequirementOk returns a tuple with the ProofOfDeliveryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryRequirement

`func (o *OrderDetails) SetProofOfDeliveryRequirement(v []ProofOfDeliveryRequirement)`

SetProofOfDeliveryRequirement sets ProofOfDeliveryRequirement field to given value.

### HasProofOfDeliveryRequirement

`func (o *OrderDetails) HasProofOfDeliveryRequirement() bool`

HasProofOfDeliveryRequirement returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *OrderDetails) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *OrderDetails) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *OrderDetails) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *OrderDetails) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetCustomerTrackingUrl

`func (o *OrderDetails) GetCustomerTrackingUrl() string`

GetCustomerTrackingUrl returns the CustomerTrackingUrl field if non-nil, zero value otherwise.

### GetCustomerTrackingUrlOk

`func (o *OrderDetails) GetCustomerTrackingUrlOk() (*string, bool)`

GetCustomerTrackingUrlOk returns a tuple with the CustomerTrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTrackingUrl

`func (o *OrderDetails) SetCustomerTrackingUrl(v string)`

SetCustomerTrackingUrl sets CustomerTrackingUrl field to given value.

### HasCustomerTrackingUrl

`func (o *OrderDetails) HasCustomerTrackingUrl() bool`

HasCustomerTrackingUrl returns a boolean if a field has been set.

### GetOrderTrackingInfo

`func (o *OrderDetails) GetOrderTrackingInfo() OrderTrackingInfo`

GetOrderTrackingInfo returns the OrderTrackingInfo field if non-nil, zero value otherwise.

### GetOrderTrackingInfoOk

`func (o *OrderDetails) GetOrderTrackingInfoOk() (*OrderTrackingInfo, bool)`

GetOrderTrackingInfoOk returns a tuple with the OrderTrackingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTrackingInfo

`func (o *OrderDetails) SetOrderTrackingInfo(v OrderTrackingInfo)`

SetOrderTrackingInfo sets OrderTrackingInfo field to given value.

### HasOrderTrackingInfo

`func (o *OrderDetails) HasOrderTrackingInfo() bool`

HasOrderTrackingInfo returns a boolean if a field has been set.

### GetHandoffType

`func (o *OrderDetails) GetHandoffType() HandoffType`

GetHandoffType returns the HandoffType field if non-nil, zero value otherwise.

### GetHandoffTypeOk

`func (o *OrderDetails) GetHandoffTypeOk() (*HandoffType, bool)`

GetHandoffTypeOk returns a tuple with the HandoffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffType

`func (o *OrderDetails) SetHandoffType(v HandoffType)`

SetHandoffType sets HandoffType field to given value.


### GetHandoffInfo

`func (o *OrderDetails) GetHandoffInfo() HandoffInfo`

GetHandoffInfo returns the HandoffInfo field if non-nil, zero value otherwise.

### GetHandoffInfoOk

`func (o *OrderDetails) GetHandoffInfoOk() (*HandoffInfo, bool)`

GetHandoffInfoOk returns a tuple with the HandoffInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffInfo

`func (o *OrderDetails) SetHandoffInfo(v HandoffInfo)`

SetHandoffInfo sets HandoffInfo field to given value.

### HasHandoffInfo

`func (o *OrderDetails) HasHandoffInfo() bool`

HasHandoffInfo returns a boolean if a field has been set.

### GetUsedCredits

`func (o *OrderDetails) GetUsedCredits() bool`

GetUsedCredits returns the UsedCredits field if non-nil, zero value otherwise.

### GetUsedCreditsOk

`func (o *OrderDetails) GetUsedCreditsOk() (*bool, bool)`

GetUsedCreditsOk returns a tuple with the UsedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCredits

`func (o *OrderDetails) SetUsedCredits(v bool)`

SetUsedCredits sets UsedCredits field to given value.


### GetProofOfDelivery

`func (o *OrderDetails) GetProofOfDelivery() []ProofOfDelivery`

GetProofOfDelivery returns the ProofOfDelivery field if non-nil, zero value otherwise.

### GetProofOfDeliveryOk

`func (o *OrderDetails) GetProofOfDeliveryOk() (*[]ProofOfDelivery, bool)`

GetProofOfDeliveryOk returns a tuple with the ProofOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDelivery

`func (o *OrderDetails) SetProofOfDelivery(v []ProofOfDelivery)`

SetProofOfDelivery sets ProofOfDelivery field to given value.

### HasProofOfDelivery

`func (o *OrderDetails) HasProofOfDelivery() bool`

HasProofOfDelivery returns a boolean if a field has been set.

### GetNotes

`func (o *OrderDetails) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderDetails) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderDetails) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMerchantName

`func (o *OrderDetails) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *OrderDetails) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *OrderDetails) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.


### GetLabelRequired

`func (o *OrderDetails) GetLabelRequired() bool`

GetLabelRequired returns the LabelRequired field if non-nil, zero value otherwise.

### GetLabelRequiredOk

`func (o *OrderDetails) GetLabelRequiredOk() (*bool, bool)`

GetLabelRequiredOk returns a tuple with the LabelRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelRequired

`func (o *OrderDetails) SetLabelRequired(v bool)`

SetLabelRequired sets LabelRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


