# aws
## [1][S3 Path Deprecation Plan Updated](https://www.reddit.com/r/aws/comments/iyd84u/s3_path_deprecation_plan_updated/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [2][Cross region Cloudformation. Now you can do it with a single file!](https://www.reddit.com/r/aws/comments/izihe8/cross_region_cloudformation_now_you_can_do_it/)
- url: https://www.reddit.com/r/aws/comments/izihe8/cross_region_cloudformation_now_you_can_do_it/
---
A few days ago I wanted to use a single file (ie a single thing for people to update in a single place) to create some resources in different regions. I started down a few dead end ideas like SSM parameters and the like. They're all region locked.

I could have scripted it, but people struggle to set up roles/etc. on their CLI so console access was preferred.

Eventually I found that AWS have added StackSets as Cloudformation objects a few days previously.

To me this is a game changer. We can now supply code inline in a template for a second template that includes variables defined in the upper level template. I'd be interested to hear other uses for that. My first thought was a template that creates an update failed alarm on the sub-stack. BUT you can't yet get the IDs of the sub-stacks to know what that alarm needs to look at. (without a lamdba/macro).

BUT now, with one file you can create resources that reference other region's resources. If you've ever tried to Cloudformation ACM and Cloudfront, you know what I'm talking about. ("But the cert needs to be in us-east-1 even though all my other infrastructure isn't in us-east-1.. Wut?")

If you want more detail I made a blog post with a simple example and a run down of some of the other things you can do to achieve this (the StackSet doesn't quite work for all scenarios, I have some alternative ideas that \_might\_ help which I still need to test! At the moment it is "parent can pass values to children only". You can't pass values between children or back to the parent.)  


