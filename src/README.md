# Identity Intelligence Services Components
The components in the Identity Intelligence Services enable the collection and storage of large amounts of identity data so it is available for further analysis.  Components are primarily written in Golang and use AWS resources for processing, data storage, and analytical capabilities.

# Requirements
* golang 1.18 or higher

# Components

## Applications
|Name|Location|Purpose|
|---|---|---|
|Command Line Tool|cmd/cli|Command line tool that reads directory data and pushes to the AWS resources for processing|
|Directory Object Receiver|cmd/lambda-directoryobjectreceiver|AWS Lambda Function that receies messages pushed from the command line tool and saves them to the proper S3 bucket|

## Packages
|Name|Location|Purpose|
|---|---|---|
|configuration|internal/pkg/configuration|Reads configuration and secret data|
|core|internal/pkg/core|Extensions to the native types in golang|
|messaging|internal/pkg/messaging|Implementations of different messaging platforms for service to service communication|
|models|internal/pkg/models|Data models / structs used in the pkg packages|
|orchestration|internal/pkg/orchestration|The main entry point for all exposed functionality.  Contains business logic, processing, and flow implementations|
|repositories|internal/pkg/repositories|Data repositories used to read and update application data|
|services|internal/pkg/services|Implementations around interacting with different directory providers and other external APIs / services|

# Infrastructure
The infrastructure used to implement this solution is defined with Terraform HCL.  See the [deployments/infra](deployments/infra/) directory for more details.

A sample deployment pipelien for this can be seen at [terraform-deploy-prd](../.github/workflows/infrastructure-deploy-prd.yml)

# Pipelines
The following pipelines are used by this repository to build, test, and deploy a sample instance.  They are:
|Name|Location|Purpose|
|---|---|---|
|application-continuous-integration|[.github/workflows/application-ci.yml](../.github/workflows/application-ci.yml)|Application code (golang) CI build|
|terraform-continuous-integration|[.github/workflows/infrastrcture-tf-ci.yml](../.github/workflows//infrastructure-tf-ci.yml)|Infrastructure code (terraform) CI Build|
|terraform-deploy-prd|[.github/workflows/infrastructure/deploy-prod.yml](../.github/workflows/infrastructure-deploy-prd.yml)|Sample deployment pipeline to an AWS Account|

# How To Use

# Development Environment Setup