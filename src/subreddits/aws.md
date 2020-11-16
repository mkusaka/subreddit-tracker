# aws
## [1][Demystifying AWS VPC](https://www.reddit.com/r/aws/comments/juojf4/demystifying_aws_vpc/)
- url: https://scorpil.com/post/aws-vpc/
---

## [2][Security group referencing in AWS Transit Gateway](https://www.reddit.com/r/aws/comments/jv5yc3/security_group_referencing_in_aws_transit_gateway/)
- url: https://www.reddit.com/r/aws/comments/jv5yc3/security_group_referencing_in_aws_transit_gateway/
---
In my company we have multiple VPCs in an AWS dev/test account that each have 2 NAT gateways and 11 VPC endpoints across two AZs (all managed via Terraform). From a cost perspective the figures are mounting up and I've been tasked with reducing these costs.

From looking at AWS Transit Gateway it seems like it is the answer to our prayers - we can move the NAT gateways and VPC endpoints into a new VPC and all the existing VPCs can route their traffic to this new VPC, however, the inability to reference security group IDs in the security group rules seems like a showstopper.

An example from our infrastructure is: we are using ECS and run all our services in the awsvpc network mode so we can create granular security group rules per service. One particular service (an email notification sending service) needs to access the SQS VPC endpoint and also a couple of public internet CIDRs from a non-AWS SMTP provider. Right now this is easily done via a rule to allow:

    egress port 443 -&gt; SQS VPC endpoint security group
    egress port 587 -&gt; CIDRs 1.2.3.4/32 and 5.6.7.8/32 

If I create multiple subnets (e.g. one for the SQS endpoints and another for the NAT gateways) I can add the ECS service into both these subnets and can understand how the routing would work, but what security group rules would I apply without opening up the outbound connectivity more than it currently is?
## [3][Dynamo Exports May Get Your Data Out… But This Is Still the Fastest Way to Move Data In!](https://www.reddit.com/r/aws/comments/jv6j61/dynamo_exports_may_get_your_data_out_but_this_is/)
- url: https://medium.com/@psingman/dynamo-exports-may-get-your-data-out-but-this-is-still-the-fastest-way-to-move-data-in-5bcd9748cc00?sk=7664e57767488ec9c5d7bd80884e5e56
---

## [4][AWS Personalize recommending recently purchased items for users](https://www.reddit.com/r/aws/comments/jv0at3/aws_personalize_recommending_recently_purchased/)
- url: https://www.reddit.com/r/aws/comments/jv0at3/aws_personalize_recommending_recently_purchased/
---
 Cross-posted at r/MachineLearning

I've noticed that AWS Personalize (user-personalization recipe) is recommending products to customers who have already recently purchased that product.

