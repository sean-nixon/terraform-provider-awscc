// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_network_performance_metric_subscription", networkPerformanceMetricSubscriptionResource)
}

// networkPerformanceMetricSubscriptionResource returns the Terraform awscc_ec2_network_performance_metric_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::NetworkPerformanceMetricSubscription resource.
func networkPerformanceMetricSubscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Destination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target Region or Availability Zone for the metric to subscribe to.",
		//	  "type": "string"
		//	}
		"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target Region or Availability Zone for the metric to subscribe to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Metric
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The metric type to subscribe to.",
		//	  "type": "string"
		//	}
		"metric": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The metric type to subscribe to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Source
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The starting Region or Availability Zone for metric to subscribe to.",
		//	  "type": "string"
		//	}
		"source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The starting Region or Availability Zone for metric to subscribe to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Statistic
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The statistic to subscribe to.",
		//	  "type": "string"
		//	}
		"statistic": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The statistic to subscribe to.",
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
		Description: "Resource Type definition for AWS::EC2::NetworkPerformanceMetricSubscription",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkPerformanceMetricSubscription").WithTerraformTypeName("awscc_ec2_network_performance_metric_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destination": "Destination",
		"metric":      "Metric",
		"source":      "Source",
		"statistic":   "Statistic",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
