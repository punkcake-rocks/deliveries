# TimeslotProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | [**ProposalLabel**](ProposalLabel.md) |  | 
**Proposal** | Pointer to [**Proposal**](Proposal.md) |  | [optional] 

## Methods

### NewTimeslotProposal

`func NewTimeslotProposal(label ProposalLabel, ) *TimeslotProposal`

NewTimeslotProposal instantiates a new TimeslotProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeslotProposalWithDefaults

`func NewTimeslotProposalWithDefaults() *TimeslotProposal`

NewTimeslotProposalWithDefaults instantiates a new TimeslotProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *TimeslotProposal) GetLabel() ProposalLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TimeslotProposal) GetLabelOk() (*ProposalLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TimeslotProposal) SetLabel(v ProposalLabel)`

SetLabel sets Label field to given value.


### GetProposal

`func (o *TimeslotProposal) GetProposal() Proposal`

GetProposal returns the Proposal field if non-nil, zero value otherwise.

### GetProposalOk

`func (o *TimeslotProposal) GetProposalOk() (*Proposal, bool)`

GetProposalOk returns a tuple with the Proposal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposal

`func (o *TimeslotProposal) SetProposal(v Proposal)`

SetProposal sets Proposal field to given value.

### HasProposal

`func (o *TimeslotProposal) HasProposal() bool`

HasProposal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


