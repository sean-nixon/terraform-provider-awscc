// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_imagebuilder_distribution_configuration", distributionConfigurationResource)
}

// distributionConfigurationResource returns the Terraform awscc_imagebuilder_distribution_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::ImageBuilder::DistributionConfiguration resource.
func distributionConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the distribution configuration.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the distribution configuration.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the distribution configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Distributions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The distributions of the distribution configuration.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The distributions of the distribution configuration.",
		//	    "properties": {
		//	      "AmiDistributionConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "The specific AMI settings (for example, launch permissions, AMI tags).",
		//	        "properties": {
		//	          "AmiTags": {
		//	            "additionalProperties": false,
		//	            "description": "The tags to apply to AMIs distributed to this Region.",
		//	            "patternProperties": {
		//	              "": {
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Description": {
		//	            "description": "The description of the AMI distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "KmsKeyId": {
		//	            "description": "The KMS key identifier used to encrypt the distributed image.",
		//	            "type": "string"
		//	          },
		//	          "LaunchPermissionConfiguration": {
		//	            "additionalProperties": false,
		//	            "description": "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
		//	            "properties": {
		//	              "OrganizationArns": {
		//	                "description": "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "OrganizationalUnitArns": {
		//	                "description": "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "UserGroups": {
		//	                "description": "The name of the group.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "UserIds": {
		//	                "description": "The AWS account ID.",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the AMI distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "TargetAccountIds": {
		//	            "description": "The ID of accounts to which you want to distribute an image.",
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ContainerDistributionConfiguration": {
		//	        "additionalProperties": false,
		//	        "description": "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
		//	        "properties": {
		//	          "ContainerTags": {
		//	            "description": "Tags that are attached to the container distribution configuration.",
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "Description": {
		//	            "description": "The description of the container distribution configuration.",
		//	            "type": "string"
		//	          },
		//	          "TargetRepository": {
		//	            "additionalProperties": false,
		//	            "description": "The destination repository for the container distribution configuration.",
		//	            "properties": {
		//	              "RepositoryName": {
		//	                "description": "The repository name of target container repository.",
		//	                "type": "string"
		//	              },
		//	              "Service": {
		//	                "description": "The service of target container repository.",
		//	                "enum": [
		//	                  "ECR"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "FastLaunchConfigurations": {
		//	        "description": "The Windows faster-launching configurations to use for AMI distribution.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The Windows faster-launching configuration to use for AMI distribution.",
		//	          "properties": {
		//	            "AccountId": {
		//	              "description": "The owner account ID for the fast-launch enabled Windows AMI.",
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "description": "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
		//	              "type": "boolean"
		//	            },
		//	            "LaunchTemplate": {
		//	              "additionalProperties": false,
		//	              "description": "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
		//	              "properties": {
		//	                "LaunchTemplateId": {
		//	                  "description": "The ID of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                },
		//	                "LaunchTemplateName": {
		//	                  "description": "The name of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                },
		//	                "LaunchTemplateVersion": {
		//	                  "description": "The version of the launch template to use for faster launching for a Windows AMI.",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "MaxParallelLaunches": {
		//	              "description": "The maximum number of parallel instances that are launched for creating resources.",
		//	              "type": "integer"
		//	            },
		//	            "SnapshotConfiguration": {
		//	              "additionalProperties": false,
		//	              "description": "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
		//	              "properties": {
		//	                "TargetResourceCount": {
		//	                  "description": "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "LaunchTemplateConfigurations": {
		//	        "description": "A group of launchTemplateConfiguration settings that apply to image distribution.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "launchTemplateConfiguration settings that apply to image distribution.",
		//	          "properties": {
		//	            "AccountId": {
		//	              "description": "The account ID that this configuration applies to.",
		//	              "type": "string"
		//	            },
		//	            "LaunchTemplateId": {
		//	              "description": "Identifies the EC2 launch template to use.",
		//	              "type": "string"
		//	            },
		//	            "SetDefaultVersion": {
		//	              "description": "Set the specified EC2 launch template as the default launch template for the specified account.",
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "LicenseConfigurationArns": {
		//	        "description": "The License Manager Configuration to associate with the AMI in the specified Region.",
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "description": "The Amazon Resource Name (ARN) of the License Manager configuration.",
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "Region": {
		//	        "description": "region",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Region"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"distributions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AmiDistributionConfiguration
					"ami_distribution_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AmiTags
							"ami_tags":          // Pattern: ""
							schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The tags to apply to AMIs distributed to this Region.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
									mapplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description of the AMI distribution configuration.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: KmsKeyId
							"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The KMS key identifier used to encrypt the distributed image.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: LaunchPermissionConfiguration
							"launch_permission_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: OrganizationArns
									"organization_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The ARN for an Amazon Web Services Organization that you want to share your AMI with.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											Multiset(),
											listplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: OrganizationalUnitArns
									"organizational_unit_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The ARN for an Organizations organizational unit (OU) that you want to share your AMI with.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											Multiset(),
											listplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: UserGroups
									"user_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The name of the group.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											Multiset(),
											listplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: UserIds
									"user_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "The AWS account ID.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											Multiset(),
											listplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Launch permissions can be used to configure which AWS accounts can use the AMI to launch instances.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the AMI distribution configuration.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TargetAccountIds
							"target_account_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "The ID of accounts to which you want to distribute an image.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The specific AMI settings (for example, launch permissions, AMI tags).",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ContainerDistributionConfiguration
					"container_distribution_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ContainerTags
							"container_tags": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "Tags that are attached to the container distribution configuration.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description of the container distribution configuration.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TargetRepository
							"target_repository": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: RepositoryName
									"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The repository name of target container repository.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Service
									"service": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The service of target container repository.",
										Optional:    true,
										Computed:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"ECR",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The destination repository for the container distribution configuration.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Container distribution settings for encryption, licensing, and sharing in a specific Region.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: FastLaunchConfigurations
					"fast_launch_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccountId
								"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The owner account ID for the fast-launch enabled Windows AMI.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "A Boolean that represents the current state of faster launching for the Windows AMI. Set to true to start using Windows faster launching, or false to stop using it.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: LaunchTemplate
								"launch_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: LaunchTemplateId
										"launch_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The ID of the launch template to use for faster launching for a Windows AMI.",
											Optional:    true,
											Computed:    true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: LaunchTemplateName
										"launch_template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name of the launch template to use for faster launching for a Windows AMI.",
											Optional:    true,
											Computed:    true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: LaunchTemplateVersion
										"launch_template_version": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The version of the launch template to use for faster launching for a Windows AMI.",
											Optional:    true,
											Computed:    true,
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: MaxParallelLaunches
								"max_parallel_launches": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The maximum number of parallel instances that are launched for creating resources.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SnapshotConfiguration
								"snapshot_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: TargetResourceCount
										"target_resource_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Description: "The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.",
											Optional:    true,
											Computed:    true,
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The Windows faster-launching configurations to use for AMI distribution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LaunchTemplateConfigurations
					"launch_template_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AccountId
								"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The account ID that this configuration applies to.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: LaunchTemplateId
								"launch_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Identifies the EC2 launch template to use.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: SetDefaultVersion
								"set_default_version": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Description: "Set the specified EC2 launch template as the default launch template for the specified account.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A group of launchTemplateConfiguration settings that apply to image distribution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LicenseConfigurationArns
					"license_configuration_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The License Manager Configuration to associate with the AMI in the specified Region.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Region
					"region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "region",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The distributions of the distribution configuration.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the distribution configuration.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the distribution configuration.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The tags associated with the component.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The tags associated with the component.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::DistributionConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::DistributionConfiguration").WithTerraformTypeName("awscc_imagebuilder_distribution_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                           "AccountId",
		"ami_distribution_configuration":       "AmiDistributionConfiguration",
		"ami_tags":                             "AmiTags",
		"arn":                                  "Arn",
		"container_distribution_configuration": "ContainerDistributionConfiguration",
		"container_tags":                       "ContainerTags",
		"description":                          "Description",
		"distributions":                        "Distributions",
		"enabled":                              "Enabled",
		"fast_launch_configurations":           "FastLaunchConfigurations",
		"kms_key_id":                           "KmsKeyId",
		"launch_permission_configuration":      "LaunchPermissionConfiguration",
		"launch_template":                      "LaunchTemplate",
		"launch_template_configurations":       "LaunchTemplateConfigurations",
		"launch_template_id":                   "LaunchTemplateId",
		"launch_template_name":                 "LaunchTemplateName",
		"launch_template_version":              "LaunchTemplateVersion",
		"license_configuration_arns":           "LicenseConfigurationArns",
		"max_parallel_launches":                "MaxParallelLaunches",
		"name":                                 "Name",
		"organization_arns":                    "OrganizationArns",
		"organizational_unit_arns":             "OrganizationalUnitArns",
		"region":                               "Region",
		"repository_name":                      "RepositoryName",
		"service":                              "Service",
		"set_default_version":                  "SetDefaultVersion",
		"snapshot_configuration":               "SnapshotConfiguration",
		"tags":                                 "Tags",
		"target_account_ids":                   "TargetAccountIds",
		"target_repository":                    "TargetRepository",
		"target_resource_count":                "TargetResourceCount",
		"user_groups":                          "UserGroups",
		"user_ids":                             "UserIds",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
