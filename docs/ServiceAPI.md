# \ServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activity**](ServiceAPI.md#Activity) | **Get** /api/v1/service/{tenant}/{environment}/{name}/activity | 
[**Create**](ServiceAPI.md#Create) | **Post** /api/v1/service | 
[**Delete**](ServiceAPI.md#Delete) | **Delete** /api/v1/service/{tenant}/{environment}/{name} | 
[**Get**](ServiceAPI.md#Get) | **Get** /api/v1/service/{tenant}/{environment}/{name} | 
[**History**](ServiceAPI.md#History) | **Get** /api/v1/service/{tenant}/{environment}/{name}/history | 
[**List**](ServiceAPI.md#List) | **Get** /api/v1/service/{tenant} | 



## Activity

> ServiceActivityResponse Activity(ctx, tenant, environment, name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	tenant := "tenant_example" // string | Name of the tenant.
	environment := "environment_example" // string | Name of the environment.
	name := "name_example" // string | Name of the service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.Activity(context.Background(), tenant, environment, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.Activity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Activity`: ServiceActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.Activity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Name of the tenant. | 
**environment** | **string** | Name of the environment. | 
**name** | **string** | Name of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceActivityResponse**](ServiceActivityResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> ServiceResponse Create(ctx).Service(service).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	service := *openapiclient.NewService("This is a hello world service.", "production", "hello-world", "Make sure this service prints hello world on /", "startrail") // Service | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.Create(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: ServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**Service**](Service.md) |  | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> StringResponse Delete(ctx, tenant, environment, name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	tenant := "tenant_example" // string | Name of the tenant.
	environment := "environment_example" // string | Name of the environment.
	name := "name_example" // string | Name of the service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.Delete(context.Background(), tenant, environment, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: StringResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Name of the tenant. | 
**environment** | **string** | Name of the environment. | 
**name** | **string** | Name of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StringResponse**](StringResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> ServiceResponse Get(ctx, tenant, environment, name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	tenant := "tenant_example" // string | Name of the tenant.
	environment := "environment_example" // string | Name of the environment.
	name := "name_example" // string | Name of the service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.Get(context.Background(), tenant, environment, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Name of the tenant. | 
**environment** | **string** | Name of the environment. | 
**name** | **string** | Name of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## History

> ServiceListResponse History(ctx, tenant, environment, name).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	tenant := "tenant_example" // string | Name of the tenant.
	environment := "environment_example" // string | Name of the environment.
	name := "name_example" // string | Name of the service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.History(context.Background(), tenant, environment, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.History``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `History`: ServiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.History`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Name of the tenant. | 
**environment** | **string** | Name of the environment. | 
**name** | **string** | Name of the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServiceListResponse**](ServiceListResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ServiceListResponse List(ctx, tenant).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/srevinsaju/startrail-go-sdk"
)

func main() {
	tenant := "tenant_example" // string | Name of the tenant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.List(context.Background(), tenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ServiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenant** | **string** | Name of the tenant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceListResponse**](ServiceListResponse.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

