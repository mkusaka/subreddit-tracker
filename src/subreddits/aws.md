# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][AWS CloudWatch colored logs for Spring apps](https://www.reddit.com/r/aws/comments/jd67la/aws_cloudwatch_colored_logs_for_spring_apps/)
- url: https://chrome.google.com/webstore/detail/aws-cloudwatch-ansi-color/feoelakkoolicldilidmgbhpgabiodcd
---

## [3][Suggested AWS tool to browse large amounts of data](https://www.reddit.com/r/aws/comments/jdb8vo/suggested_aws_tool_to_browse_large_amounts_of_data/)
- url: https://www.reddit.com/r/aws/comments/jdb8vo/suggested_aws_tool_to_browse_large_amounts_of_data/
---
I'm having a rather large elasticsearch cluster (1 month, 1TB of data spread on 9 nodes, thousand of shards). 

The data is ingested daily and supposed to be stored for 10 more years. This will obviously cause the elasticsearch cluster to grow along with the cost.

To prevent this, I'm currently implementing a snapshot policy to move old data to S3 buckets and procedures to restore them quickly when they would be needed.

I'm wondering if maybe I should use some other AWS Tool like Athena or Redshift that would allow to browse the data stored on S3 before deciding to load them back into elasticsearch. Would any of them be suitable for this and - more importantly - would not cost much when not in use?
## [4][[HELP, BUG?] Workspace deploying in public subnets (should be private)](https://www.reddit.com/r/aws/comments/jdf17i/help_bug_workspace_deploying_in_public_subnets/)
- url: https://www.reddit.com/r/aws/comments/jdf17i/help_bug_workspace_deploying_in_public_subnets/
---
Hello. This is one of my first posts on reddit, long time user and lurker though.

I've been using workspaces in my company and they're working well except for one potential bug or configuration issue. I've configured a VPC with one public subnet and 2 private subnets and I have attached a Simple AD directory to it, which uses the private subnets.

The private subnets have internet access via a gateway and the workspaces have internet access when they are in the private subnets, so everything seems to work as intended. This architecture is described here: [https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-vpc.html#configure-vpc-nat-gateway](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-vpc.html#configure-vpc-nat-gateway)

**Here's the odd part:** Whenever I launch workspaces, they seem to have a random chance of being deployed in the public subnet instead of either of the private ones. To get them to deploy in the private subnets I have to rebuild them a few times (sometimes only once but that's just luck). Once they are deployed in the private subnets they then work fully, internet access and all, as described above.

I've read the post here: [https://www.reddit.com/r/aws/comments/esw6fd/workspace\_provisioning\_in\_wrong\_subnet/](https://www.reddit.com/r/aws/comments/esw6fd/workspace_provisioning_in_wrong_subnet/)  
And the cause of the issue there doesn't seem to be the same cause for the issues I have with my setup.

Any advice is greatly appreciated, I have not been able to find any relevant articles or information online about my specific issue, and there doesn't seem to be anything obviously wrong with my setup/configuration.

EDIT: Clarifications
## [5][Can API Gateway let me know what time zone the request is coming from?](https://www.reddit.com/r/aws/comments/jdeifq/can_api_gateway_let_me_know_what_time_zone_the/)
- url: https://www.reddit.com/r/aws/comments/jdeifq/can_api_gateway_let_me_know_what_time_zone_the/
---
I am querying our database, which stores the date time data in GMT time. I want to adjust these times based on where the client is requesting the data.

Can API Gateway let me know what time zone people are in, so that I can adjust the times?
## [6][EKS + Secrets Manager](https://www.reddit.com/r/aws/comments/jczx1q/eks_secrets_manager/)
- url: https://www.reddit.com/r/aws/comments/jczx1q/eks_secrets_manager/
---
Running Kubernetes on AWS is quite a bit easier than it used to be, but even now with EKS you can still feel the clash of approaches/technologies. There's the K8s way to do things and the AWS way to do things and the mish mash world of EKS.

Anyway, I would like to come up with a sane way to synchronize secrets and even configmaps in a way that feels native to AWS and is also simple. The best solution I've seen thus far is this https://github.com/godaddy/kubernetes-external-secrets

What I'd really like is a tool that can tell me what I have in Secrets Manager and/or Parameter Store and diffs that against what I have (insensitive and maybe templated secrets) in source control and in my clusters.

I have a shitty combination of shell scripts to do this now but I'm close to writing my own tool. Just curious to hear what people are doing on this front.
## [7][Cost of NVMe SSD Instance Storage?](https://www.reddit.com/r/aws/comments/jdb4cv/cost_of_nvme_ssd_instance_storage/)
- url: https://www.reddit.com/r/aws/comments/jdb4cv/cost_of_nvme_ssd_instance_storage/
---
Hi, was checking pricing of some of the ec2 instance, and I was wondering if there is a different cost for the NVMe SSD instance storage for nitro instances (ala EBS)?  Or is the price already included into the hourly instance cost?
## [8][CDK Native Applications - Are there any established patterns for highly integrated apps?](https://www.reddit.com/r/aws/comments/jczefh/cdk_native_applications_are_there_any_established/)
- url: https://www.reddit.com/r/aws/comments/jczefh/cdk_native_applications_are_there_any_established/
---
TLDR; I argue that there should be an option to support writing and transpiling TypeScript based Lambda functions out of the box with the CDK.

