# aws
## [1][Launch virtual machine from aws console mobile app?](https://www.reddit.com/r/aws/comments/f1o6de/launch_virtual_machine_from_aws_console_mobile_app/)
- url: https://www.reddit.com/r/aws/comments/f1o6de/launch_virtual_machine_from_aws_console_mobile_app/
---
I installed the console app on my phone and thought I would be able to start an ec-2 instance. But afaict it is not possible? If I sign in with the same iam in a browser it seems possible.

Is the app only for monitoring? Looks like I can create security groups in the app.. but when I press instances there is no option for creating a new?

Thanks
## [2][Triggering AWS Lambda on a csv file upload](https://www.reddit.com/r/aws/comments/f1k3v3/triggering_aws_lambda_on_a_csv_file_upload/)
- url: https://www.reddit.com/r/aws/comments/f1k3v3/triggering_aws_lambda_on_a_csv_file_upload/
---
Hello,

I have some automation code that I want to move to aws lambda, the code uses a csv file to read the data and then execute routines based on them.

I am a bit new to whole AWS ecosystem, I was wondering if there exists some prebuilt stuff that can be useful for me?

What I need is a place where I can upload the csv file (drag and drop ideally, even non tech savvy should be able to do) something like google drive? Maybe

And then my lambda should be executed on upload. It should also be able to read the contents of the file.

The running should be done evertime I update the csv file or upload new one 🤔

Is there any simple way to achieve such functionality?

Thanks!
## [3][CI/CD tool for ECS on AWS](https://www.reddit.com/r/aws/comments/f1ozgx/cicd_tool_for_ecs_on_aws/)
- url: https://www.reddit.com/r/aws/comments/f1ozgx/cicd_tool_for_ecs_on_aws/
---
Hi, I'm looking for tool that could build our ECS container images and deploy them to ECS. I'd like to be able to move the same image from Dev to Prod environments. What are the tools that could help me(excluding jenkins)
## [4][MediaConvert and Teletext](https://www.reddit.com/r/aws/comments/f1p0kq/mediaconvert_and_teletext/)
- url: https://www.reddit.com/r/aws/comments/f1p0kq/mediaconvert_and_teletext/
---
Wondering if anyone has had any experience with using an STL file as an input then inserting this as a teletext captions output?
## [5][SSH agent forwarding vs SSM for shell access](https://www.reddit.com/r/aws/comments/f1efjf/ssh_agent_forwarding_vs_ssm_for_shell_access/)
- url: https://www.reddit.com/r/aws/comments/f1efjf/ssh_agent_forwarding_vs_ssm_for_shell_access/
---
What is the more secure method for accessing an EC2 instance in a private subnet?
## [6][Dissociate AWS account from Amazon (consumer) ?](https://www.reddit.com/r/aws/comments/f1fbuw/dissociate_aws_account_from_amazon_consumer/)
- url: https://www.reddit.com/r/aws/comments/f1fbuw/dissociate_aws_account_from_amazon_consumer/
---
Hello,

A few years ago, I've sign up on AWS using the same email / account that I use for Amazon.

I'm planning to use AWS more and wanted to dissociate both account.

- if I close my AWS account, will it close my Amazon account as well?  

In fact, I searched the forum and it wasn't possible in the past... Did it changes?  Does my only solution is to open a new AWS account with a different email?

Thanks
## [7][End-of-life announcement for CoreOS Container Linux](https://www.reddit.com/r/aws/comments/f1aq1y/endoflife_announcement_for_coreos_container_linux/)
- url: https://coreos.com/os/eol/
---

## [8][Send mongdb logs to cloudwatch's loggroup?](https://www.reddit.com/r/aws/comments/f1n8u1/send_mongdb_logs_to_cloudwatchs_loggroup/)
- url: https://www.reddit.com/r/aws/comments/f1n8u1/send_mongdb_logs_to_cloudwatchs_loggroup/
---
I have some mongodb instances in a primary/secondary model, how to push mongdb logs (like in /var/logs/mongodb/mongd.log) to a loggroup of cloudwatch?
## [9][AWS Tutoring/Advising](https://www.reddit.com/r/aws/comments/f1kbsw/aws_tutoringadvising/)
- url: https://www.reddit.com/r/aws/comments/f1kbsw/aws_tutoringadvising/
---
Hey, I’m a software engineer at a large Silicon Valley company, but I’m trying to learn some AWS concepts on the side for my own personal projects. I’m working on a server less REST API with I’m mostly just struggling with configuration and deployment with Cloudformation. I mostly just want an hour or two to talk to someone about best practices and to actually learn some of these concepts, and how to better manage resources, stacks, and my template. Would be paid time.
## [10][RDS and KMS Access - A follow up](https://www.reddit.com/r/aws/comments/f17a25/rds_and_kms_access_a_follow_up/)
- url: https://www.reddit.com/r/aws/comments/f17a25/rds_and_kms_access_a_follow_up/
---
I was recently struggling to understand why RDS had access to a CMK I created in KMS and came across [an old post on this sub](https://www.reddit.com/r/aws/comments/b8nrar/rds_and_kms_im_doing_something_wrong/) of someone asking the same question. 

The post is archived but I believe the comments are not correct and it sent me down the wrong path for a while.

The comment saying
&gt; Your $ACCOUNTID:root principal allows all principals in $ACCOUNTID to have those kms:* permissions...

is not true. From the [documentation](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam), it says:
&gt;  It does not by itself give any IAM users or roles access to the CMK, but it enables you to use IAM policies to do so

For my Aurora Serverless RDS, it was gaining access to the key from a grant. If you run `aws kms list-grants --key-id yourkey` you should see the grant which contains an encryption context with your DB instance ID. You should also see a CreateGrant entry in CloudTrail. The grant is created by the entity who creates the database.

I'm making this post in the hopes it saves someone from hours of head scratching in the future!
