# OrderExtendedStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderStatus** | [**OrderStatus**](OrderStatus.md) |  | 
**Reporter** | **string** | Who reported this status | 
**ReportingTime** | **time.Time** | The reporting time in ISO8601 format | 
**ManuallyUpdated** | **bool** | True if the status was manually changed by the reporter | 

## Methods

### NewOrderExtendedStatus1

`func NewOrderExtendedStatus1(orderStatus OrderStatus, reporter string, reportingTime time.Time, manuallyUpdated bool, ) *OrderExtendedStatus1`

NewOrderExtendedStatus1 instantiates a new OrderExtendedStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderExtendedStatus1WithDefaults

`func NewOrderExtendedStatus1WithDefaults() *OrderExtendedStatus1`

NewOrderExtendedStatus1WithDefaults instantiates a new OrderExtendedStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderStatus

`func (o *OrderExtendedStatus1) GetOrderStatus() OrderStatus`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderExtendedStatus1) GetOrderStatusOk() (*OrderStatus, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderExtendedStatus1) SetOrderStatus(v OrderStatus)`

SetOrderStatus sets OrderStatus field to given value.


### GetReporter

`func (o *OrderExtendedStatus1) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *OrderExtendedStatus1) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *OrderExtendedStatus1) SetReporter(v string)`

SetReporter sets Reporter field to given value.


### GetReportingTime

`func (o *OrderExtendedStatus1) GetReportingTime() time.Time`

GetReportingTime returns the ReportingTime field if non-nil, zero value otherwise.

### GetReportingTimeOk

`func (o *OrderExtendedStatus1) GetReportingTimeOk() (*time.Time, bool)`

GetReportingTimeOk returns a tuple with the ReportingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingTime

`func (o *OrderExtendedStatus1) SetReportingTime(v time.Time)`

SetReportingTime sets ReportingTime field to given value.


### GetManuallyUpdated

`func (o *OrderExtendedStatus1) GetManuallyUpdated() bool`

GetManuallyUpdated returns the ManuallyUpdated field if non-nil, zero value otherwise.

### GetManuallyUpdatedOk

`func (o *OrderExtendedStatus1) GetManuallyUpdatedOk() (*bool, bool)`

GetManuallyUpdatedOk returns a tuple with the ManuallyUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyUpdated

`func (o *OrderExtendedStatus1) SetManuallyUpdated(v bool)`

SetManuallyUpdated sets ManuallyUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


