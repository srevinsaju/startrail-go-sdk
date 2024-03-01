# ServiceActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diagnostics** | [**[]Diagnostic**](Diagnostic.md) |  | 
**Response** | Pointer to [**NullableServiceActivity**](ServiceActivity.md) |  | [optional] 
**Success** | **bool** |  | 

## Methods

### NewServiceActivityResponse

`func NewServiceActivityResponse(diagnostics []Diagnostic, success bool, ) *ServiceActivityResponse`

NewServiceActivityResponse instantiates a new ServiceActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceActivityResponseWithDefaults

`func NewServiceActivityResponseWithDefaults() *ServiceActivityResponse`

NewServiceActivityResponseWithDefaults instantiates a new ServiceActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiagnostics

`func (o *ServiceActivityResponse) GetDiagnostics() []Diagnostic`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *ServiceActivityResponse) GetDiagnosticsOk() (*[]Diagnostic, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *ServiceActivityResponse) SetDiagnostics(v []Diagnostic)`

SetDiagnostics sets Diagnostics field to given value.


### GetResponse

`func (o *ServiceActivityResponse) GetResponse() ServiceActivity`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ServiceActivityResponse) GetResponseOk() (*ServiceActivity, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ServiceActivityResponse) SetResponse(v ServiceActivity)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ServiceActivityResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ServiceActivityResponse) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ServiceActivityResponse) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSuccess

`func (o *ServiceActivityResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ServiceActivityResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ServiceActivityResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


