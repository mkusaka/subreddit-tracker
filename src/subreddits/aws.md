# aws
## [1][Reasons to use EKS over ECS (2020 Oct) ?](https://www.reddit.com/r/aws/comments/j5daf2/reasons_to_use_eks_over_ecs_2020_oct/)
- url: https://www.reddit.com/r/aws/comments/j5daf2/reasons_to_use_eks_over_ecs_2020_oct/
---
Is there any reason to use EKS over ECS if:

1. We are not planning to manage EC2; means we only want to use Fargate
2. We are fully invested on AWS. No plans to migrate to other cloud or on premise
3. We have enough AWS skilled people in general (not specifically ECS), but have fewer K8s skilled people (but possible to hire or train)
4. The applications we are planning to deploy are regional. Will only be used by users in the region. We will only deploy in one AWS region. The peak usage could be around 10k concurrent sessions. Will go down to few hundred users in non peak periods.
5. We need to have full control on the VPC and networking in general.

6. We will be using terraform exclusively for the deployment.

7. Our CICD pipeline will be using gitlab.
## [2][AWS EventBridge Rule complains that target is not a valid JSON format](https://www.reddit.com/r/aws/comments/j5g80w/aws_eventbridge_rule_complains_that_target_is_not/)
- url: https://www.reddit.com/r/aws/comments/j5g80w/aws_eventbridge_rule_complains_that_target_is_not/
---
hello, I am following AWS tutorial &amp; trying out input transformation on rule.  
[https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-input-transformer-tutorial.html](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-input-transformer-tutorial.html)

I am not changing anything &amp; copying straight from the tutorial but when I click on create to create the rule, I get `Input for target Idbda77dd1-bec3-4408-9e1f-9d154933b733 is not a valid JSON text.` 

Am I missing something? Is there a working example out there that I can use, maybe console or IaC?

&amp;#x200B;

https://preview.redd.it/363gw5dst8r51.png?width=774&amp;format=png&amp;auto=webp&amp;s=18dc3b49b66bfc2db62598df195361c88dcb20a2
## [3][How to properly use security group to manage accesses ?](https://www.reddit.com/r/aws/comments/j5ghkc/how_to_properly_use_security_group_to_manage/)
- url: https://www.reddit.com/r/aws/comments/j5ghkc/how_to_properly_use_security_group_to_manage/
---
Hi there

TL;DR : I've been pondering a change in the way I use security groups. From "SG allow this subnet here" to "allow this (ASG A) SG to this (other ASG B) SG". Good idea, bad idea, how do you do it ? 

Currently I've VPCs with a bunch of subnets : for each AZ I've "public", "private", "database".

I don't want to discuss that specific point which could be improved, but for clarity's sake : there's one VPC by environment (production, staging, dev...), in production's VPC you've both IT stuff like the wiki, mattermost, etc. And our apps for our clients. 

I'm managing accesses for my different services through ingress security group stating "from this subnet, you can access this port on the ASG's instances the SG is attached to". A real life example being "rabbitmq-from-vpc-private", meaning anything in the VPC's private subnets can access the rabbitmq ports on the rabbitmq ASG's instances. 

It's been good enough to begin with, but I feel it's lacking, for example to allow access from the SSH bastion in public subnet, I need a rule "allow SSH from VPC public" everywhere, which isn't exactly secure if one of my HAProxy get compromised for example.

So I'm thinking about switching to a "allow this access from this security group". The terraform module I'm using refer it as "Access from source security groups", I don't know if that's the word for it. I can see the gain, I'm even wondering if keeping subnets is useful once properly set up. But I've no idea of the downside except for the time spent.

Hence here I am asking for your opinions, experiences in the matter, any interesting article or piece of doc is welcome too.

Thanks guys &amp; gals, take care !
## [4][Best way to store logs for analysis](https://www.reddit.com/r/aws/comments/j54mo6/best_way_to_store_logs_for_analysis/)
- url: https://www.reddit.com/r/aws/comments/j54mo6/best_way_to_store_logs_for_analysis/
---
Slightly silly question, maybe, but doing silly projects is how I teach myself stuff.

