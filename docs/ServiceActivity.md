# ServiceActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **int64** |  | [optional] 

## Methods

### NewServiceActivity

`func NewServiceActivity() *ServiceActivity`

NewServiceActivity instantiates a new ServiceActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceActivityWithDefaults

`func NewServiceActivityWithDefaults() *ServiceActivity`

NewServiceActivityWithDefaults instantiates a new ServiceActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ServiceActivity) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ServiceActivity) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ServiceActivity) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ServiceActivity) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *ServiceActivity) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *ServiceActivity) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetValue

`func (o *ServiceActivity) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceActivity) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceActivity) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceActivity) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


