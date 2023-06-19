# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListPriceCents** | **int64** | The item&#39;s list price (before discount) in cents/pence. E.G. \&quot;500\&quot; &#x3D; 5 GBP | 
**DiscountPercent** | **int64** | The discount percentage | 
**Reason** | **string** | The discount reason | 

## Methods

### NewDiscount

`func NewDiscount(listPriceCents int64, discountPercent int64, reason string, ) *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListPriceCents

`func (o *Discount) GetListPriceCents() int64`

GetListPriceCents returns the ListPriceCents field if non-nil, zero value otherwise.

### GetListPriceCentsOk

`func (o *Discount) GetListPriceCentsOk() (*int64, bool)`

GetListPriceCentsOk returns a tuple with the ListPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListPriceCents

`func (o *Discount) SetListPriceCents(v int64)`

SetListPriceCents sets ListPriceCents field to given value.


### GetDiscountPercent

`func (o *Discount) GetDiscountPercent() int64`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *Discount) GetDiscountPercentOk() (*int64, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *Discount) SetDiscountPercent(v int64)`

SetDiscountPercent sets DiscountPercent field to given value.


### GetReason

`func (o *Discount) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Discount) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Discount) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


