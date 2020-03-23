# aws
## [1][Static IPs with an ALB?](https://www.reddit.com/r/aws/comments/fnjj41/static_ips_with_an_alb/)
- url: https://www.reddit.com/r/aws/comments/fnjj41/static_ips_with_an_alb/
---
We're launching a new service soon and I'd like your input on how to set up the networking infrastructure. Here are the main points that I'd like to accomplish and a bit of background

* The service runs on ECS (Fargate) with an Nginx reverse proxy in front
* Would like to use an application load balancer in front of the service, mainly for the nice monitoring capabilities you don't get with an NLB (and also built-in authentication which I'd like to use to keep our test environment private)
* Would strongly prefer having static IPs (presumably using EIPs) for the service to ease DNS setup 
* This service will be at the domain apex and unfortunately name server delegation to Route 53 is probably not feasible

Currently ALBs do not support EIPs while NLBs do. AWS has published a blog post with a solution to get around that problem, but it looks byzantine and seems like overengineering to solve what should be a simple problem: [https://aws.amazon.com/blogs/networking-and-content-delivery/using-static-ip-addresses-for-application-load-balancers/](https://aws.amazon.com/blogs/networking-and-content-delivery/using-static-ip-addresses-for-application-load-balancers/)

Does anyone have experience with that setup? If not I'll probably go with a setup of EIPs -&gt; NLB -&gt; ECS service, even though I'll miss out on the nice things that ALBs provide.
## [2][VPC with Public and Private Subnets - EC2 in public subnet has no IP](https://www.reddit.com/r/aws/comments/fnggp0/vpc_with_public_and_private_subnets_ec2_in_public/)
- url: https://www.reddit.com/r/aws/comments/fnggp0/vpc_with_public_and_private_subnets_ec2_in_public/
---
Just tried the VPC wizard with "VPC with Public and Private Subnets" but the EC2 instance I created in the public subnet has no public IP.

What am I missing?
## [3][Automating docker deployments to EKS with CDK and CodePipeline](https://www.reddit.com/r/aws/comments/fnc69t/automating_docker_deployments_to_eks_with_cdk_and/)
- url: https://www.reddit.com/r/aws/comments/fnc69t/automating_docker_deployments_to_eks_with_cdk_and/
---
Anyone successfully created a CodePipeline deployment with CDK? I found [this doc](https://github.com/aws/aws-cdk/tree/master/packages/%40aws-cdk/aws-codepipeline-actions) but it skips over EKS completely. And I couldn't find anything online that describes the ins and outs of this pipeline -- it's like they don't want us to use EKS at all (all the docs, examples, and helper classes are geared towards ECS). 

P.S. If codepipeline isn't the way to go, what would you recommend for an easy-to-setup and use CI/CD pipeline?
## [4][Problem Deploying a docker container on ECS](https://www.reddit.com/r/aws/comments/fn7zhk/problem_deploying_a_docker_container_on_ecs/)
- url: https://www.reddit.com/r/aws/comments/fn7zhk/problem_deploying_a_docker_container_on_ecs/
---
Hey all, today is my first day on AWS, so apologies in advance if this is a stupid question. I have a Dash App that I wrapped up in a docker container, which I'm not trying to deploy on AWS. I *think* that I've successfully got an ECS instance up and running, but I can't access my app access the given Public DNS. I've been following tutorials all day, but something is just not working out for me. The Cluster, ECS Instance, and Task are all running and everything appear to be in a good state. However, when I try to access either the Public DNS or Public IP, I get hit with a Problem Loading Page. I'm not really sure how to troubleshoot at this point. Does anyone have any idea what may be going on?
## [5][Is the Workspaces service getting hammered due to the increase of WFH persons?](https://www.reddit.com/r/aws/comments/fnj5ix/is_the_workspaces_service_getting_hammered_due_to/)
- url: https://www.reddit.com/r/aws/comments/fnj5ix/is_the_workspaces_service_getting_hammered_due_to/
---
I'm sure that the internet as a whole is getting hammered with the increase of persons WFH, and that AWS announcing:

&gt;We are announcing two new offers that enable you to use Amazon WorkSpaces and Amazon WorkDocs for up to 50 users at no charge. Both offers are for new customers that have not previously used these services and are available through June 30, 2020.

... probably isn't helping much.  

Is anyone seeing an increase of end users complaining of their Workspaces being "laggy" or unresponsive?  I have some users on the other side of the world, and trying to determine if there is anything I can do to help them.
## [6][Websocket](https://www.reddit.com/r/aws/comments/fniuvx/websocket/)
- url: https://www.reddit.com/r/aws/comments/fniuvx/websocket/
---
Hey guys, I'm using serverless-offline and I already have a websocket configured, but i want to know is it possible to open multiple websocket connections in different ports using serverless?
## [7][Python/boto3 question](https://www.reddit.com/r/aws/comments/fn6ank/pythonboto3_question/)
- url: https://www.reddit.com/r/aws/comments/fn6ank/pythonboto3_question/
---
Hey everyone, 

Was looking for some help on writing a simple python script that takes a user input like “us-East-1” and reboots all instances there. So like “script.py us-east-1” on a regular shell command line. Already have profile, keys, and permissions loaded up and can perform basic things just can’t narrow it down. Thanks in advance!
## [8][Regarding Amazon SFTP](https://www.reddit.com/r/aws/comments/fn8e3j/regarding_amazon_sftp/)
- url: https://www.reddit.com/r/aws/comments/fn8e3j/regarding_amazon_sftp/
---
Users of my product/app (not a webapp) needs to submit files for processing via SFTP and after processing we need to send the results back to user via same SFTP.

 I learnt that behind the scenes Amazon SFTP stores the files in S3 buckets. If I have several users and if they should not be able to see each other's file then I create multiple S3 buckets like User1, User2, etc.

Under User folder, I further create folders like User/Inbound and User/Outbound so that we can receive and send the files.

Question:

If I need to mount all these S3 buckets to the EC2 instance where my app is running do I need to mount all the S3's separately or could there be a parent folder that I can mount so that I can do it only once and need not create a input handler on my app?

Also, after processing can I pass the S3 bucket names corresponding to each user as a parameter so that I don't have to create individual S3 outputs on my app?

Note - Apologies if I am not using AWS terms and using regular NAS terms, as I am new to AWS.
## [9][EventBridge: The key component in Serverless Architectures (Serverless-Transformation Article)](https://www.reddit.com/r/aws/comments/fngqtz/eventbridge_the_key_component_in_serverless/)
- url: https://medium.com/serverless-transformation/eventbridge-the-key-component-in-serverless-architectures-e7d4e60fca2d
---

## [10][SQS Policy that only allows S3 notifications from objects from a specific user](https://www.reddit.com/r/aws/comments/fngfnq/sqs_policy_that_only_allows_s3_notifications_from/)
- url: https://www.reddit.com/r/aws/comments/fngfnq/sqs_policy_that_only_allows_s3_notifications_from/
---
Dear Experts! 

I'm building a messaging system for a customer that responds to all create events on a s3 bucket. This system has the task of responding to new files files being written to the bucket, editing them, and the re-uploading the edited files back to the same bucket; under the same key. As you might understand, this causes a problem: Since my own edits will trigger new create notifications themselves, the event systems gets into an infinite loop.

I have no control over the s3 bucket keys, so I cannot use the build-in notification filters. I could check in my code if the changes where already made and don't upload the file to stop the loop. But this still means that some unwanted events from objects end up in my SQS queue. 

The most promising solution I can think of is to have some SQS policy that only allows s3 to send a message if the s3 key was put by a certain user. That way I can prevent s3 from sending events from objects that I uploaded myself. The hope that this is even possible comes from the fact that I see the ["userIdentity.principalId"](https://docs.aws.amazon.com/AmazonS3/latest/dev/notification-content-structure.html) change in the notification message. 

Or maybe there are other options like the other [policy conditions](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-overview-of-managing-access.html#sqs-specifying-conditions-in-policy). I could freely set certain tags on the object for example.

Does anyone know if it is possible to create an SQS policy that only allows s3 to send notifications to the queue if an s3 key was created by certain user? Or maybe have another solution to my problem.

Thank you!
