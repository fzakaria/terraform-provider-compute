Terraform Provider for Oracle Compute Cloud
===========================================

PLEASE NOTE: This repository location has changed from github.com/oracle/terraform to [github.com/oracle/terraform-provider-compute](https://github.com/oracle/terraform-provider-compute/)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.7.x
-	[Oracle Compute Cloud](https://cloud.oracle.com/compute) Account
-	[Go](https://golang.org/doc/install) 1.7 (to build the provider plugin)

Building
--------

Create a directory where the provider will be built, and set the Go language `GOPATH`

```sh
$ export GOPATH=/home/opc/terraform-provider
$ cd $GOPATH
```

Fetch the source and build the provider.

```sh
$ go get -d github.com/oracle/terraform-provider-compute/provider
$ go build -o terraform-provider-opc github.com/oracle/terraform-provider-compute/provider
```

Usage
-----

Add the generated `terraform-provider-opc` executable to your `.terraformrc` configuration (`%APPDATA%/terraform.rc` on Windows), e.g.

```
providers {
    opc = "/home/opc/terraform-provider/terraform-provider-opc"
}
```

To authenticate with the Oracle Compute Cloud the provider will prompt for the required environment credentials. These credentails can be set in the following environment variables:

-	`OPC_ENDPOINT` - Endpoint provided by Oracle Public Cloud (e.g. https://api-z13.compute.em2.oraclecloud.com/\)
-	`OPC_USERNAME` - Username for Oracle Public Cloud
-	`OPC_PASSWORD` - Password for Oracle Public Cloud
-	`OPC_IDENTITY_DOMAIN` - Identity domain for Oracle Public Cloud

### Example Terraform configuration

An example [`test.tf`](test/test.tf) is provided that demonstatates the basic usage of the Oracle Compute Cloud Terraform Provider.

```sh
$ cd $GOPATH/src/github.com/oracle/terraform-provider-compute/test
$ terraform plan
$ terraform apply
$ terraform destroy
```

Running the Integration Tests
-----------------------------

An Oracle Compute Cloud Account is required to run the integration tests. The `OCP_*` variables must have been exported

```sh
$ cd $GOPATH/src/github.com/oracle/terraform-provider-compute/sdk/compute
$ go test
```
