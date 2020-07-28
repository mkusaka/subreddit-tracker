# aws
## [1][Week of July 27th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/hysmye/week_of_july_27th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/hysmye/week_of_july_27th_what_are_your_favorite_aws_tips/
---
Share your favorite AWS tips!
## [2][us-east-1 DNS issues](https://www.reddit.com/r/aws/comments/hzbcry/useast1_dns_issues/)
- url: https://www.reddit.com/r/aws/comments/hzbcry/useast1_dns_issues/
---
Is anyone else experiencing DNS resolver issues right now in US-east-1?  Started noticing it around 4:45 AM EST.
## [3][New – Amazon EC2 Instances based on AWS Graviton2 with local NVMe-based SSD storage](https://www.reddit.com/r/aws/comments/hz1iuy/new_amazon_ec2_instances_based_on_aws_graviton2/)
- url: https://aws.amazon.com/blogs/aws/new-graviton2-instance-types-c6g-r6g-and-their-d-variant/
---

## [4][Series of MINI Projects/Advanced Demos for AWS](https://www.reddit.com/r/aws/comments/hz6myp/series_of_mini_projectsadvanced_demos_for_aws/)
- url: https://www.reddit.com/r/aws/comments/hz6myp/series_of_mini_projectsadvanced_demos_for_aws/
---
&amp;#x200B;

