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
To start using the Identity Intelligence Services on your own resources, use the following steps

## 1. Create Cloud Resources
Using the infrastructure defined via HCL code in the [deployments/infra](deployments/infra/) folder, run this on an AWS Account that you own so the resources are properly configured

## 2. Configure Directories
Once the resources are created in your AWS Account, there will be a DynamoDB table named identity_intelligence_*_directories (ex:identity_intelligence_prd_directories).  For each directory that you wish to onboard, add an item to that table with the following attributes:

|Name|Description|Example Value|
|---|---|---|
|Id|Unique identifier for the directory|corp.mycompany.com|
|Name|Friendly name of the directory| MyCompany Corp|
|Host|Host where the directory can be found at|corp.mycompany.com:636|
|Base|(optional)Starting point for the directory searches|OU=mycompany,OU=com|
|Type|Type of directory|LDAP|
|Authentication Type|Authentication method to use when connecting to the host|ClientId|
|ClientIdConfigName|Path to the AWS Secret Manager secret that contains the Client Id|prd/identity-intelligence/client-id|
|ClientSecretConfigName|Path to the AWS Secret Manager secret that contains the Client Secret|prd/identity-intelligence-client-secret|

## 3. Populate Client Id and Client Secret
For every directory where there is a Cliend Id and Client Secret, create the secrets using the AWS Secret Manager and populate with the proper value.
###Example

|Secret Name|Secret Value|
|---|---|
|prd/identity-intelligence/client-id|the-user|
|prd/identity-intelligence-client-secret|superSecretValue!|

## 4. Configure Local AWS Credentials
Once the AWS resources are configured, credentials on the machine that will be running the consoel app need to be configured.  In your machine's environment variables, ensure there are valid AWS credentials for the Account where the resources created in step 1 are located.  The user these credentials are for need publish permissions on the Signal Ingestion SQS queue, Read Access on the Directories DynamoDB, and Read Access to the Secrets Manager secrets added in step 3.

### Example
```
export aws_region=us-west-2
export AWS_ACCESS_KEY_ID=AKIAGDBMSF5KR367TZX
export AWS_SECRET_ACCESS_KEY=somesupersecretvalue
```

## 5. Execute Command Line Application
When the AWS resources are setup and credentials set locally, the command line app can be executed.  This app reads data from one or more directories and pushes to the signal ingestion queue where they are then saved in S3 for later analysis.  

In the cmd/cli directory, below are examples of commands to run to read directory data.  Arguments are:
|Name|Required|Purpose|Example Value|
|directory|No|Specifies the directory to read from.  If not specified, will read all directories|corp.mycompany.com|
|object|No|What type of object to read.  Valid values are _user_ and _group_|user|

### Examples
#### Read User Data For All Directories
```
cd cmd/cli
go run main.go -object user
```

#### Read Group Data For A Single Directory
```
cd cmd/cli
go run main.go -object user -directory corp.mycompany.com
```

## 6. Tableau Analytics
Once the desired directories to be analyzed are read from, it is now time to gain insights from the data.  Using queries in the [analytics/queries](../analytics/queries/) directory, create data extracts (either text or hyperfile format) for each one.  Once data is extracted, upload or open the sample workbooks in the [analytics/tableau/workbooks](../analytics/tableau/workbooks/) directory and replace each source with the matching query in your installation.

# Development Environment Setup