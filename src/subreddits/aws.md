# aws
## [1][Automatic Capacity Reservations](https://www.reddit.com/r/aws/comments/f099cu/automatic_capacity_reservations/)
- url: https://www.reddit.com/r/aws/comments/f099cu/automatic_capacity_reservations/
---
Hi all 👋,

Today, AWS released a great new [update](https://aws.amazon.com/about-aws/whats-new/2020/02/amazon-ec2-adds-ability-to-easily-query-billing-information-of-amazon-machine-images/) which now allows us to retrieve the full platform type for running instances (indirectly via DescribeImages). With this, I was able to complete a small tool which automatically provisions capacity reservations to meet the accounts instance count.

This tool would be helpful for those that maintain a fleet of instances that are regularly stopped/started and need guaranteed capacity to be able to start on demand.

It's available as a deployable CloudFormation stack at: [https://github.com/iann0036/auto-capacity-reservations](https://github.com/iann0036/auto-capacity-reservations)

Comments, suggestions and general discussion welcomed!
## [2][Inherit aws resources from other cloudformaiton stack?](https://www.reddit.com/r/aws/comments/f08c0k/inherit_aws_resources_from_other_cloudformaiton/)
- url: https://www.reddit.com/r/aws/comments/f08c0k/inherit_aws_resources_from_other_cloudformaiton/
---
I write a cloudformation file -- a stack that contain ALL generate/static parts of aws resources (except iam, aws certs ..etc)

&amp;#x200B;

I want to add others stacks that would use directly resources that declared from above static stack as native instead of import them. Is there a way to achieve this?
## [3][ELI5: When should one use ElasticSearch as opposed to SQL/NoSql?](https://www.reddit.com/r/aws/comments/f00t52/eli5_when_should_one_use_elasticsearch_as_opposed/)
- url: https://www.reddit.com/r/aws/comments/f00t52/eli5_when_should_one_use_elasticsearch_as_opposed/
---
This is from AWS documentation on ElasticSearch

'Amazon Elasticsearch Service is a managed service that makes it easy to deploy, operate, and scale Elasticsearch in the AWS Cloud. Elasticsearch is a popular open-source search and analytics engine for use cases such as log analytics, real-time application monitoring, and click stream analytics'

I'am familiar with querying data with SQL(joining tables when you need data from multiple tables) and from NoSql by primary key

There's also indexes to speedup the lookup of both.

From my understanding, ElasticSearch is a search querying search on top of the data-store.

My question is when should you use ElastiSearch as opposed to just SQL and NoSQL with indexes? Can someone give an ELI5 example of this?

What value does ElasticSearch bring that the previously mentioned database solutions don't provide? With ElasticSearch, comes maintaining additional infrastructure(cluster) and complexity/costs as well
## [4][Dockerization of NodeJS Applications on Amazon Elastic Containers](https://www.reddit.com/r/aws/comments/f05s5b/dockerization_of_nodejs_applications_on_amazon/)
- url: https://blog.soshace.com/dockerization-of-node-js-applications-on-amazon-elastic-containers/
---

## [5][Running a bastion host / jumpbox on Fargate](https://www.reddit.com/r/aws/comments/f07wsu/running_a_bastion_host_jumpbox_on_fargate/)
- url: https://blog.deleu.dev/running-a-bastion-host-on-fargate/
---

## [6][Transitioning small objects to Glacier or Glacier Deep Archive will increase costs.](https://www.reddit.com/r/aws/comments/f09pzm/transitioning_small_objects_to_glacier_or_glacier/)
- url: https://www.reddit.com/r/aws/comments/f09pzm/transitioning_small_objects_to_glacier_or_glacier/
---
&gt; Before creating a lifecycle rule that transitions small objects to Glacier or Glacier Deep Archive, consider how many objects will be transitioned and how long you plan to keep the objects. Lifecycle request charges for these objects will increase your costs. 


I'm looking to backup my Synology NAS to the cloud. Between 100 GB and 200 GB. I'm thinking of doing backups every 3 to 6 months and keep these in Glacier for safe keeping. 

I started by trying to setup my Synology NAS to backup directly to the S3 Vault I configured. Apparently it's not possible. I can only backup to an S3 Bucket. I should then set this up in a way that it archives my backup to S3 Glacier's vault. Fine. However, like many times before, I got stuck. This time on the message in the title. 

From what I can see, Synology's Hyper Backup backs up my files in its own proprietary format, made up of many 50 MB files. I'm wondering what impact this will have on the costs of backing up my NAS this way. I know there is a link to the costs information page, but I get lost in that.

Can anyone offer some guidance? I had a really hard time getting this far, because I have 0 experience with online backup services, much less Amazon's. There's so much information! In short:

Am I doing this the right way? Can I configure my NAS to backup directly into Glacier? Can I set it up in a way that it doesn't divide the files into 50MB pieces? Can I upload these to an S3 bucket and then archive into glacier?
How can I restore from Glacier, if I can only access the bucket via Synology Hyper Backup?

Bonus question: how to minimize the time my data spends in the S3 bucket? I don't care much about versioning, but I find myself having to fiddle with that, as if it's impossible to avoid setting up...

I'm left with many other questions, but I believe that if I can take a few steps forward, I can figure those out. How does restore work, for example? Can I restore a single file or do I need to download the whole backup? I guess this will depend on Synology's Hyper Backup more than S3 bucket... How often can I, or should I backup to S3 - this one is for me to decide, but I need to make sense of the costs information first. etc, etc....

Thanks for shedding some light into this!
## [7][AWS VPC for Software Engineers](https://www.reddit.com/r/aws/comments/ezru46/aws_vpc_for_software_engineers/)
- url: https://blog.deleu.dev/aws-vpc-for-software-engineers/
---

## [8][Professional Services Consultant](https://www.reddit.com/r/aws/comments/f05fls/professional_services_consultant/)
- url: https://www.reddit.com/r/aws/comments/f05fls/professional_services_consultant/
---
Hey everyone,

I just wanted to know what a professional services consultant does at AWS. What does your day to day look like? Is it a technical role? Is it kind of a sales role? It would also be helpful to know what a proserv intern does.
## [9][Moving data from S3 to an on-premise data warehouse](https://www.reddit.com/r/aws/comments/f09jlc/moving_data_from_s3_to_an_onpremise_data_warehouse/)
- url: https://www.reddit.com/r/aws/comments/f09jlc/moving_data_from_s3_to_an_onpremise_data_warehouse/
---
The startup I work at has resources for an on-premise data warehouse. They are not willing to spend on Redshift due to the fact that they don't want it online 24/7. 

They want to see if setting up a data lake would be an efficient way of doing things. Basically, all data would reside in S3 on the lake, and the required data would be offloaded onto the warehouse. They also want to keep a backup of all data on S3. 

An alternative here would be to directly query the data on S3 using Athena but I want to explore other alternatives. 

This basic lake would just use just S3 and Glue (for conversion to Parquet). My question is what is an efficient way to transfer data from S3 to the on-premise warehouse? This would be done on a daily basis. 

Also wondering if the idea presented here is a step in the right direction, in terms of efficiency and cost-effectiveness. Thanks.
## [10][Beginner need help with AWS CLI](https://www.reddit.com/r/aws/comments/f08de9/beginner_need_help_with_aws_cli/)
- url: https://www.reddit.com/r/aws/comments/f08de9/beginner_need_help_with_aws_cli/
---
Hello. So I started preparing for AWS SAA exam and I'm at a point where I have to learn AWS CLI setup and basic operations. Problem is, after installing the program and checking the version, when I try to open the .exe file, it closes really quickly. Like, I have no time to do anything. What am I doing wrong?
