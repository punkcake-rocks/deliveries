# DateProposals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date for which these proposals are proposed in format YYYY-MM-DD | 
**Proposals** | [**[]TimeslotProposal**](TimeslotProposal.md) | A list of timeslot proposals in the given date | 

## Methods

### NewDateProposals

`func NewDateProposals(date string, proposals []TimeslotProposal, ) *DateProposals`

NewDateProposals instantiates a new DateProposals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateProposalsWithDefaults

`func NewDateProposalsWithDefaults() *DateProposals`

NewDateProposalsWithDefaults instantiates a new DateProposals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DateProposals) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DateProposals) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DateProposals) SetDate(v string)`

SetDate sets Date field to given value.


### GetProposals

`func (o *DateProposals) GetProposals() []TimeslotProposal`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *DateProposals) GetProposalsOk() (*[]TimeslotProposal, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *DateProposals) SetProposals(v []TimeslotProposal)`

SetProposals sets Proposals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


