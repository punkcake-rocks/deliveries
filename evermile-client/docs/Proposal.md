# Proposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The proposal ID | 
**Price** | [**Price**](Price.md) |  | 
**PriceVat** | Pointer to [**Price**](Price.md) |  | [optional] 
**CreditBack** | Pointer to **int64** | The amount of credits this proposal awards, in credit cents | [optional] 
**EstimatedPickup** | [**DeliverySlot**](DeliverySlot.md) |  | 
**EstimatedDropoff** | [**DeliverySlot**](DeliverySlot.md) |  | 
**DropoffWindowType** | [**DropoffWindowTypeEnum**](DropoffWindowTypeEnum.md) |  | 
**Origin** | [**Address**](Address.md) |  | 
**PickupLocationId** | Pointer to **string** | The id of an existing sender location with the sender contact details (optional) | [optional] 
**Destination** | [**Address**](Address.md) |  | 
**DeliveryCompany** | Pointer to **string** | The company that will perform the delivery (if known) | [optional] 
**Labels** | [**[]ProposalLabelEnum**](ProposalLabelEnum.md) | An array of labels for this proposal | 
**PriceLevel** | Pointer to **int32** |  | [optional] 
**HandoffInfo** | Pointer to [**HandoffInfo**](HandoffInfo.md) |  | [optional] 
**Type** | Pointer to [**ProposalType**](ProposalType.md) |  | [optional] 

## Methods

### NewProposal

`func NewProposal(id string, price Price, estimatedPickup DeliverySlot, estimatedDropoff DeliverySlot, dropoffWindowType DropoffWindowTypeEnum, origin Address, destination Address, labels []ProposalLabelEnum, ) *Proposal`

NewProposal instantiates a new Proposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProposalWithDefaults

`func NewProposalWithDefaults() *Proposal`

NewProposalWithDefaults instantiates a new Proposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Proposal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Proposal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Proposal) SetId(v string)`

SetId sets Id field to given value.


### GetPrice

`func (o *Proposal) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Proposal) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Proposal) SetPrice(v Price)`

SetPrice sets Price field to given value.


### GetPriceVat

`func (o *Proposal) GetPriceVat() Price`

GetPriceVat returns the PriceVat field if non-nil, zero value otherwise.

### GetPriceVatOk

`func (o *Proposal) GetPriceVatOk() (*Price, bool)`

GetPriceVatOk returns a tuple with the PriceVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVat

`func (o *Proposal) SetPriceVat(v Price)`

SetPriceVat sets PriceVat field to given value.

### HasPriceVat

`func (o *Proposal) HasPriceVat() bool`

HasPriceVat returns a boolean if a field has been set.

### GetCreditBack

`func (o *Proposal) GetCreditBack() int64`

GetCreditBack returns the CreditBack field if non-nil, zero value otherwise.

### GetCreditBackOk

`func (o *Proposal) GetCreditBackOk() (*int64, bool)`

GetCreditBackOk returns a tuple with the CreditBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBack

`func (o *Proposal) SetCreditBack(v int64)`

SetCreditBack sets CreditBack field to given value.

### HasCreditBack

`func (o *Proposal) HasCreditBack() bool`

HasCreditBack returns a boolean if a field has been set.

### GetEstimatedPickup

`func (o *Proposal) GetEstimatedPickup() DeliverySlot`

GetEstimatedPickup returns the EstimatedPickup field if non-nil, zero value otherwise.

### GetEstimatedPickupOk

`func (o *Proposal) GetEstimatedPickupOk() (*DeliverySlot, bool)`

GetEstimatedPickupOk returns a tuple with the EstimatedPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPickup

`func (o *Proposal) SetEstimatedPickup(v DeliverySlot)`

SetEstimatedPickup sets EstimatedPickup field to given value.


### GetEstimatedDropoff

`func (o *Proposal) GetEstimatedDropoff() DeliverySlot`

GetEstimatedDropoff returns the EstimatedDropoff field if non-nil, zero value otherwise.

### GetEstimatedDropoffOk

`func (o *Proposal) GetEstimatedDropoffOk() (*DeliverySlot, bool)`

GetEstimatedDropoffOk returns a tuple with the EstimatedDropoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDropoff

`func (o *Proposal) SetEstimatedDropoff(v DeliverySlot)`

SetEstimatedDropoff sets EstimatedDropoff field to given value.


### GetDropoffWindowType

`func (o *Proposal) GetDropoffWindowType() DropoffWindowTypeEnum`

GetDropoffWindowType returns the DropoffWindowType field if non-nil, zero value otherwise.

### GetDropoffWindowTypeOk

`func (o *Proposal) GetDropoffWindowTypeOk() (*DropoffWindowTypeEnum, bool)`

GetDropoffWindowTypeOk returns a tuple with the DropoffWindowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindowType

`func (o *Proposal) SetDropoffWindowType(v DropoffWindowTypeEnum)`

SetDropoffWindowType sets DropoffWindowType field to given value.


### GetOrigin

`func (o *Proposal) GetOrigin() Address`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Proposal) GetOriginOk() (*Address, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Proposal) SetOrigin(v Address)`

