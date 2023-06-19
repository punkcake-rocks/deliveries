# QuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateProposals** | [**[]DateProposals**](DateProposals.md) | An array (map) of proposals by date | 
**ExpressProposals** | [**[]TimeslotProposal**](TimeslotProposal.md) | An array (map) of proposals by date | 
**ValidUntil** | **time.Time** | Proposals are valid until this time in ISO8601 format | 
**CurrentPrice** | Pointer to [**Price**](Price.md) |  | [optional] 
**CancellationFee** | Pointer to [**Price**](Price.md) |  | [optional] 
**CancellationToken** | Pointer to **string** | A token that should be returned in order to confirm the cancellation of the original order. To confirm the cancellation, pass this in an X-EVERMILE-TOKEN header to createOrder. | [optional] 

## Methods

### NewQuoteResponse

`func NewQuoteResponse(dateProposals []DateProposals, expressProposals []TimeslotProposal, validUntil time.Time, ) *QuoteResponse`

NewQuoteResponse instantiates a new QuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseWithDefaults

`func NewQuoteResponseWithDefaults() *QuoteResponse`

NewQuoteResponseWithDefaults instantiates a new QuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateProposals

`func (o *QuoteResponse) GetDateProposals() []DateProposals`

GetDateProposals returns the DateProposals field if non-nil, zero value otherwise.

### GetDateProposalsOk

`func (o *QuoteResponse) GetDateProposalsOk() (*[]DateProposals, bool)`

GetDateProposalsOk returns a tuple with the DateProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProposals

`func (o *QuoteResponse) SetDateProposals(v []DateProposals)`

SetDateProposals sets DateProposals field to given value.


### GetExpressProposals

`func (o *QuoteResponse) GetExpressProposals() []TimeslotProposal`

GetExpressProposals returns the ExpressProposals field if non-nil, zero value otherwise.

### GetExpressProposalsOk

`func (o *QuoteResponse) GetExpressProposalsOk() (*[]TimeslotProposal, bool)`

GetExpressProposalsOk returns a tuple with the ExpressProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressProposals

`func (o *QuoteResponse) SetExpressProposals(v []TimeslotProposal)`

SetExpressProposals sets ExpressProposals field to given value.


### GetValidUntil

`func (o *QuoteResponse) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *QuoteResponse) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *QuoteResponse) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.


### GetCurrentPrice

`func (o *QuoteResponse) GetCurrentPrice() Price`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *QuoteResponse) GetCurrentPriceOk() (*Price, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *QuoteResponse) SetCurrentPrice(v Price)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *QuoteResponse) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetCancellationFee

`func (o *QuoteResponse) GetCancellationFee() Price`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *QuoteResponse) GetCancellationFeeOk() (*Price, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *QuoteResponse) SetCancellationFee(v Price)`

SetCancellationFee sets CancellationFee field to given value.

### HasCancellationFee

`func (o *QuoteResponse) HasCancellationFee() bool`

HasCancellationFee returns a boolean if a field has been set.

### GetCancellationToken

`func (o *QuoteResponse) GetCancellationToken() string`

GetCancellationToken returns the CancellationToken field if non-nil, zero value otherwise.

### GetCancellationTokenOk

`func (o *QuoteResponse) GetCancellationTokenOk() (*string, bool)`

GetCancellationTokenOk returns a tuple with the CancellationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationToken

`func (o *QuoteResponse) SetCancellationToken(v string)`

SetCancellationToken sets CancellationToken field to given value.

### HasCancellationToken

`func (o *QuoteResponse) HasCancellationToken() bool`

HasCancellationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


