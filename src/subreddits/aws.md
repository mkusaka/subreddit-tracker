# aws
## [1][My understanding of Aurora serverless](https://www.reddit.com/r/aws/comments/glb7pk/my_understanding_of_aurora_serverless/)
- url: https://www.reddit.com/r/aws/comments/glb7pk/my_understanding_of_aurora_serverless/
---
* You're just given an endpoint to which your application is supposed to connect to
* You set a few extra parameters while creating the instance like min, max size, etc
* Everything else is taken care of from aws side
* There's no concept of instances (at least from our side). There's no instance/replica/reader/writer creation/management from our side. Only things we can modify from our end are the cluster parameter groups,etc.

Can anyone confirm these? Also, I'm not sure how backups are done in serverless. Thanks :)
## [2][VPC desgin studio](https://www.reddit.com/r/aws/comments/gkz23i/vpc_desgin_studio/)
- url: https://houqp.github.io/vpcstudio/
---

## [3][How to monitor a Node.js Express API running on an EB.](https://www.reddit.com/r/aws/comments/gld7ar/how_to_monitor_a_nodejs_express_api_running_on_an/)
- url: https://www.reddit.com/r/aws/comments/gld7ar/how_to_monitor_a_nodejs_express_api_running_on_an/
---
Hi redditors, 

**TL;DR**

Asking for advice on how to monitor an Node.js Express API running on an Elastic Beanstalk.

I'm a mobile engineering. I'm writing this post because my apps are growing and we start having a decent amount of traffic. I'm new to AWS and devops related stuff and I'm lost in all the options available. I would like to introduce my use case to get your advice if possible. Thanks in advance!

**The Infrastructure**

* I have 4 mobile apps relaying on one server.
* The server is a Node.js backend running on an Elastic Beanstalk behind a Load Balancer with autoscale enabled. There is usually 2  (*t2.micro*) machines running.
* The server has an API built on express. 
* I'm using the AWS free tier whenever is possible.

**The Problem**

I need to start monitoring my API calls. I'm interested on: 

* Amount of request segmented by time *(I wanna understand when my servers are more active)*
* Amount of request segmented by API endpoint. *(I wanna know the distribution of my requests by endpoint)*
* Average time to respond by endpoint.
* Response status code by endpoint.

**The Requirements**

* I would like to do everything without using third party services if possible.
* I would like to keep the monitoring as cheap as possible.

**Questions**

* What is the standard way to do this?
* Should I need to add something at code level or just using cloud-watch metrics (*for example*) I will be able to achieve this?

  
Thanks for your time, 

Marcos.
## [4][How can I make the best of the AWS free tier?](https://www.reddit.com/r/aws/comments/gl7ijs/how_can_i_make_the_best_of_the_aws_free_tier/)
- url: https://www.reddit.com/r/aws/comments/gl7ijs/how_can_i_make_the_best_of_the_aws_free_tier/
---
Just starting to learn AWS and was wondering how can I get the most out of it with getting that hands-on experience and common use situations you’d be doing. Should I just be doing projects? Thanks all.
## [5][Route 53 costs for studying](https://www.reddit.com/r/aws/comments/gldo5a/route_53_costs_for_studying/)
- url: https://www.reddit.com/r/aws/comments/gldo5a/route_53_costs_for_studying/
---
Hi

I'm considering purchasing a domain via Route 53 to use for AWS cert studying.

I was wondering whether anyone would have any experience on the kind of cost I could expect in these circumstance? Is it worth doing for cert studies?

I doubt I'll be using it for much other than getting hands-on with Certificate Manager (just the free public certs).

I notice that in the Route 53 pricing, there is a $0.125 per ENI / hour charge for a Resolver and that a minimum of two ENIs a required. Am I correct in thinking that I would only use the Route 53 Resolver if I wanted to offload DNS queries to Route 53 from any EC2 instances?

Sorry if this is kind of vague but I don't want to be lumped with a $182 monthly bill that the price calculator says I'd get charged for the Resolver.
## [6][Using a Load Balancer as a reverse proxy.](https://www.reddit.com/r/aws/comments/gl9uuk/using_a_load_balancer_as_a_reverse_proxy/)
- url: https://www.reddit.com/r/aws/comments/gl9uuk/using_a_load_balancer_as_a_reverse_proxy/
---
Hi folks,

I am putting a licensing server into AWS for our organisation. Our prod environment is heavily federated and we are only allowed to create or update new resources via CloudFormation. 

I have written a CloudFormation script that spins up an EC2, auto licenses the licensing server by downloading some config files, creates the security group, creates a Route 53 based off the name of the server and sets up an instance profile. It works great.

One of the goals of the project was to add HTTPS, as the server doesn't natively support it, but it can be configured [via reverse proxy](https://www.jetbrains.com/help/license_server/configuring_secure_connection.html)  


I asked around to see if I could get the certificates for our internal tools domain (it doesn't need to be accessible from the internet) but was told that these are Amazon certificates and the best way to secure this connection is by creating a Load Balancer in front of my single EC2 (I have been told by the manufacturer of the server that it won't work in any kind of multi-node configuration) and secure this connection using the built in certificate tools.   


I have never done this before and it is confusing me a little. Am I right in thinking that I would need to:  


* Create a Application Load Balancer that depends on my EC2 Instance
* Create a Load Balancer Listener that listens for https requests on my custom port with an SSL Rule and the ARN of my certificate
* Set my Route53 to create a record based on the name of my Load Balancer and to point at it's IP, when created.  


Is that about right? Any feedback appreciated.
## [7][Updating an existing EC2's AMI ID --- why?](https://www.reddit.com/r/aws/comments/glb2ei/updating_an_existing_ec2s_ami_id_why/)
- url: https://www.reddit.com/r/aws/comments/glb2ei/updating_an_existing_ec2s_ami_id_why/
---
https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html

*For example, if you have a stack with an EC2 instance, you can update the stack to change the instance's AMI ID.*

For what reason / purpose would you update an instance's AMI ID?  AMI's are only ever used at instance launch time, right?  Would this result in an implicit Terminate --&gt; Launch with the new AMI?
## [8][Are the exams being held online now?](https://www.reddit.com/r/aws/comments/gla3hm/are_the_exams_being_held_online_now/)
- url: https://www.reddit.com/r/aws/comments/gla3hm/are_the_exams_being_held_online_now/
---
I’m thinking of signing up for the AWS Solutions Architect Exam but don’t know if it’ll be in person or online. Anyone know?
## [9][Connecting to mongodb in EC2 from Lambda functions](https://www.reddit.com/r/aws/comments/gl8dgr/connecting_to_mongodb_in_ec2_from_lambda_functions/)
- url: https://www.reddit.com/r/aws/comments/gl8dgr/connecting_to_mongodb_in_ec2_from_lambda_functions/
---
Hi, I tried connecting to mongodb in Ec2 from my lambda function using private ip, that works fine but if I connect using public ip, its not. Every port is open in security group. And both lambda and ec2 are in same vpc.
What is it that I am doing wrong.
Lambda logs is giving connection timed out.
## [10][Does AWS Savings Plan apply to spot instances?](https://www.reddit.com/r/aws/comments/gl4l3t/does_aws_savings_plan_apply_to_spot_instances/)
- url: https://www.reddit.com/r/aws/comments/gl4l3t/does_aws_savings_plan_apply_to_spot_instances/
---