[You want demos ..... I have demos](https://preview.redd.it/7rls18qqfid51.png?width=2462&amp;format=png&amp;auto=webp&amp;s=bfd747cd80fc97047d6f376300a9225c397f8046)

**TL;DR I've been working on a series of Advanced Demos/MINI Projects for AWS for the past few weeks.  I've made them for my AWS training courses .. but I want to help the community so they are all freely available for usage.. MIT licensed.**

**(Full disclosure, yes I create AWS courses at** [**https://learn.cantrill.io**](https://learn.cantrill.io) \- but these demos can be used without ever visiting that site - in my courses you just get companion theory and video guides). The demos currently have an architecture focus .. because ive been working on SAA-C02 and SAP-C01 courses, the ones i've planned for next will be more OPS/NETWORK focussed.

My aim with these demos is to give my students (and now anyone) with practical experience working with more complex scenarios and architectures within AWS. These aren't 'pure' AWS demos - they all generally simulate a 'less clean' or 'hybrid' environment between AWS and On-premises - giving you real world useful skills.

I'm busy working on a way to collate and maintain these as seperate things, for now.. I want to make this post a bit of a collection. All of the assets for these demos will be created and updated at [https://github.com/acantril/learn-cantrill-io-labs](https://github.com/acantril/learn-cantrill-io-labs) keep an eye on these if you want more as they are released.

\#1 **Implement a BGP Based, HA, Dynamic Site-to-Site VPN with no hardware**

Create a fully HA, Dynamic, BGP based, multi connection, multi Tunnel Site-to-Site VPN between AWS and On-premises. The On-premises part is simulated so you don't need any local physical hardware

[https://www.reddit.com/r/aws/comments/hjtq23/mini\_project\_implement\_a\_bgp\_based\_ha\_dynamic/](https://www.reddit.com/r/aws/comments/hjtq23/mini_project_implement_a_bgp_based_ha_dynamic/)

\#2 **Hybrid Active Directory in AWS &amp; Simulated On-Premises - No hardware (or premises) needed**

Experience how to integrate on-premises active directory with AWS. You get to integrate two Active Directory Domains with a two-way forest trust and use AWS services with on-premise identities. The demo uses Directory Service, FSx and Workspaces.

[https://www.reddit.com/r/AWSCertifications/comments/hmnyft/mini\_project\_hybrid\_active\_directory\_in\_aws/](https://www.reddit.com/r/AWSCertifications/comments/hmnyft/mini_project_hybrid_active_directory_in_aws/)

\#3 **Implement a Hybrid DNS Architecture between AWS and Simulated On-premises using R53 Endpoints**

integrate the DNS platforms of AWS and a linux based, simulated on-premises environment using Route53 inbound and outbound endpoints.

[https://www.reddit.com/r/AWSCertifications/comments/hq8iqe/mini\_project\_implement\_a\_hybrid\_dns\_architecture/](https://www.reddit.com/r/AWSCertifications/comments/hq8iqe/mini_project_implement_a_hybrid_dns_architecture/)

\#4 **Patching with AWS Systems Manager in a Hybrid Environment**

Experience using systems manager in a Hybrid Environment, including hybrid agent installation

[https://www.reddit.com/r/AWSCertifications/comments/hyigt9/mini\_project\_patching\_with\_aws\_systems\_manager\_in/](https://www.reddit.com/r/AWSCertifications/comments/hyigt9/mini_project_patching_with_aws_systems_manager_in/)

Hope you enjoy them, comment if you have any questions/comments ....if you have any cool demo ideas I'm all ears ... 

/Adrian
## [5][AWS Cognito - verify token using JWT vs cognito.getUser SDK](https://www.reddit.com/r/aws/comments/hze7z5/aws_cognito_verify_token_using_jwt_vs/)
- url: https://www.reddit.com/r/aws/comments/hze7z5/aws_cognito_verify_token_using_jwt_vs/
---
I am using the following doc by AWS for verifying the incoming token with Cognito to verify the user in cognito pool: [https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html)

Now I also need to verify one of the custom attributes (email) of that user. (Basically this is to mitigate a case where some user can change their email-id in the API payload to access someone else's details).

The Cognito SDK I found to do this was this : [https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API\_GetUser.html](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_GetUser.html)

It works just fine, however this led me to a question : Since I'm sending access token in this SDK, do I need to use the former jwt-based token verification as well? Because this SDK also handles the case where token is invalid/expired and sends error codes accordingly. Is there something I'm missing which the former case handles and the latter one doesn't ?
## [6][Can someone tell me what permissions I would have to apply for a user to see S3 Access?](https://www.reddit.com/r/aws/comments/hze3np/can_someone_tell_me_what_permissions_i_would_have/)
- url: https://www.reddit.com/r/aws/comments/hze3np/can_someone_tell_me_what_permissions_i_would_have/
---
I am setting up an S3 bucket for backup purposes for something in house, I set up a test user to see how my policy is working in real time (the policy sim is good, but not enough for my visual learning curve,) and I was wondering what permissions will allow them to see bucket access?

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/un1i5l8khld51.png?width=1557&amp;format=png&amp;auto=webp&amp;s=3882e871468c882d5ff25809aff7b465d0363f9b
## [7][AWS ECS for Kubernetes: When to use it? | ECS vs EKS](https://www.reddit.com/r/aws/comments/hzdysh/aws_ecs_for_kubernetes_when_to_use_it_ecs_vs_eks/)
- url: https://www.padok.fr/en/blog/aws-ecs-kubernetes
---

## [8][Restricting Amplify Access](https://www.reddit.com/r/aws/comments/hzduj5/restricting_amplify_access/)
- url: https://www.reddit.com/r/aws/comments/hzduj5/restricting_amplify_access/
---
Hi All,

I have been building a web app using Amplify as part of a project.

A security requirement is that access is locked down to certain IP ranges.

Has anyone worked out a way to do this? There seems to be no policy or security groups in front of Amplify that can be altered. The only 'security' feature is to password protect the site.

The only thing I can think of is using the geolocation feature on Route53 but that is not very robust.

Any ideas would be fantastic,

Thanks,
## [9][DynamoDB: TransactionConflict on ConditionCheck item in TransactWriteItems](https://www.reddit.com/r/aws/comments/hzdikr/dynamodb_transactionconflict_on_conditioncheck/)
- url: https://www.reddit.com/r/aws/comments/hzdikr/dynamodb_transactionconflict_on_conditioncheck/
---
I'm using TransactWriteItems with a ConditionCheck to check a foreign key constraint. So I have two items in my transaction, the ConditionCheck, which checks for the existence of a parent item in the table, and a Put, which inserts a child item, which contains an attribute that references the parent. When I concurrently insert multiple children that all reference the same parent, I often get a TransactionCanceledException, where the CancellationReason is "TransactionConflict" and its position in the array corresponds to the condition check. As far as I can tell, the parent item isn't being written to concurrently. Is there some other reason I would be getting a conflict in this situation?
## [10][Best way to verify an entry in Cloudwatch with nodejs?](https://www.reddit.com/r/aws/comments/hzddt3/best_way_to_verify_an_entry_in_cloudwatch_with/)
- url: https://www.reddit.com/r/aws/comments/hzddt3/best_way_to_verify_an_entry_in_cloudwatch_with/
---
I'd like to perform a test where I verify a log entry in Cloudwatch after performing an action, but I'm not sure of the best approach. Is there a recommended library for this?
## [11][Change AWS instance type from c5.large to t3.large](https://www.reddit.com/r/aws/comments/hz6bzy/change_aws_instance_type_from_c5large_to_t3large/)
- url: https://www.reddit.com/r/aws/comments/hz6bzy/change_aws_instance_type_from_c5large_to_t3large/
---
Hi - I would like to change some of my production instances from c5.large to t3.large. As this is a change between different instance families, should I just power down the instance and select the t3.large or I need to create an AMI first, power down the c5.large, then relaunch using the AMI?

I know that changing between instance families can tricky so I wanted some advice if I missed out on anything.
