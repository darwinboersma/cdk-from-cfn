using Amazon.CDK;
using Amazon.CDK.AWS.SNS;
using Constructs;
using System.Collections.Generic;

namespace ConditionSubConstruct
{
    public class ConditionSubConstructProps
    {
        public string EnvName { get; set; }

    }

    /// <summary>
    /// Exercises Fn::Sub inside a Condition
    /// </summary>
    public class ConditionSubConstruct : Construct
    {
        public ConditionSubConstruct(Construct scope, string id, ConditionSubConstructProps props = null) : base(scope, id)
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
