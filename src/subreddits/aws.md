# aws
## [1][API gateway scenario](https://www.reddit.com/r/aws/comments/g28pe6/api_gateway_scenario/)
- url: https://www.reddit.com/r/aws/comments/g28pe6/api_gateway_scenario/
---
Hi all, my team is currently trying to figure out a design to satisfy our requirements.

At the moment we have a web service hosted in IIS that our clients currently connect directly to and send data.  

We have a few problems that we would like to resolve and we think API gateway would help

1) we would like to log all incoming requests and the response code

2) currently if our service is not operating / cannot activate / IIS or the server is down we can’t detect it without adding additional monitors, having API gateway we’d like to send cloud watch events 


Our original plan was to have an API gateway with a lambda backing it, the lambda would store the incoming request (maybe S3 or Dynamo), forward the call through to the actual web service and then pass the response back.

Of note is most of the incoming data packets are quite large, usually between 1-5mb which was a limiting factor for the internal logging support.

A key consideration here is we wanted a completely severless and highly available receiving gateway for the requests.

Is this the right way to go about it?
## [2][Optimizing ECS resource allocation](https://www.reddit.com/r/aws/comments/g27vf0/optimizing_ecs_resource_allocation/)
- url: https://www.reddit.com/r/aws/comments/g27vf0/optimizing_ecs_resource_allocation/
---
Suppose you have a service on ECS with significantly low cpu utilization and/or memory utilization. What do you look at to decide between lowering the minimum task count or lowering the cpu and/or memory allocation for the container?

In other words, if CloudWatch is telling me my ECS service is over provisioned , how do I know when it is best to decrease min tasks vs when it is best to decrease resources per container?
## [3][Don't Forget to Set Budgets Monitoring in Cost Explorer! How to Detect AWS Cost Anomalies Before They Spiral Out of Control.](https://www.reddit.com/r/aws/comments/g1vaql/dont_forget_to_set_budgets_monitoring_in_cost/)
- url: https://www.cloudforecast.io/blog/how-to-detect-aws-cost-anomalies-before-they-spiral-out-of-control/
---

## [4][Deploying a Flask App with Elastic Beanstalk Gives ModuleNotFoundError](https://www.reddit.com/r/aws/comments/g2abdg/deploying_a_flask_app_with_elastic_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/g2abdg/deploying_a_flask_app_with_elastic_beanstalk/
---
This is my first time making a Flask App with Elastic Beanstalk so I probably have something configured wrong. [This](https://gist.github.com/jaredgoodman03/cbf54ae38120b21142a573b0ffe4b3d4) is my error log. I have application.py in my directory, and it works on my local machine.
## [5][Use Reserved Instances within an Auto Scaling Group?](https://www.reddit.com/r/aws/comments/g2arog/use_reserved_instances_within_an_auto_scaling/)
- url: https://www.reddit.com/r/aws/comments/g2arog/use_reserved_instances_within_an_auto_scaling/
---
Is it possible to utilise RIs within an ASG?  There doesn't seem to be an option to reference them directly but does it "just work" if the instance types of your RIs and launch template match up?

Example scenario: ASG with a minimum of 3 instances (ideally all reserved) split across 3 AZs.
## [6][Data Transfer Out (DTO) 40% Price Reduction in South America (São Paulo) Region](https://www.reddit.com/r/aws/comments/g2esea/data_transfer_out_dto_40_price_reduction_in_south/)
- url: https://aws.amazon.com/blogs/aws/aws-data-transfer-out-dto-40-price-reduction-in-south-america-sao-paulo-region/
---

## [7][Getting started on AWS? A Beginner Roadmap](https://www.reddit.com/r/aws/comments/g2eqyy/getting_started_on_aws_a_beginner_roadmap/)
- url: https://www.reddit.com/r/aws/comments/g2eqyy/getting_started_on_aws_a_beginner_roadmap/
---
Hi everyone,

I've been noticing quite a few topics lately about folks looking to get started on AWS, but not knowing where to start. I decided to put together a video outlining a beginner roadmap for those of you just beginning your journey on AWS.

The video is available here: https://youtu.be/lTyqzyk86f8

Hopefully some newer folks find this helpful. I'm curious to hear anyone else's thoughts regarding resources for getting started/better on AWS.
## [8][IAM policy for managing network ACL](https://www.reddit.com/r/aws/comments/g29b5p/iam_policy_for_managing_network_acl/)
- url: https://www.reddit.com/r/aws/comments/g29b5p/iam_policy_for_managing_network_acl/
---
I'm trying to create a policy to allow a user to manage network ACLs.  I have:

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "ManageNetworkAcl",
                "Effect": "Allow",
                "Action": [
                    "ec2:CreateNetworkAclEntry",
                    "ec2:ReplaceNetworkAclEntry",
                    "ec2:DeleteNetworkAclEntry"
                ],
                "Resource": "*"
            }
        ]
    }

