# aws
## [1][Does anyone know how S3 Intelligent Tiering works under the hood?](https://www.reddit.com/r/aws/comments/fr33ai/does_anyone_know_how_s3_intelligent_tiering_works/)
- url: https://www.reddit.com/r/aws/comments/fr33ai/does_anyone_know_how_s3_intelligent_tiering_works/
---
I know it uses machine learning, but I'd like to know more about the algorithm itself. 

I'm thinking about implementing something similar on my personal PC using Python, to move big folders (looking at you, games) from my SSD to a HDD if not touched for a while.
## [2][In-progress aws pricing calculator](https://www.reddit.com/r/aws/comments/fr4i4t/inprogress_aws_pricing_calculator/)
- url: https://www.reddit.com/r/aws/comments/fr4i4t/inprogress_aws_pricing_calculator/
---
Working on an aws pricing calculator for newbies like me who just want an easy to understand number. [https://aws-calculator.com/](https://aws-calculator.com/)

The above is a proof of concept of sorts ill be iterating on, only has EC2 Instances and Fargate atm. Next ill be adding storage followed by the ability to 'save' an estimate.

Thoughts, feedback, suggestions welcomed.
## [3][Could someone help me set up SSL on my EC2 Instance?](https://www.reddit.com/r/aws/comments/fqv4fk/could_someone_help_me_set_up_ssl_on_my_ec2/)
- url: https://www.reddit.com/r/aws/comments/fqv4fk/could_someone_help_me_set_up_ssl_on_my_ec2/
---
Hi! My name is Jack, and I am brand new to AWS, and need to set up a SSL certificate on my EC2 instance, running Amazon's Linux distro and https. Please send me a PM if you are willing to walk me through the process!
## [4][Should I use Amplify or build out my backend without it?](https://www.reddit.com/r/aws/comments/fqpio2/should_i_use_amplify_or_build_out_my_backend/)
- url: https://www.reddit.com/r/aws/comments/fqpio2/should_i_use_amplify_or_build_out_my_backend/
---
I've been working on a project and have been gearing up to work on a serverless backend. I was planning on using Cognito for authentication / user management, AppSync for the API, Lambda for backend functions, and DynamoDB for storage.

A lot of AWS' documentation and resources on the usage of these services together suggest to use Amplify.

I guess my concern with using Amplify is that I'm worried I'm not going to learn these services as well as I could if I didn't use Amplify. I'm also worried that if I implemented Amplify and wanted greater customization in the future, that it might be difficult to migrate away from Amplify.

Would you recommend using Amplify in favor of not using it? Why / why not?

Edit: I guess I'm interpreting Amplify as training wheels, or a way to quickly scaffold a backend.
## [5][DynamoDB Table design - to RDBS](https://www.reddit.com/r/aws/comments/fqq488/dynamodb_table_design_to_rdbs/)
- url: https://www.reddit.com/r/aws/comments/fqq488/dynamodb_table_design_to_rdbs/
---
Im used to relational DBs and cant quite get my head around NoSQL.

I have been looking at DDB design best practices but cant see how to model my data to a DDB table.

The application contains Events;Event have attributes:

Title (String)  
DataTime (String)  
Duration (String)  
Location (String)  
Attendees (List of String)

The access patterns are as follows;Get events by TitleGet events by date rangeGet events by locationGet attendees by event

I thought I could model it as follows;  
PK          | SK                | Data  
eventId | title              | data  
eventId | dataTime    | data  
eventId | attendee1  | data  
eventId | attendee2  | data  
eventId | eventDetails | non queryable attribute e.g. duration/created/updated/active

Now I see that if I have the eventId I can easily query for the title/date/attendee.But if I dont know the eventId and want to query by title only the above design wont work.

How best can I model the data above?
## [6][Manage in-flight transactions while RDS migration](https://www.reddit.com/r/aws/comments/fr1mqt/manage_inflight_transactions_while_rds_migration/)
- url: https://www.reddit.com/r/aws/comments/fr1mqt/manage_inflight_transactions_while_rds_migration/
---
Hi All,

I am working on a cross-region partial migration of the RDS instance. It means I will be only migrating a few tables from one region to another in the end. I have thought of taking a `dump` of DB and importing it in new RDS while selecting what all I would need in new RDS.

However, I am more worried about any transaction that would be done meanwhile to the existing RDS.  One way I could manage this is before starting the migration I would start pushing the events to an SQS (FIFO) along with existing RDS (still need this in case things go wrong) and can poll a lambda for moving the transaction out of SQS to new RDS. 

Is there any other thing I can do here?
## [7][Is AWS manually extending recertification dates?](https://www.reddit.com/r/aws/comments/fr1kgr/is_aws_manually_extending_recertification_dates/)
- url: https://www.reddit.com/r/aws/comments/fr1kgr/is_aws_manually_extending_recertification_dates/
---
I received a notification over the weekend stating my SA certs were recertified and now expiring later this year instead of in the coming months. I haven't taken *any* AWS cert exams as of late, so there's no chance I took a higher level one that automatically recertified my lower level things.

Anyone else seeing similar extensions? I assume it's due to COVID, but haven't seen anything official other than the exams being available online now.

[Certs Extended](https://imgur.com/Xp4EbGy)
## [8][Can you version control a CodePipeline pipeline that uses GitHub source?](https://www.reddit.com/r/aws/comments/fqlr3a/can_you_version_control_a_codepipeline_pipeline/)
- url: https://www.reddit.com/r/aws/comments/fqlr3a/can_you_version_control_a_codepipeline_pipeline/
---
I'm using AWS CodeStar and yesterday, after deleting an old CloudFormation stack, I go back to my CodeStar dashboard to find that my CD pipeline is gone. WTF?

Anyways, I can recreate it OK but I'd like to get it version controlled so when it randomly disappears again I can recreate it. I tried the CLI but it turns out, because I'm using GitHub for source, I can't seem to recreate the pipeline in the CLI from JSON because I can't pass in an OAuth code.

Anyone else have this problem? Anyone else able to version control their AWS CI/CD configuration?
## [9][Securing WIX subdomain connected (externally) to AWS](https://www.reddit.com/r/aws/comments/fr0zqq/securing_wix_subdomain_connected_externally_to_aws/)
- url: https://www.reddit.com/r/aws/comments/fr0zqq/securing_wix_subdomain_connected_externally_to_aws/
---
Hi, I am Leon and i'm relatively new to AWS and was trying to secure my WIX subdomain ([api.example.co.za](https://api.example.co.za)) which is connected to AWS. So far, I have an EC2 instance for the website, as well as my S3 bucket running. I have also set up my CloudFront and pointed my subdomain on wix ([api.example.co.za](https://api.example.co.za)) to CloudFront ([df\*\*\*.cloudfront.net](https://dfxxx.cloudfront.net)) within WIX DNS records. I got an AWS certificate through DNS verification and attached it to my CloudFront. In AWS CloudFront, I also routed all requests to my S3 bucket. I also redirected all traffic from http to https in CloudFront

When I type into my brower [api.example.co.za](https://api.example.co.za), I am able to access my website, but insecurely, just as before. But when I type [https://api.example.co.za](https://api.example.co.za), the website is secure, but does not load the web content. When Inspect the web content in Chrome&gt;Inspect&gt;Console, I see the following errors

Mixed Content: The page at '[https://api.example.co.za](https://api.example.co.za)' zone-evergreen.js:29\*\* was loaded over HTTPS, but requested an insecure XMLHttpRequest endpoint '[http://api.example.co.za/users/5e7a2](http://api.example.co.za/users/5e7a2)'. This request has been blocked; the content must be served over HTTPS.
## [10][Trouble renewing Let's Encrypt Certificate on Lightsail](https://www.reddit.com/r/aws/comments/fqxs5p/trouble_renewing_lets_encrypt_certificate_on/)
- url: https://www.reddit.com/r/aws/comments/fqxs5p/trouble_renewing_lets_encrypt_certificate_on/
---
I don't speak command line, so I am always copy/pasting from tutorials. Renewal of certificates was bumpy, but worked so far. Today I got a weird response from the server and cannot go on.

I am using [this tutorial](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-using-lets-encrypt-certificates-with-wordpress#request-a-lets-encrypt-certificate-wordpress) for renewal,  and started with   
[DOMAIN=mydomain.com](https://DOMAIN=mydomain.com)

WILDCARD=\*.$DOMAIN

echo $DOMAIN &amp;&amp; echo $WILDCARD

and 

sudo certbot -d $DOMAIN -d $WILDCARD --manual --preferred-challenges dns certonly

&amp;#x200B;

Today I get a response

Certbot can obtain and install HTTPS/TLS/SSL certificates.  By default,

it will attempt to use a webserver both for obtaining and installing the

certificate. 

certbot: error: argument -d/--domains/--domain: expected one argument  


I have done these steps before, but now I am lost.  


Anybody..?

  
Thanks!
