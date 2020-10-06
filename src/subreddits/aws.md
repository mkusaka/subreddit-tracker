# aws
## [1][Week of Oct 5th - What AWS questions do you have? (but were always afraid to ask :-) )](https://www.reddit.com/r/aws/comments/j5jsap/week_of_oct_5th_what_aws_questions_do_you_have/)
- url: https://www.reddit.com/r/aws/comments/j5jsap/week_of_oct_5th_what_aws_questions_do_you_have/
---
Ask your AWS questions here and help others - Also don't forget that there are almost 2500 members in the /r/aws chatroom!
## [2][Welcome to new mod u/_abhayshah!](https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/)
- url: https://www.reddit.com/r/aws/comments/j5trs4/welcome_to_new_mod_u_abhayshah/
---
Thrilled to expand the mod team in order better serve the community. As /r/aws grows, u/_abhayshah will be able to assist with AMAs, post flair, mod queue/mail, building out the Wiki, and more!   


Please give him a nice /r/aws welcome and let us know how we can improve things together going forward.
## [3][IaC comparison for event-driven architecture](https://www.reddit.com/r/aws/comments/j5uch5/iac_comparison_for_eventdriven_architecture/)
- url: https://www.reddit.com/r/aws/comments/j5uch5/iac_comparison_for_eventdriven_architecture/
---
Hey everyone!  I've been running production workloads on AWS for a long time (started when S3 was announced in 2006) and have been silently lurking on r/aws for over a year.  After many discussions with colleagues debating the future of cloud architectures, I decided to start [my own blog](https://blog.outwiththeold.info/) where I will explore new ways to solve old tasks.

