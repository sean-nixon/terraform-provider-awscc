// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_eks_nodegroup", nodegroupResourceType)
}

// nodegroupResourceType returns the Terraform awscc_eks_nodegroup resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EKS::Nodegroup resource type.
func nodegroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"ami_type": {
			// Property: AmiType
			// CloudFormation resource type schema:
			// {
			//   "description": "The AMI type for your node group.",
			//   "type": "string"
			// }
			Description: "The AMI type for your node group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"capacity_type": {
			// Property: CapacityType
			// CloudFormation resource type schema:
			// {
			//   "description": "The capacity type of your managed node group.",
			//   "type": "string"
			// }
			Description: "The capacity type of your managed node group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the cluster to create the node group in.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the cluster to create the node group in.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"disk_size": {
			// Property: DiskSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The root device disk size (in GiB) for your node group instances.",
			//   "type": "integer"
			// }
			Description: "The root device disk size (in GiB) for your node group instances.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"force_update_enabled": {
			// Property: ForceUpdateEnabled
			// CloudFormation resource type schema:
			// {
			//   "default": false,
			//   "description": "Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.",
			//   "type": "boolean"
			// }
			Description: "Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Bool{Value: false}),
				tfsdk.UseStateForUnknown(),
			},
			// ForceUpdateEnabled is a write-only property.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"instance_types": {
			// Property: InstanceTypes
			// CloudFormation resource type schema:
			// {
			//   "description": "Specify the instance types for a node group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Specify the instance types for a node group.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"labels": {
			// Property: Labels
			// CloudFormation resource type schema:
			// {
			//   "description": "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
			//   "type": "object"
			// }
			Description: "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"launch_template": {
			// Property: LaunchTemplate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing a node group's launch template specification.",
			//   "properties": {
			//     "Id": {
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Name": {
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Version": {
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object representing a node group's launch template specification.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
					"version": {
						// Property: Version
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
				},
			),
			Optional: true,
		},
		"node_role": {
			// Property: NodeRole
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"nodegroup_name": {
			// Property: NodegroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name to give your node group.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The unique name to give your node group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"release_version": {
			// Property: ReleaseVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The AMI version of the Amazon EKS-optimized AMI to use with your node group.",
			//   "type": "string"
			// }
			Description: "The AMI version of the Amazon EKS-optimized AMI to use with your node group.",
			Type:        types.StringType,
			Optional:    true,
		},
		"remote_access": {
			// Property: RemoteAccess
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The remote access (SSH) configuration to use with your node group.",
			//   "properties": {
			//     "Ec2SshKey": {
			//       "type": "string"
			//     },
			//     "SourceSecurityGroups": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     }
			//   },
			//   "required": [
			//     "Ec2SshKey"
			//   ],
			//   "type": "object"
			// }
			Description: "The remote access (SSH) configuration to use with your node group.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"ec_2_ssh_key": {
						// Property: Ec2SshKey
						Type:     types.StringType,
						Required: true,
					},
					"source_security_groups": {
						// Property: SourceSecurityGroups
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"scaling_config": {
			// Property: ScalingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The scaling configuration details for the Auto Scaling group that is created for your node group.",
			//   "properties": {
			//     "DesiredSize": {
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "MaxSize": {
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "MinSize": {
			//       "minimum": 0,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The scaling configuration details for the Auto Scaling group that is created for your node group.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"desired_size": {
						// Property: DesiredSize
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
					"max_size": {
						// Property: MaxSize
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(1),
						},
					},
					"min_size": {
						// Property: MinSize
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(0),
						},
					},
				},
			),
			Optional: true,
		},
		"subnets": {
			// Property: Subnets
			// CloudFormation resource type schema:
			// {
			//   "description": "The subnets to use for the Auto Scaling group that is created for your node group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The subnets to use for the Auto Scaling group that is created for your node group.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.",
			//   "type": "object"
			// }
			Description: "The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"taints": {
			// Property: Taints
			// CloudFormation resource type schema:
			// {
			//   "description": "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An object representing a Taint specification for AWS EKS Nodegroup.",
			//     "properties": {
			//       "Effect": {
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Key": {
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"effect": {
						// Property: Effect
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(0),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"update_config": {
			// Property: UpdateConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The node group update configuration.",
			//   "properties": {
			//     "MaxUnavailable": {
			//       "description": "The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. ",
			//       "minimum": 1,
			//       "type": "number"
			//     },
			//     "MaxUnavailablePercentage": {
			//       "description": "The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.",
			//       "maximum": 100,
			//       "minimum": 1,
			//       "type": "number"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The node group update configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"max_unavailable": {
						// Property: MaxUnavailable
						Description: "The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. ",
						Type:        types.Float64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatAtLeast(1.000000),
						},
					},
					"max_unavailable_percentage": {
						// Property: MaxUnavailablePercentage
						Description: "The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.",
						Type:        types.Float64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatBetween(1.000000, 100.000000),
						},
					},
				},
			),
			Optional: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The Kubernetes version to use for your managed nodes.",
			//   "type": "string"
			// }
			Description: "The Kubernetes version to use for your managed nodes.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::EKS::Nodegroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Nodegroup").WithTerraformTypeName("awscc_eks_nodegroup")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ami_type":                   "AmiType",
		"arn":                        "Arn",
		"capacity_type":              "CapacityType",
		"cluster_name":               "ClusterName",
		"desired_size":               "DesiredSize",
		"disk_size":                  "DiskSize",
		"ec_2_ssh_key":               "Ec2SshKey",
		"effect":                     "Effect",
		"force_update_enabled":       "ForceUpdateEnabled",
		"id":                         "Id",
		"instance_types":             "InstanceTypes",
		"key":                        "Key",
		"labels":                     "Labels",
		"launch_template":            "LaunchTemplate",
		"max_size":                   "MaxSize",
		"max_unavailable":            "MaxUnavailable",
		"max_unavailable_percentage": "MaxUnavailablePercentage",
		"min_size":                   "MinSize",
		"name":                       "Name",
		"node_role":                  "NodeRole",
		"nodegroup_name":             "NodegroupName",
		"release_version":            "ReleaseVersion",
		"remote_access":              "RemoteAccess",
		"scaling_config":             "ScalingConfig",
		"source_security_groups":     "SourceSecurityGroups",
		"subnets":                    "Subnets",
		"tags":                       "Tags",
		"taints":                     "Taints",
		"update_config":              "UpdateConfig",
		"value":                      "Value",
		"version":                    "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ForceUpdateEnabled",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