I would like to set the condition to match on a resource Project tag, e.g.

         "Condition": {
                "StringEquals": {
                      "ec2:ResourceTag/Project": "my-project"
                }
         }

But when I run the policy in the IAM Simulator, I get access denied?

How I can set a condition to match on a tag, or object id such as nacl id, or vpc id?
## [9][Glacier Vault or S3 Bucker Glacier Class](https://www.reddit.com/r/aws/comments/g2cos5/glacier_vault_or_s3_bucker_glacier_class/)
- url: https://www.reddit.com/r/aws/comments/g2cos5/glacier_vault_or_s3_bucker_glacier_class/
---
I'm really trying to understand what the difference between uploading to Glacier vault and uploading to S3 bucket with Glacier class. Obiously the later is prefer because I can see the files actually in the S3 buckets.

I know that transition of object from S3 standard to Glacier cost per-object, but what if I upload directly to Glacier ?

 aws s3 cp test.mov s3://test-videos-raw/ --storage-class DEEP\_ARCHIVE
## [10][I am charged ~$60K on AWS, without using anything](https://www.reddit.com/r/aws/comments/g1ve18/i_am_charged_60k_on_aws_without_using_anything/)
- url: https://www.reddit.com/r/aws/comments/g1ve18/i_am_charged_60k_on_aws_without_using_anything/
---
So here is what's going on.

I am web developer and my employer gave me a task one day. It was "Create reductant setup of a \*website\*".

So at first glance I don't have a clue and start reading comments. They were debating whether they should pay higher to a AWS guy to do it or just leave one of the guys research and do it. So they end up giving the task to me.

Long story short, I end up on a page about reductant setup with amazon AWS RDS. I go to AWS, follow the instructions briefly to see what happens. After an hour or so, I got switched to a higher prio task and totally forgot about this, UNTIL TODAY.

I open my email and see bunch of emails up to 3 months prior, stating that they could not c bill my card, with the amount of \~$5,000. I was "WTF is this joke" and closed the email. Deleted all from AWS, threatening to terminate my account. (Edit: After acknowledging they were not scam, I restored them on the SAME day)

After a while(Edit: 3-4hrs) I opened the deleted mails and they were even stating I owe $32,000 ... WTF...

For this month I have \~$24k and I don't even know how to stop this service! I wrote to the support and hope they do something in order to help me, because $60k is not something I will be able to pay EVER.

Have you guys experience something like this, I am very very concerned about my well being right now..

TL;DR;

Got charged \~$60,000 by AWS for a test task I worked on at my job 3 months ago.  


Edit: I am going to throw some clarifications, as I might have mislead many people with some of my words above.   
\- I was not ignoring AWS email and deleting them for months.  
\- Saying I deleted emails, only meant to express my disbelief for the mails   
\- I contacted AWS on the same day (something like 3 hours after I read the first one). I logged into the console and created a case  
\- I am not ranting against AWS, I just want to explain clearly and sincerely all my actions, as I believe it will help throw better light on this story.
