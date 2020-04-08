# aws
## [1][AWS started running Folding@Home on Spot Instances](https://www.reddit.com/r/aws/comments/fx1aw0/aws_started_running_foldinghome_on_spot_instances/)
- url: https://folding.extremeoverclocking.com/team_summary.php?s=&amp;t=238068
---

## [2][Optimizing Japanese text-to-speech with Amazon Polly | Amazon Web Services](https://www.reddit.com/r/aws/comments/fwt4ma/optimizing_japanese_texttospeech_with_amazon/)
- url: https://aws.amazon.com/blogs/machine-learning/optimizing-japanese-text-to-speech-with-amazon-polly/
---

## [3][Network configuration error and AWS not reachable anymore](https://www.reddit.com/r/aws/comments/fx6hce/network_configuration_error_and_aws_not_reachable/)
- url: https://www.reddit.com/r/aws/comments/fx6hce/network_configuration_error_and_aws_not_reachable/
---
Hello,   
I made a stupid thing :   
I tried to setup a network bridge connection between two network connections - comporting my "main" one on an AWS server (windows) but it won't be reachable again. I think the network interface crashed and an hard reset didn't fix it.   


Is there a fix to this mistake ?
## [4][S3 lifecycle not triggering](https://www.reddit.com/r/aws/comments/fx0qo4/s3_lifecycle_not_triggering/)
- url: https://www.reddit.com/r/aws/comments/fx0qo4/s3_lifecycle_not_triggering/
---
I have two "folders" in a standard bucket, where i want to move one "folder" to glacier. However since one of the folders has 9TB of 180byte sized files, i absolutely don't want to accidentally move it to glacier (because of astronomical metadata costs on such small files), so i wanted to first test out the lifecycle rule on dummy data, before committing it on the real data.

I set up a bucket with two "folders" of fake data:

