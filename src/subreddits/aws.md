# aws
## [1][Python script to pull Savings Plan prices into Excel](https://www.reddit.com/r/aws/comments/i6a5om/python_script_to_pull_savings_plan_prices_into/)
- url: https://www.reddit.com/r/aws/comments/i6a5om/python_script_to_pull_savings_plan_prices_into/
---
Been developing a Python script to pull 3Y All Upfront Savings Plan prices using the Bulk API into an Excel file. It's finally in a usable alpha state.

Hopefully others and more experienced developers may find it useful and can extend it or tailor it for their use case.

[https://github.com/longhorn09/aws_prices](https://github.com/longhorn09/aws_prices)
## [2][Cloudfront OAI access denied origin S3 website](https://www.reddit.com/r/aws/comments/i6galt/cloudfront_oai_access_denied_origin_s3_website/)
- url: https://www.reddit.com/r/aws/comments/i6galt/cloudfront_oai_access_denied_origin_s3_website/
---
I bought a custom domain from AWS and have SSL a certificate using ACM. Wanted to host S3 as a static website for my custom domain but only allowing access through  cloudfront distribution. I have done the following:

1. Created a S3 bucket and named it as my custom domain.
2. Uploaded my static files to the S3 bucket.
3. Created a cloudfront distribution, selected my S3 bucket as origin 
4. Selected Restricted access and create a new OAI user and selected Yes update bucket policy
5. In the Cloudfront distribution under CNAME added my custom domain and selected my ssl cert from ACM.
6. Added an A record alias in the route53 pointing to the Cloudfront distribution.

Waited for almost 48 hours but I am still getting access denied 403 in both the Cloudfront and my custom domain URL.
## [3][Is there a more modular (OOP) methodolgy for developing AWS Step Functions?](https://www.reddit.com/r/aws/comments/i6ducb/is_there_a_more_modular_oop_methodolgy_for/)
- url: https://www.reddit.com/r/aws/comments/i6ducb/is_there_a_more_modular_oop_methodolgy_for/
---
I have to develop out a AWS Step Function that will serve as an processing orchestrator. There's five different orchestrators (step functions) each containing about 20 states. There's so much going on in these orchestrators where a select few of states are exactly the same (minus a few details).

I'm using cloudformation scripts right now and I'm quickly realizing I'm repeating alot of code (amazon states language) for these duplicate states.

Is there a more OOP way of doing this rather than using cloudformation scripts? If I use boto3 (python) how might I implement such a solution? I've been thinking about defining a class for these duplicate states and have a json object in it where I can define the state. This way if the design changes I don't have to make changes in all these different places. Open to any direction and advice.

 ~ S
## [4][learned Terraform and built an app that provisions and deploys to AWS via a single command](https://www.reddit.com/r/aws/comments/i67dpi/learned_terraform_and_built_an_app_that/)
- url: https://www.reddit.com/r/aws/comments/i67dpi/learned_terraform_and_built_an_app_that/
---
Over the past few weeks, I plunged into learning Terraform and AWS and ended up building a full-stack IaC project. It is a containerized server+client app that is one-click-deployable to an ECS Fargate cluster. The Terraform code is split by module and environment. The environments (dev, prod, and test) have their own VPC and computing resources to avoid affecting each other. The entire system is end-to-end testable with a suite that runs Terratest and Cypress. All this is tied together into a few simple commands via Makefile. Finally, I documented it in a way that would help a developer use it for the first time.

The repo: [https://github.com/brietsparks/guestbook](https://github.com/brietsparks/guestbook)

Feedback is welcome, and I hope this can be a helpful reference for someone looking to build something similar
## [5][Can I set cloudwatch to alert if today's max value is more than double yesterday's max value?](https://www.reddit.com/r/aws/comments/i66s38/can_i_set_cloudwatch_to_alert_if_todays_max_value/)
- url: https://www.reddit.com/r/aws/comments/i66s38/can_i_set_cloudwatch_to_alert_if_todays_max_value/
---
I want to create a dynamic cloudwatch alarm that will alert if today's max value is more than double yesterday's max value, but I can't find anything which allows me to set a max value based on a previous rolling period of time.

The closest thing I can find is the anomaly detection but that doesn't quite work for me as the values do not follow any sort of pattern.
## [6][CI/CD Pipeline with AWS CloudFormation](https://www.reddit.com/r/aws/comments/i6345q/cicd_pipeline_with_aws_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/i6345q/cicd_pipeline_with_aws_cloudformation/
---
I'm trying to figure out a way to come up with a CI/CD pipeline for CloudFormation. We use Spinnaker today to deploy our infrastructure and app to the cloud, and use artifactory for our artifact respository.

Does anyone have any examples of how they have created a CI/CD pipeline using Jenkins or other types of CI tools to do some type of linting, CI, version control, and artifact deployment to Artifactory (or similar toolset)? I'd like to execute a Spinnaker pipeline once a new version of the cloudformation templates are uploaded to Artifactory. 

Looking for some suggestions, examples, or guidance on how these pipelines could be built for AWS CloudFormation. Thanks!
## [7][How do you implement unique reply-to email addresses?](https://www.reddit.com/r/aws/comments/i6aou6/how_do_you_implement_unique_replyto_email/)
- url: https://www.reddit.com/r/aws/comments/i6aou6/how_do_you_implement_unique_replyto_email/
---
When a user sends an email to support, I'd like to create a unique email reply-to for the support ticket. Can anyone point me in the direction of how to do this? I'm using AWS SES right now.

I'm not sure how to implement this. I'm thinking you create new domain (@example-reply-to.com) and have every email sent to that domain be piped to an EC2 server that reads the localname of the email and connects that localname to a record in the database? Is that correct?

edit: Looks like this exists (https://aws.amazon.com/blogs/aws/new-receive-and-process-incoming-email-with-amazon-ses/) which solves my needs.
## [8][Beginner in AWS - Simple site in Lambda or EC2?](https://www.reddit.com/r/aws/comments/i61nzd/beginner_in_aws_simple_site_in_lambda_or_ec2/)
- url: https://www.reddit.com/r/aws/comments/i61nzd/beginner_in_aws_simple_site_in_lambda_or_ec2/
---
Hello, I'm currently studying cloud concepts(mostly on azure and aws) and I want to start practice, migrating a simple site that now runs in a VPS to AWS, this site have only these functionalities:

- some static text and images(I think S3 will be good to storage these files in .html and .png).

- a contact form that sends a message to an e-mail address.

- a news feed, that have static text and some images, with a simple text formatter feature(a user logins, create a new a article, puts text and images,  and then save, this will create a new .html file and a href in home page pointing to this page, automatically).

I guess lamba is suited for this, but for news functionality, this is the best approach? 
In Lamba, I think I'll do this: create a page news that updates home page, then saving .html and .png files in S3 bucket as object, as I was going to do with static contents). I don't need a users or news database right now because it's a single login and static files only, but in future if I need one, I can use RDS for my logins and blob texts, right?

I guess Lamba is good because his billing mechanism will be cheap, and EC2 will need and instance, which is expensive and more complicated to maintain. And I need to convince my employer this will be good for us to start using AWS.

I don't have the code (that's written in PHP) but I'll replicate these functionalities in python or other language if I'm going to use Lamba.
## [9][Feedback on CI CD for EMR](https://www.reddit.com/r/aws/comments/i69o0i/feedback_on_ci_cd_for_emr/)
- url: https://www.reddit.com/r/aws/comments/i69o0i/feedback_on_ci_cd_for_emr/
---
Our team is working on migrating hundreds of spark jobs to AWS. All jobs are nightly retros that create output that is then consumed by some application.

There is really is three categories we need to address

1) Full CI CD 
2) Running Code/tuning on full dataset
3) Adhoc data analysis

