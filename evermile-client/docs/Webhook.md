# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Evermile ID for this webhook | 
**Topic** | [**WebhookTopic**](WebhookTopic.md) |  | 
**Url** | **string** | The registered webhook URL. | 
**Status** | **string** | &lt;u&gt;The webhook status&lt;/u&gt;: &lt;br&gt; &lt;table&gt;   &lt;tr&gt;&lt;td&gt;enabled&lt;/td&gt;     &lt;td&gt;The webhook is active&lt;/td&gt;&lt;/tr&gt;   &lt;tr&gt;&lt;td&gt;disabled&lt;/td&gt;    &lt;td&gt;The webhook is disabled (due to too many delivery failures&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt;  | 
**LastSuccess** | Pointer to **time.Time** | The last time the webhook call was successful, in ISO8601 format | [optional] 
**FailuresSinceLastSuccess** | **int32** | The number of failures in calling the webhook since the last success | 
**AuthHeader** | **string** | The header to use for sending a secret string authenticating a webhook call | 
**AuthHeaderKey** | **string** | The key to set in the header | 

## Methods

### NewWebhook

`func NewWebhook(id string, topic WebhookTopic, url string, status string, failuresSinceLastSuccess int32, authHeader string, authHeaderKey string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetTopic

`func (o *Webhook) GetTopic() WebhookTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Webhook) GetTopicOk() (*WebhookTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Webhook) SetTopic(v WebhookTopic)`

SetTopic sets Topic field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *Webhook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Webhook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Webhook) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSuccess

`func (o *Webhook) GetLastSuccess() time.Time`

GetLastSuccess returns the LastSuccess field if non-nil, zero value otherwise.

### GetLastSuccessOk

`func (o *Webhook) GetLastSuccessOk() (*time.Time, bool)`

GetLastSuccessOk returns a tuple with the LastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccess

`func (o *Webhook) SetLastSuccess(v time.Time)`

SetLastSuccess sets LastSuccess field to given value.

### HasLastSuccess

`func (o *Webhook) HasLastSuccess() bool`

HasLastSuccess returns a boolean if a field has been set.

### GetFailuresSinceLastSuccess

`func (o *Webhook) GetFailuresSinceLastSuccess() int32`

GetFailuresSinceLastSuccess returns the FailuresSinceLastSuccess field if non-nil, zero value otherwise.

### GetFailuresSinceLastSuccessOk

`func (o *Webhook) GetFailuresSinceLastSuccessOk() (*int32, bool)`

GetFailuresSinceLastSuccessOk returns a tuple with the FailuresSinceLastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresSinceLastSuccess

`func (o *Webhook) SetFailuresSinceLastSuccess(v int32)`

SetFailuresSinceLastSuccess sets FailuresSinceLastSuccess field to given value.


### GetAuthHeader

`func (o *Webhook) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *Webhook) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *Webhook) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.


### GetAuthHeaderKey

`func (o *Webhook) GetAuthHeaderKey() string`

GetAuthHeaderKey returns the AuthHeaderKey field if non-nil, zero value otherwise.

### GetAuthHeaderKeyOk

`func (o *Webhook) GetAuthHeaderKeyOk() (*string, bool)`

GetAuthHeaderKeyOk returns a tuple with the AuthHeaderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaderKey

`func (o *Webhook) SetAuthHeaderKey(v string)`

SetAuthHeaderKey sets AuthHeaderKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


