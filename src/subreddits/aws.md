# aws
## [1][Calling AWS Architects: How common is the request to implement ipv6?](https://www.reddit.com/r/aws/comments/f2ffha/calling_aws_architects_how_common_is_the_request/)
- url: https://www.reddit.com/r/aws/comments/f2ffha/calling_aws_architects_how_common_is_the_request/
---
Any Architects here? 

How common is the client request to implement IpV6? When the question is raised, is the client future-proofing, migrating a existing ipv6 service, or just wants to have it because its fancy? Natting pretty much solves ipv4 problems, so whats the big deal?

I did some lite research and found this thread: 

[https://www.reddit.com/r/aws/comments/af4sj3/when\_do\_i\_need\_to\_care\_about\_ipv6/](https://www.reddit.com/r/aws/comments/af4sj3/when_do_i_need_to_care_about_ipv6/) 

That basically said - Hey cloudfront is neat when it's ipv6. Linkedin uses ipv6 for mobile traffic. It's simple enough to implement and subnet but my question is:

What are the common scenarios in which a client wants ipv6 implemented? Does supporting both add any significant management overhead or make anything easier to deal with?
## [2][Billing through AWS consulting partner?](https://www.reddit.com/r/aws/comments/f2oq5r/billing_through_aws_consulting_partner/)
- url: https://www.reddit.com/r/aws/comments/f2oq5r/billing_through_aws_consulting_partner/
---
Hi, we were approached by AWS consulting Partner. They offered us to transfer billing to them mentioning that the price will be the same but we will have help from their architects. That's sounds quite good but I'm not sure is there any drawbacks/ security risks about that?

Thanks
## [3][Kinesis with enabled encryption - no data coming in Firehose and S3](https://www.reddit.com/r/aws/comments/f2r1of/kinesis_with_enabled_encryption_no_data_coming_in/)
- url: https://www.reddit.com/r/aws/comments/f2r1of/kinesis_with_enabled_encryption_no_data_coming_in/
---
Hi /r/aws,

I have CloudWatch log groups with subscription filters to Kinesis data streams and then Firehose with that Datastream as a source configured to dump data in S3.
Everything is working just fine until I turn on encryption on Kinesis Data Streams.

I can't find much about it and only link out there is this:
https://docs.aws.amazon.com/firehose/latest/dev/encryption.html

When you send data from your data producers to your data stream, Kinesis Data Streams encrypts your data using an AWS Key Management Service (AWS KMS) key before storing the data at rest. When your Kinesis Data Firehose delivery stream reads the data from your data stream, Kinesis Data Streams first decrypts the data and then sends it to Kinesis Data Firehose. Kinesis Data Firehose buffers the data in memory based on the buffering hints that you specify. It then delivers it to your destinations without storing the unencrypted data at rest.


If I'm reading this well, seems like Firehose doesn't need access to KMS key? Looks like Data Streams would take care of decryption before sending data to Firehose?

I have tried multiple things, different policies and permissions and nothing, just can't get anything to Firehose as long as I have Encryption enabled.

Any hints or ideas? Thanks!
## [4][HOW to REMOVE NON_EXISTING TARGET_GROUP from AUTO-SCALING](https://www.reddit.com/r/aws/comments/f2o7jr/how_to_remove_non_existing_target_group_from/)
- url: https://www.reddit.com/r/aws/comments/f2o7jr/how_to_remove_non_existing_target_group_from/
---
I am facing a problem explained below:

I deleted a target group after removing the respective rules from the Application Load Balancer, I forgot to detach the target group from auto-scaling due to which auto-scaling giving error while launching new instances.

*Error: Status Reason: One or more target groups not found. Validating load balancer configuration failed.*

I don't have CLI access for my AWS IAM (account).

What should I do to resolve the issue ... Reporting it to the higher authority may bring problem to me.....

&amp;#x200B;

Thanks
## [5][My Amazon Elastic Compute Cloud usage is growing with no EC2 instances running](https://www.reddit.com/r/aws/comments/f2q8qi/my_amazon_elastic_compute_cloud_usage_is_growing/)
- url: https://www.reddit.com/r/aws/comments/f2q8qi/my_amazon_elastic_compute_cloud_usage_is_growing/
---
So I am using my 750 hours of free EC2 instance, and I recently shut down all my databases and all my EC2 instances, but my "Current Usage" hours keep growing with no EC2 instances or anything.. What else causes usage of the elastic compute cloud?
## [6][Powershell and Private ECS/S3 Object Store](https://www.reddit.com/r/aws/comments/f2px5y/powershell_and_private_ecss3_object_store/)
- url: https://www.reddit.com/r/aws/comments/f2px5y/powershell_and_private_ecss3_object_store/
---
My organisation has deployed a DellEMC ECS Object store with compatibility for the S3 API and I am trying to use the AWSPowershell.Netcore module to connect to it.

I have googled for days and cannot see any method for telling the session/cmdlets that I want to talk to my own environment and not AWS in the cloud.

Is this even possible?
## [7][[ECS] CPU utilization in Cloudwatch is different from docker stats](https://www.reddit.com/r/aws/comments/f2phop/ecs_cpu_utilization_in_cloudwatch_is_different/)
- url: https://www.reddit.com/r/aws/comments/f2phop/ecs_cpu_utilization_in_cloudwatch_is_different/
---
Hi,

I noticed today that the cloudwatch metric  **AWS/ECS "CPUUtilization"** is reporting values different from what i have when i do **docker stat** command in the container instance.

For example today in Cloudwatch, the CPU utilization is between 0 and 10% while when i executed docker stat on the container i saw multiple time CPU 100% for the same period.

Why this difference ? 

BTW, i'm using a t3.2xlarge instance type.

Thanks.
## [8][Wafv2 &amp; CloudFormation](https://www.reddit.com/r/aws/comments/f2p5us/wafv2_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/f2p5us/wafv2_cloudformation/
---
Has anyone had any experience with creating WAFv2 ACLs/Rules with CloudFormation?

I am currently struggling to get WAFv2 resources (specifically Rate Based rules for Cloudfront resources) to create successfully in CloudFormation.

For example, when creating an ACL and rules and setting the scope as CLOUDFRONT I get the following error returned:

    The scope is not valid., field: SCOPE_VALUE, parameter: CLOUDFRONT (Service: Wafv2, Status Code: 400...

I'm hoping I'm doing something wrong but having exhausted the CF documentation for WAFv2 I'm starting to think that there is something up with WAFv2 and CloudFormed resources.
## [9][AWS console logout timeout](https://www.reddit.com/r/aws/comments/f2oxah/aws_console_logout_timeout/)
- url: https://www.reddit.com/r/aws/comments/f2oxah/aws_console_logout_timeout/
---
Hi,

I don't know if this is the right place to ask this but when I log in to my organization's aws console, it logs me out after 10 minutes. Is there any way to increase this time?
## [10][[CodeBuild] Deploy artifacts to different S3 directories based on Git branches](https://www.reddit.com/r/aws/comments/f2lruy/codebuild_deploy_artifacts_to_different_s3/)
- url: https://www.reddit.com/r/aws/comments/f2lruy/codebuild_deploy_artifacts_to_different_s3/
---
I'm working on CodeBuild pipeline for a project. The pipeline builds a single artifact from a GitHub repo which is then uploaded to an S3 bucket, which is then accessed at a URL using CloudFront. All of that works fine. What I would like to do is have non-master branches or tag names (i.e., the `CODEBUILD_SOURCE_VERSION` buildspec variable, be denominated by sub-directories. For example, if building from the dev branch build to `/dev/artifact` and `/artifact` for the master branch. I can't seem to get it to work, as putting slashes in the artifact name causes the build to fail.

&amp;#x200B;

Any suggestions as to what to do? Right now I'm considering a Lambda function to move/rewrite the object in S3, but I'd like to know if there's a simpler way.
