# aws
## [1][Anybody have a good (spark?) recipe for transferring from Aurora PostgreSQL 9.6 to Aurora PostgreSQL 10/11?](https://www.reddit.com/r/aws/comments/epdk3a/anybody_have_a_good_spark_recipe_for_transferring/)
- url: https://www.reddit.com/r/aws/comments/epdk3a/anybody_have_a_good_spark_recipe_for_transferring/
---
DMS still does not properly support PostgreSQL and dies on some of my tables. Unfortunately the tables that don't work on DMS contain over 250 million rows and are not CSV friendly.

I've had some success with Apache Spark reading the source DB table in to a dataframe and then writing the DF to the target DB table. I have been using a single instance with Spark-Shell. Unfortunately, even on a machine with 64gb of RAM and 50gb committed to the Spark JVM, I am getting OOM errors. Unfortunately I haven had much success getting Spark to read+write in smaller batches or to find decent batch-load examples. I'm excited about Spark because you can do a diff of two tables and it seems like the perfect solution for my FB migration.

How have you tackled the Aurora pg9 to Aurora pg10/11 migration?

Do you have a Spark/EMR recipe that worked for you?
## [2][Re:Invent 2019 DOP301 removed from YouTube](https://www.reddit.com/r/aws/comments/ep4s1u/reinvent_2019_dop301_removed_from_youtube/)
- url: https://www.reddit.com/r/aws/comments/ep4s1u/reinvent_2019_dop301_removed_from_youtube/
---
Andy Troutman's excellent talk on service teams DOP301 https://www.portal.reinvent.awsevents.com/connect/sessionDetail.ww?SESSION_ID=95548&amp;tclass=popup&amp;csrftkn=4ZQ1-JBFI-EYJ0-IRP0-3F31-5TZO-R93F-0SGJ was available on the AWS Events YouTube channel as of Monday, but is now marked as private. I was going to watch it again before a discussion at work on organizational structure but now I'll have to rely on memory. Does anyone have any idea why? Is it common for session recordings to be taken down?

UPDATE: Thanks to u/k8s_helm for posting an audio link for the talk at http://aws-reinvent-audio.s3-website.us-east-2.amazonaws.com/2019/DOP301-R1.m4a
## [3][How to leverage Aws features for sending out invoices?](https://www.reddit.com/r/aws/comments/epeter/how_to_leverage_aws_features_for_sending_out/)
- url: https://www.reddit.com/r/aws/comments/epeter/how_to_leverage_aws_features_for_sending_out/
---
So, I'm using dynamoDB for storage. It consists of payments, users and among other things. My intended task at hand is to send out an invoice to each user via email when they make a new payment or clear previous dues. 

I know I should be using AWS SES for sending out emails. I want to know how and where should I be modelling invoice to make it look like an invoice and send to users.  Hope my intention is understood. Please kindly provide some inputs.
## [4][Connecting to RDS from internet-connected Lambda](https://www.reddit.com/r/aws/comments/epe796/connecting_to_rds_from_internetconnected_lambda/)
- url: https://www.reddit.com/r/aws/comments/epe796/connecting_to_rds_from_internetconnected_lambda/
---
Hello all,

I'm trying to connect to and query my PostgreSQL DB running in RDS from a Lambda API which does some external HTTP requests (so the lambda needs internet access). Can anyone suggest a solution?

Thanks
## [5][Cost Matters! The Serverless Edition - Cloud Journey IO](https://www.reddit.com/r/aws/comments/ep7osl/cost_matters_the_serverless_edition_cloud_journey/)
- url: https://www.cloudjourney.io/articles/publiccloud/cost-matters-the-serverless-edition-ls/
---

## [6][Part 6: Building an Imgur Clone in Vue.js and Serverless - Login/Registration Flow - Now Live!](https://www.reddit.com/r/aws/comments/ep23iv/part_6_building_an_imgur_clone_in_vuejs_and/)
- url: https://tutorialedge.net/projects/building-imgur-clone-vuejs-nodejs/part-6-login-register-flow/
---

## [7][AWS Quicksight - MySQL UNION from multiple databases](https://www.reddit.com/r/aws/comments/epeoho/aws_quicksight_mysql_union_from_multiple_databases/)
- url: https://www.reddit.com/r/aws/comments/epeoho/aws_quicksight_mysql_union_from_multiple_databases/
---
I'm using AWS Quicksight for an Analytics dashboard and I have multiple databases that have the same tables. I then set a Data Source for each database. Now, I've been trying to create a Data Set for a table, let's say "products" table, which should be a UNION of all the "products" tables of all the databases.

So far, I've only been able to combine data from those databases using the UI to do a JOIN. However, for the Analysis I'm trying to create, I'll be needing to do a UNION. How do you do a UNION from multiple different databases in AWS Quicksight?

Thanks.
## [8][Choosing the right AWS services for a small business](https://www.reddit.com/r/aws/comments/ep9sor/choosing_the_right_aws_services_for_a_small/)
- url: https://www.reddit.com/r/aws/comments/ep9sor/choosing_the_right_aws_services_for_a_small/
---
Hi, I am having trouble finding the right setup for my business. I would very much appreciate some advice. We currently run our applications in a number of Docker containers on a dedicated server. This setup seems outdated and not very cost-effective, which is why we are considering moving to AWS. All our applications are built from the ground up, and we're not very big, so there is enough room to change things around now. We have been online for less than half a year, but at peak times we already notice our website getting slower. 

My business sells renders of 3D scenes in high-DPI format, for printing promotional materials and such. The setup is as follows: there is a static website which communicates with an API in order to send and receive relational business data (orders, products, accounts, etc.) from a Postgres database. Then there is a 'worker' process, which will periodically poll this database and check whether there are new orders that require rendering. Rendering is done via WebGL in a Docker container. The API, static web-server and worker are written in Go.

This leaves me with the following questions about which services to use:

1. API: This could be done in a Lightsail VPS with Docker, or I could rewrite each HTTP route as a Lambda function. I expect avg. 4 API calls per visitor to my website.
2. Static website: This could be done via S3 if I pre-render the Go/HTML templates into static pages. Does this approach allow me to bring my own routes, rather than point to specific files?
3. Database: AWS offers a managed database and a serverless database, or I could host my own on a Lightsail VPS. What considerations should I make?
4. Rendering task: This could be a Lambda function or a Fargate task. Is there any inherent benefit to Lambda over Fargate for periodic tasks?
5. Is AWS overall even worth considering if I am a small business (less than &lt;1000 serious visitors per day)? Maybe we're not the target market.

Any further suggestions or resources are much appreciated!
## [9][Authenticate logins to AWS VPN using Azure AD?](https://www.reddit.com/r/aws/comments/epcori/authenticate_logins_to_aws_vpn_using_azure_ad/)
- url: https://www.reddit.com/r/aws/comments/epcori/authenticate_logins_to_aws_vpn_using_azure_ad/
---
Hey folks,

Title kinda says it all.

We have an Azure AD instance and AADFS.

We have an AWS virtual private cloud.

We want to set up [AWS client VPN](https://aws.amazon.com/vpn/) to allow remote users to VPN into the AWS instance, and have them authenticate against their Azure AD credentials.

How can we make this work?

Thanks!
## [10][Modernizing a 3-Tier Application with Serverless on AWS - [IOD.tech.content]](https://www.reddit.com/r/aws/comments/ep4vju/modernizing_a_3tier_application_with_serverless/)
- url: https://iamondemand.com/blog/modernizing-a-3-tier-application-with-serverless-on-aws/
---

