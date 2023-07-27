# \DefaultApi

All URIs are relative to *https://wa.dup.company*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContact**](DefaultApi.md#AddContact) | **Post** /devices/{deviceId}/contacts | Add a contact using device ID
[**AddGroupMembers**](DefaultApi.md#AddGroupMembers) | **Post** /devices/{deviceId}/groups/{groupId}/addmembers | Add members to a group
[**GetAllDevices**](DefaultApi.md#GetAllDevices) | **Get** /devices | Get all registered devices
[**GetContactDetails**](DefaultApi.md#GetContactDetails) | **Get** /devices/{deviceId}/contacts/{contactId} | Get contact details by contact ID
[**GetGroupMembers**](DefaultApi.md#GetGroupMembers) | **Get** /devices/{deviceId}/groups/{groupId}/members | Get a list of members in a group
[**GetJoinedGroups**](DefaultApi.md#GetJoinedGroups) | **Get** /devices/{deviceId}/groups | Get a list of joined groups for a device
[**GetQRCodeByRefreshCode**](DefaultApi.md#GetQRCodeByRefreshCode) | **Get** /devices/{deviceId}/qrcode | Get a new QR code image using the refresh code
[**RegisterDevice**](DefaultApi.md#RegisterDevice) | **Post** /device | Register a new device for WhatsApp
[**UpdateProfile**](DefaultApi.md#UpdateProfile) | **Put** /devices/{deviceId}/profile | Update profile name or avatar of the device



## AddContact

> AddContact(ctx, deviceId).GetContactDetails200Response(getContactDetails200Response).Execute()

Add a contact using device ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    getContactDetails200Response := *openapiclient.NewGetContactDetails200Response() // GetContactDetails200Response | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddContact(context.Background(), deviceId).GetContactDetails200Response(getContactDetails200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddContact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getContactDetails200Response** | [**GetContactDetails200Response**](GetContactDetails200Response.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddGroupMembers

> AddGroupMembers(ctx, deviceId, groupId).AddGroupMembersRequest(addGroupMembersRequest).Execute()

Add members to a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    groupId := "groupId_example" // string | ID of the group
    addGroupMembersRequest := *openapiclient.NewAddGroupMembersRequest() // AddGroupMembersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AddGroupMembers(context.Background(), deviceId, groupId).AddGroupMembersRequest(addGroupMembersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addGroupMembersRequest** | [**AddGroupMembersRequest**](AddGroupMembersRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDevices

> []GetAllDevices200ResponseInner GetAllDevices(ctx).Execute()

Get all registered devices



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAllDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllDevices`: []GetAllDevices200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDevicesRequest struct via the builder pattern


### Return type

[**[]GetAllDevices200ResponseInner**](GetAllDevices200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactDetails

> GetContactDetails200Response GetContactDetails(ctx, deviceId, contactId).Execute()

Get contact details by contact ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    contactId := "contactId_example" // string | ID of the contact

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetContactDetails(context.Background(), deviceId, contactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetContactDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContactDetails`: GetContactDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetContactDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 
**contactId** | **string** | ID of the contact | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContactDetails200Response**](GetContactDetails200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMembers

> []GetGroupMembers200ResponseInner GetGroupMembers(ctx, deviceId, groupId).Execute()

Get a list of members in a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    groupId := "groupId_example" // string | ID of the group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGroupMembers(context.Background(), deviceId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMembers`: []GetGroupMembers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetGroupMembers200ResponseInner**](GetGroupMembers200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJoinedGroups

> []GetJoinedGroups200ResponseInner GetJoinedGroups(ctx, deviceId).Execute()

Get a list of joined groups for a device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetJoinedGroups(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetJoinedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJoinedGroups`: []GetJoinedGroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetJoinedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJoinedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetJoinedGroups200ResponseInner**](GetJoinedGroups200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQRCodeByRefreshCode

> GetQRCodeByRefreshCode200Response GetQRCodeByRefreshCode(ctx, deviceId).RefreshCode(refreshCode).Execute()

Get a new QR code image using the refresh code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    refreshCode := "refreshCode_example" // string | Refresh code obtained during registration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetQRCodeByRefreshCode(context.Background(), deviceId).RefreshCode(refreshCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetQRCodeByRefreshCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQRCodeByRefreshCode`: GetQRCodeByRefreshCode200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetQRCodeByRefreshCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQRCodeByRefreshCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refreshCode** | **string** | Refresh code obtained during registration | 

### Return type

[**GetQRCodeByRefreshCode200Response**](GetQRCodeByRefreshCode200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDevice

> RegisterDevice200Response RegisterDevice(ctx).RegisterDeviceRequest(registerDeviceRequest).Execute()

Register a new device for WhatsApp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registerDeviceRequest := *openapiclient.NewRegisterDeviceRequest() // RegisterDeviceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegisterDevice(context.Background()).RegisterDeviceRequest(registerDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegisterDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDevice`: RegisterDevice200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegisterDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerDeviceRequest** | [**RegisterDeviceRequest**](RegisterDeviceRequest.md) |  | 

### Return type

[**RegisterDevice200Response**](RegisterDevice200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile(ctx, deviceId).UpdateProfileRequest(updateProfileRequest).Execute()

Update profile name or avatar of the device



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deviceId := "deviceId_example" // string | ID of the registered device
    updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateProfile(context.Background(), deviceId).UpdateProfileRequest(updateProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | ID of the registered device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

