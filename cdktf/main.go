package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v16/provider"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/s3bucket"
	"github.com/cdktf/cdktf-tf-module-stack-go/tfmodulestack/v2"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewHclInteropStack(scope constructs.Construct, name string) cdktf.TerraformStack {
	stack := tfmodulestack.NewTFModuleStack(scope, &name)

	region := cdktf.NewTerraformVariable(stack, jsii.String("region"), &cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Default:     jsii.String("ca-central-1"),
		Description: jsii.String("The region to use"),
	})

	bucketName := cdktf.NewTerraformVariable(stack, jsii.String("bucket_name"), &cdktf.TerraformVariableConfig{
		Type:        jsii.String("string"),
		Default:     jsii.String(""),
		Description: jsii.String("The name of the bucket to create"),
	})

	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region: region.StringValue(),
	})

	bucket := s3bucket.NewS3Bucket(stack, jsii.String("bucket"), &s3bucket.S3BucketConfig{
		Bucket: bucketName.StringValue(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("name"), &cdktf.TerraformOutputConfig{
		Value: bucket.Bucket(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	NewHclInteropStack(app, "s3_interop")
	app.Synth()
}