For number one. This is the current plan. Any check in to Master branch kicks off code pipeline. Code build will build the spark jar. Code deploy will load the jar to s3 and use cloud formation to create a step function in our dev account. The state machine will create an EMR cluster, execute spark submits and then shut down the cluster. A lambda will then be triggered that starts the state machine in the dev environment. In dev there will me a small set QA of content. When the state machine finishes another lambda will kick off. This contains our QA test suite which will validate the output from the dev state machine. If all those tests pass we will go to a manual UAT approval state. If that is approved code deploy will load jar to s3/cloud formation to our UAT environment. That code will most likely sit in UAT for a few weeks before again another manual check  to moves that code to Prod.

2) Obviously we don’t want to bottleneck our developers so we need a second code pipeline that kicks off on branch checking. This follows the same flow as above. Dev workflow deployed. QA check is ran. This will be the time where our QA team keeps their test suite up with our developers. One big issue with spark is large datasets can require a lot of tiny tweaks to get it working properly. So after QA tests pass we deploy the same state machine but pass in the production data. We just don’t load it anywhere so you can tune/optimize your code become making a PR to master.

3) adhoc data analysis. I’m thinking just an always on EMR cluster that we can shh into and run spark shell commands.

Thoughts? We just started discussing this week!
## [10][Reset AWS account after sitting dormant?](https://www.reddit.com/r/aws/comments/i61flo/reset_aws_account_after_sitting_dormant/)
- url: https://www.reddit.com/r/aws/comments/i61flo/reset_aws_account_after_sitting_dormant/
---
We have an AWS account that we played around with for some testing a couple years back.

I have no need for anything that is in it.

So two questions as I'm a total novice with AWS.

Is there a quick way to just delete EVERYTHING under the account?

If not and if I do so manually is there anything I can delete that will screw me over?

To reiterate there's nothing in the account that is needed.
