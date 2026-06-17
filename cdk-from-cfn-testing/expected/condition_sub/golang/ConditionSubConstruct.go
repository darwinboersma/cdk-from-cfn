package main

import (
	"fmt"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	sns "github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ConditionSubConstructProps struct {
	EnvName *string
}

/// Exercises Fn::Sub inside a Condition
type ConditionSubConstruct struct {
	constructs.Construct
}

func NewConditionSubConstruct(scope constructs.Construct, id string, props *ConditionSubConstructProps) *ConditionSubConstruct {
	construct := constructs.NewConstruct(scope, &id)

	hasCustomSuffix := jsii.String(fmt.Sprintf("%v-suffix", props.EnvName)) == jsii.String("prod-suffix")

	sns.NewCfnTopic(
		construct,
		jsii.String("Topic"),
		&sns.CfnTopicProps{
		},
	)

	return &ConditionSubConstruct{
		Construct: construct,
	}
}

