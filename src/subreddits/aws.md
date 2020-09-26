# aws
## [1][S3 Path Deprecation Plan Updated](https://www.reddit.com/r/aws/comments/iyd84u/s3_path_deprecation_plan_updated/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [2][Transitioning from SysOps to DevOps](https://www.reddit.com/r/aws/comments/j041h6/transitioning_from_sysops_to_devops/)
- url: https://www.reddit.com/r/aws/comments/j041h6/transitioning_from_sysops_to_devops/
---
I am currently employed as a Systems Engineer for a consulting company which serves many clients here in Italy.
I'm mainly a Windows Admin, due to exposure, and have no formal training. I work with all the usual hassle (vmWare, networking, WS, some Linux machines, security, AD,....) but due to personal reasons I would like to relocate to a different country. 
I see many job offers as DevOps and after having a look around, I got interested in moving my focus into cloud based infrastructures, mainly AWS.
I grasp OOP concepts and have some personal experience in programming or scripting tools for my job (VBA and Powershell).
If you were in my position, how would you move ahead in order to improve your knowledge of DevOps and show a future employer that you have the skills he requires in order to work in this field?
Would you go with certs such as (AWS SysOps engineer)?
Which (paid if necessary) training would you undergo?

Thanks.
## [3][Where would you store a secret shared key?](https://www.reddit.com/r/aws/comments/izsd21/where_would_you_store_a_secret_shared_key/)
- url: https://www.reddit.com/r/aws/comments/izsd21/where_would_you_store_a_secret_shared_key/
---
If I want to break my app up using Lambdas, and I want to share a secret key to authenticate jwt tokens, I'm assuming I'm better off keeping it in one location rather than having it stored in each app on Lambda.  Does AWS have a built in tool for this?  I've heard others talk about something I think was called Zookeeper, but I believe that is a third party program.  I'm wondering if AWS has a simple solution for this, where when the Lambda awakens, it can fetch the secret key and authenticate the jwt.
## [4][Is it possible to reset AWS?](https://www.reddit.com/r/aws/comments/j04atj/is_it_possible_to_reset_aws/)
- url: https://www.reddit.com/r/aws/comments/j04atj/is_it_possible_to_reset_aws/
---
I used my personal email to kickstart my company a few years ago, before moving it to its own account. Now I want to use the account for a few personal projects and, although I removed all I could see related to my previous company, I just want to make sure that I'm starting with a clean AWS account.

I thought of deleting the account and creating a new one, but apparently you cannot open new accounts with email addresses that are associated with closed accounts.

Is there a way of resetting all of AWS to defaults, or to open a new account with my personal email?
## [5][S3 pattern for web &amp; mobile application.](https://www.reddit.com/r/aws/comments/j014hj/s3_pattern_for_web_mobile_application/)
- url: https://www.reddit.com/r/aws/comments/j014hj/s3_pattern_for_web_mobile_application/
---
The pattern I've come up with is:


(1) Send user a presigned url to post/get - assuming database reads/writes are not public



(2) When user uploads to s3, trigger lambda which posts to our own server, and then creates a database reference for its path.


My problem with this is an edge case where the presigned url hasn't expired, our server went down, and the lambda either kept running or stopped trying to contact our server - which means our database isn't in sync.



How do you guys use s3 :)
## [6][What is the largest AWS environment you have worked with? How do you manage it!???](https://www.reddit.com/r/aws/comments/j03qek/what_is_the_largest_aws_environment_you_have/)
- url: https://www.reddit.com/r/aws/comments/j03qek/what_is_the_largest_aws_environment_you_have/
---
Hello,

  
I work for a SaaS company and our infrastructure is hosted in AWS. For most part until now we have been managing our cloud with homegrown scripts, terraform, and AWS dashboards.

However, we are rapidly increasing our cloud infrastructure both in EC2 as well as server less side.   
I'm looking for the pointers from the community about when to start looking for third party solutions for management, if at all. 

Also, I'm looking for some general information about AWS environment that you have personally worked with. How many resources (EC2 Instances, Lambda instances, RDS, ECS cluster sizes - number of containers) etc? How many accounts? Did you use any third party solution,? If yes, could you please mention those too. Idea is I want to compare it with our environment and see what fits to us.

We recently started experimenting with Cloud Custodian - but at times we are facing API throttling issues. This worries me as it may impact our workloads too.

Thanks in advance.
## [7][AWS Service Catalog | Feedback and Experiences Using?](https://www.reddit.com/r/aws/comments/j05huf/aws_service_catalog_feedback_and_experiences_using/)
- url: https://www.reddit.com/r/aws/comments/j05huf/aws_service_catalog_feedback_and_experiences_using/
---
Looking into using AWS Service Catalog for a large enterprise.  Wanted to hear feedback, experiences, pros/cons, challenges, gotchas, missing features, use cases?

