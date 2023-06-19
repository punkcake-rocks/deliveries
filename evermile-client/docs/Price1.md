# Price1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int64** | The item&#39;s value in cents/pence. E.G. \&quot;500\&quot; &#x3D; 5 GBP | 
**Currency** | **string** |  | 
**Discount** | Pointer to [**Discount**](Discount.md) |  | [optional] 

## Methods

### NewPrice1

`func NewPrice1(value int64, currency string, ) *Price1`

NewPrice1 instantiates a new Price1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrice1WithDefaults

`func NewPrice1WithDefaults() *Price1`

NewPrice1WithDefaults instantiates a new Price1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Price1) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Price1) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Price1) SetValue(v int64)`

SetValue sets Value field to given value.


### GetCurrency

`func (o *Price1) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price1) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price1) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDiscount

`func (o *Price1) GetDiscount() Discount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Price1) GetDiscountOk() (*Discount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Price1) SetDiscount(v Discount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Price1) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


