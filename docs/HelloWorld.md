# HelloWorld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagline** | **string** |  | 
**Version** | [**Version**](Version.md) |  | 

## Methods

### NewHelloWorld

`func NewHelloWorld(tagline string, version Version, ) *HelloWorld`

NewHelloWorld instantiates a new HelloWorld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelloWorldWithDefaults

`func NewHelloWorldWithDefaults() *HelloWorld`

NewHelloWorldWithDefaults instantiates a new HelloWorld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagline

`func (o *HelloWorld) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *HelloWorld) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *HelloWorld) SetTagline(v string)`

SetTagline sets Tagline field to given value.


### GetVersion

`func (o *HelloWorld) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HelloWorld) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HelloWorld) SetVersion(v Version)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


