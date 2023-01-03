// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ecs_cluster_capacity_provider_associations", clusterCapacityProviderAssociationsResource)
}

// clusterCapacityProviderAssociationsResource returns the Terraform awscc_ecs_cluster_capacity_provider_associations resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECS::ClusterCapacityProviderAssociations resource.
func clusterCapacityProviderAssociationsResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CapacityProviders
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of capacity providers to associate with the cluster",
		//	  "items": {
		//	    "anyOf": [
		//	      {},
		//	      {}
		//	    ],
		//	    "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"capacity_providers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "List of capacity providers to associate with the cluster",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Cluster
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the cluster",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cluster": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the cluster",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DefaultCapacityProviderStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of capacity providers to associate with the cluster",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Base": {
		//	        "maximum": 100000,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      },
		//	      "CapacityProvider": {
		//	        "anyOf": [
		//	          {},
		//	          {}
		//	        ],
		//	        "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
		//	        "type": "string"
		//	      },
		//	      "Weight": {
		//	        "maximum": 1000,
		//	        "minimum": 0,
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "CapacityProvider"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"default_capacity_provider_strategy": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Base
					"base": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 100000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: CapacityProvider
					"capacity_provider": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Weight
					"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(0, 1000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of capacity providers to associate with the cluster",
			Required:    true,
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
		Description: "Associate a set of ECS Capacity Providers with a specified ECS Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::ClusterCapacityProviderAssociations").WithTerraformTypeName("awscc_ecs_cluster_capacity_provider_associations")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"base":                               "Base",
		"capacity_provider":                  "CapacityProvider",
		"capacity_providers":                 "CapacityProviders",
		"cluster":                            "Cluster",
		"default_capacity_provider_strategy": "DefaultCapacityProviderStrategy",
		"weight":                             "Weight",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
