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

// checks if the ServiceListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceListResponse{}

// ServiceListResponse struct for ServiceListResponse
type ServiceListResponse struct {
	Diagnostics []Diagnostic `json:"diagnostics"`
	Response []Service `json:"response,omitempty"`
	Success bool `json:"success"`
}

type _ServiceListResponse ServiceListResponse

// NewServiceListResponse instantiates a new ServiceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceListResponse(diagnostics []Diagnostic, success bool) *ServiceListResponse {
	this := ServiceListResponse{}
	this.Diagnostics = diagnostics
	this.Success = success
	return &this
}

// NewServiceListResponseWithDefaults instantiates a new ServiceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceListResponseWithDefaults() *ServiceListResponse {
	this := ServiceListResponse{}
	return &this
}

// GetDiagnostics returns the Diagnostics field value
func (o *ServiceListResponse) GetDiagnostics() []Diagnostic {
	if o == nil {
		var ret []Diagnostic
		return ret
	}

	return o.Diagnostics
}

// GetDiagnosticsOk returns a tuple with the Diagnostics field value
// and a boolean to check if the value has been set.
func (o *ServiceListResponse) GetDiagnosticsOk() ([]Diagnostic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Diagnostics, true
}

// SetDiagnostics sets field value
func (o *ServiceListResponse) SetDiagnostics(v []Diagnostic) {
	o.Diagnostics = v
}

// GetResponse returns the Response field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceListResponse) GetResponse() []Service {
	if o == nil {
		var ret []Service
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceListResponse) GetResponseOk() ([]Service, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ServiceListResponse) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []Service and assigns it to the Response field.
func (o *ServiceListResponse) SetResponse(v []Service) {
	o.Response = v
}

// GetSuccess returns the Success field value
func (o *ServiceListResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ServiceListResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ServiceListResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o ServiceListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["diagnostics"] = o.Diagnostics
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"diagnostics",
		"success",
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

	varServiceListResponse := _ServiceListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceListResponse)

	if err != nil {
		return err
	}

	*o = ServiceListResponse(varServiceListResponse)

	return err
}

type NullableServiceListResponse struct {
	value *ServiceListResponse
	isSet bool
}

func (v NullableServiceListResponse) Get() *ServiceListResponse {
	return v.value
}

func (v *NullableServiceListResponse) Set(val *ServiceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceListResponse(val *ServiceListResponse) *NullableServiceListResponse {
	return &NullableServiceListResponse{value: val, isSet: true}
}

func (v NullableServiceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

