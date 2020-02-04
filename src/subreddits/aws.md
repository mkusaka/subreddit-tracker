# aws
## [1][Password-less Amazon Cognito Registration and Authentication (through email or SMS)](https://www.reddit.com/r/aws/comments/eyhj2m/passwordless_amazon_cognito_registration_and/)
- url: https://aws.amazon.com/blogs/mobile/implementing-passwordless-email-authentication-with-amazon-cognito/
---

## [2][Getting Started with CloudWatch Synthetics - Canary testing on AWS](https://www.reddit.com/r/aws/comments/eyopgb/getting_started_with_cloudwatch_synthetics_canary/)
- url: https://zoph.me/posts/2020-02-02-cloudwatch-synthetics/
---

## [3][EC2 Auto scaling group seems to ignore weights](https://www.reddit.com/r/aws/comments/eyorjx/ec2_auto_scaling_group_seems_to_ignore_weights/)
- url: https://www.reddit.com/r/aws/comments/eyorjx/ec2_auto_scaling_group_seems_to_ignore_weights/
---
I'm attempting to use the new weighting option in ASG -  [https://aws.amazon.com/about-aws/whats-new/2019/11/amazon-ec2-auto-scaling-supports-instance-weighting/](https://aws.amazon.com/about-aws/whats-new/2019/11/amazon-ec2-auto-scaling-supports-instance-weighting/) 

I've assigned weights to the server options I added to the ASG, but they seem to be ignored.

I set the t3a.small weight to 2 and t3a.medium to 6 and when I set desired capacity to 1or 2 it always starts a t3a.medium. I would have expected it to start a small which would fulfill the capacity requirement.

I've tried adding t2 and t3 types too as well as adding large types and trying to get it to start smaller instances than the largest ones based on desired capacity, but no joy.

Any ideas how to use capacity to get ASG to launcher smaller instances?
## [4][Configuring Federation LAB for AWS SSO using AD and ADFS](https://www.reddit.com/r/aws/comments/eyo2qp/configuring_federation_lab_for_aws_sso_using_ad/)
- url: https://www.reddit.com/r/aws/comments/eyo2qp/configuring_federation_lab_for_aws_sso_using_ad/
---
Hey guys, I have a bare metal server which can connect to the internet. I can browse to AWS website and can login using an IAM user.

Recently I thought of trying to configure a lab on it with Active Directory and ADFS so I can create some users on premises and practice SSO.

So I configured a local AD environment on a Windows 2012 R2 server, "[mydomain.com](https://mydomain.com)" and I followed the instructions outlined in the article below. I can get to the Sign-In page but when I try to sign-in, I get a 400 bad request.

[https://aws.amazon.com/blogs/security/enabling-federation-to-aws-using-windows-active-directory-adfs-and-saml-2-0/](https://aws.amazon.com/blogs/security/enabling-federation-to-aws-using-windows-active-directory-adfs-and-saml-2-0/)

Do I have to forward some port on my AD server? I don't think there's anything else wrong with it as I can browse to the sites but I am not sure if I have to forward/open port on the server??? Has anyone tried this before.
## [5][awscredx: AWS role assumption made simple](https://www.reddit.com/r/aws/comments/eyo2lb/awscredx_aws_role_assumption_made_simple/)
- url: https://github.com/sam701/awscredx
---

## [6][Node.js AWS SDK for creating AWS EC2/ECS Instance](https://www.reddit.com/r/aws/comments/eyn4ih/nodejs_aws_sdk_for_creating_aws_ec2ecs_instance/)
- url: https://www.reddit.com/r/aws/comments/eyn4ih/nodejs_aws_sdk_for_creating_aws_ec2ecs_instance/
---
Does anyone here knows how to start a EC2/ECS instance using aws sdk on node.js? Is it possible to use a docker run command as well when running a docker image/container? Any working sample codes will be a lot of help. Thanks in advance!
## [7][Find cleartext passwords stored in Lambda Functions](https://www.reddit.com/r/aws/comments/eyg293/find_cleartext_passwords_stored_in_lambda/)
- url: https://www.reddit.com/r/aws/comments/eyg293/find_cleartext_passwords_stored_in_lambda/
---
Hello! Is anyone aware of a tool (preferably open source) I can use to periodically scan the code of all of the functions in an AWS account and determine if there are any passwords/high-entropy strings stored in the code of the function?

Eventually I would like to also run this as part of the CI/CD process, but for now I need to get an idea of the different risks we have in our current environment. I have looked at LambdaGuard, but it seems to be lacking this functionality. I also looked at integrating SonarQube with LambdaGuard which worked well, but still doesn't provide this level of information.

I am thinking of just writing something that does this, but want to make sure I'm not missing out on something that is already available for this purpose.

Thanks!
## [8][Has anyone hit the maximum limit of 50 tags?](https://www.reddit.com/r/aws/comments/eyly1i/has_anyone_hit_the_maximum_limit_of_50_tags/)
- url: https://www.reddit.com/r/aws/comments/eyly1i/has_anyone_hit_the_maximum_limit_of_50_tags/
---
We all rejoiced when the limit was increased from 10 to 50 (at least I did), but it makes me wonder, is anyone getting any closer to the upper limit?  What's the most you're using in your environment, and why?  30? 40? 45? 48?  

Have you hit the ceiling yet? :)
## [9][Increasing web delivery speeds across regions.](https://www.reddit.com/r/aws/comments/eylwg8/increasing_web_delivery_speeds_across_regions/)
- url: https://www.reddit.com/r/aws/comments/eylwg8/increasing_web_delivery_speeds_across_regions/
---
Hi,


I have an instance r4.4xLarge of SAP BI Platform running in Oregon US-WEST-2 serving a web-app (Lumira 2.3). The web-app is quite heavy and takes 40 seconds to load to a client in Taiwan, Asia. Are there any AwS network boosting options to reduce the latency and increase load times?



Thanks.
## [10][CDK CI/CD](https://www.reddit.com/r/aws/comments/eylswc/cdk_cicd/)
- url: https://www.reddit.com/r/aws/comments/eylswc/cdk_cicd/
---
Starting to use CDK a little more and wondering if others have implemented a CI/CD pipeline with CDK?

Would be nice to have CDK code build, test and deploy and integrate with some other workflows if possible. 

Thanks.
