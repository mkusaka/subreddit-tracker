# aws
## [1][Week of Sept 14th - What are your favorite container/serverless tips in AWS?](https://www.reddit.com/r/aws/comments/isls8o/week_of_sept_14th_what_are_your_favorite/)
- url: https://www.reddit.com/r/aws/comments/isls8o/week_of_sept_14th_what_are_your_favorite/
---
Share your container/serverless tips
## [2][Amazon CloudFront announces support for Brotli compression](https://www.reddit.com/r/aws/comments/ithz5w/amazon_cloudfront_announces_support_for_brotli/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/cloudfront-brotli-compression/
---

## [3][Fastest and most cost efficient way to copy over an S3 bucket from another AWS account](https://www.reddit.com/r/aws/comments/itvd0f/fastest_and_most_cost_efficient_way_to_copy_over/)
- url: https://www.reddit.com/r/aws/comments/itvd0f/fastest_and_most_cost_efficient_way_to_copy_over/
---
I have an S3 bucket that is 9TB and I want to copy it over to another AWS account.

What would  be the fastest and most cost efficient way to copy it? 

I know I can rsync them and also use S3 replication.

Rsync I think will take too long and I think be a bit pricey.

I have not played with S3 replication so I am not sure of its speed and cost.

Are there any other methods that I might not be aware of?

FYI - The source and destination buckets will be in the same region (but different accounts).
## [4][You can now share your Amazon CloudWatch dashboards outside the AWS console - by username/password, SSO and public URLs. Dark mode by default, can be embedded in other websites. With extremely long login/sessions timeouts perfect for TV dashboards.](https://www.reddit.com/r/aws/comments/ite4im/you_can_now_share_your_amazon_cloudwatch/)
- url: https://i.redd.it/nkrtmnavpcn51.png
---

## [5][Amazon Transcribe Now Supports Automatic Language Identification](https://www.reddit.com/r/aws/comments/iteso8/amazon_transcribe_now_supports_automatic_language/)
- url: https://aws.amazon.com/blogs/aws/amazon-transcribe-now-supports-automatic-language-identification/
---

## [6][Please, return back the shortcuts in the top panel... I like the new looks but I want the top panel back](https://www.reddit.com/r/aws/comments/itraqb/please_return_back_the_shortcuts_in_the_top_panel/)
- url: https://i.redd.it/5v316o9mtgn51.png
---

## [7][How to track RDS minor version upgrade in code?](https://www.reddit.com/r/aws/comments/ituisc/how_to_track_rds_minor_version_upgrade_in_code/)
- url: https://www.reddit.com/r/aws/comments/ituisc/how_to_track_rds_minor_version_upgrade_in_code/
---
I have an RDS, created using a basic Cloudformation template.  My question is, how do you keep the version defined in the template and the actual version in sync, if I enable automatic minor version upgrade?


Should I use some sort of lambda function to keep that value in an SSM parameter, and have the CF template refer to the param?    Or is it better to disable automatic upgrade and do it manually through update-stack?

Or is this simply a problem not worth solving?
## [8][Making a scheduled task communicate with other tasks in an ECS Fargate service](https://www.reddit.com/r/aws/comments/itszqh/making_a_scheduled_task_communicate_with_other/)
- url: https://www.reddit.com/r/aws/comments/itszqh/making_a_scheduled_task_communicate_with_other/
---
I have a web service built in Java running in ECS Fargate with 3 tasks that have 1 container each. All 3 of these containers cache a table from a DB in JVM heap inside the container and I need to refresh this cache periodically. Ideally, I would just move the cache outside the container to something like Elasticache and try to use that, but for reasons that are beyond my control I am limited to storing and updating the cache inside these containers for the time being. There is an http end point in the service that accepts requests to refresh this cache. So where I am at right now is trying to figure out how I can make a request to the refresh cache end point in the service periodically and have that request routed exactly once to all 3 of these containers. I have come across scheduled tasks in ECS, so in theory I can spin up that scheduled task on a specified time interval in the same service and have it make that refresh cache request to all the other tasks in the service, but I am unsure about how to make it communicate with the other tasks. I am also aware of the risk of one of the containers restarting or being replaced and having it's cache out of sync with the other containers, but I am willing to accept that risk for the time being.
## [9][Security September: Racing against CloudWatch Synthetics Canaries – One Cloud Please](https://www.reddit.com/r/aws/comments/itjt4j/security_september_racing_against_cloudwatch/)
- url: https://onecloudplease.com/blog/security-september-racing-against-cloudwatch-synthetics-canaries
---

## [10][Statistical simulations in the cloud?](https://www.reddit.com/r/aws/comments/itsdaq/statistical_simulations_in_the_cloud/)
- url: https://www.reddit.com/r/aws/comments/itsdaq/statistical_simulations_in_the_cloud/
---
Dear all,  
is it possible to run simulations in the AWS cloud using Stata for a person that is not an expert with cloud computing? I am looking for a machine with 16 to 64 cores and about 8GB of RAM. What are the approximate costs if this simulation runs about 12h? What type of machine can you recommend? Thanks for some info.
## [11][Are there any intentions of lowering the price of EBS volumes?](https://www.reddit.com/r/aws/comments/itkhk9/are_there_any_intentions_of_lowering_the_price_of/)
- url: https://www.reddit.com/r/aws/comments/itkhk9/are_there_any_intentions_of_lowering_the_price_of/
---
An honest question, not complaining or anything, just wondering if it's on the "roadmap".

[2014 vs 2020 pricing](https://preview.redd.it/xblvi08ween51.jpg?width=595&amp;format=pjpg&amp;auto=webp&amp;s=44d97859d0d8f393c076e95255b29a2748023cd2)

A good article talking regarding the pricing of SSD drives over here years - [https://www.minitool.com/news/ssd-prices-fall.html](https://www.minitool.com/news/ssd-prices-fall.html) (not mine, just a good ref).

**Sources**

2014-Jun [https://aws.amazon.com/blogs/aws/new-ssd-backed-elastic-block-storage/](https://aws.amazon.com/blogs/aws/new-ssd-backed-elastic-block-storage/)

2020-Sep (Updated constantly) [https://aws.amazon.com/ebs/pricing/](https://aws.amazon.com/ebs/pricing/)
