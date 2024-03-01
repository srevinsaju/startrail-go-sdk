/*
startrail-service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServiceActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceActivity{}

// ServiceActivity struct for ServiceActivity
type ServiceActivity struct {
	Date NullableString `json:"date,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

// NewServiceActivity instantiates a new ServiceActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceActivity() *ServiceActivity {
	this := ServiceActivity{}
	return &this
}

// NewServiceActivityWithDefaults instantiates a new ServiceActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceActivityWithDefaults() *ServiceActivity {
	this := ServiceActivity{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceActivity) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceActivity) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ServiceActivity) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *ServiceActivity) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ServiceActivity) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ServiceActivity) UnsetDate() {
	o.Date.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ServiceActivity) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceActivity) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ServiceActivity) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *ServiceActivity) SetValue(v int64) {
	o.Value = &v
}

func (o ServiceActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableServiceActivity struct {
	value *ServiceActivity
	isSet bool
}

func (v NullableServiceActivity) Get() *ServiceActivity {
	return v.value
}

func (v *NullableServiceActivity) Set(val *ServiceActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceActivity(val *ServiceActivity) *NullableServiceActivity {
	return &NullableServiceActivity{value: val, isSet: true}
}

func (v NullableServiceActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