[https://imgur.com/a/44xeLLV](https://imgur.com/a/44xeLLV)

And added a lifecycle rule with the following prefix:

[https://imgur.com/a/cIlQFfD](https://imgur.com/a/cIlQFfD)

I waited two days, but nothing seems to be happening. Am i declaring the prefix incorrectly? Should it have slashes from both sides?
## [5][Whitelist only traffic that has gone through my NLB, how?](https://www.reddit.com/r/aws/comments/fx5hq3/whitelist_only_traffic_that_has_gone_through_my/)
- url: https://www.reddit.com/r/aws/comments/fx5hq3/whitelist_only_traffic_that_has_gone_through_my/
---
Im working on a setup: Public web --&gt; NLB(with EIP) --&gt; EC2(with webserver)

However clients can just circumvent the NLB and connect to the EC2 directly , and by doing this avoid the logging happening in the NLB, which is a problem. Im not sure how to avoid the at the moment outside of making the EC2 private, but in this case i would also lose SSH access to it which is not acceptable.  
Ive tried:

1. Add a SG to the NLB so i can whiteliste this SG in the EC2s SG, however NLBs dont support SGs. ALBs dont support EIPs which i need, so im stuck with the NLB.
2. Whitelisting the NLBs EIP in the EC2s SG, however it seems traffic is still resolved in the EC2s by their source IP which is not whitelisted, hence no connection.

What options do i have here to block clients from just connecting to the EC2 directly given the circumstances where i need SSH access to the EC2 and EIP for my LB ?
## [6][Lambda to create jira ticket?](https://www.reddit.com/r/aws/comments/fwpm1q/lambda_to_create_jira_ticket/)
- url: https://www.reddit.com/r/aws/comments/fwpm1q/lambda_to_create_jira_ticket/
---
I have pushed ecs services log to aws cloudwatch log group for creating alarms to trigger lambda that take filtered log group’s content and send to group of email. 

I want to have a way to create jira ticket (jira cloud) from aws service (lambda maybe?)— I’m not sure how. 

So I’m asking is there away to create jira ticket from aws that contain filtered contents from aws log group? And how to make it non-repeatedly of same filtered content?
## [7][Active Directory](https://www.reddit.com/r/aws/comments/fwu93o/active_directory/)
- url: https://www.reddit.com/r/aws/comments/fwu93o/active_directory/
---
Im on a one forest and one domain structure, HQ in california, office in NV, and NY. Most servers are in HQ along with Primary Domain Controller and i have Read Only Domain Controllers at the other sites. NY and NV are both connected to CA by site to site vpn where they access a file server (there are other servers, just keeping it simple). We are currently 100% on prem. 
I cleared my AWS SAA and would like to create a cloud environment that Could potentially be used for work. As of now, it’s just a personal project but I would like to be an AWS Solutions Architect in the future so I’d like to approach this personal project as professional as possible. 
Now that I got that out of the way, on to the questions:

1. Should I create an EC2 and use win 2016 AMI      then install the roles and promote it as PDC? offices would then have RODC including HQ. 
Or
2. Should I create an EC2 with a different domain but create a trust relationship?
Or
3. Use AWS Managed Directory but would it let me have RODC for other sites and keep my users, groups, OU’s, and permissions? Probably not. 
Or
4. Any suggestions?

Also, what is the devops approach for this? Cloud Formation the EC2? versioning in S3 or Code Commit the YAML template? Use powershell to add users? 
This project is multi-purpose - something i could possibly use for work, showcase for a future interview as Cloud Solutions Architect, and self learning (with guidance from reddit fam). 

Sorry for the long post and loaded post, I have a long journey ahead. Thanks for any input!
## [8][Data Transfer Costs between ELB and EC2](https://www.reddit.com/r/aws/comments/fx46ig/data_transfer_costs_between_elb_and_ec2/)
- url: https://www.reddit.com/r/aws/comments/fx46ig/data_transfer_costs_between_elb_and_ec2/
---
Hey everyone!

I am trying to figure out, what our traffic costs will be for our setup:

We have an internet-facing ALB in region eu-central-1 and two EC2 instances in the same region. ALB is configured with two subnets in AZs eu-central-1a and eu-central-1b. In each of these AZs, there is one EC2 instance running in a private subnet with only a private IP.

Clients are loading data via the ALB, so we have mostly outgoing data.

I know how to calculate the costs for the ALB, but what I don't get is the data transfer costs between EC2 and ALB. According to the docs: *"Data transferred "in" to and "out" from Amazon Classic and Application Elastic Load Balancers using private IP addresses, between EC2 instances and the load balancer in the same AWS Region is free."*  
Do they mean "EC2 instance is connecting to the ALB via private IP" or "ALB is connecting to the EC2 via private IP"?

Another thing I don't get is this: According to the docs, traffic between ALB and EC2 within the same AZ is free. Considering our setup, does that mean the traffic will always be within the same AZ no matter what instance is chosen by the ALB as the target?

Thank you for any help!
## [9][Should I stop idle EC2 spot instances?](https://www.reddit.com/r/aws/comments/fx17k7/should_i_stop_idle_ec2_spot_instances/)
- url: https://www.reddit.com/r/aws/comments/fx17k7/should_i_stop_idle_ec2_spot_instances/
---
Sorry for the noob question. I've set up a small GPU-enabled EC2 instance that I am going to provide to four or five people in my lab for occasional use (no one has access to a GPU and we need it for our work). I don't want to keep stopping and starting the instance, and I don't know when people will need to access it. 

&amp;#x200B;

Will I actually be charged much if the instance is just sitting idle (but not switched off)? I'm under the impression that the pricing is scaled by usage, i.e. if the CPU is running at 2% I will be charged less than if I am pushing the machine at 100% utilization. Is this true, and is the charge for low usage scaled linearly by utilization? Thanks.
## [10][Deleted Lightsail Instance without detaching Static IP](https://www.reddit.com/r/aws/comments/fx39le/deleted_lightsail_instance_without_detaching/)
- url: https://www.reddit.com/r/aws/comments/fx39le/deleted_lightsail_instance_without_detaching/
---
Hi guys,

I spun up a few Amazon Lightsail Instances, installed wordpress, attached the provided Static IP addresses (5) then deleted the instances without detaching the Static IP's first. Now i have spun up new instances and when i try to attach a Static IP Im getting an error message that i have exceeded the maximum number of Static IP addresses. This despite I physically have less than 5 instances in my console. Does anyone know what I can do here?
