# aws
## [1][Hands-on exercise with Amazon AppFlow to transfer data from S3 to S3](https://www.reddit.com/r/aws/comments/ianuhe/handson_exercise_with_amazon_appflow_to_transfer/)
- url: http://aws-dojo.com/excercises/excercise4
---

## [2][Dark Mode Please!! My eyes are hurting :(](https://www.reddit.com/r/aws/comments/iaq34a/dark_mode_please_my_eyes_are_hurting/)
- url: https://www.reddit.com/r/aws/comments/iaq34a/dark_mode_please_my_eyes_are_hurting/
---
Hey AWS Team, My Dark Reader extension for some reason can't transform the aws console page into dark mode. I understand that might be due to some security reasons.   
But I request you people to please add a dark theme on AWS. My eyes really hurt a lot.

If there is an existing solution to turn the console into dark theme, I would like to know it.
## [3][Amazon S3 Path Deprecation Plan (Deprecate date: Sep 30, 2020)](https://www.reddit.com/r/aws/comments/ia5wy7/amazon_s3_path_deprecation_plan_deprecate_date/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [4][Can a Network Load Balancer be used to balance RTMP connections?](https://www.reddit.com/r/aws/comments/iamf3j/can_a_network_load_balancer_be_used_to_balance/)
- url: https://www.reddit.com/r/aws/comments/iamf3j/can_a_network_load_balancer_be_used_to_balance/
---
I've got one set up, and I can't seem to connect to it. When I check the end of the route for traffic, I'm getting a constant stream of about 2 connect/disconnects per second and I've been unable to get any further information.  Am I misunderstanding how RTMP works?  I thought since it was standard tcp I could pass it through a NLB.  

The NLB is on a VPC with an ECS cluster and I've confirmed the security groups allow traffic from NLB to ECS instances, but I can't seem to get an RTMP stream through.  

&amp;#x200B;

If I can't do it, does anyone have any ideas?  And if it *should* work, does anyone have any theories as to what I'm doing wrong?
## [5][How to have RDS replicate data to EMR Hive / EMR Hive react to RDS changes?](https://www.reddit.com/r/aws/comments/iaqevu/how_to_have_rds_replicate_data_to_emr_hive_emr/)
- url: https://www.reddit.com/r/aws/comments/iaqevu/how_to_have_rds_replicate_data_to_emr_hive_emr/
---
Hi all. I am fairly new to the AWS universe, and we have started using some of its services at work approximately a year ago. Now we find ourselves in a situation where we need EMR Hadoop's HDFS and Hive's ability to query properly structured tables via SQL. We plan to install R, RHIPE and RHIVE to define UDFs to feed to Hive in order to produce statistical data. We are looking for a way to have RDS and EMR (directly Hive, if possible) sync data for (near) real-time calculations. I have read AWS own documentation for both services but have not found answers. We also checked out some APN software, but we were looking to use only AWS services, even an intermediary service for replication / syncing if needed. Any suggestion is appreciated. Thanks!
## [6][Help with running a very basic web application](https://www.reddit.com/r/aws/comments/iamc32/help_with_running_a_very_basic_web_application/)
- url: https://www.reddit.com/r/aws/comments/iamc32/help_with_running_a_very_basic_web_application/
---
For a uni assignment, I have to create a web application with a group. The assignment is more about the project management, so we're just creating a really basic web application. We're using AWS, but I'm really confused by the services and what they offer even after doing a bunch of research. I've never really used AWS before. 

We honestly need very little functionality, just basically retrieving data from our database, displaying that to the user on the front end, and then manipulating the database based on certain user actions.

Right now we hjave S3, lambda, cognito, and dynamoDB setup with a basic landing page that lets the user register/login but I don't think this can do what we want for the actual application? I understand S3 is for static webpages and we essentially want a dynamic webpage, but doesn't lambda essentially act as the serverside code?

Would we just be better off using Elastic Beanstalk?
## [7][Accessing S3 buckets objects with Google Authentication (or any other OAuth provider)](https://www.reddit.com/r/aws/comments/iapt7p/accessing_s3_buckets_objects_with_google/)
- url: https://www.reddit.com/r/aws/comments/iapt7p/accessing_s3_buckets_objects_with_google/
---
#### Requirement
Employees need access (read-only) to specific S3 buckets, using an OAuth provider, such as Google.

