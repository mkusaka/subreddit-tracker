# aws
## [1][Week of Sept 28th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/j1dndd/week_of_sept_28th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/j1dndd/week_of_sept_28th_what_are_your_favorite_aws_tips/
---
Share your AWS Tips
## [2][AWS CloudFormation now supports StackSets Resource Type in the CloudFormation Registry](https://www.reddit.com/r/aws/comments/j2euog/aws_cloudformation_now_supports_stacksets/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/aws-cloudformation-now-supports-stacksets-resource-type-in-the-cloudformation-registry/
---

## [3][Introducing AWS cost anomaly detection (preview)](https://www.reddit.com/r/aws/comments/j1ymwz/introducing_aws_cost_anomaly_detection_preview/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/introducing-aws-cost-anomaly-detection-preview/#:~:text=AWS%20Cost%20Anomaly%20Detection%20is,increases)%20with%20minimal%20user%20intervention.
---

## [4][aws-cost-saver - A tiny CLI tool to help save costs in development environments when you're asleep and don't need them!](https://www.reddit.com/r/aws/comments/j2j590/awscostsaver_a_tiny_cli_tool_to_help_save_costs/)
- url: https://www.npmjs.com/package/aws-cost-saver
---

## [5][Have an API gateway with JWT Authentication in front of two microservices. I am wondering if I would put a load balancer in front of the API gateway or the Microservice instances.](https://www.reddit.com/r/aws/comments/j2lmq7/have_an_api_gateway_with_jwt_authentication_in/)
- url: https://www.reddit.com/r/aws/comments/j2lmq7/have_an_api_gateway_with_jwt_authentication_in/
---
Relatively new to AWS, at least building this type of architecture. An ALB would go in front of my instances, but would the NLB go in front of my API gateway? Also, I built the API gateway with Amazons service to save time, but I've seen you can do this with NGINX, does anyone recommend this?
## [6][Step Functions - ECS/Fargate integration](https://www.reddit.com/r/aws/comments/j2llwh/step_functions_ecsfargate_integration/)
- url: https://www.reddit.com/r/aws/comments/j2llwh/step_functions_ecsfargate_integration/
---
Fargate has the ability to scale based on alarms, often times based on the load volume in it's SQS (I guess?).

What would you all use to scale the number of Fargate tasks when orchestrating the data input to the Fargate task with Step Functions?

EDIT: Revised question to specifically ask about scaling out the number of Fargate tasks.
## [7][Extracting email template variables. Useful for SES.](https://www.reddit.com/r/aws/comments/j2kxp9/extracting_email_template_variables_useful_for_ses/)
- url: https://www.reddit.com/r/aws/comments/j2kxp9/extracting_email_template_variables_useful_for_ses/
---
So I thought I'd share the regex I use to pull email template variables so that I can test SES emails.

`{{[\s]*[#/]?[a-zA-Z][a-zA-Z0-9\.]*[\s]*[a-zA-Z0-9\.]*[\s]*}}`

This regex isn't perfect like I've explained [here](https://zeer0.com/b/extracting-variables-from-email-templates.html) but it does get all the  handlebar style strings out. Hope it helps someone googling this.
## [8][Using Lambda to perform get requests, but only getting part of the response, solution?](https://www.reddit.com/r/aws/comments/j2kspp/using_lambda_to_perform_get_requests_but_only/)
- url: https://www.reddit.com/r/aws/comments/j2kspp/using_lambda_to_perform_get_requests_but_only/
---
Hello there strangers, im having an issue with my Lambda function where I use it to run multiple GET requests, but one of the requests are only returning part(roughly half) of its response, which is a problem as the field i need is allways the part of the response missing //sigh... As the exact code used locally does not experience this issue, im thinking the problem is with Lambda and its limitations, but its not clear to me what exactly goes wrong and if i can just change it with a configuration change or maybe some code compression. Anyone who has a tip on how to fix this?  


The python code:

from urllib.request import Request, urlopen   

def some\_method():

   url = "some url here"

req = Request(url )

req.add\_header('authorization', mykey) 

content = urlopen(req).read()

print(content) 

print("Bytesize of response: " + str(len(content))) 

some\_method()  


The response is allways cutting off at the same position in the JSON response, and the lenght of the response is allways the same.
## [9][AWS Serverless Application Model : Guide to writing your first AWS SAM Application](https://www.reddit.com/r/aws/comments/j29vsr/aws_serverless_application_model_guide_to_writing/)
- url: https://rajanpanchal.net/complete-guide-to-writing-your-first-aws-sam-application/
---

## [10][RDS Performance](https://www.reddit.com/r/aws/comments/j2iwpk/rds_performance/)
- url: https://www.reddit.com/r/aws/comments/j2iwpk/rds_performance/
---
Hello guys,   


what's the best way to run a load test on my RDS database to analyse performance?
## [11][Should't I wait for ECS container instance to be registered when I launch a new instance in ASGs?](https://www.reddit.com/r/aws/comments/j2fqqq/shouldt_i_wait_for_ecs_container_instance_to_be/)
- url: https://www.reddit.com/r/aws/comments/j2fqqq/shouldt_i_wait_for_ecs_container_instance_to_be/
---
I found articles on draining instances automatically when terminating instances, but what about launching?
Wouldn't it be better to wait for container instances registration as well?
I couldn't find any example on this.

 https://github.com/aws-samples/ecs-cid-sample

This example only have lifecycle hook for draining.

I've never experienced cases where instance registration fails, though in theory the instance should not be considered healthy if it's not shown in container instances list.

EDIT: After searching with different keywords, I've found this. https://github.com/awslabs/ecs-cluster-manager

The python script seems to have the ability to check if new ECS container instance is healthy.
