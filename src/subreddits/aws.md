# aws
## [1][AWS GitHub Actions 🚀](https://www.reddit.com/r/aws/comments/ff1o4h/aws_github_actions/)
- url: https://www.reddit.com/r/aws/comments/ff1o4h/aws_github_actions/
---
I created some GitHub actions for common AWS deployment tasks, open to any feedback 

[https://github.com/clowdhaus/aws-github-actions](https://github.com/clowdhaus/aws-github-actions)

&amp;#x200B;

The IAM access credentials is similar to the AWS provided action except I've added the ability to assume roles from the action. Let me know what you think - unfortunately you won't find these from the GitHub marketplace because they do not support monorepos but the actions do support monorepos ([I know, confusing](https://help.github.com/en/actions/reference/workflow-syntax-for-github-actions#example-using-a-public-action-in-a-subdirectory)).
## [2][Hosting WP on AWS](https://www.reddit.com/r/aws/comments/ffaebt/hosting_wp_on_aws/)
- url: https://www.reddit.com/r/aws/comments/ffaebt/hosting_wp_on_aws/
---
I’ve successfully setup a bucket for one of our WordPress sites to store and serve images. 

Ive read aws can be utilized for website hosting so I’m wondering is AWS can be used to host an entire wp site rather then the regular hosting platforms. 

Any opinions and resources will be appreciated, thanks.
## [3][Best way to use DMS to make a local (i.e., on the developer's laptop) copy of a database?](https://www.reddit.com/r/aws/comments/ff52z6/best_way_to_use_dms_to_make_a_local_ie_on_the/)
- url: https://www.reddit.com/r/aws/comments/ff52z6/best_way_to_use_dms_to_make_a_local_ie_on_the/
---
This sort of operation is mentioned in the Database Migration Service overview but I haven't really found a good best-practices walkthrough of how to do it.  I have a Postgres database that developers need to clone to their local machine for dev/testing, and along the way I need to anonymize/scramble a few columns to remove PII.

The export from the production database doesn't need to be up-to-the-minute accurate and a simple export every 12 or 24 hours is fine; I'd also want to trigger a migration/export via our CI/CD system.  I'd be ok using DMS to go to a dev RDS server as an intermediate step.

Is DMS from Prod to Dev, then scrambling the PII in dev, then exporting a dump of the dev database to S3 the right way to think about this?  Then the developers can just run a local script to import the S3 dump files when they need it?
## [4][SNS + Discord](https://www.reddit.com/r/aws/comments/ffb4v5/sns_discord/)
- url: https://www.reddit.com/r/aws/comments/ffb4v5/sns_discord/
---
Hi everyone.   
I'm quite new to AWS and I wanted to try to setup a pipeline which will send a notification every time the build is successful or it's failed.   
But I'm not quite sure how to achieve that. Anyone has any idea?   
Thank you in advance :)
## [5][How to setup CodeBuild for all branches except 1?](https://www.reddit.com/r/aws/comments/ff8gk5/how_to_setup_codebuild_for_all_branches_except_1/)
- url: https://www.reddit.com/r/aws/comments/ff8gk5/how_to_setup_codebuild_for_all_branches_except_1/
---
Hey all! As the title reads, how can I setup CodeBuild to build for *all* branches except for the default?

Use Case:

* Developers will test their apps using a CodeBuild + buildspec.yml
* When ready for production, the CodePipeline will use a similar buildspec + a full approval process

Is there any intuitive way to do this? Ideally, we don't want the CodeBuild to invoke when the Codepipeline is invoked.
## [6][I got locked out of my AWS Lightsail instance](https://www.reddit.com/r/aws/comments/ff2r7u/i_got_locked_out_of_my_aws_lightsail_instance/)
- url: https://www.reddit.com/r/aws/comments/ff2r7u/i_got_locked_out_of_my_aws_lightsail_instance/
---
Last night, I did an SSH access to my Lightsail instance to pull some changes to my application that was hosted on there. During the process, I figured I should update the system and I noticed that my firewall (ufw) was disabled. I did `sudo ufw enable` and didn't check the status. Turns out, the OpenSSH rule wasn't in the ufw by default so my SSH access through port 22 was disabled. As it was 4am and I was so tired, I didn't bother to check `sudo ufw status` to display my rules and logged out.

Today, when I tried connecting to my Lightsail instance through SSH, I kept getting timed out and after research I figured the problem was the firewall didn't allow SSH access to the instance. Last snapshot of the server was October 29th, 2019. and I cannot restore to that checkpoint as I would lose hundreds of user accounts that registered since and their uploaded content. It was a rookie mistake and now I don't know what to do to regain access to my instance. I tried creating a new shapshot and purchasing a new instance with the plan to add another storage unit and upload the snapshot to that storage unit, as that would allow me to have ssh access to the system partition and from there, I could restore my application. Unfortunately, it seems that Lightsail does not support uploading snapshots to new instances, and if I simply create a new instance from today's snapshot, I would again get an instance with a firewall blocking my SSH access.

&amp;#x200B;

Is there anything at all that I can do? The future of my platform might depend on me solving this issue.
## [7][Amazon SDK vs Using Amazon API Gateway to interact with Lambda Functions](https://www.reddit.com/r/aws/comments/ff3ezg/amazon_sdk_vs_using_amazon_api_gateway_to/)
- url: https://www.reddit.com/r/aws/comments/ff3ezg/amazon_sdk_vs_using_amazon_api_gateway_to/
---
At my company we're looking into hosting a new application in a container within a VPC.  We're torn on whether or not we should call Lambda Functions via an API Gateway or using the SDK.  It'll be an .NET Core 3.1 Web app contacting .NET Core 2.1 Lambda Functions.  If we were to host the API Gateway it would be Private to the VPC and not publicly facing.  We plan on using IAM Roles as much as possible as well as Secrets Manager.

Our major concerns are:

* Security
* Efficiency
* Performance

Does anyone have any experience with choosing between the two?
## [8][Locked out of RDS accessed via boto3](https://www.reddit.com/r/aws/comments/ff5lc6/locked_out_of_rds_accessed_via_boto3/)
- url: https://www.reddit.com/r/aws/comments/ff5lc6/locked_out_of_rds_accessed_via_boto3/
---
I'm trying to [execute_statement](https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/rds-data.html#RDSDataService.Client.execute_statement) so I can save to my database.


I'm getting the error "Access denied for user 'root'", I believe because I'm using the wrong parameters.


What is execute_statement looking for in the parameter "SecretARN"? Where can I find that?


Currently, I'm setting SecretARN = root user db password. I know the password's right, because I'm successfully using it to connect through dbeaver (which is a desktop SQL ... connector thingy).


Any ideas?
## [9][RDS/MySQL Community Engine import advice](https://www.reddit.com/r/aws/comments/ff83hk/rdsmysql_community_engine_import_advice/)
- url: https://www.reddit.com/r/aws/comments/ff83hk/rdsmysql_community_engine_import_advice/
---
Hi All,

We have an Amazon RDS, **My**SQL 5.7 and i have received a zipped `.sql` file that has inside the `INSERT INTO` commands.

What should be the first step of importing them ? is it possible to avoid building a web server for phpmyadmin or installing local tools like mysqlworkbench?

Thank you!
## [10][Using CloudWatch for business metrics?](https://www.reddit.com/r/aws/comments/feymu2/using_cloudwatch_for_business_metrics/)
- url: https://www.reddit.com/r/aws/comments/feymu2/using_cloudwatch_for_business_metrics/
---
Is anyone in here using CloudWatch for not just operational metrics but also business metrics? I'm talking about using CW for things like: conversion %, conversion abandonment %, time to convert, new user sign-up, etc.

I'm seeing a few advantages to this:

1. A change in some of these business metrics could indicate a problem in the system not readily apparent from standard operational metrics. I could see wanting to alarm on some of these metrics.
2. A/B testing.
3. CloudWatch could be a stop-gap tool to provide MVP measurement until a decent analytics / BI platform could be built out.

**Potential use case:**

A client of mine had an insurance quote wizard with several screens. Each screen had a "Next" button that was disable until all the required fields were populated. They had a bug that under certain circumstances, would not enable the "Next" button even though required fields were filled out. Their quoting conversion rate dropped significantly. I doubt that would have shown up in operational metrics but an alarm on rolling 5 minute conversion rate would have pinged about 6 minutes after they deployed.

Curious how y'all handle this in your worlds...

(edited to add potential use case)
