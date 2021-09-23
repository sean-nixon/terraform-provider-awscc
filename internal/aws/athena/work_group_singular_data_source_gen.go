// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package athena

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_athena_work_group", workGroupDataSourceType)
}

// workGroupDataSourceType returns the Terraform awscc_athena_work_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Athena::WorkGroup resource type.
func workGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time the workgroup was created.",
			//   "type": "string"
			// }
			Description: "The date and time the workgroup was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The workgroup description.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The workgroup description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The workGroup name.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The workGroup name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"recursive_delete_option": {
			// Property: RecursiveDeleteOption
			// CloudFormation resource type schema:
			// {
			//   "description": "The option to delete the workgroup and its contents even if the workgroup contains any named queries.",
			//   "type": "boolean"
			// }
			Description: "The option to delete the workgroup and its contents even if the workgroup contains any named queries.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the workgroup: ENABLED or DISABLED.",
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The state of the workgroup: ENABLED or DISABLED.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"work_group_configuration": {
			// Property: WorkGroupConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BytesScannedCutoffPerQuery": {
			//       "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
			//       "format": "int64",
			//       "minimum": 10000000,
			//       "type": "integer"
			//     },
			//     "EnforceWorkGroupConfiguration": {
			//       "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
			//       "type": "boolean"
			//     },
			//     "EngineVersion": {
			//       "additionalProperties": false,
			//       "description": "The Athena engine version for running queries.",
			//       "properties": {
			//         "EffectiveEngineVersion": {
			//           "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
			//           "type": "string"
			//         },
			//         "SelectedEngineVersion": {
			//           "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "PublishCloudWatchMetricsEnabled": {
			//       "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
			//       "type": "boolean"
			//     },
			//     "RequesterPaysEnabled": {
			//       "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
			//       "type": "boolean"
			//     },
			//     "ResultConfiguration": {
			//       "additionalProperties": false,
			//       "description": "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
			//       "properties": {
			//         "EncryptionConfiguration": {
			//           "additionalProperties": false,
			//           "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
			//           "properties": {
			//             "EncryptionOption": {
			//               "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
			//               "enum": [
			//                 "SSE_S3",
			//                 "SSE_KMS",
			//                 "CSE_KMS"
			//               ],
			//               "type": "string"
			//             },
			//             "KmsKey": {
			//               "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "EncryptionOption"
			//           ],
			//           "type": "object"
			//         },
			//         "OutputLocation": {
			//           "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bytes_scanned_cutoff_per_query": {
						// Property: BytesScannedCutoffPerQuery
						Description: "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						Type:        types.NumberType,
						Computed:    true,
					},
					"enforce_work_group_configuration": {
						// Property: EnforceWorkGroupConfiguration
						Description: "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						Type:        types.BoolType,
						Computed:    true,
					},
					"engine_version": {
						// Property: EngineVersion
						Description: "The Athena engine version for running queries.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"effective_engine_version": {
									// Property: EffectiveEngineVersion
									Description: "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									Type:        types.StringType,
									Computed:    true,
								},
								"selected_engine_version": {
									// Property: SelectedEngineVersion
									Description: "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"publish_cloudwatch_metrics_enabled": {
						// Property: PublishCloudWatchMetricsEnabled
						Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"requester_pays_enabled": {
						// Property: RequesterPaysEnabled
						Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						Type:        types.BoolType,
						Computed:    true,
					},
					"result_configuration": {
						// Property: ResultConfiguration
						Description: "The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results. These are known as \"client-side settings\". If workgroup settings override client-side settings, then the query uses the workgroup settings.\n",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"encryption_configuration": {
									// Property: EncryptionConfiguration
									Description: "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"encryption_option": {
												// Property: EncryptionOption
												Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												Type:        types.StringType,
												Computed:    true,
											},
											"kms_key": {
												// Property: KmsKey
												Description: "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"output_location": {
									// Property: OutputLocation
									Description: "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"work_group_configuration_updates": {
			// Property: WorkGroupConfigurationUpdates
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. ",
			//   "properties": {
			//     "BytesScannedCutoffPerQuery": {
			//       "description": "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
			//       "format": "int64",
			//       "minimum": 10000000,
			//       "type": "integer"
			//     },
			//     "EnforceWorkGroupConfiguration": {
			//       "description": "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
			//       "type": "boolean"
			//     },
			//     "EngineVersion": {
			//       "additionalProperties": false,
			//       "description": "The Athena engine version for running queries.",
			//       "properties": {
			//         "EffectiveEngineVersion": {
			//           "description": "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
			//           "type": "string"
			//         },
			//         "SelectedEngineVersion": {
			//           "description": "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "PublishCloudWatchMetricsEnabled": {
			//       "description": "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
			//       "type": "boolean"
			//     },
			//     "RemoveBytesScannedCutoffPerQuery": {
			//       "description": "Indicates that the data usage control limit per query is removed.",
			//       "type": "boolean"
			//     },
			//     "RequesterPaysEnabled": {
			//       "description": "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
			//       "type": "boolean"
			//     },
			//     "ResultConfigurationUpdates": {
			//       "additionalProperties": false,
			//       "description": "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
			//       "properties": {
			//         "EncryptionConfiguration": {
			//           "additionalProperties": false,
			//           "description": "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
			//           "properties": {
			//             "EncryptionOption": {
			//               "description": "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
			//               "enum": [
			//                 "SSE_S3",
			//                 "SSE_KMS",
			//                 "CSE_KMS"
			//               ],
			//               "type": "string"
			//             },
			//             "KmsKey": {
			//               "description": "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "EncryptionOption"
			//           ],
			//           "type": "object"
			//         },
			//         "OutputLocation": {
			//           "description": "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
			//           "type": "string"
			//         },
			//         "RemoveEncryptionConfiguration": {
			//           "type": "boolean"
			//         },
			//         "RemoveOutputLocation": {
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. ",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bytes_scanned_cutoff_per_query": {
						// Property: BytesScannedCutoffPerQuery
						Description: "The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.",
						Type:        types.NumberType,
						Computed:    true,
					},
					"enforce_work_group_configuration": {
						// Property: EnforceWorkGroupConfiguration
						Description: "If set to \"true\", the settings for the workgroup override client-side settings. If set to \"false\", client-side settings are used",
						Type:        types.BoolType,
						Computed:    true,
					},
					"engine_version": {
						// Property: EngineVersion
						Description: "The Athena engine version for running queries.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"effective_engine_version": {
									// Property: EffectiveEngineVersion
									Description: "Read only. The engine version on which the query runs. If the user requests a valid engine version other than Auto, the effective engine version is the same as the engine version that the user requested. If the user requests Auto, the effective engine version is chosen by Athena. When a request to update the engine version is made by a CreateWorkGroup or UpdateWorkGroup operation, the EffectiveEngineVersion field is ignored.",
									Type:        types.StringType,
									Computed:    true,
								},
								"selected_engine_version": {
									// Property: SelectedEngineVersion
									Description: "The engine version requested by the user. Possible values are determined by the output of ListEngineVersions, including Auto. The default is Auto.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"publish_cloudwatch_metrics_enabled": {
						// Property: PublishCloudWatchMetricsEnabled
						Description: "Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"remove_bytes_scanned_cutoff_per_query": {
						// Property: RemoveBytesScannedCutoffPerQuery
						Description: "Indicates that the data usage control limit per query is removed.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"requester_pays_enabled": {
						// Property: RequesterPaysEnabled
						Description: "If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries. If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. ",
						Type:        types.BoolType,
						Computed:    true,
					},
					"result_configuration_updates": {
						// Property: ResultConfigurationUpdates
						Description: "The result configuration information about the queries in this workgroup that will be updated. Includes the updated results location and an updated option for encrypting query results. ",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"encryption_configuration": {
									// Property: EncryptionConfiguration
									Description: "If query results are encrypted in Amazon S3, indicates the encryption option used (for example, SSE-KMS or CSE-KMS) and key information.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"encryption_option": {
												// Property: EncryptionOption
												Description: "Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.",
												Type:        types.StringType,
												Computed:    true,
											},
											"kms_key": {
												// Property: KmsKey
												Description: "For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID. ",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"output_location": {
									// Property: OutputLocation
									Description: "The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/. To run the query, you must specify the query results location using one of the ways: either for individual queries using either this setting (client-side), or in the workgroup, using WorkGroupConfiguration",
									Type:        types.StringType,
									Computed:    true,
								},
								"remove_encryption_configuration": {
									// Property: RemoveEncryptionConfiguration
									Type:     types.BoolType,
									Computed: true,
								},
								"remove_output_location": {
									// Property: RemoveOutputLocation
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
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
		Description: "Data Source schema for AWS::Athena::WorkGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::WorkGroup").WithTerraformTypeName("awscc_athena_work_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bytes_scanned_cutoff_per_query":        "BytesScannedCutoffPerQuery",
		"creation_time":                         "CreationTime",
		"description":                           "Description",
		"effective_engine_version":              "EffectiveEngineVersion",
		"encryption_configuration":              "EncryptionConfiguration",
		"encryption_option":                     "EncryptionOption",
		"enforce_work_group_configuration":      "EnforceWorkGroupConfiguration",
		"engine_version":                        "EngineVersion",
		"key":                                   "Key",
		"kms_key":                               "KmsKey",
		"name":                                  "Name",
		"output_location":                       "OutputLocation",
		"publish_cloudwatch_metrics_enabled":    "PublishCloudWatchMetricsEnabled",
		"recursive_delete_option":               "RecursiveDeleteOption",
		"remove_bytes_scanned_cutoff_per_query": "RemoveBytesScannedCutoffPerQuery",
		"remove_encryption_configuration":       "RemoveEncryptionConfiguration",
		"remove_output_location":                "RemoveOutputLocation",
		"requester_pays_enabled":                "RequesterPaysEnabled",
		"result_configuration":                  "ResultConfiguration",
		"result_configuration_updates":          "ResultConfigurationUpdates",
		"selected_engine_version":               "SelectedEngineVersion",
		"state":                                 "State",
		"tags":                                  "Tags",
		"value":                                 "Value",
		"work_group_configuration":              "WorkGroupConfiguration",
		"work_group_configuration_updates":      "WorkGroupConfigurationUpdates",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
