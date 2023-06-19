# OrderListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The order ID | 
**PickupWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**Carrier** | **string** | The carrier name | 
**TotalItems** | **int32** | The total number of items | 
**Value** | [**Price**](Price.md) |  | 
**Status** | [**OrderExtendedStatus**](OrderExtendedStatus.md) |  | 
**CreationTime** | **time.Time** | The order original creation time in ISO8601 format | 
**UpdatedTime** | Pointer to **time.Time** | The order last update time in ISO8601 format | [optional] 
**LabelPrintTime** | Pointer to **time.Time** | The last time a label was printed in ISO8601 format | [optional] 
**DropoffAddress** | **string** | The dropoff address | 
**DropoffWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**Customer** | [**Customer**](Customer.md) |  | 
**Notes** | Pointer to **string** | The dropoff note | [optional] 
**AggregationId** | **string** | The aggregation ID | 
**ParcelsCount** | **int32** | The number of parcels in this order | 
**ExternalOrderId** | **string** | An external order ID attached to this order | 
**StoreId** | Pointer to **string** | The ID of the e-commerce store for this order (if exists) | [optional] 
**OrderTrackingInfo** | Pointer to [**OrderTrackingInfo**](OrderTrackingInfo.md) |  | [optional] 
**HandoffType** | [**HandoffType**](HandoffType.md) |  | 
**HandoffInfo** | Pointer to [**HandoffInfo**](HandoffInfo.md) |  | [optional] 
**UsedCredits** | **bool** | Whether credits were used to pay for this order | 
**LabelRequired** | **bool** | Whether label is required | 

## Methods

### NewOrderListItem

`func NewOrderListItem(id string, pickupWindow TimeWindow, carrier string, totalItems int32, value Price, status OrderExtendedStatus, creationTime time.Time, dropoffAddress string, customer Customer, aggregationId string, parcelsCount int32, externalOrderId string, handoffType HandoffType, usedCredits bool, labelRequired bool, ) *OrderListItem`

NewOrderListItem instantiates a new OrderListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListItemWithDefaults

`func NewOrderListItemWithDefaults() *OrderListItem`

NewOrderListItemWithDefaults instantiates a new OrderListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderListItem) SetId(v string)`

SetId sets Id field to given value.


### GetPickupWindow

`func (o *OrderListItem) GetPickupWindow() TimeWindow`

GetPickupWindow returns the PickupWindow field if non-nil, zero value otherwise.

### GetPickupWindowOk

`func (o *OrderListItem) GetPickupWindowOk() (*TimeWindow, bool)`

GetPickupWindowOk returns a tuple with the PickupWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupWindow

`func (o *OrderListItem) SetPickupWindow(v TimeWindow)`

SetPickupWindow sets PickupWindow field to given value.


### GetCarrier

`func (o *OrderListItem) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *OrderListItem) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *OrderListItem) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.


### GetTotalItems

`func (o *OrderListItem) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *OrderListItem) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *OrderListItem) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetValue

`func (o *OrderListItem) GetValue() Price`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrderListItem) GetValueOk() (*Price, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrderListItem) SetValue(v Price)`

SetValue sets Value field to given value.


### GetStatus

`func (o *OrderListItem) GetStatus() OrderExtendedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListItem) GetStatusOk() (*OrderExtendedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListItem) SetStatus(v OrderExtendedStatus)`

SetStatus sets Status field to given value.


### GetCreationTime

`func (o *OrderListItem) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *OrderListItem) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *OrderListItem) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetUpdatedTime

`func (o *OrderListItem) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *OrderListItem) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *OrderListItem) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *OrderListItem) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetLabelPrintTime

`func (o *OrderListItem) GetLabelPrintTime() time.Time`

GetLabelPrintTime returns the LabelPrintTime field if non-nil, zero value otherwise.

### GetLabelPrintTimeOk

`func (o *OrderListItem) GetLabelPrintTimeOk() (*time.Time, bool)`

GetLabelPrintTimeOk returns a tuple with the LabelPrintTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPrintTime

`func (o *OrderListItem) SetLabelPrintTime(v time.Time)`

SetLabelPrintTime sets LabelPrintTime field to given value.

### HasLabelPrintTime

`func (o *OrderListItem) HasLabelPrintTime() bool`

HasLabelPrintTime returns a boolean if a field has been set.

### GetDropoffAddress

`func (o *OrderListItem) GetDropoffAddress() string`

GetDropoffAddress returns the DropoffAddress field if non-nil, zero value otherwise.

### GetDropoffAddressOk

`func (o *OrderListItem) GetDropoffAddressOk() (*string, bool)`

GetDropoffAddressOk returns a tuple with the DropoffAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffAddress

`func (o *OrderListItem) SetDropoffAddress(v string)`

SetDropoffAddress sets DropoffAddress field to given value.


### GetDropoffWindow

