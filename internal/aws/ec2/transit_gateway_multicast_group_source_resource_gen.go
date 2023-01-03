// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_transit_gateway_multicast_group_source", transitGatewayMulticastGroupSourceResource)
}

// transitGatewayMulticastGroupSourceResource returns the Terraform awscc_ec2_transit_gateway_multicast_group_source resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastGroupSource resource.
func transitGatewayMulticastGroupSourceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupIpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address assigned to the transit gateway multicast group.",
		//	  "type": "string"
		//	}
		"group_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address assigned to the transit gateway multicast group.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GroupMember
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that the resource is a transit gateway multicast group member.",
		//	  "type": "boolean"
		//	}
		"group_member": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GroupSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that the resource is a transit gateway multicast group member.",
		//	  "type": "boolean"
		//	}
		"group_source": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MemberType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The member type (for example, static).",
		//	  "type": "string"
		//	}
		"member_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The member type (for example, static).",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway attachment.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway attachment.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the resource.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of resource, for example a VPC attachment.",
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of resource, for example a VPC attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source type.",
		//	  "type": "string"
		//	}
		"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source type.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet.",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway attachment.",
		//	  "type": "string"
		//	}
		"transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayMulticastDomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway multicast domain.",
		//	  "type": "string"
		//	}
		"transit_gateway_multicast_domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway multicast domain.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "The AWS::EC2::TransitGatewayMulticastGroupSource registers and deregisters members and sources (network interfaces) with the transit gateway multicast group",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastGroupSource").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_group_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_ip_address":                    "GroupIpAddress",
		"group_member":                        "GroupMember",
		"group_source":                        "GroupSource",
		"member_type":                         "MemberType",
		"network_interface_id":                "NetworkInterfaceId",
		"resource_id":                         "ResourceId",
		"resource_type":                       "ResourceType",
		"source_type":                         "SourceType",
		"subnet_id":                           "SubnetId",
		"transit_gateway_attachment_id":       "TransitGatewayAttachmentId",
		"transit_gateway_multicast_domain_id": "TransitGatewayMulticastDomainId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