People keep mentioning how the lines between Application and IaC are getting blurred (Darko Mesaros touches on this in [this video](https://youtu.be/fWtuwGSoSOU?t=225). However, besides a very small number of posts on Medium, this idea doesn't appear to be fully embraced. Every video, repo, or article I've come across shows some interesting IaC solution using the CDK, but with the actual application logic being a completely decoupled, trivial, hello-world.js. I cannot find any example of a TypeScript based Lambda being integrated in the CDK docs / example repos and for me this would really be the icing on the cake for the CDK.

This goes beyond simply wanting to write handlers in TypeScript. For most non-trivial Serverless applications, some of the most important services that Lambda integrates with are SQS and SES. Defining a Lambda handler that consumes from an SQS queue is, obviously, very well supported in the CDK. However, having the same handler dispatch messages to other queues, or to SES, or write to a DB, has no type-safe integration with the constructs defined in the CDK. So to write a Serverless application in the CDK, you either have to define some configuration in JSON / JavaScript and share that between the application logic and the IaC logic, or define your configuration twice.

I have written about 5-6 Serverless stacks with the CDK which are 100% TypeScript based and it's been great. Type-safe interfaces / configurations that are shared between CDK constructs and App handlers have made complex stacks a piece of cake with little scope for misconfiguration or bugs. All that's required is a fairly straightforward Webpack configuration. I'm happy to share a template repo if people are interested.

I want to know if this is some sort of crazy anti-pattern, or the next logical step in Serverless applications. Are other people building the same stack over and over, or is it just me? If this is in fact a popular model then why isn't it in official docs or templates, and what's the scope for having an official convention / template added somewhere?
## [9][Managing my domain through AWS?](https://www.reddit.com/r/aws/comments/jd8xnf/managing_my_domain_through_aws/)
- url: https://www.reddit.com/r/aws/comments/jd8xnf/managing_my_domain_through_aws/
---
Hi

I just found out how awesome Elastic Beanstalk is, I was looking for the best way to deploy a Django project and came across EBT, read the documentation and found out that it is super duper easy to get up and going!

I have a domain registered through namecheap that I want to manage through AWS as well, is it possible? By manage I mean adding / removing records, alleviating the option to ever log in into namecheap's control panel.
## [10][[Noob Question] Network Routing through the internet](https://www.reddit.com/r/aws/comments/jctgv4/noob_question_network_routing_through_the_internet/)
- url: https://www.reddit.com/r/aws/comments/jctgv4/noob_question_network_routing_through_the_internet/
---
Hi, Noob here, just have a question how the network is routed through the internet when the client hits the url.

I want to know how the traffic is routed using with/without WAF, Route 53(hostedzone) and cloudfront

This is my understanding so far

with cloudfront, route 53, Waf

CLient(http request) ----&gt; Domain registrar DNS servers ----&gt; Route53 DNS ----&gt; cloudfront ----&gt; WAF ----&gt; EC2 live server

THIS MIGHT VERY WELL BE A BLUNDER. I want to know the correct traffic travel path. Thank You. Experts please help!!
## [11][Extending Budget with Chatbot to Slack (CDK + CF)](https://www.reddit.com/r/aws/comments/jctbz4/extending_budget_with_chatbot_to_slack_cdk_cf/)
- url: https://www.reddit.com/r/aws/comments/jctbz4/extending_budget_with_chatbot_to_slack_cdk_cf/
---
&amp;#x200B;

https://preview.redd.it/cauqo9w81nt51.jpg?width=6000&amp;format=pjpg&amp;auto=webp&amp;s=2dfb83a73c96dfaa8b750aaf15f5edcebe58e2bd

Hi All,

My previous post discussed AWS Budget ([https://www.reddit.com/r/aws/comments/jan2x2/aws\_cdk\_for\_aws\_budget\_with\_notification\_and\_cf/](https://www.reddit.com/r/aws/comments/jan2x2/aws_cdk_for_aws_budget_with_notification_and_cf/)) and this one builds on that to integrate Budget with Chatbot for notifications via Slack.

I think Chatbot is a pretty cool free service from AWS that will be one to watch in the future. This basic example will get you up and running with CF deploy stack button in minutes or you use the CDK stack to see how its built and make it better.

There github repo for this is open source and available here:

[https://github.com/talkncloud/aws/tree/main/essential-billing-bot](https://github.com/talkncloud/aws/tree/main/essential-billing-bot)

Question for everyone: Does anybody know how to script or automate the configuration between Chatbot and Slack? At the moment this is manual step in the console. I'd be interested to hear if this was something you can include in the stack? I've seen a couple of other services like this in AWS.

Suggestions &amp; feedback welcome, let me know how you're using Chatbot!

Thanks,

Mick
