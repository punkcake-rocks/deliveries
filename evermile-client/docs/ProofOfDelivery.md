# ProofOfDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the proof of delivery | 
**Urls** | Pointer to **[]string** | urls list of the specific proof type | [optional] 

## Methods

### NewProofOfDelivery

`func NewProofOfDelivery(type_ string, ) *ProofOfDelivery`

NewProofOfDelivery instantiates a new ProofOfDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProofOfDeliveryWithDefaults

`func NewProofOfDeliveryWithDefaults() *ProofOfDelivery`

NewProofOfDeliveryWithDefaults instantiates a new ProofOfDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProofOfDelivery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProofOfDelivery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProofOfDelivery) SetType(v string)`

SetType sets Type field to given value.


### GetUrls

`func (o *ProofOfDelivery) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *ProofOfDelivery) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *ProofOfDelivery) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *ProofOfDelivery) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


