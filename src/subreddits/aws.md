# aws
## [1][Don't let this happen to you!](https://www.reddit.com/r/aws/comments/fkwrxv/dont_let_this_happen_to_you/)
- url: https://www.reddit.com/r/aws/comments/fkwrxv/dont_let_this_happen_to_you/
---
At work I am developing a proof-of-concept application based around exporting an Aurora RDS snapshot to S3 and querying the exported Parquet files with AWS Athena. As part of this process I had  to set up an AWS Glue database and a crawler to populate same. I kept getting "Access Denied" errors, though, when I ran the crawler, even though the crawler had the appropriate S3 permissions with regard to the bucket.

Turns out the problem was KMS. Aurora encrypts the exported files, so the IAM Role for the crawler needs the additional permission of `kms:Decrypt` for the KMS key used to encrypt the Parquet files. I had to get AWS support to look at the back-end S3 logs to figure that out. :)

Don't let this happen to you! If the thing that's writing to S3 is writing encrypted, then the thing that's reading from S3 has to have permission to decrypt stuff.
## [2][Converting to AWS: Advice and Best Practices](https://www.reddit.com/r/aws/comments/fkujsv/converting_to_aws_advice_and_best_practices/)
- url: https://www.reddit.com/r/aws/comments/fkujsv/converting_to_aws_advice_and_best_practices/
---
I am a Systems Engineer who has been given a task to prototype conversion of our physical system to AWS. I can't go into details, except to say it involves multiple servers and micro-services. Are there any common pitfalls I can avoid or best practices I should be following? I've a small amount of AWS experience, enough to launch an instance, but AWS is pretty daunting. Is there anywhere you would recommend starting?
## [3][AWSome Day](https://www.reddit.com/r/aws/comments/fl86kd/awsome_day/)
- url: https://www.reddit.com/r/aws/comments/fl86kd/awsome_day/
---
Can somebody explain to me why people feel the need to be dumping links to their LinkedIn/Instagram accounts in the chat ? What do they plan on achieving?
## [4][Lambda Synchronous](https://www.reddit.com/r/aws/comments/fl7lda/lambda_synchronous/)
- url: https://www.reddit.com/r/aws/comments/fl7lda/lambda_synchronous/
---
The lamda statement is querying DynamoDB, emailing customers with info and then updating the database. However, I need this to happen not in parallel, so each request is placed in a queue, as the database update may affect the next customers query. Any thoughts? I'm using Python for the lambda.
## [5][WFH office NAS to AWS?](https://www.reddit.com/r/aws/comments/fl31mp/wfh_office_nas_to_aws/)
- url: https://www.reddit.com/r/aws/comments/fl31mp/wfh_office_nas_to_aws/
---
OneDrive/Dropbox is probably a better solution for my wife's office, but as a thought experiment, to get her office laptop mounting her familiar NAS volumes from home, what is the AWS answer for this?

Lets assume the current local NAS could be wholly hosted on AWS instead of the office. So no VPN server at the office. Lets assume that there is a mix of Windows and MacOS clients at her office. Files are typically ~100M architecture drawings that multiple people might work on.

What are the steps. S3 sync to s3 and then setup FSx and then setup VPN client on the PCs? Or is it more involved with Active Directory?
## [6][Cortex - Open source alternative to SageMaker for model serving](https://www.reddit.com/r/aws/comments/fkmpnr/cortex_open_source_alternative_to_sagemaker_for/)
- url: https://github.com/cortexlabs/cortex
---

## [7][[Support query] How to access private IP publicly?](https://www.reddit.com/r/aws/comments/fl3ynl/support_query_how_to_access_private_ip_publicly/)
- url: https://www.reddit.com/r/aws/comments/fl3ynl/support_query_how_to_access_private_ip_publicly/
---
So i have some content that can only run on the private IP and no matter what hosts file chicanery i do i can't get it to resolve on the public IP, so how can i make it so the private IP is like the public IP?
## [8][API gateway v2 and SQS?](https://www.reddit.com/r/aws/comments/fkv9aq/api_gateway_v2_and_sqs/)
- url: https://www.reddit.com/r/aws/comments/fkv9aq/api_gateway_v2_and_sqs/
---
Does anyone know if you can set up SQS as the target for the new apigatewayv2 API AWS_PROXY? Or maybe you’d use HTTP_PROXY with the queue URL?
## [9][Is Workmail still being developed?](https://www.reddit.com/r/aws/comments/fko8vm/is_workmail_still_being_developed/)
- url: https://www.reddit.com/r/aws/comments/fko8vm/is_workmail_still_being_developed/
---
We use Workmail as our company's email service. It is generally ok, but it doesn't seem to have had any new features since its release 3 years ago. The interface is clunky but we expected this to improve. 

We are concerned that there seems to be a lack of development on the platform. It is very basic but we pay almost the same amount as we would pay for Gmail. 

Also, we receive calendar invites for AWS events, from AWS, and the only options we have are Outlook, Google or iCal.  No mention of Workmail or any web based integration to their events. This just comes across as a product that AWS doesn't even recognise.

Does anybody use Workmail and is there any reason not to move onto Gmail?
## [10][Any reasons why the /r/aws moderators are removing the link posts?](https://www.reddit.com/r/aws/comments/fl0pv3/any_reasons_why_the_raws_moderators_are_removing/)
- url: https://www.reddit.com/r/aws/comments/fl0pv3/any_reasons_why_the_raws_moderators_are_removing/
---
Recently I have submitted few blog links (redshift related resources) but after a day the post got removed by the moderator. It is an AWS related external link. But I don't know why they are removing it.
