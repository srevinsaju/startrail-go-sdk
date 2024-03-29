/*
startrail-service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package startrail

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Access type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Access{}

// Access struct for Access
type Access struct {
	Auth bool `json:"auth"`
	Endpoint string `json:"endpoint"`
	Internal bool `json:"internal"`
}

type _Access Access

// NewAccess instantiates a new Access object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccess(auth bool, endpoint string, internal bool) *Access {
	this := Access{}
	this.Auth = auth
	this.Endpoint = endpoint
	this.Internal = internal
	return &this
}

// NewAccessWithDefaults instantiates a new Access object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessWithDefaults() *Access {
	this := Access{}
	return &this
}

// GetAuth returns the Auth field value
func (o *Access) GetAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *Access) GetAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *Access) SetAuth(v bool) {
	o.Auth = v
}

// GetEndpoint returns the Endpoint field value
func (o *Access) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *Access) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *Access) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetInternal returns the Internal field value
func (o *Access) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *Access) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *Access) SetInternal(v bool) {
	o.Internal = v
}

func (o Access) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Access) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auth"] = o.Auth
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["internal"] = o.Internal
	return toSerialize, nil
}

func (o *Access) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auth",
		"endpoint",
		"internal",
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

	varAccess := _Access{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccess)

	if err != nil {
		return err
	}

	*o = Access(varAccess)

	return err
}

type NullableAccess struct {
	value *Access
	isSet bool
}

func (v NullableAccess) Get() *Access {
	return v.value
}

func (v *NullableAccess) Set(val *Access) {
	v.value = val
	v.isSet = true
}

func (v NullableAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccess(val *Access) *NullableAccess {
	return &NullableAccess{value: val, isSet: true}
}

func (v NullableAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


