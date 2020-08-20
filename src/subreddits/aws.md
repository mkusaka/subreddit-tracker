# aws
## [1][AMA - I miss working in AWS after switching to GCP](https://www.reddit.com/r/aws/comments/id3suu/ama_i_miss_working_in_aws_after_switching_to_gcp/)
- url: https://www.reddit.com/r/aws/comments/id3suu/ama_i_miss_working_in_aws_after_switching_to_gcp/
---
Me: Software Engineer who designs systems &amp; supporting infrastructure, 5 years exp using AWS, joined new company which is all in on GCP (1 year). The company I work for is enterprise-sized &amp; petabyte scaled.

-----------------

Basically, I miss working with AWS' extensive services which are mature, feature rich and have really good documentation. AWS is just miles ahead of GCP, imo.

Ask Me Anything!
## [2][Introducing the AWS Controllers for Kubernetes (ACK)](https://www.reddit.com/r/aws/comments/ictvt3/introducing_the_aws_controllers_for_kubernetes_ack/)
- url: https://aws.amazon.com/blogs/containers/aws-controllers-for-kubernetes-ack/
---

## [3][HowTo: Setup up a virtual call centre in under 30 mins using AWS Connect](https://www.reddit.com/r/aws/comments/id3gnn/howto_setup_up_a_virtual_call_centre_in_under_30/)
- url: https://medium.com/@joel.lutman/set-up-a-virtual-call-centre-in-30-minutes-with-amazon-connect-2771bbc52ee
---

## [4][RDS Mysql admin account privileges](https://www.reddit.com/r/aws/comments/ida71c/rds_mysql_admin_account_privileges/)
- url: https://www.reddit.com/r/aws/comments/ida71c/rds_mysql_admin_account_privileges/
---
Hi, I have created a Mysql RDS instance with the master user account named 'admin'.  Shouldn't this account have DBA access?  I am trying to run an infile script and am running into privileges errors.  When I look at my users my admin account does not have dba privileges however the rdsadmin account does.  I don't know the password for 'rdsadmin' because it was set as admin.  I'm very confused at how I can access a dba role on an RDS if it doesn't give you this as the master account?  Hoping someone can help my understand.
## [5][EFS for Shared Web Code (PHP)](https://www.reddit.com/r/aws/comments/id9l2b/efs_for_shared_web_code_php/)
- url: https://www.reddit.com/r/aws/comments/id9l2b/efs_for_shared_web_code_php/
---
I'm currently running my own NFS server for all my shared web PHP code between different autoscaling groups of Web servers, but I'd like to move this to EFS.

Previous "attempts" to do this always resulted in poor performance metrics and I've abandoned the idea more than once.  But this was a while back, and I know that EFS has made great strides in performance configuration options to support such a scenario.  

Anyone doing this?  If so, what EFS configuration are you using to accomplish this?  The previous process to increase EFS performance was to create huge "dummy files" to provision additional storage throughput.

Would love to hear who's doing this and how they've configured EFS for this type of workload.
## [6][SQL to DynamoDB Transformation Trainer](https://www.reddit.com/r/aws/comments/id9k20/sql_to_dynamodb_transformation_trainer/)
- url: https://bukov-ka.github.io/dynamodb/
---

## [7][VPN access to internal load-balancer blocked by security group ingress rule](https://www.reddit.com/r/aws/comments/id8whz/vpn_access_to_internal_loadbalancer_blocked_by/)
- url: https://www.reddit.com/r/aws/comments/id8whz/vpn_access_to_internal_loadbalancer_blocked_by/
---
Heya!

I'm experiencing a somewhat unexpected issue when only allowing VPC traffic to reach port 443 on our internal load balancer. Our set-up is the following:

We have a VPN connection into the VPC via a transit-gateway. The traffic is routed into 3 private subnets (one in each AZ). We have an internal load-balancer in those private subnets, which does not have a public IP and whose public dns name definitely (checked a million times) resolves to a private IP address that corresponds to the cidr range of the private subnets. One can only connect to the load-balancer via the vpn tunnel (or from within the VPC) where those private ip addresses are routed through. So everything works as expected here.

If, however, we configure the load-balancer's security group to only allow the VPC's cidr range to reach port 443, requests through the vpn tunnel start to time-out - that seems quite counter-intuitive to me, since in my understanding both the private subnets as well as the fact that the load-balancer is of type internal already ensure that only traffic from within the VPC can reach it. Am I missing something here? 

Many thanks in advance!
## [8][End to End Machine Learning Tutorial — From Data Collection to Deployment on AWS 🚀](https://www.reddit.com/r/aws/comments/ickq5y/end_to_end_machine_learning_tutorial_from_data/)
- url: https://www.reddit.com/r/aws/comments/ickq5y/end_to_end_machine_learning_tutorial_from_data/
---
\[Updated Post: End To End machine learning and beyond notebooks 🚀\]

Hello everyone!

With the great collaboration of a friend of mine, we built and deployed a machine learning application to AWS using Python.

To tell you more about this fun journey, we wrote a post where we'll walk you through the steps to:

* Collect and scrape data with Scrapy / Selenium
* Train a deep character CNN for (English) sentiment analysis using PyTorch
* Build an interactive web app with Dash to serve the model in real-time
* Put everything in Docker Compose
* Deploy to AWS on a custom domain name

Here the link: [https://medium.com/datadriveninvestor/end-to-end-machine-learning-from-data-collection-to-deployment-ce74f51ca203](https://medium.com/datadriveninvestor/end-to-end-machine-learning-from-data-collection-to-deployment-ce74f51ca203)

All the code is available in Github, the blocks are independent and reusable in other projects [https://github.com/MarwanDebbiche/post-tuto-deployment](https://github.com/MarwanDebbiche/post-tuto-deployment)

Feel free to reach out to me if you have any question,
## [9][Who knows how to work with AWS China?](https://www.reddit.com/r/aws/comments/icy6hu/who_knows_how_to_work_with_aws_china/)
- url: https://www.reddit.com/r/aws/comments/icy6hu/who_knows_how_to_work_with_aws_china/
---
I've worked very hard to get our AWS Organizations and control tower squared away for central account management.  We have great cost optimization and security practices being enforced and life was good... until I realized my company had a big AWS account in China.  I figured I would just get that added to our current master payer, start getting the spend credit applied to our EDP, and implement all of our same policies centrally managed.  Boy was I wrong!  It turns out that with AWS China accounts you do not even get the root account as they are retained by the reseller.  This means I can't just link them in like a regular account.

I'm not exactly sure how this all works but does someone have experience managing accounts in China?  It looks like I might have to ask the vendor to create a separate master payer account where I have to mirror everything that I have setup globally.  Separate control tower, SSO, security hub, etc.  I want to go to the vendor and make sure I ask for the right thing but I am not quite sure what that is.  Has anyone played this game before?
## [10][Foreign data wrapper for PostgreSQL](https://www.reddit.com/r/aws/comments/icxoqh/foreign_data_wrapper_for_postgresql/)
- url: https://www.reddit.com/r/aws/comments/icxoqh/foreign_data_wrapper_for_postgresql/
---
Is there a FDW for S3 in PostgreSQL (RDS)? I'd like to be able to create an indexed MTQ using the S3 FDW without loading the data into a table. I read that this capability is on the roadmap but haven't heard if it has been implemented yet.
