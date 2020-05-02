# aws
## [1][[CloudFormation] - Using export and cross referencing stacks in templates](https://www.reddit.com/r/aws/comments/gc1rav/cloudformation_using_export_and_cross_referencing/)
- url: https://www.reddit.com/r/aws/comments/gc1rav/cloudformation_using_export_and_cross_referencing/
---
Greetings,

I'm in a project that requires us to build a serverless application in AWS. For this purpose, we have chosen SAM.

I made 2 resources of type \`\`\`Serverless\`\`\` within the SAM template, referencing other templates. That was fine until I had to configure networking and ElastiCache.  


**The problem**

I'm looking to cross reference templates **at deploy time for a new region bootstrapping**. I found that for me to use Export/Import of resources, the template (stack) that has to do the export must be created first in AWS. Both nested stacks can't be created at the same time, which made me do a dumb solution to comment out the other stack, deploy the networking, uncomment, re-deploy so it gets the exported values.  


I've read around the subreddit and people seem to dislike nested stacks? What would be a good approach to tackle this situation?  


I don't mind deploying each stack individually (Not sure if the exports will work between independent stacks tho) but It'd be nice to just reference a single template that references the rest.
## [2][How can I make sure that EC2 instances with dynamic IPs are available through stable subdomains?](https://www.reddit.com/r/aws/comments/gc552p/how_can_i_make_sure_that_ec2_instances_with/)
- url: https://www.reddit.com/r/aws/comments/gc552p/how_can_i_make_sure_that_ec2_instances_with/
---
Hello!

Imagine I have a host of EC2 instances with dynamic IPs. Every day, at the end of office hours, most of them are turned off. In the morning of every business day, they are turned on again. This means that their IP addresses change every day.

I want to have subdomains which do not change and always point to the right EC2 instance. For example, in the image below `inst-1.mydomain.com` should always point to `EC2 instance 1`, regardless of its IP address.

[Here](https://i.imgur.com/UTxXeS1.png) is a sketch showing what I mean.

These domains must work for all ports, including, but not limited to SSH, HTTP, HTTPS.

Buying static IPs from Amazon is not an option.

**How can I implement it in the best way?**

One obvious solution is this: Take the script that turns the instances on and off again, and modify it so that it updates the subdomain DNS records. That is, whenever the instances are started, the script would take their new IP address and then update the DNS record of `inst-1.mydomain.com` so that it points to the new IP address. Some providers allow you to update DNS via an API (here is an [example](https://help.dreamhost.com/hc/en-us/articles/217555707-DNS-API-commands)).

Are there other, better ways to achieve the same result?

Thanks in advance
## [3][[Blog post] Tracking Amazonians (AWS) on Twitter](https://www.reddit.com/r/aws/comments/gc3pfl/blog_post_tracking_amazonians_aws_on_twitter/)
- url: https://zoph.me/posts/2020-04-19-tracking-amazonian-on-twitter/
---

## [4][Secret Manager - RDS Password Rotation](https://www.reddit.com/r/aws/comments/gbrorb/secret_manager_rds_password_rotation/)
- url: https://www.reddit.com/r/aws/comments/gbrorb/secret_manager_rds_password_rotation/
---
Good evening, 

I have "stored" the master password for a Postgres RDS instance in Secret Manager. I know it is working correctly as I can access the secret from an EC2 instance to connect to the database. I have tried enabling the rotate secret feature, but it does not seem to be working. It created a lambda but I cannot find a way to look at the logs to see what went wrong. When I click "Rotate Secret Immediately", it says: "Fail to rotate the secret "master\_password\_prod" A  previous rotation isn't complete. That rotation will be reattempted." It doesn't matter how long I wait, it never succeeds. 

Any advice would be appreciated :)
## [5][Terminate idle EC2 instances which have a specific tag?](https://www.reddit.com/r/aws/comments/gc2atq/terminate_idle_ec2_instances_which_have_a/)
- url: https://www.reddit.com/r/aws/comments/gc2atq/terminate_idle_ec2_instances_which_have_a/
---
I want to do something like the one described here https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/UsingAlarmActions.html but the instances are not in an ASG. Instead they'll have a common tag.
## [6][Unable to login to console on Chrome](https://www.reddit.com/r/aws/comments/gbrnif/unable_to_login_to_console_on_chrome/)
- url: https://www.reddit.com/r/aws/comments/gbrnif/unable_to_login_to_console_on_chrome/
---
Anyone else have this issue? I can use safari and Firefox but using chrome it days “bad request”.

I presume I must’ve some bad cookies or something
## [7][Serverless Thought Leaders Converge Online on May 21 &amp; May 28](https://www.reddit.com/r/aws/comments/gbuhaj/serverless_thought_leaders_converge_online_on_may/)
- url: https://www.reddit.com/r/aws/comments/gbuhaj/serverless_thought_leaders_converge_online_on_may/
---
Hi, r/aws!

The first-ever AWS Serverless-First Function is a free, virtual event intended to help you learn the latest about serverless approaches on AWS. The event occurs over two Thursdays, May 21 and May 28, and each day focuses on specific topics.

* May 21: Learn about adopting serverless across your organization
* May 28: Learn about building serverless into your applications

Live chat is included on both days – so bring all of your questions to be answered by serverless experts!

You can check out the full session agenda and speaker and event details on the info page [here](https://pages.awscloud.com/GLOBAL-event-OE-serverless-first-function-2020-reg-event.html).
## [8][Anyone having acquired Outposts, and willing to share some feedback on usage, so far?](https://www.reddit.com/r/aws/comments/gbh5aj/anyone_having_acquired_outposts_and_willing_to/)
- url: https://www.reddit.com/r/aws/comments/gbh5aj/anyone_having_acquired_outposts_and_willing_to/
---
Among things of interest:

- reason for acquisition &amp; expectations being met?!?

- performance compared to the region, both vis-a-vis on-prem

- LGW - w/CoIP, or else?!?

- services on-prem delivered for workloads in Outposts (e.g. are you load balancing with on-prem devices like Citrix, F5, A10?)

- ... anything else that comes to mind ...

**Edit#2**: ~~the amount of reminders and~~ quite a few downvotes (AWS bots?!?), with no true feedback to the points I've asked about, tells me something, already ... + just marketing and re:invent re:peats. Isn't anyone having one of these in production???
## [9][Healthcheck for UDP app in ECS using dynamic port routing](https://www.reddit.com/r/aws/comments/gc0lyj/healthcheck_for_udp_app_in_ecs_using_dynamic_port/)
- url: https://www.reddit.com/r/aws/comments/gc0lyj/healthcheck_for_udp_app_in_ecs_using_dynamic_port/
---
So I am working on a udp application running in docker. I created an ECS service and am able to connect to it, but it quickly fail it's healthcheck and be terminated by the Network Load Balancer.

I created an HTTP endpoint on port 80 for the healthcheck and can confirm it runs properly in the image (via hitting when running locally).

Things I've tried:

1. Manually set the healthcheck port, it return false due to using dynamic routing
2. Setting the protocol to just UDP and removing the tcp port mapping. Still failed.

I can work around this by not using dynamic routing, but I would really prefer to have it.

Currently, the target group is just setup as following:

`NlbTargetGroup:`  
 `Type: AWS::ElasticLoadBalancingV2::TargetGroup`  
 `Properties:`  
 `Port: 7777`  
 `Protocol: TCP_UDP`  
 `TargetGroupAttributes:`  
`- Key: deregistration_delay.timeout_seconds`  
 `Value: '20'`  
 `VpcId: !Ref 'VpcId'`

The task looks like this:

`TaskDefinition:`  
 `Type: AWS::ECS::TaskDefinition`  
 `Properties:`  
   `ContainerDefinitions:`  
   `- Name: main`  
`Essential: 'true'`  
`Image: !Ref DockerImage`  
`MemoryReservation: 250`  
`PortMappings:`  
`- HostPort: 0`  
`ContainerPort: 7777`  
`Protocol: udp`  
`- HostPort: 0`  
`ContainerPort: 80`  
`Protocol: tcp`  
   `LogConfiguration:`  
`LogDriver: awslogs`  
`Options:`  
`awslogs-region:`  
`Ref: AWS::Region`  
`awslogs-group:`  
`Ref: LogGroup`

Thanks for your time, I appreciate it!
## [10][Cfn import Route43::RecordSet records](https://www.reddit.com/r/aws/comments/gbygn6/cfn_import_route43recordset_records/)
- url: https://www.reddit.com/r/aws/comments/gbygn6/cfn_import_route43recordset_records/
---
Hello,

I have a template that is comprised of a hosted zone and many record sets. The template/stackname was created with the wrong name.

What’s the best way to fix this without changing name servers away from Route53? All of the RecordSets have DependsOn hosted zone.

I was thinking this, but it sounds risky.

1.	Add DeletionPolicy of Retain to HostedZone and all record sets.
2.	Remove DependsOn from all resources in template.
3.	Delete hosted zone record from the template.
4.	Create new hosted zone record in correctly named template.
5.	Change TTL for all record sets in original template to 60 seconds.
6.	Delete records during off hours.
7.	Run new template with the records readded.

I’m not really a fan of doing this or using another DNS provider to migrate away and back.

Support for one or both of the following would be amazing.

1.	Support to rename cfn stack
2.	Cfn import resource Route53::RecordSet

Thanks!