Thanks in advance for any help!
## [8][Best practise for deployments](https://www.reddit.com/r/aws/comments/j04cpc/best_practise_for_deployments/)
- url: https://www.reddit.com/r/aws/comments/j04cpc/best_practise_for_deployments/
---
Hey all 👋

I’ve been asked with getting various applications up and running in our new AWS Environment.  I’m looking to build Jenkins, Vault etc (I know, why not use Codepipeljne/KMS etc, it’s a migration at the moment, well be swapping to these in the future).

I’m looking to build this up via CDK, but unsure on the best way to approach, which of these would you suggest?

1. Building the instances and configuring with Ansible (either tower, or Systems Manager - playbooks / roles will be stored in GitHub)
2. Build the AMI with Jenkins Master/Slave baked into the image? 

Ideally I’d like this to be immutable, slaves with be in a ASG to auto scale when needed. 

Many thanks,
## [9][Question on S3 lifecycle policies](https://www.reddit.com/r/aws/comments/j0462w/question_on_s3_lifecycle_policies/)
- url: https://www.reddit.com/r/aws/comments/j0462w/question_on_s3_lifecycle_policies/
---
I have a small collection of stuff (few hundred GB) that I keep on a Synology in my house, and I want to keep an offline copy in AWS. I don't expect to access it unless the Synology goes kablooey, so Glacier Deep Archive is fine with me.  I'm confused about lifecycle policies, though.

I have an S3 bucket and upload music to it via Cloud Sync.  I want to transition stuff out of this bucket and into glacier, so I have a policy which transitions everything in the bucket into deep archive after 7 days.  I know this might be a (relatively) expensive way to do it, but even the expensive way is like $3 so I don't care.  But after I've moved everything to deep archive, I want to delete all versions in the S3 bucket so I'm not storing stuff there permanently. 

If I create a lifecycle policy to transition to deep archive, do I also need to expire current and previous versions from the S3 bucket, or does the transition to deep archive remove things from the bucket as part of the policy?  I assume "transition" means "remove from one place and put into another" but ass-u-me and all that.

Thanks!
## [10][Can we make EC2 instances in the web tier as Private?](https://www.reddit.com/r/aws/comments/j04134/can_we_make_ec2_instances_in_the_web_tier_as/)
- url: https://www.reddit.com/r/aws/comments/j04134/can_we_make_ec2_instances_in_the_web_tier_as/
---
We have Typical 3 tier architecture having Web, App and DB. Can we make EC2 instances in the web tier as Private? and allow incoming traffic only through ALB? AFAIK we can apply an SG only allowing connections from the SG of the ALB. But What if our Private EC2 instance has to return response back to the client? How it'll be routed through ALB as ALB is mostly used for managing incoming traffic. Also for outgoing traffic can we configure something like Private EC2 instance -&gt; ALB -&gt; Internet? If yes then how? So, is there any way for private EC2 instances to communicate to internet without assigning them public IP?
## [11][Cross region Cloudformation. Now you can do it with a single file!](https://www.reddit.com/r/aws/comments/izihe8/cross_region_cloudformation_now_you_can_do_it/)
- url: https://www.reddit.com/r/aws/comments/izihe8/cross_region_cloudformation_now_you_can_do_it/
---
A few days ago I wanted to use a single file (ie a single thing for people to update in a single place) to create some resources in different regions. I started down a few dead end ideas like SSM parameters and the like. They're all region locked.

I could have scripted it, but people struggle to set up roles/etc. on their CLI so console access was preferred.

Eventually I found that AWS have added StackSets as Cloudformation objects a few days previously.

To me this is a game changer. We can now supply code inline in a template for a second template that includes variables defined in the upper level template. I'd be interested to hear other uses for that. My first thought was a template that creates an update failed alarm on the sub-stack. BUT you can't yet get the IDs of the sub-stacks to know what that alarm needs to look at. (without a lamdba/macro).

BUT now, with one file you can create resources that reference other region's resources. If you've ever tried to Cloudformation ACM and Cloudfront, you know what I'm talking about. ("But the cert needs to be in us-east-1 even though all my other infrastructure isn't in us-east-1.. Wut?")

If you want more detail I made a blog post with a simple example and a run down of some of the other things you can do to achieve this (the StackSet doesn't quite work for all scenarios, I have some alternative ideas that \_might\_ help which I still need to test! At the moment it is "parent can pass values to children only". You can't pass values between children or back to the parent.)  


[https://surevine.com/creating-cloudformation-stacks-in-multiple-aws-regions-with-common-resources/](https://surevine.com/creating-cloudformation-stacks-in-multiple-aws-regions-with-common-resources/)
