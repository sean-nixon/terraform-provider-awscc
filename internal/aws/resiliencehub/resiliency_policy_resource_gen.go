// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_resiliencehub_resiliency_policy", resiliencyPolicyResource)
}

// resiliencyPolicyResource returns the Terraform awscc_resiliencehub_resiliency_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::ResilienceHub::ResiliencyPolicy resource.
func resiliencyPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DataLocationConstraint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Data Location Constraint of the Policy.",
		//	  "enum": [
		//	    "AnyLocation",
		//	    "SameContinent",
		//	    "SameCountry"
		//	  ],
		//	  "type": "string"
		//	}
		"data_location_constraint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Data Location Constraint of the Policy.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"AnyLocation",
					"SameContinent",
					"SameCountry",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Failure Policy.",
		//	      "properties": {
		//	        "RpoInSecs": {
		//	          "description": "RPO in seconds.",
		//	          "type": "integer"
		//	        },
		//	        "RtoInSecs": {
		//	          "description": "RTO in seconds.",
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "RtoInSecs",
		//	        "RpoInSecs"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"policy":                  // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RpoInSecs
					"rpo_in_secs": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "RPO in seconds.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: RtoInSecs
					"rto_in_secs": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "RTO in seconds.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of Resiliency Policy.",
		//	  "maxLength": 500,
		//	  "type": "string"
		//	}
		"policy_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of Resiliency Policy.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Resiliency Policy.",
		//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Resiliency Policy.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Resiliency Policy Tier.",
		//	  "enum": [
		//	    "MissionCritical",
		//	    "Critical",
		//	    "Important",
		//	    "CoreServices",
		//	    "NonCritical"
		//	  ],
		//	  "type": "string"
		//	}
		"tier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Resiliency Policy Tier.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"MissionCritical",
					"Critical",
					"Important",
					"CoreServices",
					"NonCritical",
				),
			}, /*END VALIDATORS*/
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
		Description: "Resource Type Definition for Resiliency Policy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::ResiliencyPolicy").WithTerraformTypeName("awscc_resiliencehub_resiliency_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location_constraint": "DataLocationConstraint",
		"policy":                   "Policy",
		"policy_arn":               "PolicyArn",
		"policy_description":       "PolicyDescription",
		"policy_name":              "PolicyName",
		"rpo_in_secs":              "RpoInSecs",
		"rto_in_secs":              "RtoInSecs",
		"tags":                     "Tags",
		"tier":                     "Tier",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
