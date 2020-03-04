# aws
## [1][RDS CA... are you prepared?](https://www.reddit.com/r/aws/comments/fd2m10/rds_ca_are_you_prepared/)
- url: https://www.reddit.com/r/aws/comments/fd2m10/rds_ca_are_you_prepared/
---
So, we’re ~36 hours from the RDS CA expiration as detailed in the [AWS blog](https://aws.amazon.com/blogs/database/amazon-rds-customers-update-your-ssl-tls-certificates-by-february-5-2020/).

TL;DR.. if you use SSL/TLS to connect to RDS or DoumentDB, you need to update your client AND your RDS instance by 9:11am UTC on 5th March. 

Are you prepared?

* edit: link formatting.
* edit 2: fun fact: MySQL Java Connector v5.1.38 or later, and MySQL Java Connector v8.0.9 or later all default to SSSL/TLS. So you might be using it without knowing!! 
They perform partial certificate verification and will fail to connect if the database server certificate is expired.
## [2][How do you manage SSO access into multi AWS accounts?](https://www.reddit.com/r/aws/comments/fdaerx/how_do_you_manage_sso_access_into_multi_aws/)
- url: https://www.reddit.com/r/aws/comments/fdaerx/how_do_you_manage_sso_access_into_multi_aws/
---
As the title (poorly) asks, how does your org manage SSO access across multiple AWS accounts, with multiple potential roles within each account. For example, our org has 20 AWS accounts, with multiple different platforms and different teams that manage each. Further within each account, there may be some people that have admin access, some people that have read only access, and others with something in between.

The easy answer is to create an AD group for each account/role combination and then assign them to the users as necessary. However, this becomes a management nightmare, as we'll easily end up with 40-60 different groups, and some users being added to 5-20 different groups depending on their level of access.

Is there a more clever way to sort this? We tried thinking of a dynamic way, but are hardpressed to think of anything to base it on.
## [3][AWS Secrets Manager - Anyone using it?](https://www.reddit.com/r/aws/comments/fd08ft/aws_secrets_manager_anyone_using_it/)
- url: https://www.reddit.com/r/aws/comments/fd08ft/aws_secrets_manager_anyone_using_it/
---
Hi AWS Community - 

I have been investigating AWS Secrets Manager as a solution for us to store, retrieve, and rotate secrets. I've done some digging on what the product is and I've tried using it a bit. However, it seems that the real magic of this product is the ability to rotate the secret and I was hoping for something, I guess, a little more magical? I am trying to setup a PoC for rotating SQL Server credentials in RDS and as I'm working through this solution (building the lambda, getting the lambda to import pymssql, adding iam policies so everything can talk to everything it needs to) I am asking myself is this really worth it? To me, what is this product really other than a place to store, retrieve some secrets? It seems the magic from the rotation is just the user writing their own lambda function to do everything. Why not just use Parameter store and still have your own lambda rotating secrets?

Is anyone out there using Secrets Manager and can share their experiences with it? For me, it seems like the work required to rotate the secrets makes this product less appealing. I've tried searching the terraform module registry (we are a terraform shop) for some example modules and I'm not seeing anything that has rotation build in.

Thanks!
## [4][RDS SecretARN](https://www.reddit.com/r/aws/comments/fdcb7w/rds_secretarn/)
- url: https://www.reddit.com/r/aws/comments/fdcb7w/rds_secretarn/
---
AWS is keeping it secret even from me.

I'm trying to use boto3 to access RDS. Where can I find the ARNsecret?
## [5][Filling CDK gaps](https://www.reddit.com/r/aws/comments/fd1m2n/filling_cdk_gaps/)
- url: https://www.reddit.com/r/aws/comments/fd1m2n/filling_cdk_gaps/
---
For those of you using the AWS CDK, how do you go about filling the gaps for parts of it that are incomplete? Do you have a separate CFN file that you deploy first? Do you just use the console and create those resources that you need?

I really like using the CDK, it makes so much more sense to me than writing a giant CFN file but it seems like it's not really production-ready yet.
## [6][Cannot create more than 3 customized metrics filter](https://www.reddit.com/r/aws/comments/fd9reu/cannot_create_more_than_3_customized_metrics/)
- url: https://www.reddit.com/r/aws/comments/fd9reu/cannot_create_more_than_3_customized_metrics/
---
It's weird but I can only create three customized metrics filter: CloudWatch -&gt; CloudWatch Logs -&gt; Log groups -&gt; Metrics filter

When I try to create more than 3 (the 4th one) it will overwrite the last created metric.

&amp;#x200B;

As this aws doc [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_limits.html) the quota for Custom metrics is "no quota" -- unlimited. How's that?
## [7][Is Athena a suitable option to use it as a webapp db?](https://www.reddit.com/r/aws/comments/fcz9vd/is_athena_a_suitable_option_to_use_it_as_a_webapp/)
- url: https://www.reddit.com/r/aws/comments/fcz9vd/is_athena_a_suitable_option_to_use_it_as_a_webapp/
---
I mean, i know the purpose of athena is defined as an analytics tools to query your data stored in s3.

But i dont know if this description fits the scenario in wich i use it as a common db to receive queries made by users in a web application for example.

Is this possible or athena is more orientated to make querys over your data with analytics purposes?  


Thanks.
## [8][is it possible to leave no trail behind in this case?](https://www.reddit.com/r/aws/comments/fcthjq/is_it_possible_to_leave_no_trail_behind_in_this/)
- url: https://www.reddit.com/r/aws/comments/fcthjq/is_it_possible_to_leave_no_trail_behind_in_this/
---
Hello!

My instances are locked behind a security group that only allows traffic through ports 80 and 443. When I need access, I use a custom batch script to allow traffic through ports 22 and 5432 exclusively to my IP address. Then I proceed to access it with putty using my key pair. Once I'm done, I use another custom script to close ports 22 and 5432.

AWS has CloudTrail, which records all activity for your account. I've noticed that I can monitor security group changes (such as those that I explained above) and I want to know if having these records is enough to tell if someone got into my instance.

So, my questions are:

1) Can anyone access the instances behind that security group without having to open port 22 AND physically having access to my key pair file?