`func (o *OrderListItem) GetDropoffWindow() TimeWindow`

GetDropoffWindow returns the DropoffWindow field if non-nil, zero value otherwise.

### GetDropoffWindowOk

`func (o *OrderListItem) GetDropoffWindowOk() (*TimeWindow, bool)`

GetDropoffWindowOk returns a tuple with the DropoffWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindow

`func (o *OrderListItem) SetDropoffWindow(v TimeWindow)`

SetDropoffWindow sets DropoffWindow field to given value.

### HasDropoffWindow

`func (o *OrderListItem) HasDropoffWindow() bool`

HasDropoffWindow returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderListItem) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderListItem) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderListItem) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.


### GetNotes

`func (o *OrderListItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderListItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderListItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderListItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAggregationId

`func (o *OrderListItem) GetAggregationId() string`

GetAggregationId returns the AggregationId field if non-nil, zero value otherwise.

### GetAggregationIdOk

`func (o *OrderListItem) GetAggregationIdOk() (*string, bool)`

GetAggregationIdOk returns a tuple with the AggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationId

`func (o *OrderListItem) SetAggregationId(v string)`

SetAggregationId sets AggregationId field to given value.


### GetParcelsCount

`func (o *OrderListItem) GetParcelsCount() int32`

GetParcelsCount returns the ParcelsCount field if non-nil, zero value otherwise.

### GetParcelsCountOk

`func (o *OrderListItem) GetParcelsCountOk() (*int32, bool)`

GetParcelsCountOk returns a tuple with the ParcelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelsCount

`func (o *OrderListItem) SetParcelsCount(v int32)`

SetParcelsCount sets ParcelsCount field to given value.


### GetExternalOrderId

`func (o *OrderListItem) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *OrderListItem) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *OrderListItem) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.


### GetStoreId

`func (o *OrderListItem) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *OrderListItem) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *OrderListItem) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *OrderListItem) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetOrderTrackingInfo

`func (o *OrderListItem) GetOrderTrackingInfo() OrderTrackingInfo`

GetOrderTrackingInfo returns the OrderTrackingInfo field if non-nil, zero value otherwise.

### GetOrderTrackingInfoOk

`func (o *OrderListItem) GetOrderTrackingInfoOk() (*OrderTrackingInfo, bool)`

GetOrderTrackingInfoOk returns a tuple with the OrderTrackingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTrackingInfo

`func (o *OrderListItem) SetOrderTrackingInfo(v OrderTrackingInfo)`

SetOrderTrackingInfo sets OrderTrackingInfo field to given value.

### HasOrderTrackingInfo

`func (o *OrderListItem) HasOrderTrackingInfo() bool`

HasOrderTrackingInfo returns a boolean if a field has been set.

### GetHandoffType

`func (o *OrderListItem) GetHandoffType() HandoffType`

GetHandoffType returns the HandoffType field if non-nil, zero value otherwise.

### GetHandoffTypeOk

`func (o *OrderListItem) GetHandoffTypeOk() (*HandoffType, bool)`

GetHandoffTypeOk returns a tuple with the HandoffType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffType

`func (o *OrderListItem) SetHandoffType(v HandoffType)`

SetHandoffType sets HandoffType field to given value.


### GetHandoffInfo

`func (o *OrderListItem) GetHandoffInfo() HandoffInfo`

GetHandoffInfo returns the HandoffInfo field if non-nil, zero value otherwise.

### GetHandoffInfoOk

`func (o *OrderListItem) GetHandoffInfoOk() (*HandoffInfo, bool)`

GetHandoffInfoOk returns a tuple with the HandoffInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffInfo

`func (o *OrderListItem) SetHandoffInfo(v HandoffInfo)`

SetHandoffInfo sets HandoffInfo field to given value.

### HasHandoffInfo

`func (o *OrderListItem) HasHandoffInfo() bool`

HasHandoffInfo returns a boolean if a field has been set.

### GetUsedCredits

`func (o *OrderListItem) GetUsedCredits() bool`

GetUsedCredits returns the UsedCredits field if non-nil, zero value otherwise.

### GetUsedCreditsOk

`func (o *OrderListItem) GetUsedCreditsOk() (*bool, bool)`

GetUsedCreditsOk returns a tuple with the UsedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCredits

`func (o *OrderListItem) SetUsedCredits(v bool)`

SetUsedCredits sets UsedCredits field to given value.


### GetLabelRequired

`func (o *OrderListItem) GetLabelRequired() bool`

GetLabelRequired returns the LabelRequired field if non-nil, zero value otherwise.

### GetLabelRequiredOk

`func (o *OrderListItem) GetLabelRequiredOk() (*bool, bool)`

GetLabelRequiredOk returns a tuple with the LabelRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelRequired

`func (o *OrderListItem) SetLabelRequired(v bool)`

SetLabelRequired sets LabelRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


