# aws
## [1][Week of Sept 14th - What are your favorite container/serverless tips in AWS?](https://www.reddit.com/r/aws/comments/isls8o/week_of_sept_14th_what_are_your_favorite/)
- url: https://www.reddit.com/r/aws/comments/isls8o/week_of_sept_14th_what_are_your_favorite/
---
Share your container/serverless tips
## [2][New EC2 T4g Instances – Burstable Performance Powered by AWS Graviton2 – Try Them for Free | Amazon Web Services](https://www.reddit.com/r/aws/comments/isvuym/new_ec2_t4g_instances_burstable_performance/)
- url: https://aws.amazon.com/blogs/aws/new-t4g-instances-burstable-performance-powered-by-aws-graviton2/
---

## [3][Not sure where to start on AWS? Check out these 25 services. (Getting started and video tutorial links included)](https://www.reddit.com/r/aws/comments/isrbo2/not_sure_where_to_start_on_aws_check_out_these_25/)
- url: https://www.reddit.com/r/aws/comments/isrbo2/not_sure_where_to_start_on_aws_check_out_these_25/
---
See a service you’re interested in learning about? Click the link below for the links to getting started: [Getting started with AWS services](https://www.allcode.com/top-aws-services/)  


Did I miss any? Let me know and I will add them to the list with links to getting started. Thanks for your help!

&amp;#x200B;

1. Amazon EC2 (Elastic Compute Cloud)
2. Amazon RDS (Relational Database Services)
3. Amazon S3 (Simple Storage Service)
4. Amazon Lambda
5. Amazon CloudFront
6. Amazon Glacier
7. Amazon SNS (Simple Notification Service)
8. Amazon EBS (Elastic Block Store)
9. Amazon VPC (Virtual Private Cloud)
10. Amazon Kinesis
11. Amazon Auto-scaling
12. Amazon IAM (Identity and Access Management)
13. Amazon SQS (Simple Queue Service)
14. Amazon Elastic Beanstalk
15. Dynamo DB
16. Amazon ElastiCache
17. Amazon Redshift
18. Amazon Sagemaker
19. Amazon Lightsail
20. Amazon EFS (Elastic File System)
21. Amazon Cloudwatch
22. Amazon Chime
23. Amazon Cloud Directory
24. Amazon Cognito
25. Amazon Inspector
## [4][[Glacier] Restoring a 130GB file from Glacier, using MultipartCopy to transfer back to standard returns an "internal error" after a while?](https://www.reddit.com/r/aws/comments/it8ete/glacier_restoring_a_130gb_file_from_glacier_using/)
- url: https://www.reddit.com/r/aws/comments/it8ete/glacier_restoring_a_130gb_file_from_glacier_using/
---
There is a file sitting in S3 under Glacier. We trigger a RestoreObject on this key. Once the Restore has completed, we trigger a MultipartCopy back to the Standard class - as the Restore will only remain for 2 days.

However - it seems like many large files are hitting errors during this MultipartCopy:

&gt; An exception occurred while uploading parts to a multipart upload. The following parts had errors:
- Part 6289: Error executing "UploadPartCopy" on "https://mybucket.s3.amazonaws.com/projects/ID/source-files/ID/master/ID.mov?partNumber=6289&amp;uploadId=uploadIdToken";

&gt; AWS HTTP error: Server error: `PUT https://mybucket.s3.amazonaws.com/projects/ID/source-files/ID/master/ID.mov?partNumber=6289&amp;uploadId=uploadIdToken` resulted in a `500 Internal Server Error` response:

&gt; &lt;Error&gt;&lt;Code&gt;InternalError&lt;/Code&gt;&lt;Message&gt;**We encountered an internal error. Please try again.**&lt;/Message&gt;

Is there anything obvious here?
## [5][Wondering about S3 supplements/alternatives for the high bandwidth website](https://www.reddit.com/r/aws/comments/it7lmm/wondering_about_s3_supplementsalternatives_for/)
- url: https://www.reddit.com/r/aws/comments/it7lmm/wondering_about_s3_supplementsalternatives_for/
---
Hi guys! I've recently watched a lecture on the differences between geo.distributed and centralized storage models ([https://youtu.be/TOnZ78ay3rs](https://youtu.be/TOnZ78ay3rs)). My personal project to create and deliver a 4k resolution stock video through the private subscription.  I am with S3 for 4+years now, but the transfer costs are killing me.

1. I was wondering if anyone used geographically distributed storage for the heavy files storage and delivery as a replacement or alongside AWS S3?
2. Talking about geo-distributed storage services, any recommendations? I've spent hours searching and found only two so far: JUCE ([https://juce.cloud/](https://juce.cloud/)) and Eco4cloud ([http://www.eco4cloud.com/](http://www.eco4cloud.com/)).
3. The first one offers 20tb storage and unlimited transfer for €1800 per month. It would cut my expenses more than 5 times, but I couldn't find any reviews on them online, and this kind of puts me off.  Do you think it worth trying out?

If you have any other suggestions/solutions not related to geo-distributed storage - I'd be more than happy to hear them. The main idea is to cut the transfer-related costs while delivering 4k videos to the USA and Asia (this is where most of my clients are).
## [6][To "lift-and-shift" or not to "lift-and-shift"?](https://www.reddit.com/r/aws/comments/it7fdu/to_liftandshift_or_not_to_liftandshift/)
- url: https://www.reddit.com/r/aws/comments/it7fdu/to_liftandshift_or_not_to_liftandshift/
---
Under the very high pressure of time constraints (migration out of data centers by mid-2021) a lot of application owners are engaging consultants specialized in their specific areas, who - of course - offer almost exclusively "lift-and-shift" solutions, i.e. acquire "x" EC2/AMI machines of sizes "a", "b", ..., databases (PostegreSQL, Oracle, SQL) of other machine sizes, and "connect them together" identically to how they look on-prem. The only area where I was able to "insert" something more cloud-native, was in the replacement of very crude LB functions, presently conducted on F5s, which could be easily addressed via simple ELB configurations.

Question: while an approach of "like-to-like" seems logical, is there anything in the world of tooling, for AWS, which would allow an analysis of an app architecture and/or data flow, with ability to recommend maybe some AWS native solutions, where the traditional server/DB/storage/connectivity package is suboptimal for that specific usage?
## [7][SES Not Delivering to outlook.com email](https://www.reddit.com/r/aws/comments/it80r5/ses_not_delivering_to_outlookcom_email/)
- url: https://www.reddit.com/r/aws/comments/it80r5/ses_not_delivering_to_outlookcom_email/
---
Anyone else having issues with Amazon SES not delivering emails to [outlook.com](https://outlook.com) mailboxes?  I just set the service up and cannot get emails to deliver to my two [outlook.com](https://outlook.com) email accounts.  Other accounts seems to be working.  SPF is passing.
## [8][Lambda@edge returns JS as content-type: text/html](https://www.reddit.com/r/aws/comments/it7u86/lambdaedge_returns_js_as_contenttype_texthtml/)
- url: https://www.reddit.com/r/aws/comments/it7u86/lambdaedge_returns_js_as_contenttype_texthtml/
---
I have a website (Django site) deployed with Lambda and CloudFront and it works fine with no problems. I decided to try Lambda@edge with origin request set and it deployed without errors but when I visit it my site it will load all my JS except one script.

[https://xxxx.cloudfront.net/js/main.js](https://xxxx.cloudfront.net/js/main.js)

**content-type: text/html**

# 502 ERROR

## The request could not be satisfied.

The Lambda function returned invalid json: The json output must be an object type.
## [9][Seeing t4g instances going live](https://www.reddit.com/r/aws/comments/isulkr/seeing_t4g_instances_going_live/)
- url: https://www.reddit.com/r/aws/comments/isulkr/seeing_t4g_instances_going_live/
---
Seeing them available for spin up in my accounts. The pricing information is also live in the usual places, but no formal announcement as at the time of this post.
## [10][Starting stopped Gs.xlarge instance gives limits error, but allows me to start it.](https://www.reddit.com/r/aws/comments/it6rmw/starting_stopped_gsxlarge_instance_gives_limits/)
- url: https://www.reddit.com/r/aws/comments/it6rmw/starting_stopped_gsxlarge_instance_gives_limits/
---
We have one instance that we're working on and today I've got this error message.

But I was allowed to start it.

What does this mean actually?

&amp;#x200B;

https://preview.redd.it/10i7hd9span51.png?width=852&amp;format=png&amp;auto=webp&amp;s=a8d405688d7a5b8b07802e02f7c7be8da909b63b
## [11][Centrally enabling AWS Config across all accounts/regions with Organizations?](https://www.reddit.com/r/aws/comments/it6fe9/centrally_enabling_aws_config_across_all/)
- url: https://www.reddit.com/r/aws/comments/it6fe9/centrally_enabling_aws_config_across_all/
---
We have a situation where we can't use the Organizations integration for AWS config, but want to easily enable it across a number of accounts and all of the regions in those accounts. Is there any easy/preferred way to do this with Terraform? The thought of having X number of providers for each region multiplied accounts seems completely out of hand.

We don't use much Cloudformation, but I'd consider it for the better stack support across regions/accounts.

Is there a good CLI example? Also, is it simply just creating an S3 bucket in a central account, granting the permissions to to the other accounts and then enabling the recorders from there?
