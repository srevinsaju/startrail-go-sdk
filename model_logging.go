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

// checks if the Logging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Logging{}

// Logging struct for Logging
type Logging struct {
	Labels map[string]string `json:"labels"`
}

type _Logging Logging

// NewLogging instantiates a new Logging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogging(labels map[string]string) *Logging {
	this := Logging{}
	this.Labels = labels
	return &this
}

// NewLoggingWithDefaults instantiates a new Logging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingWithDefaults() *Logging {
	this := Logging{}
	return &this
}

// GetLabels returns the Labels field value
func (o *Logging) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Logging) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *Logging) SetLabels(v map[string]string) {
	o.Labels = v
}

func (o Logging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Logging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

func (o *Logging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labels",
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

	varLogging := _Logging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogging)

	if err != nil {
		return err
	}

	*o = Logging(varLogging)

	return err
}

type NullableLogging struct {
	value *Logging
	isSet bool
}

func (v NullableLogging) Get() *Logging {
	return v.value
}

func (v *NullableLogging) Set(val *Logging) {
	v.value = val
	v.isSet = true
}

func (v NullableLogging) IsSet() bool {
	return v.isSet
}

func (v *NullableLogging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogging(val *Logging) *NullableLogging {
	return &NullableLogging{value: val, isSet: true}
}

func (v NullableLogging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