I run an IRC bouncer in AWS that saves logs in JSON format. I want to archive said logs somewhere that will make for easy searching/analysis.

My original thought was Elasticsearch, but a managed cluster seems like overkill. Standing up an ES container in my Nomad cluster is problematic because EBS doesn't play well with ephemeral EC2 hosts (that may not always be in the same AZ) and EFS has too much latency to host ES.

Thoughts?
## [5][CI/CD for Elastic Beanstalk](https://www.reddit.com/r/aws/comments/j5ikem/cicd_for_elastic_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/j5ikem/cicd_for_elastic_beanstalk/
---
I will be creating a beanstalk app and need guidance on best practices for CI/CD.

I am well versed with the AWS CLI, and Cloud Formation, and have a bit of experience with the EB CLI. The last project I worked on that used beanstalk had an environment that was set up manually then updated via pre-written scripts that used the EB CLI. Unfortunately this made it difficult to stand up clones of the environment due to the manual step.

The environment will consist of a multi-instance java app and will be using RDS. Standard public web app stuff like CNAME, redirect HTTP to HTTPS, putting the ec2 instances in a private subnet while having the load balancer public, etc.

What is the best way to do this? In my serverless environments i'm able to stand up a new env front/backend, s3, api gateway, dynamo, route53, etc. in about an hour using cloud formation/SAM. I'm looking for something similar with beanstalk. I haven't found any great docs on how to do this for more advanced setups other than using the console. Or is the best way to stand up new beanstalk environments actually to use the console, then make modifications with the CLI/config files as I was doing in my other project?
## [6][Scheduling/Autoscaling Pet instances on AWS EC2](https://www.reddit.com/r/aws/comments/j5ii8m/schedulingautoscaling_pet_instances_on_aws_ec2/)
- url: https://www.reddit.com/r/aws/comments/j5ii8m/schedulingautoscaling_pet_instances_on_aws_ec2/
---
Hi all,

I’m on a Move-to-Cloud project for one of our costumer. It is a very big migration that needs to be executed fast to go out their existing DC  ASAP (both for deadline DC closure reason and avoid double run costs  Cloud/On-Prem). So, refactoring all the applications is not an option due to the short  timeline. Currently, they are sized to the peak and migrating as is, could be  inefficient from a cost perspective. They can’t use AWS AutoScaling groups  yet due to their existing ‘Pet’ model on OnPrem VMware.

I’m looking for a solution that will do some kind of autoscaling on AWS EC2 instances  but orchestrating ‘Pet’ ones (so start/stop instances only and not create/terminate ones) based  on some custom metrics or CPU...

Thanks for any advice!
## [7][Is it correct to use Serverless Lambda functions to render templates?](https://www.reddit.com/r/aws/comments/j5ifbr/is_it_correct_to_use_serverless_lambda_functions/)
- url: https://www.reddit.com/r/aws/comments/j5ifbr/is_it_correct_to_use_serverless_lambda_functions/
---
I was transitioned to project using Serverless/Lambdas and needed to create a PDF generator.  It's a node project, using tsoa (TypeScript api), so I went with Puppeteer and Handlebar.js.

I'm running into an issue where locally the template directory exists and resolves but on the Lambda server it doesn't exist.

Our set up is, we have a parent directory in the root and every child directory to it is a service.  In a child service, I'm trying to access and read template files into memory to build a template in Handlebars.  Locally, the it finds the paths - but upon deploy - it cannot find the templates directory.  First templates was a sibling to `/proj_root/api`, but then I moved it to a shared folder under `api`.  Again, works locally but on server I get an error saying " `/var/task/api/shared/templates` could not be found."

