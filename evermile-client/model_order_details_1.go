/*
Evermile Commercial Quotes and Booking API

This is Evermile's commercial API for handling delivery quotes and orders

API version: 1.0
Contact: support@evermile.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EvermileClient

import (
	"encoding/json"
	"time"
)

// checks if the OrderDetails1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetails1{}

// OrderDetails1 The details of a booked order
type OrderDetails1 struct {
	// The ID of the order
	Id string `json:"id"`
	// An external order ID attached to this order
	ExternalOrderId string `json:"externalOrderId"`
	// The ID of the e-commerce store for this order (if exists)
	StoreId *string `json:"storeId,omitempty"`
	// The name of the store for this order (if exists)
	StoreName *string `json:"storeName,omitempty"`
	Status OrderExtendedStatus1 `json:"status"`
	CommunicatedPickup DeliverySlot `json:"communicatedPickup"`
	CommunicatedDropoff DeliverySlot `json:"communicatedDropoff"`
	// The actual pickup time in ISO8601 format (for packages picked up)
	ActualPickup *time.Time `json:"actualPickup,omitempty"`
	// The actual dropoff time in ISO8601 format (for packages dropped off)
	ActualDropoff *time.Time `json:"actualDropoff,omitempty"`
	// The expected pickup time (if known)
	PickupEta *time.Time `json:"pickupEta,omitempty"`
	// The expected dropoff time (if known)
	DropoffEta *time.Time `json:"dropoffEta,omitempty"`
	CourierLocation *GeoLocation `json:"courierLocation,omitempty"`
	Origin Address2 `json:"origin"`
	Destination Address2 `json:"destination"`
	Customer Customer `json:"customer"`
	// The company that will perform the delivery (if known)
	DeliveryCompany *string `json:"deliveryCompany,omitempty"`
	// The list of parcels in this order
	Parcels []Parcel1 `json:"parcels"`
	Value Price1 `json:"value"`
	// Additional handling instructions for the packages
	Instructions *string `json:"instructions,omitempty"`
	// Handling instructions for a parcel
	Handling []HandlingInstruction `json:"handling,omitempty"`
	// Proof of delivery requirements for a parcel
	ProofOfDeliveryRequirement []ProofOfDeliveryRequirement `json:"proofOfDeliveryRequirement,omitempty"`
	// The carrier tracking URL for this order (if exists)
	TrackingUrl *string `json:"trackingUrl,omitempty"`
	// The Evermile customer tracking URL for this order (if exists)
	CustomerTrackingUrl *string `json:"customerTrackingUrl,omitempty"`
	OrderTrackingInfo *OrderTrackingInfo1 `json:"orderTrackingInfo,omitempty"`
	HandoffType HandoffType `json:"handoffType"`
	HandoffInfo *HandoffInfo1 `json:"handoffInfo,omitempty"`
	// Whether credits were used to pay for this order
	UsedCredits bool `json:"usedCredits"`
	// Proof of delivery for an order
	ProofOfDelivery []ProofOfDelivery `json:"proofOfDelivery,omitempty"`
	// Courier notes from each delivery in this order
	Notes []string `json:"notes,omitempty"`
	// The name of the merchant for this order
	MerchantName string `json:"merchantName"`
	// Whether label is required
	LabelRequired bool `json:"labelRequired"`
}

// NewOrderDetails1 instantiates a new OrderDetails1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetails1(id string, externalOrderId string, status OrderExtendedStatus1, communicatedPickup DeliverySlot, communicatedDropoff DeliverySlot, origin Address2, destination Address2, customer Customer, parcels []Parcel1, value Price1, handoffType HandoffType, usedCredits bool, merchantName string, labelRequired bool) *OrderDetails1 {
	this := OrderDetails1{}
	this.Id = id
	this.ExternalOrderId = externalOrderId
	this.Status = status
	this.CommunicatedPickup = communicatedPickup
	this.CommunicatedDropoff = communicatedDropoff
	this.Origin = origin
	this.Destination = destination
	this.Customer = customer
	this.Parcels = parcels
	this.Value = value
	this.HandoffType = handoffType
	this.UsedCredits = usedCredits
	this.MerchantName = merchantName
	this.LabelRequired = labelRequired
	return &this
}

// NewOrderDetails1WithDefaults instantiates a new OrderDetails1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetails1WithDefaults() *OrderDetails1 {
	this := OrderDetails1{}
	return &this
}

// GetId returns the Id field value
func (o *OrderDetails1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderDetails1) SetId(v string) {
	o.Id = v
}

// GetExternalOrderId returns the ExternalOrderId field value
func (o *OrderDetails1) GetExternalOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalOrderId
}

// GetExternalOrderIdOk returns a tuple with the ExternalOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetExternalOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalOrderId, true
}

// SetExternalOrderId sets field value
func (o *OrderDetails1) SetExternalOrderId(v string) {
	o.ExternalOrderId = v
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *OrderDetails1) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *OrderDetails1) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *OrderDetails1) SetStoreId(v string) {
	o.StoreId = &v
}

// GetStoreName returns the StoreName field value if set, zero value otherwise.
func (o *OrderDetails1) GetStoreName() string {
	if o == nil || IsNil(o.StoreName) {
		var ret string
		return ret
	}
	return *o.StoreName
}

// GetStoreNameOk returns a tuple with the StoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.StoreName) {
		return nil, false
	}
	return o.StoreName, true
}

// HasStoreName returns a boolean if a field has been set.
func (o *OrderDetails1) HasStoreName() bool {
	if o != nil && !IsNil(o.StoreName) {
		return true
	}

	return false
}

// SetStoreName gets a reference to the given string and assigns it to the StoreName field.
func (o *OrderDetails1) SetStoreName(v string) {
	o.StoreName = &v
}

// GetStatus returns the Status field value
func (o *OrderDetails1) GetStatus() OrderExtendedStatus1 {
	if o == nil {
		var ret OrderExtendedStatus1
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetStatusOk() (*OrderExtendedStatus1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderDetails1) SetStatus(v OrderExtendedStatus1) {
	o.Status = v
}

// GetCommunicatedPickup returns the CommunicatedPickup field value
func (o *OrderDetails1) GetCommunicatedPickup() DeliverySlot {
	if o == nil {
		var ret DeliverySlot
		return ret
	}

	return o.CommunicatedPickup
}

// GetCommunicatedPickupOk returns a tuple with the CommunicatedPickup field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetCommunicatedPickupOk() (*DeliverySlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicatedPickup, true
}

// SetCommunicatedPickup sets field value
func (o *OrderDetails1) SetCommunicatedPickup(v DeliverySlot) {
	o.CommunicatedPickup = v
}

// GetCommunicatedDropoff returns the CommunicatedDropoff field value
func (o *OrderDetails1) GetCommunicatedDropoff() DeliverySlot {
	if o == nil {
		var ret DeliverySlot
		return ret
	}

	return o.CommunicatedDropoff
}

// GetCommunicatedDropoffOk returns a tuple with the CommunicatedDropoff field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetCommunicatedDropoffOk() (*DeliverySlot, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicatedDropoff, true
}

// SetCommunicatedDropoff sets field value
func (o *OrderDetails1) SetCommunicatedDropoff(v DeliverySlot) {
	o.CommunicatedDropoff = v
}

// GetActualPickup returns the ActualPickup field value if set, zero value otherwise.
func (o *OrderDetails1) GetActualPickup() time.Time {
	if o == nil || IsNil(o.ActualPickup) {
		var ret time.Time
		return ret
	}
	return *o.ActualPickup
}

// GetActualPickupOk returns a tuple with the ActualPickup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetActualPickupOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActualPickup) {
		return nil, false
	}
	return o.ActualPickup, true
}

// HasActualPickup returns a boolean if a field has been set.
func (o *OrderDetails1) HasActualPickup() bool {
	if o != nil && !IsNil(o.ActualPickup) {
		return true
	}

	return false
}

// SetActualPickup gets a reference to the given time.Time and assigns it to the ActualPickup field.
func (o *OrderDetails1) SetActualPickup(v time.Time) {
	o.ActualPickup = &v
}

// GetActualDropoff returns the ActualDropoff field value if set, zero value otherwise.
func (o *OrderDetails1) GetActualDropoff() time.Time {
	if o == nil || IsNil(o.ActualDropoff) {
		var ret time.Time
		return ret
	}
	return *o.ActualDropoff
}

// GetActualDropoffOk returns a tuple with the ActualDropoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetActualDropoffOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActualDropoff) {
		return nil, false
	}
	return o.ActualDropoff, true
}

// HasActualDropoff returns a boolean if a field has been set.
func (o *OrderDetails1) HasActualDropoff() bool {
	if o != nil && !IsNil(o.ActualDropoff) {
		return true
	}

	return false
}

// SetActualDropoff gets a reference to the given time.Time and assigns it to the ActualDropoff field.
func (o *OrderDetails1) SetActualDropoff(v time.Time) {
	o.ActualDropoff = &v
}

// GetPickupEta returns the PickupEta field value if set, zero value otherwise.
func (o *OrderDetails1) GetPickupEta() time.Time {
	if o == nil || IsNil(o.PickupEta) {
		var ret time.Time
		return ret
	}
	return *o.PickupEta
}

// GetPickupEtaOk returns a tuple with the PickupEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetPickupEtaOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PickupEta) {
		return nil, false
	}
	return o.PickupEta, true
}

// HasPickupEta returns a boolean if a field has been set.
func (o *OrderDetails1) HasPickupEta() bool {
	if o != nil && !IsNil(o.PickupEta) {
		return true
	}

	return false
}

// SetPickupEta gets a reference to the given time.Time and assigns it to the PickupEta field.
func (o *OrderDetails1) SetPickupEta(v time.Time) {
	o.PickupEta = &v
}

// GetDropoffEta returns the DropoffEta field value if set, zero value otherwise.
func (o *OrderDetails1) GetDropoffEta() time.Time {
	if o == nil || IsNil(o.DropoffEta) {
		var ret time.Time
		return ret
	}
	return *o.DropoffEta
}

// GetDropoffEtaOk returns a tuple with the DropoffEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetDropoffEtaOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DropoffEta) {
		return nil, false
	}
	return o.DropoffEta, true
}

// HasDropoffEta returns a boolean if a field has been set.
func (o *OrderDetails1) HasDropoffEta() bool {
	if o != nil && !IsNil(o.DropoffEta) {
		return true
	}

	return false
}

// SetDropoffEta gets a reference to the given time.Time and assigns it to the DropoffEta field.
func (o *OrderDetails1) SetDropoffEta(v time.Time) {
	o.DropoffEta = &v
}

// GetCourierLocation returns the CourierLocation field value if set, zero value otherwise.
func (o *OrderDetails1) GetCourierLocation() GeoLocation {
	if o == nil || IsNil(o.CourierLocation) {
		var ret GeoLocation
		return ret
	}
	return *o.CourierLocation
}

// GetCourierLocationOk returns a tuple with the CourierLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetCourierLocationOk() (*GeoLocation, bool) {
	if o == nil || IsNil(o.CourierLocation) {
		return nil, false
	}
	return o.CourierLocation, true
}

// HasCourierLocation returns a boolean if a field has been set.
func (o *OrderDetails1) HasCourierLocation() bool {
	if o != nil && !IsNil(o.CourierLocation) {
		return true
	}

	return false
}

// SetCourierLocation gets a reference to the given GeoLocation and assigns it to the CourierLocation field.
func (o *OrderDetails1) SetCourierLocation(v GeoLocation) {
	o.CourierLocation = &v
}

// GetOrigin returns the Origin field value
func (o *OrderDetails1) GetOrigin() Address2 {
	if o == nil {
		var ret Address2
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetOriginOk() (*Address2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *OrderDetails1) SetOrigin(v Address2) {
	o.Origin = v
}

// GetDestination returns the Destination field value
func (o *OrderDetails1) GetDestination() Address2 {
	if o == nil {
		var ret Address2
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetDestinationOk() (*Address2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *OrderDetails1) SetDestination(v Address2) {
	o.Destination = v
}

// GetCustomer returns the Customer field value
func (o *OrderDetails1) GetCustomer() Customer {
	if o == nil {
		var ret Customer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetCustomerOk() (*Customer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *OrderDetails1) SetCustomer(v Customer) {
	o.Customer = v
}

// GetDeliveryCompany returns the DeliveryCompany field value if set, zero value otherwise.
func (o *OrderDetails1) GetDeliveryCompany() string {
	if o == nil || IsNil(o.DeliveryCompany) {
		var ret string
		return ret
	}
	return *o.DeliveryCompany
}

// GetDeliveryCompanyOk returns a tuple with the DeliveryCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetDeliveryCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryCompany) {
		return nil, false
	}
	return o.DeliveryCompany, true
}

// HasDeliveryCompany returns a boolean if a field has been set.
func (o *OrderDetails1) HasDeliveryCompany() bool {
	if o != nil && !IsNil(o.DeliveryCompany) {
		return true
	}

	return false
}

// SetDeliveryCompany gets a reference to the given string and assigns it to the DeliveryCompany field.
func (o *OrderDetails1) SetDeliveryCompany(v string) {
	o.DeliveryCompany = &v
}

// GetParcels returns the Parcels field value
func (o *OrderDetails1) GetParcels() []Parcel1 {
	if o == nil {
		var ret []Parcel1
		return ret
	}

	return o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetParcelsOk() ([]Parcel1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parcels, true
}

// SetParcels sets field value
func (o *OrderDetails1) SetParcels(v []Parcel1) {
	o.Parcels = v
}

// GetValue returns the Value field value
func (o *OrderDetails1) GetValue() Price1 {
	if o == nil {
		var ret Price1
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetValueOk() (*Price1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OrderDetails1) SetValue(v Price1) {
	o.Value = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *OrderDetails1) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *OrderDetails1) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *OrderDetails1) SetInstructions(v string) {
	o.Instructions = &v
}

// GetHandling returns the Handling field value if set, zero value otherwise.
func (o *OrderDetails1) GetHandling() []HandlingInstruction {
	if o == nil || IsNil(o.Handling) {
		var ret []HandlingInstruction
		return ret
	}
	return o.Handling
}

// GetHandlingOk returns a tuple with the Handling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetHandlingOk() ([]HandlingInstruction, bool) {
	if o == nil || IsNil(o.Handling) {
		return nil, false
	}
	return o.Handling, true
}

// HasHandling returns a boolean if a field has been set.
func (o *OrderDetails1) HasHandling() bool {
	if o != nil && !IsNil(o.Handling) {
		return true
	}

	return false
}

// SetHandling gets a reference to the given []HandlingInstruction and assigns it to the Handling field.
func (o *OrderDetails1) SetHandling(v []HandlingInstruction) {
	o.Handling = v
}

// GetProofOfDeliveryRequirement returns the ProofOfDeliveryRequirement field value if set, zero value otherwise.
func (o *OrderDetails1) GetProofOfDeliveryRequirement() []ProofOfDeliveryRequirement {
	if o == nil || IsNil(o.ProofOfDeliveryRequirement) {
		var ret []ProofOfDeliveryRequirement
		return ret
	}
	return o.ProofOfDeliveryRequirement
}

// GetProofOfDeliveryRequirementOk returns a tuple with the ProofOfDeliveryRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetProofOfDeliveryRequirementOk() ([]ProofOfDeliveryRequirement, bool) {
	if o == nil || IsNil(o.ProofOfDeliveryRequirement) {
		return nil, false
	}
	return o.ProofOfDeliveryRequirement, true
}

// HasProofOfDeliveryRequirement returns a boolean if a field has been set.
func (o *OrderDetails1) HasProofOfDeliveryRequirement() bool {
	if o != nil && !IsNil(o.ProofOfDeliveryRequirement) {
		return true
	}

	return false
}

// SetProofOfDeliveryRequirement gets a reference to the given []ProofOfDeliveryRequirement and assigns it to the ProofOfDeliveryRequirement field.
func (o *OrderDetails1) SetProofOfDeliveryRequirement(v []ProofOfDeliveryRequirement) {
	o.ProofOfDeliveryRequirement = v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *OrderDetails1) GetTrackingUrl() string {
	if o == nil || IsNil(o.TrackingUrl) {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingUrl) {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *OrderDetails1) HasTrackingUrl() bool {
	if o != nil && !IsNil(o.TrackingUrl) {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *OrderDetails1) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

// GetCustomerTrackingUrl returns the CustomerTrackingUrl field value if set, zero value otherwise.
func (o *OrderDetails1) GetCustomerTrackingUrl() string {
	if o == nil || IsNil(o.CustomerTrackingUrl) {
		var ret string
		return ret
	}
	return *o.CustomerTrackingUrl
}

// GetCustomerTrackingUrlOk returns a tuple with the CustomerTrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetCustomerTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerTrackingUrl) {
		return nil, false
	}
	return o.CustomerTrackingUrl, true
}

// HasCustomerTrackingUrl returns a boolean if a field has been set.
func (o *OrderDetails1) HasCustomerTrackingUrl() bool {
	if o != nil && !IsNil(o.CustomerTrackingUrl) {
		return true
	}

	return false
}

// SetCustomerTrackingUrl gets a reference to the given string and assigns it to the CustomerTrackingUrl field.
func (o *OrderDetails1) SetCustomerTrackingUrl(v string) {
	o.CustomerTrackingUrl = &v
}

// GetOrderTrackingInfo returns the OrderTrackingInfo field value if set, zero value otherwise.
func (o *OrderDetails1) GetOrderTrackingInfo() OrderTrackingInfo1 {
	if o == nil || IsNil(o.OrderTrackingInfo) {
		var ret OrderTrackingInfo1
		return ret
	}
	return *o.OrderTrackingInfo
}

// GetOrderTrackingInfoOk returns a tuple with the OrderTrackingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetOrderTrackingInfoOk() (*OrderTrackingInfo1, bool) {
	if o == nil || IsNil(o.OrderTrackingInfo) {
		return nil, false
	}
	return o.OrderTrackingInfo, true
}

// HasOrderTrackingInfo returns a boolean if a field has been set.
func (o *OrderDetails1) HasOrderTrackingInfo() bool {
	if o != nil && !IsNil(o.OrderTrackingInfo) {
		return true
	}

	return false
}

// SetOrderTrackingInfo gets a reference to the given OrderTrackingInfo1 and assigns it to the OrderTrackingInfo field.
func (o *OrderDetails1) SetOrderTrackingInfo(v OrderTrackingInfo1) {
	o.OrderTrackingInfo = &v
}

// GetHandoffType returns the HandoffType field value
func (o *OrderDetails1) GetHandoffType() HandoffType {
	if o == nil {
		var ret HandoffType
		return ret
	}

	return o.HandoffType
}

// GetHandoffTypeOk returns a tuple with the HandoffType field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetHandoffTypeOk() (*HandoffType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandoffType, true
}

// SetHandoffType sets field value
func (o *OrderDetails1) SetHandoffType(v HandoffType) {
	o.HandoffType = v
}

// GetHandoffInfo returns the HandoffInfo field value if set, zero value otherwise.
func (o *OrderDetails1) GetHandoffInfo() HandoffInfo1 {
	if o == nil || IsNil(o.HandoffInfo) {
		var ret HandoffInfo1
		return ret
	}
	return *o.HandoffInfo
}

// GetHandoffInfoOk returns a tuple with the HandoffInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetHandoffInfoOk() (*HandoffInfo1, bool) {
	if o == nil || IsNil(o.HandoffInfo) {
		return nil, false
	}
	return o.HandoffInfo, true
}

// HasHandoffInfo returns a boolean if a field has been set.
func (o *OrderDetails1) HasHandoffInfo() bool {
	if o != nil && !IsNil(o.HandoffInfo) {
		return true
	}

	return false
}

// SetHandoffInfo gets a reference to the given HandoffInfo1 and assigns it to the HandoffInfo field.
func (o *OrderDetails1) SetHandoffInfo(v HandoffInfo1) {
	o.HandoffInfo = &v
}

// GetUsedCredits returns the UsedCredits field value
func (o *OrderDetails1) GetUsedCredits() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsedCredits
}

// GetUsedCreditsOk returns a tuple with the UsedCredits field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetUsedCreditsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedCredits, true
}

// SetUsedCredits sets field value
func (o *OrderDetails1) SetUsedCredits(v bool) {
	o.UsedCredits = v
}

// GetProofOfDelivery returns the ProofOfDelivery field value if set, zero value otherwise.
func (o *OrderDetails1) GetProofOfDelivery() []ProofOfDelivery {
	if o == nil || IsNil(o.ProofOfDelivery) {
		var ret []ProofOfDelivery
		return ret
	}
	return o.ProofOfDelivery
}

// GetProofOfDeliveryOk returns a tuple with the ProofOfDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetProofOfDeliveryOk() ([]ProofOfDelivery, bool) {
	if o == nil || IsNil(o.ProofOfDelivery) {
		return nil, false
	}
	return o.ProofOfDelivery, true
}

// HasProofOfDelivery returns a boolean if a field has been set.
func (o *OrderDetails1) HasProofOfDelivery() bool {
	if o != nil && !IsNil(o.ProofOfDelivery) {
		return true
	}

	return false
}

// SetProofOfDelivery gets a reference to the given []ProofOfDelivery and assigns it to the ProofOfDelivery field.
func (o *OrderDetails1) SetProofOfDelivery(v []ProofOfDelivery) {
	o.ProofOfDelivery = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderDetails1) GetNotes() []string {
	if o == nil || IsNil(o.Notes) {
		var ret []string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderDetails1) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []string and assigns it to the Notes field.
func (o *OrderDetails1) SetNotes(v []string) {
	o.Notes = v
}

// GetMerchantName returns the MerchantName field value
func (o *OrderDetails1) GetMerchantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetMerchantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantName, true
}

// SetMerchantName sets field value
func (o *OrderDetails1) SetMerchantName(v string) {
	o.MerchantName = v
}

// GetLabelRequired returns the LabelRequired field value
func (o *OrderDetails1) GetLabelRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LabelRequired
}

// GetLabelRequiredOk returns a tuple with the LabelRequired field value
// and a boolean to check if the value has been set.
func (o *OrderDetails1) GetLabelRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelRequired, true
}

// SetLabelRequired sets field value
func (o *OrderDetails1) SetLabelRequired(v bool) {
	o.LabelRequired = v
}

func (o OrderDetails1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetails1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["externalOrderId"] = o.ExternalOrderId
	if !IsNil(o.StoreId) {
		toSerialize["storeId"] = o.StoreId
	}
	if !IsNil(o.StoreName) {
		toSerialize["storeName"] = o.StoreName
	}
	toSerialize["status"] = o.Status
	toSerialize["communicatedPickup"] = o.CommunicatedPickup
	toSerialize["communicatedDropoff"] = o.CommunicatedDropoff
	if !IsNil(o.ActualPickup) {
		toSerialize["actualPickup"] = o.ActualPickup
	}
	if !IsNil(o.ActualDropoff) {
		toSerialize["actualDropoff"] = o.ActualDropoff
	}
	if !IsNil(o.PickupEta) {
		toSerialize["pickupEta"] = o.PickupEta
	}
	if !IsNil(o.DropoffEta) {
		toSerialize["dropoffEta"] = o.DropoffEta
	}
	if !IsNil(o.CourierLocation) {
		toSerialize["courierLocation"] = o.CourierLocation
	}
	toSerialize["origin"] = o.Origin
	toSerialize["destination"] = o.Destination
	toSerialize["customer"] = o.Customer
	if !IsNil(o.DeliveryCompany) {
		toSerialize["deliveryCompany"] = o.DeliveryCompany
	}
	toSerialize["parcels"] = o.Parcels
	toSerialize["value"] = o.Value
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.Handling) {
		toSerialize["handling"] = o.Handling
	}
	if !IsNil(o.ProofOfDeliveryRequirement) {
		toSerialize["proofOfDeliveryRequirement"] = o.ProofOfDeliveryRequirement
	}
	if !IsNil(o.TrackingUrl) {
		toSerialize["trackingUrl"] = o.TrackingUrl
	}
	if !IsNil(o.CustomerTrackingUrl) {
		toSerialize["customerTrackingUrl"] = o.CustomerTrackingUrl
	}
	if !IsNil(o.OrderTrackingInfo) {
		toSerialize["orderTrackingInfo"] = o.OrderTrackingInfo
	}
	toSerialize["handoffType"] = o.HandoffType
	if !IsNil(o.HandoffInfo) {
		toSerialize["handoffInfo"] = o.HandoffInfo
	}
	toSerialize["usedCredits"] = o.UsedCredits
	if !IsNil(o.ProofOfDelivery) {
		toSerialize["proofOfDelivery"] = o.ProofOfDelivery
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["merchantName"] = o.MerchantName
	toSerialize["labelRequired"] = o.LabelRequired
	return toSerialize, nil
}

type NullableOrderDetails1 struct {
	value *OrderDetails1
	isSet bool
}

func (v NullableOrderDetails1) Get() *OrderDetails1 {
	return v.value
}

func (v *NullableOrderDetails1) Set(val *OrderDetails1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetails1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetails1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetails1(val *OrderDetails1) *NullableOrderDetails1 {
	return &NullableOrderDetails1{value: val, isSet: true}
}

func (v NullableOrderDetails1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetails1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


