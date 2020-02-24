# aws
## [1][Architecting a Successful SaaS](https://www.reddit.com/r/aws/comments/f8ok5z/architecting_a_successful_saas/)
- url: https://www.reddit.com/r/aws/comments/f8ok5z/architecting_a_successful_saas/
---
This series explores how to properly plan and architect a SaaS product offering, designed for hosting on the public cloud. This part answers basic questions, such as, what is a SaaS, what are the alternatives to SaaS for software distribution, and what are the most common SaaS product models:

https://medium.com/@GaryStafford/architecting-a-successful-saas-eaa24c5ad6d7?source=friends_link&amp;sk=4f53d1d8ea0ac87cbb30491e49861c89
## [2][Knowing when a solution is a good fit](https://www.reddit.com/r/aws/comments/f8idur/knowing_when_a_solution_is_a_good_fit/)
- url: https://www.reddit.com/r/aws/comments/f8idur/knowing_when_a_solution_is_a_good_fit/
---
Disclaimer, been only working with aws full time since September.

I joined the cloud team and come to realize that only recently we reached a tipping point at which we can explore new approaches and services.

However, almost daily I find new services that make me question my recent architecture choices for, numer one, a static website and, number two, a fullstack nodejs app (connected to a google service).

For example, I could run the web app using an EC2, elastic beanstalk, amplify and probably one or two additional ways. It's funny at first, but it becomes annoying when you start questioning yourself with almost every announcement.

So I'm wondering what others would do for a) a static website, quarterly updated with a zip upload and b) an all-in-one demo application that uses server side rendered views and only needs to show off a chatbot service from google
## [3][ECS task in internal subnet cannot connect to the internet](https://www.reddit.com/r/aws/comments/f8oiy6/ecs_task_in_internal_subnet_cannot_connect_to_the/)
- url: https://www.reddit.com/r/aws/comments/f8oiy6/ecs_task_in_internal_subnet_cannot_connect_to_the/
---
Configured ECS service to have only internal subnets as allowed subnets because this service cannot be accessed from the internet but is should be able to access internet.

However, the issue is that the service cannot access internet anymore but it can access other services in the same VPC.

Here is the terraform configuration for the VPC and ECS service - https://gist.github.com/genert/7ae99dbff9c7c1ba64edc08694fe460c

As you can see, the route table configuration should make this work but it does not.

Any ideas what the issue could be?

I would also point out that the instance itself where the ECS task is running on is in an external subnet and the other tasks that do not have "subnet allowed" configured can access internet without any problems so security groups are fine.
## [4][Any update on AWS South Africa region general availability?](https://www.reddit.com/r/aws/comments/f8p47z/any_update_on_aws_south_africa_region_general/)
- url: https://www.reddit.com/r/aws/comments/f8p47z/any_update_on_aws_south_africa_region_general/
---
Hi Team,

Does anyone know when South Africa region is going online for general availability?  We're doing some things that having a POP doing in south Africa would help with.

Thanks,
Warren
## [5][Datasync and KMS?](https://www.reddit.com/r/aws/comments/f8p5zs/datasync_and_kms/)
- url: https://www.reddit.com/r/aws/comments/f8p5zs/datasync_and_kms/
---
Hi,

I'm currently reading up on DataSync for some future transfers to our AWS environment. One thing I find a bit contradictory are the following statements in the docs:

* DataSync supports using default encryption for S3 buckets using Amazon S3-Managed Encryption Keys (SSE-S3)
* DataSync has integrations with AWS KMS

So, what is it? Can you use buckets encrypted with your own KMS keys, or is only the SSE-S3 key supported?
## [6][AWS Secrets Manager Issue](https://www.reddit.com/r/aws/comments/f8bqou/aws_secrets_manager_issue/)
- url: https://www.reddit.com/r/aws/comments/f8bqou/aws_secrets_manager_issue/
---
I've created a secret in Secrets Manager and a custom lambda to rotate a bearer token I need to call some APIs.   


My issue is that sometimes... The rotation doesn't kick off at all. I have the rotation rules to automatically kick off every day (value set to 1). Am I missing something? Why would the rotation just not kick off some days?   


The lambda it invokes is within a VPC but I don't think that has anything to do with this but thought it might be worth mentioning. Whenever I kick off the rotation via the console everything works fine.  