SetOrigin sets Origin field to given value.


### GetPickupLocationId

`func (o *Proposal) GetPickupLocationId() string`

GetPickupLocationId returns the PickupLocationId field if non-nil, zero value otherwise.

### GetPickupLocationIdOk

`func (o *Proposal) GetPickupLocationIdOk() (*string, bool)`

GetPickupLocationIdOk returns a tuple with the PickupLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLocationId

`func (o *Proposal) SetPickupLocationId(v string)`

SetPickupLocationId sets PickupLocationId field to given value.

### HasPickupLocationId

`func (o *Proposal) HasPickupLocationId() bool`

HasPickupLocationId returns a boolean if a field has been set.

### GetDestination

`func (o *Proposal) GetDestination() Address`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Proposal) GetDestinationOk() (*Address, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Proposal) SetDestination(v Address)`

SetDestination sets Destination field to given value.


### GetDeliveryCompany

`func (o *Proposal) GetDeliveryCompany() string`

GetDeliveryCompany returns the DeliveryCompany field if non-nil, zero value otherwise.

### GetDeliveryCompanyOk

`func (o *Proposal) GetDeliveryCompanyOk() (*string, bool)`

GetDeliveryCompanyOk returns a tuple with the DeliveryCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCompany

`func (o *Proposal) SetDeliveryCompany(v string)`

SetDeliveryCompany sets DeliveryCompany field to given value.

### HasDeliveryCompany

`func (o *Proposal) HasDeliveryCompany() bool`

HasDeliveryCompany returns a boolean if a field has been set.

### GetLabels

`func (o *Proposal) GetLabels() []ProposalLabelEnum`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Proposal) GetLabelsOk() (*[]ProposalLabelEnum, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Proposal) SetLabels(v []ProposalLabelEnum)`

SetLabels sets Labels field to given value.


### GetPriceLevel

`func (o *Proposal) GetPriceLevel() int32`

GetPriceLevel returns the PriceLevel field if non-nil, zero value otherwise.

### GetPriceLevelOk

`func (o *Proposal) GetPriceLevelOk() (*int32, bool)`

GetPriceLevelOk returns a tuple with the PriceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLevel

`func (o *Proposal) SetPriceLevel(v int32)`

SetPriceLevel sets PriceLevel field to given value.

### HasPriceLevel

`func (o *Proposal) HasPriceLevel() bool`

HasPriceLevel returns a boolean if a field has been set.

### GetHandoffInfo

`func (o *Proposal) GetHandoffInfo() HandoffInfo`

GetHandoffInfo returns the HandoffInfo field if non-nil, zero value otherwise.

### GetHandoffInfoOk

`func (o *Proposal) GetHandoffInfoOk() (*HandoffInfo, bool)`

GetHandoffInfoOk returns a tuple with the HandoffInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffInfo

`func (o *Proposal) SetHandoffInfo(v HandoffInfo)`

SetHandoffInfo sets HandoffInfo field to given value.

### HasHandoffInfo

`func (o *Proposal) HasHandoffInfo() bool`

HasHandoffInfo returns a boolean if a field has been set.

### GetType

`func (o *Proposal) GetType() ProposalType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Proposal) GetTypeOk() (*ProposalType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Proposal) SetType(v ProposalType)`

SetType sets Type field to given value.

### HasType

`func (o *Proposal) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


