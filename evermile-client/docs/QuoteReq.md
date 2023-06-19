# QuoteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupLocations** | **[]string** | A list of location ids from which the items can be picked up. | 
**DestinationLocations** | [**[]DestinationLocation**](DestinationLocation.md) | A list of locations to which the items can be delivered | 
**Parcels** | [**[]Parcel**](Parcel.md) | The list of parcels to pick up in this delivery from any **single** source&lt;br&gt;(must all be available at all of the pickup locations in *pickupLocations*) | 
**Handling** | Pointer to [**[]HandlingInstruction**](HandlingInstruction.md) | Handling instructions for a parcel | [optional] 
**ProofOfDeliveryRequirement** | Pointer to [**[]ProofOfDeliveryRequirement**](ProofOfDeliveryRequirement.md) | Proof of delivery requirements for a parcel | [optional] [default to ["none"]]
**Instructions** | Pointer to **string** | Additional handling instructions for the packages | [optional] 
**OriginalOrderId** | Pointer to **string** | The id of the order that this quote is based on, usually for cases when editing an existing order | [optional] 
**ExcludeCarriers** | Pointer to [**[]CarriersType**](CarriersType.md) | An array of carriers names to exclude from the proposals results | [optional] 
**ExcludedVehicleTypes** | Pointer to [**[]VehicleType**](VehicleType.md) | Which vehicle types to exclude from the request. **note** excluding vehicle types may result in higher prices. | [optional] 
**AdditionalProposalTypes** | Pointer to [**[]ProposalType**](ProposalType.md) |  | [optional] 

## Methods

### NewQuoteReq

`func NewQuoteReq(pickupLocations []string, destinationLocations []DestinationLocation, parcels []Parcel, ) *QuoteReq`

NewQuoteReq instantiates a new QuoteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteReqWithDefaults

`func NewQuoteReqWithDefaults() *QuoteReq`

NewQuoteReqWithDefaults instantiates a new QuoteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupLocations

`func (o *QuoteReq) GetPickupLocations() []string`

GetPickupLocations returns the PickupLocations field if non-nil, zero value otherwise.

### GetPickupLocationsOk

`func (o *QuoteReq) GetPickupLocationsOk() (*[]string, bool)`

GetPickupLocationsOk returns a tuple with the PickupLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLocations

`func (o *QuoteReq) SetPickupLocations(v []string)`

SetPickupLocations sets PickupLocations field to given value.


### GetDestinationLocations

`func (o *QuoteReq) GetDestinationLocations() []DestinationLocation`

GetDestinationLocations returns the DestinationLocations field if non-nil, zero value otherwise.

### GetDestinationLocationsOk

`func (o *QuoteReq) GetDestinationLocationsOk() (*[]DestinationLocation, bool)`

GetDestinationLocationsOk returns a tuple with the DestinationLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocations

`func (o *QuoteReq) SetDestinationLocations(v []DestinationLocation)`

SetDestinationLocations sets DestinationLocations field to given value.


### GetParcels

`func (o *QuoteReq) GetParcels() []Parcel`

GetParcels returns the Parcels field if non-nil, zero value otherwise.

### GetParcelsOk

`func (o *QuoteReq) GetParcelsOk() (*[]Parcel, bool)`

GetParcelsOk returns a tuple with the Parcels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcels

`func (o *QuoteReq) SetParcels(v []Parcel)`

SetParcels sets Parcels field to given value.


### GetHandling

`func (o *QuoteReq) GetHandling() []HandlingInstruction`

GetHandling returns the Handling field if non-nil, zero value otherwise.

### GetHandlingOk

`func (o *QuoteReq) GetHandlingOk() (*[]HandlingInstruction, bool)`

GetHandlingOk returns a tuple with the Handling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandling

`func (o *QuoteReq) SetHandling(v []HandlingInstruction)`

SetHandling sets Handling field to given value.

### HasHandling

`func (o *QuoteReq) HasHandling() bool`

HasHandling returns a boolean if a field has been set.

### GetProofOfDeliveryRequirement

