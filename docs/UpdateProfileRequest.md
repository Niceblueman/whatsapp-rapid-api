# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileName** | Pointer to **string** | Updated profile name for the device | [optional] 
**AvatarUrl** | Pointer to **string** | Valid image URL for the device&#39;s avatar | [optional] 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileName

`func (o *UpdateProfileRequest) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *UpdateProfileRequest) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *UpdateProfileRequest) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *UpdateProfileRequest) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UpdateProfileRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UpdateProfileRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UpdateProfileRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UpdateProfileRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


