# Diagnostic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** |  | 
**Detail** | **string** |  | 
**Severity** | [**DiagSeverity**](DiagSeverity.md) |  | 
**Summary** | **string** |  | 

## Methods

### NewDiagnostic

`func NewDiagnostic(context string, detail string, severity DiagSeverity, summary string, ) *Diagnostic`

NewDiagnostic instantiates a new Diagnostic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticWithDefaults

`func NewDiagnosticWithDefaults() *Diagnostic`

NewDiagnosticWithDefaults instantiates a new Diagnostic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Diagnostic) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Diagnostic) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Diagnostic) SetContext(v string)`

SetContext sets Context field to given value.


### GetDetail

`func (o *Diagnostic) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Diagnostic) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Diagnostic) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetSeverity

`func (o *Diagnostic) GetSeverity() DiagSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Diagnostic) GetSeverityOk() (*DiagSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Diagnostic) SetSeverity(v DiagSeverity)`

SetSeverity sets Severity field to given value.


### GetSummary

`func (o *Diagnostic) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Diagnostic) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Diagnostic) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