[https://surevine.com/creating-cloudformation-stacks-in-multiple-aws-regions-with-common-resources/](https://surevine.com/creating-cloudformation-stacks-in-multiple-aws-regions-with-common-resources/)
## [3][Amazon Elasticsearch Service now offers T3 Instances](https://www.reddit.com/r/aws/comments/iz48gt/amazon_elasticsearch_service_now_offers_t3/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/amazon-elasticsearch-service-offers-t3-instances/
---

## [4][Someone uses the AWS App?](https://www.reddit.com/r/aws/comments/izfiiq/someone_uses_the_aws_app/)
- url: https://www.reddit.com/r/aws/comments/izfiiq/someone_uses_the_aws_app/
---
Every time I use the AWS App I see it really useless, you can hardly do anything.

you feel the same?
## [5][New EC2 Experience - need New Permission to control resources?](https://www.reddit.com/r/aws/comments/izg0v0/new_ec2_experience_need_new_permission_to_control/)
- url: https://www.reddit.com/r/aws/comments/izg0v0/new_ec2_experience_need_new_permission_to_control/
---
Hello,

One of our IAM user has permissions only for starting and stopping tagged EC2 instances.

(based on this document : [https://docs.aws.amazon.com/IAM/latest/UserGuide/reference\_policies\_examples\_ec2\_tag-owner.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_examples_ec2_tag-owner.html))

I found out that in new ec2 experience, the error  occured while changing instance state saying that I need describe vpcs permission and others.

It doesn't give me any error when I switched back to old console and instance is started successfully

I had to put those action in the user's policy if I want to use new console.

Is there any API changes when it's operating on new console?

I couldn't find any documents about this.

Thank you.
## [6][What's your CICD process for API Gateway backed by Lambdas using Versions &amp; Aliases?](https://www.reddit.com/r/aws/comments/iz8nq2/whats_your_cicd_process_for_api_gateway_backed_by/)
- url: https://www.reddit.com/r/aws/comments/iz8nq2/whats_your_cicd_process_for_api_gateway_backed_by/
---

## [7][Using AWS cognito to deal with user authentication and api key token creation](https://www.reddit.com/r/aws/comments/izicyu/using_aws_cognito_to_deal_with_user/)
- url: https://www.reddit.com/r/aws/comments/izicyu/using_aws_cognito_to_deal_with_user/
---
Hello all,

We have the following use case. We have an application running on AWS where we do the authentication of users manually^[1]. We are looking to migrate to using AWS Cognito to handle the user authentication and authorization. So far this all seems pretty easy and doable. The only roadblock is the generation of api_keys. When users login into our application they have the option to generate api_keys so that they can use our developer API from their own application. Picture something like [stripe](https://stripe.com/en-gb-nl) where you can make an account and login and within the application lets you generate api keys.

Is it possible to leverage Cognito to do handle the creation of api keys (or something similar like client credentials in Oauth2) as well? The thing we tried are User Pool App Clients for every user but there is a limit of 1000 clients per user pool so it doesn't seem like this is meant to be used for every single user.

Another thing we looked at is the client credentials flow on a single app client. So we create a single app client for our application and turn on client credentials and let users login using that. However a cursory glance makes it seem like client credentials are for our own machines and not so much third party developers?

[1] With manually I mean that we have an endpoint where people sign up with a username and password, save those in an RDS and when people login we simply check if the user exists and give them a JWT token
## [8][AWS instance shortage in US-EAST-1](https://www.reddit.com/r/aws/comments/iz2mmf/aws_instance_shortage_in_useast1/)
- url: https://www.reddit.com/r/aws/comments/iz2mmf/aws_instance_shortage_in_useast1/
---
Recently we ran into big difficulties to obtain some type of instances in us-east-1 (whatever the az).

m5.xlarge, r5.4xlarge etc.. very basic one. 

Are we alone in this situation ? are all customers already reserved all instances for holidays seasons ?
## [9][Dedicated hosts Vs dedicated instances](https://www.reddit.com/r/aws/comments/izcp0w/dedicated_hosts_vs_dedicated_instances/)
- url: https://www.reddit.com/r/aws/comments/izcp0w/dedicated_hosts_vs_dedicated_instances/
---
Hi guys are dedicated hosts basically an entire physical server rack that customers can select and can choose which individual physical server in rack that the instances can be deployed on ? Is dedicated instances basically instances that are running on the same physical server on a shared rack ? Each time a dedicated instance is stopped started , it goes onto a new physical server which only the company can spin up new vms ?
## [10][Using lamba to replace bash scripts in theory.](https://www.reddit.com/r/aws/comments/izgcay/using_lamba_to_replace_bash_scripts_in_theory/)
- url: https://www.reddit.com/r/aws/comments/izgcay/using_lamba_to_replace_bash_scripts_in_theory/
---
Looking at the possibility on moving a legacy solution up to the cloud

Current situation:
Multiple times a day a dmz linux vm rsyncs some files from a third party ftp share. It then does some remote renaming operations to tag the files previously downloaded and unzips downloaded files and stores them locally. 

At different multiple times a different linux vm inside our main network running an oracle database then downloads and processes these files from the first vm to oracle readable files, injests the updates then deletes the files from the first vm

Lock files are used to stop the two processes conflicting. Honestly yes it's as bad as it sounds. 

This is all triggered by cron using bash scripts. 

I was considering the benefits of migrating the oracle database to oracle rds but would then need something to replace the function of said bash scripts. 

Very new to aws, but was thinking in theory could this be achieved with a couple of lambda functions. 

The logic would need some reworking and I am guessing it would need to be written in python and attach some s3 storage to hold the files but in theory is it possible and a good idea?
## [11][Code that refuses to let the AWS SDK initialize itself and breaks if I don't feed it access keys](https://www.reddit.com/r/aws/comments/iz51lk/code_that_refuses_to_let_the_aws_sdk_initialize/)
- url: https://www.reddit.com/r/aws/comments/iz51lk/code_that_refuses_to_let_the_aws_sdk_initialize/
---
I keep running into code written by devs who just assume that I'm going to be feeding in a key/secret pair to a library that wants to talk to AWS. I set `AWS_PROFILE` to test locally and the whole library just throws a fit and tells me to go RTFM.

Right now, I'm dealing with Ruby code that does:

     AWS::EC2.new(
        access_key_id: fetch(:ec2_access_key_id),
        secret_access_key: fetch(:ec2_secret_access_key),
        region: region
      )

And when those are `nil`, the whole thing dies trying to talk to the metadata service. On my home network.

Am I doing something wrong that I keep seeing this out in the wild, or am I doomed to keep forking source and fixing it by removing code?
