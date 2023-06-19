# WebhookPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | [**Webhook**](Webhook.md) |  | 
**Secret** | **string** | The public key from a keypair used to validate the HMAC signature for a call from this webhook. | 

## Methods

### NewWebhookPostResponse

`func NewWebhookPostResponse(webhook Webhook, secret string, ) *WebhookPostResponse`

NewWebhookPostResponse instantiates a new WebhookPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPostResponseWithDefaults

`func NewWebhookPostResponseWithDefaults() *WebhookPostResponse`

NewWebhookPostResponseWithDefaults instantiates a new WebhookPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhook

`func (o *WebhookPostResponse) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookPostResponse) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookPostResponse) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.


### GetSecret

`func (o *WebhookPostResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookPostResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookPostResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


