// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_elasticbeanstalk_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_elasticbeanstalk_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticBeanstalk::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the application that is associated with this environment.",
		//	  "type": "string"
		//	}
		"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the application that is associated with this environment.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CNAMEPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.",
		//	  "type": "string"
		//	}
		"cname_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Your description for this environment.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Your description for this environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointURL
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"endpoint_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique name for the environment.",
		//	  "type": "string"
		//	}
		"environment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique name for the environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OperationsRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.",
		//	  "type": "string"
		//	}
		"operations_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OptionSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Key-value pairs defining configuration options for this environment, such as the instance type.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Namespace": {
		//	        "description": "A unique namespace that identifies the option's associated AWS resource.",
		//	        "type": "string"
		//	      },
		//	      "OptionName": {
		//	        "description": "The name of the configuration option.",
		//	        "type": "string"
		//	      },
		//	      "ResourceName": {
		//	        "description": "A unique resource name for the option setting. Use it for a time–based scaling configuration option.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The current value for the configuration option.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Namespace",
		//	      "OptionName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"option_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Namespace
					"namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique namespace that identifies the option's associated AWS resource.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: OptionName
					"option_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the configuration option.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceName
					"resource_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A unique resource name for the option setting. Use it for a time–based scaling configuration option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The current value for the configuration option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Key-value pairs defining configuration options for this environment, such as the instance type.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// OptionSettings is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PlatformArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the custom platform to use with the environment.",
		//	  "type": "string"
		//	}
		"platform_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the custom platform to use with the environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SolutionStackName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.",
		//	  "type": "string"
		//	}
		"solution_stack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.",
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
		//	  "description": "Specifies the tags applied to resources in the environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Description: "The key name of the tag.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies the tags applied to resources in the environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Elastic Beanstalk configuration template to use with the environment.",
		//	  "type": "string"
		//	}
		"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Elastic Beanstalk configuration template to use with the environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// TemplateName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.",
		//	  "properties": {
		//	    "Name": {
		//	      "description": "The name of this environment tier.",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "The type of this environment tier.",
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tier": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of this environment tier.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of this environment tier.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionLabel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the application version to deploy.",
		//	  "type": "string"
		//	}
		"version_label": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the application version to deploy.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::ElasticBeanstalk::Environment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::Environment").WithTerraformTypeName("awscc_elasticbeanstalk_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":    "ApplicationName",
		"cname_prefix":        "CNAMEPrefix",
		"description":         "Description",
		"endpoint_url":        "EndpointURL",
		"environment_name":    "EnvironmentName",
		"key":                 "Key",
		"name":                "Name",
		"namespace":           "Namespace",
		"operations_role":     "OperationsRole",
		"option_name":         "OptionName",
		"option_settings":     "OptionSettings",
		"platform_arn":        "PlatformArn",
		"resource_name":       "ResourceName",
		"solution_stack_name": "SolutionStackName",
		"tags":                "Tags",
		"template_name":       "TemplateName",
		"tier":                "Tier",
		"type":                "Type",
		"value":               "Value",
		"version":             "Version",
		"version_label":       "VersionLabel",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateName",
		"/properties/OptionSettings",
		"/properties/OptionSettings/*/OptionName",
		"/properties/OptionSettings/*/ResourceName",
		"/properties/OptionSettings/*/Namespace",
		"/properties/OptionSettings/*/Value",
	})
	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(210)

	opts = opts.WithUpdateTimeoutInMinutes(300)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
