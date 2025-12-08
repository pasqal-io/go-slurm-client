# SlurmV0041PostNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**CpuBind** | Pointer to **int32** | Default method for binding tasks to allocated CPUs | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Features** | Pointer to **[]string** | Available features | [optional] 
**FeaturesAct** | Pointer to **[]string** | Currently active features | [optional] 
**Gres** | Pointer to **string** | Generic resources | [optional] 
**Address** | Pointer to **[]string** | NodeAddr, used to establish a communication path | [optional] 
**Hostname** | Pointer to **[]string** | NodeHostname | [optional] 
**Name** | Pointer to **[]string** | NodeName | [optional] 
**State** | Pointer to **[]string** | New state to assign to the node | [optional] 
**Reason** | Pointer to **string** | Reason for node being DOWN or DRAINING | [optional] 
**ReasonUid** | Pointer to **string** | User ID to associate with the reason (needed if user root is sending message) | [optional] 
**ResumeAfter** | Pointer to [**SlurmV0041PostNodeRequestResumeAfter**](SlurmV0041PostNodeRequestResumeAfter.md) |  | [optional] 
**Weight** | Pointer to [**SlurmV0041PostNodeRequestWeight**](SlurmV0041PostNodeRequestWeight.md) |  | [optional] 

## Methods

### NewSlurmV0041PostNodeRequest

`func NewSlurmV0041PostNodeRequest() *SlurmV0041PostNodeRequest`

NewSlurmV0041PostNodeRequest instantiates a new SlurmV0041PostNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041PostNodeRequestWithDefaults

`func NewSlurmV0041PostNodeRequestWithDefaults() *SlurmV0041PostNodeRequest`

NewSlurmV0041PostNodeRequestWithDefaults instantiates a new SlurmV0041PostNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SlurmV0041PostNodeRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SlurmV0041PostNodeRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SlurmV0041PostNodeRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SlurmV0041PostNodeRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCpuBind

`func (o *SlurmV0041PostNodeRequest) GetCpuBind() int32`

GetCpuBind returns the CpuBind field if non-nil, zero value otherwise.

### GetCpuBindOk

`func (o *SlurmV0041PostNodeRequest) GetCpuBindOk() (*int32, bool)`

GetCpuBindOk returns a tuple with the CpuBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBind

`func (o *SlurmV0041PostNodeRequest) SetCpuBind(v int32)`

SetCpuBind sets CpuBind field to given value.

### HasCpuBind

`func (o *SlurmV0041PostNodeRequest) HasCpuBind() bool`

HasCpuBind returns a boolean if a field has been set.

### GetExtra

`func (o *SlurmV0041PostNodeRequest) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SlurmV0041PostNodeRequest) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SlurmV0041PostNodeRequest) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *SlurmV0041PostNodeRequest) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetFeatures

`func (o *SlurmV0041PostNodeRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SlurmV0041PostNodeRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SlurmV0041PostNodeRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SlurmV0041PostNodeRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFeaturesAct

`func (o *SlurmV0041PostNodeRequest) GetFeaturesAct() []string`

GetFeaturesAct returns the FeaturesAct field if non-nil, zero value otherwise.

### GetFeaturesActOk

`func (o *SlurmV0041PostNodeRequest) GetFeaturesActOk() (*[]string, bool)`

GetFeaturesActOk returns a tuple with the FeaturesAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturesAct

`func (o *SlurmV0041PostNodeRequest) SetFeaturesAct(v []string)`

SetFeaturesAct sets FeaturesAct field to given value.

### HasFeaturesAct

`func (o *SlurmV0041PostNodeRequest) HasFeaturesAct() bool`

HasFeaturesAct returns a boolean if a field has been set.

### GetGres

`func (o *SlurmV0041PostNodeRequest) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *SlurmV0041PostNodeRequest) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *SlurmV0041PostNodeRequest) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *SlurmV0041PostNodeRequest) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetAddress

`func (o *SlurmV0041PostNodeRequest) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SlurmV0041PostNodeRequest) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SlurmV0041PostNodeRequest) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SlurmV0041PostNodeRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *SlurmV0041PostNodeRequest) GetHostname() []string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SlurmV0041PostNodeRequest) GetHostnameOk() (*[]string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SlurmV0041PostNodeRequest) SetHostname(v []string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SlurmV0041PostNodeRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetName

`func (o *SlurmV0041PostNodeRequest) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlurmV0041PostNodeRequest) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlurmV0041PostNodeRequest) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *SlurmV0041PostNodeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *SlurmV0041PostNodeRequest) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SlurmV0041PostNodeRequest) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SlurmV0041PostNodeRequest) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *SlurmV0041PostNodeRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *SlurmV0041PostNodeRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SlurmV0041PostNodeRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SlurmV0041PostNodeRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SlurmV0041PostNodeRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonUid

`func (o *SlurmV0041PostNodeRequest) GetReasonUid() string`

GetReasonUid returns the ReasonUid field if non-nil, zero value otherwise.

### GetReasonUidOk

`func (o *SlurmV0041PostNodeRequest) GetReasonUidOk() (*string, bool)`

GetReasonUidOk returns a tuple with the ReasonUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUid

`func (o *SlurmV0041PostNodeRequest) SetReasonUid(v string)`

SetReasonUid sets ReasonUid field to given value.

### HasReasonUid

`func (o *SlurmV0041PostNodeRequest) HasReasonUid() bool`

HasReasonUid returns a boolean if a field has been set.

### GetResumeAfter

`func (o *SlurmV0041PostNodeRequest) GetResumeAfter() SlurmV0041PostNodeRequestResumeAfter`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *SlurmV0041PostNodeRequest) GetResumeAfterOk() (*SlurmV0041PostNodeRequestResumeAfter, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *SlurmV0041PostNodeRequest) SetResumeAfter(v SlurmV0041PostNodeRequestResumeAfter)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *SlurmV0041PostNodeRequest) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetWeight

`func (o *SlurmV0041PostNodeRequest) GetWeight() SlurmV0041PostNodeRequestWeight`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SlurmV0041PostNodeRequest) GetWeightOk() (*SlurmV0041PostNodeRequestWeight, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SlurmV0041PostNodeRequest) SetWeight(v SlurmV0041PostNodeRequestWeight)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SlurmV0041PostNodeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


