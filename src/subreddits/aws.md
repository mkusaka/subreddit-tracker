# aws
## [1][Beginning work that will involve AWS in the next week at work - what are some good resources for getting up to speed?](https://www.reddit.com/r/aws/comments/fgd3q4/beginning_work_that_will_involve_aws_in_the_next/)
- url: https://www.reddit.com/r/aws/comments/fgd3q4/beginning_work_that_will_involve_aws_in_the_next/
---
Hey, everyone. I will be joining a new team next week where I will be helping with migration from Azure to AWS. The most experience I have in AWS right now is setting up a Lightsail instance and playing around with it, so obviously what I know about AWS is very limited.

The team I have been on ended their iteration at the end of last week and it will take me about a week to get all my credentials set up(my company moves slowly) so I have about a week where I'm just sitting on my hands. I'd like to use this time to start getting more familiar with the AWS landscape and its' tools so I don't have such a high learning curve once I dive in next week.

Any assistance would be great. Thank you!
## [2][Do S3 buckets live on EC2's under the hood? Can't find this information anywhere.](https://www.reddit.com/r/aws/comments/fg2rg8/do_s3_buckets_live_on_ec2s_under_the_hood_cant/)
- url: https://www.reddit.com/r/aws/comments/fg2rg8/do_s3_buckets_live_on_ec2s_under_the_hood_cant/
---
Let's say you're using an S3 bucket to serve a static website and AWS creates an endpoint for you to view your bucket. Does clicking that endpoint (aka going to the website) trigger some kind of EC2 activity? Or maybe Lambda -&gt; EC2 activity? How are S3 files actually served?

Thanks everyone.
## [3][Is it possible to compile and execute a C program using AWS lambda?](https://www.reddit.com/r/aws/comments/fg7pkq/is_it_possible_to_compile_and_execute_a_c_program/)
- url: https://www.reddit.com/r/aws/comments/fg7pkq/is_it_possible_to_compile_and_execute_a_c_program/
---
I am working on a personal project of making an online IDE . And i have implemented the compilation and Execution using a EC2 instance. But, Recently came to know about Lambda and i am wondering if its possible to do the same (compile and execute the c program and return the output ). IF not what would be the better way to approach this.On EC2 i went for a nodejs application to get the code from user and then send the output as the response to the user.
## [4][AWS EC2 shell request failed on channel 0](https://www.reddit.com/r/aws/comments/fg9nmi/aws_ec2_shell_request_failed_on_channel_0/)
- url: https://www.reddit.com/r/aws/comments/fg9nmi/aws_ec2_shell_request_failed_on_channel_0/
---
Previously i can easily access my EC2 instance but now i get this message over and over again and i can't access it anymore.

`shell request failed on channel 0`
## [5][Benchmarking the AWS Graviton2 with KeyDB – M6g up to 65% faster](https://www.reddit.com/r/aws/comments/ffupap/benchmarking_the_aws_graviton2_with_keydb_m6g_up/)
- url: https://docs.keydb.dev/blog/2020/03/02/blog-post/
---

## [6][I have a simple WordPress.org blog — should I continue using regular webhosting services (using iPage currently) or switch to AWS? I'm looking for the cheapest option.](https://www.reddit.com/r/aws/comments/ffv4dy/i_have_a_simple_wordpressorg_blog_should_i/)
- url: https://www.reddit.com/r/aws/comments/ffv4dy/i_have_a_simple_wordpressorg_blog_should_i/
---
I'm paying 7.99 - 9.99USD monthly.

Also, what is the estimated price per month if I were to go with AWS? Expected visitors per month are no more than 5k. As for storage I can go with 100MB. Because I can upload media on Imgur and other services. So that's not a huge deal. I mostly want the cheapest option possible.

And I already have a domain on namesilo so don't count that. It will be linked to AWS if that's what you guys are recommending.

Thank you.

UPDATE: Thank you guys. Honestly overwhelmingly great responses. I can't reply to everyone but I appreciate all your inputs.

I have several options now:

1. Amazon's LightSail.
2. A static website that scrapes the content from an offline (on my PC) server. This one is particularly great because it can be hosted free of charge.
3. And DigitalOcean.

These are all great suggestions and I'll probably try them all and see which method/service provider suits me the most.
## [7][Massive siteground account to AWS?](https://www.reddit.com/r/aws/comments/fg1wsb/massive_siteground_account_to_aws/)
- url: https://www.reddit.com/r/aws/comments/fg1wsb/massive_siteground_account_to_aws/
---
Hello, I currently have a large Siteground account with 100+ wordpress sites. It works pretty well, but the cost is very high. 
Is AWS or lightsail worth looking into? 
My servers stats are:
4 CPU cores centOS
10 gb ram
130 gb ssd storage 
I need to upgrade the ram, so I want to explore other options first. 
Any thoughts?
## [8][Best practices for api logging with ALB and lambda](https://www.reddit.com/r/aws/comments/ffxq3t/best_practices_for_api_logging_with_alb_and_lambda/)
- url: https://www.reddit.com/r/aws/comments/ffxq3t/best_practices_for_api_logging_with_alb_and_lambda/
---
Hi everyone,
I have a lambda function (a fat one) behind an ALB for a public API. I'm looking for an efficient way of sending custom logs to my rds (postgres) instance. Ideally O would like to go for a serverless/cheap solution (so dynamodb is a no).
The problem is that for my use case api latency is a real issue, it's unacceptable for it to go over 60-70ms. 
I've tried sending them to SQS and Kinesis, but sometimes it takes over 100ms just to send the events (there could be multiple events per request).
Currently I'm lookong into sendong them to redis (using oub/sub) and having another lambda run every minute or so to get that data and send it to rds.

Does anyone have a better way of doing this? Oe do you think this idea makes sense?
## [9][Average Cloud Support Engineer Salary in Oregon, Portland?](https://www.reddit.com/r/aws/comments/fg8edk/average_cloud_support_engineer_salary_in_oregon/)
- url: https://www.reddit.com/r/aws/comments/fg8edk/average_cloud_support_engineer_salary_in_oregon/
---
I'm getting an offer for working as an Cloud Support Engineer for Amazon in Portland and I cannot find enough salaries online to justify a reasonable price.

Anyone here work in Amazon Cloud or knows someone and can think of a good salary range to negotiate for in the city of Portland?

Note: it isn't about experience and other factors \*on the topic of salary\*, They're simply looking if you have the skills for it or not. 20 years or 5 years, do not matter for this one. The initial 'entry' salary is claimed to be relatively the same from the recruiter's knowledge.

I have 0 experience with AWS Cloud. Only 2 years of Azure and a history of Help desk / Entry Sys Admin work for 6 years.

\- Or am I being mislead??

I'm from a low CoL state in the south so I have no clue what's a good number to make up for those strenuous gas prices and rent up in the west
## [10][Need to track downloads on S3](https://www.reddit.com/r/aws/comments/fg355n/need_to_track_downloads_on_s3/)
- url: https://www.reddit.com/r/aws/comments/fg355n/need_to_track_downloads_on_s3/
---
 My boss wants me to find a way to track client downloads. We host client data on S3 (our servers are on AWS). There's no native way to do this, it seems. I found an app called S3Stat. Anybody ever use it or have other (simple, extremely user friendly) suggestions?
