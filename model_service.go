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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	Access []Access `json:"access,omitempty"`
	Description string `json:"description"`
	Disabled *bool `json:"disabled,omitempty"`
	Environment string `json:"environment"`
	Logging map[string]Logging `json:"logging,omitempty"`
	Metadata NullableMetadata `json:"metadata,omitempty"`
	Name string `json:"name"`
	Remarks string `json:"remarks"`
	Sources map[string]Source `json:"sources,omitempty"`
	Tenant string `json:"tenant"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	UpdatedBy NullableString `json:"updated_by,omitempty"`
	UpdatedDate NullableString `json:"updated_date,omitempty"`
}

type _Service Service

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(description string, environment string, name string, remarks string, tenant string) *Service {
	this := Service{}
	this.Description = description
	this.Environment = environment
	this.Name = name
	this.Remarks = remarks
	this.Tenant = tenant
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Service) GetAccess() []Access {
	if o == nil || IsNil(o.Access) {
		var ret []Access
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessOk() ([]Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *Service) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []Access and assigns it to the Access field.
func (o *Service) SetAccess(v []Access) {
	o.Access = v
}

// GetDescription returns the Description field value
func (o *Service) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Service) SetDescription(v string) {
	o.Description = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *Service) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *Service) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *Service) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEnvironment returns the Environment field value
func (o *Service) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Service) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Service) SetEnvironment(v string) {
	o.Environment = v
}

// GetLogging returns the Logging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetLogging() map[string]Logging {
	if o == nil {
		var ret map[string]Logging
		return ret
	}
	return o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetLoggingOk() (*map[string]Logging, bool) {
	if o == nil || IsNil(o.Logging) {
		return nil, false
	}
	return &o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *Service) HasLogging() bool {
	if o != nil && !IsNil(o.Logging) {
		return true
	}

	return false
}

// SetLogging gets a reference to the given map[string]Logging and assigns it to the Logging field.
func (o *Service) SetLogging(v map[string]Logging) {
	o.Logging = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret Metadata
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *Service) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableMetadata and assigns it to the Metadata field.
func (o *Service) SetMetadata(v Metadata) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *Service) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *Service) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetRemarks returns the Remarks field value
func (o *Service) GetRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value
// and a boolean to check if the value has been set.
func (o *Service) GetRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remarks, true
}

// SetRemarks sets field value
func (o *Service) SetRemarks(v string) {
	o.Remarks = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetSources() map[string]Source {
	if o == nil {
		var ret map[string]Source
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetSourcesOk() (*map[string]Source, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *Service) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given map[string]Source and assigns it to the Sources field.
func (o *Service) SetSources(v map[string]Source) {
	o.Sources = v
}

// GetTenant returns the Tenant field value
func (o *Service) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *Service) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *Service) SetTenant(v string) {
	o.Tenant = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Service) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *Service) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Service) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Service) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Service) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *Service) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *Service) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *Service) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Service) GetUpdatedDate() string {
	if o == nil || IsNil(o.UpdatedDate.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedDate.Get()
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Service) GetUpdatedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedDate.Get(), o.UpdatedDate.IsSet()
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *Service) HasUpdatedDate() bool {
	if o != nil && o.UpdatedDate.IsSet() {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given NullableString and assigns it to the UpdatedDate field.
func (o *Service) SetUpdatedDate(v string) {
	o.UpdatedDate.Set(&v)
}
// SetUpdatedDateNil sets the value for UpdatedDate to be an explicit nil
func (o *Service) SetUpdatedDateNil() {
	o.UpdatedDate.Set(nil)
}

// UnsetUpdatedDate ensures that no value is present for UpdatedDate, not even an explicit nil
func (o *Service) UnsetUpdatedDate() {
	o.UpdatedDate.Unset()
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	toSerialize["environment"] = o.Environment
	if o.Logging != nil {
		toSerialize["logging"] = o.Logging
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["remarks"] = o.Remarks
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["tenant"] = o.Tenant
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}
	if o.UpdatedDate.IsSet() {
		toSerialize["updated_date"] = o.UpdatedDate.Get()
	}
	return toSerialize, nil
}

func (o *Service) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"environment",
		"name",
		"remarks",
		"tenant",
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

	varService := _Service{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varService)

	if err != nil {
		return err
	}

	*o = Service(varService)

	return err
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


