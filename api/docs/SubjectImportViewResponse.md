# SubjectImportViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**JwtSettings** | [**Import**](Import.md) |  | 
**Name** | **string** |  | 
**RemoteAccountNkeyPublic** | **string** |  | 
**RemoteSubject** | **string** |  | 

## Methods

### NewSubjectImportViewResponse

`func NewSubjectImportViewResponse(created time.Time, id string, jwtSettings Import, name string, remoteAccountNkeyPublic string, remoteSubject string, ) *SubjectImportViewResponse`

NewSubjectImportViewResponse instantiates a new SubjectImportViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectImportViewResponseWithDefaults

`func NewSubjectImportViewResponseWithDefaults() *SubjectImportViewResponse`

NewSubjectImportViewResponseWithDefaults instantiates a new SubjectImportViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SubjectImportViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubjectImportViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubjectImportViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *SubjectImportViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectImportViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectImportViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJwtSettings

`func (o *SubjectImportViewResponse) GetJwtSettings() Import`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SubjectImportViewResponse) GetJwtSettingsOk() (*Import, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SubjectImportViewResponse) SetJwtSettings(v Import)`

SetJwtSettings sets JwtSettings field to given value.


### GetName

`func (o *SubjectImportViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubjectImportViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubjectImportViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteAccountNkeyPublic

`func (o *SubjectImportViewResponse) GetRemoteAccountNkeyPublic() string`

GetRemoteAccountNkeyPublic returns the RemoteAccountNkeyPublic field if non-nil, zero value otherwise.

### GetRemoteAccountNkeyPublicOk

`func (o *SubjectImportViewResponse) GetRemoteAccountNkeyPublicOk() (*string, bool)`

GetRemoteAccountNkeyPublicOk returns a tuple with the RemoteAccountNkeyPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccountNkeyPublic

`func (o *SubjectImportViewResponse) SetRemoteAccountNkeyPublic(v string)`

SetRemoteAccountNkeyPublic sets RemoteAccountNkeyPublic field to given value.


### GetRemoteSubject

`func (o *SubjectImportViewResponse) GetRemoteSubject() string`

GetRemoteSubject returns the RemoteSubject field if non-nil, zero value otherwise.

### GetRemoteSubjectOk

`func (o *SubjectImportViewResponse) GetRemoteSubjectOk() (*string, bool)`

GetRemoteSubjectOk returns a tuple with the RemoteSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubject

`func (o *SubjectImportViewResponse) SetRemoteSubject(v string)`

SetRemoteSubject sets RemoteSubject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


