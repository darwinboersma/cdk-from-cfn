from aws_cdk import Stack
import aws_cdk as cdk
import aws_cdk.aws_sns as sns
from constructs import Construct

"""
  Exercises Fn::Sub inside a Condition
"""
class ConditionSubConstruct(Construct):
  def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
    super().__init__(scope, construct_id)

    # Conditions
    has_custom_suffix = f"""{props['envName']}-suffix""" == 'prod-suffix'

    # Resources
    topic = sns.CfnTopic(self, 'Topic',
        ) if has_custom_suffix else None
    if (topic is not None):


