# aws
## [1][New – Amazon Simple Email Service (SES) for VPC Endpoints](https://www.reddit.com/r/aws/comments/gai57v/new_amazon_simple_email_service_ses_for_vpc/)
- url: https://aws.amazon.com/blogs/aws/new-amazon-simple-email-service-ses-for-vpc-endpoints/
---

## [2][Pen-test company says our serverless API's (APIGateway+Lambda) are vulnerable to http desync / request smuggling. AWS says cloudfront protects against that. I am lost here.](https://www.reddit.com/r/aws/comments/gaoj9k/pentest_company_says_our_serverless_apis/)
- url: https://www.reddit.com/r/aws/comments/gaoj9k/pentest_company_says_our_serverless_apis/
---
I am having a very painful back and forth with AWS about this. I keep being told that Cloudfront which sits infront of API Gateway will protect against these kinds of attacks. Great.

However i have a pen-test company with results that says the opposite. Unfortunately this pen-test company isn't working on our dime and instead the pen-test is being done by a customer as part of their technical vetting. 

This leaves me in a really akward position. Can anyone give me insight?
## [3][Why is SES still not available in most of the regions?](https://www.reddit.com/r/aws/comments/gavfmm/why_is_ses_still_not_available_in_most_of_the/)
- url: https://www.reddit.com/r/aws/comments/gavfmm/why_is_ses_still_not_available_in_most_of_the/
---
 I posted this as a comment in another thread earlier. 

SES has been around since 2011, it's a very basic service,  and probably every customer can find a use case with it. But why is it only available in few regions? 

Anyone have any insights about this ? Any technical limitations that could cause this ?
## [4][Fargate ECS task takes ages to stop in step function](https://www.reddit.com/r/aws/comments/gauuqm/fargate_ecs_task_takes_ages_to_stop_in_step/)
- url: https://www.reddit.com/r/aws/comments/gauuqm/fargate_ecs_task_takes_ages_to_stop_in_step/
---
I wonder if I'm doing something wrong, since my simple docker container in ECS takes up to 30 seconds to move from stopping to stopped. It just runs a test python script which runs in 1 second, but then the task in the step function takes about a minute to complete due to the task not stopping.

I can't modify any kind of stop timeout afaik since it's a fargate task. Any ideas?
## [5][CodePipeline and Cloudformation stack limits](https://www.reddit.com/r/aws/comments/gauq1i/codepipeline_and_cloudformation_stack_limits/)
- url: https://www.reddit.com/r/aws/comments/gauq1i/codepipeline_and_cloudformation_stack_limits/
---
I'm doing some experimenting on using CodePipeline for CI/CD of our microservices (as well as moving some of them over to Lambda with SAM).    My assumption is that each codebase (i.e. microservice as we don't use monorepos) needs it's own pipeline and therefore it's own stack.   Right now we have maybe 80 microservices but that number is only going to increase as well as we are going to be using CF stacks for other parts of our infrastructure it seems we'll hit the 200 stack per account limit sooner than later.

Am I missing something?  Is the answer as simple as just contacting AWS to increase the limit and I don't need to think about this?   Nested stacks for CI/CD sounds weird on the face of it but is that an approach?   I've found no information about hitting this limit from CodePipeline stacks over the last day of searching.

Thanks for any advice.
## [6][Service and price estimate for AWS](https://www.reddit.com/r/aws/comments/gauoa4/service_and_price_estimate_for_aws/)
- url: https://www.reddit.com/r/aws/comments/gauoa4/service_and_price_estimate_for_aws/
---
I am working on a E-commerce website project and need to come up with accurate estimates for a budget. 

We are hosting 3 servers on AWS: a Web server, database server, and API. Which AWS services would I need? AWS RDS Cloud Database, Amazon EC2Cloud Server, API Gateway or just Lightsail? I am confused as to what each do

How could I get a price estimate for custom e-commerce site selling less than 5 products using AWS? 

 Thank you!
## [7][Best way to separate Dev and Prod](https://www.reddit.com/r/aws/comments/gacoju/best_way_to_separate_dev_and_prod/)
- url: https://www.reddit.com/r/aws/comments/gacoju/best_way_to_separate_dev_and_prod/
---
Took on a new client who has both Dev and Prod instances, all in the same VPC and Subnet.  I feel this needs to be split out.  Where Dev is in one place and Prod another.  Is the best approach to this, to create an entirely different AWS account, a Sub Account or maybe just on a different VPC?

One of the issues, is the people creating the Dev environment should not be able to touch Prod.
## [8][Best Serverless Solution for Handling User Notification State](https://www.reddit.com/r/aws/comments/gar4fd/best_serverless_solution_for_handling_user/)
- url: https://www.reddit.com/r/aws/comments/gar4fd/best_serverless_solution_for_handling_user/
---
I have an iOS app that sends push notifications when users post,like,comment. I have a lambda that handles the logic of that new activity, create list of users I want to send a push notification to and sends it.

What I want to do next is send the amount of pending notifications (the badge number) with each  push. As my system is now I store notifications in a dynamo table, which works okay when pulling from the client every so often to render the notification view, but I'm not sure that's going to work out too well when sending notifications. I first create the notification records in dynamo for every user that's going to get one. I then get every user record (for device token), \*I would then need do a GSI query on the notifications table and get a pending count\*. I feel like that's going to add a lot of each time inside the lambda especially since the amount of users can be variable. I can separate the sending notifications into a new lambda and process async from the activity post, which I plan on doing, but I want to make sure I'm not going down the wrong path before I make that change.

I'm wondering if an in memory data store may work best for user notifications and possible their device token, that way I don't have to make so many reads and write to the DB in a single workflow. I am also looking into DAX but I'm not sure if the eventually consistent data would be a problem or not in this case.
## [9][Lambda to download on premisses bitbucket](https://www.reddit.com/r/aws/comments/gar1fr/lambda_to_download_on_premisses_bitbucket/)
- url: https://www.reddit.com/r/aws/comments/gar1fr/lambda_to_download_on_premisses_bitbucket/
---
 Currently I have a vpc with a vpg connecting to my on-premises tolls, I need to download from the bitbucket hosted there but when i do i receive connection time out. Following some links lead me to here [https://aws.amazon.com/pt/premiumsupport/knowledge-center/internet-access-lambda-function/](https://aws.amazon.com/pt/premiumsupport/knowledge-center/internet-access-lambda-function/) , do I need a nat associated with the vpg?
## [10][Any way to remove queued AWS Lambda requests on throttled function?](https://www.reddit.com/r/aws/comments/gaqoqx/any_way_to_remove_queued_aws_lambda_requests_on/)
- url: https://www.reddit.com/r/aws/comments/gaqoqx/any_way_to_remove_queued_aws_lambda_requests_on/
---
So, I have a Lambda function that has a lot of requests queued. The requests contain bad data and I no longer want to process them. So, I have throttled the function, but when I un-throttle it, the queued requests continue processing. Is there any way to remove these queued requests completely before I un-throttle the function? There are too many requests and I feel it would take a long time to wait it out.
