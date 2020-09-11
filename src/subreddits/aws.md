# aws
## [1][Elastic Beanstalk introduces support for shared load balancers!](https://www.reddit.com/r/aws/comments/iqjxcw/elastic_beanstalk_introduces_support_for_shared/)
- url: https://aws.amazon.com/blogs/containers/amazon-elastic-beanstalk-introduces-support-shared-load-balancers/
---

## [2][What is something you guys have learned while using AWS? (For me, establishing a disaster recovery plan is a must)](https://www.reddit.com/r/aws/comments/iqerhh/what_is_something_you_guys_have_learned_while/)
- url: https://www.reddit.com/r/aws/comments/iqerhh/what_is_something_you_guys_have_learned_while/
---
&amp;#x200B;

https://preview.redd.it/abdi3ga9lem51.png?width=825&amp;format=png&amp;auto=webp&amp;s=3c48af1f718e128711f6ca599a2d752c27a75bd6

In March 2017, Amazon had an outage which affected businesses in many ways, including ours. We had about 25 separate user-facing systems that were unavailable, affecting users number in the 100s of thousands. When the outage first occurred, we tried to switch to an alternate region for those systems that are in multiple regions (not all of them are), but could not because we could not use the AWS load balancing service, which was also impacted by the outage. In the end, we just had to wait for Amazon to resolve the problem and test it.

Now, I'm not saying Amazon is going to shut down. Our organization still uses and loves Amazon, but did it have an impact on our business? Absolutely.

We have obviously recovered since this but have established a well-structured business continuity plan in the case that this ever happens again and are prepared. You should too!

For those of you not familiar with business continuity plans, here is an example: [AWS Business Continuity Plan &amp; Template](https://www.allcode.com/aws-business-continuity-plan/)

What is some advice that you would recommend to somebody starting on AWS? Or even something you have learned along the way that could help some people out?
## [3][How big to size production VPC subnets with 3 tiers?](https://www.reddit.com/r/aws/comments/iqnndm/how_big_to_size_production_vpc_subnets_with_3/)
- url: https://www.reddit.com/r/aws/comments/iqnndm/how_big_to_size_production_vpc_subnets_with_3/
---
Currently I'm thinking of the following setup where everything is balanced:

* Public tier (load balancers + bastion): 3 subnets across 3 AZ with a /20 mask. (12,288 IP)
* Private tier (application servers connected to NAT gateway): 3 subnets across 3 AZ with a /20 mask. (12,288 IP)
* Restricted tier (database/cache servers): 3 subnets across 3 AZ with a /20 mask. (12,288 IP)

The application servers auto-scale with ECS Fargate. I decided to make the public tier the same size as the private tier since I may remove NAT gateways and move the app servers into the public subnet in the future. This leaves [10.0.143.255](http://10.0.143.255/)\-10.0.255.255 un-allocated. 

Anything I should consider/resize before rolling with this?
## [4][AWS SSO adds some APIs and CloudFormation support for multi-account setups](https://www.reddit.com/r/aws/comments/iqggtx/aws_sso_adds_some_apis_and_cloudformation_support/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/aws-single-sign-on-adds-account-assignment-apis-and-aws-cloudformation-support-to-automate-multi-account-access-management/
---

## [5][In praise of S3, the greatest cloud service of all time](https://www.reddit.com/r/aws/comments/iq20le/in_praise_of_s3_the_greatest_cloud_service_of_all/)
- url: https://acloudguru.com/blog/engineering/brazeal-in-praise-of-s3-the-greatest-cloud-service-of-all-time
---

## [6][Lambda function check health application](https://www.reddit.com/r/aws/comments/iqjy9q/lambda_function_check_health_application/)
- url: https://www.reddit.com/r/aws/comments/iqjy9q/lambda_function_check_health_application/
---
Let's assume I completed a python script that checks an external web application outside of AWS network. From the lambda logs, I can see http status code of 200 which is a success. I'd like to record/store the date/timestamp of the code that checked the health of the remote application as well as the http status code into cloudwatch. That way, we might be able to pull those logs from cloudwatch and be able to create a chart/graph of the health history of the web application. Is this possible?
## [7][Introducing security groups for pods](https://www.reddit.com/r/aws/comments/iq9rvu/introducing_security_groups_for_pods/)
- url: https://aws.amazon.com/blogs/containers/introducing-security-groups-for-pods/
---

## [8][Why are s3 egress costs lesser than cloudfront ?](https://www.reddit.com/r/aws/comments/iqov30/why_are_s3_egress_costs_lesser_than_cloudfront/)
- url: https://www.reddit.com/r/aws/comments/iqov30/why_are_s3_egress_costs_lesser_than_cloudfront/
---
I have the notion that when cloudfront is used the overall costs reduce because the number of origin hits decrease. But now when I see the pricing of S3 and Cloudfront, for India the egress costs of S3 are 0.1093$ while cloudfront costs 0.17$ which means I will be charged less if I just use S3. 

&amp;#x200B;

Was confused what has gone wrong here ?
## [9][Having trouble sending outbound SMTP from ec2 to office 365 - does not seem to be an issue of throttling, happens on 25 and 587](https://www.reddit.com/r/aws/comments/iqo8j4/having_trouble_sending_outbound_smtp_from_ec2_to/)
- url: https://www.reddit.com/r/aws/comments/iqo8j4/having_trouble_sending_outbound_smtp_from_ec2_to/
---
Hi, so, I am using an application that sends mail using my office 365 email address. However the mail is failing to send

Error sending email to example.user@mydomain.com, error:Could not open socket: stream\_socket\_client(): unable to connect to tcp://mydomain-com.mail.proection.outlook.com:25

I have security group rules for outbound - [0.0.0.0/0](https://0.0.0.0/0) all traffic and I have also added for smtp over port 25 from the same address, as well as port 587 tcp

I have created a connector in office 365 with the public IP of the ec2 instance used to send the mail whitelisted....but the error persists

A telnet to my mx record just hangs at 'trying'

What am I doing wrong?
## [10][In Redshift how can I create an user which which has only user creation privileges?](https://www.reddit.com/r/aws/comments/iqo8e2/in_redshift_how_can_i_create_an_user_which_which/)
- url: https://www.reddit.com/r/aws/comments/iqo8e2/in_redshift_how_can_i_create_an_user_which_which/
---
The primary requirement is to automatically create or alter users based on user information stored in a table. I am planning to write a stored procedure / use Databricks notebook to do the automation. The Redshift user which I use for creating other users in this automation should only user create/alter permissions. It is the only which should be performed from that user. I read that the option CREATEUSER option give super user like privileges to the user. How can i restrict/revoke other superuser privileges?
