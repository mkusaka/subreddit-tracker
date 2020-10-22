# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][Presgined URL'S vs static hosting](https://www.reddit.com/r/aws/comments/jfxsny/presgined_urls_vs_static_hosting/)
- url: https://www.reddit.com/r/aws/comments/jfxsny/presgined_urls_vs_static_hosting/
---
I am developing an application with users and photos. 

Currently, my implementation is, that I have a one-to-many relationship between users and photos.

When I serve the images I use the AWS-SDK to generate the resigned URL's and it works great!

&amp;#x200B;

My issue now is that I need to load the user photos relatively fast, and when I retrieve users.

Right now I have the photos and users and separate endpoints (this is the client can load conditionally if the presigned URLs are not generated yet or vice versa if the user data is not fetched yet).

&amp;#x200B;

But the operation quickly becomes very expensive because it may have to generate 50 resigned URLs at once (10 users \* 5 images), and now I'm considering just hosting static images (or static private images). 

&amp;#x200B;

What is you guys' input on this? Should I stick with presigned URLs?

Furthermore is it necessary to have photos and users on two separate endpoints?
## [3][Public Preview – AWS Distro for OpenTelemetry](https://www.reddit.com/r/aws/comments/jfg2il/public_preview_aws_distro_for_opentelemetry/)
- url: https://aws.amazon.com/blogs/aws/public-preview-aws-distro-open-telemetry/
---

## [4][Lightsail Instance vanished - No technical support available](https://www.reddit.com/r/aws/comments/jfx5ps/lightsail_instance_vanished_no_technical_support/)
- url: https://www.reddit.com/r/aws/comments/jfx5ps/lightsail_instance_vanished_no_technical_support/
---
I am not sure Amazon will let me complain on their board so here it is:

&amp;#x200B;

 I tried building a wordpress website using AWS lightsail.  


This is the sequence of events.  


(1) At about 2 PM on Oct 21, while I was editing my wordpress site, it stopped getting updated  
(2) I login in to console, my AWS instance is gone, DNS is gone, static ip is gone, just nothing there and my site is not working either  
(3) I scramble around and try to get some technical support, they ask me to $29.99 else all I get is to talk to banking associates who don't understand static ip and namerservers.  
Those associate keep telling me that I never had created anything on AWS ever even though I have nameserver and static ip from AWS.  
(4) At about 8PM, my site is suddenly live again  
(5) I go to AWS console, I still don't have an instance  
(6) I was able to login using SSH and FileZila and started downloading files  
(7) Today at 5AM on OCt 22, the instance is still on available on console but I am able to SSH into my instance  


And AWS wants to still charge me $29.99 to give me support else all I get is talk to banking associates. 

This was my first time working with AWS. What kind of platform it is where things break and there is no one to help even though I am a paid client.
## [5][Alternative to AWS Lambda for data processing](https://www.reddit.com/r/aws/comments/jfwrnr/alternative_to_aws_lambda_for_data_processing/)
- url: https://www.reddit.com/r/aws/comments/jfwrnr/alternative_to_aws_lambda_for_data_processing/
---
Hi all,

I've been running some Lambda functions which are dealt and created from a queue. It's a python job which loads a couple of layers, imports some data, does some analytics/transforms, then pushes the output to another place on AWS.

I've got a feeling I could be using something better and that Lambda functions aren't really designed for this type of thing? I've been looking at Glue. Maybe Batch?

If anyone has any insight into this, I'd love to hear it!
## [6][How private is the traffic within for example VPC or within a single subnet?](https://www.reddit.com/r/aws/comments/jfzjmr/how_private_is_the_traffic_within_for_example_vpc/)
- url: https://www.reddit.com/r/aws/comments/jfzjmr/how_private_is_the_traffic_within_for_example_vpc/
---
This questions should be looked with encryption between services in mind. The reason for asking this is that we have an application that cannot support TLS (dont ask........) but it is critical for the big feature itself. There will be a solution later but right now no TLS support.

So I am trying to understand the risks and how private the traffic really is in the three scenarios below:  


1. Traffic between two services WITHIN same subnet. What are the risks?
2. Traffic between two services between two subnets? What are the risks?
3. Traffic between two subnets in two different AZs?

In my head I think "between AZs it will pass multiple physical and logical routers, switches, racks, machine and if something along that infrastructure is poisioned they can extract the data".   
For the same subnet communication my instinctive thoughts are "that should be quite safe".

But I cannot find any information on the nitty gritty details of how things actually work, mainly best practices and AWS logical views which doesnt help much.

