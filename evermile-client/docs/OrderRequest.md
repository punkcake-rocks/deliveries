# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProposalId** | **string** | The proposal ID to create order from | 
**ExternalOrderId** | Pointer to **string** | An external order ID to attach to this order | [optional] [default to ""]
**PickupContactDetails** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**PickupLocationId** | Pointer to **string** | The id of an existing sender location with the sender contact details | [optional] 
**DropoffContactDetails** | [**ContactDetails**](ContactDetails.md) |  | 
**UseCredits** | **bool** | Whether to use available credits to pay for order | [default to false]
**ProofOfDeliveryRequirement** | Pointer to [**[]ProofOfDeliveryRequirement**](ProofOfDeliveryRequirement.md) | Proof of delivery requirements for a parcel, Optional. Will override any POD passed in on the quote | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest(proposalId string, dropoffContactDetails ContactDetails, useCredits bool, ) *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposalId

`func (o *OrderRequest) GetProposalId() string`

GetProposalId returns the ProposalId field if non-nil, zero value otherwise.

### GetProposalIdOk

`func (o *OrderRequest) GetProposalIdOk() (*string, bool)`

GetProposalIdOk returns a tuple with the ProposalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposalId

`func (o *OrderRequest) SetProposalId(v string)`

SetProposalId sets ProposalId field to given value.


### GetExternalOrderId

`func (o *OrderRequest) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *OrderRequest) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *OrderRequest) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *OrderRequest) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.

### GetPickupContactDetails

`func (o *OrderRequest) GetPickupContactDetails() ContactDetails`

GetPickupContactDetails returns the PickupContactDetails field if non-nil, zero value otherwise.

### GetPickupContactDetailsOk

`func (o *OrderRequest) GetPickupContactDetailsOk() (*ContactDetails, bool)`

GetPickupContactDetailsOk returns a tuple with the PickupContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupContactDetails

`func (o *OrderRequest) SetPickupContactDetails(v ContactDetails)`

SetPickupContactDetails sets PickupContactDetails field to given value.

### HasPickupContactDetails

`func (o *OrderRequest) HasPickupContactDetails() bool`

HasPickupContactDetails returns a boolean if a field has been set.

### GetPickupLocationId

`func (o *OrderRequest) GetPickupLocationId() string`

GetPickupLocationId returns the PickupLocationId field if non-nil, zero value otherwise.

### GetPickupLocationIdOk

`func (o *OrderRequest) GetPickupLocationIdOk() (*string, bool)`

GetPickupLocationIdOk returns a tuple with the PickupLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLocationId

`func (o *OrderRequest) SetPickupLocationId(v string)`

SetPickupLocationId sets PickupLocationId field to given value.

### HasPickupLocationId

`func (o *OrderRequest) HasPickupLocationId() bool`

HasPickupLocationId returns a boolean if a field has been set.

### GetDropoffContactDetails

`func (o *OrderRequest) GetDropoffContactDetails() ContactDetails`

GetDropoffContactDetails returns the DropoffContactDetails field if non-nil, zero value otherwise.

### GetDropoffContactDetailsOk

`func (o *OrderRequest) GetDropoffContactDetailsOk() (*ContactDetails, bool)`

GetDropoffContactDetailsOk returns a tuple with the DropoffContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffContactDetails

`func (o *OrderRequest) SetDropoffContactDetails(v ContactDetails)`

SetDropoffContactDetails sets DropoffContactDetails field to given value.


### GetUseCredits

`func (o *OrderRequest) GetUseCredits() bool`

GetUseCredits returns the UseCredits field if non-nil, zero value otherwise.

### GetUseCreditsOk

`func (o *OrderRequest) GetUseCreditsOk() (*bool, bool)`

GetUseCreditsOk returns a tuple with the UseCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCredits

`func (o *OrderRequest) SetUseCredits(v bool)`

SetUseCredits sets UseCredits field to given value.


### GetProofOfDeliveryRequirement

`func (o *OrderRequest) GetProofOfDeliveryRequirement() []ProofOfDeliveryRequirement`

GetProofOfDeliveryRequirement returns the ProofOfDeliveryRequirement field if non-nil, zero value otherwise.

### GetProofOfDeliveryRequirementOk

`func (o *OrderRequest) GetProofOfDeliveryRequirementOk() (*[]ProofOfDeliveryRequirement, bool)`

GetProofOfDeliveryRequirementOk returns a tuple with the ProofOfDeliveryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfDeliveryRequirement

`func (o *OrderRequest) SetProofOfDeliveryRequirement(v []ProofOfDeliveryRequirement)`

SetProofOfDeliveryRequirement sets ProofOfDeliveryRequirement field to given value.

### HasProofOfDeliveryRequirement

`func (o *OrderRequest) HasProofOfDeliveryRequirement() bool`

HasProofOfDeliveryRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


