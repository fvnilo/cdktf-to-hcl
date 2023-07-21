terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

module "s3_bucket" {
  source = "./module"

  region = "ca-central-1"
  bucket_name = "fanilo-buckt-cool"
}

output "bucket_name" {
  value = module.s3_bucket.name
}