I thought maybe applying a filter like below, following the suggestion here [https://aws.amazon.com/blogs/machine-learning/enhancing-recommendation-filters-by-filtering-on-item-metadata-with-amazon-personalize/](https://aws.amazon.com/blogs/machine-learning/enhancing-recommendation-filters-by-filtering-on-item-metadata-with-amazon-personalize/) would help, but no dice.

EXCLUDE ItemID WHERE interactions.event\_type in ("\\\*\\")

Can anyone offer some guidance on how to only recommend products that the customer has not purchased?
## [5][IAM Tutorial: How to Copy between Encrypted S3 Buckets Cross Account](https://www.reddit.com/r/aws/comments/jutvvo/iam_tutorial_how_to_copy_between_encrypted_s3/)
- url: https://www.reddit.com/r/aws/comments/jutvvo/iam_tutorial_how_to_copy_between_encrypted_s3/
---
My [last post](https://www.reddit.com/r/aws/comments/jhu004/aws_iam_introduction/) here was pretty successful, so I wanted to do a follow up tutorial for something I get asked a lot about at work

&amp;#x200B;

[https://towardsdatascience.com/how-to-copy-between-encrypted-s3-buckets-cross-account-e4e3096d1a8a](https://towardsdatascience.com/how-to-copy-between-encrypted-s3-buckets-cross-account-e4e3096d1a8a)

&amp;#x200B;

If you're past your limit on medium, you can also read on [my personal website](https://evankozliner.com/texts/how-to-copy-between-encrypted-s3-buckets-cross-account-e4e3096d1a8a/)
## [6][Hey all, trying to gauge interest here. How many of you would be interested in a codesandbox-like website for Serverless development? I know that I've put a lot of time and effort into learning the serverless ecosystem, and I feel like something like a codesandbox would be helpful.](https://www.reddit.com/r/aws/comments/jup247/hey_all_trying_to_gauge_interest_here_how_many_of/)
- url: https://www.reddit.com/r/aws/comments/jup247/hey_all_trying_to_gauge_interest_here_how_many_of/
---
Just trying to see if anyone else would be interested. In the same way that codesandbox can make it easier to get started and collaborate with teams, I imagine a sandbox for serverless development would be equally helpful. It can be hard to get started with all the configuration files and IAM permissions when you're new to the AWS ecosystem (not that all serverless is AWS, but a good portion of it is). There are also just a ton of complex services, like API Gateway, RDS proxy, Batch, and Kinesis, Appsync, AWS Organizations, X Ray, etc. that could be a bit easier to work with.

All code would be exportable and would sync up with github just like codesandbox.

Thoughts?  


Edit: Some template ideas:

* Serverless API: API Gateway, Lambda, Dynamodb with optional X Ray
* Serverless API Gateway with RDS Proxy: The same as above but with RDS Proxy instead of Dynamodb
* Serverless ETL: SQS -&gt; Lambda -&gt; Dynamodb with Reddis for Deduplication &amp; a dead letter queue
* Serverless order processing with a manual verification step: Step function, Lambdas, Dynamodb  

## [7][Best way to connect Marketo to S3](https://www.reddit.com/r/aws/comments/juyrg3/best_way_to_connect_marketo_to_s3/)
- url: https://www.reddit.com/r/aws/comments/juyrg3/best_way_to_connect_marketo_to_s3/
---
I am trying to transfer all the "Leads " and "Activities" data from Marketo to AWS S3. (I haven't chosen a database yet, would like some advice on that too if available. There is &lt;25GB of data and its purported use is to query and build dashboards.)  
I looked into AppFlow which has a connector for Marketo but that does not allow me to transfer "Activities" and only allows "Leads". So

1) Is there any other ready made connector option for me within or outside of AWS?

2) I have looked into using AWS Lambda function for this purpose. Will using them work or should I go for something like Beanstalk?   
3) How else can I build my own connector and run it on the cloud? Hopefully using something like Python. I would also need to then transfer the data I pull (potentially using the Lambda function) into S3. Are Lambda functions viable for that operation too? 

Please provide any possible and I am new to AWS. I have looked around but I found it hard to get enough specific information on this.
## [8][move large zip files in S3 to RDS](https://www.reddit.com/r/aws/comments/juw7zi/move_large_zip_files_in_s3_to_rds/)
- url: https://www.reddit.com/r/aws/comments/juw7zi/move_large_zip_files_in_s3_to_rds/
---
I have two large(100GB each) zipped csv files in a bucket and I need to load them in our postgres DB instance. Previously I got sliced csv files, so I extracted them on EC2 and used psycopg2 python package and bytesIO to load those csv files, into postgres. Now I only have these large files and hereon I'll be receiving large files only instead of chunks. Does anyone know of a better way to load this data into my DB?

&amp;#x200B;

I've tried using the same approach as before but its very slow and after a day of running the python script, I could see the data present there but not accessible. I haven't looked into this problem yet but I'm wondering if there are better ways to do this. Thanks.
## [9][Does AWS publish an explanation of potential of increase prices if you chose to stay on previous/old Instance Type EC2?](https://www.reddit.com/r/aws/comments/juyrmw/does_aws_publish_an_explanation_of_potential_of/)
- url: https://www.reddit.com/r/aws/comments/juyrmw/does_aws_publish_an_explanation_of_potential_of/
---
As in, is it implied if you look at the pricing, that there is an increase in price?  I was just trying to find something official or even psuedo-official like @QuinnyPig that explains that retired instance types *will* (or may) go up in price as they are consuming valuable resources which would be better used by something newer/better.
Thanks all!
## [10][How can I remove all records from a dynamodb table at the same time everyday?](https://www.reddit.com/r/aws/comments/juu789/how_can_i_remove_all_records_from_a_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/juu789/how_can_i_remove_all_records_from_a_dynamodb/
---
Is there some kind of way I can set a trigger on AWS side rather than in my python code to auto delete all records (to be clear NOT delete the table, just delete all the records from the table) at the same time every single day?
