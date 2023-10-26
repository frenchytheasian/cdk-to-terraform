import aws_cdk as core
import aws_cdk.assertions as assertions

from test_cdk.test_cdk_stack import TestCdkStack

# example tests. To run these tests, uncomment this file along with the example
# resource in test_cdk/test_cdk_stack.py
def test_sqs_queue_created():
    app = core.App()
    stack = TestCdkStack(app, "test-cdk")
    template = assertions.Template.from_stack(stack)

#     template.has_resource_properties("AWS::SQS::Queue", {
#         "VisibilityTimeout": 300
#     })
