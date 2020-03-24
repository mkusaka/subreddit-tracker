# aws
## [1][Covid-19 AWS Certification Update: You can now take all AWS Certification exams with online proctoring](https://www.reddit.com/r/aws/comments/fnri6t/covid19_aws_certification_update_you_can_now_take/)
- url: https://aws.amazon.com/certification/faqs/
---

## [2][How to hide credentials from aws-sdk in lambda when running untrusted code?](https://www.reddit.com/r/aws/comments/fo3mnt/how_to_hide_credentials_from_awssdk_in_lambda/)
- url: https://www.reddit.com/r/aws/comments/fo3mnt/how_to_hide_credentials_from_awssdk_in_lambda/
---
I am running Node.js sandbox as a service (part of bigger SaaS product). 

Currently, I am invoking a lambda handler which injects the untrusted code(passed from SQS as plaintext) into the excellent [VM2 sandbox](https://github.com/patriksimek/vm2). I would like to whitelist the `aws-sdk` module for the untrusted code. 

The problem I noticed is that when the untrusted code requires `aws-sdk` it is populated with AWS credentials of the lambda execution role (fortunately it can only publish to CloudWatch, I can live with that). Is there any way to hide/remove the credentials?
## [3][Adding Cloudwatch Alarms to EC2s via Terraform](https://www.reddit.com/r/aws/comments/fo2hhb/adding_cloudwatch_alarms_to_ec2s_via_terraform/)
- url: https://www.reddit.com/r/aws/comments/fo2hhb/adding_cloudwatch_alarms_to_ec2s_via_terraform/
---
The author wanted to add Cloudwatch Status Check alarms to his ec2 instances and used terraform to do this. Here’s how he did it: https://medium.com/kareemery_78433/adding-cloudwatch-alarms-to-ec2s-via-terraform-d19c006ad95a?source=friends_link&amp;sk=fe6d7e0ced10c7b2354ed71339f13c1f
## [4][ECS ComposeX](https://www.reddit.com/r/aws/comments/fo1fku/ecs_composex/)
- url: https://docs.ecs-composex.lambda-my-aws.io/
---

## [5][CloudFormation stack errors](https://www.reddit.com/r/aws/comments/fo21kf/cloudformation_stack_errors/)
- url: https://www.reddit.com/r/aws/comments/fo21kf/cloudformation_stack_errors/
---
I had a user yesterday whose stack was failing when trying to create a new VPC.  The underlying issue was that the account they were using was at their VPC limit.  I happened to have seen this situation before, so I was able to delete some unused dev VPCs, but the user had no indication of the underlying reasoning.  A few hours were lost troubleshooting where or how the code may be failing, checking bucket permissions, role permissions, etc.  It would've likely been several more hours if they hadn't reached out to me for help.  In these situations, I haven't been able to find a way to find a log or warning that tells the user why they failed to create resources.  Is there some way to find these sorts of errors either in the console or via CLI?
## [6][Does separate organization per client make sense?](https://www.reddit.com/r/aws/comments/fnle81/does_separate_organization_per_client_make_sense/)
- url: https://www.reddit.com/r/aws/comments/fnle81/does_separate_organization_per_client_make_sense/
---
Hello, I am potentially working on multiple projects with different clients and want to setup billing separately with their own credit card for each account. So for example, lets say I have 2 clients. Each with their own infra since they have diff systems and apps. What is the best setup for something like this? What I am leaning towards is "Project based account structure" as defined in this article:

 [https://aws.amazon.com/answers/account-management/aws-multi-account-billing-strategy/](https://aws.amazon.com/answers/account-management/aws-multi-account-billing-strategy/)

Does this make sense for my use case? Is there a better way? If I need to turn over the keys how would I do that?
## [7][Are all t2 EC2's free?](https://www.reddit.com/r/aws/comments/fo18fq/are_all_t2_ec2s_free/)
- url: https://www.reddit.com/r/aws/comments/fo18fq/are_all_t2_ec2s_free/
---
Today I launched the t2.2xlarge just for funzies since Ive only been using the free tier and need to upgrade.  I was expecting to get billed but haven't yet despite pricing being $0.148 per hour.  On the reserved instances pricing page, there were no t2 types.. Im a little confused.

I haven't been billed and I replaced my debit card, so Im not too worried
## [8][How to check website speed from different locations?](https://www.reddit.com/r/aws/comments/fo18do/how_to_check_website_speed_from_different/)
- url: https://www.reddit.com/r/aws/comments/fo18do/how_to_check_website_speed_from_different/
---
Which AWS Service can be used or how can I setup AWS to check the website speed from different locations?

We are looking at building some custom tool for our requirement which needs to check the website speed at certain intervals. A script that can check website speed giving details about TTFB, as I understand, that would be different for each location. So how to check and setup this?

Any help or reference to any article/video would be of great support.

Thanks.
## [9][Has there been any changes to how Performance Insights works for MySQL RDS?](https://www.reddit.com/r/aws/comments/fo0vk7/has_there_been_any_changes_to_how_performance/)
- url: https://www.reddit.com/r/aws/comments/fo0vk7/has_there_been_any_changes_to_how_performance/
---
I used to be able to see the full mysql queries in my Performance insights. Now I’m only seeing database admin operations unlike

Use database;
COMMIT
Etc..
## [10][Updating AWS Workspace Image](https://www.reddit.com/r/aws/comments/fnzzn8/updating_aws_workspace_image/)
- url: https://www.reddit.com/r/aws/comments/fnzzn8/updating_aws_workspace_image/
---
I've created a custom AWS Workspace Image and successfully created the follow on bundle and generated an AWS Workspace from it. The image contains some software downloaded to the C drive, and some reference files stored on the Desktop. In trying to update the image, I've created a new workspace from the base image, installed some additional software, and added new files to the reference folder on the desktop. I then go through and create a new image and bundle based on the updated Workspace. When I try to spin up a new workspace from this new bundle, the software is installed, but the changes to the desktop file do not persist. Am I missing a step in the process?

Thanks!
