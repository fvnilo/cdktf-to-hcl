#cdktf-to-hcl
> A simple test of interop between CDKTF and HCL

# Getting Started

## Prerequisites

- [terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
- [cdktf](https://developer.hashicorp.com/terraform/tutorials/cdktf/cdktf-install)
- [go](https://go.dev/doc/install)

## Generate Terraform Module From CDKTF Code
```
cd cdktf
go mod tidy
cdkth synth
mkdir -p ../hcl/module
cp cdktf.out/stacks/s3_interop/cdk.tf.json ../hcl/module
```

## Run HCL Config Using CDKTF Module
```
cd hcl
terraform init
terraform plan
terraform apply
```