2) Can I trust CloudTrail records, so that all breaches are guaranteed to be logged just like normal access?

Thanks in advance!
## [9][Help explaining strange S3 copy behavior from Java SDK?](https://www.reddit.com/r/aws/comments/fd7v2h/help_explaining_strange_s3_copy_behavior_from/)
- url: https://www.reddit.com/r/aws/comments/fd7v2h/help_explaining_strange_s3_copy_behavior_from/
---
Hi All, I observed some strange behavior using the Java SDK to copy objects between a source bucket and destination bucket (within or across regions), and was wondering if anyone could help shed some insight into this behavior.


Basically, we have an app that is containerized and runs in an EKS cluster, and uses the 1.X Java AWS SDK for S3, particularly the [copy](https://docs.aws.amazon.com/AWSJavaSDK/latest/javadoc/com/amazonaws/services/s3/transfer/TransferManager.html) method to do the copy. The app also pulls messages off a input SQS queue, each message containing the source object and destination bucket and object to be copied to.

The interesting part and the crux of this post, is the performance of copying objects. Copying a 10.8GB file within a region, serially (one at a time, after another), resulted in a average throughput of **~365 MB/s** (sample size of 10). However, when multiple copies originate **from the same app** (the apps can pull up to 10 messages at a time off the queue and submits them to a thread pool with calls the copy method concurrently), the throughput average drops to **~100 MB/s.**. This is a bit strange to me, since behind the scene the copy method is calling the [upload-part-copy](https://docs.aws.amazon.com/cli/latest/reference/s3api/upload-part-copy.html) api of S3. I am almost certain that the object content *does not* get streamed through our application's memory space and is instead handling by S3, so it is a bit perplexing as to why adding concurrent uploads really affects our speeds. I also monitored the CPU/memory usage of the pod and it didn't seem to grow as concurrent uploads were added.

Interestingly as well, when concurrent uploads are performed instead on different pods/instances (instead of from the same app), the throughput is very high (matching or exceeding our serial times).  Just was wondering if anyone has every experienced something like this and how it could be explained.
## [10][Looking for a script/process to download user details and their tags.](https://www.reddit.com/r/aws/comments/fd7nys/looking_for_a_scriptprocess_to_download_user/)
- url: https://www.reddit.com/r/aws/comments/fd7nys/looking_for_a_scriptprocess_to_download_user/
---
So we have several dozen users logging into our AWS dev account working on various development-related projects. Every user when added to the environment are tagged, the tag contains team id and their manager id, so they are just the tags feature in the IAM for that user.

I realise the feature exists within the AWS console, but I want to automate this process, for this, I have created an EC2 instance on Ubuntu, the plan is to evoke the scripts through a regular shell/python.. (sic..) method, and run it through a cron job. The excel report is then emailed out to the auditors bi-monthly.
