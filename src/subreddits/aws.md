# aws
## [1][Welcome to new mod u/_abhayshah!](https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/)
- url: https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/
---
Thrilled to expand the mod team in order better serve the community. As /r/aws grows, u/_abhayshah will be able to assist with AMAs, post flair, mod queue/mail, building out the Wiki, and more!   


Please give him a nice /r/aws welcome and let us know how we can improve things together going forward.
## [2][Introducing Distributed Load Testing Solution from AWS](https://www.reddit.com/r/aws/comments/j6ga4c/introducing_distributed_load_testing_solution/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/introducing-distributed-load-testing-v1-1/
---

## [3][AWS IAM role impersonation to HashiCorp Vault vulnerability](https://www.reddit.com/r/aws/comments/j6cmb0/aws_iam_role_impersonation_to_hashicorp_vault/)
- url: https://googleprojectzero.blogspot.com/2020/10/enter-the-vault-auth-issues-hashicorp-vault.html
---

## [4][Cloudwatch Event Pattern Event Rule Limit](https://www.reddit.com/r/aws/comments/j6qso5/cloudwatch_event_pattern_event_rule_limit/)
- url: https://www.reddit.com/r/aws/comments/j6qso5/cloudwatch_event_pattern_event_rule_limit/
---
We are trying to setup our cloudwatch to monitor 250+ servers. However, there are ec2 instances that are only started on-demand basis. This doesn't allow us to use "Any resource" as the it will include the instances that are intentionally stopped. Is there a way to exclude those stopped instances? Or is our only choice is to create multiple pattern event rules?
## [5][Revisiting CDK project](https://www.reddit.com/r/aws/comments/j6q06v/revisiting_cdk_project/)
- url: https://www.reddit.com/r/aws/comments/j6q06v/revisiting_cdk_project/
---
It's been 6 months since I last looked at my first CDK project that is now running in production. However, I have a slight problem.

We want to expand with a development setup that is identical to production. I made everything into reusable components back in March (with DTAP tags and everything). I haven't made any changes to the code, but when I run

    cdk diff

it shows that it wants to add a bunch of resources. I'm a bit scared it will completely fuck up my production environment or that it will create a new production environment. Does anyone have any tips or guidelines that could help me? Or is it better if I clone this repository and start fresh to make sure my Dev deployment won't accidentally mess with production?
## [6][What was your AHA moment using AWS?](https://www.reddit.com/r/aws/comments/j6f58g/what_was_your_aha_moment_using_aws/)
- url: https://www.reddit.com/r/aws/comments/j6f58g/what_was_your_aha_moment_using_aws/
---
Either the moment when you understood its full potential or when you knew that AWS was the right set of tools for your path?
## [7][What's the proper way to troubleshoot site-to-site VPN (GRE stage)?](https://www.reddit.com/r/aws/comments/j6pf36/whats_the_proper_way_to_troubleshoot_sitetosite/)
- url: https://www.reddit.com/r/aws/comments/j6pf36/whats_the_proper_way_to_troubleshoot_sitetosite/
---
Hello,

I'm configuring subj from Customer Gateway (CG) to Virtual Private Gateway (VPG).

Ipsec tunnel as it is starts normal, on VPG side both redundant tunnels are in state "IPSEC IS UP", from CG side (strongswan) - ipsec status also in state established.

