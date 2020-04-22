# aws
## [1][AWS Africa (Cape Town) region goes live](https://www.reddit.com/r/aws/comments/g5xch0/aws_africa_cape_town_region_goes_live/)
- url: https://mybroadband.co.za/news/cloud-hosting/349001-aws-cape-town-region-goes-live.html
---

## [2][AWS launches machine learning enabled search capabilities for COVID-19 dataset](https://www.reddit.com/r/aws/comments/g5jj4t/aws_launches_machine_learning_enabled_search/)
- url: https://aws.amazon.com/blogs/publicsector/aws-launches-machine-learning-enabled-search-capabilities-covid-19-dataset/
---

## [3][How to remove S3 bucket with period character in name?](https://www.reddit.com/r/aws/comments/g5yp0f/how_to_remove_s3_bucket_with_period_character_in/)
- url: https://www.reddit.com/r/aws/comments/g5yp0f/how_to_remove_s3_bucket_with_period_character_in/
---
Hello!

I am a teacher and my student somehow created a bucket with a period in name.

I have tried remove it using AWS web console:

&gt;Unknown Error. An unexpected error occurred.

I have tried AWS CLI:

&gt;aws s3 rb s3://bucket-name22.04... blah-blah-blah ...An error occurred (NoSuchBucket) when calling the ListObjectsV2 operation: The specified bucket does not exist

After that I realized how great was sin that my student just done (｡•́︿•̀｡)

Please help me to earn  forgiveness of AWS gods and remove this bucket!
## [4][Main things to consider when comparing RDS to Aurora](https://www.reddit.com/r/aws/comments/g5t8l7/main_things_to_consider_when_comparing_rds_to/)
- url: https://www.reddit.com/r/aws/comments/g5t8l7/main_things_to_consider_when_comparing_rds_to/
---
Hi all,

What are the main things to take into consideration when choosing between RDS and Aurora ? Is the price difference between having a completely managed service in Aurora versus managing your RDS yourself worth it ? It's been years since I've had to discuss relational databases so I need a bit of insight ... Thanks !
## [5][Creating your AWS resources in a declarative and programmatic way with AWS CDK](https://www.reddit.com/r/aws/comments/g5fpco/creating_your_aws_resources_in_a_declarative_and/)
- url: https://www.reddit.com/r/aws/comments/g5fpco/creating_your_aws_resources_in_a_declarative_and/
---
Hey guys, 

Here's my latest video on a topic that impacts almost anyone in this group: Infrastructure as Code. 

[https://youtu.be/zSwwwTupoAk](https://youtu.be/zSwwwTupoAk)

The idea is simple -- you can use multiple frameworks to create your AWS services in a repeatable and idempotent way, but I found CDK to be most robust and easy to learn. 

BTW, I still prefer the [Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/intro/) and [SAM](https://aws.amazon.com/serverless/sam/) for my simple, de-coupled, Lambda functions, but when it comes to more complex coupling then CDK is the go-to framework for me. As an example, check out the [Cognito + Lambda functions usage here](https://github.com/vbudilov/aws-cdk-starter/blob/master/src/main/java/com/budilov/cdk/CognitoStack.java#L37). 

Let me know if you have topic recommendations for me for my next explainer video, although I have an itch to scratch when it comes to streaming data ingestion. 

Vladimir
## [6][Question about Cost Categories (aka wtf are cost categories)](https://www.reddit.com/r/aws/comments/g616dc/question_about_cost_categories_aka_wtf_are_cost/)
- url: https://www.reddit.com/r/aws/comments/g616dc/question_about_cost_categories_aka_wtf_are_cost/
---
Every time I read the [intro page](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/manage-cost-categories.html) for Cost Categories, I feel like I understand what it's saying, but I can't figure out how it's different than what already exists.

Like the example they give about creating a **Team** category, and then defining values for `Team123` and `Team456`. I can do that already in Cost Explorer on the fly, what do I gain by doing it as a Cost Category? It actually seems _more_ cumbersome to have to pre-define it. Are they basically just a way to save Cost Explorer queries, or is there something else that I'm missing?
## [7][Security group with rules for both public and private IP adress possible?](https://www.reddit.com/r/aws/comments/g614zk/security_group_with_rules_for_both_public_and/)
- url: https://www.reddit.com/r/aws/comments/g614zk/security_group_with_rules_for_both_public_and/
---
I have an EC2 behind an NLB. I want the public IP of the EC2 only to be accessible on port 22, and on the private IP which is used by the NLB illl use port 80 and 443. Is this possible to do in a Security Group?

Eg  
SG:  
Ingress port 22 Destination: PublicIP

Ingress port 80, 443: Destination: PrivateIP

Or something like that.
## [8][Lambda - Secrets Manager - VPC Endpoint](https://www.reddit.com/r/aws/comments/g60sde/lambda_secrets_manager_vpc_endpoint/)
- url: https://www.reddit.com/r/aws/comments/g60sde/lambda_secrets_manager_vpc_endpoint/
---
I have a Lambda running in a VPC (Account A), and want to access Secrets Manager from Account B. There's an Interface VPC Endpoint for Secrets Manager in Account B (some other services in Account B VPC uses Secrets Manager as well). Currently all calls to Secrets Manager endpoint times out from Lambda.

Do I need an Interface VPC Endpoint for Secrets Manager in Account A as well for Lambda to work?
## [9][Deep-learning AMI nvidia drivers not talking to the hardware](https://www.reddit.com/r/aws/comments/g60jby/deeplearning_ami_nvidia_drivers_not_talking_to/)
- url: https://www.reddit.com/r/aws/comments/g60jby/deeplearning_ami_nvidia_drivers_not_talking_to/
---
I have been getting the following error when I look at the GPU i.e. input using

    nvidia-smi

`NVIDIA-SMI has failed because it couldn't communicate with the NVIDIA driver. Make sure that the latest NVIDIA driver is installed and running.`

Looking into this it turns out that they did an update of the AMI I amusing in [March](https://slack-redir.net/link?url=https%3A%2F%2Fforums.aws.amazon.com%2Fann.jspa%3FannID%3D7521) and I had not noticed.

I made a new instance with the updated AMI like suggested here ([https://docs.aws.amazon.com/dlami/latest/devguide/upgrading-dlami.html](https://slack-redir.net/link?url=https%3A%2F%2Fdocs.aws.amazon.com%2Fdlami%2Flatest%2Fdevguide%2Fupgrading-dlami.html)), but I am still getting the same problem.

Does anyone have any ideas?
## [10][AWS Lambda and Go](https://www.reddit.com/r/aws/comments/g5zyb8/aws_lambda_and_go/)
- url: https://www.reddit.com/r/aws/comments/g5zyb8/aws_lambda_and_go/
---
We have a Go application with Mux as the router wrapped in Apex/Gateway all zipped and deployed as a lambda and incoming requests gets routed to the appropriate function, but as I understand this is more of a hack and that each function should have its own Lambda. In refactoring this, would I need to init the DB, SQS, and other required things in each function that needs those or is there a better way of handling this? Or is our way of zipping up the whole application and using it with the Mux router fine?

I'm fairly new to working with AWS so I want to make sure I set this up correctly.
