import * as cdk from 'aws-cdk-lib';
import * as sns from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';

export interface ConditionSubConstructProps {
  /**
   */
  readonly envName: string;
}

/**
 * Exercises Fn::Sub inside a Condition
 */
export class ConditionSubConstruct extends Construct {
  public constructor(scope: Construct, id: string, props: ConditionSubConstructProps) {
    super(scope, id);

    // Conditions
    const hasCustomSuffix = `${props.envName!}-suffix` === 'prod-suffix';

    // Resources
    const topic = hasCustomSuffix
      ? new sns.CfnTopic(this, 'Topic', {
        })
      : undefined;
    if (topic != null) {
    }
  }
}
