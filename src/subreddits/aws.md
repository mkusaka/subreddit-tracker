# aws
## [1][Lambda Overview in 2020](https://www.reddit.com/r/aws/comments/f908xx/lambda_overview_in_2020/)
- url: https://www.reddit.com/r/aws/comments/f908xx/lambda_overview_in_2020/
---
Hey folks, wanted to share a youtube channel that I've been working on dedicated to providing simple and easy to digest tutorials on various AWS services.

My newest video talks about overviews Lambda functions and the improvements they've received in the past few years. 

The video is available here: https://youtu.be/iUIWG0h2D84

Thank you!
## [2][tech u intern program](https://www.reddit.com/r/aws/comments/f98edo/tech_u_intern_program/)
- url: https://www.reddit.com/r/aws/comments/f98edo/tech_u_intern_program/
---
Hi!

Has anyone gone through the process of interviewing for the tech u intern program? it seems like a fairly new program.. I cant find alot of information online. was just wondering how the interview process works and what to expect at the interview. 

Thanks! :)
## [3][Monitor AWS Managed IAM Policies Changes - Now on Twitter](https://www.reddit.com/r/aws/comments/f8ymn7/monitor_aws_managed_iam_policies_changes_now_on/)
- url: https://www.reddit.com/r/aws/comments/f8ymn7/monitor_aws_managed_iam_policies_changes_now_on/
---
Hey Folks,

Recently, I was working on a newer version of my bot who checks changes on AWS IAM Managed Policies, to push changes to a dedicated Twitter account:

\- [https://zoph.me/posts/2019-09-08-mamip/](https://zoph.me/posts/2019-09-08-mamip/)

Please tell me if you see any enhancement or issue.

A quick sneak peeks of the architecture schema below:

https://preview.redd.it/rcey42ubxxi41.png?width=571&amp;format=png&amp;auto=webp&amp;s=e6708b83a52fbda71d189931a8a12ea39d810c7c

Hope this helps :)

zoph
## [4][How to define Models for Serverless Api in CloudFormation template](https://www.reddit.com/r/aws/comments/f9a60e/how_to_define_models_for_serverless_api_in/)
- url: https://www.reddit.com/r/aws/comments/f9a60e/how_to_define_models_for_serverless_api_in/
---
See below for my template. I don't know how to set Models for Api. If I leave the Models part out, the 'sam deploy' says: "the related API does not define any Models". 

Can the models be defined in external json/yaml files? 

What is the best approach to defining more complex requests?   
How can I define model for the function response?

Can I introduce models in separate template file? 

Thanks. 

&amp;#x200B;

`Resources:`  
  `MyApi:`  
`Type: AWS::Serverless::Api`  
`Properties:`  
`StageName: test`  
`Models:`  
`???`

`PostNewItem:`  
`Type: AWS::ApiGateway::Model`  
`Properties:`  
`RestApiId: !Ref MyApi`  
`Name: PostNewItem`  
`ContentType: application/json`  
`Schema:`  
`$schema: '`[`http://json-schema.org/draft-04/schema#`](http://json-schema.org/draft-04/schema#)`'`  
`title: NewItemModel`  
`type: object`  
`properties:`  
`name:`   
`type: string`  
`description:`  
`type: string`  
`....`

  `MyFunction:`  
`Type: AWS::Serverless::Function`   
`Properties:`  
`...`  
`Events:`  
`AddItem:`  
`Type: Api`  
`Properties:`  
`Path: /item`  
`Method: post`  
`RestApiId:`   
`!Ref MyApi`  
`RequestModel:`  
`Model: !Ref PostNewItem`  
`Required: true`
## [5][Any starting tips?](https://www.reddit.com/r/aws/comments/f97ojz/any_starting_tips/)
- url: https://www.reddit.com/r/aws/comments/f97ojz/any_starting_tips/
---
I just want to host a wordpress site on amazon with some scale-able file storage.

In layman's terms - I want to create a private video hosting site with control over storage. Hosting this at somewhere like Shitdaddy would be a bad idea I assume Amazon is better value when it comes to storage?

This whole Amazon thing has been made way too confusing, why can't it be clear how to set up a website like every other hosting company has managed to do?
## [6][What's the best AWS Compute option for your project?](https://www.reddit.com/r/aws/comments/f97m7b/whats_the_best_aws_compute_option_for_your_project/)
- url: https://cloudonaut.io/whats-the-best-aws-compute-option-for-your-project/
---

## [7][Does an Application Load Balancer automatically route traffic away from a terminating Spot Instance?](https://www.reddit.com/r/aws/comments/f93ae3/does_an_application_load_balancer_automatically/)
- url: https://www.reddit.com/r/aws/comments/f93ae3/does_an_application_load_balancer_automatically/
---
Title. Wondering how you deal with interruptions; currently considering the viability of running an NGINX+Wordpress deployment on ECS with Fargate Spot. If traffic is routed away from the terminating server for 120s prior to deletion, there shouldn't be any active connections at that point, right? If so, how do you deal with that?
## [8][Recently asked this in an interview: How do you handle error if S3 event does not trigger lambda function - nodejs](https://www.reddit.com/r/aws/comments/f90lzq/recently_asked_this_in_an_interview_how_do_you/)
- url: https://www.reddit.com/r/aws/comments/f90lzq/recently_asked_this_in_an_interview_how_do_you/
---
This was a tough one and still I am struggling to find the right answer to it.

Lets assume a simple node app, which takes a file and uploads it to S3. 

a) after upload to S3, event will trigger lambda

b) lambda will parse the file and do an api call to push the contents of the file

Now, how can I handle the error if S3 event for whatever reason fails to trigger the lambda? Looking for a simple solution if possible.
## [9][Would adding CloudFront help in this case?](https://www.reddit.com/r/aws/comments/f96e5r/would_adding_cloudfront_help_in_this_case/)
- url: https://www.reddit.com/r/aws/comments/f96e5r/would_adding_cloudfront_help_in_this_case/
---
Dear r/aws,

I have an EKS pod getting an item from Dynamodb. The DB item contains a S3 URL of large string.  So the flow goes like:

1. user request -&gt; EKS get-item from DynamoDB
2. EKS get-object from S3
3. EKS response to user both info: 1) info from DDB 2) large string from S3

When I tried loadtest on this flow, I got like 4000 rps.  I am wondering if there is a way to improve capacity without adding cache servers.  Would adding CloudFront help in this case?
## [10][AWS Lambda behaviour when encountering execution limits?](https://www.reddit.com/r/aws/comments/f9629y/aws_lambda_behaviour_when_encountering_execution/)
- url: https://www.reddit.com/r/aws/comments/f9629y/aws_lambda_behaviour_when_encountering_execution/
---
So I need to create a system where 600k+ files in S3 are to be processed using AWS Lambda. I have planned to implement this by having two Lambda functions: One to split the files into batches of ~3000 and send SNS notifications with the batch information. Another function would be triggered by these notifications, fetch the objects using the key in the SNS message body and process them.

Since there is a maximum limit of 1000 instances of any given Lambda function, I wanted to know that in case that the invocations exceed 1000, will the remaining invocations get queued or would they be denied?
