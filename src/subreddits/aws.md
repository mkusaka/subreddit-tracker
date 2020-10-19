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
## [2][Is it possible to replicate an entire environment? Or copy every setting from every service?](https://www.reddit.com/r/aws/comments/je0thd/is_it_possible_to_replicate_an_entire_environment/)
- url: https://www.reddit.com/r/aws/comments/je0thd/is_it_possible_to_replicate_an_entire_environment/
---
Essentially I have an environment setup with every service I need, (VPCs / subnets / code pipeline / code build / subnets / DNS etc etc) and I want a QA environment with the exact same settings.

Is there a known way to do this?
## [3][S3 replicate](https://www.reddit.com/r/aws/comments/je1a1j/s3_replicate/)
- url: https://www.reddit.com/r/aws/comments/je1a1j/s3_replicate/
---
Hi devs,

I using S3 with a PHP system. Has some way to make a "backup" from the production bucket? Replicate it to new bucket with diferent type - like glacier - could be a good option?
## [4][SNS automated statistics via email(?)](https://www.reddit.com/r/aws/comments/jdzzbt/sns_automated_statistics_via_email/)
- url: https://www.reddit.com/r/aws/comments/jdzzbt/sns_automated_statistics_via_email/
---
Hi,

We use an SNS Topic to send out SMS to customers on a daily basis.

I wanted to know if there is a way to get an automated statistics summary/report sent via email - to save users having to sign in and navigate to the S3 bucket where the daily usage reports are? (not much of a hardship I know)

Would also mean we don't have to create Users in IAM for those that just wanted to see stats.

Thanks!
## [5][Can I host a JavaScript website using S3?](https://www.reddit.com/r/aws/comments/jdj2ox/can_i_host_a_javascript_website_using_s3/)
- url: https://www.reddit.com/r/aws/comments/jdj2ox/can_i_host_a_javascript_website_using_s3/
---
Hi I am completely new to AWS and found [this](https://cloudirregular.substack.com/p/the-cloud-resume-challenge) resume challenge that I want to try. I also have very limited knowledge of web dev so this might be a dumb question. I read that S3 can only host "Static" web pages. Does that mean JavaScript and/or React is out of the question? I want to make my resume page have some minor CSS animations if possible.
## [6][Cognito: resend MFA code (Java)](https://www.reddit.com/r/aws/comments/jdxkyz/cognito_resend_mfa_code_java/)
- url: https://www.reddit.com/r/aws/comments/jdxkyz/cognito_resend_mfa_code_java/
---
This was asked for JavaScript library, but no solution was provided. 

Scenario:

1. user signs in with MFA enabled
2. SMS is sent and user needs to confirm challenge
3. for whatever reason user didn’t receive SMS

Is there any possibility to resend SMS without log in again?
## [7][Where do I even begin?](https://www.reddit.com/r/aws/comments/jdzdvo/where_do_i_even_begin/)
- url: https://www.reddit.com/r/aws/comments/jdzdvo/where_do_i_even_begin/
---
It’s become apparent that all of a sudden the business I work for is moving at breakneck speed up into the cloud. Everything that used to sit on bare-metal in a giant CAR (Central Apparatus Room) is going up into AWS. We have S3 buckets, EC2 instances, an RDS database supporting it and all manor of other things, and full disclosure... I am entirely lost on what it is, what it all means, how to work within it and therefore how to work, period.

I am terrified that if I don’t get a mastery of the basics soon, this will entirely pass me by and before I know it I will no longer be able to do my job and subsequently won’t be able to call myself an Engineer in the new world we have been forced to create.

Where do I start with this stuff? How do I work my way through it from the very basics?

TLDR; I’m old and new things scare me. Help.
## [8][Intermittently unable to reach web server over VPN tunnel](https://www.reddit.com/r/aws/comments/jdvz8h/intermittently_unable_to_reach_web_server_over/)
- url: https://www.reddit.com/r/aws/comments/jdvz8h/intermittently_unable_to_reach_web_server_over/
---
Hi, I have some machines running an application/website in ec2, which accesses a filestore located on premise. In order to access this filestore I have created a VPN tunnel.

Occasionally - for seemingly no reason, I am unable to hit my website from on premise. At the same time I will be able to hit it from off-premise, and every other network function still works on premise.

Is there some basic traffic I have forgotten to tunnel?

&amp;#x200B;

Thanks
## [9][Why would a connection from an EC2 instance to Aurora serverless in the same VPC / subnet group be in a totally different CIDR?](https://www.reddit.com/r/aws/comments/jdu75w/why_would_a_connection_from_an_ec2_instance_to/)
- url: https://www.reddit.com/r/aws/comments/jdu75w/why_would_a_connection_from_an_ec2_instance_to/
---
My VPC subnet is CIDR 172.31.0.0/16 and Aurora Serverless DB and an EC2 instance are both assigned to the same subnet group within it. The EC2 instance has a 172.31.x.x IP address as expected.

However when I connect from the EC2 instance the Aurora DB sees the connection as coming from 10.1.10.181 ?!? 

Has anyone encountered this before? Why is this happening?
## [10][Cloudfront + S3 static-site DNS question](https://www.reddit.com/r/aws/comments/jdp6bj/cloudfront_s3_staticsite_dns_question/)
- url: https://www.reddit.com/r/aws/comments/jdp6bj/cloudfront_s3_staticsite_dns_question/
---
Hey everyone,
I apologize that this is likely more of a failure in understanding of DNS than an issue with AWS, but I could also be wrong.

I have a Cloudfront served static-site with S3 origin, an AWS generated SSL certificate, and use a NameCheap domain.

Issue:

With my current configuration:
*domain*.com and www.*domain*.com are directed to https://www.*domain*.com where my content is rendered. 
https://*domain*.com, however returns a 'This site cannot provide a secure connection' error.


My DNS configuration:

CNAME Record, @, *mycloudfronturl*

CNAME Record, www, *mycloudfronturl*

CNAME Record,*www ssl record name*, *www ssl record value*

CNAME Record,*non-www ssl record name*, *non-www ssl record name*



According to AWS Certificate Manager, both the *domain*.com and www.*domain*.com domains have been validated, and is in use by Cloudfront.

Please let me know if there is anything else I can supply that might help.
## [11][Is possibile that AWS SES notify of MX DNS errors.](https://www.reddit.com/r/aws/comments/jdoc3d/is_possibile_that_aws_ses_notify_of_mx_dns_errors/)
- url: https://www.reddit.com/r/aws/comments/jdoc3d/is_possibile_that_aws_ses_notify_of_mx_dns_errors/
---
I was testing SES events by sending some bad emails.

When sending an email to nonexistent domain SES does not notify about it.

Does this type of failure counts to email quota?
