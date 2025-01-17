// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_enclave_certificate_iam_role_association", enclaveCertificateIamRoleAssociationDataSource)
}

// enclaveCertificateIamRoleAssociationDataSource returns the Terraform awscc_ec2_enclave_certificate_iam_role_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::EnclaveCertificateIamRoleAssociation resource.
func enclaveCertificateIamRoleAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
		//	  "maxLength": 1283,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$",
		//	  "type": "string"
		//	}
		"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateS3BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon S3 bucket to which the certificate was uploaded.",
		//	  "type": "string"
		//	}
		"certificate_s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon S3 bucket to which the certificate was uploaded.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateS3ObjectKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
		//	  "type": "string"
		//	}
		"certificate_s3_object_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
		//	  "type": "string"
		//	}
		"encryption_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
		//	  "maxLength": 1283,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::EnclaveCertificateIamRoleAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EnclaveCertificateIamRoleAssociation").WithTerraformTypeName("awscc_ec2_enclave_certificate_iam_role_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_arn":            "CertificateArn",
		"certificate_s3_bucket_name": "CertificateS3BucketName",
		"certificate_s3_object_key":  "CertificateS3ObjectKey",
		"encryption_kms_key_id":      "EncryptionKmsKeyId",
		"role_arn":                   "RoleArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
