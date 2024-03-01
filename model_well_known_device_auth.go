/*
startrail-service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WellKnownDeviceAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownDeviceAuth{}

// WellKnownDeviceAuth struct for WellKnownDeviceAuth
type WellKnownDeviceAuth struct {
	Audience NullableString `json:"audience,omitempty"`
	AuthorizationUrl string `json:"authorization_url"`
	ClientId string `json:"client_id"`
	DeviceCodeUrl string `json:"device_code_url"`
	Enabled bool `json:"enabled"`
	Scopes []string `json:"scopes,omitempty"`
	TokenUrl NullableString `json:"token_url,omitempty"`
}

type _WellKnownDeviceAuth WellKnownDeviceAuth

// NewWellKnownDeviceAuth instantiates a new WellKnownDeviceAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownDeviceAuth(authorizationUrl string, clientId string, deviceCodeUrl string, enabled bool) *WellKnownDeviceAuth {
	this := WellKnownDeviceAuth{}
	this.AuthorizationUrl = authorizationUrl
	this.ClientId = clientId
	this.DeviceCodeUrl = deviceCodeUrl
	this.Enabled = enabled
	return &this
}

// NewWellKnownDeviceAuthWithDefaults instantiates a new WellKnownDeviceAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownDeviceAuthWithDefaults() *WellKnownDeviceAuth {
	this := WellKnownDeviceAuth{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WellKnownDeviceAuth) GetAudience() string {
	if o == nil || IsNil(o.Audience.Get()) {
		var ret string
		return ret
	}
	return *o.Audience.Get()
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WellKnownDeviceAuth) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Audience.Get(), o.Audience.IsSet()
}

// HasAudience returns a boolean if a field has been set.
func (o *WellKnownDeviceAuth) HasAudience() bool {
	if o != nil && o.Audience.IsSet() {
		return true
	}

	return false
}

// SetAudience gets a reference to the given NullableString and assigns it to the Audience field.
func (o *WellKnownDeviceAuth) SetAudience(v string) {
	o.Audience.Set(&v)
}
// SetAudienceNil sets the value for Audience to be an explicit nil
func (o *WellKnownDeviceAuth) SetAudienceNil() {
	o.Audience.Set(nil)
}

// UnsetAudience ensures that no value is present for Audience, not even an explicit nil
func (o *WellKnownDeviceAuth) UnsetAudience() {
	o.Audience.Unset()
}

// GetAuthorizationUrl returns the AuthorizationUrl field value
func (o *WellKnownDeviceAuth) GetAuthorizationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationUrl
}

// GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field value
// and a boolean to check if the value has been set.
func (o *WellKnownDeviceAuth) GetAuthorizationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationUrl, true
}

// SetAuthorizationUrl sets field value
func (o *WellKnownDeviceAuth) SetAuthorizationUrl(v string) {
	o.AuthorizationUrl = v
}

// GetClientId returns the ClientId field value
func (o *WellKnownDeviceAuth) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *WellKnownDeviceAuth) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *WellKnownDeviceAuth) SetClientId(v string) {
	o.ClientId = v
}

// GetDeviceCodeUrl returns the DeviceCodeUrl field value
func (o *WellKnownDeviceAuth) GetDeviceCodeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceCodeUrl
}

// GetDeviceCodeUrlOk returns a tuple with the DeviceCodeUrl field value
// and a boolean to check if the value has been set.
func (o *WellKnownDeviceAuth) GetDeviceCodeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceCodeUrl, true
}

// SetDeviceCodeUrl sets field value
func (o *WellKnownDeviceAuth) SetDeviceCodeUrl(v string) {
	o.DeviceCodeUrl = v
}

// GetEnabled returns the Enabled field value
func (o *WellKnownDeviceAuth) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WellKnownDeviceAuth) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WellKnownDeviceAuth) SetEnabled(v bool) {
	o.Enabled = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WellKnownDeviceAuth) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WellKnownDeviceAuth) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *WellKnownDeviceAuth) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *WellKnownDeviceAuth) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WellKnownDeviceAuth) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TokenUrl.Get()
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WellKnownDeviceAuth) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl.Get(), o.TokenUrl.IsSet()
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *WellKnownDeviceAuth) HasTokenUrl() bool {
	if o != nil && o.TokenUrl.IsSet() {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given NullableString and assigns it to the TokenUrl field.
func (o *WellKnownDeviceAuth) SetTokenUrl(v string) {
	o.TokenUrl.Set(&v)
}
// SetTokenUrlNil sets the value for TokenUrl to be an explicit nil
func (o *WellKnownDeviceAuth) SetTokenUrlNil() {
	o.TokenUrl.Set(nil)
}

// UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
func (o *WellKnownDeviceAuth) UnsetTokenUrl() {
	o.TokenUrl.Unset()
}

func (o WellKnownDeviceAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownDeviceAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience.IsSet() {
		toSerialize["audience"] = o.Audience.Get()
	}
	toSerialize["authorization_url"] = o.AuthorizationUrl
	toSerialize["client_id"] = o.ClientId
	toSerialize["device_code_url"] = o.DeviceCodeUrl
	toSerialize["enabled"] = o.Enabled
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.TokenUrl.IsSet() {
		toSerialize["token_url"] = o.TokenUrl.Get()
	}
	return toSerialize, nil
}

func (o *WellKnownDeviceAuth) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorization_url",
		"client_id",
		"device_code_url",
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWellKnownDeviceAuth := _WellKnownDeviceAuth{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWellKnownDeviceAuth)

	if err != nil {
		return err
	}

	*o = WellKnownDeviceAuth(varWellKnownDeviceAuth)

	return err
}

type NullableWellKnownDeviceAuth struct {
	value *WellKnownDeviceAuth
	isSet bool
}

func (v NullableWellKnownDeviceAuth) Get() *WellKnownDeviceAuth {
	return v.value
}

func (v *NullableWellKnownDeviceAuth) Set(val *WellKnownDeviceAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownDeviceAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownDeviceAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownDeviceAuth(val *WellKnownDeviceAuth) *NullableWellKnownDeviceAuth {
	return &NullableWellKnownDeviceAuth{value: val, isSet: true}
}

func (v NullableWellKnownDeviceAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownDeviceAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


