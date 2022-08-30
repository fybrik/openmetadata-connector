# ShardSearchFailureCauseStackTraceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassLoaderName** | Pointer to **string** |  | [optional] 
**ModuleName** | Pointer to **string** |  | [optional] 
**ModuleVersion** | Pointer to **string** |  | [optional] 
**MethodName** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**NativeMethod** | Pointer to **bool** |  | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 

## Methods

### NewShardSearchFailureCauseStackTraceInner

`func NewShardSearchFailureCauseStackTraceInner() *ShardSearchFailureCauseStackTraceInner`

NewShardSearchFailureCauseStackTraceInner instantiates a new ShardSearchFailureCauseStackTraceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardSearchFailureCauseStackTraceInnerWithDefaults

`func NewShardSearchFailureCauseStackTraceInnerWithDefaults() *ShardSearchFailureCauseStackTraceInner`

NewShardSearchFailureCauseStackTraceInnerWithDefaults instantiates a new ShardSearchFailureCauseStackTraceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassLoaderName

`func (o *ShardSearchFailureCauseStackTraceInner) GetClassLoaderName() string`

GetClassLoaderName returns the ClassLoaderName field if non-nil, zero value otherwise.

### GetClassLoaderNameOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetClassLoaderNameOk() (*string, bool)`

GetClassLoaderNameOk returns a tuple with the ClassLoaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassLoaderName

`func (o *ShardSearchFailureCauseStackTraceInner) SetClassLoaderName(v string)`

SetClassLoaderName sets ClassLoaderName field to given value.

### HasClassLoaderName

`func (o *ShardSearchFailureCauseStackTraceInner) HasClassLoaderName() bool`

HasClassLoaderName returns a boolean if a field has been set.

### GetModuleName

`func (o *ShardSearchFailureCauseStackTraceInner) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *ShardSearchFailureCauseStackTraceInner) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *ShardSearchFailureCauseStackTraceInner) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetModuleVersion

`func (o *ShardSearchFailureCauseStackTraceInner) GetModuleVersion() string`

GetModuleVersion returns the ModuleVersion field if non-nil, zero value otherwise.

### GetModuleVersionOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetModuleVersionOk() (*string, bool)`

GetModuleVersionOk returns a tuple with the ModuleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleVersion

`func (o *ShardSearchFailureCauseStackTraceInner) SetModuleVersion(v string)`

SetModuleVersion sets ModuleVersion field to given value.

### HasModuleVersion

`func (o *ShardSearchFailureCauseStackTraceInner) HasModuleVersion() bool`

HasModuleVersion returns a boolean if a field has been set.

### GetMethodName

`func (o *ShardSearchFailureCauseStackTraceInner) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *ShardSearchFailureCauseStackTraceInner) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.

### HasMethodName

`func (o *ShardSearchFailureCauseStackTraceInner) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.

### GetFileName

`func (o *ShardSearchFailureCauseStackTraceInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ShardSearchFailureCauseStackTraceInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ShardSearchFailureCauseStackTraceInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetLineNumber

`func (o *ShardSearchFailureCauseStackTraceInner) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ShardSearchFailureCauseStackTraceInner) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *ShardSearchFailureCauseStackTraceInner) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetNativeMethod

`func (o *ShardSearchFailureCauseStackTraceInner) GetNativeMethod() bool`

GetNativeMethod returns the NativeMethod field if non-nil, zero value otherwise.

### GetNativeMethodOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetNativeMethodOk() (*bool, bool)`

GetNativeMethodOk returns a tuple with the NativeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeMethod

`func (o *ShardSearchFailureCauseStackTraceInner) SetNativeMethod(v bool)`

SetNativeMethod sets NativeMethod field to given value.

### HasNativeMethod

`func (o *ShardSearchFailureCauseStackTraceInner) HasNativeMethod() bool`

HasNativeMethod returns a boolean if a field has been set.

### GetClassName

`func (o *ShardSearchFailureCauseStackTraceInner) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ShardSearchFailureCauseStackTraceInner) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ShardSearchFailureCauseStackTraceInner) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ShardSearchFailureCauseStackTraceInner) HasClassName() bool`

HasClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


