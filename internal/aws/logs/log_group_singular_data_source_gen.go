// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_log_group", logGroupDataSource)
}

// logGroupDataSource returns the Terraform awscc_logs_log_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::LogGroup resource.
func logGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The CloudWatch log group ARN.",
			//	  "type": "string"
			//	}
			Description: "The CloudWatch log group ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_protection_policy": {
			// Property: DataProtectionPolicy
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per topic.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
			//	  "type": "object"
			//	}
			Description: "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per topic.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			//	  "maxLength": 256,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			//	  "maxLength": 512,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"retention_in_days": {
			// Property: RetentionInDays
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			//	  "enum": [
			//	    1,
			//	    3,
			//	    5,
			//	    7,
			//	    14,
			//	    30,
			//	    60,
			//	    90,
			//	    120,
			//	    150,
			//	    180,
			//	    365,
			//	    400,
			//	    545,
			//	    731,
			//	    1827,
			//	    2192,
			//	    2557,
			//	    2922,
			//	    3288,
			//	    3653
			//	  ],
			//	  "type": "integer"
			//	}
			Description: "The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, and 3653.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this resource.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, - and @.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Logs::LogGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogGroup").WithTerraformTypeName("awscc_logs_log_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"data_protection_policy": "DataProtectionPolicy",
		"key":                    "Key",
		"kms_key_id":             "KmsKeyId",
		"log_group_name":         "LogGroupName",
		"retention_in_days":      "RetentionInDays",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
