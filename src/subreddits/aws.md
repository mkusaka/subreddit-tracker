# aws
## [1][S3 Path Deprecation Plan Updated](https://www.reddit.com/r/aws/comments/iyd84u/s3_path_deprecation_plan_updated/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [2][We are the AWS EC2 Team - Ask the Experts - Sep 24th @ 9AM PT / 12PM ET / 4PM GMT!](https://www.reddit.com/r/aws/comments/iu0c8d/we_are_the_aws_ec2_team_ask_the_experts_sep_24th/)
- url: https://www.reddit.com/r/aws/comments/iu0c8d/we_are_the_aws_ec2_team_ask_the_experts_sep_24th/
---
Hey r/aws! u/AmazonWebServices here.

The AWS EC2 team will be hosting an Ask the Experts session here **in this thread** to answer any questions you may have about **deploying your machine learning models to Amazon EC2 Inf1 instances powered by the** [**AWS Inferentia chip**](https://aws.amazon.com/machine-learning/inferentia/), which is custom designed by AWS to provide high performance and cost-effective machine learning inference in the cloud. These instances provide up to 30% higher throughput, and 45% lower cost per inference over comparable GPU-based instances for a wide variety of machine learning use cases such as image and video analysis, conversational agents, fraud detection, financial forecasting, healthcare automation, recommendation engines, text analytics, and transcription. It's easy to get started and popular frameworks such as TensorFlow, PyTorch, and MXNet are supported.

Already have questions? Post them below and we'll answer them starting at 9AM PT on Sep 24, 2020!
## [3][Security issue in CloudFormation was leaking role session credentials for customer accounts](https://www.reddit.com/r/aws/comments/iyr4yr/security_issue_in_cloudformation_was_leaking_role/)
- url: https://onecloudplease.com/blog/security-september-cataclysms-in-the-cloud-formations?reup
---

## [4][Can AWS bring back the pinned favorites option in the navigation bar?](https://www.reddit.com/r/aws/comments/iyree9/can_aws_bring_back_the_pinned_favorites_option_in/)
- url: https://www.reddit.com/r/aws/comments/iyree9/can_aws_bring_back_the_pinned_favorites_option_in/
---
It was a good feature. Now there's a lot of wasted space in the nav bar :(
## [5][IAM Console us-east-1 not working?](https://www.reddit.com/r/aws/comments/iywyxm/iam_console_useast1_not_working/)
- url: https://www.reddit.com/r/aws/comments/iywyxm/iam_console_useast1_not_working/
---
Is anyone else having problems bringing up IAM in us-east-1?
## [6][WAF rules](https://www.reddit.com/r/aws/comments/iywh6s/waf_rules/)
- url: https://www.reddit.com/r/aws/comments/iywh6s/waf_rules/
---
Hey Guys,

I've recently been upgraded from a junior system administrator to a cloud engineer and in need of some help.

&amp;#x200B;

So I need to implement some WAF rules to better protect against DDoS, XSS and the usual fun you get on the web, but I am struggling to find some good resources on some 'Good starter rules'

Suffice to say I am slightly in over my head

Does anyone have a good document or suggestions I should look in to?

Thanks,

\-Chin
## [7][Why AWS DocumentDB charges hourly for instances and DynamoDB only charges for read, write and storage?](https://www.reddit.com/r/aws/comments/iytkxq/why_aws_documentdb_charges_hourly_for_instances/)
- url: https://www.reddit.com/r/aws/comments/iytkxq/why_aws_documentdb_charges_hourly_for_instances/
---
I'm new to AWS and little confused on Dynamo and Document pricing?


Thanks in advance
## [8][Sharing Access to Elastic Beanstalk Application](https://www.reddit.com/r/aws/comments/iyw81b/sharing_access_to_elastic_beanstalk_application/)
- url: https://www.reddit.com/r/aws/comments/iyw81b/sharing_access_to_elastic_beanstalk_application/
---
Sorry if this a stupid question, but I'm not sure how to do it. I'm working with a group and have used Elastic Beanstalk to create an web application. Now I would like to give everyone else in the group the ability to go in and edit the environment's configuration. How do I do this?
## [9][Setting Up a Redundant VPN Tunnel Between AWS &amp; CheckPoint](https://www.reddit.com/r/aws/comments/iyw5hu/setting_up_a_redundant_vpn_tunnel_between_aws/)
- url: https://www.reddit.com/r/aws/comments/iyw5hu/setting_up_a_redundant_vpn_tunnel_between_aws/
---
Hi all. I've been asked if it's possible to setup a redundant VPN tunnel between AWS and our office. I'm not familiar with AWS so have been looking at documentation and from what I've read, AWS VPCs can be setup with 2 gateways. On the CheckPoint side, these AWS gateways can be put in a star VPN community with the on-prem gateways.

What I would like confirming is, if our on-prem CheckPoint fails over to our backup line, would the tunnel to AWS still stay up? Is there a way of defining a primary and a backup IP for the CheckPoint side on AWS?
## [10][Connecting to Oracle DB using Python via Lambda](https://www.reddit.com/r/aws/comments/iyuxc5/connecting_to_oracle_db_using_python_via_lambda/)
- url: https://www.reddit.com/r/aws/comments/iyuxc5/connecting_to_oracle_db_using_python_via_lambda/
---
I am trying to connect to Oracle RDS from my python code using lamda handler but the due to huge size of  my code owing mainly to Oracle library that's coming upto 250 MB, I am unable to deploy it. 
1) I tried using layers for libraries but they also have a size limitation.
2) zipping the files is also not helping. It's still huge.

