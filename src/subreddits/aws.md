# aws
## [1][US-EAST-1 outage use1-az2](https://www.reddit.com/r/aws/comments/j7xvib/useast1_outage_use1az2/)
- url: https://www.reddit.com/r/aws/comments/j7xvib/useast1_outage_use1az2/
---
We are facing issue for new instance scaled from 1:20 AM PDT having netowrk connectivity issue.. Tried other 2 AZ B,A &amp; D..

Is anyone else facing problem? What other service.. I believe all major service calling EC2 or network interface will be affected..

&gt;Network connectivity issues  \[01:48 AM PDT\] We are investigating networking connectivity issues for a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. We have identified root cause and are working towards resolution. Network connectivity for existing instances is not affected by this issue. For newly launched instances that are affected, relaunching a new instance may resolve the issue.  
&gt;  
&gt;\[02:47 AM PDT\] We continue to work toward recovery for the networking connectivity issues affecting a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. Network connectivity for existing instances remains unaffected by this issue. For newly launched instances that are affected, relaunching a new instance may resolve the issue.  
&gt;  
&gt;\[04:53 AM PDT\] We are still working toward recovery for the networking connectivity issues affecting a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. Network connectivity for existing instances remains unaffected by this issue. For newly launched instances that are affected, launching a replacement instance may resolve the issue.

We beleive the issue started well before this.. probably 3-4 hrs before,  when we started noticing the higher error rates but where not sure.. Can AWS come clean on when it actually started &amp; how soon can this be resolved??  sitting ducks for now!..
## [2][Introducing AWS Lambda Extensions](https://www.reddit.com/r/aws/comments/j7gqsy/introducing_aws_lambda_extensions/)
- url: https://aws.amazon.com/blogs/compute/introducing-aws-lambda-extensions-in-preview/
---

## [3][Amazon ElastiCache Day](https://www.reddit.com/r/aws/comments/j7ngkb/amazon_elasticache_day/)
- url: https://pages.awscloud.com/AmazonElasticacheDay.html
---

## [4][Deploying application configuration to serverless: introducing the AWS AppConfig Lambda extension | Amazon Web Services](https://www.reddit.com/r/aws/comments/j7pync/deploying_application_configuration_to_serverless/)
- url: https://aws.amazon.com/blogs/mt/introducing-aws-appconfig-lambda-extension-deploying-application-configuration-serverless/
---

