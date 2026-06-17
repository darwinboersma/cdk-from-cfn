using Amazon.CDK;
using Amazon.CDK.AWS.SNS;
using Constructs;
using System.Collections.Generic;

namespace ConditionSubStack
{
    public class ConditionSubStackProps : StackProps
    {
        public string EnvName { get; set; }

    }

    /// <summary>
    /// Exercises Fn::Sub inside a Condition
    /// </summary>
    public class ConditionSubStack : Stack
    {
        public ConditionSubStack(Construct scope, string id, ConditionSubStackProps props = null) : base(scope, id, props)
        {

            // Conditions
            bool hasCustomSuffix = $"{props.EnvName}-suffix" == "prod-suffix";

            // Resources
            var topic = new CfnTopic(this, "Topic", new CfnTopicProps
            {
            });
        }
    }
}
