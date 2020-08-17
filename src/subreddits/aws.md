# aws
## [1][Week of Aug 17th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/
---
Share your best AWS tips with the community!
## [2][Amazon MSK (Kafka) vs EventBridge](https://www.reddit.com/r/aws/comments/ibateh/amazon_msk_kafka_vs_eventbridge/)
- url: https://www.reddit.com/r/aws/comments/ibateh/amazon_msk_kafka_vs_eventbridge/
---
I need to create a message broker in my app (pub/sub model).

I haven't found any comparison between MSK and EventBridge - description of both claims it's a good choice.

Nevertheless, as I know life, there are some major differences. Can someone say which of this system is better? Or what are the major flaw of each of them? I have no experience in either of them.
## [3][Creating a new Dev env in the cloud](https://www.reddit.com/r/aws/comments/ibaplk/creating_a_new_dev_env_in_the_cloud/)
- url: https://www.reddit.com/r/aws/comments/ibaplk/creating_a_new_dev_env_in_the_cloud/
---
I just need a single server dev environment, but I am unsure about cloud security. This might end up needing two servers networked together. 

If I spin up an EC2 server and control who has access, it that considered secure enough? Should I create a proper VPC with a private subnet?

The code isn’t super sensitive, but I am not sure what open people do.

Edit: I haven’t used AWS before as our company is on-premise.
## [4][Did anyone get their credits from attending the EMEA Summit online??](https://www.reddit.com/r/aws/comments/ibafn5/did_anyone_get_their_credits_from_attending_the/)
- url: https://www.reddit.com/r/aws/comments/ibafn5/did_anyone_get_their_credits_from_attending_the/
---
I just remembered I should have had some free credits from AWS for attending the online Summit. Did anyone else get theirs yet??

[https://aws.amazon.com/events/summits/online/emea/faqs/](https://aws.amazon.com/events/summits/online/emea/faqs/)

Says 31st July. 

I have my certificate of attendance, which I feel I should frame and hang on the wall.... /s
## [5][Quick hands-on exercise to understand Amazon Lambda Destinations](https://www.reddit.com/r/aws/comments/iba91t/quick_handson_exercise_to_understand_amazon/)
- url: http://aws-dojo.com/excercises/excercise9
---

## [6][Controlling the storage volume of a bucket in S3](https://www.reddit.com/r/aws/comments/ibcutw/controlling_the_storage_volume_of_a_bucket_in_s3/)
- url: https://www.reddit.com/r/aws/comments/ibcutw/controlling_the_storage_volume_of_a_bucket_in_s3/
---
Hi r/aws,

I'm looking around for options on controlling the storage volume of a given S3 bucket.Take for example we are contracted to provide no more than 100GB of storage to a specific client.

I would like to control the capacity of the bucket.I'm finding CloudWatch metrics to be a little unreliable (not that I've really spent more than 10 minutes with it tonight).

Does anyone have any guidance on how I might enforce a storage quota?

Ideally - twice a day I'd like to get the current storage volume and if the size exceeds the value defined I would restrict the user or apply a new policy/acl on the bucket to restrict upload until the volume is reduced.
## [7][Is there a way to leverage IAM authentication in non-AWS services?](https://www.reddit.com/r/aws/comments/ib7p4h/is_there_a_way_to_leverage_iam_authentication_in/)
- url: https://www.reddit.com/r/aws/comments/ib7p4h/is_there_a_way_to_leverage_iam_authentication_in/
---
Here is the use-case. Let's say I am building a new backend API, I would like my clients to obtain an auth token for their IAM roles. Then the API will validate the token from request header and extrat iam role ARN from it for authentication purpose.

Is this a pattern that can be implemented using AWS SDKs?
## [8][API Gateway to mirror traffic](https://www.reddit.com/r/aws/comments/ib8xbt/api_gateway_to_mirror_traffic/)
- url: https://www.reddit.com/r/aws/comments/ib8xbt/api_gateway_to_mirror_traffic/
---
I have an OLTP system using lambdas and the databases. It's accessed through the API Gateway. I'd like to mirror the traffic that goes to the OLTP to an OLAP system. Is this possible with the API Gateway?  


I'm quite limited in the number of different services I can use on AWS due to restrictions by the sys admins, hence my question. The alternative is to have each lambda in the OLTP system shoot off a POST to the OLAP system at the start of each service, but I'm trying to avoid that.
## [9][Is there a way to create a CloudWatch Event trigger to fire an event exactly at the first and second weekdays of a month?](https://www.reddit.com/r/aws/comments/ibczdh/is_there_a_way_to_create_a_cloudwatch_event/)
- url: https://www.reddit.com/r/aws/comments/ibczdh/is_there_a_way_to_create_a_cloudwatch_event/
---
I have two events that I want to execute sequentially with an interval of one weekday between them (so one could be in Friday and the other in Monday).

I've come up with those cron expressions: "0 15 1W \* ? \*" and "0 15 2W \* ? \*". The problem is that in the second one, with "2W", it gets executed at 1st January 2021 (Friday), the same day as the one that has "1W", because per the definition the W wildcard is the 'closest day'. That’s straight just an horrible idea, I can't think of a single scenario that this is any useful. One could argue, "oh no, it can be useful because if you put 31W, then in February it's gonna execute on the closest weekday, 26th, 27th or 28th if it's a leap year"—wrong, "0 15 31W \* ? \*" is called on 29th January 2021 and next time on 31th March 2021.

"3W" and "4W" also gives 4th January 2021 for both schedules... this idea of "closest" in this scenario is just completely no sense.
## [10][Dark Mode Please!! My eyes are hurting :(](https://www.reddit.com/r/aws/comments/iaq34a/dark_mode_please_my_eyes_are_hurting/)
- url: https://www.reddit.com/r/aws/comments/iaq34a/dark_mode_please_my_eyes_are_hurting/
---
Hey AWS Team, My Dark Reader extension for some reason can't transform the aws console page into dark mode. I understand that might be due to some security reasons.   
But I request you people to please add a dark theme on AWS. My eyes really hurt a lot.

If there is an existing solution to turn the console into dark theme, I would like to know it.
## [11][Should I use DynamoDB for this use-case?](https://www.reddit.com/r/aws/comments/ib95mx/should_i_use_dynamodb_for_this_usecase/)
- url: https://www.reddit.com/r/aws/comments/ib95mx/should_i_use_dynamodb_for_this_usecase/
---
a list of events with attributes:

* twitterLink
* city
* state
* date
* time
* description

I want to be able to filter on any of those - most importantly the city and state.

Some events won't have any of those attributes filled in yet. They'll be added by real people looking at the twitterLink and extracting meaningful information.

I was planning on using Lambda to serve events to the frontend.