Is there any way to make it work? A different library I can use? Please let me know.
## [11][[EFS] Trouble mounting EFS Access Point to ECS Volume](https://www.reddit.com/r/aws/comments/iyux0b/efs_trouble_mounting_efs_access_point_to_ecs/)
- url: https://www.reddit.com/r/aws/comments/iyux0b/efs_trouble_mounting_efs_access_point_to_ecs/
---
Hi guys,

I ran into a problem trying to mount an ECS Volume to EFS through an EFS access point. 

The task role is set up with ClientWrite, ClientRead, and ClientRootAccess to that file system.

The access point is setup with posix userid 1001 and groupid 1001 with permission 755.

The cluster and the file system are in the correct VPC. 

But ECS failed to spin up a task with this error:

Error response from daemon: failed to copy file info for /var/lib/ecs/volumes/{{task-name}}: failed to chown /var/lib/ecs/volumes/{{task-name}}

I was able to spin up the task if i set the access point's POSIX userid and groupid  to 0 as in root. But i feel like its not the best choice for security reason in a shared FS.
## [12][Trouble launching g4dn instances in US-East-2.](https://www.reddit.com/r/aws/comments/iyp0m2/trouble_launching_g4dn_instances_in_useast2/)
- url: https://www.reddit.com/r/aws/comments/iyp0m2/trouble_launching_g4dn_instances_in_useast2/
---
So at the moment I am trying to launch an instance to customize an AMI. It's a g4dn.xlarge in particular (or 2xlarge) so that I can install Nvidia drivers and the Nvidia container runtime. Most of today I haven't been able to launch the instance (US-East-2). I receive a "We currently do not have sufficient g4dn.xlarge capacity in the Availability Zone you requested (us-east-2a)."

It doesn't matter which availability zone I choose as it goes, and I can come back later and it will eventually work.

My actual question revolves around elastic beanstalk. I have an app that uses this same instance type and I run into this same error when auto scaling kicks in, except it cycles through all the availability zones and tries both spot and on demand (to no avail). How do I account for this? Do I need to make the app in another region, and if so, what happens when this occurs there at some point in the future? Scaling up through the instance types probably works, but the lack of extra GPUs (until the g4dn.12xlarge) means I'm not scaling anything except my budget.

Thoughts?
