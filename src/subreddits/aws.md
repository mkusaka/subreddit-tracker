# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Week of Nov 9th - What have you learned recently about AWS?](https://www.reddit.com/r/aws/comments/jqya79/week_of_nov_9th_what_have_you_learned_recently/)
- url: https://www.reddit.com/r/aws/comments/jqya79/week_of_nov_9th_what_have_you_learned_recently/
---
Share your learnings with the community
## [3][Welcome to AWS Storage Day 2020](https://www.reddit.com/r/aws/comments/jr4ggk/welcome_to_aws_storage_day_2020/)
- url: https://aws.amazon.com/blogs/aws/welcome-to-aws-storage-day-2020/
---

## [4][New – Export Amazon DynamoDB Table Data to Your Data Lake in Amazon S3, No Code Writing Required](https://www.reddit.com/r/aws/comments/jr7m8e/new_export_amazon_dynamodb_table_data_to_your/)
- url: https://aws.amazon.com/blogs/aws/new-export-amazon-dynamodb-table-data-to-data-lake-amazon-s3/
---

## [5][Amazon ECS now supports Internet Protocol Version 6 (IPv6) in awsvpc networking mode](https://www.reddit.com/r/aws/comments/jr88ao/amazon_ecs_now_supports_internet_protocol_version/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/amazon-ecs-supports-ipv6-in-awsvpc-networking-mode/
---

## [6][Transition from VPS and cost comparison](https://www.reddit.com/r/aws/comments/jrjxcg/transition_from_vps_and_cost_comparison/)
- url: https://www.reddit.com/r/aws/comments/jrjxcg/transition_from_vps_and_cost_comparison/
---
Did you have experience deploying a web application or service on VPS that costed you anything less than $100 per VPS, and transitioned to AWS non vps or lightspeed offering, where it now costs you less per month, or you are getting any additional benefit from the transition?

if you transitioned to aws lambda or any other per request + bandwidth + storage options.

if you have any numbers to share, like average requests per day or storage use per day or db operations.

it seems to me, that it will cost more, and be more limiting, as I can't just run any long running tasks to generate a report or create chart data when the server is idling.
## [7][Exporting DynamoDB table data to Amazon S3](https://www.reddit.com/r/aws/comments/jr7mze/exporting_dynamodb_table_data_to_amazon_s3/)
- url: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataExport.html
---

## [8][Athena CTAS query with column names](https://www.reddit.com/r/aws/comments/jrhfpc/athena_ctas_query_with_column_names/)
- url: https://www.reddit.com/r/aws/comments/jrhfpc/athena_ctas_query_with_column_names/
---
Hey hey! 

I’m working on a project that mainly uses Athena to query and transform data in S3, and allowing a client to download the transformed data from a S3 bucket. I’m currently using Athena’s CTAS query to output the transformed data to a destination S3 bucket. However, the column names are not present. I tried hard coding the column names and UNION ALL the table but Athena CTAS doesn’t preserve the order due to it’s distributed nature so the column headers doesn’t appear as the first row (even after ORDER BY). Does anyone have experience working on similar set up? Keen to pick your brains on it! 

Cheers

Edit: I realised I only need the CSV query result from a SELECT query since it comes with column headers and nicely formatted. So simply triggering an athena SELECT query with a specified output bucket achieves my goal. An S3 event is triggered upon the .csv file landing in the bucket.
## [9][New – Export Amazon DynamoDB Table Data to Your Data Lake in Amazon S3, No Code Writing Required | Amazon Web Services](https://www.reddit.com/r/aws/comments/jr7s0i/new_export_amazon_dynamodb_table_data_to_your/)
- url: https://aws.amazon.com/blogs/aws/new-export-amazon-dynamodb-table-data-to-data-lake-amazon-s3/
---

## [10][AWS Migration with CloudEndure](https://www.reddit.com/r/aws/comments/jrj4ml/aws_migration_with_cloudendure/)
- url: https://www.reddit.com/r/aws/comments/jrj4ml/aws_migration_with_cloudendure/
---

How does CloudEndure decide what type of server to create in the target environment?

(Since CloudEndure does not generate any EC2 recommendations based on the existing server)
## [11][Remote Developing with VsCode on AWS Cloud9 EC2 Instance](https://www.reddit.com/r/aws/comments/jrinwe/remote_developing_with_vscode_on_aws_cloud9_ec2/)
- url: https://www.reddit.com/r/aws/comments/jrinwe/remote_developing_with_vscode_on_aws_cloud9_ec2/
---
With this setup, all the CPU used to run/debug, all the docker images are just running in the cloud and are not exhausting thee laptop resources anymore.  --&gt; [Article](https://itnext.io/remote-developing-with-vscode-on-aws-cloud9-ec2-instance-e18134af353f?source=friends_link&amp;sk=9a4362dc515ca8d090fdbf69047c9b23)
## [12][AWS Guard Duty Delegation to Security account](https://www.reddit.com/r/aws/comments/jrcqk8/aws_guard_duty_delegation_to_security_account/)
- url: https://www.reddit.com/r/aws/comments/jrcqk8/aws_guard_duty_delegation_to_security_account/
---
Hi,

How to automate guard duty reporting from all aws accounts to central security aws account that security team owned ?

  
Ideally, if we add new aws accounts to our environment how should be make sure new aws account guard duty service is added to central security account for reporting finding into central secuirty account ?

can someone advise how can we achieve above ?

Here few thoughts :

1. can we implement using SCP policies at organization level ? and automate for new aws accounts that might added in future ?  
if so, can you point to resources , or sample scp policies please 
2. Or run daily job to add new account guard duty service reporting to security account ?  
(probably not ideal i guess )
3. Definitely avoid manually .
