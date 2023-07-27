# RegisterDevice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCode** | Pointer to **string** | Base64-encoded QR code image | [optional] 
**RefreshCode** | Pointer to **string** | Refresh code for getting a new QR code | [optional] 

## Methods

### NewRegisterDevice200Response

`func NewRegisterDevice200Response() *RegisterDevice200Response`

NewRegisterDevice200Response instantiates a new RegisterDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDevice200ResponseWithDefaults

`func NewRegisterDevice200ResponseWithDefaults() *RegisterDevice200Response`

NewRegisterDevice200ResponseWithDefaults instantiates a new RegisterDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCode

`func (o *RegisterDevice200Response) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *RegisterDevice200Response) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *RegisterDevice200Response) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *RegisterDevice200Response) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetRefreshCode

`func (o *RegisterDevice200Response) GetRefreshCode() string`

GetRefreshCode returns the RefreshCode field if non-nil, zero value otherwise.

### GetRefreshCodeOk

`func (o *RegisterDevice200Response) GetRefreshCodeOk() (*string, bool)`

GetRefreshCodeOk returns a tuple with the RefreshCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshCode

`func (o *RegisterDevice200Response) SetRefreshCode(v string)`

SetRefreshCode sets RefreshCode field to given value.

### HasRefreshCode

`func (o *RegisterDevice200Response) HasRefreshCode() bool`

HasRefreshCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


