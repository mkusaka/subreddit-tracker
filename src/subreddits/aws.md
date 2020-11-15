# aws
## [1][Moving AWS S3 objects to an Infrequent Access storage class without going bankrupt](https://www.reddit.com/r/aws/comments/juhz2m/moving_aws_s3_objects_to_an_infrequent_access/)
- url: https://blog.doit-intl.com/moving-aws-s3-objects-to-an-infrequent-access-storage-class-without-going-bankrupt-e4de168cd05c
---

## [2][AWS Network Load Balancer now supports IPv6](https://www.reddit.com/r/aws/comments/ju10mh/aws_network_load_balancer_now_supports_ipv6/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/network-load-balancer-supports-ipv6/
---

## [3][Name a security challenge you want to see automation for](https://www.reddit.com/r/aws/comments/jugyc9/name_a_security_challenge_you_want_to_see/)
- url: https://www.reddit.com/r/aws/comments/jugyc9/name_a_security_challenge_you_want_to_see/
---
Wondering what security challenges you guys have encountered, that you wish automation already existed.

In short, interested in understanding any security tasks you have to do manually because there is no current automation within the AWS platform.

Thanks a lot in advance!
## [4][Would Cassandra be a better option than Aurora for daily bulk inserts and frequent API calls?](https://www.reddit.com/r/aws/comments/juh3gd/would_cassandra_be_a_better_option_than_aurora/)
- url: https://www.reddit.com/r/aws/comments/juh3gd/would_cassandra_be_a_better_option_than_aurora/
---
I have Spark running on AWS EMR, and I probably insert 10k+ new and upsert entries in Aurora right now. The problem is Spark takes around 10% of the time for the transformations, around 90% goes on with the insertion in Aurora. This happens once every day, rest there are regular API calls going on with Aurora to fetch data. Would Cassandra be a quicker and feasible option, performance-wise?   


PS: I cannot overspend. Have to work on the budget around what is being costed for Aurora
## [5][I assume there is no way for an AWS to see what process/services you're running on your outside AWS PC?](https://www.reddit.com/r/aws/comments/juj2sp/i_assume_there_is_no_way_for_an_aws_to_see_what/)
- url: https://www.reddit.com/r/aws/comments/juj2sp/i_assume_there_is_no_way_for_an_aws_to_see_what/
---
Or is there?

&amp;#x200B;

I.e, if you're running AWS on a desktop, your own one
## [6][Amazon Athena adds support for running SQL queries across relational, non-relational, object, and custom data sources.](https://www.reddit.com/r/aws/comments/jtxnwl/amazon_athena_adds_support_for_running_sql/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/aws-what-s-new-for-athena-federated-query/
---

## [7][AWS master account](https://www.reddit.com/r/aws/comments/juhyft/aws_master_account/)
- url: https://www.reddit.com/r/aws/comments/juhyft/aws_master_account/
---
Recently I've seen the AWS master account is shifting to the AWS management account. 

Is it going to bring any changes to the functionalities? or they are just changing the names or everything will be the same as before?
## [8][Build a Serverless GraphQL API on AWS [Tutorial]](https://www.reddit.com/r/aws/comments/ju33fe/build_a_serverless_graphql_api_on_aws_tutorial/)
- url: https://lucas-le-ray.com/blog/build-serverless-graphql-api
---

## [9][Is there a way for my Lambda function to hit an external API as well as my Redis ElastiCache?](https://www.reddit.com/r/aws/comments/jubgju/is_there_a_way_for_my_lambda_function_to_hit_an/)
- url: https://www.reddit.com/r/aws/comments/jubgju/is_there_a_way_for_my_lambda_function_to_hit_an/
---
Hi,

I want my site to hit an API, providing an ID. I want to check my ElastiCache, which lives in a VPC, for that ID. If it exists, return the value associated with the key; if not, hit an external API and put the result in the cache. 

My understanding is that a Lambda function needs to live in a VPC to access the cache, since that also lives in a VPC. And a Lambda in a VPC does not have Internet access, so it can't hit the external API.

Are there any workarounds? Is there anything I can use other than a Lambda function that allows me to achieve this functionality?
## [10][Using AWS S3 to download files from same origin](https://www.reddit.com/r/aws/comments/judckz/using_aws_s3_to_download_files_from_same_origin/)
- url: https://www.reddit.com/r/aws/comments/judckz/using_aws_s3_to_download_files_from_same_origin/
---
I have an S3 bucket that backs a website where videos are uploaded.

The bucket is configured properly, objects can be uploaded and accessed, etc.

The issue I'm running into is how to download these files from S3. When I expose the link to the S3 object in an HTML a ref download, it doesn't prompt a download in Firefox but instead just opens the file in the browser.

Some troubleshooting tells me this is because Firefox does not allow cross origin downloading, e.g. from an `amazonaws` domain on my domain.

I added the following convention to my S3 bucket: `media.&lt;mydomain&gt;.com` which points to `s3.amazonaws.com` which works but because it's simply a redirect to something in the effect of [`media.mydomain.com.s3-us-west-1.amazonaws.com`](https://media.mydomain.com.s3-us-west-1.amazonaws.com), it's still considered cross origin.

I've looked for solutions to this but can't seem to find any. I know obviously this is done frequently, as I've personally seen lots of downloads coming from AWS services but I'm not sure what to use at this point to make this work.

&amp;#x200B;

Solved Update: If you set `content-dispostion` in the `metadata` as `attachment`, it will prompt the download. 
