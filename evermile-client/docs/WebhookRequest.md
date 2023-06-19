# WebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to [**WebhookTopic**](WebhookTopic.md) |  | [optional] 
**Url** | **string** | The URL to send the webhook to. Must be an **https** endpoint | 
**AuthHeader** | **string** | A header to use for sending a secret string authenticating a webhook call | 
**AuthHeaderKey** | **string** | The key to set in the header | 

## Methods

### NewWebhookRequest

`func NewWebhookRequest(url string, authHeader string, authHeaderKey string, ) *WebhookRequest`

NewWebhookRequest instantiates a new WebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestWithDefaults

`func NewWebhookRequestWithDefaults() *WebhookRequest`

NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *WebhookRequest) GetTopic() WebhookTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookRequest) GetTopicOk() (*WebhookTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookRequest) SetTopic(v WebhookTopic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthHeader

`func (o *WebhookRequest) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *WebhookRequest) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *WebhookRequest) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.


### GetAuthHeaderKey

`func (o *WebhookRequest) GetAuthHeaderKey() string`

GetAuthHeaderKey returns the AuthHeaderKey field if non-nil, zero value otherwise.

### GetAuthHeaderKeyOk

`func (o *WebhookRequest) GetAuthHeaderKeyOk() (*string, bool)`

GetAuthHeaderKeyOk returns a tuple with the AuthHeaderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderKey

`func (o *WebhookRequest) SetAuthHeaderKey(v string)`

SetAuthHeaderKey sets AuthHeaderKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


