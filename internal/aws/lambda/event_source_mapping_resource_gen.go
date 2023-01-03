// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_lambda_event_source_mapping", eventSourceMappingResource)
}

// eventSourceMappingResource returns the Terraform awscc_lambda_event_source_mapping resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lambda::EventSourceMapping resource.
func eventSourceMappingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AmazonManagedKafkaEventSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specific configuration settings for an MSK event source.",
		//	  "properties": {
		//	    "ConsumerGroupId": {
		//	      "description": "The identifier for the Kafka Consumer Group to join.",
		//	      "maxLength": 200,
		//	      "minLength": 1,
		//	      "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"amazon_managed_kafka_event_source_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConsumerGroupId
				"consumer_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier for the Kafka Consumer Group to join.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 200),
						stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9-\\/*:_+=.@-]*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specific configuration settings for an MSK event source.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BatchSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum number of items to retrieve in a single batch.",
		//	  "maximum": 10000,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"batch_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum number of items to retrieve in a single batch.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 10000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BisectBatchOnFunctionError
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) If the function returns an error, split the batch in two and retry.",
		//	  "type": "boolean"
		//	}
		"bisect_batch_on_function_error": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "(Streams) If the function returns an error, split the batch in two and retry.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
		//	  "properties": {
		//	    "OnFailure": {
		//	      "additionalProperties": false,
		//	      "description": "The destination configuration for failed invocations.",
		//	      "properties": {
		//	        "Destination": {
		//	          "description": "The Amazon Resource Name (ARN) of the destination resource.",
		//	          "maxLength": 1024,
		//	          "minLength": 12,
		//	          "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"destination_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OnFailure
				"on_failure": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Destination
						"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the destination resource.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 1024),
								stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The destination configuration for failed invocations.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Disables the event source mapping to pause polling and invocation.",
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Disables the event source mapping to pause polling and invocation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventSourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the event source.",
		//	  "maxLength": 1024,
		//	  "minLength": 12,
		//	  "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
		//	  "type": "string"
		//	}
		"event_source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the event source.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(12, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FilterCriteria
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The filter criteria to control event filtering.",
		//	  "properties": {
		//	    "Filters": {
		//	      "description": "List of filters of this FilterCriteria",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The filter object that defines parameters for ESM filtering.",
		//	        "properties": {
		//	          "Pattern": {
		//	            "description": "The filter pattern that defines which events should be passed for invocations.",
		//	            "maxLength": 4096,
		//	            "minLength": 0,
		//	            "pattern": ".*",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"filter_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Filters
				"filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Pattern
							"pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The filter pattern that defines which events should be passed for invocations.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(0, 4096),
									stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of filters of this FilterCriteria",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 20),
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The filter criteria to control event filtering.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FunctionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Lambda function.",
		//	  "maxLength": 140,
		//	  "minLength": 1,
		//	  "pattern": "(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\\d{1}:)?(\\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\\$LATEST|[a-zA-Z0-9-_]+))?",
		//	  "type": "string"
		//	}
		"function_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Lambda function.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 140),
				stringvalidator.RegexMatches(regexp.MustCompile("(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\\d{1}:)?(\\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\\$LATEST|[a-zA-Z0-9-_]+))?"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: FunctionResponseTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) A list of response types supported by the function.",
		//	  "items": {
		//	    "enum": [
		//	      "ReportBatchItemFailures"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "maxLength": 1,
		//	  "minLength": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"function_response_types": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "(Streams) A list of response types supported by the function.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"ReportBatchItemFailures",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Event Source Mapping Identifier UUID.",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Event Source Mapping Identifier UUID.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaximumBatchingWindowInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
		//	  "maximum": 300,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"maximum_batching_window_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(0, 300),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaximumRecordAgeInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
		//	  "maximum": 604800,
		//	  "minimum": -1,
		//	  "type": "integer"
		//	}
		"maximum_record_age_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(-1, 604800),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaximumRetryAttempts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) The maximum number of times to retry when the function returns an error.",
		//	  "maximum": 10000,
		//	  "minimum": -1,
		//	  "type": "integer"
		//	}
		"maximum_retry_attempts": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "(Streams) The maximum number of times to retry when the function returns an error.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(-1, 10000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParallelizationFactor
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) The number of batches to process from each shard concurrently.",
		//	  "maximum": 10,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"parallelization_factor": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "(Streams) The number of batches to process from each shard concurrently.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 10),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Queues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(ActiveMQ) A list of ActiveMQ queues.",
		//	  "items": {
		//	    "maxLength": 1000,
		//	    "minLength": 1,
		//	    "pattern": "[\\s\\S]*",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"queues": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "(ActiveMQ) A list of ActiveMQ queues.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 1000),
					stringvalidator.RegexMatches(regexp.MustCompile("[\\s\\S]*"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScalingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The scaling configuration for the event source.",
		//	  "properties": {
		//	    "MaximumConcurrency": {
		//	      "description": "The maximum number of concurrent functions that the event source can invoke.",
		//	      "maximum": 1000,
		//	      "minimum": 2,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scaling_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaximumConcurrency
				"maximum_concurrency": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of concurrent functions that the event source can invoke.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(2, 1000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The scaling configuration for the event source.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SelfManagedEventSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Self-managed event source endpoints.",
		//	  "properties": {
		//	    "Endpoints": {
		//	      "additionalProperties": false,
		//	      "description": "The endpoints for a self-managed event source.",
		//	      "properties": {
		//	        "KafkaBootstrapServers": {
		//	          "description": "A list of Kafka server endpoints.",
		//	          "items": {
		//	            "description": "The URL of a Kafka server.",
		//	            "maxLength": 300,
		//	            "minLength": 1,
		//	            "pattern": "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9]):[0-9]{1,5}",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"self_managed_event_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Endpoints
				"endpoints": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: KafkaBootstrapServers
						"kafka_bootstrap_servers": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "A list of Kafka server endpoints.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 10),
								listvalidator.UniqueValues(),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 300),
									stringvalidator.RegexMatches(regexp.MustCompile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9]):[0-9]{1,5}"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The endpoints for a self-managed event source.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Self-managed event source endpoints.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SelfManagedKafkaEventSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specific configuration settings for a Self-Managed Apache Kafka event source.",
		//	  "properties": {
		//	    "ConsumerGroupId": {
		//	      "description": "The identifier for the Kafka Consumer Group to join.",
		//	      "maxLength": 200,
		//	      "minLength": 1,
		//	      "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"self_managed_kafka_event_source_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConsumerGroupId
				"consumer_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier for the Kafka Consumer Group to join.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 200),
						stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9-\\/*:_+=.@-]*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specific configuration settings for a Self-Managed Apache Kafka event source.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceAccessConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of SourceAccessConfiguration.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The configuration used by AWS Lambda to access event source",
		//	    "properties": {
		//	      "Type": {
		//	        "description": "The type of source access configuration.",
		//	        "enum": [
		//	          "BASIC_AUTH",
		//	          "VPC_SUBNET",
		//	          "VPC_SECURITY_GROUP",
		//	          "SASL_SCRAM_512_AUTH",
		//	          "SASL_SCRAM_256_AUTH",
		//	          "VIRTUAL_HOST",
		//	          "CLIENT_CERTIFICATE_TLS_AUTH",
		//	          "SERVER_ROOT_CA_CERTIFICATE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "URI": {
		//	        "description": "The URI for the source access configuration resource.",
		//	        "maxLength": 200,
		//	        "minLength": 1,
		//	        "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 22,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"source_access_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of source access configuration.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"BASIC_AUTH",
								"VPC_SUBNET",
								"VPC_SECURITY_GROUP",
								"SASL_SCRAM_512_AUTH",
								"SASL_SCRAM_256_AUTH",
								"VIRTUAL_HOST",
								"CLIENT_CERTIFICATE_TLS_AUTH",
								"SERVER_ROOT_CA_CERTIFICATE",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: URI
					"uri": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The URI for the source access configuration resource.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 200),
							stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9-\\/*:_+=.@-]*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of SourceAccessConfiguration.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 22),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StartingPosition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
		//	  "maxLength": 12,
		//	  "minLength": 6,
		//	  "pattern": "(LATEST|TRIM_HORIZON|AT_TIMESTAMP)+",
		//	  "type": "string"
		//	}
		"starting_position": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(6, 12),
				stringvalidator.RegexMatches(regexp.MustCompile("(LATEST|TRIM_HORIZON|AT_TIMESTAMP)+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StartingPositionTimestamp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
		//	  "type": "number"
		//	}
		"starting_position_timestamp": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
				float64planmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Topics
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Kafka) A list of Kafka topics.",
		//	  "items": {
		//	    "maxLength": 249,
		//	    "minLength": 1,
		//	    "pattern": "^[^.]([a-zA-Z0-9\\-_.]+)",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"topics": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "(Kafka) A list of Kafka topics.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 249),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^.]([a-zA-Z0-9\\-_.]+)"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TumblingWindowInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
		//	  "maximum": 900,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"tumbling_window_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(0, 900),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::Lambda::EventSourceMapping",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::EventSourceMapping").WithTerraformTypeName("awscc_lambda_event_source_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_managed_kafka_event_source_config": "AmazonManagedKafkaEventSourceConfig",
		"batch_size":                             "BatchSize",
		"bisect_batch_on_function_error":         "BisectBatchOnFunctionError",
		"consumer_group_id":                      "ConsumerGroupId",
		"destination":                            "Destination",
		"destination_config":                     "DestinationConfig",
		"enabled":                                "Enabled",
		"endpoints":                              "Endpoints",
		"event_source_arn":                       "EventSourceArn",
		"filter_criteria":                        "FilterCriteria",
		"filters":                                "Filters",
		"function_name":                          "FunctionName",
		"function_response_types":                "FunctionResponseTypes",
		"id":                                     "Id",
		"kafka_bootstrap_servers":                "KafkaBootstrapServers",
		"maximum_batching_window_in_seconds":     "MaximumBatchingWindowInSeconds",
		"maximum_concurrency":                    "MaximumConcurrency",
		"maximum_record_age_in_seconds":          "MaximumRecordAgeInSeconds",
		"maximum_retry_attempts":                 "MaximumRetryAttempts",
		"on_failure":                             "OnFailure",
		"parallelization_factor":                 "ParallelizationFactor",
		"pattern":                                "Pattern",
		"queues":                                 "Queues",
		"scaling_config":                         "ScalingConfig",
		"self_managed_event_source":              "SelfManagedEventSource",
		"self_managed_kafka_event_source_config": "SelfManagedKafkaEventSourceConfig",
		"source_access_configurations":           "SourceAccessConfigurations",
		"starting_position":                      "StartingPosition",
		"starting_position_timestamp":            "StartingPositionTimestamp",
		"topics":                                 "Topics",
		"tumbling_window_in_seconds":             "TumblingWindowInSeconds",
		"type":                                   "Type",
		"uri":                                    "URI",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
