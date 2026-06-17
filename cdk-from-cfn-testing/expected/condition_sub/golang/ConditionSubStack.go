package main

import (
	"fmt"

	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	sns "github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ConditionSubStackProps struct {
	cdk.StackProps
	EnvName *string
}

/// Exercises Fn::Sub inside a Condition
type ConditionSubStack struct {
	cdk.Stack
}

func NewConditionSubStack(scope constructs.Construct, id string, props *ConditionSubStackProps) *ConditionSubStack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	hasCustomSuffix := jsii.String(fmt.Sprintf("%v-suffix", props.EnvName)) == jsii.String("prod-suffix")

	sns.NewCfnTopic(
		stack,
		jsii.String("Topic"),
		&sns.CfnTopicProps{
		},
	)

	return &ConditionSubStack{
		Stack: stack,
	}
}

