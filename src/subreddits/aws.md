# aws
## [1][Week of Aug 3rd - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/i2ydcv/week_of_aug_3rd_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/i2ydcv/week_of_aug_3rd_what_are_you_building_this_week/
---
Discuss your challenges and successes
## [2][Quick Exercise - Amazon AppFlow to transfer data from Salesforce to S3](https://www.reddit.com/r/aws/comments/i40w3q/quick_exercise_amazon_appflow_to_transfer_data/)
- url: http://aws-dojo.com/excercises/excercise5
---

## [3][AWS Identity and Access Management (IAM) Reading and Resource List](https://www.reddit.com/r/aws/comments/i413i3/aws_identity_and_access_management_iam_reading/)
- url: https://github.com/straithe/ReadingAndResourcesLists/blob/master/AWS_IAM.md
---

## [4][AWS Monitoring: Dashboard/Metrics/Alerts/NOC Style - Options/Opinions?](https://www.reddit.com/r/aws/comments/i3ywua/aws_monitoring_dashboardmetricsalertsnoc_style/)
- url: https://www.reddit.com/r/aws/comments/i3ywua/aws_monitoring_dashboardmetricsalertsnoc_style/
---
So I guess there is a few options and a bit of a love/hate relationship amongst them.

I've been tasked with establishing a complete and all emcompasing network/systems/infrastructure monitoring solution for our businesses. We've purchased a few companies in the last few years, and have scattered things across a few different AWS accounts, and even one Azure account right now. 

Ideally I'd like to be able to have a single monitoring solution that can track/manage all of this; Both at an application level (website up/available, etc.) and infrastructure level (disk space, ram usage, CPU usage, etc.) 

The business owns a license (and not cheap!) for OpManager &amp; AppManager tools. The managers woujld ideally like to use this, but I'm finding it seems like pretty good product for on-prem/DC monitoring, but not really cloud-based. While the 2 parts SSO between each other, there is no other integration. I can hook up some levels of EC2 data into AppManager, but then not see it in opManager. This leaves me without a "single dashboard" solution. The boss's want to put up 50" screens in each of the IT based offices and display this output, so having multiple things doing different parts is a mess. 

We also need alerts/alarms and escalation paths based on these as well.

Personally I've seen DataDog and Site24x7, based off past-year exposure to AWS Summits, and some limited messing around, but I'm curious what do others do/use? Do you like the tools you have?
## [5][Applying DISA STIG GPOs to AWS Workspaces?](https://www.reddit.com/r/aws/comments/i42mj5/applying_disa_stig_gpos_to_aws_workspaces/)
- url: https://www.reddit.com/r/aws/comments/i42mj5/applying_disa_stig_gpos_to_aws_workspaces/
---
Has anyone managed to apply the DISA STIG GPOs to Workspaces Server 2016 (Windows 10) ?

I have been struggling to get it working. There seems to be multiple settings that are breaking the internal PCOIP mechanisms and no documentation. 

The logon banner needs to be disabled, as that is a default setting in the provided GPOs, but what else is incompatible and would break it? When applied, you can still RDP in, but can't access via the client. I have read all documentation and made sure it's configured correctly.

I don't want to reinvent the wheel here. If anyone has gone through this please help me out!
## [6][How to calculate the resources required by my spring boot application(REST API) to run on ec2?](https://www.reddit.com/r/aws/comments/i44nl4/how_to_calculate_the_resources_required_by_my/)
- url: https://www.reddit.com/r/aws/comments/i44nl4/how_to_calculate_the_resources_required_by_my/
---
I developed my private REST API using Spring-boot and planning to deploy it on the AWS ec2. But having no experience in this, I have no idea how to calculate resources required by my boot application and then decide which instance would I require. In general, how should one calculate the resources required by a java application? My REST API talks to the database and returns the JSON response. The database is MySQL and I am using AWS managed RDS MySQL instance.

