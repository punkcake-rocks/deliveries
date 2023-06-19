# ProposalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proposal** | [**Proposal**](Proposal.md) |  | 
**ValidUntil** | **time.Time** | Proposals are valid until this time in ISO8601 format | 

## Methods

### NewProposalResponse

`func NewProposalResponse(proposal Proposal, validUntil time.Time, ) *ProposalResponse`

NewProposalResponse instantiates a new ProposalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProposalResponseWithDefaults

`func NewProposalResponseWithDefaults() *ProposalResponse`

NewProposalResponseWithDefaults instantiates a new ProposalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposal

`func (o *ProposalResponse) GetProposal() Proposal`

GetProposal returns the Proposal field if non-nil, zero value otherwise.

### GetProposalOk

`func (o *ProposalResponse) GetProposalOk() (*Proposal, bool)`

GetProposalOk returns a tuple with the Proposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposal

`func (o *ProposalResponse) SetProposal(v Proposal)`

SetProposal sets Proposal field to given value.


### GetValidUntil

`func (o *ProposalResponse) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ProposalResponse) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ProposalResponse) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


