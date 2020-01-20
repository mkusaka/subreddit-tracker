# aws
## [1][Osaka Local Region to be upgraded to a full region in early 2021](https://www.reddit.com/r/aws/comments/ercig3/osaka_local_region_to_be_upgraded_to_a_full/)
- url: https://aws.amazon.com/blogs/aws/in-the-works-aws-osaka-local-region-expansion-to-full-region/
---

## [2][Building a virtual cloud engineer](https://www.reddit.com/r/aws/comments/ercuhe/building_a_virtual_cloud_engineer/)
- url: https://www.totalcloud.io/create-your-own-virtual-cloud-engineer
---

## [3][Fargate price calculation check](https://www.reddit.com/r/aws/comments/er3sjw/fargate_price_calculation_check/)
- url: https://www.reddit.com/r/aws/comments/er3sjw/fargate_price_calculation_check/
---
Hi all, sorry if I'm way off. I'm new and want some validation of my calculations and maybe learn something new.

Aim is to host a container, 1 running instance, always on, using ECS Fargate.

All prices below are in AUD for Syndey.

2 vCPU and 4 GB of ram. I'm assuming 730 hours in a month. **This should cost $87/month**?

I'll be hosting the docker image in ECR to spin-up the container instance. And I'm assuming I need to leave it there for Fargate to auto-recover in case the instance dies. Also I'm going with assumption that 10GB ECR storage and 10GB ECR data transfer per month will be used. This **should cost about $2/month**

Fargate container instance itself with 1 free elastic IP, would probably use about 50GB of bandwidth, **which should cost about $6/month**?

All in all: $95/month

Have I have missed something?
## [4][Fargate for analytics service?](https://www.reddit.com/r/aws/comments/er2n05/fargate_for_analytics_service/)
- url: https://www.reddit.com/r/aws/comments/er2n05/fargate_for_analytics_service/
---
I am looking at hosting Matomo analytics on Fargate + RDS. I have never used Fargate before. This would be for very low-traffic websites (0 to 10 clicks per day). I estimate the smallest MySQL RDS will be about $10/month, and if I went with EC2 over Fargate, that would be another $10/month. I am planning to deploy the Matomo Docker image, so Fargate also looks easier to configure in that regard.

&amp;#x200B;

Is Fargate the right choice? Will it bootup fast enough? Could I end up spending more than EC2?
## [5][Hosting a web game for 1.7 Million users](https://www.reddit.com/r/aws/comments/er62pp/hosting_a_web_game_for_17_million_users/)
- url: https://www.reddit.com/r/aws/comments/er62pp/hosting_a_web_game_for_17_million_users/
---
Hi !   


I'm considering to host a game on **S3**, the static files could weight between **20Mo and 50Mo**.   
My client is excepting **14-28k connections a day** and around **1.7 Million connections over the course of 6 week.**  


I end up with this calculation for the cost of the hosting :  
0,02GB \* 1 700 000 \* 0,09$/GB = 3060$  
0,05GB \* 1 700 000 \* 0,09$/GB = 7650$  


Am i right that the cost could be between **3060$** and **7650$** ?  
Is it worth it to host this kind of project in S3 ? Or I should consider to host it directly on a dedicated server (i imagine the cost could be much lower) ?

