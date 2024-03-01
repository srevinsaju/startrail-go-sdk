# WellKnownDeviceAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **NullableString** |  | [optional] 
**AuthorizationUrl** | **string** |  | 
**ClientId** | **string** |  | 
**DeviceCodeUrl** | **string** |  | 
**Enabled** | **bool** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 
**TokenUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWellKnownDeviceAuth

`func NewWellKnownDeviceAuth(authorizationUrl string, clientId string, deviceCodeUrl string, enabled bool, ) *WellKnownDeviceAuth`

NewWellKnownDeviceAuth instantiates a new WellKnownDeviceAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownDeviceAuthWithDefaults

`func NewWellKnownDeviceAuthWithDefaults() *WellKnownDeviceAuth`

NewWellKnownDeviceAuthWithDefaults instantiates a new WellKnownDeviceAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *WellKnownDeviceAuth) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *WellKnownDeviceAuth) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *WellKnownDeviceAuth) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *WellKnownDeviceAuth) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *WellKnownDeviceAuth) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *WellKnownDeviceAuth) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetAuthorizationUrl

`func (o *WellKnownDeviceAuth) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *WellKnownDeviceAuth) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *WellKnownDeviceAuth) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.


### GetClientId

`func (o *WellKnownDeviceAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WellKnownDeviceAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WellKnownDeviceAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetDeviceCodeUrl

`func (o *WellKnownDeviceAuth) GetDeviceCodeUrl() string`

GetDeviceCodeUrl returns the DeviceCodeUrl field if non-nil, zero value otherwise.

### GetDeviceCodeUrlOk

`func (o *WellKnownDeviceAuth) GetDeviceCodeUrlOk() (*string, bool)`

GetDeviceCodeUrlOk returns a tuple with the DeviceCodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCodeUrl

`func (o *WellKnownDeviceAuth) SetDeviceCodeUrl(v string)`

SetDeviceCodeUrl sets DeviceCodeUrl field to given value.


### GetEnabled

`func (o *WellKnownDeviceAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WellKnownDeviceAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WellKnownDeviceAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetScopes

`func (o *WellKnownDeviceAuth) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *WellKnownDeviceAuth) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *WellKnownDeviceAuth) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *WellKnownDeviceAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *WellKnownDeviceAuth) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *WellKnownDeviceAuth) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetTokenUrl

`func (o *WellKnownDeviceAuth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *WellKnownDeviceAuth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *WellKnownDeviceAuth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *WellKnownDeviceAuth) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *WellKnownDeviceAuth) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *WellKnownDeviceAuth) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


