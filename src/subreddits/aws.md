# aws
## [1][Dear AWS, Please stop spamming me](https://www.reddit.com/r/aws/comments/iwzrqz/dear_aws_please_stop_spamming_me/)
- url: https://www.reddit.com/r/aws/comments/iwzrqz/dear_aws_please_stop_spamming_me/
---
Dear AWS, 

Like many, I'm trying to follow 'best practices' - using AWS Control Tower and Multi-Account configuration.

However you're making this rather painful - every account created results in me getting yet another subscription to your marketing emails. 

Today I got a dozen more "Welcome to your Getting Started series" emails.   
As if that wasn't painful enough, un-subscribing is made unnecessarily difficult - I need to copy the email address over and tell you, once again, that I really don't want more marketing material. 

Given the amount of tracking in all those links, the least you could do would be to provide a one-click unsubscribe.  

Better still - how about for every account created from AWS Orgs / Control Tower you opt them out of marketing emails. 

Sincerely Frustrated, 

L. Extension.
## [2][I passed the AWS solutions architect exam about an year ago, wanted to share my journey and my notes. It has already helped a lot of people, hope it also helps some of you.](https://www.reddit.com/r/aws/comments/iwfi95/i_passed_the_aws_solutions_architect_exam_about/)
- url: https://github.com/SkullTech/aws-solutions-architect-associate-notes
---

## [3][Tool for managing parameters in AWS SSM Parameter Store](https://www.reddit.com/r/aws/comments/iwylqu/tool_for_managing_parameters_in_aws_ssm_parameter/)
- url: https://www.reddit.com/r/aws/comments/iwylqu/tool_for_managing_parameters_in_aws_ssm_parameter/
---
[https://github.com/kevinglasson/goss](https://github.com/kevinglasson/goss)

I developed a tool to help managing config for various environments in AWS SSM PS. I though some people here might be interested.

I am also interested to know how everyone is managing their own config? We use SSM PS to manage local dev config as well as prod / stage (under different paths of course)
## [4][Can you receive emails from an external domain on an Amazon SES domain?](https://www.reddit.com/r/aws/comments/iwz5g1/can_you_receive_emails_from_an_external_domain_on/)
- url: https://www.reddit.com/r/aws/comments/iwz5g1/can_you_receive_emails_from_an_external_domain_on/
---
Hi,

Haven't really worked much with SES in the past but I have some integrations I need to implement. I have used SES before to send automated emails but now I'm trying to set it up so it can receive emails.

I am able to configure it so it receives emails from inside the domain, but I can't receive from outside the domain.

Is this due to how SES is setup or is there something I'm missing, can you only receive emails on SES from inside the domain?
## [5][Route53 Geolocation Routing Policy returning values in random order](https://www.reddit.com/r/aws/comments/iwxmor/route53_geolocation_routing_policy_returning/)
- url: https://www.reddit.com/r/aws/comments/iwxmor/route53_geolocation_routing_policy_returning/
---
I set up 4 records with geolocation routing policy on them: one record for my instance in London (I set Europe region), one for Ohio (North America region), one for Sydney (Oceania region), and one for default. On the default record, I put the London instance IP address on it. So there would be two records with the London Instance IP address.

When I refresh every 5-6 minutes, I'll get my Ohio instance, and then my Sydney instance, and then London or something. Just like a simple routing policy.

I tried redoing the records just in case and it's the same.

Is this the correct setup?

(My instances are stopped, so the IP's don't work atm)

https://preview.redd.it/z026ugs77ho51.png?width=850&amp;format=png&amp;auto=webp&amp;s=3644a5fda630c599ba94fc5d797c7232c7134f57
## [6][Week of Sept 21st - What are you building this week in AWS?](https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/
---
Share what you are working on
## [7][How do i validate a Cloudfront signed url in Lambda@Edge?](https://www.reddit.com/r/aws/comments/iwzgj7/how_do_i_validate_a_cloudfront_signed_url_in/)
- url: https://www.reddit.com/r/aws/comments/iwzgj7/how_do_i_validate_a_cloudfront_signed_url_in/
---
Pretty much my scenario is as follows, right now i have a Cloudfront distribution where access is controlled by a Lambda@Edge function. Basically it will validate a token in the Authorization header. This has worked fine. 

Recently i ran into a use case where i need to use cloudfront signed urls for certain distributions protected by the above. 

The problem is the Lambda@Edge function prevents signed urls from working:

* If i remove the lambda@Edge function the signed urls work fine.
* If i add logic to allow any requests with Expires/Key-Pair-Id/Signature params in the Lambda and act as a simple pass through they will return 200. HOWEVER it doesn't seem like Cloudfront is actually validating the signed url. 
* I cannot find anything in the AWS-SDK Cloudfront stuff that indicates how to validate these urls.
* I really don't want to create duplicate distributions for this.

I have spent the last few hours digging for anything and i cannot find anything. Does anyone have an idea?
## [8][HELP: Is uploading to S3 with the Amplify GraphQL client possible?](https://www.reddit.com/r/aws/comments/iwwjwl/help_is_uploading_to_s3_with_the_amplify_graphql/)
- url: https://www.reddit.com/r/aws/comments/iwwjwl/help_is_uploading_to_s3_with_the_amplify_graphql/
---
I want to include S3Objects in my schemas like this example from the documentation:  [https://docs.amplify.aws/lib/graphqlapi/advanced-workflows/q/platform/js](https://docs.amplify.aws/lib/graphqlapi/advanced-workflows/q/platform/js) 

According to this, the feature isn't supported yet  [https://github.com/aws-amplify/amplify-js/issues/2706](https://github.com/aws-amplify/amplify-js/issues/2706) but there appears to be someone at the bottom who was able to do it. Anyone have any luck with this? Am I missing something?
## [9][Multi AWS Account Manager](https://www.reddit.com/r/aws/comments/iwu57p/multi_aws_account_manager/)
- url: https://www.reddit.com/r/aws/comments/iwu57p/multi_aws_account_manager/
---
I have just created about 30 AWS Organizations for a project I am working on.   
is there a application I can use to manage all of the 30 accounts without logging into each one?

Basically I want to view how many EC2 instances are running at any given time and turn them on or off on demand without having to manually switch to the organization account. 

Any recommendations welcome.
## [10][AWS product suggestion](https://www.reddit.com/r/aws/comments/iwxs1q/aws_product_suggestion/)
- url: https://www.reddit.com/r/aws/comments/iwxs1q/aws_product_suggestion/
---
Hello

I might have a new client who wants to work in cloud instead of onpremise. They have a Windows app which runs on a Windows server (and old accounting app) and everyone is connecting to it via a shortcut in their own computers. Something like shops used years ago. They will be uploading files to this server too. Not like as pure file server but only to specific folders which this app will use. They also want to use their local printers too.

How can I achieve that? App DB size is around 5 GB, and user count is around 80ish