On [https://calculator.s3.amazonaws.com/index.html](https://calculator.s3.amazonaws.com/index.html) i end up with **3823.45$** to **8799.85$.**  
Is 0,09$/GB of data out is still relevant in my case ?

Also do you have some tips for hosting this kind of project that are really demanding on data transfer ?

Thanks for the help
## [6][AWS Step Functions tasks limit](https://www.reddit.com/r/aws/comments/er0c9i/aws_step_functions_tasks_limit/)
- url: https://www.reddit.com/r/aws/comments/er0c9i/aws_step_functions_tasks_limit/
---
0

Does AWS Step Functions have a limit on the number of tasks or a limit on the number of state transitions per state machine? I have a use case where I to orchestrate workflows with hundreds maybe even a million of tasks and I can't find anywhere if Step Functions are suitable for my case.
## [7][VPC Interface Endpoints - Subnet Association Clarification](https://www.reddit.com/r/aws/comments/er57sx/vpc_interface_endpoints_subnet_association/)
- url: https://www.reddit.com/r/aws/comments/er57sx/vpc_interface_endpoints_subnet_association/
---
I need to get confirmation if I understood the limitation of the VPC Inteface endpoint according to the documentation here - [https://docs.aws.amazon.com/vpc/latest/userguide/vpce-interface.html](https://docs.aws.amazon.com/vpc/latest/userguide/vpce-interface.html)

&amp;#x200B;

It says in terms of subnet association

&gt;For each interface endpoint, you can choose only one subnet per Availability Zone.

So assume I create an SSM VPC endpoint, and I have 4 subnets

* Subnet A in Zone 1
* Subnet B in Zone 2
* Subnet X in Zone 1
* Subnet Y in Zone 2

Does this mean that 

* I can only associate ONE single subnet to the SSM VPC endpoint I created
* OR can I associate multiple subnets to ONE SSM VPC endpoint as along as they are in diff AZ. Example I can associate Subnet A and Subnet B to the SSM VPC endpoint?
## [8][RDS DMS Data migration and scrubbing.](https://www.reddit.com/r/aws/comments/er39q2/rds_dms_data_migration_and_scrubbing/)
- url: https://www.reddit.com/r/aws/comments/er39q2/rds_dms_data_migration_and_scrubbing/
---
Hey all,

Been wanting to use some data migration to scrub sensitive data off prod and push to dev.

The database in question is just a mysql one.

Does anyone know how to achieve this,
## [9][Cognito is confusing...](https://www.reddit.com/r/aws/comments/eqsi1u/cognito_is_confusing/)
- url: https://www.reddit.com/r/aws/comments/eqsi1u/cognito_is_confusing/
---
I had set-up my web-app to use cognito for user sign-in, sign-up, MFA and password management some time ago using  `amazon-cognito-identity-js` library. Import statement: 

`import * as AmazonCognitoIdentity from "amazon-cognito-identity-js";`

Now, after some time, I've come back to add few other implementations like fetching user attributes and others following resources like: 

[https://aws.amazon.com/blogs/mobile/accessing-your-user-pools-using-the-amazon-cognito-identity-sdk-for-javascript/](https://aws.amazon.com/blogs/mobile/accessing-your-user-pools-using-the-amazon-cognito-identity-sdk-for-javascript/)

As I was writing code, most of the things are not recognized by above imported library. Example: 

    AWS.config.credentials = new AWS.CognitoIdentityCredentials({})  &lt;-- seems like these are not part of amazon-cognito-identity-js library.

Another resource, from aws js docs: [https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#getUser-property](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#getUser-property)

    var cognitoidentityserviceprovider = new AWS.CognitoIdentityServiceProvider(); &lt;-- same error.

What am I missing here? Note that I've tried importing library as AWS.

Edit: I don't want to use amplify sdk since most of my code is already written using js sdk.
## [10][What's the correct parameters to pass to AWS to set Glue Crawler to ignore changes?](https://www.reddit.com/r/aws/comments/eqzvw4/whats_the_correct_parameters_to_pass_to_aws_to/)
- url: https://www.reddit.com/r/aws/comments/eqzvw4/whats_the_correct_parameters_to_pass_to_aws_to/
---
I am passing the following parameters to AWS using Terraform:

    resource "aws_glue_crawler" "my_crawler" {
      database_name = var.catalogDb
      name          = "my_crawler"
      role          = var.crawlerRole
    
      schedule      = "cron(0 */6 * * ? *)"
      table_prefix = "md_"
    
      s3_target {
        path = "s3://my_crawler_data_test/"
      }
    
      schema_change_policy {
        delete_behavior = "LOG"
      }
    
      configuration = &lt;&lt;EOF
    {
      "Version":1.0,
      "CrawlerOutput": {
         "Partitions": { "AddOrUpdateBehavior": "InheritFromTable" },
         "Tables": { "AddOrUpdateBehavior": "MergeNewColumns" }
      },
      "Grouping": {
         "TableGroupingPolicy": "CombineCompatibleSchemas"
      }
    }
    EOF

to generate a Glue/Crawler. On the UI I can see that my configuration generates:

[What I am currently getting with above Json configuration](https://preview.redd.it/bw66e3py0sb41.png?width=1328&amp;format=png&amp;auto=webp&amp;s=19522c2da40ba1b8882008b0345570fc5eb63c1d)

&amp;#x200B;

but I would like to get this:

[What I would like to get](https://preview.redd.it/njezv2511sb41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=2a83a9e60ac115a4cf2b4743ee4053d781916f1f)

The only change is the option "Add new columns" to "Ignore...". I tried different options in my Json conf but nothing worked, and from the documentation it is not clear how to get that. Any idea?
