# \SlurmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbV0041DeleteAccount**](SlurmdbAPI.md#SlurmdbV0041DeleteAccount) | **Delete** /slurmdb/v0.0.41/account/{account_name} | Delete account
[**SlurmdbV0041DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociation) | **Delete** /slurmdb/v0.0.41/association/ | Delete association
[**SlurmdbV0041DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociations) | **Delete** /slurmdb/v0.0.41/associations/ | Delete associations
[**SlurmdbV0041DeleteCluster**](SlurmdbAPI.md#SlurmdbV0041DeleteCluster) | **Delete** /slurmdb/v0.0.41/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0041DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0041DeleteSingleQos) | **Delete** /slurmdb/v0.0.41/qos/{qos} | Delete QOS
[**SlurmdbV0041DeleteUser**](SlurmdbAPI.md#SlurmdbV0041DeleteUser) | **Delete** /slurmdb/v0.0.41/user/{name} | Delete user
[**SlurmdbV0041DeleteWckey**](SlurmdbAPI.md#SlurmdbV0041DeleteWckey) | **Delete** /slurmdb/v0.0.41/wckey/{id} | Delete wckey
[**SlurmdbV0041GetAccount**](SlurmdbAPI.md#SlurmdbV0041GetAccount) | **Get** /slurmdb/v0.0.41/account/{account_name} | Get account info
[**SlurmdbV0041GetAccounts**](SlurmdbAPI.md#SlurmdbV0041GetAccounts) | **Get** /slurmdb/v0.0.41/accounts/ | Get account list
[**SlurmdbV0041GetAssociation**](SlurmdbAPI.md#SlurmdbV0041GetAssociation) | **Get** /slurmdb/v0.0.41/association/ | Get association info
[**SlurmdbV0041GetAssociations**](SlurmdbAPI.md#SlurmdbV0041GetAssociations) | **Get** /slurmdb/v0.0.41/associations/ | Get association list
[**SlurmdbV0041GetCluster**](SlurmdbAPI.md#SlurmdbV0041GetCluster) | **Get** /slurmdb/v0.0.41/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0041GetClusters**](SlurmdbAPI.md#SlurmdbV0041GetClusters) | **Get** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041GetConfig**](SlurmdbAPI.md#SlurmdbV0041GetConfig) | **Get** /slurmdb/v0.0.41/config | Dump all configuration information
[**SlurmdbV0041GetDiag**](SlurmdbAPI.md#SlurmdbV0041GetDiag) | **Get** /slurmdb/v0.0.41/diag/ | Get slurmdb diagnostics
[**SlurmdbV0041GetInstance**](SlurmdbAPI.md#SlurmdbV0041GetInstance) | **Get** /slurmdb/v0.0.41/instance/ | Get instance info
[**SlurmdbV0041GetInstances**](SlurmdbAPI.md#SlurmdbV0041GetInstances) | **Get** /slurmdb/v0.0.41/instances/ | Get instance list
[**SlurmdbV0041GetJob**](SlurmdbAPI.md#SlurmdbV0041GetJob) | **Get** /slurmdb/v0.0.41/job/{job_id} | Get job info
[**SlurmdbV0041GetJobs**](SlurmdbAPI.md#SlurmdbV0041GetJobs) | **Get** /slurmdb/v0.0.41/jobs/ | Get job list
[**SlurmdbV0041GetQos**](SlurmdbAPI.md#SlurmdbV0041GetQos) | **Get** /slurmdb/v0.0.41/qos/ | Get QOS list
[**SlurmdbV0041GetSingleQos**](SlurmdbAPI.md#SlurmdbV0041GetSingleQos) | **Get** /slurmdb/v0.0.41/qos/{qos} | Get QOS info
[**SlurmdbV0041GetTres**](SlurmdbAPI.md#SlurmdbV0041GetTres) | **Get** /slurmdb/v0.0.41/tres/ | Get TRES info
[**SlurmdbV0041GetUser**](SlurmdbAPI.md#SlurmdbV0041GetUser) | **Get** /slurmdb/v0.0.41/user/{name} | Get user info
[**SlurmdbV0041GetUsers**](SlurmdbAPI.md#SlurmdbV0041GetUsers) | **Get** /slurmdb/v0.0.41/users/ | Get user list
[**SlurmdbV0041GetWckey**](SlurmdbAPI.md#SlurmdbV0041GetWckey) | **Get** /slurmdb/v0.0.41/wckey/{id} | Get wckey info
[**SlurmdbV0041GetWckeys**](SlurmdbAPI.md#SlurmdbV0041GetWckeys) | **Get** /slurmdb/v0.0.41/wckeys/ | Get wckey list
[**SlurmdbV0041PostAccounts**](SlurmdbAPI.md#SlurmdbV0041PostAccounts) | **Post** /slurmdb/v0.0.41/accounts/ | Add/update list of accounts
[**SlurmdbV0041PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0041PostAccountsAssociation) | **Post** /slurmdb/v0.0.41/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0041PostAssociations**](SlurmdbAPI.md#SlurmdbV0041PostAssociations) | **Post** /slurmdb/v0.0.41/associations/ | Set associations info
[**SlurmdbV0041PostClusters**](SlurmdbAPI.md#SlurmdbV0041PostClusters) | **Post** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041PostConfig**](SlurmdbAPI.md#SlurmdbV0041PostConfig) | **Post** /slurmdb/v0.0.41/config | Load all configuration information
[**SlurmdbV0041PostQos**](SlurmdbAPI.md#SlurmdbV0041PostQos) | **Post** /slurmdb/v0.0.41/qos/ | Add or update QOSs
[**SlurmdbV0041PostTres**](SlurmdbAPI.md#SlurmdbV0041PostTres) | **Post** /slurmdb/v0.0.41/tres/ | Add TRES
[**SlurmdbV0041PostUsers**](SlurmdbAPI.md#SlurmdbV0041PostUsers) | **Post** /slurmdb/v0.0.41/users/ | Update users
[**SlurmdbV0041PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0041PostUsersAssociation) | **Post** /slurmdb/v0.0.41/users_association/ | Add users with conditional association
[**SlurmdbV0041PostWckeys**](SlurmdbAPI.md#SlurmdbV0041PostWckeys) | **Post** /slurmdb/v0.0.41/wckeys/ | Add or update wckeys



## SlurmdbV0041DeleteAccount

> SlurmdbV0041DeleteAccount200Response SlurmdbV0041DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAccount`: SlurmdbV0041DeleteAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteAccount200Response**](SlurmdbV0041DeleteAccount200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociation

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociation`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociations

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociations`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteCluster

> SlurmdbV0041DeleteCluster200Response SlurmdbV0041DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteCluster`: SlurmdbV0041DeleteCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**SlurmdbV0041DeleteCluster200Response**](SlurmdbV0041DeleteCluster200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteSingleQos

> SlurmdbV0041DeleteSingleQos200Response SlurmdbV0041DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteSingleQos`: SlurmdbV0041DeleteSingleQos200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteSingleQos200Response**](SlurmdbV0041DeleteSingleQos200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteUser

> V0041OpenapiResp SlurmdbV0041DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteUser`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteWckey

> SlurmdbV0041DeleteWckey200Response SlurmdbV0041DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteWckey`: SlurmdbV0041DeleteWckey200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteWckey200Response**](SlurmdbV0041DeleteWckey200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccount

> V0041OpenapiAccountsResp SlurmdbV0041GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccount`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccounts

> V0041OpenapiAccountsResp SlurmdbV0041GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccounts`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociation

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociation`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociations

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociations`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetCluster

> V0041OpenapiClustersResp SlurmdbV0041GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetCluster`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetClusters

> V0041OpenapiClustersResp SlurmdbV0041GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetClusters`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetConfig

> V0041OpenapiSlurmdbdConfigResp SlurmdbV0041GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetConfig`: V0041OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetConfigRequest struct via the builder pattern


### Return type

[**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetDiag

> SlurmdbV0041GetDiag200Response SlurmdbV0041GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetDiag`: SlurmdbV0041GetDiag200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetDiagRequest struct via the builder pattern


### Return type

[**SlurmdbV0041GetDiag200Response**](SlurmdbV0041GetDiag200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstance

> V0041OpenapiInstancesResp SlurmdbV0041GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstance`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstances

> V0041OpenapiInstancesResp SlurmdbV0041GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstances`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJob

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	jobId := "jobId_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJob`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJobs

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV wckey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJobs`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV wckey list | 

### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetSingleQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetSingleQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetTres

> V0041OpenapiTresResp SlurmdbV0041GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetTres`: V0041OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetTresRequest struct via the builder pattern


### Return type

[**V0041OpenapiTresResp**](V0041OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUser

> V0041OpenapiUsersResp SlurmdbV0041GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUser`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include wckeys | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUsers

> V0041OpenapiUsersResp SlurmdbV0041GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default wckey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With wckeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUsers`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default wckey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With wckeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckey

> V0041OpenapiWckeyResp SlurmdbV0041GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckey`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckeys

> V0041OpenapiWckeyResp SlurmdbV0041GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckeys`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 

### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccounts

> V0041OpenapiResp SlurmdbV0041PostAccounts(ctx).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	v0041OpenapiAccountsResp := *openapiclient.NewV0041OpenapiAccountsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAccountsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAccountsInner("Description_example", "Name_example", "Organization_example")}) // V0041OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccounts(context.Background()).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccounts`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAccountsResp** | [**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccountsAssociation

> SlurmdbV0041PostAccountsAssociation200Response SlurmdbV0041PostAccountsAssociation(ctx).SlurmdbV0041PostAccountsAssociationRequest(slurmdbV0041PostAccountsAssociationRequest).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	slurmdbV0041PostAccountsAssociationRequest := *openapiclient.NewSlurmdbV0041PostAccountsAssociationRequest() // SlurmdbV0041PostAccountsAssociationRequest | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccountsAssociation(context.Background()).SlurmdbV0041PostAccountsAssociationRequest(slurmdbV0041PostAccountsAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccountsAssociation`: SlurmdbV0041PostAccountsAssociation200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmdbV0041PostAccountsAssociationRequest** | [**SlurmdbV0041PostAccountsAssociationRequest**](SlurmdbV0041PostAccountsAssociationRequest.md) | Add list of accounts with conditional association | 

### Return type

[**SlurmdbV0041PostAccountsAssociation200Response**](SlurmdbV0041PostAccountsAssociation200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAssociations

> V0041OpenapiResp SlurmdbV0041PostAssociations(ctx).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	v0041OpenapiAssocsResp := *openapiclient.NewV0041OpenapiAssocsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAssociationsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAssociationsInner("User_example")}) // V0041OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAssociations(context.Background()).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAssociations`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAssocsResp** | [**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostClusters

> V0041OpenapiResp SlurmdbV0041PostClusters(ctx).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)
	v0041OpenapiClustersResp := *openapiclient.NewV0041OpenapiClustersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespClustersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespClustersInner()}) // V0041OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostClusters(context.Background()).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostClusters`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 
 **v0041OpenapiClustersResp** | [**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostConfig

> V0041OpenapiResp SlurmdbV0041PostConfig(ctx).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	v0041OpenapiSlurmdbdConfigResp := *openapiclient.NewV0041OpenapiSlurmdbdConfigResp() // V0041OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostConfig(context.Background()).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostConfig`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiSlurmdbdConfigResp** | [**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostQos

> V0041OpenapiResp SlurmdbV0041PostQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)
	v0041OpenapiSlurmdbdQosResp := *openapiclient.NewV0041OpenapiSlurmdbdQosResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespQosInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespQosInner()}) // V0041OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostQos`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 
 **v0041OpenapiSlurmdbdQosResp** | [**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostTres

> V0041OpenapiResp SlurmdbV0041PostTres(ctx).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	v0041OpenapiTresResp := *openapiclient.NewV0041OpenapiTresResp([]openapiclient.SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner{*openapiclient.NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner("Type_example")}) // V0041OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostTres(context.Background()).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostTres`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiTresResp** | [**V0041OpenapiTresResp**](V0041OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsers

> V0041OpenapiResp SlurmdbV0041PostUsers(ctx).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	v0041OpenapiUsersResp := *openapiclient.NewV0041OpenapiUsersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInner("Name_example")}) // V0041OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsers(context.Background()).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsers`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiUsersResp** | [**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsersAssociation

> SlurmdbV0041PostUsersAssociation200Response SlurmdbV0041PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).SlurmdbV0041PostUsersAssociationRequest(slurmdbV0041PostUsersAssociationRequest).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)
	slurmdbV0041PostUsersAssociationRequest := *openapiclient.NewSlurmdbV0041PostUsersAssociationRequest(*openapiclient.NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition([]string{"Users_example"}), *openapiclient.NewSlurmdbV0041PostUsersAssociationRequestUser()) // SlurmdbV0041PostUsersAssociationRequest | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).SlurmdbV0041PostUsersAssociationRequest(slurmdbV0041PostUsersAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsersAssociation`: SlurmdbV0041PostUsersAssociation200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 
 **slurmdbV0041PostUsersAssociationRequest** | [**SlurmdbV0041PostUsersAssociationRequest**](SlurmdbV0041PostUsersAssociationRequest.md) | Create users with conditional association | 

### Return type

[**SlurmdbV0041PostUsersAssociation200Response**](SlurmdbV0041PostUsersAssociation200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostWckeys

> V0041OpenapiResp SlurmdbV0041PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)
	v0041OpenapiWckeyResp := *openapiclient.NewV0041OpenapiWckeyResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner("Cluster_example", "Name_example", "User_example")}) // V0041OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostWckeys`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 
 **v0041OpenapiWckeyResp** | [**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