Thanks for any input or comments.
## [7][You can configure CloudWatch dashboard sharing from the console, but NOT from the cli/SDK. WTF?](https://www.reddit.com/r/aws/comments/jfwiwl/you_can_configure_cloudwatch_dashboard_sharing/)
- url: https://www.reddit.com/r/aws/comments/jfwiwl/you_can_configure_cloudwatch_dashboard_sharing/
---
Don't AWS always say that they eat their own dogfood? That the console is just a web page using the same public API as customers?

Why can't I share a dashboard using the cli/SDK? What API does the console use to do this, if not *the* AWS API?

Are there any other services/features like this? I can't think of any.
## [8][Getting a simple Reservation report on all regions from the command line](https://www.reddit.com/r/aws/comments/jfzhti/getting_a_simple_reservation_report_on_all/)
- url: https://www.reddit.com/r/aws/comments/jfzhti/getting_a_simple_reservation_report_on_all/
---
I've just added a cool command to get the account reservation status.

Just simply run `awsctl get ri --region all`

[https://github.com/omerh/awsctl/releases/tag/v0.0.8](https://github.com/omerh/awsctl/releases/tag/v0.0.8)

Report looks like this:

Reservation status in eu-west-2 (Only for running instances  
=================================================================  
\- You have 7 t3.medium with a resevation of 9  
\- You have 9 m5.large with a resevation of 8  
\- You have 1 t3.small with a resevation of 0

Appreciate the feedbacks
## [9][IAM Roles with AWS SSO](https://www.reddit.com/r/aws/comments/jfzh9f/iam_roles_with_aws_sso/)
- url: https://www.reddit.com/r/aws/comments/jfzh9f/iam_roles_with_aws_sso/
---
We are currently evaluating, if we should use AWS SSO with PermissionSets or IAM users and roles to grant access to AWS accounts in our AWS Organization.

&amp;#x200B;

Previously we had an identity account that held all our IAM users and allowed different groups to assume roles in different accounts we had.

We now migrated these accounts into a single AWS Organization using AWS ControlTower and created some AWS SSO Users with the predefined PermissionSets AWS SSO was configured with.

We are using Terraform to provision our infrastructure (where possible.... looking at you AWS SSO).

&amp;#x200B;

How is the AWS SSO model supposed to work with custom policies in each account. Let's say, we have a Frontend-Developer Role that needs access to CloudFront and S3 in some of our accounts. Previously I could create an IAM role that was specific to a set of ressources in an account and deploy that to the accounts holding those ressources. Now with SSO I'd have to create an SSO PermissionSet using the Console (API support was just released I believe) with the specific resources, right? That gets quite messy when we do that for each role in each type of account we have. On the other hand having the comfort of the SSO portal would be great.

How do you others handle this?
## [10][Anyone have experience setting up Cognito Hosted UI with Authorization Code Grant?](https://www.reddit.com/r/aws/comments/jfzffl/anyone_have_experience_setting_up_cognito_hosted/)
- url: https://www.reddit.com/r/aws/comments/jfzffl/anyone_have_experience_setting_up_cognito_hosted/
---
MY goal is to setup the Cognito Hosted UI to validate users after login. I have followed the steps laid out in the OAuth2 blog here: [https://developer.okta.com/blog/2018/04/10/oauth-authorization-code-grant-type](https://developer.okta.com/blog/2018/04/10/oauth-authorization-code-grant-type)

&amp;#x200B;

My steps are as follows.

&amp;#x200B;

1. User logs into the AWS provided login screen.

2. It redirects to my website and I pull down the authorization code in Angular.

3. I send the code as part of my headers to the backend Nodejs

4. I use the code to get a token and then validate the token

&amp;#x200B;

This stream works but then what? I want to validate the AWS token for each API call but I have no idea how to access the token.

&amp;#x200B;

I am guessing that I am either missing the point of this procedure or that the token is somewhere I am unaware of. 

&amp;#x200B;

Any help would be greatly appreciated.
## [11][From AWS noob to AWS Specialist? - Data Analytics and Machine Learning](https://www.reddit.com/r/aws/comments/jfzd7q/from_aws_noob_to_aws_specialist_data_analytics/)
- url: https://www.reddit.com/r/aws/comments/jfzd7q/from_aws_noob_to_aws_specialist_data_analytics/
---
Completely new to AWS but considering pursuing specialist qualifications for AWS Data Analytics and Machine Learning. I'm aware I'll have to go through other stages but wanted to know the recommended approach and resources. I'm aware there is also recommended experience but I read up on it a bit and it's possible to clear the certification without it. 

I'd appreciate if anyone could offer some insights as to how long it could approximately take from Foundational to Specialist and how much financial investment is required for studying materials and the exams. I understand learning speed is different for each individual but knowing estimated learning hours would help me balance other responsibilities as I already work and study. 

Also, besides the training tutorials, what other resources/courses could I use to aid learning for the different certifications?

I'm UK based, if it helps.

Please let me know if there are any questions.
