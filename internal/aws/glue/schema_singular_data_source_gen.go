// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_glue_schema", schemaDataSource)
}

// schemaDataSource returns the Terraform awscc_glue_schema data source.
// This Terraform data source corresponds to the CloudFormation AWS::Glue::Schema resource.
func schemaDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name for the Schema.",
		//	  "pattern": "arn:aws(-(cn|us-gov|iso(-[bef])?))?:glue:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name for the Schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CheckpointVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specify checkpoint version for update. This is only required to update the Compatibility.",
		//	  "properties": {
		//	    "IsLatest": {
		//	      "description": "Indicates if the latest version needs to be updated.",
		//	      "type": "boolean"
		//	    },
		//	    "VersionNumber": {
		//	      "description": "Indicates the version number in the schema to update.",
		//	      "maximum": 100000,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"checkpoint_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IsLatest
				"is_latest": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates if the latest version needs to be updated.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VersionNumber
				"version_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Indicates the version number in the schema to update.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specify checkpoint version for update. This is only required to update the Compatibility.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Compatibility
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Compatibility setting for the schema.",
		//	  "enum": [
		//	    "NONE",
		//	    "DISABLED",
		//	    "BACKWARD",
		//	    "BACKWARD_ALL",
		//	    "FORWARD",
		//	    "FORWARD_ALL",
		//	    "FULL",
		//	    "FULL_ALL"
		//	  ],
		//	  "type": "string"
		//	}
		"compatibility": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Compatibility setting for the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DataFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
		//	  "enum": [
		//	    "AVRO",
		//	    "JSON",
		//	    "PROTOBUF"
		//	  ],
		//	  "type": "string"
		//	}
		"data_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the schema. If description is not provided, there will not be any default value for this.",
		//	  "maxLength": 1000,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the schema. If description is not provided, there will not be any default value for this.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InitialSchemaVersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents the version ID associated with the initial schema version.",
		//	  "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
		//	  "type": "string"
		//	}
		"initial_schema_version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents the version ID associated with the initial schema version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the schema.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Registry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Identifier for the registry which the schema is part of.",
		//	  "properties": {
		//	    "Arn": {
		//	      "description": "Amazon Resource Name for the Registry.",
		//	      "pattern": "arn:aws(-(cn|us-gov|iso(-[bef])?))?:glue:.*",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "Name of the registry in which the schema will be created.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"registry": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Amazon Resource Name for the Registry.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Name of the registry in which the schema will be created.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Identifier for the registry which the schema is part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Definition for the initial schema version in plain-text.",
		//	  "maxLength": 170000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"schema_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Definition for the initial schema version in plain-text.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of tags to tag the schema",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key to identify the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Corresponding tag value for the key.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key to identify the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Corresponding tag value for the key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of tags to tag the schema",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Glue::Schema",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Schema").WithTerraformTypeName("awscc_glue_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"checkpoint_version":        "CheckpointVersion",
		"compatibility":             "Compatibility",
		"data_format":               "DataFormat",
		"description":               "Description",
		"initial_schema_version_id": "InitialSchemaVersionId",
		"is_latest":                 "IsLatest",
		"key":                       "Key",
		"name":                      "Name",
		"registry":                  "Registry",
		"schema_definition":         "SchemaDefinition",
		"tags":                      "Tags",
		"value":                     "Value",
		"version_number":            "VersionNumber",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
