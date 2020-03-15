# aws
## [1][Be careful if you redeploy any HTTP API Gateways via SAM](https://www.reddit.com/r/aws/comments/fireba/be_careful_if_you_redeploy_any_http_api_gateways/)
- url: https://www.reddit.com/r/aws/comments/fireba/be_careful_if_you_redeploy_any_http_api_gateways/
---
It seems like since HTTP APIs went GA a couple days ago, the default Lambda event format has changed, and it **will change for your existing APIs** after you redeploy them (I can confirm this with SAM deploys, not sure about all other methods).

This will make any code you have accessing the headers, HTTP method, etc not work without changing to the new format, or opting back into the old format (which I don't think is supported via SAM at this point).

Also don't believe the [docs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html), the event will include `pathParameters` in v2.
## [2][Lambda Destinations - Not Sending MSG to SQS](https://www.reddit.com/r/aws/comments/fj0wzp/lambda_destinations_not_sending_msg_to_sqs/)
- url: https://www.reddit.com/r/aws/comments/fj0wzp/lambda_destinations_not_sending_msg_to_sqs/
---
For the life of me, I can’t get my lambda function to deliver a completion message as a destination set in the Lambda console to another SQS queue on success or failure.

I have SQS queue A which takes a message of an S3 upload, Lambda which does some process and movement, and then a SQS queue B which should receive success messages and an SQS queue C which should take failure messages.

Any ideas?
## [3][[BEGINNER HELP] Connecting RESTful API to RDS Help!](https://www.reddit.com/r/aws/comments/fiw9gl/beginner_help_connecting_restful_api_to_rds_help/)
- url: https://www.reddit.com/r/aws/comments/fiw9gl/beginner_help_connecting_restful_api_to_rds_help/
---
I'm looking for some help with AWS and connecting to it. Currently I have a Spring Backend, MySQL Database (on aws rds) and a React front-end. 

I can build it all together, and connect to AWS from my computer when I have it built locally, but as soon as I deploy it on Heroku, I get hit with 403 errors and I'm assuming that AWS is blocking Heroku from being able to connect to it. I do have public accessibility on it set to everyone.

I've been lost for the past couple hours trying to sort this out.
## [4][EKS - Why nginx ingress controller is using classic load balancer as default LB? Will it be deprecated soon ?](https://www.reddit.com/r/aws/comments/fiy004/eks_why_nginx_ingress_controller_is_using_classic/)
- url: https://www.reddit.com/r/aws/comments/fiy004/eks_why_nginx_ingress_controller_is_using_classic/
---
I am bit worried about the "Migrate Now" message in classic load balancer setting page , Is AWS planning to discontinue it any time soon ?   

Since 1.15 EKS supports tls termination at NLB , So is it better to switch the NLB since we are  planning to host a web socket based application in future ?
## [5][Does EC2 + ElasticBeanstalk essentially make docker containers obsolete for web apps?](https://www.reddit.com/r/aws/comments/fiud7y/does_ec2_elasticbeanstalk_essentially_make_docker/)
- url: https://www.reddit.com/r/aws/comments/fiud7y/does_ec2_elasticbeanstalk_essentially_make_docker/
---
Still new to aws so could be a dumb question. It seems like Docker containers and EC2 instances basically achieve the same end result though.

Could someone clarify for me the difference between the two?
## [6][Logging into an app with Device Farm?](https://www.reddit.com/r/aws/comments/fj1b7i/logging_into_an_app_with_device_farm/)
- url: https://www.reddit.com/r/aws/comments/fj1b7i/logging_into_an_app_with_device_farm/
---
Been working on a project where users can authentic themselves by either creating an account or logging with an existing Google account (Angular/Capacitor using Firebase authentication). I would like to test the app using AWS Device Farm and have a couple of questions:

1. Is it possible to test signup flow with device farm?
2. Will Device Farm be able to login to the app and continue testing?
3. If so, how do I set it up?
4. Do you have any other suggestions in general?

I have limited access to devices to test on, both Android and iOS. Device Farm seems like a perfect solution for me if it is able to login. Thanks in advance!
## [7][Limit users by account? Is that a strategy?](https://www.reddit.com/r/aws/comments/fit2fp/limit_users_by_account_is_that_a_strategy/)
- url: https://www.reddit.com/r/aws/comments/fit2fp/limit_users_by_account_is_that_a_strategy/
---
I think you're supposed to create users in the root (master?) account and setup different accounts for say development and *security*. 
But how do you say just want user123 to be able *only access the security log account* and not the others like dev? I know you could effectively limit user123 to the *security account* with IAM policy, though you don't want user123 logging into accounts they don't have any business in.
## [8][Transfert files directly from server to S3 Bucket](https://www.reddit.com/r/aws/comments/fih9a2/transfert_files_directly_from_server_to_s3_bucket/)
- url: https://www.reddit.com/r/aws/comments/fih9a2/transfert_files_directly_from_server_to_s3_bucket/
---
Hello guys,

I’m using a service who provide many .gz files and I want to transfer them to my S3 bucket. The client update these files every day. I have to stay in sync with his server and my S3 Bucket.

The idea is to create an Lambda  (called every day) who would be in charge to transfer these .gz files directly from his server to my Bucket.

I’ve no idea how to do that.

They provide an API with all endpoints where these files are stored.

`{`  
`fileName: "xxxxx",`  
`url: "https://domaine.com/file.gz"`  
`}`

**Do you have articles or documentation about this use case?**
## [9][Spatial queries with AWS Lambda](https://www.reddit.com/r/aws/comments/fivfae/spatial_queries_with_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/fivfae/spatial_queries_with_aws_lambda/
---
Hi all

I need to make a small web API that performs a few spatial queries (e.g. in which polygons is the submitted point).

What databases have people used for this in the past? It's not a high traffic system so would love to avoid having to spin up a full Postgres DB for the task.
## [10][What access do you get with a developer plan?](https://www.reddit.com/r/aws/comments/fir7hl/what_access_do_you_get_with_a_developer_plan/)
- url: https://www.reddit.com/r/aws/comments/fir7hl/what_access_do_you_get_with_a_developer_plan/
---
I just signed up to free tier of AWS to experiment with EC2s.  But if I pay for a Dev plan will I get access to the G&amp;P type cpus? 

Sorry if this has already been asked.
