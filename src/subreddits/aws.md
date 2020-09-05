# aws
## [1][PROTIP: Watch out for stranded multipart uploads - I was billed for 220GB of "invisible" objects for months](https://www.reddit.com/r/aws/comments/immer3/protip_watch_out_for_stranded_multipart_uploads_i/)
- url: https://www.reddit.com/r/aws/comments/immer3/protip_watch_out_for_stranded_multipart_uploads_i/
---
I used rclone to sync gdrive to a new S3 AWS bucket. Due to network connectivity issues, the rclone had to be retried multiple times (it was a big gdrive)

Once the data was cleaned up and confirmed ready for deep archive, I transitioned the storage class to S3 DEEPARCHIVE for all visible objects.

I recently looked at the bucket and noticed 220GB of Standard storage class objects - but, here is the kicker, there were none visible ! 

Googled the issue and found this "advisory". Basically, I think I paid peanuts for the TB+ of data I stored in DEEPARCHIVE but paid like 30X more for the "invisible" multipart uploads. I didn't know they needed to be cleaned up, and AWS would have happily billed me for years more.

[https://aws.amazon.com/premiumsupport/knowledge-center/s3-reduce-costs/](https://aws.amazon.com/premiumsupport/knowledge-center/s3-reduce-costs/)
## [2][Pentagon says it will stick with Microsoft for $10 billion JEDI cloud contract](https://www.reddit.com/r/aws/comments/imo110/pentagon_says_it_will_stick_with_microsoft_for_10/)
- url: https://www.cnbc.com/2020/09/04/pentagon-says-it-will-stick-with-microsoft-for-jedi-cloud-contract.html
---

## [3][Aurora with autoscaling vs. Aurora Serverless](https://www.reddit.com/r/aws/comments/imwyk5/aurora_with_autoscaling_vs_aurora_serverless/)
- url: https://www.reddit.com/r/aws/comments/imwyk5/aurora_with_autoscaling_vs_aurora_serverless/
---
Could anyone explain the main differences between a Postgres database on Aurora with autoscaling switched on vs. Aurora Serverless? I've read through a lot of docs and the difference is not clear to me. Thanks!
## [4][AWS Named as a Cloud Leader for the 10th Consecutive Year in Gartner’s Infrastructure &amp; Platform Services Magic Quadrant](https://www.reddit.com/r/aws/comments/imlnux/aws_named_as_a_cloud_leader_for_the_10th/)
- url: https://aws.amazon.com/blogs/aws/aws-named-as-a-cloud-leader-for-the-10th-consecutive-year-in-gartners-infrastructure-platform-services-magic-quadrant/
---

## [5][Migrating a Windows VM to EC2 using Server Migration Service](https://www.reddit.com/r/aws/comments/in14r9/migrating_a_windows_vm_to_ec2_using_server/)
- url: https://www.reddit.com/r/aws/comments/in14r9/migrating_a_windows_vm_to_ec2_using_server/
---
I've been migrating linux and windows vms using Server Migration Service. How or where can I get more detailed error information as to why a replication job fails with the error below?

**"Job status message**FirstBootFailure: This import request failed because the instance failed to boot and establish network connectivity."

I've already turned off the antivirus running on the source. However, the error is not helpful. This is my 5th attempt. It worked on other Windows vms I've moved.
## [6][What is the difference between elastic beanstalk and fargate?](https://www.reddit.com/r/aws/comments/in0f5o/what_is_the_difference_between_elastic_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/in0f5o/what_is_the_difference_between_elastic_beanstalk/
---
What is the difference between elastic beanstalk and fargate? We can deploy docker on both services and we don't need to provison anything in both of these services as everything is managed by them (PaaS).
## [7][Add security policy to allow ELB/ALB operations - HELP](https://www.reddit.com/r/aws/comments/imzxvt/add_security_policy_to_allow_elbalb_operations/)
- url: https://www.reddit.com/r/aws/comments/imzxvt/add_security_policy_to_allow_elbalb_operations/
---
Hey guys,

I primarily work as a Cloud Engineer in a staging environment at my job and I've now been tasked to update our security policy in our production environment to allow elbv2 automation. Now, our current IAM role through our cloud suite portal has all permissions needed, but I have to go in and update our IAM role template. 

I have to make these changes  to the security policy to reflect these updates:  

*Internal Microservices ELB/ALB*  
*ConfigureHealthCheck*

*Web server ALBs*  
*RegisterTargets*  
*DeregisterTargets*  
*DescribeTargetHealth*

I have a copy of  production iam-template that I believe has what I need, but I have no idea what I need to do to make these changes on my iam-template and deploy it. Done some research online and through aws documentation, just stumped on what I'm trying to accomplish. 

Any assistance?
## [8][upload custom logs in s3 to cloudwatch for metrics monitoring](https://www.reddit.com/r/aws/comments/imxdcq/upload_custom_logs_in_s3_to_cloudwatch_for/)
- url: https://www.reddit.com/r/aws/comments/imxdcq/upload_custom_logs_in_s3_to_cloudwatch_for/
---
Hi, 

I created a custom app that automatically uploads logs to s3.

Is there a way to push those logs to cloudwatch from s3 for analysis and alerting?

I'm aware that I can use a cloudwatch agent to push directly to cloudwatch from the app but there are complications involved in that option.

Thank you!
## [9][[QUESTION] About CloudFront and livestreaming a video for 20000+](https://www.reddit.com/r/aws/comments/imw426/question_about_cloudfront_and_livestreaming_a/)
- url: https://www.reddit.com/r/aws/comments/imw426/question_about_cloudfront_and_livestreaming_a/
---
I'm doing some research for an upcoming event, and one part is installing Mozilla Hubs-Cloud on AWS, and another streaming video inside Hubs-Cloud. I was thinking on using CloudFront to stream a video for the visitors.

I'm still trying to understand how the costs are going to look like for the streaming (I already figured out the costs for Moz. Hubs-Cloud), since CloudFront says it's free to go up to 50 gb, and if the stream is going to be requested from inside Mozilla Hubs-Cloud to the visitors.  


Anyone has insight on how the calculations for each viewer is made?  


Thanks in advance.
## [10][Changes for AWS customers in Brazil](https://www.reddit.com/r/aws/comments/imfifi/changes_for_aws_customers_in_brazil/)
- url: https://i.redd.it/jmtles60q4l51.png
---