#### Flow
1. A static website served through S3 bucket (publicly available) - https://s3.company.com
1. Employees (end-users) must go through a login process (Google Sign-In) to view the static website
1. The end-user sees a list of buckets that he can view, the end-user clicks on a bucket and sees the list of objects in that bucket
1. The end-user clicks on an object to download it (same experience as in AWS Console &gt; S3)

#### Pros
1. Avoid setting up ap SSO with AWS, no need to create a user per employee
1. Avoid using VPN, using OAuth provider instead
1. Providing a quick access to S3 objects, with great user experience - sign-in and download, no need to go through the AWS Console

#### Cons (discussion)
This is where you come in - I need your help to understand if there's something fundamentally wrong with this process


Refs- 
1. A great example of how the frontend should like - https://github.com/awslabs/aws-js-s3-explorer
1. https://aws.amazon.com/blogs/security/how-to-set-up-federated-single-sign-on-to-aws-using-google-apps/
## [8][WordPress on Elastic Beanstalk + RDS](https://www.reddit.com/r/aws/comments/iajx49/wordpress_on_elastic_beanstalk_rds/)
- url: https://www.reddit.com/r/aws/comments/iajx49/wordpress_on_elastic_beanstalk_rds/
---
Hi All,

I'm looking for some advice on the best way to migrate a WordPress site off of a managed provider (some random vendor my company uses) onto AWS using Elastic Beanstalk + RDS.

I understand the fundamentals of doing this, connecting S3 for storage, using RDS as the database, and having Elastic Beanstalk be completely stateless. But I'm curious as to the nuances of keeping WordPress and all of the plugins up to date across the environment. My current thought is to do all of the updates locally and push them with GitHub + CodeDeploy but I've never done this exact thing before so I thought I'd ask.

Any feedback would be helpful and appreciated.
## [9][Global Accelerator vs CloudFront for public ALB](https://www.reddit.com/r/aws/comments/iahog4/global_accelerator_vs_cloudfront_for_public_alb/)
- url: https://www.reddit.com/r/aws/comments/iahog4/global_accelerator_vs_cloudfront_for_public_alb/
---
Hi all

Been wondering which service if any, would be useful, here is the relevant parts of my tech stack:

*  JavaScript SPA in a S3 bucket with CloudFront for all the wonders a CDN does for static content.
* A ECS cluster exposed by a public ALB, inside the cluster lives the business logic in the form of microservices, currently only in one region us-east1.

The front-end makes HTTP requests to the alb, typical stuff, all the data gets send as JSON.

The situation I'm facing is that even if using CloudFront helps for delivering the static content (especially when the user is far from the US) I still don't get much use of the aws backbone for networking, because the api calls are still going form the browser through the public internet to my alb (right?).

This is were I think Global Accelerator or putting CloudFront in front of my alb would help, if I'm getting it correctly the flow would go something like this:

* User enters example.com
* The requests gets routed to the nearest PoP and the content gets delivered from there
* User authenticates with Cognito, not much I can do here since the Cognito user pool and all the routing is taking care by aws
* User starts using example.com, and when needed the browser makes requests to api.example.com 
* Here ideally by using CloudFront or Global Accelerator the request would get routed to the nearest PoP and once there the request would go through the aws backbone all the way to my alb and back

That last point is what is important and where I'm having trouble with, what difference would make to use either service? Would CloudFront actually cache any content from the alb? If yes then I would have to have a TTL=0 as the content varies by user. 

What about the TLS handshake that occurs on each request to api.example.com are there any mechanisms to reuse the handshakes needed if i would use any of theses services?

I'm a gaining anything by this? Or just wasting money? Would like to hear what everyone has to say.
## [10][Does API Gateway / Lambda support Server-Sent Events? Or just WebSockets?](https://www.reddit.com/r/aws/comments/iaj5ea/does_api_gateway_lambda_support_serversent_events/)
- url: https://www.reddit.com/r/aws/comments/iaj5ea/does_api_gateway_lambda_support_serversent_events/
---
Does API Gateway / Lambda support Server-Sent Events? Or just WebSockets?

I’ve found one resource that said the lambdas would have to be kept alive for the duration of the connection which kind of defeats the purpose.