## [5][Oneliner to map aws AZ ids to AZ names](https://www.reddit.com/r/aws/comments/j7vcup/oneliner_to_map_aws_az_ids_to_az_names/)
- url: https://www.reddit.com/r/aws/comments/j7vcup/oneliner_to_map_aws_az_ids_to_az_names/
---
[https://gist.github.com/awsiv/bc7a2a28679a7bfc59b988134a8597bd](https://gist.github.com/awsiv/bc7a2a28679a7bfc59b988134a8597bd)  


\`aws ec2 describe-availability-zones --region us-east-1 | jq -r '.AvailabilityZones\[\] | .ZoneId + "=" + .ZoneName'\`   


With the current DNS issue in \`us-east-1\` - this has proven quite helpful
## [6][AWS Pen Testing and Whitelisting](https://www.reddit.com/r/aws/comments/j7xi2z/aws_pen_testing_and_whitelisting/)
- url: https://www.reddit.com/r/aws/comments/j7xi2z/aws_pen_testing_and_whitelisting/
---
I manage security for an organization that has an AWS environment that is used for eCom as well as some other functions. We are required to have a pen test annually and to have the segmentation of the network tested (technically two separate tests) as well. 

The setup for the test includes placing pen test jump hosts at various places within the environment so that the tester can simulate having a foothold inside various internal networks in the environment. We have a standoff over a configuration point: I say that while we modified the test platforms and applied a SG that will allow all outbound traffic from the test platforms/jumphosts to all IPs in our environment, we have not modified the SGs on any other hosts in the environment. Similarly, we have not modified the NACLs in any way. It is our position that these modifications would invalidate the test since the test should be to determine whether or not our controls are effectively protecting and/or segmenting the environment. 

The tester's position is that his test platform should be whitelisted everywhere so that he can reach all the hosts and access anything in the environment. 

I'm looking for any sort of references that are out there that would provide guidance on which position would be considered best practice. I've looked all over and all I can find is references that call out physical networks and the practices there. Can anyone point to cloud-specific guidance on this? I'd prefer AWS specific, but I'll take what I can get.

&amp;#x200B;

Thanks
## [7][Insufficient capacity for lambda + API Gateway](https://www.reddit.com/r/aws/comments/j7uxk0/insufficient_capacity_for_lambda_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/j7uxk0/insufficient_capacity_for_lambda_api_gateway/
---
Hi there, I’ve started receiving this 5xx error in the API Gateway + Lambda in the Sydney region a day ago.

“We currently do not have sufficient capacity in the region you requested. Our system will be working on provisioning additional capacity. You can avoid getting this error by temporarily reducing your request rate.”

As a result, our services were down for hours. Is it possible for lambdas to fail to start because of capacity issues? If so, how can I prevent this from happening by routing requests to backup instances say an ECS or EC2.
## [8][AWS CodePipeline Docker on existing EC2 Machine](https://www.reddit.com/r/aws/comments/j7uu19/aws_codepipeline_docker_on_existing_ec2_machine/)
- url: https://www.reddit.com/r/aws/comments/j7uu19/aws_codepipeline_docker_on_existing_ec2_machine/
---
Hi everyone. I'm pretty new to AWS concepts and DevOps concepts, I'm currently studying both of them worlds. Right now I'm free tiering for everything since we're trying to bootstrap a startup here in my country.   
I'm trying to make a Pipeline that deploys a Docker image on an existing EC2 machine that we are already using for other things ( such as Redmine ). 

I think that the Pipeline should be something like this:  


1) CodeCommit for the source code   
2) CodeBuild that run the build process for the docker images  
3) ECR that contains the built docker image  
4) CodeDeploy that runs a script which pull the new built image and run it on the existing EC2 machine  


Up until point 3 everything is cool and fine, I don't really know how to make point 4. I was thinking of maybe running a remote script from a ssh shell ( I don't even know if it's possible on CodeDeploy ) with the .pem that pull the images and run them.  
Have you ever came across a problem like this?
## [9][How to point root record to an ALB?](https://www.reddit.com/r/aws/comments/j7rfiq/how_to_point_root_record_to_an_alb/)
- url: https://www.reddit.com/r/aws/comments/j7rfiq/how_to_point_root_record_to_an_alb/
---
The ALB’s IPv4 constantly changes, and I can’t use a CNAME for the root record (@).  

I’m on Google Domains and managed to forward the @ record to the www subdomain. The problem is the client doesn’t want the www as the root is the perfect name for their brand.  

Now I’m thinking, would it be possible to point the root record at an NLB, then have the NLB point to the ALB? It’s a lot more rerouting than I was expecting, so wondering if it’s possible and if so, is it a good practice to do so?
## [10][Problem with AWS Fargate and Elastic ALB's - Any Help Would be Appreciated...](https://www.reddit.com/r/aws/comments/j7r0kv/problem_with_aws_fargate_and_elastic_albs_any/)
- url: https://www.reddit.com/r/aws/comments/j7r0kv/problem_with_aws_fargate_and_elastic_albs_any/
---
Hey Friends!

I have a question/problem with AWS Fargate and Elastic ALB's that I could use some help with. I've been banging my head against a figurative brick wall all day and now I have a massive headache. I feel like this should be simple but for some reason it's just not. Here's the problem I have a node js application that I've just finished putting into a docker container. I exposed a port in the Dockerfile, login into ECR, build the container, tag it and push into ECR and then finally deployed the container into a Fargate task. When it's deployed I can access the container using both the Public IP and the port that I exposed. Ideally I would like to have this container behind an Application Load Balancer (ALB) so I create a cname with Route 53 and point it the ALB's Public DNS name. This is my though process of how It should go:

1. deploy container into AWS Fargate
2. open port xxxx in security rule attached to eni-xxxx that is attached to AWS Fargate.
   1. Test to see if I can access the container by using both the public IPv4 address and port.
3. create target group that has private IPv4 Address of container hosted in AWS Fargate.
4. create security group to be attached to the ALB that accepts traffic from everywhere but only on port 443.
5. modify security group attached to eni-xxxx that's attached to AWS Fargate to only accept traffic from the security group attached to the ALB. 
6. deploy ALB attaching security group and target groups created.
7. copy Public DNS name from ALB and create a cname that points to the ALB's public DNS name.

When I test... I get 502 bad gateway... which doesn't make that much sense to me. This in my mind should work for me, it looks like the ALB is rejecting my packet... so r/aws what am I doing wrong? I've added logging to the ALB and nothing of useful is displaying in the logs, as well I've used both dig and traceroute... I can confirm dig is pointing authoritatively to the ALB and traceroute confirms packet is being dropping off at the ALB...