I followed [this guide](https://www.serverless.com/framework/docs/providers/aws/guide/packaging/#examples) to try to include the directory with:

    package:
      include: api/shared/templates/**

Still to no avail.

Is this a misuse of serverless to include static files with a function?

Is it "best practice" to use an S3 bucket and make an api call for every static template file?

Is there a serverless way to include static files with a function? (AWS Layers?)
## [8][Slow AWS Client VPN Performance](https://www.reddit.com/r/aws/comments/j5hpyb/slow_aws_client_vpn_performance/)
- url: https://www.reddit.com/r/aws/comments/j5hpyb/slow_aws_client_vpn_performance/
---
We are testing AWS Client VPN before we completely migrate from using Windows Server as our VPN server. We're using certificates generated from easyRSA as authentication and  everything seems to be working as expected. I tried using the ovpn config through AWS' own VPN app and OpenVPN client.  


However, I  noticed VPN clients are getting significantly lower speeds compared to our old setup. I'm talking about around 6Mbps where it should be 50Mbps.  


To  troubleshoot, I ran the speedtest on a Windows Server EC2 instance on  the same region as the Client VPN endpoint (Tokyo). The speeds are upwards 100Mbps as expected.  


I've read about disabling "Hardware Checksum  Offloading Hardware TCP Segmentation Offloading" but I can't find a way  to configure this on AWS Client VPN. I checked laptop's NIC and I can't  find that option in 'advanced' tab either. Help?
## [9][Configuring a white label name server using Route 53 - help](https://www.reddit.com/r/aws/comments/j59we9/configuring_a_white_label_name_server_using_route/)
- url: https://www.reddit.com/r/aws/comments/j59we9/configuring_a_white_label_name_server_using_route/
---
Hello there

First of all, forgive me for m poor knowledge about this subject, but I have a few questions:

I successfully followed this tutorial ([https://www.youtube.com/watch?v=ahnhoFgd5K4&amp;ab\_channel=AmazonWebServices](https://www.youtube.com/watch?v=ahnhoFgd5K4&amp;ab_channel=AmazonWebServices)) on how to configure a white label name server using Route 53 (I wasn't able to confirm the changes on WHOIS because there wasn't enough time for the DNS changer to propagate probably). However, after doing that I got some questions that I hope someone can help me out:

1. Since I created a new hosted zone with a new set of name servers for the domain, should I also update the old name servers on the domain information page Route 53 panel &gt; Domains &gt; Registered domains &gt; mydomain &gt; name servers? I did that by the way
2. After the changes have propagated and I see them on WHOIS, if I click on the IP addresses of my name servers, I can check they are from Amazon. That would confirm to me that I have set all the IP addresses correctly right? (i'm concerned that I used a wrong Ip address during the process or something like that, maybe I have OCD hahah)
3. I have two domains, a primary one to actually use it and secondary one just to keep. I created this white label name server on my secondary domain since it was my first time just to try. Now I'm thinking about using this secondary domain white label name servers for my primary domain, just because it's done and easier to configure. Is that a bad practice or maybe not recommended for some reason? (besides the fact that I'll have to pay for a second hosted zone)

Thanks, any help is appreciated
## [10][Redshift/Quicksight: Most performant way to handle querying arrays within json?](https://www.reddit.com/r/aws/comments/j5525n/redshiftquicksight_most_performant_way_to_handle/)
- url: https://www.reddit.com/r/aws/comments/j5525n/redshiftquicksight_most_performant_way_to_handle/
---
Hi all,

I currently have a large set of json data that I'd like to import into Amazon Athena for visualization in Amazon Quicksight. each json contains two fields: one is a comma separated string of ids (orderlist), and the other field is an array of strings(locations). What is the best way to handle/store this data to produce the most performant queries when generating dashboards on Quicksight?

 Because Quicksight doesn't support array searching, I'm currently resorting to creating a view where I generate crossjoins across the two string arrays:

    select id,
     try_CAST(orderid AS bigint) orderid_targeting,
     location
    from advertising_json 
    CROSS JOIN UNNEST(split(orderlist, ',')) as x(orderid)
    CROSS JOIN UNNEST(locations) t (location)

I've also considered flattening out the data beforehand, would that make a difference in query performance? Considering that I have two crossjoins currently, that could 30x my storage space if I explode out the two arrays, I'm hesitant to go down this route if it doesn't make a significant difference in performance.
