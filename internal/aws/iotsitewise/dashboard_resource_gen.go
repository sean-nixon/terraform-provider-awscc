// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iotsitewise_dashboard", dashboardResource)
}

// dashboardResource returns the Terraform awscc_iotsitewise_dashboard resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTSiteWise::Dashboard resource.
func dashboardResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DashboardArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dashboard.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DashboardDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dashboard definition specified in a JSON literal.",
		//	  "type": "string"
		//	}
		"dashboard_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The dashboard definition specified in a JSON literal.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the dashboard.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the dashboard.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DashboardName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A friendly name for the dashboard.",
		//	  "type": "string"
		//	}
		"dashboard_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A friendly name for the dashboard.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the project in which to create the dashboard.",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the project in which to create the dashboard.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the dashboard.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the dashboard.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::IoTSiteWise::Dashboard",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Dashboard").WithTerraformTypeName("awscc_iotsitewise_dashboard")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dashboard_arn":         "DashboardArn",
		"dashboard_definition":  "DashboardDefinition",
		"dashboard_description": "DashboardDescription",
		"dashboard_id":          "DashboardId",
		"dashboard_name":        "DashboardName",
		"key":                   "Key",
		"project_id":            "ProjectId",
		"tags":                  "Tags",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
