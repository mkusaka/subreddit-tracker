# aws
## [1][Amazon Lightsail now supports monitoring, notifications and alarming!](https://www.reddit.com/r/aws/comments/fakfj9/amazon_lightsail_now_supports_monitoring/)
- url: https://www.reddit.com/r/aws/comments/fakfj9/amazon_lightsail_now_supports_monitoring/
---
Hello Lightsailors,

Today we announced support for monitoring, notification and alarming for various metrics of your Lightsail instances, databases and load balancers. Choose to be alerted with emails or SMS! Comes included in the bundles - no extra charge :) 

Do try it and let us know any feedback you may have! Thank you.

Announcement below:

[https://aws.amazon.com/about-aws/whats-new/2020/02/amazon-lightsail-now-supports-resource-monitoring-alarming-and-notifications/](https://aws.amazon.com/about-aws/whats-new/2020/02/amazon-lightsail-now-supports-resource-monitoring-alarming-and-notifications/)

There's also a blog to walk you through this-

[https://aws.amazon.com/blogs/compute/configuring-and-using-monitoring-and-notifications-in-amazon-lightsail/](https://aws.amazon.com/blogs/compute/configuring-and-using-monitoring-and-notifications-in-amazon-lightsail/)

PS: If you are wondering what Lightsail is: [https://aws.amazon.com/lightsail/](https://aws.amazon.com/lightsail/)  and our easy to understand pricing plans - [https://aws.amazon.com/lightsail/pricing/](https://aws.amazon.com/lightsail/pricing/)
## [2][CodePipeline + CloudFormation + Lambda](https://www.reddit.com/r/aws/comments/faumli/codepipeline_cloudformation_lambda/)
- url: https://www.reddit.com/r/aws/comments/faumli/codepipeline_cloudformation_lambda/
---
Hey all,

I created this post on medium to help the ones who like me had difficulties to understand how to implement a DevOps pipeline in AWS using AWS developer tools (CodeCommit, CodeBuild, CodeDeploy and CodePipeline). And from understanding the DevOps pipeline create an real pipeline using cloudformation and deploying the code to Lambda (Python and Java).

https://link.medium.com/blysiuvJr4

I shared in the post files that I created that you can use to start an sample project in your account.
## [3][10 Tips to Optimize Your DynamoDB Costs](https://www.reddit.com/r/aws/comments/fai0fi/10_tips_to_optimize_your_dynamodb_costs/)
- url: https://dynobase.dev/dynamodb-cost-optimization/
---

## [4][Do I lose aws credits when changing the root email address of my account](https://www.reddit.com/r/aws/comments/faum6k/do_i_lose_aws_credits_when_changing_the_root/)
- url: https://www.reddit.com/r/aws/comments/faum6k/do_i_lose_aws_credits_when_changing_the_root/
---
Heya! Sorry this might be a bit of a too obvious to answer kinda question, but in case anyone knows: Does an account lose any aws credits when switching the root email address and/or billing information of said account? (not switching accounts just the email address and/or billing information- the accountId remains the same)  


I would assume that those credits are connected to the accountId rather than an account's email address but someone suggested differently, so just wanted to make sure.  


Thanks a lot in advance!
## [5][AWS Transit Gateway Examined – Part II | Caylent](https://www.reddit.com/r/aws/comments/fau9ht/aws_transit_gateway_examined_part_ii_caylent/)
- url: https://caylent.com/aws-transit-gateway-examined-part-ii
---

## [6][How to use API gateway to handle thousands of users each with their own throttling?](https://www.reddit.com/r/aws/comments/fattfu/how_to_use_api_gateway_to_handle_thousands_of/)
- url: https://www.reddit.com/r/aws/comments/fattfu/how_to_use_api_gateway_to_handle_thousands_of/
---
I am creating a service that lets users sign up, get an api key, and then can ping my api a maximum amount of 200 times per second.

My plan is to use a lambda for the logic of the api, and api gateway to handle the invokations.

I want to hard enforce the 200 / second on each user to avoid anyone abusing my lambdas and preventing access for other users.

I am expecting to have roughly 3000 users initially, with this number increasing over time.

I am looking into a solution as follows:

User signs up -&gt; a new api gateway is created that would be &lt;domain&gt;/&lt;username&gt;/endpoint with a limit of 200 requests per second, api key is generated from api gateway and given to the user.

I'm seeing on the limits doc a limit of 500 "API keys per account per region" - But I imagine I would be needing 1 key for each user I have, so I would need to get this increased?

I'm also seeing a limit of 300 "Usage plans per account per region" - Again I imagine I would need one usage plan per user if I'm wanting to lower the max requests to 200 / second?

Do people think this is the best solution or is there a different way to tackle this?
## [7][Is it possible to register an API gateway in a Private Hosted Zone?](https://www.reddit.com/r/aws/comments/fat7f7/is_it_possible_to_register_an_api_gateway_in_a/)
- url: https://www.reddit.com/r/aws/comments/fat7f7/is_it_possible_to_register_an_api_gateway_in_a/
---
I have a private hosted zone in route53, I'd like to register my Lambda API gateways in there so other applications can route to them via some consistent, sensible DNS entry within the VPC. Is this possible or does the zone need to be public?We're using SAM to create our Lambdas and API gateways, but the section to assign A-records within same requires a Certificate ARN which our private Route53 doesnt have.  


Edit:   
I'd create a public one, but we use cloudflare to manage all our certs and I'd rather not make the double hop
## [8][CodePipeline : Get commit-name and message which I can pass to Lambda in Environment Vars](https://www.reddit.com/r/aws/comments/fas1b3/codepipeline_get_commitname_and_message_which_i/)
- url: https://www.reddit.com/r/aws/comments/fas1b3/codepipeline_get_commitname_and_message_which_i/
---
Hello friends,

&amp;#x200B;

I am working on CodePipeline, without the CodeBuild phase. I am using codeDeploy to deploy applications on our server. Before starting the deployment and after finishing it, I am sending messages to Slack.   
The messages are not that useful, as they don't contain the commit name or the message. Any idea how I can access in CodePipeline the commit-name and message? Right now, I can access Environment variables from CodePipeline as follows :   


urlMessage = event\['CodePipeline.job'\]\['data'\]\['actionConfiguration'\]\['configuration'\]\['UserParameters'\]

&amp;#x200B;

But these are just custom params. I need from Github. Thank you. :-)
## [9][How do I search for an Intel 3Ghz machine on https://calculator.aws ?](https://www.reddit.com/r/aws/comments/fammil/how_do_i_search_for_an_intel_3ghz_machine_on/)
- url: https://www.reddit.com/r/aws/comments/fammil/how_do_i_search_for_an_intel_3ghz_machine_on/
---
Whilst searching for 6 core 16 GB RAM machine in https://calculator.aws/ it recommends me a a1.2xlarge, which I think is not X86 compatible. Whoops!!
Also the customer expects at least 3GHz clock speed. I looked at https://aws.amazon.com/ec2/faqs/ and does that mean I need to choose a Xeon processor? Or how can I better track the clock rate of EC2 instance types?
## [10][Error when running the CloudWatch Canaries to test website up/down](https://www.reddit.com/r/aws/comments/farjz1/error_when_running_the_cloudwatch_canaries_to/)
- url: https://www.reddit.com/r/aws/comments/farjz1/error_when_running_the_cloudwatch_canaries_to/
---
I got a canary exception when setting up the canaries for testing URL loading or not. I don't see any data regarding the loading success percentage on the cloudwatch metric. Anyone know why?
