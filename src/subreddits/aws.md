# aws
## [1][boto3 users, I've built a (hopefully) better reference site for us](https://www.reddit.com/r/aws/comments/g74giy/boto3_users_ive_built_a_hopefully_better/)
- url: https://www.reddit.com/r/aws/comments/g74giy/boto3_users_ive_built_a_hopefully_better/
---
[https://botodocs.com](https://botodocs.com)

&amp;#x200B;

I've released an open-source doc site for boto3. It focuses on making the doc more navigable and clearer docs. Let me know what you think!

It solves some frustrations with the current doc site, e.g:

1. opening some pages like ec2 freezes the page and I have to wait to use it. (terrible for usability)
2. redundant pseudocode and documentation that make pages extremely long (print preview ec2 shows it's 1700+ pages!)
3. unclear what parameters are required or optional (unless you like to scroll a lot)
4. with full refreshes, pages are often slow to load (which disrupts my flow)

Leveraging docsify, it caches pages, have offline search and a great default UI.

It's has full coverage of the boto3 features, viz clients, resources, paginators, etc and it will be deployed frequently so that you can enjoy up-to-date docs.

&amp;#x200B;

Let me know what you think. It's also on github if you want to contribute (or give a star, much appreciated!)

[https://github.com/jeshan/botodocs](https://github.com/jeshan/botodocs)
## [2][issues with resetting MFA in order to access AWS console](https://www.reddit.com/r/aws/comments/g75t11/issues_with_resetting_mfa_in_order_to_access_aws/)
- url: https://www.reddit.com/r/aws/comments/g75t11/issues_with_resetting_mfa_in_order_to_access_aws/
---
Hi,

its a pity AWS doesnt offer support for multiple U2F keys or at least to have two options available at the same time TOTP and U2F, or even recovery codes option..

but I was trying to do some tests and bypass my MFA with resetting it with e-mail &amp; phone verification, the problem is that even though I have added phone to the "contact information " section in account settings, after step 1(e-mail) on step 2 (phone) i have the following error: "Phone verification could not be completed", no matter what number and country code I put there,

did you also have such issue ?

BR
## [3][My visual notes on the AWS Transfer Family w/support for FTP, FTPS &amp; SFTP](https://www.reddit.com/r/aws/comments/g6uzsk/my_visual_notes_on_the_aws_transfer_family/)
- url: https://www.awsgeek.com/AWS-Transfer-Family/
---

## [4][CloudWatch Synthetics - Monitor Sites, API Endpoints, Web Workflows, and More](https://www.reddit.com/r/aws/comments/g6u4kc/cloudwatch_synthetics_monitor_sites_api_endpoints/)
- url: https://aws.amazon.com/blogs/aws/new-use-cloudwatch-synthetics-to-monitor-sites-api-endpoints-web-workflows-and-more/
---

## [5][Route my home traffic through AWS, too expensive?](https://www.reddit.com/r/aws/comments/g6sqvt/route_my_home_traffic_through_aws_too_expensive/)
- url: https://www.reddit.com/r/aws/comments/g6sqvt/route_my_home_traffic_through_aws_too_expensive/
---
Im thinking about routing my home networks trafic through an AWS hostet EC2 which runs analysis on my trafic. RIght now i have some security applications on-prem(at home) like Pfsense and the ELK stack, id like to have this in AWS and have my traffic run through it there, however im not sure how expensive something like this would be considering it would be a fuckton of traffic as i play games, stream, download frequently. Anyone who has more knowledge on AWS traffic cost than me who has some input on this?
## [6][RDS Postgres IAM Authentication w/ Password for users](https://www.reddit.com/r/aws/comments/g77tl7/rds_postgres_iam_authentication_w_password_for/)
- url: https://www.reddit.com/r/aws/comments/g77tl7/rds_postgres_iam_authentication_w_password_for/
---
Hi everyone, 

I would like to know if it is possible to set up a Postgres RDS to use IAM users so that users can login to the database as they normally would, just with the IAM username and password? 

I have made an attempt, but I feel something is missing. I created an IAM user with a password. I enabled IAM users on the database. I then added a new user in the database with the same name as the IAM user and gave it the role iam\_user. When I try to login, it says the password is incorrect. 

PS: I want to use passwords instead of the token that expires every 15 minutes because the database clients do not accept the security keys and secrets, they are username/password only.
## [7][Using Amazon CloudWatch Synthetics and Service Lens | Caylent](https://www.reddit.com/r/aws/comments/g77tgc/using_amazon_cloudwatch_synthetics_and_service/)
- url: https://caylent.com/using-amazon-cloudwatch-synthetics-and-service-lens
---

## [8][Is there a kinda “factory reset” option for Lightsail?](https://www.reddit.com/r/aws/comments/g6xqqn/is_there_a_kinda_factory_reset_option_for/)
- url: https://www.reddit.com/r/aws/comments/g6xqqn/is_there_a_kinda_factory_reset_option_for/
---
I’ve messed my site up and want to just start over. I have a backup but a lot of the backend stuff is broken so I just need to set it up as new. I still own the domain through AWS.
## [9][Amazon Appflow](https://www.reddit.com/r/aws/comments/g6u4yy/amazon_appflow/)
- url: https://www.reddit.com/r/aws/comments/g6u4yy/amazon_appflow/
---
 [https://techcrunch.com/2020/04/22/aws-launches-amazon-appflow-its-new-saas-integration-service/](https://techcrunch.com/2020/04/22/aws-launches-amazon-appflow-its-new-saas-integration-service/) 

Is this a direct attack against companies like [Segment.io](https://Segment.io) and [Funnel.Io](https://Funnel.Io)?
## [10][Amazon Keyspaces (Managed Cassandra) Now GA](https://www.reddit.com/r/aws/comments/g70w93/amazon_keyspaces_managed_cassandra_now_ga/)
- url: https://aws.amazon.com/blogs/aws/new-amazon-keyspaces-for-apache-cassandra-is-now-generally-available/
---

