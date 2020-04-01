# aws
## [1][Should I switch my ALB to use the Least Outstanding Requests algorithm? YES!](https://www.reddit.com/r/aws/comments/fsy2sq/should_i_switch_my_alb_to_use_the_least/)
- url: https://medium.com/dazn-tech/aws-application-load-balancer-algorithms-765be2eca158?source=friends_link&amp;sk=2da467951a30524ef398b9b333707d43
---

## [2][Considering using AWS to host a Minecraft server. Would this be a bad idea?](https://www.reddit.com/r/aws/comments/fss6nx/considering_using_aws_to_host_a_minecraft_server/)
- url: https://www.reddit.com/r/aws/comments/fss6nx/considering_using_aws_to_host_a_minecraft_server/
---
I always run servers on my own computer, but I have a group of fifteen friends who all want to play on a server that would be running 24/7 and I can't commit to that for my own computer. We've used MC hosting services in the past, but I feel like using something like AWS would be better because there wouldn't be any middleman. Is there any reason why I shouldn't use AWS? If I should use it, do you have any suggestions?

Edit: Our current option costs $30/month and we're looking to reduce that.
## [3][Introducing Amazon Detective Following The Breadcrumbs](https://www.reddit.com/r/aws/comments/fsp836/introducing_amazon_detective_following_the/)
- url: https://engineering.kablamo.com.au/posts/2020/amazon-detective-following-the-breadcrumbs
---

## [4][AWS Global Accelerator: Terrible Name, Awesome Service](https://www.reddit.com/r/aws/comments/fsba7m/aws_global_accelerator_terrible_name_awesome/)
- url: https://www.sentiatechblog.com/aws-global-accelerator-terrible-name-awesome-service?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=global_accelerator
---

## [5][CDK Templates for Elasticsearch, Kinesis Streams, DynamoDB, Cognito, and Lambda](https://www.reddit.com/r/aws/comments/fskejj/cdk_templates_for_elasticsearch_kinesis_streams/)
- url: https://github.com/vbudilov/aws-cdk-starter
---

## [6][Do you delete the default VPC?](https://www.reddit.com/r/aws/comments/fshd0f/do_you_delete_the_default_vpc/)
- url: https://www.reddit.com/r/aws/comments/fshd0f/do_you_delete_the_default_vpc/
---
After setting up your VPC do you delete the default AWS one? I've read conflicting information on it.
## [7][AWS NOOB: AWS EC2 or RDS?](https://www.reddit.com/r/aws/comments/fsztmm/aws_noob_aws_ec2_or_rds/)
- url: https://www.reddit.com/r/aws/comments/fsztmm/aws_noob_aws_ec2_or_rds/
---
Our DB needs to be accessible 24/7, but it has low load.  We are a company of 10 people using the DB during an 8-hour workday, but we have some users (maybe 50) lightly touching the system outside of business hours.

Right now, we have a SQL Server 2012 on-premises.  The physical device is at least 5 years old. The VM has 64 GB RAM.  The total DB size is 200GB.  I don't know the total data throughput or the total number of queries.  It's been running fine like this for years.  I think we're over specced at 64 GB of RAM, but since we had it available, we assigned it. We use no stored procedures or T-SQL.  Just a simple RDMS that serves our web site which is our primary CMS/operations system.

I'm looking to migrate to AWS while spending as little money as I can feasibly spend, but with a solution that I won't have to worry about and will be fast. I am not a DBA... I'm a hack programmer.

Here's what I'm considering...

**Convert to either postgresql or mysql**

I've used mysql before and have been comfortable with it.  I'm a little worried that some of the queries that have been working in sql server will break.  Perhaps I might've unknowingly used some sql server specific function or syntax. I'm considering this to avoid paying sql server licensing fees.  Is this wise?

**Server options**

I'm using 16 GB RAM as a baseline for comparison.  

\-Get an EC2 instance and run that as our DB server.  16GB Linux instance costs about $140 per month. 

\-An on-demand RDS instance - m4.2xlarge has 16 GB ram and costs about $280 / mo.

