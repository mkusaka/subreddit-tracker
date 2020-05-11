# aws
## [1][Awesome AWS Workshops](https://www.reddit.com/r/aws/comments/ghfgx7/awesome_aws_workshops/)
- url: https://www.reddit.com/r/aws/comments/ghfgx7/awesome_aws_workshops/
---
Following the idea of Awesome Lists, I went about our public workshops that we could find and created this weekend project to help the community. Feel free to add your own workshops, guides, and walkthroughs.

[https://awesome-aws-workshops.com/](https://awesome-aws-workshops.com/)
## [2][What does engineering for region-failure mean?](https://www.reddit.com/r/aws/comments/ghfai8/what_does_engineering_for_regionfailure_mean/)
- url: https://www.reddit.com/r/aws/comments/ghfai8/what_does_engineering_for_regionfailure_mean/
---
Correct me if I say something wrong, but here's my understanding:

When you engineer for multiple regions, you might have active/active or active/passive or whatever. Say you've got an API on both ends, you're running EKS/ECS/Whatever in each region. You've got a Primary DB that's multi AZ in one region, and a replica in the other. 

When you engineer for multi-regions and DR, what's supposed to...go down? Are you assuming literally every single boto3 call, LB-&gt;Container request, DB connection, everything fails in a region?  

So then your engineering involves both regions checking the other region for downtime? Are there assumptions you're making on how you know to fail the app layer and db layer over completely to the other region? (Automated, I'm sure, but still). 

What about services you use to perform DR? Lambda is scoped to a region, EC2 is scoped to a region. RDS scoped to a region. Is everyone architecting solutions that constantly ping the other region for any signs of an outage to immediately promote itself?
## [3][Any reason *not* to deploy an S3 gateway in a VPC?](https://www.reddit.com/r/aws/comments/gh7lml/any_reason_not_to_deploy_an_s3_gateway_in_a_vpc/)
- url: https://www.reddit.com/r/aws/comments/gh7lml/any_reason_not_to_deploy_an_s3_gateway_in_a_vpc/
---
We're looking into making our S3 buckets only accessible from inside a VPC in the same account, and arrived at the (obvious) VPC Privatelink solution of deploying an S3 Gateway in the account and working and modifying the bucket policy for said S3 buckets.

Then we were reading more about the S3 Gateway, and started realizing "why don't we just dump an S3 gateway in all our VPC's?"

Given that there is no cost to an S3 Gateway, and that you're eliminating some internet traffic costs, is there any reason \*not\* to deploy an S3 Gateway in all our accounts? It seems like such a no-brainer we must be missing something.

(Note that it's easier for use to just dump one in all our VPC's instead of targeted VPC's actually using S3 given the way we manage our AWS Organisation.)

Thanks!
## [4][Ireland availability zone locations](https://www.reddit.com/r/aws/comments/ghnd4y/ireland_availability_zone_locations/)
- url: https://www.reddit.com/r/aws/comments/ghnd4y/ireland_availability_zone_locations/
---
Does anyone know in which town or city the availability zones in Ireland are located? All other regions are named after a city but Ireland is just named Ireland. I want to know because there is a policy in my company to name machines by which city they are in and we have EC2 instances in each of the availability zones in Ireland.
## [5][Could AWS Artifical Intelligence services be used to categorise financial transactions in a simple accounting app?](https://www.reddit.com/r/aws/comments/ghl6bk/could_aws_artifical_intelligence_services_be_used/)
- url: https://www.reddit.com/r/aws/comments/ghl6bk/could_aws_artifical_intelligence_services_be_used/
---
I'm developing a very simple web app for logging financial transactions. The next feature I want to add is for transactions entered by user (or imported from bank via open-banking API) to be automatically categorised. I could do this by coming up with a simple set of rules, which look at how previus transactions made at the same merchant were categorised. Eg if 90% of transactions made at Tesco's were categorised as Shopping, then future transactions made at Tesco's could automatically be categorised as Shopping.

Are there any AWS AI services that could do this more cleverly? The whole project is for self learning, so even if AI is not the best solution, it could be a learning opportunity for me.
## [6][RDS vs Aurora big price difference](https://www.reddit.com/r/aws/comments/gh1nqw/rds_vs_aurora_big_price_difference/)
- url: https://www.reddit.com/r/aws/comments/gh1nqw/rds_vs_aurora_big_price_difference/
---
Difference between RDS and Aurora, when hosting Java application with PostgreSQL database. 

&amp;#x200B;

https://preview.redd.it/nniyg3hdxxx41.png?width=1737&amp;format=png&amp;auto=webp&amp;s=ff09d0ba6e32beb6e3a313eb88495439a9a6adf7

Here i have my estimated pricing for using RDS. I suppose this is a separate server for actually hosting the database (hence the price) I found this alternative *Aurora*, which seemed a little bit better in regards to pricing. 

&amp;#x200B;

https://preview.redd.it/oo6b3f4mxxx41.png?width=1748&amp;format=png&amp;auto=webp&amp;s=1bb1d730d97b83474b536a3c8dfa247829c6269e

This is much cheaper and also allows for almost the same amount of data (and up 1 million requests). 

&amp;#x200B;

Can anyone explain to me the major differences in regards to these two services?
## [7][Adding conditions to lambda policy](https://www.reddit.com/r/aws/comments/ghjlfh/adding_conditions_to_lambda_policy/)
- url: https://www.reddit.com/r/aws/comments/ghjlfh/adding_conditions_to_lambda_policy/
---
Hello! I'm trying to add a condition to my lambda policy. Seeing that to do this I have to use CLI rather than console, I'm having a hard time modifying my lambda policy. 

I've read the docs and it seems that there is no explicitly stated way to add conditions, like adding principals, to my policy.

Is there something I'm missing here?
## [8][Unable to SSH into an EC2 instance](https://www.reddit.com/r/aws/comments/ghiz5s/unable_to_ssh_into_an_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/ghiz5s/unable_to_ssh_into_an_ec2_instance/
---
I'm in a bit of a bind. I have a Linux EC2 instance that I need to SSH into. I am stuck because the person who had the private key left the company. I tried using EC2 Instance Connect but SSM is not installed on this instance.  
Is there any way for me to get around this?

TIA
## [9][Changing display name when using AWS SSO?](https://www.reddit.com/r/aws/comments/ghihrr/changing_display_name_when_using_aws_sso/)
- url: https://www.reddit.com/r/aws/comments/ghihrr/changing_display_name_when_using_aws_sso/
---
We're just about to roll out AWS SSO via Azure AD, and one of the things I'll miss the most about our current method of IAM/switch roles is the display name in the upper right corner of the screen, when we use the AWS Extend Switch Roles plugin.

Has anyone found a similar plugin or method to change to a friendly display name when logging in via AWS SSO? AWSReservedSSO_XXXXXXXXX_XXXXXXXXXXXXXXX isn't the most friendly of names when easily trying to determine account I left off on.
## [10][Packagemanager for ECS?](https://www.reddit.com/r/aws/comments/ghibwb/packagemanager_for_ecs/)
- url: https://www.reddit.com/r/aws/comments/ghibwb/packagemanager_for_ecs/
---
Hello!

Is there a package manager for Amazon ECS similar to how Helm works with Kubernetes? If not, is there another way to use templates, charts or similar to configure and manage ECS, preferably from the command line?
