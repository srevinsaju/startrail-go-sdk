# WellKnownAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **NullableString** |  | [optional] 
**ClientId** | **string** |  | 
**Device** | [**WellKnownDeviceAuth**](WellKnownDeviceAuth.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Issuer** | **string** |  | 
**JwksUrl** | **string** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWellKnownAuth

`func NewWellKnownAuth(clientId string, device WellKnownDeviceAuth, issuer string, jwksUrl string, ) *WellKnownAuth`

NewWellKnownAuth instantiates a new WellKnownAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownAuthWithDefaults

`func NewWellKnownAuthWithDefaults() *WellKnownAuth`

NewWellKnownAuthWithDefaults instantiates a new WellKnownAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *WellKnownAuth) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *WellKnownAuth) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *WellKnownAuth) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *WellKnownAuth) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *WellKnownAuth) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *WellKnownAuth) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetClientId

`func (o *WellKnownAuth) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WellKnownAuth) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WellKnownAuth) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetDevice

`func (o *WellKnownAuth) GetDevice() WellKnownDeviceAuth`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WellKnownAuth) GetDeviceOk() (*WellKnownDeviceAuth, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WellKnownAuth) SetDevice(v WellKnownDeviceAuth)`

SetDevice sets Device field to given value.


### GetEnabled

`func (o *WellKnownAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WellKnownAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WellKnownAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WellKnownAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIssuer

`func (o *WellKnownAuth) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *WellKnownAuth) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *WellKnownAuth) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksUrl

`func (o *WellKnownAuth) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *WellKnownAuth) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *WellKnownAuth) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.


### GetScopes

`func (o *WellKnownAuth) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *WellKnownAuth) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *WellKnownAuth) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *WellKnownAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *WellKnownAuth) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *WellKnownAuth) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