`func (o *QuoteReq) GetProofOfDeliveryRequirement() []ProofOfDeliveryRequirement`

GetProofOfDeliveryRequirement returns the ProofOfDeliveryRequirement field if non-nil, zero value otherwise.

### GetProofOfDeliveryRequirementOk

`func (o *QuoteReq) GetProofOfDeliveryRequirementOk() (*[]ProofOfDeliveryRequirement, bool)`

GetProofOfDeliveryRequirementOk returns a tuple with the ProofOfDeliveryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryRequirement

`func (o *QuoteReq) SetProofOfDeliveryRequirement(v []ProofOfDeliveryRequirement)`

SetProofOfDeliveryRequirement sets ProofOfDeliveryRequirement field to given value.

### HasProofOfDeliveryRequirement

`func (o *QuoteReq) HasProofOfDeliveryRequirement() bool`

HasProofOfDeliveryRequirement returns a boolean if a field has been set.

### GetInstructions

`func (o *QuoteReq) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *QuoteReq) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *QuoteReq) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *QuoteReq) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetOriginalOrderId

`func (o *QuoteReq) GetOriginalOrderId() string`

GetOriginalOrderId returns the OriginalOrderId field if non-nil, zero value otherwise.

### GetOriginalOrderIdOk

`func (o *QuoteReq) GetOriginalOrderIdOk() (*string, bool)`

GetOriginalOrderIdOk returns a tuple with the OriginalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalOrderId

`func (o *QuoteReq) SetOriginalOrderId(v string)`

SetOriginalOrderId sets OriginalOrderId field to given value.

### HasOriginalOrderId

`func (o *QuoteReq) HasOriginalOrderId() bool`

HasOriginalOrderId returns a boolean if a field has been set.

### GetExcludeCarriers

`func (o *QuoteReq) GetExcludeCarriers() []CarriersType`

GetExcludeCarriers returns the ExcludeCarriers field if non-nil, zero value otherwise.

### GetExcludeCarriersOk

`func (o *QuoteReq) GetExcludeCarriersOk() (*[]CarriersType, bool)`

GetExcludeCarriersOk returns a tuple with the ExcludeCarriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCarriers

`func (o *QuoteReq) SetExcludeCarriers(v []CarriersType)`

SetExcludeCarriers sets ExcludeCarriers field to given value.

### HasExcludeCarriers

`func (o *QuoteReq) HasExcludeCarriers() bool`

HasExcludeCarriers returns a boolean if a field has been set.

### GetExcludedVehicleTypes

`func (o *QuoteReq) GetExcludedVehicleTypes() []VehicleType`

GetExcludedVehicleTypes returns the ExcludedVehicleTypes field if non-nil, zero value otherwise.

### GetExcludedVehicleTypesOk

`func (o *QuoteReq) GetExcludedVehicleTypesOk() (*[]VehicleType, bool)`

GetExcludedVehicleTypesOk returns a tuple with the ExcludedVehicleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVehicleTypes

`func (o *QuoteReq) SetExcludedVehicleTypes(v []VehicleType)`

SetExcludedVehicleTypes sets ExcludedVehicleTypes field to given value.

### HasExcludedVehicleTypes

`func (o *QuoteReq) HasExcludedVehicleTypes() bool`

HasExcludedVehicleTypes returns a boolean if a field has been set.

### GetAdditionalProposalTypes

`func (o *QuoteReq) GetAdditionalProposalTypes() []ProposalType`

GetAdditionalProposalTypes returns the AdditionalProposalTypes field if non-nil, zero value otherwise.

### GetAdditionalProposalTypesOk

`func (o *QuoteReq) GetAdditionalProposalTypesOk() (*[]ProposalType, bool)`

GetAdditionalProposalTypesOk returns a tuple with the AdditionalProposalTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalProposalTypes

`func (o *QuoteReq) SetAdditionalProposalTypes(v []ProposalType)`

SetAdditionalProposalTypes sets AdditionalProposalTypes field to given value.

### HasAdditionalProposalTypes

`func (o *QuoteReq) HasAdditionalProposalTypes() bool`

HasAdditionalProposalTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


