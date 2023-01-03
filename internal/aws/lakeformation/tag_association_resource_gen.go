// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_lakeformation_tag_association", tagAssociationResource)
}

// tagAssociationResource returns the Terraform awscc_lakeformation_tag_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::LakeFormation::TagAssociation resource.
func tagAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LFTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Lake Formation Tags to associate with the Lake Formation Resource",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "CatalogId": {
		//	        "maxLength": 12,
		//	        "minLength": 12,
		//	        "type": "string"
		//	      },
		//	      "TagKey": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "TagValues": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "maxLength": 256,
		//	          "minLength": 0,
		//	          "type": "string"
		//	        },
		//	        "maxItems": 50,
		//	        "minItems": 1,
		//	        "type": "array"
		//	      }
		//	    },
		//	    "required": [
		//	      "CatalogId",
		//	      "TagKey",
		//	      "TagValues"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"lf_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CatalogId
					"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(12, 12),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: TagKey
					"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: TagValues
					"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Required:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.SizeBetween(1, 50),
							listvalidator.ValueStringsAre(
								stringvalidator.LengthBetween(0, 256),
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							Multiset(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of Lake Formation Tags to associate with the Lake Formation Resource",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Resource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Resource to tag with the Lake Formation Tags",
		//	  "properties": {
		//	    "Catalog": {
		//	      "additionalProperties": false,
		//	      "type": "object"
		//	    },
		//	    "Database": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Table": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TableWildcard": {
		//	          "additionalProperties": false,
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TableWithColumns": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "ColumnNames": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 255,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName",
		//	        "Name",
		//	        "ColumnNames"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Catalog
				"catalog": schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Database
				"database": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Table
				"table": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: TableWildcard
						"table_wildcard": schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TableWithColumns
				"table_with_columns": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: ColumnNames
						"column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 255),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								Multiset(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Resource to tag with the Lake Formation Tags",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
		//	  "type": "string"
		//	}
		"resource_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TagsIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
		//	  "type": "string"
		//	}
		"tags_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
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
		Description: "A resource schema representing a Lake Formation Tag Association. While tag associations are not explicit Lake Formation resources, this CloudFormation resource can be used to associate tags with Lake Formation entities.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::TagAssociation").WithTerraformTypeName("awscc_lakeformation_tag_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":             "Catalog",
		"catalog_id":          "CatalogId",
		"column_names":        "ColumnNames",
		"database":            "Database",
		"database_name":       "DatabaseName",
		"lf_tags":             "LFTags",
		"name":                "Name",
		"resource":            "Resource",
		"resource_identifier": "ResourceIdentifier",
		"table":               "Table",
		"table_wildcard":      "TableWildcard",
		"table_with_columns":  "TableWithColumns",
		"tag_key":             "TagKey",
		"tag_values":          "TagValues",
		"tags_identifier":     "TagsIdentifier",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