I'm considering creating a cloudwatch event which will kick off the rotation (reinventing the wheel here) so I don't have to worry about this flaky behavior.

 
Response from AWS support (I'll continue to update the post as I hear from them): 

Thank you for contacting AWS Support, my name is Michael and I will be assisting you with this request.

I have gone through your CloudTrail Logs and can see the secret rotation triggered automatically on the 20th(01:07), 21st(08:08), 22nd(01:08) UTC time. On the 23rd I can see no automatic rotation and at 16:27 that day I can see that you manually triggered Rotate Secret from the Secrets Manager Console. I have attached the CloudTrail for each of these events.
I have also gone through the Lambda Function CloudTrail related API calls and could see no errors hinting at what could have caused Secrets Manager not to trigger the Lambda Rotation Function. Additionally, I could see no permission errors when the Lambda function was run. When invoked, the Lambda function was able to successfully rotate your secret.

To help me investigate further I have opened an Internal Ticket with the Secrets Manager Service Team to investigate why the Auto Rotation is not being triggered. While we wait for a response from the service team I will move this case into Pending Amazon Action and will update you as soon as the Service Team responds.
In the meantime, if you have additional questions please let me know.
## [7][Should EKS cluster be in the same subnet as other resources?](https://www.reddit.com/r/aws/comments/f8lvwf/should_eks_cluster_be_in_the_same_subnet_as_other/)
- url: https://www.reddit.com/r/aws/comments/f8lvwf/should_eks_cluster_be_in_the_same_subnet_as_other/
---
**EDIT:** Title should say same VPC

I used [eksctl](https://eksctl.io/) to create an EKS cluster. By default, it put the cluster into its own VPC and configured the subnets.

&amp;#x200B;

I have other resources in the same region on a different VPC that I would like my EKS cluster to have access to (Aurora, Redis, EFS, etc), but this is harder when they are not in the same VPC.

Is the correct way to handle this to put the EKS cluster in the existing VPC? The [documentation](https://eksctl.io/usage/vpc-networking/#use-existing-vpc-any-custom-configuration) for eksctl mentions that you can use an existing VPC, but then you need to create your own subnets and make sure they are configured correctly, which I think seems error prone (I wasn't even sure how to fill in the IPv4 CIDR blocks, let alone any tagging). Is there a better way to solve this, or maybe a reliable guide on how to create the subnets for the EKS cluster?
## [8][Using AWS services with OVH Object Storage](https://www.reddit.com/r/aws/comments/f8o61r/using_aws_services_with_ovh_object_storage/)
- url: https://www.reddit.com/r/aws/comments/f8o61r/using_aws_services_with_ovh_object_storage/
---
Hello !

For my current project, I need to use OVH Object Storage for their cheap bandwidth. However, I'd like to use AWS Amplify for the integrated mobile experience.  


My needs are :

* Users should be able to upload files only if they are connected (authentication is handled by AWS Cognito)
* They should be able to share files to other users, in a secure way

How am I supposed to handle the authentication flow? I looked through OVH and AWS documentation, but I didn't find any clue on how to do that. Is that even possible?
## [9][AWS + Kotlin : DynamoDB with superpowers](https://www.reddit.com/r/aws/comments/f8my8h/aws_kotlin_dynamodb_with_superpowers/)
- url: https://blog.yudiz.com/aws-kotlin-dynamodb-with-superpowers/
---

## [10][Anyone successfully deploying Rails 6 apps on Elastic Beanstalk?](https://www.reddit.com/r/aws/comments/f8j8v4/anyone_successfully_deploying_rails_6_apps_on/)
- url: https://www.reddit.com/r/aws/comments/f8j8v4/anyone_successfully_deploying_rails_6_apps_on/
---
I've been using Beanstalk for a side project for a few months and, to be honest, overlooking the fact that something isn't working quite right.

In my case, once the deploy is complete, the Puma logs show that /var/app/current/log/production.log isn't accessible. When I look at that directory, I can see the logfile... but it's set to root:root ownership. Doing a simple 'chown' to webapp:webapp makes it all work just fine.

So... here's my question: why isn't a standard Rails app deploy working right out of the box?

I can see where the symlinks are supposed to be set (via '/opt/elasticbeanstalk/hooks/appdeploy/post/01\_rails\_support.sh') but I'm not quite sure how to debug this...
