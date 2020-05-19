# aws
## [1][How to setup AWS Organizations with AWS SSO using G Suite as an identity provider. Made account management, centralized billing and resource sharing much easier in my own company. Hope this helps :) !](https://www.reddit.com/r/aws/comments/gmn67i/how_to_setup_aws_organizations_with_aws_sso_using/)
- url: https://medium.com/serverless-transformation/setup-aws-organizations-with-google-suite-saml-sso-7e676f5ed3e1
---

## [2][AWS YouTube deep dives and code walkthroughs](https://www.reddit.com/r/aws/comments/gmdyy9/aws_youtube_deep_dives_and_code_walkthroughs/)
- url: https://www.reddit.com/r/aws/comments/gmdyy9/aws_youtube_deep_dives_and_code_walkthroughs/
---
Hey guys, given the lockdown and an ample amount of free time I decided to create a YouTube channel and dive deep into a lot of the most popular AWS topics. Here's the channel:
https://www.youtube.com/VladimirBudilov

Here's my GitHub account with a lot of sample code: 
https://github.com/vbudilov/

Let me know if there's a topic that you're really interested in and if there's additional demand I can do a deep dive on that topic. 

Vladimir

P.S. This isn't sponsored by Amazon
## [3][A Multithreaded Fork of Redis that is 5X faster](https://www.reddit.com/r/aws/comments/gm4lve/a_multithreaded_fork_of_redis_that_is_5x_faster/)
- url: https://github.com/JohnSully/KeyDB
---

## [4][Setting up LAMP stack with load balancing - is Lightsail right for me?](https://www.reddit.com/r/aws/comments/gmohl3/setting_up_lamp_stack_with_load_balancing_is/)
- url: https://www.reddit.com/r/aws/comments/gmohl3/setting_up_lamp_stack_with_load_balancing_is/
---
I need to set up/write a very basic LAMP app -- some static pages, and some PHP / MySQL CRUD logic, all over SSL.  The catch is it needs to be highly available, lots of inserts to MySQL, and it may need load balancing in the future.

I've run regular EC2 instances with no load balancers for other stuff and am comfortable with that -- but I want to get the architecture for this app correct right off the bat.  It seems Lightsail would be ideal for this, but I'm wondering what other alternatives there are?
## [5][Sagemaker Ground Truth - Cannot get active learning to auto label my data](https://www.reddit.com/r/aws/comments/gmnmmr/sagemaker_ground_truth_cannot_get_active_learning/)
- url: https://www.reddit.com/r/aws/comments/gmnmmr/sagemaker_ground_truth_cannot_get_active_learning/
---
Hi guys,