[In my first post](https://blog.outwiththeold.info/posts/event-driven-iac/) I compare two architectures, both built with the AWS CDK, that process a file dropped into an S3 bucket. One architecture uses ECS/Fargate and the other Lambda.

I'd really like to hear what you think about this post and get suggestions for additional topics you would like to see covered.

You can read my first post here:

* [https://blog.outwiththeold.info/posts/event-driven-iac/](https://blog.outwiththeold.info/posts/event-driven-iac/)
## [4][How to change a variable in a JavaScript file, depending on the build project used in CodeBuild?](https://www.reddit.com/r/aws/comments/j63sbr/how_to_change_a_variable_in_a_javascript_file/)
- url: https://www.reddit.com/r/aws/comments/j63sbr/how_to_change_a_variable_in_a_javascript_file/
---
I am building a serverless website using HTML, CSS and JavaScirpt that is hosted on an S3 bucket.

I have a CodeCommit repository, with two branches: develop and master. I do all my development work in develop, and when I'm happy with the changes I merge into master, which is my production environment.

I have two CodeBuild projects, one for the development environment and one for the production environment. I have a build spec file, which tells CodeBuild to sync the CodeCommit files with the S3 bucket that is hosting the website.

There is a variable inside the JavaScript file, that contains the URL for a HTTP endpoint. I need to dynamically change this variable, depending on which environment I am deploying to - development or production. For example, if I deploy to develop, I want this variable to automatically be "[http://develop](http://develop)". Alternatively, if I deploy to master, I want it to be "[http://production](http://production)".

How can I do this? I know there is something about environment variables, but I am not sure on how to tie this in with everything else.
## [5][I've deleted my account but Amazon keeps billing me for Route 53 charges.](https://www.reddit.com/r/aws/comments/j5nh4w/ive_deleted_my_account_but_amazon_keeps_billing/)
- url: https://www.reddit.com/r/aws/comments/j5nh4w/ive_deleted_my_account_but_amazon_keeps_billing/
---
I deleted my AWS account because I wasn't using it and apparently I forgot to terminate the Route 53 service.

Now Amazon won't let me access the account proper to finish it and their "customer support" form is utter garbage so, what should I do? Is there a direct e-mail address? Am I just screwed?

Update that should have been qrittrn yrsterday but forgot to: I put a case in for the recovery of the account and I'm now waiting a response/update on it

Update 2: They've gotten back to me, they're reopenning the account to let me fix it. Thanks for the help, everyone! Once it's done I'll officially mark this as closed.
## [6][AWS development environment/workflows for large teams](https://www.reddit.com/r/aws/comments/j5xkqv/aws_development_environmentworkflows_for_large/)
- url: https://www.reddit.com/r/aws/comments/j5xkqv/aws_development_environmentworkflows_for_large/
---
Just curious how different development teams working with aws design their development environment when needing to test changes against real aws resources.

For example, does each developer on your team create their own version of resources to test their changes against (ex: using cloudformation)? Does software like localstack usually work as an alternative?
## [7][Most efficient way to connect to RDS without an own instance](https://www.reddit.com/r/aws/comments/j64g2i/most_efficient_way_to_connect_to_rds_without_an/)
- url: https://www.reddit.com/r/aws/comments/j64g2i/most_efficient_way_to_connect_to_rds_without_an/
---
Hi everybody!

We are starting to use AWS in our company, but we are still newbies, so I hope I don't bother if any question is too obvious, I have tried to document as much as possible as with the rest of the things we have already done.

We have a Datalake in S3 made up of 3 buckets (in a dev account):

* In the first bucket, we have the raw data, which we filter and normalize through Glue, to save it in the second bucket.
* In the second bucket, we have the filtered data that is sent to EMR to perform several actions (data merge, variable creation and transformation)... and later saved in the third bucket.
* In the third bucket, we have the final data that is consulted through Athena and sent to Quicksight.

We need to incorporate data that is stored in an Oracle RDS instance into another AWS account. To incorporate it into our datalake, the data needs to be stored in S3 in the first cube. 

I have seen that is possible to export RDS snapshots to S3 from another account, but this RDS database will be updated frequently. You can automate the daily creation of snapshots in S3, but only in the same account in which the RDS instance is placed. To send them to another account, I think it must be done manually.

Therefore, I thought of creating an RDS instance in our account and connect that instance to the instance in the other account, and automate the generation of snapshots from our account. However, the creation of the RDS instance in our account has a very high cost in relation to the rest of services.

Do you know if this task could be done in an efficient way without having to export the snapshots manually or without having to create an RDS instance in our account?

Thanks a lot for your time!
## [8][AWS S3 Browser-Based Uploads Question](https://www.reddit.com/r/aws/comments/j63ior/aws_s3_browserbased_uploads_question/)
- url: https://www.reddit.com/r/aws/comments/j63ior/aws_s3_browserbased_uploads_question/
---
Hi Everyone,

I'm relatively new with very little experience in using AWS. I was wondering if it was possible to allow public users to upload files to a bucket I've setup using just a public URL? 

What I am currently doing is extracting the files people upload to a dropbox then manually uploading the file to the s3 and setting up folders within the bucket.

Thanks in advance!
## [9][How should I deploy my AWS for a simple app?](https://www.reddit.com/r/aws/comments/j62rue/how_should_i_deploy_my_aws_for_a_simple_app/)
- url: https://www.reddit.com/r/aws/comments/j62rue/how_should_i_deploy_my_aws_for_a_simple_app/
---
Hello,

I am going to make a simple web app + Android app that contains a MySQL (Or MariaDB) database with not too many records. The app does not have registration, but it stores information provided by users. Both the web and the Android interface should send the data from the app to the database.

I'm not expecting too much traffic, so I want to know how and what should I deploy, and make it scalable, if some miracle happens and the traffic grows. And also how can I secure my deployment?

I was thinking about a single EC2 instance and a single RDS. But there are many instances, and my budget is not too big right now, so I was thinking which is the best machine to deploy?

Also, in order to save costs should I install the database on the EC2 itself and get a larger SSD space, or it's will be more expensive and less secured?

Thanks
## [10][Hosting thousands of static websites?](https://www.reddit.com/r/aws/comments/j60ze6/hosting_thousands_of_static_websites/)
- url: https://www.reddit.com/r/aws/comments/j60ze6/hosting_thousands_of_static_websites/
---
What would be the best way to host a lot (thousands) static websites with different domains using AWS? Ideally with instant updates when content changes and without having to handle scaling ourselves. 

I’ve looked into just using S3 and CloudFront via API, but that would mean thousands of CloudFront distributions, as I understand it, since one distribution only can be associated with one certificate. This approach feels a bit hard to manage in the long run, but maybe it’s the way to go?
## [11][Limiting the S3 PUT file size using pre-signed URLs](https://www.reddit.com/r/aws/comments/j5lhhn/limiting_the_s3_put_file_size_using_presigned_urls/)
- url: https://www.reddit.com/r/aws/comments/j5lhhn/limiting_the_s3_put_file_size_using_presigned_urls/
---
I am generating S3 pre signed URLs so that the client(mobile app) can PUT an image directly to S3 instead of going through a service. For my use case the expiry time of the pre signed URL needs to be configured for a longer window (10-20 mins). I want to limit the size of file upload to S3 so that any malicious attacker can not upload large files to the S3 bucket. The client will get the URL from a service which has access to the S3 bucket. I am using AWS Java SDK.

I found that this can be done using POST forms for browser uploads but how can I do this using just signed S3 URL PUTs? I have found conflicting answers for this on different platforms.
## [12][Lightsail data transfer (egress) - what's the catch?](https://www.reddit.com/r/aws/comments/j5wgfi/lightsail_data_transfer_egress_whats_the_catch/)
- url: https://www.reddit.com/r/aws/comments/j5wgfi/lightsail_data_transfer_egress_whats_the_catch/
---
If I wanted to download 2TB from S3, I could proxy it through a lightsail instance for $5.

Downloading it directly would cost $180. That's 36x more expensive.

Is this true? What's the catch? Why does this exist?

I was reading https://forums.aws.amazon.com/thread.jspa?threadID=253162 to find out, and it appears that the data transfer is prorated to the month - so if you have a lightsail instance for 15 days (half a month), you only get half of the month's data transfer limit (quite reasonable!). So it isn't even more ridiculous than that, you need the full month's $5.

Even if this is technically possible, is this the sort of thing that will cause AWS to terminate me for abusing their pricing model / services? I'm considering spinning up about two dozen of these $5 instances to download about 25 terabytes from S3.
