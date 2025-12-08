# \SlurmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmV0041DeleteJob**](SlurmAPI.md#SlurmV0041DeleteJob) | **Delete** /slurm/v0.0.41/job/{job_id} | cancel or signal job
[**SlurmV0041DeleteJobs**](SlurmAPI.md#SlurmV0041DeleteJobs) | **Delete** /slurm/v0.0.41/jobs/ | send signal to list of jobs
[**SlurmV0041DeleteNode**](SlurmAPI.md#SlurmV0041DeleteNode) | **Delete** /slurm/v0.0.41/node/{node_name} | delete node
[**SlurmV0041GetDiag**](SlurmAPI.md#SlurmV0041GetDiag) | **Get** /slurm/v0.0.41/diag/ | get diagnostics
[**SlurmV0041GetJob**](SlurmAPI.md#SlurmV0041GetJob) | **Get** /slurm/v0.0.41/job/{job_id} | get job info
[**SlurmV0041GetJobs**](SlurmAPI.md#SlurmV0041GetJobs) | **Get** /slurm/v0.0.41/jobs/ | get list of jobs
[**SlurmV0041GetJobsState**](SlurmAPI.md#SlurmV0041GetJobsState) | **Get** /slurm/v0.0.41/jobs/state/ | get list of job states
[**SlurmV0041GetLicenses**](SlurmAPI.md#SlurmV0041GetLicenses) | **Get** /slurm/v0.0.41/licenses/ | get all Slurm tracked license info
[**SlurmV0041GetNode**](SlurmAPI.md#SlurmV0041GetNode) | **Get** /slurm/v0.0.41/node/{node_name} | get node info
[**SlurmV0041GetNodes**](SlurmAPI.md#SlurmV0041GetNodes) | **Get** /slurm/v0.0.41/nodes/ | get node(s) info
[**SlurmV0041GetPartition**](SlurmAPI.md#SlurmV0041GetPartition) | **Get** /slurm/v0.0.41/partition/{partition_name} | get partition info
[**SlurmV0041GetPartitions**](SlurmAPI.md#SlurmV0041GetPartitions) | **Get** /slurm/v0.0.41/partitions/ | get all partition info
[**SlurmV0041GetPing**](SlurmAPI.md#SlurmV0041GetPing) | **Get** /slurm/v0.0.41/ping/ | ping test
[**SlurmV0041GetReconfigure**](SlurmAPI.md#SlurmV0041GetReconfigure) | **Get** /slurm/v0.0.41/reconfigure/ | request slurmctld reconfigure
[**SlurmV0041GetReservation**](SlurmAPI.md#SlurmV0041GetReservation) | **Get** /slurm/v0.0.41/reservation/{reservation_name} | get reservation info
[**SlurmV0041GetReservations**](SlurmAPI.md#SlurmV0041GetReservations) | **Get** /slurm/v0.0.41/reservations/ | get all reservation info
[**SlurmV0041GetShares**](SlurmAPI.md#SlurmV0041GetShares) | **Get** /slurm/v0.0.41/shares | get fairshare info
[**SlurmV0041PostJob**](SlurmAPI.md#SlurmV0041PostJob) | **Post** /slurm/v0.0.41/job/{job_id} | update job
[**SlurmV0041PostJobAllocate**](SlurmAPI.md#SlurmV0041PostJobAllocate) | **Post** /slurm/v0.0.41/job/allocate | submit new job allocation without any steps that must be signaled to stop
[**SlurmV0041PostJobSubmit**](SlurmAPI.md#SlurmV0041PostJobSubmit) | **Post** /slurm/v0.0.41/job/submit | submit new job
[**SlurmV0041PostNode**](SlurmAPI.md#SlurmV0041PostNode) | **Post** /slurm/v0.0.41/node/{node_name} | update node properties



## SlurmV0041DeleteJob

> V0041OpenapiResp SlurmV0041DeleteJob(ctx, jobId).Signal(signal).Flags(flags).Execute()

cancel or signal job

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
	jobId := "jobId_example" // string | Job ID
	signal := "signal_example" // string | Signal to send to Job (optional)
	flags := "flags_example" // string | Signalling flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJob(context.Background(), jobId).Signal(signal).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJob`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **string** | Signal to send to Job | 
 **flags** | **string** | Signalling flags | 

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


## SlurmV0041DeleteJobs

> SlurmV0041DeleteJobs200Response SlurmV0041DeleteJobs(ctx).SlurmV0041DeleteJobsRequest(slurmV0041DeleteJobsRequest).Execute()

send signal to list of jobs

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
	slurmV0041DeleteJobsRequest := *openapiclient.NewSlurmV0041DeleteJobsRequest() // SlurmV0041DeleteJobsRequest | Signal or cancel jobs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteJobs(context.Background()).SlurmV0041DeleteJobsRequest(slurmV0041DeleteJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteJobs`: SlurmV0041DeleteJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041DeleteJobsRequest** | [**SlurmV0041DeleteJobsRequest**](SlurmV0041DeleteJobsRequest.md) | Signal or cancel jobs | 

### Return type

[**SlurmV0041DeleteJobs200Response**](SlurmV0041DeleteJobs200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041DeleteNode

> V0041OpenapiResp SlurmV0041DeleteNode(ctx, nodeName).Execute()

delete node

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
	nodeName := "nodeName_example" // string | Node name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041DeleteNode(context.Background(), nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041DeleteNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041DeleteNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041DeleteNodeRequest struct via the builder pattern


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


## SlurmV0041GetDiag

> SlurmV0041GetDiag200Response SlurmV0041GetDiag(ctx).Execute()

get diagnostics

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetDiag`: SlurmV0041GetDiag200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetDiagRequest struct via the builder pattern


### Return type

[**SlurmV0041GetDiag200Response**](SlurmV0041GetDiag200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJob

> V0041OpenapiJobInfoResp SlurmV0041GetJob(ctx, jobId).UpdateTime(updateTime).Flags(flags).Execute()

get job info

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
	jobId := "jobId_example" // string | Job ID
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJob(context.Background(), jobId).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJob`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJobs

> V0041OpenapiJobInfoResp SlurmV0041GetJobs(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of jobs

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
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJobs(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJobs`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetJobsState

> V0041OpenapiJobInfoResp SlurmV0041GetJobsState(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get list of job states

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
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetJobsState(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetJobsState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetJobsState`: V0041OpenapiJobInfoResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetJobsState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetJobsStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiJobInfoResp**](V0041OpenapiJobInfoResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetLicenses

> SlurmV0041GetLicenses200Response SlurmV0041GetLicenses(ctx).Execute()

get all Slurm tracked license info

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetLicenses`: SlurmV0041GetLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetLicensesRequest struct via the builder pattern


### Return type

[**SlurmV0041GetLicenses200Response**](SlurmV0041GetLicenses200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNode

> V0041OpenapiNodesResp SlurmV0041GetNode(ctx, nodeName).UpdateTime(updateTime).Flags(flags).Execute()

get node info

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
	nodeName := "nodeName_example" // string | Node name
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNode(context.Background(), nodeName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNode`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetNodes

> V0041OpenapiNodesResp SlurmV0041GetNodes(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get node(s) info

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
	updateTime := "updateTime_example" // string | Filter jobs since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetNodes(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetNodes`: V0041OpenapiNodesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter jobs since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiNodesResp**](V0041OpenapiNodesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartition

> V0041OpenapiPartitionResp SlurmV0041GetPartition(ctx, partitionName).UpdateTime(updateTime).Flags(flags).Execute()

get partition info

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
	partitionName := "partitionName_example" // string | Partition name
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartition(context.Background(), partitionName).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartition`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Partition name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPartitions

> V0041OpenapiPartitionResp SlurmV0041GetPartitions(ctx).UpdateTime(updateTime).Flags(flags).Execute()

get all partition info

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPartitions(context.Background()).UpdateTime(updateTime).Flags(flags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPartitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPartitions`: V0041OpenapiPartitionResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 

### Return type

[**V0041OpenapiPartitionResp**](V0041OpenapiPartitionResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetPing

> SlurmV0041GetPing200Response SlurmV0041GetPing(ctx).Execute()

ping test

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetPing`: SlurmV0041GetPing200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetPingRequest struct via the builder pattern


### Return type

[**SlurmV0041GetPing200Response**](SlurmV0041GetPing200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReconfigure

> V0041OpenapiResp SlurmV0041GetReconfigure(ctx).Execute()

request slurmctld reconfigure

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReconfigure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReconfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReconfigure`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReconfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReconfigureRequest struct via the builder pattern


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


## SlurmV0041GetReservation

> V0041OpenapiReservationResp SlurmV0041GetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

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
	reservationName := "reservationName_example" // string | Reservation name
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservation`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Reservation name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetReservations

> V0041OpenapiReservationResp SlurmV0041GetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

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
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetReservations(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetReservations`: V0041OpenapiReservationResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiReservationResp**](V0041OpenapiReservationResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041GetShares

> SlurmV0041GetShares200Response SlurmV0041GetShares(ctx).Accounts(accounts).Users(users).Execute()

get fairshare info

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
	accounts := "accounts_example" // string | Accounts to query (optional)
	users := "users_example" // string | Users to query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041GetShares(context.Background()).Accounts(accounts).Users(users).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041GetShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041GetShares`: SlurmV0041GetShares200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041GetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041GetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **string** | Accounts to query | 
 **users** | **string** | Users to query | 

### Return type

[**SlurmV0041GetShares200Response**](SlurmV0041GetShares200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJob

> SlurmV0041PostJob200Response SlurmV0041PostJob(ctx, jobId).SlurmV0041PostJobSubmitRequestJobsInner(slurmV0041PostJobSubmitRequestJobsInner).Execute()

update job

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
	jobId := "jobId_example" // string | Job ID
	slurmV0041PostJobSubmitRequestJobsInner := *openapiclient.NewSlurmV0041PostJobSubmitRequestJobsInner() // SlurmV0041PostJobSubmitRequestJobsInner | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJob(context.Background(), jobId).SlurmV0041PostJobSubmitRequestJobsInner(slurmV0041PostJobSubmitRequestJobsInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJob`: SlurmV0041PostJob200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slurmV0041PostJobSubmitRequestJobsInner** | [**SlurmV0041PostJobSubmitRequestJobsInner**](SlurmV0041PostJobSubmitRequestJobsInner.md) | Job update description | 

### Return type

[**SlurmV0041PostJob200Response**](SlurmV0041PostJob200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobAllocate

> SlurmV0041PostJobAllocate200Response SlurmV0041PostJobAllocate(ctx).SlurmV0041PostJobAllocateRequest(slurmV0041PostJobAllocateRequest).Execute()

submit new job allocation without any steps that must be signaled to stop

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
	slurmV0041PostJobAllocateRequest := *openapiclient.NewSlurmV0041PostJobAllocateRequest() // SlurmV0041PostJobAllocateRequest | Job allocation description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobAllocate(context.Background()).SlurmV0041PostJobAllocateRequest(slurmV0041PostJobAllocateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobAllocate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobAllocate`: SlurmV0041PostJobAllocate200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobAllocate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobAllocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041PostJobAllocateRequest** | [**SlurmV0041PostJobAllocateRequest**](SlurmV0041PostJobAllocateRequest.md) | Job allocation description | 

### Return type

[**SlurmV0041PostJobAllocate200Response**](SlurmV0041PostJobAllocate200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostJobSubmit

> SlurmV0041PostJobSubmit200Response SlurmV0041PostJobSubmit(ctx).SlurmV0041PostJobSubmitRequest(slurmV0041PostJobSubmitRequest).Execute()

submit new job

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
	slurmV0041PostJobSubmitRequest := *openapiclient.NewSlurmV0041PostJobSubmitRequest() // SlurmV0041PostJobSubmitRequest | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostJobSubmit(context.Background()).SlurmV0041PostJobSubmitRequest(slurmV0041PostJobSubmitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostJobSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostJobSubmit`: SlurmV0041PostJobSubmit200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostJobSubmit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostJobSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmV0041PostJobSubmitRequest** | [**SlurmV0041PostJobSubmitRequest**](SlurmV0041PostJobSubmitRequest.md) | Job description | 

### Return type

[**SlurmV0041PostJobSubmit200Response**](SlurmV0041PostJobSubmit200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmV0041PostNode

> V0041OpenapiResp SlurmV0041PostNode(ctx, nodeName).SlurmV0041PostNodeRequest(slurmV0041PostNodeRequest).Execute()

update node properties

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
	nodeName := "nodeName_example" // string | Node name
	slurmV0041PostNodeRequest := *openapiclient.NewSlurmV0041PostNodeRequest() // SlurmV0041PostNodeRequest | Node update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmAPI.SlurmV0041PostNode(context.Background(), nodeName).SlurmV0041PostNodeRequest(slurmV0041PostNodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmV0041PostNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmV0041PostNode`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmV0041PostNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Node name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmV0041PostNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slurmV0041PostNodeRequest** | [**SlurmV0041PostNodeRequest**](SlurmV0041PostNodeRequest.md) | Node update description | 

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

