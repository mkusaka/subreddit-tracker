# aws
## [1][How to develop your Lambda Functions like a rockstar - our firsthand experience](https://www.reddit.com/r/aws/comments/f5w9hm/how_to_develop_your_lambda_functions_like_a/)
- url: https://www.reddit.com/r/aws/comments/f5w9hm/how_to_develop_your_lambda_functions_like_a/
---
Hey all - [thought I'd share some learnings and experiences](https://medium.com/@joel.tbarna/how-to-develop-your-lambda-functions-like-a-rockstar-7d9422259d57) we've had getting up-to-speed developing our application with just AWS Lambda. It was pretty slow at first but we've created a pretty solid strategy around locally developing and testing that may be helpful to anyone taking on the challenge of Serverless development.

Let me know if you have any questions! Happy to help where I can.
## [2][RDS Data API Inconsistency and weird behavior](https://www.reddit.com/r/aws/comments/f6ag4c/rds_data_api_inconsistency_and_weird_behavior/)
- url: https://www.reddit.com/r/aws/comments/f6ag4c/rds_data_api_inconsistency_and_weird_behavior/
---
We are using RDS Data API for communication with the Postgres Serverless cluster from AWS Lambdas. Everything works well with small loads, like 1-10 simultaneous connections, but when load grows to 20+ simultaneous connections, RDS Data API becomes unstable and fails with weird things.

&amp;#x200B;

1. RDS Data API requests are mixed with their respective responses. When addressing about 50+ connections with SQL queries by RDS Data API, about 5% of requests can return results from another parallel request. For example, 50 lambdas invoke SQL query over table A and different 50 lambdas invoke SQL query over table B. As a result, some queries to table A returns result from table B and vice versa.
2. RDS Data API proxy connections can hang permanently. When 50+ lambdas have issued even one SQL query and even it has been responded with success, the following queries fails with different errors, but they can be summarized like Too many connections. Only connecting as administrator and force killing RDS connections can help.

Second issue can be avoided by wrapping requests into loop with jitter, and terminating old stuck connections from active connections. But first issue is terrible - when results from SQL requests are mixed, there is no way to workaround it from RDS-client side. What you can recommend in this situation?

Today we found new RDS Data API issue. If somebody tries to access RDS Data from different AWS account, he runs into error cluster &lt;ARN&gt; does not belong to the calling account id &lt;ACC-ID&gt;  
, which is good behavior. But after reaching RDS Data from different account, it stop works from current legitimate AWS account too with same error.

Unfortunately, reaching RDS Data API team via Rds-data-api-feedback@amazon.com  
 has had no result via last month. Maybe there are any AWS Data API developers here and can contact us about this really weird issue. Thanks
## [3][Exposing stable DNS/IP for Kube cluster (EKS) is only possible throuh AWS managed LBs?](https://www.reddit.com/r/aws/comments/f6ajju/exposing_stable_dnsip_for_kube_cluster_eks_is/)
- url: https://www.reddit.com/r/aws/comments/f6ajju/exposing_stable_dnsip_for_kube_cluster_eks_is/
---
Hi everyone, i hope you can guide me a bit.

We have Managed Kubernetes (EKS) and exposing services through Classic LB  or ALB. 

However i would like to reduce usage of AWS managed LBs but do not see a way to use ingress or somethiong else running on the cluster that would allow me to have publiv and stable DNS name or IP Adress. 

is there as way to do so except running your own EC2 as LB which sounds not much better than managed LBs?
## [4][Does anyone have a link to the presentation (pdf/ppt/mp4/other) for NET340 - networking for AWS Outposts?](https://www.reddit.com/r/aws/comments/f6abg0/does_anyone_have_a_link_to_the_presentation/)
- url: https://www.reddit.com/r/aws/comments/f6abg0/does_anyone_have_a_link_to_the_presentation/
---
AWS' link of links [here](http://aws-reinvent-audio.s3-website.us-east-2.amazonaws.com/2019/2019_presentations.html) points to an almost empty (four useless slides) on the [ENT340 topic](https://d1.awsstatic.com/events/reinvent/2019/REPEAT_1_Networking_for_AWS_Outposts_NET340-R1.pdf).
## [5][Monitor JVM in AWS Fargate](https://www.reddit.com/r/aws/comments/f69ntd/monitor_jvm_in_aws_fargate/)
- url: https://www.reddit.com/r/aws/comments/f69ntd/monitor_jvm_in_aws_fargate/
---
Hi guys!

I have been currently trying to connect VisualVM (A program which monitors the JVM, heap and memory usage etc) to a Spring Boot Application (Java App) running on AWS Fargate in Docker containers.

I have been exposing the JMX ports accordingly and I am able to connect through the JMX ports when running the Docker container locally. However, when running the Java App on Fargate, I have not found a way to connect to the Container through JMX. I have tried setting the VM argument -Djava.rmi.server.hostname to the IP Address of the container, but when I try to connect through JMX it still fails to do so. Has anyone had any experience with this?

JMX commands for reference:

*ENTRYPOINT* java \\  
 *-*Dcom.sun.management.jmxremote \\  
 *-*Dcom.sun.management.jmxremote.local.only=false \\  
 *-*Dcom.sun.management.jmxremote.authenticate=false \\  
 *-*Dcom.sun.management.jmxremote.ssl=false \\  
 *-*Djava.rmi.server.hostname=172.17.0.2 \\  
 *-*Dcom.sun.management.jmxremote.port=9090\\  
 *-*Dcom.sun.management.jmxremote.rmi.port=9090\\  
 *-*jar research-api-0.11.7*-*alpha-1.jar server
## [6][Cloudfront Distribution - selective 404 responses depending on path?](https://www.reddit.com/r/aws/comments/f6968h/cloudfront_distribution_selective_404_responses/)
- url: https://www.reddit.com/r/aws/comments/f6968h/cloudfront_distribution_selective_404_responses/
---
Using Cloudfront to host a static React site (S3 origin) - how would I go about sending a visitor a 404 response vs. a 200 response, depending on the path?

To be clear, all requests need to hit the index.html in the S3 bucket.

Some requests (root, or /about-us, for example) - need a 200 response, however most paths would simply receive a 404 response (however, still rendering the React app using the index file).

TIA!
## [7][If you use statuscake, try this command line tool I've built](https://www.reddit.com/r/aws/comments/f68wdt/if_you_use_statuscake_try_this_command_line_tool/)
- url: https://github.com/omerh/statuscakectl
---

## [8]["No Instances In Tag" on run-command for patch baseline.](https://www.reddit.com/r/aws/comments/f68obf/no_instances_in_tag_on_runcommand_for_patch/)
- url: https://www.reddit.com/r/aws/comments/f68obf/no_instances_in_tag_on_runcommand_for_patch/
---
Can anyone offer advice on using Systems Manager / Patch Management / Maintenance Windows for patching EC2 instances.

Specifically the fact that every time my maintenance window triggers i get the response **"No Instances In Tag"** even though there definately ARE (or should be) EC2 instances with that Tag on them.

More details below:

I have set the "Patch Group" tag on my EC2 instances to be "Tuesday23to7"  ...
https://i.imgur.com/55vnrPJ.png

I have registered a maintenance window and set the target as being the "Patch Group" - "Tuesday23to7"
https://i.imgur.com/27f2FbV.png

The runcommand task is set to run on "registered target groups"
https://i.imgur.com/DoYwFhW.png

The patch baseline has had "modify patch groups" set to the same "Tuesday23to7"
https://i.imgur.com/Cg3v7LX.png

Every time the maintenance window triggers or even if i try to MANUALLY run the command, the command runs successfully, but I i get the declaration that there are "No Instances In Tag" and no targets read 'zero'.
https://i.imgur.com/SHlnpCt.png


Thanks!
## [9][Geo-Blocking](https://www.reddit.com/r/aws/comments/f68kh1/geoblocking/)
- url: https://www.reddit.com/r/aws/comments/f68kh1/geoblocking/
---
Hi there,

I am using Amazon Lightsail with 3CX. I would like to set up Geo-Blocking. Does anyone know if Geo-Blocking is supported on lightsail?
## [10][AWS Organization and Billing](https://www.reddit.com/r/aws/comments/f67sk8/aws_organization_and_billing/)
- url: https://www.reddit.com/r/aws/comments/f67sk8/aws_organization_and_billing/
---
I was using an AWS account with my credit card info and my personal account. Since I'm learning my way around AWS services I am using free tier as much as I can but I still have some small charges sent to my credit card at the end of the month. 

Recently I started a freelance project and they sent me a link and I became a part of the organization. Under My Organization I see the organization with the following note: "*Organization features enabled... All features enabled...*".

My question is how is billing done and if I spin up an instance not in the free tier who will pay for it and how can I manage and define who is paying for what?
