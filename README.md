# Go slurm client SDK

## Supported version:

This sdk currently supports the following version of the Slurm API:
- v0.0.37
- v0.0.39
- v0.0.41
- v0.0.44

## How to install

If you want to install the sdk for version `v0.0.41` of the Slurm API, you need to use the `v41` path, and similarly for other versions:

```sh
go get github.com/pasqal-io/go-slurm-client/v41
```

Then, use as such:

```golang

import (
	slurmclient "github.com/pasqal-io/go-slurm-client/v41"
)

func main() {
    cfg := slurmclient.NewConfiguration()
    // ...

    myClient := slurmclient.NewAPIClient(cfg)
    response, httpResponse, err := myClient.SlurmAPI.SlurmctldSubmitJob(
		ctx,
	).V0041JobSubmission(*preparedV0041JobSubmission).Execute()
}
```
