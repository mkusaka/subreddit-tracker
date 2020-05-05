# aws
## [1][AWS VP &amp; Distinguished Engineer Tim Bray resigns over worker treatment](https://www.reddit.com/r/aws/comments/gde37f/aws_vp_distinguished_engineer_tim_bray_resigns/)
- url: https://www.tbray.org/ongoing/When/202x/2020/04/29/Leaving-Amazon
---

## [2][Cache for JSON data?](https://www.reddit.com/r/aws/comments/gdwx9h/cache_for_json_data/)
- url: https://www.reddit.com/r/aws/comments/gdwx9h/cache_for_json_data/
---
I am developing an app and one of the processes (which is critical) is invoking a web service which returns JSON information. Every invocation of this JSON service involves a cost (say 10p).

The information returned (about a vehicle) I want to cache in AWS. If a customer returns to the app with the same vehicle details, then it can be retrieved from the cache and avoid the company being charged the fee.

Is there a suitable service for this that would also be more performant? I was/am looking at something like S3 or Dynamo DB.

Thanks!
## [3][AWS ELB in two AZ's costs](https://www.reddit.com/r/aws/comments/gdwrt7/aws_elb_in_two_azs_costs/)
- url: https://www.reddit.com/r/aws/comments/gdwrt7/aws_elb_in_two_azs_costs/
---
In Simply Month calculator i can type the number of ALB's that i want to calculate. Is ALB with a NIC in AZ1 and  a NIC in AZ2 for Multi AZ one ALB Loadbalancer regarding calculations or counts this as 2 ALBs?
## [4][NoSQL Workbench for DynamoDB adds support for Linux--data modeling, querying, testing](https://www.reddit.com/r/aws/comments/gdm9ok/nosql_workbench_for_dynamodb_adds_support_for/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/05/nosql-workbench-for-dynamodb-adds-support-for-linux/
---

## [5][Having Trouble Mounting S3 Bucket to EC2 Instance Using s3fs](https://www.reddit.com/r/aws/comments/gdxgh3/having_trouble_mounting_s3_bucket_to_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/gdxgh3/having_trouble_mounting_s3_bucket_to_ec2_instance/
---
* Ubuntu 18.04 (ARM)
* s3fs v1.86
* I've double checked to ensure the user has a policy that allows access to the S3 bucket (I've actually given the account admin while I try to work through this)
* EC2 instance has a role that allows access to the S3 bucket (again admin as I try to resolve this)
* I created the password file with my access key ID and secret access key and CHMOD 600 for the password file
* I run **sudo s3fs my-bucket /home/ubuntu/efs\_uploads -o passwd\_file=${HOME}/.passwd-s3fs -o dbglevel=info -f -o curldbg** and it runs until it hangs up after the following two lines 

\[INF\]  curl.cpp:RequestPerform(2455): HTTP response code 200

\[INF\] curl.cpp:ReturnHandler(318): Pool full: destroy the oldest handler

* In version 1.84 of s3fs it stopped on the first line and never reached the line about curl.cpp.ReturnHandler(318)...

I've seen others post about this on GitHub but I haven't seen any solutions
## [6][Clarity on EBS Incremental Snapshots &amp; Cross-Region Copy Pricing](https://www.reddit.com/r/aws/comments/gdxge4/clarity_on_ebs_incremental_snapshots_crossregion/)
- url: https://www.reddit.com/r/aws/comments/gdxge4/clarity_on_ebs_incremental_snapshots_crossregion/
---
Let's say I have a 10GB EBS volume, using 5GB of capacity. I take daily backups, where there is negligible change to the disk, and I copy these backups from Singapore to Sydney.

Assumption #1: The original snapshot size is \~5GB as it only takes consumed data.

Assumption #2: My daily Singapore snapshots, all together, consume only \~5GB as no changes are being made.

Assumption #3: My daily Sydney snapshot copies, all together, consume only \~5GB as no changes are made **and** I'm only charged \~5GB of transfer to Sydney as all future backups have no change.

I'm fairly confident in #1 and #2, and I'm completely unsure of #3.

Any assistance would be wonderful, thanks!
## [7][AWS VPC for Software Engineers](https://www.reddit.com/r/aws/comments/gdv6kv/aws_vpc_for_software_engineers/)
- url: https://blog.deleu.dev/aws-vpc-for-software-engineers/
---

## [8][Question on cost effective exporting of CloudWatch Logs to S3](https://www.reddit.com/r/aws/comments/gdo6mj/question_on_cost_effective_exporting_of/)
- url: https://www.reddit.com/r/aws/comments/gdo6mj/question_on_cost_effective_exporting_of/
---
Hello,

I setup a new service that has to write it's logs to AWS CloudWatch in order to be ingested into our SIEM.  I don't want to keep them inside CloudWatch for more than a few days due to the sheer size of the logs, so I want to automate exporting them to S3 for long term storage.

I know there are several methods to do this, with the following coming to mind:

* Lambda/Step Function to automate exporting to S3.
* Kinesis Firehose

Which option is the most cost effective, at about \~200 events per second on average and growing?

  
Thanks!
## [9][Webscraper on steroids, using 2,000 Lambda invokes to scan 1,000,000 websites in under 7 minutes.](https://www.reddit.com/r/aws/comments/gd6xss/webscraper_on_steroids_using_2000_lambda_invokes/)
- url: /r/Python/comments/gcq18f/a_serverless_web_scraper_built_on_the_lambda/
---

## [10][Allowing external users to query redshift in browser](https://www.reddit.com/r/aws/comments/gdt8uh/allowing_external_users_to_query_redshift_in/)
- url: https://www.reddit.com/r/aws/comments/gdt8uh/allowing_external_users_to_query_redshift_in/
---
We have our data stored in Redshift but recently were asked to find a solution to allow clients to query their own databases that we host (read access). Are there any  browser tools that allow this direct querying where we don't need to pay per query? (since we're already paying for redshift...) Let me know if I need to provide any additional information.
