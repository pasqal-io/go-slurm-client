# Contributing

This document will mostly focus on how to move this repository forward

## Path in repo

For each version `v0.0.XX` of the slurm api, always work in a subdirectory `vXX`, so that it can be used correctly from a client's perspective.

For the rest of this document, we assume that all the command will be run in this `vXX` sub-directory - in our examples `v40` for the version `v0.0.40`.

## Extracting a new version of the API

Any version above v0.0.40 can be generated using a Slurm version that supports it, and running the following command:

```sh
env SLURM_CONF=/dev/null slurmrestd --generate-openapi-spec -s slurmctld,slurmdbd -d v0.0.40 > openapi.json
```

(Replace `v0.0.40` with any version you might be interested in)

### Generate a docker container using any version of Slurm

You can use the repository [giovtorres/slurm-docker-cluster](https://github.com/giovtorres/slurm-docker-cluster) to build a docker image of a specific Slurm version.

## Generate the client sdk code for a version

Using [openapi-generator-cli](https://openapi-generator.tech/docs/installation/) (you can also use the devcontainer):

```sh
openapi-generator-cli generate -i openapi.json -g go 
```

Proceed to replace all occurences of `github.com/GIT_USER_ID/GIT_REPO_ID` with `github.com/pasqal-io/go-slurm-client/v40`, replacing `v40` with the version you are currently generating.
