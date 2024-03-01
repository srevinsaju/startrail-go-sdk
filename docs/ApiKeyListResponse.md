# ApiKeyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diagnostics** | [**[]Diagnostic**](Diagnostic.md) |  | 
**Response** | Pointer to [**[]ApiKey**](ApiKey.md) |  | [optional] 
**Success** | **bool** |  | 

## Methods

### NewApiKeyListResponse

`func NewApiKeyListResponse(diagnostics []Diagnostic, success bool, ) *ApiKeyListResponse`

NewApiKeyListResponse instantiates a new ApiKeyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyListResponseWithDefaults

`func NewApiKeyListResponseWithDefaults() *ApiKeyListResponse`

NewApiKeyListResponseWithDefaults instantiates a new ApiKeyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiagnostics

`func (o *ApiKeyListResponse) GetDiagnostics() []Diagnostic`

GetDiagnostics returns the Diagnostics field if non-nil, zero value otherwise.

### GetDiagnosticsOk

`func (o *ApiKeyListResponse) GetDiagnosticsOk() (*[]Diagnostic, bool)`

GetDiagnosticsOk returns a tuple with the Diagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostics

`func (o *ApiKeyListResponse) SetDiagnostics(v []Diagnostic)`

SetDiagnostics sets Diagnostics field to given value.


### GetResponse

`func (o *ApiKeyListResponse) GetResponse() []ApiKey`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ApiKeyListResponse) GetResponseOk() (*[]ApiKey, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ApiKeyListResponse) SetResponse(v []ApiKey)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ApiKeyListResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ApiKeyListResponse) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ApiKeyListResponse) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSuccess

`func (o *ApiKeyListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiKeyListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiKeyListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


