# RegisterDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Name of the device | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number with extension (e.g., +1234567890) | [optional] 

## Methods

### NewRegisterDeviceRequest

`func NewRegisterDeviceRequest() *RegisterDeviceRequest`

NewRegisterDeviceRequest instantiates a new RegisterDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeviceRequestWithDefaults

`func NewRegisterDeviceRequestWithDefaults() *RegisterDeviceRequest`

NewRegisterDeviceRequestWithDefaults instantiates a new RegisterDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *RegisterDeviceRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *RegisterDeviceRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *RegisterDeviceRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *RegisterDeviceRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *RegisterDeviceRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RegisterDeviceRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RegisterDeviceRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *RegisterDeviceRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