\-Aurora Serverless - 100% utilization with 8 ACU 16 GB RAM costs $371.   
**My questions**

Are these reasonable solutions?  Are my price estimates accurate?  Which of these sounds like the best solution for my use-case?    
Thanks so much for your help.  Love you all.
## [8][Elasticbeanstalk single docker environment, customized awslogs overwritten on ConfigDeploy.](https://www.reddit.com/r/aws/comments/fszsnn/elasticbeanstalk_single_docker_environment/)
- url: https://www.reddit.com/r/aws/comments/fszsnn/elasticbeanstalk_single_docker_environment/
---
I've noticed an issue when updating config on Elasticbeanstalk environment and I'm looking for a clean way to overcome it.

When deploying configuration changes,  my customized awslogs.conf gets overwritten with addon.

&gt;AddonsBefore/ConfigCWLAgent/10-config.sh

When I deploy my app, I use container command to modify the logs and container command is running after the addons therefore awslogs.conf is the way I want it, that is not the case on ConfigUpdate.

Consequence is that my logs are not streamed anymore. As per AWS documentation using files in post hooks is not advised.

Anyone has a suggestion how can I resolve this?
## [9][Making billing dashboard publically accessible?](https://www.reddit.com/r/aws/comments/fsz4rc/making_billing_dashboard_publically_accessible/)
- url: https://www.reddit.com/r/aws/comments/fsz4rc/making_billing_dashboard_publically_accessible/
---
Hey, I want to make it transparent to my customers how much the infrastructure costs. Is there any way how I can do it? Preferable the whole dashboard.
## [10][Publish WebSite with https in AWS nightmare](https://www.reddit.com/r/aws/comments/fsz2rk/publish_website_with_https_in_aws_nightmare/)
- url: https://www.reddit.com/r/aws/comments/fsz2rk/publish_website_with_https_in_aws_nightmare/
---
I can see the website with [http://www.devops.engineering/](http://www.devops.engineering/)

&amp;#x200B;

but I am unable to use [https://www.devops.engineering/](https://www.devops.engineering/)

I am getting **www.devops.engineering** unexpectedly closed the connection.

For certificate

&amp;#x200B;

https://preview.redd.it/2oxvvu32c7q41.png?width=345&amp;format=png&amp;auto=webp&amp;s=9caab69219fa7cf6dee1d94498745d9efad35979

&amp;#x200B;

What did I do?

Created a certificated using and passed validation

[https://console.aws.amazon.com/acm/home?region=us-east-1#/](https://console.aws.amazon.com/acm/home?region=us-east-1#/)

&amp;#x200B;

https://preview.redd.it/ob5x2xa8c7q41.png?width=1911&amp;format=png&amp;auto=webp&amp;s=30a784a7169db48168c382228ce18b7280bf09be

i checked the routes

[https://console.aws.amazon.com/route53/home#resource-record-sets:Z2QK9LSHRP09KZ](https://console.aws.amazon.com/route53/home#resource-record-sets:Z2QK9LSHRP09KZ)

&amp;#x200B;

https://preview.redd.it/kogxfuyac7q41.png?width=1557&amp;format=png&amp;auto=webp&amp;s=f8affe1f0c718b6103f198cc0a39e5f8ce88e29e

I have created the CloudFront distributions

[https://console.aws.amazon.com/cloudfront/home#distributions:](https://console.aws.amazon.com/cloudfront/home#distributions:)

https://preview.redd.it/jfybem5ec7q41.png?width=1869&amp;format=png&amp;auto=webp&amp;s=7fff009ed078e7e1f2b81796db07049cbdfb4e89

One with a redirect

&amp;#x200B;

https://preview.redd.it/q960p4dhc7q41.png?width=1212&amp;format=png&amp;auto=webp&amp;s=05603b1f49d554cf3e81685c44b528ca700065c3

&amp;#x200B;

https://preview.redd.it/78vj95uic7q41.png?width=1313&amp;format=png&amp;auto=webp&amp;s=712a6bde323142fcb8c530bf081dbd6df9db5de2
