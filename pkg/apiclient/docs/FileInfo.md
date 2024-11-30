# FileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**IsDir** | Pointer to **bool** |  | [optional] 
**ModTime** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileInfo

`func NewFileInfo() *FileInfo`

NewFileInfo instantiates a new FileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileInfoWithDefaults

`func NewFileInfoWithDefaults() *FileInfo`

NewFileInfoWithDefaults instantiates a new FileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *FileInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FileInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FileInfo) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FileInfo) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetIsDir

`func (o *FileInfo) GetIsDir() bool`

GetIsDir returns the IsDir field if non-nil, zero value otherwise.

### GetIsDirOk

`func (o *FileInfo) GetIsDirOk() (*bool, bool)`

GetIsDirOk returns a tuple with the IsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDir

`func (o *FileInfo) SetIsDir(v bool)`

SetIsDir sets IsDir field to given value.

### HasIsDir

`func (o *FileInfo) HasIsDir() bool`

HasIsDir returns a boolean if a field has been set.

### GetModTime

`func (o *FileInfo) GetModTime() string`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FileInfo) GetModTimeOk() (*string, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FileInfo) SetModTime(v string)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FileInfo) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMode

`func (o *FileInfo) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FileInfo) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FileInfo) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FileInfo) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *FileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *FileInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FileInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FileInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FileInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPermissions

`func (o *FileInfo) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FileInfo) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FileInfo) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FileInfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSize

`func (o *FileInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileInfo) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


