# aws
## [1][Fluent Bit supports Amazon S3 as a destination to route container logs](https://www.reddit.com/r/aws/comments/jbcjb9/fluent_bit_supports_amazon_s3_as_a_destination_to/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/fluent-bit-supports-amazon-s3-destination-route-container-logs/
---

## [2][Need to updated my Elastic Beanstalk environment from PHP 7.2 to 7.4. Whats the best way to change platforms given that I cannot seem to clone my existing configurations into the new platform.](https://www.reddit.com/r/aws/comments/jbnki2/need_to_updated_my_elastic_beanstalk_environment/)
- url: https://www.reddit.com/r/aws/comments/jbnki2/need_to_updated_my_elastic_beanstalk_environment/
---
I have an EB for a site with a lot of environment variables, and some specified security groups ect.  That environment is using an older platform on php 7.2 which reaches its EOL in a couple months but php 7.4 is on a new EB platform.  It doesnt seem that I can either clone my environment and rebuild it with a new platform, nor can I save the configs and then launch a new environment with the new platform.

Is there some automated, or semi automated way to get the various settings from the old environment/platform (things like environment variables, instance counts and autoscaling settings, security groups, load balancer settings) to the new platform?  Is EB cli the way to go? Or will I have to manually reset everything?
## [3][Restricting Access to a Website Without Updating a SG Constantly](https://www.reddit.com/r/aws/comments/jbnc2l/restricting_access_to_a_website_without_updating/)
- url: https://www.reddit.com/r/aws/comments/jbnc2l/restricting_access_to_a_website_without_updating/
---
In a typical on-prem office network I would have a website running on a private subnet so any employee connected to the LAN would have access to it.

Our company is all remote and all in AWS. I have a website currently in a public subnet and port 80/443 are open to all. I'd like to restrict this somehow so only our employees can access it. The issue is we can't constantly manually update the security group for everyone's public IP as they change.

Anyone have any suggestions or unique ways to handle this?
## [4][Best was to backup EC2 instance.](https://www.reddit.com/r/aws/comments/jbj2bi/best_was_to_backup_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/jbj2bi/best_was_to_backup_ec2_instance/
---
Hello, I am new to AWS,

I Wanted to know what are the ways to backup an EC2 instance and,

What is the best way to backup an EC2 instance.

&amp;#x200B;

thank you in advance
## [5][Saas Manual, building products on AWS](https://www.reddit.com/r/aws/comments/jbl6ii/saas_manual_building_products_on_aws/)
- url: https://www.reddit.com/r/aws/comments/jbl6ii/saas_manual_building_products_on_aws/
---
Hi awesome AWS folks. I have launched a course, [saasmanual.com](https://saasmanual.com) which teaches building SaaS products from scratch, on AWS. I previously was part of Cloud9 which was acquired by AWS and spent three years there. I thought it might be nice to share my lessons. Would love to get your feedback, comments, and ideas for what you'd like to learn. 

From an AWS perspective, I'll be using:  
\* CDK for infra as code, it is super cool.  
\* Serverless (Lambda, DynamoDB, API Gateway, etc)  
\* Developer tools (Pipelines, CodeBuild, etc.)

Have a wonderful day ☀️
## [6][Sending collected cpu/memory metrics from ec2 is instance to local file?](https://www.reddit.com/r/aws/comments/jbkvqi/sending_collected_cpumemory_metrics_from_ec2_is/)
- url: https://www.reddit.com/r/aws/comments/jbkvqi/sending_collected_cpumemory_metrics_from_ec2_is/
---
I use lambda to send mail that triggered by sns topic. 

With alarms created by filtering cloudwatch log I can do that. 

But with alarms created from cloudwatch metrics I can’t (or not find a way yet). Alarm with action to send to sns topic but sns topic did not trigger lambda, then lambda didn’t send mail to user to notify about metrics over threshold. 

Why with metrics filter on cloudwatch log, lambda sent mail and with metrics it didn’t?

How to have alarms from metrics to trigger lambda and send mail as with cloudwatch log?

If not, how to send metrics memory/cpu of cloudwatch log agent locally to ec2 instance and then send to cloudwatch log for later filtering?

Thanks in advanced!
## [7][Deciding between ECS, EKS, Fargate, and Lambda](https://www.reddit.com/r/aws/comments/jbketo/deciding_between_ecs_eks_fargate_and_lambda/)
- url: https://www.reddit.com/r/aws/comments/jbketo/deciding_between_ecs_eks_fargate_and_lambda/
---
We have been asked many times about which compute platforms fit to which use case. In order to share our findings, we have put together a paper that compares the different computing platforms of AWS from operational and cost optimization perspectives. 

We're looking for feedbacks and other inputs to keep this paper live from now on: [https://www.thundra.io/whitepaper/the-guide-to-aws-computing-2020](https://www.thundra.io/whitepaper/the-guide-to-aws-computing-2020)
## [8][Can i delete a stack created with serverless framework directly from the console?](https://www.reddit.com/r/aws/comments/jbkakr/can_i_delete_a_stack_created_with_serverless/)
- url: https://www.reddit.com/r/aws/comments/jbkakr/can_i_delete_a_stack_created_with_serverless/
---
I created a stack, deployed, changed the name, redeployed, and now i have two stack. I want to delete the old one directly from the aws console, avoiding editing the yaml file again to do sls remove. Can i use the aws console to remove it without leaving resourse floating around, if it was created by serverless?
## [9][Access denied 403 for AWS Lambda when trying to do multi-account and multi-region copy](https://www.reddit.com/r/aws/comments/jbk85a/access_denied_403_for_aws_lambda_when_trying_to/)
- url: https://www.reddit.com/r/aws/comments/jbk85a/access_denied_403_for_aws_lambda_when_trying_to/
---
My lambda function was working correctly within one account and one region. I wanted to extend that a bit and add support multi-regions and multi-accounts, so I modified my setup a bit.

I am seeing this error message in my Lambda CloudWatch logs:

&gt;AccessDenied: Access Denied status code: 403

This is the source and the destination bucket policy (with very wide permission set for actions)

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "arn:aws:iam::123456789:role/aws-lambda-role"
                },
                "Action": "*",
                "Resource": [         
                    "arn:aws:s3:::bucketA/*", &lt;-- Different for bucket B
                    "arn:aws:s3:::bucketA"
                ]
            }
        ]
    }

The role for the Lambda function is `AWSLambdaBasicExecutionRole`

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Action": [
                    "logs:CreateLogGroup",
                    "logs:CreateLogStream",
                    "logs:PutLogEvents"
                ],
                "Resource": "*"
            }
        ]
    }

Any ideas?
## [10][Add custom attributes to Cloudmap via ECS service deploy](https://www.reddit.com/r/aws/comments/jbk2cd/add_custom_attributes_to_cloudmap_via_ecs_service/)
- url: https://www.reddit.com/r/aws/comments/jbk2cd/add_custom_attributes_to_cloudmap_via_ecs_service/
---
Is it possible to instruct ECS service(I’m using fargate) to add custom attributes to Cloudmap instance?[ECS adds certain attributes by default](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html), but I was thinking of adding additional attributes. Can be done with external tooling ofc, but then that needs to be made aware of scale-in and scale-out operations as well. ie. update every task entry in Cloudmap everytime scale-out happens :(

Are there any best-practices/recommended solutions for this?

  
UPDATE: Unfortunately it's not supported out-of-the-box. Created a feature request: [https://github.com/aws/containers-roadmap/issues/1119](https://github.com/aws/containers-roadmap/issues/1119)  