But when i try to ping inner GRE address from CG to VPG - no luck, no reply packets. Tcpdump shows no activity from VPG side.  
No specific firewall rules on CG side. This machine keeps other ipsec tunnels with no problem, and new endpoints added as new items in ipset lists.  
What options do i have to troubleshoot the problem?  
Because aws troubleshooting guide has only one advise for this kind of situation "review your tunnel interface configuration to make sure that the proper IP address is configured" :)  
[https://docs.aws.amazon.com/vpn/latest/s2svpn/Generic\_Troubleshooting.html](https://docs.aws.amazon.com/vpn/latest/s2svpn/Generic_Troubleshooting.html)
## [8][ECS + Windows Containers](https://www.reddit.com/r/aws/comments/j6oz1y/ecs_windows_containers/)
- url: https://www.reddit.com/r/aws/comments/j6oz1y/ecs_windows_containers/
---
I have a bunch of Windows instances that aren’t important but need to be available for end users who need them, which is very infrequently but could be anytime of the day.

Considering containerising them and running the workloads inside ECS, does anyone have experience with a similar setup or other suggestions?

I use instance scheduler to schedule systems down when possible but these are low utilisation systems that need to be available 24x7.
## [9][Feature request: target S3 static site with an ALB](https://www.reddit.com/r/aws/comments/j6orzo/feature_request_target_s3_static_site_with_an_alb/)
- url: https://www.reddit.com/r/aws/comments/j6orzo/feature_request_target_s3_static_site_with_an_alb/
---
Hi everyone,

I want to preface the following by noting that this is not a support request.

I'm finding myself in the situation where I have an ALB with a path prefix that proxies to some Docker containers in Fargate. However the rest of the site is static so I would like to just make a static S3 site, offload TLS at the ALB, proxy what doesn't go to Fargate to the S3 site, and Bob would be your proverbial uncle.

I am surprised to find that that is not possible out of the box. I can set up an EC2 instance, or something like Fargate, or redirect to another site, but if I want to have:

- a single domain - one (1) domain,
- with an S3 site
- and an API endpoint of some kind in ECS
- and cheap or free TLS
- and I want to do all of that serverlessly,

...then I need to spin up CloudFront in addition to the ALB I need for the ECS stuff. For various reasons I prefer an ALB for this over CloudFront.

I know folks who work for AWS read this, so I'd like to submit the feature request that directly adding a static S3 bucket as a target to an ALB is something I for one would very much appreciate - as a matter of fact I actually totally expected it to be a thing, because it honestly seems like a super obvious feature to me. Of course, that may just be me because apparently it's not obvious to AWS. :)
## [10][How can I guarantee performance globally?](https://www.reddit.com/r/aws/comments/j6hhb4/how_can_i_guarantee_performance_globally/)
- url: https://www.reddit.com/r/aws/comments/j6hhb4/how_can_i_guarantee_performance_globally/
---
I've been given the task of making sure our Website/Mobile application performs well globally.  All our backend infrastructure runs in AWS (with Cloudflare sitting infront of everything). We are using Lambdas and API Gateway to run our api and backend. Our frontend application run in ECS. And the bulk of our content sits in S3 (\~5TB), with \~500gb in databases

There seems to be a suggestion from colleagues of duplicating infrastructure (e.g. RDS, DyanmoDB, Compute platforms) and running it in different global regions, so that requests can be processed closer to the user/client . And then try and sync data between the different geographical regions to keep things up to date. But to me this sounds like a headache, not to mention expensive. 

The easier option (and my more preferred option) would be to serve media content (e.g. images, large files) via a CDN, which can serve our content from different nodes around the world. The data processing would still be happening in the primary region (e.g. us-east-1) but large files would be served from a CDN located close to the clients location.

Would be great to get some recommendations.

Thanks
## [11][SQS - fifo queues and retention behaviour](https://www.reddit.com/r/aws/comments/j6nyxe/sqs_fifo_queues_and_retention_behaviour/)
- url: https://www.reddit.com/r/aws/comments/j6nyxe/sqs_fifo_queues_and_retention_behaviour/
---
Hi all,

We are looking at utilising SQS for our pipeline. 

I was hoping somebody would be able to answer a quick hypothetical question.

If we wanted to guarantee the order in which messages are processed using fifo queue type how does message retention period play a part in this process?

For example, we have a few messages come into the queue and we have the first message fail to be handled by the application properly and the message is not removed.

Will this result in the other messages waiting till the failed message reached the end of its retention time?

I have had a look online and the aws document and couldn’t find anything referencing this sort of behaviour. 

I could setup some tests to check this behaviour but I thought maybe someone who has had dealing with this before maybe able to advise to save me sometime.