On what factors the required resources are calculated? The number of API calls? I am assuming(on my website which uses this API) per page 5 API calls and per day 1000 visitors.
## [7][Reason not to use Fargate + public IPs?](https://www.reddit.com/r/aws/comments/i3uslj/reason_not_to_use_fargate_public_ips/)
- url: https://www.reddit.com/r/aws/comments/i3uslj/reason_not_to_use_fargate_public_ips/
---
Is there any reason not to use Fargate with public IPs? You have to lock container access down, but you have to do that anyway. It makes things marginally easier, because you don't have to add NAT and the routes to the VPC/subnets. But really, are there any downsides besides the obvious security aspects?
## [8][Why You Should Never, Ever print() in a Lambda Function](https://www.reddit.com/r/aws/comments/i3h6rn/why_you_should_never_ever_print_in_a_lambda/)
- url: https://medium.com/@psingman/why-you-should-never-ever-print-in-a-lambda-function-f997d684a705
---

## [9][Multi-Region S3 Buckets for Lower Latency on Cloudfront](https://www.reddit.com/r/aws/comments/i3p271/multiregion_s3_buckets_for_lower_latency_on/)
- url: https://www.reddit.com/r/aws/comments/i3p271/multiregion_s3_buckets_for_lower_latency_on/
---
The AWS Blog has a post, [Using Amazon CloudFront with Multi-Region Amazon S3 Origins](https://aws.amazon.com/blogs/apn/using-amazon-cloudfront-with-multi-region-amazon-s3-origins/), that outlines how to use a Lambda@Edge function to route Cloudfront requests to an S3 Bucket based on DNS TTL. I have several questions related to this approach:

&amp;#x200B;

1. Why not use Route53 [latency-based routing](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency) instead of a Lambda@Edge function?
2. Can impacts on origin latency be estimated?
   1. Specifically, transitioning from a single region (us-east-1) to double region (us-east-1 &amp; ap-south-1) configuration?
3. What other approaches can reduce origin latency?

For additional background on my use case, I would use [cross-region replication in S3](https://aws.amazon.com/blogs/aws/new-cross-region-replication-for-amazon-s3/) to replicate my bucket from us-east-1 to ap-south-1. My goal is to reduce origin latency, not necessarily to vary content by geography.
## [10][ALT key sometimes stops working while Workspace running in background](https://www.reddit.com/r/aws/comments/i40sd3/alt_key_sometimes_stops_working_while_workspace/)
- url: https://www.reddit.com/r/aws/comments/i40sd3/alt_key_sometimes_stops_working_while_workspace/
---
This is exceptionally annoying.  For some reason numerous times a day with my Workspace running int he background, my ALT key stops working on my primary desktop environment (not my Workspace).  This prevents me from doing things like ALT-TABing.  It seems to happen sometime after the Workspace locks.

&amp;#x200B;

The only way I know to fix it is to unlock my Workspace and then I can switch back to my primary desktop and it will work for some period of time.
## [11][testing Workmail with the free domain. Is DMARC/MAIL FROM possible?](https://www.reddit.com/r/aws/comments/i3xy19/testing_workmail_with_the_free_domain_is/)
- url: https://www.reddit.com/r/aws/comments/i3xy19/testing_workmail_with_the_free_domain_is/
---
here's something i have been unable to do and i'm suspecting it's not possible with the free domain. So i know that if i add my own domain to workmail, i need to add some records on the DNS resolver to validate the domain and add extra functionality, like dmarc / mail from.

So i sent me an email from the free mydomain.awsapps.com domain and i don't see the DMARC header. I go to the SES console to set mail from so that amazonses.com is not used, thinking this might setup DMARC as well. However the documentation is not clear.

When workmail created and verified the free domain, i see that route 53 was not setup, which is my first suspicion that suggests that you can't create DNS records with the free domain option. No matter what i do, the mail from option can't be enabled because it gets stuck in pending. I tried to create the hosting zone on R53 and add the required MX record to the subdomain, but i can never see  it with nslookup despite seeing it correctly in the R53 console.

Has anyone been able to configure mail from and DMARC on free workmail domains?
