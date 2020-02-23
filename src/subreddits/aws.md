# aws
## [1][Cognito + Lambda + ?? = userID](https://www.reddit.com/r/aws/comments/f861hq/cognito_lambda_userid/)
- url: https://www.reddit.com/r/aws/comments/f861hq/cognito_lambda_userid/
---
I want to have a unique userID as a Cognito custom attribute that will also live in RDS.

So I need a Lambda function that picks an ID, saves it as a custom attribute to Cognito, and saves it in RDS. Pretty straight forward. (Hopefully.)

But then I stick that function to the .... Sign Up button? But what happens if the user enters a password that's too short, or whatever? Then the endpoint will fire multiple times, right? And I'll just have a bunch of IDs with no activity.

Am I on the right track? What am I missing?

Also, the Cognito UI runs smoothly out of the box. So besides this, I don't need any new Lambda/API Gateway functions, right?

Thanks.
## [2][Scaling up DynamoDB](https://www.reddit.com/r/aws/comments/f86dir/scaling_up_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/f86dir/scaling_up_dynamodb/
---
My company started using DynamoDB when we were pretty small. We used it more as a dump, keeping useful information in Postgres, and keeping the rest in DynamoDB. As the business grew, we kept changing the Postgres meeting the immediate requirements.
Now, I am sitting on motherlode of data, but can’t monetise it as querying in DynamoDB costs a lot. We are in fact planning on moving away from it. 
What are your experiences with scaling DynamoDB, pitfalls, and suggestions?
## [3][Need help to move from MongoDB+mongoose to DynamoDB.](https://www.reddit.com/r/aws/comments/f88gv4/need_help_to_move_from_mongodbmongoose_to_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/f88gv4/need_help_to_move_from_mongodbmongoose_to_dynamodb/
---
Hi everyone,

lately, I've started developing Alexa voice apps. My current stack is Jovo (a nodejs framework for voice app), nodejs, typescript, and MongoDB with mongoose.

All my backend logic at the moment is hosted on a digital ocean server.

I would like to fully switch to AWS + DynamoDB but to do that I should replace my current MongoDB+mongoose logic.

Could you suggest me some library/guide that could allow me to do this process? I would like to find something like mongoose to using on top of row DynamoDB and that would speed up my development and help me with modelling data.

If it support TypeScript it would be a huge plus.
## [4]["MySQL server has gone away" errors during scale down events in Aurora DB Serverless](https://www.reddit.com/r/aws/comments/f85tcw/mysql_server_has_gone_away_errors_during_scale/)
- url: https://www.reddit.com/r/aws/comments/f85tcw/mysql_server_has_gone_away_errors_during_scale/
---
I've just migrated an ElasticBeanstalk app last week. Lately, I've noticed some errors and when I investigated the logs, the timestamps correspond to scale down events in Aurora DB. The web app is written in PHP. It's still weekend and I don't have the SSH key to read the code, so I want to gather ideas for Monday.

Question: what are the usual reasons for PHP apps to get MySQL errors during scale downs? Does the PHP app have a persistent connection to Aurora? To my understanding, Aurora waits until there are no connections before it scales down; otherwise, it forces a disconnection, and hence a "MySQL server has gone away" error. 

UPDATE: I forgot that I can just download the code without the need for SSH keys. The connection is just a standard call to mysqli().
## [5][What are you folks building using AWS Lambda?](https://www.reddit.com/r/aws/comments/f7u598/what_are_you_folks_building_using_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/f7u598/what_are_you_folks_building_using_aws_lambda/
---
I see the use of AWS Lambda but I'm not really sure what the right use-cases are?  


If there's any open source Lambda based projects someone's got, I'd love to take a look!
## [6][Announce list using SNS, by creating a form to subscribe by email](https://www.reddit.com/r/aws/comments/f814in/announce_list_using_sns_by_creating_a_form_to/)
- url: https://github.com/kaihendry/sns-subscribe
---

## [7][Getting the right filter for ec2 Instance](https://www.reddit.com/r/aws/comments/f8568x/getting_the_right_filter_for_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/f8568x/getting_the_right_filter_for_ec2_instance/
---
Hey All, 

I'm trying to get the right filter I should use to get the 'AvailabilityZone', or Region of the Instance. Can you all help me with this?

    import boto3
    ec2 = boto3.resource('ec2')
    for instance in ec2.instances.all():
        print (instance.id , instance.state)

\^ That is supposed to give me my instance id, and state. 

Is there a way to get the region of the instance?

Thanks in advance!
## [8][Marketplace side hustle?](https://www.reddit.com/r/aws/comments/f7whn3/marketplace_side_hustle/)
- url: https://www.reddit.com/r/aws/comments/f7whn3/marketplace_side_hustle/
---
Any freelancers out there making any income with things in the marketplace?
## [9][S3 as a personal remote FS](https://www.reddit.com/r/aws/comments/f7yqd6/s3_as_a_personal_remote_fs/)
- url: https://www.reddit.com/r/aws/comments/f7yqd6/s3_as_a_personal_remote_fs/
---
Is it worth it to create and use a S3 account as a personal backup File System for long term storage like STL files, documents , etc?
## [10][S3 Service requests from US East (N. Virginia) region even though there are no buckets or no one there to access it](https://www.reddit.com/r/aws/comments/f84lni/s3_service_requests_from_us_east_n_virginia/)
- url: https://www.reddit.com/r/aws/comments/f84lni/s3_service_requests_from_us_east_n_virginia/
---
I have few private buckets in ap-southeast-1 region. These are new buckets, so there isn't much traffic to it either. When looking through the bill for last month, I notice that there are a number of Tier 1 and Tier 2 requests coming from US East region, which should not be. It does not incur any charge as the number is in only in hundreds, but I am curious what is the reason for it. Is it something that is done by AWS themselves ? 

I have few other accounts, and I checked them too, also to find that they all receive services requests in US East region. Does anyone know the reason for these requests ? (Bucket logging is not enabled, so cannot check them)