I’m trying to use Sagemaker Groundtruth’s active learning capability to label some data but can not figure out how to get the auto-labeling part to work. Do I need to start with 1,000 labeled objects in a labeling job? Do I need to have an initial model? Thanks guys.
## [6][Load tests with Gatling's Frontline on AWS](https://www.reddit.com/r/aws/comments/gmngz5/load_tests_with_gatlings_frontline_on_aws/)
- url: https://www.reddit.com/r/aws/comments/gmngz5/load_tests_with_gatlings_frontline_on_aws/
---
I deployed an infrastructure in five days on AWS (CloudFrint, WAF, S3, API-Gateway, NLB, EKS, RDS). This infrastructure had to sustain 50 thousand calls in one minute on the backend. To make sure everything would go smoothly when we go live, I deployed a Gatling frontline instance. [Here is how I have done it](https://www.padok.fr/en/blog/load-test-gatling-frontline-aws). If you have experience with another tool, I would love to hear it, please share it with me in comments.
## [7][Api gateway private integration](https://www.reddit.com/r/aws/comments/gmmyu7/api_gateway_private_integration/)
- url: https://www.reddit.com/r/aws/comments/gmmyu7/api_gateway_private_integration/
---
Hi all,

I'm trying to setup a public api gateway with a private integration to forward requests to our on-premise services. I'm using vpc link and an internal network load balancer within a vpc to achieve this, I also want the tls termination to occur at the nlb layer. I've setup the https listener on the nlb and configured a certificate to use.  


When I go to test the connection, I get this error: 

`Tue May 19 10:21:25 UTC 2020 : Execution failed due to configuration error: Host name 'vpce-xxxxxxxxxxxxxxxxx-xxxxxxxx.vpce-svc-xxxxxxxxxxxxxxxx.eu-west-1.vpce.amazonaws.com' does not match the certificate subject provided by the peer`  


Has anyone else encountered this problem? Is it simply a case of adding `*.vpce-svc-xxxxxxxxxxxxxxxx.eu-west-1.vpce.amazonaws.com` as an alternate in the certificate, or configuring a cname in my dns?
## [8][Why am I being charged for DynamoDB read and write capacity under the free tier limit?](https://www.reddit.com/r/aws/comments/gmgthz/why_am_i_being_charged_for_dynamodb_read_and/)
- url: https://www.reddit.com/r/aws/comments/gmgthz/why_am_i_being_charged_for_dynamodb_read_and/
---
I am being charged money daily for DynamoDB "WriteCapacityUnit-Hrs" and "ReadCapacityUnit-Hrs". However, my DynamoDB tables only use a maximum of 25 provisioned capacity units for both read and write, which is under the free tier limit. I am even being charged when i'm not actually consuming any units at all. You can see it in the images. Why am I being charged for read and write capacity units if I am still under the free tier limit?

&amp;#x200B;

[Being charged for read and write capacity unity hours](https://preview.redd.it/yxqr1vr95nz41.png?width=3787&amp;format=png&amp;auto=webp&amp;s=2b54519e4a82569a16d001dbb53683ea0d850619)

&amp;#x200B;

[My tables only use a total of 25 provisioned capacity \(free tier limit\)](https://preview.redd.it/nff6jv3c5nz41.png?width=2555&amp;format=png&amp;auto=webp&amp;s=26c0c1a589a15d628d60b600c11a49bcd56c7a08)
## [9][This Week in DevOps](https://www.reddit.com/r/aws/comments/gmm72z/this_week_in_devops/)
- url: https://www.reddit.com/r/aws/comments/gmm72z/this_week_in_devops/
---
This Week in DevOps – AWS reduced prices on some reserved instances and inter-region data transfer for some regions. They also rolled out a new Enterprise Search service called Kendra and a new class of instances called M6g which are based on Graviton2. Spot Virtual Machines came to Azure this week, Digital Ocean opened a new data center in San Francisco and HashiCorp released consul 1.8 for public beta.

Get the details and more here: [https://thisweekindevops.com/2020/05/15/weekly-roundup-may-15th-2020/](https://thisweekindevops.com/2020/05/15/weekly-roundup-may-15th-2020/)
## [10][Access EC2 (behind VPC) by SSM](https://www.reddit.com/r/aws/comments/gmcqdo/access_ec2_behind_vpc_by_ssm/)
- url: https://www.reddit.com/r/aws/comments/gmcqdo/access_ec2_behind_vpc_by_ssm/
---
Hi all,  


I'm working on a project that needs an EC2 instance to poll a website (in this case, PRAW for Reddit), and it will need to create SQS messages. It's my understanding (I could be wrong) that the best way to do this is to put the EC2 behind a VPC along with all the other services that will be running (elasticache, aurora).  


Where I'm running into trouble is I'm trying to make this EC2 instance a managed service, but it is not working. I made sure it is on the proper VPC, that there are security groups with endoints for SSM, SSMMessages, EC2, EC2 messages ([https://docs.aws.amazon.com/systems-manager/latest/userguide/setup-create-vpc.html](https://docs.aws.amazon.com/systems-manager/latest/userguide/setup-create-vpc.html)), and I added the default SSM Quickstart role for the EC2 instance.  


Still, I cannot start an SSM session, nor does it show up on the managed instances. My other EC2 instance (on the default VPC) shows up, no problem and I can connect.

Any advice or suggestions for trouble shooting? Should I even be using a VPC for this kind of work?  


Thanks! And let me know if you'd like to hear more about the project :)
