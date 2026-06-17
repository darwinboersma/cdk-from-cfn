import * as cdk from 'aws-cdk-lib';
import * as sns from 'aws-cdk-lib/aws-sns';

export interface ConditionSubStackProps extends cdk.StackProps {
  /**
   */
  readonly envName: string;
}

/**
 * Exercises Fn::Sub inside a Condition
 */
export class ConditionSubStack extends cdk.Stack {
  public constructor(scope: cdk.App, id: string, props: ConditionSubStackProps) {
    super(scope, id, props);

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
